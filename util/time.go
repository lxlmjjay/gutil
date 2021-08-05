package util

import (
	"fmt"
	"time"
)

func TimeFormatGet() string {
	return "2006-01-02 15:04:05"
}

func TimeParseFromString(tt string) time.Time {
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", tt, time.Local)
	return end
}
func TimeParseFromStringWx(tt string) time.Time {
	end, _ := time.ParseInLocation("20060102150405", tt, time.Local)
	return end
}
func TimeFormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func TimeFormatDateOnly(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

func TimeFormatDateOrderNo(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("20060102150405" + GenRand(6))
}

func TimeFormatNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimestampFormat(t string) time.Time {
	all := "2006-01-02 15:04:05"
	//easy := "2006-01-02"
	location, err := time.ParseInLocation(all, t, time.Local)
	if err != nil {
		fmt.Println("时间转化错误", err)
	}
	return location
}

func TimeIntFormatDate(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
