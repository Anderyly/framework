package lib

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"framework/ay"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var _ Str = (*str)(nil)

type Str interface {
	LastTime(t int64) string
	MD5(str string) string
	AuthCode(str, operation, key string, expiry int64) string
	MakeCoupon(coupon string) float64
	Summary(content string, count int) string
}

type str struct {
}

func NewStr() Str {
	return &str{}
}

func (con *str) LastTime(t int64) (msg string) {
	s := int(time.Now().Unix()-t) / 60

	switch {
	case s < 60:
		msg = strconv.Itoa(s) + "分钟前"

	case s >= 60 && s < (60*24):
		msg = strconv.Itoa(s/60) + "小时前"
	case s >= (60*24) && s < (60*24*3):
		msg = strconv.Itoa(s/24/60) + "天前"
	default:
		msg = time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
	}
	return
}

func (con *str) MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (con *str) AuthCode(str, operation, key string, expiry int64) string {
	cKeyLength := 1
	if len(str) < cKeyLength {
		return ""
	}

	if len(str) < 11 && operation == "DECODE" {
		return ""
	}
	if key == "" {
		key = ay.Yaml.GetString("key")
	}
	key = con.MD5(key)

	keyA := con.MD5(key[:16])
	keyB := con.MD5(key[16:])
	keyC := ""
	if operation == "DECODE" {
		keyC = str[:cKeyLength]
	} else {
		sTime := con.MD5(time.Now().String())
		sLen := 32 - cKeyLength
		keyC = sTime[sLen:]
	}
	cryptKey := fmt.Sprintf("%s%s", keyA, con.MD5(keyA+keyC))
	keyLength := len(cryptKey)
	if operation == "DECODE" {
		str = strings.Replace(str, "-", "+", -1)
		str = strings.Replace(str, "_", "/", -1)
		strByte, err := base64.StdEncoding.DecodeString(str[cKeyLength:])
		if err != nil {
			return ""
		}
		str = string(strByte)
	} else {
		if expiry != 0 {
			expiry = expiry + time.Now().Unix()
		}
		tmpMd5 := con.MD5(str + keyB)
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
	tmp = 0
	for i = 0; i < stringLength; i++ {
		a = (a + 1) % 256
		j = (j + box[a]) % 256
		tmp = box[a]
		box[a] = box[j]
		box[j] = tmp
		resData = append(resData, byte(int(str[i])^box[(box[a]+box[j])%256]))
	}
	result := string(resData)
	if operation == "DECODE" {
		frontTen, _err := strconv.ParseInt(result[:10], 10, 0)
		if _err != nil {
			return ""
		}
		if (frontTen == 0 || frontTen-time.Now().Unix() > 0) && result[10:26] == con.MD5(result[26:] + keyB)[:16] {
			return result[26:]
		} else {
			return ""
		}
	} else {
		result = keyC + base64.StdEncoding.EncodeToString([]byte(result))
		result = strings.Replace(result, "+", "-", -1)
		result = strings.Replace(result, "/", "_", -1)
		return result
	}
}

// MakeCoupon 优惠价
func (con *str) MakeCoupon(coupon string) float64 {

	couponArr := strings.Split(coupon, "-")

	log.Println(couponArr)

	maxPrice, err := strconv.ParseFloat(couponArr[1], 64)
	if err != nil {
		return 0
	}
	minPrice, err1 := strconv.ParseFloat(couponArr[0], 64)
	if err1 != nil {
		return 0
	}
	cha := maxPrice - minPrice
	var price float64
	for {
		p := 0.01 + rand.Float64()*(cha-0.01)
		price, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", p), 64)
		if price <= cha {
			break
		}
	}
	return price + minPrice
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
