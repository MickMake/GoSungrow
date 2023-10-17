package valueTypes

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MickMake/GoUnify/Only"
)

var inputDateLayout = []string{
	DateTimeFullLayout,
	DateTimeLayout,
	DateHumanLayoutLayout,
	DateLayout,
	DateHumanLayout,
	DateTimeLayoutSecond,
	DateTimeLayoutMinute,
	DateTimeLayoutHour,
	DateTimeLayoutDay,
	DateTimeLayoutMonth,
	DateTimeLayoutYear,
	DateTimeLayoutYear,
}

var DateLayoutMap = map[string]string{
	"DateTimeFullLayout":        DateTimeFullLayout,
	"DateTimeLayout":            DateTimeLayout,
	"DateTimeAltLayout":         DateTimeAltLayout,
	"DateHumanLayoutLayout":     DateHumanLayoutLayout,
	"DateTimeLayoutZeroSeconds": DateTimeLayoutZeroSeconds,
	"DateTimeLayoutSecond":      DateTimeLayoutSecond,
	"DateTimeLayoutMinute":      DateTimeLayoutMinute,
	"DateTimeLayoutHour":        DateTimeLayoutHour,
	"DateTimeLayoutDay":         DateTimeLayoutDay,
	"DateTimeLayoutMonth":       DateTimeLayoutMonth,
	"DateTimeLayoutYear":        DateTimeLayoutYear,

	"DateLayout":      DateLayout,
	"DateHumanLayout": DateHumanLayout,
	"DateLayoutDay":   DateLayoutDay,
	"DateLayoutMonth": DateLayoutMonth,
	"DateLayoutYear":  DateLayoutYear,

	"TimeLayout":            TimeLayout,
	"TimeLayoutZeroSeconds": TimeLayoutZeroSeconds,
	"TimeLayoutHourColon":   TimeLayoutHourColon,
	"TimeLayoutSecond":      TimeLayoutSecond,
	"TimeLayoutMinute":      TimeLayoutMinute,
	"TimeLayoutHour":        TimeLayoutHour,
}

const (
	DateTimeFullLayout        = time.RFC3339
	DateTimeLayout            = DateLayout + " " + TimeLayout
	DateTimeAltLayout         = DateLayoutDay + "-" + TimeLayoutSecond
	DateHumanLayoutLayout     = DateHumanLayout + " " + TimeLayout
	DateTimeLayoutZeroSeconds = DateLayout + " " + TimeLayoutZeroSeconds
	DateTimeLayoutSecond      = DateLayoutDay + TimeLayoutSecond
	DateTimeLayoutMinute      = DateLayoutDay + TimeLayoutMinute
	DateTimeLayoutHour        = DateLayoutDay + TimeLayoutHour
	DateTimeLayoutDay         = DateLayoutDay
	DateTimeLayoutMonth       = DateLayoutMonth
	DateTimeLayoutYear        = DateLayoutYear

	DateLayout      = "2006-01-02"
	DateHumanLayout = "2006/01/02"
	DateHumanMonth  = "2006/01"
	DateHumanYear   = "2006"
	DateLayoutDay   = "20060102"
	DateLayoutMonth = "200601"
	DateLayoutYear  = "2006"

	TimeLayout            = "15:04:05"
	TimeLayoutZeroSeconds = "15:04:00"
	TimeLayoutHourColon   = "15:04"
	TimeLayoutSecond      = "150405"
	TimeLayoutMinute      = "1504"
	TimeLayoutHour        = "15"

	DateTypeDay   = "1"
	DateTypeMonth = "2"
	DateTypeYear  = "3"
	DateTypeTotal = "4"
)

type DateTime struct {
	string    `json:"string,omitempty"`
	time.Time `json:"time,omitempty"`
	DateType  string
	format    string
	Error     error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from string
		dt.Error = json.Unmarshal(data, &dt.string)
		if dt.Error == nil {
			dt.SetString(dt.string)
			break
		}

		// Store result from time
		dt.Error = json.Unmarshal(data, &dt.Time)
		if dt.Error == nil {
			dt.SetValue(dt.Time)
			break
		}

		for _, f := range inputDateLayout {
			dt.Time, dt.Error = time.Parse(f, string(data))
			if dt.Error == nil {
				dt.SetDateType(string(data))
				dt.format = f
				dt.string = dt.Time.Format(DateTimeLayout)
				break
			}
		}

		if dt.Error != nil {
			fmt.Printf("Error:UnmarshalJSON DateTime(%s) - %s\n", string(data), dt.Error)
		}
	}

	return dt.Error
}

// MarshalJSON - Convert value to JSON
func (dt DateTime) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		// data = []byte("\"" + dt.Time.Format(DateTimeLayout) + "\"")
		// data = []byte("\"" + dt.string + "\"")
		// data = []byte("\"" + dt.Original() + "\"")
		data = []byte("\"" + dt.Time.Format(dt.format) + "\"")
	}

	return data, dt.Error
}

func (dt DateTime) Value() time.Time {
	return dt.Time
}

func (dt DateTime) String() string {
	// return dt.Original()
	if dt.IsZero() {
		return "--"
	}
	return dt.Time.Format(DateTimeLayout)
}

func (dt DateTime) Match(comp time.Time) bool {
	if dt.Time == comp {
		return true
	}
	return false
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

		// if value == "." {
		// 	dt.format = DateTimeLayout
		// 	dt.string = time.Now().Format(dt.format)
		// 	dt.SetDateType(dt.string)
		// 	break
		// }

		for _, f := range inputDateLayout {
			dt.Time, dt.Error = time.Parse(f, value)
			if dt.Error == nil {
				dt.SetDateType(value)
				dt.format = f
				dt.string = dt.Time.Format(f)
				break
			}
		}

		if dt.Error != nil {
			fmt.Printf("Error:SetString DateTime(%s) - %s\n", value, dt.Error)
		}
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

		dt.format = DateTimeLayout
		dt.string = value.Format(dt.format)
		dt.DateType = "3"
	}

	return *dt
}

func (dt *DateTime) SetDateType(value string) {
	for range Only.Once {
		l := len(value)
		switch {
		case l == len(DateTimeLayout):
			dt.DateType = "1"
			dt.format = DateTimeLayout
		case l == len(DateTimeLayoutYear):
			dt.DateType = "3"
			dt.format = DateTimeLayoutYear
		case l == len(DateTimeLayoutMonth):
			dt.DateType = "2"
			dt.format = DateTimeLayoutMonth
		case l == len(DateTimeLayoutDay):
			dt.DateType = "1"
			dt.format = DateTimeLayoutDay
		case l == len(DateTimeLayoutHour):
			dt.DateType = "1"
			dt.format = DateTimeLayoutHour
		case l == len(DateTimeLayoutMinute):
			dt.DateType = "1"
			dt.format = DateTimeLayoutMinute
		case l == len(DateTimeLayoutSecond):
			dt.DateType = "1"
			dt.format = DateTimeLayoutSecond
		case value == "total":
			dt.DateType = "4"
			dt.format = DateTimeLayoutYear
		}
		dt.string = dt.Time.Format(dt.format)
	}
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
	f1 := time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(), 0, 0, 0, 0, dt.Time.Location())
	ret = f1.Format(DateTimeLayoutSecond)
	return ret
	// return fmt.Sprintf("%s000000", dt.Time.Format(DtLayoutDay))
}

func (dt *DateTime) SetDayStart() {
	// dt.Time = dt.Time.Truncate(time.Hour * 24)
	dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(), 0, 0, 0, 0, dt.Time.Location())
	dt.string = dt.Time.Format(dt.format)
}

func (dt *DateTime) GetDayEndTimestamp() string {
	var ret string
	f1 := time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(), 23, 59, 59, 0, dt.Time.Location())
	ret = f1.Format(DateTimeLayoutSecond)
	return ret
	// return fmt.Sprintf("%s235900", dt.Time.Format(DtLayoutDay))
}

func (dt *DateTime) SetDayEnd() {
	// dt.Time = dt.Time.Truncate(time.Hour * 24).Add(time.Hour * 24).Add(-time.Second)
	dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(), 23, 59, 59, 0, dt.Time.Location())
	dt.string = dt.Time.Format(dt.format)
}

func (dt DateTime) PrintFull() string {
	return dt.Time.Format(DateTimeLayout)
}

func (dt DateTime) Original() string {
	var ret string
	switch dt.DateType {
	case "3":
		ret = dt.Time.Format(DateTimeLayoutYear)
	case "2":
		ret = dt.Time.Format(DateTimeLayoutMonth)
	case "1":
		ret = dt.Time.Format(DateTimeLayoutDay)
	default:
		ret = dt.Time.Format(DateTimeLayout)
	}
	return ret
}

func (dt *DateTime) GetRanges(count int, dur time.Duration, format string) []string {
	var ret []string
	for range Only.Once {
		if format == "" {
			format = DateTimeLayout
		}
		next := dt.Time
		for i := 0; i < count; i++ {
			ret = append(ret, next.Format(format))
			next = next.Add(dur)
		}
	}
	return ret
}

const Now = "now"

func NewDateTime(value string) DateTime {
	var ret DateTime
	for range Only.Once {
		if (value == Now) || (value == "") || (value == ".") {
			// value = time.Now().Format(DateTimeLayout)
			ret.SetValue(time.Now())
			ret.SetDateType(ret.string)
			break
		}
		for _, f := range inputDateLayout {
			ret.Time, ret.Error = time.Parse(f, value)
			if ret.Error == nil {
				ret.SetValue(ret.Time)
				ret.SetDateType(value)
				ret.format = f
				break
			}
		}

		if ret.Error != nil {
			fmt.Printf("Error:NewDateTime DateTime(%s) - %s\n", value, ret.Error)
		}
	}
	return ret
}

func ParseDateTime(value string) (time.Time, string, error) {
	var ret time.Time
	var format string
	var err error

	for range Only.Once {
		if (value == Now) || (value == "") || (value == ".") {
			ret = time.Now()
			break
		}

		for _, format = range inputDateLayout {
			ret, err = time.Parse(format, value)
			if err == nil {
				break
			}
		}
	}

	return ret, format, err
}

func TimeNowString() string {
	return time.Now().Format(DateTimeLayout)
}
