package queryPowerStationInfo

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/queryPowerStationInfo"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id"`
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
	ArrearsStatus         int64         `json:"arrears_status"`
	BatteryType           string        `json:"battery_type"`
	CityCode              interface{}   `json:"city_code"`
	CityName              interface{}   `json:"city_name"`
	ComponentArea         interface{}   `json:"component_area"`
	ComponentStatus       int64         `json:"component_status"`
	ConnectType           int64         `json:"connect_type"`
	ConnectTypeDesc       string        `json:"connect_type_desc"`
	ContactName           string        `json:"contact_name"`
	CountryID             int64         `json:"country_id"`
	Description           interface{}   `json:"description"`
	DesignCapacity        float64       `json:"design_capacity"`
	DesignCapacityBattery string        `json:"design_capacity_battery"`
	DistrictCode          interface{}   `json:"district_code"`
	DistrictName          interface{}   `json:"district_name"`
	DivisionCode          interface{}   `json:"division_code"`
	Email                 string        `json:"email"`
	EnergyScheme          interface{}   `json:"energy_scheme"`
	ExpectInstallDate     string        `json:"expect_install_date"`
	GcjLatitude           string        `json:"gcj_latitude"`
	GcjLongitude          string        `json:"gcj_longitude"`
	GprsLatitude          interface{}   `json:"gprs_latitude"`
	GprsLongitude         interface{}   `json:"gprs_longitude"`
	GridLevel             interface{}   `json:"grid_level"`
	Images                []interface{} `json:"images"`
	InstallDate           string        `json:"install_date"`
	InstallDateZone       string        `json:"install_date_zone"`
	InverterCount         int64         `json:"inverter_count"`
	InvestmentType        int64         `json:"investment_type"`
	InvestmentTypeDesc    string        `json:"investment_type_desc"`
	IsAgreeGdpr           int64         `json:"is_agree_gdpr"`
	IsGdpr                int64         `json:"is_gdpr"`
	IsNewVersion          int64         `json:"is_new_version"`
	IsOpenProtocol        int64         `json:"is_open_protocol"`
	IsPsCreateUser        string        `json:"is_ps_create_user"`
	IsPsOwner             string        `json:"is_ps_owner"`
	IsReceiveNotice       int64         `json:"is_receive_notice"`
	IsSharePosition       int64         `json:"is_share_position"`
	IsValidMobileEmail    int64         `json:"is_valid_mobile_email"`
	Latitude              float64       `json:"latitude"`
	Longitude             float64       `json:"longitude"`
	MapLatitude           string        `json:"map_latitude"`
	MapLongitude          string        `json:"map_longitude"`
	MlpeFlag              int64         `json:"mlpe_flag"`
	MobleTel              interface{}   `json:"moble_tel"`
	MobleTelBak           interface{}   `json:"moble_tel_bak"`
	ModuleModelID         interface{}   `json:"module_model_id"`
	ModuleModelName       interface{}   `json:"module_model_name"`
	Nmi                   string        `json:"nmi"`
	OperationBusName      interface{}   `json:"operation_bus_name"`
	OrgIndexCode          []string      `json:"org_index_code"`
	OwnerContact          interface{}   `json:"owner_contact"`
	ParamIncomeUnit       int64         `json:"param_income_unit"`
	ParamIncomeUnitName   string        `json:"param_income_unit_name"`
	ProvinceCode          interface{}   `json:"province_code"`
	ProvinceName          interface{}   `json:"province_name"`
	PsBuildDate           string        `json:"ps_build_date"`
	PsCountryID           int64         `json:"ps_country_id"`
	PsCreateUserID        int64         `json:"ps_create_user_id"`
	PsCurrentTimeZone     string        `json:"ps_current_time_zone"`
	PsDirectOrgList       []struct {
		OrgID        int64  `json:"org_id"`
		OrgIndexCode string `json:"org_index_code"`
		OrgName      string `json:"org_name"`
	} `json:"ps_direct_org_list"`
	PsHolder         string `json:"ps_holder"`
	PsID             int64  `json:"ps_id"`
	PsInstalledPower float64 `json:"ps_installed_power"`
	PsKey            string `json:"ps_key"`
	PsLocation       string `json:"ps_location"`
	PsName           string `json:"ps_name"`
	PsOrgInfo        []struct {
		DealerOrgCode   string `json:"dealer_org_code"`
		Installer       string `json:"installer"`
		InstallerEmail  string `json:"installer_email"`
		InstallerPhone  string `json:"installer_phone"`
		OrgID           int64  `json:"org_id"`
		OrgIndexCode    string `json:"org_index_code"`
		OrgName         string `json:"org_name"`
		PsDealerOrgCode string `json:"ps_dealer_org_code"`
		UpOrgID         int64  `json:"up_org_id"`
	} `json:"ps_org_info"`
	PsPrice          string `json:"ps_price"`
	PsPriceKwh       string `json:"ps_price_kwh"`
	PsType           int64  `json:"ps_type"`
	PsTypeDesc       string `json:"ps_type_desc"`
	PsTypeName       string `json:"ps_type_name"`
	PsUserID         int64  `json:"ps_user_id"`
	RecoreCreateTime string `json:"recore_create_time"`
	SafeStartDate    string `json:"safe_start_date"`
	SelectedOrgList  []struct {
		OrgID        int64  `json:"org_id"`
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
		CommunicateDeviceType     int64  `json:"communicate_device_type"`
		CommunicateDeviceTypeName string `json:"communicate_device_type_name"`
		ID                        int64  `json:"id"`
		IsEnable                  int64  `json:"is_enable"`
		Sn                        string `json:"sn"`
	} `json:"sn_detail_list"`
	SummerTimeState    int64       `json:"summer_time_state"`
	SummerTimeZone     string      `json:"summer_time_zone"`
	SummerTimeZoneID   int64       `json:"summer_time_zone_id"`
	TimeZoneID         int64       `json:"time_zone_id"`
	Timezone           string      `json:"timezone"`
	UserAccount        string      `json:"user_account"`
	UserCapacityStatus int64       `json:"user_capacity_status"`
	UserEnglishName    interface{} `json:"user_english_name"`
	UserLanguage       string      `json:"user_language"`
	UserMobleTel       interface{} `json:"user_moble_tel"`
	UserName           string      `json:"user_name"`
	UserTelNationCode  interface{} `json:"user_tel_nation_code"`
	ValidFlag          int64       `json:"valid_flag"`
	WgsLatitude        float64     `json:"wgs_latitude"`
	WgsLongitude       float64     `json:"wgs_longitude"`
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
