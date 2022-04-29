package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"strconv"
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


type DataMap struct {
	Entries map[string]DataEntry
	Order []string
}

type DataEntry struct {
	EndPoint   string   `json:"endpoint"`
	Point      *Point   `json:"point"`
	Date       DateTime `json:"date"`
	// Id        string `json:"id"`
	// GroupName string `json:"group_name"`
	// Name      string `json:"name"`
	// Unit       string   `json:"unit"`
	Value      string   `json:"value"`
	ValueFloat float64  `json:"value_float"`
	Index      int      `json:"index"`
}

func NewDataMap() DataMap {
	return DataMap{ Entries: make(map[string]DataEntry)}
}

// func (dm *DataMap) Add(point string, entry DataEntry) {
// 	dm.Entries[point] = entry
// 	dm.Order = append(dm.Order, point)
// }


func (dm *DataMap) StructToPoints(endpoint string, ref interface{}) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		now := NewDateTime(time.Now().Round(5 * time.Minute).Format(DtLayoutZeroSeconds))

		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			valueTo := vo.Field(i).Interface()
			spew.Dump(&fieldTo)

			j := fieldTo.Tag.Get("json")
			pid := fieldTo.Tag.Get("json")
			if (j != "") && (pid == "") {
				pid = j
			} else if (pid != "") && (j == "") {
				j = pid
			}

			name := fieldTo.Tag.Get("PointName")
			if name == "" {
				name = PointToName(pid)
			}

			device := fieldTo.Tag.Get("PointDevice")
			if device == "" {
				device = "virtual"
			}

			var ignore bool
			if fieldTo.Tag.Get("PointIgnore") != "" {
				ignore = true
			}

			fullName := JoinDevicePoint(device, pid)

			unit := fieldTo.Tag.Get("PointUnit")
			var uv UnitValue
			if unit == "" {
				bar := fieldTo.Type.Name()
				fmt.Printf("bar:%v\n", bar)
				fmt.Println("")
				switch bar {
					case "int":
						uv.Unit = "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)
					case "int32":
						uv.Unit = "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)
					case "int64":
						uv.Unit = "integer"
						uv.ValueInt = valueTo.(int64)
						uv.Value = strconv.FormatInt(uv.ValueInt, 10)

					case "float32":
						uv.Unit = "float"
						uv.ValueFloat = float64(valueTo.(float32))
						uv.Value = Float64ToString(uv.ValueFloat)
					case "float64":
						uv.Unit = "float"
						uv.ValueFloat = valueTo.(float64)
						uv.Value = Float64ToString(uv.ValueFloat)

					case "string":
						uv.Unit = "string"
						uv.Value = valueTo.(string)

					case "UnitValue":
						fallthrough
					case "api.UnitValue":
						uv = valueTo.(UnitValue)
						uv = uv.UnitValueFix()

					default:
						ignore = true
				}
			}

			if ignore {
				continue
			}

			// foo := apiReflect.DataStructure {
			// 	Json:        j,
			// 	PointDevice: device,
			// 	PointId:     pid,
			// 	PointName:   name,
			// 	PointUnit:   uv.Unit,
			// 	PointType:   fieldTo.Tag.Get("PointType"),
			// }
			// ret[fieldTo.Name] = foo

			p := Point {
				EndPoint:  endpoint,
				FullId:    fullName,
				PsKey:     device,
				Id:        pid,
				GroupName: "",
				Name:      name,
				Unit:      uv.Unit,
				Type:      fieldTo.Tag.Get("PointType"),
				Valid:     true,
			}
			dm.AddEntry(p, now, uv.Value)

			alias := fieldTo.Tag.Get("PointAlias")
			if alias != "" {
				p.FullId = NameDevicePoint(device, alias)
				p.Id = alias
				dm.AddEntry(p, now, uv.Value)
			}

			// spew.Dump(&foo)
			// fieldVo := vo.Field(i)
			// value := fmt.Sprintf("%v", fieldVo.Interface())
			// if value == "" {
			// 	err = errors.New(fmt.Sprintf("option '%s' is empty", fieldTo.Name))
			// 	break
			// }
		}
	}
}

func (dm *DataMap) GetEntry(entry string) DataEntry {
	ret := dm.Entries[entry]
	return ret
}

func (dm *DataMap) CopyEntry(entry string) *DataEntry {
	ret := dm.Entries[entry]
	return &ret
	// return &DataEntry {
	// 	Date:       DateTime{},
	// 	Point:      nil,
	// 	Value:      "",
	// 	ValueFloat: 0,
	// 	Index:      0,
	// }
}

func (dm *DataMap) GetFloatValue(entry string) float64 {
	return dm.Entries[entry].ValueFloat
}

func (dm *DataMap) LowerUpper(lower string, upper string) float64 {
	l := dm.GetEntry(lower)
	u := dm.GetEntry(upper)

	if l.ValueFloat > 0 {
		return 0 - l.ValueFloat
	}
	return u.ValueFloat
}

func (dm *DataMap) GetPercent(value string, max string) float64 {
	v := dm.GetEntry(value)
	m := dm.GetEntry(max)
	return GetPercent(v.ValueFloat, m.ValueFloat)
}

func (dm *DataMap) GetValue(refname string) float64 {
	v := dm.GetEntry(refname)
	return v.ValueFloat
}


// func (dm *DataMap) FromRefAddState(refname string, psId string, point string, name string) {
// 	v := dm.GetEntry(refname)
// 	dm.Entries[point] = v.CreateState(psId, point, name)
// 	dm.Order = append(dm.Order, point)
// }

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


func (dm *DataMap) Add(point string, de DataEntry) {
	for range Only.Once {
		// point := de.ValueType.Id
		de.Index = len(dm.Entries)
		dm.Entries[point] = de
		dm.Order = append(dm.Order, point)

		if _, ok := Points[point]; ok {
			break
		}
		Points[point] = *de.Point
	}
}

func (dm *DataMap) AddEntry(point Point, date DateTime, value string) {
	for range Only.Once {
		unit := point.Unit	// Save unit.

		// Match to a previously defined point.
		p := GetPoint(point.PsKey, point.Id)
		if p == nil {
			point = *p
		}

		if point.PsKey == "" {
			point.PsKey = "virtual"
		}
		if point.Name == "" {
			point.Name = PointToName(point.Id)
		}
		if point.FullId == "" {
			point.FullId = JoinDevicePoint(point.PsKey, point.Id)
		}
		ref := CreateUnitValue(value, unit)
		point.Unit = ref.Unit
		point.Valid = true

		// foo := Point {
		// 	PsKey:     psId,
		// 	Id:        "total_income",
		// 	Unit:      p.TotalIncome.Unit,
		// 	Type:      PointTypeTotal,
		// }
		//
		// p := GetPoint(point.PsKey, point.Id)
		// if p == nil {
		// 	fmt.Printf("Found point already: %s.%s\n", p.PsKey, p.Id)
		// 	fmt.Println("&point")
		// 	spew.Dump(&point)
		// 	fmt.Println("&p")
		// 	spew.Dump(&p)
		// 	// dm.Add(point.Id, CreateDataEntryUnitValue(date, point.PsKey, point.Id, point.Name, ref))
		// 	// if p.Name == "" {
		// 	// 	p.Name = PointToName(point.Id)
		// 	// }
		// 	// p = CreatePoint(psId, point, name, value.Unit)
		// 	// break
		// }

		dm.Add(JoinDevicePoint(point.EndPoint, point.Id), DataEntry {
			EndPoint:   point.EndPoint,
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

func (dm *DataMap) FromRefAddAlias2(refname string, psId string, point string, name string) {
	de := dm.GetEntry(refname)
	dm.Add(point, de.CreateAlias(psId, point, name))
}

func (dm *DataMap) AddEntryFromRef(refPoint Point, point Point, date DateTime, value string) {
	for range Only.Once {
		p := GetPoint(refPoint.PsKey, refPoint.Id)
		if p != nil {
			fmt.Printf("Found point already: %s.%s\n", p.PsKey, p.Id)
			fmt.Println("&point")
			spew.Dump(&point)
			fmt.Println("&p")
			spew.Dump(&p)
			break
		}

		if point.PsKey == "" {
			point.PsKey = "virtual"
		}
		if point.Name == "" {
			point.Name = PointToName(point.Id)
		}
		if point.FullId == "" {
			point.FullId = JoinDevicePoint(point.PsKey, point.Id)
		}
		ref := CreateUnitValue(value, point.Unit)
		point.Unit = ref.Unit
		point.Valid = true

		dm.Add(point.Id, DataEntry {
			Date:       date,
			Point:      p,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}
}


func JoinDevicePoint(device string, point string) string {
	var ret string
	for range Only.Once {
		if device == "" {
			device = "virtual"
		}
		ret = device + "." + point
	}
	return ret
}

func (dm *DataMap) AddUnitValue(endpoint string, psId string, point string, name string, date DateTime, ref UnitValue) {
	for range Only.Once {
		if endpoint == "" {
			endpoint = apiReflect.GetCallerPackage(2)
		}

		ref = ref.UnitValueFix()

		p := GetPoint(psId, point)
		if p == nil {
			// No UV found. Create one.
			dm.Add(point, CreateDataEntryUnitValue(date, psId, point, name, ref))
			break
		}

		if p.Unit == "" {
			p.Unit = ref.Unit
		}
		if p.Name == "" {
			p.Name = name
		}
		if p.Name == "" {
			p.Name = PointToName(point)
		}
		dm.Add(point, DataEntry {
			Date:       date,
			Point:      p,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}
}

func (dm *DataMap) AddFloat(psId string, point string, name string, date DateTime, value float64) {
	for range Only.Once {
		fvs := Float64ToString(value)
		p := GetPoint(psId, point)
		if p == nil {
			// No UV found. Create one.
			dm.Add(point, CreateDataEntryUnitValue(date, psId, point, name, CreateUnitValue(fvs, "float")))
			break
		}
		ref := CreateUnitValue(fvs, p.Unit)
		if ref.Unit != p.Unit {
			fmt.Printf("OOOPS: Unit mismatch - %f %s != %f %s\n", value, p.Unit, ref.ValueFloat, ref.Unit)
			p.Unit = ref.Unit
		}
		dm.Add(point, DataEntry {
			Date:       date,
			Point:      p,
			Value:      ref.Value,
			ValueFloat: ref.ValueFloat,
		})
	}

	de := CreateDataEntryUnitValue(date, psId, point, name, UnitValue {
		Unit:       "float",
		Value:      fmt.Sprintf("%f", value),
		ValueFloat: 0,
	})
	dm.Add(point, de)
}

func (dm *DataMap) AddString(psId string, point string, name string, date DateTime, value string) {
	dm.Add(point, CreateDataEntryString(date, psId, point, name, value))
}

func (dm *DataMap) AddInt(psId string, point string, name string, date DateTime, value int64) {
	de := CreateDataEntryUnitValue(date, psId, point, name, UnitValue {
		Unit:       "int",
		Value:      fmt.Sprintf("%d", value),
		ValueFloat: float64(value),
	})
	dm.Add(point, de)
}


// func (dm *DataMap) AddVirtualAliasFromRef(refname string, point string, name string) {
// 	de := dm.GetEntry(refname)
// 	dm.Add(point, de.CreateAlias("virtual", point, name))
// }

func (dm *DataMap) FromRefAddAlias(refname string, psId string, point string, name string) {
	de := dm.GetEntry(refname)
	dm.Add(point, de.CreateAlias(psId, point, name))
}

func (dm *DataMap) FromRefAddState(refname string, psId string, point string, name string) {
	v := dm.GetEntry(refname)
	dm.Add(point, v.CreateState(psId, point, name))
}

func (dm *DataMap) FromRefAddFloat(refname string, psId string, point string, name string, value float64) {
	de := dm.GetEntry(refname)
	dm.Add(point, de.CreateFloat(psId, point, name, value))
}

func (dm *DataMap) CopyPoint(refname string, psId string, point string, name string, value float64) {
	de := dm.GetEntry(refname)
	dm.Add(point, de.CreateFloat(psId, point, name, value))
}


func (de *DataEntry) CreateAlias(psId string, point string, name string) DataEntry {
	if name == "" {
		name = PointToName(point)
	}

	de.Point.FullId = NameDevicePoint(psId, point)
	de.Point.PsKey = psId
	de.Point.Id = point
	de.Point.Name = name
	de.Point.GroupName = psId
	de.Point.Valid = true
	de.Index = 0

	return *de
}

func (de *DataEntry) CreateFloat(psId string, point string, name string, value float64) DataEntry {
	if name == "" {
		name = PointToName(point)
	}

	de2 := de.CreateAlias(psId, point, name)
	uv := CreateUnitValue(Float64ToString(value), de2.Point.Unit)
	de2.Value = uv.Value
	de2.ValueFloat = uv.ValueFloat

	return de2
	// if name == "" {
	// 	name = PointToName(point)
	// }
	//
	// return DataEntry {
	// 	Date:       de.Date,
	// 	Value:      Float64ToString(v),
	// 	ValueFloat: v,
	// 	Point:      &Point {
	// 		FullId:     NameDevicePoint(psId, point),
	// 		GroupName:  psId,
	// 		PsKey: psId,
	// 		Id:    point,
	// 		Name:  name,
	// 		Unit:  u,
	// 		Type:  "",
	// 		Valid: true,
	// 	},
	// 	Index:          0,
	// }
}

func (de *DataEntry) CreateState(psId string, point string, name string) DataEntry {
	if name == "" {
		name = PointToName(point)
	}

	de2 := de.CreateAlias(psId, point, name)
	de2.Value = fmt.Sprintf("%v", IsActive(de.ValueFloat))
	de2.ValueFloat = 0
	de2.Point.Unit = "binary"

	return de2

	// return DataEntry {
	// 	Date:       de.Date,
	// 	Value:      fmt.Sprintf("%v", IsActive(de.ValueFloat)),
	// 	ValueFloat: 0,
	// 	Point:      &Point {
	// 		FullId:         NameDevicePoint(psId, point),
	// 		PsKey: psId,
	// 		Id:    point,
	// 		GroupName:  psId,
	// 		Name:  name,
	// 		Unit:  "binary",
	// 		Type:  "",
	// 		Valid: true,
	// 	},
	// 	Index:      0,
	// }
}

func (de *DataEntry) UpdateMeta(date *DateTime, psId string, point string, name string) {

	if date != nil {
		de.Date = *date
	}
	if name == "" {
		name = PointToName(point)
	}

	de.Point.FullId = NameDevicePoint(psId, point)
	de.Point.PsKey = psId
	de.Point.Id = point
	de.Point.Name = name
	de.Point.GroupName = psId
	de.Index = 0
}


func CreateDataEntryActive(date DateTime, psId string, point string, name string, value float64) DataEntry {
	p := GetPoint(psId, point)
	if p == nil {
		if name == "" {
			name = PointToName(point)
		}
		p = CreatePoint(psId, point, name, "state")
	}

	return DataEntry {
		Date:       date,
		Value:      fmt.Sprintf("%v", IsActive(value)),
		ValueFloat: 0,
		Point:      p,
		Index:      0,
	}
}

func CreateDataEntryString(date DateTime, psId string, point string, name string, value string) DataEntry {
	p := GetPoint(psId, point)
	if p == nil {
		if name == "" {
			name = PointToName(point)
		}
		p = CreatePoint(psId, point, name, "string")
	}

	return DataEntry {
		Date:       date,
		Value:      value,
		ValueFloat: 0,
		Point:      p,
		Index:      0,
	}
}

func CreateDataEntryUnitValue(date DateTime, psId string, point string, name string, value UnitValue) DataEntry {
	value = value.UnitValueFix()

	p := GetPoint(psId, point)
	if p == nil {
		if name == "" {
			name = PointToName(point)
		}
		p = CreatePoint(psId, point, name, value.Unit)
	}

	return DataEntry {
		Date:       date,
		Value:      value.Value,
		ValueFloat: value.ValueFloat,
		Point:      p,
		Index:      0,
	}
}

func CreatePoint(psId string, point string, name string, unit string) *Point {
	if name == "" {
		name = PointToName(point)
	}

	return &Point {
			FullId:    NameDevicePoint(psId, point),
			PsKey:     psId,
			Id:        point,
			GroupName: psId,
			Name:      name,
			Unit:      unit,
			Type:      "",
			Valid:     true,
		}
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

func (ref *UnitValue) UnitValueToPoint(psId string, point string, name string) *Point {
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
		name = PointToName(point)
	}

	vt := GetPoint(psId, point)
	if !vt.Valid {
		vt = &Point {
			PsKey: psId,
			Id:    point,
			Name:  name,
			Unit:  uv.Unit,
			Type:  "PointTypeInstant",
			Valid: true,
		}
	}

	return vt
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