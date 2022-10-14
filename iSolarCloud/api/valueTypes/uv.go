package valueTypes

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"sort"
	"strconv"
	"strings"
)


type UnitValue struct {
	UnitValue   string `json:"unit"`  			// Primary iSolarCloud entity.
	StringValue string `json:"value"` 			// Primary iSolarCloud entity.

	TypeValue   string `json:"type_value"`

	float64     `json:"value_float,omitempty"`
	int64       `json:"value_int,omitempty"`
	bool        `json:"value_bool,omitempty"`

	isFloat     bool
	Valid       bool `json:"valid"`
}

func (t *UnitValue) UnitValueFix() UnitValue {
	switch t.UnitValue {
	case "w":
		t.UnitValue = "W"
	}

	switch t.UnitValue {
		case "Wp":
			fallthrough
		case "Wh":
			fallthrough
		case "W":
			fv := t.Value() / 1000
			t.SetFloat(fv)
			t.SetUnit("k" + t.UnitValue)

			// fv, err := strconv.ParseFloat(value, 64)
			// if err == nil {
			// 	fv = fv / 1000
			// 	value, _ = DivideByThousand(value)
			// 	UnitValue = "k" + UnitValue
			// }
	}

	return *t
}

// UnmarshalJSON - Convert JSON to value
func (t *UnitValue) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		t.Valid = false

		if len(data) == 0 {
			break
		}

		var resString struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		}
		// Store result from JSON string
		err = json.Unmarshal(data, &resString)
		if err == nil {
			t.SetString(resString.Value)
			t.SetUnit(resString.Unit)
			t.UnitValueFix()
			break
		}

		var resInteger struct {
			Unit  string `json:"unit"`
			Value int64  `json:"value"`
		}
		// Store result from JSON string
		err = json.Unmarshal(data, &resInteger)
		if err == nil {
			t.SetInteger(resInteger.Value)
			t.SetUnit(resInteger.Unit)
			t.UnitValueFix()
			break
		}

		var resFloat struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		}
		// Store result from JSON string
		err = json.Unmarshal(data, &resFloat)
		if err == nil {
			t.SetFloat(resFloat.Value)
			t.SetUnit(resFloat.Unit)
			t.UnitValueFix()
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t UnitValue) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		t.Valid = false

		if t.isFloat {
			// Store result to JSON string
			data, err = json.Marshal(&struct {
				Unit  string  `json:"unit"`
				Value float64 `json:"value"`
			}{
				Unit:  t.UnitValue,
				Value: t.float64,
			})
			if err != nil {
				break
			}

			t.Valid = true
			break
		}

		// Store result to JSON string
		data, err = json.Marshal(&struct {
			Unit  string `json:"unit"`
			Value int64  `json:"value"`
		}{
			Unit:  t.UnitValue,
			Value: t.int64,
		})
		if err != nil {
			break
		}

		t.Valid = true
	}

	return data, err
}

func (t UnitValue) Value() float64 {
	if t.isFloat {
		return t.float64
	}
	return float64(t.int64)
}

func (t UnitValue) ValueFloat() float64 {
	return t.float64
}

func (t UnitValue) ValueInt() int64 {
	return t.int64
}

func (t UnitValue) ValueBool() bool {
	return t.bool
}

func (t UnitValue) String() string {
	return t.StringValue
}

func (t UnitValue) Unit() string {
	return t.UnitValue
}

func (t UnitValue) Type() string {
	return t.TypeValue
}

func (t *UnitValue) SetString(value string) UnitValue {
	for range Only.Once {
		t.StringValue = value
		t.int64 = 0
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}

		if strings.Contains(value, ".") {
			v, err := strconv.ParseFloat(t.StringValue, 64)
			if err != nil {
				break
			}
			t.SetFloat(v)
			break
		}

		v, err := strconv.Atoi(t.StringValue)
		if err != nil {
			break
		}
		t.SetInteger(int64(v))
	}

	return *t
}

func (t *UnitValue) SetInteger(value int64) UnitValue {
	for range Only.Once {
		t.int64 = value
		t.float64 = float64(value)
		t.isFloat = false
		t.Valid = true
		t.StringValue = strconv.FormatInt(t.int64, 10)
	}

	return *t
}

func (t *UnitValue) SetFloat(value float64) UnitValue {
	for range Only.Once {
		t.int64 = int64(value)
		t.float64 = value
		t.isFloat = true
		t.Valid = true
		// t.String = strconv.FormatFloat(t.float64, 'f', 12, 64)
		// t.String = strings.TrimRight(t.String, "0")
		t.StringValue = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return *t
}

func (t *UnitValue) SetBool(value bool) UnitValue {
	for range Only.Once {
		t.isFloat = false
		t.Valid = true
		if value {
			t.bool = value
			t.float64 = 0
			t.int64 = 0
			t.StringValue = "true"
			break
		}
		t.bool = value
		t.float64 = 1
		t.int64 = 1
		t.StringValue = "false"
	}

	return *t
}

func (t *UnitValue) SetUnit(unit string) UnitValue {
	for range Only.Once {
		t.UnitValue = unit
		// t.Valid = true
	}

	return *t
}

func (t *UnitValue) SetType(Type string) UnitValue {
	for range Only.Once {
		t.TypeValue = Type
	}

	return *t
}

func SetUnitValueString(value string, unit string, Type string) UnitValue {
	var t UnitValue
	t = t.SetString(value)
	t = t.SetUnit(unit)
	t = t.SetType(Type)
	return t.UnitValueFix()
}

func SetUnitValueInteger(value int64, unit string, Type string) UnitValue {
	var t UnitValue
	t = t.SetInteger(value)
	t = t.SetUnit(unit)
	t = t.SetType(Type)
	return t.UnitValueFix()
}

func SetUnitValueFloat(value float64, unit string, Type string) UnitValue {
	var t UnitValue
	t = t.SetFloat(value)
	t = t.SetUnit(unit)
	t = t.SetType(Type)
	return t.UnitValueFix()
}

func SetUnitValueBool(value bool) UnitValue {
	var t UnitValue
	t = t.SetBool(value)
	t = t.SetUnit("")
	t = t.SetType("Bool")
	return t
}

type UnitValues []UnitValue
type UnitValueMap map[PointId]UnitValue

func (u *UnitValueMap) Sort() []string {
	var ret []string
	for n := range *u {
		ret = append(ret, n.String())
	}
	sort.Strings(ret)
	return ret
}


func (t UnitValues) Unit() string {
	var ret string
	for _, v := range t {
		ret = v.Unit()
		break
	}
	return ret
}

func (t UnitValues) Type() string {
	var ret string
	for _, v := range t {
		ret = v.Type()
		break
	}
	return ret
}
