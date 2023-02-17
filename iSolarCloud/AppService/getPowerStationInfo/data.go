package getPowerStationInfo

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/Common"
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationInfo"
const Disabled = false
const EndPointName = "AppService.getPowerStationInfo"

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
	ParamIncome valueTypes.Float `json:"param_income"`
	RemindType  Common.Unknown   `json:"remindType" PointId:"remind_type"`

	PowerStationMap struct {
		AreaId            Common.Unknown      `json:"area_id"`
		ConnectType       valueTypes.Integer  `json:"connect_type"`
		County            Common.Unknown      `json:"county"`
		CountyCode        Common.Unknown      `json:"county_code"`
		DesignCapacity2   valueTypes.Float    `json:"designCapacity" PointIgnore:"true"`
		DesignCapacity    valueTypes.Float    `json:"design_capacity" PointUnit:"W"`
		Email             valueTypes.String   `json:"email"`
		EnergyScheme      Common.Unknown      `json:"energy_scheme"`
		ExpectInstallDate valueTypes.DateTime `json:"expect_install_date" PointNameDateFormat:"DateTimeLayout"`
		FaultSendType     Common.Unknown      `json:"fault_send_type"`
		GcjLatitude       valueTypes.Float    `json:"gcj_latitude"`
		GcjLongitude      valueTypes.Float    `json:"gcj_longitude"`
		Latitude          valueTypes.Float    `json:"latitude"`
		Longitude         valueTypes.Float    `json:"longitude"`
		Name              Common.Unknown      `json:"name"`
		NameCode          Common.Unknown      `json:"name_code"`
		Nation            Common.Unknown      `json:"nation"`
		NationCode        Common.Unknown      `json:"nation_code"`
		ParamIncome       valueTypes.Float    `json:"param_income"`
		Prov              Common.Unknown      `json:"prov"`
		ProvCode          Common.Unknown      `json:"prov_code"`
		PsBuildDate       valueTypes.DateTime `json:"ps_build_date" PointNameDateFormat:"DateTimeLayout"`
		PsCountryId       valueTypes.Integer  `json:"ps_country_id"`
		PsDesc            Common.Unknown      `json:"ps_desc"`
		PsHolder          valueTypes.String   `json:"ps_holder"`
		PsId              valueTypes.PsId     `json:"ps_id"`
		PsLocation        valueTypes.String   `json:"ps_location"`
		PsName            valueTypes.String   `json:"ps_name"`
		PsType            valueTypes.Integer  `json:"ps_type"`
		RecordCreateTime  valueTypes.DateTime `json:"recore_create_time" PointId:"record_create_time" PointNameDateFormat:"DateTimeLayout"`
		ReportType        Common.Unknown      `json:"report_type"`
		ShippingAddress   valueTypes.String   `json:"shipping_address"`
		ShippingZipCode   valueTypes.String   `json:"shipping_zip_code"`
		TimeZoneId        valueTypes.Integer  `json:"time_zone_id"`
		ValidFlag         valueTypes.Bool     `json:"valid_flag"`
		Village           Common.Unknown      `json:"village"`
		VillageCode       Common.Unknown      `json:"village_code"`
		WgsLatitude       valueTypes.Float    `json:"wgs_latitude"`
		WgsLongitude      valueTypes.Float    `json:"wgs_longitude"`
		ZipCode           valueTypes.String   `json:"zip_code"`
	} `json:"powerStationMap" PointId:"power_station_map"`

	InstallProviderInfo struct {
		Installer      valueTypes.String  `json:"installer"`
		InstallerEmail valueTypes.String  `json:"installer_email"`
		InstallerPhone valueTypes.String  `json:"installer_phone"`
		OrgURL         Common.Unknown     `json:"org_url"`
		UpOrgId        valueTypes.Integer `json:"up_org_id"`
	} `json:"installProviderInfo" PointId:"install_provider_info"`

	PowerChargeDataMap struct {
		CodeType            valueTypes.Integer  `json:"code_type"`
		DefaultCharge       valueTypes.Float    `json:"default_charge"`
		ElectricChargeId    valueTypes.Integer  `json:"electric_charge_id"`
		EndTime             valueTypes.DateTime `json:"end_time" PointNameDateFormat:"DateTimeLayout"`
		IntervalTimeCharge  Common.Unknown      `json:"interval_time_charge"`
		ParamIncomeUnitName valueTypes.String   `json:"param_income_unit_name"`
		StartTime           valueTypes.DateTime `json:"start_time" PointNameDateFormat:"DateTimeLayout"`
	} `json:"powerChargeDataMap" PointId:"power_charge_data_map"`

	SysTimeZones []struct {
		Id           valueTypes.Integer `json:"id"`
		TimezoneName valueTypes.String  `json:"timezone_name"`
		TimezoneUtc  valueTypes.String  `json:"timezone_utc"`
	} `json:"sysTimeZones" PointId:"sys_time_zones" PointArrayFlatten:"false" DataTable:"true"`

	PowerPictureList []struct {
		PicID       valueTypes.Integer `json:"pic_id"`
		PicLanguage valueTypes.Integer `json:"pic_language"`
		PicType     valueTypes.Integer `json:"pic_type"`
		PictureName valueTypes.String  `json:"picture_name"`
		PictureURL  valueTypes.String  `json:"picture_url"`
		PsUnitUUID  interface{}        `json:"ps_unit_uuid"`
	} `json:"powerPictureList" PointId:"power_picture_list" PointArrayFlatten:"false"`

	SendReportConfigList []valueTypes.String `json:"sendReportConfigList" PointId:"send_report_config_list"`

	CurrencyTypeList []struct {
		CodeName  valueTypes.String  `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"currencyTypeList" PointId:"currency_type_list" PointArrayFlatten:"false" DataTable:"true"`

	ParallelTypes []struct {
		CodeName  valueTypes.String  `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"parallelTypes" PointId:"parallel_types" PointArrayFlatten:"false" DataTable:"true"`

	StatusList []struct {
		CodeName  valueTypes.String  `json:"code_name"`
		CodeValue valueTypes.Integer `json:"code_value"`
	} `json:"statusList" PointId:"status_list" PointArrayFlatten:"false" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
