package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"regexp"
	"strconv"
	"strings"
)


type PointId struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	isInt  bool
	Valid  bool `json:"valid"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PointId) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from int
		err = json.Unmarshal(data, &t.int64)
		if err == nil {
			t.SetValue(t.int64)
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &t.string)
		if err == nil {
			t.SetString(t.string)
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t PointId) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		if t.isInt {
			data, err = json.Marshal(t.int64)
			break
		}

		data, err = json.Marshal(t.string)
	}

	return data, err
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

		var err error
		var v int
		v, err = strconv.Atoi(t.string)
		if err != nil {
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
		_, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
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
