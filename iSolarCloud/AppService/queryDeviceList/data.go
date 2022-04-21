package queryDeviceList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"strconv"
	"strings"
)

const Url = "/v1/devService/queryDeviceList"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
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
		FaultCount   int64 `json:"fault_count"`
		OfflineCount int64 `json:"offline_count"`
		RunCount     int64 `json:"run_count"`
		WarningCount int64 `json:"warning_count"`
	} `json:"dev_count_by_status_map"`
	DevCountByTypeMap struct {
		One4 int64 `json:"14"`
		Two2 int64 `json:"22"`
	} `json:"dev_count_by_type_map"`
	DevTypeDefinition struct {
		One    string `json:"1"`
		One0   string `json:"10"`
		One1   string `json:"11"`
		One2   string `json:"12"`
		One3   string `json:"13"`
		One4   string `json:"14"`
		One5   string `json:"15"`
		One6   string `json:"16"`
		One7   string `json:"17"`
		One8   string `json:"18"`
		One9   string `json:"19"`
		Two0   string `json:"20"`
		Two1   string `json:"21"`
		Two2   string `json:"22"`
		Two3   string `json:"23"`
		Two4   string `json:"24"`
		Two5   string `json:"25"`
		Two6   string `json:"26"`
		Two8   string `json:"28"`
		Two9   string `json:"29"`
		Three  string `json:"3"`
		Three0 string `json:"30"`
		Three1 string `json:"31"`
		Three2 string `json:"32"`
		Three3 string `json:"33"`
		Three4 string `json:"34"`
		Three5 string `json:"35"`
		Three6 string `json:"36"`
		Three7 string `json:"37"`
		Three8 string `json:"38"`
		Three9 string `json:"39"`
		Four   string `json:"4"`
		Four0  string `json:"40"`
		Four1  string `json:"41"`
		Four2  string `json:"42"`
		Four3  string `json:"43"`
		Four4  string `json:"44"`
		Four5  string `json:"45"`
		Four6  string `json:"46"`
		Four7  string `json:"47"`
		Four8  string `json:"48"`
		Five   string `json:"5"`
		Five0  string `json:"50"`
		Six    string `json:"6"`
		Seven  string `json:"7"`
		Eight  string `json:"8"`
		Nine   string `json:"9"`
		Nine9  string `json:"99"`
	} `json:"dev_type_definition"`
	PageList []struct {
		AlarmCount              int64       `json:"alarm_count"`
		ChnnlID                 int64       `json:"chnnl_id"`
		CommandStatus           int64       `json:"command_status"`
		ComponentAmount         int64       `json:"component_amount"`
		DataFlag                int64       `json:"data_flag"`
		DataFlagDetail          int64       `json:"data_flag_detail"`
		DeviceArea              string      `json:"device_area"`
		DeviceAreaName          string      `json:"device_area_name"`
		DeviceCode              int64       `json:"device_code"`
		DeviceID                int64       `json:"device_id"`
		DeviceModelCode         string      `json:"device_model_code"`
		DeviceModelID           string      `json:"device_model_id"`
		DeviceName              string      `json:"device_name"`
		DeviceStatus            int64       `json:"device_status"`
		DeviceType              int64       `json:"device_type"`
		FaultCount              int64       `json:"fault_count"`
		FaultStatus             string      `json:"fault_status"`
		FunctionEnum            string      `json:"function_enum"`
		InstallerAlarmCount     int64       `json:"installer_alarm_count"`
		InstallerDevFaultStatus int64       `json:"installer_dev_fault_status"`
		InstallerFaultCount     int64       `json:"installer_fault_count"`
		InverterModelType       int64       `json:"inverter_model_type"`
		IsDeveloper             string      `json:"is_developer"`
		IsG2point5Module        int64       `json:"is_g2point5_module"`
		IsInit                  int64       `json:"is_init"`
		IsSecond                int64       `json:"is_second"`
		IsSupportParamset       int64       `json:"is_support_paramset"`
		NodeTimestamps          interface{} `json:"node_timestamps"`
		OwnerAlarmCount         int64       `json:"owner_alarm_count"`
		OwnerDevFaultStatus     int64       `json:"owner_dev_fault_status"`
		OwnerFaultCount         int64       `json:"owner_fault_count"`
		PointData               PointData   `json:"point_data"`
		Points                  interface{} `json:"points"`
		PsTimezoneInfo          struct {
			IsDst    string `json:"is_dst"`
			TimeZone string `json:"time_zone"`
		} `json:"psTimezoneInfo"`
		PsID          int64       `json:"ps_id"`
		PsKey         string      `json:"ps_key"`
		RelState      int64       `json:"rel_state"`
		Sn            string      `json:"sn"`
		StringAmount  int64       `json:"string_amount"`
		TypeName      string      `json:"type_name"`
		UnitName      interface{} `json:"unit_name"`
		UUID          string      `json:"uuid"`
		UUIDIndexCode string      `json:"uuid_index_code"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
}

type PointData []struct {
	CodeID                 int64  `json:"code_id"`
	CodeIDOrderID          string `json:"code_id_order_id"`
	CodeName               string `json:"code_name"`
	DevPointLastUpdateTime string `json:"dev_point_last_update_time"`
	IsPlatformDefaultUnit  int64  `json:"is_platform_default_unit"`
	IsShow                 int64  `json:"is_show"`
	OrderID                int64  `json:"order_id"`
	OrderNum               int64  `json:"order_num"`
	PointGroupID           int64  `json:"point_group_id"`
	PointGroupIDOrderID    string `json:"point_group_id_order_id"`
	PointGroupName         string `json:"point_group_name"`
	PointID                int64  `json:"point_id"`
	PointName              string `json:"point_name"`
	PointSign              string `json:"point_sign"`
	Relate                 int64  `json:"relate"`
	TimeStamp              string `json:"time_stamp"`
	Unit                   string `json:"unit"`
	ValIsFixd              string `json:"val_is_fixd"`
	ValidSize              int64  `json:"valid_size"`
	Value                  string `json:"value"`
	ValueDescription       string `json:"value_description"`
}

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
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		_ = table.SetHeader(
			"Date",
			"Point Id",
			"Point Group Name",
			"Description",
			"Value",
			"Unit",
		)

		for _, d := range e.Response.ResultData.PageList {
			for _, p := range d.PointData {
				p.Value, p.Unit = api.DivideByThousandIfRequired(p.Value, p.Unit)
				// gp := api.GetPointInt("", p.PointID)
				// if gp != nil {
				// 	_ = table.AddRow(
				// 		api.NewDateTime(p.TimeStamp).PrintFull(),
				// 		api.NameDevicePointInt(d.PsKey, p.PointID),
				// 		p.PointGroupName,
				// 		p.PointName,
				// 		gp.Description,
				// 		p.Value,
				// 		p.Unit,
				// 		gp.Unit,
				// 	)
				// 	continue
				// }

				_ = table.AddRow(
					api.NewDateTime(p.TimeStamp).PrintFull(),
					api.NameDevicePointInt(d.PsKey, p.PointID),
					p.PointGroupName,
					p.PointName,
					p.Value,
					p.Unit,
				)
			}
		}

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

func (e *EndPoint) GetData() api.Data {
	var ret api.Data

	for range Only.Once {
		// Used for virtual entries.
		// 0 - sungrow_battery_charging_power
		var BatteryChargingPower float64
		// sensor.sungrow_battery_discharging_power
		var BatteryDischargingPower float64
		// 0 - sensor.sungrow_total_export_active_power
		var TotalExportActivePower float64
		// sensor.sungrow_purchased_power
		var PurchasedPower float64

		// 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		var DailyBatteryChargingEnergyFromPv float64
		// sensor.sungrow_daily_battery_discharging_energy
		var DailyBatteryDischargingEnergy float64
		// 0 - sensor.sungrow_daily_feed_in_energy_pv
		var DailyFeedInEnergyPv float64
		// sensor.sungrow_daily_purchased_energy
		var DailyPurchasedEnergy float64

		var TotalDcPower float64

		var TotalLoadActivePower float64

		// index := 0
		for _, d := range e.Response.ResultData.PageList {
			for _, p := range d.PointData {
				if p.Unit == "W" {
					fv, err := api.DivideByThousand(p.Value)
					// fv, err := strconv.ParseFloat(p.Value, 64)
					// fv = fv / 1000
					if err == nil {
						// p.Value = fmt.Sprintf("%.3f", fv)
						p.Value = fv
						p.Unit = "kW"
					}
				}

				if p.Unit == "Wh" {
					fv, err := api.DivideByThousand(p.Value)
					// fv, err := strconv.ParseFloat(p.Value, 64)
					// fv = fv / 1000
					if err == nil {
						// p.Value = fmt.Sprintf("%.3f", fv)
						p.Value = fv
						p.Unit = "kWh"
					}
				}

				vt := api.GetPointInt(d.PsKey, p.PointID)
				if !vt.Valid {
					vt = &api.Point{
						PsKey:       d.PsKey,
						Id:          "p" + strings.TrimPrefix(strconv.FormatInt(p.PointID, 10), "p"),
						Description: p.PointName,
						Unit:        p.Unit,
						Type:        "",
					}
				}
				ret.Entries = append(ret.Entries, api.DataEntry {
					Date:           api.NewDateTime(p.TimeStamp),
					PointId:        api.NameDevicePointInt(d.PsKey, p.PointID),
					PointGroupName: p.PointGroupName,
					PointName:      p.PointName,
					Value:          p.Value,
					Unit:           p.Unit,
					ValueType:      vt,
					Index:          len(ret.Entries),
				})

				// Handle virtual results.
				switch strings.ReplaceAll(p.PointName, " ", "") {
					case "BatteryChargingPower":
						BatteryChargingPower, _ = strconv.ParseFloat(p.Value, 64)
					case "BatteryDischargingPower":
						BatteryDischargingPower, _ = strconv.ParseFloat(p.Value, 64)
					case "TotalExportActivePower":
						TotalExportActivePower, _ = strconv.ParseFloat(p.Value, 64)
					case "PurchasedPower":
						PurchasedPower, _ = strconv.ParseFloat(p.Value, 64)
					case "TotalDCPower":
						TotalDcPower, _ = strconv.ParseFloat(p.Value, 64)
					case "TotalLoadActivePower":
						TotalLoadActivePower, _ = strconv.ParseFloat(p.Value, 64)

					case "DailyBatteryChargingEnergyFromPv":
						DailyBatteryChargingEnergyFromPv, _ = strconv.ParseFloat(p.Value, 64)
					case "DailyBatteryDischargingEnergy":
						DailyBatteryDischargingEnergy, _ = strconv.ParseFloat(p.Value, 64)
					case "DailyFeedInEnergyPv":
						DailyFeedInEnergyPv, _ = strconv.ParseFloat(p.Value, 64)
					case "DailyPurchasedEnergy":
						DailyPurchasedEnergy, _ = strconv.ParseFloat(p.Value, 64)
				}
			}
		}

		if len(ret.Entries) == 0 {
			break
		}

		// Add virtual entries.
		ts := ret.Entries[0].Date
		var value float64

		/*
			PVPower				- TotalDcPower
			PVPowerToBattery	- BatteryChargingPower
			PVPowerToLoad		- TotalDcPower - BatteryChargingPower - TotalExportActivePower
			PVPowerToGrid		- TotalExportActivePower

			LoadPower			- TotalLoadActivePower
			BatteryToLoad		- BatteryDischargingPower
			BatteryToGrid		- ?

			GridPower			- TotalDcPower
			GridToLoad			- PurchasedPower
			GridToBattery		- ?
		*/


		PVPower				:= TotalDcPower
		PVPowerToBattery	:= BatteryChargingPower
		PVPowerToLoad		:= TotalDcPower - BatteryChargingPower - TotalExportActivePower
		PVPowerToGrid		:= TotalExportActivePower
		ret.Entries = append(ret.Entries, addState(ts, "pv_power_active", "PV Power Active", isActive(PVPower), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power", "PV Power", PVPower, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "pv_power_to_battery_active", "PV Power To Battery Active", isActive(PVPowerToBattery), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power_to_battery", "PV Power To Battery", PVPowerToBattery, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "pv_power_to_load_active", "PV Power To Load Active", isActive(PVPowerToLoad), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power_to_load", "PV Power To Load", PVPowerToLoad, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "pv_power_to_grid_active", "PV Power To Grid Active", isActive(PVPowerToGrid), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power_to_grid", "PV Power To Grid", PVPowerToGrid, "kW", len(ret.Entries)))


		BatteryPower		:= lowerUpper(BatteryChargingPower, BatteryDischargingPower)
		BatteryToLoad		:= BatteryDischargingPower
		BatteryToGrid		:= 0.0
		ret.Entries = append(ret.Entries, addState(ts, "battery_power_active", "Battery Power Active", isActive(BatteryPower), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "battery_power", "Battery Power", BatteryPower, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "battery_power_to_load_active", "Battery Power To Load Active", isActive(BatteryToLoad), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "battery_power_to_load", "Battery Power To Load", BatteryToLoad, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "battery_power_to_grid_active", "Battery Power To Grid Active", isActive(BatteryToGrid), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "battery_power_to_grid", "Battery Power To Grid", BatteryToGrid, "kW", len(ret.Entries)))


		GridPower			:= lowerUpper(TotalExportActivePower, PurchasedPower)
		GridToLoad			:= PurchasedPower
		GridToBattery		:= 0.0
		ret.Entries = append(ret.Entries, addState(ts, "grid_power_active", "Grid Power Active", isActive(GridPower), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "grid_power", "Grid Power", GridPower, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "grid_power_to_load_active", "Grid Power To Load Active", isActive(GridToLoad), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "grid_power_to_load", "Grid Power To Load", GridToLoad, "kW", len(ret.Entries)))

		ret.Entries = append(ret.Entries, addState(ts, "grid_power_to_battery_active", "Grid Power To Battery Active", isActive(GridToBattery), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "grid_power_to_battery", "Grid Power To Battery", GridToBattery, "kW", len(ret.Entries)))


		LoadPower			:= TotalLoadActivePower
		ret.Entries = append(ret.Entries, addState(ts, "load_power_active", "Load Power Active", isActive(LoadPower), len(ret.Entries)))
		ret.Entries = append(ret.Entries, addFloatValue(ts, "load_power", "Load Power", LoadPower, "kW", len(ret.Entries)))


		// {
		// 	if isActive(BatteryChargingPower) {
		// 		value = 0 - BatteryChargingPower
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_power_active", "Battery Power Active", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_battery", "PV Flow To Battery", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_flow_to_load", "Battery Flow To Load", false, len(ret.Entries)))
		//
		// 	} else if isActive(BatteryDischargingPower) {
		// 		value = BatteryDischargingPower
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_power_active", "Battery Power Active", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_battery", "PV Flow To Battery", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_flow_to_load", "Battery Flow To Load", true, len(ret.Entries)))
		//
		// 	} else {
		// 		value = 0
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_power_active", "Battery Power Active", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_battery", "PV Flow To Battery", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "battery_flow_to_load", "Battery Flow To Load", false, len(ret.Entries)))
		// 	}
		//
		// 	ret.Entries = append(ret.Entries, addFloatValue(ts, "battery_power", "Battery Power", value, len(ret.Entries)))
		// }
		//
		// {
		// 	if isActive(TotalExportActivePower) {
		// 		value = 0 - TotalExportActivePower
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_power_active", "Grid Power Active", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_grid", "PV Flow To Grid", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_flow_to_load", "Grid Flow To Load", false, len(ret.Entries)))
		//
		// 	} else if isActive(PurchasedPower) {
		// 		value = PurchasedPower
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_power_active", "Grid Power Active", true, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_grid", "PV Flow To Grid", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_flow_to_load", "Grid Flow To Load", true, len(ret.Entries)))
		//
		// 	} else {
		// 		value = 0
		//
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_power_active", "Grid Power Active", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "pv_flow_to_grid", "PV Flow To Grid", false, len(ret.Entries)))
		// 		ret.Entries = append(ret.Entries, addState(ts, "grid_flow_to_load", "Grid Flow To Load", false, len(ret.Entries)))
		// 	}
		//
		// 	ret.Entries = append(ret.Entries, addFloatValue(ts, "grid_power", "Grid Power", value, len(ret.Entries)))
		// }
		//
		// {
		// 	ret.Entries = append(ret.Entries, addState(ts, "pv_power_active", "PV Power Active", isActive(TotalDcPower), len(ret.Entries)))
		// 	ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power", "PV Power", TotalDcPower, len(ret.Entries)))
		//
		// 	value = TotalDcPower - BatteryChargingPower - TotalExportActivePower
		// 	ret.Entries = append(ret.Entries, addFloatValue(ts, "pv_power_to_load", "PV Power To Load", value, len(ret.Entries)))
		// }
		//
		// {
		// 	ret.Entries = append(ret.Entries, addState(ts, "load_power_active", "Load Power Active", isActive(TotalLoadActivePower), len(ret.Entries)))
		//
		// 	ret.Entries = append(ret.Entries, api.DataEntry{
		// 		Date:           ts,
		// 		PointId:        "virtual.pv_power_to_grid",
		// 		PointGroupName: "Virtual",
		// 		PointName:      "PV Power To Grid",
		// 		Value:          api.Float64ToString(TotalExportActivePower),
		// 		Unit:           "kW",
		// 		ValueType: &api.Point{
		// 			PsKey:       "virtual",
		// 			Id:          "pv_power_to_grid",
		// 			Description: "PV Power To Grid",
		// 			Unit:        "kW",
		// 			Type:        "PointTypeInstant",
		// 		},
		// 		Index: len(ret.Entries),
		// 	})
		// }

		{
			if DailyBatteryChargingEnergyFromPv > 0 {
				value = 0 - DailyBatteryChargingEnergyFromPv
			} else {
				value = DailyBatteryDischargingEnergy
			}
			ret.Entries = append(ret.Entries, api.DataEntry{
				Date:           ts,
				PointId:        "virtual.battery_energy",
				PointGroupName: "Virtual",
				PointName:      "Battery Energy",
				Value:          api.Float64ToString(value),
				Unit:           "kWh",
				ValueType: &api.Point{
					PsKey:       "virtual",
					Id:          "battery_energy",
					Description: "Battery Energy",
					Unit:        "kWh",
					Type:        "PointTypeInstant",
				},
				Index: len(ret.Entries),
			})
		}

		{
			if DailyFeedInEnergyPv > 0 {
				value = 0 - DailyFeedInEnergyPv
			} else {
				value = DailyPurchasedEnergy
			}
			ret.Entries = append(ret.Entries, api.DataEntry{
				Date:           ts,
				PointId:        "virtual.grid_energy",
				PointGroupName: "Virtual",
				PointName:      "Grid Energy",
				Value:          api.Float64ToString(value),
				Unit:           "kWh",
				ValueType: &api.Point{
					PsKey:       "virtual",
					Id:          "grid_energy",
					Description: "Grid Energy",
					Unit:        "kWh",
					Type:        "PointTypeInstant",
				},
				Index: len(ret.Entries),
			})
		}

	}

	return ret
}

func lowerUpper(lower float64, upper float64) float64 {
	if lower > 0 {
		return 0 - lower
	}
	return upper
}


func addState(now api.DateTime, point string, name string, state bool, index int) api.DataEntry {
	return add(now, "virtual", point, name, api.UnitValue{ Value: fmt.Sprintf("%v", state), Unit: "binary"}, index)

	// return api.DataEntry {
	// 	Date:           now,
	// 	PointId:        api.NameDevicePoint("virtual", point),
	// 	PointGroupName: "Virtual",
	// 	PointName:      name,
	// 	Value:          fmt.Sprintf("%v", state),
	// 	Unit:           "binary",
	// 	ValueType: &api.Point{
	// 		PsKey:       "virtual",
	// 		Id:          point,
	// 		Description: name,
	// 		Unit:        "binary",
	// 		Type:        "PointTypeInstant",
	// 	},
	// 	Index: index,
	// }
}

func addValue(now api.DateTime, point string, name string, value string, unit string, index int) api.DataEntry {
	return add(now, "virtual", point, name, api.UnitValue{ Value: value, Unit: unit}, index)

	// vt := api.GetPoint(psId, point)
	// if !vt.Valid {
	// 	vt = &api.Point{
	// 		PsKey:       psId,
	// 		Id:          point,
	// 		Description: name,
	// 		Unit:        "",
	// 		Type:        "PointTypeInstant",
	// 	}
	// }
	// return api.DataEntry {
	// 	Date:           now,
	// 	PointId:        api.NameDevicePoint(psId, point),
	// 	PointGroupName: "Summary",
	// 	PointName:      name,
	// 	Value:          value,
	// 	Unit:           "",
	// 	ValueType:      vt,
	// 	Index:          index,
	// }
}

func addIntValue(now api.DateTime, point string, name string, value int64, unit string, index int) api.DataEntry {
	return add(now, "virtual", point, name, api.UnitValue{ Value: strconv.FormatInt(value, 10), Unit: unit }, index)
}

func addFloatValue(now api.DateTime, point string, name string, value float64, unit string, index int) api.DataEntry {
	return add(now, "virtual", point, name, api.UnitValue{ Value: api.Float64ToString(value), Unit: unit }, index)
}

func add(now api.DateTime, psId string, point string, name string, value api.UnitValue, index int) api.DataEntry {
	vt := api.GetPoint(psId, point)
	if !vt.Valid {
		vt = &api.Point{
			PsKey:       psId,
			Id:          point,
			Description: name,
			Unit:        value.Unit,
			Type:        "PointTypeInstant",
		}
	}
	return api.DataEntry {
		Date:           now,
		PointId:        api.NameDevicePoint(psId, point),
		PointGroupName: "Virtual",
		PointName:      name,
		Value:          value.Value,
		Unit:           value.Unit,
		ValueType:      vt,
		Index:          index,
	}
}

// func addState(now api.DateTime, point string, name string, state bool, index int) api.DataEntry {
// 	return api.DataEntry {
// 		Date:           now,
// 		PointId:        api.NameDevicePoint("virtual", point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          fmt.Sprintf("%v", state),
// 		Unit:           "binary",
// 		ValueType: &api.Point{
// 			PsKey:       "virtual",
// 			Id:          point,
// 			Description: name,
// 			Unit:        "binary",
// 			Type:        "PointTypeInstant",
// 		},
// 		Index: index,
// 	}
// }

func isActive(value float64) bool {
	if (value > 0.01) || (value < -0.01) {
		return true
	}
	return false
}