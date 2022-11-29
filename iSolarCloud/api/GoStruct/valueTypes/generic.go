package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strconv"
)


type Generic struct {
	float64     `json:"value_float,omitempty"`
	string      `json:"value_string,omitempty"`

	Valid   bool `json:"valid"`
	Error   error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Generic) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from float
		t.Error = json.Unmarshal(data, &t.float64)
		if t.Error == nil {
			t.SetFloat(t.float64)
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error == nil {
			t.SetString(t.string)
			break
		}

		t.Error = nil
		t.SetString(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t Generic) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.float64)
		if t.Error != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return data, t.Error
}

func (t Generic) Value() float64 {
	return t.float64
}

func (t Generic) Match(comp float64) bool {
	if t.float64 == comp {
		return true
	}
	return false
}

func (t Generic) String() string {
	return t.string
}

func (t *Generic) SetString(value string) Generic {
	for range Only.Once {
		t.string = value
		t.float64 = 0
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}

		t.float64, t.Error = strconv.ParseFloat(t.string, 64)
		if t.Error == nil {
			t.Valid = true
			break
		}

		var v int64
		v, t.Error = strconv.ParseInt(t.string, 10, 64)
		if t.Error == nil {
			t.float64 = float64(v)
			t.Valid = true
			break
		}
	}

	return *t
}

func (t *Generic) SetFloat(value float64) Generic {
	for range Only.Once {
		t.float64 = value
		t.Valid = true
		t.string = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return *t
}

func (t *Generic) SetInteger(value int64) Generic {
	for range Only.Once {
		t.float64 = float64(value)
		t.Valid = true
		t.string = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return *t
}

func (t *Generic) ToUnitValue() UnitValue {
	return SetUnitValueFloat("", "", t.float64)
}

func SetGenericString(value string) Generic {
	var t Generic
	return t.SetString(value)
}

func SetGenericFloat(value float64) Generic {
	var t Generic
	return t.SetFloat(value)
}

func SetGenericInteger(value int64) Generic {
	var t Generic
	return t.SetInteger(value)
}
