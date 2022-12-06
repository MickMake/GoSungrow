package GoStruct

import (
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"sort"
	"strconv"
)


type StructTables map[string]*StructTable

func (sm *StructTables) GetTableNames() []string {
	var ret []string

	for range Only.Once {
		for name := range *sm {
			ret = append(ret, name)
		}
	}

	return ret
}


type StructTable struct {
	Area       string
	Name       string
	Current    *Reflect

	MapName    string
	Reflects   ReflectArray
	Values     StructValues
	Table      output.Table
	SortOn     string
	ShowIndex  bool
	IndexTitle string
	IsValid    bool
	Columns    []string
	Rows       int
	Cols       int
	ActualRows int
	ActualCols int

	Debug      bool
	Error      error
}

func (ta *StructTable) PrintDebug(format string, args ...interface{})  {
	if ta.Debug { _, _ = fmt.Fprintf(os.Stderr, format, args...) }
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
	return ta.Columns
}

func (ta *StructTable) Get() ReflectArray {
	return ta.Reflects
}

func (ta *StructTable) Process(name string, Current *Reflect) error {

	for range Only.Once {
		if Current == nil {
			ta.Error = errors.New("*Reflect is nil")
			break
		}

		ta.Current = Current
		ta.MapName = name
		ta.Name = ta.Current.Name()
		ta.IsValid = true

		if ta.Current.DataStructure.DataTableIndex {
			ta.ShowIndex = ta.Current.DataStructure.DataTableIndex
		}

		ta.IndexTitle = "Index"
		if ta.Current.DataStructure.DataTableIndexTitle != "" {
			ta.IndexTitle = ta.Current.DataStructure.DataTableIndexTitle
		}


		ta.Rows, ta.Cols = ta.Current.CountChildren()
		var isPivot bool
		if ta.Current.DataStructure.DataTablePivot {
			isPivot = true
		}
		if ta.Cols <= 1 {
			isPivot = true
		}

		// if rows == 0 {
		// 	// var refs ReflectArray
		// 	for row, Child := range ret.Current.Value.Range(true) {
		// 		fmt.Printf("GetTableData() row[%d]: %s - %s == %s\n", row, Child, Child.ValueKey(), Child.StringValue)
		//
		// 		// var refRow ReflectArrayRow
		// 		// refRow = append(refRow, ChildStruct)
		// 		//
		// 		// if len(refRow) > 0 {
		// 		// 	refs = refs.AddRow(refRow...)
		// 		// 	continue
		// 		// }
		//
		// 		// Single column.
		// 		ret.ShowIndex = true
		// 		// refs = refs.AddRow(Child)
		// 	}
		// 	break
		// }

		ta.PrintDebug("GetTableData(%s) - path:%s type:%s rows:%d cols:%d\n", name, ta.Current.FieldPath, ta.Current.Kind, ta.Rows, ta.Cols)
		if ta.Current.IsPointIgnore() {
			break
		}

		var refs ReflectArray
		for row, Child := range ta.Current.ChildReflect {
			ta.PrintDebug("GetTableData() row[%d]: %s\n", row, Child)
			if Child.IsPointIgnore() {
				continue
			}

			var refRow ReflectArrayRow

			for col, ChildStruct := range Child.ChildReflect {
				ta.PrintDebug("GetTableData() cell[%d][%d]: %s\n", row, col, Child)
				if ChildStruct.IsPointIgnore() {
					continue
				}

				// Make sure we have a valid sort column name.
				if ta.Current.DataStructure.DataTableSortOn == ChildStruct.FieldName {
					ta.SortOn = ChildStruct.DataStructure.PointName
				}

				if ChildStruct.IsKnown() {
					refRow = append(refRow, ChildStruct)
					continue
				}
				refRow = append(refRow, ChildStruct)
			}

			if len(refRow) > 0 {
				if ta.ActualCols < len(refRow) {
					ta.ActualCols = len(refRow)
				}
				refs = refs.AddRow(refRow...)
				continue
			}

			// Single column.
			ta.ShowIndex = true
			if Child.IsPointIgnore() {
				continue
			}
			if ta.ActualCols < len(refRow) {
				ta.ActualCols = len(refRow)
			}
			refs = refs.AddRow(Child)
		}

		if refs == nil {
			break
		}

		if !isPivot {
			ta.Reflects = refs
			// ret.AddHeader(ret.Reflects[0]...)
			break
		}

		// Handle table pivots here.
		for row := 0; row < len(refs[0]); row++ {
			var refRow ReflectArrayRow
			for col := 0; col < len(refs); col++ {
				refRow = append(refRow, refs[col][row])
			}
			if len(refRow) > 0 {
				ta.Reflects = ta.Reflects.AddRow(refRow...)
			}
		}
	}

	return ta.Error
}

func (ta *StructTable) GetValues() StructValues {

	for range Only.Once {
		ta.Values = make(StructValues, 0)
		if !ta.IsValid {
			break
		}

		colOrder := make(map[string]int)
		var colOrderIndex int

		var addCol = func(name string) {
			if _, ok := colOrder[name]; !ok {
				colOrder[name] = colOrderIndex
				colOrderIndex++
			}
		}

		var colName = func(sub *Reflect, value *valueTypes.UnitValue, length int) string {
			name := sub.DataStructure.PointName
			if value.ValueKey() == "" {
				if name == "" {
					name = "Column " + strconv.Itoa(length)
				}
			} else {
				name += " " + value.ValueKey()
			}
			switch value.Unit() {
			case "--":
			case "":
			default:
				if !sub.DataStructure.PointVariableUnit {
					name += " (" + value.Unit() + ")"
				}
			}
			return name
		}

		// ta.Reflects - contains the rows.
		// ta.Reflects == 0 - something wrong.
		// ta.Reflects == 1 - Single row, .

		if len(ta.Reflects) == 0 {
			// fmt.Println("len(ta.Reflects) == 0")
			// Probs an array of values.
			// cm := make(map[string][]valueTypes.UnitValue)
			// var length int
			//
			// if ta.ShowIndex {
			// 	addCol(ta.IndexTitle)
			// }
			// for _, sub := range ta.Current.Value.Range(true) {
			// 	name := sub.ValueKey()
			// 	if !ta.Current.IsPointVariableUnit() {
			// 		switch sub.Unit() {
			// 			case "--":
			// 			case "":
			// 			default:
			// 				name += " (" + sub.Unit() + ")"
			// 		}
			// 		addCol(name)
			// 	}
			//
			// 	cm[name] = sub.Range(valueTypes.SortOrder)
			// 	l := sub.Value.Length()
			// 	if l > length {
			// 		length = l
			// 	}
			// }
			//
			// for index := 0; index < length; index++ {
			// 	data := make(StructValue)
			//
			// 	if ta.ShowIndex {
			// 		vi := valueTypes.SetUnitValueInteger(int64(index), ta.IndexTitle, "")
			// 		data[ta.IndexTitle] = vi
			// 	}
			//
			// 	for name, value := range cm {
			// 		if index >= len(value) {
			// 			data[name] = valueTypes.UnitValue{}
			// 		}
			// 		data[name] = value[index]
			// 	}
			// 	ret = append(ret, data)
			// }
			//
			// ta.Columns = sortMapByValues(colOrder)

			// cm := make(map[string][]valueTypes.UnitValue)
			// cm[ta.Current.DataStructure.PointName] = ta.Current.Value.Range(valueTypes.SortOrder)
			// length := ta.Current.Value.Length()
			// addCol("Key")
			// addCol("Value")
			// // data := make(StructValue)
			// // data["Key"] = value[name]
			// // ret = append(ret, data)
			// for index := 0; index < length; index++ {
			// 	data := make(StructValue)
			//
			// 	if ta.ShowIndex {
			// 		vi := valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(index))
			// 		data[ta.IndexTitle] = vi
			// 	}
			//
			// 	for name, value := range cm {
			// 		data[name] = value[index]
			// 	}
			// 	ret = append(ret, data)
			// }
			break
		}


		if len(ta.Reflects) == 1 {
			// Probs an array of values - sw we want to pivot the data.
			cm := make(map[string][]valueTypes.UnitValue)
			var length int

			if ta.ShowIndex {
				addCol(ta.IndexTitle)
			}
			for _, sub := range ta.Reflects[0] {
				name := sub.DataStructure.PointName
				if !sub.DataStructure.PointVariableUnit {
					switch sub.Value.GetUnit() {
					case "--":
					case "":
					default:
						name += " (" + sub.Value.GetUnit() + ")"
					}
					addCol(name)
				} else {
					// addCol(name)
					// addCol("Units")
					// cm[name] = sub.Value.Range(valueTypes.SortOrder)
				}

				cm[name] = sub.Value.Range(valueTypes.SortOrder)
				l := sub.Value.Length()
				if l > length {
					length = l
				}
			}

			tin := ta.Current.GetDataTableIndexNames()
			for index := 0; index < length; index++ {
				data := make(StructValue)

				if ta.ShowIndex {
					var vi valueTypes.UnitValue
					if len(tin) > 0 {
						vi = valueTypes.SetUnitValueString(ta.IndexTitle, "", tin[index])
					} else {
						vi = valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(index))
					}
					data[ta.IndexTitle] = vi
				}

				for name, value := range cm {
					if index >= len(value) {
						data[name] = valueTypes.UnitValue{}
						continue
						// @TODO - potential issue here.
					}
					data[name] = value[index]
				}
				ta.Values = append(ta.Values, data)
			}

			ta.Columns = sortMapByValues(colOrder)
			break
		}


		if ta.ShowIndex {
			addCol(ta.IndexTitle)
		}

		for rowIndex, _ := range ta.Reflects {
			// fmt.Printf("ROW[%d] - size:%d\n", rowIndex, len(ta.Reflects[rowIndex]))
			data := make(StructValue)
			// fmt.Printf("DEBUG[0].FieldPath == %s\n", row[0].FieldPath.String())

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
				var vi valueTypes.UnitValue
				vi = valueTypes.SetUnitValueInteger(ta.IndexTitle, "", int64(rowIndex))
				data[ta.IndexTitle] = vi
			}

			for colIndex, col := range ta.Reflects[rowIndex] {
				// It's important that the values are sorted by table header.
				// This is so that the headers match with data.
				if len(col.ChildReflect) > 0 {
					// Handles
					for _, sub := range col.ChildReflect {
						for _, val := range sub.Value.Range(valueTypes.SortOrder) {
							name := colName(sub, &val, len(data))
							data[name] = val
							addCol(name)
						}
					}
					continue
				}

				if col.IsKnown() {
					value := ta.Reflects[rowIndex][colIndex].Value
					// data = append(data, value.Range(valueTypes.SortOrder)...)
					for _, val := range value.Range(valueTypes.SortOrder) {
						name := colName(col, &val, len(data))
						data[name] = val
						addCol(name)
					}
					continue
				}

				dateFormat := col.DataStructure.PointNameDateFormat
				if dateFormat == "" {
					dateFormat = valueTypes.DateTimeLayout
				}
				value, _, _ := valueTypes.AnyToUnitValue(col.Value, "", col.DataStructure.PointUnit,
					col.DataStructure.PointValueType, dateFormat)

				// data = append(data, value.Range(valueTypes.SortOrder)...)
				for _, val := range value.Range(valueTypes.SortOrder) {
					name := colName(col, &val, len(data))
					data[name] = val
					addCol(name)
				}
			}
			ta.Values = append(ta.Values, data)
		}

		ta.Columns = sortMapByValues(colOrder)

		// @TODO - Add sorting capability here.
	}

	return ta.Values
}

func (ta *StructTable) CreateTable() (output.Table, error) { // (output.Tables, GoStruct.StructValuesMap) {

	for range Only.Once {
		if ta.Current == nil {
			ta.Error = errors.New("*Reflect is nil")
			break
		}

		values := ta.GetValues()
		if (values == nil) || (len(values) == 0) {
			fmt.Printf("No data table results for '%s'\n", ta.MapName)
			break
		}

		headers := ta.GetHeaders()
		ta.Table = output.NewTable(headers...)
		for row := range values {
			var items []interface{}
			for _, col := range ta.Columns {
				items = append(items, values.GetCell(row, col).String())
			}
			ta.Error = ta.Table.AddRow(items...)
			if ta.Error != nil {
				break
			}
		}
		if ta.Error != nil {
			break
		}

		title := ta.Current.DataStructure.DataTableTitle
		if title == "" {
			title = ta.Current.DataStructure.DataTableName
		}
		if title == "" {
			title = valueTypes.PointToName(ta.Current.DataStructure.DataTableId)
		}
		// if title == "" {
		// 	title = valueTypes.PointToName(ta.Current.DataStructure.PointId)
		// }
		// dm.EndPoint.GetRequestArgNames()

		ta.Table.SetName(ta.MapName)
		if title == "" {
			ta.Table.SetTitle("DataTable %s.%s", ta.Area, ta.Name)
			ta.Table.SetFilePrefix("%s.%s", ta.Area, ta.Name)
		} else {
			ta.Table.SetTitle("DataTable %s.%s (%s)", ta.Area, ta.Name, title)
			ta.Table.SetFilePrefix("%s.%s-%s", ta.Area, ta.Name, ta.Current.DataStructure.DataTableId)
		}

		// table.Sort(td.SortOn)
		ta.Table.SetJson(nil)
		ta.Table.SetRaw(nil)

		ta.Table.SetGraphFilter("")	// @TODO - Consider setting graph options here instead of iSolarCloud/data.go:487

		// if sgd.Options.GraphRequest.TimeColumn == nil {
		// 	for _, col := range table.GetHeaders() {
		// 		val := value.GetCell(0, col)
		// 		if val.Type() == "DateTime" {
		// 			sgd.Options.GraphRequest.TimeColumn = &col
		// 			break
		// 		}
		// 	}
		// }
		//
		// if sgd.Options.GraphRequest.DataColumn == nil {
		// 	for _, col := range table.GetHeaders() {
		// 		val := value.GetCell(0, col)
		// 		if val.IsNumber() {
		// 			sgd.Options.GraphRequest.DataColumn = &col
		// 			break
		// 		}
		// 	}
		// }
		//
		// if sgd.Options.GraphRequest.ValueColumn == nil {
		// 	for _, col := range table.GetHeaders() {
		// 		val := value.GetCell(0, col)
		// 		if val.IsNumber() {
		// 			sgd.Options.GraphRequest.ValueColumn = &col
		// 			break
		// 		}
		// 	}
		// }
	}

	return ta.Table, ta.Error
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


type StructValuesMap map[string]StructValues
type StructValue map[string]valueTypes.UnitValue

type StructValues []StructValue

func (ta *StructValues) GetCell(row int, col string) valueTypes.UnitValue {
	var ret valueTypes.UnitValue
	for range Only.Once {
		if row >= len(*ta) {
			ret = valueTypes.SetUnitValueString("error", "error", "row > size")
			break
		}

		if _, ok := (*ta)[row][col]; !ok {
			ret = valueTypes.SetUnitValueString("", "", "")
			break
		}

		ret = (*ta)[row][col]
	}
	return ret
}


func sortMapByValues(data map[string]int) []string {
	var ret []string
	keys := make([]string, 0, len(data))

	for key := range data {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool{
		return data[keys[i]] < data[keys[j]]
	})

	for _, k := range keys{
		ret = append(ret, k)
	}
	return ret
}
