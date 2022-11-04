package GoStruct

import (
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"reflect"
	"strings"
	"time"
)


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
	DataTableSortOn           string

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
			ret := reflection.GetStringFrom(current, pointIgnoreIfNil)
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

		pointJson := reflection.GetJsonTag(fieldTo)
		pointId := fieldTo.Tag.Get(PointId)
		if pointId == "" {
			pointId = pointJson
		}

		pointUnit := fieldTo.Tag.Get(PointUnit)
		pointUnitFrom := fieldTo.Tag.Get(PointUnitFrom)
		pointUnitFromParent := fieldTo.Tag.Get(PointUnitFromParent)
		if pointUnitFrom != "" {
			pointUnit = reflection.GetStringFrom(current, pointUnitFrom)
		}
		if pointUnitFromParent != "" {
			pointUnit = reflection.GetStringFrom(parent, pointUnitFromParent)
		}

		pointGroupName := fieldTo.Tag.Get(PointGroupName)
		pointGroupNameFrom := fieldTo.Tag.Get(PointGroupNameFrom)
		if pointGroupNameFrom != "" {
			pointGroupName = reflection.GetStringFrom(current, pointGroupNameFrom)
		}

		pointTimestamp := time.Now()
		pointTimestampFrom := fieldTo.Tag.Get(PointTimestampFrom)
		if pointTimestampFrom != "" {
			pointTimestamp = reflection.GetTimestampFrom(current, pointTimestampFrom, valueTypes.DateTimeLayout)
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

			DataTable:          dataTable,
			DataTableId:        fieldTo.Tag.Get(DataTableId),
			DataTableName:      fieldTo.Tag.Get(DataTableName),
			DataTableTitle:     fieldTo.Tag.Get(DataTableTitle),
			DataTableMerge:     dataTableMerge,
			DataTableShowIndex: dataTableShowIndex,
			DataTableSortOn:    fieldTo.Tag.Get(DataTableSortOn),

			Value:     fieldVo.Interface(),
			ValueType: valueType,
			ValueKind: fieldVo.Kind().String(),
			Endpoint:  EndPointPath{},
		}
	}
	return *ds
}

type DataStructures struct {
	DataMap    map[string]DataStructure
	DataTables DataTables
	ShowEmpty  bool
	Debug      bool
}

func (dss *DataStructures) GetPointTags(Parent *Reflect, Current *Reflect, name EndPointPath) DataStructures {

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

func (dss *DataStructures) ProcessUnsupported(_ *Reflect, Current *Reflect, name EndPointPath) {
	for range Only.Once {
		if dss.ShowEmpty {
			dss.Add(Current.DataStructure)
			break
		}
		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field '%s' type not supported (%s): Type %s\n", Current.FieldName, name, Current.Kind.String())
	}
}

func (dss *DataStructures) ProcessSlice(Parent *Reflect, Current *Reflect, name EndPointPath) {
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
			Child.SetByIndex(*Parent, *Current, si, reflect.Value{}, name)
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
			}
			name2 := Child.DataStructure.Endpoint.Copy()
			if Current.DataStructure.PointNameFromChild != "" {
				name2 = Current.PointNameFromChild(Child, name)
			}

			if Child.IsUnknown() {
				dss.GetPointTags(Current, &Child, name2)
				continue
			}

			if dss.PointSplitOn(Current, &Child, name2) {
				continue
			}

			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) ProcessStruct(Parent *Reflect, Current *Reflect, name EndPointPath) {
	for range Only.Once {
		// Iterate over all available fields and read the tag value
		for si := 0; si < Current.Length; si++ {
			var Child Reflect
			Child.SetByIndex(*Parent, *Current, si, reflect.Value{}, name)
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
			}
			name2 := Child.DataStructure.Endpoint.Copy()
			if Current.DataStructure.PointNameFromChild != "" {
				name2 = Current.PointNameFromChild(Child, name)
			}

			if !Child.IsExported {
				_, _ = fmt.Fprintf(os.Stderr, "WARNING: Field '%s' type not exported (%s): Type %s\n", Child.FieldName, name2, Child.Kind.String())
				continue
			}

			if dss.GoStructOptions(Parent, Current, &Child, name) {
				continue
			}

			if Child.IsUnknown() {
				dss.GetPointTags(Current, &Child, name2)
				continue
			}

			if dss.PointSplitOn(Current, &Child, name2) {
				continue
			}

			// COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) ProcessMap(Parent *Reflect, Current *Reflect, name EndPointPath) {
	for range Only.Once {
		for si, sm := range Current.FieldVo.MapKeys() {
			// @TODO - Implement pointNameFromChild / pointNameFromParent.
			// @TODO - Need to look at other types, besides known types.
			var Child Reflect
			Child.SetByIndex(*Parent, *Current, si, sm, name)
			if dss.Debug {
				_, _ = fmt.Fprintf(os.Stderr,"SetByIndex() Child: %s\n", Child)
			}
			name2 := Child.DataStructure.Endpoint.Copy()
			if Current.DataStructure.PointNameFromChild != "" {
				name2 = Current.PointNameFromChild(Child, name)
			}

			if Child.IsUnknown() {
				dss.GetPointTags(Current, &Child, name2)
				continue
			}

			if dss.PointSplitOn(Current, &Child, name2) {
				continue
			}

			// COMPARE(Child.DataStructure.Endpoint, Child.DataStructure.Value, Child.Interface)
			dss.Add(Child.DataStructure)
		}
	}
}

func (dss *DataStructures) PointNameAppend(_ *Reflect, Current *Reflect, name EndPointPath) EndPointPath {
	for range Only.Once {
		if Current.DataStructure.PointNameAppend == false {
			if len(name) == 0 {
				break
			}
			name = name.PopLast()
		}
	}
	return name
}

func (dss *DataStructures) PointArrayFlatten(_ *Reflect, Current *Reflect, name EndPointPath) bool {
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

func (dss *DataStructures) PointIgnoreZero(_ *Reflect, Current *Reflect, _ EndPointPath) bool {
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

func (dss *DataStructures) PointIgnoreIfNilFromChild(Parent *Reflect, Current *Reflect, _ EndPointPath) bool {
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
		ret := reflection.GetStringFrom(Current.Interface, Parent.DataStructure.PointIgnoreIfNilFromChild)
		if ret == "" {
			yes = false
			break
		}
		yes = true
	}
	return yes
}

func (dss *DataStructures) PointSplitOn(_ *Reflect, Current *Reflect, _ EndPointPath) bool {
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

func (dss *DataStructures) GoStructOptions(Parent *Reflect, Current *Reflect, Child *Reflect, _ EndPointPath) bool {
	var yes bool
	for range Only.Once {
		if Child.FieldName != NameGoStruct {
			break
		}

		// @TODO - Need to check here if the parent is a slice.
		// If so - then "parent" is actually Parent.
		// If not - then "parent" is actually Current.

		if Child.DataStructure.DataTable {
			Parent.DataStructure.DataTable = Child.DataStructure.DataTable
		}
		if Child.DataStructure.DataTableMerge {
			Parent.DataStructure.DataTableMerge = Child.DataStructure.DataTableMerge
		}
		if Child.DataStructure.DataTableShowIndex {
			Parent.DataStructure.DataTableShowIndex = Child.DataStructure.DataTableShowIndex
		}
		if Child.DataStructure.DataTableId != "" {
			Parent.DataStructure.DataTableId = Child.DataStructure.DataTableId
		}
		if Child.DataStructure.DataTableName != "" {
			Parent.DataStructure.DataTableName = Child.DataStructure.DataTableName
		}
		if Child.DataStructure.DataTableTitle != "" {
			Parent.DataStructure.DataTableTitle = Child.DataStructure.DataTableTitle
		}
		if Child.DataStructure.DataTableSortOn != "" {
			Parent.DataStructure.DataTableSortOn = Child.DataStructure.DataTableSortOn
		}

		dss.AddTable(Parent)
		yes = true
	}
	return yes
}


func COMPARE(name EndPointPath, ref1 interface{}, ref2 interface{}) {
	t1 := fmt.Sprintf("%v", ref1)
	t2 := fmt.Sprintf("%v", ref2)
	if t1 != t2 {
		fmt.Printf("[%s] VALUE ERROR: '%s' != '%s'\n", name, t1, t2)
	}
}
