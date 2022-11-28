package getPsDetailWithPsType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/powerStationService/getPsDetailWithPsType"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	// GoStruct                    GoStruct.GoStruct    `json:"-" PointIdFrom:"PsKey" PointIdReplace:"true"`

	BatteryLevelPercent         valueTypes.Float     `json:"battery_level_percent" PointId:"battery_level_percent" PointUnit:"%" PointUpdateFreq:"UpdateFreqInstant"`
	ChargingDischargingPowerMap valueTypes.UnitValue `json:"charging_discharging_power_map" PointId:"charging_discharging_power_map" PointUpdateFreq:"UpdateFreqInstant"`	// Holds the battery charge/discharge amount
	Co2ReduceTotal              valueTypes.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	CoalReduceTotal             valueTypes.UnitValue `json:"coal_reduce_total" PointId:"coal_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	ConnectType                 valueTypes.Integer   `json:"connect_type" PointId:"connect_type" PointUpdateFreq:"UpdateFreqBoot"`
	CurrPower                   valueTypes.UnitValue `json:"curr_power" PointId:"curr_power" PointUpdateFreq:"UpdateFreqInstant"`
	DesignCapacity              valueTypes.UnitValue `json:"design_capacity" PointId:"design_capacity" PointUpdateFreq:"UpdateFreqBoot"`
	EnergyScheme                interface{}          `json:"energy_scheme" PointId:"energy_scheme"`
	GcjLatitude                 valueTypes.Float     `json:"gcj_latitude" PointId:"gcj_latitude" PointUpdateFreq:"UpdateFreqBoot"`
	GcjLongitude                valueTypes.Float     `json:"gcj_longitude" PointId:"gcj_longitude" PointUpdateFreq:"UpdateFreqBoot"`
	HasAmmeter                  valueTypes.Bool      `json:"has_ammeter" PointId:"has_ammeter" PointUpdateFreq:"UpdateFreqBoot"`
	HouseholdInverterData       interface{}          `json:"household_inverter_data" PointId:"household_inverter_data"`
	InstallerPsFaultStatus      valueTypes.Integer   `json:"installer_ps_fault_status" PointId:"installer_ps_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
	IsHaveEsInverter            valueTypes.Bool      `json:"is_have_es_inverter" PointId:"is_have_es_inverter" PointUpdateFreq:"UpdateFreqBoot"`
	IsSingleInverter            valueTypes.Bool      `json:"is_single_inverter" PointId:"is_single_inverter" PointUpdateFreq:"UpdateFreqBoot"`
	IsTransformSystem           valueTypes.Bool      `json:"is_transform_system" PointId:"is_transform_system" PointUpdateFreq:"UpdateFreqBoot"`
	Latitude                    valueTypes.Float     `json:"latitude" PointId:"latitude" PointUpdateFreq:"UpdateFreqBoot"`
	LoadPowerMap                valueTypes.UnitValue `json:"load_power_map" PointId:"load_power_map" PointUpdateFreq:"UpdateFreqInstant"`
	LoadPowerMapVirgin          valueTypes.UnitValue `json:"load_power_map_virgin"  PointIgnore:"true"`
	Longitude                   valueTypes.Float     `json:"longitude" PointId:"longitude" PointUpdateFreq:"UpdateFreqBoot"`
	MapLatitude                 valueTypes.Float     `json:"map_latitude" PointId:"map_latitude" PointUpdateFreq:"UpdateFreqBoot"`
	MapLongitude                valueTypes.Float     `json:"map_longitude" PointId:"map_longitude" PointUpdateFreq:"UpdateFreqBoot"`
	MeterReduceTotal            valueTypes.UnitValue `json:"meter_reduce_total" PointId:"meter_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	MobileTel                   valueTypes.String    `json:"moble_tel" PointId:"mobile_tel" PointUpdateFreq:"UpdateFreqBoot"`
	MonthEnergy                 valueTypes.UnitValue `json:"month_energy" PointId:"month_energy" PointUpdateFreq:"UpdateFreqMonth"`
	MonthEnergyVirgin           valueTypes.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthIncome                 valueTypes.UnitValue `json:"month_income" PointId:"month_income" PointUpdateFreq:"UpdateFreqMonth"`
	NegativeLoadMsg             interface{}          `json:"negative_load_msg" PointId:"negative_load_msg"`
	OwnerPsFaultStatus          valueTypes.Integer   `json:"owner_ps_fault_status" PointId:"owner_ps_fault_status" PointUpdateFreq:"UpdateFreqInstant"`

	P83081Map                   valueTypes.UnitValue `json:"p83081_map" PointId:"p83081" PointName:"Load Power" PointUpdateFreq:"UpdateFreq5Mins"`
	P83081MapVirgin             valueTypes.UnitValue `json:"p83081_map_virgin"  PointIgnore:"true"`
	P83102Map                   valueTypes.UnitValue `json:"p83102_map" PointId:"p83102" PointName:"Energy Purchased" PointUpdateFreq:"UpdateFreq5Mins"`
	P83102MapVirgin             valueTypes.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent               valueTypes.Float     `json:"p83102_percent" PointId:"p83102_percent" PointName:"Energy Purchased Percent" PointUnit:"%" PointUpdateFreq:"UpdateFreq5Mins"`
	P83118Map                   valueTypes.UnitValue `json:"p83118_map" PointId:"p83118" PointName:"Energy Used" PointUpdateFreq:"UpdateFreq5Mins"`
	P83118MapVirgin             valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map                   valueTypes.UnitValue `json:"p83119_map" PointId:"p83119" PointName:"Energy Feed-In" PointUpdateFreq:"UpdateFreq5Mins"`
	P83119MapVirgin             valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map                   valueTypes.UnitValue `json:"p83120_map" PointId:"p83120" PointName:"Energy Battery Charge" PointUpdateFreq:"UpdateFreq5Mins"`
	P83120MapVirgin             valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83122                      valueTypes.Float     `json:"p83122" PointId:"p83122" PointName:"Self Sufficiency Percent" PointUnit:"%" PointUpdateFreq:"UpdateFreq5Mins"`
	P83124Map                   valueTypes.UnitValue `json:"p83124_map" PointId:"p83124" PointUpdateFreq:"UpdateFreq5Mins"`
	P83124MapVirgin             valueTypes.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83202Map                   valueTypes.UnitValue `json:"p83202_map" PointId:"p83202" PointName:"Installed Power" PointUpdateFreq:"UpdateFreq5Mins"`
	P83202MapVirgin             valueTypes.UnitValue `json:"p83202_map_virgin" PointIgnore:"true"`
	P83532MapVirgin             valueTypes.UnitValue `json:"p83532_map_virgin"  PointIgnore:"true"`

	PowerChargeSetted             valueTypes.Bool      `json:"power_charge_setted" PointId:"power_charge_set" PointUpdateFreq:"UpdateFreqBoot"`
	PowerGridPowerMap             valueTypes.UnitValue `json:"power_grid_power_map" PointId:"power_grid_power_map" PointUpdateFreq:"UpdateFreq5Mins"`
	PowerGridPowerMapVirgin       valueTypes.UnitValue `json:"power_grid_power_map_virgin"  PointIgnore:"true"`
	PsCountryId                   valueTypes.Integer   `json:"ps_country_id" PointId:"ps_country_id" PointUpdateFreq:"UpdateFreqBoot"`
	PsDeviceType                  valueTypes.Integer   `json:"ps_device_type" PointId:"ps_device_type" PointUpdateFreq:"UpdateFreqBoot"`
	PsFaultStatus                 valueTypes.Integer   `json:"ps_fault_status" PointId:"ps_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
	PsHealthStatus                valueTypes.Integer   `json:"ps_health_status" PointId:"ps_health_status" PointUpdateFreq:"UpdateFreqInstant"`
	PsLocation                    valueTypes.String    `json:"ps_location" PointId:"ps_location" PointUpdateFreq:"UpdateFreqBoot"`
	PsName                        valueTypes.String    `json:"ps_name" PointId:"ps_name" PointUpdateFreq:"UpdateFreqBoot"`
	PsKey                         valueTypes.PsKey     `json:"ps_ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
	PsState                       valueTypes.Bool      `json:"ps_state" PointId:"ps_state" PointUpdateFreq:"UpdateFreqInstant"`
	PsType                        valueTypes.Integer   `json:"ps_type" PointId:"ps_type" PointUpdateFreq:"UpdateFreqBoot"`
	PvPowerMap                    valueTypes.UnitValue `json:"pv_power_map" PointId:"pv_power_map" PointUpdateFreq:"UpdateFreq5Mins"`
	PvPowerMapVirgin              valueTypes.UnitValue `json:"pv_power_map_virgin"  PointIgnore:"true"`
	SelfConsumptionOffsetReminder valueTypes.Integer   `json:"self_consumption_offset_reminder" PointId:"self_consumption_offset_reminder" PointUpdateFreq:"UpdateFreqBoot"`
	So2ReduceTotal                valueTypes.UnitValue `json:"so2_reduce_total" PointId:"so2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	TodayEnergy                   valueTypes.UnitValue `json:"today_energy" PointId:"today_energy" PointUpdateFreq:"UpdateFreqDay"`
	TodayEnergyVirgin             valueTypes.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome                   valueTypes.UnitValue `json:"today_income" PointId:"today_income" PointUpdateFreq:"UpdateFreqDay"`
	TotalEnergy                   valueTypes.UnitValue `json:"total_energy" PointId:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`
	TotalEnergyVirgin             valueTypes.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalIncome                   valueTypes.UnitValue `json:"total_income" PointId:"total_income" PointUpdateFreq:"UpdateFreqTotal"`
	TreeReduceTotal               valueTypes.UnitValue `json:"tree_reduce_total" PointId:"tree_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	ValidFlag                     valueTypes.Bool      `json:"valid_flag" PointId:"valid_flag" PointUpdateFreq:"UpdateFreqBoot"`
	WgsLatitude                   valueTypes.Float     `json:"wgs_latitude" PointId:"wgs_latitude" PointUpdateFreq:"UpdateFreqBoot"`
	WgsLongitude                  valueTypes.Float     `json:"wgs_longitude" PointId:"wgs_longitude" PointUpdateFreq:"UpdateFreqBoot"`
	ZfzyMap                       valueTypes.UnitValue `json:"zfzy_map" PointId:"zfzy_map" PointName:"Self Consumption Of PV" PointUpdateFreq:"UpdateFreq5Mins"`
	ZfzyMapVirgin                 valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap                       valueTypes.UnitValue `json:"zjzz_map" PointId:"zjzz_map" PointName:"Self Sufficiency" PointUpdateFreq:"UpdateFreq5Mins"`
	ZjzzMapVirgin                 valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`

	StorageInverterData           []struct {
		GoStruct                GoStruct.GoStruct    `json:"-" PointIdFrom:"PsKey" PointIdReplace:"true"`

		CommunicationDevSn      valueTypes.String    `json:"communication_dev_sn" PointId:"communication_dev_sn" PointName:"Serial No" PointUpdateFreq:"UpdateFreqBoot"`
		DevStatus               valueTypes.Integer   `json:"dev_status" PointId:"dev_status" PointUpdateFreq:"UpdateFreqInstant"`
		DeviceCode              valueTypes.Integer   `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceModelCode         valueTypes.String    `json:"device_model_code" PointId:"device_model_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceName              valueTypes.String    `json:"device_name" PointId:"device_name" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceState             valueTypes.Integer   `json:"device_state" PointId:"device_state" PointUpdateFreq:"UpdateFreqInstant"`
		DeviceType              valueTypes.Integer   `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
		DrmStatus               valueTypes.Integer   `json:"drm_status" PointId:"drm_status" PointUpdateFreq:"UpdateFreqBoot"`
		DrmStatusName           valueTypes.String    `json:"drm_status_name" PointId:"drm_status_name" PointUpdateFreq:"UpdateFreqBoot"`
		EnergyFlow              []valueTypes.Integer `json:"energy_flow" PointId:"energy_flow" PointArrayFlatten:"true" PointUpdateFreq:"UpdateFreq5Mins"`
		HasAmmeter              valueTypes.Bool      `json:"has_ammeter" PointId:"has_ammeter" PointUpdateFreq:"UpdateFreqBoot"`
		InstallerDevFaultStatus valueTypes.Integer   `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
		InverterSn              valueTypes.String    `json:"inverter_sn" PointId:"inverter_sn" PointUpdateFreq:"UpdateFreqBoot"`
		OwnerDevFaultStatus     valueTypes.Integer   `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
		P13003Map               valueTypes.UnitValue `json:"p13003_map" PointId:"p13003" PointName:"PV Power To Load" PointUpdateFreq:"UpdateFreq5Mins"`
		P13003MapVirgin         valueTypes.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13119Map               valueTypes.UnitValue `json:"p13119_map" PointId:"p13119" PointName:"Load Power" PointUpdateFreq:"UpdateFreq5Mins"`
		P13119MapVirgin         valueTypes.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map               valueTypes.UnitValue `json:"p13121_map" PointId:"p13121" PointName:"PV Power To Grid" PointUpdateFreq:"UpdateFreq5Mins"`
		P13121MapVirgin         valueTypes.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map               valueTypes.UnitValue `json:"p13126_map" PointId:"p13126" PointName:"PV Power To Battery" PointUpdateFreq:"UpdateFreq5Mins"`
		P13126MapVirgin         valueTypes.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141                  valueTypes.Float     `json:"p13141" PointId:"p13141" PointName:"Battery Charge Percent" PointUnit:"%" PointUpdateFreq:"UpdateFreq5Mins"`
		P13149Map               valueTypes.UnitValue `json:"p13149_map" PointId:"p13149" PointName:"Grid Power To Load" PointUpdateFreq:"UpdateFreq5Mins"`
		P13149MapVirgin         valueTypes.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map               valueTypes.UnitValue `json:"p13150_map" PointId:"p13150" PointUpdateFreq:"UpdateFreq5Mins"`
		P13150MapVirgin         valueTypes.UnitValue `json:"p13150_map_virgin" PointIgnore:"true"`
		PsKey                   valueTypes.PsKey     `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
		UUID                    valueTypes.Integer   `json:"uuid" PointId:"uuid" PointUpdateFreq:"UpdateFreqBoot"`
	} `json:"storage_inverter_data" PointId:"inverter" PointIdReplace:"true"`	// PointIdFromChild:"PsKey" PointIdReplace:"true"`
	RobotNumSweepCapacity       struct {
		Num           valueTypes.Integer `json:"num" PointId:"num" PointUpdateFreq:"UpdateFreqBoot"`
		SweepCapacity valueTypes.Float   `json:"sweep_capacity" PointId:"sweep_capacity" PointUpdateFreq:"UpdateFreqBoot"`
	} `json:"robot_num_sweep_capacity" PointId:"robot"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// name := api.JoinWithDots(0, valueTypes.DateTimeLayoutDay, pkg, e.Response.ResultData.PsPsKey)
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))	// , e.Response.ResultData.PsKey.String()

		// dstEndpoint := "virtual." + e.Request.PsId.String()
		// srcEndpoint := fmt.Sprintf("%s.%s", pkg, e.Response.ResultData.PsPsKey.Value())

		for pn, device := range e.Response.ResultData.StorageInverterData {
			fmt.Println(pn)
			fmt.Println(device.PsKey)
		}

		// var devices []string
		for _, device := range e.Response.ResultData.StorageInverterData {
			if !device.DeviceType.Match(api.DeviceNameEnergyStorageSystem) {
				// Only looking for a Battery.
				continue
			}
		// 	devices = append(devices, device.PsKey.String())
		// }
		//
		// for _, device := range devices {
			// dstEndpoint = "virtual." + device
			entries.CopyPointFromName("P13003Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13003", "PV Power To Load")
			entries.CopyPointFromName("P13119Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13119", "Load Power")
			entries.CopyPointFromName("P13121Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13121", "PV Power To Grid")
			entries.CopyPointFromName("P13126Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13126", "PV Power To Battery")
			entries.CopyPointFromName("P13141", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13141", "Battery Charge Percent")
			entries.CopyPointFromName("P13149Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13149", "Grid Power To Load") // ??
			entries.CopyPointFromName("P13150Map", GoStruct.NewEndPointPath("virtual", device.PsKey.String()), "p13150", "")
		}

		// dstEndpoint = "virtual." + e.Response.ResultData.PsPsKey.Value()
		pskey := e.Response.ResultData.PsKey.Value()
		entries.CopyPointFromName("P83081Map", GoStruct.NewEndPointPath("virtual", pskey), "p83081", "Load Power")	// ?? - also getPsDetailWithPsType.1171348_11_0_0.curr_power
		entries.CopyPointFromName("P83102Map", GoStruct.NewEndPointPath("virtual", pskey), "p83102", "Energy Purchased")
		entries.CopyPointFromName("P83102Percent", GoStruct.NewEndPointPath("virtual", pskey, "p83102"), "percent", "Energy Purchased Percent")
		entries.CopyPointFromName("P83118Map", GoStruct.NewEndPointPath("virtual", pskey), "p83118", "Energy Used")
		entries.CopyPointFromName("P83119Map", GoStruct.NewEndPointPath("virtual", pskey), "p83119", "Energy Feed-In")
		entries.CopyPointFromName("P83120Map", GoStruct.NewEndPointPath("virtual", pskey), "p83120", "Energy Battery Charge")
		entries.CopyPointFromName("P83122", GoStruct.NewEndPointPath("virtual", pskey), "p83122", "Self Sufficiency Percent")
		entries.CopyPointFromName("P83124Map", GoStruct.NewEndPointPath("virtual", pskey), "p83124", "")
		entries.CopyPointFromName("P83202Map", GoStruct.NewEndPointPath("virtual", pskey), "p83202", "Installed Power")

		entries.CopyPointFromName("ZjzzMap", GoStruct.NewEndPointPath("virtual", pskey), "zjzz_map", "Self Sufficiency")
		entries.CopyPointFromName("ZfzyMap", GoStruct.NewEndPointPath("virtual", pskey), "zfzy_map", "Self Consumption Of PV")
	}

	return entries
}

func (e *EndPoint) GetPsKeys() []string {
	ret := []string{e.Response.ResultData.PsKey.Value()}
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = append(ret, l.PsKey.Value())
	}
	return ret
}

func (e *EndPoint) GetPsName() string {
	return e.Response.ResultData.PsName.Value()
}

func (e *EndPoint) GetPsState() string {
	return e.Response.ResultData.PsState.String()
}

func (e *EndPoint) GetPsKey() string {
	return e.Response.ResultData.PsKey.Value()
}

func (e *EndPoint) GetDeviceModelCode() string {
	ret := e.Response.ResultData.PsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.DeviceModelCode.Value()
		break
	}
	return ret
}

func (e *EndPoint) GetDeviceName() string {
	ret := e.Response.ResultData.PsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.DeviceName.Value()
		break
	}
	return ret
}

func (e *EndPoint) GetDeviceSerial() string {
	ret := e.Response.ResultData.PsKey.Value()
	for _, l := range e.Response.ResultData.StorageInverterData {
		ret = l.InverterSn.Value()
		break
	}
	return ret
}
