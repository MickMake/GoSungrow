package getPsList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
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
		AlarmCount                int64         `json:"alarm_count" PointId:"alarm_count" PointType:"PointTypeBoot"`
		AlarmDevCount             int64         `json:"alarm_dev_count" PointId:"alarm_dev_count" PointType:"PointTypeBoot"`
		AreaID                    interface{}   `json:"area_id" PointId:"area_id" PointType:""`
		AreaType                  interface{}   `json:"area_type" PointId:"area_type" PointType:""`
		ArrearsStatus             int64         `json:"arrears_status" PointId:"arrears_status" PointType:""`
		BuildDate                 string        `json:"build_date" PointId:"build_date" PointType:""`
		BuildStatus               int64         `json:"build_status" PointId:"build_status" PointType:""`
		Co2Reduce                 api.UnitValue `json:"co2_reduce" PointId:"co2_reduce" PointType:""`
		Co2ReduceTotal            api.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointType:"PointTypeTotal"`
		CurrPower                 api.UnitValue `json:"curr_power" PointId:"curr_power" PointType:""`
		DailyIrradiation          api.UnitValue `json:"daily_irradiation" PointId:"daily_irradiation" PointType:"PointTypeDaily"`
		DailyIrradiationVirgin    interface{}   `json:"daily_irradiation_virgin" PointId:"daily_irradiation_virgin" PointType:"PointTypeDaily"`
		DesignCapacity            string        `json:"design_capacity" PointId:"design_capacity" PointType:""`
		DesignCapacityUnit        string        `json:"design_capacity_unit" PointId:"design_capacity_unit" PointType:""`
		DesignCapacityVirgin      float64       `json:"design_capacity_virgin" PointId:"design_capacity_virgin" PointType:""`
		EquivalentHour            api.UnitValue `json:"equivalent_hour" PointId:"equivalent_hour" PointType:"PointTypeDaily"`
		EsDisenergy               api.UnitValue `json:"es_discharge_energy" PointId:"es_discharge_energy" PointAlias:"p83089" PointType:""`
		EsEnergy                  api.UnitValue `json:"es_energy" PointId:"es_energy" PointAlias:"p83120" PointType:""`
		EsPower                   api.UnitValue `json:"es_power" PointId:"es_power" PointAlias:"p83081" PointType:""`
		EsTotalDisenergy          api.UnitValue `json:"es_total_discharge_energy" PointId:"es_total_discharge_energy" PointAlias:"p83095" PointType:"PointTypeTotal"`
		EsTotalEnergy             api.UnitValue `json:"es_total_energy" PointId:"es_total_energy" PointAlias:"p83127" PointType:"PointTypeTotal"`
		ExpectInstallDate         string        `json:"expect_install_date" PointId:"expect_install_date" PointType:""`
		FaultAlarmOfflineDevCount int64         `json:"fault_alarm_offline_dev_count" PointId:"fault_alarm_offline_dev_count" PointType:""`
		FaultCount                int64         `json:"fault_count" PointId:"fault_count" PointType:""`
		FaultDevCount             int64         `json:"fault_dev_count" PointId:"fault_dev_count" PointType:""`
		GcjLatitude               string        `json:"gcj_latitude" PointId:"gcj_latitude" PointType:""`
		GcjLongitude              string        `json:"gcj_longitude" PointId:"gcj_longitude" PointType:""`
		GprsLatitude              interface{}   `json:"gprs_latitude" PointId:"gprs_latitude" PointType:""`
		GprsLongitude             interface{}   `json:"gprs_longitude" PointId:"gprs_longitude" PointType:""`
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
		InstallDate            string        `json:"install_date" PointId:"install_date" PointType:""`
		InstalledPowerMap      api.UnitValue `json:"installed_power_map" PointId:"installed_power_map" PointType:""`
		InstalledPowerVirgin   float64       `json:"installed_power_virgin" PointId:"installed_power_virgin" PointType:""`
		InstallerAlarmCount    int64         `json:"installer_alarm_count" PointId:"installer_alarm_count" PointType:""`
		InstallerFaultCount    int64         `json:"installer_fault_count" PointId:"installer_fault_count" PointType:""`
		InstallerPsFaultStatus int64         `json:"installer_ps_fault_status" PointId:"installer_ps_fault_status" PointType:""`
		IsBankPs               int64         `json:"is_bank_ps" PointId:"is_bank_ps" PointType:""`
		IsTuv                  int64         `json:"is_tuv" PointId:"is_tuv" PointType:""`
		JoinYearInitElec       float64       `json:"join_year_init_elec" PointId:"join_year_init_elec" PointType:""`
		Latitude               float64       `json:"latitude" PointId:"latitude" PointType:""`
		Location               string        `json:"location" PointId:"location" PointType:""`
		Longitude              float64       `json:"longitude" PointId:"longitude" PointType:""`
		MapLatitude            string        `json:"map_latitude" PointId:"map_latitude" PointType:""`
		MapLongitude           string        `json:"map_longitude" PointId:"map_longitude" PointType:""`
		MlpeFlag               int64         `json:"mlpe_flag" PointId:"mlpe_flag" PointType:""`
		Nmi                    interface{}   `json:"nmi" PointId:"nmi" PointType:""`
		OfflineDevCount        int64         `json:"offline_dev_count" PointId:"offline_dev_count" PointType:""`
		OperateYear            interface{}   `json:"operate_year" PointId:"operate_year" PointType:""`
		OperationBusName       interface{}   `json:"operation_bus_name" PointId:"operation_bus_name" PointType:""`
		OwnerAlarmCount        int64         `json:"owner_alarm_count" PointId:"owner_alarm_count" PointType:""`
		OwnerFaultCount        int64         `json:"owner_fault_count" PointId:"owner_fault_count" PointType:""`
		OwnerPsFaultStatus     int64         `json:"owner_ps_fault_status" PointId:"owner_ps_fault_status" PointType:""`
		P83022y                string        `json:"p83022y" PointId:"P83022" PointType:""`
		P83046                 interface{}   `json:"p83046" PointId:"p83046" PointType:""`
		P83048                 interface{}   `json:"p83048" PointId:"p83048" PointType:""`
		P83049                 interface{}   `json:"p83049" PointId:"p83049" PointType:""`
		P83050                 interface{}   `json:"p83050" PointId:"p83050" PointType:""`
		P83051                 interface{}   `json:"p83051" PointId:"p83051" PointType:""`
		P83054                 interface{}   `json:"p83054" PointId:"p83054" PointType:""`
		P83055                 interface{}   `json:"p83055" PointId:"p83055" PointType:""`
		P83067                 interface{}   `json:"p83067" PointId:"p83067" PointType:""`
		P83070                 interface{}   `json:"p83070" PointId:"p83070" PointType:""`
		P83076                 float64       `json:"p83076" PointId:"p83076" PointType:""`		// Dupe of PvPower
		P83077                 float64       `json:"p83077" PointId:"p83077" PointType:""`
		P83081                 float64       `json:"p83081" PointId:"p83081" PointType:""`		// Dupe of EsPower
		P83089                 float64       `json:"p83089" PointId:"p83089" PointType:""`		// Dupe of EsDisenergy
		P83095                 float64       `json:"p83095" PointId:"p83095" PointType:""`		// Dupe of EsTotalDisenergy
		P83118                 float64       `json:"p83118" PointId:"p83118" PointType:""`		// Dupe of UseEnergy
		P83120                 float64       `json:"p83120" PointId:"p83120" PointType:""`		// Dupe of EsEnergy
		P83127                 float64       `json:"p83127" PointId:"p83127" PointType:""`		// Dupe of EsTotalEnergy
		ParamCo2               float64       `json:"param_co2" PointId:"param_co2" PointType:""`
		ParamCoal              float64       `json:"param_coal" PointId:"param_coal" PointType:""`
		ParamIncome            float64       `json:"param_income" PointId:"param_income" PointType:""`
		ParamMeter             float64       `json:"param_meter" PointId:"param_meter" PointType:""`
		ParamNox               float64       `json:"param_nox" PointId:"param_nox" PointType:""`
		ParamPowder            float64       `json:"param_powder" PointId:"param_powder" PointType:""`
		ParamSo2               float64       `json:"param_so2" PointId:"param_so2" PointType:""`
		ParamTree              float64       `json:"param_tree" PointId:"param_tree" PointType:""`
		ParamWater             float64       `json:"param_water" PointId:"param_water" PointType:""`
		PrScale                string        `json:"pr_scale" PointId:"pr_scale" PointType:""`
		Producer               interface{}   `json:"producer" PointId:"producer" PointType:""`
		PsCountryID            int64         `json:"ps_country_id" PointId:"ps_country_id" PointType:""`
		PsFaultStatus          int64         `json:"ps_fault_status" PointId:"ps_fault_status" PointType:""`
		PsHealthStatus         string        `json:"ps_health_status" PointId:"ps_health_status" PointType:""`
		PsHolder               string        `json:"ps_holder" PointId:"ps_holder" PointType:""`
		PsID                   int64         `json:"ps_id" PointId:"ps_id" PointType:""`
		PsIsNotInit            string        `json:"ps_is_not_init" PointId:"ps_is_not_init" PointType:""`
		PsName                 string        `json:"ps_name" PointId:"ps_name" PointType:""`
		PsShortName            string        `json:"ps_short_name" PointId:"ps_short_name" PointType:""`
		PsStatus               int64         `json:"ps_status" PointId:"ps_status" PointType:""`
		PsTimezone             string        `json:"ps_timezone" PointId:"ps_timezone" PointType:""`
		PsType                 int64         `json:"ps_type" PointId:"ps_type" PointType:""`
		PvEnergy               api.UnitValue `json:"pv_energy" PointId:"pv_energy" PointAlias:"p83077" PointType:""`
		PvPower                api.UnitValue `json:"pv_power" PointId:"pv_power" PointAlias:"p83076" PointType:""`
		Radiation              api.UnitValue `json:"radiation" PointId:"radiation" PointType:""`
		RadiationVirgin        interface{}   `json:"radiation_virgin" PointId:"radiation_virgin" PointType:""`
		RecoreCreateTime       string        `json:"recore_create_time" PointId:"recore_create_time" PointType:""`
		SafeStartDate          string        `json:"safe_start_date" PointId:"safe_start_date" PointType:""`
		ShareType              string        `json:"share_type" PointId:"share_type" PointType:""`
		ShippingAddress        string        `json:"shipping_address" PointId:"shipping_address" PointType:""`
		ShippingZipCode        string        `json:"shipping_zip_code" PointId:"shipping_zip_code" PointType:""`
		TodayEnergy            api.UnitValue `json:"today_energy" PointId:"today_energy" PointType:"PointTypeDaily"`
		TodayIncome            api.UnitValue `json:"today_income" PointId:"today_income" PointType:"PointTypeDaily"`
		TotalCapcity           api.UnitValue `json:"total_capacity" PointId:"total_capacity" PointType:"PointTypeTotal"`
		TotalEnergy            api.UnitValue `json:"total_energy" PointId:"total_energy" PointType:"PointTypeTotal"`
		TotalIncome            api.UnitValue `json:"total_income" PointId:"total_income" PointType:"PointTypeTotal"`
		TotalInitCo2Accelerate float64       `json:"total_init_co2_accelerate" PointId:"total_init_co2_accelerate" PointType:"PointTypeTotal"`
		TotalInitElec          float64       `json:"total_init_elec" PointId:"total_init_elec" PointType:"PointTypeTotal"`
		TotalInitProfit        float64       `json:"total_init_profit" PointId:"total_init_profit" PointType:"PointTypeTotal"`
		UseEnergy              api.UnitValue `json:"use_energy" PointId:"use_energy" PointAlias:"p83118" PointType:""`
		ValidFlag              int64         `json:"valid_flag" PointId:"valid_flag" PointType:""`
		WgsLatitude            float64       `json:"wgs_latitude" PointId:"wgs_latitude" PointType:""`
		WgsLongitude           float64       `json:"wgs_longitude" PointId:"wgs_longitude" PointType:""`
		ZipCode                string        `json:"zip_code" PointId:"zip_code" PointType:""`
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

func (e *EndPoint) GetData() api.DataMap {
	return e.Response.ResultData.GetData()
}

func (e *ResultData) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}

		// now := api.NewDateTime(time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds))

		for _, p := range e.PageList {
			// psId := strconv.FormatInt(p.PsID, 10)
			entries.StructToPoints("", p)

			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83077"}, now, p.PvEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83089"}, now, p.EsDisenergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83095"}, now, p.EsTotalDisenergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83118"}, now, p.UseEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83120"}, now, p.EsEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83127"}, now, p.EsTotalEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83076"}, now, p.PvPower.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"p83081"}, now, p.EsPower.Value)
			//
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"co2_reduce", Unit:p.Co2Reduce.Unit}, now, p.Co2Reduce.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"co2_reduce_total", Unit:p.Co2ReduceTotal.Unit}, now, p.Co2ReduceTotal.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"curr_power", Unit:p.CurrPower.Unit}, now, p.CurrPower.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"daily_irradiation", Unit:p.DailyIrradiation.Unit}, now, p.DailyIrradiation.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"equivalent_hour", Unit:p.EquivalentHour.Unit}, now, p.EquivalentHour.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"installed_power_map", Unit:p.InstalledPowerMap.Unit}, now, p.InstalledPowerMap.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"radiation", Unit:p.Radiation.Unit}, now, p.Radiation.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"today_energy", Unit:p.TodayEnergy.Unit}, now, p.TodayEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"today_income", Unit:p.TodayIncome.Unit}, now, p.TodayIncome.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"total_capacity", Unit:p.TotalCapcity.Unit}, now, p.TotalCapcity.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"total_energy", Unit:p.TotalEnergy.Unit}, now, p.TotalEnergy.Value)
			// entries.AddEntry(api.Point{PsKey: "virtual", Id:"total_income", Unit:p.TotalIncome.Unit}, now, p.TotalIncome.Value)
			//
			// entries.AddString(now, psId, "build_date", "", p.BuildDate)
			// entries.AddInt(now, psId, "build_status", "", p.BuildStatus)
			// entries.AddFloat(now, psId, "latitude", "", p.Latitude)
			// entries.AddFloat(now, psId, "longitude", "", p.Longitude)
			// entries.AddString(now, psId, "location", "", p.Location)
			//
			// entries.AddInt(now, psId, "installer_ps_fault_status", "", p.InstallerPsFaultStatus)
			// entries.AddInt(now, psId, "owner_ps_fault_status", "", p.OwnerPsFaultStatus)
			// entries.AddInt(now, psId, "ps_fault_status", "", p.PsFaultStatus)
			// entries.AddString(now, psId, "ps_health_status", "", p.PsHealthStatus)
			//
			// entries.AddString(now, psId, "ps_holder", "", p.PsHolder)
			// entries.AddInt(now, psId, "ps_id", "", p.PsID)
			// entries.AddString(now, psId, "ps_name", "", p.PsName)
			// entries.AddString(now, psId, "ps_short_name", "", p.PsShortName)
			// entries.AddInt(now, psId, "ps_status", "", p.PsStatus)
			// entries.AddInt(now, psId, "ps_type", "", p.PsType)

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
	return entries
}

// func add(now api.DateTime, psId string, pid string, name string, p api.UnitValue) api.DataEntry {
//
// 	vt := api.GetPoint(psId, pid)
// 	if !vt.Valid {
// 		vt = &api.Point {
// 			PsKey:       psId,
// 			Id:          pid,
// 			Description: name,
// 			Unit:        p.Unit,
// 			Type:        "PointTypeInstant",
// 		}
// 	}
// 	fv, _ := strconv.ParseFloat(p.Value, 64)
// 	return api.DataEntry {
// 		Date:           now,
// 		PointId:        api.NameDevicePoint(psId, pid),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          p.Value,
// 		ValueFloat:     fv,
// 		Unit:           p.Unit,
// 		ValueType:      vt,
// 		Index:          0,
// 	}
// }
//
// func addState(now api.DateTime, point string, name string, state bool) api.DataEntry {
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: fmt.Sprintf("%v", state) })
//
// 	// return api.DataEntry {
// 	// 	Date:           now,
// 	// 	PointId:        api.NameDevicePoint("virtual", point),
// 	// 	PointGroupName: "Virtual",
// 	// 	PointName:      name,
// 	// 	Value:          fmt.Sprintf("%v", state),
// 	// 	Unit:           "binary",
// 	// 	ValueType: &api.Point{
// 	// 		PsKey:       "virtual",
// 	// 		Id:          point,
// 	// 		Description: name,
// 	// 		Unit:        "binary",
// 	// 		Type:        "PointTypeInstant",
// 	// 	},
// 	// 	Index: index,
// 	// }
// }
//
// func addValue(now api.DateTime, psId string, point string, name string, value string) api.DataEntry {
// 	return add(now, psId, point, name, api.UnitValue{ Value: value })
// 	// vt := api.GetPoint(psId, point)
// 	// if !vt.Valid {
// 	// 	vt = &api.Point{
// 	// 		PsKey:       psId,
// 	// 		Id:          point,
// 	// 		Description: name,
// 	// 		Unit:        "",
// 	// 		Type:        "PointTypeInstant",
// 	// 	}
// 	// }
// 	// return api.DataEntry {
// 	// 	Date:           now,
// 	// 	PointId:        api.NameDevicePoint(psId, point),
// 	// 	PointGroupName: "Summary",
// 	// 	PointName:      name,
// 	// 	Value:          value,
// 	// 	Unit:           "",
// 	// 	ValueType:      vt,
// 	// 	Index:          index,
// 	// }
// }
//
// func addIntValue(now api.DateTime, psId string, point string, name string, value int64) api.DataEntry {
// 	return add(now, psId, point, name, api.UnitValue{ Value: strconv.FormatInt(value, 10) })
// }
//
// func addFloatValue(now api.DateTime, psId string, point string, name string, value float64) api.DataEntry {
// 	return add(now, psId, point, name, api.UnitValue{ Value: api.Float64ToString(value) })
// }
//
// func add2(now api.DateTime, psId string, point string, name string, value api.UnitValue) api.DataEntry {
// 	vt := api.GetPoint(psId, point)
// 	if !vt.Valid {
// 		vt = &api.Point{
// 			PsKey:       psId,
// 			Id:          point,
// 			Description: name,
// 			Unit:        value.Unit,
// 			Type:        "PointTypeInstant",
// 		}
// 	}
// 	return api.DataEntry {
// 		Date:           now,
// 		PointId:        api.NameDevicePoint(psId, point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          value.Value,
// 		Unit:           value.Unit,
// 		ValueType:      vt,
// 		Index:          0,
// 	}
// }
//
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
