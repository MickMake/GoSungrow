package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
)


type String struct {
	string `json:"string,omitempty"`
	Valid   bool `json:"valid"`
}

// UnmarshalJSON - Convert JSON to value
func (t *String) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &t.string)
		if err != nil {
			break
		}
		t.SetString(t.string)
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t String) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		t.Valid = false

		data, err = json.Marshal(t.string)
		if err != nil {
			break
		}
		t.Valid = true
	}

	return data, err
}

func (t String) Value() string {
	return t.string
}

func (t String) String() string {
	return t.string
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
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKey) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &t.string)
		if err != nil {
			break
		}
		t.SetPsKey(t.string)
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t PsKey) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		t.Valid = false

		data, err = json.Marshal(t.string)
		if err != nil {
			break
		}
		t.Valid = true
	}

	return data, err
}

func (t PsKey) Value() string {
	return t.string
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
