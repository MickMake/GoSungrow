package queryDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/queryDeviceList"
const Disabled = false
const EndPointName = "AppService.queryDeviceList"

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
	PageList []struct {
		GoStruct                GoStruct.GoStruct  `json:"-" PointIdFrom:"PsKey" PointIdReplace:"true" PointDeviceFrom:"PsKey"`

		PointData               []PointStruct      `json:"point_data" PointId:"data" PointIdReplace:"true" DataTable:"true"`	// PointIdFromChild:"PointId" PointIdReplace:"true" PointId:"data" DataTable:"true"`
		PsTimezoneInfo          struct {
			IsDst    valueTypes.Bool   `json:"is_dst" PointUpdateFreq:"UpdateFreqInstant"`
			TimeZone valueTypes.String `json:"time_zone" PointUpdateFreq:"UpdateFreqInstant"`
		}       `json:"psTimezoneInfo" PointId:"ps_timezone_info"`

		AlarmCount              valueTypes.Count   `json:"alarm_count" PointId:"alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
		ChannelId               valueTypes.Integer `json:"chnnl_id" PointId:"channel_id" PointUpdateFreq:"UpdateFreqBoot"`
		CommandStatus           valueTypes.Integer `json:"command_status" PointId:"command_status" PointUpdateFreq:"UpdateFreqInstant"`
		ComponentAmount         valueTypes.Integer `json:"component_amount" PointId:"component_amount"`
		DataFlag                valueTypes.Integer `json:"data_flag" PointId:"data_flag" PointUpdateFreq:"UpdateFreqBoot"`
		DataFlagDetail          valueTypes.Integer `json:"data_flag_detail" PointId:"data_flag_detail"`
		DeviceArea              valueTypes.Integer `json:"device_area" PointId:"device_area" PointUpdateFreq:"UpdateFreqBoot"` // References UUID and referenced by UUIDIndexCode
		DeviceAreaName          valueTypes.String  `json:"device_area_name" PointId:"device_area_name" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceCode              valueTypes.Integer `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceId                valueTypes.Integer `json:"device_id" PointId:"device_id" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceModelCode         valueTypes.String  `json:"device_model_code" PointId:"device_model_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceModelId           valueTypes.Integer `json:"device_model_id" PointId:"device_model_id" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceName              valueTypes.String  `json:"device_name" PointId:"device_name" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceStatus            valueTypes.Bool    `json:"device_status" PointId:"device_status" PointUpdateFreq:"UpdateFreqInstant"`
		DeviceType              valueTypes.Integer `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
		FaultCount              valueTypes.Count   `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		FaultStatus             valueTypes.String  `json:"fault_status" PointId:"fault_status" PointUpdateFreq:"UpdateFreqInstant"`
		FunctionEnum            valueTypes.String  `json:"function_enum" PointId:"function_enum" PointUpdateFreq:"UpdateFreqInstant"`
		InstallerAlarmCount     valueTypes.Count   `json:"installer_alarm_count" PointId:"installer_alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
		InstallerDevFaultStatus valueTypes.Integer `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
		InstallerFaultCount     valueTypes.Count   `json:"installer_fault_count" PointId:"installer_fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		InverterModelType       valueTypes.Integer `json:"inverter_model_type" PointId:"inverter_model_type" PointUpdateFreq:"UpdateFreqBoot"`
		IsDeveloper             valueTypes.Bool    `json:"is_developer" PointId:"is_developer" PointUpdateFreq:"UpdateFreqBoot"`
		IsG2point5Module        valueTypes.Bool    `json:"is_g2point5_module" PointId:"is_g2point5_module" PointUpdateFreq:"UpdateFreqBoot"`
		IsInit                  valueTypes.Bool    `json:"is_init" PointId:"is_init" PointUpdateFreq:"UpdateFreqBoot"`
		IsSecond                valueTypes.Bool    `json:"is_second" PointId:"is_second" PointUpdateFreq:"UpdateFreqBoot"`
		IsSupportParamset       valueTypes.Bool    `json:"is_support_paramset" PointId:"is_support_paramset" PointUpdateFreq:"UpdateFreqBoot"`
		NodeTimestamps          interface{}        `json:"node_timestamps" PointId:"node_timestamps"`
		OwnerAlarmCount         valueTypes.Count   `json:"owner_alarm_count" PointId:"owner_alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
		OwnerDevFaultStatus     valueTypes.Integer `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointUpdateFreq:"UpdateFreqInstant"`
		OwnerFaultCount         valueTypes.Count   `json:"owner_fault_count" PointId:"owner_fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		Points                  interface{}        `json:"points" PointId:"points"`
		PsId                    valueTypes.PsId    `json:"ps_id" PointId:"ps_id" PointUpdateFreq:"UpdateFreqBoot"`
		PsKey                   valueTypes.PsKey   `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
		RelState                valueTypes.Integer `json:"rel_state" PointId:"rel_state" PointUpdateFreq:"UpdateFreqInstant"`
		Sn                      valueTypes.String  `json:"sn" PointId:"sn" PointName:"Serial Number" PointUpdateFreq:"UpdateFreqBoot"`
		StringAmount            valueTypes.Integer `json:"string_amount" PointId:"string_amount"`
		TypeName                valueTypes.String  `json:"type_name" PointId:"type_name" PointUpdateFreq:"UpdateFreqBoot"`
		UnitName                valueTypes.String  `json:"unit_name" PointId:"unit_name" PointUpdateFreq:"UpdateFreqBoot"`
		UUID                    valueTypes.Integer `json:"uuid" PointId:"uuid" PointUpdateFreq:"UpdateFreqBoot"`                       // Referenced by DeviceArea
		UUIDIndexCode           valueTypes.String  `json:"uuid_index_code" PointId:"uuid_index_code" PointUpdateFreq:"UpdateFreqBoot"` // Referenced by DeviceArea
	} `json:"pageList" PointId:"devices"`

	DevCountByStatusMap struct {
		FaultCount   valueTypes.Count `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		OfflineCount valueTypes.Count `json:"offline_count" PointId:"offline_count" PointUpdateFreq:"UpdateFreqTotal"`
		RunCount     valueTypes.Count `json:"run_count" PointId:"run_count" PointUpdateFreq:"UpdateFreqTotal"`
		WarningCount valueTypes.Count `json:"warning_count" PointId:"warning_count" PointUpdateFreq:"UpdateFreqTotal"`
	} `json:"dev_count_by_status_map" PointId:"device_status_count"`
	DevCountByTypeMap map[string]valueTypes.Integer `json:"dev_count_by_type_map" PointId:"device_type_count" PointUpdateFreq:"UpdateFreqBoot"`
	DevTypeDefinition map[string]valueTypes.String `json:"dev_type_definition" PointId:"device_types" PointUpdateFreq:"UpdateFreqBoot"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}
// DevCountByTypeMap struct {
// 	One4 valueTypes.Integer `json:"14"`
// 	Two2 valueTypes.Integer `json:"22"`
// } `json:"dev_count_by_type_map"`
// DevTypeDefinition struct {
// 	One    valueTypes.String `json:"1"`
// 	One0   valueTypes.String `json:"10"`
// 	One1   valueTypes.String `json:"11"`
// 	One2   valueTypes.String `json:"12"`
// 	One3   valueTypes.String `json:"13"`
// 	One4   valueTypes.String `json:"14"`
// 	One5   valueTypes.String `json:"15"`
// 	One6   valueTypes.String `json:"16"`
// 	One7   valueTypes.String `json:"17"`
// 	One8   valueTypes.String `json:"18"`
// 	One9   valueTypes.String `json:"19"`
// 	Two0   valueTypes.String `json:"20"`
// 	Two1   valueTypes.String `json:"21"`
// 	Two2   valueTypes.String `json:"22"`
// 	Two3   valueTypes.String `json:"23"`
// 	Two4   valueTypes.String `json:"24"`
// 	Two5   valueTypes.String `json:"25"`
// 	Two6   valueTypes.String `json:"26"`
// 	Two8   valueTypes.String `json:"28"`
// 	Two9   valueTypes.String `json:"29"`
// 	Three  valueTypes.String `json:"3"`
// 	Three0 valueTypes.String `json:"30"`
// 	Three1 valueTypes.String `json:"31"`
// 	Three2 valueTypes.String `json:"32"`
// 	Three3 valueTypes.String `json:"33"`
// 	Three4 valueTypes.String `json:"34"`
// 	Three5 valueTypes.String `json:"35"`
// 	Three6 valueTypes.String `json:"36"`
// 	Three7 valueTypes.String `json:"37"`
// 	Three8 valueTypes.String `json:"38"`
// 	Three9 valueTypes.String `json:"39"`
// 	Four   valueTypes.String `json:"4"`
// 	Four0  valueTypes.String `json:"40"`
// 	Four1  valueTypes.String `json:"41"`
// 	Four2  valueTypes.String `json:"42"`
// 	Four3  valueTypes.String `json:"43"`
// 	Four4  valueTypes.String `json:"44"`
// 	Four5  valueTypes.String `json:"45"`
// 	Four6  valueTypes.String `json:"46"`
// 	Four7  valueTypes.String `json:"47"`
// 	Four8  valueTypes.String `json:"48"`
// 	Five   valueTypes.String `json:"5"`
// 	Five0  valueTypes.String `json:"50"`
// 	Six    valueTypes.String `json:"6"`
// 	Seven  valueTypes.String `json:"7"`
// 	Eight  valueTypes.String `json:"8"`
// 	Nine   valueTypes.String `json:"9"`
// 	Nine9  valueTypes.String `json:"99"`
// } `json:"dev_type_definition"`

type PointStruct struct {
	GoStruct               GoStruct.GoStruct   `json:"-" PointIdFrom:"PointId" PointIdReplace:"true" PointTimestampFrom:"TimeStamp" PointDeviceFromParent:"PsKey"`
	// GoStruct               GoStruct.GoStruct   `json:"-" PointDeviceFromParent:"PsKey"`

	TimeStamp              valueTypes.DateTime `json:"time_stamp" PointUpdateFreq:"UpdateFreq5Mins" PointNameDateFormat:"2006/01/02 15:04:05"`
	PointId                valueTypes.PointId  `json:"point_id" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupName         valueTypes.String   `json:"point_group_name" PointUpdateFreq:"UpdateFreqBoot"`
	PointName              valueTypes.String   `json:"point_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	Value                  valueTypes.Float    `json:"value" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUnitFrom:"Unit" PointVariableUnit:"true" PointUpdateFreq:"UpdateFreq5Mins"`
	Unit                   valueTypes.String   `json:"unit" PointUpdateFreq:"UpdateFreqBoot"`
	PointSign              valueTypes.String   `json:"point_sign" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	ValueDescription       valueTypes.String   `json:"value_description" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`

	CodeId                 valueTypes.Integer  `json:"code_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeIdOrderId          valueTypes.String   `json:"code_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeName               valueTypes.String   `json:"code_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	DevPointLastUpdateTime valueTypes.DateTime `json:"dev_point_last_update_time" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreq5Mins" PointNameDateFormat:"2006/01/02 15:04:05"`
	IsPlatformDefaultUnit  valueTypes.Bool     `json:"is_platform_default_unit" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	IsShow                 valueTypes.Bool     `json:"is_show" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderId                valueTypes.Integer  `json:"order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderNum               valueTypes.Integer  `json:"order_num" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupId           valueTypes.Integer  `json:"point_group_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupIdOrderId    valueTypes.Integer  `json:"point_group_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	Relate                 valueTypes.Integer  `json:"relate" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	ValIsFixed             valueTypes.Bool     `json:"val_is_fixd" PointId:"value_is_fixed" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	ValidSize              valueTypes.Integer  `json:"valid_size" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *ResultData) GetDataByName(name string) []PointStruct {
	var ret []PointStruct
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.DeviceName.Value() != name {
				continue
			}
			ret = p.PointData
			break
		}
	}
	return ret
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

		// // Used for virtual entries.
		// // 0 - sungrow_battery_charging_power
		// var PVPowerToBattery
		// // sensor.sungrow_battery_discharging_power
		// var BatteryPowerToLoad
		// // 0 - sensor.sungrow_total_export_active_power
		// var PVPowerToGrid
		// // sensor.sungrow_purchased_power
		// var GridPowerToLoad
		// // 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		// var YieldBatteryCharge
		// // var DailyBatteryChargingEnergy
		// // sensor.sungrow_daily_battery_discharging_energy
		// var DailyBatteryDischargingEnergy
		// // 0 - sensor.sungrow_daily_feed_in_energy_pv
		// var YieldFeedIn
		// // sensor.sungrow_daily_purchased_energy
		// var DailyPurchasedEnergy
		// var PVPower
		// var LoadPower
		// var YieldSelfConsumption
		// // var DailyFeedInEnergy
		// var TotalPvYield
		// var DailyTotalLoad
		// var TotalEnergyConsumption

		e.GetEnergyStorageSystem(entries)
		e.GetCommunicationModule(entries)
		e.GetBattery(entries)
	}

	return entries
}

func (e *EndPoint) GetEnergyStorageSystem(entries api.DataMap) {
	for range Only.Once {
		/*
			PVPower				- TotalDcPower
			PVPowerToBattery	- BatteryChargingPower
			PVPowerToLoad		- TotalDcPower - BatteryChargingPower - TotalExportActivePower
			PVPowerToGrid		- TotalExportActivePower

			LoadPower			- TotalLoadActivePower
			BatteryPowerToLoad	- BatteryDischargingPower
			BatteryPowerToGrid	- ?

			GridPower			- lowerUpper(PVPowerToGrid, GridPowerToLoad)
			GridPowerToLoad		- PurchasedPower
			GridPowerToBattery	- ?

			YieldSelfConsumption	- DailyLoadEnergyConsumptionFromPv
			YieldBatteryCharge		- DailyBatteryChargingEnergyFromPv
			YieldFeedIn				- DailyFeedInEnergyPv
		*/

		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameEnergyStorageSystem) {
				// Only looking for a Solar Storage System.
				continue
			}
			epp := GoStruct.NewEndPointPath("virtual", device.PsId.String(), device.PsKey.String())
			// Points are embedded within []PointStruct. So manually add virtuals instead of using the structure.

			// BatteryChargingPower
			batteryChargePower := entries.CopyPointFromName("p13126.value", epp, "battery_charge_power", "")

			// BatteryDischargingPower
			batteryDischargePower := entries.CopyPointFromName("p13150.value", epp, "battery_discharge_power", "")

			// Daily PV Yield
			dailyPvEnergy := entries.CopyPointFromName("p13112.value", epp, "daily_pv_energy", "")

			// DailyBatteryChargingEnergyFromPv
			pvBatteryChargeEnergy := entries.CopyPointFromName("p13174.value", epp, "pv_battery_charge_energy", "")

			// DailyBatteryDischargingEnergy
			batteryDischarge := entries.CopyPointFromName("p13029.value", epp, "battery_discharge", "")

			// DailyFeedInEnergy - @TODO - This may differ from DailyFeedInEnergyPv
			_ = entries.CopyPointFromName("p13122.value", epp, "pv_feed_in2", "")

			// DailyFeedInEnergyPv
			pvFeedIn := entries.CopyPointFromName("p13173.value", epp, "pv_feed_in", "")

			// DailyPurchasedEnergy
			dailyPurchasedEnergy := entries.CopyPointFromName("p13147.value", epp, "daily_purchased_energy", "")

			// DailyLoadEnergyConsumptionFromPv
			pvSelfConsumption := entries.CopyPointFromName("p13116.value", epp, "pv_self_consumption", "")

			// TotalPvYield
			_ = entries.CopyPointFromName("p13134.value", epp, "pv_total_yield", "")

			// Daily Load Energy Consumption
			dailyTotalEnergy := entries.CopyPointFromName("p13199.value", epp, "daily_total_energy", "")

			// Total Load Energy Consumption
			_ = entries.CopyPointFromName("p13130.value", epp, "total_energy_consumption", "")

			pvDailyYield := entries.CopyPointFromName(pvSelfConsumption.PointId(), epp, "pv_daily_yield", "")
			pvDailyYield.SetValue(GoStruct.AddFloatValues(pvSelfConsumption, pvBatteryChargeEnergy, pvFeedIn))

			pvSelfConsumptionPercent := entries.CopyPoint(dailyPvEnergy, epp, "pv_self_consumption_percent", "")
			pvSelfConsumptionPercent.SetValue(entries.GetPercent(pvSelfConsumption, dailyPvEnergy))

			batteryEnergy := entries.CopyPoint(pvBatteryChargeEnergy, epp, "battery_energy", "")
			batteryEnergy.SetValue(entries.LowerUpper(pvBatteryChargeEnergy, batteryDischarge))

			pvBatteryChargePercent := entries.CopyPoint(dailyPvEnergy, epp, "pv_battery_charge_percent", "")
			pvBatteryChargePercent.SetValue(entries.LowerUpper(pvBatteryChargeEnergy, dailyPvEnergy))

			pvFeedInPercent := entries.CopyPoint(dailyPvEnergy, epp, "pv_feed_in_percent", "")
			pvFeedInPercent.SetValue(entries.LowerUpper(pvFeedIn, dailyPvEnergy))

			dailyPvEnergyPercent := entries.CopyPoint(dailyTotalEnergy, epp, "daily_pv_energy_percent", "")
			DailyPvEnergy := dailyTotalEnergy.Value.First().ValueFloat() - dailyPurchasedEnergy.Value.First().ValueFloat()
			dailyPvEnergyPercent.SetValue(api.GetPercent(DailyPvEnergy, dailyTotalEnergy.Value.First().ValueFloat()))

			dailyPurchasedEnergyPercent := entries.CopyPoint(dailyTotalEnergy, epp, "daily_purchased_energy_percent", "")
			dailyPurchasedEnergyPercent.SetValue(entries.LowerUpper(dailyPurchasedEnergy, dailyTotalEnergy))


			// PV src
			powerPv := entries.CopyPointFromName("p13003.value", epp, "power_pv", "Total DC Power")
			_ = entries.MakeState(powerPv, epp, "power_pv_active", "")

			powerPvToBattery := entries.CopyPoint(batteryChargePower, epp, "power_pv_to_battery", "")
			powerPvToBattery.SetValue(batteryChargePower.Value.First().ValueFloat())
			_ = entries.MakeState(powerPvToBattery, epp, "power_pv_to_battery_active", "")

			powerPvToGrid := entries.CopyPointFromName("p13121.value", epp, "power_pv_to_grid", "Total Export Active Power")
			_ = entries.MakeState(powerPvToGrid, epp, "power_pv_to_grid_active", "")

			powerPvToLoad := entries.CopyPoint(powerPv, epp, "power_pv_to_load", "")
			powerPvToLoad.SetValue(powerPv.Value.First().ValueFloat() - batteryChargePower.Value.First().ValueFloat() - powerPvToGrid.Value.First().ValueFloat())
			_ = entries.MakeState(powerPvToLoad, epp, "power_pv_to_load_active", "")


			// Battery src
			powerBattery := entries.CopyPoint(batteryChargePower, epp, "power_battery", "")
			powerBattery.SetValue(entries.LowerUpper(batteryDischargePower, batteryChargePower))
			_ = entries.MakeState(powerBattery, epp, "power_battery_active", "")

			powerBatteryToLoad := entries.CopyPoint(batteryDischargePower, epp, "power_battery_to_load", "")
			powerBatteryToLoad.SetValue(batteryDischargePower.Value.First().ValueFloat())
			_ = entries.MakeState(powerBatteryToLoad, epp, "power_battery_to_load_active", "")

			powerBatteryToGrid := entries.CopyPoint(batteryChargePower, epp, "power_battery_to_grid", "")
			powerBatteryToGrid.SetValue(0.0)
			_ = entries.MakeState(powerBatteryToGrid, epp, "power_battery_to_grid_active", "")


			// Grid src
			powerGridToLoad := entries.CopyPointFromName("p13149.value", epp, "power_grid_to_load", "Purchased Power")
			_ = entries.MakeState(powerGridToLoad, epp, "power_grid_to_load_active", "")

			powerGrid := entries.CopyPoint(powerGridToLoad, epp, "power_grid", "")
			powerGrid.SetValue(entries.LowerUpper(powerPvToGrid, powerGridToLoad))
			_ = entries.MakeState(powerGrid, epp, "power_grid_active", "")

			powerGridToBattery := entries.CopyPoint(powerGridToLoad, epp, "power_grid_to_battery", "")
			powerGridToBattery.SetValue(0.0)
			_ = entries.MakeState(powerGridToBattery, epp, "power_grid_to_battery_active", "")

			gridEnergy := entries.CopyPoint(pvFeedIn, epp, "grid_energy", "")
			gridEnergy.SetValue(entries.LowerUpper(pvFeedIn, dailyPurchasedEnergy))


			// Load src
			powerLoad := entries.CopyPointFromName("p13119.value", epp, "power_load", "Total Load Active Power")
			_ = entries.MakeState(powerLoad, epp, "power_load_active", "")
		}
	}
}

func (e *EndPoint) GetCommunicationModule(entries api.DataMap) {
	for range Only.Once {
		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameCommunicationModule) {
				// Only looking for a Communication Module.
				continue
			}
			epp := GoStruct.NewEndPointPath("virtual", device.PsId.String(), device.PsKey.String())
			// Points are embedded within []PointStruct. So manually add virtuals instead of using the structure.

			// WLAN Signal Strength
			_ = entries.CopyPointFromName("p23014.value", epp, "wlan_signal_strength", "")
		}
	}
}

func (e *EndPoint) GetBattery(entries api.DataMap) {
	for range Only.Once {
		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameBattery) {
				// Only looking for a Battery.
				continue
			}
			epp := GoStruct.NewEndPointPath("virtual", device.PsId.String(), device.PsKey.String())
			// Points are embedded within []PointStruct. So manually add virtuals instead of using the structure.

			// Battery Voltage
			_ = entries.CopyPointFromName("p58601.value", epp, "battery_voltage", "")
			// Battery Current
			_ = entries.CopyPointFromName("p58602.value", epp, "battery_current", "")
			// Battery Temperature
			_ = entries.CopyPointFromName("p58603.value", epp, "battery_temperature", "")
			// Battery Level
			_ = entries.CopyPointFromName("p58604.value", epp, "battery_level", "")
			// Battery Health (SOH)
			_ = entries.CopyPointFromName("p58605.value", epp, "battery_health", "")
			// Total Battery Charging Energy
			_ = entries.CopyPointFromName("p58606.value", epp, "total_battery_charging_energy", "")
			// Total Battery Discharging Energy
			_ = entries.CopyPointFromName("p58607.value", epp, "total_battery_discharging_energy", "")
		}
	}
}
