package getPowerStationInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPowerStationInfo"
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
	CurrencyTypeList []struct {
		CodeName  string `json:"code_name"`
		CodeValue string `json:"code_value"`
	} `json:"currencyTypeList"`
	InstallProviderInfo struct {
		Installer      string      `json:"installer"`
		InstallerEmail string      `json:"installer_email"`
		InstallerPhone string      `json:"installer_phone"`
		OrgURL         interface{} `json:"org_url"`
		UpOrgID        api.Integer `json:"up_org_id"`
	} `json:"installProviderInfo"`
	ParallelTypes []struct {
		CodeName  string `json:"code_name"`
		CodeValue string `json:"code_value"`
	} `json:"parallelTypes"`
	ParamIncome        api.Float `json:"param_income"`
	PowerChargeDataMap struct {
		CodeType            api.Integer `json:"code_type"`
		DefaultCharge       api.Float   `json:"default_charge"`
		ElectricChargeID    api.Integer `json:"electric_charge_id"`
		EndTime             string      `json:"end_time"`
		IntervalTimeCharge  interface{} `json:"interval_time_charge"`
		ParamIncomeUnitName string      `json:"param_income_unit_name"`
		StartTime           string      `json:"start_time"`
	} `json:"powerChargeDataMap"`
	PowerPictureList []interface{} `json:"powerPictureList"`
	PowerStationMap  struct {
		AreaID            interface{} `json:"area_id"`
		ConnectType       api.Integer `json:"connect_type"`
		County            interface{} `json:"county"`
		CountyCode        interface{} `json:"county_code"`
		DesignCapacity2   api.Float   `json:"designCapacity"`
		DesignCapacity    api.Float   `json:"design_capacity"`
		Email             string      `json:"email"`
		EnergyScheme      interface{} `json:"energy_scheme"`
		ExpectInstallDate api.DateTime      `json:"expect_install_date"`
		FaultSendType     interface{} `json:"fault_send_type"`
		GcjLatitude       api.Float   `json:"gcj_latitude"`
		GcjLongitude      api.Float   `json:"gcj_longitude"`
		Latitude          api.Float   `json:"latitude"`
		Longitude         api.Float   `json:"longitude"`
		Name              interface{} `json:"name"`
		NameCode          interface{} `json:"name_code"`
		Nation            interface{} `json:"nation"`
		NationCode        interface{} `json:"nation_code"`
		ParamIncome       api.Float   `json:"param_income"`
		Prov              interface{} `json:"prov"`
		ProvCode          interface{} `json:"prov_code"`
		PsBuildDate       api.DateTime      `json:"ps_build_date"`
		PsCountryID       api.Integer `json:"ps_country_id"`
		PsDesc            interface{} `json:"ps_desc"`
		PsHolder          string      `json:"ps_holder"`
		PsID              api.Integer `json:"ps_id"`
		PsLocation        string      `json:"ps_location"`
		PsName            string      `json:"ps_name"`
		PsType            api.Integer `json:"ps_type"`
		RecordCreateTime  string      `json:"recore_create_time"`
		ReportType        interface{} `json:"report_type"`
		ShippingAddress   string      `json:"shipping_address"`
		ShippingZipCode   string      `json:"shipping_zip_code"`
		TimeZoneID        api.Integer `json:"time_zone_id"`
		ValidFlag         api.Integer `json:"valid_flag"`
		Village           interface{} `json:"village"`
		VillageCode       interface{} `json:"village_code"`
		WgsLatitude       api.Float   `json:"wgs_latitude"`
		WgsLongitude      api.Float   `json:"wgs_longitude"`
		ZipCode           string      `json:"zip_code"`
	} `json:"powerStationMap"`
	RemindType           interface{} `json:"remindType"`
	SendReportConfigList []string    `json:"sendReportConfigList"`
	StatusList           []struct {
		CodeName  string `json:"code_name"`
		CodeValue string `json:"code_value"`
	} `json:"statusList"`
	SysTimeZones []struct {
		ID           api.Integer  `json:"id"`
		TimezoneName string `json:"timezone_name"`
		TimezoneUtc  string `json:"timezone_utc"`
	} `json:"sysTimeZones"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
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
