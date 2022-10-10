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
			dt.SetString(dt.string)
			break
		}

		// Store result from time
		err = json.Unmarshal(data, &dt.Time)
		if err == nil {
			dt.SetValue(dt.Time)
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

func (dt *DateTime) SetString(value string) DateTime {
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

		dt.DateType = dt.getDateType()
		for _, f := range inputLayout {
			dt.Time, dt.Error = time.Parse(f, value)
			if dt.Error == nil {
				break
			}
		}

		// switch dt.DateType {
		// 	case "3":
		// 		ret = dt.Time.Format(DtLayoutYear)
		// 	case "2":
		// 		ret = dt.Time.Format(DtLayoutMonth)
		// 	case "1":
		// 		ret = dt.Time.Format(DtLayoutDay)
		// }
	}

	return *dt
}

func (dt *DateTime) SetValue(value time.Time) DateTime {
	for range Only.Once {
		dt.string = ""
		dt.Time = value

		if value.IsZero() {
			break
		}

		dt.string = value.Format(DtLayout)
		dt.DateType = "3"
	}

	return *dt
}

func SetDateTimeString(value string) DateTime {
	var t DateTime
	return t.SetString(value)
}

func SetDateTimeValue(value time.Time) DateTime {
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

// func (dt DateTime) String() string {
// 	// return dt.Time.Format(DtLayout)
// 	var ret string
// 	switch dt.DateType {
// 		case "3":
// 			ret = dt.Time.Format(DtLayoutYear)
// 		case "2":
// 			ret = dt.Time.Format(DtLayoutMonth)
// 		case "1":
// 			ret = dt.Time.Format(DtLayoutDay)
// 	}
// 	return ret
// }

func (dt DateTime) PrintFull() string {
	return dt.Time.Format(DtLayout)
}

func (dt DateTime) getDateType() string {
	var ret string
	switch len(dt.string) {
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

func NewDateTime(date string) DateTime {
	var ret DateTime
	if date == "" {
		date = time.Now().Format(DtLayoutDay)
	}
	ret.string = date
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
