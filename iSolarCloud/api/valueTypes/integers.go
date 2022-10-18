package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strconv"
)


type Integer struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	Valid   bool `json:"valid"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Integer) UnmarshalJSON(data []byte) error {
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
func (t Integer) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		t.Valid = false

		data, err = json.Marshal(t.int64)
		if err != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatInt(t.int64, 10)
	}

	return data, err
}

func (t Integer) Value() int64 {
	return t.int64
}

func (t Integer) String() string {
	return t.string
}

func (t Integer) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *Integer) SetString(value string) Integer {
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

		var err error
		var v int
		v, err = strconv.Atoi(t.string)
		if err != nil {
			break
		}
		t.int64 = int64(v)
		t.Valid = true
	}

	return *t
}

func (t *Integer) SetValue(value int64) Integer {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}

func SetIntegerString(value string) Integer {
	var t Integer
	return t.SetString(value)
}

func SetIntegerValue(value int64) Integer {
	var t Integer
	return t.SetValue(value)
}


type Count struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	Valid   bool `json:"valid"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Count) UnmarshalJSON(data []byte) error {
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
func (t Count) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		t.Valid = false

		data, err = json.Marshal(t.int64)
		if err != nil {
			break
		}
		t.Valid = true
		// t.string = strconv.FormatInt(t.int64, 10)
	}

	return data, err
}

func (t Count) Value() int64 {
	return t.int64
}

func (t Count) String() string {
	return t.string
}

func (t Count) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *Count) SetString(value string) Count {
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

		var err error
		var v int
		v, err = strconv.Atoi(t.string)
		if err != nil {
			break
		}
		t.int64 = int64(v)
		t.Valid = true
	}

	return *t
}

func (t *Count) SetValue(value int64) Count {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}

func SetCountString(value string) Count {
	var t Count
	return t.SetString(value)
}

func SetCountValue(value int64) Count {
	var t Count
	return t.SetValue(value)
}
