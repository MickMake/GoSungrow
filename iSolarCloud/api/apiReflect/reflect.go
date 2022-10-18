package apiReflect

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)


type DataStructures struct {
	Map map[string]DataStructure
}
type DataStructure struct {
	// PointType      string
	Json               string
	PointId            string
	PointParentId      string
	PointUnit          string
	PointUnitFrom      string
	PointTimestamp     time.Time
	PointName          string
	PointUpdateFreq    string
	PointValueType     string
	PointAliasTo       string
	PointIgnore        bool
	PointGroupName     string
	PointGroupNameFrom string
	PointTimestampFrom string

	Value              interface{}
	ValueType          string
	ValueKind          string
	Endpoint           string
}

const (
	PointId             = "PointId"             // Point id in the form p\d+ or \d+
	PointUnit           = "PointUnit"           // Units: Wh, kWh, C, h.
	PointUnitFrom       = "PointUnitFrom"       // Get PointUnit from another field structure.
	PointParentId       = "PointParentId"       // Associated parent of point.
	PointName           = "PointName"           // Human-readable name of point.
	PointUpdateFreq     = "PointUpdateFreq"     // Point update frequency - Total, Yearly, Monthly, Day.
	PointValueType      = "PointValueType"      // Value type of point: energy, date, battery, temperature.
	PointAliasTo        = "PointAliasTo"        // Alias this point to another point.
	PointAliasFrom      = "PointAliasFrom"      // Alias this point from another point.
	PointIgnore         = "PointIgnore"         // Ignore this point.
	PointGroupName      = "PointGroupName"      // Point group name.
	PointGroupNameFrom  = "PointGroupNameFrom"  // Get PointGroupName from another field structure.
	PointNameFromChild  = "PointNameFromChild"  // Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameFromParent = "PointNameFromParent" // Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameDateFormat = "PointNameDateFormat" // Date format when using PointNameFrom, (if the field is a time.Time type).
	PointArrayFlatten   = "PointArrayFlatten"   // Flatten an array into a string. EG: ["one", "two", "three"]
	PointTimestampFrom  = "PointTimestampFrom"  // Pull timestamp from another field structure.
)

func (dss *DataStructures) Add(name string, ds DataStructure)  {
	for range Only.Once {
		if dss.Map == nil {
			dss.Map = make(map[string]DataStructure)
		}
		// fmt.Printf("DEBUG DataStructures.Add() %s - Kind:'%s' Type:'%s'\n", name, ds.ValueKind, ds.ValueType)
		dss.Map[name] = ds
	}
}

func (dss *DataStructures) Exists(name string) bool {
	var yes bool
	for range Only.Once {
		if dss.Map == nil {
			break
		}
		if _, ok := dss.Map[name]; ok {
			yes = ok
			break
		}
	}
	return yes
}

func (dss *DataStructures) Append(dsm DataStructures)  {
	for range Only.Once {
		if dss.Map == nil {
			dss.Map = make(map[string]DataStructure)
		}
		// fmt.Printf("Map BEFORE:[%d] (adding %d)\n", len(dss.Map), len(dsm.Map))
		for name, ds := range dsm.Map {
			dss.Map[name] = ds
		}
		// fmt.Printf("Map AFTER:[%d]\n", len(dss.Map))
	}
}


func (dss *DataStructures) GetPointTags(parentRef interface{}, ref interface{}, name ...string) DataStructures {
	// var ret DataStructures
	// ret.Map = make(map[string]DataStructure)

	for range Only.Once {
		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			fieldVo := vo.Field(i)

			if !fieldTo.IsExported() {
				fmt.Printf("DEBUG: NOTEXPORTED(%s): Type %s\n", name, fieldTo.Name)
				continue
			}

			pointJson := fieldTo.Tag.Get("json")
			pointId := fieldTo.Tag.Get(PointId)
			if pointId == "" {
				pointId = pointJson
			}

			// fmt.Printf("DEBUG[%d]: %s.%s: Key[%s]:\tKind:'%s' Type:'%s'\n",
			// 	len(dss.Map),
			// 	strings.Join(name, "."), pointId,
			// 	fieldTo.Name,
			// 	fieldVo.Kind(),
			// 	fieldTo.Type.String(),
			// )
			// if strings.Contains(pointId, "actual_energy") || strings.Contains(pointJson, "actual_energy") {
			// 	fmt.Printf("F:%v\n", pointId)
			// 	fmt.Println("")
			// }

			switch fieldVo.Kind() {
				case reflect.Uintptr:
					fallthrough
				case reflect.Complex64:
					fallthrough
				case reflect.Complex128:
					fallthrough
				case reflect.Chan:
					fallthrough
				case reflect.Func:
					fallthrough
				case reflect.Pointer:
					fallthrough
				case reflect.UnsafePointer:
					fmt.Printf("Unsupported type: '%s.%s' (%s)\n", name, pointId, fieldVo.Type().String())
					continue

				case reflect.Slice:
					// Handle slices here.
					// Adds more Point* tags - PointNameFromChild, PointNameDateFormat
					// Replicates the JoinWithDots 1st and 2nd arguments.
					// intSize int, dateFormat string
					// fmt.Printf("Kind: %s ##########################################\n", fieldVo.Kind().String())
					pointNameFromChild := fieldTo.Tag.Get(PointNameFromChild)
					pointNameFromParent := fieldTo.Tag.Get(PointNameFromParent)
					pointNameDateFormat := fieldTo.Tag.Get(PointNameDateFormat)
					intSize := valueTypes.SizeOfInt(fieldVo.Len())
					pointArrayFlatten := fieldTo.Tag.Get(PointArrayFlatten)
					n2 := append(name, pointId)

					if pointArrayFlatten != "" {
						// We want to flatten a slice down to EG "[1, 2, 3]"
						endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, name)
						ds.Value = valueTypes.AnyToValueString(fieldVo.Interface(), 0, "")
						dss.Add(endPointName + "." + ds.PointId, ds)
						continue
					}

					if valueTypes.IsUnknownStruct(fieldVo.Interface()) {
						for si := 0; si < fieldVo.Len(); si++ {
							// Are we using an index number for name or field key value?
							pn := strconv.Itoa(si)
							n3 := append(n2, pn)
							if pointNameFromChild != "" {
								// PointNameFromChild - In this case points to a field within a CHILD struct.
								pn = GetPointNameFrom(fieldVo.Index(si).Interface(), pointNameFromChild, intSize, pointNameDateFormat)
								n3 = append(n2[:len(n2) - 1], pn)
							}
							if pointNameFromParent != "" {
								// PointNameFromChild - In this case points to a field within a CHILD struct.
								pn = GetPointNameFrom(fieldVo.Interface(), pointNameFromParent, intSize, pointNameDateFormat)
								n3 = append(n2[:len(n2) - 1], pn)
							}

							dss.GetPointTags(fieldVo.Index(si).Interface(), fieldVo.Index(si).Interface(), n3...)
						}
						continue
					}

					// Flatten slice for []Integer / []Float objects.
					endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, name)
					ds.Value = valueTypes.AnyToValueString(fieldVo.Interface(), 0, "")
					dss.Add(endPointName + "." + ds.PointId, ds)

					// This commented-out section can handle []Integer fields.
					// for si := 0; si < fieldVo.Len(); si++ {
					// 	// Are we using an index number for name or field key value?
					// 	pn := strconv.Itoa(si)
					// 	n3 := append(n2, pn)
					// 	if pointNameFromChild != "" {
					// 		// PointNameFromChild - In this case points to a field within a CHILD struct.
					// 		pn = GetPointNameFrom(fieldVo.Index(si).Interface(), pointNameFromChild, intSize, pointNameDateFormat)
					// 		n3 = append(n2[:len(n2) - 1], pn)
					// 	}
					// 	if pointNameFromParent != "" {
					// 		// PointNameFromChild - In this case points to a field within a CHILD struct.
					// 		pn = GetPointNameFrom(fieldVo.Interface(), pointNameFromParent, intSize, pointNameDateFormat)
					// 		n3 = append(n2[:len(n2) - 1], pn)
					// 	}
					//
					// 	if valueTypes.IsUnknownStruct(fieldVo.Index(si).Interface()) {
					// 		dss.GetPointTags(fieldVo.Index(si).Interface(), fieldVo.Index(si).Interface(), n3...)
					// 		continue
					// 	}
					//
					// 	// // We want to flatten a slice down to EG "[1, 2, 3]"
					// 	// endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, n2)
					// 	// ds.Value = valueTypes.AnyToValueString(fieldVo.Interface(), 0, "")
					// 	// dss.Add(endPointName + "." + pn, ds)
					// 	endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, n3)
					// 	// ds.Json = pointId + pn
					// 	ds.PointId = pn
					// 	val := fieldVo.Index(si)
					// 	ds.Value = val.Interface()
					// 	dss.Add(endPointName, ds)
					// }
					continue

				case reflect.Array:
					// @TODO - Handle arrays here.
					// Adds more Point* tags - PointNameFromChild, PointNameDateFormat
					// Replicates the JoinWithDots 1st and 2nd arguments.
					// intSize int, dateFormat string
					// fmt.Printf("Kind: %s ##########################################\n", fieldVo.Kind().String())
					pointNameFromChild := fieldTo.Tag.Get(PointNameFromChild)
					pointNameFromParent := fieldTo.Tag.Get(PointNameFromParent)
					pointNameDateFormat := fieldTo.Tag.Get(PointNameDateFormat)
					intSize := valueTypes.SizeOfInt(fieldVo.Len())
					pointArrayFlatten := fieldTo.Tag.Get(PointArrayFlatten)
					n2 := append(name, pointId)

					if pointArrayFlatten != "" {
						// We want to flatten a slice down to EG "[1, 2, 3]"
						endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, name)
						ds.Value = valueTypes.AnyToValueString(fieldVo.Interface(), 0, "")
						dss.Add(endPointName + "." + ds.PointId, ds)
						continue
					}

					for si := 0; si < fieldVo.Len(); si++ {
						// Are we using an index number for name or field key value?
						pn := strconv.Itoa(si)
						n3 := append(n2, pn)
						if pointNameFromChild != "" {
							// PointNameFromChild - In this case points to a field within a CHILD struct.
							pn = GetPointNameFrom(fieldVo.Index(si).Interface(), pointNameFromChild, intSize, pointNameDateFormat)
							n3 = append(n2[:len(n2) - 1], pn)
						}
						if pointNameFromParent != "" {
							// PointNameFromChild - In this case points to a field within a CHILD struct.
							pn = GetPointNameFrom(fieldVo.Interface(), pointNameFromParent, intSize, pointNameDateFormat)
							n3 = append(n2[:len(n2) - 1], pn)
						}

						if valueTypes.IsUnknownStruct(fieldVo.Index(si).Interface()) {
							dss.GetPointTags(fieldVo.Index(si).Interface(), fieldVo.Index(si).Interface(), n3...)
							continue
						}

						// // We want to flatten a slice down to EG "[1, 2, 3]"
						// endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, n2)
						// ds.Value = valueTypes.AnyToValueString(fieldVo.Interface(), 0, "")
						// dss.Add(endPointName + "." + pn, ds)
						endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, n3)
						// ds.Json = pointId + pn
						ds.PointId = pointId + pn
						val := fieldVo.Index(si)
						ds.Value = val.Interface()
						dss.Add(endPointName, ds)
					}

					// @TODO - Can't quite figure out what to do here. So will just flatten the array.
					// for si := 0; si < fieldVo.Len(); si++ {
					// 	// Are we using an index number for name or field key value?
					// 	pn := strconv.Itoa(si)
					// 	n3 := append(name)	// , pointId + pn)
					// 	if pointNameFrom != "" {
					// 		pn = GetPointNameFromChild(parentRef, pointNameFrom, intSize, pointNameDateFormat)
					// 		n3 = append(n2[:len(n2) - 1], pn)
					// 	}
					//
					// 	endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, n3)
					// 	// ds.Json = pointId + pn
					// 	ds.PointId = pointId + pn
					// 	val := fieldVo.Index(si)
					// 	ds.Value = val.Interface()
					// 	dss.Add(endPointName, ds)
					// }
					continue

				case reflect.Map:
					// fmt.Printf("Kind: %s ##########################################\n", fieldVo.Kind().String())
					n2 := append(name, pointId)
					// pointNameFromChild := fieldTo.Tag.Get(PointNameFromChild)
					// pointNameFromParent := fieldTo.Tag.Get(PointNameFromParent)

					for _, key := range fieldVo.MapKeys() {
						// @TODO - Implement this.
						// if pointNameFromChild != "" {
						// 	// PointNameFromChild - In this case points to a field within a CHILD struct.
						// 	pn = GetPointNameFrom(fieldVo.Index(si).Interface(), pointNameFromChild, intSize, pointNameDateFormat)
						// 	n3 = append(n2[:len(n2) - 1], pn)
						// }
						// if pointNameFromParent != "" {
						// 	// PointNameFromChild - In this case points to a field within a CHILD struct.
						// 	pn = GetPointNameFrom(fieldVo.Interface(), pointNameFromParent, intSize, pointNameDateFormat)
						// 	n3 = append(n2[:len(n2) - 1], pn)
						// }

						endPointName, ds := makeDataStructure(fieldVo.Interface(), fieldTo, fieldVo, n2)
						ds.Json = key.String()
						ds.PointId = key.String()
						val := fieldVo.MapIndex(key)
						ds.Value = val.Interface()
						dss.Add(endPointName + "." + ds.PointId, ds)
					}
					continue

				case reflect.Struct:
					// fmt.Printf("Kind: %s ##########################################\n", fieldVo.Kind().String())
					if valueTypes.IsUnknownStruct(fieldVo.Interface()) {
						n2 := append(name, pointId)
						dss.GetPointTags(parentRef, fieldVo.Interface(), n2...)
						continue
					}

					// fmt.Printf("[%s.%s] => %v\n", strings.Join(name, "."), pointId, fieldVo.Interface())
					endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, name)
					ds.Value = fieldVo.Interface()
					dss.Add(endPointName + "." + ds.PointId, ds)
					continue
			}

			endPointName, ds := makeDataStructure(parentRef, fieldTo, fieldVo, name)
			ds.Value = fieldVo.Interface()
			dss.Add(endPointName + "." + ds.PointId, ds)
		}
	}

	return *dss
}

func GetUnitFrom(ref interface{}, pointUnitFrom string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Struct {
			break
		}

		// Iterate over all available fields, looking for the field name.
		for i := 0; i < vo.NumField(); i++ {
			if vo.Type().Field(i).Name != pointUnitFrom {
				continue
			}

			// fmt.Printf("GetUnitFrom: %v\n", fieldVo.Interface())
			ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
			break
		}
	}

	return ret
}

func GetGroupNameFrom(ref interface{}, pointGroupNameFrom string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Struct {
			break
		}

		// Iterate over all available fields, looking for the field name.
		for i := 0; i < vo.NumField(); i++ {
			if vo.Type().Field(i).Name != pointGroupNameFrom {
				continue
			}

			// fmt.Printf("GetGroupNameFrom: %v\n", fieldVo.Interface())
			ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
			break
		}
	}

	return ret
}

func GetTimestampFrom(ref interface{}, pointTimestampFrom string, dateFormat string) time.Time {
	var ret time.Time
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Struct {
			break
		}

		// Iterate over all available fields, looking for the field name.
		for i := 0; i < vo.NumField(); i++ {
			if vo.Type().Field(i).Name != pointTimestampFrom {
				continue
			}

			// fmt.Printf("GetTimestampFrom: %v\n", fieldVo.Interface())
			foo2 := fmt.Sprintf("%v", vo.Field(i).Interface())
			ret = valueTypes.SetDateTimeString(foo2).Time
			break
		}
	}

	return ret
}

func GetPointNameFrom(ref interface{}, pointNameFrom string, intSize int, dateFormat string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)
		if vo.Kind() != reflect.Struct {
			break
		}

		// Iterate over all available fields, looking for the field name.
		for i := 0; i < vo.NumField(); i++ {
			if vo.Type().Field(i).Name != pointNameFrom {
				continue
			}

			// fmt.Printf("GetPointNameFrom: %v\n", fieldVo.Interface())
			ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), intSize, dateFormat)
			break
		}
	}

	return ret
}

func makeDataStructure(parentRef interface{}, fieldTo reflect.StructField, fieldVo reflect.Value, name []string) (string, DataStructure) {
	var endpoint string
	var ds DataStructure

	for range Only.Once {
		if !fieldTo.IsExported() {
			fmt.Printf("DEBUG: NOTEXPORTED(%s): %s\n", fieldTo.Name, fieldTo.Tag.Get("json"))
			break
		}

		ignore := false
		if fieldTo.Tag.Get(PointIgnore) != "" {
			ignore = true
		}

		// if valueTypes.IsNil(ref) {
		// 	pointValueType = "NIL"
		// }

		pointJson := fieldTo.Tag.Get("json")
		pointId := fieldTo.Tag.Get(PointId)
		if pointId == "" {
			pointId = pointJson
		}

		pointValueType := fieldTo.Tag.Get(PointValueType)
		if pointValueType == "" {
			// pointValueType = fieldVo.Kind().String()
		}

		pointUnit := fieldTo.Tag.Get(PointUnit)
		pointUnitFrom := fieldTo.Tag.Get(PointUnitFrom)
		if pointUnitFrom != "" {
			pointUnit = GetUnitFrom(parentRef, pointUnitFrom)
		}

		pointGroupName := fieldTo.Tag.Get(PointGroupName)
		pointGroupNameFrom := fieldTo.Tag.Get(PointGroupNameFrom)
		if pointGroupNameFrom != "" {
			pointGroupName = GetGroupNameFrom(parentRef, pointGroupNameFrom)
		}

		pointTimestamp := time.Now()
		pointTimestampFrom := fieldTo.Tag.Get(PointTimestampFrom)
		if pointTimestampFrom != "" {
			pointTimestamp = GetTimestampFrom(parentRef, pointTimestampFrom, valueTypes.DateTimeLayout)
		}

		endpoint = strings.TrimPrefix(strings.Join(name, "."), ".")

		ds = DataStructure {
			// PointType:      fieldTo.Tag.Get(PointType),
			Json:               pointJson,
			PointId:            pointId,
			PointParentId:      fieldTo.Tag.Get(PointParentId),

			PointUnit:          pointUnit,
			PointUnitFrom:      pointUnitFrom,

			PointTimestamp:     pointTimestamp,
			PointTimestampFrom: pointTimestampFrom,

			PointGroupName:     pointGroupName,
			PointGroupNameFrom: pointGroupNameFrom,

			PointName:       fieldTo.Tag.Get(PointName),
			PointUpdateFreq: fieldTo.Tag.Get(PointUpdateFreq),
			PointValueType:  pointValueType,
			PointAliasTo:    fieldTo.Tag.Get(PointAliasTo),
			PointIgnore:     ignore,
			Value:           nil,
			ValueType:       fieldTo.Type.String(),
			ValueKind:       fieldVo.Kind().String(),
			Endpoint:        endpoint,
		}
	}

	return endpoint, ds
}

// func (dss *DataStructures) GetUnitFrom(ref ...string) string {
// 	var ret string
// 	for range Only.Once {
// 		r := strings.Join(ref, ".")
// 		if s, ok := dss.Map[r]; ok {
// 			ret = s.PointUnit
// 		}
// 	}
// 	return ret
// }
//
// func (dss *DataStructures) GetPointGroupNameFrom(ref ...string) string {
// 	var ret string
// 	for range Only.Once {
// 		r := strings.Join(ref, ".")
// 		if s, ok := dss.Map[r]; ok {
// 			ret = s.PointUnit
// 		}
// 	}
// 	return ret
// }


// GetArea Return an Area name if we are given an Area or EndPoint struct.
func GetArea(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret1 = strings.TrimPrefix(ret1, trim)
		ret2 := val.Type().Name()
		// ret3 := val.Type().String()
		// fmt.Printf("%s\t%s\t%s\n", ret1, ret2, "")

		if ret2 == "Area" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-1]
			break
		}

		if ret2 == "EndPoint" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-2]
			break
		}

		ret = ret1
	}
	return ret
}

// GetName Return an endpoint name if we are given an Area or EndPoint struct.
func GetName(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret1 = strings.TrimPrefix(ret1, trim)
		ret2 := val.Type().Name()
		// ret3 := val.Type().String()
		// fmt.Printf("%s\t%s\t%s\n", ret1, ret2, "")

		if ret2 == "Area" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-2]
			break
		}

		if ret2 == "EndPoint" {
			s := strings.Split(ret1, "/")
			ret = s[len(s)-1]
			break
		}

		ret = ret1
	}
	return ret
}

func GetType(v interface{}) string {
	return reflect.ValueOf(v).Type().Name()
	// var ret string
	// for range Only.Once {
	// 	if v == nil {
	// 		break
	// 	}
	//
	// 	val := reflect.ValueOf(v)
	// 	ret = val.Type().Name()
	// }
	// return ret
}

func GetPkgType(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
	// var ret string
	// for range Only.Once {
	// 	if v == nil {
	// 		break
	// 	}
	//
	// 	val := reflect.ValueOf(v)
	// 	// ret = val.Type().Name()
	// 	ret = val.Type().String()
	// }
	// return ret
}

func GetStructName(v interface{}) (string, string) {
	var area string
	var endpoint string
	for range Only.Once {
		val := reflect.ValueOf(v)
		// ret = val.Type().Name()		// Returns structure, (EndPoint name).
		// ret = val.Type().PkgPath()	// Returns structure path.
		// ret = val.Type().String()	// Returns

		// @TODO - Need to check for pointers to struct
		// 	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
		// 		ret = strings.ToLower(t.Elem().Name())
		// 	} else {
		// 		ret = strings.ToLower(t.Name())
		// 	}

		s := strings.Split(val.Type().String(), ".")
		if len(s) < 2 {
			break
		}
		area = s[0]
		endpoint = s[1]
	}
	return area, endpoint
}

func DoTypesMatch(a interface{}, b interface{}) error {
	var err error
	for range Only.Once {
		aName := GetType(a)
		bName := GetType(b)
		if aName == bName {
			break
		}
		err = errors.New(fmt.Sprintf("interface '%s' doesn't match '%s'", aName, bName))
	}
	return err
}

func DoPkgTypesMatch(a interface{}, b interface{}) error {
	var err error
	for range Only.Once {
		aName := GetPkgType(a)
		bName := GetPkgType(b)
		if aName == bName {
			break
		}
		err = errors.New(fmt.Sprintf("interface '%s' doesn't match '%s'", aName, bName))
	}
	return err
}

func PackageName(trim string, v interface{}) string {
	var ret string
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			ret = val.Elem().Type().PkgPath()
			break
		}

		ret2 := val.Type().Name()
		ret3 := val.Type().String()
		ret = val.Type().PkgPath()
		ret = strings.TrimPrefix(ret, trim)

		ret = fmt.Sprintf("%s\t%s\t%s\n", ret, ret2, ret3)
	}
	return ret
}

func Query(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	for id := 0; id < s.NumField(); id++ {
		value := fmt.Sprintf("%v", s.Field(id).Interface())
		if value == "" {
			continue
		}
		ret += fmt.Sprintf("&%s=%s",
			typeOf.Field(id).Tag.Get("json"),
			value,
		)
		// fmt.Printf("%d: %s %s = %v\n",
		//	i,
		//	typeOfT.Field(i).Name,
		//	s.Field(i).Type(),
		//	s.Field(i).Interface(),
		// )
	}

	return ret
}

func PrintHeader(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
		case "string":
			ret = fmt.Sprintf("%v", s)
		default:
			for id := 0; id < s.NumField(); id++ {
				// value := fmt.Sprintf("%v", s.Field(id).Interface())
				// if value == "" {
				//	continue
				// }
				ret += fmt.Sprintf("%s (%s),",
					typeOf.Field(id).Name,
					typeOf.Field(id).Tag.Get("json"),
				)
				// fmt.Printf("%d: %s %s = %v\n",
				//	i,
				//	typeOfT.Field(i).Name,
				//	s.Field(i).Type(),
				//	s.Field(i).Interface(),
				// )
			}
	}

	return ret
}

func PrintValue(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = fmt.Sprintf("%v", s)
	default:
		for id := 0; id < s.NumField(); id++ {
			ret += fmt.Sprintf("%v,", s.Field(id).Interface())
		}
	}

	return ret
}

func PrintAsJson(ref interface{}) (string, error) {
	var j string
	var err error

	for range Only.Once {
		var ret []byte
		ret, err = json.MarshalIndent(ref, "", "\t")
		if err != nil {
			break
		}

		j = string(ret)
	}

	return j, err
}

func HeaderAsArray(i interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, "Name")
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%s (%s)",
				typeOf.Field(id).Name,
				typeOf.Field(id).Tag.Get("json"),
			))
		}
	}

	return ret
}

func AsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, fmt.Sprintf("%v", s))
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
		}
	}

	return ret
}

//goland:noinspection GoUnusedFunction
func rowAsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	for id := 0; id < s.NumField(); id++ {
		ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
	}

	return ret
}

// var headerStyleTable = map[string]json2csv.KeyStyle{
//	"jsonpointer": json2csv.JSONPointerStyle,
//	"slash":       json2csv.SlashStyle,
//	"dot":         json2csv.DotNotationStyle,
//	"dot-bracket": json2csv.DotBracketStyle,
// }
//
// func PrintAsCsv(ref interface{}) (string, error) {
//	var c string
//	var err error
//
//	for range Once.Once {
//		var ret []byte
//		ret, err = json.Marshal(ref)
//		if err != nil {
//			break
//		}
//
//		var results []json2csv.KeyValue
//		results, err = json2csv.JSON2CSV(string(ret))
//		if err != nil {
//			log.Fatal(err)
//		}
//		if len(results) == 0 {
//			break
//		}
//
//		csv := json2csv.NewCSVWriter(os.Stdout)
//		// header style (jsonpointer, slash, dot, dot-bracket
//		csv.HeaderStyle = headerStyleTable["dot"]
//		csv.Transpose = true
//
//		err = csv.WriteCSV(results)
//		if err != nil {
//			break
//		}
//
//	}
//
//	return c, err
// }

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func ReflectAsJson(ref interface{}) string {
	var ret string

	for range Only.Once {
		switch reflect.TypeOf(ref).Kind() {
			case reflect.Slice:
			case reflect.Array:
				fmt.Println("The interface is a slice.")
				s := reflect.ValueOf(ref)
				ret += "["
				for i := 0; i < s.Len(); i++ {
					ret += ReflectAsJson(s.Index(i))
				}
				ret += "]"

			case reflect.Struct:
				s := reflect.ValueOf(ref) // .Elem()
				typeOf := s.Type()
				for i := 0; i < s.NumField(); i++ {
					value := fmt.Sprintf("%v", s.Field(i).Interface())
					if value == "" {
						continue
					}
					ret += fmt.Sprintf("%s:%s\n",
						typeOf.Field(i).Tag.Get("json"),
						value,
					)
					// fmt.Printf("%d: %s %s = %v\n",
					//	i,
					//	typeOfT.Field(i).Name,
					//	s.Field(i).Type(),
					//	s.Field(i).Interface(),
					// )
			}
		}
	}

	return ret
}

// FindInStruct Search a given structure for the State object and return its pointer.
//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func FindInStruct(ref interface{}, name string) interface{} {
	var ret interface{}

	for range Only.Once {
		v := reflect.ValueOf(ref)
		var e reflect.Value

		kind := v.Kind()
		// We're doing these checks to ensure ease of future expansion.
		if kind == reflect.Ptr {
			e = v.Elem()
			if e.Kind().String() == "ptr" {
				// PrintflnCyan("POINTER TO POINTER")	@TODO - DEBUG
				ret = FindInStruct(e.Addr().Elem().Interface(), name)
				break
			}
		} else if kind == reflect.Struct {
			// We can't handle a non-pointer, otherwise we get this...
			// reflect.flag.mustBeAssignable using unaddressable value
			e = v
			// Panic(PanicErrorNotGivenAPointer, v.String())
		} else {
			break
		}

		typeOfT := e.Type()
		for i := 0; i < e.NumField(); i++ {
			if typeOfT.Field(i).Name == name {
				ret = typeOfT.Field(i)
				break
			}

			// if typeOfT.Field(i).Name == name {
			//	state = e.Field(i).Interface().(*State)
			//	if state == nil {
			//		e.Field(i).Set(reflect.ValueOf(state))
			//	}
			//	break
			// }
		}
	}

	return ret
}

func GetNameOld(ref interface{}) (string, string) {
	var packageName string
	var structName string

	str := reflect.TypeOf(ref).String()
	str = strings.TrimPrefix(str, "*")
	str = strings.ToLower(str)
	sa := strings.SplitN(str, ".", 2)
	switch len(sa) {
		case 0:
		case 1:
			packageName = sa[0]
		case 2:
			packageName = sa[0]
			structName = sa[1]
	}
	return packageName, structName
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetPackageName(ref interface{}) string {
	str, _ := GetNameOld(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetStructName2(ref interface{}) string {
	str, _ := GetNameOld(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetCmdName2(ref interface{}) string {
	str := reflect.TypeOf(ref).PkgPath()
	str = filepath.Base(str)
	str = strings.ToLower(str)
	return str
}

// func GetStructName(ref interface{}) string {
// 	var ret string
//
// 	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
// 		ret = strings.ToLower(t.Elem().Name())
// 	} else {
// 		ret = strings.ToLower(t.Name())
// 	}
//
// 	return ret
// }

// VerifyOptionsRequired Verify fields within the structure that are required.
func VerifyOptionsRequired(ref interface{}) error {
	var err error

	for range Only.Once {
		// required := GetOptionsRequired(ref)

		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			required := fieldTo.Tag.Get("required")
			if required == "" {
				continue
			}

			fieldVo := vo.Field(i)
			value := fmt.Sprintf("%v", fieldVo.Interface())
			if value == "" {
				err = errors.New(fmt.Sprintf("option '%s' is empty", fieldTo.Name))
				break
			}
		}
	}

	return err
}

func HelpOptions(ref interface{}) string {
	var ret string

	for range Only.Once {
		t := reflect.TypeOf(ref)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			required := field.Tag.Get("required")
			if required == "" {
				ret += fmt.Sprintf("%s: optional\n", field.Name)
				continue
			}

			ret += fmt.Sprintf("%s: required\n", field.Name)
		}
	}

	return ret
}

func FindRequestData(ref interface{}) string {
	var ret string

	for range Only.Once {
		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			// required := fieldTo.Tag.GetByJson("required")
			fmt.Printf(">%s\t", fieldTo.Name)

			fieldVo := vo.Field(i)

			fmt.Printf(">%s\n", fieldVo.String())
			value := fmt.Sprintf("%v", fieldVo.Interface())
			if value == "" {
				break
			}
		}
	}

	return ret
}

func GetRequestString(ref interface{}) string {
	var ret string

	for range Only.Once {
		vo := reflect.ValueOf(ref)
		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldVo := vo.Field(i)
			ret += fmt.Sprintf("-%v", fieldVo.Interface())
		}
	}

	return ret
}

func GetFingerprint(ref interface{}) string {
	var ret string

	for range Only.Once {
		// h := hash(GetRequestString(ref))
		h := md5.Sum([]byte(GetRequestString(ref)))
		ret = fmt.Sprintf("%x", h)
	}

	return ret
}

func GetCallerPackage(skip int) string {
	var ret string
	if pc, _, _, ok := runtime.Caller(skip); ok {
		funcName := runtime.FuncForPC(pc).Name()
		slash := strings.LastIndexByte(funcName, '/')
		if slash < 0 {
			slash = 0
		}
		dot := strings.IndexByte(funcName, '.')
		ret = funcName[slash+1:dot]
	}
	return ret
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

type Required []string

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetOptionsRequired(ref interface{}) Required {
	var ret []string

	for range Only.Once {
		t := reflect.TypeOf(ref)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			required := field.Tag.Get("required")
			if required == "" {
				continue
			}

			ret = append(ret, field.Name)
		}
	}

	return ret
}

type StructKey struct {
	Name      string
	JsonName  string
	JsonValue string
	Required  string
	Value     string
	Type      reflect.Type
	Interface interface{}
}
type StructKeys map[string]StructKey

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetStructKeys(ref interface{}, keys ...string) StructKeys {
	ret := make(StructKeys)

	s := New(ref)
	if len(keys) == 0 {
		keys = s.Names()
	}

	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	for _, k := range New(ref).Fields() {
		if _, ok := keyMap[k.Name()]; !ok {
			continue
		}

		jValue := ""
		value := ""
		switch k.value.Type().Name() {
			case "string":
				value = k.Value().(string)
			case "int":
				v := k.Value().(int)
				value = strconv.FormatInt(int64(v), 10)
			case "int64":
				value = strconv.FormatInt(k.Value().(int64), 10)
			case "float64":
				value = strconv.FormatFloat(k.Value().(float64), 'f', -1, 64)
			default:
				j, e := json.Marshal(k.Value())
				if e == nil {
					jValue = string(j)
				} else {
					jValue = fmt.Sprintf("%v", k.Value())
				}
		}

		ret[k.Name()] = StructKey {
			Name:      k.Name(),
			JsonName:  k.Tag("json"),
			JsonValue: jValue,
			Value:     value,
			Required:  k.Tag("required"),
			Type:      k.value.Type(),
			Interface: k.Value(),
		}
	}

	return ret
}

func (r *Required) IsRequired(field string) bool {
	var ok bool
	for _, f := range *r {
		if f == field {
			ok = true
		}
	}
	return ok
}

func (r *Required) IsNotRequired(field string) bool {
	return !r.IsRequired(field)
}
