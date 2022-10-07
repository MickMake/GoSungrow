package api

import (
	"GoSungrow/Only"
	"fmt"
	"strconv"
	"strings"
	"time"
)


type Point struct {
	// EndPoint  string            `json:"endpoint"`
	// FullId    PointId           `json:"full_id"`
	// Parent    ParentDevice      `json:"parent"`
	Parents   ParentDevices     `json:"parents"`
	Id        PointId           `json:"id"`
	GroupName string            `json:"group_name"`
	Name      string            `json:"name"`
	Unit      string            `json:"unit"`
	Type      string            `json:"type"`
	Valid     bool              `json:"valid"`
	States    map[string]string `json:"states"`
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
	return p.Type
}

func (p Point) IsInstant() bool {
	if p.Type == PointTypeInstant {
		return true
	}
	return false
}

func (p Point) IsDaily() bool {
	if p.Type == PointTypeDaily {
		return true
	}
	return false
}

func (p Point) IsMonthly() bool {
	if p.Type == PointTypeMonthly {
		return true
	}
	return false
}

func (p Point) IsYearly() bool {
	if p.Type == PointTypeYearly {
		return true
	}
	return false
}

func (p Point) IsTotal() bool {
	if p.Type == PointTypeTotal {
		return true
	}
	return false
}


func GetPoint(device string, point PointId) *Point {
	return Points.Get(device, point)
}

func GetPointInt(device string, point int64) *Point {
	return Points.Get(device, PointId(strconv.FormatInt(point, 10)))
}

func GetDevicePoint(devicePoint string) *Point {
	return Points.GetDevicePoint(devicePoint)
}

// func GetPointName(device string, point int64) string {
// 	return fmt.Sprintf("%s.%d", device, point)
// }

func NameDevicePointInt(device string, point int64) PointId {
	return PointId(fmt.Sprintf("%s.%d", device, point))
}

func NameDevicePoint(device string, point PointId) PointId {
	return PointId(fmt.Sprintf("%s.%s", device, point))
}

func SetPoint(point PointId) PointId {
	for range Only.Once {
		p := strings.TrimPrefix(string(point), "p")
		_, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			point = PointId("p" + p)
			break
		}
	}
	return point
}

func SetPointInt(point int64) PointId {
	return PointId("p" + strconv.FormatInt(point, 10))
}

func PointToName(s PointId) string {
	ret := string(s)
	ret = CleanString(ret)
	ret = strings.ReplaceAll(ret, "_", " ")
	ret = strings.Title(ret)
	return ret
}

func (p *PointId) Fix() PointId {
	for range Only.Once {
		p := strings.TrimPrefix(string(*p), "p")
		_, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			p = "p" + p
			break
		}
	}
	return *p
}


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
