package lib

import "time"

var _ Time = (*timer)(nil)

const cstLayout = "2006-01-02 15:04:05"

type Time interface {
	RFC3339ToDateTime(value string) (string, error)
	DateTimeToTime(value string) (time.Time, error)
}

type timer struct {
	Cst *time.Location
}

func NewTime() Time {
	cst, _ := time.LoadLocation("Asia/Shanghai")
	return &timer{
		Cst: cst,
	}
}

func (con *timer) RFC3339ToDateTime(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return ts.In(con.Cst).Format(cstLayout), nil
}

func (con *timer) DateTimeToTime(value string) (time.Time, error) {
	stamp, err := time.ParseInLocation(cstLayout, value, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return stamp, nil
}
