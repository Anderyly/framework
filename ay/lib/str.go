/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"framework/ay"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var _ Str = (*str)(nil)

type Str interface {
	LastTime(t int64) string
	Md5(str string) string
	AuthCode(str string, operation bool, key string, expiry int64) (string, error) // false is decode
	MakeCoupon(coupon string) (float64, error)                                     //
	Summary(content string, count int) string
}

type str struct {
}

func NewStr() Str {
	return &str{}
}

func (con *str) LastTime(t int64) (msg string) {
	minute := int(time.Minute)
	s := int(time.Now().Unix()-t) / minute

	switch {
	case s < minute:
		msg = strconv.Itoa(s) + "分钟前"
	case s >= minute && s < (minute*24):
		msg = strconv.Itoa(s/minute) + "小时前"
	case s >= (minute*24) && s < (minute*24*3):
		msg = strconv.Itoa(s/24/minute) + "天前"
	default:
		msg = time.Unix(t, 0).Format("2006-01-02 15:04:05")
	}
	return
}

func (con *str) Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (con *str) AuthCode(str string, operation bool, key string, expiry int64) (result string, err error) {
	cKeyLength := 1
	if len(str) < cKeyLength {
		err = errors.New("str length less than key length")
		return
	}

	if len(str) < 11 && !operation {
		err = errors.New("decode length less than 11")
		return
	}
	if key == "" {
		key = ay.Yaml.GetString("key")
	}
	key = con.Md5(key)

	keyA := con.Md5(key[:16])
	keyB := con.Md5(key[16:])
	keyC := ""
	if !operation {
		keyC = str[:cKeyLength]
	} else {
		sTime := con.Md5(time.Now().String())
		sLen := 32 - cKeyLength
		keyC = sTime[sLen:]
	}
	cryptKey := fmt.Sprintf("%s%s", keyA, con.Md5(keyA+keyC))
	keyLength := len(cryptKey)
	if !operation {
		str = strings.Replace(str, "-", "+", -1)
		str = strings.Replace(str, "_", "/", -1)
		str = strings.Replace(str, "=", ".", -1)
		var strByte []byte
		strByte, err = base64.StdEncoding.DecodeString(str[cKeyLength:])
		if err != nil {
			return
		}
		str = string(strByte)
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		tmpMd5 := con.Md5(str + keyB)
		str = fmt.Sprintf("%010d%s%s", expiry, tmpMd5[:16], str)
	}
	stringLength := len(str)
	resData := make([]byte, 0, stringLength)
	var randKey, box [256]int
	j := 0
	a := 0
	i := 0
	tmp := 0
	for i = 0; i < 256; i++ {
		randKey[i] = int(cryptKey[i%keyLength])
		box[i] = i
	}
	for i = 0; i < 256; i++ {
		j = (j + box[i] + randKey[i]) % 256
		tmp = box[i]
		box[i] = box[j]
		box[j] = tmp
	}
	a = 0
	j = 0
	for i = 0; i < stringLength; i++ {
		a = (a + 1) % 256
		j = (j + box[a]) % 256
		tmp = box[a]
		box[a] = box[j]
		box[j] = tmp
		resData = append(resData, byte(int(str[i])^box[(box[a]+box[j])%256]))
	}
	result = string(resData)
	if !operation {
		var frontTen int64
		frontTen, err = strconv.ParseInt(result[:10], 10, 0)
		if err != nil {
			return
		}
		if (frontTen == 0 || frontTen-time.Now().Unix() > 0) && result[10:26] == con.Md5(result[26:] + keyB)[:16] {
			result = result[26:]
			return
		} else {
			err = errors.New("decode error")
			return
		}
	} else {
		result = keyC + base64.StdEncoding.EncodeToString([]byte(result))
		result = strings.Replace(result, "+", "-", -1)
		result = strings.Replace(result, "/", "_", -1)
		result = strings.Replace(result, ".", "=", -1)
		return
	}
}

// MakeCoupon 优惠价
func (con *str) MakeCoupon(coupon string) (amount float64, err error) {
	var maxPrice float64
	var minPrice float64
	couponArr := strings.Split(coupon, "-")

	if maxPrice, err = strconv.ParseFloat(couponArr[1], 64); err != nil {
		return
	}
	if minPrice, err = strconv.ParseFloat(couponArr[0], 64); err != nil {
		return
	}
	cha := maxPrice - minPrice
	var price float64
	for {
		p := 0.01 + rand.Float64()*(cha-0.01)
		if price, err = strconv.ParseFloat(fmt.Sprintf("%.2f", p), 64); err != nil {
			return
		}
		if price <= cha {
			break
		}
	}
	amount = price + minPrice
	return
}

// Summary 过滤特殊字符
func (con *str) Summary(content string, count int) string {

	content = strings.Replace(content, `@<script(.*?)</script>@is`, "", -1)
	content = strings.Replace(content, `@<iframe(.*?)</iframe>@is`, "", -1)
	content = strings.Replace(content, `@<style(.*?)</style>@is`, "", -1)
	content = strings.Replace(content, `\`, "", -1)

	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	content = re.ReplaceAllString(content, "")

	content = strings.Replace(content, " ", "", -1)
	content = strings.Replace(content, "　　", "", -1)
	content = strings.Replace(content, "\t", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	cont := []rune(content)

	return string(cont[:count])
}
