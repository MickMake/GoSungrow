package GoStruct

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
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

	SortOn     string
	Start      *Reflect
	Map        map[string]*Reflect
	TableMap   map[string]*Reflect
	VirtualMap map[string]*Reflect
	_Timestamp valueTypes.DateTime

	Error      error
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

		sm.Map = make(map[string]*Reflect)
		sm.VirtualMap = make(map[string]*Reflect)
		sm.TableMap = make(map[string]*Reflect)

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

		ok = true	// ok = Current.IsKnown()
		for index, key := range Current.FieldVo.MapKeys() {
			var Child Reflect
			Child.SetByIndex(Parent, Current, index, key)
			sm.PrintDebug("# ScanMap().SetByIndex() Child: %s\n", Child)

			// if strings.Contains(Child.FieldPath.String(), "DevTypeDefinition") {
			// 	for index2, key2 := range Child.FieldVo.MapKeys() {
			// 		var Child2 Reflect
			// 		Child2.SetByIndex(Current, &Child, index2, key2)
			// 		sm.Add(&Child2)
			// 	}
			// 	fmt.Println("")
			// }

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

		ok = true	// ok = Current.IsKnown()
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

		ok = true	// ok = Current.IsKnown()
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
		// fmt.Printf("[%s]\n", Child.FieldPath.String())
		sm.PrintDebug("Check() Child: %s\n", Child)
		if Child.Kind == reflect.Invalid {
			ok = true
			break
		}

		if Child.FieldName == sm.StructMapOptions.StartAt {
			// fmt.Printf("# SEARCHING(FOUND): %s\n", Child)
			Child.IsStart = true
			Child.ParentReflect = Child
			sm.Start = Child
			sm.StructMapOptions.StartAt = ""
			sm.PrintDebug("DEBUG: sm.Start.EndPointPath(): %s\n", sm.Start.EndPointPath())
			if Child.Kind == reflect.Array || Child.Kind == reflect.Slice {
				Child.SetGoStructOptions(2)
				Child.DataStructure.UpdateTags(Child.CurrentReflect, Child)
			}
			// @TODO - FIX THIS UP - need to check if the point_id is not this...
			if sm.Start.EndPointPath().Last() == "result_data" {
				// if sm.Start.FieldPath.Last() == sm.StructMapOptions.StartAt {
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

		if Child.IsGoStruct() {
			// 	sm.SaveGoStructOptions(Child)
			break
		}

		if Child.DataStructure.PointTimestamp.IsZero() {
			if sm._Timestamp.IsZero() {
				Child.DataStructure.PointTimestamp = time.Now()
			} else {
				Child.DataStructure.PointTimestamp = sm.TimeStamp
			}
		}


		// Start processing the Child here on...
		if sm.Exists(Child) {
			// Already processed.
			break
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

func (sm *StructMap) Exists(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		if Current.IsGoStruct() {
			break
		}

		if sm.Start == nil {
			break
		}

		if sm.Map == nil {
			break
		}

		name := Current.EndPointPath().String()
		if _, yes = sm.Map[name]; yes {
			sm.PrintDebug("\t- Exists() Current EXISTS: %s\n", Current)
			break
		}
	}
	return yes
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

		sm.AddVirtual(Current)

		if sm.Map == nil {
			sm.Map = make(map[string]*Reflect)
		}
		name := Current.EndPointPath().String()
		if _, ok := sm.Map[name]; ok {
			sm.PrintDebug("\t- Add() Current EXISTS: %s\n", Current)
			break
		}
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
			sm.PrintDebug("# AddTable(STILL SEARCHING): %s\n", Current)
			break
		}

		if sm.TableMap == nil {
			sm.TableMap = make(map[string]*Reflect)
		}
		name := Current.Name()
		if _, ok := sm.TableMap[name]; ok {
			sm.PrintDebug("\t- AddTable() Current: %s\n", Current)
			break
		}
		sm.TableMap[name] = Current
		sm.PrintDebug("\t- AddTable() Current: %s\n", Current)
	}
}

func (sm *StructMap) AddVirtual(Current *Reflect)  {
	for range Only.Once {
		if Current.IsGoStruct() {
			break
		}

		if sm.StructMapOptions.StartAt != "" {
			sm.PrintDebug("# AddVirtual(STILL SEARCHING): %s\n", Current)
			break
		}

		if !Current.DataStructure.PointVirtual {
			break
		}

		if sm.VirtualMap == nil {
			sm.VirtualMap = make(map[string]*Reflect)
		}
		name := Current.Name()
		if _, ok := sm.VirtualMap[name]; ok {
			sm.PrintDebug("\t- AddVirtual() Current: %s\n", Current)
			break
		}
		sm.VirtualMap[name] = Current
		sm.PrintDebug("\t- AddVirtual() Current: %s\n", Current)
	}
}

func (sm *StructMap) IsNil(Current *Reflect) bool {
	var yes bool
	for range Only.Once {
		yes = Current.PointIgnoreIfChildFromNil()	// true - Current is nil.

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
			// uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
			// // uv = Current.AsJson()
			// Current.Value.Set(uv, "", Current.DataStructure.PointUnit, Current.DataStructure.PointValueType)
			sm.Add(Current)
			break
		}

		// Check for a table child.
		// ok, iterate := Current.IsTableChild()
		ok, _ := Current.IsTableChild()
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
		sm.Add(Current)

		// First order types get flattened.
		// switch Current.Kind {
		// 	case reflect.Struct:
		// 		if yes {
		// 			break
		// 		}
		// 		if iterate == 1 {
		// 			sm.Add(Current)
		// 			break
		// 		}
		//
		// 	case reflect.Array:
		// 		fallthrough
		// 	case reflect.Slice:
		// 		if yes {
		// 			break
		// 		}
		// 		if iterate == 1 {
		// 			if Current.Length <= 1 {
		// 				uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
		// 				Current.Value.AddString(uv, "", Current.DataStructure.PointUnit, Current.DataStructure.PointValueType)
		// 			}
		// 			sm.Add(Current)
		// 			break
		// 		}
		//
		// 	case reflect.Map:
		// 		if yes {
		// 			break
		// 		}
		// 		if iterate == 1 {
		// 			if Current.Length <= 1 {
		// 				uv := valueTypes.AnyToValueString(Current.InterfaceValue, 0, Current.DataStructure.PointNameDateFormat)
		// 				Current.Value.AddString(uv, "", Current.DataStructure.PointUnit, Current.DataStructure.PointValueType)
		// 			}
		// 			sm.Add(Current)
		// 			break
		// 		}
		// }
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
		Current.Value.AddString("", Current.DataStructure.PointUnit, Current.DataStructure.PointValueType, uv)
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
