package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"

	"strings"
	"time"
)


type Point struct {
	Parents   ParentDevices      `json:"parents,omitempty"`
	Id        valueTypes.PointId `json:"id,omitempty"`
	GroupName string             `json:"group_name,omitempty"`
	Name      string             `json:"name,omitempty"`
	Unit      string             `json:"unit,omitempty"`
	TimeSpan  string             `json:"time_span,omitempty"`
	ValueType string             `json:"value_type,omitempty"`
	Valid     bool               `json:"valid,omitempty"`
	States    map[string]string  `json:"states,omitempty"`
}


func (p *Point) FixUnitType() Point {
	for range Only.Once {
		switch p.Unit {
			case "Wh":
				fallthrough
			case "kWh":
				fallthrough
			case "MWh":
				p.ValueType = "Energy"

			case "kWp":
				fallthrough
			case "W":
				fallthrough
			case "kW":
				fallthrough
			case "MW":
				p.ValueType = "Power"

			case "AUD":
				p.ValueType = "Currency"

			case "kg":
				p.ValueType = "Weight"

			case "mV":
				fallthrough
			case "V":
				p.ValueType = "Voltage"

			case "A":
				p.ValueType = "Current"

			case "Hz":
				p.ValueType = "Frequency"

			case "kvar":
				p.ValueType = "Reactive Power"

			case "kΩ":
				p.ValueType = "Resistance"

			case "%":
				p.ValueType = "Percent"

			case "℃":
				p.ValueType = "Temperature"
		}
	}
	return *p
}

func (p *Point) WhenReset() string {
	var ret string

	for range Only.Once {
		var err error
		now := time.Now()

		switch {
			case p.IsInstant():
				ret = ""

			case p.IsDaily():
				now, err = time.Parse("2006-01-02T15:04:05", now.Format("2006-01-02") + "T00:00:00")
				// ret = fmt.Sprintf("%d", now.Unix())
				ret = now.Format("2006-01-02T15:04:05") + ""

			case p.IsMonthly():
				now, err = time.Parse("2006-01-02T15:04:05", now.Format("2006-01") + "-01T00:00:00")
				ret = fmt.Sprintf("%d", now.Unix())
				ret = now.Format("2006-01-02T15:04:05") + ""

			case p.IsYearly():
				now, err = time.Parse("2006-01-02T15:04:05", now.Format("2006") + "-01-01T00:00:00")
				ret = fmt.Sprintf("%d", now.Unix())
				ret = now.Format("2006-01-02T15:04:05") + ""

			case p.IsTotal():
				// ret = "0"
				ret = "1970-01-01T00:00:00"

			default:
				// ret = "0"
				ret = "1970-01-01T00:00:00"
		}
		if err != nil {
			now := time.Now()
			ret = fmt.Sprintf("%d", now.Unix())
		}
	}

	return ret
}

func (p Point) String() string {
	return fmt.Sprintf("Id:%s\tName:%s\tUnits:%s\tTimespan:%s", p.Id, p.Name, p.Unit, p.TimeSpan)
	// return p.TimeSpan
}

func (p Point) IsInstant() bool {
	if p.TimeSpan == PointTimeSpanInstant {
		return true
	}
	return false
}

func (p Point) IsDaily() bool {
	if p.TimeSpan == PointTimeSpanDaily {
		return true
	}
	return false
}

func (p Point) IsMonthly() bool {
	if p.TimeSpan == PointTimeSpanMonthly {
		return true
	}
	return false
}

func (p Point) IsYearly() bool {
	if p.TimeSpan == PointTimeSpanYearly {
		return true
	}
	return false
}

func (p Point) IsTotal() bool {
	if p.TimeSpan == PointTimeSpanTotal {
		return true
	}
	return false
}


func GetPoint(point string) *Point {
	return Points.Get(point)
}

func GetDevicePoint(devicePoint string) *Point {
	return Points.GetDevicePoint(devicePoint)
}

// func GetPointInt(device string, point int64) *Point {
// 	return Points.Get(device, valueTypes.PointId(strconv.FormatInt(point, 10)))
// }
//
// func GetPointName(device string, point int64) string {
// 	return fmt.Sprintf("%s.%d", device, point)
// }
//
// func NameDevicePointInt(device string, point int64) valueTypes.PointId {
// 	return valueTypes.PointId(fmt.Sprintf("%s.%d", device, point))
// }
//
// func SetPointInt(point int64) valueTypes.PointId {
// 	// return valueTypes.PointId("p" + strconv.FormatInt(point, 10))
// 	return valueTypes.SetPointIdValue(point)
// }
//
// func NameDevicePoint(point valueTypes.DataPoint) valueTypes.DataPoint {
// 	return point
// }
//
// func SetPoint(point valueTypes.DataPoint) valueTypes.DataPoint {
// 	// for range Only.Once {
// 	// 	p := strings.TrimPrefix(string(point), "p")
// 	// 	_, err := strconv.ParseInt(p, 10, 64)
// 	// 	if err == nil {
// 	// 		point = valueTypes.PointId("p" + p)
// 	// 		break
// 	// 	}
// 	// }
// 	// return point
// 	return point
// }


type ParentDevice struct {
	Key  string `json:"ps_key"`
	PsId string `json:"ps_id"`
	Type string `json:"parent_type"`
	Code string  `json:"parent_code"`
}
type ParentDevices struct {
	Map map[string]*ParentDevice
	Index []string
}

func (pd *ParentDevices) Add(device ParentDevice) {
	for range Only.Once {
		if len(pd.Map) == 0 {
			pd.Map = make(map[string]*ParentDevice)
		}
		if device.Type == "" {
			device.Split()
		}
		pd.Index = append(pd.Index, device.Key)
		if _, ok := pd.Map[device.Key]; ok {
			break
		}
		pd.Map[device.Key] = &device
	}
}

func (pd ParentDevices) String() string {
	var ret string
	for _, l := range pd.Index {
		ret += fmt.Sprintf("%s\n", pd.Map[l].Key)
	}
	ret = strings.TrimSuffix(ret, "\n")
	return ret
}

func (pd *ParentDevices) Keys() string {
	var ret string
	for _, l := range pd.Index {
		ret += fmt.Sprintf("%s\n", pd.Map[l].Key)
	}
	ret = strings.TrimSuffix(ret, "\n")
	return ret
}

func (pd *ParentDevices) PsIds() string {
	var ret string
	for _, l := range pd.Index {
		ret += fmt.Sprintf("%s\n", pd.Map[l].PsId)
	}
	ret = strings.TrimSuffix(ret, "\n")
	return ret
}

func (pd *ParentDevices) Codes() string {
	var ret string
	for _, l := range pd.Index {
		ret += fmt.Sprintf("%s\n", pd.Map[l].Code)
	}
	ret = strings.TrimSuffix(ret, "\n")
	return ret
}

func (pd *ParentDevices) Types() string {
	var ret string
	for _, l := range pd.Index {
		ret += fmt.Sprintf("%s\n", pd.Map[l].Type)
	}
	ret = strings.TrimSuffix(ret, "\n")
	return ret
}


func (pd *ParentDevice) Set(key string) {
	for range Only.Once {
		pd.Key = key
	}
}

func (pd *ParentDevice) Split() {
	for range Only.Once {
		// if pd.Key == "virtual" {
		// 	break
		// }
		if pd.Key == "" {
			pd.Key = "virtual"
			break
		}

		if !strings.Contains(pd.Key, "_") {
			pd.PsId = pd.Key
			break
		}
		s := strings.Split(pd.Key, "_")
		if len(s) > 0 {
			pd.PsId = s[0]
		}
		if len(s) > 1 {
			pd.Type = s[1]
		}
		if len(s) > 2 {
			pd.Code = s[2]
		}
	}
}

// type ParentDevice struct {
// 	Key  string `json:"ps_key"`
// 	PsId string `json:"ps_id"`
// 	Type string `json:"parent_type"`
// 	Code string  `json:"parent_code"`
// }
// type ParentDevices struct {
// 	Map map[string]ParentDevice
// }
//
// func (pd *ParentDevices) Add(device ParentDevice) {
// 	if len(pd.Map) == 0 {
// 		pd.Map = make(map[string]ParentDevice)
// 	}
// 	if device.PsId == "" {
// 		device = device.Split()
// 	}
// 	pd.Map[device.Key] = device
// }
//
// func (pd ParentDevices) String() string {
// 	var ret string
// 	for _, l := range pd.Map {
// 		ret += fmt.Sprintf("%s\n", l.Key)
// 	}
// 	ret = strings.TrimSuffix(ret, "\n")
// 	return ret
// }
//
// func (pd *ParentDevices) PsIds() string {
// 	var ret string
// 	for _, l := range pd.Map {
// 		ret += fmt.Sprintf("%s\n", l.PsId)
// 	}
// 	ret = strings.TrimSuffix(ret, "\n")
// 	return ret
// }
//
// func (pd *ParentDevices) Codes() string {
// 	var ret string
// 	for _, l := range pd.Map {
// 		ret += fmt.Sprintf("%s\n", l.Code)
// 	}
// 	ret = strings.TrimSuffix(ret, "\n")
// 	return ret
// }
//
// func (pd *ParentDevices) Types() string {
// 	var ret string
// 	for _, l := range pd.Map {
// 		ret += fmt.Sprintf("%s\n", l.Type)
// 	}
// 	ret = strings.TrimSuffix(ret, "\n")
// 	return ret
// }
//
// func (pd *ParentDevice) Split() ParentDevice {
// 	for range Only.Once {
// 		if pd.Key == "" {
// 			pd.Key = "virtual"
// 			break
// 		}
//
// 		if !strings.Contains(pd.Key, "_") {
// 			pd.PsId = pd.Key
// 			break
// 		}
// 		s := strings.Split(pd.Key, "_")
// 		if len(s) > 0 {
// 			pd.PsId = s[0]
// 		}
// 		if len(s) > 1 {
// 			pd.Type = s[1]
// 		}
// 		if len(s) > 2 {
// 			pd.Code = s[2]
// 		}
// 	}
// 	return *pd
// }
