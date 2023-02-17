package api

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

type DataMap struct {
	Map map[string]*DataEntries

	parentDeviceId string
	TimeStamp      time.Time
	EndPointPath   GoStruct.EndPointPath
	StructMap      GoStruct.StructMap
	EndPoint       EndPoint
	Debug          bool
	Error          error
}

func NewDataMap() DataMap {
	return DataMap{
		Map: make(map[string]*DataEntries),
	}
}

func (dm *DataMap) StructToDataMap(endpoint EndPoint, parentDeviceId string, name GoStruct.EndPointPath) DataMap {
	for range Only.Once {
		epName := GoStruct.NewEndPointPath(reflection.GetName("", endpoint))
		epName.Append(name.Strings()...)
		name = epName.Copy()

		dm.EndPoint = endpoint
		dm.parentDeviceId = parentDeviceId
		dm.EndPointPath = name
		dm.TimeStamp = time.Now() // .Round(5 * time.Minute)
		dm.Debug = endpoint.IsDebug()

		// Parse response structure.
		dm.StructMap.InitScan(endpoint.ResponseRef(), GoStruct.StructMapOptions{
			StartAt:        "ResultData",
			Name:           name,
			TimeStamp:      dm.TimeStamp,
			Debug:          dm.Debug,
			AddUnexported:  false,
			AddUnsupported: false,
			AddInvalid:     false,
			AddNil:         false,
			AddEmpty:       false,
		})
		if dm.parentDeviceId == "" {
			dm.parentDeviceId = name.First()
		}
	}

	return *dm
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

		if Current.Value.GetUnit() != point.Unit {
			if dm.Debug {
				_, _ = fmt.Fprintf(os.Stderr, "OOOPS FP['%s'] - Point/Value unit mismatch - Point:%s != Value:%s (%f)\n",
					Current.FieldPath.String(), point.Unit, Current.Value.GetUnit(), Current.Value.First().Value())
			}
			point.Unit = Current.Value.GetUnit()
		}

		if Current.Value.Length() == 0 {
			if dm.Debug {
				_, _ = fmt.Fprintf(os.Stderr, "OOOPS FP['%s'] - UVS is nil\n", Current.FieldPath.String())
			}
			break
		}

		// Add arrays as multiple entries.
		if Current.Value.Length() > 1 {
			entries := CreatePointDataEntries(Current, parentDeviceId, point, date)
			dm.Add(entries.Entries...)
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

func (dm *DataMap) CopyPoint(refEndpoint *GoStruct.Reflect, endpoint GoStruct.EndPointPath, pointId string, pointName string) *GoStruct.Reflect {
	var Current *GoStruct.Reflect
	for range Only.Once {
		var tmp GoStruct.Reflect
		tmp = *refEndpoint
		Current = &tmp

		if pointId != "" {
			Current.DataStructure.PointId = pointId
		}

		if pointName != "" {
			Current.DataStructure.PointName = pointName
			// Current.DataStructure.PointName += " (" + Current.DataStructure.PointId + ")"
			// } else {

		}

		if !endpoint.IsZero() {
			Current.DataStructure.Endpoint.Clear()
			Current.DataStructure.Endpoint = endpoint
			Current.DataStructure.Endpoint.Append(Current.DataStructure.PointId)
		}
		Current.DataStructure.DataTable = false
		Current.DataStructure.DataTableChild = false
		Current.DataStructure.PointValueType = Current.DataStructure.ValueType

		dm.StructMap.Add(Current)
	}
	return Current
}

func (dm *DataMap) CopyPointFromName(refEndpoint string, endpoint GoStruct.EndPointPath, pointId string, pointName string) *GoStruct.Reflect {
	var Current *GoStruct.Reflect
	for range Only.Once {
		if ref, ok := dm.StructMap.Map[refEndpoint]; ok {
			Current = dm.CopyPoint(ref, endpoint, pointId, pointName)
			break
		}

		for key := range dm.StructMap.Map {
			if strings.HasSuffix(dm.StructMap.Map[key].FieldPath.String(), refEndpoint) {
				Current = dm.CopyPoint(dm.StructMap.Map[key], endpoint, pointId, pointName)
				break
			}
			if strings.HasSuffix(key, refEndpoint) {
				Current = dm.CopyPoint(dm.StructMap.Map[key], endpoint, pointId, pointName)
				break
			}
		}
	}

	return Current
}

func (dm *DataMap) GetReflect(refEndpoint string) *GoStruct.Reflect {
	var Current *GoStruct.Reflect
	for range Only.Once {
		if ref, ok := dm.StructMap.Map[refEndpoint]; ok {
			Current = ref
			break
		}

		for key := range dm.StructMap.Map {
			if strings.HasSuffix(dm.StructMap.Map[key].FieldPath.String(), refEndpoint) {
				Current = dm.StructMap.Map[key]
				break
			}
			if strings.HasSuffix(key, refEndpoint) {
				Current = dm.StructMap.Map[key]
				break
			}
		}
	}

	return Current
}

func (dm *DataMap) MakeState(Current *GoStruct.Reflect) *GoStruct.Reflect {
	// func (dm *DataMap) MakeState(refEndpoint *GoStruct.Reflect, endpoint GoStruct.EndPointPath, pointId string, pointName string) *GoStruct.Reflect {
	for range Only.Once {
		// Current = dm.CopyPoint(refEndpoint, endpoint, pointId, pointName)
		bv := Current.Value.First().IsNotZero()
		Current.Value.Reset()
		vt := valueTypes.SetUnitValueBool(bv)
		Current.Value.AppendUnitValue(vt)
		Current.Interface = vt
		Current.InterfaceValue = vt
		Current.DataStructure.PointUnit = "--"
		Current.DataStructure.ValueType = "Bool"
	}
	return Current
}

func (dm *DataMap) LowerUpper(lowerEntry *GoStruct.Reflect, upperEntry *GoStruct.Reflect) float64 {
	var ret float64
	for range Only.Once {
		if lowerEntry.Value.First().ValueFloat() > 0 {
			ret = 0 - lowerEntry.Value.First().ValueFloat()
			break
		}
		ret = upperEntry.Value.First().ValueFloat()
	}
	return ret
}

func (dm *DataMap) GetPercent(entry *GoStruct.Reflect, max *GoStruct.Reflect, precision int) float64 {
	return GetPercent(entry.Value.First().ValueFloat(), max.Value.First().ValueFloat(), precision)
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
			// endpoint := de.EndPoint

			if dm.Map[de.EndPoint] == nil {
				dm.Map[de.EndPoint] = &DataEntries{Entries: []DataEntry{}}
			}

			// We end up using the last entry in this list for MQTT.
			entries := dm.Map[de.EndPoint]
			if entries.Add(de) == nil {
				continue
			}
			// dm.Order = append(dm.Order, endpoint)

			if Points.Exists(de.EndPoint) {
				fmt.Printf("EXISTS: %s\n", de.EndPoint)
			}
			Points.Add(*de.Point)
		}
	}
}

func (dm *DataMap) ProcessMap() {
	for range Only.Once {
		// Convert Struct.Map to DataMap
		for _, Child := range dm.StructMap.Map {
			if Child.IsPointIgnore() {
				continue
			}

			var when valueTypes.DateTime
			if Child.IsPointTimestampNotZero() {
				when = valueTypes.SetDateTimeValue(Child.DataStructure.PointTimestamp)
			} else {
				when = valueTypes.SetDateTimeValue(dm.TimeStamp)
			}

			dm.AddPointUnitValues(Child, dm.parentDeviceId, when)
		}

		// Convert Struct.VirtualMap to DataMap
		for _, Child := range dm.StructMap.VirtualMap {
			if Child.IsPointIgnore() {
				continue
			}

			var when valueTypes.DateTime
			if Child.IsPointTimestampNotZero() {
				when = valueTypes.SetDateTimeValue(Child.DataStructure.PointTimestamp)
			} else {
				when = valueTypes.SetDateTimeValue(dm.TimeStamp)
			}

			switch {
			case Child.DataStructure.Endpoint.IsBeginsWith("virtual"):
				// Don't prepend path - fixes the double virtual path issue.
			case Child.DataStructure.PointVirtualShift > 0:
				Child.DataStructure.Endpoint.ShiftLeft(Child.DataStructure.PointVirtualShift)
				Child.DataStructure.Endpoint.InsertFirst("virtual")
			default:
				Child.DataStructure.Endpoint.ReplaceFirst("virtual")
			}

			dm.AddPointUnitValues(Child, dm.parentDeviceId, when)
		}
	}
}

type Tables GoStruct.StructTables

func (dm *DataMap) CreateDataTables() Tables {
	tables := make(Tables, 0)

	for range Only.Once {
		for name := range dm.StructMap.TableMap {
			var ret GoStruct.StructTable
			dm.Error = ret.Process(dm.EndPoint.GetArea().String(), name, dm.StructMap.TableMap[name])
			if dm.Error != nil {
				break
			}

			_, dm.Error = ret.CreateTable()
			if dm.Error != nil {
				break
			}

			tables[name] = &ret
		}
	}

	return tables
}

func (dm *DataMap) CreateResultTable(full bool) output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable(
			"Date",
			"Point Id",
			"Value",
			"Unit",
			"Unit Type",
			"Group Name",
			"Description",
			"Update Freq",
		)

		for p := range dm.Map {
			entries := dm.Map[p].Entries
			for _, de := range entries {
				if full {
					if de.Current.DataStructure.DataTable {
						continue // We are a datatable parent.
					}

					if de.Current.CurrentReflect.IsPointListFlatten() {
						// if len(de.Current.ChildReflect) > 0 {
						// 	continue // We already have the children.
						// }
						if len(de.Current.ChildReflect) == 0 {
							continue // Ignore points with no children.
						}
					}
				} else {
					if de.Hide {
						continue // Ignore hidden entries.
					}
					if de.Current.DataStructure.DataTableChild {
						continue // Ignore data table children.
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
				dm.Error = table.AddRow(
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
		table.SetTitle("EndPoint Data %s.%s", dm.EndPoint.GetArea(), dm.StructMap.Name.String()) // endpoint.GetArea(), endpoint.GetName()
		if dm.StructMap.Start.DataStructure.DataTableName != "" {
			table.AppendTitle(" - %s", dm.StructMap.Start.DataStructure.DataTableName)
		}
		table.SetFilePrefix("%s.%s", dm.EndPoint.GetArea(), dm.StructMap.Name.String())
		table.SetGraphFilter("")
		table.Sort("Point Id")
		if full {
			table.SetJson([]byte(dm.EndPoint.GetJsonData(false)))
			table.SetRaw([]byte(dm.EndPoint.GetJsonData(true)))
		}
	}

	return table
}

func (dm DataMap) String() string {
	var ret string
	for range Only.Once {
		table := output.NewTable(
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

				dm.Error = table.AddRow(
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
				if dm.Error != nil {
					break
				}
			}
		}

		ret = table.String()

		// table := datatable.New("utf8-heavy")
		// table.AddHeaders(
		// 	"Index",
		// 	"EndPoint",
		//
		// 	"Id",
		// 	"Name",
		// 	"Unit",
		// 	"Type",
		// 	"Value",
		// 	"(Value)",
		// 	"Valid",
		//
		// 	"GroupName",
		// 	"Parent Ids",
		// 	"Parent Types",
		// 	"Parent Codes",
		// )
		// // dm.Order - Produces double the amount of entries for some reason.
		// i := 0
		// for k := range dm.Map {
		// 	for _, v := range dm.Map[k].Entries {
		// 		i++
		//
		// 		table.AddRowItems(
		// 			i,
		// 			v.EndPoint,
		//
		// 			v.Point.Id,
		// 			v.Point.Description,
		// 			v.Point.Unit,
		// 			v.Point.UpdateFreq,
		// 			v.Value,
		// 			v.Current.Value.First(),
		// 			v.Point.Valid,
		//
		// 			v.Point.GroupName,
		// 			v.Point.Parents.PsIds(),
		// 			v.Point.Parents.Types(),
		// 			v.Point.Parents.Codes(),
		// 		)
		// 	}
		// }
		//
		// ret, _ := table.Render()
		// fmt.Println(ret)
	}
	return ret
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

func CreatePointDataEntry(Current *GoStruct.Reflect, parentDeviceId string, point Point, dateTime valueTypes.DateTime, uv valueTypes.UnitValue) DataEntry {
	// CreatePointDataEntry(Current, Current.EndPointPath().String(), parentDeviceId, point, date, *Current.Value.First())
	var ret DataEntry
	for range Only.Once {
		if uv.DeviceId() != "" {
			parentDeviceId = uv.DeviceId()
		} else {
			uv.SetDeviceId(parentDeviceId)
		}

		ret = DataEntry{
			Current:  Current,
			EndPoint: Current.EndPointPath().String(),
			Point:    &point,
			Parent:   NewParentDevice(parentDeviceId),
			Date:     dateTime,
			Value:    uv,
			Valid:    true,
			Hide:     false,
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
		res := valueTypes.SizeOfInt(sorted)
		for i, uv := range sorted {
			var epn string
			if uv.ValueKey() == "" {
				epn = JoinWithDots(res, valueTypes.DateTimeLayoutDay, Current.EndPointPath().String(), i)
			} else {
				epn = JoinWithDots(res, valueTypes.DateTimeLayoutDay, Current.EndPointPath().String(), uv.ValueKey())
			}

			if uv.DeviceId() != "" {
				parentDeviceId = uv.DeviceId()
			} else {
				uv.SetDeviceId(parentDeviceId)
			}

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
	// unit string				- Current.Value.GetUnit()
	// Type string				- Current.Value.Type()
	// timeSpan string			- Current.DataStructure.PointUpdateFreq

	var point Point
	for range Only.Once {
		// pid := valueTypes.SetPointIdString(parentDeviceId, Current.PointId())
		// name := Current.DataStructure.PointName
		// if name == "" {
		// 	name = valueTypes.PointToName(Current.PointId())
		// }

		var parent ParentDevice
		parent.Set(parentDeviceId)
		var parents ParentDevices
		parents.Add(parent)

		point = Point{
			Parents:     parents,
			Id:          Current.PointId(),
			GroupName:   Current.PointGroupName(),
			Description: Current.PointName(),
			Unit:        Current.ValueUnit(),
			UpdateFreq:  Current.PointUpdateFreq(),
			ValueType:   Current.ValueType(),
			Valid:       Current.IsOk,
			States:      nil,
		}
		point.FixUnitType()
	}

	return point
}

func GetPercent(value float64, max float64, precision int) float64 {
	if max == 0 {
		return 0
	}

	percent := valueTypes.SetPrecision((value/max)*100, precision)
	return percent
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
