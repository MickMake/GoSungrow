package getPowerStationForHousehold

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPowerStationForHousehold"
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
	LbsAreaCode           api.String   `json:"LbsAreaCode"`
	LbsCountry            api.String   `json:"LbsCountry"`
	AccessType            interface{}  `json:"access_type"`
	AreaID                interface{}  `json:"area_id"`
	ArrearsStatus         api.Integer  `json:"arrears_status"`
	BatteryType           string       `json:"battery_type"`
	CityCode              api.String   `json:"city_code"`
	CityName              api.String   `json:"city_name"`
	ComponentArea         interface{}  `json:"component_area"`
	ComponentStatus       api.Integer  `json:"component_status"`
	ConnectType           api.Integer  `json:"connect_type"`
	ConnectTypeDesc       api.String   `json:"connect_type_desc"`
	ContactName           api.String   `json:"contact_name"`
	CountryID             api.Integer  `json:"country_id"`
	Description           api.String   `json:"description"`
	DesignCapacity        api.Float    `json:"design_capacity"`
	DesignCapacityBattery api.Float    `json:"design_capacity_battery"`
	DistrictCode          api.String   `json:"district_code"`
	DistrictName          api.String   `json:"district_name"`
	DivisionCode          api.String   `json:"division_code"`
	Email                 api.String   `json:"email"`
	EnergyScheme          interface{}  `json:"energy_scheme"`
	ExpectInstallDate     api.DateTime `json:"expect_install_date"`
	GcjLatitude           api.Float    `json:"gcj_latitude"`
	GcjLongitude          api.Float    `json:"gcj_longitude"`
	GprsLatitude          api.Float    `json:"gprs_latitude"`
	GprsLongitude         api.Float    `json:"gprs_longitude"`
	GridLevel             interface{}  `json:"grid_level"`
	Images                []api.String `json:"images"`
	InstallDate           api.DateTime `json:"install_date"`
	InstallDateZone       api.DateTime `json:"install_date_zone"`
	InverterCount         api.Integer  `json:"inverter_count"`
	InvestmentType        api.Integer  `json:"investment_type"`
	InvestmentTypeDesc    api.String   `json:"investment_type_desc"`
	IsAgreeGdpr           api.Bool     `json:"is_agree_gdpr"`
	IsGdpr                api.Bool     `json:"is_gdpr"`
	IsNewVersion          api.Bool     `json:"is_new_version"`
	IsOpenProtocol        api.Bool     `json:"is_open_protocol"`
	IsPsCreateUser        api.Bool     `json:"is_ps_create_user"`
	IsPsOwner             api.Bool     `json:"is_ps_owner"`
	IsReceiveNotice       api.Bool     `json:"is_receive_notice"`
	IsSharePosition       api.Bool     `json:"is_share_position"`
	IsValidMobileEmail    api.Bool     `json:"is_valid_mobile_email"`
	Latitude              api.Float    `json:"latitude"`
	Longitude             api.Float    `json:"longitude"`
	MapLatitude           api.Float    `json:"map_latitude"`
	MapLongitude          api.Float    `json:"map_longitude"`
	MlpeFlag              api.Integer  `json:"mlpe_flag"`
	MobileTel             api.String   `json:"moble_tel"`
	MobileTelBak          api.String   `json:"moble_tel_bak"`
	ModuleModelID         interface{}  `json:"module_model_id"`
	ModuleModelName       api.String   `json:"module_model_name"`
	Nmi                   api.String   `json:"nmi"`
	OperationBusName      api.String   `json:"operation_bus_name"`
	OrgIndexCode          []api.String `json:"org_index_code"`
	OwnerContact          api.String   `json:"owner_contact"`
	ParamIncomeUnit       api.Integer  `json:"param_income_unit"`
	ParamIncomeUnitName   api.String   `json:"param_income_unit_name"`
	ProvinceCode          api.String   `json:"province_code"`
	ProvinceName          api.String   `json:"province_name"`
	PsBuildDate           api.DateTime `json:"ps_build_date"`
	PsCountryID           api.Integer  `json:"ps_country_id"`
	PsCreateUserID        api.Integer  `json:"ps_create_user_id"`
	PsCurrentTimeZone     string       `json:"ps_current_time_zone"`
	PsDirectOrgList       []struct {
		OrgID        api.Integer `json:"org_id"`
		OrgIndexCode api.String  `json:"org_index_code"`
		OrgName      api.String  `json:"org_name"`
	} `json:"ps_direct_org_list"`
	PsHolder         api.String  `json:"ps_holder"`
	PsID             api.Integer `json:"ps_id"`
	PsInstalledPower api.Float   `json:"ps_installed_power"`
	PsKey            api.PsKey   `json:"ps_key"`
	PsLocation       api.String  `json:"ps_location"`
	PsName           api.String  `json:"ps_name"`
	PsOrgInfo        []struct {
		DealerOrgCode   api.String  `json:"dealer_org_code"`
		Installer       api.String  `json:"installer"`
		InstallerEmail  api.String  `json:"installer_email"`
		InstallerPhone  api.String  `json:"installer_phone"`
		OrgID           api.Integer `json:"org_id"`
		OrgIndexCode    api.String  `json:"org_index_code"`
		OrgName         api.String  `json:"org_name"`
		PsDealerOrgCode api.String  `json:"ps_dealer_org_code"`
		UpOrgID         api.Integer `json:"up_org_id"`
	} `json:"ps_org_info"`
	PsPrice          api.Float    `json:"ps_price"`
	PsPriceKwh       api.Float    `json:"ps_price_kwh"`
	PsType           api.Integer  `json:"ps_type"`
	PsTypeDesc       api.String   `json:"ps_type_desc"`
	PsTypeName       api.String   `json:"ps_type_name"`
	PsUserID         api.Integer  `json:"ps_user_id"`
	RecordCreateTime api.DateTime `json:"recore_create_time"`
	SafeStartDate    api.DateTime `json:"safe_start_date"`
	SelectedOrgList  []struct {
		OrgID        api.Integer `json:"org_id"`
		OrgIndexCode api.String  `json:"org_index_code"`
		OrgName      api.String  `json:"org_name"`
	} `json:"selectedOrgList"`
	SetUserOrg      string      `json:"set_user_org"`
	ShareType       string      `json:"share_type"`
	ShareUserType   interface{} `json:"share_user_type"`
	ShippingAddress api.String  `json:"shipping_address"`
	ShippingZipCode api.String  `json:"shipping_zip_code"`
	Sn              api.String  `json:"sn"`
	SnDetailList    []struct {
		CommunicateDeviceType     api.Integer `json:"communicate_device_type"`
		CommunicateDeviceTypeName api.String  `json:"communicate_device_type_name"`
		ID                        api.Integer `json:"id"`
		IsEnable                  api.Bool    `json:"is_enable"`
		Sn                        api.String  `json:"sn"`
	} `json:"sn_detail_list"`
	SummerTimeState    api.Integer `json:"summer_time_state"`
	SummerTimeZone     api.String  `json:"summer_time_zone"`
	SummerTimeZoneID   api.Integer `json:"summer_time_zone_id"`
	TimeZoneID         api.Integer `json:"time_zone_id"`
	Timezone           api.String  `json:"timezone"`
	UserAccount        api.String  `json:"user_account"`
	UserCapacityStatus api.Integer `json:"user_capacity_status"`
	UserEnglishName    api.String  `json:"user_english_name"`
	UserLanguage       api.String  `json:"user_language"`
	UserMobileTel      api.String  `json:"user_moble_tel"`
	UserName           api.String  `json:"user_name"`
	UserTelNationCode  api.String  `json:"user_tel_nation_code"`
	ValidFlag          api.Integer `json:"valid_flag"`
	WgsLatitude        api.Float   `json:"wgs_latitude"`
	WgsLongitude       api.Float   `json:"wgs_longitude"`
	ZipCode            api.String  `json:"zip_code"`
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
