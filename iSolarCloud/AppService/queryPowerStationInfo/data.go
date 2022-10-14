package queryPowerStationInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/powerStationService/queryPowerStationInfo"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id"`
	Sn   string `json:"sn"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData   struct {
	LbsAreaCode           interface{}   `json:"LbsAreaCode"`
	LbsCountry            interface{}   `json:"LbsCountry"`
	AccessType            interface{}   `json:"access_type"`
	AreaID                interface{}   `json:"area_id"`
	ArrearsStatus         valueTypes.Integer   `json:"arrears_status"`
	BatteryType           string        `json:"battery_type"`
	CityCode              interface{}   `json:"city_code"`
	CityName              interface{}   `json:"city_name"`
	ComponentArea         interface{}   `json:"component_area"`
	ComponentStatus       valueTypes.Integer   `json:"component_status"`
	ConnectType           valueTypes.Integer   `json:"connect_type"`
	ConnectTypeDesc       string        `json:"connect_type_desc"`
	ContactName           string        `json:"contact_name"`
	CountryID             valueTypes.Integer   `json:"country_id"`
	Description           interface{}   `json:"description"`
	DesignCapacity        valueTypes.Float     `json:"design_capacity"`
	DesignCapacityBattery string        `json:"design_capacity_battery"`
	DistrictCode          interface{}   `json:"district_code"`
	DistrictName          interface{}   `json:"district_name"`
	DivisionCode          interface{}   `json:"division_code"`
	Email                 string        `json:"email"`
	EnergyScheme          interface{}   `json:"energy_scheme"`
	ExpectInstallDate     valueTypes.DateTime        `json:"expect_install_date"`
	GcjLatitude           valueTypes.Float     `json:"gcj_latitude"`
	GcjLongitude          valueTypes.Float     `json:"gcj_longitude"`
	GprsLatitude          valueTypes.Float     `json:"gprs_latitude"`
	GprsLongitude         valueTypes.Float     `json:"gprs_longitude"`
	GridLevel             interface{}   `json:"grid_level"`
	Images                []interface{} `json:"images"`
	InstallDate           valueTypes.DateTime        `json:"install_date"`
	InstallDateZone       string        `json:"install_date_zone"`
	InverterCount         valueTypes.Integer   `json:"inverter_count"`
	InvestmentType        valueTypes.Integer   `json:"investment_type"`
	InvestmentTypeDesc    string        `json:"investment_type_desc"`
	IsAgreeGdpr           valueTypes.Bool      `json:"is_agree_gdpr"`
	IsGdpr                valueTypes.Bool      `json:"is_gdpr"`
	IsNewVersion          valueTypes.Bool      `json:"is_new_version"`
	IsOpenProtocol        valueTypes.Bool      `json:"is_open_protocol"`
	IsPsCreateUser        valueTypes.Bool      `json:"is_ps_create_user"`
	IsPsOwner             valueTypes.Bool      `json:"is_ps_owner"`
	IsReceiveNotice       valueTypes.Bool      `json:"is_receive_notice"`
	IsSharePosition       valueTypes.Bool      `json:"is_share_position"`
	IsValidMobileEmail    valueTypes.Bool      `json:"is_valid_mobile_email"`
	Latitude              valueTypes.Float     `json:"latitude"`
	Longitude             valueTypes.Float     `json:"longitude"`
	MapLatitude           valueTypes.Float     `json:"map_latitude"`
	MapLongitude          valueTypes.Float     `json:"map_longitude"`
	MlpeFlag              valueTypes.Integer   `json:"mlpe_flag"`
	MobleTel              interface{}   `json:"moble_tel"`
	MobleTelBak           interface{}   `json:"moble_tel_bak"`
	ModuleModelID         interface{}   `json:"module_model_id"`
	ModuleModelName       interface{}   `json:"module_model_name"`
	Nmi                   string        `json:"nmi"`
	OperationBusName      interface{}   `json:"operation_bus_name"`
	OrgIndexCode          []string      `json:"org_index_code"`
	OwnerContact          interface{}   `json:"owner_contact"`
	ParamIncomeUnit       valueTypes.Integer   `json:"param_income_unit"`
	ParamIncomeUnitName   string        `json:"param_income_unit_name"`
	ProvinceCode          interface{}   `json:"province_code"`
	ProvinceName          interface{}   `json:"province_name"`
	PsBuildDate           valueTypes.DateTime        `json:"ps_build_date"`
	PsCountryID           valueTypes.Integer   `json:"ps_country_id"`
	PsCreateUserID        valueTypes.Integer   `json:"ps_create_user_id"`
	PsCurrentTimeZone     string        `json:"ps_current_time_zone"`
	PsDirectOrgList       []struct {
		OrgID        valueTypes.Integer  `json:"org_id"`
		OrgIndexCode string `json:"org_index_code"`
		OrgName      string `json:"org_name"`
	} `json:"ps_direct_org_list"`
	PsHolder         string `json:"ps_holder"`
	PsID             valueTypes.Integer  `json:"ps_id"`
	PsInstalledPower valueTypes.Float `json:"ps_installed_power"`
	PsKey            valueTypes.PsKey `json:"ps_key"`
	PsLocation       string `json:"ps_location"`
	PsName           string `json:"ps_name"`
	PsOrgInfo        []struct {
		DealerOrgCode   string `json:"dealer_org_code"`
		Installer       string `json:"installer"`
		InstallerEmail  string `json:"installer_email"`
		InstallerPhone  string `json:"installer_phone"`
		OrgID           valueTypes.Integer  `json:"org_id"`
		OrgIndexCode    string `json:"org_index_code"`
		OrgName         string `json:"org_name"`
		PsDealerOrgCode string `json:"ps_dealer_org_code"`
		UpOrgID         valueTypes.Integer  `json:"up_org_id"`
	} `json:"ps_org_info"`
	PsPrice          string `json:"ps_price"`
	PsPriceKwh       string `json:"ps_price_kwh"`
	PsType           valueTypes.Integer  `json:"ps_type"`
	PsTypeDesc       string `json:"ps_type_desc"`
	PsTypeName       string `json:"ps_type_name"`
	PsUserID         valueTypes.Integer  `json:"ps_user_id"`
	RecoreCreateTime string `json:"recore_create_time"`
	SafeStartDate    valueTypes.DateTime `json:"safe_start_date"`
	SelectedOrgList  []struct {
		OrgID        valueTypes.Integer  `json:"org_id"`
		OrgIndexCode string `json:"org_index_code"`
		OrgName      string `json:"org_name"`
	} `json:"selectedOrgList"`
	SetUserOrg      string      `json:"set_user_org"`
	ShareType       string      `json:"share_type"`
	ShareUserType   interface{} `json:"share_user_type"`
	ShippingAddress string      `json:"shipping_address"`
	ShippingZipCode string      `json:"shipping_zip_code"`
	Sn              string      `json:"sn"`
	SnDetailList    []struct {
		CommunicateDeviceType     valueTypes.Integer  `json:"communicate_device_type"`
		CommunicateDeviceTypeName string `json:"communicate_device_type_name"`
		ID                        valueTypes.Integer  `json:"id"`
		IsEnable                  valueTypes.Bool     `json:"is_enable"`
		Sn                        string `json:"sn"`
	} `json:"sn_detail_list"`
	SummerTimeState    valueTypes.Integer `json:"summer_time_state"`
	SummerTimeZone     string      `json:"summer_time_zone"`
	SummerTimeZoneID   valueTypes.Integer `json:"summer_time_zone_id"`
	TimeZoneID         valueTypes.Integer `json:"time_zone_id"`
	Timezone           string      `json:"timezone"`
	UserAccount        string      `json:"user_account"`
	UserCapacityStatus valueTypes.Integer `json:"user_capacity_status"`
	UserEnglishName    interface{} `json:"user_english_name"`
	UserLanguage       string      `json:"user_language"`
	UserMobleTel       interface{} `json:"user_moble_tel"`
	UserName           string      `json:"user_name"`
	UserTelNationCode  interface{} `json:"user_tel_nation_code"`
	ValidFlag          valueTypes.Integer `json:"valid_flag"`
	WgsLatitude        valueTypes.Float   `json:"wgs_latitude"`
	WgsLongitude       valueTypes.Float   `json:"wgs_longitude"`
	ZipCode            string      `json:"zip_code"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
