package GoStruct


// -------------------------------------------------------------------------------- //
// From struct_reflect.go
//
// func (r *Reflect) ConvertToState() {
// 	for range Only.Once {
// 		var ok bool
//
// 		switch {
// 			case r.Value.First().ValueFloat() > 0:
// 				ok = true
// 			case r.Value.First().ValueFloat() < 0:
// 				ok = true
// 			case r.Value.First().ValueInt() > 0:
// 				ok = true
// 			case r.Value.First().ValueInt() < 0:
// 				ok = true
// 		}
// 		r.Value = valueTypes.SetUnitValueBool(ok)
// 	}
// }
//
// func (r *Reflect) SetValues(values ...interface{}) {
// 	for range Only.Once {
// 		r.InterfaceValue = values
// 		var uvs valueTypes.UnitValues
// 		for _, value := range values {
// 			var uvs2 valueTypes.UnitValues
// 			uvs2, r.IsNil, r.IsOk = valueTypes.AnyToUnitValue(
// 				value, "", r.DataStructure.PointUnit,
// 				r.DataStructure.PointValueType, r.DataStructure.PointNameDateFormat)
// 			uvs.AddUnitValues(uvs2)
// 		}
// 	}
// }


// func FindStart(fieldName string, Parent *Reflect, Current *Reflect, name EndPointPath) *Reflect {
// 	var ret Reflect
//
// 	for range Only.Once {
// 		if Current.Kind == reflect.Pointer {
// 			// Special case:
// 			// We're going to change the pointer to a proper object reference.
// 			if Current.IsNil {
// 				break
// 			}
// 			ref2 := Current.ValueOf.Elem().Interface()
// 			if valueTypes.IsNil(ref2) {
// 				break
// 			}
// 			Current.Init(Current, ref2, Current.DataStructure.Endpoint)
// 			if Current.IsNil {
// 				break
// 			}
// 			// DO NOT BREAK!
// 			// KEEP FIRST!
// 		}
//
// 		if Current.Kind == reflect.Struct {
// 			// Iterate over all available fields and read the tag value
// 			for index := 0; index < Current.Length; index++ {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, reflect.Value{})
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Struct {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		if Current.Kind == reflect.Slice {
// 			// Iterate over all available fields and read the tag value
// 			for index := 0; index < Current.Length; index++ {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, reflect.Value{})
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Slice {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		if Current.Kind == reflect.Map {
// 			// Iterate over all available fields and read the tag value
// 			for index, key := range Current.FieldVo.MapKeys() {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, key)
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Map {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported: Type %s\n", Current.FieldName, Current.Kind.String())
// 	}
//
// 	return &ret
// }

// -------------------------------------------------------------------------------- //
// From structmap.go
//
// func (sm *StructMap) GetTables() StructTables {
// 	var ret StructTables
//
// 	for range Only.Once {
// 		st := sm.GetResultTableData()
// 		ret = append(ret, st)
//
// 		names := sm.GetTableNames()
// 		for _, name := range names {
// 			st = sm.GetTableData(name)
// 			if !st.IsValid {
// 				continue
// 			}
// 			ret = append(ret, st)
// 		}
//
// 	}
//
// 	return ret
// }
//
// func (sm *StructMap) GetTableData(name string) *StructTable {
// 	var ret StructTable
//
// 	for range Only.Once {
// 		if current, ok := sm.TableMap[name]; ok {
// 			sm.Error = ret.Process(name, current)
// 			_, _ = ret.CreateTable()
// 			break
// 		}
//
// 		sm.PrintDebug("GetTableData(%s) - UNKNOWN path:%s type:%s\n", name, ret.Current.FieldPath, ret.Current.Kind)
// 	}
//
// 	// for range Only.Once {
// 	// 	var ok bool
// 	// 	if ret.Current, ok = sm.TableMap[name]; !ok {
// 	// 		break
// 	// 	}
// 	//
// 	// 	ret.Name = ret.Current.Name()
// 	// 	ret.IsValid = true
// 	//
// 	// 	if ret.Current.DataStructure.DataTableIndex {
// 	// 		ret.ShowIndex = ret.Current.DataStructure.DataTableIndex
// 	// 	}
// 	//
// 	// 	ret.IndexTitle = "Index"
// 	// 	if ret.Current.DataStructure.DataTableIndexTitle != "" {
// 	// 		ret.IndexTitle = ret.Current.DataStructure.DataTableIndexTitle
// 	// 	}
// 	//
// 	//
// 	// 	ret.Rows, ret.Cols = ret.Current.CountChildren()
// 	// 	var isPivot bool
// 	// 	if ret.Current.DataStructure.DataTablePivot {
// 	// 		isPivot = true
// 	// 	}
// 	// 	if ret.Cols <= 1 {
// 	// 		isPivot = true
// 	// 	}
// 	//
// 	// 	// if rows == 0 {
// 	// 	// 	// var refs ReflectArray
// 	// 	// 	for row, Child := range ret.Current.Value.Range(true) {
// 	// 	// 		fmt.Printf("GetTableData() row[%d]: %s - %s == %s\n", row, Child, Child.ValueKey(), Child.StringValue)
// 	// 	//
// 	// 	// 		// var refRow ReflectArrayRow
// 	// 	// 		// refRow = append(refRow, ChildStruct)
// 	// 	// 		//
// 	// 	// 		// if len(refRow) > 0 {
// 	// 	// 		// 	refs = refs.AddRow(refRow...)
// 	// 	// 		// 	continue
// 	// 	// 		// }
// 	// 	//
// 	// 	// 		// Single column.
// 	// 	// 		ret.ShowIndex = true
// 	// 	// 		// refs = refs.AddRow(Child)
// 	// 	// 	}
// 	// 	// 	break
// 	// 	// }
// 	//
// 	// 	sm.PrintDebug("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ret.Current.FieldPath, ret.Current.Kind, ret.Rows, ret.Cols)
// 	// 	if ret.Current.IsPointIgnore() {
// 	// 		break
// 	// 	}
// 	//
// 	// 	var refs ReflectArray
// 	// 	for row, Child := range ret.Current.ChildReflect {
// 	// 		sm.PrintDebug("GetTableData() row[%d]: %s\n", row, Child)
// 	// 		if Child.IsPointIgnore() {
// 	// 			continue
// 	// 		}
// 	//
// 	// 		var refRow ReflectArrayRow
// 	//
// 	// 		for col, ChildStruct := range Child.ChildReflect {
// 	// 			sm.PrintDebug("GetTableData() cell[%d][%d]: %s\n", row, col, Child)
// 	// 			if ChildStruct.IsPointIgnore() {
// 	// 				continue
// 	// 			}
// 	//
// 	// 			// Make sure we have a valid sort column name.
// 	// 			if ret.Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
// 	// 				ret.SortOn = ChildStruct.DataStructure.PointName
// 	// 			}
// 	//
// 	// 			if ChildStruct.IsKnown() {
// 	// 				refRow = append(refRow, ChildStruct)
// 	// 				continue
// 	// 			}
// 	// 			refRow = append(refRow, ChildStruct)
// 	// 		}
// 	//
// 	// 		if len(refRow) > 0 {
// 	// 			if ret.ActualCols < len(refRow) {
// 	// 				ret.ActualCols = len(refRow)
// 	// 			}
// 	// 			refs = refs.AddRow(refRow...)
// 	// 			continue
// 	// 		}
// 	//
// 	// 		// Single column.
// 	// 		ret.ShowIndex = true
// 	// 		if Child.IsPointIgnore() {
// 	// 			continue
// 	// 		}
// 	// 		if ret.ActualCols < len(refRow) {
// 	// 			ret.ActualCols = len(refRow)
// 	// 		}
// 	// 		refs = refs.AddRow(Child)
// 	// 	}
// 	//
// 	// 	if refs == nil {
// 	// 		break
// 	// 	}
// 	//
// 	// 	if !isPivot {
// 	// 		ret.Reflects = refs
// 	// 		// ret.AddHeader(ret.Reflects[0]...)
// 	// 		break
// 	// 	}
// 	//
// 	// 	// Handle table pivots here.
// 	// 	for row := 0; row < len(refs[0]); row++ {
// 	// 		var refRow ReflectArrayRow
// 	// 		for col := 0; col < len(refs); col++ {
// 	// 			refRow = append(refRow, refs[col][row])
// 	// 		}
// 	// 		if len(refRow) > 0 {
// 	// 			ret.Reflects = ret.Reflects.AddRow(refRow...)
// 	// 		}
// 	// 	}
// 	// }
//
// 	return &ret
// }
//
// func (sm *StructMap) GetResultTableData() StructTable {
// 	var ret StructTable
//
// 	fmt.Printf("NOT IMPLEMENTED YET.\n")
// 	// for range Only.Once {
// 	// 	var sorted []string
// 	// 	// @TODO - Add in other column sorting options here.
// 	// 	for name := range sm.Map {
// 	// 		sorted = append(sorted, name)
// 	// 	}
// 	//
// 	// 	ret.Name = "Results"
// 	// 	ret.IsValid = true
// 	//
// 	// 	for _, name := range sorted {
// 	// 		Current := sm.Map[name]
// 	//
// 	// 		rows, cols := Current.CountChildren()
// 	// 		// fmt.Printf("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ret.Current.FieldPath, ret.Current.Kind, rows, cols)
// 	// 		sm.PrintDebug("GetResultTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, Current.FieldPath, Current.Kind, rows, cols)
// 	// 		if Current.IsPointIgnore() {
// 	// 			break
// 	// 		}
// 	//
// 	// 		// if len(refs) > 0 {
// 	// 		// 	ret.AddRow(refs...)
// 	// 		// }
// 	//
// 	// 		// for row, Child := range Current.ChildReflect {
// 	// 		// 	fmt.Printf("[%s]%s - Known:%t Current:%d / Child:%d\n", Child.FieldPath, ret.Current.Kind,
// 	// 		// 		Child.IsKnown(),
// 	// 		// 		len(ret.Current.ChildReflect),
// 	// 		// 		len(Child.ChildReflect))
// 	// 		// 	if sm.Debug {
// 	// 		// 		_, _ = fmt.Fprintf(os.Stderr, "GetResultTableData() row[%d]: %s\n", row, Child)
// 	// 		// 	}
// 	// 		// 	if Child.IsPointIgnore() {
// 	// 		// 		continue
// 	// 		// 	}
// 	// 		//
// 	// 		// 	var refs []*Reflect
// 	// 		//
// 	// 		// 	for col, ChildStruct := range Child.ChildReflect {
// 	// 		// 		// fmt.Printf("[%s]%s - Known:%t Current:%d / Child:%d / ChildStruct:%d\n", ChildStruct.FieldPath, ret.Current.Kind,
// 	// 		// 		// 	ChildStruct.IsKnown(),
// 	// 		// 		// 	len(ret.Current.ChildReflect),
// 	// 		// 		// 	len(Child.ChildReflect),
// 	// 		// 		// 	len(ChildStruct.ChildReflect))
// 	// 		// 		if sm.Debug {
// 	// 		// 			_, _ = fmt.Fprintf(os.Stderr, "GetResultTableData() cell[%d][%d]: %s\n", row, col, Child)
// 	// 		// 		}
// 	// 		// 		if ChildStruct.IsPointIgnore() {
// 	// 		// 			continue
// 	// 		// 		}
// 	// 		//
// 	// 		// 		// Make sure we have a valid sort column name.
// 	// 		// 		if Current.DataStructure.DataTableSortOn != "" {
// 	// 		// 			if Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
// 	// 		// 				ret.SortOn = ChildStruct.DataStructure.PointName
// 	// 		// 			}
// 	// 		// 		}
// 	// 		//
// 	// 		// 		if ChildStruct.IsKnown() {
// 	// 		// 			refs = append(refs, ChildStruct)
// 	// 		// 			continue
// 	// 		// 		}
// 	// 		// 		refs = append(refs, ChildStruct)
// 	// 		// 	}
// 	// 		//
// 	// 		// 	if len(refs) > 0 {
// 	// 		// 		ret.AddRow(refs...)
// 	// 		// 	}
// 	// 	}
// 	// }
//
// 	return ret
// }
//
// func (sm *StructMap) GetTableNames() []string {
// 	var ret []string
// 	for name := range sm.TableMap {
// 		ret = append(ret, name)
// 	}
// 	return ret
// }
//
//
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

