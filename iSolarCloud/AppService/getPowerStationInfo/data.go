package getPowerStationInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationInfo"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
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
		CodeName  valueTypes.String `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"currencyTypeList"`
	InstallProviderInfo struct {
		Installer      valueTypes.String      `json:"installer"`
		InstallerEmail valueTypes.String      `json:"installer_email"`
		InstallerPhone valueTypes.String      `json:"installer_phone"`
		OrgURL         interface{} `json:"org_url"`
		UpOrgID        valueTypes.Integer `json:"up_org_id"`
	} `json:"installProviderInfo"`
	ParallelTypes []struct {
		CodeName  valueTypes.String `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"parallelTypes"`
	ParamIncome        valueTypes.Float `json:"param_income"`
	PowerChargeDataMap struct {
		CodeType            valueTypes.Integer `json:"code_type"`
		DefaultCharge       valueTypes.Float   `json:"default_charge"`
		ElectricChargeID    valueTypes.Integer `json:"electric_charge_id"`
		EndTime             valueTypes.DateTime      `json:"end_time"`
		IntervalTimeCharge  interface{} `json:"interval_time_charge"`
		ParamIncomeUnitName valueTypes.String      `json:"param_income_unit_name"`
		StartTime           valueTypes.DateTime      `json:"start_time"`
	} `json:"powerChargeDataMap"`
	PowerPictureList []interface{} `json:"powerPictureList"`
	PowerStationMap  struct {
		AreaID            interface{} `json:"area_id"`
		ConnectType       valueTypes.Integer `json:"connect_type"`
		County            interface{} `json:"county"`
		CountyCode        interface{} `json:"county_code"`
		DesignCapacity2   valueTypes.Float   `json:"designCapacity" PointIgnore:"true"`
		DesignCapacity    valueTypes.Float   `json:"design_capacity"`
		Email             valueTypes.String      `json:"email"`
		EnergyScheme      interface{} `json:"energy_scheme"`
		ExpectInstallDate valueTypes.DateTime      `json:"expect_install_date"`
		FaultSendType     interface{} `json:"fault_send_type"`
		GcjLatitude       valueTypes.Float   `json:"gcj_latitude"`
		GcjLongitude      valueTypes.Float   `json:"gcj_longitude"`
		Latitude          valueTypes.Float   `json:"latitude"`
		Longitude         valueTypes.Float   `json:"longitude"`
		Name              interface{} `json:"name"`
		NameCode          interface{} `json:"name_code"`
		Nation            interface{} `json:"nation"`
		NationCode        interface{} `json:"nation_code"`
		ParamIncome       valueTypes.Float   `json:"param_income"`
		Prov              interface{} `json:"prov"`
		ProvCode          interface{} `json:"prov_code"`
		PsBuildDate       valueTypes.DateTime      `json:"ps_build_date"`
		PsCountryID       valueTypes.Integer `json:"ps_country_id"`
		PsDesc            interface{} `json:"ps_desc"`
		PsHolder          valueTypes.String      `json:"ps_holder"`
		PsID              valueTypes.Integer `json:"ps_id"`
		PsLocation        valueTypes.String      `json:"ps_location"`
		PsName            valueTypes.String      `json:"ps_name"`
		PsType            valueTypes.Integer `json:"ps_type"`
		RecordCreateTime  valueTypes.DateTime      `json:"recore_create_time"`
		ReportType        interface{} `json:"report_type"`
		ShippingAddress   valueTypes.String      `json:"shipping_address"`
		ShippingZipCode   valueTypes.String      `json:"shipping_zip_code"`
		TimeZoneID        valueTypes.Integer `json:"time_zone_id"`
		ValidFlag         valueTypes.Integer `json:"valid_flag"`
		Village           interface{} `json:"village"`
		VillageCode       interface{} `json:"village_code"`
		WgsLatitude       valueTypes.Float   `json:"wgs_latitude"`
		WgsLongitude      valueTypes.Float   `json:"wgs_longitude"`
		ZipCode           valueTypes.String      `json:"zip_code"`
	} `json:"powerStationMap"`
	RemindType           interface{} `json:"remindType"`
	SendReportConfigList []valueTypes.String    `json:"sendReportConfigList"`
	StatusList           []struct {
		CodeName  valueTypes.String `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"statusList"`
	SysTimeZones []struct {
		ID           valueTypes.Integer  `json:"id"`
		TimezoneName valueTypes.String `json:"timezone_name"`
		TimezoneUtc  valueTypes.String `json:"timezone_utc"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e) + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, pkg, e.Request.PsId.String(), valueTypes.NewDateTime(""))

		for i, v := range e.Response.ResultData.CurrencyTypeList {
			name := fmt.Sprintf("%s.CurrencyTypeList.%.2d", pkg, i)
			entries.StructToPoints(v, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		}

		for i, v := range e.Response.ResultData.ParallelTypes {
			name := fmt.Sprintf("%s.ParallelTypes.%.2d", pkg, i)
			entries.StructToPoints(v, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		}

		for i, v := range e.Response.ResultData.StatusList {
			name := fmt.Sprintf("%s.StatusList.%.2d", pkg, i)
			entries.StructToPoints(v, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		}

		for i, v := range e.Response.ResultData.SysTimeZones {
			name := fmt.Sprintf("%s.SysTimeZones.%.2d", pkg, i)
			entries.StructToPoints(v, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		}

	}

	return entries
}
