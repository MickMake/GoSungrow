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

	*float64    `json:"value_float,omitempty"`
	*int64      `json:"value_int,omitempty"`
	*bool       `json:"value_bool,omitempty"`

	key         string
	deviceId    string
	Valid       bool  `json:"valid"`
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

func (t *UnitValue) IsFloat() bool {
	if t.float64 != nil {
		return true
	}
	return false
}

func (t *UnitValue) IsInt() bool {
	if t.int64 != nil {
		return true
	}
	return false
}

func (t *UnitValue) IsNumber() bool {
	if t.float64 != nil {
		return true
	}
	if t.int64 != nil {
		return true
	}
	return false
}

func (t *UnitValue) IsBool() bool {
	if t.float64 != nil {
		return true
	}
	return false
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

func (t *UnitValue) IsTypeDateTime() bool {
	if t.TypeValue == "DateTime" {
		return true
	}
	return false
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
		t.SetType(UnitValueType(unit))
	}

	return *t
}

func (t *UnitValue) SetType(Type string) UnitValue {
	for range Only.Once {
		if Type == "" {
			break
		}
		t.TypeValue = Type
	}

	return *t
}

func (t *UnitValue) SetDeviceId(deviceId string) UnitValue {
	for range Only.Once {
		if deviceId == "" {
			break
		}
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

func (t *UnitValue) SetPrecision(precision int) UnitValue {
	for range Only.Once {
		if t.float64 != nil {
			v := SetPrecision(*t.float64, precision)
			t.SetFloat(v)
		}
	}
	return *t
}

func (t *UnitValue) IsZero() bool {
	var yes bool
	switch {
		case t.float64 != nil:
			if *t.float64 == 0 {
				yes = true
				break
			}
		case t.int64 != nil:
			if *t.int64 == 0 {
				yes = true
				break
			}
	}
	return yes
}

func (t *UnitValue) IsNotZero() bool {
	return !t.IsZero()
}


func SetUnitValueString(unit string, Type string, value string) UnitValue {
	var t UnitValue
	t = t.SetString(value)
	t = t.SetUnit(unit)
	t = t.SetType(Type)
	return t.UnitValueFix()
}

func SetUnitValueInteger(unit string, Type string, value int64) UnitValue {
	var t UnitValue
	t = t.SetInteger(value)
	t = t.SetUnit(unit)
	t = t.SetType(Type)
	return t.UnitValueFix()
}

func SetUnitValueFloat(unit string, Type string, value float64) UnitValue {
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


// type UnitValueMap map[PointId]UnitValue
//
// func (u *UnitValueMap) Sort() []string {
// 	var ret []string
// 	for n := range *u {
// 		ret = append(ret, n.String())
// 	}
// 	sort.Strings(ret)
// 	return ret
// }


type UnitValues struct {
	arrayValues []*UnitValue
	mapValues   map[string]*UnitValue
	mapOrder    []string

	Unit      string `json:"unit"`
	TypeValue string `json:"type_value"`
}

func (t UnitValues) String() string {
	var ret string
	for range Only.Once {
		if t.IsArray() {
			f := fmt.Sprintf("%%.%dd", SizeOfInt(len(t.arrayValues)))
			for k, v := range t.arrayValues {
				ret += fmt.Sprintf(f + ": %s\n", k, v.String())
			}
			break
		}

		if t.IsMap() {
			for _, k := range t.mapOrder {
				ret += fmt.Sprintf("%s: %s\n", k, t.mapValues[k].String())
			}
			break
		}
	}
	return ret
}

func (t *UnitValues) IsInit() bool {
	if t.arrayValues != nil {
		return false
	}
	if t.mapValues != nil {
		return false
	}
	return true
}

func (t *UnitValues) IsArray() bool {
	if t.arrayValues != nil {
		return true
	}
	return false
	// return !*t.ismapValues
}

func (t *UnitValues) IsMap() bool {
	if t.mapValues != nil {
		return true
	}
	return false
	// return *t.ismapValues
}

// func (t UnitValues) firstKey() string {
// 	var ret string
// 	for range Only.Once {
// 		if len(t.values) == 0 {
// 			break
// 		}
// 		for k := range t.values {
// 			ret = k
// 			break
// 		}
// 	}
// 	return ret
// }

func (t *UnitValues) GetUnit() string {
	var ret string
	for range Only.Once {
		if t.IsArray() {
			for _, v := range t.arrayValues {
				ret = v.Unit()
				break
			}
			break
		}

		if t.IsMap() {
			for _, k := range t.mapOrder {
				ret = t.mapValues[k].Unit()
				break
			}
			break
		}
	}
	return ret
}

func (t *UnitValues) GetmapValues() map[string]*UnitValue {
	return t.mapValues
}

func (t *UnitValues) GetarrayValues() []*UnitValue {
	return t.arrayValues
}

const SortOrder = false
const LoadOrder = true
func (t *UnitValues) Range(loadOrder bool) []UnitValue {
	var ret []UnitValue
	for range Only.Once {
		if t.IsArray() {
			for _, k := range t.arrayValues {
				ret = append(ret, *k)
			}
			break
		}

		if t.IsMap() {
			if loadOrder {
				for _, k := range t.mapOrder {
					ret = append(ret, *t.mapValues[k])
				}
				break
			}

			var keys []string
			for k := range t.mapValues {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				ret = append(ret, *t.mapValues[k])
			}
			break
		}
	}
	return ret
}


func (t *UnitValues) Keys(loadOrder bool) []string {
	var ret []string
	for range Only.Once {
		if t.IsArray() {
			for k := range t.arrayValues {
				ret = append(ret, strconv.Itoa(k))
			}
			break
		}

		if t.IsMap() {
			if loadOrder {
				ret = t.mapOrder
				break
			}

			for k := range t.mapValues {
				ret = append(ret, k)
			}
			sort.Strings(ret)
			break
		}
	}
	return ret
}

func (t *UnitValues) Type() string {
	var ret string
	for range Only.Once {
		if t.IsArray() {
			for _, v := range t.arrayValues {
				ret = v.Type()
			}
			break
		}

		if t.IsMap() {
			for _, v := range t.mapValues {
				ret = v.Type()
			}
			break
		}
	}
	return ret
}

func (t *UnitValues) SetType(Type string) *UnitValues {
	for range Only.Once {
		if Type == "" {
			break
		}

		t.TypeValue = Type

		if t.IsArray() {
			for k := range t.arrayValues {
				t.arrayValues[k].SetType(Type)
			}
			break
		}

		if t.IsMap() {
			for k := range t.mapValues {
				t.mapValues[k].SetType(Type)
			}
			break
		}
	}
	return t
}

func (t *UnitValues) SetUnit(unit string) *UnitValues {
	for range Only.Once {
		t.Unit = unit

		if t.IsArray() {
			for k := range t.arrayValues {
				t.arrayValues[k].SetUnit(unit)
			}
			break
		}

		if t.IsMap() {
			for k := range t.mapValues {
				t.mapValues[k].SetUnit(unit)
			}
			break
		}
	}
	return t
}

func (t *UnitValues) SetDeviceId(deviceId string) *UnitValues {
	for range Only.Once {
		if t.IsArray() {
			for k := range t.arrayValues {
				t.arrayValues[k].SetDeviceId(deviceId)
			}
			break
		}

		if t.IsMap() {
			for k := range t.mapValues {
				t.mapValues[k].SetDeviceId(deviceId)
			}
			break
		}
	}

	return t
}

func (t *UnitValues) GetIndex(index int) *UnitValue {
	var ret *UnitValue
	for range Only.Once {
		if t.IsArray() {
			if index >= len(t.arrayValues) {
				break
			}
			ret = t.arrayValues[index]
			break
		}

		if t.IsMap() {
			if index >= len(t.mapOrder) {
				break
			}
			key := t.mapOrder[index]
			ret = t.mapValues[key]
			break
		}
	}
	return ret
}

func (t *UnitValues) GetKey(key string) *UnitValue {
	var ret *UnitValue
	for range Only.Once {
		if t.IsArray() {
			// Doesn't make sense to return anything.
			break
		}

		if t.IsMap() {
			if m, ok := t.mapValues[key]; ok {
				ret = m
			}
			break
		}
	}
	return ret
}

func (t *UnitValues) First() *UnitValue {
	var ret *UnitValue
	for range Only.Once {
		if t.IsArray() {
			if len(t.arrayValues) == 0 {
				break
			}
			ret = t.arrayValues[0]
			break
		}

		if t.IsMap() {
			if len(t.mapOrder) == 0 {
				break
			}
			key := t.mapOrder[0]
			ret = t.mapValues[key]
			break
		}
	}
	return ret
}

func (t *UnitValues) Last() *UnitValue {
	var ret *UnitValue
	for range Only.Once {
		if t.IsArray() {
			if len(t.arrayValues) == 0 {
				break
			}
			ret = t.arrayValues[len(t.arrayValues) - 1]
			break
		}

		if t.IsMap() {
			if len(t.mapOrder) == 0 {
				break
			}
			key := t.mapOrder[len(t.mapOrder) - 1]
			ret = t.mapValues[key]
			break
		}
	}
	return ret
}

func (t *UnitValues) SetPrecision(precision int) {
	for range Only.Once {
		if t.IsArray() {
			for k := range t.arrayValues {
				t.arrayValues[k].SetPrecision(precision)
			}
			break
		}

		if t.IsMap() {
			for k := range t.mapValues {
				t.mapValues[k].SetPrecision(precision)
			}
			break
		}
	}
}


// mapValuess.

func (t *UnitValues) addMap(key string, value UnitValue) *UnitValues {
	for range Only.Once {
		t.setupMap()

		if t.IsArray() {
			// Adding a map to an array!
			// What to do, what to do?
			break
		}

		if t.mapOrder == nil {
			t.mapOrder = make([]string, 0)
		}
		t.mapOrder = append(t.mapOrder, key) // Keep track of the order.

		if t.mapValues == nil {
			t.mapValues = make(map[string]*UnitValue)
		}
		if value.key == "" {
			value.key = key
		}
		t.mapValues[key] = &value
	}
	return t
}

func (t *UnitValues) setupMap() bool {
	var yes bool
	for range Only.Once {
		if !t.IsInit() {
			break
		}
		t.mapValues = make(map[string]*UnitValue)
		t.mapOrder = make([]string, 0)
		yes = true
	}
	return yes
}

func (t *UnitValues) AddUnitValues(key string, uvs UnitValues) *UnitValues {
	for range Only.Once {
		t.AddUnitValue(key, uvs.Range(true)...)
	}
	return t
}

func (t *UnitValues) AddUnitValue(key string, uvs ...UnitValue) *UnitValues {
	for range Only.Once {
		t.setupMap()

		if t.IsArray() {
			// Doesn't make sense.
			// for _, uv := range uvs {
			// 	t.addmapValues(key, uv)
			// }
			break
		}

		if t.IsMap() {
			for _, uv := range uvs {
				t.addMap(key, uv)
			}
		}
	}
	return t
}

const IsarrayValues = "array"
func (t *UnitValues) AddString(key string, unit string, Type string, value ...string) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueString(unit, Type, v)
		if key == IsarrayValues {
			t.appendArray(uv)
			break
		}
		t.addMap(key, uv)
	}
	return t
}

func (t *UnitValues) AddBool(key string, value ...bool) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueBool(v)
		if key == IsarrayValues {
			t.appendArray(uv)
			break
		}
		t.addMap(key, uv)
	}
	return t
}

func (t *UnitValues) AddFloat(key string, unit string, Type string, value ...float64) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueFloat(unit, Type, v)
		if key == IsarrayValues {
			t.appendArray(uv)
			break
		}
		t.addMap(key, uv)
	}
	return t
}

func (t *UnitValues) AddInteger(key string, unit string, Type string, value ...int64) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueInteger(unit, Type, v)
		if key == IsarrayValues {
			t.appendArray(uv)
			break
		}
		t.addMap(key, uv)
	}
	return t
}


// arrayValuess.

func (t *UnitValues) appendArray(value UnitValue) *UnitValues {
	for range Only.Once {
		t.setupArray()

		if t.IsMap() {
			// Adding an array to a map!
			// What to do, what to do?
			break
		}

		if t.arrayValues == nil {
			t.arrayValues = make([]*UnitValue, 0)
		}
		if value.key == "" {
			value.key = strconv.Itoa(len(t.arrayValues))
		}
		t.arrayValues = append(t.arrayValues, &value)
	}
	return t
}

func (t *UnitValues) setupArray() bool {
	var yes bool
	for range Only.Once {
		if !t.IsInit() {
			break
		}
		t.arrayValues = make([]*UnitValue, 0)
		yes = true
	}
	return yes
}

func (t *UnitValues) AppendUnitValues(uvs UnitValues) *UnitValues {
	for range Only.Once {
		t.AppendUnitValue(uvs.Range(true)...)
	}
	return t
}

func (t *UnitValues) AppendUnitValue(uvs ...UnitValue) *UnitValues {
	for range Only.Once {
		t.setupArray()

		if t.IsArray() {
			for _, uv := range uvs {
				t.appendArray(uv)
			}
		}

		if t.IsMap() {
			for _, uv := range uvs {
				t.addMap(uv.key, uv)
			}
		}
	}
	return t
}

// func (t *UnitValues) Append(unit string, Type string, value string) *UnitValues {
// 	for range Only.Once {
// 		t.Reset()
// 		uv := SetUnitValueString(unit, Type, value)
// 		t.addMap(key, uv)
// 	}
// 	return t
// }

func (t *UnitValues) AppendString(unit string, Type string, value ...string) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueString(unit, Type, v)
		t.appendArray(uv)
	}
	return t
}

func (t *UnitValues) AppendBool(value ...bool) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueBool(v)
		t.appendArray(uv)
	}
	return t
}

func (t *UnitValues) AppendFloat(unit string, Type string, value ...float64) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueFloat(unit, Type, v)
		t.appendArray(uv)
	}
	return t
}

func (t *UnitValues) AppendInteger(unit string, Type string, value ...int64) *UnitValues {
	for _, v := range value {
		uv := SetUnitValueInteger(unit, Type, v)
		t.appendArray(uv)
	}
	return t
}


func (t *UnitValues) Reset() *UnitValues {
	t.arrayValues = nil
	t.mapValues = nil
	t.mapOrder = nil
	t.Unit = ""
	t.TypeValue = ""
	return t
}

// func (t *UnitValues) AddUnitValue(value UnitValue) *UnitValues {
// 	return t.add(value)
// }

// func (t *UnitValues) SetUnitValue(uv UnitValue) *UnitValues {
// 	for range Only.Once {
// 		t.values = map[string]UnitValue{}
// 		t.order = []*UnitValue{}
// 		t.UnitValue = uv.UnitValue
// 		t.TypeValue = uv.TypeValue
// 		t.add(uv)
// 	}
// 	return t
// }

func (t *UnitValues) Length() int {
	var ret int
	for range Only.Once {
		if t.IsArray() {
			ret = len(t.arrayValues)
			break
		}

		if t.IsMap() {
			ret = len(t.mapValues)
			break
		}
	}
	return ret
}


// func Float32ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }

// func Float64ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }
