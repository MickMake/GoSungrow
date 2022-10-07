package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	datatable "go.pennock.tech/tabular/auto"
	"reflect"
	"strconv"
	"strings"
	"time"
)


const (
	PointTypeInstant = "instant"
	PointTypeBoot = "boot"
	PointTypeDaily = "daily"
	PointTypeMonthly = "monthly"
	PointTypeYearly = "yearly"
	PointTypeTotal = "total"
)

type PointId string
type DataPointEntries []DataEntry
type DataPoints map[PointId]DataPointEntries

type DataMap struct {
	DataPoints DataPoints
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
	Index      int          `json:"index"`
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
	return DataMap{ DataPoints: make(DataPoints)}
}

// func (dm *DataMap) Add(point string, entry DataEntry) {
// 	dm.Entries[point] = entry
// 	dm.Order = append(dm.Order, point)
// }


func (dm *DataMap) StructToPoints(ref interface{}, endpoint string, parentId string, timestamp time.Time) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			valueTo := vo.Field(i).Interface()
			// spew.Dump(&fieldTo)

			if fieldTo.Tag.Get("PointIgnore") != "" {
				continue
			}

			j := fieldTo.Tag.Get("json")
			pid := fieldTo.Tag.Get("PointId")

			switch {
				case pid != "":
					j = pid
				case j != "" && pid == "":
					pid = j
			}

			name := fieldTo.Tag.Get("PointName")
			if name == "" {
				name = PointToName(PointId(pid))
			}

			device := fieldTo.Tag.Get("PointDevice")
			if device == "" {
				if parentId != "" {
					device = parentId
				} else {
					device = "virtual"
				}
			}

			unit := fieldTo.Tag.Get("PointUnit")
			var uv UnitValue
			if unit == "" {
				bar := fieldTo.Type.Name()
				// fmt.Printf("bar:%v\n", bar)
				var ignore bool
				switch bar {
					case "int":
						uv.Unit = ""	// "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)
					case "int32":
						uv.Unit = ""	// "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)
					case "int64":
						uv.Unit = ""	// "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)

					case "float32":
						uv.Unit = ""	// "float"
						uv.ValueFloat = float64(valueTo.(float32))
						uv.Value = Float64ToString(uv.ValueFloat)
					case "float64":
						uv.Unit = ""	// "float"
						uv.ValueFloat = valueTo.(float64)
						uv.Value = Float64ToString(uv.ValueFloat)

					case "string":
						uv.Unit = ""	// "string"
						uv.Value = valueTo.(string)

					case "UnitValue":
						fallthrough
					case "api.UnitValue":
						uv = valueTo.(UnitValue)
						uv = uv.UnitValueFix()

					case "bool":
						uv.Unit = ""	// "binary"
						uv.Value = fmt.Sprintf("%v", valueTo)

					default:
						ignore = true
				}

				if ignore {
					continue
				}
			}

			pType := fieldTo.Tag.Get("PointType")
			switch pType {
				case "PointTypeInstant":
					pType = PointTypeInstant
				case "PointTypeBoot":
					pType = PointTypeBoot
				case "PointTypeDaily":
					pType = PointTypeDaily
				case "PointTypeMonthly":
					pType = PointTypeMonthly
				case "PointTypeYearly":
					pType = PointTypeYearly
				case "PointTypeTotal":
					pType = PointTypeTotal
			}

			// fullName := JoinDevicePoint(device, PointId(pid))
			p := Point {
				// EndPoint:  endpoint,
				// FullId:    fullName,
				// Parent:    parentDevice.Split(),
				// Parents:   parents,
				Id:        PointId(pid),
				GroupName: "",
				Name:      name,
				Unit:      uv.Unit,
				Type:      pType,
				Valid:     true,
			}

			var now DateTime
			if timestamp.IsZero() {
				now = NewDateTime(time.Now().Round(5 * time.Minute).Format(DtLayoutZeroSeconds))
			} else {
				now = NewDateTime(timestamp.String())
			}

			if parentId == "virtual" {
				fmt.Sprintf("")
			}
			dm.AddEntry(endpoint, device, p, now, uv.Value)

			alias := fieldTo.Tag.Get("PointAlias")
			if alias != "" {
				// fullName = NameDevicePoint(device, PointId(alias))
				p.Id = PointId(alias)
				dm.AddEntry(endpoint, device, p, now, uv.Value)
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
		// sp := strings.Split(entry, ".")
		// entry = sp[len(sp)-1]
		// pe = dm.DataPoints[PointId(entry)]
		// if pe != nil {
		// 	ret = pe.GetEntry(index)
		// 	break
		// }
	}
	return ret
}

// func (dm *DataMap) CopyEntry(entry string, index int) *DataEntry {
// 	l := len(dm.DataPoints[PointName(entry)]) - 1
// 	if index > l {
// 		index = l
// 	}
// 	ret := dm.DataPoints[PointName(entry)][index]
// 	return &ret
// }

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

func (dm *DataMap) LowerUpper(lower string, upper string, index int) float64 {
	var ret float64
	for range Only.Once {
		l := dm.GetEntry(lower, index)
		if l.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lower, upper, index)
			break
		}

		u := dm.GetEntry(upper, index)
		if u.IsNotValid() {
			fmt.Printf("ERROR: LowerUpper('%s', '%s', %d)\n", lower, upper, index)
			break
		}

		if l.ValueFloat > 0 {
			ret = 0 - l.ValueFloat
		}
		ret = u.ValueFloat
	}
	return ret
}

func (dm *DataMap) GetPercent(value string, max string, index int) float64 {
	var ret float64
	for range Only.Once {
		v := dm.GetEntry(value, index)
		if v.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", value, max, index)
			break
		}

		m := dm.GetEntry(max, index)
		if m.IsNotValid() {
			fmt.Printf("ERROR: GetPercent('%s', '%s', %d)\n", value, max, index)
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


// func (dm *DataMap) FromRefAddState(refname string, psId string, point string, name string) {
// 	v := dm.GetEntry(refname)
// 	dm.Entries[point] = v.CreateState(psId, point, name)
// 	dm.Order = append(dm.Order, point)
// }
//
// func (dm *DataMap) AddVirtualValue(refname string, point string, name string, value float64) {
// 	v := dm.GetEntry(refname)
// 	dm.Entries[point] = v.CreateFloat(VirtualPsId, point, name, value)
// 	dm.Order = append(dm.Order, point)
// }
//
// func (ref *DataMap) AddUnitValue(refname string, point string, name string, value UnitValue) {
// 	v := ref.GetEntry(refname)
// 	ref.Entries[point] = v.FromRefAddFloat("virtual", point, name, value.Value)
// 	ref.Order = append(ref.Order, point)
// }


func (dm *DataMap) AppendMap(add DataMap) {
	for range Only.Once {
		if dm.DataPoints == nil {
			dm.DataPoints = make(DataPoints)
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
			// if ep, ok := Points[point]; ok {
			// 	jep, _ := json.Marshal(ep)
			// 	jde, _ := json.Marshal(ep)
			// 	fmt.Printf("EXISTS[%s]:\n%s\n%s\n", point, jde, jep)
			// 	continue
			// }
			// Points[point] = *de[len(de)-1].Point
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

		if parentId == "virtual" {
			fmt.Sprintf("")
		}

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
		ref := CreateUnitValue(value, unit)
		point.Unit = ref.Unit
		point.Valid = true

		dm.Add(JoinDevicePoint(endpoint, point.Id), DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(parent.Key, point.Id),
			Parent:     parent,

			Point:      &point,
			Date:       date,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}
}

// func (dm *DataMap) AddPointAlias(refPoint Point, point Point) {
// 	for range Only.Once {
// 		p := GetPoint(refPoint.PsKey, refPoint.Id)
// 		if p != nil {
// 			fmt.Printf("Found point already: %s.%s\n", p.PsKey, p.Id)
// 			fmt.Println("&point")
// 			spew.Dump(&point)
// 			fmt.Println("&p")
// 			spew.Dump(&p)
// 			break
// 		}
//
// 		if p.PsKey == "" {
// 			p.PsKey = "virtual"
// 		}
// 		if p.Name == "" {
// 			p.Name = PointToName(point.Id)
// 		}
// 		if p.FullId == "" {
// 			p.FullId = JoinDevicePoint(point.PsKey, point.Id)
// 		}
// 		ref := CreateUnitValue(value, p.Unit)
// 		p.Unit = ref.Unit
// 		p.Valid = true
//
// 		dm.Add(point.Id, DataEntry {
// 			Date:       date,
// 			Point:      p,
// 			Value:      ref.Value,
// 			ValueFloat: ref.ValueFloat,
// 		})
// 	}
// }

// func (dm *DataMap) FromRefAddAlias2(refname string, psId string, point string, name string) {
// 	de := dm.GetEntry(refname)
// 	dm.Add(point, de.CreateAlias(psId, point, name))
// }

// func (dm *DataMap) AddEntryFromRef(refPoint Point, point Point, date DateTime, value string) {
// 	for range Only.Once {
// 		p := GetPoint(refPoint.Parent.Id, refPoint.Id)
// 		if p != nil {
// 			fmt.Printf("Found point already: %s.%s\n", p.Parent.Id, p.Id)
// 			fmt.Println("&point")
// 			spew.Dump(&point)
// 			fmt.Println("&p")
// 			spew.Dump(&p)
// 			break
// 		}
//
// 		point.Parent = point.Parent.Split()
// 		if point.Name == "" {
// 			point.Name = PointToName(point.Id)
// 		}
// 		if point.FullId == "" {
// 			point.FullId = JoinDevicePoint(point.Parent.Key, point.Id)
// 		}
// 		ref := CreateUnitValue(value, point.Unit)
// 		point.Unit = ref.Unit
// 		point.Valid = true
//
// 		dm.Add(point.Id, DataEntry {
// 			EndPoint:   p.EndPoint,
// 			Date:       date,
// 			Point:      p,
// 			Value:      ref.Value,
// 			ValueFloat: ref.ValueFloat,
// 		})
// 	}
// }


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

func (dm *DataMap) AddUnitValue(endpoint string, parentId string, pid PointId, name string, date DateTime, ref UnitValue) {
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
			point.Unit = ref.Unit
		}
		if point.Name == "" {
			point.Name = name
		}
		if point.Name == "" {
			point.Name = PointToName(pid)
		}

		dm.Add(NameDevicePoint(endpoint, pid), DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(parent.Key, pid),
			Parent:     parent,

			Point:      point,
			Date:       date,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}
}

func (dm *DataMap) AddFloat(endpoint string, parentId string, pid PointId, name string, date DateTime, value float64) {
	for range Only.Once {
		fvs := Float64ToString(value)
		point := GetPoint(parentId, pid)
		if point == nil {
			// No UV found. Create one.
			dm.Add(pid, CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, CreateUnitValue(fvs, "float")))
			break
		}

		ref := CreateUnitValue(fvs, point.Unit)
		if ref.Unit != point.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, point.Unit, ref.ValueFloat, ref.Unit)
			point.Unit = ref.Unit
		}

		var parent ParentDevice
		parent.Set(parentId)
		point.Parents.Add(parent)

		dm.Add(pid, DataEntry {
			EndPoint:   endpoint,
			FullId:     JoinDevicePoint(parent.Key, pid),
			Parent:     parent,

			Date:       date,
			Point:      point,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}

	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
		Unit:       "float",
		Value:      fmt.Sprintf("%f", value),
		ValueFloat: 0,
	})
	dm.Add(pid, de)
}

func (dm *DataMap) AddString(endpoint string, parentId string, pid PointId, name string, date DateTime, value string) {
	dm.Add(pid, CreateDataEntryString(date, endpoint, parentId, pid, name, value))
}

func (dm *DataMap) AddInt(endpoint string, parentId string, pid PointId, name string, date DateTime, value int64) {
	de := CreateDataEntryUnitValue(date, endpoint, parentId, pid, name, UnitValue {
		Unit:       "int",
		Value:      fmt.Sprintf("%d", value),
		ValueFloat: float64(value),
	})
	dm.Add(pid, de)
}

func (dm *DataMap) FromRefAddAlias(refname string, parentId string, pid PointId, name string) {
	pe := dm.GetEntry(refname, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddAlias('%s', '%s', '%s', '%s')\n", refname, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateAlias(pe.EndPoint, parentId, pid, name))
}

func (dm *DataMap) FromRefAddState(refname string, parentId string, pid PointId, name string) {
	pe := dm.GetEntry(refname, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddState('%s', '%s', '%s', '%s')\n", refname, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateState(pe.EndPoint, parentId, pid, name))
}

func (dm *DataMap) FromRefAddFloat(refname string, parentId string, pid PointId, name string, value float64) {
	pe := dm.GetEntry(refname, 0)
	if pe.IsNotValid() {
		fmt.Printf("ERROR: FromRefAddFloat('%s', '%s', '%s', '%s')\n", refname, parentId, pid, name)
		return
	}
	dm.Add(pid, pe.CreateFloat(pe.EndPoint, parentId, pid, name, value))
}

// func (dm *DataMap) CopyPoint(refname string, psId string, point PointId, name string, value float64) {
// 	pe := dm.GetEntry(refname, 0)
// 	if pe.IsNotValid() {
// 		fmt.Printf("ERROR: CopyPoint('%s', '%s', '%s', '%s')\n", refname, psId, point, name)
// 		return
// 	}
// 	dm.Add(point, pe.CreateFloat(pe.EndPoint, psId, point, name, value))
// }

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
					v.Point.Type,
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

	var parent ParentDevice
	parent.Set(parentId)
	de.Point.Parents.Add(parent)

	de.FullId = NameDevicePoint(parent.Key, pid)
	// de.Point.FullId = NameDevicePoint(psId, point)
	// de.Point.Parent.Key = psId
	// de.Point.Parent = de.Point.Parent.Split()
	de.Point.Id = pid
	de.Point.Name = name
	de.Point.GroupName = parentId
	de.Point.Valid = true
	de.EndPoint = endpoint
	de.Index = 0

	return *de
}

func (de *DataEntry) CreateFloat(endpoint string, parentId string, pid PointId, name string, value float64) DataEntry {
	if name == "" {
		name = PointToName(pid)
	}

	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	uv := CreateUnitValue(Float64ToString(value), de2.Point.Unit)
	de2.Value = uv.Value
	de2.ValueFloat = uv.ValueFloat

	return de2
}

func (de *DataEntry) CreateState(endpoint string, parentId string, pid PointId, name string) DataEntry {
	if name == "" {
		name = PointToName(pid)
	}

	de2 := de.CreateAlias(endpoint, parentId, pid, name)
	de2.Value = fmt.Sprintf("%v", IsActive(de.ValueFloat))
	de2.ValueFloat = 0
	de2.Point.Unit = "binary"

	return de2
}

func (de *DataEntry) UpdateMeta(date *DateTime, parentId string, pid PointId, name string) {

	if date != nil {
		de.Date = *date
	}
	if name == "" {
		name = PointToName(pid)
	}

	var parent ParentDevice
	parent.Set(parentId)
	de.Point.Parents.Add(parent)

	de.FullId = NameDevicePoint(parentId, pid)
	de.Parent.Key = parentId
	de.Parent = parent
	de.Point.Id = pid
	de.Point.Name = name
	de.Point.GroupName = parentId
	de.Index = 0
}


func CreateDataEntryActive(date DateTime, endpoint string, parentId string, pid PointId, name string, value float64) DataEntry {
	point := GetPoint(parentId, pid)
	if point == nil {
		if name == "" {
			name = PointToName(pid)
		}
		point = CreatePoint(endpoint, parentId, pid, name, "state")
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(parent.Key, pid),
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
		point = CreatePoint(endpoint, parentId, pid, name, "string")
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(parent.Key, pid),
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
		point = CreatePoint(endpoint, parentId, pid, name, value.Unit)
	}

	var parent ParentDevice
	parent.Set(parentId)
	point.Parents.Add(parent)

	return DataEntry {
		EndPoint:   endpoint,
		FullId:     JoinDevicePoint(parent.Key, pid),
		Parent:     parent,

		Point:      point,
		Date:       date,
		Value:      value.Value,
		ValueFloat: value.ValueFloat,
		Index:      0,
	}
}

func CreatePoint(endpoint string, parentId string, pid PointId, name string, unit string) *Point {
	if name == "" {
		name = PointToName(pid)
	}

	ret := &Point {
		// EndPoint:  endpoint,
		// FullId:    NameDevicePoint(psId, pid),
		// Parent:    ParentDevice{ Key: psId },
		Id:        PointId(pid),
		GroupName: parentId,
		Name:      name,
		Unit:      unit,
		Type:      "",
		Valid:     true,
	}
	// ret.Parent = ret.Parent.Split()
	return ret
}


func CreateUnitValue(value string, unit string) UnitValue {
	ret := UnitValue {
		Unit:       unit,
		Value:      value,
	}
	return ret.UnitValueFix()
}


func (ref *UnitValue) UnitValueFix() UnitValue {
	if ref.Unit == "W" {
		fvs, err := DivideByThousand(ref.Value)
		// fv, err := strconv.ParseFloat(p.Value, 64)
		// fv = fv / 1000
		if err == nil {
			// p.Value = fmt.Sprintf("%.3f", fv)
			ref.Value = fvs
			ref.Unit = "kW"
		}
	}

	if ref.Unit == "Wh" {
		fvs, err := DivideByThousand(ref.Value)
		// fv, err := strconv.ParseFloat(p.Value, 64)
		// fv = fv / 1000
		if err == nil {
			// p.Value = fmt.Sprintf("%.3f", fv)
			ref.Value = fvs
			ref.Unit = "kWh"
		}
	}

	ref.ValueFloat, _ = strconv.ParseFloat(ref.Value, 64)

	return *ref
}

func (ref *UnitValue) UnitValueToPoint(endpoint string, parentId string, pid PointId, name string) *Point {
	uv := ref.UnitValueFix()

	// u := ref.Unit
	//
	// if ref.Unit == "W" {
	// 	fvs, err := DivideByThousand(ref.Value)
	// 	// fv, err := strconv.ParseFloat(p.Value, 64)
	// 	// fv = fv / 1000
	// 	if err == nil {
	// 		// p.Value = fmt.Sprintf("%.3f", fv)
	// 		ref.Value = fvs
	// 		ref.Unit = "kW"
	// 	}
	// }
	//
	// if ref.Unit == "Wh" {
	// 	fvs, err := DivideByThousand(ref.Value)
	// 	// fv, err := strconv.ParseFloat(p.Value, 64)
	// 	// fv = fv / 1000
	// 	if err == nil {
	// 		// p.Value = fmt.Sprintf("%.3f", fv)
	// 		ref.Value = fvs
	// 		ref.Unit = "kWh"
	// 	}
	// }

	if name == "" {
		name = PointToName(pid)
	}

	ret := GetPoint(parentId, pid)
	if !ret.Valid {
		ret = &Point {
			// EndPoint:  endpoint,
			// FullId:    "",
			// Parent:    ParentDevice{ Key: psId },
			Id:        pid,
			GroupName: "",
			Name:      name,
			Unit:      uv.Unit,
			Type:      "PointTypeInstant",
			Valid:     true,
			States:    nil,
		}
		// ret.Parent = ret.Parent.Split()
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

// // 		entries.AddVirtualValue("DailyTotalLoad", "DailyPvEnergyPercent", "daily_pv_energy_percent", "Daily PV Energy Percent", value)
// func NameToRefName(name string) string {
// 	var ret string	// "Daily PV Energy Percent"
// 	ret = strings.ReplaceAll(name, " ", "_")
// 	ret = strings.ToLower(ret)
// 	return ret
// }
//
// func (upper *DataEntry) LowerUpper(lower DataEntry) float64 {
// 	if lower.ValueFloat > 0 {
// 		return 0 - lower.ValueFloat
// 	}
// 	return upper.ValueFloat
// }
//
// 	Type       string
//	Name       string
//	SubName    string
//
//	ParentId   string
//	ParentName string
//
//	UniqueId   string
//	FullName   string
//	Units      string
//	ValueName  string
//	DeviceClass      string
//
//	Value      string
