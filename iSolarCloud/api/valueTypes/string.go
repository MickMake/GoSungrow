package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
)


type String struct {
	string `json:"string,omitempty"`
	Valid   bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *String) UnmarshalJSON(data []byte) error {

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error != nil {
			break
		}
		t.SetString(t.string)
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t String) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.string)
		if t.Error != nil {
			break
		}
		t.Valid = true
	}

	return data, t.Error
}

func (t String) Value() string {
	return t.string
}

func (t String) String() string {
	return t.string
}

func (t String) Match(comp string) bool {
	if t.string == comp {
		return true
	}
	return false
}

func (t *String) SetString(value string) String {
	for range Only.Once {
		t.string = value
		t.Valid = true
	}

	return *t
}

func (t *String) SetValue(value string) String {
	for range Only.Once {
		t.string = value
		t.Valid = true
	}

	return *t
}

func SetStringValue(value string) String {
	var t String
	return t.SetValue(value)
}


type PsKey struct {
	string `json:"string,omitempty"`
	Valid   bool `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKey) UnmarshalJSON(data []byte) error {

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from string
		t.Error = json.Unmarshal(data, &t.string)
		if t.Error != nil {
			break
		}
		t.SetPsKey(t.string)
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t PsKey) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		data, t.Error = json.Marshal(t.string)
		if t.Error != nil {
			break
		}
		t.Valid = true
	}

	return data, t.Error
}

func (t PsKey) Value() string {
	return t.string
}

func (t PsKey) String() string {
	return t.string
}

func (t PsKey) Match(comp string) bool {
	if t.string == comp {
		return true
	}
	return false
}

func (t PsKey) PsKey() string {
	return t.string
}

func (t *PsKey) SetPsKey(value string) PsKey {
	for range Only.Once {
		t.string = value
		t.Valid = true
	}

	return *t
}

func (t *PsKey) SetValue(value string) PsKey {
	for range Only.Once {
		t.string = value
		t.Valid = true
	}

	return *t
}

func SetPsKeyValue(value string) PsKey {
	var t PsKey
	return t.SetValue(value)
}
