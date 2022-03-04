package api

import (
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

func TimeNowString() string {
	return time.Now().Format(DtLayout)
}

func (dt *DateTime) GetDayStartTimestamp() string {
	var ret string
	f1 := dt.Time.Round(time.Hour * 24)
	ret = f1.Format(DtLayoutSecond)
	return ret
	// return fmt.Sprintf("%s000000", dt.Time.Format(DtLayoutDay))
}

func (dt *DateTime) GetDayEndTimestamp() string {
	var ret string
	f1 := dt.Time.Round(time.Hour * 24).Add(time.Hour * 24)
	ret = f1.Format(DtLayoutSecond)
	return ret
	// return fmt.Sprintf("%s235900", dt.Time.Format(DtLayoutDay))
}

func (dt DateTime) String() string {
	// return dt.Time.Format(DtLayout)
	var ret string
	switch dt.DateType {
		case "3":
			ret = dt.Time.Format(DtLayoutYear)
		case "2":
			ret = dt.Time.Format(DtLayoutMonth)
		case "1":
			ret = dt.Time.Format(DtLayoutDay)
	}
	return ret
}

func (dt DateTime) PrintFull() string {
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
		case len(DtLayoutHour):
			ret = "1"
		case len(DtLayoutMinute):
			ret = "1"
		case len(DtLayoutSecond):
			ret = "1"
	}
	return ret
}

// func (dt DateTime) GetFilenameSuffix() string {
// 	var ret string
// 	switch dt.DateType {
// 	case "3":
// 		ret = dt.Time.Format(DtLayoutYear)
// 	case "2":
// 		ret = dt.Time.Format(DtLayoutMonth)
// 	case "1":
// 		ret = dt.Time.Format(DtLayoutDay)
// 	}
// 	return ret
// }
