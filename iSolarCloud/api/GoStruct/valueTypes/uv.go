package valueTypes

import (
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"sort"
	"strconv"
	"strings"
)


type UnitValue struct {
	UnitValue   string `json:"unit"`  			// Primary iSolarCloud entity.
	StringValue string `json:"value"` 			// Primary iSolarCloud entity.

	TypeValue   string `json:"type_value"`

	*float64     `json:"value_float,omitempty"`
	*int64       `json:"value_int,omitempty"`
	*bool        `json:"value_bool,omitempty"`

	key         string
	deviceId    string
	Valid       bool `json:"valid"`
	Error       error `json:"-"`
}

var zero = int64(0)

func (t *UnitValue) UnitValueFix() UnitValue {
	switch t.UnitValue {
		case "w":
			t.UnitValue = "W"
	}

	switch t.UnitValue {
		case "g":
			fallthrough
		case "Wp":
			fallthrough
		case "Wh":
			fallthrough
		case "W":
			if t.float64 == nil {
				// Only if we have a float.
				break
			}
			fv := t.Value() / 1000
			t.SetFloat(fv)
			t.SetUnit("k" + t.UnitValue)

			// fv, dt.Error := strconv.ParseFloat(value, 64)
			// if dt.Error == nil {
			// 	fv = fv / 1000
			// 	value, _ = DivideByThousand(value)
			// 	UnitValue = "k" + UnitValue
			// }
	}

	return *t
}

func UnitValueType(unit string) string {
	var ret string
	switch unit {
		case "Wh":
			fallthrough
		case "kWh":
			fallthrough
		case "MWh":
			ret = "Energy"

		case "kWp":
			fallthrough
		case "W":
			fallthrough
		case "kW":
			fallthrough
		case "MW":
			ret = "Power"

		case "AUD":
			ret = "Currency"

		case "g":
			fallthrough
		case "kg":
			ret = "Weight"

		case "mV":
			fallthrough
		case "V":
			ret = "Voltage"

		case "mA":
			fallthrough
		case "A":
			ret = "Current"

		case "Hz":
			ret = "Frequency"

		case "kvar":
			ret = "Reactive Power"

		case "Ω":
			fallthrough
		case "kΩ":
			ret = "Resistance"

		case "%":
			ret = "Percent"

		case "F":
			fallthrough
		case "°F":
			fallthrough
		case "℉":
			fallthrough
		case "C":
			fallthrough
		case "°C":
			fallthrough
		case "℃":
			ret = "Temperature"
	}
	return ret
}

// UnmarshalJSON - Convert JSON to value
func (t *UnitValue) UnmarshalJSON(data []byte) error {
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
		t.Error = json.Unmarshal(data, &resString)
		if t.Error == nil {
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
		t.Error = json.Unmarshal(data, &resInteger)
		if t.Error == nil {
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
		t.Error = json.Unmarshal(data, &resFloat)
		if t.Error == nil {
			t.SetFloat(resFloat.Value)
			t.SetUnit(resFloat.Unit)
			t.UnitValueFix()
			break
		}

		t.Error = nil
	}

	return t.Error
}

// MarshalJSON - Convert value to JSON
func (t UnitValue) MarshalJSON() ([]byte, error) {
	var data []byte

	for range Only.Once {
		t.Valid = false

		if t.float64 != nil {
			// Store result to JSON string
			data, t.Error = json.Marshal(&struct {
				Unit  string  `json:"unit"`
				Value float64 `json:"value"`
			}{
				Unit:  t.UnitValue,
				Value: *t.float64,
			})
			if t.Error != nil {
				break
			}

			t.Valid = true
			break
		}

		if t.int64 != nil {
			// Store result to JSON string
			data, t.Error = json.Marshal(&struct {
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
			}{
				Unit:  t.UnitValue,
				Value: *t.int64,
			})
			if t.Error != nil {
				break
			}
		}

		if t.bool != nil {
			// Store result to JSON string
			data, t.Error = json.Marshal(&struct {
				Unit  string `json:"unit"`
				Value bool   `json:"value"`
			}{
				Unit:  t.UnitValue,
				Value: *t.bool,
			})
			if t.Error != nil {
				break
			}
		}

		// t = t.UnitValueFix()

		data = []byte(fmt.Sprintf(`{"unit":"%s","value":"--","key":"%s"}`, t.UnitValue, t.key))

		t.Error = nil
		t.Valid = true
	}

	return data, t.Error
}

func (t *UnitValue) Value() float64 {
	var ret float64
	for range Only.Once {
		if t.float64 != nil {
			ret = *t.float64
			break
		}

		if t.int64 != nil {
			ret = float64(*t.int64)
			break
		}

		if t.bool != nil {
			if *t.bool {
				ret = 1
				break
			}
			ret = 0
			break
		}
	}
	return ret
}

func (t *UnitValue) ValueFloat() float64 {
	if t.float64 == nil {
		return 0
	}
	return *t.float64
}

func (t *UnitValue) ValueInt() int64 {
	if t.int64 == nil {
		return 0
	}
	return *t.int64
}

func (t *UnitValue) ValueBool() bool {
	if t.float64 == nil {
		return false
	}
	return *t.bool
}

func (t *UnitValue) ValueKey() string {
	return t.key
}

func (t UnitValue) String() string {
	var ret string
	for range Only.Once {
		switch {
			case t.float64 != nil:
				ret = strconv.FormatFloat(*t.float64, 'f', -1, 64)

			case t.int64 != nil:
				ret = strconv.FormatInt(*t.int64, 10)

			case t.bool != nil:
				ret = strconv.FormatBool(*t.bool)
			default:
				ret = t.StringValue
		}

		if t.key != "" {
			ret = fmt.Sprintf(`"%s": "%s"`, t.key, ret)
		}
	}
	return ret
}

func (t *UnitValue) MatchFloat(comp float64) bool {
	if t.float64 == nil {
		return false
	}
	if *t.float64 == comp {
		return true
	}
	return false
}

func (t *UnitValue) MatchInt(comp int64) bool {
	if t.int64 == nil {
		return false
	}
	if *t.int64 == comp {
		return true
	}
	return false
}

func (t *UnitValue) MatchBool(comp bool) bool {
	if t.bool == nil {
		return false
	}
	if *t.bool == comp {
		return true
	}
	return false
}

func (t *UnitValue) Unit() string {
	return t.UnitValue
}

func (t *UnitValue) Type() string {
	return t.TypeValue
}

func (t *UnitValue) DeviceId() string {
	return t.deviceId
}

var varTrue = true
var varFalse = false
func (t *UnitValue) SetString(value string) UnitValue {
	for range Only.Once {
		t.StringValue = value
		t.float64 = nil
		t.int64 = nil
		t.bool = nil
		t.Valid = false

		if value == "" {
			break
		}

		if value == "--" {
			// value = ""
			break
		}

		if value == "true" {
			t.SetBool(true)
			break
		}
		if value == "false" {
			t.SetBool(false)
			break
		}

		if strings.Contains(value, ".") {
			var v float64
			v, t.Error = strconv.ParseFloat(t.StringValue, 64)
			if t.Error != nil {
				t.Error = nil	// Less than useful.
				break
			}
			t.SetFloat(v)
			break
		}

		var v int
		v, t.Error = strconv.Atoi(t.StringValue)
		if t.Error != nil {
			t.Error = nil	// Less than useful.
			break
		}
		t.SetInteger(int64(v))
	}

	return *t
}

func (t *UnitValue) SetInteger(value int64) UnitValue {
	for range Only.Once {
		t.int64 = &value
		// fv := float64(value); t.float64 = &fv
		t.float64 = nil
		t.bool = nil
		t.Valid = true
		t.StringValue = strconv.FormatInt(*t.int64, 10)
	}

	return *t
}

func (t *UnitValue) SetFloat(value float64) UnitValue {
	for range Only.Once {
		// iv := int64(value)
		// t.int64 = &iv
		t.int64 = nil
		t.float64 = &value
		t.bool = nil
		t.Valid = true
		// t.String = strconv.FormatFloat(t.float64, 'f', 12, 64)
		// t.String = strings.TrimRight(t.String, "0")
		t.StringValue = strconv.FormatFloat(*t.float64, 'f', -1, 64)
	}

	return *t
}

func (t *UnitValue) SetBool(value bool) UnitValue {
	for range Only.Once {
		t.Valid = true
		t.float64 = nil
		t.int64 = nil
		t.bool = &value
		t.StringValue = strconv.FormatBool(value)
	}

	return *t
}

func (t *UnitValue) SetBoolString(value string) UnitValue {
	for range Only.Once {
		t.Valid = true
		t.float64 = nil
		t.int64 = nil
		// t.StringValue = strconv.FormatBool(value)

		switch strings.ToLower(value) {
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
				t.bool = &varFalse
				t.StringValue = "false"
				t.Valid = true

			case "true":
				fallthrough
			case "yes":
				fallthrough
			case "on":
				fallthrough
			case "1":
				t.bool = &varTrue
				t.StringValue = "true"
				t.Valid = true
		}
	}

	return *t
}

func (t *UnitValue) SetUnit(unit string) UnitValue {
	for range Only.Once {
		t.UnitValue = unit
		t.TypeValue = UnitValueType(unit)
	}

	return *t
}

func (t *UnitValue) SetType(Type string) UnitValue {
	for range Only.Once {
		t.TypeValue = Type
	}

	return *t
}

func (t *UnitValue) SetDeviceId(deviceId string) UnitValue {
	for range Only.Once {
		t.deviceId = deviceId
	}

	return *t
}

func (t *UnitValue) SetKey(key string) UnitValue {
	for range Only.Once {
		t.key = key
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


type UnitValueMap map[PointId]UnitValue

func (u *UnitValueMap) Sort() []string {
	var ret []string
	for n := range *u {
		ret = append(ret, n.String())
	}
	sort.Strings(ret)
	return ret
}


type UnitValues []UnitValue

func (t UnitValues) String() string {
	var ret string
	for range Only.Once {
		if len(t) == 0 {
			break
		}

		if len(t) == 1 {
			ret = t[0].String()
			break
		}

		for i, v := range t {
			ret += fmt.Sprintf("%d: %s\n", i, v.String())
		}
	}
	return ret
}

func (t *UnitValues) Unit() string {
	var ret string
	for _, v := range *t {
		ret = v.Unit()
		break
	}
	return ret
}

func (t *UnitValues) Type() string {
	var ret string
	for _, v := range *t {
		ret = v.Type()
		break
	}
	return ret
}

func (t *UnitValues) SetUnit(unit string) *UnitValue {
	return &(*t)[0]
}

func (t *UnitValues) First() *UnitValue {
	if len(*t) == 0 {
		return &UnitValue{}
	}
	return &(*t)[0]
}

func (t *UnitValues) Set(uv UnitValue) *UnitValues {
	*t = UnitValues{uv}
	return t
}

func (t *UnitValues) Append(uv UnitValue) *UnitValues {
	*t = append(*t, uv)
	return t
}

func (t *UnitValues) Length() int {
	return len(*t)
}
