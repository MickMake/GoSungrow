package apiReflect

import (
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"reflect"
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

	PointAliasTo   = "PointAliasTo"   				// Alias this point to another point.
	PointAliasFrom = "PointAliasFrom" 				// Alias this point from another point.

	PointUnit           = "PointUnit"           	// Units: Wh, kWh, C, h.
	PointUnitFrom       = "PointUnitFrom"       	// Get PointUnit from another field structure.
	PointUnitFromParent = "PointUnitFromParent" 	// Get PointUnit from another parent field structure.

	PointGroupName     = "PointGroupName"     		// Point group name.
	PointGroupNameFrom = "PointGroupNameFrom" 		// Get PointGroupName from another field structure.

	PointName           = "PointName"           	// Human-readable name of point.
	PointNameFromChild  = "PointNameFromChild"  	// Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameFromParent = "PointNameFromParent" 	// Searches child for field value to use for naming when hitting a slice, (as opposed to using an index).
	PointNameDateFormat = "PointNameDateFormat" 	// Date format when using PointNameFrom, (if the field is a time.Time type).
	PointNameAppend     = "PointNameAppend"     	// Append PointNameFrom instead of replace.

	PointArrayFlatten     = "PointArrayFlatten"     // Flatten an array into a string. EG: ["one", "two", "three"]
	PointSplitOn          = "PointSplitOn"          // Split a point into an array separating by defined string.
	PointSplitOnType      = "PointSplitOnType"      // What valueTypes will be used for a split.
	PointIgnoreZero       = "PointIgnoreZero"       // Ignore arrays with zero size, (default true).

	PointTimestampFrom = "PointTimestampFrom" 		// Pull timestamp from another field structure.

	IsDataTable        = "DataTable"				// This entity is a data table - Will only traverse down one child.
	DataTableId        = "DataTableId"				// Table id, (defaults to Json tag).
	DataTableName      = "DataTableName"			// Table Name, (defaults to DataTableId).
	DataTableTitle     = "DataTableTitle"			// Table Title, (defaults to DataTableId in name format).
	DataTableMerge     = "DataTableMerge"			// Merge rows together - useful for when we use, for EG: []valueTypes.Float
	DataTableShowIndex = "DataTableShowIndex" 		// Show index on table.
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


type EndPointPath []string

func NewEndPointPath(path ...string) EndPointPath {
	var epp EndPointPath
	return epp.Append(path...)
}

func (e *EndPointPath) Copy() EndPointPath {
	ret := make(EndPointPath, len(*e))
	copy(ret, *e)
	return ret
}

func (e *EndPointPath) Append(path ...string) EndPointPath {
	ret := make(EndPointPath, len(*e))
	copy(ret, *e)
	for _, p := range path {
		ret = append(ret, p)
	}
	return ret
}

func (e EndPointPath) String() string {
	return strings.Join(e, ".")
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

	DataTable                 bool
	DataTableId               string
	DataTableName             string
	DataTableTitle            string
	DataTableMerge            bool
	DataTableShowIndex        bool

	Value                     interface{}
	ValueType                 string
	ValueKind                 string
	Endpoint                  EndPointPath
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

		required := false
		if fieldTo.Tag.Get("required") == "true" {
			required = true
		}

		pointNameAppend := true		// Defaults to true
		if fieldTo.Tag.Get(PointNameAppend) == "false" {
			pointNameAppend = false
		}

		pointArrayFlatten := false
		if fieldTo.Tag.Get(PointArrayFlatten) == "true" {
			pointArrayFlatten = true
		}

		dataTable := false
		if fieldTo.Tag.Get(IsDataTable) == "true" {
			dataTable = true
		}

		dataTableMerge := false
		if fieldTo.Tag.Get(DataTableMerge) == "true" {
			dataTableMerge = true
		}

		dataTableShowIndex := false
		if fieldTo.Tag.Get(DataTableShowIndex) == "true" {
			dataTableShowIndex = true
		}

		dtn := fieldTo.Tag.Get(DataTableName)
		if dtn == "" {
			dtn = valueTypes.PointToName(pointId)
		}

		did := fieldTo.Tag.Get(DataTableId)
		if did == "" {
			did = pointId
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

			DataTable:      dataTable,
			DataTableId:    did,
			DataTableName:  dtn,
			DataTableTitle: fieldTo.Tag.Get(DataTableTitle),
			DataTableMerge: dataTableMerge,
			DataTableShowIndex: dataTableShowIndex,

			Value:     fieldVo.Interface(),
			ValueType: valueType,
			ValueKind: fieldVo.Kind().String(),
			Endpoint:  EndPointPath{},
		}
	}
	return *ds
}

func COMPARE(name EndPointPath, ref1 interface{}, ref2 interface{}) {
	t1 := fmt.Sprintf("%v", ref1)
	t2 := fmt.Sprintf("%v", ref2)
	if t1 != t2 {
		fmt.Printf("[%s] VALUE ERROR: '%s' != '%s'\n", name, t1, t2)
	}
}

type DataStructures struct {
	DataMap    map[string]DataStructure
	DataTables DataTables
	ShowEmpty  bool
	Debug      bool
}

func (dss *DataStructures) GetPointTags(Parent Reflect, Current Reflect, name EndPointPath) DataStructures {

	for range Only.Once {
		if Current.DataStructure.DataTable {
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"GetPointTags() Adding DataTable - Current: %s\n", Current)
			}
			dss.AddTable(Current)
		}
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

		if !Current.IsExported {
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"GetPointTags(%s) Current: %s is NOT exported - Name:%s\n", Current.FieldName, Current, name)
			}
			break
		}

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

		switch Current.Kind {
			case reflect.Slice:
				fallthrough
			case reflect.Array:
				dss.ProcessSlice(Parent, Current, name)

			case reflect.Map:
				dss.ProcessMap(Parent, Current, name)

			case reflect.Struct:
				dss.ProcessStruct(Parent, Current, name)

			case reflect.Invalid:
				dss.ProcessUnsupported(Parent, Current, name)

			case reflect.Interface:
				dss.ProcessUnsupported(Parent, Current, name)

			default:
				_, _ = fmt.Fprintf(os.Stderr,"GetPointTags(%s) Current: %s type is NOT supported - Name:%s Kind:%s\n", Current.FieldName, Current, name, Current.Kind.String())

			// case reflect.Uintptr:
			// case reflect.Complex64:
			// case reflect.Complex128:
			// case reflect.Chan:
			// case reflect.Func:
			// case reflect.UnsafePointer:
		}
	}

	return *dss
}

func (dss *DataStructures) Add(ds DataStructure)  {
	for range Only.Once {
		if dss.DataMap == nil {
			dss.DataMap = make(map[string]DataStructure)
		}
		dss.DataMap[ds.Endpoint.String()] = ds
		if dss.Debug {
			_, _ = fmt.Fprintf(os.Stderr, "DEBUG DataStructures.Add() %s - Kind:'%s' Type:'%s'\n", ds.Endpoint.String(), ds.ValueKind, ds.ValueType)
		}
	}
}

func (dss *DataStructures) Exists(name string) bool {
	var yes bool
	for range Only.Once {
		if dss.DataMap == nil {
			break
		}
		if _, ok := dss.DataMap[name]; ok {
			yes = ok
			break
		}
	}
	return yes
}

func (dss *DataStructures) Append(dsm DataStructures)  {
	for range Only.Once {
		if dss.DataMap == nil {
			dss.DataMap = make(map[string]DataStructure)
		}
		for name, ds := range dsm.DataMap {
			dss.DataMap[name] = ds
		}
	}
}

func (dss *DataStructures) ProcessUnsupported(_ Reflect, Current Reflect, name EndPointPath) {
	for range Only.Once {
		if dss.ShowEmpty {
			dss.Add(Current.DataStructure)
			break
		}
		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Current.FieldName, name, Current.Kind.String())
	}
}

func (dss *DataStructures) ProcessSlice(Parent Reflect, Current Reflect, name EndPointPath) {
	for range Only.Once {
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

			if Child.Kind == reflect.Slice {
				if Child.IsUnknown() {
					dss.GetPointTags(Current, Child, name2)
					continue
				}

				if dss.PointSplitOn(Current, Child, name2) {
					continue
				}

				COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
				dss.Add(Child.DataStructure)
				continue
			}

			if Child.IsUnknown() {
				dss.GetPointTags(Current, Child, name2)
				continue
			}
			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) ProcessStruct(Parent Reflect, Current Reflect, name EndPointPath) {
	for range Only.Once {
		// Iterate over all available fields and read the tag value
		for si := 0; si < Current.Length; si++ {
			var Child Reflect
			Child.SetByIndex(Parent, Current, si, name)
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
			}
			name2 := Child.DataStructure.Endpoint.Copy()

			if !Child.IsExported {
				_, _ = fmt.Fprintf(os.Stderr, "WARNING: Field '%s' type not exported (%s): Type %s\n", Child.FieldName, name2, Child.Kind.String())
				continue
			}

			if Child.Kind == reflect.Struct {
				if Child.IsUnknown() {
					dss.GetPointTags(Current, Child, name2)
					continue
				}

				if dss.PointSplitOn(Current, Child, name2) {
					continue
				}

				COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
				dss.Add(Child.DataStructure)
				continue
			}

			if Child.IsUnknown() {
				dss.GetPointTags(Current, Child, name2)
				continue
			}
			COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) ProcessMap(Parent Reflect, Current Reflect, name EndPointPath) {
	for range Only.Once {
		for si := range Current.FieldVo.MapKeys() {
			// @TODO - Implement pointNameFromChild / pointNameFromParent.
			// @TODO - Need to look at other types, besides known types.
			var Child Reflect
			Child.SetByIndex(Parent, Current, si, name)
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
			}
			name2 := Child.DataStructure.Endpoint.Copy()

			if Child.IsUnknown() {
				dss.GetPointTags(Current, Child, name2)
				continue
			}

			if dss.PointSplitOn(Current, Child, name2) {
				continue
			}

			COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) PointNameAppend(_ Reflect, Current Reflect, name EndPointPath) []string {
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

func (dss *DataStructures) PointArrayFlatten(_ Reflect, Current Reflect, name EndPointPath) bool {
	var yes bool
	for range Only.Once {
		if Current.DataStructure.PointArrayFlatten == true {
			// We want to flatten a slice down to EG "[1, 2, 3]"
			Current.DataStructure.Value = valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
			Current.DataStructure.Endpoint = name.Copy()
			dss.Add(Current.DataStructure)
			yes = true
		}
	}
	return yes
}

func (dss *DataStructures) PointIgnoreZero(_ Reflect, Current Reflect, _ EndPointPath) bool {
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

func (dss *DataStructures) PointIgnoreIfNilFromChild(Parent Reflect, Current Reflect, _ EndPointPath) bool {
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

func (dss *DataStructures) PointSplitOn(_ Reflect, Current Reflect, _ EndPointPath) bool {
	var yes bool
	for range Only.Once {
		if Current.DataStructure.PointSplitOn == "" {
			break
		}
		// We want to split a string into separate points - currently only handles string types.
		// @TODO - Fix this up! - Use PointSplitType to define a particular type.
		soVal := valueTypes.AnyToValueString(Current.FieldVo.Interface(), 0, "")
		soEP := Current.DataStructure.Endpoint
		soSplit := strings.Split(soVal, Current.DataStructure.PointSplitOn)
		for soI, soV := range soSplit {
			Current.DataStructure.Value = strings.ReplaceAll(soV, "&", ".p")
			Current.DataStructure.Endpoint = append(soEP, valueTypes.PrintInt(2, soI))
			dss.Add(Current.DataStructure)
		}
		yes = true
	}
	return yes
}


type DataTables struct {
	Map   []*DataTable
	Merge bool
	Index bool
}

func (dt *DataTables) Get() []*DataTable {
	return dt.Map
}


type DataTable struct {
	Reflect Reflect
	Name    string
	Merge   bool
	Index   bool
	Headers []string
	Data    [][]Reflect
	Debug   bool
}

func (dss *DataStructures) AddTable(ref Reflect) *DataTable {
	var ret *DataTable
	for range Only.Once {
		if dss.DataTables.Map == nil {
			dss.DataTables.Map = []*DataTable{}
		}
		ret = &DataTable {
			Reflect: ref,
			Name:    ref.DataStructure.DataTableId,
			Merge:   ref.DataStructure.DataTableMerge,
			Index:   ref.DataStructure.DataTableShowIndex,
			Headers: nil,
			Data:    nil,
		}
		dss.DataTables.Map = append(dss.DataTables.Map, ret)
		if ref.DataStructure.DataTableMerge {
			dss.DataTables.Merge = true
		}
		if ref.DataStructure.DataTableShowIndex {
			dss.DataTables.Index = true
		}
		if dss.Debug {
			_, _ = fmt.Fprintf(os.Stderr, "DEBUG DataStructures.AddTable() %s - Kind:'%s' Type:'%s'\n",
				ref.DataStructure.Endpoint.String(), ref.DataStructure.ValueKind, ref.DataStructure.ValueType)
		}
	}
	return ret
}

func (dt *DataTable) GetTable() DataTable {

	for range Only.Once {
		if dt.Debug {
			_, _ = fmt.Fprintf(os.Stderr,"GetTable() Current[%s]: %s\n", dt.Reflect.DataStructure.DataTableId, dt.Reflect)
		}
		if !dt.Reflect.DataStructure.DataTable {
			break
		}

		if dt.Reflect.Kind == reflect.Pointer {
			// Special case:
			// We're going to change the pointer to a proper object reference.
			if dt.Reflect.IsNil {
				break
			}
			ref2 := dt.Reflect.ValueOf.Elem().Interface()
			if valueTypes.IsNil(ref2) {
				break
			}
			dt.Reflect.SetByFieldName(dt.Reflect.Interface, ref2, "")
			if dt.Reflect.IsNil {
				break
			}

			// DO NOT BREAK!
			// KEEP FIRST!
		}

		if dt.Reflect.Kind == reflect.Slice {
			// Handle slices here.
			for row := 0; row < dt.Reflect.Length; row++ {
				var Child Reflect
				Child.SetByIndex(dt.Reflect, dt.Reflect, row, EndPointPath{})
				if dt.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child[%s]: %s\n", dt.Reflect.DataStructure.DataTableId, Child)
				}

				if Child.IsKnown() {
					// We have a known value
					if row == 0 {
						dt.AddHeader(dt.Reflect)
					}
					dt.AddRow(Child)
					continue
				}

				if Child.Kind == reflect.Struct {
					var refs []Reflect

					for col := 0; col < Child.Length; col++ {
						var ChildStruct Reflect
						ChildStruct.SetByIndex(dt.Reflect, Child, col, EndPointPath{})
						if dt.Debug {
							_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", ChildStruct)
						}

						if !ChildStruct.IsExported {
							_, _ = fmt.Fprintf(os.Stderr, "WARNING: Field '%s' type not exported: Type %s\n", ChildStruct.FieldName, ChildStruct.Kind.String())
							continue
						}
						refs = append(refs, ChildStruct)
					}

					if row == 0 {
						dt.AddHeader(refs...)
					} else {
						dt.AddRow(refs...)
					}
				}
			}
			break
		}

		if dt.Reflect.Kind == reflect.Map {
			// Handle maps here.
			for row := range dt.Reflect.FieldVo.MapKeys() {
				var Child Reflect
				Child.SetByIndex(dt.Reflect, dt.Reflect, row, EndPointPath{})
				if dt.Debug {
					_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child[%s]: %s\n", dt.Reflect.DataStructure.DataTableId, Child)
				}

				if Child.IsKnown() {
					// We have a known value
					if row == 0 {
						dt.AddHeader(dt.Reflect)
					}
					dt.AddRow(Child)
					continue
				}

				if dt.Reflect.Kind == reflect.Map {
					var refs []Reflect

					for col := range dt.Reflect.FieldVo.MapKeys() {
						var ChildStruct Reflect
						ChildStruct.SetByIndex(dt.Reflect, Child, col, EndPointPath{})
						if dt.Debug {
							_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", ChildStruct)
						}

						if !ChildStruct.IsExported {
							_, _ = fmt.Fprintf(os.Stderr, "WARNING: Field '%s' type not exported: Type %s\n", ChildStruct.FieldName, ChildStruct.Kind.String())
							continue
						}
						refs = append(refs, ChildStruct)
					}

					if row == 0 {
						dt.AddHeader(refs...)
					} else {
						dt.AddRow(refs...)
					}
				}
			}
			break
		}

		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported (%s): Type %s\n",
			dt.Reflect.FieldName, dt.Reflect.DataStructure.DataTableId, dt.Reflect.Kind.String())
	}

	return *dt
}

func (dt *DataTable) AddHeader(headers ...Reflect) {
	for range Only.Once {
		for _, header := range headers {
			name := valueTypes.PointToName(header.DataStructure.PointName)
			if header.DataStructure.PointUnit != "" {
				name += "\n" + header.DataStructure.PointUnit
			}
			dt.Headers = append(dt.Headers, name)
		}
	}
}

func (dt *DataTable) AddRow(refs ...Reflect) {
	for range Only.Once {
		if dt.Data == nil {
			dt.Data = make([][]Reflect, 0)
		}

		var row []Reflect
		row = append(row, refs...)
		// for _, ref := range refs {
		// }
		dt.Data = append(dt.Data, row)
	}
}

func (dt *DataTable) GetRow(row int) []Reflect {
	return dt.Data[row]
}

func (dt *DataTable) Get() [][]Reflect {
	return dt.Data
}
