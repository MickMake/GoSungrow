package GoStruct

import (
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
	"reflect"
)


type DataTables struct {
	Map    []*DataTable
	Merge  bool
	Index  bool
}

func (dts *DataTables) Get() []*DataTable {
	// for index, val := range dt.Map {
	// 	val.Reflect.String()
	// }
	return dts.Map
}

func (dts *DataTables) GetTables(timestamp valueTypes.DateTime, parentDeviceId string) output.Tables {
	ret := make(output.Tables)
	for range Only.Once {
		for k, v := range dts.GetDataTables() {
			ret[k] = v
		}
		for k, v := range dts.GetDataMergedTables(timestamp, parentDeviceId) {
			ret[k] = v
		}
	}
	return ret
}

func (dts *DataTables) GetDataMergedTables(timestamp valueTypes.DateTime, parentDeviceId string) output.Tables {
	ret := make(output.Tables)
	for range Only.Once {
		if !dts.Merge {
			break
		}

		var data [][]interface{}
		var headers []string

		if dts.Index {
			headers = append(headers, "Date/Time", "Index")
		}

		// var names []string
		var name string
		for _, f := range dts.Get() {
			if f == nil {
				continue
			}
			if !f.Reflect.DataStructure.DataTable {
				continue
			}

			// names = append(names, f.Name)

			dt := f.GetTable()
			for index := range dt.Headers {
				headers = append(headers, dt.Headers[index])
			}
			name = dt.EndPoint.String()
			fmt.Printf("Name:%s\n", name)
			for rowIndex := range dt.Data {
				if data == nil {
					data = make([][]interface{}, 0)
				}
				for _, col := range dt.Data[rowIndex] {
					if col.DataStructure.PointNameDateFormat == "" {
						col.DataStructure.PointNameDateFormat = valueTypes.DateTimeLayout
					}
					value, _, _ := valueTypes.AnyToUnitValue(col.DataStructure.Value, col.DataStructure.PointUnit,
						col.DataStructure.PointValueType, col.DataStructure.PointNameDateFormat)

					if (len(data) - 1) < rowIndex {
						var d []interface{}
						d = append(d, value[0].String())
						data = append(data, d)
						continue
					}
					// data[rowIndex] = append(data[rowIndex], uvs[0].String())
					data[rowIndex] = append(data[rowIndex], value[0].String())
				}
			}
		}

		if dts.Index {
			for i := range data {
				var t []interface{}
				t = append(t, timestamp.String(), i)
				data[i] = append(t, data[i]...)
			}
		}

		table := output.NewTable(headers...)
		for i := range data {
			_ = table.AddRow(data[i]...)
		}
		// _ = table.SetHeader(headers...)
		table.SetTitle("Table %s-%s - (%s)", name, parentDeviceId)
		table.SetFilePrefix("%s-%s-%s", name, parentDeviceId)
		table.SetGraphFilter("")
		ret[name] = table
	}
	return ret
}

func (dts *DataTables) GetDataTables() output.Tables {
	ret := make(output.Tables)
	for range Only.Once {
		if dts.Merge {
			break
		}

		for _, f := range dts.Map {
			if !f.Reflect.DataStructure.DataTable {
				continue
			}

			dt := f.GetTable()
			table := output.NewTable(dt.Headers...)
			for rowIndex := range dt.Data {
				// fmt.Printf("ROW[%d] - %v\n", rowIndex, dt.Data[rowIndex])
				var data []interface{}
				for _, col := range dt.Data[rowIndex] {
					// fmt.Printf("\tCOL[%d][%d] - %s - %s - %v\n", rowIndex, colIndex,
					// 	dt.Data[rowIndex][colIndex].FieldName,
					// 	dt.Data[rowIndex][colIndex].DataStructure.PointId,
					// 	dt.Data[rowIndex][colIndex].DataStructure.Value,
					// 	)
					if col.DataStructure.PointNameDateFormat == "" {
						col.DataStructure.PointNameDateFormat = valueTypes.DateTimeLayout
					}
					value, _, _ := valueTypes.AnyToUnitValue(col.DataStructure.Value, col.DataStructure.PointUnit,
						col.DataStructure.PointValueType, col.DataStructure.PointNameDateFormat)

					data = append(data, value[0].String())
				}
				_ = table.AddRow(data...)
			}
			title := dt.Reflect.DataStructure.DataTableTitle
			if title == "" {
				title = dt.Reflect.DataStructure.DataTableName
			}
			if title == "" {
				title = dt.Reflect.DataStructure.DataTableId
			}

			table.SetTitle("Data table %s - %s - %s", dt.Name, dt.EndPoint.String(), title)
			table.SetFilePrefix("%s-%s", dt.EndPoint.String(), dt.Reflect.DataStructure.DataTableId)
			table.SetGraphFilter("")
			table.Sort(dt.SortOn)
			ret[dt.Reflect.DataStructure.DataTableId] = table
		}
	}
	return ret
}


type DataTable struct {
	Reflect  *Reflect
	Name     string
	EndPoint EndPointPath
	SortOn   string
	Headers  []string
	Data     [][]*Reflect
	Debug    bool
}

func (dss *DataStructures) AddTable(ref *Reflect) *DataTable {
	var ret *DataTable
	for range Only.Once {
		if ref.FieldName == NameGoStruct {
			break
		}

		if dss.DataTables.Map == nil {
			dss.DataTables.Map = []*DataTable{}
		}
		ret = &DataTable {
			Reflect:  ref,
			Name:     ref.DataStructure.DataTableId,
			EndPoint: ref.DataStructure.Endpoint,
			SortOn:   ref.DataStructure.DataTableSortOn,
			Headers:  nil,
			Data:     nil,
		}
		dss.DataTables.Map = append(dss.DataTables.Map, ret)
		// dss.DataTables.SortOn = ref.DataStructure.DataTableSortOn
		dss.DataTables.Merge = ref.DataStructure.DataTableMerge
		dss.DataTables.Index = ref.DataStructure.DataTableShowIndex
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

		if dt.Reflect.Kind == reflect.Struct {
			// Handle slices here.
			for row := 0; row < dt.Reflect.Length; row++ {
				var Child Reflect
				Child.SetByIndex(*dt.Reflect, *dt.Reflect, row, reflect.Value{}, EndPointPath{})

				if Child.IsKnown() {
					if row == 0 {
						dt.AddHeader(dt.Reflect)
					}
					dt.AddRow(&Child)
					continue
				}

				ok, refs := dt.GetTableStruct2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableMap2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableSlice2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}
			}
			break
		}

		if dt.Reflect.Kind == reflect.Slice {
			// Handle slices here.
			for row := 0; row < dt.Reflect.Length; row++ {
				var Child Reflect
				Child.SetByIndex(*dt.Reflect, *dt.Reflect, row, reflect.Value{}, EndPointPath{})

				if Child.IsKnown() {
					if row == 0 {
						dt.AddHeader(dt.Reflect)
					}
					dt.AddRow(&Child)
					continue
				}

				ok, refs := dt.GetTableStruct2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableMap2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableSlice2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}
			}
			break
		}

		if dt.Reflect.Kind == reflect.Map {
			// Handle maps here.
			for row, rowMap := range dt.Reflect.FieldVo.MapKeys() {
				var Child Reflect
				Child.SetByIndex(*dt.Reflect, *dt.Reflect, row, rowMap, EndPointPath{})

				if Child.IsKnown() {
					if row == 0 {
						dt.AddHeader(dt.Reflect)
					}
					dt.AddRow(&Child)
					continue
				}

				ok, refs := dt.GetTableStruct2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableMap2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}

				ok, refs = dt.GetTableSlice2(dt.Reflect, &Child)
				if ok {
					if row == 0 {
						dt.AddHeader(refs...)
					}
					dt.AddRow(refs...)
					continue
				}
			}
			break
		}

		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported (%s): Type %s\n",
			dt.Reflect.FieldName, dt.Reflect.DataStructure.DataTableId, dt.Reflect.Kind.String())
	}

	return *dt
}

func (dt *DataTable) GetTableMap2(Parent *Reflect, Child *Reflect) (bool, []*Reflect) {
	var ok bool
	var refs []*Reflect

	for range Only.Once {
		if Child.Kind != reflect.Map {
			break
		}

		ok = true
		for col, colMap := range dt.Reflect.FieldVo.MapKeys() {
			var ChildStruct Reflect
			ChildStruct.SetByIndex(*dt.Reflect, *Child, col, colMap, EndPointPath{})

			if dt.GoStructOptions(Parent, Child, &ChildStruct) {
				continue
			}

			if dt.Reflect.DataStructure.DataTableSortOn != "" {
				// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
				if dt.Reflect.DataStructure.DataTableSortOn == ChildStruct.FieldName {
					dt.SortOn = ChildStruct.DataStructure.PointName
					dt.Reflect.DataStructure.DataTableSortOn = ""
				}
			}
			refs = append(refs, &ChildStruct)
		}
	}

	return ok, refs
}

func (dt *DataTable) GetTableSlice2(Parent *Reflect, Child *Reflect) (bool, []*Reflect) {
	var ok bool
	var refs []*Reflect

	for range Only.Once {
		if Child.Kind != reflect.Slice {
			break
		}

		ok = true
		for col := 0; col < Child.Length; col++ {
			var ChildStruct Reflect
			ChildStruct.SetByIndex(*dt.Reflect, *Child, col, reflect.Value{}, EndPointPath{})

			if dt.GoStructOptions(Parent, Child, &ChildStruct) {
				continue
			}

			if dt.Reflect.DataStructure.DataTableSortOn != "" {
				// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
				if dt.Reflect.DataStructure.DataTableSortOn == ChildStruct.FieldName {
					dt.SortOn = ChildStruct.DataStructure.PointName
					dt.Reflect.DataStructure.DataTableSortOn = ""
				}
			}
			refs = append(refs, &ChildStruct)
		}
	}

	return ok, refs
}

func (dt *DataTable) GetTableStruct2(Parent *Reflect, Child *Reflect) (bool, []*Reflect) {
	var ok bool
	var refs []*Reflect

	for range Only.Once {
		if Child.Kind != reflect.Struct {
			break
		}

		ok = true
		for col := 0; col < Child.Length; col++ {
			var ChildStruct Reflect
			ChildStruct.SetByIndex(*dt.Reflect, *Child, col, reflect.Value{}, EndPointPath{})

			if dt.GoStructOptions(Parent, Child, &ChildStruct) {
				continue
			}

			if dt.Reflect.DataStructure.DataTableSortOn != "" {
				// fmt.Printf("\tChildStruct> %s / %s\n", ChildStruct.FieldName, dt.Reflect.DataStructure.DataTableSortOn)
				if dt.Reflect.DataStructure.DataTableSortOn == ChildStruct.FieldName {
					dt.SortOn = ChildStruct.DataStructure.PointName
					dt.Reflect.DataStructure.DataTableSortOn = ""
				}
			}
			refs = append(refs, &ChildStruct)
		}
	}

	return ok, refs
}

func (dt *DataTable) GoStructOptions(Parent *Reflect, Current *Reflect, Child *Reflect) bool {
	var yes bool
	for range Only.Once {
		if Parent.Kind == reflect.Slice {
			// @TODO - Need to check here if the parent is a slice.
			// If so - then "parent" is actually Parent.
			// If not - then "parent" is actually Current.
		}

		if Child.FieldName != NameGoStruct {
			break
		}

		yes = true
	}
	return yes
}


func (dt *DataTable) AddHeader(headers ...*Reflect) {
	for range Only.Once {
		for _, header := range headers {
			name := valueTypes.PointToName(header.DataStructure.PointId)
			if header.DataStructure.PointUnit != "" {
				name += " (" + header.DataStructure.PointUnit + ")"
			}
			dt.Headers = append(dt.Headers, name)
		}
	}
}

func (dt *DataTable) AddRow(refs ...*Reflect) {
	for range Only.Once {
		if dt.Data == nil {
			dt.Data = make([][]*Reflect, 0)
		}

		var row []*Reflect
		row = append(row, refs...)
		// for _, ref := range refs {
		// }
		dt.Data = append(dt.Data, row)
	}
}

func (dt *DataTable) GetRow(row int) []*Reflect {
	return dt.Data[row]
}

func (dt *DataTable) Get() [][]*Reflect {
	return dt.Data
}

