package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strconv"
)


type PsId struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
	Valid  bool  `json:"valid"`
	Error  error `json:"-"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsId) UnmarshalJSON(data []byte) error {
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
func (t PsId) MarshalJSON() ([]byte, error) {
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

func (t PsId) Value() int64 {
	return t.int64
}

func (t PsId) String() string {
	return t.string
}

func (t PsId) Match(comp int64) bool {
	if t.int64 == comp {
		return true
	}
	return false
}

func (t *PsId) SetString(value string) PsId {
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

func (t *PsId) SetValue(value int64) PsId {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.Valid = true
		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}


func SetPsIdString(value string) PsId {
	var t PsId
	return t.SetString(value)
}

func SetPsIdValue(value int64) PsId {
	var t PsId
	return t.SetValue(value)
}

func SetPsIdStrings(values []string) PsIds {
	var t PsIds
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range values {
			if psId == "" {
				continue
			}
			t = append(t, SetPsIdString(psId))
		}
	}
	return t
}

func SetPsIdValues(values []int64) PsIds {
	var t PsIds
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range values {
			if psId == 0 {
				continue
			}
			t = append(t, SetPsIdValue(psId))
		}
	}
	return t
}


type PsIds []PsId

func (t *PsIds) Strings() []string {
	var ret []string
	for range Only.Once {
		// sgd.PsId = valueTypes.SetPsIdString(psId)
		for _, psId := range *t {
			ret = append(ret, psId.String())
		}
	}
	return ret
}
