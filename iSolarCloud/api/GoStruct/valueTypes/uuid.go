package valueTypes

import (
	"encoding/json"
	"strconv"

	"github.com/MickMake/GoUnify/Only"
)

type Uuid struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	Valid  bool  `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Uuid) UnmarshalJSON(data []byte) error {
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
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t Uuid) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.int64)
		if t.Error != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatInt(t.int64, 10)
	}

	return data, t.Error
}

func (t Uuid) Value() int64 {
	return t.int64
}

func (t Uuid) String() string {
	return t.string
}

func (t Uuid) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *Uuid) SetString(value string) Uuid {
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

		var v int
		v, t.Error = strconv.Atoi(t.string)
		if t.Error != nil {
			break
		}
		t.int64 = int64(v)
		t.Valid = true
	}

	return *t
}

func (t *Uuid) SetValue(value int64) Uuid {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}

func SetUuidString(value string) Uuid {
	var t Uuid
	return t.SetString(value)
}

func SetUuidValue(value int64) Uuid {
	var t Uuid
	return t.SetValue(value)
}
