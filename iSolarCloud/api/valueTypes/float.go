package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strconv"
)


type Float struct {
	string  `json:"string,omitempty"`
	float64 `json:"float,omitempty"`
	Valid   bool `json:"valid"`
	Error   error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Float) UnmarshalJSON(data []byte) error {
	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from int
		t.Error = json.Unmarshal(data, &t.float64)
		if t.Error == nil {
			t.SetValue(t.float64)
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
func (t Float) MarshalJSON() ([]byte, error) {
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

func (t Float) Value() float64 {
	return t.float64
}

func (t Float) Match(comp float64) bool {
	if t.float64 == comp {
		return true
	}
	return false
}

func (t Float) String() string {
	return t.string
}

func (t *Float) SetString(value string) Float {
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
		if t.Error != nil {
			// @TODO - Figure out a good way to handle fields that *could be* float, but aren't always.
			t.Error = nil
			break
		}
		t.Valid = true
	}

	return *t
}

func (t *Float) SetValue(value float64) Float {
	for range Only.Once {
		t.string = ""
		t.float64 = value
		t.Valid = true
		t.string = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return *t
}

func SetFloatString(value string) Float {
	var t Float
	return t.SetString(value)
}

func SetFloatValue(value float64) Float {
	var t Float
	return t.SetValue(value)
}
