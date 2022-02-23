package api

import (
	"fmt"
	"time"
)

var inputLayout = []string{
	DtLayout,
	"2006/01/02 15:04:05",
	DtLayoutSecond,
	DtLayoutMinute,
	DtLayoutHour,
	DtLayoutDay,
	DtLayoutMonth,
	DtLayoutYear,
}

const (
	DtLayout            = "2006-01-02 15:04:05"
	DtLayoutZeroSeconds = "2006-01-02 15:04:00"
	DtLayoutSecond      = "20060102150405"
	DtLayoutMinute      = "200601021504"
	DtLayoutHour        = "2006010215"
	DtLayoutDay         = "20060102"
	DtLayoutMonth       = "200601"
	DtLayoutYear        = "2006"
)

type DateTime struct {
	date string
	time.Time
	DateType string
	Error    error
}

func NewDateTime(date string) DateTime {
	var ret DateTime
	if date == "" {
		date = time.Now().Format(DtLayoutDay)
	}
	ret.date = date
	ret.DateType = ret.getDateType()
	for _, f := range inputLayout {
		ret.Time, ret.Error = time.Parse(f, date)
		if ret.Error == nil {
			break
		}
	}
	return ret
}

func (dt *DateTime) GetDayStartTimestamp() string {
	return fmt.Sprintf("%s000000", dt.Time.Format(DtLayoutDay))
}

func (dt *DateTime) GetDayEndTimestamp() string {
	return fmt.Sprintf("%s235900", dt.Time.Format(DtLayoutDay))
}

func (dt DateTime) String() string {
	return dt.Time.Format(DtLayout)
}

func (dt DateTime) getDateType() string {
	var ret string
	switch len(dt.date) {
	case len(DtLayout):
		ret = "1"
	case len(DtLayoutYear):
		ret = "3"
	case len(DtLayoutMonth):
		ret = "2"
	case len(DtLayoutDay):
		ret = "1"
	}
	return ret
}
