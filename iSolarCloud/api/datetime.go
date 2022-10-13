package api

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"time"
)


var inputLayout = []string{
	DtLayout,
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
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
	string  `json:"string,omitempty"`
	time.Time `json:"time,omitempty"`
	DateType string
	Error    error
}

// UnmarshalJSON - Convert JSON to value
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &dt.string)
		if err == nil {
			dt = dt.SetString(dt.string)
			break
		}

		// Store result from time
		err = json.Unmarshal(data, &dt.Time)
		if err == nil {
			dt = dt.SetValue(dt.Time)
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (dt DateTime) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		// data, err = json.Marshal(dt.string)
		// if err != nil {
		// 	break
		// }
		data = []byte("\"" + dt.Time.Format(DtLayout) + "\"")
	}

	return data, err
}

func (dt DateTime) Value() time.Time {
	return dt.Time
}

func (dt DateTime) String() string {
	return dt.string
}

func (dt *DateTime) SetString(value string) *DateTime {
	for range Only.Once {
		dt.string = value
		dt.Time = time.Time{}

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		for _, f := range inputLayout {
			dt.Time, dt.Error = time.Parse(f, value)
			if dt.Error == nil {
				dt.string = dt.Time.Format(DtLayout)
				dt.SetDateType(value)
				break
			}
		}
	}

	return dt
}

func (dt *DateTime) SetValue(value time.Time) *DateTime {
	for range Only.Once {
		dt.string = ""
		dt.Time = value

		if value.IsZero() {
			break
		}

		dt.string = value.Format(DtLayout)
		dt.DateType = "3"
	}

	return dt
}

func (dt *DateTime) SetDateType(value string) {
	switch len(value) {
	case len(DtLayout):
		dt.DateType = "1"
	case len(DtLayoutYear):
		dt.DateType = "3"
	case len(DtLayoutMonth):
		dt.DateType = "2"
	case len(DtLayoutDay):
		dt.DateType = "1"
	case len(DtLayoutHour):
		dt.DateType = "1"
	case len(DtLayoutMinute):
		dt.DateType = "1"
	case len(DtLayoutSecond):
		dt.DateType = "1"
	}
}

func SetDateTimeString(value string) *DateTime {
	var t DateTime
	return t.SetString(value)
}

func SetDateTimeValue(value time.Time) *DateTime {
	var t DateTime
	return t.SetValue(value)
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

func (dt DateTime) PrintFull() string {
	return dt.Time.Format(DtLayout)
}

func (dt DateTime) Original() string {
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

const Now = "now"
func NewDateTime(value string) DateTime {
	var ret DateTime
	if (value == Now) || (value == "") {
		value = time.Now().Format(DtLayout)
	}
	for _, f := range inputLayout {
		ret.Time, ret.Error = time.Parse(f, value)
		if ret.Error == nil {
			ret.SetValue(ret.Time)
			ret.SetDateType(value)
			break
		}
	}
	return ret
}

func TimeNowString() string {
	return time.Now().Format(DtLayout)
}
