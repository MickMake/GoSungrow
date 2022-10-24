package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


type Bool struct {
	string `json:"string,omitempty"`
	bool   `json:"bool,omitempty"`
	Valid   bool `json:"valid"`
	Error   error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Bool) UnmarshalJSON(data []byte) error {

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		var ret2 int64
		// Store result from int
		t.Error = json.Unmarshal(data, &ret2)
		if t.Error == nil {
			t.SetInteger(ret2)
			break
		}

		var ret1 bool
		// Store result from int
		t.Error = json.Unmarshal(data, &ret1)
		if t.Error == nil {
			t.SetValue(ret1)
			break
		}

		var ret3 string
		// Store result from string
		t.Error = json.Unmarshal(data, &ret3)
		if t.Error == nil {
			t.SetString(ret3)
			break
		}

		t.SetString(string(data))
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t Bool) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.bool)
		if t.Error != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatFloat(t.bool, 'f', -1, 64)
	}

	return data, t.Error
}

func (t Bool) Value() bool {
	return t.bool
}

func (t Bool) Match(comp bool) bool {
	if t.bool == comp {
		return true
	}
	return false
}

func (t Bool) String() string {
	return t.string
}

func (t *Bool) SetString(value string) Bool {
	for range Only.Once {
		t.string = value
		t.Valid = false

		switch strings.ToLower(t.string) {
			case "false":
				fallthrough
			case "no":
				fallthrough
			case "off":
				fallthrough
			case "0":
				fallthrough
			case "":
				// 	fallthrough
				// case "--":
				t.bool = false
				t.string = "false"
				t.Valid = true

			case "true":
				fallthrough
			case "yes":
				fallthrough
			case "on":
				fallthrough
			case "1":
				t.bool = true
				t.string = "true"
				t.Valid = true
		}
	}

	return *t
}

func (t *Bool) SetValue(value bool) Bool {
	for range Only.Once {
		t.bool = value
		t.Valid = true

		if t.bool {
			t.string = "true"
			break
		}

		t.string = "false"
	}

	return *t
}

func (t *Bool) SetInteger(value int64) Bool {
	for range Only.Once {
		t.Valid = true

		if value == 0 {
			t.bool = false
			t.string = "false"
			break
		}

		t.bool = true
		t.string = "true"
	}

	return *t
}

func SetBoolString(value string) Bool {
	var t Bool
	return t.SetString(value)
}

func SetBoolValue(value bool) Bool {
	var t Bool
	return t.SetValue(value)
}
