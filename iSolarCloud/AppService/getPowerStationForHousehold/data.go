package getPowerStationForHousehold

import (
	"GoSungrow/iSolarCloud/Common"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationForHousehold"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Images          Common.PowerStationImages `json:"images" PointArrayFlatten:"false"`
	PsDirectOrgList Common.PsDirectOrgList    `json:"ps_direct_org_list" PointArrayFlatten:"false"`
	PsOrgInfo       Common.PsOrgInfo          `json:"ps_org_info" PointArrayFlatten:"false"`
	SelectedOrgList Common.SelectedOrgList    `json:"selectedOrgList" PointId:"selected_org_list" PointArrayFlatten:"false"`
	SnDetailList    Common.SnDetailList       `json:"sn_detail_list" PointArrayFlatten:"false"`

	LbsAreaCode           valueTypes.String   `json:"LbsAreaCode" PointId:"lbs_area_code"`
	LbsCountry            valueTypes.String   `json:"LbsCountry" PointId:"lbs_country"`
	AccessType            interface{}         `json:"access_type"`
	AreaId                interface{}         `json:"area_id"`
	ArrearsStatus         valueTypes.Integer  `json:"arrears_status"`
	BatteryType           valueTypes.Integer  `json:"battery_type"`
	CityCode              valueTypes.String   `json:"city_code"`
	CityName              valueTypes.String   `json:"city_name"`
	ComponentArea         interface{}         `json:"component_area"`
	ComponentStatus       valueTypes.Integer  `json:"component_status"`
	ConnectType           valueTypes.Integer  `json:"connect_type"`
	ConnectTypeDesc       valueTypes.String   `json:"connect_type_desc"`
	ContactName           valueTypes.String   `json:"contact_name"`
	CountryId             valueTypes.Integer  `json:"country_id"`
	Description           valueTypes.String   `json:"description"`
	DesignCapacity        valueTypes.Float    `json:"design_capacity" PointUnit:"W"` // @TODO - When this is set to valueTypes.Integer, we get a failure.
	DesignCapacityBattery valueTypes.Float    `json:"design_capacity_battery" PointUnit:"W"`
	DistrictCode          valueTypes.String   `json:"district_code"`
	DistrictName          valueTypes.String   `json:"district_name"`
	DivisionCode          valueTypes.String   `json:"division_code"`
	Email                 valueTypes.String   `json:"email"`
	EnergyScheme          interface{}         `json:"energy_scheme"`
	ExpectInstallDate     valueTypes.DateTime `json:"expect_install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	GcjLatitude           valueTypes.Float    `json:"gcj_latitude"`
	GcjLongitude          valueTypes.Float    `json:"gcj_longitude"`
	GprsLatitude          valueTypes.Float    `json:"gprs_latitude"`
	GprsLongitude         valueTypes.Float    `json:"gprs_longitude"`
	GridLevel             interface{}         `json:"grid_level"`
	InstallDate           valueTypes.DateTime `json:"install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	InstallDateZone       valueTypes.DateTime `json:"install_date_zone" PointNameDateFormat:"2006/01/02 15:04:05"`
	InverterCount         valueTypes.Integer  `json:"inverter_count"`
	InvestmentType        valueTypes.Integer  `json:"investment_type"`
	InvestmentTypeDesc    valueTypes.String   `json:"investment_type_desc"`
	IsAgreeGdpr           valueTypes.Bool     `json:"is_agree_gdpr"`
	IsGdpr                valueTypes.Bool     `json:"is_gdpr"`
	IsNewVersion          valueTypes.Bool     `json:"is_new_version"`
	IsOpenProtocol        valueTypes.Bool     `json:"is_open_protocol"`
	IsPsCreateUser        valueTypes.Bool     `json:"is_ps_create_user"`
	IsPsOwner             valueTypes.Bool     `json:"is_ps_owner"`
	IsReceiveNotice       valueTypes.Bool     `json:"is_receive_notice"`
	IsSharePosition       valueTypes.Bool     `json:"is_share_position"`
	IsValidMobileEmail    valueTypes.Bool     `json:"is_valid_mobile_email"`
	Latitude              valueTypes.Float    `json:"latitude"`
	Longitude             valueTypes.Float    `json:"longitude"`
	MapLatitude           valueTypes.Float    `json:"map_latitude"`
	MapLongitude          valueTypes.Float    `json:"map_longitude"`
	MlpeFlag              valueTypes.Bool     `json:"mlpe_flag"`
	MobileTel             valueTypes.String   `json:"moble_tel" PointId:"mobile_tel"`
	MobileTelBak          valueTypes.String   `json:"moble_tel_bak" PointId:"mobile_tel_bak"`
	ModuleModelId         interface{}         `json:"module_model_id"`
	ModuleModelName       valueTypes.String   `json:"module_model_name"`
	Nmi                   valueTypes.String   `json:"nmi"`
	OperationBusName      valueTypes.String   `json:"operation_bus_name"`
	OrgIndexCode          []valueTypes.String `json:"org_index_code"`
	OwnerContact          valueTypes.String   `json:"owner_contact"`
	ParamIncomeUnit       valueTypes.Integer  `json:"param_income_unit"`
	ParamIncomeUnitName   valueTypes.String   `json:"param_income_unit_name"`
	ProvinceCode          valueTypes.String   `json:"province_code"`
	ProvinceName          valueTypes.String   `json:"province_name"`
	PsBuildDate           valueTypes.DateTime `json:"ps_build_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	PsCountryId           valueTypes.Integer  `json:"ps_country_id"`
	PsCreateUserId        valueTypes.Integer  `json:"ps_create_user_id"`
	PsCurrentTimeZone     valueTypes.String   `json:"ps_current_time_zone"`
	PsHolder              valueTypes.String   `json:"ps_holder"`
	PsId                  valueTypes.PsId     `json:"ps_id"`
	PsInstalledPower      valueTypes.Float    `json:"ps_installed_power" PointUnit:"W"`
	PsKey                 valueTypes.PsKey    `json:"ps_key"`
	PsLocation            valueTypes.String   `json:"ps_location"`
	PsName                valueTypes.String   `json:"ps_name"`
	PsPrice               valueTypes.Float    `json:"ps_price" PointUnitFrom:"ParamIncomeUnitName"`
	PsPriceKwh            valueTypes.Float    `json:"ps_price_kwh" PointUnitFrom:"ParamIncomeUnitName"`
	PsType                valueTypes.Integer  `json:"ps_type"`
	PsTypeDesc            valueTypes.String   `json:"ps_type_desc"`
	PsTypeName            valueTypes.String   `json:"ps_type_name"`
	PsUserId              valueTypes.Integer  `json:"ps_user_id"`
	RecordCreateTime      valueTypes.DateTime `json:"recore_create_time" PointId:"record_create_time" PointNameDateFormat:"2006/01/02 15:04:05"`
	SafeStartDate         valueTypes.DateTime `json:"safe_start_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	SetUserOrg            valueTypes.Integer  `json:"set_user_org"`
	ShareType             valueTypes.Integer  `json:"share_type"`
	ShareUserType         interface{}         `json:"share_user_type"`
	ShippingAddress       valueTypes.String   `json:"shipping_address"`
	ShippingZipCode       valueTypes.String   `json:"shipping_zip_code"`
	Sn                    valueTypes.String   `json:"sn" PointName:"Serial Number"`
	SummerTimeState       valueTypes.Bool     `json:"summer_time_state"`
	SummerTimeZone        valueTypes.String   `json:"summer_time_zone"`
	SummerTimeZoneId      valueTypes.Integer  `json:"summer_time_zone_id"`
	TimeZoneId            valueTypes.Integer  `json:"time_zone_id"`
	Timezone              valueTypes.String   `json:"timezone"`
	UserAccount           valueTypes.String   `json:"user_account"`
	UserCapacityStatus    valueTypes.Integer  `json:"user_capacity_status"`
	UserEnglishName       valueTypes.String   `json:"user_english_name"`
	UserLanguage          valueTypes.String   `json:"user_language"`
	UserMobileTel         valueTypes.String   `json:"user_moble_tel" PointId:"user_mobile_tel"`
	UserName              valueTypes.String   `json:"user_name"`
	UserTelNationCode     valueTypes.String   `json:"user_tel_nation_code"`
	ValidFlag             valueTypes.Bool     `json:"valid_flag"`
	WgsLatitude           valueTypes.Float    `json:"wgs_latitude"`
	WgsLongitude          valueTypes.Float    `json:"wgs_longitude"`
	ZipCode               valueTypes.String   `json:"zip_code"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String() + "." + e.Response.ResultData.PsKey.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
