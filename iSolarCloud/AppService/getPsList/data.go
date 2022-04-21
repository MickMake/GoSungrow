package getPsList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"strconv"
	"time"
)

const Url = "/v1/powerStationService/getPsList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		AlarmCount                int64         `json:"alarm_count"`
		AlarmDevCount             int64         `json:"alarm_dev_count"`
		AreaID                    interface{}   `json:"area_id"`
		AreaType                  interface{}   `json:"area_type"`
		ArrearsStatus             int64         `json:"arrears_status"`
		BuildDate                 string        `json:"build_date"`
		BuildStatus               int64         `json:"build_status"`
		Co2Reduce                 api.UnitValue `json:"co2_reduce"`
		Co2ReduceTotal            api.UnitValue `json:"co2_reduce_total"`
		CurrPower                 api.UnitValue `json:"curr_power"`
		DailyIrradiation          api.UnitValue `json:"daily_irradiation"`
		DailyIrradiationVirgin    interface{}   `json:"daily_irradiation_virgin"`
		DesignCapacity            string        `json:"design_capacity"`
		DesignCapacityUnit        string        `json:"design_capacity_unit"`
		DesignCapacityVirgin      float64       `json:"design_capacity_virgin"`
		EquivalentHour            api.UnitValue `json:"equivalent_hour"`
		EsDisenergy               api.UnitValue `json:"es_disenergy"`
		EsEnergy                  api.UnitValue `json:"es_energy"`
		EsPower                   api.UnitValue `json:"es_power"`
		EsTotalDisenergy          api.UnitValue `json:"es_total_disenergy"`
		EsTotalEnergy             api.UnitValue `json:"es_total_energy"`
		ExpectInstallDate         string        `json:"expect_install_date"`
		FaultAlarmOfflineDevCount int64         `json:"fault_alarm_offline_dev_count"`
		FaultCount                int64         `json:"fault_count"`
		FaultDevCount             int64         `json:"fault_dev_count"`
		GcjLatitude               string        `json:"gcj_latitude"`
		GcjLongitude              string        `json:"gcj_longitude"`
		GprsLatitude              interface{}   `json:"gprs_latitude"`
		GprsLongitude             interface{}   `json:"gprs_longitude"`
		Images                    []struct {
			FileID      int64       `json:"file_id"`
			ID          int64       `json:"id"`
			PicLanguage int64       `json:"pic_language"`
			PicType     int64       `json:"pic_type"`
			PictureName string      `json:"picture_name"`
			PictureURL  string      `json:"picture_url"`
			PsID        int64       `json:"ps_id"`
			PsUnitUUID  interface{} `json:"ps_unit_uuid"`
		} `json:"images"`
		InstallDate            string        `json:"install_date"`
		InstalledPowerMap      api.UnitValue `json:"installed_power_map"`
		InstalledPowerVirgin   float64       `json:"installed_power_virgin"`
		InstallerAlarmCount    int64         `json:"installer_alarm_count"`
		InstallerFaultCount    int64         `json:"installer_fault_count"`
		InstallerPsFaultStatus int64         `json:"installer_ps_fault_status"`
		IsBankPs               int64         `json:"is_bank_ps"`
		IsTuv                  int64         `json:"is_tuv"`
		JoinYearInitElec       float64       `json:"join_year_init_elec"`
		Latitude               float64       `json:"latitude"`
		Location               string        `json:"location"`
		Longitude              float64       `json:"longitude"`
		MapLatitude            string        `json:"map_latitude"`
		MapLongitude           string        `json:"map_longitude"`
		MlpeFlag               int64         `json:"mlpe_flag"`
		Nmi                    interface{}   `json:"nmi"`
		OfflineDevCount        int64         `json:"offline_dev_count"`
		OperateYear            interface{}   `json:"operate_year"`
		OperationBusName       interface{}   `json:"operation_bus_name"`
		OwnerAlarmCount        int64         `json:"owner_alarm_count"`
		OwnerFaultCount        int64         `json:"owner_fault_count"`
		OwnerPsFaultStatus     int64         `json:"owner_ps_fault_status"`
		P83022y                string        `json:"p83022y"`
		P83046                 interface{}   `json:"p83046"`
		P83048                 interface{}   `json:"p83048"`
		P83049                 interface{}   `json:"p83049"`
		P83050                 interface{}   `json:"p83050"`
		P83051                 interface{}   `json:"p83051"`
		P83054                 interface{}   `json:"p83054"`
		P83055                 interface{}   `json:"p83055"`
		P83067                 interface{}   `json:"p83067"`
		P83070                 interface{}   `json:"p83070"`
		P83076                 float64       `json:"p83076"`
		P83077                 float64       `json:"p83077"`
		P83081                 float64       `json:"p83081"`
		P83089                 float64       `json:"p83089"`
		P83095                 float64       `json:"p83095"`
		P83118                 float64       `json:"p83118"`
		P83120                 float64       `json:"p83120"`
		P83127                 float64       `json:"p83127"`
		ParamCo2               float64       `json:"param_co2"`
		ParamCoal              float64       `json:"param_coal"`
		ParamIncome            float64       `json:"param_income"`
		ParamMeter             float64       `json:"param_meter"`
		ParamNox               float64       `json:"param_nox"`
		ParamPowder            float64       `json:"param_powder"`
		ParamSo2               float64       `json:"param_so2"`
		ParamTree              float64       `json:"param_tree"`
		ParamWater             float64       `json:"param_water"`
		PrScale                string        `json:"pr_scale"`
		Producer               interface{}   `json:"producer"`
		PsCountryID            int64         `json:"ps_country_id"`
		PsFaultStatus          int64         `json:"ps_fault_status"`
		PsHealthStatus         string        `json:"ps_health_status"`
		PsHolder               string        `json:"ps_holder"`
		PsID                   int64         `json:"ps_id"`
		PsIsNotInit            string        `json:"ps_is_not_init"`
		PsName                 string        `json:"ps_name"`
		PsShortName            string        `json:"ps_short_name"`
		PsStatus               int64         `json:"ps_status"`
		PsTimezone             string        `json:"ps_timezone"`
		PsType                 int64         `json:"ps_type"`
		PvEnergy               api.UnitValue `json:"pv_energy"`
		PvPower                api.UnitValue `json:"pv_power"`
		Radiation              api.UnitValue `json:"radiation"`
		RadiationVirgin        interface{}   `json:"radiation_virgin"`
		RecoreCreateTime       string        `json:"recore_create_time"`
		SafeStartDate          string        `json:"safe_start_date"`
		ShareType              string        `json:"share_type"`
		ShippingAddress        string        `json:"shipping_address"`
		ShippingZipCode        string        `json:"shipping_zip_code"`
		TodayEnergy            api.UnitValue `json:"today_energy"`
		TodayIncome            api.UnitValue `json:"today_income"`
		TotalCapcity           api.UnitValue `json:"total_capcity"`
		TotalEnergy            api.UnitValue `json:"total_energy"`
		TotalIncome            api.UnitValue `json:"total_income"`
		TotalInitCo2Accelerate float64       `json:"total_init_co2_accelerate"`
		TotalInitElec          float64       `json:"total_init_elec"`
		TotalInitProfit        float64       `json:"total_init_profit"`
		UseEnergy              api.UnitValue `json:"use_energy"`
		ValidFlag              int64         `json:"valid_flag"`
		WgsLatitude            float64       `json:"wgs_latitude"`
		WgsLongitude           float64       `json:"wgs_longitude"`
		ZipCode                string        `json:"zip_code"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
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

func (e *ResultData) GetPsId() int64 {
	var ret int64
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID != 0 {
				ret = p.PsID
				break
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsName() string {
	var ret string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID != 0 {
				ret = p.PsName
				break
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsSerial() string {
	var ret string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID != 0 {
				ret = p.PsShortName
				break
			}
		}
	}
	return ret
}

func (e *EndPoint) GetPsId() int64 {
	return e.Response.ResultData.GetPsId()
}

func (e *EndPoint) GetData() api.Data {
	return e.Response.ResultData.GetData()
}

func (e *ResultData) GetData() api.Data {
	var ret api.Data

	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}

		now := api.NewDateTime(time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds))

		for _, p := range e.PageList {
			psId := strconv.FormatInt(p.PsID, 10)

			ret.Entries = append(ret.Entries, add(now, psId, "p83077", "Pv Energy",					p.PvEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83089", "Es Discharge Energy",		p.EsDisenergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83095", "Es Total Discharge Energy",	p.EsTotalDisenergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83118", "Use Energy", 				p.UseEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83120", "Es Energy", 				p.EsEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83127", "Es Total Energy", 			p.EsTotalEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83076", "Pv Power", 					p.PvPower, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "p83081", "Es Power", 					p.EsPower, len(ret.Entries)))

			ret.Entries = append(ret.Entries, add(now, psId, "co2_reduce", "Co2 Reduce", 					p.Co2Reduce, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "co2_reduce_total", "Co2 Reduce Total", 		p.Co2ReduceTotal, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "curr_power", "Curr Power", 					p.CurrPower, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "daily_irradiation", "Daily Irradiation", 		p.DailyIrradiation, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "equivalent_hour", "Equivalent Hour", 			p.EquivalentHour, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "installed_power_map", "Installed Power Map",	p.InstalledPowerMap, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "radiation", "Radiation", 						p.Radiation, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "today_energy", "Today Energy", 				p.TodayEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "today_income", "Today Income", 				p.TodayIncome, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "total_capacity", "Total Capacity", 			p.TotalCapcity, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "total_energy", "Total Energy", 				p.TotalEnergy, len(ret.Entries)))
			ret.Entries = append(ret.Entries, add(now, psId, "total_income", "Total Income", 				p.TotalIncome, len(ret.Entries)))

			ret.Entries = append(ret.Entries, addValue(now, psId, "build_date", "Build Date",			p.BuildDate, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "build_status", "Build Status",	p.BuildStatus, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addFloatValue(now, psId, "latitude", "Latitude",			p.Latitude, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addFloatValue(now, psId, "longitude", "Longitude",		p.Longitude, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addValue(now, psId, "location", "Location",				p.Location, len(ret.Entries)))

			ret.Entries = append(ret.Entries, addIntValue(now, psId, "installer_ps_fault_status", "Installer PS Fault Status",	p.InstallerPsFaultStatus, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "owner_ps_fault_status", "Owner PS Fault Status",			p.OwnerPsFaultStatus, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "ps_fault_status", "PS Fault Status",						p.PsFaultStatus, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addValue(now, psId, "ps_health_status", "PS Health Status",						p.PsHealthStatus, len(ret.Entries)))

			ret.Entries = append(ret.Entries, addValue(now, psId, "ps_holder", "PS Holder",			p.PsHolder, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "ps_id", "PS Id",				p.PsID, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addValue(now, psId, "ps_name", "PS Name",				p.PsName, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addValue(now, psId, "ps_short_name", "PS Short Name",	p.PsShortName, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "ps_status", "PS Status",		p.PsStatus, len(ret.Entries)))
			ret.Entries = append(ret.Entries, addIntValue(now, psId, "ps_type", "PS Type",			p.PsType, len(ret.Entries)))

			// ret = append(ret, []string{now, "Pv Energy", p.PvEnergy.Value, p.PvEnergy.Unit})		// p83077
			// ret = append(ret, []string{now, "Es Discharge Energy", p.EsDisenergy.Value, p.EsDisenergy.Unit})	// p83089
			// ret = append(ret, []string{now, "Es Total Discharge Energy", p.EsTotalDisenergy.Value, p.EsTotalDisenergy.Unit})	// p83095
			// ret = append(ret, []string{now, "Use Energy", p.UseEnergy.Value, p.UseEnergy.Unit})		// p83118
			// ret = append(ret, []string{now, "Es Energy", p.EsEnergy.Value, p.EsEnergy.Unit})		// p83120
			// ret = append(ret, []string{now, "Es Total Energy", p.EsTotalEnergy.Value, p.EsTotalEnergy.Unit})	// p83127
			// ret = append(ret, []string{now, "Pv Power", p.PvPower.Value, p.PvPower.Unit})			// p83076
			// ret = append(ret, []string{now, "Es Power", p.EsPower.Value, p.EsPower.Unit})			// p83081
			//
			// ret = append(ret, []string{now, "Co2 Reduce", p.Co2Reduce.Value, p.Co2Reduce.Unit})
			// ret = append(ret, []string{now, "Co2 Reduce Total", p.Co2ReduceTotal.Value, p.Co2ReduceTotal.Unit})
			// ret = append(ret, []string{now, "Curr Power", p.CurrPower.Value, p.CurrPower.Unit})
			// ret = append(ret, []string{now, "Daily Irradiation", p.DailyIrradiation.Value, p.DailyIrradiation.Unit})
			// ret = append(ret, []string{now, "Equivalent Hour", p.EquivalentHour.Value, p.EquivalentHour.Unit})
			// ret = append(ret, []string{now, "Installed Power Map", p.InstalledPowerMap.Value, p.InstalledPowerMap.Unit})
			// ret = append(ret, []string{now, "Radiation", p.Radiation.Value, p.Radiation.Unit})
			// ret = append(ret, []string{now, "Today Energy", p.TodayEnergy.Value, p.TodayEnergy.Unit})
			// ret = append(ret, []string{now, "Today Income", p.TodayIncome.Value, p.TodayIncome.Unit})
			// ret = append(ret, []string{now, "Total Capacity", p.TotalCapcity.Value, p.TotalCapcity.Unit})
			// ret = append(ret, []string{now, "Total Energy", p.TotalEnergy.Value, p.TotalEnergy.Unit})
			// ret = append(ret, []string{now, "Total Income", p.TotalIncome.Value, p.TotalIncome.Unit})
		}
	}
	return ret
}

func addState(now api.DateTime, point string, name string, state bool, index int) api.DataEntry {
	return add(now, "virtual", point, name, api.UnitValue{ Value: fmt.Sprintf("%v", state) }, index)

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

func addValue(now api.DateTime, psId string, point string, name string, value string, index int) api.DataEntry {
	return add(now, psId, point, name, api.UnitValue{ Value: value }, index)

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

func addIntValue(now api.DateTime, psId string, point string, name string, value int64, index int) api.DataEntry {
	return add(now, psId, point, name, api.UnitValue{ Value: strconv.FormatInt(value, 10) }, index)
}

func addFloatValue(now api.DateTime, psId string, point string, name string, value float64, index int) api.DataEntry {
	return add(now, psId, point, name, api.UnitValue{ Value: api.Float64ToString(value) }, index)
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

// func (e *ResultData) GetData() [][]string {
// 	var ret [][]string
// 	for range Only.Once {
// 		i := len(e.PageList)
// 		if i == 0 {
// 			break
// 		}
// 		now := time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds)
// 		for _, p := range e.PageList {
// 			ret = append(ret, []string{now, "Co2 Reduce", p.Co2Reduce.Value, p.Co2Reduce.Unit})
// 			ret = append(ret, []string{now, "Co2 Reduce Total", p.Co2ReduceTotal.Value, p.Co2ReduceTotal.Unit})
// 			ret = append(ret, []string{now, "Curr Power", p.CurrPower.Value, p.CurrPower.Unit})
// 			ret = append(ret, []string{now, "Daily Irradiation", p.DailyIrradiation.Value, p.DailyIrradiation.Unit})
// 			ret = append(ret, []string{now, "Equivalent Hour", p.EquivalentHour.Value, p.EquivalentHour.Unit})
// 			ret = append(ret, []string{now, "Es Discharge Energy", p.EsDisenergy.Value, p.EsDisenergy.Unit})
// 			ret = append(ret, []string{now, "Es Energy", p.EsEnergy.Value, p.EsEnergy.Unit})
// 			ret = append(ret, []string{now, "Es Power", p.EsPower.Value, p.EsPower.Unit})
// 			ret = append(ret, []string{now, "Es Total Discharge Energy", p.EsTotalDisenergy.Value, p.EsTotalDisenergy.Unit})
// 			ret = append(ret, []string{now, "Es Total Energy", p.EsTotalEnergy.Value, p.EsTotalEnergy.Unit})
// 			ret = append(ret, []string{now, "Installed Power Map", p.InstalledPowerMap.Value, p.InstalledPowerMap.Unit})
// 			ret = append(ret, []string{now, "Pv Energy", p.PvEnergy.Value, p.PvEnergy.Unit})
// 			ret = append(ret, []string{now, "Pv Power", p.PvPower.Value, p.PvPower.Unit})
// 			ret = append(ret, []string{now, "Radiation", p.Radiation.Value, p.Radiation.Unit})
// 			ret = append(ret, []string{now, "Today Energy", p.TodayEnergy.Value, p.TodayEnergy.Unit})
// 			ret = append(ret, []string{now, "Today Income", p.TodayIncome.Value, p.TodayIncome.Unit})
// 			ret = append(ret, []string{now, "Total Capacity", p.TotalCapcity.Value, p.TotalCapcity.Unit})
// 			ret = append(ret, []string{now, "Total Energy", p.TotalEnergy.Value, p.TotalEnergy.Unit})
// 			ret = append(ret, []string{now, "Total Income", p.TotalIncome.Value, p.TotalIncome.Unit})
// 			ret = append(ret, []string{now, "Use Energy", p.UseEnergy.Value, p.UseEnergy.Unit})
// 		}
// 	}
// 	return ret
// }

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
			"Point Name",
			"Value",
			"Unit",
		)

		now := time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds)

		for _, p := range e.Response.ResultData.PageList {
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Co2 Reduce", p.Co2Reduce.Value, p.Co2Reduce.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Co2 Reduce Total", p.Co2ReduceTotal.Value, p.Co2ReduceTotal.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Curr Power", p.CurrPower.Value, p.CurrPower.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Daily Irradiation", p.DailyIrradiation.Value, p.DailyIrradiation.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Equivalent Hour", p.EquivalentHour.Value, p.EquivalentHour.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Es Discharge Energy", p.EsDisenergy.Value, p.EsDisenergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Es Energy", p.EsEnergy.Value, p.EsEnergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Es Power", p.EsPower.Value, p.EsPower.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Es Total Discharge Energy", p.EsTotalDisenergy.Value, p.EsTotalDisenergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Es Total Energy", p.EsTotalEnergy.Value, p.EsTotalEnergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Installed Power Map", p.InstalledPowerMap.Value, p.InstalledPowerMap.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Pv Energy", p.PvEnergy.Value, p.PvEnergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Pv Power", p.PvPower.Value, p.PvPower.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Radiation", p.Radiation.Value, p.Radiation.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Today Energy", p.TodayEnergy.Value, p.TodayEnergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Today Income", p.TodayIncome.Value, p.TodayIncome.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Total Capacity", p.TotalCapcity.Value, p.TotalCapcity.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Total Energy", p.TotalEnergy.Value, p.TotalEnergy.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Total Income", p.TotalIncome.Value, p.TotalIncome.Unit)
			_ = table.AddRow(now, fmt.Sprintf("%d.%d", p.PsID, 0), "Use Energy", p.UseEnergy.Value, p.UseEnergy.Unit)
		}

		table.InitGraph(output.GraphRequest {
			Title:        "",
			TimeColumn:   output.SetInteger(1),
			SearchColumn: output.SetInteger(2),
			NameColumn:   output.SetInteger(3),
			ValueColumn:  output.SetInteger(4),
			UnitsColumn:  output.SetInteger(5),
			SearchString: output.SetString(""),
			MinLeftAxis:  output.SetFloat(0),
			MaxLeftAxis:  output.SetFloat(0),
		})

	}
	return table
}
