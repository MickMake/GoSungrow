package valueTypes

import (
	"GoSungrow/Only"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


const (
	TypeBool      = "Bool"
	TypeCount     = "Count"
	TypeDateTime  = "DateTime"
	TypeFloat     = "Float"
	TypeInteger   = "Integer"
	TypePointId   = "PointId"
	TypePsKey     = "PsKey"
	TypePsId      = "PsId"
	TypeString    = "String"
	TypeTime      = "Time"
	TypeUnitValue = "UnitValue"

	TypeArrayBool      = "[]Bool"
	TypeArrayCount     = "[]Count"
	TypeArrayDateTime  = "[]DateTime"
	TypeArrayFloat     = "[]Float"
	TypeArrayInteger   = "[]Integer"
	TypeArrayPointId   = "[]PointId"
	TypeArrayPsKey     = "[]PsKey"
	TypeArrayPsId      = "[]PsId"
	TypeArrayString    = "[]String"
	TypeArrayTime      = "[]Time"
	TypeArrayUnitValue = "[]UnitValue"
	TypeUnitValues     = "UnitValues"

	// TypeArrayValueTypesBool      = "[]valueTypes.Bool"
	// TypeArrayValueTypesCount     = "[]valueTypes.Count"
	// TypeArrayValueTypesDateTime  = "[]valueTypes.DateTime"
	// TypeArrayValueTypesFloat     = "[]valueTypes.Float"
	// TypeArrayValueTypesInteger   = "[]valueTypes.Integer"
	// TypeArrayValueTypesPointId   = "[]valueTypes.PointId"
	// TypeArrayValueTypesPsKey     = "[]valueTypes.PsKey"
	// TypeArrayValuePsId           = "[]valueTypes.PsId"
	// TypeArrayValueTypesString    = "[]valueTypes.String"
	// TypeArrayValueTypesTime      = "[]valueTypes.Time"
	// TypeArrayValueTypesUnitValue = "[]valueTypes.UnitValue"
	//
	// TypeValueTypesBool      = "valueTypes.Bool"
	// TypeValueTypesCount     = "valueTypes.Count"
	// TypeValueTypesDateTime  = "valueTypes.DateTime"
	// TypeValueTypesFloat     = "valueTypes.Float"
	// TypeValueTypesInteger   = "valueTypes.Integer"
	// TypeValueTypesPointId   = "valueTypes.PointId"
	// TypeValueTypesPsKey     = "valueTypes.PsKey"
	// TypeValueTypesPsId      = "valueTypes.PsId"
	// TypeValueTypesString    = "valueTypes.String"
	// TypeValueTypesTime      = "valueTypes.Time"
	// TypeValueTypesUnitValue = "valueTypes.UnitValue"
)


func IsUnknownStruct(ref interface{}) bool {
	var ok bool

	for range Only.Once {
		fieldVo := reflect.ValueOf(ref)
		// fieldTo := reflect.TypeOf(ref)
		// fmt.Printf("fieldVo.Type().String(): %s\n", fieldVo.Type().String())
		// fmt.Printf("fieldVo.Type().Name(): %s\n", fieldVo.Type().Name())
		// fmt.Printf("fieldVo.Kind().String(): %s\n", fieldVo.Kind().String())
		// fmt.Printf("fieldTo.String(): %s\n", fieldTo.String())
		// fmt.Printf("fieldTo.Name(): %s\n", fieldTo.Name())
		// fmt.Printf("fieldTo.Kind().String(): %s\n", fieldTo.Kind().String())

		kindy := fieldVo.Kind()
		// fmt.Printf("DEBUYg:    K:%s / T:%v\n", kindy.String(), fieldVo)
		if kindy == reflect.Interface {
			ok = false
			break
		}

		if kindy == reflect.Slice {
			if fieldVo.Len() > 0 {
				ok = IsUnknownStruct(fieldVo.Index(0).Interface())
			}
			break
		}

		if kindy == reflect.Array {
			if fieldVo.Len() > 0 {
				ok = IsUnknownStruct(fieldVo.Index(0).Interface())
			}
			break
		}

		if kindy == reflect.Map {
			mk := fieldVo.MapKeys()
			if len(mk) > 0 {
				ok = IsUnknownStruct(fieldVo.MapIndex(mk[0]).Interface())
			}
			break
		}

		if kindy == reflect.Struct {
			Type := fieldVo.Type().String()
			Type = strings.ReplaceAll(Type, "valueTypes.", "")
			switch Type {
				case TypeBool:
				case TypeCount:
				case TypeDateTime:
				case TypeFloat:
				case TypeInteger:
				case TypePointId:
				case TypePsId:
				case TypePsKey:
				case TypeString:
				case TypeTime:
				case TypeUnitValue:
				case TypeArrayBool:
				case TypeArrayCount:
				case TypeArrayDateTime:
				case TypeArrayFloat:
				case TypeArrayInteger:
				case TypeArrayPointId:
				case TypeArrayPsId:
				case TypeArrayPsKey:
				case TypeArrayString:
				case TypeArrayTime:
				case TypeArrayUnitValue:
				case TypeUnitValues:

				default:
					ok = true
			}
		}
	}

	return ok
}

func IsNil(i interface{}) bool {
	var ok bool
	for range Only.Once {
		if i == nil {
			ok = true
			break
		}

		switch reflect.TypeOf(i).Kind() {
			case reflect.Ptr:
				fallthrough
			case reflect.Map:
				fallthrough
			case reflect.Array:
				fallthrough
			case reflect.Chan:
				fallthrough
			case reflect.Slice:
				ok = reflect.ValueOf(i).IsNil()
				break
		}
	}
	return ok
}

func PrintInt(s int, i interface{}) string {
	var ret string
	for range Only.Once {
		var val int64
		Type := reflect.TypeOf(i).String()
		Type = strings.ReplaceAll(Type, "valueTypes.", "")
		switch Type {
			case "int":
				val = int64(i.(int))
			case "int32":
				val = int64(i.(int32))
			case "int64":
				val = i.(int64)

			case TypeInteger:
				val = i.(Integer).Value()

			case TypeCount:
				val = i.(Count).Value()
		}

		if s == 0 {
			ret = fmt.Sprintf("%d", val)
			break
		}

		ret = fmt.Sprintf("%." + strconv.Itoa(s) + "d", val)
	}
	return ret
}

func SizeOfInt(ref interface{}) int {
	var ret int
	for range Only.Once {
		var val int64
		Type := reflect.TypeOf(ref).String()
		Type = strings.ReplaceAll(Type, "valueTypes.", "")
		switch Type {
			case "int":
				val = int64(ref.(int))
			case "int32":
				val = int64(ref.(int32))
			case "int64":
				val = ref.(int64)

			case TypeInteger:
				val = ref.(Integer).Value()

			case TypeCount:
				val = ref.(Count).Value()
		}
		switch {
			case val > 9999:
				ret = 5
			case val > 999:
				ret = 4
			case val > 99:
				ret = 3
			case val > 9:
				ret = 2
			default:
				ret = 1
		}
	}
	return ret
}

func GetIntFormatForPrintf(i interface{}) string {
	var ret string
	for range Only.Once {
		s := SizeOfInt(i)
		if s == 0 {
			ret = "%d"
			break
		}
		ret = "%." + strconv.Itoa(s) + "d"
	}
	return ret
}

func ArrayLength(i interface{}) int {
	return reflect.ValueOf(i).Len()
}

func SizeOfArrayLength(i interface{}) int {
	return SizeOfInt(reflect.ValueOf(i).Len())
}

func AnyToUnitValue(ref interface{}, unit string, typeString string, dateFormat string) (UnitValues, bool, bool) {
	var uv UnitValues
	ok := true
	isNil := false
	for range Only.Once {
		if IsNil(ref) {
			// fmt.Println("DEBUG: AnyToUnitValue(): NIL")
			uv = append(uv, SetUnitValueString("", unit, typeString + "(unknown)"))
			isNil = true
			break
		}
		// fmt.Printf("DEBUG: AnyToUnitValue(): %s\n", reflect.TypeOf(e).String())

		if dateFormat == "" {
			dateFormat = DateTimeAltLayout
		}

		Type := reflect.TypeOf(ref).String()
		Type = strings.ReplaceAll(Type, "valueTypes.", "")
		switch Type {
			case "int":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueInteger(int64(ref.(int)), unit, typeString))
			case "int32":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueInteger(int64(ref.(int32)), unit, typeString))
			case "int64":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueInteger(ref.(int64), unit, typeString))
			case "float32":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueFloat(float64(ref.(float32)), unit, typeString))
			case "float64":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueFloat(ref.(float64), unit, typeString))
			case "string":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueString(ref.(string), unit, typeString))
			case "[]string":
				// v := strings.Join(e.([]string), ",")
				if typeString == "" {
					typeString = "--"
				}
				j, err := json.Marshal(ref.([]string))
				if err != nil {
					j = []byte(fmt.Sprintf("%v", ref.([]string)))
				}
				uv = append(uv, SetUnitValueString(string(j), unit, typeString))
			case "bool":
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, SetUnitValueBool(ref.(bool)))

			case TypeUnitValue:
				if typeString == "" {
					typeString = "--"
				}
				uv = append(uv, ref.(UnitValue))
				// uv = uv.UnitValueFix()
			case TypeUnitValues:
				fallthrough
			case TypeArrayUnitValue:
				for _, val := range ref.([]UnitValue) {
					uv = append(uv, val)
				}

			case TypeFloat:
				if typeString == "" {
					typeString = TypeFloat
				}
				v := ref.(Float)
				uv = append(uv, SetUnitValueFloat(v.Value(), unit, typeString))
			case TypeArrayFloat:
				if typeString == "" {
					typeString = TypeFloat
				}
				v := ref.([]Float)
				for _, val := range v {
					uv = append(uv, SetUnitValueFloat(val.Value(), unit, typeString))
				}

			case TypeInteger:
				if typeString == "" {
					typeString = TypeInteger
				}
				v := ref.(Integer)
				uv = append(uv, SetUnitValueInteger(v.Value(), unit, typeString))
			case TypeArrayInteger:
				if typeString == "" {
					typeString = TypeInteger
				}
				v := ref.([]Integer)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, typeString))
				}
				// HERE IS THE PROBLEM - need to return SOMETHING, even if it's null!

			case TypeCount:
				if typeString == "" {
					typeString = TypeCount
				}
				v := ref.(Count)
				uv = append(uv, SetUnitValueInteger(v.Value(), unit, typeString))
			case TypeArrayCount:
				if typeString == "" {
					typeString = TypeCount
				}
				v := ref.([]Count)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, typeString))
				}

			case TypeBool:
				if typeString == "" {
					typeString = TypeBool
				}
				v := ref.(Bool)
				uv = append(uv, SetUnitValueBool(v.Value()))
			case TypeArrayBool:
				if typeString == "" {
					typeString = TypeBool
				}
				v := ref.([]Bool)
				for _, val := range v {
					uv = append(uv, SetUnitValueBool(val.Value()))
				}

			case TypeString:
				if typeString == "" {
					typeString = TypeString
				}
				v := ref.(String)
				uv = append(uv, SetUnitValueString(v.String(), unit, typeString))
			case TypeArrayString:
				if typeString == "" {
					typeString = TypeString
				}
				v := ref.([]String)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, typeString))
				}

			case TypePsId:
				if typeString == "" {
					typeString = TypePsId
				}
				v := ref.(PsId)
				uv = append(uv, SetUnitValueString(v.String(), unit, typeString))

			case TypeArrayPsId:
				if typeString == "" {
					typeString = TypePsId
				}
				v := ref.([]PsId)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, typeString))
				}

			case TypePsKey:
				if typeString == "" {
					typeString = TypePsKey
				}
				v := ref.(PsKey)
				uv = append(uv, SetUnitValueString(v.Value(), unit, typeString))

			case TypeArrayPsKey:
				if typeString == "" {
					typeString = TypePsKey
				}
				v := ref.([]PsKey)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, typeString))
				}

			case TypePointId:
				if typeString == "" {
					typeString = TypePointId
				}
				v := ref.(PointId)
				uv = append(uv, SetUnitValueString(v.String(), unit, typeString))
			case TypeArrayPointId:
				if typeString == "" {
					typeString = TypePointId
				}
				v := ref.([]PointId)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, typeString))
				}

			case TypeDateTime:
				if typeString == "" {
					typeString = TypeDateTime
				}
				v := ref.(DateTime)
				uv = append(uv, SetUnitValueString(v.Format(dateFormat), unit, typeString))
			case TypeArrayDateTime:
				if typeString == "" {
					typeString = TypeDateTime
				}
				v := ref.([]DateTime)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Format(dateFormat), unit, typeString))
				}

			case TypeTime:
				if typeString == "" {
					typeString = TypeTime
				}
				v := ref.(Time)
				uv = append(uv, SetUnitValueString(v.Format(TimeLayout), unit, typeString))
			case TypeArrayTime:
				if typeString == "" {
					typeString = TypeTime
				}
				v := ref.([]Time)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Format(TimeLayout), unit, typeString))
				}

			default:
				uv = append(uv, SetUnitValueString("", unit, typeString+ "(unknown)"))
				ok = false
		}
	}
	return uv, isNil, ok
}

func AnyToValueString(ref interface{}, intSize int, dateFormat string) string {
	var ret string

	for range Only.Once {
		if IsNil(ref) {
			break
		}

		// fmt.Printf("DEBUG TYPE: %s\n", reflect.TypeOf(e).String())
		Type := reflect.TypeOf(ref).String()
		Type = strings.ReplaceAll(Type, "valueTypes.", "")
		switch Type {
			case "bool":
				ret = fmt.Sprintf("%v", ref.(bool))
			case "int":
				ret = PrintInt(intSize, ref.(int))
			case "int32":
				ret = PrintInt(intSize, ref.(int32))
			case "int64":
				ret = PrintInt(intSize, ref.(int64))
			case "float32":
				// ret = float64(s.(float32))
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)
			case "float64":
				// ret = s.(float64)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)
			case "string":
				ret = ref.(string)
				// ret = strings.Trim(ref.(string), ".")
			case "[]string":
				// v := strings.Join(s.([]string), ",")
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeUnitValue:
				ret = ref.(UnitValue).String()
			case TypeUnitValues:
				fallthrough
			case TypeArrayUnitValue:
				// ret = s.([]UnitValue)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeFloat:
				ret = ref.(Float).String()
				// v , err := json.Marshal(ref)
				// if err != nil {
				// 	break
				// }
				// ret = string(v)
			case TypeArrayFloat:
				// ret = s.([]Float)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeInteger:
				ret = PrintInt(intSize, ref.(Integer))
			case TypeArrayInteger:
				// ret = s.([]Integer)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeCount:
				ret = PrintInt(intSize, ref.(Count))
			case TypeArrayCount:
				// ret = s.([]Count)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeBool:
				ret = ref.(Bool).String()
			case TypeArrayBool:
				// ret = s.([]Bool)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeString:
				ret = ref.(String).String()
				// ret = strings.Trim(ref.(String).String(), ".")
			case TypeArrayString:
				// ret = s.([]String)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypePsKey:
				ret = ref.(PsKey).String()
			case TypeArrayPsKey:
				// ret = s.([]PsKey)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypePsId:
				ret = ref.(PsId).String()
			case TypeArrayPsId:
				// ret = s.([]PsId)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypePointId:
				ret = ref.(PointId).String()
			case TypeArrayPointId:
				// ret = s.([]PointId)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeDateTime:
				if dateFormat == "" {
					dateFormat = DateTimeAltLayout
				}
				v := ref.(DateTime)
				if v.IsZero() {
					ret = ""
					break
				}
				ret = v.Format(dateFormat)
			case TypeArrayDateTime:
				// ret = s.([]DateTime)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeTime:
				if dateFormat == "" {
					dateFormat = TimeLayoutSecond
				}
				v := ref.(Time)
				if v.IsZero() {
					ret = ""
					break
				}
				ret = v.Format(dateFormat)
			case TypeArrayTime:
				// ret = s.([]Time)
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)

			default:
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
