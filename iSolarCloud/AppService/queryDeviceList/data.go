package queryDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/queryDeviceList"
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
	DevCountByStatusMap struct {
		FaultCount   valueTypes.Count `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		OfflineCount valueTypes.Count `json:"offline_count" PointId:"offline_count" PointUpdateFreq:"UpdateFreqTotal"`
		RunCount     valueTypes.Count `json:"run_count" PointId:"run_count" PointUpdateFreq:"UpdateFreqTotal"`
		WarningCount valueTypes.Count `json:"warning_count" PointId:"warning_count" PointUpdateFreq:"UpdateFreqTotal"`
	} `json:"dev_count_by_status_map" PointId:"device_status_count" DataTable:"true"`
	DevCountByTypeMap map[string]valueTypes.Integer `json:"dev_count_by_type_map" PointId:"device_type_count" PointUpdateFreq:"UpdateFreqBoot" DataTable:"true"`
	// DevCountByTypeMap struct {
	// 	One4 valueTypes.Integer `json:"14"`
	// 	Two2 valueTypes.Integer `json:"22"`
	// } `json:"dev_count_by_type_map"`
	DevTypeDefinition map[string]valueTypes.String `json:"dev_type_definition" PointId:"device_types" PointUpdateFreq:"UpdateFreqBoot" DataTable:"true"`
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
	PageList []struct {
		AlarmCount              valueTypes.Count   `json:"alarm_count" PointId:"alarm_count" PointUpdateFreq:"UpdateFreqTotal"`
		ChannelId               valueTypes.Integer `json:"chnnl_id" PointId:"channel_id" PointUpdateFreq:"UpdateFreqBoot"`
		CommandStatus           valueTypes.Integer `json:"command_status" PointId:"command_status" PointUpdateFreq:"UpdateFreqInstant"`
		ComponentAmount         valueTypes.Integer `json:"component_amount" PointId:"component_amount"`
		DataFlag                valueTypes.Integer `json:"data_flag" PointId:"data_flag" PointUpdateFreq:"UpdateFreqBoot"`
		DataFlagDetail          valueTypes.Integer `json:"data_flag_detail" PointId:"data_flag_detail"`
		DeviceArea              valueTypes.Integer `json:"device_area" PointId:"device_area" PointUpdateFreq:"UpdateFreqBoot"`			// References UUID and referenced by UUIDIndexCode
		DeviceAreaName          valueTypes.String  `json:"device_area_name" PointId:"device_area_name" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceCode              valueTypes.Integer `json:"device_code" PointId:"device_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceId                valueTypes.Integer `json:"device_id" PointId:"device_id" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceModelCode         valueTypes.String  `json:"device_model_code" PointId:"device_model_code" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceModelId           valueTypes.Integer `json:"device_model_id" PointId:"device_model_id" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceName              valueTypes.String  `json:"device_name" PointId:"device_name" PointUpdateFreq:"UpdateFreqBoot"`
		DeviceStatus            valueTypes.Bool    `json:"device_status" PointId:"device_status" PointUpdateFreq:"UpdateFreqInstant"`
		DeviceType              valueTypes.Integer `json:"device_type" PointId:"device_type" PointUpdateFreq:"UpdateFreqBoot"`
		FaultCount              valueTypes.Count   `json:"fault_count" PointId:"fault_count" PointUpdateFreq:"UpdateFreqTotal"`
		FaultStatus             string             `json:"fault_status" PointId:"fault_status" PointUpdateFreq:"UpdateFreqInstant"`
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
		PointData               []PointStruct      `json:"point_data" PointNameFromChild:"PointId" PointNameAppend:"false" DataTable:"true"`
		Points                  interface{}        `json:"points" PointId:"points"`
		PsTimezoneInfo          struct {
			IsDst    valueTypes.Bool   `json:"is_dst" PointUpdateFreq:"UpdateFreqInstant"`
			TimeZone valueTypes.String `json:"time_zone" PointUpdateFreq:"UpdateFreqInstant"`
		} `json:"psTimezoneInfo"`
		PsId          valueTypes.PsId `json:"ps_id" PointId:"ps_id" PointUpdateFreq:"UpdateFreqBoot"`
		PsKey         valueTypes.PsKey   `json:"ps_key" PointId:"ps_key" PointUpdateFreq:"UpdateFreqBoot"`
		RelState      valueTypes.Integer `json:"rel_state" PointId:"rel_state" PointUpdateFreq:"UpdateFreqInstant"`
		Sn            valueTypes.String  `json:"sn" PointId:"sn" PointName:"Serial Number" PointUpdateFreq:"UpdateFreqBoot"`
		StringAmount  valueTypes.Integer `json:"string_amount" PointId:"string_amount"`
		TypeName      valueTypes.String  `json:"type_name" PointId:"type_name" PointUpdateFreq:"UpdateFreqBoot"`
		UnitName      valueTypes.String  `json:"unit_name" PointId:"unit_name" PointUpdateFreq:"UpdateFreqBoot"`
		UUID          valueTypes.Integer `json:"uuid" PointId:"uuid" PointUpdateFreq:"UpdateFreqBoot"`							// Referenced by DeviceArea
		UUIDIndexCode valueTypes.String  `json:"uuid_index_code" PointId:"uuid_index_code" PointUpdateFreq:"UpdateFreqBoot"`	// Referenced by DeviceArea
	} `json:"pageList" PointId:"page_list" PointNameFromChild:"PsKey" PointNameAppend:"false" PointArrayFlatten:"false"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count" PointIgnore:"true"`
}

type PointStruct struct {
	PointId                valueTypes.PointId  `json:"point_id" PointIgnore:"true" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupName         valueTypes.String   `json:"point_group_name" PointIgnore:"true" PointUpdateFreq:"UpdateFreqBoot"`
	PointName              valueTypes.String   `json:"point_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	TimeStamp              valueTypes.DateTime `json:"time_stamp" PointIgnore:"true" PointUpdateFreq:"UpdateFreq5Mins"`
	Value                  valueTypes.Float    `json:"value" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUnitFrom:"Unit" PointUpdateFreq:"UpdateFreq5Mins"`
	Unit                   valueTypes.String   `json:"unit" PointIgnore:"true" PointUpdateFreq:"UpdateFreqBoot"`
	ValueDescription       valueTypes.String   `json:"value_description" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`

	CodeId                 valueTypes.Integer  `json:"code_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeIdOrderId          valueTypes.String   `json:"code_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	CodeName               valueTypes.String   `json:"code_name" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	DevPointLastUpdateTime valueTypes.DateTime `json:"dev_point_last_update_time" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreq5Mins"`
	IsPlatformDefaultUnit  valueTypes.Bool     `json:"is_platform_default_unit" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	IsShow                 valueTypes.Bool     `json:"is_show" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderId                valueTypes.Integer  `json:"order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	OrderNum               valueTypes.Integer  `json:"order_num" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupId           valueTypes.Integer  `json:"point_group_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointGroupIdOrderId    valueTypes.Integer  `json:"point_group_id_order_id" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
	PointSign              valueTypes.String   `json:"point_sign" PointGroupNameFrom:"PointGroupName" PointTimestampFrom:"TimeStamp" PointUpdateFreq:"UpdateFreqBoot"`
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
		// // Used for virtual entries.
		// // 0 - sungrow_battery_charging_power
		// var PVPowerToBattery VirtualPointStruct
		//
		// // sensor.sungrow_battery_discharging_power
		// var BatteryPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_total_export_active_power
		// var PVPowerToGrid VirtualPointStruct
		//
		// // sensor.sungrow_purchased_power
		// var GridPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		// var YieldBatteryCharge VirtualPointStruct
		// // var DailyBatteryChargingEnergy VirtualPointStruct
		//
		// // sensor.sungrow_daily_battery_discharging_energy
		// var DailyBatteryDischargingEnergy VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_feed_in_energy_pv
		// var YieldFeedIn VirtualPointStruct
		//
		// // sensor.sungrow_daily_purchased_energy
		// var DailyPurchasedEnergy VirtualPointStruct
		//
		// var PVPower VirtualPointStruct
		//
		// var LoadPower VirtualPointStruct
		//
		// var YieldSelfConsumption VirtualPointStruct
		// // var DailyFeedInEnergy VirtualPointStruct
		// var TotalPvYield VirtualPointStruct
		//
		// var DailyTotalLoad VirtualPointStruct
		//
		// var TotalEnergyConsumption VirtualPointStruct

		// // pkg := reflection.GetName("", *e)
		// name := api.JoinWithDots(0, "", pkg, e.Request.PsId)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

		if len(entries.Map) == 0 {
			break
		}

		e.GetEnergyStorageSystem(entries)
		e.GetCommunicationModule(entries)
		e.GetBattery(entries)
	}

	return entries
}

func (e *EndPoint) GetEnergyStorageSystem(entries api.DataMap) {

	for range Only.Once {
		// // Used for virtual entries.
		// // 0 - sungrow_battery_charging_power
		// var PVPowerToBattery VirtualPointStruct
		//
		// // sensor.sungrow_battery_discharging_power
		// var BatteryPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_total_export_active_power
		// var PVPowerToGrid VirtualPointStruct
		//
		// // sensor.sungrow_purchased_power
		// var GridPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		// var YieldBatteryCharge VirtualPointStruct
		// // var DailyBatteryChargingEnergy VirtualPointStruct
		//
		// // sensor.sungrow_daily_battery_discharging_energy
		// var DailyBatteryDischargingEnergy VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_feed_in_energy_pv
		// var YieldFeedIn VirtualPointStruct
		//
		// // sensor.sungrow_daily_purchased_energy
		// var DailyPurchasedEnergy VirtualPointStruct
		//
		// var PVPower VirtualPointStruct
		//
		// var LoadPower VirtualPointStruct
		//
		// var YieldSelfConsumption VirtualPointStruct
		// // var DailyFeedInEnergy VirtualPointStruct
		// var TotalPvYield VirtualPointStruct
		//
		// var DailyTotalLoad VirtualPointStruct
		//
		// var TotalEnergyConsumption VirtualPointStruct

		pkg := reflection.GetName("", *e)

		var devices []string
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
			devices = append(devices, api.JoinWithDots(0, "", device.PsId, device.PsKey))
		}

		// Points are in an array. So manually add virtuals instead of using the structure.
		for _, device := range devices {
			// fmt.Printf("endpoint: %s\n", device)
			dstEndpoint := "virtual." + device
			srcEndpoint := fmt.Sprintf("%s.%s", pkg, device)

			// BatteryChargingPower
			battery_charge_power := entries.CopyPoint(srcEndpoint + ".p13126.value", dstEndpoint, "battery_charge_power", "")

			// BatteryDischargingPower
			battery_discharge_power := entries.CopyPoint(srcEndpoint + ".p13150.value", dstEndpoint, "battery_discharge_power", "")

			// Daily PV Yield
			daily_pv_energy := entries.CopyPoint(srcEndpoint + ".p13112.value", dstEndpoint, "daily_pv_energy", "")

			// DailyBatteryChargingEnergyFromPv
			pv_battery_charge_energy := entries.CopyPoint(srcEndpoint + ".p13174.value", dstEndpoint, "pv_battery_charge_energy", "")

			// DailyBatteryDischargingEnergy
			battery_discharge := entries.CopyPoint(srcEndpoint + ".p13029.value", dstEndpoint, "battery_discharge", "")

			// DailyFeedInEnergy - @TODO - This may differ from DailyFeedInEnergyPv
			_ = entries.CopyPoint(srcEndpoint + ".p13122.value", dstEndpoint, "pv_feed_in2", "")
			// fmt.Println(pv_feed_in2)

			// DailyFeedInEnergyPv
			pv_feed_in := entries.CopyPoint(srcEndpoint + ".p13173.value", dstEndpoint, "pv_feed_in", "")

			// DailyPurchasedEnergy
			daily_purchased_energy := entries.CopyPoint(srcEndpoint + ".p13147.value", dstEndpoint, "daily_purchased_energy", "")

			// DailyLoadEnergyConsumptionFromPv
			pv_self_consumption := entries.CopyPoint(srcEndpoint + ".p13116.value", dstEndpoint, "pv_self_consumption", "")

			// TotalPvYield
			_ = entries.CopyPoint(srcEndpoint + ".p13134.value", dstEndpoint, "pv_total_yield", "")
			// fmt.Println(pv_total_yield)

			// Daily Load Energy Consumption
			daily_total_energy := entries.CopyPoint(srcEndpoint + ".p13199.value", dstEndpoint, "daily_total_energy", "")

			// Total Load Energy Consumption
			_ = entries.CopyPoint(srcEndpoint + ".p13130.value", dstEndpoint, "total_energy_consumption", "")
			// fmt.Println(total_energy_consumption)

			pv_daily_yield := entries.CopyDataEntries(*pv_self_consumption, dstEndpoint, "pv_daily_yield", "")
			pv_daily_yield.SetFloat(pv_self_consumption.GetFloat() + pv_battery_charge_energy.GetFloat() + pv_feed_in.GetFloat(), "", "")

			pv_self_consumption_percent := entries.CopyDataEntries(*daily_pv_energy, dstEndpoint, "pv_self_consumption_percent", "")
			pv_self_consumption_percent.SetFloat(entries.GetPercent(*pv_self_consumption, *daily_pv_energy), "", "")

			battery_energy := entries.CopyDataEntries(*pv_battery_charge_energy, dstEndpoint, "battery_energy", "")
			battery_energy.SetFloat(entries.LowerUpper(*pv_battery_charge_energy, *battery_discharge), "", "")

			pv_battery_charge_percent := entries.CopyDataEntries(*daily_pv_energy, dstEndpoint, "pv_battery_charge_percent", "")
			pv_battery_charge_percent.SetFloat(entries.LowerUpper(*pv_battery_charge_energy, *daily_pv_energy), "", "")

			pv_feed_in_percent := entries.CopyDataEntries(*daily_pv_energy, dstEndpoint, "pv_feed_in_percent", "")
			pv_feed_in_percent.SetFloat(entries.LowerUpper(*pv_feed_in, *daily_pv_energy), "", "")

			daily_pv_energy_percent := entries.CopyDataEntries(*daily_total_energy, dstEndpoint, "daily_pv_energy_percent", "")
			DailyPvEnergy := daily_total_energy.GetFloat() - daily_purchased_energy.GetFloat()
			daily_pv_energy_percent.SetFloat(api.GetPercent(DailyPvEnergy, daily_total_energy.GetFloat()), "", "")

			daily_purchased_energy_percent := entries.CopyDataEntries(*daily_total_energy, dstEndpoint, "daily_purchased_energy_percent", "")
			daily_purchased_energy_percent.SetFloat(entries.LowerUpper(*daily_purchased_energy, *daily_total_energy), "", "")


			// PV src
			power_pv := entries.CopyPoint(srcEndpoint + ".p13003.value", dstEndpoint, "power_pv", "") // TotalDcPower
			power_pv_active := entries.CopyDataEntries(*power_pv, dstEndpoint, "power_pv_active", "")
			power_pv_active.FloatToState(power_pv_active.GetFloat())

			power_pv_to_battery := entries.CopyDataEntries(*battery_charge_power, dstEndpoint, "power_pv_to_battery", "")
			power_pv_to_battery.SetFloat(battery_charge_power.GetFloat(), "", "")
			power_pv_to_battery_active := entries.CopyDataEntries(*power_pv_to_battery, dstEndpoint, "power_pv_to_battery_active", "")
			power_pv_to_battery_active.FloatToState(power_pv_to_battery_active.GetFloat())

			power_pv_to_grid := entries.CopyPoint(srcEndpoint + ".p13121.value", dstEndpoint, "power_pv_to_grid", "") // TotalExportActivePower
			power_pv_to_grid_active := entries.CopyDataEntries(*power_pv_to_grid, dstEndpoint, "power_pv_to_grid_active", "")
			power_pv_to_grid_active.FloatToState(power_pv_to_grid_active.GetFloat())

			power_pv_to_load := entries.CopyDataEntries(*power_pv, dstEndpoint, "power_pv_to_load", "")
			power_pv_to_load.SetFloat(power_pv.GetFloat() - battery_charge_power.GetFloat() - power_pv_to_grid.GetFloat(), "", "")
			power_pv_to_load_active := entries.CopyDataEntries(*power_pv_to_load, dstEndpoint, "power_pv_to_load_active", "")
			power_pv_to_load_active.FloatToState(power_pv_to_load_active.GetFloat())


			// Battery src
			power_battery := entries.CopyDataEntries(*battery_charge_power, dstEndpoint, "power_battery", "")
			power_battery.SetFloat(entries.LowerUpper(*battery_discharge_power, *battery_charge_power), "", "")
			power_battery_active := entries.CopyDataEntries(*power_battery, dstEndpoint, "power_battery_active", "")
			power_battery_active.FloatToState(power_battery_active.GetFloat())

			power_battery_to_load := entries.CopyDataEntries(*battery_discharge_power, dstEndpoint, "power_battery_to_load", "")
			power_battery_to_load.SetFloat(battery_discharge_power.GetFloat(), "", "")
			power_battery_to_load_active := entries.CopyDataEntries(*power_battery_to_load, dstEndpoint, "power_battery_to_load_active", "")
			power_battery_to_load_active.FloatToState(power_battery_to_load_active.GetFloat())

			power_battery_to_grid := entries.CopyDataEntries(*battery_charge_power, dstEndpoint, "power_battery_to_grid", "")
			power_battery_to_grid.SetFloat(0.0, "", "")
			power_battery_to_grid_active := entries.CopyDataEntries(*power_battery_to_grid, dstEndpoint, "power_battery_to_grid_active", "")
			power_battery_to_grid_active.FloatToState(power_battery_to_grid_active.GetFloat())


			// Grid src
			power_grid_to_load := entries.CopyPoint(srcEndpoint + ".p13149.value", dstEndpoint, "power_grid_to_load", "") // PurchasedPower
			power_grid_to_load_active := entries.CopyDataEntries(*power_grid_to_load, dstEndpoint, "power_grid_to_load_active", "")
			power_grid_to_load_active.FloatToState(power_grid_to_load_active.GetFloat())

			power_grid := entries.CopyDataEntries(*power_grid_to_load, dstEndpoint, "power_grid", "")
			power_grid.SetFloat(entries.LowerUpper(*power_pv_to_grid, *power_grid_to_load), "", "")
			power_grid_active := entries.CopyDataEntries(*power_grid, dstEndpoint, "power_grid_active", "")
			power_grid_active.FloatToState(power_grid_active.GetFloat())

			power_grid_to_battery := entries.CopyDataEntries(*power_grid_to_load, dstEndpoint, "power_grid_to_battery", "")
			power_grid_to_battery.SetFloat(0.0, "", "")
			power_grid_to_battery_active := entries.CopyDataEntries(*power_grid_to_battery, dstEndpoint, "power_grid_to_battery_active", "")
			power_grid_to_battery_active.FloatToState(power_grid_to_battery_active.GetFloat())

			grid_energy := entries.CopyDataEntries(*pv_feed_in, dstEndpoint, "grid_energy", "")
			grid_energy.SetFloat(entries.LowerUpper(*pv_feed_in, *daily_purchased_energy), "", "")


			// Load src
			power_load := entries.CopyPoint(srcEndpoint + ".p13119.value", dstEndpoint, "power_load", "") // TotalLoadActivePower
			power_load_active := entries.CopyDataEntries(*power_load, dstEndpoint, "power_load_active", "")
			power_load_active.FloatToState(power_load_active.GetFloat())
		}
	}
}

func (e *EndPoint) GetCommunicationModule(entries api.DataMap) {

	for range Only.Once {
		pkg := reflection.GetName("", *e)

		var devices []string
		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameCommunicationModule) {
				// Only looking for a Communication Module.
				continue
			}
			devices = append(devices, api.JoinWithDots(0, "", device.PsId, device.PsKey))
		}

		// Points are in an array. So manually add virtuals instead of using the structure.
		for _, device := range devices {
			// fmt.Printf("endpoint: %s\n", device)
			dstEndpoint := "virtual." + device
			srcEndpoint := fmt.Sprintf("%s.%s", pkg, device)

			// WLAN Signal Strength
			_ = entries.CopyPoint(srcEndpoint + ".p23014.value", dstEndpoint, "wlan_signal_strength", "")
		}
	}
}

func (e *EndPoint) GetBattery(entries api.DataMap) {

	for range Only.Once {
		pkg := reflection.GetName("", *e)

		var devices []string
		for _, device := range e.Response.ResultData.PageList {
			if !device.DeviceType.Match(api.DeviceNameBattery) {
				// Only looking for a Battery.
				continue
			}
			devices = append(devices, api.JoinWithDots(0, "", device.PsId, device.PsKey))
		}

		// Points are in an array. So manually add virtuals instead of using the structure.
		for _, device := range devices {
			// fmt.Printf("endpoint: %s\n", device)
			dstEndpoint := "virtual." + device
			srcEndpoint := fmt.Sprintf("%s.%s", pkg, device)

			// Battery Voltage
			_ = entries.CopyPoint(srcEndpoint + ".p58601.value", dstEndpoint, "battery_voltage", "")
			// Battery Current
			_ = entries.CopyPoint(srcEndpoint + ".p58602.value", dstEndpoint, "battery_current", "")
			// Battery Temperature
			_ = entries.CopyPoint(srcEndpoint + ".p58603.value", dstEndpoint, "battery_temperature", "")
			// Battery Level
			_ = entries.CopyPoint(srcEndpoint + ".p58604.value", dstEndpoint, "battery_level", "")
			// Battery Health (SOH)
			_ = entries.CopyPoint(srcEndpoint + ".p58605.value", dstEndpoint, "battery_health", "")
			// Total Battery Charging Energy
			_ = entries.CopyPoint(srcEndpoint + ".p58606.value", dstEndpoint, "total_battery_charging_energy", "")
			// Total Battery Discharging Energy
			_ = entries.CopyPoint(srcEndpoint + ".p58607.value", dstEndpoint, "total_battery_discharging_energy", "")
		}
	}
}
