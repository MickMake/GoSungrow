package getPsList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
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
