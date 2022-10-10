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
		AlarmCount                api.Integer   `json:"alarm_count" PointId:"alarm_count" PointType:"PointTypeBoot"`
		AlarmDevCount             api.Integer   `json:"alarm_dev_count" PointId:"alarm_dev_count" PointType:"PointTypeBoot"`
		AreaID                    interface{}   `json:"area_id" PointId:"area_id" PointType:""`
		AreaType                  interface{}   `json:"area_type" PointId:"area_type" PointType:""`
		ArrearsStatus             api.Integer   `json:"arrears_status" PointId:"arrears_status" PointType:""`
		BuildDate                 api.DateTime        `json:"build_date" PointId:"build_date" PointType:""`
		BuildStatus               api.Integer   `json:"build_status" PointId:"build_status" PointType:""`
		Co2Reduce                 api.UnitValue `json:"co2_reduce" PointId:"co2_reduce" PointType:""`
		Co2ReduceTotal            api.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointType:"PointTypeTotal"`
		CurrPower                 api.UnitValue `json:"curr_power" PointId:"curr_power" PointType:""`
		DailyIrradiation          api.UnitValue `json:"daily_irradiation" PointId:"daily_irradiation" PointType:"PointTypeDaily"`
		DailyIrradiationVirgin    api.Float   `json:"daily_irradiation_virgin" PointId:"daily_irradiation_virgin" PointType:"PointTypeDaily"`
		DesignCapacity            api.Float        `json:"design_capacity" PointId:"design_capacity" PointType:""`
		DesignCapacityUnit        api.String        `json:"design_capacity_unit" PointId:"design_capacity_unit" PointType:""`
		DesignCapacityVirgin      api.Float     `json:"design_capacity_virgin" PointId:"design_capacity_virgin" PointType:""`
		EquivalentHour            api.UnitValue `json:"equivalent_hour" PointId:"equivalent_hour" PointType:"PointTypeDaily"`
		EsDisenergy               api.UnitValue `json:"es_disenergy" PointId:"es_discharge_energy" PointAlias:"p83089" PointType:""`
		EsEnergy                  api.UnitValue `json:"es_energy" PointId:"es_energy" PointAlias:"p83120" PointType:""`
		EsPower                   api.UnitValue `json:"es_power" PointId:"es_power" PointAlias:"p83081" PointType:""`
		EsTotalDisenergy          api.UnitValue `json:"es_total_disenergy" PointId:"es_total_discharge_energy" PointAlias:"p83095" PointType:"PointTypeTotal"`
		EsTotalEnergy             api.UnitValue `json:"es_total_energy" PointId:"es_total_energy" PointAlias:"p83127" PointType:"PointTypeTotal"`
		ExpectInstallDate         api.DateTime        `json:"expect_install_date" PointId:"expect_install_date" PointType:""`
		FaultAlarmOfflineDevCount api.Integer   `json:"fault_alarm_offline_dev_count" PointId:"fault_alarm_offline_dev_count" PointType:""`
		FaultCount                api.Integer   `json:"fault_count" PointId:"fault_count" PointType:""`
		FaultDevCount             api.Integer   `json:"fault_dev_count" PointId:"fault_dev_count" PointType:""`
		GcjLatitude               api.Float     `json:"gcj_latitude" PointId:"gcj_latitude" PointType:""`
		GcjLongitude              api.Float     `json:"gcj_longitude" PointId:"gcj_longitude" PointType:""`
		GprsLatitude              api.Float     `json:"gprs_latitude" PointId:"gprs_latitude" PointType:""`
		GprsLongitude             api.Float     `json:"gprs_longitude" PointId:"gprs_longitude" PointType:""`
		Images                    []struct {
			FileID      api.Integer `json:"file_id"`
			ID          api.Integer `json:"id"`
			PicLanguage api.Integer `json:"pic_language"`
			PicType     api.Integer `json:"pic_type"`
			PictureName api.String      `json:"picture_name"`
			PictureURL  api.String      `json:"picture_url"`
			PsID        api.Integer `json:"ps_id"`
			PsUnitUUID  interface{} `json:"ps_unit_uuid"`
		} `json:"images"`
		InstallDate            api.DateTime        `json:"install_date" PointId:"install_date" PointType:""`
		InstalledPowerMap      api.UnitValue `json:"installed_power_map" PointId:"installed_power_map" PointType:""`
		InstalledPowerVirgin   api.Float     `json:"installed_power_virgin" PointId:"installed_power_virgin" PointType:""`
		InstallerAlarmCount    api.Integer   `json:"installer_alarm_count" PointId:"installer_alarm_count" PointType:""`
		InstallerFaultCount    api.Integer   `json:"installer_fault_count" PointId:"installer_fault_count" PointType:""`
		InstallerPsFaultStatus api.Integer   `json:"installer_ps_fault_status" PointId:"installer_ps_fault_status" PointType:""`
		IsBankPs               api.Bool      `json:"is_bank_ps" PointId:"is_bank_ps" PointType:""`
		IsTuv                  api.Bool      `json:"is_tuv" PointId:"is_tuv" PointType:""`
		JoinYearInitElec       api.Float     `json:"join_year_init_elec" PointId:"join_year_init_elec" PointType:""`
		Latitude               api.Float     `json:"latitude" PointId:"latitude" PointType:""`
		Location               api.String        `json:"location" PointId:"location" PointType:""`
		Longitude              api.Float     `json:"longitude" PointId:"longitude" PointType:""`
		MapLatitude            api.Float     `json:"map_latitude" PointId:"map_latitude" PointType:""`
		MapLongitude           api.Float     `json:"map_longitude" PointId:"map_longitude" PointType:""`
		MlpeFlag               api.Integer   `json:"mlpe_flag" PointId:"mlpe_flag" PointType:""`
		Nmi                    api.String   `json:"nmi" PointId:"nmi" PointType:""`
		OfflineDevCount        api.Integer   `json:"offline_dev_count" PointId:"offline_dev_count" PointType:""`
		OperateYear            interface{}   `json:"operate_year" PointId:"operate_year" PointType:""`
		OperationBusName       api.String   `json:"operation_bus_name" PointId:"operation_bus_name" PointType:""`
		OwnerAlarmCount        api.Integer   `json:"owner_alarm_count" PointId:"owner_alarm_count" PointType:""`
		OwnerFaultCount        api.Integer   `json:"owner_fault_count" PointId:"owner_fault_count" PointType:""`
		OwnerPsFaultStatus     api.Integer   `json:"owner_ps_fault_status" PointId:"owner_ps_fault_status" PointType:""`
		P83022y                api.String        `json:"p83022y" PointId:"P83022" PointType:""`
		P83046                 api.Float   `json:"p83046" PointId:"p83046" PointType:""`
		P83048                 api.Float   `json:"p83048" PointId:"p83048" PointType:""`
		P83049                 api.Float   `json:"p83049" PointId:"p83049" PointType:""`
		P83050                 api.Float   `json:"p83050" PointId:"p83050" PointType:""`
		P83051                 api.Float   `json:"p83051" PointId:"p83051" PointType:""`
		P83054                 api.Float   `json:"p83054" PointId:"p83054" PointType:""`
		P83055                 api.Float   `json:"p83055" PointId:"p83055" PointType:""`
		P83067                 api.Float   `json:"p83067" PointId:"p83067" PointType:""`
		P83070                 api.Float   `json:"p83070" PointId:"p83070" PointType:""`
		P83076                 api.Float     `json:"p83076" PointId:"p83076" PointIgnore:"true"`		// Dupe of PvPower
		P83077                 api.Float     `json:"p83077" PointId:"p83077" PointType:""`
		P83081                 api.Float     `json:"p83081" PointId:"p83081" PointIgnore:"true"`		// Dupe of EsPower
		P83089                 api.Float     `json:"p83089" PointId:"p83089" PointIgnore:"true"`		// Dupe of EsDisenergy
		P83095                 api.Float     `json:"p83095" PointId:"p83095" PointIgnore:"true"`		// Dupe of EsTotalDisenergy
		P83118                 api.Float     `json:"p83118" PointId:"p83118" `		// Dupe of UseEnergy
		P83120                 api.Float     `json:"p83120" PointId:"p83120" PointIgnore:"true"`		// Dupe of EsEnergy
		P83127                 api.Float     `json:"p83127" PointId:"p83127" PointIgnore:"true"`		// Dupe of EsTotalEnergy
		ParamCo2               api.Float     `json:"param_co2" PointId:"param_co2" PointType:""`
		ParamCoal              api.Float     `json:"param_coal" PointId:"param_coal" PointType:""`
		ParamIncome            api.Float     `json:"param_income" PointId:"param_income" PointType:""`
		ParamMeter             api.Float     `json:"param_meter" PointId:"param_meter" PointType:""`
		ParamNox               api.Float     `json:"param_nox" PointId:"param_nox" PointType:""`
		ParamPowder            api.Float     `json:"param_powder" PointId:"param_powder" PointType:""`
		ParamSo2               api.Float     `json:"param_so2" PointId:"param_so2" PointType:""`
		ParamTree              api.Float     `json:"param_tree" PointId:"param_tree" PointType:""`
		ParamWater             api.Float     `json:"param_water" PointId:"param_water" PointType:""`
		PrScale                string        `json:"pr_scale" PointId:"pr_scale" PointType:""`
		Producer               interface{}   `json:"producer" PointId:"producer" PointType:""`
		PsCountryID            api.Integer   `json:"ps_country_id" PointId:"ps_country_id" PointType:""`
		PsFaultStatus          api.Integer   `json:"ps_fault_status" PointId:"ps_fault_status" PointType:""`
		PsHealthStatus         string        `json:"ps_health_status" PointId:"ps_health_status" PointType:""`
		PsHolder               api.String        `json:"ps_holder" PointId:"ps_holder" PointType:""`
		PsID                   api.Integer   `json:"ps_id" PointId:"ps_id" PointType:""`
		PsIsNotInit            api.Bool        `json:"ps_is_not_init" PointId:"ps_is_not_init" PointType:""`
		PsName                 api.String        `json:"ps_name" PointId:"ps_name" PointType:""`
		PsShortName            api.String        `json:"ps_short_name" PointId:"ps_short_name" PointType:""`
		PsStatus               api.Integer   `json:"ps_status" PointId:"ps_status" PointType:""`
		PsTimezone             api.String        `json:"ps_timezone" PointId:"ps_timezone" PointType:""`
		PsType                 api.Integer   `json:"ps_type" PointId:"ps_type" PointType:""`
		PvEnergy               api.UnitValue `json:"pv_energy" PointId:"pv_energy" PointAlias:"p83077" PointType:""`
		PvPower                api.UnitValue `json:"pv_power" PointId:"pv_power" PointAlias:"p83076" PointType:""`
		Radiation              api.UnitValue `json:"radiation" PointId:"radiation" PointType:""`
		RadiationVirgin        api.Float   `json:"radiation_virgin" PointId:"radiation_virgin" PointType:""`
		RecordCreateTime       api.DateTime        `json:"recore_create_time" PointId:"recore_create_time" PointType:""`
		SafeStartDate          api.DateTime        `json:"safe_start_date" PointId:"safe_start_date" PointType:""`
		ShareType              string        `json:"share_type" PointId:"share_type" PointType:""`
		ShippingAddress        api.String        `json:"shipping_address" PointId:"shipping_address" PointType:""`
		ShippingZipCode        api.String        `json:"shipping_zip_code" PointId:"shipping_zip_code" PointType:""`
		TodayEnergy            api.UnitValue `json:"today_energy" PointId:"today_energy" PointType:"PointTypeDaily"`
		TodayIncome            api.UnitValue `json:"today_income" PointId:"today_income" PointType:"PointTypeDaily"`
		TotalCapacity           api.UnitValue `json:"total_capcity" PointId:"total_capacity" PointType:"PointTypeTotal"`
		TotalEnergy            api.UnitValue `json:"total_energy" PointId:"total_energy" PointType:"PointTypeTotal"`
		TotalIncome            api.UnitValue `json:"total_income" PointId:"total_income" PointType:"PointTypeTotal"`
		TotalInitCo2Accelerate api.Float     `json:"total_init_co2_accelerate" PointId:"total_init_co2_accelerate" PointType:"PointTypeTotal"`
		TotalInitElec          api.Float     `json:"total_init_elec" PointId:"total_init_elec" PointType:"PointTypeTotal"`
		TotalInitProfit        api.Float     `json:"total_init_profit" PointId:"total_init_profit" PointType:"PointTypeTotal"`
		UseEnergy              api.UnitValue `json:"use_energy" PointId:"use_energy" PointAlias:"p83118" PointType:""`
		ValidFlag              api.Integer   `json:"valid_flag" PointId:"valid_flag" PointType:""`
		WgsLatitude            api.Float     `json:"wgs_latitude" PointId:"wgs_latitude" PointType:""`
		WgsLongitude           api.Float     `json:"wgs_longitude" PointId:"wgs_longitude" PointType:""`
		ZipCode                api.String        `json:"zip_code" PointId:"zip_code" PointType:""`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
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

func (e *ResultData) GetPsIds() []api.Integer {
	var ret []api.Integer
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID.Value() != 0 {
				ret = append(ret, p.PsID)
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsName() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID.Value() != 0 {
				ret = append(ret, p.PsName.Value())
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsSerial() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsID.Value() != 0 {
				ret = append(ret, p.PsShortName.Value())
			}
		}
	}
	return ret
}

func (e *EndPoint) GetPsIds() []api.Integer {
	return e.Response.ResultData.GetPsIds()
}

func (e *EndPoint) GetData() api.DataMap {
	return e.Response.ResultData.GetData()
}

func (e *ResultData) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		if len(e.PageList) == 0 {
			break
		}

		// now := api.NewDateTime(time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds))

		for _, p := range e.PageList {
			psId := p.PsID.String()		// psId := strconv.FormatInt(p.PsID.Value(), 10)
			name := "getPsList." + psId
			entries.StructToPoints(p, name, psId, time.Time{})
		}
	}
	return entries
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
			"Point Name",
			"Value",
			"Unit",
		)

		now := time.Now().Round(5 * time.Minute).Format(api.DtLayoutZeroSeconds)

		for _, p := range e.Response.ResultData.PageList {
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Co2 Reduce", p.Co2Reduce.Value(), p.Co2Reduce.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Co2 Reduce Total", p.Co2ReduceTotal.Value(), p.Co2ReduceTotal.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Curr Power", p.CurrPower.Value(), p.CurrPower.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Daily Irradiation", p.DailyIrradiation.Value(), p.DailyIrradiation.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Equivalent Hour", p.EquivalentHour.Value(), p.EquivalentHour.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Es Discharge Energy", p.EsDisenergy.Value(), p.EsDisenergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Es Energy", p.EsEnergy.Value(), p.EsEnergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Es Power", p.EsPower.Value(), p.EsPower.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Es Total Discharge Energy", p.EsTotalDisenergy.Value(), p.EsTotalDisenergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Es Total Energy", p.EsTotalEnergy.Value(), p.EsTotalEnergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Installed Power Map", p.InstalledPowerMap.Value(), p.InstalledPowerMap.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Pv Energy", p.PvEnergy.Value(), p.PvEnergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Pv Power", p.PvPower.Value(), p.PvPower.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Radiation", p.Radiation.Value(), p.Radiation.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Today Energy", p.TodayEnergy.Value(), p.TodayEnergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Today Income", p.TodayIncome.Value(), p.TodayIncome.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Total Capacity", p.TotalCapacity.Value(), p.TotalCapacity.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Total Energy", p.TotalEnergy.Value(), p.TotalEnergy.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Total Income", p.TotalIncome.Value(), p.TotalIncome.Unit())
			_ = table.AddRow(now, fmt.Sprintf("%s.%d", p.PsID.String(), 0), "Use Energy", p.UseEnergy.Value(), p.UseEnergy.Unit())
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
