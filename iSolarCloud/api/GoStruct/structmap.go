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


type StructMapOptions struct {
	StartAt        string
	Name           EndPointPath
	TimeStamp      time.Time
	Debug          bool
	AddUnexported  bool // Add unexported values, (default: false).
	AddUnsupported bool // Add unsupported types, (default: false).
	AddInvalid     bool // Add invalid types, (default: false).
	AddEmpty       bool // Add empty values, (default: false).
	AddNil         bool // Add nil values, (default: false).
}

type StructMap struct {
	StructMapOptions

	SortOn                string
	Start                 *Reflect
	Map                   map[string]*Reflect
	TableMap              map[string]*Reflect
	_Timestamp            valueTypes.DateTime
}

func (sm *StructMap) PrintDebug(format string, args ...interface{})  {
	if sm.Debug { _, _ = fmt.Fprintf(os.Stderr, format, args...) }
}

func (sm *StructMap) InitScan(current interface{}, options StructMapOptions) {
	for range Only.Once {
		sm.StructMapOptions = options

		if sm.StructMapOptions.Name.IsZero() {
			sm.StructMapOptions.Name = NewEndPointPath(reflection.GetName("", current))
		}

		sm._Timestamp = valueTypes.SetDateTimeValue(sm.StructMapOptions.TimeStamp)
		if sm._Timestamp.IsZero() {
			sm._Timestamp = valueTypes.SetDateTimeValue(time.Now())
		}

		var Parent Reflect
		Parent.Init(current, current, sm.StructMapOptions.Name)
		Parent.SetEndPointPath(sm.StructMapOptions.Name.Copy())

		sm.PrintDebug("InitScan() EndPoint: '%s' Current: %s\n", sm.StructMapOptions.Name, sm.Start)

		sm.Scan(&Parent, &Parent)
		// fmt.Printf("sm.Current size: %v\n", unsafe.Sizeof(sm.Current))
		// fmt.Printf("sm.Map size: %v\n", unsafe.Sizeof(sm.Map))
	}
}

func (sm *StructMap) Scan(Parent *Reflect, Current *Reflect) *StructMap {
	for range Only.Once {
		sm.PrintDebug("# Scan() Current: %s\n", Current)

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
			Current.Init(Current, ref2, *Current.EndPointPath())
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
				sm.ScanSlice(Parent, Current)

			case reflect.Map:
				sm.ScanMap(Parent, Current)

			case reflect.Struct:
				sm.ScanStruct(Parent, Current)

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
			case reflect.UnsafePointer:
				sm.PrintDebug("Scan() IsUnsupported: %s\n", Current)
				sm.IsUnsupported(Current)

			case reflect.Interface:
				sm.PrintDebug("Scan() IsUnsupported: %s\n", Current)
				sm.IsUnsupported(Current)

			case reflect.Invalid:
				sm.PrintDebug("Scan() IsInvalid: %s\n", Current)
				sm.IsInvalid(Current)

			default:
				if Current.IsGoStruct() {
					break
				}
				sm.PrintDebug("Scan() IsUnsupported: %s\n", Current)
				sm.IsUnsupported(Current)
		}
	}

	return sm
}

func (sm *StructMap) ScanMap(Parent *Reflect, Current *Reflect) bool {
	var ok bool

	for range Only.Once {
		if Current.Kind != reflect.Map {
			break
		}
		sm.PrintDebug("# ScanMap() Current: %s\n", Current)

		if sm.Process(Current) {
			break
		}

		ok = true
		for index, key := range Current.FieldVo.MapKeys() {
			var Child Reflect
			Child.SetByIndex(Parent, Current, index, key)
			sm.PrintDebug("# ScanMap().SetByIndex() Child: %s\n", Child)

			if sm.Process(&Child) {
				continue
			}

			sm.Scan(Current, &Child)
		}
	}

	return ok
}

func (sm *StructMap) ScanSlice(Parent *Reflect, Current *Reflect) bool {
	var ok bool

	for range Only.Once {
		if Current.Kind != reflect.Slice {
			break
		}
		sm.PrintDebug("# ScanSlice() Current: %s\n", Current)

		if sm.Process(Current) {
			break
		}

		ok = true
		for index := 0; index < Current.Length; index++ {
			var Child Reflect
			Child.SetByIndex(Parent, Current, index, reflect.Value{})
			sm.PrintDebug("# ScanSlice().SetByIndex() Child: %s\n", Child)

			if sm.Process(&Child) {
				continue
			}

			sm.Scan(Current, &Child)
		}
	}

	return ok
}

func (sm *StructMap) ScanStruct(Parent *Reflect, Current *Reflect) bool {
	var ok bool

	for range Only.Once {
		if Current.Kind != reflect.Struct {
			break
		}
		sm.PrintDebug("# ScanStruct() Current: %s\n", Current)

		if sm.Process(Current) {
			break
		}

		ok = true
		for index := 0; index < Current.Length; index++ {
			var Child Reflect
			Child.SetByIndex(Parent, Current, index, reflect.Value{})
			sm.PrintDebug("# ScanStruct().SetByIndex() Child: %s\n", Child)

			if sm.Process(&Child) {
				continue
			}

			sm.Scan(Current, &Child)
		}
	}

	return ok
}

func (sm *StructMap) Process(Child *Reflect) bool {
	var ok bool

	for range Only.Once {
		sm.PrintDebug("Check() Child: %s\n", Child)

		if Child.Kind == reflect.Invalid {
			ok = true
			break
		}

		if Child.IsGoStruct() {
			// 	sm.SaveGoStructOptions(Child)
			break
		}

		// 	// @TODO - sigh - A real hack. We're trying to avoid updating the direct child of the table.
		// 	// if !Child.CurrentReflect.DataStructure.DataTable {
		// 	// 	if Child.SetGoStructOptions() {
		// 	// 		sm.IsTable(sm.GoStructOptionCurrent)
		// 	// 	}
		// 	// }
		// }
		//
		// if !Child.CurrentReflect.DataStructure.DataTable {
		// 	if Child.SetGoStructOptions(sm.GoStructOptions, sm.GoStructOptionCurrent) {
		// 		sm.IsTable(sm.GoStructOptionCurrent)
		// 		// Child.DataStructure.UpdateTags(Child.CurrentReflect, Child)
		// 		Child.SetPointId()
		// 		// Never exit.
		// 	}
		// }

		if Child.FieldName == sm.StructMapOptions.StartAt {
			// fmt.Printf("# SEARCHING(FOUND): %s\n", Child)
			Child.IsStart = true
			Child.ParentReflect = Child
			sm.Start = Child
			sm.StructMapOptions.StartAt = ""
			sm.PrintDebug("DEBUG: sm.Start.EndPointPath(): %s\n", sm.Start.EndPointPath())
			// @TODO - FIX THIS UP - need to check if the point_id is not this...
			if sm.Start.EndPointPath().Last() == "result_data" {
				sm.Start.EndPointPath().PopLast() // Remove "ResultData" from end.
				sm.PrintDebug("DEBUG: sm.Start.EndPointPath().PopLast(): %s\n", sm.Start.EndPointPath())
			}

			sm.PrintDebug( "DEBUG: sm.StructMapOptions.Name1: %s\n", sm.StructMapOptions.Name)
			sm.StructMapOptions.Name = sm.Start.EndPointPath().Copy()
			sm.PrintDebug("DEBUG: sm.StructMapOptions.Name2: %s\n", sm.StructMapOptions.Name)
			break
		}

		if sm.Start == nil {
			ok = Child.IsKnown()
			break	// Wait until we've started.
		}

		// IsUnexported - will add to map only if sm.AddEmpty == true
		ok = sm.IsUnexported(Child)
		if ok {
			break
		}

		// IsNil - will add to map only if sm.AddEmpty == true
		ok = sm.IsNil(Child)
		if ok {
			break
		}

		// IsEmpty - will add to map only if sm.AddEmpty == true
		ok = sm.IsEmpty(Child)
		if ok {
			break
		}

		// IsPointArrayFlatten -
		ok = sm.IsPointArrayFlatten(Child)
		if ok {
			// Will flatten a type - don't need to parse children.
			ok = true
			break
		}

		// IsPointSplitOn -
		ok = sm.IsPointSplitOn(Child)
		if ok {
			// Will split a type - don't need to parse children as it only applies to known types.
			break
		}

		// IsTable -
		ok = sm.IsTable(Child)
		if ok {
			// Never exit with true, but break - we still want to parse the children.
			break
		}

		// This will add any type that we know about
		// Basically all valueTypes and builtins, except those we want to parse - map, array, slice or struct.
		ok = Child.IsKnown()
		if ok {
			sm.Add(Child)
			break
		}
	}

	return ok
}

func (sm *StructMap) Add(Current *Reflect)  {
	for range Only.Once {
		if Current.IsGoStruct() {
			break
		}

		if sm.StructMapOptions.StartAt != "" {
			sm.PrintDebug("# Add(STILL SEARCHING): %s\n", Current)
			break
		}

		if sm.Map == nil {
			sm.Map = make(map[string]*Reflect)
		}
		name := Current.EndPointPath().String()
		sm.Map[name] = Current
		sm.PrintDebug("\t- Add() Current: %s\n", Current)
	}
}

func (sm *StructMap) AddTable(Current *Reflect)  {
	for range Only.Once {
		if Current.IsGoStruct() {
			break
		}

		if sm.StructMapOptions.StartAt != "" {
			sm.PrintDebug("# Add(STILL SEARCHING): %s\n", Current)
			break
		}

		if sm.TableMap == nil {
			sm.TableMap = make(map[string]*Reflect)
		}
		name := Current.Name()
		sm.TableMap[name] = Current
		sm.PrintDebug("\t- Add() Current: %s\n", Current)
	}
}

func (sm *StructMap) IsNil(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		yes = Current.PointIgnoreChildIfFromNil()	// true - Current is nil.

		if sm.AddNil {
			sm.Add(Current)
			_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field type is nil: %s\n", Current)
			break
		}
	}
	return yes
}

func (sm *StructMap) IsEmpty(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		if Current.Length == valueTypes.IgnoreLength {
			break
		}

		if Current.Length > 0 {
			break
		}

		if Current.IsPointIgnoreZero() {
			break
		}

		yes = true
		if sm.AddEmpty {
			sm.Add(Current)
			_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field type is nil: %s\n", Current)
			break
		}
	}
	return yes
}

func (sm *StructMap) IsUnexported(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		if Current.IsExported {
			break
		}

		yes = true
		if sm.AddUnexported {
			sm.Add(Current)
			_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field type is not exported: %s\n", Current)
		}
	}
	return yes
}

func (sm *StructMap) IsUnsupported(Current *Reflect) {
	for range Only.Once {
		if !sm.AddUnsupported {
			break
		}
		sm.Add(Current)
		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field type is unsupported: %s\n", Current)
	}
}

func (sm *StructMap) IsInvalid(Current *Reflect) {
	for range Only.Once {
		if !sm.AddInvalid {
			break
		}
		sm.Add(Current)
		_, _ = fmt.Fprintf(os.Stderr,"WARNING: Field type is invalid: %s\n", Current)
	}
}

func (sm *StructMap) IsTable(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		// Primary table type, gets added to main map, but not table map.
		if Current.IsTable() {
			// fmt.Printf("Current[%s] - Current.IsTable\n", Current.FieldPath)
			if Current.DataStructure.PointUnit == "" {
				Current.DataStructure.PointUnit = "--"
			}
			if Current.DataStructure.PointValueType == "" {
				Current.DataStructure.PointValueType = "Table"
			}
			if Current.DataStructure.ValueType == "" {
				Current.DataStructure.ValueType = "Table"
			}
			if Current.DataStructure.PointGroupName == "" {
				Current.DataStructure.PointGroupName = Current.FieldName
			}
			// if Current.DataStructure.PointName != "" {
			// 	Current.DataStructure.PointName = Current.FieldName + " " + Current.DataStructure.PointName
			// }

			// @TODO - Consider standardizing points to a known format.
			// queryUserCurveTemplateData is a good example of the structure.

			sm.PrintDebug("\t- IsTable(): %s\n", Current)
			sm.AddTable(Current)
			yes = Current.IsKnown()

			// We want to flatten a slice down to EG "[1, 2, 3]"
			uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
			// uv = Current.AsJson()
			Current.Value.Set(valueTypes.SetUnitValueString(uv, Current.DataStructure.PointUnit, Current.DataStructure.PointValueType))
			sm.Add(Current)
			break
		}

		// Check for a table child.
		ok, iterate := Current.IsTableChild()
		if !ok {
			break
		}
		// fmt.Printf("Current[%s] IsTableChild - iterate:%d kind:%s\n", Current.FieldPath, iterate, Current.Kind.String())

		// @TODO - Should we instead, recurse back up the parent tree to ensure we aren't a member of another table?
		sm.PrintDebug("\t- IsTable() Child: %s\n", Current)
		Current.DataStructure.DataTableChild = true
		if Current.DataStructure.PointUnit == "" {
			Current.DataStructure.PointUnit = "--"
		}
		if Current.DataStructure.PointValueType == "" {
			Current.DataStructure.PointValueType = "Table"
		}
		if Current.DataStructure.ValueType == "" {
			Current.DataStructure.ValueType = "Table"
		}
		if Current.DataStructure.PointGroupName == "" {
			Current.DataStructure.PointGroupName = Current.CurrentReflect.DataStructure.PointGroupName
		}

		// We want to flatten a slice down to EG "[1, 2, 3]"
		yes = Current.IsKnown()
		// First order types get flattened.
		switch Current.Kind {
			case reflect.Struct:
				if yes {
					break
				}
				if iterate == 1 {
					sm.Add(Current)
					break
				}

			case reflect.Array:
				fallthrough
			case reflect.Slice:
				if yes {
					break
				}
				if iterate == 1 {
					if Current.Length <= 1 {
						uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
						Current.Value.Set(valueTypes.SetUnitValueString(uv, Current.DataStructure.PointUnit, Current.DataStructure.PointValueType))
					}
					sm.Add(Current)
					break
				}

			case reflect.Map:
				if yes {
					break
				}
				if iterate == 1 {
					if Current.Length <= 1 {
						uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
						Current.Value.Set(valueTypes.SetUnitValueString(uv, Current.DataStructure.PointUnit, Current.DataStructure.PointValueType)) }
					sm.Add(Current)
					break
				}
		}
	}
	return yes
}

func (sm *StructMap) IsPointArrayFlatten(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		for range Only.Once {
			if Current.IsPointArrayFlatten() {
				yes = true
				break
			}

			yes = false
		}
		if !yes {
			break
		}
		sm.PrintDebug("\t- PointArrayFlatten() Child: %s\n", Current)

		// // @TODO - Here we need to convert the type to json, but it's not that simple.
		// // We need to scan the structure and resolve some of the tags.
		// // Maybe we do this in the table scanner...
		// var tp DataStructures
		// tp.Debug = false
		// tp.ShowEmpty = false
		// tp.GetPointTags(Parent, Current, EndPointPath{})
		//
		// fmt.Printf("TP: %v\n", tp)
		// uv := valueTypes.AnyToValueString(tp, 0, Current.DataStructure.PointNameDateFormat)

		// We want to flatten a slice down to EG "[1, 2, 3]"
		if Current.DataStructure.PointUnit == "" {
			Current.DataStructure.PointUnit = "--"
		}
		if Current.DataStructure.PointValueType == "" {
			Current.DataStructure.PointValueType = "--"
		}
		if Current.DataStructure.ValueType == "" {
			Current.DataStructure.ValueType = "--"
		}
		if Current.DataStructure.PointGroupName == "" {
			Current.DataStructure.PointGroupName = Current.FieldName
		}
		// if Current.DataStructure.PointName != "" {
		// 	Current.DataStructure.PointName = Current.FieldName + " " + Current.DataStructure.PointName
		// }

		// @TODO - Consider providing a "collapse" of an array/slice/map where:
		// the array elements end up becoming elements within a UV array.

		// Current.Value, Current.IsNil, Current.IsOk = valueTypes.AnyToUnitValue(
		// Current.InterfaceValue, Current.DataStructure.PointUnit,
		// Current.DataStructure.PointValueType, Current.DataStructure.PointNameDateFormat)
		uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
		Current.Value.Set(valueTypes.SetUnitValueString(uv, Current.DataStructure.PointUnit, Current.DataStructure.PointValueType))
		sm.Add(Current)
	}
	return yes
}

func (sm *StructMap) IsPointSplitOn(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		if Current.DataStructure.PointSplitOn == "" {
			break
		}
		sm.PrintDebug("\t- PointSplitOn() Child: %s\n", Current)

		// We want to split a string into separate points - currently only handles string types.
		// @TODO - Fix this up! - Use PointSplitType to define a particular type.
		soSplit := strings.Split(Current.Value.First().String(), Current.DataStructure.PointSplitOn)
		Current.SetValue(soSplit)
		for soI, soV := range soSplit {
			Child := Current.Copy()
			if Child.DataStructure.PointValueReplace != "" {
				soV = strings.ReplaceAll(soV, Child.DataStructure.PointValueReplace, Child.DataStructure.PointValueReplaceWith)
			}
			Child.SetValue(soV)
			Child.EndPointPath().Append(valueTypes.PrintInt(2, soI))
			if Child.DataStructure.PointUnit == "" {
				Child.DataStructure.PointUnit = "--"
			}
			if Child.DataStructure.PointValueType == "" {
				Child.DataStructure.PointValueType = "Split"
			}
			if Current.DataStructure.ValueType == "" {
				Current.DataStructure.ValueType = "Split"
			}
			if Child.DataStructure.PointGroupName == "" {
				Child.DataStructure.PointGroupName = Child.FieldName
			}
			// if Child.DataStructure.PointName != "" {
			// 	Child.DataStructure.PointName = Child.FieldName + " " + Child.DataStructure.PointName
			// }
			sm.Add(&Child)
		}
		yes = true
	}
	return yes
}


func (sm *StructMap) GetTables() StructTables {
	var ret StructTables

	for range Only.Once {
		st := sm.GetResultTableData()
		ret = append(ret, st)

		names := sm.GetTableNames()
		for _, name := range names {
			st = sm.GetTableData(name)
			if !st.IsValid {
				continue
			}
			ret = append(ret, st)
		}

	}

	return ret
}

func (sm *StructMap) GetTableNames() []string {
	var ret []string

	for range Only.Once {
		for name := range sm.TableMap {
			ret = append(ret, name)
		}
	}

	return ret
}

func (sm *StructMap) GetTableData(name string) StructTable {
	var ret StructTable

	for range Only.Once {
		var ok bool
		if ret.Current, ok = sm.TableMap[name]; !ok {
			break
		}

		ret.Name = ret.Current.Name()
		ret.IsValid = true

		rows, cols := ret.Current.CountChildren()
		// fmt.Printf("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ret.Current.FieldPath, ret.Current.Kind, rows, cols)
		sm.PrintDebug("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ret.Current.FieldPath, ret.Current.Kind, rows, cols)
		if ret.Current.IsPointIgnore() {
			break
		}

		if rows == 0 {
			break
		}

		if ret.Current.DataStructure.DataTableIndex {
			ret.ShowIndex = ret.Current.DataStructure.DataTableIndex
		}

		ret.IndexTitle = "Index"
		if ret.Current.DataStructure.DataTableIndexTitle != "" {
			ret.IndexTitle = ret.Current.DataStructure.DataTableIndexTitle
		}

		var refs ReflectArray
		for row, Child := range ret.Current.ChildReflect {
			sm.PrintDebug("GetTableData() row[%d]: %s\n", row, Child)
			if Child.IsPointIgnore() {
				continue
			}

			var refRow ReflectArrayRow

			for col, ChildStruct := range Child.ChildReflect {
				sm.PrintDebug("GetTableData() cell[%d][%d]: %s\n", row, col, Child)
				if ChildStruct.IsPointIgnore() {
					continue
				}

				// Make sure we have a valid sort column name.
				if ret.Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
					ret.SortOn = ChildStruct.DataStructure.PointName
				}

				if ChildStruct.IsKnown() {
					refRow = append(refRow, ChildStruct)
					continue
				}
				refRow = append(refRow, ChildStruct)
			}

			if len(refRow) > 0 {
				refs = refs.AddRow(refRow...)
			}
		}

		if !ret.Current.DataStructure.DataTablePivot {
			ret.Reflects = refs
			ret.AddHeader(ret.Reflects[0]...)
			break
		}

		// Handle table pivots here.
		for row := 0; row < len(refs[0]); row++ {
			var refRow ReflectArrayRow
			for col := 0; col < len(refs); col++ {
				refRow = append(refRow, refs[col][row])
			}
			if len(refRow) > 0 {
				ret.Reflects = ret.Reflects.AddRow(refRow...)
			}
		}
		ret.AddHeader(ret.Reflects[0]...)
	}

	return ret
}

func (sm *StructMap) GetResultTableData() StructTable {
	var ret StructTable

	for range Only.Once {
		fmt.Printf("NOT IMPLEMENTED YET.\n")
		break
		var sorted []string
		// @TODO - Add in other column sorting options here.
		for name := range sm.Map {
			sorted = append(sorted, name)
		}

		ret.Name = "Results"
		ret.IsValid = true

		for _, name := range sorted {
			Current := sm.Map[name]

			rows, cols := Current.CountChildren()
			// fmt.Printf("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ret.Current.FieldPath, ret.Current.Kind, rows, cols)
			sm.PrintDebug("GetResultTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, Current.FieldPath, Current.Kind, rows, cols)
			if Current.IsPointIgnore() {
				break
			}

			// if len(refs) > 0 {
			// 	ret.AddRow(refs...)
			// }

			// for row, Child := range Current.ChildReflect {
			// 	fmt.Printf("[%s]%s - Known:%t Current:%d / Child:%d\n", Child.FieldPath, ret.Current.Kind,
			// 		Child.IsKnown(),
			// 		len(ret.Current.ChildReflect),
			// 		len(Child.ChildReflect))
			// 	if sm.Debug {
			// 		_, _ = fmt.Fprintf(os.Stderr, "GetResultTableData() row[%d]: %s\n", row, Child)
			// 	}
			// 	if Child.IsPointIgnore() {
			// 		continue
			// 	}
			//
			// 	var refs []*Reflect
			//
			// 	for col, ChildStruct := range Child.ChildReflect {
			// 		// fmt.Printf("[%s]%s - Known:%t Current:%d / Child:%d / ChildStruct:%d\n", ChildStruct.FieldPath, ret.Current.Kind,
			// 		// 	ChildStruct.IsKnown(),
			// 		// 	len(ret.Current.ChildReflect),
			// 		// 	len(Child.ChildReflect),
			// 		// 	len(ChildStruct.ChildReflect))
			// 		if sm.Debug {
			// 			_, _ = fmt.Fprintf(os.Stderr, "GetResultTableData() cell[%d][%d]: %s\n", row, col, Child)
			// 		}
			// 		if ChildStruct.IsPointIgnore() {
			// 			continue
			// 		}
			//
			// 		// Make sure we have a valid sort column name.
			// 		if Current.DataStructure.DataTableSortOn != "" {
			// 			if Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
			// 				ret.SortOn = ChildStruct.DataStructure.PointName
			// 			}
			// 		}
			//
			// 		if ChildStruct.IsKnown() {
			// 			refs = append(refs, ChildStruct)
			// 			continue
			// 		}
			// 		refs = append(refs, ChildStruct)
			// 	}
			//
			// 	if len(refs) > 0 {
			// 		ret.AddRow(refs...)
			// 	}
		}
	}

	return ret
}


type StructTable struct {
	Name      string
	Current   *Reflect
	Headers   []string
	Reflects  ReflectArray
	SortOn    string
	ShowIndex bool
	IndexTitle string
	IsValid   bool
}


type ReflectArray []ReflectArrayRow
type ReflectArrayRow []*Reflect

func (ta *ReflectArray) AddRow(refs ...*Reflect) ReflectArray {
	for range Only.Once {
		if ta == nil {
			*ta = make(ReflectArray, 0)
		}

		var row ReflectArrayRow
		row = append(row, refs...)
		*ta = append(*ta, row)
	}
	return *ta
}

func (ta *ReflectArray) GetRow(row int) ReflectArrayRow {
	if row >= len(*ta) {
		return ReflectArrayRow{}
	}
	return (*ta)[row]
}


type StructTables []StructTable
type StructValues [][]valueTypes.UnitValue

func (ta *StructTable) AddHeader(headers ...*Reflect) {
	for range Only.Once {
		for index, header := range headers {
			if header == nil {
				name := fmt.Sprintf("Column %d", index)
				ta.Headers = append(ta.Headers, name)
				continue
			}

			name := valueTypes.PointToName(header.DataStructure.PointId)
			switch header.Value.Unit() {	// header.DataStructure.PointUnit {
				case "--":
					//
				case "":
					//

				default:
					// name += " (" + header.DataStructure.PointUnit + ")"
					name += " (" + header.Value.Unit() + ")"
			}
			ta.Headers = append(ta.Headers, name)
		}
	}
}

func (ta *StructTable) AddRow(refs ...*Reflect) {
	for range Only.Once {
		if ta.Reflects == nil {
			ta.Reflects = make(ReflectArray, 0)
		}

		// var row ReflectArrayRow
		// row = append(row, refs...)
		ta.Reflects = append(ta.Reflects, append(ReflectArrayRow{}, refs...))
	}
}

func (ta *StructTable) GetRow(row int) ReflectArrayRow {
	return ta.Reflects.GetRow(row)
}

func (ta *StructTable) GetHeaders() []string {

	if ta.ShowIndex {
		ret := []string{ta.IndexTitle}
		ret = append(ret, ta.Headers...)
		return ret
	}
	return ta.Headers
}

func (ta *StructTable) Get() ReflectArray {
	return ta.Reflects
}

func (ta *StructTable) GetValues() StructValues {
	ret := make(StructValues, 0)

	for range Only.Once {
		if !ta.IsValid {
			break
		}

		for rowIndex := range ta.Reflects {
			// fmt.Printf("ROW[%d] - size:%d\n", rowIndex, len(ta.Reflects[rowIndex]))
			var data []valueTypes.UnitValue

			// size := len(ta.Reflects[rowIndex])
			// if size == 1 {
			// 	Current := ta.Reflects[rowIndex][0]
			// 	for _, value := range Current.Value {
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Date"
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Point Id"
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Value"
			// 		data = append(data, valueTypes.UnitValue{StringValue: value.Unit()})	// "Unit"
			// 		data = append(data, valueTypes.UnitValue{StringValue: value.TypeValue})	// "Unit Type"
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Group Name"
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Description"
			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Update Freq"
			// 		ret = append(ret, data)
			// 	}
			// 	continue
			// }

			if ta.ShowIndex {
				vi := valueTypes.SetUnitValueInteger(int64(rowIndex), ta.IndexTitle, "")
				data = append(data, vi)
			}

			for colIndex, col := range ta.Reflects[rowIndex] {
				value := ta.Reflects[rowIndex][colIndex].Value

				if col.IsUnknown() {
					dateFormat := col.DataStructure.PointNameDateFormat
					if dateFormat == "" {
						dateFormat = valueTypes.DateTimeLayout
					}
					value, _, _ = valueTypes.AnyToUnitValue(col.Value, col.DataStructure.PointUnit,
						col.DataStructure.PointValueType, dateFormat)
				}

				for _, v := range value {
					data = append(data, v)
				}
			}
			ret = append(ret, data)
		}

		// @TODO - Add sorting capability here.
	}

	return ret
}


// func (sm *StructMap) GetMergedTableData(name string) StructTable {
// 	var ret StructTable
//
// 	for range Only.Once {
// 		var ok bool
// 		if ret.Current, ok = sm.TableMap[name]; !ok {
// 			break
// 		}
//
// 		ret.Name = ret.Current.Name()
// 		ret.IsValid = true
//
// 		if ret.Current.Kind == reflect.Struct {
// 			if sm.Debug {
// 				_, _ = fmt.Fprintf(os.Stderr,"GetTableArray(%s) - STRUCT\n", name)
// 			}
//
// 			if ret.Current.IsPointIgnore() {
// 				continue
// 			}
//
// 			fmt.Printf("[%s]STRUCT - Current:%d\n", ret.Current.FieldPath,
// 				len(ret.Current.ChildReflect),
// 			)
//
// 			var refs []*Reflect
// 			for _, Child := range ret.Current.ChildReflect {
// 				if Child.IsPointIgnore() {
// 					continue
// 				}
//
// 				if ret.Current.DataStructure.DataTableSortOn != "" {
// 					// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
// 					if ret.Current.DataStructure.DataTableSortOn == Child.FieldName {
// 						ret.SortOn = Child.DataStructure.PointName
// 						// dt.Reflect.DataStructure.DataTableSortOn = ""
// 					}
// 				}
//
// 				fmt.Printf("[%s]STRUCT - Known:%t Current:%d / Child:%d\n", Child.FieldPath,
// 					Child.IsKnown(),
// 					len(ret.Current.ChildReflect),
// 					len(Child.ChildReflect),
// 				)
// 				if Child.IsKnown() {
// 					refs = append(refs, Child)
// 					continue
// 				}
//
// 				refs = append(refs, Child)
// 			}
//
// 			if len(refs) > 0 {
// 				ret.AddRow(refs...)
// 			}
// 			break
// 		}
//
// 		if ret.Current.Kind == reflect.Map {
// 			if sm.Debug {
// 				_, _ = fmt.Fprintf(os.Stderr,"GetTableArray(%s) - MAP\n", name)
// 			}
//
// 			if ret.Current.IsPointIgnore() {
// 				continue
// 			}
//
// 			fmt.Printf("[%s]STRUCT - Current:%d\n", ret.Current.FieldPath,
// 				len(ret.Current.ChildReflect),
// 			)
//
// 			var refs []*Reflect
// 			for _, Child := range ret.Current.ChildReflect {
// 				if Child.IsPointIgnore() {
// 					continue
// 				}
//
// 				if ret.Current.DataStructure.DataTableSortOn != "" {
// 					// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
// 					if ret.Current.DataStructure.DataTableSortOn == Child.FieldName {
// 						ret.SortOn = Child.DataStructure.PointName
// 						// dt.Reflect.DataStructure.DataTableSortOn = ""
// 					}
// 				}
//
// 				fmt.Printf("[%s]STRUCT - Known:%t Current:%d / Child:%d\n", Child.FieldPath,
// 					Child.IsKnown(),
// 					len(ret.Current.ChildReflect),
// 					len(Child.ChildReflect),
// 				)
// 				if Child.IsKnown() {
// 					refs = append(refs, Child)
// 					continue
// 				}
// 				refs = append(refs, Child)
// 			}
//
// 			if len(refs) > 0 {
// 				ret.AddRow(refs...)
// 			}
// 			break
// 		}
//
// 		_, _ = fmt.Fprintf(os.Stderr,"GetTableArray(%s) - ARRAY\n", name)
// 		for row, Child := range ret.Current.ChildReflect {
// 			if sm.Debug {
// 				_, _ = fmt.Fprintf(os.Stderr,"GetTableArray() Child[%d]: %s\n", row, Child)
// 			}
// 			if Child.IsPointIgnore() {
// 				continue
// 			}
//
// 			fmt.Printf("[%s]SLICE - Current:%d / Child:%d\n", Child.FieldPath,
// 				len(ret.Current.ChildReflect),
// 				len(Child.ChildReflect),
// 			)
//
// 			var refs []*Reflect
// 			for _, ChildStruct := range Child.ChildReflect {
// 				if ChildStruct.IsPointIgnore() {
// 					continue
// 				}
//
// 				if ret.Current.DataStructure.DataTableSortOn != "" {
// 					// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
// 					if ret.Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
// 						ret.SortOn = ChildStruct.DataStructure.PointName
// 						// dt.Reflect.DataStructure.DataTableSortOn = ""
// 					}
// 				}
//
// 				fmt.Printf("[%s]STRUCT - Known:%t Current:%d / Child:%d / ChildStruct:%d\n", Child.FieldPath,
// 					Child.IsKnown(),
// 					len(ret.Current.ChildReflect),
// 					len(Child.ChildReflect),
// 					len(ChildStruct.ChildReflect),
// 				)
// 				if ChildStruct.IsKnown() {
// 					refs = append(refs, ChildStruct)
// 					continue
// 				}
//
// 				refs = append(refs, ChildStruct)
// 			}
//
// 			if len(refs) > 0 {
// 				ret.AddRow(refs...)
// 			}
// 			break
// 		}
// 	}
//
// 	return ret
// }

// func (sm *StructMap) SaveGoStructOptions(Child *Reflect) bool {
// 	var yes bool
//
// 	for range Only.Once {
// 		var ds DataTags
// 		ds = Child.DataStructure
// 		ds.Json = ""
// 		ds.PointId = ""
// 		ds.PointName = ""
// 		ds.ValueType = ""
// 		ds.ValueKind = ""
// 		ds.Endpoint.Clear()
//
// 		sm.GoStructOptions = &ds
// 		sm.GoStructOptionCurrent = Child.CurrentReflect.CurrentReflect		// @TODO - Need to sort out this mess.
// 		if Child.IsTable() {
// 			sm.GoStructOptionCurrent = Child.CurrentReflect.CurrentReflect
// 		}
// 		Child.CurrentReflect.CurrentReflect.GoStruct = &ds
// 	}
//
// 	return yes
// }
//
// func (sm *StructMap) ClearGoStructOptions() {
// 	sm.GoStructOptions = nil
// 	sm.GoStructOptionCurrent = nil
// }
