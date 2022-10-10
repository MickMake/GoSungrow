package getPsDetailWithPsType

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"time"
)

const Url = "/v1/powerStationService/getPsDetailWithPsType"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BatteryLevelPercent         api.Integer   `json:"battery_level_percent" PointId:"BatteryLevelPercent" PointType:"PointTypeBoot"`
	ChargingDischargingPowerMap api.UnitValue `json:"charging_discharging_power_map" PointId:"ChargingDischargingPowerMap" PointType:"PointTypeInstant"`
	Co2ReduceTotal              api.UnitValue `json:"co2_reduce_total" PointId:"Co2ReduceTotal" PointType:"PointTypeTotal"`
	CoalReduceTotal             api.UnitValue `json:"coal_reduce_total" PointId:"CoalReduceTotal" PointType:"PointTypeTotal"`
	ConnectType                 string        `json:"connect_type" PointId:"ConnectType" PointType:"PointTypeBoot"`
	CurrPower                   api.UnitValue `json:"curr_power" PointId:"CurrPower" PointType:"PointTypeInstant"`
	DesignCapacity              api.UnitValue `json:"design_capacity" PointId:"DesignCapacity" PointType:"PointTypeBoot"`
	EnergyScheme                interface{}   `json:"energy_scheme"`
	GcjLatitude                 api.Float     `json:"gcj_latitude" PointId:"GcjLatitude" PointType:"PointTypeBoot"`
	GcjLongitude                api.Float     `json:"gcj_longitude" PointId:"GcjLongitude" PointType:"PointTypeBoot"`
	HasAmmeter                  api.Bool      `json:"has_ammeter" PointId:"HasAmmeter" PointType:"PointTypeBoot"`
	HouseholdInverterData       interface{}   `json:"household_inverter_data"`
	InstallerPsFaultStatus      string        `json:"installer_ps_fault_status" PointId:"InstallerPsFaultStatus" PointType:"PointTypeBoot"`
	IsHaveEsInverter            api.Bool      `json:"is_have_es_inverter" PointId:"IsHaveEsInverter" PointType:"PointTypeBoot"`
	IsSingleInverter            api.Bool      `json:"is_single_inverter" PointId:"IsSingleInverter" PointType:"PointTypeBoot"`
	IsTransformSystem           api.Bool      `json:"is_transform_system" PointId:"IsTransformSystem" PointType:"PointTypeBoot"`
	Latitude                    api.Float     `json:"latitude" PointId:"Latitude" PointType:"PointTypeBoot"`
	LoadPowerMap                api.UnitValue `json:"load_power_map" PointId:"LoadPowerMap" PointType:"PointTypeInstant"`
	LoadPowerMapVirgin          api.UnitValue `json:"load_power_map_virgin"  PointIgnore:"true"`
	Longitude                   api.Float     `json:"longitude" PointId:"Longitude" PointType:"PointTypeBoot"`
	MapLatitude                 api.Float     `json:"map_latitude" PointId:"MapLatitude" PointType:"PointTypeBoot"`
	MapLongitude                api.Float     `json:"map_longitude" PointId:"MapLongitude" PointType:"PointTypeBoot"`
	MeterReduceTotal            api.UnitValue `json:"meter_reduce_total" PointId:"MeterReduceTotal" PointType:"PointTypeTotal"`
	MobileTel                   api.String    `json:"moble_tel" PointId:"MobleTel" PointType:"PointTypeBoot"`
	MonthEnergy                 api.UnitValue `json:"month_energy" PointId:"MonthEnergy" PointType:"PointTypeMonthly"`
	MonthEnergyVirgin           api.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthIncome                 api.UnitValue `json:"month_income" PointId:"MonthIncome" PointType:"PointTypeMonthly"`
	NegativeLoadMsg             interface{}   `json:"negative_load_msg"`
	OwnerPsFaultStatus          string        `json:"owner_ps_fault_status" PointId:"OwnerPsFaultStatus" PointType:"PointTypeBoot"`
	P83081Map                   api.UnitValue `json:"p83081_map" PointId:"P83081Map" PointType:"PointTypeInstant"`
	P83081MapVirgin             api.UnitValue `json:"p83081_map_virgin"  PointIgnore:"true"`
	P83102Map                   api.UnitValue `json:"p83102_map" PointId:"P83102Map" PointType:"PointTypeInstant"`
	P83102MapVirgin             api.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent               api.Float     `json:"p83102_percent" PointId:"P83102Percent" PointType:"PointTypeInstant"`
	P83118Map                   api.UnitValue `json:"p83118_map" PointId:"P83118Map" PointType:"PointTypeInstant"`
	P83118MapVirgin             api.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map                   api.UnitValue `json:"p83119_map" PointId:"P83119Map" PointType:"PointTypeInstant"`
	P83119MapVirgin             api.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map                   api.UnitValue `json:"p83120_map" PointId:"P83120Map" PointType:"PointTypeInstant"`
	P83120MapVirgin             api.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83122                      api.Float     `json:"p83122" PointId:"P83122" PointType:"PointTypeInstant"`
	P83124Map                   api.UnitValue `json:"p83124_map" PointId:"P83124Map" PointType:"PointTypeInstant"`
	P83124MapVirgin             api.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83202Map                   api.UnitValue `json:"p83202_map" PointId:"P83202Map" PointType:"PointTypeInstant"`
	P83202MapVirgin             api.UnitValue `json:"p83202_map_virgin"  PointIgnore:"true"`
	P83532MapVirgin             api.UnitValue `json:"p83532_map_virgin"  PointIgnore:"true"`
	PowerChargeSetted           api.Integer   `json:"power_charge_setted" PointId:"PowerChargeSetted" PointType:"PointTypeBoot"`
	PowerGridPowerMap           api.UnitValue `json:"power_grid_power_map" PointId:"PowerGridPowerMap" PointType:"PointTypeInstant"`
	PowerGridPowerMapVirgin     api.UnitValue `json:"power_grid_power_map_virgin"  PointIgnore:"true"`
	PsCountryID                 api.Integer   `json:"ps_country_id" PointId:"PsCountryID" PointType:"PointTypeBoot"`
	PsDeviceType                api.Integer   `json:"ps_device_type" PointId:"PsDeviceType" PointType:"PointTypeBoot"`
	PsFaultStatus               string        `json:"ps_fault_status" PointId:"PsFaultStatus" PointType:"PointTypeBoot"`
	PsHealthStatus              string        `json:"ps_health_status" PointId:"PsHealthStatus" PointType:"PointTypeBoot"`
	PsLocation                  api.String    `json:"ps_location" PointId:"PsLocation" PointType:"PointTypeBoot"`
	PsName                      api.String    `json:"ps_name" PointId:"PsName" PointType:"PointTypeBoot"`
	PsPsKey                     api.PsKey     `json:"ps_ps_key" PointId:"PsPsKey" PointType:"PointTypeBoot"`
	PsState                     string        `json:"ps_state" PointId:"PsState" PointType:"PointTypeBoot"`
	PsType                      api.Integer   `json:"ps_type" PointId:"PsType" PointType:"PointTypeBoot"`
	PvPowerMap                  api.UnitValue `json:"pv_power_map" PointId:"PvPowerMap" PointType:"PointTypeInstant"`
	PvPowerMapVirgin            api.UnitValue `json:"pv_power_map_virgin"  PointIgnore:"true"`
	RobotNumSweepCapacity       struct {
		Num           api.Integer `json:"num" PointId:"Num" PointType:"PointTypeBoot"`
		SweepCapacity api.Float   `json:"sweep_capacity" PointId:"SweepCapacity" PointType:"PointTypeBoot"`
	} `json:"robot_num_sweep_capacity"`
	SelfConsumptionOffsetReminder api.Integer   `json:"self_consumption_offset_reminder" PointId:"SelfConsumptionOffsetReminder" PointType:"PointTypeBoot"`
	So2ReduceTotal                api.UnitValue `json:"so2_reduce_total" PointId:"So2ReduceTotal" PointType:"PointTypeTotal"`
	StorageInverterData           []struct {
		CommunicationDevSn      api.String    `json:"communication_dev_sn" PointId:"CommunicationDevSn" PointType:"PointTypeBoot"`
		DevStatus               api.Integer   `json:"dev_status" PointId:"DevStatus" PointType:"PointTypeBoot"`
		DeviceCode              api.Integer   `json:"device_code" PointId:"DeviceCode" PointType:"PointTypeBoot"`
		DeviceModelCode         api.String    `json:"device_model_code" PointId:"DeviceModelCode" PointType:"PointTypeBoot"`
		DeviceName              api.String    `json:"device_name" PointId:"DeviceName" PointType:"PointTypeBoot"`
		DeviceState             string        `json:"device_state" PointId:"DeviceState" PointType:"PointTypeBoot"`
		DeviceType              api.Integer   `json:"device_type" PointId:"DeviceType" PointType:"PointTypeBoot"`
		DrmStatus               api.Integer   `json:"drm_status" PointId:"DrmStatus" PointType:"PointTypeBoot"`
		DrmStatusName           api.String    `json:"drm_status_name" PointId:"DrmStatusName" PointType:"PointTypeBoot"`
		EnergyFlow              []string      `json:"energy_flow"`
		HasAmmeter              api.Integer   `json:"has_ammeter" PointId:"HasAmmeter" PointType:"PointTypeBoot"`
		InstallerDevFaultStatus api.Integer   `json:"installer_dev_fault_status" PointId:"InstallerDevFaultStatus" PointType:"PointTypeBoot"`
		InverterSn              api.String    `json:"inverter_sn" PointId:"InverterSn" PointType:"PointTypeBoot"`
		OwnerDevFaultStatus     api.Integer   `json:"owner_dev_fault_status" PointId:"OwnerDevFaultStatus" PointType:"PointTypeBoot"`
		P13003Map               api.UnitValue `json:"p13003_map" PointId:"P13003Map" PointType:"PointTypeInstant"`
		P13003MapVirgin         api.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13119Map               api.UnitValue `json:"p13119_map" PointId:"P13119Map" PointType:"PointTypeInstant"`
		P13119MapVirgin         api.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map               api.UnitValue `json:"p13121_map" PointId:"P13121Map" PointType:"PointTypeInstant"`
		P13121MapVirgin         api.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map               api.UnitValue `json:"p13126_map" PointId:"P13126Map" PointType:"PointTypeInstant"`
		P13126MapVirgin         api.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141                  api.Float     `json:"p13141" PointId:"P13141" PointType:"PointTypeInstant"`
		P13149Map               api.UnitValue `json:"p13149_map" PointId:"P13149Map" PointType:"PointTypeInstant"`
		P13149MapVirgin         api.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map               api.UnitValue `json:"p13150_map" PointId:"P13150Map" PointType:"PointTypeInstant"`
		P13150MapVirgin         api.UnitValue `json:"p13150_map_virgin"  PointIgnore:"true"`
		PsKey                   api.PsKey     `json:"ps_key" PointId:"PsKey" PointType:"PointTypeBoot"`
		UUID                    api.Integer   `json:"uuid" PointId:"UUID" PointType:"PointTypeBoot"`
	} `json:"storage_inverter_data"`
	TodayEnergy       api.UnitValue `json:"today_energy" PointId:"TodayEnergy" PointType:"PointTypeDaily"`
	TodayEnergyVirgin api.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome       api.UnitValue `json:"today_income" PointId:"TodayIncome" PointType:"PointTypeDaily"`
	TotalEnergy       api.UnitValue `json:"total_energy" PointId:"TotalEnergy" PointType:"PointTypeTotal"`
	TotalEnergyVirgin api.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalIncome       api.UnitValue `json:"total_income" PointId:"TotalIncome" PointType:"PointTypeTotal"`
	TreeReduceTotal   api.UnitValue `json:"tree_reduce_total" PointId:"TreeReduceTotal" PointType:"PointTypeTotal"`
	ValidFlag         api.Integer   `json:"valid_flag" PointId:"ValidFlag" PointType:"PointTypeBoot"`
	WgsLatitude       api.Float     `json:"wgs_latitude" PointId:"WgsLatitude" PointType:"PointTypeBoot"`
	WgsLongitude      api.Float     `json:"wgs_longitude" PointId:"WgsLongitude" PointType:"PointTypeBoot"`
	ZfzyMap           api.UnitValue `json:"zfzy_map" PointId:"ZfzyMap" PointType:"PointTypeInstant"`
	ZfzyMapVirgin     api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap           api.UnitValue `json:"zjzz_map" PointId:"ZjzzMap" PointType:"PointTypeInstant"`
	ZjzzMapVirgin     api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Date",
			"Point Id",
			"Description",
			"Value",
			"(Unit)",
			"(Unit)",
		)

		now := api.TimeNowString()

		keys := api.GetStructKeys(e.Response.ResultData)
		for _, n := range keys.Sort() {
			pn := api.PointId(n)
			p := api.GetPoint(e.Response.ResultData.PsPsKey.Value(), pn)
			if p.Valid {
				_ = table.AddRow(
					now,
					api.NameDevicePoint(e.Response.ResultData.PsPsKey.Value(), pn),
					p.Name,
					keys[pn].Value,
					p.Unit,
					keys[pn].Unit,
				)
				continue
			}

			_ = table.AddRow(
				now,
				api.NameDevicePoint(e.Response.ResultData.PsPsKey.Value(), pn),
				api.PointToName(pn),
				keys[pn].Value,
				keys[pn].Unit,
				keys[pn].Unit,
			)
		}

		if len(e.Response.ResultData.StorageInverterData) == 0 {
			break
		}

		for _, sid := range e.Response.ResultData.StorageInverterData {
			keys = api.GetStructKeys(sid)
			for _, n := range keys.Sort() {
				pn := api.PointId(n)
				p := api.GetPoint(sid.PsKey.Value(), pn)
				if p.Valid {
					_ = table.AddRow(
						now,
						api.NameDevicePoint(sid.PsKey.Value(), pn),
						p.Name,
						keys[pn].Value,
						p.Unit,
						keys[pn].Unit,
					)
					continue
				}

				_ = table.AddRow(
					now,
					api.NameDevicePoint(sid.PsKey.Value(), pn),
					api.PointToName(pn),
					keys[pn].Value,
					keys[pn].Unit,
					keys[pn].Unit,
				)
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(4),
		// 	ValueColumn:  output.SetInteger(5),
		// 	UnitsColumn:  output.SetInteger(6),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}

	return table
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		name := fmt.Sprintf("getPsDetailWithPsType.%s", e.Response.ResultData.PsPsKey.Value())
		entries.StructToPoints(e.Response.ResultData, name, e.Response.ResultData.PsPsKey.Value(), time.Time{})

		for _, sid := range e.Response.ResultData.StorageInverterData {
			entries.StructToPoints(sid, name + ".StorageInverterData." + sid.PsKey.Value(), sid.PsKey.Value(), time.Time{})
		}
	}
	return entries
}

func (e *EndPoint) GetPsKeys() []string {
	ret := []string{e.Response.ResultData.PsPsKey.Value()}
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = append(ret, l.PsKey.Value())
	}
	return ret
}

func (e *EndPoint) GetPsName() string {
	return e.Response.ResultData.PsName.Value()
}

func (e *EndPoint) GetPsState() string {
	return e.Response.ResultData.PsState
}

func (e *EndPoint) GetPsKey() string {
	return e.Response.ResultData.PsPsKey.Value()
}

func (e *EndPoint) GetDeviceModelCode() string {
	ret := e.Response.ResultData.PsPsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.DeviceModelCode.Value()
		break
	}
	return ret
}

func (e *EndPoint) GetDeviceName() string {
	ret := e.Response.ResultData.PsPsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.DeviceName.Value()
		break
	}
	return ret
}

func (e *EndPoint) GetDeviceSerial() string {
	ret := e.Response.ResultData.PsPsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.InverterSn.Value()
		break
	}
	return ret
}

// ChargingDischargingPowerMap
// Co2ReduceTotal
// CoalReduceTotal
// ConnectType
// CurrPower
// DesignCapacity
// EnergyScheme
// GcjLatitude
// GcjLongitude
// HasAmmeter
// HouseholdInverterData
// InstallerPsFaultStatus
// IsHaveEsInverter
// IsSingleInverter
// IsTransformSystem
// Latitude
// LoadPowerMap
// LoadPowerMapVirgin
// Longitude
// MapLatitude
// MapLongitude
// MeterReduceTotal
// MobleTel
// MonthEnergy
// MonthEnergyVirgin
// MonthIncome
// NegativeLoadMsg
// OwnerPsFaultStatus
// P83081Map
// P83081MapVirgin
// P83102Map
// P83102MapVirgin
// P83102Percent
// P83118Map
// P83118MapVirgin
// P83119Map
// P83119MapVirgin
// P83120Map
// P83120MapVirgin
// P83122
// P83124Map
// P83124MapVirgin
// P83202Map
// P83202MapVirgin
// P83532MapVirgin
// PowerChargeSetted
// PowerGridPowerMap
// PowerGridPowerMapVirgin
// PsCountryID
// PsDeviceType
// PsFaultStatus
// PsHealthStatus
// PsLocation
// PsName
// PsPsKey
// PsState
// PsType
// PvPowerMap
// PvPowerMapVirgin
// RobotNumSweepCapacity
// Num
// SweepCapacity
// }
// SelfConsumptionOffsetReminder
// So2ReduceTotal
// StorageInverterData
// CommunicationDevSn
// DevStatus
// DeviceCode
// DeviceModelCode
// DeviceName
// DeviceState
// DeviceType
// DrmStatus
// DrmStatusName
// EnergyFlow
// HasAmmeter
// InstallerDevFaultStatus
// InverterSn
// OwnerDevFaultStatus
// P13003Map
// P13003MapVirgin
// P13119Map
// P13119MapVirgin
// P13121Map
// P13121MapVirgin
// P13126Map
// P13126MapVirgin
// P13141
// P13149Map
// P13149MapVirgin
// P13150Map
// P13150MapVirgin
// PsKey
// UUID
// }
// TodayEnergy
// TodayEnergyVirgin
// TodayIncome
// TotalEnergy
// TotalEnergyVirgin
// TotalIncome
// TreeReduceTotal
// ValidFlag
// WgsLatitude
// WgsLongitude
// ZfzyMap
// ZfzyMapVirgin
// ZjzzMap
// ZjzzMapVirgin
