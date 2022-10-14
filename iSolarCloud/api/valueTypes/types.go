package valueTypes

import (
	"GoSungrow/Only"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


func IsUnknownStruct(fieldTo reflect.StructField, fieldVo reflect.Value) bool {
	var ok bool

	for range Only.Once {
		if fieldVo.Kind() == reflect.Struct {
			switch fieldTo.Type.String() {
				case "Bool":
				case "Count":
				case "DateTime":
				case "Float":
				case "Integer":
				case "PointId":
				case "PsKey":
				case "String":
				case "Time":
				case "UnitValue":
				case "valueTypes.Bool":
				case "valueTypes.Count":
				case "valueTypes.DateTime":
				case "valueTypes.Float":
				case "valueTypes.Integer":
				case "valueTypes.PointId":
				case "valueTypes.PsKey":
				case "valueTypes.String":
				case "valueTypes.Time":
				case "valueTypes.UnitValue":

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

			case "valueTypes.Integer":
				fallthrough
			case "Integer":
				val = i.(Integer).Value()

			case "valueTypes.Count":
				fallthrough
			case "Count":
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

			case "valueTypes.Integer":
				fallthrough
			case "Integer":
				val = i.(Integer).Value()

			case "valueTypes.Count":
				fallthrough
			case "Count":
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

			case "valueTypes.UnitValue":
				fallthrough
			case "UnitValue":
				ret = e.(UnitValue).String()

			case "UnitValues":
				fallthrough
			case "[]UnitValue":
				fallthrough
			case "[]valueTypes.UnitValue":
				// ret = s.([]UnitValue)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.Float":
				fallthrough
			case "Float":
				// ret = s.(Float)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "[]valueTypes.Float":
				fallthrough
			case "[]Float":
				// ret = s.([]Float)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.Integer":
				fallthrough
			case "Integer":
				ret = PrintInt(intSize, e.(Integer))

			case "[]valueTypes.Integer":
				fallthrough
			case "[]Integer":
				// ret = s.([]Integer)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.Count":
				fallthrough
			case "Count":
				ret = PrintInt(intSize, e.(Count))

			case "[]valueTypes.Count":
				fallthrough
			case "[]Count":
				// ret = s.([]Count)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.Bool":
				fallthrough
			case "Bool":
				ret = e.(Bool).String()

			case "[]valueTypes.Bool":
				fallthrough
			case "[]Bool":
				// ret = s.([]Bool)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.String":
				fallthrough
			case "String":
				ret = strings.Trim(e.(String).String(), ".")

			case "[]valueTypes.String":
				fallthrough
			case "[]String":
				// ret = s.([]String)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.PsKey":
				fallthrough
			case "PsKey":
				ret = e.(PsKey).Value()

			case "[]valueTypes.PsKey":
				fallthrough
			case "[]PsKey":
				// ret = s.([]PsKey)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.PointId":
				fallthrough
			case "PointId":
				ret = e.(PointId).String()

			case "[]valueTypes.PointId":
				fallthrough
			case "[]PointId":
				// ret = s.([]PointId)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "valueTypes.DateTime":
				fallthrough
			case "DateTime":
				ret = e.(DateTime).Format(dateFormat)

			case "[]valueTypes.DateTime":
				fallthrough
			case "[]DateTime":
				// ret = s.([]DateTime)
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)

			case "[]string":
				// v := strings.Join(s.([]string), ",")
				v , err := json.Marshal(e)
				if err != nil {
					break
				}
				ret = string(v)
			}
	}
	return ret
}

func AnyToUnitValue(e interface{}, unit string, Type string) (UnitValues, bool, bool) {
	var uv UnitValues
	ok := true
	isNil := false
	for range Only.Once {
		if IsNil(e) {
			// fmt.Println("DEBUG: AnyToUnitValue(): NIL")
			uv = append(uv, SetUnitValueString("", unit, Type + "(nil)"))
			isNil = true
			break
		}
		// fmt.Printf("DEBUG: AnyToUnitValue(): %s\n", reflect.TypeOf(e).String())

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

			case "valueTypes.UnitValue":
				fallthrough
			case "UnitValue":
				if Type == "" {
					Type = "--"
				}
				uv = append(uv, e.(UnitValue))
				// uv = uv.UnitValueFix()

			case "UnitValues":
				fallthrough
			case "[]valueTypes.UnitValue":
				fallthrough
			case "[]UnitValue":
				for _, val := range e.([]UnitValue) {
					uv = append(uv, val)
				}

			case "valueTypes.Float":
				fallthrough
			case "Float":
				if Type == "" {
					Type = "Float"
				}
				v := e.(Float)
				uv = append(uv, SetUnitValueFloat(v.Value(), unit, Type))

			case "[]valueTypes.Float":
				fallthrough
			case "[]Float":
				if Type == "" {
					Type = "Float"
				}
				v := e.([]Float)
				for _, val := range v {
					uv = append(uv, SetUnitValueFloat(val.Value(), unit, Type))
				}

			case "valueTypes.Integer":
				fallthrough
			case "Integer":
				if Type == "" {
					Type = "Integer"
				}
				v := e.(Integer).Value()
				uv = append(uv, SetUnitValueInteger(v, unit, Type))

			case "[]valueTypes.Integer":
				fallthrough
			case "[]Integer":
				if Type == "" {
					Type = "Integer"
				}
				v := e.([]Integer)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, Type))
				}

			case "valueTypes.Count":
				fallthrough
			case "Count":
				if Type == "" {
					Type = "Count"
				}
				v := e.(Count).Value()
				uv = append(uv, SetUnitValueInteger(v, unit, Type))

			case "[]valueTypes.Count":
				fallthrough
			case "[]Count":
				if Type == "" {
					Type = "Count"
				}
				v := e.([]Count)
				for _, val := range v {
					uv = append(uv, SetUnitValueInteger(val.Value(), unit, Type))
				}

			case "valueTypes.Bool":
				fallthrough
			case "Bool":
				if Type == "" {
					Type = "Bool"
				}
				v := e.(Bool)
				uv = append(uv, SetUnitValueBool(v.Value()))

			case "[]valueTypes.Bool":
				fallthrough
			case "[]Bool":
				if Type == "" {
					Type = "Bool"
				}
				v := e.([]Bool)
				for _, val := range v {
					// uv = append(uv, SetUnitValueString(val.String(), unit, Type))
					uv = append(uv, SetUnitValueBool(val.Value()))
				}

			case "valueTypes.String":
				fallthrough
			case "String":
				if Type == "" {
					Type = "String"
				}
				v := e.(String).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case "[]valueTypes.String":
				fallthrough
			case "[]String":
				if Type == "" {
					Type = "String"
				}
				v := e.([]String)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, Type))
				}

			case "valueTypes.PsKey":
				fallthrough
			case "PsKey":
				if Type == "" {
					Type = "PsKey"
				}
				v := e.(PsKey).Value()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case "[]valueTypes.PsKey":
				fallthrough
			case "[]PsKey":
				if Type == "" {
					Type = "PsKey"
				}
				v := e.([]PsKey)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.Value(), unit, Type))
				}

			case "valueTypes.PointId":
				fallthrough
			case "PointId":
				if Type == "" {
					Type = "PointId"
				}
				v := e.(PointId).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case "[]valueTypes.PointId":
				fallthrough
			case "[]PointId":
				if Type == "" {
					Type = "PointId"
				}
				v := e.([]PointId)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, Type))
				}

			case "valueTypes.DateTime":
				fallthrough
			case "DateTime":
				if Type == "" {
					Type = "DateTime"
				}
				v := e.(DateTime).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case "[]valueTypes.DateTime":
				fallthrough
			case "[]DateTime":
				if Type == "" {
					Type = "DateTime"
				}
				v := e.([]DateTime)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, Type))
				}

			case "valueTypes.Time":
				fallthrough
			case "Time":
				if Type == "" {
					Type = "Time"
				}
				v := e.(Time).String()
				uv = append(uv, SetUnitValueString(v, unit, Type))

			case "[]valueTypes.Time":
				fallthrough
			case "[]Time":
				if Type == "" {
					Type = "Time"
				}
				v := e.([]Time)
				for _, val := range v {
					uv = append(uv, SetUnitValueString(val.String(), unit, Type))
				}

			default:
				ok = false
		}
	}
	return uv, isNil, ok
}


// func Float32ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }

// func Float64ToString(num float64) string {
// 	s := fmt.Sprintf("%.6f", num)
// 	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
// }
