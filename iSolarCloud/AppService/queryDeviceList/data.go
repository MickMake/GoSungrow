package queryDeviceList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"sort"
	"time"
)

const Url = "/v1/devService/queryDeviceList"
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
	DevCountByStatusMap struct {
		FaultCount   api.Count   `json:"fault_count" PointId:"fault_count" PointValueType:"" PointTimeSpan:""`
		OfflineCount api.Count   `json:"offline_count" PointId:"offline_count" PointValueType:"" PointTimeSpan:""`
		RunCount     api.Count   `json:"run_count" PointId:"run_count" PointValueType:"" PointTimeSpan:""`
		WarningCount api.Count   `json:"warning_count" PointId:"warning_count" PointValueType:"" PointTimeSpan:""`
	} `json:"dev_count_by_status_map"`
	DevCountByTypeMap map[string]api.Integer `json:"dev_count_by_type_map"`
	// DevCountByTypeMap struct {
	// 	One4 api.Integer `json:"14"`
	// 	Two2 api.Integer `json:"22"`
	// } `json:"dev_count_by_type_map"`
	DevTypeDefinition map[string]api.String `json:"dev_type_definition"`
	// DevTypeDefinition struct {
	// 	One    string `json:"1"`
	// 	One0   string `json:"10"`
	// 	One1   string `json:"11"`
	// 	One2   string `json:"12"`
	// 	One3   string `json:"13"`
	// 	One4   string `json:"14"`
	// 	One5   string `json:"15"`
	// 	One6   string `json:"16"`
	// 	One7   string `json:"17"`
	// 	One8   string `json:"18"`
	// 	One9   string `json:"19"`
	// 	Two0   string `json:"20"`
	// 	Two1   string `json:"21"`
	// 	Two2   string `json:"22"`
	// 	Two3   string `json:"23"`
	// 	Two4   string `json:"24"`
	// 	Two5   string `json:"25"`
	// 	Two6   string `json:"26"`
	// 	Two8   string `json:"28"`
	// 	Two9   string `json:"29"`
	// 	Three  string `json:"3"`
	// 	Three0 string `json:"30"`
	// 	Three1 string `json:"31"`
	// 	Three2 string `json:"32"`
	// 	Three3 string `json:"33"`
	// 	Three4 string `json:"34"`
	// 	Three5 string `json:"35"`
	// 	Three6 string `json:"36"`
	// 	Three7 string `json:"37"`
	// 	Three8 string `json:"38"`
	// 	Three9 string `json:"39"`
	// 	Four   string `json:"4"`
	// 	Four0  string `json:"40"`
	// 	Four1  string `json:"41"`
	// 	Four2  string `json:"42"`
	// 	Four3  string `json:"43"`
	// 	Four4  string `json:"44"`
	// 	Four5  string `json:"45"`
	// 	Four6  string `json:"46"`
	// 	Four7  string `json:"47"`
	// 	Four8  string `json:"48"`
	// 	Five   string `json:"5"`
	// 	Five0  string `json:"50"`
	// 	Six    string `json:"6"`
	// 	Seven  string `json:"7"`
	// 	Eight  string `json:"8"`
	// 	Nine   string `json:"9"`
	// 	Nine9  string `json:"99"`
	// } `json:"dev_type_definition"`
	PageList []struct {
		AlarmCount              api.Count   `json:"alarm_count" PointId:"alarm_count" PointValueType:"" PointTimeSpan:""`
		ChannelId               api.Integer `json:"chnnl_id" PointId:"channel_id" PointValueType:"" PointTimeSpan:""`
		CommandStatus           api.Integer `json:"command_status" PointId:"command_status" PointValueType:"" PointTimeSpan:""`
		ComponentAmount         api.Integer `json:"component_amount" PointId:"component_amount" PointValueType:"" PointTimeSpan:""`
		DataFlag                api.Integer `json:"data_flag" PointId:"data_flag" PointValueType:"" PointTimeSpan:""`
		DataFlagDetail          api.Integer `json:"data_flag_detail" PointId:"data_flag_detail" PointValueType:"" PointTimeSpan:""`
		DeviceArea              api.Integer `json:"device_area" PointId:"device_area" PointValueType:"" PointTimeSpan:""`
		DeviceAreaName          api.String  `json:"device_area_name" PointId:"device_area_name" PointValueType:"" PointTimeSpan:""`
		DeviceCode              api.Integer `json:"device_code" PointId:"device_code" PointValueType:"" PointTimeSpan:""`
		DeviceID                api.Integer `json:"device_id" PointId:"device_id" PointValueType:"" PointTimeSpan:""`
		DeviceModelCode         api.String  `json:"device_model_code" PointId:"device_model_code" PointValueType:"" PointTimeSpan:""`
		DeviceModelID           api.Integer `json:"device_model_id" PointId:"device_model_id" PointValueType:"" PointTimeSpan:""`
		DeviceName              api.String  `json:"device_name" PointId:"device_name" PointValueType:"" PointTimeSpan:""`
		DeviceStatus            api.Bool    `json:"device_status" PointId:"device_status" PointValueType:"" PointTimeSpan:""`
		DeviceType              api.Integer `json:"device_type" PointId:"device_type" PointValueType:"" PointTimeSpan:""`
		FaultCount              api.Count   `json:"fault_count" PointId:"fault_count" PointValueType:"" PointTimeSpan:""`
		FaultStatus             string      `json:"fault_status" PointId:"fault_status" PointValueType:"" PointTimeSpan:""`
		FunctionEnum            api.String  `json:"function_enum" PointId:"function_enum" PointValueType:"" PointTimeSpan:""`
		InstallerAlarmCount     api.Count   `json:"installer_alarm_count" PointId:"installer_alarm_count" PointValueType:"" PointTimeSpan:""`
		InstallerDevFaultStatus api.Integer `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointValueType:"" PointTimeSpan:""`
		InstallerFaultCount     api.Count   `json:"installer_fault_count" PointId:"installer_fault_count" PointValueType:"" PointTimeSpan:""`
		InverterModelType       api.Integer `json:"inverter_model_type" PointId:"inverter_model_type" PointValueType:"" PointTimeSpan:""`
		IsDeveloper             api.Bool    `json:"is_developer" PointId:"is_developer" PointValueType:"" PointTimeSpan:""`
		IsG2point5Module        api.Bool    `json:"is_g2point5_module" PointId:"is_g2point5_module" PointValueType:"" PointTimeSpan:""`
		IsInit                  api.Bool    `json:"is_init" PointId:"is_init" PointValueType:"" PointTimeSpan:""`
		IsSecond                api.Bool    `json:"is_second" PointId:"is_second" PointValueType:"" PointTimeSpan:""`
		IsSupportParamset       api.Bool    `json:"is_support_paramset" PointId:"is_support_paramset" PointValueType:"" PointTimeSpan:""`
		NodeTimestamps          interface{} `json:"node_timestamps" PointId:"node_timestamps" PointValueType:"" PointTimeSpan:""`
		OwnerAlarmCount         api.Count   `json:"owner_alarm_count" PointId:"owner_alarm_count" PointValueType:"" PointTimeSpan:""`
		OwnerDevFaultStatus     api.Integer `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointValueType:"" PointTimeSpan:""`
		OwnerFaultCount         api.Count   `json:"owner_fault_count" PointId:"owner_fault_count" PointValueType:"" PointTimeSpan:""`
		PointData               PointData   `json:"point_data"`
		Points                  interface{} `json:"points" PointId:"points" PointValueType:"" PointTimeSpan:""`
		PsTimezoneInfo          struct {
			IsDst    api.Bool   `json:"is_dst"`
			TimeZone api.String `json:"time_zone"`
		} `json:"psTimezoneInfo"`
		PsID          api.Integer `json:"ps_id" PointId:"ps_id" PointValueType:"" PointTimeSpan:""`
		PsKey         api.PsKey   `json:"ps_key" PointId:"ps_key" PointValueType:"" PointTimeSpan:""`
		RelState      api.Integer `json:"rel_state" PointId:"rel_state" PointValueType:"" PointTimeSpan:""`
		Sn            api.String  `json:"sn" PointId:"sn" PointValueType:"" PointTimeSpan:""`
		StringAmount  api.Integer `json:"string_amount" PointId:"string_amount" PointValueType:"" PointTimeSpan:""`
		TypeName      api.String  `json:"type_name" PointId:"type_name" PointValueType:"" PointTimeSpan:""`
		UnitName      api.String  `json:"unit_name" PointId:"unit_name" PointValueType:"" PointTimeSpan:""`
		UUID          api.Integer `json:"uuid" PointId:"uuid" PointValueType:"" PointTimeSpan:""`
		UUIDIndexCode api.String  `json:"uuid_index_code" PointId:"uuid_index_code" PointValueType:"" PointTimeSpan:""`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
}

type PointData []PointStruct
type PointStruct struct {
	CodeID                 api.Integer  `json:"code_id"`
	CodeIDOrderID          api.String   `json:"code_id_order_id"`
	CodeName               api.String   `json:"code_name"`
	DevPointLastUpdateTime api.DateTime `json:"dev_point_last_update_time"`
	IsPlatformDefaultUnit  api.Bool     `json:"is_platform_default_unit"`
	IsShow                 api.Bool     `json:"is_show"`
	OrderID                api.Integer  `json:"order_id"`
	OrderNum               api.Integer  `json:"order_num"`
	PointGroupID           api.Integer  `json:"point_group_id"`
	PointGroupIDOrderID    string       `json:"point_group_id_order_id"`
	PointGroupName         api.String   `json:"point_group_name"`
	PointID                api.Integer  `json:"point_id"`
	PointName              api.String   `json:"point_name"`
	PointSign              api.String   `json:"point_sign"`
	Relate                 api.Integer  `json:"relate"`
	TimeStamp              api.DateTime `json:"time_stamp"`
	Unit                   api.String   `json:"unit"`
	ValIsFixd              string       `json:"val_is_fixd"`
	ValidSize              api.Integer  `json:"valid_size"`
	Value                  api.Float    `json:"value"`
	ValueDescription       api.String   `json:"value_description"`
}

// type VirtualPointStruct struct {
// 	api.DataEntry
// 	ValueFloat api.Float
// }


func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
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
// }

func (e *ResultData) GetDataByName(name string) PointData {
	var ret PointData
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
			// "Parents",
			"Group Name",
			"Description",
			"Value",
			"Unit",
		)

		data := e.GetData()
		var sorted []string
		for p := range data.DataPoints {
			sorted = append(sorted, string(p))
		}
		sort.Strings(sorted)

		for _, p := range sorted {
			entries := data.DataPoints[api.PointId(p)]
			for _, de := range entries {
				if de.Hide {
					continue
				}

				_ = table.AddRow(
					de.Date.Format(api.DtLayout),
					// api.NameDevicePointInt(de.Point.Parents, p.PointID.Value()),
					// de.Point.Id,
					p,
					// de.Point.Parents.String(),
					de.Point.GroupName,
					de.Point.Name,
					de.Value,
					de.Point.Unit,
				)
			}
		}

		// for _, d := range e.Response.ResultData.PageList {
		// 	for _, p := range d.PointData {
		// 		// p.Value, p.Unit = api.DivideByThousandIfRequired(p.Value, p.Unit)
		// 		uv := api.SetUnitValueFloat(p.Value.Value(), p.Unit)
		// 		_ = table.AddRow(
		// 			api.NewDateTime(p.TimeStamp).PrintFull(),
		// 			api.NameDevicePointInt(d.PsKey, p.PointID.Value()),
		// 			p.PointGroupName,
		// 			p.PointName,
		// 			uv.Value(),
		// 			uv.Unit(),
		// 		)
		// 	}
		// }

		table.InitGraph(output.GraphRequest {
			Title:        "",
			TimeColumn:   output.SetInteger(1),
			SearchColumn: output.SetInteger(2),
			NameColumn:   output.SetInteger(4),
			ValueColumn:  output.SetInteger(5),
			UnitsColumn:  output.SetInteger(6),
			SearchString: output.SetString(""),
			MinLeftAxis:  output.SetFloat(0),
			MaxLeftAxis:  output.SetFloat(0),
		})

	}
	return table
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

		name1 := "queryDeviceList." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData.DevCountByStatusMap, name1 + ".status", e.Request.PsId.String(), time.Time{})

		for _, d := range e.Response.ResultData.PageList {
			name2 := fmt.Sprintf("queryDeviceList.%s", d.PsKey.Value())
			entries.StructToPoints(d, name2, d.PsKey.Value(), time.Time{})

			// entries.StructToPoints(d.PsTimezoneInfo, name2 + "PsTimezoneInfo", d.PsKey.Value(), time.Time{})

			for _, p := range d.PointData {
				pid := api.SetPointInt(p.PointID.Value())
				uv := api.SetUnitValueFloat(p.Value.Value(), p.Unit.Value())
				// name2 := fmt.Sprintf("%s.PointData.%s", name, pid)
				// name3 := fmt.Sprintf("%s.PointData", name2)
				entries.AddUnitValue(name2, d.PsKey.Value(), pid, p.PointName.Value(), p.PointGroupName.Value(), p.TimeStamp, uv)

				// Handle virtual results.
				// switch pid {
				// 	case "13126":
				// 		// BatteryChargingPower
				// 		entries["PVPowerToBattery"] = entries[pid]
				// 	case "13150":
				// 		// BatteryDischargingPower
				// 		entries["BatteryPowerToLoad"] = entries[pid]
				// 	case "13121":
				// 		// TotalExportActivePower
				// 		entries["PVPowerToGrid"] = entries[pid]
				// 	case "13149":
				// 		// PurchasedPower
				// 		entries["GridPowerToLoad"] = entries[pid]
				// 	case "13003":
				// 		// TotalDcPower
				// 		entries["PVPower"] = addVirtualAlias(entries[pid], "pv_power", "PV Power")
				// 	case "13119":
				// 		// TotalLoadActivePower
				// 		entries["LoadPower"] = addVirtualAlias(entries[pid], "load_power", "Load Power")
				//
				// 		// addVirtualAlias(entries[pid], "FOO", "FOO")
				//
				// 	case "13112":
				// 		// Daily PV Yield
				// 		entries["DailyPvEnergy"] = addVirtualAlias(entries["DailyPvEnergy"], "daily_pv_energy", "Daily PV Energy")
				// 	case "13174":
				// 		// DailyBatteryChargingEnergyFromPv
				// 		entries["YieldBatteryCharge"] = addVirtualAlias(entries[pid], "pv_battery_charge", "PV Battery Charge")
				// 	case "13029":
				// 		// DailyBatteryDischargingEnergy
				// 		entries["DailyBatteryDischargingEnergy"] = entries[pid]
				// 	case "13122":
				// 		// entries["DailyFeedInEnergy"] = addVirtualAlias(entries[pid], "pv_feed_in", "PV Feed In")
				// 		// @TODO - This may differ from DailyFeedInEnergyPv
				// 	case "13173":
				// 		// DailyFeedInEnergyPv
				// 		entries["YieldFeedIn"] = addVirtualAlias(entries[pid], "pv_feed_in", "PV Feed In")
				// 	case "13147":
				// 		// DailyPurchasedEnergy
				// 		entries["DailyPurchasedEnergy"] = addVirtualAlias(entries[pid], "daily_purchased_energy", "Daily Purchased Energy")
				//
				// 	case "13116":
				// 		// DailyLoadEnergyConsumptionFromPv
				// 		entries["YieldSelfConsumption"] = addVirtualAlias(entries[pid], "pv_self_consumption", "PV Self Consumption")
				// 	case "13134":
				// 		// TotalPvYield
				// 		entries["TotalPvYield"] = addVirtualAlias(entries[pid], "pv_total_yield", "PV Total Yield")
				//
				// 	case "13199":
				// 		// Daily Load Energy Consumption
				// 		entries["DailyTotalLoad"] = addVirtualAlias(entries[pid], "daily_total_energy", "Daily Total Energy")
				//
				// 	case "13130":
				// 		// Total Load Energy Consumption
				// 		entries["TotalEnergyConsumption"] = addVirtualAlias(entries[pid], "total_energy_consumption", "Total Energy Consumption"
				// }
			}
		}

		if len(entries.DataPoints) == 0 {
			break
		}

		// TotalDcPower
		entries.FromRefAddAlias("p13003", api.VirtualPsId, "power_pv", "")
		// BatteryChargingPower
		entries.FromRefAddAlias("p13126", api.VirtualPsId, "battery_charge_power", "")
		// BatteryDischargingPower
		entries.FromRefAddAlias("p13150", api.VirtualPsId, "battery_discharge_power", "")
		// TotalExportActivePower
		entries.FromRefAddAlias("p13121", api.VirtualPsId, "power_pv_to_grid", "")

		// TotalLoadActivePower
		entries.FromRefAddAlias("p13119", api.VirtualPsId, "power_load", "")

		// PurchasedPower
		entries.FromRefAddAlias("p13149", api.VirtualPsId, "power_grid_to_load", "")

		// Daily PV Yield
		entries.FromRefAddAlias("p13112", api.VirtualPsId, "daily_pv_energy", "")
		// DailyPvEnergy := entries.getFloatValue("DailyTotalLoad") - entries.getFloatValue("DailyPurchasedEnergy")
		// DailyBatteryChargingEnergyFromPv
		entries.FromRefAddAlias("p13174", api.VirtualPsId, "pv_battery_charge_energy", "")
		// DailyBatteryDischargingEnergy
		entries.FromRefAddAlias("p13029", api.VirtualPsId, "battery_discharge", "")

		// @TODO - This may differ from DailyFeedInEnergyPv
		// entries["DailyFeedInEnergy"] = entries.AddVirtualAliasFromRef("13122", "pv_feed_in", "PV Feed In")

		// DailyFeedInEnergyPv
		entries.FromRefAddAlias("p13173", api.VirtualPsId, "pv_feed_in", "")
		// DailyPurchasedEnergy
		entries.FromRefAddAlias("p13147", api.VirtualPsId, "daily_purchased_energy", "")
		// DailyLoadEnergyConsumptionFromPv
		entries.FromRefAddAlias("p13116", api.VirtualPsId, "pv_self_consumption", "")
		// TotalPvYield
		entries.FromRefAddAlias("p13134", api.VirtualPsId, "pv_total_yield", "")
		// Daily Load Energy Consumption
		entries.FromRefAddAlias("p13199", api.VirtualPsId, "daily_total_energy", "")
		// Total Load Energy Consumption
		entries.FromRefAddAlias("p13130", api.VirtualPsId, "total_energy_consumption", "")
		// entries.AddPointFromRef(api.Point{ Id:"queryDeviceList.p13130" }, api.Point{ PsKey:api.VirtualPsId, Id:"total_energy_consumption" })

		// entries.CopyEntry("p13130").CreateAlias()
		// 		entries.GetEntry(api.Point{PsKey:psId, Id:"total_income", Unit:p.TotalIncome.Unit, Type:api.PointTypeTotal}, now, p.TotalIncome.Value)

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

		// Add virtual entries.
		// ts := ret.Entries[0].Date
		// var value api.Float

		entries.FromRefAddFloat("pv_self_consumption",
			api.VirtualPsId,"pv_daily_yield", "",
			entries.GetFloatValue("pv_self_consumption", api.LastEntry) + entries.GetFloatValue("pv_battery_charge_energy", api.LastEntry) + entries.GetFloatValue("pv_feed_in", api.LastEntry))

		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId,"pv_self_consumption_percent", "",
			entries.GetPercent("pv_self_consumption", "daily_pv_energy", api.LastEntry))
		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId,"pv_battery_charge_percent", "",
			entries.GetPercent("pv_battery_charge_energy", "daily_pv_energy", api.LastEntry))
		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId,"pv_feed_in_percent", "",
			entries.GetPercent("pv_feed_in", "daily_pv_energy", api.LastEntry))

		// @TODO - Add this calculation.
		DailyPvEnergy := entries.GetFloatValue("daily_total_energy", api.LastEntry) - entries.GetFloatValue("daily_purchased_energy", api.LastEntry)
		// fmt.Sprintf("%f", DailyPvEnergy)
		entries.FromRefAddFloat("daily_total_energy",
			api.VirtualPsId,"daily_pv_energy_percent", "",
			api.GetPercent(DailyPvEnergy, entries.GetValue("daily_total_energy", api.LastEntry)))
		entries.FromRefAddFloat("daily_total_energy",
			api.VirtualPsId,"daily_purchased_energy_percent", "",
			entries.GetPercent("daily_purchased_energy", "daily_total_energy", api.LastEntry))

		entries.FromRefAddFloat("power_pv",
			api.VirtualPsId,"power_pv_to_load", "",
			entries.GetFloatValue("power_pv", api.LastEntry) - entries.GetFloatValue("battery_charge_power", api.LastEntry) - entries.GetFloatValue("power_pv_to_grid", api.LastEntry))

		// Battery
		entries.FromRefAddFloat("battery_charge_power",
			api.VirtualPsId,"power_battery", "",
			entries.LowerUpper("battery_discharge_power", "battery_charge_power", api.LastEntry))
		entries.FromRefAddFloat("battery_charge_power",
			api.VirtualPsId,"power_pv_to_battery", "",
			entries.GetFloatValue("battery_charge_power", api.LastEntry))
		entries.FromRefAddFloat("battery_discharge_power",
			api.VirtualPsId,"power_battery_to_load", "",
			entries.GetFloatValue("battery_charge_power", api.LastEntry))
		entries.FromRefAddFloat("battery_charge_power",
			api.VirtualPsId,"power_battery_to_grid", "",
			0.0)

		// Grid
		entries.FromRefAddFloat("power_grid_to_load",
			api.VirtualPsId,"power_grid", "",
			entries.LowerUpper("power_pv_to_grid", "power_grid_to_load", api.LastEntry))
		entries.FromRefAddFloat("power_grid_to_load",
			api.VirtualPsId,"power_grid_to_battery", "",
			0.0)


		entries.FromRefAddState("power_pv", api.VirtualPsId,"power_pv_active", "")
		entries.FromRefAddState("power_battery", api.VirtualPsId,"power_battery_active", "")
		entries.FromRefAddState("power_grid", api.VirtualPsId,"power_grid_active", "")
		entries.FromRefAddState("power_load", api.VirtualPsId,"power_load_active", "")

		entries.FromRefAddState("power_pv_to_battery", api.VirtualPsId,"power_pv_to_battery_active", "")
		entries.FromRefAddState("power_pv_to_load", api.VirtualPsId,"power_pv_to_load_active", "")
		entries.FromRefAddState("power_pv_to_grid", api.VirtualPsId,"power_pv_to_grid_active", "")

		entries.FromRefAddState("power_battery_to_load", api.VirtualPsId,"power_battery_to_load_active", "")
		entries.FromRefAddState("power_battery_to_grid", api.VirtualPsId,"power_battery_to_grid_active", "")

		entries.FromRefAddState("power_grid_to_load", api.VirtualPsId,"power_grid_to_load_active", "")
		entries.FromRefAddState("power_grid_to_battery", api.VirtualPsId,"power_grid_to_battery_active", "")


		entries.FromRefAddFloat("pv_battery_charge_energy",
			api.VirtualPsId, "battery_energy", "",
			entries.LowerUpper("pv_battery_charge_energy", "battery_discharge", api.LastEntry))

		entries.FromRefAddFloat("pv_feed_in",
			api.VirtualPsId,"grid_energy", "",
			entries.LowerUpper("pv_feed_in", "daily_purchased_energy", api.LastEntry))
	}

	return entries
}
