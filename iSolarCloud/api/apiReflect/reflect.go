package apiReflect

// import "C"
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


const (
	PointId                   = "PointId"                   // Point id in the form p\d+ or \d+
	PointParentId             = "PointParentId"             // Associated parent of point.
	PointUpdateFreq           = "PointUpdateFreq"           // Point update frequency - Total, Yearly, Monthly, Day.
	PointValueType            = "PointValueType"            // Value type of point: energy, date, battery, temperature.
	PointIgnore               = "PointIgnore"               // Ignore this point.
	PointIgnoreIfNil          = "PointIgnoreIfNil"          // Ignore this point if a child is nil or empty.
	PointIgnoreIfNilFromChild = "PointIgnoreIfNilFromChild" // Ignore this point if a child is nil or empty.

	PointAliasTo   = "PointAliasTo"   // Alias this point to another point.
	PointAliasFrom = "PointAliasFrom" // Alias this point from another point.

	PointUnit           = "PointUnit"           // Units: Wh, kWh, C, h.
	PointUnitFrom       = "PointUnitFrom"       // Get PointUnit from another field structure.
	PointUnitFromParent = "PointUnitFromParent" // Get PointUnit from another parent field structure.

	PointGroupName     = "PointGroupName"     // Point group name.
	PointGroupNameFrom = "PointGroupNameFrom" // Get PointGroupName from another field structure.

	PointName           = "PointName"           // Human-readable name of point.
	PointNameFromChild  = "PointNameFromChild"  // Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameFromParent = "PointNameFromParent" // Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameDateFormat = "PointNameDateFormat" // Date format when using PointNameFrom, (if the field is a time.Time type).
	PointNameFromAppend = "PointNameFromAppend" // Append PointNameFrom instead of replace.

	PointArrayFlatten     = "PointArrayFlatten"     // Flatten an array into a string. EG: ["one", "two", "three"]
	PointSplitOn          = "PointSplitOn"          // Split a point into an array separating by defined string.
	PointSplitOnType      = "PointSplitOnType"      // What valueTypes will be used for a split.
	PointIgnoreZero       = "PointIgnoreZero"       // Ignore arrays with zero size, (default true).

	PointTimestampFrom = "PointTimestampFrom" // Pull timestamp from another field structure.
)

const (
	UpdateFreqInstant = "instant"
	UpdateFreq5Mins   = "5mins"
	UpdateFreqBoot    = "boot"
	UpdateFreqDay     = "daily"
	UpdateFreqMonth   = "monthly"
	UpdateFreqYear    = "yearly"
	UpdateFreqTotal   = "total"
)


type Reflect struct {
	Valid         bool
	DataStructure DataStructure
	Parent        interface{}
	Interface     interface{}
	IsNil         bool
	IsExported    bool
	IsUnknown     bool
	Kind          reflect.Kind
	TypeOf        reflect.Type
	ValueOf       reflect.Value

	Length        int
	FieldName     string
	FieldTo       reflect.StructField
	FieldVo       reflect.Value
}

func (r *Reflect) SetByFieldName(parent interface{}, current interface{}, fieldName string) { // , fieldTo reflect.StructField, fieldVo reflect.Value) {
	for range Only.Once {
		r.Valid = true
		r.Parent = parent
		r.Interface = current
		r.IsNil = valueTypes.IsNil(current)
		r.IsUnknown = valueTypes.IsUnknownStruct(current)
		r.TypeOf = reflect.TypeOf(current)
		r.ValueOf = reflect.ValueOf(current)
		r.Kind = r.TypeOf.Kind()
		r.FieldName = fieldName

		if r.Kind == reflect.Struct {
			r.Length = r.ValueOf.NumField()
		}

		if r.Kind == reflect.Slice {
			r.Length = r.ValueOf.Len()
		}

		if r.Kind == reflect.Array {
			r.Length = r.ValueOf.Len()
		}

		if r.Kind == reflect.Map {
			r.Length = len(r.ValueOf.MapKeys())
		}

		r.SetFieldName(parent, current, fieldName)
		// r.DataStructure = r.DataStructure.Set(ref, r.FieldTo, r.FieldVo)
	}
}

func (r *Reflect) SetFieldName(parent interface{}, current interface{}, fieldName string) { // , fieldTo reflect.StructField, fieldVo reflect.Value) {
	for range Only.Once {
		if fieldName == "" {
			break
		}

		p := reflect.TypeOf(current)
		if p.Kind() != reflect.Struct {
			break
		}


		sf, ok := p.FieldByName(fieldName)
		if !ok {
			break
		}

		r.FieldTo = sf
		r.IsExported = r.FieldTo.IsExported()
		r.FieldVo = reflect.ValueOf(current).FieldByName(fieldName)

		r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
	}
}

func (r *Reflect) SetByIndex(parent interface{}, current interface{}, fieldIndex int) {
	for range Only.Once {
		// Get child interface from parent.
		pt := reflect.TypeOf(current)
		pv := reflect.ValueOf(current)
		pk := pt.Kind()
		switch pk {
			case reflect.Struct:
				r.Interface = pv.Field(fieldIndex).Interface()
			case reflect.Slice:
				r.Interface = pv.Index(fieldIndex).Interface()
			case reflect.Array:
				r.Interface = pv.Index(fieldIndex).Interface()
			case reflect.Map:
				mk := pv.MapKeys()
				r.Interface = pv.MapIndex(mk[fieldIndex]).Interface()
		}

		r.Valid = true
		r.Parent = parent
		r.IsNil = valueTypes.IsNil(r.Interface)
		r.IsUnknown = valueTypes.IsUnknownStruct(r.Interface)
		r.TypeOf = reflect.TypeOf(r.Interface)
		r.ValueOf = reflect.ValueOf(r.Interface)
		if r.IsNil {
			r.Kind = reflect.Invalid
		} else {
			r.Kind = r.TypeOf.Kind()
		}

		r.Length = -1
		if r.Kind == reflect.Struct {
			r.Length = r.ValueOf.NumField()
		}
		if r.Kind == reflect.Slice {
			r.Length = r.ValueOf.Len()
		}
		if r.Kind == reflect.Array {
			r.Length = r.ValueOf.Len()
		}
		if r.Kind == reflect.Map {
			r.Length = len(r.ValueOf.MapKeys())
		}

		r.SetFieldNameByIndex(parent, current, fieldIndex)
		// r.DataStructure = r.DataStructure.Set(ref, r.FieldTo, r.FieldVo)
	}
}

func (r *Reflect) SetFieldNameByIndex(parent interface{}, current interface{}, fieldIndex int) { // , fieldTo reflect.StructField, fieldVo reflect.Value) {
	for range Only.Once {
		p := reflect.TypeOf(current)
		if p.Kind() == reflect.Struct {
			r.FieldTo = p.Field(fieldIndex)
			r.IsExported = r.FieldTo.IsExported()
			r.FieldVo = reflect.ValueOf(current).Field(fieldIndex)
			r.FieldName = r.FieldTo.Name

			r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
			break
		}

		if p.Kind() == reflect.Array {
			r.FieldTo = p.Field(fieldIndex)
			r.IsExported = r.FieldTo.IsExported()
			r.FieldVo = reflect.ValueOf(current).Field(fieldIndex)
			r.FieldName = r.FieldTo.Name

			r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
			break
		}
	}
}

// setPointName - Are we using an index number for name or field key value?
func (r *Reflect) setPointName(parent interface{}, current interface{}, name []string, index int) []string {
	for range Only.Once {
		// pointTimestamp := time.Now()
		// pointTimestampFrom := fieldTo.Tag.Get(PointTimestampFrom)
		// if pointTimestampFrom != "" {
		// 	pointTimestamp = GetTimestampFrom(parentRef, pointTimestampFrom, valueTypes.DateTimeLayout)
		// }

		ft := valueTypes.GetIntFormatForPrintf(r.Length)
		pn := fmt.Sprintf(ft, index)
		intSize := valueTypes.SizeOfInt(r.Length)

		switch {
			case r.DataStructure.PointNameFromChild != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(current, r.DataStructure.PointNameFromChild, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameFromAppend == "true" {
					name = append(name, pn)
				} else {
					name = append(name[:len(name) - 1], pn)
				}

			case r.DataStructure.PointNameFromParent != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(parent, r.DataStructure.PointNameFromParent, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameFromAppend == "true" {
					name = append(name, pn)
				} else {
					name = append(name[:len(name) - 1], pn)
				}

			default:
				name = append(name, pn)
		}
	}

	return name
}


type DataStructure struct {
	Required                  bool
	Json                      string
	PointId                   string
	PointParentId             string
	PointUpdateFreq           string
	PointValueType            string
	PointAliasTo              string

	PointTimestamp            time.Time
	PointTimestampFrom        string

	PointUnit                 string
	PointUnitFrom             string
	PointUnitFromParent       string

	PointName                 string
	PointNameFromAppend       string
	PointNameFromChild        string
	PointNameFromParent       string
	PointNameDateFormat       string

	PointIgnore               bool
	PointIgnoreZero           bool
	PointIgnoreIfNil          string
	PointIgnoreIfNilFromChild string

	PointGroupName            string
	PointGroupNameFrom        string

	PointArrayFlatten         string
	PointSplitOn              string
	PointSplitOnType          string

	Value                     interface{}
	ValueType                 string
	ValueKind                 string
	Endpoint                  string
}

func (ds *DataStructure) Set(parent interface{}, current interface{}, fieldTo reflect.StructField, fieldVo reflect.Value) DataStructure {
	for range Only.Once {
		// sf, ok := reflect.TypeOf(current).FieldByName(fieldName)
		// if !ok {
		// 	break
		// }
		// fieldTo := sf
		// fieldVo := reflect.ValueOf(current).FieldByName(fieldName)

		ignore := false
		if fieldTo.Tag.Get(PointIgnore) != "" {
			ignore = true
		}

		pointIgnoreIfNil := fieldTo.Tag.Get(PointIgnoreIfNil)
		if pointIgnoreIfNil != "" {
			ret := GetStringFrom(current, pointIgnoreIfNil)
			if (ret == "") || (ret == "--") {
				ignore = true
			}
		}

		// pointIgnoreIfNilFromChild := fieldTo.Tag.Get(PointIgnoreIfNilFromChild)
		// if pointIgnoreIfNilFromChild != "" {
		// 	ret := GetStringFrom(fieldVo.Interface(), pointIgnoreIfNilFromChild)
		// 	if (ret == "") || (ret == "--") {
		// 		ignore = true
		// 	}
		// }

		var pointIgnoreZero bool
		switch fieldTo.Tag.Get(PointIgnoreZero) {
			case "false":
				pointIgnoreZero = false
			case "true":
				pointIgnoreZero = true
			default:
				pointIgnoreZero = true
		}

		// if valueTypes.IsNil(ref) {
		// 	pointValueType = "NIL"
		// }

		pointJson := getJsonTag(fieldTo)
		pointId := fieldTo.Tag.Get(PointId)
		if pointId == "" {
			pointId = pointJson
		}

		pointUnit := fieldTo.Tag.Get(PointUnit)
		pointUnitFrom := fieldTo.Tag.Get(PointUnitFrom)
		pointUnitFromParent := fieldTo.Tag.Get(PointUnitFromParent)
		if pointUnitFrom != "" {
			pointUnit = GetStringFrom(current, pointUnitFrom)
		}
		if pointUnitFromParent != "" {
			pointUnit = GetStringFrom(parent, pointUnitFromParent)
		}

		pointGroupName := fieldTo.Tag.Get(PointGroupName)
		pointGroupNameFrom := fieldTo.Tag.Get(PointGroupNameFrom)
		if pointGroupNameFrom != "" {
			pointGroupName = GetStringFrom(current, pointGroupNameFrom)
		}

		pointTimestamp := time.Now()
		pointTimestampFrom := fieldTo.Tag.Get(PointTimestampFrom)
		if pointTimestampFrom != "" {
			pointTimestamp = GetTimestampFrom(current, pointTimestampFrom, valueTypes.DateTimeLayout)
		}

		var valueType string
		if fieldTo.Type != nil {
			valueType = fieldTo.Type.String()
		}

		pointName := fieldTo.Tag.Get(PointName)
		if pointName == "" {
			pointName = valueTypes.PointToName(pointId)
		}

		pointNameDateFormat := fieldTo.Tag.Get(PointNameDateFormat)
		if pointNameDateFormat == "" {
			pointNameDateFormat = valueTypes.DateTimeAltLayout
		}

		pointUpdateFreq := fieldTo.Tag.Get(PointUpdateFreq)
		switch pointUpdateFreq {
			case "UpdateFreqInstant":
				pointUpdateFreq = UpdateFreqInstant
			case "UpdateFreq5Mins":
				pointUpdateFreq = UpdateFreq5Mins
			case "UpdateFreqBoot":
				pointUpdateFreq = UpdateFreqBoot
			case "UpdateFreqDay":
				pointUpdateFreq = UpdateFreqDay
			case "UpdateFreqMonth":
				pointUpdateFreq = UpdateFreqMonth
			case "UpdateFreqYear":
				pointUpdateFreq = UpdateFreqYear
			case "UpdateFreqTotal":
				pointUpdateFreq = UpdateFreqTotal
		}

		var required bool
		req := fieldTo.Tag.Get("required")
		if req == "true" {
			required = true
		}

		*ds = DataStructure {
			Required:           required,
			Json:               pointJson,
			PointId:            pointId,
			PointParentId:      fieldTo.Tag.Get(PointParentId),

			PointUnit:           pointUnit,
			PointUnitFrom:       pointUnitFrom,
			PointUnitFromParent: pointUnitFromParent,

			PointTimestamp:     pointTimestamp,
			PointTimestampFrom: pointTimestampFrom,

			PointGroupName:     pointGroupName,
			PointGroupNameFrom: pointGroupNameFrom,

			PointName:           pointName,
			PointNameFromAppend: fieldTo.Tag.Get(PointNameFromAppend),
			PointNameFromChild:  fieldTo.Tag.Get(PointNameFromChild),
			PointNameFromParent: fieldTo.Tag.Get(PointNameFromParent),
			PointNameDateFormat: pointNameDateFormat,

			PointUpdateFreq:           pointUpdateFreq,
			PointValueType:            fieldTo.Tag.Get(PointValueType),
			PointAliasTo:              fieldTo.Tag.Get(PointAliasTo),
			PointIgnore:               ignore,
			PointIgnoreIfNil:          pointIgnoreIfNil,
			PointIgnoreIfNilFromChild: fieldTo.Tag.Get(PointIgnoreIfNilFromChild),
			PointArrayFlatten:         fieldTo.Tag.Get(PointArrayFlatten),
			PointSplitOn:              fieldTo.Tag.Get(PointSplitOn),
			PointSplitOnType:          fieldTo.Tag.Get(PointSplitOnType),
			PointIgnoreZero:           pointIgnoreZero,

			Value:           nil,
			ValueType:       valueType,
			ValueKind:       fieldVo.Kind().String(),
			Endpoint:        "",	// strings.TrimPrefix(strings.Join(name, "."), "."),
		}

	}
	return *ds
}


type DataStructures struct {
	Map map[string]DataStructure
}

func (dss *DataStructures) Add(ds DataStructure)  {
	for range Only.Once {
		if dss.Map == nil {
			dss.Map = make(map[string]DataStructure)
		}
		// fmt.Printf("DEBUG DataStructures.Add() %s - Kind:'%s' Type:'%s'\n", name, ds.ValueKind, ds.ValueType)
		dss.Map[ds.Endpoint] = ds
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

func (dss *DataStructures) GetPointTags(Parent Reflect, Current Reflect, name ...string) DataStructures {

	for range Only.Once {
		// fmt.Printf("%s - Parent Kind: %s\tCurrent Kind: %s\n", strings.Join(name, "."), Parent.Kind, Current.Kind)

		if Current.Kind == reflect.Pointer {
			// Special case:
			// We're going to change the pointer to a proper object reference.
			if Current.IsNil {
				break
			}
			ref2 := Current.ValueOf.Elem().Interface()
			if valueTypes.IsNil(ref2) {
				break
			}
			Current.SetByFieldName(Current.Interface, ref2, "")
			if Current.IsNil {
				break
			}

			// DO NOT BREAK!
			// KEEP FIRST!
		}

		if Current.Kind == reflect.Slice {
			// Handle slices here.
			if Current.Length == 0 {
				if Current.DataStructure.PointIgnoreZero {
					break
				}
			}

			if Parent.DataStructure.PointIgnoreIfNilFromChild != "" {
				ret := GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
				if ret == "" {
					break
				}
			}

			if Current.DataStructure.PointArrayFlatten != "" {
				// We want to flatten a slice down to EG "[1, 2, 3]"
				Current.DataStructure.Value = valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
				Current.DataStructure.Endpoint = strings.Join(name, ".")	// + "." + Current.DataStructure.PointId
				dss.Add(Current.DataStructure)
				break
			}

			if Current.IsUnknown {
				for si := 0; si < Current.Length; si++ {
					// @TODO - Need to cover types other than struct that may be referenced.

					var Child Reflect
					Child.SetByIndex(Parent.Interface, Current.Interface, si)
					name3 := Current.setPointName(Parent.Interface, Child.Interface, name, si)
					dss.GetPointTags(Current, Child, name3...)
				}
				break
			}

			// Flatten slice for []Integer / []Float objects.
			Current.DataStructure.Value = valueTypes.AnyToValueString(Current.Interface, 0, "")
			Current.DataStructure.Endpoint = strings.Join(name, ".")
			dss.Add(Current.DataStructure)
			break
		}

		if Current.Kind == reflect.Array {
			// Handle arrays here.
			if Current.Length == 0 {
				if Current.DataStructure.PointIgnoreZero {
					break
				}
			}

			if Parent.DataStructure.PointIgnoreIfNilFromChild != "" {
				ret := GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
				if ret == "" {
					break
				}
			}

			if Current.DataStructure.PointArrayFlatten != "" {
				// We want to flatten a slice down to EG "[1, 2, 3]"
				Current.DataStructure.Value = valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
				Current.DataStructure.Endpoint = strings.Join(name, ".")	// + "." + Current.DataStructure.PointId
				dss.Add(Current.DataStructure)
				break
			}

			if Current.IsUnknown {
				for si := 0; si < Current.Length; si++ {
					// @TODO - Need to cover types other than struct that may be referenced.

					var Child Reflect
					Child.SetByIndex(Parent.Interface, Current.Interface, si)
					name3 := Current.setPointName(Parent.Interface, Child.Interface, name, si)
					dss.GetPointTags(Current, Child, name3...)
				}
				break
			}

			// Flatten slice for []Integer / []Float objects.
			Current.DataStructure.Value = valueTypes.AnyToValueString(Current.Interface, 0, "")
			Current.DataStructure.Endpoint = strings.Join(name, ".")
			dss.Add(Current.DataStructure)
			break
		}

		if Current.Kind == reflect.Map {
			if Current.Length == 0 {
				if Current.DataStructure.PointIgnoreZero {
					break
				}
			}

			if Parent.DataStructure.PointIgnoreIfNilFromChild != "" {
				ret := GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
				if ret == "" {
					break
				}
			}

			// Parent.SetByFieldName(Parent.Interface, Current.Interface, "")
			// n2 := append(name, Current.DataStructure.PointId)
			for _, key := range Current.FieldVo.MapKeys() {
				// @TODO - Implement pointNameFromChild / pointNameFromParent.
				Current.DataStructure.Json = key.String()
				Current.DataStructure.PointId = key.String()
				val := Current.FieldVo.MapIndex(key)
				Current.DataStructure.Value = val.Interface()
				Current.DataStructure.Endpoint = strings.Join(name, ".") + "." + Current.DataStructure.PointId
				// @TODO - For integers, it'd be nice to format them with a 0 prefix.
				// ft := valueTypes.GetIntFormatForPrintf(r.Length)
				// pn := fmt.Sprintf(ft, index)
				// intSize := valueTypes.SizeOfInt(r.Length)
				// name3 := Current.setPointName(Parent.Interface, Current.Interface, name, si)
				// fmt.Printf("DEBUG: %s / %s\n", Current.DataStructure.Endpoint, name3)
				// @TODO - Need to look at other types, besides known types.
				dss.Add(Current.DataStructure)

				if Current.DataStructure.PointSplitOn != "" {
					// We want to split a string into separate points - currently only handles string types.
					// @TODO - Fix this up! - Use PointSplitOnType
					soEP := Current.DataStructure.Endpoint
					soVal := valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
					for soI, soV := range strings.Split(soVal, Current.DataStructure.PointSplitOn) {
						Current.DataStructure.Value = soV
						Current.DataStructure.Endpoint = fmt.Sprintf("%s.%d", soEP, soI)
						dss.Add(Current.DataStructure)
					}
				}
			}
			break
		}

		if Current.Kind == reflect.Struct {
			if Current.Length == 0 {
				if Current.DataStructure.PointIgnoreZero {
					break
				}
			}

			if Parent.DataStructure.PointIgnoreIfNilFromChild != "" {
				ret := GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
				if ret == "" {
					break
				}
			}

			// Parent.SetByFieldName(Parent.Interface, Current.Interface, "")
			// Iterate over all available fields and read the tag value
			for i := 0; i < Current.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Current.Parent, Current.Interface, i)
				name2 := append(name, Child.DataStructure.PointId)

				if !Child.IsExported {
					fmt.Printf("DEBUG: NOTEXPORTED(%s): Type %s\n", name, Child.FieldName)
					continue
				}

				// fmt.Printf("DEBUG[%d]: %s.%s: Key[%s]:\tKind:'%s' Type:'%s'\n",
				// 	len(dss.Map),
				// 	strings.Join(name, "."), pointId,
				// 	fieldTo.Name,
				// 	fieldVo.Kind(),
				// 	fieldTo.Type.String(),
				// )

				if Child.Kind == reflect.Struct {
					if Child.IsUnknown {
						dss.GetPointTags(Current, Child, name2...)
						continue
					}

					Child.DataStructure.Value = Child.Interface
					Child.DataStructure.Endpoint = strings.Join(name, ".")	+ "." + Child.DataStructure.PointId
					dss.Add(Child.DataStructure)

					if Child.DataStructure.PointSplitOn != "" {
						// We want to split a string into separate points - currently only handles string types.
						// @TODO - Fix this up! - Use PointSplitOnType
						soEP := Child.DataStructure.Endpoint
						soVal := valueTypes.AnyToValueString(Child.FieldVo.Interface(), 0, "")
						for soI, soV := range strings.Split(soVal, Child.DataStructure.PointSplitOn) {
							Child.DataStructure.Value = soV
							Child.DataStructure.Endpoint = fmt.Sprintf("%s.%d", soEP, soI)
							dss.Add(Child.DataStructure)
						}
					}
					continue
				}

				switch Child.Kind {
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
					case reflect.Map:
						fallthrough
					case reflect.Pointer:
						fallthrough
					case reflect.Slice:
						fallthrough
					case reflect.Array:
						fallthrough
					case reflect.UnsafePointer:
						dss.GetPointTags(Current, Child, name2...)

					default:
						// @TODO - Need to fix this!
						// This parses ordinary builtin types.
						Child.DataStructure.Value = Child.Interface
						Child.DataStructure.Endpoint = strings.Join(name, ".") + "." + Child.DataStructure.PointId
						dss.Add(Child.DataStructure)

						if Child.DataStructure.PointSplitOn != "" {
							// We want to split a string into separate points - currently only handles string types.
							// @TODO - Fix this up! - Use PointSplitOnType
							soEP := Current.DataStructure.Endpoint
							soVal := valueTypes.AnyToValueString(Child.FieldVo.Interface(), 0, "")
							for soI, soV := range strings.Split(soVal, Child.DataStructure.PointSplitOn) {
								Child.DataStructure.Value = soV
								Child.DataStructure.Endpoint = fmt.Sprintf("%s.%d", soEP, soI)
								dss.Add(Child.DataStructure)
							}
						}
					}
			}
			break
		}

		fmt.Printf("ERROR: Unsupported type: '%s.%s' (%s)\n", name, Current.DataStructure.PointId, Current.FieldVo.Type().String())
	}

	return *dss
}


func GetStructFields(ref interface{}) map[string]string {
	ret := make(map[string]string)

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref.Interface, Ref.Interface, i)
				// SetFieldNameByIndex

				if !Child.IsExported {
					continue
				}

				ret[Child.FieldName] = fmt.Sprintf("%v", Child.DataStructure.Required)
			}
			break
		}
	}

	return ret
}

func GetStructFieldsAsArray(ref interface{}) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref.Interface, Ref.Interface, i)
				// SetFieldNameByIndex

				if !Child.IsExported {
					continue
				}

				ret = append(ret, Child.FieldName)
			}
		}
	}

	return ret
}

func GetStructValuesAsArray(ref interface{}) []string {
	var ret []string

	for range Only.Once {
		var Ref Reflect
		Ref.SetByFieldName(ref, ref, "")

		if Ref.Kind == reflect.Struct {
			if Ref.Length == 0 {
				if Ref.DataStructure.PointIgnoreZero {
					break
				}
			}

			// Iterate over all available fields and read the tag value
			for i := 0; i < Ref.Length; i++ {
				var Child Reflect
				Child.SetByIndex(Ref.Interface, Ref.Interface, i)
				// SetFieldNameByIndex

				if !Child.IsExported {
					continue
				}

				ret = append(ret, fmt.Sprintf("%v", Child.Interface))
			}
		}
	}

	return ret
}

func GetStringFrom(ref interface{}, name string) string {
	var ret string
	for range Only.Once {
		vo := reflect.ValueOf(ref)

		switch vo.Kind() {
			case reflect.Struct:
				// Iterate over all available fields, looking for the field name.
				for i := 0; i < vo.NumField(); i++ {
					if vo.Type().Field(i).Name == name {
						ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
						break
					}
				}

			case reflect.Map:
				// Iterate over all available fields, looking for the field name.
				for _, key := range vo.MapKeys() {
					if key.String() == name {
						ret = valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), 0, "")
						break
					}
				}
		}
	}

	return ret
}

func GetTimestampFrom(ref interface{}, name string, dateFormat string) time.Time {
	var ret time.Time
	for range Only.Once {
		if dateFormat == "" {
			dateFormat = valueTypes.DateTimeAltLayout
		}
		vo := reflect.ValueOf(ref)

		switch vo.Kind() {
			case reflect.Struct:
				// Iterate over all available fields, looking for the field name.
				for i := 0; i < vo.NumField(); i++ {
					if vo.Type().Field(i).Name == name {
						v := fmt.Sprintf("%v", vo.Field(i).Interface())
						ret = valueTypes.SetDateTimeString(v).Time
						break
					}
				}

			case reflect.Map:
				// Iterate over all available fields, looking for the field name.
				for _, key := range vo.MapKeys() {
					if key.String() == name {
						v := fmt.Sprintf("%v", vo.MapIndex(key).Interface())
						ret = valueTypes.SetDateTimeString(v).Time
						break
					}
				}
		}
	}

	return ret
}

func GetPointNameFrom(ref interface{}, name string, intSize int, dateFormat string) string {
	var ret string
	for range Only.Once {
		if dateFormat == "" {
			dateFormat = valueTypes.DateTimeAltLayout
		}
		vo := reflect.ValueOf(ref)

		var ra []string
		switch vo.Kind() {
			case reflect.Struct:
				for _, pnf := range strings.Split(name, ".") {
					// Iterate over all available fields, looking for the field name.
					for i := 0; i < vo.NumField(); i++ {
						fn := vo.Type().Field(i).Name
						if fn == pnf {
							ra = append(ra, valueTypes.AnyToValueString(vo.Field(i).Interface(), intSize, dateFormat))
							break
						}
					}
				}

			case reflect.Map:
				for _, pnf := range strings.Split(name, ".") {
					// Iterate over all available keys, looking for the key name.
					for _, key := range vo.MapKeys() {
						if key.String() == pnf {
							ra = append(ra, valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), intSize, dateFormat))
							break
						}
					}
				}
		}
		ret = strings.Join(ra, ".")
	}

	return ret
}

func getJsonTag(fieldTo reflect.StructField) string {
	var ret string

	for range Only.Once {
		ret = fieldTo.Tag.Get("json")
		ret = strings.ReplaceAll(ret, "omitempty", "")
		ret = strings.TrimSuffix(ret, ",")
	}

	return ret
}


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
}

func GetPkgType(v interface{}) string {
	return reflect.ValueOf(v).Type().String()
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
		// @TODO - Move over to using Reflect structure.
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
