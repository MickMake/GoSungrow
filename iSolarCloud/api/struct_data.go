package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	datatable "go.pennock.tech/tabular/auto"
	"strings"
	"time"
)


const (
	PointTimeSpanInstant = "instant"
	PointTimeSpanBoot    = "boot"
	PointTimeSpanDaily   = "daily"
	PointTimeSpanMonthly = "monthly"
	PointTimeSpanYearly  = "yearly"
	PointTimeSpanTotal   = "total"
)

type PointId string
type DataPointEntries []DataEntry
// type DataPoints struct {
// 	Map map[PointId]DataPointEntries
// }

func (de *DataPointEntries) Hide() {
	for range Only.Once {
		for i := range *de {
			(*de)[i].Hide = true
		}
	}
}


type DataMap struct {
	DataPoints map[PointId]DataPointEntries
	Order      []PointId
}

type DataEntry struct {
	Point      *Point       `json:"point"`
	Date       DateTime     `json:"date"`

	EndPoint   string       `json:"endpoint"`
	FullId     PointId      `json:"full_id"`
	Parent     ParentDevice `json:"parent"`

	Value      string       `json:"value"`
	ValueFloat float64      `json:"value_float"`
	ValueBool  bool         `json:"value_bool"`
	Index      int          `json:"index"`
	Valid      bool         `json:"valid"`
	Hide       bool         `json:"hide"`
}

func (de *DataEntry) IsValid() bool {
	var ok bool
	for range Only.Once {
		if de.Point == nil {
			break
		}
		if de.Point.Valid == false {
			break
		}
		ok = true
	}
	return ok
}
func (de *DataEntry) IsNotValid() bool {
	return !de.IsValid()
}


func NewDataMap() DataMap {
	return DataMap {
		DataPoints: make(map[PointId]DataPointEntries),
	}
}

// func (dm *DataMap) Add(point string, entry DataEntry) {
// 	dm.Entries[point] = entry
// 	dm.Order = append(dm.Order, point)
// }


func (dm *DataMap) StructToPoints(ref interface{}, endpoint string, parentId string, timestamp DateTime) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		// Iterate over all available fields and read the tags value
		tp := apiReflect.GetPointTags(ref, endpoint)
		// fmt.Printf("TP: %v\n", tp)

		for _, f := range tp {
			// if strings.Contains(strings.ToLower(f.Json), "p83012") {
			// 	fmt.Sprintf("")
			// }

			if f.PointIgnore {
				// fmt.Printf("IGNORE: %s\n", f.PointId)
				continue
			}

			if f.PointName == "" {
				f.PointName = PointToName(PointId(f.PointId))
			}

			if f.PointDevice == "" {
				if parentId != "" {
					f.PointDevice = parentId
				} else {
					f.PointDevice = "virtual"
				}
			}
			var parents ParentDevices
			parents.Add(ParentDevice{Key: f.PointDevice})

			// pUnit := fieldTo.Tag.Get("PointUnit")
			var uv UnitValue
			var uvs UnitValues
			var ignore bool
			switch f.ValueType {
				case "int":
					uv = SetUnitValueInteger(int64(f.Value.(int)), f.PointUnit)
				case "int32":
					uv = SetUnitValueInteger(int64(f.Value.(int32)), f.PointUnit)
				case "int64":
					uv = SetUnitValueInteger(f.Value.(int64), f.PointUnit)

				case "float32":
					uv = SetUnitValueFloat(float64(f.Value.(float32)), f.PointUnit)
				case "float64":
					uv = SetUnitValueFloat(f.Value.(float64), f.PointUnit)

				case "string":
					uv = SetUnitValueString(f.Value.(string), f.PointUnit)

				case "bool":
					v := f.Value.(bool)
					uv = SetUnitValueString(fmt.Sprintf("%v", v), "binary")

				case "api.UnitValue":
					fallthrough
				case "UnitValue":
					uv = f.Value.(UnitValue)
					// uv = uv.UnitValueFix()

				case "UnitValues":
					fallthrough
				case "[]api.UnitValue":
					fallthrough
				case "[]UnitValue":
					uvs = f.Value.([]UnitValue)

				case "api.Float":
					fallthrough
				case "Float":
					v := f.Value.(Float)
					uv = SetUnitValueFloat(v.Value(), f.PointUnit)

				case "[]api.Float":
					fallthrough
				case "[]Float":
					v := f.Value.([]Float)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueFloat(val.Value(), f.PointUnit))
					}

				case "api.Integer":
					fallthrough
				case "Integer":
					v := f.Value.(Integer).Value()
					uv = SetUnitValueInteger(v, f.PointUnit)

				case "[]api.Integer":
					fallthrough
				case "[]Integer":
					v := f.Value.([]Integer)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueInteger(val.Value(), f.PointUnit))
					}

				case "api.Count":
					fallthrough
				case "Count":
					v := f.Value.(Count).Value()
					uv = SetUnitValueInteger(v, "counter")

				case "[]api.Count":
					fallthrough
				case "[]Count":
					v := f.Value.([]Count)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueInteger(val.Value(), f.PointUnit))
					}

				case "api.Bool":
					fallthrough
				case "Bool":
					v := f.Value.(Bool).String()
					uv = SetUnitValueString(v, "binary")

				case "[]api.Bool":
					fallthrough
				case "[]Bool":
					v := f.Value.([]Bool)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueString(val.String(), f.PointUnit))
					}

				case "api.String":
					fallthrough
				case "String":
					v := f.Value.(String).String()
					uv = SetUnitValueString(v, f.PointUnit)

				case "[]api.String":
					fallthrough
				case "[]String":
					v := f.Value.([]String)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueString(val.Value(), f.PointUnit))
					}

				case "api.PsKey":
					fallthrough
				case "PsKey":
					v := f.Value.(PsKey).Value()
					uv = SetUnitValueString(v, f.PointUnit)

				case "[]api.PsKey":
					fallthrough
				case "[]PsKey":
					v := f.Value.([]PsKey)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueString(val.Value(), f.PointUnit))
					}

				case "api.DateTime":
					fallthrough
				case "DateTime":
					v := f.Value.(DateTime).String()
					uv = SetUnitValueString(v, "date")

				case "[]api.DateTime":
					fallthrough
				case "[]DateTime":
					v := f.Value.([]DateTime)
					for _, val := range v {
						uvs = append(uvs, SetUnitValueString(val.String(), f.PointUnit))
					}

				case "[]string":
					// v := strings.Join(f.Value.([]string), ",")
					j, err := json.Marshal(f.Value.([]string))
					if err != nil {
						j = []byte(fmt.Sprintf("%v", f.Value.([]string)))
					}
					uv = SetUnitValueString(string(j), f.PointUnit)

				default:
					ignore = true
			}
			// fmt.Printf("TYPE: %s.%s (%s)\n", f.Endpoint, f.PointId, f.ValueType)
			if ignore {
				// fmt.Printf("IGNORE: %s.%s (%s)\n", f.Endpoint, f.PointId, f.ValueType)
				continue
			}

			switch f.PointTimeSpan {
				case "PointTimeSpanInstant":
					f.PointTimeSpan = PointTimeSpanInstant
				case "PointTimeSpanBoot":
					f.PointTimeSpan = PointTimeSpanBoot
				case "PointTimeSpanDaily":
					f.PointTimeSpan = PointTimeSpanDaily
				case "PointTimeSpanMonthly":
					f.PointTimeSpan = PointTimeSpanMonthly
				case "PointTimeSpanYearly":
					f.PointTimeSpan = PointTimeSpanYearly
				case "PointTimeSpanTotal":
					f.PointTimeSpan = PointTimeSpanTotal
			}

			var now DateTime
			if timestamp.IsZero() {
				now = NewDateTime(time.Now().Round(5 * time.Minute).Format(DtLayoutZeroSeconds))
			} else {
				now = NewDateTime(timestamp.String())
			}

			p := Point {
				Parents:   parents,
				Id:        PointId(f.PointId),
				GroupName: f.PointGroupName,
				Name:      f.PointName,
				Unit:      uv.Unit(),
				TimeSpan:  f.PointTimeSpan,
				ValueType: f.PointValueType,
				Valid:     true,
				States:    nil,
			}

			// Add arrays as multiple entries.
			if len(uvs) > 0 {
				// @TODO - Think about adding in arrays of values OR just marshal arrays into JSON.
				res := "%s.%d"
				switch  {
					case len(uvs) > 99:
						res = "%s.%.3d"
					case len(uvs) > 9:
						res = "%s.%.2d"
				}
				for i, val := range uvs {
					// fmt.Sprintf(res, f.Endpoint, i)
					dm.AddEntry(fmt.Sprintf(res, f.Endpoint, i), f.PointDevice, p, now, val.String())
				}
				continue
			}

			dm.AddEntry(f.Endpoint, f.PointDevice, p, now, uv.String())
			if f.PointAlias != "" {
				// fullName = NameDevicePoint(device, PointId(alias))
				p.Id = PointId(f.PointAlias)
				dm.AddEntry(f.Endpoint, f.PointDevice, p, now, uv.String())
			}
		}

		for _, f := range tp {
			// Reference Units from another data point.
			if f.PointUnitFrom != "" {
				sdp := dm.GetEntryFromPointId(f.PointUnitFrom)
				if sdp == nil {
					continue
				}
				ddp := dm.GetEntryFromPointId(f.PointId)
				if ddp == nil {
					continue
				}
				ddp.SetUnits(sdp.GetEntry(0).Value)
				sdp.Hide()
			}
		}
	}
}

const LastEntry = -1
func (de *DataPointEntries) GetEntry(index int) DataEntry {
	for range Only.Once {
		l := len(*de) - 1
		if index > l {
			index = l
			break
		}
		if index < 0 {
			index = l + index + 1
			if index < 0 {
				index = 0
			}
		}
	}
	return (*de)[index]
}

func (de *DataPointEntries) GetUnits() string {
	var unit string
	for range Only.Once {
		for _, v := range *de {
			unit = v.Point.Unit
			break
		}
	}
	return unit
}

func (de *DataPointEntries) SetUnits(units string) {
	for range Only.Once {
		for i := range *de {
			(*de)[i].Point.Unit = units
		}
	}
}

func (dm *DataMap) GetEntry(entry string, index int) DataEntry {
	var ret DataEntry
	for range Only.Once {
		pe := dm.DataPoints[PointId(entry)]
		if pe != nil {
			ret = pe.GetEntry(index)
			break
		}

		for k, v := range dm.DataPoints {
			if strings.HasSuffix(string(k), "." + entry) {
				ret = v.GetEntry(index)
				break
			}
		}
	}
	return ret
}

func (dm *DataMap) GetEntryFromPointId(pointId string) *DataPointEntries {
	var ret *DataPointEntries
	for range Only.Once {
		for _, v := range dm.DataPoints {
			pe := v.GetEntry(0)
			if pe.IsNotValid() {
				continue
			}

			if string(pe.Point.Id) == pointId {
				ret = &v
				break
			}
		}
	}
	return ret
}

func (dm *DataMap) HideEntry(pointId string) {
	for range Only.Once {
		de := dm.GetEntryFromPointId(pointId)
		de.Hide()
	}
}

func (dm *DataMap) GetFloatValue(entry string, index int) float64 {
	var ret float64
	for range Only.Once {
		pe := dm.GetEntry(entry, index)
		if pe.IsNotValid() {
			fmt.Printf("ERROR: GetFloatValue('%s', '%d')\n", entry, index)
			break
		}
		ret = pe.ValueFloat
	}
	return ret
}

func (dm *DataMap) LowerUpper(lowerEntry string, upperEntry string, index int) float64 {
	var ret float64
	for range Only.Once {
		l := dm.GetEntry(lowerEntry, index)
		if l.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lowerEntry, upperEntry, index)
			break
		}

		u := dm.GetEntry(upperEntry, index)
		if u.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lowerEntry, upperEntry, index)
			break
		}

		if l.ValueFloat > 0 {
			ret = 0 - l.ValueFloat
			break
		}
		ret = u.ValueFloat
	}
	return ret
}

func (dm *DataMap) GetPercent(entry string, max string, index int) float64 {
	var ret float64
	for range Only.Once {
		v := dm.GetEntry(entry, index)
		if v.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", entry, max, index)
			break
		}

		m := dm.GetEntry(max, index)
		if m.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", entry, max, index)
			break
		}

		ret = GetPercent(v.ValueFloat, m.ValueFloat)
	}
	return ret
}

func (dm *DataMap) GetValue(entry string, index int) float64 {
	var ret float64
	for range Only.Once {
		v := dm.GetEntry(entry, index)
		if v.IsNotValid() {
			fmt.Printf("ERROR: GetValue('%s', %d)\n", entry, index)
			break
		}

		ret = v.ValueFloat
	}
	return ret
}


func (dm *DataMap) AppendMap(add DataMap) {
	for range Only.Once {
		if dm.DataPoints == nil {
			dm.DataPoints = make(map[PointId]DataPointEntries)
		}

		for point, de := range add.DataPoints {
			if dd, ok := dm.DataPoints[point]; ok {
				jde, _ := json.Marshal(de)
				jdd, _ := json.Marshal(dd)
				if string(jdd) != string(jde) {
					fmt.Printf("DIFF ")
				}
				fmt.Printf("Duplicate[%s]:\n%s\n%s\n", point, jde, jdd)
				continue
			}
			dm.DataPoints[point] = de
			dm.Order = append(dm.Order, point)

			if Points.Exists(point) {
				fmt.Printf("EXISTS: %s\n", point)
			}
			Points.Add(point, *de[len(de)-1].Point)
		}
	}
}

func (dm *DataMap) Add(pid PointId, de DataEntry) {
	for range Only.Once {
		if !strings.Contains(string(pid), ".") {
			pid = PointId(de.EndPoint + "." + string(pid))
		}

		de.Index = len(dm.Order)
		dm.DataPoints[pid] = append(dm.DataPoints[pid], de)
		dm.Order = append(dm.Order, pid)

		if Points.Exists(pid) {
			fmt.Printf("EXISTS: %s\n", pid)
		}
		Points.Add(pid, *de.Point)
		// if ep, ok := Points[point]; ok {
		// 	jep, _ := json.Marshal(ep)
		// 	jde, _ := json.Marshal(ep)
		// 	fmt.Printf("EXISTS[%s]:\n%s\n%s\n", point, jde, jep)
		// 	continue
		// }
		// Points[point] = *de.Point
	}
}

func (dm *DataMap) AddEntry(endpoint string, parentId string, point Point, date DateTime, value string) {
	for range Only.Once {
		unit := point.Unit	// Save unit.

		// Match to a previously defined point.
		p := GetPoint(endpoint, point.Id)
		if p == nil {
			point = *p
		}

		// var parents ParentDevices
		// parents.Add(ParentDevice{Key: device})
		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		if point.Name == "" {
			point.Name = PointToName(point.Id)
		}
		// fid := JoinDevicePoint(parent.Key, point.Id)
		ref := SetUnitValueString(value, unit)
		point.Unit = ref.Unit()
		point.Valid = true

		if _, ok := dm.DataPoints[point.Id]; ok {
			point.Id += ".BARF"
		}
		// dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
		dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(endpoint, point.Id),
			// FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Point:      &point,
			Date:       date,
			Value:      ref.String(),
			ValueFloat: ref.Value(),
		})
	}
}

func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid PointId, name string, groupName string, date DateTime, ref UnitValue) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		ref = ref.UnitValueFix()

		if name == "" {
			name = string(pid)
		}

		point := GetPoint(parentId, pid)
		if point == nil {
			// No UV found. Create one.
			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, ref))
			break
		}

		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		if point.Unit == "" {
			point.Unit = ref.Unit()
		}
		if point.Name == "" {
			point.Name = name
		}
		if point.Name == "" {
			point.Name = PointToName(pid)
		}
		if point.GroupName == "" {
			point.GroupName = groupName
		}

		dm.Add(NameDevicePoint(endpoint, pid), DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(endpoint, point.Id),
			// FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Point:      point,
			Date:       date,
			Value:      ref.String(),
			ValueFloat: ref.Value(),
		})
	}
}

func (dm *DataMap) AddFloat(endpoint string, parentId string, pid PointId, name string, date DateTime, value float64) {
	for range Only.Once {
		// fvs := Float64ToString(value)
		point := GetPoint(parentId, pid)
		if point == nil {
			// No UV found. Create one.
			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, SetUnitValueFloat(value, "float")))
			break
		}

		ref := SetUnitValueFloat(value, point.Unit)
		if ref.Unit() != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat(), ref.Unit())
			point.Unit = ref.Unit()
		}

		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		dm.Add(pid, DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(endpoint, point.Id),
			// FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Date:       date,
			Point:      point,
			Value:      ref.String(),
			ValueFloat: ref.Value(),
		})
	}

	uv := SetUnitValueFloat(value, "float")
	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
	// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
	// 	Unit:       "float",
	// 	Value:      fmt.Sprintf("%f", value),
	// 	ValueFloat: 0,
	// })
	dm.Add(pid, de)
}

func (dm *DataMap) AddString(endpoint string, parentId string, pid PointId, name string, date DateTime, value string) {
	dm.Add(pid, CreateDataEntryString(date, endpoint, parentId, pid, name, value))
}

func (dm *DataMap) AddInt(endpoint string, parentId string, pid PointId, name string, date DateTime, value int64) {
	uv := SetUnitValueInteger(value, "int")
	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, uv)
	// de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
	// 	Unit:       "int",
	// 	Value:      fmt.Sprintf("%d", value),
	// 	ValueFloat: float64(value),
	// })
	dm.Add(pid, de)
}


func (dm *DataMap) FromRefAddAlias(entry string, parentId string, pid PointId, name string) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddAlias('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateAlias(pe.EndPoint, parentId, pid, name))
}

func (dm *DataMap) FromRefAddState(entry string, parentId string, pid PointId, name string) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddState('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateState(pe.EndPoint, parentId, pid, name))
}

func (dm *DataMap) FromRefAddFloat(entry string, parentId string, pid PointId, name string, value float64) {
	pe := dm.GetEntry(entry, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddFloat('%s', '%s', '%s', '%s')\n", entry, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateFloat(pe.EndPoint, parentId, pid, name, value))
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
			for _, v := range dm.DataPoints[k] {
				table.AddRowItems(
					i,
					v.EndPoint,

					v.Point.Id,
					v.Point.Name,
					v.Point.Unit,
					v.Point.TimeSpan,
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


func (de *DataEntry) CreateAlias(endpoint string, parentId string, pid PointId, name string) DataEntry {
	if name == "" {
		name = PointToName(pid)
	}

	ret := DataEntry {
		Point:      &Point {
			Parents:   de.Point.Parents,
			Id:        pid,
			GroupName: "alias",
			Name:      name,
			Unit:      de.Point.Unit,
			TimeSpan:  de.Point.TimeSpan,
			Valid:     true,
			States:    nil,
		},
		Date:       de.Date,
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(endpoint, pid),
		// FullId:     JoinDevicePoint(parentId, pid),
		Parent:     de.Parent,		// ParentDevice{},
		Value:      de.Value,
		ValueFloat: de.ValueFloat,
		ValueBool:  de.ValueBool,
		Index:      de.Index,
	}

	ret.Parent.Set(parentId)
	de.Point.Parents.Add(ret.Parent)

	// de.FullId = JoinDevicePoint(endpoint, pid)
	// de.FullId = NameDevicePoint(ret.Parent.Key, pid)
	// de.Point.Id = pid
	// de.Point.Name = name
	// de.Point.GroupName = parentId
	// de.Point.Valid = true
	// de.EndPoint = endpoint
	// de.Index = 0

	return ret
}

func (de *DataEntry) CreateFloat(endpoint string, parentId string, pid PointId, name string, value float64) DataEntry {
	if name == "" {
		name = PointToName(pid)
	}

	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	uv := SetUnitValueFloat(value, de2.Point.Unit)
	de2.Value = uv.String()
	de2.ValueFloat = uv.Value()

	return de2
}

func (de *DataEntry) CreateState(endpoint string, parentId string, pid PointId, name string) DataEntry {
	if name == "" {
		name = PointToName(pid)
	}

	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	if de2.ValueFloat == 0 {
		de2.Value = "false"
		de2.ValueBool = false
		de2.ValueFloat = 0
	} else {
		de2.Value = "true"
		de2.ValueBool = true
		de2.ValueFloat = 1
	}
	de2.Point.Unit = "binary"

	return de2
}


func JoinDevicePoint(device string, pid PointId) PointId {
	var ret PointId
	for range Only.Once {
		if device == "" {
			device = "virtual"
		}
		ret = PointId(fmt.Sprintf("%s.%s", device, pid))
	}
	return ret
}

func CreateDataEntryActive(date DateTime, endpoint string, parentId string, pid PointId, name string, value float64) DataEntry {
	point := GetPoint(parentId, pid)
	if point == nil {
		if name == "" {
			name = PointToName(pid)
		}
		point = CreatePoint(parentId, pid, name, "state")
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(endpoint, point.Id),
		// FullId:     JoinDevicePoint(parent.Key, point.Id),
		Parent:     parent,

		Point:      point,
		Date:       date,
		Value:      fmt.Sprintf("%v", IsActive(value)),
		ValueFloat: 0,
		Index:      0,
	}
}

func CreateDataEntryString(date DateTime, endpoint string, parentId string, pid PointId, name string, value string) DataEntry {
	point := GetPoint(parentId, pid)
	if point == nil {
		if name == "" {
			name = PointToName(pid)
		}
		point = CreatePoint(parentId, pid, name, "string")
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(endpoint, pid),
		// FullId:     JoinDevicePoint(parent.Key, pid),
		Parent:     parent,

		Point:      point,
		Date:       date,
		Value:      value,
		ValueFloat: 0,
		Index:      0,
	}
}

func CreateDataEntryUnitValue(date DateTime, endpoint string, parentId string, pid PointId, name string, value UnitValue) DataEntry {
	value = value.UnitValueFix()

	point := GetPoint(parentId, pid)
	if point == nil {
		if name == "" {
			name = PointToName(pid)
		}
		point = CreatePoint(parentId, pid, name, value.Unit())
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(endpoint, pid),
		// FullId:     JoinDevicePoint(parent.Key, pid),
		Parent:     parent,

		Point:      point,
		Date:       date,
		Value:      value.String(),
		ValueFloat: value.Value(),
		Index:      0,
	}
}

func CreatePoint(parentId string, pid PointId, name string, unit string) *Point {
	if name == "" {
		name = PointToName(pid)
	}

	ret := &Point {
		Id:        pid,
		GroupName: parentId,
		Name:      name,
		Unit:      unit,
		TimeSpan:  "",
		Valid:     true,
	}
	return ret
}

func IsActive(value float64) bool {
	if (value > 0.01) || (value < -0.01) {
		return true
	}
	return false
}

func GetPercent(value float64, max float64) float64 {
	if max == 0 {
		return 0
	}
	return (value / max) * 100
}
