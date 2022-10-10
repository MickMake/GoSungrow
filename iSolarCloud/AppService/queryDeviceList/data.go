package queryDeviceList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
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
		FaultCount   api.Integer `json:"fault_count" PointId:"fault_count" PointType:""`
		OfflineCount api.Integer `json:"offline_count" PointId:"offline_count" PointType:""`
		RunCount     api.Integer `json:"run_count" PointId:"run_count" PointType:""`
		WarningCount api.Integer `json:"warning_count" PointId:"warning_count" PointType:""`
	} `json:"dev_count_by_status_map"`
	DevCountByTypeMap map[string]api.Integer `json:"dev_count_by_type_map"`
	// DevCountByTypeMap struct {
	// 	One4 api.Integer `json:"14"`
	// 	Two2 api.Integer `json:"22"`
	// } `json:"dev_count_by_type_map"`
	DevTypeDefinition map[string]string `json:"dev_type_definition"`
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
		AlarmCount              api.Integer `json:"alarm_count" PointId:"alarm_count" PointType:""`
		ChannelId               api.Integer `json:"chnnl_id" PointId:"channel_id" PointType:""`
		CommandStatus           api.Integer `json:"command_status" PointId:"command_status" PointType:""`
		ComponentAmount         api.Integer `json:"component_amount" PointId:"component_amount" PointType:""`
		DataFlag                api.Integer `json:"data_flag" PointId:"data_flag" PointType:""`
		DataFlagDetail          api.Integer `json:"data_flag_detail" PointId:"data_flag_detail" PointType:""`
		DeviceArea              string      `json:"device_area" PointId:"device_area" PointType:""`
		DeviceAreaName          string      `json:"device_area_name" PointId:"device_area_name" PointType:""`
		DeviceCode              api.Integer `json:"device_code" PointId:"device_code" PointType:""`
		DeviceID                api.Integer `json:"device_id" PointId:"device_id" PointType:""`
		DeviceModelCode         string      `json:"device_model_code" PointId:"device_model_code" PointType:""`
		DeviceModelID           string      `json:"device_model_id" PointId:"device_model_id" PointType:""`
		DeviceName              string      `json:"device_name" PointId:"device_name" PointType:""`
		DeviceStatus            api.Integer `json:"device_status" PointId:"device_status" PointType:""`
		DeviceType              api.Integer `json:"device_type" PointId:"device_type" PointType:""`
		FaultCount              api.Integer `json:"fault_count" PointId:"fault_count" PointType:""`
		FaultStatus             string      `json:"fault_status" PointId:"fault_status" PointType:""`
		FunctionEnum            string      `json:"function_enum" PointId:"function_enum" PointType:""`
		InstallerAlarmCount     api.Integer `json:"installer_alarm_count" PointId:"installer_alarm_count" PointType:""`
		InstallerDevFaultStatus api.Integer `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointType:""`
		InstallerFaultCount     api.Integer `json:"installer_fault_count" PointId:"installer_fault_count" PointType:""`
		InverterModelType       api.Integer `json:"inverter_model_type" PointId:"inverter_model_type" PointType:""`
		IsDeveloper             api.Bool    `json:"is_developer" PointId:"is_developer" PointType:""`
		IsG2point5Module        api.Bool    `json:"is_g2point5_module" PointId:"is_g2point5_module" PointType:""`
		IsInit                  api.Bool    `json:"is_init" PointId:"is_init" PointType:""`
		IsSecond                api.Bool    `json:"is_second" PointId:"is_second" PointType:""`
		IsSupportParamset       api.Bool    `json:"is_support_paramset" PointId:"is_support_paramset" PointType:""`
		NodeTimestamps          interface{} `json:"node_timestamps" PointId:"node_timestamps" PointType:""`
		OwnerAlarmCount         api.Integer `json:"owner_alarm_count" PointId:"owner_alarm_count" PointType:""`
		OwnerDevFaultStatus     api.Integer `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointType:""`
		OwnerFaultCount         api.Integer `json:"owner_fault_count" PointId:"owner_fault_count" PointType:""`
		PointData               PointData   `json:"point_data"`
		Points                  interface{} `json:"points" PointId:"points" PointType:""`
		PsTimezoneInfo          struct {
			IsDst    api.Bool `json:"is_dst"`
			TimeZone string   `json:"time_zone"`
		} `json:"psTimezoneInfo"`
		PsID                    api.Integer `json:"ps_id" PointId:"ps_id" PointType:""`
		PsKey                   api.PsKey   `json:"ps_key" PointId:"ps_key" PointType:""`
		RelState                api.Integer `json:"rel_state" PointId:"rel_state" PointType:""`
		Sn                      string      `json:"sn" PointId:"sn" PointType:""`
		StringAmount            api.Integer `json:"string_amount" PointId:"string_amount" PointType:""`
		TypeName                string      `json:"type_name" PointId:"type_name" PointType:""`
		UnitName                interface{} `json:"unit_name" PointId:"unit_name" PointType:""`
		UUID                    string      `json:"uuid" PointId:"uuid" PointType:""`
		UUIDIndexCode           string      `json:"uuid_index_code" PointId:"uuid_index_code" PointType:""`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
}

type PointData []PointStruct
type PointStruct struct {
	CodeID                 api.Integer `json:"code_id"`
	CodeIDOrderID          string      `json:"code_id_order_id"`
	CodeName               string      `json:"code_name"`
	DevPointLastUpdateTime string      `json:"dev_point_last_update_time"`
	IsPlatformDefaultUnit  api.Bool    `json:"is_platform_default_unit"`
	IsShow                 api.Bool    `json:"is_show"`
	OrderID                api.Integer `json:"order_id"`
	OrderNum               api.Integer `json:"order_num"`
	PointGroupID           api.Integer `json:"point_group_id"`
	PointGroupIDOrderID    string      `json:"point_group_id_order_id"`
	PointGroupName         string      `json:"point_group_name"`
	PointID                api.Integer `json:"point_id"`
	PointName              string      `json:"point_name"`
	PointSign              string      `json:"point_sign"`
	Relate                 api.Integer `json:"relate"`
	TimeStamp              string      `json:"time_stamp"`
	Unit                   string      `json:"unit"`
	ValIsFixd              string      `json:"val_is_fixd"`
	ValidSize              api.Integer `json:"valid_size"`
	Value                  api.Float   `json:"value"`
	ValueDescription       string      `json:"value_description"`
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
			if p.DeviceName != name {
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

		for _, p := range data.Order {
			entries := data.DataPoints[p]
			for _, de := range entries {
				if (de.Parent.Key == "virtual1") ||
					(de.Point.GroupName == "virtual1") ||
					(de.Parent.PsId == "virtual1") ||
					(de.Point.Unit == "binary") {
					fmt.Sprintf("")
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

// type DataEntry api.DataEntry
// type EntryMap api.DataMap

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

		entries.StructToPoints(e.Response.ResultData.DevCountByStatusMap, "queryDeviceList." + e.Request.PsId.String() + ".status", e.Request.PsId.String(), time.Time{})

		for _, d := range e.Response.ResultData.PageList {
			name := fmt.Sprintf("queryDeviceList.%s", d.PsKey)
			entries.StructToPoints(d, name, d.PsKey.Value(), time.Time{})

			for _, p := range d.PointData {
				pid := api.SetPointInt(p.PointID.Value())
				uv := api.SetUnitValueFloat(p.Value.Value(), p.Unit)
				// name2 := fmt.Sprintf("%s.PointData.%s", name, pid)
				name2 := fmt.Sprintf("%s.PointData", name)
				entries.AddUnitValue(name2, d.PsKey.Value(), pid, p.PointName, api.NewDateTime(p.TimeStamp), uv)

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
		entries.FromRefAddAlias("p13003", api.VirtualPsId, "pv_power", "")
		// BatteryChargingPower
		entries.FromRefAddAlias("p13126", api.VirtualPsId, "pv_power_to_battery", "")
		// BatteryDischargingPower
		entries.FromRefAddAlias("p13150", api.VirtualPsId, "battery_power_to_load", "")
		// TotalExportActivePower
		entries.FromRefAddAlias("p13121", api.VirtualPsId, "pv_power_to_grid", "")

		// TotalLoadActivePower
		entries.FromRefAddAlias("p13119", api.VirtualPsId, "load_power", "")

		// PurchasedPower
		entries.FromRefAddAlias("p13149", api.VirtualPsId, "grid_power_to_load", "")

		// Daily PV Yield
		entries.FromRefAddAlias("p13112", api.VirtualPsId, "daily_pv_energy", "")
		// DailyPvEnergy := entries.getFloatValue("DailyTotalLoad") - entries.getFloatValue("DailyPurchasedEnergy")
		// DailyBatteryChargingEnergyFromPv
		entries.FromRefAddAlias("p13174", api.VirtualPsId, "pv_battery_charge", "")
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
			entries.GetFloatValue("pv_self_consumption", api.LastEntry) + entries.GetFloatValue("pv_battery_charge", api.LastEntry) + entries.GetFloatValue("pv_feed_in", api.LastEntry))

		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId,"pv_self_consumption_percent", "",
			entries.GetPercent("pv_self_consumption", "daily_pv_energy", api.LastEntry))
		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId,"pv_battery_charge_percent", "",
			entries.GetPercent("pv_battery_charge", "daily_pv_energy", api.LastEntry))
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

		entries.FromRefAddFloat("pv_power",
			api.VirtualPsId,"pv_power_to_load", "",
			entries.GetFloatValue("pv_power", api.LastEntry) - entries.GetFloatValue("pv_power_to_battery", api.LastEntry) - entries.GetFloatValue("pv_power_to_grid", api.LastEntry))

		// Battery
		entries.FromRefAddFloat("pv_power_to_battery",
			api.VirtualPsId,"battery_power", "",
			entries.LowerUpper("pv_power_to_battery", "battery_power_to_load", api.LastEntry))
		entries.FromRefAddFloat("pv_power_to_battery",
			api.VirtualPsId,"battery_power_to_grid", "",
			0.0)

		// Grid
		entries.FromRefAddFloat("grid_power_to_load",
			api.VirtualPsId,"grid_power", "",
			entries.LowerUpper("pv_power_to_grid", "grid_power_to_load", api.LastEntry))
		entries.FromRefAddFloat("grid_power_to_load",
			api.VirtualPsId,"grid_power_to_battery", "",
			0.0)


		entries.FromRefAddState("pv_power", api.VirtualPsId,"pv_power_active", "")
		entries.FromRefAddState("battery_power", api.VirtualPsId,"battery_power_active", "")
		entries.FromRefAddState("grid_power", api.VirtualPsId,"grid_power_active", "")
		entries.FromRefAddState("load_power", api.VirtualPsId,"load_power_active", "")

		entries.FromRefAddState("pv_power_to_battery", api.VirtualPsId,"pv_power_to_battery_active", "")
		entries.FromRefAddState("pv_power_to_load", api.VirtualPsId,"pv_power_to_load_active", "")
		entries.FromRefAddState("pv_power_to_grid", api.VirtualPsId,"pv_power_to_grid_active", "")

		entries.FromRefAddState("battery_power_to_load", api.VirtualPsId,"battery_power_to_load_active", "")
		entries.FromRefAddState("battery_power_to_grid", api.VirtualPsId,"battery_power_to_grid_active", "")

		entries.FromRefAddState("grid_power_to_load", api.VirtualPsId,"grid_power_to_load_active", "")
		entries.FromRefAddState("grid_power_to_battery", api.VirtualPsId,"grid_power_to_battery_active", "")


		entries.FromRefAddFloat("pv_battery_charge",
			api.VirtualPsId, "battery_energy", "",
			entries.LowerUpper("pv_battery_charge", "battery_discharge", api.LastEntry))

		entries.FromRefAddFloat("pv_feed_in",
			api.VirtualPsId,"grid_energy", "",
			entries.LowerUpper("pv_feed_in", "daily_purchased_energy", api.LastEntry))


		// for _, pid := range entries.Order {
		// 	// entries[pid].Index = i
		// 	ret.Entries = append(ret.Entries, entries.Entries[pid])
		// }
	}

	return entries
}
