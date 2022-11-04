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
	"sort"
	"strings"
	"time"
)


type DataMap struct {
	Map        map[string]*DataEntries
	Table      output.Table
	DataTables output.Tables
	Order      []string
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

		timestamp := valueTypes.SetDateTimeValue(time.Now().Round(5 * time.Minute))

		// Iterate over all available fields and read the tag values
		var Parent GoStruct.Reflect
		Parent.SetByFieldName(endpoint.ResponseRef(), endpoint.ResponseRef(), name.Index(0))
		Parent.DataStructure.Endpoint = name.Copy()
		Current := GoStruct.FindStart("ResultData", Parent, Parent, name)
		name = Current.DataStructure.Endpoint.Copy()

		var tp GoStruct.DataStructures
		tp.Debug = false
		tp.ShowEmpty = true
		tp.GetPointTags(&Parent, Current, name)

		// Convert to DataMap
		for _, f := range tp.DataMap {
			if f.PointIgnore {
				continue
			}

			if f.PointName == "" {
				f.PointName = valueTypes.PointToName(f.PointId)
			}

			if f.PointParentId == "" {
				if parentDeviceId != "" {
					f.PointParentId = parentDeviceId
				} else {
					f.PointParentId = "virtual"
				}
			}

			// fmt.Printf("DEBUG: StructToPoints(): %s / %s\n", f.Endpoint, f.PointId)
			if f.PointNameDateFormat == "" {
				f.PointNameDateFormat = valueTypes.DateTimeLayout
			}
			uvs, _, ok := valueTypes.AnyToUnitValue(f.Value, f.PointUnit, f.PointValueType, f.PointNameDateFormat)
			if !ok {
				continue
			}

			var when valueTypes.DateTime
			if !f.PointTimestamp.IsZero() {
				when = valueTypes.SetDateTimeValue(f.PointTimestamp)
			} else {
				when = timestamp
			}

			var point Point
			p := GetPoint(f.Endpoint.String() + "." + f.PointId)
			if p == nil {
				// No point found. Create one.
				point = CreatePoint(parentDeviceId, valueTypes.SetPointIdString(f.PointId), f.PointName, f.PointGroupName, uvs.Unit(), uvs.Type(), f.PointUpdateFreq)
			} else {
				point = *p
			}
			point.UpdateFreq = f.PointUpdateFreq
			point.SetName(f.PointName)

			// Add arrays as multiple entries.
			if len(uvs) > 1 {
				// @TODO - Think about adding in arrays of values OR just marshal arrays into JSON.
				res := valueTypes.SizeOfArrayLength(uvs)
				for i, uv := range uvs {
					dm.AddPointUnitValue(JoinWithDots(res, valueTypes.DateTimeLayoutDay, f.Endpoint, i), f.PointParentId, point, when, uv)
				}
				continue
			}

			if len(uvs) == 0 {
				fmt.Printf("OOOPS - UVS is nil for %s\n", f.PointId)
				continue
			}
			dm.AddPointUnitValue(f.Endpoint.String(), f.PointParentId, point, when, uvs[0])

			if f.PointAliasTo != "" {
				f.PointId = f.PointAliasTo
				dm.AddPointUnitValue(f.Endpoint.String(), f.PointParentId, point, when, uvs[0])
			}
		}

		// Create data tables.
		dm.Table = dm.CreateEndPointResultTable()
		dm.Table.SetTitle("EndPoint Data: %s", name)
		dm.Table.SetFilePrefix("%s_%s", endpoint.GetArea(), endpoint.GetName())
		dm.Table.SetJson([]byte(endpoint.GetJsonData(false)))
		dm.Table.SetRaw([]byte(endpoint.GetJsonData(true)))

		dm.DataTables = output.NewTables()
		dm.DataTables = tp.DataTables.GetDataMergedTables(timestamp, parentDeviceId)
		dm.DataTables = tp.DataTables.GetDataTables()
	}
	return *dm
}

func (dm *DataMap) AddPointUnitValue(endpoint string, parentDeviceId string, point Point, date valueTypes.DateTime, uv valueTypes.UnitValue) {

	for range Only.Once {
		if uv.Unit() != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.Unit())
			point.Unit = uv.Unit()
		}

		de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
		dm.Add(de)
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
			dm.Order = append(dm.Order, point)

			if Points.Exists(point) {
				fmt.Printf("EXISTS: %s\n", point)
			}
			Points.Add(*(de.Entries[de.Len()-1].Point))
		}
	}
}

func (dm *DataMap) Add(de DataEntry) {
	for range Only.Once {
		// fmt.Printf("DEBUG DataMap.Add() %s - Value(%s):'%s' Parent:'%s'\n", de.FullId(), de.Point.ValueType, de.Value, de.Parent)
		endpoint := de.FullId()
		de.Index = len(dm.Order)

		if dm.Map[endpoint] == nil {
			dm.Map[endpoint] = &DataEntries{ Entries: []DataEntry{} }
		}
		entries := dm.Map[endpoint]
		if entries.Add(de) == nil {
			break
		}
		dm.Order = append(dm.Order, endpoint)

		if Points.Exists(endpoint) {
			fmt.Printf("EXISTS: %s\n", endpoint)
		}
		Points.Add(*de.Point)
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

func (dm *DataMap) CreateEndPointResultTable() output.Table {
	table := output.NewTable(
		"Date",
		"Point Id",
		"Value",
		"Unit",
		"Unit Type",
		"Group Name",
		"Description",
		"Update Freq",
	)

	for range Only.Once {
		for _, p := range dm.Sort() {
			// _ = table.SetHeader(
			// 	"Date",
			// 	"Point Id",
			// 	"Value",
			// 	"Unit",
			// 	"Unit Type",
			// 	"Group Name",
			// 	"Description",
			// 	"Update Freq",
			// )

			entries := dm.Map[p].Entries
			for _, de := range entries {
				if de.Hide {
					continue
				}

				_ = table.AddRow(de.Date.Format(valueTypes.DateTimeLayout),
					p,
					de.Value.String(),
					de.Point.Unit,
					de.Point.ValueType,
					de.Point.GroupName,
					de.Point.Description,
					de.Point.UpdateFreq,
				)
			}
		}
	}

	return table
}

func (dm *DataMap) AddAny(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, value interface{}, unit string, Type string, timeSpan string) {

	for range Only.Once {
		var point Point
		p := GetPoint(parentDeviceId + "." + pid.String())
		if p == nil {
			// No point found. Create one.
			point = CreatePoint(parentDeviceId, pid, name, groupName, unit, Type, timeSpan)
		} else {
			point = *p
		}

		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, unit, Type, valueTypes.DateTimeLayout)
		if !ok {
			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
				endpoint, parentDeviceId, pid, name, date, value)
			break
		}
		if isNil {
			point.ValueType += "(NIL)"
		}

		for _, uv := range uvs {
			if uv.Unit() != point.Unit {
				fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, uv.ValueFloat(), uv.Unit())
				point.Unit = uv.Unit()
			}

			var parent ParentDevice
			parent.Set(parentDeviceId)
			point.Parents.Add(parent)

			de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
			de.Point = &point
			dm.Add(de)
		}
	}
}

func (dm *DataMap) AddUnitValue(endpoint string, parentDeviceId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, uv valueTypes.UnitValue, timeSpan string) {

	for range Only.Once {
		var point Point
		p := GetPoint(parentDeviceId + "." + pid.String())
		if p == nil {
			// No point found. Create one.
			point = CreatePoint(parentDeviceId, pid, name, groupName, uv.Unit(), uv.Type(), timeSpan)
		} else {
			point = *p
		}

		if uv.Unit() != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.Unit())
			point.Unit = uv.Unit()
		}

		var parent ParentDevice
		parent.Set(parentDeviceId)
		point.Parents.Add(parent)

		de := CreatePointDataEntry(endpoint, parentDeviceId, point, date, uv)
		de.Point = &point
		dm.Add(de)
	}
}


func CreatePointDataEntry(endpoint string, parentDeviceId string, point Point, dateTime valueTypes.DateTime, uv valueTypes.UnitValue) DataEntry {
	var ret DataEntry
	for range Only.Once {
		ret = DataEntry {
			EndPoint:   endpoint,
			Point:      &point,
			Parent:     NewParentDevice(parentDeviceId),
			Date:       dateTime,
			Value:      uv,
			Valid:      true,
			Hide:       false,
			Index:      0,
		}
	}

	return ret
}

func CreatePoint(parentId string, pid valueTypes.PointId, name string, groupName string, unit string, Type string, timeSpan string) Point {
	var point Point
	for range Only.Once {
		if name == "" {
			name = pid.PointToName()
		}

		var parent ParentDevice
		parent.Set(parentId)
		var parents ParentDevices
		parents.Add(parent)

		point = Point {
			Parents:     parents,
			Id:          pid,
			GroupName:   groupName,
			Description: name,
			Unit:        unit,
			UpdateFreq:  timeSpan,
			ValueType:   Type,
			Valid:       true,
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
