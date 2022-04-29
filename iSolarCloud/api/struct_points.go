package api

import (
	"GoSungrow/Only"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	VirtualPsId = "virtual"
)


type Point struct {
	EndPoint    string  `json:"endpoint"`
	FullId      string	`json:"full_id"`
	PsKey       string	`json:"ps_key"`
	Id   		string	`json:"id"`
	GroupName 	string  `json:"group_name"`
	Name 		string	`json:"name"`
	Unit 		string	`json:"unit"`
	Type        string	`json:"type"`
	Valid       bool	`json:"valid"`
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

	// Added manually
	// "p83022y": { PsKey: "virtual", Id: "p83022", Name: "", Unit: "kWh", Type: PointTypeInstant },
	// "p83076": { PsKey: "virtual", Id: "p83076", Name: "Pv Power", Unit: "kWh", Type: PointTypeInstant },
	// "p83077": { PsKey: "virtual", Id: "p83077", Name: "Pv Energy", Unit: "kWh", Type: PointTypeInstant },
	// "p83081": { PsKey: "virtual", Id: "p83081", Name: "Es Power", Unit: "kWh", Type: PointTypeInstant },
	// "p83089": { PsKey: "virtual", Id: "p83089", Name: "Es Discharge Energy", Unit: "kWh", Type: PointTypeInstant },
	// "p83095": { PsKey: "virtual", Id: "p83095", Name: "Es Total Discharge Energy", Unit: "kWh", Type: PointTypeTotal },
	// "p83118": { PsKey: "virtual", Id: "p83118", Name: "Use Energy", Unit: "kWh", Type: PointTypeInstant },
	// "p83120": { PsKey: "virtual", Id: "p83120", Name: "Es Energy", Unit: "kWh", Type: PointTypeInstant },
	// "p83127": { PsKey: "virtual", Id: "p83127", Name: "Es Total Energy", Unit: "kWh", Type: PointTypeTotal },

	// "co2_reduce": { PsKey: "virtual", Id: "co2_reduce", Name: "co2_reduce", Unit: "FOO", Type: PointTypeInstant },
	// "co2_reduce_total": { PsKey: "virtual", Id: "co2_reduce_total", Name: "co2_reduce_total", Unit: "FOO", Type: PointTypeTotal },
	// "curr_power": { PsKey: "virtual", Id: "curr_power", Name: "curr_power", Unit: "FOO", Type: PointTypeInstant },
	// "daily_irradiation": { PsKey: "virtual", Id: "daily_irradiation", Name: "daily_irradiation", Unit: "FOO", Type: PointTypeDaily },
	// "equivalent_hour": { PsKey: "virtual", Id: "equivalent_hour", Name: "equivalent_hour", Unit: "FOO", Type: PointTypeDaily },
	// "installed_power_map": { PsKey: "virtual", Id: "installed_power_map", Name: "installed_power_map", Unit: "FOO", Type: PointTypeInstant },
	// "radiation": { PsKey: "virtual", Id: "radiation", Name: "radiation", Unit: "FOO", Type: PointTypeInstant },
	// "today_energy": { PsKey: "virtual", Id: "today_energy", Name: "today_energy", Unit: "FOO", Type: PointTypeDaily },
	// "today_income": { PsKey: "virtual", Id: "today_income", Name: "today_income", Unit: "FOO", Type: PointTypeDaily },
	// "total_capacity": { PsKey: "virtual", Id: "total_capacity", Name: "total_capacity", Unit: "FOO", Type: PointTypeTotal },
	// "total_energy": { PsKey: "virtual", Id: "total_energy", Name: "total_energy", Unit: "FOO", Type: PointTypeTotal },
	// "total_income": { PsKey: "virtual", Id: "total_income", Name: "total_income", Unit: "FOO", Type: PointTypeTotal },
	// Added manually


	"p13001": { PsKey: "1129147_14_1_1", Id: "p13001", Name: "MPPT1 Voltage", Unit: "V", Type: PointTypeInstant },

	"p13012": { PsKey: "1129147_14_1_1", Id: "p13012", Name: "Total Reactive Power", Unit: "kvar", Type: PointTypeDaily },

	"p13105": { PsKey: "1129147_14_1_1", Id: "p13105", Name: "MPPT2 Voltage", Unit: "V", Type: PointTypeInstant },

	"p13122": { PsKey: "1129147_14_1_1", Id: "p13122", Name: "Daily Feed-in Energy", Unit: "kWh", Type: PointTypeDaily },

	"p13125": { PsKey: "1129147_14_1_1", Id: "p13125", Name: "Total Feed-in Energy", Unit: "kWh", Type: PointTypeTotal },

	"p13138": { PsKey: "1129147_14_1_1", Id: "p13138", Name: "Battery Voltage", Unit: "V", Type: PointTypeInstant },

	"p13144": { PsKey: "1129147_14_1_1", Id: "p13144", Name: "Daily Self-consumption Rate", Unit: "%", Type: PointTypeDaily },

	"p13157": { PsKey: "1129147_14_1_1", Id: "p13157", Name: "Phase A Voltage", Unit: "V", Type: PointTypeInstant },

	"p13158": { PsKey: "1129147_14_1_1", Id: "p13158", Name: "Phase B Voltage", Unit: "V", Type: PointTypeInstant },

	"p13159": { PsKey: "1129147_14_1_1", Id: "p13159", Name: "Phase C Voltage", Unit: "V", Type: PointTypeInstant },

	"p13161": { PsKey: "1129147_14_1_1", Id: "p13161", Name: "Bus Voltage", Unit: "V", Type: PointTypeInstant },
	"p13173": { PsKey: "1129147_14_1_1", Id: "p13173", Name: "Daily Feed-in Energy (PV)", Unit: "kWh", Type: PointTypeDaily },
	"p13175": { PsKey: "1129147_14_1_1", Id: "p13175", Name: "Total Feed-in Energy (PV)", Unit: "kWh", Type: PointTypeTotal },
	"p13002": { PsKey: "1129147_14_1_1", Id: "p13002", Name: "MPPT1 Current", Unit: "A", Type: PointTypeInstant },
	"p13003": { PsKey: "1129147_14_1_1", Id: "p13003", Name: "Total DC Power", Unit: "kW" },
	"p13007": { PsKey: "1129147_14_1_1", Id: "p13007", Name: "Grid Frequency", Unit: "Hz", Type: PointTypeInstant },
	"p13008": { PsKey: "1129147_14_1_1", Id: "p13008", Name: "Phase A Current", Unit: "A", Type: PointTypeInstant },
	"p13009": { PsKey: "1129147_14_1_1", Id: "p13009", Name: "Phase B Current", Unit: "A", Type: PointTypeInstant },
	"p13010": { PsKey: "1129147_14_1_1", Id: "p13010", Name: "Phase C Current", Unit: "A", Type: PointTypeInstant },
	"p13011": { PsKey: "1129147_14_1_1", Id: "p13011", Name: "Total Active Power", Unit: "kW" },
	"p13013": { PsKey: "1129147_14_1_1", Id: "p13013", Name: "Total Power Factor", Unit: "" },
	"p13018": { PsKey: "1129147_14_1_1", Id: "p13018", Name: "Total Apparent Power", Unit: "VA" },
	"p13019": { PsKey: "1129147_14_1_1", Id: "p13019", Name: "Internal Air Temperature", Unit: "℃", Type: PointTypeInstant },
	"p13028": { PsKey: "1129147_14_1_1", Id: "p13028", Name: "Daily Battery Charging Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13029": { PsKey: "1129147_14_1_1", Id: "p13029", Name: "Daily Battery Discharging Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13034": { PsKey: "1129147_14_1_1", Id: "p13034", Name: "Total Battery Charging Energy", Unit: "kWh" , Type: PointTypeTotal },
	"p13035": { PsKey: "1129147_14_1_1", Id: "p13035", Name: "Total Battery Discharging Energy", Unit: "kWh" , Type: PointTypeTotal },
	"p13106": { PsKey: "1129147_14_1_1", Id: "p13106", Name: "MPPT2 Current", Unit: "A", Type: PointTypeInstant },
	"p13112": { PsKey: "1129147_14_1_1", Id: "p13112", Name: "Daily PV Yield", Unit: "kWh", Type: PointTypeDaily },
	"p13116": { PsKey: "1129147_14_1_1", Id: "p13116", Name: "Daily Load Energy Consumption from PV", Unit: "kWh", Type: PointTypeDaily },
	"p13119": { PsKey: "1129147_14_1_1", Id: "p13119", Name: "Total Load Active Power", Unit: "kW" },
	"p13121": { PsKey: "1129147_14_1_1", Id: "p13121", Name: "Total Export Active  Power", Unit: "kW" },
	"p13126": { PsKey: "1129147_14_1_1", Id: "p13126", Name: "Battery Charging Power", Unit: "kW" },
	"p13130": { PsKey: "1129147_14_1_1", Id: "p13130", Name: "Total Load Energy Consumption", Unit: "kWh" , Type: PointTypeTotal },
	"p13134": { PsKey: "1129147_14_1_1", Id: "p13134", Name: "Total PV Yield", Unit: "kWh" , Type: PointTypeTotal },
	"p13137": { PsKey: "1129147_14_1_1", Id: "p13137", Name: "Total Load Energy Consumption from PV", Unit: "kWh" , Type: PointTypeTotal },
	"p13139": { PsKey: "1129147_14_1_1", Id: "p13139", Name: "Battery Current", Unit: "A", Type: PointTypeInstant },
	"p13140": { PsKey: "1129147_14_1_1", Id: "p13140", Name: "Battery Capacity(kWh)", Unit: "kWh" },
	"p13141": { PsKey: "1129147_14_1_1", Id: "p13141", Name: "Battery Level (SOC)", Unit: "%", Type: PointTypeInstant },
	"p13142": { PsKey: "1129147_14_1_1", Id: "p13142", Name: "Battery Health (SOH)", Unit: "%", Type: PointTypeInstant },
	"p13143": { PsKey: "1129147_14_1_1", Id: "p13143", Name: "Battery Temperature", Unit: "℃", Type: PointTypeInstant },
	"p13147": { PsKey: "1129147_14_1_1", Id: "p13147", Name: "Daily Purchased Energy", Unit: "kWh", Type: PointTypeDaily },
	"p13148": { PsKey: "1129147_14_1_1", Id: "p13148", Name: "Total Purchased Energy", Unit: "kWh", Type: PointTypeTotal },
	"p13149": { PsKey: "1129147_14_1_1", Id: "p13149", Name: "Purchased Power", Unit: "kW" },
	"p13150": { PsKey: "1129147_14_1_1", Id: "p13150", Name: "Battery Discharging Power", Unit: "kW" },
	"p13160": { PsKey: "1129147_14_1_1", Id: "p13160", Name: "Array Insulation Resistance", Unit: "kΩ", Type: PointTypeInstant },
	"p13162": { PsKey: "1129147_14_1_1", Id: "p13162", Name: "Max. Charging Current (BMS)", Unit: "A", Type: PointTypeInstant },
	"p13163": { PsKey: "1129147_14_1_1", Id: "p13163", Name: "Max. Discharging Current (BMS)", Unit: "A", Type: PointTypeInstant },
	"p13174": { PsKey: "1129147_14_1_1", Id: "p13174", Name: "Daily Battery Charging Energy from PV", Unit: "kWh", Type: PointTypeDaily },
	"p13176": { PsKey: "1129147_14_1_1", Id: "p13176", Name: "Total Battery Charging Energy from PV", Unit: "kWh", Type: PointTypeTotal },
	"p13199": { PsKey: "1129147_14_1_1", Id: "p13199", Name: "Daily Load Energy Consumption", Unit: "kWh", Type: PointTypeDaily },
	"p18062": { PsKey: "1129147_14_1_1", Id: "p18062", Name: "Phase A Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18063": { PsKey: "1129147_14_1_1", Id: "p18063", Name: "Phase B Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18064": { PsKey: "1129147_14_1_1", Id: "p18064", Name: "Phase C Backup Current", Unit: "A", Type: PointTypeInstant },
	"p18065": { PsKey: "1129147_14_1_1", Id: "p18065", Name: "Phase A Backup Power", Unit: "kW" },
	"p18066": { PsKey: "1129147_14_1_1", Id: "p18066", Name: "Phase B Backup Power", Unit: "kW" },
	"p18067": { PsKey: "1129147_14_1_1", Id: "p18067", Name: "Phase C Backup Power", Unit: "kW" },
	"p18068": { PsKey: "1129147_14_1_1", Id: "p18068", Name: "Total Backup Power", Unit: "kW" },
	"p83001": { PsKey: "1129147_11_0_0", Id: "p83001", Name: "Inverter AC Power Normalization", Unit: "kW/kWp" },
	"p83002": { PsKey: "1129147_11_0_0", Id: "p83002", Name: "Inverter AC Power", Unit: "kW" },
	"p83004": { PsKey: "1129147_11_0_0", Id: "p83004", Name: "Inverter Total Yield", Unit: "kWh" },
	"p83005": { PsKey: "1129147_11_0_0", Id: "p83005", Name: "Daily Equivalent Hours of Meter", Unit: "h", Type: PointTypeDaily },
	"p83006": { PsKey: "1129147_11_0_0", Id: "p83006", Name: "Meter Daily Yield", Unit: "kWh" },
	"p83007": { PsKey: "1129147_11_0_0", Id: "p83007", Name: "Meter PR", Unit: "%", Type: PointTypeInstant },
	"p83008": { PsKey: "1129147_11_0_0", Id: "p83008", Name: "Daily Equivalent Hours of Inverter", Unit: "h", Type: PointTypeDaily },
	"p83009": { PsKey: "1129147_11_0_0", Id: "p83009", Name: "Daily Yield by Inverter", Unit: "kWh", Type: PointTypeDaily },
	"p83010": { PsKey: "1129147_11_0_0", Id: "p83010", Name: "Inverter PR", Unit: "%", Type: PointTypeInstant },
	"p83013": { PsKey: "1129147_11_0_0", Id: "p83013", Name: "Daily Irradiation", Unit: "Wh/m2" },
	"p83016": { PsKey: "1129147_11_0_0", Id: "p83016", Name: "Plant Ambient Temperature", Unit: "℃", Type: PointTypeInstant },
	"p83017": { PsKey: "1129147_11_0_0", Id: "p83017", Name: "Plant Module Temperature", Unit: "℃", Type: PointTypeInstant },
	"p83018": { PsKey: "1129147_11_0_0", Id: "p83018", Name: "Daily Yield (Theoretical)", Unit: "kWh", Type: PointTypeDaily },
	"p83019": { PsKey: "1129147_11_0_0", Id: "p83019", Name: "Power/Installed Power of Plant", Unit: "%", Type: PointTypeInstant },
	"p83020": { PsKey: "1129147_11_0_0", Id: "p83020", Name: "Meter Total Yield", Unit: "kWh" },
	"p83021": { PsKey: "1129147_11_0_0", Id: "p83021", Name: "Accumulative Power Consumption by Meter", Unit: "kWh" },
	"p83022": { PsKey: "1129147_11_0_0", Id: "p83022", Name: "Daily Yield of Plant", Unit: "kWh", Type: PointTypeDaily },
	"p83023": { PsKey: "1129147_11_0_0", Id: "p83023", Name: "Plant PR", Unit: "%", Type: PointTypeInstant },
	"p83024": { PsKey: "1129147_11_0_0", Id: "p83024", Name: "Plant Total Yield", Unit: "kWh" },
	"p83025": { PsKey: "1129147_11_0_0", Id: "p83025", Name: "Plant Equivalent Hours", Unit: "h" },
	"p83032": { PsKey: "1129147_11_0_0", Id: "p83032", Name: "Meter AC Power", Unit: "kW" },
	"p83033": { PsKey: "1129147_11_0_0", Id: "p83033", Name: "Plant Power", Unit: "kW" },
	"p83097": { PsKey: "1129147_11_0_0", Id: "p83097", Name: "Daily Load Energy Consumption from PV", Unit: "kWh", Type: PointTypeDaily },
	"p83100": { PsKey: "1129147_11_0_0", Id: "p83100", Name: "Total Load Energy Consumption from PV", Unit: "kWh" },
	"p83102": { PsKey: "1129147_11_0_0", Id: "p83102", Name: "Daily Purchased Energy", Unit: "kWh", Type: PointTypeDaily },
	"p83105": { PsKey: "1129147_11_0_0", Id: "p83105", Name: "Total Purchased Energy", Unit: "kWh" },
	"p83106": { PsKey: "1129147_11_0_0", Id: "p83106", Name: "Load Power", Unit: "kW" },
	"p83124": { PsKey: "1129147_11_0_0", Id: "p83124", Name: "Total Load Energy Consumption", Unit: "MWh" },
	"p83128": { PsKey: "1129147_11_0_0", Id: "p83128", Name: "Total Active Power of Optical Storage", Unit: "kW" },
	"p83129": { PsKey: "1129147_11_0_0", Id: "p83129", Name: "Battery SOC", Unit: "%", Type: PointTypeInstant },
	"p83233": { PsKey: "1129147_11_0_0", Id: "p83233", Name: "Total field maximum rechargeable power", Unit: "MW" },
	"p83234": { PsKey: "1129147_11_0_0", Id: "p83234", Name: "Total field maximum dischargeable power", Unit: "MW" },
	"p83235": { PsKey: "1129147_11_0_0", Id: "p83235", Name: "Total field chargeable energy", Unit: "MWh" },
	"p83236": { PsKey: "1129147_11_0_0", Id: "p83236", Name: "Total field dischargeable energy", Unit: "MWh" },
	"p83237": { PsKey: "1129147_11_0_0", Id: "p83237", Name: "Total field energy storage maximum reactive power", Unit: "MW" },
	"p83238": { PsKey: "1129147_11_0_0", Id: "p83238", Name: "Total field energy storage active power", Unit: "MW" },
	"p83239": { PsKey: "1129147_11_0_0", Id: "p83239", Name: "Total field reactive power", Unit: "Mvar" },
	"p83241": { PsKey: "1129147_11_0_0", Id: "p83241", Name: "Total field charge capacity", Unit: "MWh" },
	"p83242": { PsKey: "1129147_11_0_0", Id: "p83242", Name: "Total field discharge capacity", Unit: "MWh" },
	"p83243": { PsKey: "1129147_11_0_0", Id: "p83243", Name: "Total field daily charge capacity", Unit: "MWh" },
	"p83244": { PsKey: "1129147_11_0_0", Id: "p83244", Name: "Total field daily discharge capacity", Unit: "MWh" },
	"p83252": { PsKey: "1129147_11_0_0", Id: "p83252", Name: "Battery Level (SOC)", Unit: "%", Type: PointTypeInstant },
	"p83419": { PsKey: "1129147_11_0_0", Id: "p83419", Name: "Daily Highest Inverter Power/Inverter Installed Capacity", Unit: "%" },
	"p83420": { PsKey: "1129147_11_0_0", Id: "p83420", Name: "Current Power/Inverter Installed Capacity", Unit: "%", Type: PointTypeInstant },
	"p83549": { PsKey: "1129147_11_0_0", Id: "p83549", Name: "Grid active power", Unit: "kW" },

	"p23014": { PsKey: "1129147_22_247_1", Id: "p23014", Name: "WLAN Signal Strength", Unit: "" },

}

func (pm PointsMap) Get(device string, point string) *Point {
	dp := device + "." + SetPoint(point)
	if p, ok := pm[dp]; ok {
		p.Valid = true
		p.FullId = NameDevicePoint(device, point)
		return &p
	}

	dp = SetPoint(point)
	if p, ok := pm[dp]; ok {
		p.Valid = true
		p.FullId = NameDevicePoint(device, point)
		return &p
	}

	return &Point {
		FullId: NameDevicePoint(device, point),
		PsKey: device,
		Id:    dp,
		Name:  "",
		Unit:  "",
		Type:  "",
		Valid: false,
	}
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

func SetPoint(point string) string {
	for range Only.Once {
		p := strings.TrimPrefix(point, "p")
		_, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			point = "p" + p
			break
		}
	}
	return point
}

func PointToName(s string) string {
	s = CleanString(s)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.Title(s)
	return s
}
