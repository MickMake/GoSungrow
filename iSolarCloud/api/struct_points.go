package api

import (
	"GoSungrow/Only"
	"fmt"
	"strconv"
	"strings"
	"time"
)


const (
	PointTypeInstant = "instant"
	PointTypeDaily = "daily"
	PointTypeMonthly = "monthly"
	PointTypeYearly = "yearly"
	PointTypeTotal = "total"
)


type Point struct {
	PsKey       string
	Id          string
	Description string
	Unit        string
	Type        string
}
type PointsMap map[string]Point


func (pm *PointsMap) Resolve(point string) *Point {
	if ret, ok := (*pm)[point]; ok {
		return &ret
	}
	return nil
}

func ResolvePoint(point string) *Point {
	return Points.Resolve(point)
}

// p1
// p1001
// p1002
// p1005
// p1302
// p14
// p2
// p202
// p210
// p24
// p24001
// p24004
// p66
// p8018
// p8030
// p8031
// p8062
// p8063
// p81002
// p81022
// p81023
// p81202
// p81203
// p83011
// p83012
// p83039
// p83046
// p83048
// p83049
// p83050
// p83051
// p83054
// p83055
// p83067
// p83070
// p83072
// p83075
// p83076
// p83077
// p83080
// p83081
// p83086
// p83087
// p83088
// p83089
// p83095
// p83096
// p83101
// p83118
// p83119
// p83120
// p83121
// p83122
// p83127
// p83202
// p83532

// Points Discovered points from the API
var Points = PointsMap {

	"p13001": { PsKey: "1129147_14_1_1", Id: "p13001", Description: "MPPT1 Voltage", Unit: "V", Type: PointTypeInstant },

	"p13012": { PsKey: "1129147_14_1_1", Id: "p13012", Description: "Total Reactive Power", Unit: "kvar", Type: PointTypeDaily },

	"p13105": { PsKey: "1129147_14_1_1", Id: "p13105", Description: "MPPT2 Voltage", Unit: "V", Type: PointTypeInstant },

	"p13122": { PsKey: "1129147_14_1_1", Id: "p13122", Description: "Daily Feed-in Energy", Unit: "kWh", Type: PointTypeDaily },

	"p13125": { PsKey: "1129147_14_1_1", Id: "p13125", Description: "Total Feed-in Energy", Unit: "kWh", Type: PointTypeTotal },

	"p13138": { PsKey: "1129147_14_1_1", Id: "p13138", Description: "Battery Voltage", Unit: "V", Type: PointTypeInstant },

	"p13144": { PsKey: "1129147_14_1_1", Id: "p13144", Description: "Daily Self-consumption Rate", Unit: "%", Type: PointTypeDaily },

	"p13157": { PsKey: "1129147_14_1_1", Id: "p13157", Description: "Phase A Voltage", Unit: "V", Type: PointTypeInstant },

	"p13158": { PsKey: "1129147_14_1_1", Id: "p13158", Description: "Phase B Voltage", Unit: "V", Type: PointTypeInstant },

	"p13159": { PsKey: "1129147_14_1_1", Id: "p13159", Description: "Phase C Voltage", Unit: "V", Type: PointTypeInstant },

	"p13161": { PsKey: "1129147_14_1_1", Id: "p13161", Description: "Bus Voltage", Unit: "V", Type: PointTypeInstant },
	"p13173": { PsKey: "1129147_14_1_1", Id: "p13173", Description: "Daily Feed-in Energy (PV)", Unit: "kWh", Type: PointTypeDaily },
	"p13175": { PsKey: "1129147_14_1_1", Id: "p13175", Description: "Total Feed-in Energy (PV)", Unit: "kWh", Type: PointTypeTotal },
	"p13002": { PsKey: "1129147_14_1_1", Id: "p13002", Description: "MPPT1 Current", Unit: "A", Type: PointTypeInstant },
	"p13003": { PsKey: "1129147_14_1_1", Id: "p13003", Description: "Total DC Power", Unit: "kW" },
	"p13007": { PsKey: "1129147_14_1_1", Id: "p13007", Description: "Grid Frequency", Unit: "Hz", Type: PointTypeInstant },
	"p13008": { PsKey: "1129147_14_1_1", Id: "p13008", Description: "Phase A Current", Unit: "A", Type: PointTypeInstant },
	"p13009": { PsKey: "1129147_14_1_1", Id: "p13009", Description: "Phase B Current", Unit: "A", Type: PointTypeInstant },
	"p13010": { PsKey: "1129147_14_1_1", Id: "p13010", Description: "Phase C Current", Unit: "A", Type: PointTypeInstant },
	"p13011": { PsKey: "1129147_14_1_1", Id: "p13011", Description: "Total Active Power", Unit: "kW" },
	"p13013": { PsKey: "1129147_14_1_1", Id: "p13013", Description: "Total Power Factor", Unit: "" },
	"p13018": { PsKey: "1129147_14_1_1", Id: "p13018", Description: "Total Apparent Power", Unit: "VA" },
	"p13019": { PsKey: "1129147_14_1_1", Id: "p13019", Description: "Internal Air Temperature", Unit: "℃", Type: PointTypeInstant },
	"p13028": { PsKey: "1129147_14_1_1", Id: "p13028", Description: "Daily Battery Charging Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13029": { PsKey: "1129147_14_1_1", Id: "p13029", Description: "Daily Battery Discharging Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13034": { PsKey: "1129147_14_1_1", Id: "p13034", Description: "Total Battery Charging Energy", Unit: "kWh" , Type: PointTypeTotal },
	"p13035": { PsKey: "1129147_14_1_1", Id: "p13035", Description: "Total Battery Discharging Energy", Unit: "kWh" , Type: PointTypeTotal },
	"p13106": { PsKey: "1129147_14_1_1", Id: "p13106", Description: "MPPT2 Current", Unit: "A", Type: PointTypeInstant },
	"p13112": { PsKey: "1129147_14_1_1", Id: "p13112", Description: "Daily PV Yield", Unit: "kWh", Type: PointTypeDaily },
	"p13116": { PsKey: "1129147_14_1_1", Id: "p13116", Description: "Daily Load Energy Consumption from PV", Unit: "kWh", Type: PointTypeDaily },
	"p13119": { PsKey: "1129147_14_1_1", Id: "p13119", Description: "Total Load Active Power", Unit: "kW" },
	"p13121": { PsKey: "1129147_14_1_1", Id: "p13121", Description: "Total Export Active  Power", Unit: "kW" },
	"p13126": { PsKey: "1129147_14_1_1", Id: "p13126", Description: "Battery Charging Power", Unit: "kW" },
	"p13130": { PsKey: "1129147_14_1_1", Id: "p13130", Description: "Total Load Energy Consumption", Unit: "kWh" , Type: PointTypeTotal },
	"p13134": { PsKey: "1129147_14_1_1", Id: "p13134", Description: "Total PV Yield", Unit: "kWh" , Type: PointTypeTotal },
	"p13137": { PsKey: "1129147_14_1_1", Id: "p13137", Description: "Total Load Energy Consumption from PV", Unit: "kWh" , Type: PointTypeTotal },
	"p13139": { PsKey: "1129147_14_1_1", Id: "p13139", Description: "Battery Current", Unit: "A", Type: PointTypeInstant },
	"p13140": { PsKey: "1129147_14_1_1", Id: "p13140", Description: "Battery Capacity(kWh)", Unit: "kWh" },
	"p13141": { PsKey: "1129147_14_1_1", Id: "p13141", Description: "Battery Level (SOC)", Unit: "%", Type: PointTypeInstant },
	"p13142": { PsKey: "1129147_14_1_1", Id: "p13142", Description: "Battery Health (SOH)", Unit: "%", Type: PointTypeInstant },
	"p13143": { PsKey: "1129147_14_1_1", Id: "p13143", Description: "Battery Temperature", Unit: "℃", Type: PointTypeInstant },
	"p13147": { PsKey: "1129147_14_1_1", Id: "p13147", Description: "Daily Purchased Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13148": { PsKey: "1129147_14_1_1", Id: "p13148", Description: "Total Purchased Energy", Unit: "kWh", Type: PointTypeTotal },
	"p13149": { PsKey: "1129147_14_1_1", Id: "p13149", Description: "Purchased Power", Unit: "kW" },
	"p13150": { PsKey: "1129147_14_1_1", Id: "p13150", Description: "Battery Discharging Power", Unit: "kW" },
	"p13160": { PsKey: "1129147_14_1_1", Id: "p13160", Description: "Array Insulation Resistance", Unit: "kΩ", Type: PointTypeInstant },
	"p13162": { PsKey: "1129147_14_1_1", Id: "p13162", Description: "Max. Charging Current (BMS)", Unit: "A", Type: PointTypeInstant },
	"p13163": { PsKey: "1129147_14_1_1", Id: "p13163", Description: "Max. Discharging Current (BMS)", Unit: "A", Type: PointTypeInstant },
	"p13174": { PsKey: "1129147_14_1_1", Id: "p13174", Description: "Daily Battery Charging Energy from PV", Unit: "kWh", Type: PointTypeDaily },
	"p13176": { PsKey: "1129147_14_1_1", Id: "p13176", Description: "Total Battery Charging Energy from PV", Unit: "kWh", Type: PointTypeTotal },
	"p13199": { PsKey: "1129147_14_1_1", Id: "p13199", Description: "Daily Load Energy Consumption", Unit: "kWh", Type: PointTypeDaily },
	"p18062": { PsKey: "1129147_14_1_1", Id: "p18062", Description: "Phase A Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18063": { PsKey: "1129147_14_1_1", Id: "p18063", Description: "Phase B Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18064": { PsKey: "1129147_14_1_1", Id: "p18064", Description: "Phase C Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18065": { PsKey: "1129147_14_1_1", Id: "p18065", Description: "Phase A Backup Power", Unit: "kW" },
	"p18066": { PsKey: "1129147_14_1_1", Id: "p18066", Description: "Phase B Backup Power", Unit: "kW" },
	"p18067": { PsKey: "1129147_14_1_1", Id: "p18067", Description: "Phase C Backup Power", Unit: "kW" },
	"p18068": { PsKey: "1129147_14_1_1", Id: "p18068", Description: "Total Backup Power", Unit: "kW" },
	"p83001": { PsKey: "1129147_11_0_0", Id: "p83001", Description: "Inverter AC Power Normalization", Unit: "kW/kWp" },
	"p83002": { PsKey: "1129147_11_0_0", Id: "p83002", Description: "Inverter AC Power", Unit: "kW" },
	"p83004": { PsKey: "1129147_11_0_0", Id: "p83004", Description: "Inverter Total Yield", Unit: "kWh" },
	"p83005": { PsKey: "1129147_11_0_0", Id: "p83005", Description: "Daily Equivalent Hours of Meter", Unit: "h", Type: PointTypeDaily },
	"p83006": { PsKey: "1129147_11_0_0", Id: "p83006", Description: "Meter Daily Yield", Unit: "kWh" },
	"p83007": { PsKey: "1129147_11_0_0", Id: "p83007", Description: "Meter PR", Unit: "%", Type: PointTypeInstant },
	"p83008": { PsKey: "1129147_11_0_0", Id: "p83008", Description: "Daily Equivalent Hours of Inverter", Unit: "h", Type: PointTypeDaily },
	"p83009": { PsKey: "1129147_11_0_0", Id: "p83009", Description: "Daily Yield by Inverter", Unit: "kWh", Type: PointTypeDaily },
	"p83010": { PsKey: "1129147_11_0_0", Id: "p83010", Description: "Inverter PR", Unit: "%", Type: PointTypeInstant },
	"p83013": { PsKey: "1129147_11_0_0", Id: "p83013", Description: "Daily Irradiation", Unit: "Wh/m2" },
	"p83016": { PsKey: "1129147_11_0_0", Id: "p83016", Description: "Plant Ambient Temperature", Unit: "℃", Type: PointTypeInstant },
	"p83017": { PsKey: "1129147_11_0_0", Id: "p83017", Description: "Plant Module Temperature", Unit: "℃", Type: PointTypeInstant },
	"p83018": { PsKey: "1129147_11_0_0", Id: "p83018", Description: "Daily Yield (Theoretical)", Unit: "kWh", Type: PointTypeDaily },
	"p83019": { PsKey: "1129147_11_0_0", Id: "p83019", Description: "Power/Installed Power of Plant", Unit: "%", Type: PointTypeInstant },
	"p83020": { PsKey: "1129147_11_0_0", Id: "p83020", Description: "Meter Total Yield", Unit: "kWh" },
	"p83021": { PsKey: "1129147_11_0_0", Id: "p83021", Description: "Accumulative Power Consumption by Meter", Unit: "kWh" },
	"p83022": { PsKey: "1129147_11_0_0", Id: "p83022", Description: "Daily Yield of Plant", Unit: "kWh", Type: PointTypeDaily },
	"p83023": { PsKey: "1129147_11_0_0", Id: "p83023", Description: "Plant PR", Unit: "%", Type: PointTypeInstant },
	"p83024": { PsKey: "1129147_11_0_0", Id: "p83024", Description: "Plant Total Yield", Unit: "kWh" },
	"p83025": { PsKey: "1129147_11_0_0", Id: "p83025", Description: "Plant Equivalent Hours", Unit: "h" },
	"p83032": { PsKey: "1129147_11_0_0", Id: "p83032", Description: "Meter AC Power", Unit: "kW" },
	"p83033": { PsKey: "1129147_11_0_0", Id: "p83033", Description: "Plant Power", Unit: "kW" },
	"p83097": { PsKey: "1129147_11_0_0", Id: "p83097", Description: "Daily Load Energy Consumption from PV", Unit: "kWh", Type: PointTypeDaily },
	"p83100": { PsKey: "1129147_11_0_0", Id: "p83100", Description: "Total Load Energy Consumption from PV", Unit: "kWh" },
	"p83102": { PsKey: "1129147_11_0_0", Id: "p83102", Description: "Daily Purchased Energy", Unit: "kWh", Type: PointTypeDaily },
	"p83105": { PsKey: "1129147_11_0_0", Id: "p83105", Description: "Total Purchased Energy", Unit: "kWh" },
	"p83106": { PsKey: "1129147_11_0_0", Id: "p83106", Description: "Load Power", Unit: "kW" },
	"p83124": { PsKey: "1129147_11_0_0", Id: "p83124", Description: "Total Load Energy Consumption", Unit: "MWh" },
	"p83128": { PsKey: "1129147_11_0_0", Id: "p83128", Description: "Total Active Power of Optical Storage", Unit: "kW" },
	"p83129": { PsKey: "1129147_11_0_0", Id: "p83129", Description: "Battery SOC", Unit: "%", Type: PointTypeInstant },
	"p83233": { PsKey: "1129147_11_0_0", Id: "p83233", Description: "Total field maximum rechargeable power", Unit: "MW" },
	"p83234": { PsKey: "1129147_11_0_0", Id: "p83234", Description: "Total field maximum dischargeable power", Unit: "MW" },
	"p83235": { PsKey: "1129147_11_0_0", Id: "p83235", Description: "Total field chargeable energy", Unit: "MWh" },
	"p83236": { PsKey: "1129147_11_0_0", Id: "p83236", Description: "Total field dischargeable energy", Unit: "MWh" },
	"p83237": { PsKey: "1129147_11_0_0", Id: "p83237", Description: "Total field energy storage maximum reactive power", Unit: "MW" },
	"p83238": { PsKey: "1129147_11_0_0", Id: "p83238", Description: "Total field energy storage active power", Unit: "MW" },
	"p83239": { PsKey: "1129147_11_0_0", Id: "p83239", Description: "Total field reactive power", Unit: "Mvar" },
	"p83241": { PsKey: "1129147_11_0_0", Id: "p83241", Description: "Total field charge capacity", Unit: "MWh" },
	"p83242": { PsKey: "1129147_11_0_0", Id: "p83242", Description: "Total field discharge capacity", Unit: "MWh" },
	"p83243": { PsKey: "1129147_11_0_0", Id: "p83243", Description: "Total field daily charge capacity", Unit: "MWh" },
	"p83244": { PsKey: "1129147_11_0_0", Id: "p83244", Description: "Total field daily discharge capacity", Unit: "MWh" },
	"p83252": { PsKey: "1129147_11_0_0", Id: "p83252", Description: "Battery Level (SOC)", Unit: "%", Type: PointTypeInstant },
	"p83419": { PsKey: "1129147_11_0_0", Id: "p83419", Description: "Daily Highest Inverter Power/Inverter Installed Capacity", Unit: "%" },
	"p83420": { PsKey: "1129147_11_0_0", Id: "p83420", Description: "Current Power/Inverter Installed Capacity", Unit: "%", Type: PointTypeInstant },
	"p83549": { PsKey: "1129147_11_0_0", Id: "p83549", Description: "Grid active power", Unit: "kW" },

	"p23014": { PsKey: "1129147_22_247_1", Id: "p23014", Description: "WLAN Signal Strength", Unit: "" },

}


func GetPoint(device string, point string) *Point {
	return Points.Get(device, point)
}

func GetPointInt(device string, point int64) *Point {
	return Points.Get(device, strconv.FormatInt(point, 10))
}

func GetDevicePoint(devicePoint string) *Point {
	return Points.GetDevicePoint(devicePoint)
}

// func GetPointName(device string, point int64) string {
// 	return fmt.Sprintf("%s.%d", device, point)
// }

func NameDevicePointInt(device string, point int64) string {
	return fmt.Sprintf("%s.%d", device, point)
}

func NameDevicePoint(device string, point string) string {
	return fmt.Sprintf("%s.%s", device, point)
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

func (pm PointsMap) Get(device string, point string) *Point {
	dp := device + ".p" + strings.TrimPrefix(point, "p")
	if p, ok := pm[dp]; ok {
		return &p
	}

	dp = "p" + strings.TrimPrefix(point, "p")
	if p, ok := pm[dp]; ok {
		return &p
	}

	return nil
}

func (pm PointsMap) GetDevicePoint(devicePoint string) *Point {
	ret := &Point{}
	for range Only.Once {
		s := strings.Split(devicePoint, ".")
		if len(s) < 2 {
			break
		}
		ret = pm.Get(s[0], s[1])
	}
	return ret
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


func UpperCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	return s
}


type Data struct {
	Entries []DataEntry
}

type DataEntry struct {
	Date           DateTime  `json:"date"`
	PointId        string    `json:"point_id"`
	PointGroupName string    `json:"point_group_name"`
	PointName      string    `json:"point_name"`
	Value          string    `json:"value"`
	Unit           string    `json:"unit"`
	ValueType      *Point    `json:"value_type"`
	Index          int       `json:"index"`
}
