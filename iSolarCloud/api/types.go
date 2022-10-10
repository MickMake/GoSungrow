package api

import (
	"GoSungrow/Only"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)


type UnitValue struct {
	unit  string `json:"unit"`
	string `json:"value"`
	float64 `json:"value_float,omitempty"`
	int64 `json:"value_int,omitempty"`
	isFloat bool
}

func (t *UnitValue) UnitValueFix() UnitValue {
	switch t.unit {
		case "w":
			t.unit = "W"
	}

	switch t.unit {
		case "Wp":
			fallthrough
		case "Wh":
			fallthrough
		case "W":
			fv := t.Value() / 1000
			t.SetFloat(fv)
			t.SetUnit("k" + t.unit)

			// fv, err := strconv.ParseFloat(value, 64)
			// if err == nil {
			// 	fv = fv / 1000
			// 	value, _ = DivideByThousand(value)
			// 	unit = "k" + unit
			// }
	}

	return *t
}

func (t *UnitValue) UnitValueToPoint(endpoint string, parentId string, pid PointId, name string) *Point {
	uv := t.UnitValueFix()

	if name == "" {
		name = PointToName(pid)
	}

	ret := GetPoint(parentId, pid)
	if !ret.Valid {
		ret = &Point {
			// EndPoint:  endpoint,
			// FullId:    "",
			// Parent:    ParentDevice{ Key: psId },
			Id:        pid,
			GroupName: "",
			Name:      name,
			Unit:      uv.unit,
			Type:      "PointTypeInstant",
			Valid:     true,
			States:    nil,
		}
	}

	return ret
}

// UnmarshalJSON - Convert JSON to value
func (t *UnitValue) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
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
		if t.isFloat {
			// Store result to JSON string
			data, err = json.Marshal(&struct {
				Unit  string  `json:"unit"`
				Value float64 `json:"value"`
			}{
				Unit:  t.unit,
				Value: t.float64,
			})
			break
		}

		// Store result to JSON string
		data, err = json.Marshal(&struct {
			Unit  string `json:"unit"`
			Value int64  `json:"value"`
		}{
			Unit:  t.unit,
			Value: t.int64,
		})
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

func (t UnitValue) String() string {
	return t.string
}

func (t UnitValue) Unit() string {
	return t.unit
}

func (t *UnitValue) SetString(value string) UnitValue {
	for range Only.Once {
		t.string = value
		t.int64 = 0

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		if strings.Contains(value, ".") {
			v, err := strconv.ParseFloat(t.string, 64)
			if err != nil {
				break
			}
			t.SetFloat(v)
			break
		}

		v, err := strconv.Atoi(t.string)
		if err != nil {
			break
		}
		t.SetInteger(int64(v))
	}

	return *t
}

func (t *UnitValue) SetInteger(value int64) UnitValue {
	for range Only.Once {
		t.string = ""
		t.int64 = value
		t.float64 = float64(value)
		t.isFloat = false

		if value == 0 {
			break
		}

		t.string = strconv.FormatInt(t.int64, 10)
	}

	return *t
}

func (t *UnitValue) SetFloat(value float64) UnitValue {
	for range Only.Once {
		t.string = ""
		t.int64 = int64(value)
		t.float64 = value
		t.isFloat = true

		if value == 0 {
			break
		}

		// t.string = strconv.FormatFloat(t.float64, 'f', 12, 64)
		// t.string = strings.TrimRight(t.string, "0")
		t.string = strconv.FormatFloat(t.float64, 'f', -1, 64)
	}

	return *t
}

func (t *UnitValue) SetUnit(unit string) UnitValue {
	for range Only.Once {
		t.unit = unit
	}

	return *t
}

func SetUnitValueString(value string, unit string) UnitValue {
	var t UnitValue
	t = t.SetString(value)
	t = t.SetUnit(unit)
	return t.UnitValueFix()
}

func SetUnitValueInteger(value int64, unit string) UnitValue {
	var t UnitValue
	t = t.SetInteger(value)
	t = t.SetUnit(unit)
	return t.UnitValueFix()
}

func SetUnitValueFloat(value float64, unit string) UnitValue {
	var t UnitValue
	t = t.SetFloat(value)
	t = t.SetUnit(unit)
	return t.UnitValueFix()
}


type UnitValues []UnitValue
type UnitValueMap map[PointId]UnitValue

func (u *UnitValueMap) Sort() []string {
	var ret []string
	for n := range *u {
		ret = append(ret, string(n))
	}
	sort.Strings(ret)
	return ret
}


type Integer struct {
	string `json:"string,omitempty"`
	int64  `json:"integer,omitempty"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Integer) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
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
		data, err = json.Marshal(t.int64)
		if err != nil {
			break
		}
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

func (t *Integer) SetString(value string) Integer {
	for range Only.Once {
		t.string = value
		t.int64 = 0

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		var err error
		var v int
		v, err = strconv.Atoi(t.string)
		if err != nil {
			break
		}
		t.int64 = int64(v)
	}

	return *t
}

func (t *Integer) SetValue(value int64) Integer {
	for range Only.Once {
		t.string = ""
		t.int64 = value

		if value == 0 {
			break
		}

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


type Float struct {
	string  `json:"string,omitempty"`
	float64 `json:"float,omitempty"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Float) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from int
		err = json.Unmarshal(data, &t.float64)
		if err == nil {
			t.SetValue(t.float64)
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
func (t Float) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		data, err = json.Marshal(t.float64)
		if err != nil {
			break
		}
		// t.string = strconv.FormatFloat(t.float64, 'f', 12, 64)
	}

	return data, err
}

func (t Float) Value() float64 {
	return t.float64
}

func (t Float) String() string {
	return t.string
}

func (t *Float) SetString(value string) Float {
	for range Only.Once {
		t.string = value
		t.float64 = 0

		if value == "" {
			break
		}

		if value == "--" {
			value = ""
			break
		}

		var err error
		t.float64, err = strconv.ParseFloat(t.string, 64)
		if err == nil {
			break
		}
	}

	return *t
}

func (t *Float) SetValue(value float64) Float {
	for range Only.Once {
		t.string = ""
		t.float64 = value

		if value == 0 {
			break
		}

		t.string = strconv.FormatFloat(t.float64, 'f', 12, 64)
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


type Bool struct {
	string `json:"string,omitempty"`
	bool   `json:"bool,omitempty"`
}

// UnmarshalJSON - Convert JSON to value
func (t *Bool) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		var ret2 int64
		// Store result from int
		err = json.Unmarshal(data, &ret2)
		if err == nil {
			t.SetInteger(ret2)
			break
		}

		var ret1 bool
		// Store result from int
		err = json.Unmarshal(data, &ret1)
		if err == nil {
			t.SetValue(ret1)
			break
		}

		var ret3 string
		// Store result from string
		err = json.Unmarshal(data, &ret3)
		if err == nil {
			t.SetString(ret3)
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t Bool) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		data, err = json.Marshal(t.bool)
		if err != nil {
			break
		}
		// t.string = strconv.FormatFloat(t.bool, 'f', 12, 64)
	}

	return data, err
}

func (t Bool) Value() bool {
	return t.bool
}

func (t Bool) String() string {
	return t.string
}

func (t *Bool) SetString(value string) Bool {
	for range Only.Once {
		t.string = value

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
				fallthrough
			case "--":
				t.bool = false
				t.string = "false"

			case "true":
				fallthrough
			case "yes":
				fallthrough
			case "on":
				fallthrough
			case "1":
				t.bool = true
				t.string = "true"
		}
	}

	return *t
}

func (t *Bool) SetValue(value bool) Bool {
	for range Only.Once {
		t.bool = value

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


type String struct {
	string `json:"string,omitempty"`
}

// UnmarshalJSON - Convert JSON to value
func (t *String) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
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
func (t String) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		data, err = json.Marshal(t.string)
		if err != nil {
			break
		}
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
	}

	return *t
}

func (t *String) SetValue(value string) String {
	for range Only.Once {
		t.string = value
	}

	return *t
}

func SetStringValue(value string) String {
	var t String
	return t.SetValue(value)
}


type PsKey struct {
	string `json:"string,omitempty"`
}

// UnmarshalJSON - Convert JSON to value
func (t *PsKey) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		// Store result from string
		err = json.Unmarshal(data, &t.string)
		if err == nil {
			t.SetPsKey(t.string)
			break
		}
	}

	return err
}

// MarshalJSON - Convert value to JSON
func (t PsKey) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error

	for range Only.Once {
		data, err = json.Marshal(t.string)
		if err != nil {
			break
		}
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
	}

	return *t
}

func (t *PsKey) SetValue(value string) PsKey {
	for range Only.Once {
		t.string = value
	}

	return *t
}

func SetPsKeyValue(value string) PsKey {
	var t PsKey
	return t.SetValue(value)
}


// func JsonToUnitValue(j string) UnitValue {
// 	var ret UnitValue
//
// 	for range Only.Once {
// 		err := json.Unmarshal([]byte(j), &ret)
// 		if err != nil {
// 			break
// 		}
// 	}
//
// 	return ret
// }

// func Float32ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }

// func Float64ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }

// func (u *UnitValue) GetStructName() string {
// 	var ret string
// 	apiReflect.GetName()
// 	return ret
// }
//
// func (u *UnitValue) GetJsonName() string {
// 	var ret string
// 	apiReflect.GetOptionsRequired()
// 	return ret
// }
