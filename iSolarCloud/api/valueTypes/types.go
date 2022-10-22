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
	TypeArrayString    = "[]String"
	TypeArrayTime      = "[]Time"
	TypeArrayUnitValue = "[]UnitValue"
	TypeUnitValues     = "UnitValues"

	TypeArrayValueTypesBool      = "[]valueTypes.Bool"
	TypeArrayValueTypesCount     = "[]valueTypes.Count"
	TypeArrayValueTypesDateTime  = "[]valueTypes.DateTime"
	TypeArrayValueTypesFloat     = "[]valueTypes.Float"
	TypeArrayValueTypesInteger   = "[]valueTypes.Integer"
	TypeArrayValueTypesPointId   = "[]valueTypes.PointId"
	TypeArrayValueTypesPsKey     = "[]valueTypes.PsKey"
	TypeArrayValueTypesString    = "[]valueTypes.String"
	TypeArrayValueTypesTime      = "[]valueTypes.Time"
	TypeArrayValueTypesUnitValue = "[]valueTypes.UnitValue"

	TypeValueTypesBool      = "valueTypes.Bool"
	TypeValueTypesCount     = "valueTypes.Count"
	TypeValueTypesDateTime  = "valueTypes.DateTime"
	TypeValueTypesFloat     = "valueTypes.Float"
	TypeValueTypesInteger   = "valueTypes.Integer"
	TypeValueTypesPointId   = "valueTypes.PointId"
	TypeValueTypesPsKey     = "valueTypes.PsKey"
	TypeValueTypesString    = "valueTypes.String"
	TypeValueTypesTime      = "valueTypes.Time"
	TypeValueTypesUnitValue = "valueTypes.UnitValue"
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
			switch fieldVo.Type().String() {
				case TypeBool:
				case TypeCount:
				case TypeDateTime:
				case TypeFloat:
				case TypeInteger:
				case TypePointId:
				case TypePsKey:
				case TypeString:
				case TypeTime:
				case TypeUnitValue:
				case TypeValueTypesBool:
				case TypeValueTypesCount:
				case TypeValueTypesDateTime:
				case TypeValueTypesFloat:
				case TypeValueTypesInteger:
				case TypeValueTypesPointId:
				case TypeValueTypesPsKey:
				case TypeValueTypesString:
				case TypeValueTypesTime:
				case TypeValueTypesUnitValue:

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
		switch reflect.TypeOf(i).String() {
			case "int":
				val = int64(i.(int))
			case "int32":
				val = int64(i.(int32))
			case "int64":
				val = i.(int64)

			case TypeValueTypesInteger:
				fallthrough
			case TypeInteger:
				val = i.(Integer).Value()

			case "valueTypes.Count":
				fallthrough
			case TypeCount:
				val = i.(Integer).Value()
		}

		if s == 0 {
			ret = fmt.Sprintf("%d", val)
			break
		}

		ret = fmt.Sprintf("%." + strconv.Itoa(s) + "d", val)
	}
	return ret
}

func SizeOfInt(i interface{}) int {
	var ret int
	for range Only.Once {
		var val int64
		switch reflect.TypeOf(i).String() {
			case "int":
				val = int64(i.(int))
			case "int32":
				val = int64(i.(int32))
			case "int64":
				val = i.(int64)

			case TypeValueTypesInteger:
				fallthrough
			case TypeInteger:
				val = i.(Integer).Value()

			case TypeValueTypesCount:
				fallthrough
			case TypeCount:
				val = i.(Integer).Value()
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

func TypeToString(intSize int, dateFormat string, e interface{}) string {
	var ret string
	for range Only.Once {
		if IsNil(e) {
			fmt.Println("NIL")
			ret = "NIL"
			break
		}

		if dateFormat == "" {
			dateFormat = DateTimeAltLayout
		}

		// fmt.Printf("DEBUG: %s\n", reflect.TypeOf(e).String())
		switch reflect.TypeOf(e).String() {
			case "int":
				ret = PrintInt(intSize, e.(int))
			case "int32":
				ret = PrintInt(intSize, e.(int32))
			case "int64":
				ret = PrintInt(intSize, e.(int64))
			case "float32":
				// ret = float64(s.(float32))
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			case "float64":
				// ret = s.(float64)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			case "string":
				ret = strings.Trim(e.(string), ".")
			case "bool":
				ret = fmt.Sprintf("%v", e.(bool))
			case "[]string":
				// v := strings.Join(s.([]string), ",")
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesUnitValue:
				fallthrough
			case TypeUnitValue:
				ret = e.(UnitValue).String()

			case TypeUnitValues:
				fallthrough
			case TypeArrayUnitValue:
				fallthrough
			case TypeArrayValueTypesUnitValue:
				// ret = s.([]UnitValue)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesFloat:
				fallthrough
			case TypeFloat:
				// ret = s.(Float)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeArrayValueTypesFloat:
				fallthrough
			case TypeArrayFloat:
				// ret = s.([]Float)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesInteger:
				fallthrough
			case TypeInteger:
				ret = PrintInt(intSize, e.(Integer))

			case TypeArrayValueTypesInteger:
				fallthrough
			case TypeArrayInteger:
				// ret = s.([]Integer)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesCount:
				fallthrough
			case TypeCount:
				ret = PrintInt(intSize, e.(Count))

			case TypeArrayValueTypesCount:
				fallthrough
			case TypeArrayCount:
				// ret = s.([]Count)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesBool:
				fallthrough
			case TypeBool:
				ret = e.(Bool).String()

			case TypeArrayValueTypesBool:
				fallthrough
			case TypeArrayBool:
				// ret = s.([]Bool)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesString:
				fallthrough
			case TypeString:
				ret = strings.Trim(e.(String).String(), ".")

			case TypeArrayValueTypesString:
				fallthrough
			case TypeArrayString:
				// ret = s.([]String)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesPsKey:
				fallthrough
			case TypePsKey:
				ret = e.(PsKey).Value()

			case TypeArrayValueTypesPsKey:
				fallthrough
			case TypeArrayPsKey:
				// ret = s.([]PsKey)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesPointId:
				fallthrough
			case TypePointId:
				ret = e.(PointId).String()

			case TypeArrayValueTypesPointId:
				fallthrough
			case TypeArrayPointId:
				// ret = s.([]PointId)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesDateTime:
				fallthrough
			case TypeDateTime:
				ret = e.(DateTime).Format(dateFormat)

			case TypeArrayValueTypesDateTime:
				fallthrough
			case TypeArrayDateTime:
				// ret = s.([]DateTime)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			}
	}
	return ret
}

func AnyToUnitValue(e interface{}, unit string, Type string, dateFormat string) (UnitValues, bool, bool) {
	var uv UnitValues
	ok := true
	isNil := false
	for range Only.Once {
		if IsNil(e) {
			// fmt.Println("DEBUG: AnyToUnitValue(): NIL")
			uv = append(uv, SetUnitValueString("", unit, Type + "(unknown)"))
			isNil = true
			break
		}
		// fmt.Printf("DEBUG: AnyToUnitValue(): %s\n", reflect.TypeOf(e).String())

		if dateFormat == "" {
			dateFormat = DateTimeAltLayout
		}

		switch reflect.TypeOf(e).String() {
			case "int":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueInteger(int64(e.(int)), unit, Type))
			case "int32":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueInteger(int64(e.(int32)), unit, Type))
			case "int64":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueInteger(e.(int64), unit, Type))
			case "float32":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueFloat(float64(e.(float32)), unit, Type))
			case "float64":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueFloat(e.(float64), unit, Type))
			case "string":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueString(e.(string), unit, Type))
			case "[]string":
				// v := strings.Join(e.([]string), ",")
				if Type == "" {
					Type = "--"
				}
				j, err := json.Marshal(e.([]string))
				if err != nil {
					j = []byte(fmt.Sprintf("%v", e.([]string)))
				}
				uv = append(uv, SetUnitValueString(string(j), unit, Type))
			case "bool":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, SetUnitValueBool(e.(bool)))

			case TypeValueTypesUnitValue:
				fallthrough
			case TypeUnitValue:
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, e.(UnitValue))
				// uv = uv.UnitValueFix()

			case TypeUnitValues:
				fallthrough
			case TypeArrayValueTypesUnitValue:
				fallthrough
			case TypeArrayUnitValue:
				for _, val := range e.([]UnitValue) {
					uv = append(uv, val)
				}

			case TypeValueTypesFloat:
				fallthrough
			case TypeFloat:
				if Type == "" {
					Type = TypeFloat
				}
				v := e.(Float)
				uv = append(uv, SetUnitValueFloat(v.Value(), unit, Type))

			case TypeArrayValueTypesFloat:
				fallthrough
			case TypeArrayFloat:
				if Type == "" {
					Type = TypeFloat
				}
				v := e.([]Float)
				for _, val := range v {
					uv = append(uv, SetUnitValueFloat(val.Value(), unit, Type))
				}

			case TypeValueTypesInteger:
				fallthrough
			case TypeInteger:
				if Type == "" {
					Type = TypeInteger
				}
				v := e.(Integer).Value()
				uv = append(uv, SetUnitValueInteger(v, unit, Type))

			case TypeArrayValueTypesInteger:
				fallthrough
			case TypeArrayInteger:
				if Type == "" {
					Type = TypeInteger
				}
				v := e.([]Integer)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, Type))
				}
				// HERE IS THE PROBLEM - need to return SOMETHING, even if it's null!

			case TypeValueTypesCount:
				fallthrough
			case TypeCount:
				if Type == "" {
					Type = TypeCount
				}
				v := e.(Count).Value()
				uv = append(uv, SetUnitValueInteger(v, unit, Type))

			case TypeArrayValueTypesCount:
				fallthrough
			case TypeArrayCount:
				if Type == "" {
					Type = TypeCount
				}
				v := e.([]Count)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, Type))
				}

			case TypeValueTypesBool:
				fallthrough
			case TypeBool:
				if Type == "" {
					Type = TypeBool
				}
				v := e.(Bool)
				uv = append(uv, SetUnitValueBool(v.Value()))

			case TypeArrayValueTypesBool:
				fallthrough
			case TypeArrayBool:
				if Type == "" {
					Type = TypeBool
				}
				v := e.([]Bool)
				for _, val := range v {
					// uv = append(uv, SetUnitValueString(val.String(), unit, Type))
					uv = append(uv, SetUnitValueBool(val.Value()))
				}

			case TypeValueTypesString:
				fallthrough
			case TypeString:
				if Type == "" {
					Type = TypeString
				}
				v := e.(String).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case TypeArrayValueTypesString:
				fallthrough
			case TypeArrayString:
				if Type == "" {
					Type = TypeString
				}
				v := e.([]String)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, Type))
				}

			case TypeValueTypesPsKey:
				fallthrough
			case TypePsKey:
				if Type == "" {
					Type = TypePsKey
				}
				v := e.(PsKey).Value()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case TypeArrayValueTypesPsKey:
				fallthrough
			case TypeArrayPsKey:
				if Type == "" {
					Type = TypePsKey
				}
				v := e.([]PsKey)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, Type))
				}

			case TypeValueTypesPointId:
				fallthrough
			case TypePointId:
				if Type == "" {
					Type = TypePointId
				}
				v := e.(PointId).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case TypeArrayValueTypesPointId:
				fallthrough
			case TypeArrayPointId:
				if Type == "" {
					Type = TypePointId
				}
				v := e.([]PointId)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, Type))
				}

			case TypeValueTypesDateTime:
				fallthrough
			case TypeDateTime:
				if Type == "" {
					Type = TypeDateTime
				}
				v := e.(DateTime).Format(dateFormat)
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case TypeArrayValueTypesDateTime:
				fallthrough
			case TypeArrayDateTime:
				if Type == "" {
					Type = TypeDateTime
				}
				v := e.([]DateTime)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Format(dateFormat), unit, Type))
				}

			case TypeValueTypesTime:
				fallthrough
			case TypeTime:
				if Type == "" {
					Type = TypeTime
				}
				v := e.(Time).Format(TimeLayout)
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case TypeArrayValueTypesTime:
				fallthrough
			case TypeArrayTime:
				if Type == "" {
					Type = TypeTime
				}
				v := e.([]Time)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Format(TimeLayout), unit, Type))
				}

			default:
				uv = append(uv, SetUnitValueString("", unit, Type + "(unknown)"))
				ok = false
		}
	}
	return uv, isNil, ok
}

func AnyToValueString(e interface{}, intSize int, dateFormat string) string {
	var ret string

	for range Only.Once {
		if IsNil(e) {
			break
		}

		if dateFormat == "" {
			dateFormat = DateTimeAltLayout
		}

		// fmt.Printf("DEBUG TYPE: %s\n", reflect.TypeOf(e).String())
		switch reflect.TypeOf(e).String() {
			case "bool":
				ret = fmt.Sprintf("%v", e.(bool))
			case "int":
				ret = PrintInt(intSize, e.(int))
			case "int32":
				ret = PrintInt(intSize, e.(int32))
			case "int64":
				ret = PrintInt(intSize, e.(int64))
			case "float32":
				// ret = float64(s.(float32))
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			case "float64":
				// ret = s.(float64)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			case "string":
				ret = e.(string)
				// ret = strings.Trim(e.(string), ".")
			case "[]string":
				// v := strings.Join(s.([]string), ",")
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesUnitValue:
				fallthrough
			case TypeUnitValue:
				ret = e.(UnitValue).String()

			case TypeUnitValues:
				fallthrough
			case TypeArrayUnitValue:
				fallthrough
			case TypeArrayValueTypesUnitValue:
				// ret = s.([]UnitValue)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesFloat:
				fallthrough
			case TypeFloat:
				ret = e.(Float).String()

			case TypeArrayValueTypesFloat:
				fallthrough
			case TypeArrayFloat:
				// ret = s.([]Float)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesInteger:
				fallthrough
			case TypeInteger:
				ret = PrintInt(intSize, e.(Integer))

			case TypeArrayValueTypesInteger:
				fallthrough
			case TypeArrayInteger:
				// ret = s.([]Integer)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesCount:
				fallthrough
			case TypeCount:
				ret = PrintInt(intSize, e.(Count))

			case TypeArrayValueTypesCount:
				fallthrough
			case TypeArrayCount:
				// ret = s.([]Count)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesBool:
				fallthrough
			case TypeBool:
				ret = e.(Bool).String()

			case TypeArrayValueTypesBool:
				fallthrough
			case TypeArrayBool:
				// ret = s.([]Bool)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesString:
				fallthrough
			case TypeString:
				ret = e.(String).String()

			case TypeArrayValueTypesString:
				fallthrough
			case TypeArrayString:
				// ret = s.([]String)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesPsKey:
				fallthrough
			case TypePsKey:
				ret = e.(PsKey).Value()

			case TypeArrayValueTypesPsKey:
				fallthrough
			case TypeArrayPsKey:
				// ret = s.([]PsKey)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesPointId:
				fallthrough
			case TypePointId:
				ret = e.(PointId).String()

			case TypeArrayValueTypesPointId:
				fallthrough
			case TypeArrayPointId:
				// ret = s.([]PointId)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesDateTime:
				fallthrough
			case TypeDateTime:
				v := e.(DateTime)
				if v.IsZero() {
					ret = ""
					break
				}
				ret = v.Format(dateFormat)

			case TypeArrayValueTypesDateTime:
				fallthrough
			case TypeArrayDateTime:
				// ret = s.([]DateTime)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case TypeValueTypesTime:
				fallthrough
			case TypeTime:
				v := e.(Time)
				if v.IsZero() {
					ret = ""
					break
				}
				ret = v.Format(dateFormat)

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
