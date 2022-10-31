package apiReflect

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"os"
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
	PointNameAppend     = "PointNameAppend"     // Append PointNameFrom instead of replace.

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

func (r Reflect) String() string {
	var ret string
	for range Only.Once {
		fn := r.FieldName
		if strings.HasPrefix(fn, "<struct {") {
			fn = "struct {}"
		}
		switch r.Kind {
			case reflect.Array:
				fallthrough
			case reflect.Map:
				fallthrough
			case reflect.Slice:
				ret = fmt.Sprintf("%s (%s)\tFieldName:%s Kind:%s Length:%d (IsNil:%v IsUnknown:%v IsExported:%v)",
					strings.Join(r.DataStructure.Endpoint, "."),
					r.DataStructure.PointId,
					fn,
					r.Kind.String(),
					r.Length,
					r.IsNil, r.IsUnknown, r.IsExported,
				)

			default:
				ret = fmt.Sprintf("%s (%s)\tFieldName:%s Kind:%s (IsNil:%v IsUnknown:%v IsExported:%v)",
					strings.Join(r.DataStructure.Endpoint, "."),
					r.DataStructure.PointId,
					fn,
					r.Kind.String(),
					r.IsNil, r.IsUnknown, r.IsExported,
				)
		}
	}
	return ret
}

func (r *Reflect) SetByFieldName(parent interface{}, current interface{}, name string) { // , fieldTo reflect.StructField, fieldVo reflect.Value) {
	for range Only.Once {
		r.Valid = true
		r.Parent = parent
		r.Interface = current
		r.IsNil = valueTypes.IsNil(current)
		r.IsUnknown = valueTypes.IsUnknownStruct(current)
		r.TypeOf = reflect.TypeOf(current)
		r.ValueOf = reflect.ValueOf(current)
		r.Kind = r.TypeOf.Kind()
		r.FieldName = name
		r.DataStructure.PointNameAppend = true

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

		// if name == "" {
		// 	break
		// }

		p := reflect.TypeOf(current)
		if p.Kind() != reflect.Struct {
			break
		}


		sf, ok := p.FieldByName(name)
		if !ok {
			break
		}

		r.FieldTo = sf
		r.IsExported = r.FieldTo.IsExported()
		r.FieldVo = reflect.ValueOf(current).FieldByName(name)

		r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
		// r.SetFieldName(parent, current, fieldName)
		// r.DataStructure = r.DataStructure.Set(ref, r.FieldTo, r.FieldVo)
	}
}
// func (r *Reflect) SetFieldName(parent interface{}, current interface{}, fieldName string) { // , fieldTo reflect.StructField, fieldVo reflect.Value) {
// 	for range Only.Once {
// 		if fieldName == "" {
// 			break
// 		}
//
// 		p := reflect.TypeOf(current)
// 		if p.Kind() != reflect.Struct {
// 			break
// 		}
//
//
// 		sf, ok := p.FieldByName(fieldName)
// 		if !ok {
// 			break
// 		}
//
// 		r.FieldTo = sf
// 		r.IsExported = r.FieldTo.IsExported()
// 		r.FieldVo = reflect.ValueOf(current).FieldByName(fieldName)
//
// 		r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
// 	}
// }

func (r *Reflect) SetByIndex(parent Reflect, current Reflect, index int, name Endpoint) {
	for range Only.Once {
		// Get child interface from parent.
		// pt := current.TypeOf
		// pv := current.ValueOf
		switch current.Kind {
			case reflect.Struct:
				r.Interface = current.ValueOf.Field(index).Interface()
			case reflect.Slice:
				r.Interface = current.ValueOf.Index(index).Interface()
			case reflect.Array:
				r.Interface = current.ValueOf.Index(index).Interface()
			case reflect.Map:
				mk := current.ValueOf.MapKeys()
				r.Interface = current.ValueOf.MapIndex(mk[index]).Interface()
		}

		r.Valid = true
		r.Parent = parent.Interface
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

		if current.TypeOf.Kind() == reflect.Struct {
			r.FieldTo = current.TypeOf.Field(index)
			r.IsExported = r.FieldTo.IsExported()
			r.FieldVo = current.ValueOf.Field(index)
			r.FieldName = r.FieldTo.Name

			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			// r.DataStructure.Endpoint = append(name, r.DataStructure.PointId)
			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, r.DataStructure.PointId)
			break
		}

		if (current.TypeOf.Kind() == reflect.Array) || (current.TypeOf.Kind() == reflect.Slice) {
			// pt = reflect.TypeOf(parent)
			// r.FieldTo = current.TypeOf.Field(index)
			// r.IsExported = r.FieldTo.IsExported()
			r.FieldTo = reflect.StructField{}
			r.IsExported = true
			r.FieldVo = current.ValueOf.Index(index)
			r.FieldName = r.FieldVo.String()
			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			r.DataStructure.Value = r.FieldVo.Interface()
			if r.Length == 0 {
				r.DataStructure.PointNameAppend = false
			}

			f := "%d"
			if current.Length > 0 {
				f = fmt.Sprintf("%%.%dd", valueTypes.SizeOfInt(current.Length))
			}
			f = fmt.Sprintf(f, index)

			// r.DataStructure.Endpoint = append(name, f)
			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, f)

			if r.IsUnknown {
				r.DataStructure.Json = current.DataStructure.PointId
				r.DataStructure.PointId = current.DataStructure.PointId
				r.DataStructure.PointUnit = current.DataStructure.PointUnit
				break
			}
			r.DataStructure.Json = current.DataStructure.PointId + "_" + f
			r.DataStructure.PointId = current.DataStructure.PointId + "_" + f
			r.DataStructure.PointUnit = current.DataStructure.PointUnit
			break
		}

		if current.TypeOf.Kind() == reflect.Map {
			// pt = reflect.TypeOf(parent)
			// r.FieldTo = current.TypeOf.Field(index)
			// r.IsExported = r.FieldTo.IsExported()
			r.FieldTo = reflect.StructField{}
			r.IsExported = true
			mk := current.ValueOf.MapKeys()
			r.FieldVo = current.ValueOf.MapIndex(mk[index])
			r.FieldName = mk[index].String()

			r.DataStructure = r.DataStructure.Set(parent.Interface, current.Interface, r.FieldTo, r.FieldVo)
			r.DataStructure.Json = mk[index].String()
			r.DataStructure.PointId = mk[index].String()
			r.DataStructure.Value = r.FieldVo.Interface()

			// r.DataStructure.Endpoint = append(name, r.DataStructure.PointId)
			r.DataStructure.Endpoint = name.Copy()
			r.DataStructure.Endpoint = append(r.DataStructure.Endpoint, r.DataStructure.PointId)
			break
		}
	}
}
// func (r *Reflect) SetByIndex(parent interface{}, current interface{}, index int) {
// 	for range Only.Once {
// 		// Get child interface from parent.
// 		pt := reflect.TypeOf(current)
// 		pv := reflect.ValueOf(current)
// 		pk := pt.Kind()
// 		switch pk {
// 		case reflect.Struct:
// 			r.Interface = pv.Field(index).Interface()
// 		case reflect.Slice:
// 			r.Interface = pv.Index(index).Interface()
// 		case reflect.Array:
// 			r.Interface = pv.Index(index).Interface()
// 		case reflect.Map:
// 			mk := pv.MapKeys()
// 			r.Interface = pv.MapIndex(mk[index]).Interface()
// 		}
//
// 		r.Valid = true
// 		r.Parent = parent
// 		r.IsNil = valueTypes.IsNil(r.Interface)
// 		r.IsUnknown = valueTypes.IsUnknownStruct(r.Interface)
// 		r.TypeOf = reflect.TypeOf(r.Interface)
// 		r.ValueOf = reflect.ValueOf(r.Interface)
// 		if r.IsNil {
// 			r.Kind = reflect.Invalid
// 		} else {
// 			r.Kind = r.TypeOf.Kind()
// 		}
//
// 		r.Length = -1
// 		if r.Kind == reflect.Struct {
// 			r.Length = r.ValueOf.NumField()
// 		}
// 		if r.Kind == reflect.Slice {
// 			r.Length = r.ValueOf.Len()
// 		}
// 		if r.Kind == reflect.Array {
// 			r.Length = r.ValueOf.Len()
// 		}
// 		if r.Kind == reflect.Map {
// 			r.Length = len(r.ValueOf.MapKeys())
// 		}
//
// 		if pt.Kind() == reflect.Struct {
// 			r.FieldTo = pt.Field(index)
// 			r.IsExported = r.FieldTo.IsExported()
// 			r.FieldVo = pv.Field(index)
// 			r.FieldName = r.FieldTo.Name
//
// 			r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
// 			break
// 		}
//
// 		if (pt.Kind() == reflect.Array) || (pt.Kind() == reflect.Slice) {
// 			// pt = reflect.TypeOf(parent)
// 			// r.FieldTo = pt.Field(index)
// 			// r.IsExported = r.FieldTo.IsExported()
// 			r.FieldTo = reflect.StructField{}
// 			r.IsExported = true
// 			r.FieldVo = pv.Index(index)
// 			r.FieldName = r.FieldVo.String()
//
// 			r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
// 			r.DataStructure.Json = strconv.Itoa(index)
// 			r.DataStructure.PointId = strconv.Itoa(index)
// 			r.DataStructure.Value = r.FieldVo.Interface()
// 			if r.Length == 0 {
// 				r.DataStructure.PointNameAppend = false
// 			}
// 			break
// 		}
//
// 		if pt.Kind() == reflect.Map {
// 			// pt = reflect.TypeOf(parent)
// 			// r.FieldTo = pt.Field(index)
// 			// r.IsExported = r.FieldTo.IsExported()
// 			r.FieldTo = reflect.StructField{}
// 			r.IsExported = true
// 			mk := pv.MapKeys()
// 			r.FieldVo = pv.MapIndex(mk[index])
// 			r.FieldName = mk[index].String()
//
// 			r.DataStructure = r.DataStructure.Set(parent, current, r.FieldTo, r.FieldVo)
// 			r.DataStructure.Json = mk[index].String()
// 			r.DataStructure.PointId = mk[index].String()
// 			r.DataStructure.Value = r.FieldVo.Interface()
// 			break
// 		}
// 	}
// }

// setPointName - Are we using an index number for name or field key value?
func (r *Reflect) setPointName(parent Reflect, current Reflect, name []string, index int) []string {
	for range Only.Once {
		// pointTimestamp := time.Now()
		// pointTimestampFrom := fieldTo.Tag.Get(PointTimestampFrom)
		// if pointTimestampFrom != "" {
		// 	pointTimestamp = GetTimestampFrom(parentRef, pointTimestampFrom, valueTypes.DateTimeLayout)
		// }

		var pn string
		var intSize int
		switch current.Kind {
			case reflect.Array:
				fallthrough
			case reflect.Slice:
				ft := valueTypes.GetIntFormatForPrintf(r.Length)
				pn = fmt.Sprintf(ft, index)
				intSize = valueTypes.SizeOfInt(r.Length)

			case reflect.Map:
				// mk := reflect.ValueOf(current).MapKeys()
				// pn = mk[index].String()
				pn = r.DataStructure.PointId

			case reflect.Struct:
				pn = r.DataStructure.PointId
		}

		switch {
			case r.DataStructure.PointNameFromChild != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(current.Interface, r.DataStructure.PointNameFromChild, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}

			case r.DataStructure.PointNameFromParent != "":
				// PointNameFromChild - In this case points to a field within a CHILD struct.
				pn = GetPointNameFrom(parent.Interface, r.DataStructure.PointNameFromParent, intSize, r.DataStructure.PointNameDateFormat)
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}

			default:
				if r.DataStructure.PointNameAppend == false {
					name = append(name[:len(name) - 1], pn)
				} else {
					name = append(name, pn)
				}
		}
	}

	return name
}

func (r *Reflect) PointNameFromChild(child Reflect, name Endpoint) []string {
	for range Only.Once {
		// var pn string
		// var intSize int
		// switch current.Kind {
		// 	case reflect.Array:
		// 			fallthrough
		// 	case reflect.Slice:
		// 		ft := valueTypes.GetIntFormatForPrintf(r.Length)
		// 		pn = fmt.Sprintf(ft, index)
		// 		intSize = valueTypes.SizeOfInt(r.Length)
		//
		// 	case reflect.Map:
		// 		// mk := reflect.ValueOf(current).MapKeys()
		// 		// pn = mk[index].String()
		// 		pn = r.DataStructure.PointId
		//
		// 	case reflect.Struct:
		// 		pn = r.DataStructure.PointId
		// }

		if r.DataStructure.PointNameFromChild != "" {
			// PointNameFromChild - In this case points to a field within a CHILD struct.
			pn := GetPointNameFrom(child.Interface, r.DataStructure.PointNameFromChild, 0, r.DataStructure.PointNameDateFormat)
			// if r.DataStructure.PointNameAppend == false {
			// 	name = append(name[:len(name) - 1], pn)
			// } else {
				name = append(name, pn)
			// }
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
	PointNameAppend           bool
	PointNameFromChild        string
	PointNameFromParent       string
	PointNameDateFormat       string

	PointIgnore               bool
	PointIgnoreZero           bool
	PointIgnoreIfNil          string
	PointIgnoreIfNilFromChild string

	PointGroupName            string
	PointGroupNameFrom        string

	PointArrayFlatten         bool
	PointSplitOn              string
	PointSplitOnType          string

	Value                     interface{}
	ValueType                 string
	ValueKind                 string
	Endpoint                  Endpoint
}

type Endpoint []string

func (e *Endpoint) Copy() Endpoint {
	ret := make(Endpoint, len(*e))
	copy(ret, *e)
	return ret
}

func (e Endpoint) String() string {
	return strings.Join(e, ".")
}

func (ds *DataStructure) Set(parent interface{}, current interface{}, fieldTo reflect.StructField, fieldVo reflect.Value) DataStructure {
	for range Only.Once {
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

		pointNameAppend := true
		pna := fieldTo.Tag.Get(PointNameAppend)
		if pna == "false" {
			pointNameAppend = false
		}

		pointArrayFlatten := false
		paf := fieldTo.Tag.Get(PointArrayFlatten)
		if paf == "true" {
			pointArrayFlatten = true
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
			PointNameAppend:     pointNameAppend,
			PointNameFromChild:  fieldTo.Tag.Get(PointNameFromChild),
			PointNameFromParent: fieldTo.Tag.Get(PointNameFromParent),
			PointNameDateFormat: pointNameDateFormat,

			PointUpdateFreq:           pointUpdateFreq,
			PointValueType:            fieldTo.Tag.Get(PointValueType),
			PointAliasTo:              fieldTo.Tag.Get(PointAliasTo),
			PointIgnore:               ignore,
			PointIgnoreIfNil:          pointIgnoreIfNil,
			PointIgnoreIfNilFromChild: fieldTo.Tag.Get(PointIgnoreIfNilFromChild),
			PointArrayFlatten:         pointArrayFlatten,
			PointSplitOn:              fieldTo.Tag.Get(PointSplitOn),
			PointSplitOnType:          fieldTo.Tag.Get(PointSplitOnType),
			PointIgnoreZero:           pointIgnoreZero,

			Value:           nil,
			ValueType:       valueType,
			ValueKind:       fieldVo.Kind().String(),
			Endpoint:        Endpoint{},	// strings.TrimPrefix(strings.Join(name, "."), "."),
		}

	}
	return *ds
}


type DataStructures struct {
	Map       map[string]DataStructure
	ShowEmpty bool
	Debug     bool
}

func (dss *DataStructures) Add(ds DataStructure)  {
	for range Only.Once {
		if dss.Map == nil {
			dss.Map = make(map[string]DataStructure)
		}
		dss.Map[ds.Endpoint.String()] = ds
		if dss.Debug {
			_, _ = fmt.Fprintf(os.Stderr, "DEBUG DataStructures.Add() %s - Kind:'%s' Type:'%s'\n", ds.Endpoint.String(), ds.ValueKind, ds.ValueType)
		}
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

func (dss *DataStructures) GetPointTags(Parent Reflect, Current Reflect, name Endpoint) DataStructures {

	for range Only.Once {
		if dss.Debug {
			_, _ = fmt.Fprintf(os.Stderr,"GetPointTags() Current: %s\n", Current)
		}

		name = dss.PointNameAppend(Parent, Current, name)

		if dss.PointIgnoreZero(Parent, Current, name) {
			break
		}

		if dss.PointIgnoreIfNilFromChild(Parent, Current, name) {
			break
		}

		// Parent.SetByFieldName(Parent.Interface, Current.Interface, "")

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
			if dss.ShowEmpty {
				if Current.Length == 0 {
					Current.DataStructure.PointArrayFlatten = true
				}
			}
			if dss.PointArrayFlatten(Parent, Current, name) {
				break
			}

			for si := 0; si < Current.Length; si++ {
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, name)
				if dss.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
				}
				name2 := Child.DataStructure.Endpoint.Copy()

				if Current.DataStructure.PointNameFromChild != "" {
					name2 = Current.PointNameFromChild(Child, name)
				}

				if Child.IsUnknown {
					dss.GetPointTags(Current, Child, name2)
					continue
				}

				if Child.Kind == reflect.Slice {
					if Child.IsUnknown {
						dss.GetPointTags(Current, Child, name2)
						continue
					}

					if dss.PointSplitOn(Current, Child, name2) {
						continue
					}

					Child.DataStructure.Value = Child.Interface
					// Child.DataStructure.Endpoint = append(name2)	//	+ "." + Child.DataStructure.PointId
					dss.Add(Child.DataStructure)
					continue
				}
				dss.Add(Child.DataStructure)
			}
			break
		}

		if Current.Kind == reflect.Array {
			// Handle arrays here.
			if dss.ShowEmpty {
				if Current.Length == 0 {
					Current.DataStructure.PointArrayFlatten = true
				}
			}
			if dss.PointArrayFlatten(Parent, Current, name) {
				break
			}

			for si := 0; si < Current.Length; si++ {
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, name)
				if dss.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
				}
				name2 := Child.DataStructure.Endpoint.Copy()

				if Child.IsUnknown {
					dss.GetPointTags(Current, Child, name2)
					continue
				}

				if Child.Kind == reflect.Slice {
					if Child.IsUnknown {
						dss.GetPointTags(Current, Child, name2)
						continue
					}

					if dss.PointSplitOn(Current, Child, name2) {
						continue
					}

					Child.DataStructure.Value = Child.Interface
					// Child.DataStructure.Endpoint = append(name2)	//	+ "." + Child.DataStructure.PointId
					dss.Add(Child.DataStructure)
					continue
				}
				dss.Add(Child.DataStructure)
			}
			break
		}

		if Current.Kind == reflect.Map {
			for si := range Current.FieldVo.MapKeys() {
				// @TODO - Implement pointNameFromChild / pointNameFromParent.
				// @TODO - Need to look at other types, besides known types.
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, name)
				if dss.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
				}
				name2 := Child.DataStructure.Endpoint.Copy()

				if Child.IsUnknown {
					dss.GetPointTags(Current, Child, name2)
					continue
				}

				if dss.PointSplitOn(Current, Child, name2) {
					continue
				}

				dss.Add(Child.DataStructure)
			}
			break
		}

		if Current.Kind == reflect.Struct {
			// Iterate over all available fields and read the tag value
			for si := 0; si < Current.Length; si++ {
				var Child Reflect
				Child.SetByIndex(Parent, Current, si, name)
				if dss.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
				}
				name2 := Child.DataStructure.Endpoint.Copy()
				if Child.DataStructure.PointId == "images" {
					fmt.Printf("")
				}

				if !Child.IsExported {
					_, _ = fmt.Fprintf(os.Stderr, "WARNING: Field '%s' type not exported (%s): Type %s\n", Child.FieldName, name2, Child.Kind.String())
					continue
				}

				if Child.Kind == reflect.Struct {
					if Child.IsUnknown {
						dss.GetPointTags(Current, Child, name2)
						continue
					}

					if dss.PointSplitOn(Current, Child, name2) {
						continue
					}

					Child.DataStructure.Value = Child.Interface
					// Child.DataStructure.Endpoint = append(name2)	//	+ "." + Child.DataStructure.PointId
					dss.Add(Child.DataStructure)
					continue
				}

				dss.GetPointTags(Current, Child, name2)

				// switch Child.Kind {
				// 	case reflect.Uintptr:
				// 		fallthrough
				// 	case reflect.Complex64:
				// 		fallthrough
				// 	case reflect.Complex128:
				// 		fallthrough
				// 	case reflect.Chan:
				// 		fallthrough
				// 	case reflect.Func:
				// 		fallthrough
				// 	case reflect.UnsafePointer:
				// 		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Child.FieldName, name2, Child.Kind.String())
				// 		continue
				//
				// 	case reflect.Invalid:
				// 		fallthrough
				// 	case reflect.Interface:
				// 		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Child.FieldName, name2, Child.Kind.String())
				// 		if dss.ShowNull {
				// 			dss.GetPointTags(Current, Child, name2)
				// 		}
				// 		continue
				//
				// 	case reflect.Map:
				// 		fallthrough
				// 	case reflect.Slice:
				// 		fallthrough
				// 	case reflect.Array:
				// 		fallthrough
				// 	case reflect.Pointer:
				// 		dss.GetPointTags(Current, Child, name2)
				// 		continue
				// }
				//
				// if dss.PointSplitOn(Current, Child, name2) {
				// 	continue
				// }
				//
				// // @TODO - Need to fix this!
				// // This parses ordinary builtin types.
				// Child.DataStructure.Value = Child.Interface
				// // Child.DataStructure.Endpoint = append(name)		// + "." + Child.DataStructure.PointId
				// dss.Add(Child.DataStructure)
			}
			break
		}

		if Current.Kind == reflect.Invalid {
			if dss.ShowEmpty {
				dss.Add(Current.DataStructure)
				break
			}
			_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Current.FieldName, name, Current.Kind.String())
			break
		}

		if Current.Kind == reflect.Interface {
			if dss.ShowEmpty {
				dss.Add(Current.DataStructure)
				break
			}
			_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Current.FieldName, name, Current.Kind.String())
			break
		}

		// case reflect.Uintptr:
		// case reflect.Complex64:
		// case reflect.Complex128:
		// case reflect.Chan:
		// case reflect.Func:
		// case reflect.UnsafePointer:

		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported (%s): Type %s\n", Current.FieldName, name, Current.Kind.String())
	}

	return *dss
}

func (dss *DataStructures) PointNameAppend(_ Reflect, Current Reflect, name Endpoint) []string {
	for range Only.Once {
		if Current.DataStructure.PointNameAppend == false {
			if len(name) == 0 {
				break
			}
			name = name[:len(name) - 1]
		}
	}
	return name
}

func (dss *DataStructures) PointArrayFlatten(_ Reflect, Current Reflect, name Endpoint) bool {
	var yes bool
	for range Only.Once {
		if Current.DataStructure.PointArrayFlatten == true {
			// We want to flatten a slice down to EG "[1, 2, 3]"
			Current.DataStructure.Value = valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
			Current.DataStructure.Endpoint = name.Copy()	// + "." + Current.DataStructure.PointId
			dss.Add(Current.DataStructure)
			yes = true
		}
	}
	return yes
}

func (dss *DataStructures) PointIgnoreZero(_ Reflect, Current Reflect, _ Endpoint) bool {
	var yes bool
	for range Only.Once {
		if !Current.DataStructure.PointIgnoreZero {
			yes = false
			break
		}
		if dss.ShowEmpty {
			yes = false
			break
		}
		if Current.Length == 0 {
			yes = true
		}
		yes = false
	}
	return yes
}

func (dss *DataStructures) PointIgnoreIfNilFromChild(Parent Reflect, Current Reflect, _ Endpoint) bool {
	var yes bool
	for range Only.Once {
		if Parent.DataStructure.PointIgnoreIfNilFromChild == "" {
			yes = false
			break
		}
		if dss.ShowEmpty {
			yes = false
			break
		}
		ret := GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
		if ret == "" {
			yes = false
			break
		}
		yes = true
	}
	return yes
}

func (dss *DataStructures) PointSplitOn(_ Reflect, Current Reflect, _ Endpoint) bool {
	var yes bool
	for range Only.Once {
		if Current.DataStructure.PointSplitOn == "" {
			break
		}
		// We want to split a string into separate points - currently only handles string types.
		// @TODO - Fix this up! - Use PointSplitOnType
		// soEP := Current.DataStructure.Endpoint
		soVal := valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
		soSplit := strings.Split(soVal, Current.DataStructure.PointSplitOn)
		for soI, soV := range soSplit {
			Current.DataStructure.Value = soV
			// Current.DataStructure.Endpoint = fmt.Sprintf("%s.%s", soEP, valueTypes.PrintInt(2, soI))
			Current.DataStructure.Endpoint = append(Current.DataStructure.Endpoint, valueTypes.PrintInt(2, soI))
			// Ref.DataStructure.Endpoint = fmt.Sprintf("%s.%s", soEP, valueTypes.PrintInt(len(soSplit), soI))
			// Ref.DataStructure.Endpoint = fmt.Sprintf("%s.%d", soEP, soI)
			dss.Add(Current.DataStructure)
		}
		yes = true
	}
	return yes
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
				Child.SetByIndex(Ref, Ref, i, Endpoint{})
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
				Child.SetByIndex(Ref, Ref, i, Endpoint{})
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
				Child.SetByIndex(Ref, Ref, i, Endpoint{})
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
			// value := valueTypes.AnyToValueString(fieldVo.Interface(), 0, valueTypes.DateTimeLayoutSecond)
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
