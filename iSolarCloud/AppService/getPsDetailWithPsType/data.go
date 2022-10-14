package getPsDetailWithPsType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/powerStationService/getPsDetailWithPsType"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	BatteryLevelPercent         valueTypes.Integer   `json:"battery_level_percent" PointId:"BatteryLevelPercent" PointUnit:"%" PointTimeSpan:"PointTimeSpanInstant"`
	ChargingDischargingPowerMap valueTypes.UnitValue `json:"charging_discharging_power_map" PointId:"ChargingDischargingPowerMap" PointTimeSpan:"PointTimeSpanInstant"`
	Co2ReduceTotal              valueTypes.UnitValue `json:"co2_reduce_total" PointId:"Co2ReduceTotal" PointTimeSpan:"PointTimeSpanTotal"`
	CoalReduceTotal             valueTypes.UnitValue `json:"coal_reduce_total" PointId:"CoalReduceTotal" PointTimeSpan:"PointTimeSpanTotal"`
	ConnectType                 string        `json:"connect_type" PointId:"ConnectType" PointTimeSpan:"PointTimeSpanBoot"`
	CurrPower                   valueTypes.UnitValue `json:"curr_power" PointId:"CurrPower" PointTimeSpan:"PointTimeSpanInstant"`
	DesignCapacity              valueTypes.UnitValue `json:"design_capacity" PointId:"DesignCapacity" PointTimeSpan:"PointTimeSpanBoot"`
	EnergyScheme                interface{}   `json:"energy_scheme"`
	GcjLatitude                 valueTypes.Float     `json:"gcj_latitude" PointId:"GcjLatitude" PointTimeSpan:"PointTimeSpanBoot"`
	GcjLongitude                valueTypes.Float     `json:"gcj_longitude" PointId:"GcjLongitude" PointTimeSpan:"PointTimeSpanBoot"`
	HasAmmeter                  valueTypes.Bool      `json:"has_ammeter" PointId:"HasAmmeter" PointTimeSpan:"PointTimeSpanBoot"`
	HouseholdInverterData       interface{}   `json:"household_inverter_data"`
	InstallerPsFaultStatus      string        `json:"installer_ps_fault_status" PointId:"InstallerPsFaultStatus" PointTimeSpan:"PointTimeSpanBoot"`
	IsHaveEsInverter            valueTypes.Bool      `json:"is_have_es_inverter" PointId:"IsHaveEsInverter" PointTimeSpan:"PointTimeSpanBoot"`
	IsSingleInverter            valueTypes.Bool      `json:"is_single_inverter" PointId:"IsSingleInverter" PointTimeSpan:"PointTimeSpanBoot"`
	IsTransformSystem           valueTypes.Bool      `json:"is_transform_system" PointId:"IsTransformSystem" PointTimeSpan:"PointTimeSpanBoot"`
	Latitude                    valueTypes.Float     `json:"latitude" PointId:"Latitude" PointTimeSpan:"PointTimeSpanBoot"`
	LoadPowerMap                valueTypes.UnitValue `json:"load_power_map" PointId:"LoadPowerMap" PointTimeSpan:"PointTimeSpanInstant"`
	LoadPowerMapVirgin          valueTypes.UnitValue `json:"load_power_map_virgin"  PointIgnore:"true"`
	Longitude                   valueTypes.Float     `json:"longitude" PointId:"Longitude" PointTimeSpan:"PointTimeSpanBoot"`
	MapLatitude                 valueTypes.Float     `json:"map_latitude" PointId:"MapLatitude" PointTimeSpan:"PointTimeSpanBoot"`
	MapLongitude                valueTypes.Float     `json:"map_longitude" PointId:"MapLongitude" PointTimeSpan:"PointTimeSpanBoot"`
	MeterReduceTotal            valueTypes.UnitValue `json:"meter_reduce_total" PointId:"MeterReduceTotal" PointTimeSpan:"PointTimeSpanTotal"`
	MobileTel                   valueTypes.String    `json:"moble_tel" PointId:"MobleTel" PointTimeSpan:"PointTimeSpanBoot"`
	MonthEnergy                 valueTypes.UnitValue `json:"month_energy" PointId:"MonthEnergy" PointTimeSpan:"PointTimeSpanMonthly"`
	MonthEnergyVirgin           valueTypes.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthIncome                 valueTypes.UnitValue `json:"month_income" PointId:"MonthIncome" PointTimeSpan:"PointTimeSpanMonthly"`
	NegativeLoadMsg             interface{}   `json:"negative_load_msg"`
	OwnerPsFaultStatus          string        `json:"owner_ps_fault_status" PointId:"OwnerPsFaultStatus" PointTimeSpan:"PointTimeSpanBoot"`
	P83081Map                   valueTypes.UnitValue `json:"p83081_map" PointId:"p83081" PointTimeSpan:"PointTimeSpanInstant"`
	P83081MapVirgin             valueTypes.UnitValue `json:"p83081_map_virgin"  PointIgnore:"true"`
	P83102Map                   valueTypes.UnitValue `json:"p83102_map" PointId:"p83102" PointTimeSpan:"PointTimeSpanInstant"`
	P83102MapVirgin             valueTypes.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent               valueTypes.Float     `json:"p83102_percent" PointId:"p83102" PointUnit:"%" PointTimeSpan:"PointTimeSpanInstant"`
	P83118Map                   valueTypes.UnitValue `json:"p83118_map" PointId:"p83118" PointTimeSpan:"PointTimeSpanInstant"`
	P83118MapVirgin             valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map                   valueTypes.UnitValue `json:"p83119_map" PointId:"p83119" PointTimeSpan:"PointTimeSpanInstant"`
	P83119MapVirgin             valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map                   valueTypes.UnitValue `json:"p83120_map" PointId:"p83120" PointTimeSpan:"PointTimeSpanInstant"`
	P83120MapVirgin             valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83122                      valueTypes.Float     `json:"p83122" PointId:"p83122" PointTimeSpan:"PointTimeSpanInstant"`
	P83124Map                   valueTypes.UnitValue `json:"p83124_map" PointId:"p83124" PointTimeSpan:"PointTimeSpanInstant"`
	P83124MapVirgin             valueTypes.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83202Map                   valueTypes.UnitValue `json:"p83202_map" PointId:"p83202" PointTimeSpan:"PointTimeSpanInstant"`
	P83202MapVirgin             valueTypes.UnitValue `json:"p83202_map_virgin"  PointIgnore:"true"`
	P83532MapVirgin             valueTypes.UnitValue `json:"p83532_map_virgin"  PointIgnore:"true"`
	PowerChargeSetted           valueTypes.Integer   `json:"power_charge_setted" PointId:"PowerChargeSetted" PointTimeSpan:"PointTimeSpanBoot"`
	PowerGridPowerMap           valueTypes.UnitValue `json:"power_grid_power_map" PointId:"PowerGridPowerMap" PointTimeSpan:"PointTimeSpanInstant"`
	PowerGridPowerMapVirgin     valueTypes.UnitValue `json:"power_grid_power_map_virgin"  PointIgnore:"true"`
	PsCountryID                 valueTypes.Integer   `json:"ps_country_id" PointId:"PsCountryID" PointTimeSpan:"PointTimeSpanBoot"`
	PsDeviceType                valueTypes.Integer   `json:"ps_device_type" PointId:"PsDeviceType" PointTimeSpan:"PointTimeSpanBoot"`
	PsFaultStatus               string        `json:"ps_fault_status" PointId:"PsFaultStatus" PointTimeSpan:"PointTimeSpanBoot"`
	PsHealthStatus              string        `json:"ps_health_status" PointId:"PsHealthStatus" PointTimeSpan:"PointTimeSpanBoot"`
	PsLocation                  valueTypes.String    `json:"ps_location" PointId:"PsLocation" PointTimeSpan:"PointTimeSpanBoot"`
	PsName                      valueTypes.String    `json:"ps_name" PointId:"PsName" PointTimeSpan:"PointTimeSpanBoot"`
	PsPsKey                     valueTypes.PsKey     `json:"ps_ps_key" PointId:"PsPsKey" PointTimeSpan:"PointTimeSpanBoot"`
	PsState                     string        `json:"ps_state" PointId:"PsState" PointTimeSpan:"PointTimeSpanBoot"`
	PsType                      valueTypes.Integer   `json:"ps_type" PointId:"PsType" PointTimeSpan:"PointTimeSpanBoot"`
	PvPowerMap                  valueTypes.UnitValue `json:"pv_power_map" PointId:"PvPowerMap" PointTimeSpan:"PointTimeSpanInstant"`
	PvPowerMapVirgin            valueTypes.UnitValue `json:"pv_power_map_virgin"  PointIgnore:"true"`
	RobotNumSweepCapacity       struct {
		Num           valueTypes.Integer `json:"num" PointId:"Num" PointTimeSpan:"PointTimeSpanBoot"`
		SweepCapacity valueTypes.Float   `json:"sweep_capacity" PointId:"SweepCapacity" PointTimeSpan:"PointTimeSpanBoot"`
	} `json:"robot_num_sweep_capacity"`
	SelfConsumptionOffsetReminder valueTypes.Integer   `json:"self_consumption_offset_reminder" PointId:"SelfConsumptionOffsetReminder" PointTimeSpan:"PointTimeSpanBoot"`
	So2ReduceTotal                valueTypes.UnitValue `json:"so2_reduce_total" PointId:"So2ReduceTotal" PointTimeSpan:"PointTimeSpanTotal"`
	StorageInverterData           []struct {
		CommunicationDevSn      valueTypes.String    `json:"communication_dev_sn" PointId:"CommunicationDevSn" PointName:"Serial No" PointTimeSpan:"PointTimeSpanBoot"`
		DevStatus               valueTypes.Integer   `json:"dev_status" PointId:"DevStatus" PointTimeSpan:"PointTimeSpanBoot"`
		DeviceCode              valueTypes.Integer   `json:"device_code" PointId:"DeviceCode" PointTimeSpan:"PointTimeSpanBoot"`
		DeviceModelCode         valueTypes.String    `json:"device_model_code" PointId:"DeviceModelCode" PointTimeSpan:"PointTimeSpanBoot"`
		DeviceName              valueTypes.String    `json:"device_name" PointId:"DeviceName" PointTimeSpan:"PointTimeSpanBoot"`
		DeviceState             string        `json:"device_state" PointId:"DeviceState" PointTimeSpan:"PointTimeSpanBoot"`
		DeviceType              valueTypes.Integer   `json:"device_type" PointId:"DeviceType" PointTimeSpan:"PointTimeSpanBoot"`
		DrmStatus               valueTypes.Integer   `json:"drm_status" PointId:"DrmStatus" PointTimeSpan:"PointTimeSpanBoot"`
		DrmStatusName           valueTypes.String    `json:"drm_status_name" PointId:"DrmStatusName" PointTimeSpan:"PointTimeSpanBoot"`
		EnergyFlow              []valueTypes.Integer `json:"energy_flow"`
		HasAmmeter              valueTypes.Bool      `json:"has_ammeter" PointId:"HasAmmeter" PointTimeSpan:"PointTimeSpanBoot"`
		InstallerDevFaultStatus valueTypes.Integer   `json:"installer_dev_fault_status" PointId:"InstallerDevFaultStatus" PointTimeSpan:"PointTimeSpanBoot"`
		InverterSn              valueTypes.String    `json:"inverter_sn" PointId:"InverterSn" PointTimeSpan:"PointTimeSpanBoot"`
		OwnerDevFaultStatus     valueTypes.Integer   `json:"owner_dev_fault_status" PointId:"OwnerDevFaultStatus" PointTimeSpan:"PointTimeSpanBoot"`
		P13003Map               valueTypes.UnitValue `json:"p13003_map" PointId:"p13003" PointTimeSpan:"PointTimeSpanInstant"`
		P13003MapVirgin         valueTypes.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13119Map               valueTypes.UnitValue `json:"p13119_map" PointId:"p13119" PointTimeSpan:"PointTimeSpanInstant"`
		P13119MapVirgin         valueTypes.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map               valueTypes.UnitValue `json:"p13121_map" PointId:"p13121" PointTimeSpan:"PointTimeSpanInstant"`
		P13121MapVirgin         valueTypes.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map               valueTypes.UnitValue `json:"p13126_map" PointId:"p13126" PointTimeSpan:"PointTimeSpanInstant"`
		P13126MapVirgin         valueTypes.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141                  valueTypes.Float     `json:"p13141" PointId:"p13141" PointTimeSpan:"PointTimeSpanInstant"`
		P13149Map               valueTypes.UnitValue `json:"p13149_map" PointId:"p13149" PointTimeSpan:"PointTimeSpanInstant"`
		P13149MapVirgin         valueTypes.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map               valueTypes.UnitValue `json:"p13150_map" PointId:"p13150" PointTimeSpan:"PointTimeSpanInstant"`
		P13150MapVirgin         valueTypes.UnitValue `json:"p13150_map_virgin"  PointIgnore:"true"`
		PsKey                   valueTypes.PsKey     `json:"ps_key" PointId:"PsKey" PointTimeSpan:"PointTimeSpanBoot"`
		UUID                    valueTypes.Integer   `json:"uuid" PointId:"UUID" PointTimeSpan:"PointTimeSpanBoot"`
	} `json:"storage_inverter_data"`
	TodayEnergy       valueTypes.UnitValue `json:"today_energy" PointId:"TodayEnergy" PointTimeSpan:"PointTimeSpanDaily"`
	TodayEnergyVirgin valueTypes.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome       valueTypes.UnitValue `json:"today_income" PointId:"TodayIncome" PointTimeSpan:"PointTimeSpanDaily"`
	TotalEnergy       valueTypes.UnitValue `json:"total_energy" PointId:"TotalEnergy" PointTimeSpan:"PointTimeSpanTotal"`
	TotalEnergyVirgin valueTypes.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalIncome       valueTypes.UnitValue `json:"total_income" PointId:"TotalIncome" PointTimeSpan:"PointTimeSpanTotal"`
	TreeReduceTotal   valueTypes.UnitValue `json:"tree_reduce_total" PointId:"TreeReduceTotal" PointTimeSpan:"PointTimeSpanTotal"`
	ValidFlag         valueTypes.Integer   `json:"valid_flag" PointId:"ValidFlag" PointTimeSpan:"PointTimeSpanBoot"`
	WgsLatitude       valueTypes.Float     `json:"wgs_latitude" PointId:"WgsLatitude" PointTimeSpan:"PointTimeSpanBoot"`
	WgsLongitude      valueTypes.Float     `json:"wgs_longitude" PointId:"WgsLongitude" PointTimeSpan:"PointTimeSpanBoot"`
	ZfzyMap           valueTypes.UnitValue `json:"zfzy_map" PointId:"ZfzyMap" PointTimeSpan:"PointTimeSpanInstant"`
	ZfzyMapVirgin     valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap           valueTypes.UnitValue `json:"zjzz_map" PointId:"ZjzzMap" PointTimeSpan:"PointTimeSpanInstant"`
	ZjzzMapVirgin     valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
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

func (e *EndPoint) GetData() api.DataMap {
	return e.Response.ResultData.GetData()
}

func (e *ResultData) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		name := fmt.Sprintf("%s.%s", pkg, e.PsPsKey.Value())
		entries.StructToPoints(*e, name, e.PsPsKey.Value(), valueTypes.NewDateTime(""))

		for _, sid := range e.StorageInverterData {
			name = fmt.Sprintf("%s.%s", pkg, sid.PsKey.Value())
			entries.StructToPoints(sid, name, sid.PsKey.Value(), valueTypes.NewDateTime(""))
		}
	}

	// api.Points.Print()
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
