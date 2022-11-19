package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"fmt"
	datatable "go.pennock.tech/tabular/auto"
	"os"
	"sort"
	"strings"
	"time"
)


type DataMap struct {
	Map        map[string]*DataEntries
	List       output.Table
	Table      output.Table
	DataTables output.Tables
	// Order      []string
	Error      error
}

func NewDataMap() DataMap {
	return DataMap {
		Map: make(map[string]*DataEntries),
	}
}

func (dm *DataMap) StructToDataMap(endpoint EndPoint, parentDeviceId string, name GoStruct.EndPointPath) DataMap {
	for range Only.Once {
		epName := GoStruct.NewEndPointPath(reflection.GetName("", endpoint))
		epName.Append(name.Strings()...)
		name = epName.Copy()

		timestamp := time.Now().Round(5 * time.Minute)

		// Parse response structure.
		var sm GoStruct.StructMap
		sm.InitScan(endpoint.ResponseRef(), GoStruct.StructMapOptions {
			StartAt:        "ResultData",
			Name:           name,
			TimeStamp:      timestamp,
			Debug:          false,
			AddUnexported:  false,
			AddUnsupported: false,
			AddInvalid:     false,
			AddNil:         false,
			AddEmpty:       false,
		})

		// Convert to DataMap
		for _, Child := range sm.Map {
			pdi := parentDeviceId
			if Child.IsPointIgnore() {
				continue
			}

			var when valueTypes.DateTime
			if Child.IsPointTimestampNotZero() {
				when = valueTypes.SetDateTimeValue(Child.DataStructure.PointTimestamp)
			} else {
				when = valueTypes.SetDateTimeValue(timestamp)
			}

			if Child.DataStructure.PointDevice != "" {
				pdi = Child.DataStructure.PointDevice
			}
			dm.AddPointUnitValues(Child, pdi, when)
		}

		dm.List, dm.Error = dm.CreateResultTable(endpoint, sm, true)
		if dm.Error != nil {
			break
		}

		dm.Table, dm.Error = dm.CreateResultTable(endpoint, sm, false)
		if dm.Error != nil {
			break
		}

		dm.Error = dm.CreateDataTables(sm)
		if dm.Error != nil {
			break
		}
	}
	return *dm
}


func (dm *DataMap) CreateDataTables(sm GoStruct.StructMap) error {
	for range Only.Once {
		dm.DataTables = output.NewTables()

		names := sm.GetTableNames()
		for _, name := range names {
			td := sm.GetTableData(name)
			if !td.IsValid {
				continue
			}

			values := td.GetValues()
			headers := td.GetHeaders()
			table := output.NewTable(headers...)
			for row := range values {
				var items []interface{}
				for _, col := range td.Columns {
					items = append(items, values.GetCell(row, col))
				}
				dm.Error = table.AddRow(items...)
				if dm.Error != nil {
					break
				}
			}
			if dm.Error != nil {
				break
			}

			title := td.Current.DataStructure.DataTableTitle
			if title == "" {
				title = td.Current.DataStructure.DataTableName
			}
			if title == "" {
				title = valueTypes.PointToName(td.Current.DataStructure.DataTableId)
			}
			// if title == "" {
			// 	title = valueTypes.PointToName(td.Current.DataStructure.PointId)
			// }

			table.SetTitle("Data Table %s - %s", td.Name, title)
			if td.Current.DataStructure.DataTableId == "" {
				table.SetFilePrefix(td.Name)
			} else {
				table.SetFilePrefix("%s-%s", td.Name, td.Current.DataStructure.DataTableId)
			}
			table.SetGraphFilter("")
			table.Sort(td.SortOn)
			table.SetJson(nil)
			table.SetRaw(nil)

			dm.DataTables[name] = table
		}
	}

	return dm.Error
}

func (dm *DataMap) CreateResultTable(endpoint EndPoint, sm GoStruct.StructMap, full bool) (output.Table, error) {
	var ret output.Table
	for range Only.Once {
		ret = output.NewTable(
			"Date",
			"Point Id",
			"Value",
			// "(Value)",
			"Unit",
			"Unit Type",
			"Group Name",
			"Description",
			"Update Freq",
		)

		for _, p := range dm.Sort() {
			entries := dm.Map[p].Entries
			for _, de := range entries {
				if full {
					if de.Current.DataStructure.DataTable {
						continue	// We already have the children.
					}
				} else {
					if de.Hide {
						continue	// Ignore hidden entries.
					}
					if de.Current.DataStructure.DataTableChild {
						continue	// Ignore data table children.
					}
					// child, i := de.Current.IsTableChild()
					// fmt.Printf("%t[%d]\n", child, i)
					// if child {
					// 	if !de.Current.IsTable() {
					// 		continue
					// 	}
					// }
				}

				v := de.Value.String()
				if de.Current.IsTable() {
					v = "See table: " + de.Current.Name()
				}
				dm.Error = ret.AddRow(
					de.Date.Format(valueTypes.DateTimeLayout),
					p,
					v,
					// de.Value.String(),
					de.Point.Unit,
					de.Point.ValueType,
					de.Point.GroupName,
					de.Point.Description,
					de.Point.UpdateFreq,
				)
				if dm.Error != nil {
					break
				}
			}
		}

		var title string
		// title = sm.Name.String()
		// title = valueTypes.PointToName(sm.Start.DataStructure.DataTableId)

		// ret.SetTitle("EndPoint Data: %s.%s", endpoint.GetArea(), endpoint.GetName())
		ret.SetTitle("EndPoint Data %s - %s", sm.Name.String(), title)
		// ret.SetFilePrefix("%s_%s", endpoint.GetArea(), endpoint.GetName())
		ret.SetFilePrefix(sm.Name.String())
		ret.SetGraphFilter("")
		ret.Sort("Point Id")
		if full {
			ret.SetJson([]byte(endpoint.GetJsonData(false)))
			ret.SetRaw([]byte(endpoint.GetJsonData(true)))
		}
	}

	return ret, dm.Error
}

func (dm *DataMap) AddPointUnitValues(Current *GoStruct.Reflect, parentDeviceId string, date valueTypes.DateTime) {

	for range Only.Once {
		if parentDeviceId == "" {
			parentDeviceId = "system"
		}
		if Current.DataStructure.PointDevice != "" {
			parentDeviceId = Current.DataStructure.PointDevice
		}

		var point Point
		p := GetPoint(Current.EndPointPath().String())
		if p == nil {
			// No point found. Create one.
			point = CreatePoint(Current, parentDeviceId)
		} else {
			point = *p
		}
		point.UpdateFreq = Current.DataStructure.PointUpdateFreq
		point.SetName(Current.DataStructure.PointName)

		if Current.Value.Unit() != point.Unit {
			_, _ = fmt.Fprintf(os.Stderr,"OOOPS FP['%s'] - Point/Value unit mismatch - Point:%s != Value:%s (%f)\n",
				Current.FieldPath.String(), point.Unit, Current.Value.Unit(), Current.Value.First().ValueFloat())
			point.Unit = Current.Value.Unit()
		}

		// Add arrays as multiple entries.
		if Current.Value.Length() > 1 {
			entries := CreatePointDataEntries(Current, parentDeviceId, point, date)
			dm.Add(entries.Entries...)
			break
		}

		if Current.Value.Length() == 0 {
			_, _ = fmt.Fprintf(os.Stderr,"OOOPS FP['%s'] - UVS is nil\n", Current.FieldPath.String())
			break
		}

		de := CreatePointDataEntry(Current, parentDeviceId, point, date, *Current.Value.First())
		dm.Add(de)

		if Current.DataStructure.PointAliasTo != "" {
			Current.DataStructure.PointId = Current.DataStructure.PointAliasTo
			de = CreatePointDataEntry(Current, parentDeviceId, point, date, *Current.Value.First())
			dm.Add(de)
		}
	}
}

func (dm *DataMap) CopyPoint(refEndpoint string, endpoint string, pointId string, name string) *DataEntries {
	var ret *DataEntries
	for range Only.Once {
		var dep *DataEntries
		var ok bool
		if dep, ok = dm.Map[refEndpoint]; !ok {
			fmt.Printf("ERROR: Can't find point '%s'\n", refEndpoint)
			break
		}

		ret = dm.CopyDataEntries(*dep, endpoint, pointId, name)
	}
	return ret
}

func (dm *DataMap) CopyDataEntries(dep DataEntries, endpoint string, pointId string, name string) *DataEntries {
	var ret *DataEntries
	for range Only.Once {
		var des DataEntries
		des = dep.Copy()
		for i := range des.Entries {
			des.Entries[i].SetEndpoint(endpoint, pointId)
			des.Entries[i].SetPointName(name)
			dm.Add(des.Entries[i])
		}

		if len(des.Entries) == 0 {
			fmt.Printf("OOOPS\n")
		}
		epn := des.Entries[0].EndPoint
		ret = dm.Map[epn]
	}
	return ret
}

func (dm *DataMap) LowerUpper(lowerEntry DataEntries, upperEntry DataEntries) float64 {
	var ret float64
	for range Only.Once {
		if lowerEntry.GetFloat() > 0 {
			ret = 0 - lowerEntry.GetFloat()
			break
		}
		ret = upperEntry.GetFloat()
	}
	return ret
}

func (dm *DataMap) GetPercent(entry DataEntries, max DataEntries) float64 {
	var ret float64
	for range Only.Once {
		ret = GetPercent(entry.GetFloat(), max.GetFloat())
	}
	return ret
}

func (dm *DataMap) AppendMap(add DataMap) {
	for range Only.Once {
		if dm.Map == nil {
			dm.Map = make(map[string]*DataEntries)
		}

		for point, de := range add.Map {
			if dd, ok := dm.Map[point]; ok {
				jde, _ := json.Marshal(de)
				jdd, _ := json.Marshal(dd)
				if string(jdd) != string(jde) {
					fmt.Printf("DIFF ")
				}
				fmt.Printf("Duplicate[%s]:\n%s\n%s\n", point, jde, jdd)
				continue
			}
			dm.Map[point] = de

			if Points.Exists(point) {
				fmt.Printf("EXISTS: %s\n", point)
			}
			Points.Add(*(de.Entries[de.Len()-1].Point))
		}
	}
}

func (dm *DataMap) Add(des ...DataEntry) {
	for range Only.Once {
		for _, de := range des {
			// fmt.Printf("DEBUG DataMap.Add() %s - Value(%s):'%s' Parent:'%s'\n", de.FullId(), de.Point.ValueType, de.Value, de.Parent)
			endpoint := de.FullId()
			// de.Index = len(dm.Order)

			if dm.Map[endpoint] == nil {
				dm.Map[endpoint] = &DataEntries{Entries: []DataEntry{}}
			}
			entries := dm.Map[endpoint]
			if entries.Add(de) == nil {
				continue
			}
			// dm.Order = append(dm.Order, endpoint)

			if Points.Exists(endpoint) {
				fmt.Printf("EXISTS: %s\n", endpoint)
			}
			Points.Add(*de.Point)
		}
	}
}

func (dm *DataMap) Print() {
	for range Only.Once {
		table := datatable.New("utf8-heavy")
		table.AddHeaders(
			"Index",
			"EndPoint",

			"Id",
			"Name",
			"Unit",
			"Type",
			"Value",
			"(Value)",
			"Valid",

			"GroupName",
			"Parent Ids",
			"Parent Types",
			"Parent Codes",
		)

		// dm.Order - Produces double the amount of entries for some reason.
		i := 0
		for k := range dm.Map {
			for _, v := range dm.Map[k].Entries {
				i++

				table.AddRowItems(
					i,
					v.EndPoint,

					v.Point.Id,
					v.Point.Description,
					v.Point.Unit,
					v.Point.UpdateFreq,
					v.Value,
					v.Current.Value.First(),
					v.Point.Valid,

					v.Point.GroupName,
					v.Point.Parents.PsIds(),
					v.Point.Parents.Types(),
					v.Point.Parents.Codes(),
				)
			}
		}

		ret, _ := table.Render()
		fmt.Println(ret)
	}
}

func (dm *DataMap) Sort() []string {
	var sorted []string

	for range Only.Once {
		for p := range dm.Map {
			sorted = append(sorted, p)
		}
		sort.Strings(sorted)
	}
	return sorted
}

func (dm *DataMap) TableSort() []string {
	var sorted []string

	for range Only.Once {
		for p := range dm.DataTables {
			sorted = append(sorted, p)
		}
		sort.Strings(sorted)
	}
	return sorted
}

// func (dm *DataMap) AddAny(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, value interface{}, unit string, Type string, timeSpan string) {
//
// 	for range Only.Once {
// 		var point Point
// 		p := GetPoint(parentDeviceId + "." + pid.String())
// 		if p == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentDeviceId, pid, name, groupName, unit, Type, timeSpan)
// 		} else {
// 			point = *p
// 		}
//
// 		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, unit, Type, valueTypes.DateTimeLayout)
// 		if !ok {
// 			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
// 				endpoint, parentDeviceId, pid, name, date, value)
// 			break
// 		}
// 		if isNil {
// 			point.ValueType += "(NIL)"
// 		}
//
// 		for _, uv := range uvs {
// 			if uv.Unit() != point.Unit {
// 				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.Unit())
// 				point.Unit = uv.Unit()
// 			}
//
// 			var parent ParentDevice
// 			parent.Set(parentDeviceId)
// 			point.Parents.Add(parent)
//
// 			de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
// 			de.Point = &point
// 			dm.Add(de)
// 		}
// 	}
// }

// func (dm *DataMap) AddUnitValue(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, uv valueTypes.UnitValue, timeSpan string) {
//
// 	for range Only.Once {
// 		var point Point
// 		p := GetPoint(parentDeviceId + "." + pid.String())
// 		if p == nil {
// 			// No point found. Create one.
// 			point = CreatePoint(parentDeviceId, pid, name, groupName, uv.Unit(), uv.Type(), timeSpan)
// 		} else {
// 			point = *p
// 		}
//
// 		if uv.Unit() != point.Unit {
// 			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.Unit())
// 			point.Unit = uv.Unit()
// 		}
//
// 		var parent ParentDevice
// 		parent.Set(parentDeviceId)
// 		point.Parents.Add(parent)
//
// 		de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
// 		de.Point = &point
// 		dm.Add(de)
// 	}
// }


func CreatePointDataEntry(Current *GoStruct.Reflect, parentDeviceId string, point Point, dateTime valueTypes.DateTime, uv valueTypes.UnitValue) DataEntry {
	// CreatePointDataEntry(Current, Current.EndPointPath().String(), parentDeviceId, point, date, *Current.Value.First())
	var ret DataEntry
	for range Only.Once {
		ret = DataEntry {
			Current:    Current,
			EndPoint:   Current.EndPointPath().String(),
			Point:      &point,
			Parent:     NewParentDevice(parentDeviceId),
			Date:       dateTime,
			Value:      uv,
			Valid:      true,
			Hide:       false,
			// Index:      0,
		}
	}

	return ret
}

func CreatePointDataEntries(Current *GoStruct.Reflect, parentDeviceId string, point Point, dateTime valueTypes.DateTime) DataEntries {
	var ret DataEntries
	for range Only.Once {
		if Current.Value.Length() <= 1 {
			break
		}

		// res := valueTypes.SizeOfArrayLength(Current.Value.Length())
		// res := Current.Value.Length()
		sorted := Current.Value.Range(valueTypes.SortOrder)
		res := len(sorted)
		for i, uv := range sorted {
			epn := JoinWithDots(res, valueTypes.DateTimeLayoutDay, Current.EndPointPath().String(), i)
			ret.Entries = append(ret.Entries, DataEntry{
				Current:  Current,
				EndPoint: epn, // Current.EndPointPath().String(),
				Point:    &point,
				Parent:   NewParentDevice(parentDeviceId),
				Date:     dateTime,
				Value:    uv,
				Valid:    true,
				Hide:     false,
				// Index:      0,
			})
		}
	}

	return ret
}

func CreatePoint(Current *GoStruct.Reflect, parentDeviceId string) Point {
	// parentDeviceId string	- parentDeviceId
	// pid valueTypes.PointId	- valueTypes.SetPointIdString(Current.PointId())
	// name string				- Current.DataStructure.PointName
	// groupName string			- Current.DataStructure.PointGroupName
	// unit string				- Current.Value.Unit()
	// Type string				- Current.Value.Type()
	// timeSpan string			- Current.DataStructure.PointUpdateFreq

	var point Point
	for range Only.Once {
		pid := valueTypes.SetPointIdString(Current.PointId())
		name := Current.DataStructure.PointName
		if name == "" {
			name = pid.PointToName()
		}

		var parent ParentDevice
		parent.Set(parentDeviceId)
		var parents ParentDevices
		parents.Add(parent)

		point = Point {
			Parents:     parents,
			Id:          pid,
			GroupName:   Current.DataStructure.PointGroupName,
			Description: name,
			Unit:        Current.Value.Unit(),
			UpdateFreq:  Current.DataStructure.PointUpdateFreq,
			ValueType:   Current.Value.Type(),
			Valid:       Current.IsOk,
			States:      nil,
		}
		point.FixUnitType()
	}

	return point
}

func GetPercent(value float64, max float64) float64 {
	if max == 0 {
		return 0
	}
	return (value / max) * 100
}

func JoinWithDots(intSize int, dateFormat string, args ...interface{}) string {
	var ret string
	for range Only.Once {
		var a []string
		for _, ref := range args {
			v := valueTypes.AnyToValueString(ref, intSize, dateFormat)
			if v == "" {
				continue
			}
			v = strings.Trim(v, ".")
			a = append(a, v)
		}
		ret = strings.Join(a, ".")
	}
	return ret
}
