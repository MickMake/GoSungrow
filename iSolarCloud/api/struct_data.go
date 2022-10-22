package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"fmt"
	datatable "go.pennock.tech/tabular/auto"
	"strings"
	"time"
)


const (
	UpdateFreqInstant = "instant"
	UpdateFreq5Mins   = "5mins"
	UpdateFreqBoot    = "boot"
	UpdateFreqDay   = "daily"
	UpdateFreqMonth = "monthly"
	UpdateFreqYear  = "yearly"
	UpdateFreqTotal   = "total"
)


type DataMap struct {
	Map   map[string]*DataEntries
	Order []string
}

func NewDataMap() DataMap {
	return DataMap {
		Map: make(map[string]*DataEntries),
	}
}

func (dm *DataMap) StructToPoints(ref interface{}, endpoint string, parentId string, timestamp valueTypes.DateTime) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		// Iterate over all available fields and read the tag values
		var tp apiReflect.DataStructures
		var Ref apiReflect.Reflect
		Ref.SetByFieldName(ref, ref, "")
		tp.GetPointTags(Ref, Ref, endpoint)

		for _, f := range tp.Map {
			if f.PointIgnore {
				continue
			}

			// if strings.Contains(f.PointId, "p83095") || strings.Contains(f.PointId, "es_total_disenergy") {
			// 	fmt.Printf("F:%v\n", f)
			// 	fmt.Println("")
			// }

			if f.PointName == "" {
				f.PointName = valueTypes.PointToName(f.PointId)
			}

			if f.PointParentId == "" {
				if parentId != "" {
					f.PointParentId = parentId
				} else {
					f.PointParentId = "virtual"
				}
			}

			// fmt.Printf("DEBUG: StructToPoints(): %s / %s\n", f.Endpoint, f.PointId)
			uvs, _, ok := valueTypes.AnyToUnitValue(f.Value, f.PointUnit, f.PointValueType, valueTypes.DateTimeLayout)
			if !ok {
				continue
			}

			switch f.PointUpdateFreq {
				case "UpdateFreqInstant":
					f.PointUpdateFreq = UpdateFreqInstant
				case "UpdateFreq5Mins":
					f.PointUpdateFreq = UpdateFreq5Mins
				case "UpdateFreqBoot":
					f.PointUpdateFreq = UpdateFreqBoot
				case "UpdateFreqDay":
					f.PointUpdateFreq = UpdateFreqDay
				case "UpdateFreqMonth":
					f.PointUpdateFreq = UpdateFreqMonth
				case "UpdateFreqYear":
					f.PointUpdateFreq = UpdateFreqYear
				case "UpdateFreqTotal":
					f.PointUpdateFreq = UpdateFreqTotal
			}

			var when valueTypes.DateTime
			if !f.PointTimestamp.IsZero() {
				dt := valueTypes.SetDateTimeValue(f.PointTimestamp)
				when = *dt
			} else {
				if timestamp.IsZero() {
					dt := valueTypes.SetDateTimeValue(time.Now().Round(5 * time.Minute))
					when = *dt
				} else {
					when = timestamp
				}
			}

			var point Point
			p := GetPoint(f.Endpoint + "." + f.PointId)
			if p == nil {
				// No point found. Create one.
				point = CreatePoint(parentId, valueTypes.SetPointIdString(f.PointId), f.PointName, f.PointGroupName, uvs.Unit(), uvs.Type(), f.PointUpdateFreq)
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
			dm.AddPointUnitValue(f.Endpoint, f.PointParentId, point, when, uvs[0])

			if f.PointAliasTo != "" {
				f.PointId = f.PointAliasTo
				dm.AddPointUnitValue(f.Endpoint, f.PointParentId, point, when, uvs[0])
			}
		}
	}
}

func (dm *DataMap) AddPointUnitValue(endpoint string, parentId string, point Point, date valueTypes.DateTime, uv valueTypes.UnitValue) {

	for range Only.Once {
		if uv.Unit() != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.Unit())
			point.Unit = uv.Unit()
		}

		de := CreatePointDataEntry(endpoint, parentId, point, date, uv)
		dm.Add(de)
	}
}

func CreatePointDataEntry(endpoint string, parentId string, point Point, dateTime valueTypes.DateTime, uv valueTypes.UnitValue) DataEntry {
	var ret DataEntry
	for range Only.Once {
		ret = DataEntry {
			EndPoint:   endpoint,
			Point:      &point,
			Parent:     NewParentDevice(parentId),
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

		epn := des.Entries[0].EndPoint	// + "." + des.Entries[0].Point.Id.String()
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
		// DataEntry {
		// 	Point:      Point{
		// 		Parents:   ParentDevices{},
		// 		Id:        valueTypes.PointId{},
		// 		GroupName: "",
		// 		Name:      "",
		// 		Unit:      "",
		// 		UpdateFreq:  "",
		// 		ValueType: "",
		// 		Valid:     false,
		// 		States:    nil,
		// 	},
		// 	Date:       valueTypes.DateTime{},
		// 	EndPoint:   "",
		// 	FullId:     valueTypes.DataPoint{},
		// 	Parent:     ParentDevice{},
		// 	Value:      "",
		// 	ValueFloat: 0,
		// 	ValueBool:  false,
		// 	Index:      0,
		// 	Valid:      false,
		// 	Hide:       false,
		// }

		// fmt.Printf("DEBUG DataMap.Add() %s - Value(%s):'%s' Parent:'%s'\n", de.FullId(), de.Point.ValueType, de.Value, de.Parent)
		endpoint := de.FullId()
		de.Index = len(dm.Order)

		if dm.Map[endpoint] == nil {
			// make(
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

		for i, k := range dm.Order {
			for _, v := range dm.Map[k].Entries {
				table.AddRowItems(
					i,
					v.EndPoint,

					v.Point.Id,
					v.Point.Description,
					v.Point.Unit,
					v.Point.UpdateFreq,
					v.Value,
					v.Point.Valid,
					// fmt.Sprintf("%s\n%s\n", v.FullId, v.Value),

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
		for _, e := range args {
			v := valueTypes.TypeToString(intSize, dateFormat, e)
			if v == "" {
				continue
			}
			a = append(a, v)
		}
		ret = strings.Join(a, ".")
	}
	return ret
}


func (dm *DataMap) AddAny(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, value interface{}, unit string, Type string, timeSpan string) {

	for range Only.Once {
		var point Point
		p := GetPoint(parentId + "." + pid.String())
		if p == nil {
			// No point found. Create one.
			point = CreatePoint(parentId, pid, name, groupName, unit, Type, timeSpan)
		} else {
			point = *p
		}

		// ref := valueTypes.SetUnitValueFloat(value, point.Unit, point.ValueType)
		// if ref.Unit() != point.Unit {
		// 	fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
		// 	point.Unit = ref.Unit()
		// }

		uvs, isNil, ok := valueTypes.AnyToUnitValue(value, unit, Type, valueTypes.DateTimeLayout)
		if !ok {
			fmt.Printf("ERROR: AddAny(endpoint '%s', parentId '%s', pid '%s', name '%s', date '%s', value '%v')",
				endpoint, parentId, pid, name, date, value)
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
			parent.Set(parentId)
			point.Parents.Add(parent)

			de := CreatePointDataEntry(endpoint, parentId, point, date, uv)
			// de := CreateDataEntry(endpoint, parentId, pid, name, groupName, date, uv)
			de.Point = &point
			// de := DataEntry {
			// 	EndPoint: endpoint,
			// 	// FullId:     valueTypes.JoinDataPoint(endpoint, point.Id.String()),
			// 	Parent:   parent,
			//
			// 	Date:       date,
			// 	Point:      point,
			// 	Value:      uv.String(),
			// 	ValueFloat: uv.Value(),
			// 	ValueBool:  uv.ValueBool(),
			// 	Index:      0,
			// 	Valid:      true,
			// 	Hide:       false,
			// }
			dm.Add(de)
		}

		// for _, uv := range uvs {
		// 	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
		// 	dm.Add(de)
		// }
	}
}

func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, date valueTypes.DateTime, uv valueTypes.UnitValue, timeSpan string) {

	for range Only.Once {
		var point Point
		p := GetPoint(parentId + "." + pid.String())
		if p == nil {
			// No point found. Create one.
			point = CreatePoint(parentId, pid, name, groupName, uv.Unit(), uv.Type(), timeSpan)
		} else {
			point = *p
		}

		if uv.Unit() != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %s %s != %f %s\n", uv.String(), point.Unit, uv.ValueFloat(), uv.Unit())
			point.Unit = uv.Unit()
		}

		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		de := CreatePointDataEntry(endpoint, parentId, point, date, uv)
		de.Point = &point
		dm.Add(de)
	}
}

// func CreateDataEntry(endpoint string, parentId string, pid valueTypes.PointId, name string, groupName string, dateTime valueTypes.DateTime, uv valueTypes.UnitValue, timeSpan string) DataEntry {
// 	var ret DataEntry
// 	for range Only.Once {
// 		if name == "" {
// 			name = pid.PointToName()
// 		}
// 		point := CreatePoint(parentId, pid, name, groupName, uv.Unit(), uv.Type(), timeSpan)
//
// 		// point = &Point {
// 		// 	Parents:   de.Point.Parents,
// 		// 	Id:        pid,
// 		// 	GroupName: "alias",
// 		// 	Name:      name,
// 		// 	Unit:      de.Point.Unit,
// 		// 	UpdateFreq:  de.Point.UpdateFreq,
// 		// 	ValueType: de.Point.ValueType,
// 		// 	Valid:     true,
// 		// 	States:    de.Point.States,
// 		// }
// 		// var parent ParentDevice
// 		// parent.Set(parentId)
// 		// point.Parents.Add(parent)
// 		// point.Unit = "binary"
// 		// if point.Unit == "" {
// 		// 	point.Unit = ref.Unit()
// 		// }
// 		// point.Name = name
// 		// if point.Name == "" {
// 		// 	point.Name = pid.PointToName()
// 		// }
// 		// // if de2.Point.GroupName == "" {
// 		// // 	de2.Point.GroupName = groupName
// 		// // }
// 		// point.FixUnitType()
// 		// point.Valid = true
//
// 		var parent ParentDevice
// 		parent.Set(parentId)
//
// 		ret = DataEntry {
// 			Point:      point,
// 			Date:       dateTime,
// 			EndPoint:   endpoint,
// 			Parent:     parent, // ParentDevice{},
// 			Value:      uv,
// 			// Value:      uv.String(),
// 			// ValueFloat: uv.ValueFloat(),
// 			// ValueBool:  uv.ValueBool(),
// 			Index:      0,
// 			Valid:      true,
// 			Hide:       false,
// 		}
// 	}
//
// 	return ret
// }
