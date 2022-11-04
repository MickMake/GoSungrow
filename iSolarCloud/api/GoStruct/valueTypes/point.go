package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"regexp"
	"strconv"
	"strings"
)


// @TODO - Consider standardizing points to a known format.
// @TODO - queryUserCurveTemplateData is a good example of the structure.
// type Point struct {
//	PointId    valueTypes.PointId `json:"point_id"`
//	PointName  valueTypes.String  `json:"point_name"`
//	PsId       valueTypes.PsId    `json:"ps_id"`
//	PsKey      valueTypes.PsKey   `json:"ps_key"`
//	Color      valueTypes.String  `json:"color"`
//	DetailId   valueTypes.String  `json:"detail_id"`
//	PsName     valueTypes.String  `json:"ps_name"`
//	Statistics valueTypes.String  `json:"statistics"`
//	Style      valueTypes.String  `json:"style"`
//	Unit       valueTypes.String  `json:"unit"`
//	DataList   []struct {
//		TimeStamp valueTypes.String `json:"time_stamp"`
//		Value     valueTypes.String `json:"value"`
//	} `json:"data_list" PointArrayFlatten:"false"`
// }


type PointId struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	isInt  bool
	Valid  bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointId) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from int
		t.Error = json.Unmarshal(data, &t.int64)
		if t.Error == nil {
			t.SetValue(t.int64)
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error == nil {
			t.SetString(t.string)
			break
		}

		t.SetString(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PointId) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		if t.isInt {
			data, t.Error = json.Marshal(t.int64)
			break
		}

		data, t.Error = json.Marshal(t.string)
	}

	return data, t.Error
}

func (t PointId) Value() int64 {
	return t.int64
}

func (t PointId) String() string {
	if t.isInt {
		return "p" + t.string
	}
	return t.string
}

func (t PointId) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *PointId) SetString(value string) PointId {
	for range Only.Once {
		t.string = value
		t.int64 = 0
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}
		t.Valid = true

		var v int
		v, t.Error = strconv.Atoi(t.string)
		if t.Error != nil {
			break
		}
		t.int64 = int64(v)
		t.isInt = true
	}

	return *t
}

func (t *PointId) SetValue(value int64) PointId {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.isInt = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}


func (t *PointId) Fix() PointId {
	for range Only.Once {
		p := strings.TrimPrefix(t.string, "p")
		_, t.Error = strconv.ParseInt(p, 10, 64)
		if t.Error != nil {
			t.Valid = false
			break
		}
	}
	return *t
}

func (t *PointId) PointToName() string {
	return PointToName(t.String())
}

func CleanString(s string) string {
	// var ret string
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			(b == '-') ||
			(b == '_') ||
			(b == '.') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	// dupes := regexp.MustCompile(`\s+`)
	// ret = dupes.ReplaceAllString(result.String(), )
	return result.String()
}

func PointToName(ret string) string {
	ret = CleanString(ret)
	ret = strings.ReplaceAll(ret, "_", " ")
	ret = strings.Title(ret)

	// Add space between lowercase and uppercase letters.
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	ret = re.ReplaceAllString(ret, "$1 $2")

	// Lowercase point ids.
	re = regexp.MustCompile("P([0-9]+)")
	ret = re.ReplaceAllString(ret, "p$1")

	return ret
}


func SetPointIdString(value string) PointId {
	var t PointId
	return t.SetString(value)
}

func SetPointIdValue(value int64) PointId {
	var t PointId
	return t.SetValue(value)
}
