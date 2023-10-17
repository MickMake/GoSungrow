package api

import (
	"fmt"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	VirtualPsId = "virtual"
)

func init() {
	for mi := range Points.Map {
		for pi := range Points.Map[mi].Parents.Map {
			device := Points.Map[mi].Parents.Map[pi]
			device.Split()
			Points.Map[mi].Parents.Map[pi] = device
			// fmt.Sprintf("")
		}
	}
}

type PointsMap struct {
	Map map[string]*Point
}

func (pm *PointsMap) Resolve(point string) *Point {
	if ret, ok := pm.Map[point]; ok {
		return ret
	}
	return nil
}

// func ResolvePoint(point string) *Point {
// 	return Points.Resolve(point)
// }

func (pm *PointsMap) Get(point string) *Point {
	var ok bool
	var p *Point
	if p, ok = pm.Map[point]; ok {
		p.Valid = true
		return p
	}

	return p
	// // parentDevice := ParentDevice{Key: device}
	// return &Point {
	// 	// FullId: NameDevicePoint(device, point),
	// 	// Parent: parentDevice.Split(),
	// 	Id:       valueTypes.SetPointIdString(point),
	// 	Name:     "",
	// 	Unit:     "",
	// 	UpdateFreq: "",
	// 	Valid:    false,
	// }
}

func (pm *PointsMap) GetDevicePoint(devicePoint string) *Point {
	ret := &Point{}
	for range only.Once {
		s := strings.Split(devicePoint, ".")
		if len(s) < 2 {
			ret = pm.Get(s[0])
			break
		}
		ret = pm.Get(s[0] + "." + s[1])
	}
	return ret
}

func (pm *PointsMap) Exists(pid string) bool {
	var ok bool
	for range only.Once {
		if _, ok = Points.Map[pid]; !ok {
			break
		}
		ok = true
	}
	return ok
}

func (pm *PointsMap) Print() {
	for range only.Once {
		for k, v := range pm.Map {
			fmt.Printf("[%s] => %s\n", k, v)
		}
		fmt.Println("")
	}
}

// func (pm *PointsMap) Add(pid valueTypes.PointId, point Point) bool {

func (pm *PointsMap) Add(point Point) bool {
	var ok bool
	for range only.Once {
		point.Parents = ParentDevices{}

		if !pm.Exists(point.Id) {
			pm.Map[point.Id] = &point
			break
		}

		var ep Point
		ep = *pm.Map[point.Id]
		ep.Parents = ParentDevices{}

		if strings.ToLower(ep.Description) == ep.Id {
			ep.Description = point.Description
		}
		if ep.Description == "" {
			ep.Description = point.Description
		}

		if ep.Unit == "" {
			ep.Unit = point.Unit
		}
		if ep.UpdateFreq == "" {
			ep.UpdateFreq = point.UpdateFreq
		}
		if ep.GroupName == "" {
			ep.GroupName = point.GroupName
		}
		ep.Parents = ParentDevices{}
		ep.Valid = true
		ep.FixUnitType()

		var ep1 Point
		ep1 = *pm.Map[point.Id]
		ep1.Parents = ParentDevices{}
		if ep.HasNotChanged(ep1) {
			break
		}

		// jep, _ := json.Marshal(ep1)
		// jep2, _ := json.Marshal(ep)
		// fmt.Printf("EXISTS[%s]:\n", pid)
		// fmt.Printf("CURRENT: '%s'\n", jep)
		// fmt.Printf("NOW:     '%s'\n", jep2)

		pm.Map[ep.Id] = &ep
	}
	return ok
}

func (pd *ParentDevice) HasChanged(comp ParentDevice) bool {
	ok := true
	for range only.Once {
		// if pd.Key != comp.Key {
		// 	break
		// }

		if pd.Type != comp.Type {
			break
		}

		if pd.Code != comp.Code {
			break
		}

		ok = false
	}
	return ok
}

func (p *Point) HasChanged(comp Point) bool {
	ok := true
	for range only.Once {
		// Can also use Marshal method
		// jep, _ := json.Marshal(pm.Map[point.Id])
		// jep2, _ := json.Marshal(ep)
		// fmt.Printf("EXISTS[%s]:\n", pid)
		// fmt.Printf("CURRENT: '%s'\n", jep)
		// // fmt.Printf("CHANGE:  '%s'\n", jp)
		// fmt.Printf("NOW:     '%s'\n", jep2)

		if p.GroupName != comp.GroupName {
			break
		}

		if p.Description != comp.Description {
			break
		}

		if p.Unit != comp.Unit {
			break
		}

		if p.UpdateFreq != comp.UpdateFreq {
			break
		}

		ok = false
	}
	return ok
}

func (p *Point) HasNotChanged(comp Point) bool {
	return !p.HasChanged(comp)
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
var Points = PointsMap{
	Map: map[string]*Point{

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

		// "foo": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "", Name: "", Unit: "", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13001": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13001", Name: "MPPT1 Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13012": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13012", Name: "Total Reactive Power", Unit: "kvar", UpdateFreq: PointUpdateFreqDay},
		//
		// "p13105": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13105", Name: "MPPT2 Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13122": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13122", Name: "Daily Feed-in Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		//
		// "p13125": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13125", Name: "Total Feed-in Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		//
		// "p13138": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13138", Name: "Battery Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13144": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13144", Name: "Daily Self-consumption Rate", Unit: "%", UpdateFreq: PointUpdateFreqDay},
		//
		// "p13157": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13157", Name: "Phase A Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13158": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13158", Name: "Phase B Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13159": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13159", Name: "Phase C Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		//
		// "p13161": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13161", Name: "Bus Voltage", Unit: "V", UpdateFreq: PointUpdateFreqInstant},
		// "p13173": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13173", Name: "Daily Feed-in Energy (PV)", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13175": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13175", Name: "Total Feed-in Energy (PV)", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13002": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13002", Name: "MPPT1 Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13003": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13003", Name: "Total DC Power", Unit: "kW"},
		// "p13007": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13007", Name: "Grid Frequency", Unit: "Hz", UpdateFreq: PointUpdateFreqInstant},
		// "p13008": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13008", Name: "Phase A Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13009": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13009", Name: "Phase B Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13010": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13010", Name: "Phase C Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13011": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13011", Name: "Total Active Power", Unit: "kW"},
		// "p13013": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13013", Name: "Total Power Factor", Unit: ""},
		// "p13018": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13018", Name: "Total Apparent Power", Unit: "VA"},
		// "p13019": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13019", Name: "Internal Air Temperature", Unit: "℃", UpdateFreq: PointUpdateFreqInstant},
		// "p13028": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13028", Name: "Daily Battery Charging Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13029": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13029", Name: "Daily Battery Discharging Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13034": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13034", Name: "Total Battery Charging Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13035": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13035", Name: "Total Battery Discharging Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13106": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13106", Name: "MPPT2 Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13112": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13112", Name: "Daily PV Yield", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13116": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13116", Name: "Daily Load Energy Consumption from PV", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13119": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13119", Name: "Total Load Active Power", Unit: "kW"},
		// "p13121": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13121", Name: "Total Export Active  Power", Unit: "kW"},
		// "p13126": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13126", Name: "Battery Charging Power", Unit: "kW"},
		// "p13130": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13130", Name: "Total Load Energy Consumption", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13134": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13134", Name: "Total PV Yield", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13137": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13137", Name: "Total Load Energy Consumption from PV", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13139": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13139", Name: "Battery Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13140": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13140", Name: "Battery Capacity(kWh)", Unit: "kWh"},
		// "p13141": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13141", Name: "Battery Level (SOC)", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p13142": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13142", Name: "Battery Health (SOH)", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p13143": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13143", Name: "Battery Temperature", Unit: "℃", UpdateFreq: PointUpdateFreqInstant},
		// "p13147": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13147", Name: "Daily Purchased Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13148": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13148", Name: "Total Purchased Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13149": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13149", Name: "Purchased Power", Unit: "kW"},
		// "p13150": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13150", Name: "Battery Discharging Power", Unit: "kW"},
		// "p13160": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13160", Name: "Array Insulation Resistance", Unit: "kΩ", UpdateFreq: PointUpdateFreqInstant},
		// "p13162": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13162", Name: "Max. Charging Current (BMS)", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13163": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13163", Name: "Max. Discharging Current (BMS)", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p13174": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13174", Name: "Daily Battery Charging Energy from PV", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p13176": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13176", Name: "Total Battery Charging Energy from PV", Unit: "kWh", UpdateFreq: PointUpdateFreqTotal},
		// "p13199": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p13199", Name: "Daily Load Energy Consumption", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p18062": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18062", Name: "Phase A Backup Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p18063": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18063", Name: "Phase B Backup Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p18064": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18064", Name: "Phase C Backup Current", Unit: "A", UpdateFreq: PointUpdateFreqInstant},
		// "p18065": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18065", Name: "Phase A Backup Power", Unit: "kW"},
		// "p18066": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18066", Name: "Phase B Backup Power", Unit: "kW"},
		// "p18067": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18067", Name: "Phase C Backup Power", Unit: "kW"},
		// "p18068": {Parents: ParentDevices{Map: map[string]*ParentDevice{"14_1":{PsId: "manual", Type: "14", Code: "1"}}}, Id: "p18068", Name: "Total Backup Power", Unit: "kW"},
		// "p83001": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83001", Name: "Inverter AC Power Normalization", Unit: "kW/kWp"},
		// "p83002": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83002", Name: "Inverter AC Power", Unit: "kW"},
		// "p83004": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83004", Name: "Inverter Total Yield", Unit: "kWh"},
		// "p83005": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83005", Name: "Daily Equivalent Hours of Meter", Unit: "h", UpdateFreq: PointUpdateFreqDay},
		// "p83006": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83006", Name: "Meter Daily Yield", Unit: "kWh"},
		// "p83007": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83007", Name: "Meter PR", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83008": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83008", Name: "Daily Equivalent Hours of Inverter", Unit: "h", UpdateFreq: PointUpdateFreqDay},
		// "p83009": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83009", Name: "Daily Yield by Inverter", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p83010": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83010", Name: "Inverter PR", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83013": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83013", Name: "Daily Irradiation", Unit: "Wh/m2"},
		// "p83016": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83016", Name: "Plant Ambient Temperature", Unit: "℃", UpdateFreq: PointUpdateFreqInstant},
		// "p83017": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83017", Name: "Plant Module Temperature", Unit: "℃", UpdateFreq: PointUpdateFreqInstant},
		// "p83018": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83018", Name: "Daily Yield (Theoretical)", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p83019": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83019", Name: "Power/Installed Power of Plant", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83020": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83020", Name: "Meter Total Yield", Unit: "kWh"},
		// "p83021": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83021", Name: "Accumulative Power Consumption by Meter", Unit: "kWh"},
		// "p83022": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83022", Name: "Daily Yield of Plant", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p83023": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83023", Name: "Plant PR", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83024": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83024", Name: "Plant Total Yield", Unit: "kWh"},
		// "p83025": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83025", Name: "Plant Equivalent Hours", Unit: "h"},
		// "p83032": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83032", Name: "Meter AC Power", Unit: "kW"},
		// "p83033": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83033", Name: "Plant Power", Unit: "kW"},
		// "p83097": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83097", Name: "Daily Load Energy Consumption from PV", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p83100": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83100", Name: "Total Load Energy Consumption from PV", Unit: "kWh"},
		// "p83102": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83102", Name: "Daily Purchased Energy", Unit: "kWh", UpdateFreq: PointUpdateFreqDay},
		// "p83105": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83105", Name: "Total Purchased Energy", Unit: "kWh"},
		// "p83106": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83106", Name: "Load Power", Unit: "kW"},
		// "p83124": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83124", Name: "Total Load Energy Consumption", Unit: "MWh"},
		// "p83128": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83128", Name: "Total Active Power of Optical Storage", Unit: "kW"},
		// "p83129": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83129", Name: "Battery SOC", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83233": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83233", Name: "Total field maximum rechargeable power", Unit: "MW"},
		// "p83234": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83234", Name: "Total field maximum dischargeable power", Unit: "MW"},
		// "p83235": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83235", Name: "Total field chargeable energy", Unit: "MWh"},
		// "p83236": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83236", Name: "Total field dischargeable energy", Unit: "MWh"},
		// "p83237": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83237", Name: "Total field energy storage maximum reactive power", Unit: "MW"},
		// "p83238": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83238", Name: "Total field energy storage active power", Unit: "MW"},
		// "p83239": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83239", Name: "Total field reactive power", Unit: "Mvar"},
		// "p83241": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83241", Name: "Total field charge capacity", Unit: "MWh"},
		// "p83242": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83242", Name: "Total field discharge capacity", Unit: "MWh"},
		// "p83243": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83243", Name: "Total field daily charge capacity", Unit: "MWh"},
		// "p83244": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83244", Name: "Total field daily discharge capacity", Unit: "MWh"},
		// "p83252": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83252", Name: "Battery Level (SOC)", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83419": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83419", Name: "Daily Highest Inverter Power/Inverter Installed Capacity", Unit: "%"},
		// "p83420": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83420", Name: "Current Power/Inverter Installed Capacity", Unit: "%", UpdateFreq: PointUpdateFreqInstant},
		// "p83549": {Parents: ParentDevices{Map: map[string]*ParentDevice{"11_0":{PsId: "manual", Type: "11", Code: "0"}}}, Id: "p83549", Name: "Grid active power", Unit: "kW"},
		//
		// "p23014": {Parents: ParentDevices{Map: map[string]*ParentDevice{"22_247":{PsId: "manual", Type: "22", Code: "247"}}}, Id: "p23014", Name: "WLAN Signal Strength", Unit: ""},
	},
}
