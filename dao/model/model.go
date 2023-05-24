package model

import (
	"context"
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"framework/ay"
	"math/big"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Model struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
}

type Sure int

const (
	SureNil Sure = 0
	SureYes Sure = 1
	SureNo  Sure = 2
)

type GetImage func(ids ...int) (urls []string, err error)

type FileModel interface {
	SetImageUrl(fn GetImage) error
}

func BuildOrderId() (string, error) {
	// 当前时间格式
	now := time.Now().Format("20060102150405")
	n, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return "", err
	}
	rd := fmt.Sprintf("%0.4d", n.Int64())
	// redis自增保证唯一
	key := fmt.Sprintf("%s%s%s", "transaction_id_inc", now, rd)
	id, err := ay.Redis.Incr(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	if err := ay.Redis.Expire(context.Background(), key, 2*time.Second).Err(); err != nil {
		return "", err
	}
	no := fmt.Sprintf("%s%0.4d%s", now, id, rd)
	return no, nil
}

type Decimal string

func NewFromString(s string) Decimal {
	v, _ := decimal.NewFromString(s)
	return Decimal(v.String())
}

func NewFromFloat(f float64) Decimal {
	v := decimal.NewFromFloat(f)
	return Decimal(v.String())
}

func (dec Decimal) Decimal() decimal.Decimal {
	d, _ := decimal.NewFromString(string(dec))
	return d
}

// gorm 数据类型
func (Decimal) GormDataType() string {
	return "decimal(18,3)"
}

// gorm Value
func (dec Decimal) Value() (driver.Value, error) {
	v := string(dec)
	return v, nil
}

// gorm Scan
func (dec *Decimal) Scan(value interface{}) error {
	s := fmt.Sprintf("%s", value)
	d, _ := decimal.NewFromString(s)
	*dec = Decimal(d.String())
	return nil
}

// json UnmarshalJSON
func (dec *Decimal) UnmarshalJSON(data []byte) error {
	s := string(data)
	d, _ := decimal.NewFromString(strings.Trim(s, `"`))
	*dec = Decimal(d.String())
	return nil
}

type Urls []string

func (u *Urls) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, u)
	return err
}

// Value return json value, implement driver.Valuer interface
func (u Urls) Value() (driver.Value, error) {
	if len(u) == 0 {
		return nil, nil
	}
	return json.Marshal(u)
}

// gorm 数据类型
func (Urls) GormDataType() string {
	return "json"
}
