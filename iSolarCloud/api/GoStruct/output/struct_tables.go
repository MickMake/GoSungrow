package output


// type StructTable struct {
// 	Name       string
// 	Current    *Reflect
// 	Reflects   ReflectArray
// 	SortOn     string
// 	ShowIndex  bool
// 	IndexTitle string
// 	IsValid    bool
// 	Columns    []string
// 	Rows       int
// 	Cols       int
// 	ActualRows int
// 	ActualCols int
// }
//
//
// type ReflectArray []ReflectArrayRow
// type ReflectArrayRow []*Reflect
//
// func (ta *ReflectArray) AddRow(refs ...*Reflect) ReflectArray {
// 	for range Only.Once {
// 		if ta == nil {
// 			*ta = make(ReflectArray, 0)
// 		}
//
// 		var row ReflectArrayRow
// 		row = append(row, refs...)
// 		*ta = append(*ta, row)
// 	}
// 	return *ta
// }
//
// func (ta *ReflectArray) GetRow(row int) ReflectArrayRow {
// 	if row >= len(*ta) {
// 		return ReflectArrayRow{}
// 	}
// 	return (*ta)[row]
// }
//
//
// type StructTables []StructTable
// type StructValuesMap map[string]StructValues
// type StructValues []StructValue
// type StructValue map[string]valueTypes.UnitValue
//
// func (ta *StructValues) GetCell(row int, col string) valueTypes.UnitValue {
// 	var ret valueTypes.UnitValue
// 	for range Only.Once {
// 		if row >= len(*ta) {
// 			ret = valueTypes.SetUnitValueString("error", "error", "row > size")
// 			break
// 		}
//
// 		if _, ok := (*ta)[row][col]; !ok {
// 			ret = valueTypes.SetUnitValueString("", "", "")
// 			break
// 		}
//
// 		ret = (*ta)[row][col]
// 	}
// 	return ret
// }
//
// func (ta *StructTable) AddRow(refs ...*Reflect) {
// 	for range Only.Once {
// 		if ta.Reflects == nil {
// 			ta.Reflects = make(ReflectArray, 0)
// 		}
//
// 		// var row ReflectArrayRow
// 		// row = append(row, refs...)
// 		ta.Reflects = append(ta.Reflects, append(ReflectArrayRow{}, refs...))
// 	}
// }
//
// func (ta *StructTable) GetRow(row int) ReflectArrayRow {
// 	return ta.Reflects.GetRow(row)
// }
//
// func (ta *StructTable) GetHeaders() []string {
// 	return ta.Columns
// }
//
// func (ta *StructTable) Get() ReflectArray {
// 	return ta.Reflects
// }
//
// func (ta *StructTable) GetValues() StructValues {
// 	ret := make(StructValues, 0)
//
// 	for range Only.Once {
// 		if !ta.IsValid {
// 			break
// 		}
//
// 		colOrder := make(map[string]int)
// 		var colOrderIndex int
//
// 		var addCol = func(name string) {
// 			if _, ok := colOrder[name]; !ok {
// 				colOrder[name] = colOrderIndex
// 				colOrderIndex++
// 			}
// 		}
//
// 		var colName = func(sub *Reflect, value *valueTypes.UnitValue, length int) string {
// 			name := sub.DataStructure.PointName
// 			if value.ValueKey() == "" {
// 				if name == "" {
// 					name = "Column " + strconv.Itoa(length)
// 				}
// 			} else {
// 				name += " " + value.ValueKey()
// 			}
// 			switch value.Unit() {
// 			case "--":
// 			case "":
// 			default:
// 				if !sub.DataStructure.PointVariableUnit {
// 					name += " (" + value.Unit() + ")"
// 				}
// 			}
// 			return name
// 		}
//
// 		// ta.Reflects - contains the rows.
// 		// ta.Reflects == 0 - something wrong.
// 		// ta.Reflects == 1 - Single row, .
//
// 		if len(ta.Reflects) == 0 {
// 			// fmt.Println("len(ta.Reflects) == 0")
// 			// Probs an array of values.
// 			// cm := make(map[string][]valueTypes.UnitValue)
// 			// var length int
// 			//
// 			// if ta.ShowIndex {
// 			// 	addCol(ta.IndexTitle)
// 			// }
// 			// for _, sub := range ta.Current.Value.Range(true) {
// 			// 	name := sub.ValueKey()
// 			// 	if !ta.Current.IsPointVariableUnit() {
// 			// 		switch sub.Unit() {
// 			// 			case "--":
// 			// 			case "":
// 			// 			default:
// 			// 				name += " (" + sub.Unit() + ")"
// 			// 		}
// 			// 		addCol(name)
// 			// 	}
// 			//
// 			// 	cm[name] = sub.Range(valueTypes.SortOrder)
// 			// 	l := sub.Value.Length()
// 			// 	if l > length {
// 			// 		length = l
// 			// 	}
// 			// }
// 			//
// 			// for index := 0; index < length; index++ {
// 			// 	data := make(StructValue)
// 			//
// 			// 	if ta.ShowIndex {
// 			// 		vi := valueTypes.SetUnitValueInteger(int64(index), ta.IndexTitle, "")
// 			// 		data[ta.IndexTitle] = vi
// 			// 	}
// 			//
// 			// 	for name, value := range cm {
// 			// 		if index >= len(value) {
// 			// 			data[name] = valueTypes.UnitValue{}
// 			// 		}
// 			// 		data[name] = value[index]
// 			// 	}
// 			// 	ret = append(ret, data)
// 			// }
// 			//
// 			// ta.Columns = sortMapByValues(colOrder)
//
// 			// cm := make(map[string][]valueTypes.UnitValue)
// 			// cm[ta.Current.DataStructure.PointName] = ta.Current.Value.Range(valueTypes.SortOrder)
// 			// length := ta.Current.Value.Length()
// 			// addCol("Key")
// 			// addCol("Value")
// 			// // data := make(StructValue)
// 			// // data["Key"] = value[name]
// 			// // ret = append(ret, data)
// 			// for index := 0; index < length; index++ {
// 			// 	data := make(StructValue)
// 			//
// 			// 	if ta.ShowIndex {
// 			// 		vi := valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(index))
// 			// 		data[ta.IndexTitle] = vi
// 			// 	}
// 			//
// 			// 	for name, value := range cm {
// 			// 		data[name] = value[index]
// 			// 	}
// 			// 	ret = append(ret, data)
// 			// }
// 			break
// 		}
//
//
// 		if len(ta.Reflects) == 1 {
// 			// Probs an array of values - sw we want to pivot the data.
// 			cm := make(map[string][]valueTypes.UnitValue)
// 			var length int
//
// 			if ta.ShowIndex {
// 				addCol(ta.IndexTitle)
// 			}
// 			for _, sub := range ta.Reflects[0] {
// 				name := sub.DataStructure.PointName
// 				if !sub.DataStructure.PointVariableUnit {
// 					switch sub.Value.GetUnit() {
// 					case "--":
// 					case "":
// 					default:
// 						name += " (" + sub.Value.GetUnit() + ")"
// 					}
// 					addCol(name)
// 				} else {
// 					// addCol(name)
// 					// addCol("Units")
// 					// cm[name] = sub.Value.Range(valueTypes.SortOrder)
// 				}
//
// 				cm[name] = sub.Value.Range(valueTypes.SortOrder)
// 				l := sub.Value.Length()
// 				if l > length {
// 					length = l
// 				}
// 			}
//
// 			tin := ta.Current.GetDataTableIndexNames()
// 			for index := 0; index < length; index++ {
// 				data := make(StructValue)
//
// 				if ta.ShowIndex {
// 					var vi valueTypes.UnitValue
// 					if len(tin) > 0 {
// 						vi = valueTypes.SetUnitValueString(ta.IndexTitle, "", tin[index])
// 					} else {
// 						vi = valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(index))
// 					}
// 					data[ta.IndexTitle] = vi
// 				}
//
// 				for name, value := range cm {
// 					if index >= len(value) {
// 						data[name] = valueTypes.UnitValue{}
// 						continue
// 						// @TODO - potential issue here.
// 					}
// 					data[name] = value[index]
// 				}
// 				ret = append(ret, data)
// 			}
//
// 			ta.Columns = sortMapByValues(colOrder)
// 			break
// 		}
//
//
// 		if ta.ShowIndex {
// 			addCol(ta.IndexTitle)
// 		}
//
// 		for rowIndex, _ := range ta.Reflects {
// 			// fmt.Printf("ROW[%d] - size:%d\n", rowIndex, len(ta.Reflects[rowIndex]))
// 			data := make(StructValue)
// 			// fmt.Printf("DEBUG[0].FieldPath == %s\n", row[0].FieldPath.String())
//
// 			// size := len(ta.Reflects[rowIndex])
// 			// if size == 1 {
// 			// 	Current := ta.Reflects[rowIndex][0]
// 			// 	for _, value := range Current.Value {
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Date"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Point Id"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Value"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: value.Unit()})	// "Unit"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: value.TypeValue})	// "Unit Type"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Group Name"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Description"
// 			// 		data = append(data, valueTypes.UnitValue{StringValue: ""})	// "Update Freq"
// 			// 		ret = append(ret, data)
// 			// 	}
// 			// 	continue
// 			// }
//
// 			if ta.ShowIndex {
// 				var vi valueTypes.UnitValue
// 				vi = valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(rowIndex))
// 				data[ta.IndexTitle] = vi
// 			}
//
// 			for colIndex, col := range ta.Reflects[rowIndex] {
// 				// It's important that the values are sorted by table header.
// 				// This is so that the headers match with data.
// 				if len(col.ChildReflect) > 0 {
// 					// Handles
// 					for _, sub := range col.ChildReflect {
// 						for _, val := range sub.Value.Range(valueTypes.SortOrder) {
// 							name := colName(sub, &val, len(data))
// 							data[name] = val
// 							addCol(name)
// 						}
// 					}
// 					continue
// 				}
//
// 				if col.IsKnown() {
// 					value := ta.Reflects[rowIndex][colIndex].Value
// 					// data = append(data, value.Range(valueTypes.SortOrder)...)
// 					for _, val := range value.Range(valueTypes.SortOrder) {
// 						name := colName(col, &val, len(data))
// 						data[name] = val
// 						addCol(name)
// 					}
// 					continue
// 				}
//
// 				dateFormat := col.DataStructure.PointNameDateFormat
// 				if dateFormat == "" {
// 					dateFormat = valueTypes.DateTimeLayout
// 				}
// 				value, _, _ := valueTypes.AnyToUnitValue(col.Value, "", col.DataStructure.PointUnit,
// 					col.DataStructure.PointValueType, dateFormat)
//
// 				// data = append(data, value.Range(valueTypes.SortOrder)...)
// 				for _, val := range value.Range(valueTypes.SortOrder) {
// 					name := colName(col, &val, len(data))
// 					data[name] = val
// 					addCol(name)
// 				}
// 			}
// 			ret = append(ret, data)
// 		}
//
// 		ta.Columns = sortMapByValues(colOrder)
//
// 		// @TODO - Add sorting capability here.
// 	}
//
// 	return ret
// }
//
// func sortMapByValues(data map[string]int) []string {
// 	var ret []string
// 	keys := make([]string, 0, len(data))
//
// 	for key := range data {
// 		keys = append(keys, key)
// 	}
// 	sort.SliceStable(keys, func(i, j int) bool{
// 		return data[keys[i]] < data[keys[j]]
// 	})
//
// 	for _, k := range keys{
// 		ret = append(ret, k)
// 	}
// 	return ret
// }
