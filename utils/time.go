package utils

import "time"

func GetTimeNowFormat() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func GetTimeFormat(time2 time.Time) string {
	return time2.Format("2006-01-02 15:04:05")
}

// 2006-01-02
func GetDateFormat(time2 time.Time) string {
	return time2.Format("2006-01-02")
}

func GetMonthFormat(time2 time.Time) string {
	return time2.Format("2006-01")
}

func ParseDate(time2 string) time.Time {
	t, e := time.ParseInLocation("2006-01-02", time2, time.Local)
	if e != nil {
		return time.Now()
	}
	return t
}

func ParseDateTime(time2 string) time.Time {
	t, e := time.ParseInLocation("2006-01-02 15:04:05", time2, time.Local)
	if e != nil {
		return time.Now()
	}
	return t
}

// 计算月相差
// t1-t2
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval = monthInterval % 12
	month = yearInterval*12 + monthInterval
	return
}
