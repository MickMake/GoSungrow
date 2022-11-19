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
		// fmt.Printf("DEBUG:    K:%s / T:%v\n", kindy.String(), fieldVo)
		if kindy == reflect.Interface {
			ok = false
			break
		}

		if kindy == reflect.Slice {
			if fieldVo.Len() > 0 {
				fieldVo = reflect.ValueOf(fieldVo.Index(0).Interface())
				ok = IsTypeUnknown(fieldVo)
			}
			// ok = true
			break
		}

		if kindy == reflect.Array {
			if fieldVo.Len() > 0 {
				fieldVo = reflect.ValueOf(fieldVo.Index(0).Interface())
				ok = IsTypeUnknown(fieldVo)
			}
			// ok = true
			break
		}

		if kindy == reflect.Map {
			mk := fieldVo.MapKeys()
			if len(mk) > 0 {
				// ok = IsUnknownStruct(fieldVo.MapIndex(mk[0]).Interface())
				fieldVo = reflect.ValueOf(fieldVo.MapIndex(mk[0]).Interface())
				ok = IsTypeUnknown(fieldVo)
			}
			break
		}

		if kindy == reflect.Struct {
			ok = IsTypeUnknown(fieldVo)
		}
	}

	return ok
}

func IsKnownStruct(ref interface{}) bool {
	return !IsUnknownStruct(ref)
}

func IsTypeUnknown(fieldVo reflect.Value) bool {
	var ok bool

	for range Only.Once {
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

func SizeOfArrayLength(ref interface{}) int {
	ValueOf := reflect.ValueOf(ref)
	Len := ValueOf.Len()
	return SizeOfInt(Len)
}

func AnyToUnitValue(ref interface{}, key string, unit string, typeString string, dateFormat string) (UnitValues, bool, bool) {
	var uvs UnitValues
	ok := true
	isNil := false
	for range Only.Once {
		if IsNil(ref) {
			// fmt.Println("DEBUG: AnyToUnitValue(): NIL")
			uvs.AddString("", key, unit, typeString + "(unknown)")
			isNil = true
			break
		}
		// fmt.Printf("DEBUG: AnyToUnitValue(): %s\n", reflect.TypeOf(e).String())

		if dateFormat == "" {
			dateFormat = DateTimeAltLayout
		}

		ValueOf := reflect.ValueOf(ref)
		Kind := ValueOf.Kind()
		if IsKnownStruct(ref) {
			if Kind == reflect.Map {
				for _, mk := range ValueOf.MapKeys() {
					var uv UnitValues
					// fmt.Printf("Map[%s]\n", key)
					val := ValueOf.MapIndex(mk).Interface()
					uv, isNil, ok = AnyToUnitValue(val, mk.String(), unit, typeString, dateFormat) // uvs.AddUnitValue(ref.(UnitValue))
					uvs.Append(uv)
				}
				break
			}

			if Kind == reflect.Slice || Kind == reflect.Array {
				for index := 0; index < ValueOf.Len(); index++ {
					var uv UnitValues
					// fmt.Printf("Slice[%d]\n", index)
					val := ValueOf.Index(index).Interface()
					uv, isNil, ok = AnyToUnitValue(val, key, unit, typeString, dateFormat)
					uvs.Append(uv)
				}
				break
			}
		}

		Type := reflect.TypeOf(ref).String()
		Type = strings.ReplaceAll(Type, "valueTypes.", "")
		if typeString == "" {
			typeString = Type
		}

		switch Type {
			case "int":
				uvs.AddInteger(int64(ref.(int)), key, unit, typeString)
			case "int32":
				uvs.AddInteger(int64(ref.(int32)), key, unit, typeString)
			case "int64":
				uvs.AddInteger(ref.(int64), key, unit, typeString)
			case "float32":
				uvs.AddFloat(float64(ref.(float32)), key, unit, typeString)
			case "float64":
				uvs.AddFloat(ref.(float64), key, unit, typeString)
			case "string":
				uvs.AddString(ref.(string), key, unit, typeString)
			case "[]string":
				// j, err := json.Marshal(ref.([]string))
				// if err != nil {
				// 	j = []byte(fmt.Sprintf("%v", ref.([]string)))
				// }
				uvs.AddStrings(ref.([]string), key, unit, typeString)
			case "bool":
				uvs.AddBool(ref.(bool), key)

			case TypeUnitValue:
				val := ref.(UnitValue)
				val.SetKey(key)
				uvs.AddUnitValue(val)

			case TypeUnitValues:
				fallthrough
			case TypeArrayUnitValue:
				vals := ref.(UnitValues)
				for _, val := range vals.values {
					val.SetKey(key)
				}
				uvs.Append(vals)

			case TypeFloat:
				v := ref.(Float)
				uvs.AddFloat(v.Value(), key, unit, typeString)

			case TypeArrayFloat:
				v := ref.([]Float)
				for _, val := range v {
					uvs.AddFloat(val.Value(), key, unit, typeString)
				}

			case TypeInteger:
				v := ref.(Integer)
				uvs.AddInteger(v.Value(), key, unit, typeString)

			case TypeArrayInteger:
				v := ref.([]Integer)
				for _, val := range v {
					uvs.AddInteger(val.Value(), key, unit, typeString)
				}
				// HERE IS THE PROBLEM - need to return SOMETHING, even if it's null!

			case TypeCount:
				v := ref.(Count)
				uvs.AddInteger(v.Value(), key, unit, typeString)

			case TypeArrayCount:
				v := ref.([]Count)
				for _, val := range v {
					uvs.AddInteger(val.Value(), key, unit, typeString)
				}

			case TypeBool:
				v := ref.(Bool)
				uvs.AddBool(v.Value(), key)

			case TypeArrayBool:
				v := ref.([]Bool)
				for _, val := range v {
					uvs.AddBool(val.Value(), key)
				}

			case TypeString:
				v := ref.(String)
				uvs.AddString(v.String(), key, unit, typeString)

			case TypeArrayString:
				v := ref.([]String)
				for _, val := range v {
					uvs.AddString(val.String(), key, unit, typeString)
				}

			case TypePsId:
				v := ref.(PsId)
				uvs.AddString(v.String(), key, unit, typeString)

			case TypeArrayPsId:
				v := ref.([]PsId)
				for _, val := range v {
					uvs.AddString(val.String(), key, unit, typeString)
				}

			case TypePsKey:
				v := ref.(PsKey)
				uvs.AddString(v.String(), key, unit, typeString)

			case TypeArrayPsKey:
				v := ref.([]PsKey)
				for _, val := range v {
					uvs.AddString(val.String(), key, unit, typeString)
				}

			case TypePointId:
				v := ref.(PointId)
				uvs.AddString(v.String(), key, unit, typeString)

			case TypeArrayPointId:
				v := ref.([]PointId)
				for _, val := range v {
					uvs.AddString(val.String(), key, unit, typeString)
				}

			case TypeDateTime:
				v := ref.(DateTime)
				uvs.AddString(v.Format(dateFormat), key, unit, typeString)

			case TypeArrayDateTime:
				v := ref.([]DateTime)
				for _, val := range v {
					uvs.AddString(val.Format(dateFormat), key, unit, typeString)
				}

			case TypeTime:
				v := ref.(Time)
				uvs.AddString(v.Format(TimeLayout), key, unit, typeString)

			case TypeArrayTime:
				v := ref.([]Time)
				for _, val := range v {
					uvs.AddString(val.Format(TimeLayout), key, unit, typeString)
				}

			default:
				typeString = ""
				j, err := json.Marshal(ref)
				if err != nil {
					j = []byte(Type + "(unknown)")
				}
				uvs.AddString(string(j), key, unit, typeString)
				ok = false
		}
	}
	return uvs, isNil, ok
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
				v , err := json.Marshal(ref)
				if err != nil {
					break
				}
				ret = string(v)
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
