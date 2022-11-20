package getPowerStationInfo

import (
	"GoSungrow/iSolarCloud/Common"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationInfoForBackSys"
const Disabled = false

type RequestData struct {
	PsId       valueTypes.PsId     `json:"psId" require:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	RemindType Common.Unknown `json:"remindType" PointId:"remind_type"`

	PsList []struct {
		BatteryPlateArea      Common.Unknown     `json:"battery_plate_area"`
		BatteryType           valueTypes.Integer `json:"battery_type"`
		ConnectGrid           Common.Unknown     `json:"connect_grid"`
		ContactMobile         Common.Unknown     `json:"contact_mobile"`
		ContactMobileBak      Common.Unknown     `json:"contact_mobile_bak"`
		ContactName           valueTypes.String  `json:"contact_name"`
		ContactTel            Common.Unknown     `json:"contact_tel"`
		CreateTime            valueTypes.String  `json:"create_time"`
		CreateUserID          valueTypes.Integer `json:"create_user_id"`
		DesignCapacity        valueTypes.Float   `json:"design_capacity"`
		DesignCapacityBattery valueTypes.Float   `json:"design_capacity_battery"`
		EquivalentHour        valueTypes.Integer `json:"equivalent_hour"`
		GetCostCycle          valueTypes.Float   `json:"get_cost_cycle"`
		InverterLoadSum       Common.Unknown     `json:"inverter_load_sum"`
		IsNewVersion          valueTypes.String  `json:"isNewVersion"`
		IsAgreeGdpr           valueTypes.String  `json:"is_agree_gdpr"`
		IsGdpr                valueTypes.String  `json:"is_gdpr"`
		IsOpenProtocol        valueTypes.String  `json:"is_open_protocol"`
		IsReceiveNotice       valueTypes.String  `json:"is_receive_notice"`
		IsSharePosition       valueTypes.String  `json:"is_share_position"`
		MobileTel             Common.Unknown     `json:"moble_tel" PointId:"mobile_tel"`
		Monetary              Common.Unknown     `json:"monetary"`
		ParamCo2              valueTypes.Float   `json:"param_co2"`
		ParamConvertRate      valueTypes.Float   `json:"param_conver_rate" PointId:"param_convert_rate"`
		ParamIncome           valueTypes.Float   `json:"param_income"`
		ParamIncomeUnit       valueTypes.Integer `json:"param_income_unit"`
		ParamTemperature      Common.Unknown     `json:"param_temperature"`
		ParamTree             valueTypes.Float   `json:"param_tree"`
		PrAddition            valueTypes.Float   `json:"pr_addition"`
		PrMax                 valueTypes.Float   `json:"pr_max"`
		PrMin                 valueTypes.Float   `json:"pr_min"`
		PrRatio               valueTypes.Float   `json:"pr_ratio"`
		PsEmail               valueTypes.String  `json:"ps_email"`
		PsId                  valueTypes.PsId    `json:"ps_id"`
		PsType                valueTypes.Integer `json:"ps_type"`
		PwCost                valueTypes.Float   `json:"pw_cost"`
		RadiationMax          valueTypes.Float   `json:"radiation_max"`
		UserAccount           valueTypes.String  `json:"user_account"`
		UserEmail             valueTypes.String  `json:"user_email"`
		UserName              valueTypes.String  `json:"user_name"`
	} `json:"psList" PointId:"ps_list" PointIdFromChild:"PsId" PointIdReplace:"true"`

	PsMap struct {
		AccessType        Common.Unknown      `json:"access_type"`
		AreaType          Common.Unknown      `json:"area_type"`
		AreaId            Common.Unknown      `json:"areaid" PointId:"area_id"`
		BuildDate         valueTypes.DateTime `json:"build_date"`
		BuildStatus       valueTypes.Integer  `json:"buildstatus" PointId:"build_status"`
		CapitalType       valueTypes.Integer  `json:"capitaltype" PointId:"capital_type"`
		City              Common.Unknown      `json:"city"`
		Country           Common.Unknown      `json:"country"`
		CountyCode        Common.Unknown      `json:"county_code"`
		DistrictFlag      valueTypes.Integer  `json:"district_flag"`
		DivisionCode      Common.Unknown      `json:"division_code"`
		Email             valueTypes.String   `json:"email"`
		ExpectInstallDate valueTypes.DateTime `json:"expect_install_date"`
		FaultSendType     Common.Unknown      `json:"fault_send_type"`
		GcjLatitude       valueTypes.Float    `json:"gcj_latitude"`
		GcjLongitude      valueTypes.Float    `json:"gcj_longitude"`
		GprsLatitude      valueTypes.Float    `json:"gprs_latitude"`
		GprsLongitude     valueTypes.Float    `json:"gprs_longitude"`
		InstallDate       valueTypes.DateTime `json:"installdate"`
		InvestmentType    valueTypes.Integer  `json:"investment_type"`
		Latitude          valueTypes.Float    `json:"latitude"`
		Longitude         valueTypes.Float    `json:"longitude"`
		MapLatitude       valueTypes.Float    `json:"map_latitude"`
		MapLongitude      valueTypes.Float    `json:"map_longitude"`
		MonitorUrl        Common.Unknown      `json:"montorurl" PointId:"monitor_url"`
		Name              Common.Unknown      `json:"name"`
		NameCode          Common.Unknown      `json:"name_code"`
		Nation            Common.Unknown      `json:"nation"`
		NationCode        Common.Unknown      `json:"nation_code"`
		Nmi               valueTypes.String   `json:"nmi"`
		OperateYear       Common.Unknown      `json:"operate_year"`
		OperationBusName  Common.Unknown      `json:"operationbusname" PointId:"operation_bus_name"`
		OrgIndexCode      Common.Unknown      `json:"org_index_code"`
		OrgIndexCodeName  valueTypes.String   `json:"org_index_code_name"`
		OrganizationID    Common.Unknown      `json:"organization_id"`
		OrganizationName  Common.Unknown      `json:"organization_name"`
		PanoramaLevel     Common.Unknown      `json:"panorama_level"`
		Producer          Common.Unknown      `json:"producer"`
		Prov              Common.Unknown      `json:"prov"`
		ProvCode          Common.Unknown      `json:"prov_code"`
		PsBuildDate       valueTypes.DateTime `json:"ps_build_date"`
		PsCountryID       valueTypes.Integer  `json:"ps_country_id"`
		PsCode            valueTypes.String   `json:"pscode" PointId:"ps_code"`
		PsDesc            Common.Unknown      `json:"psdesc" PointId:"ps_desc"`
		PsGuid            valueTypes.String   `json:"psguid" PointId:"ps_guid"`
		PsHolder          valueTypes.String   `json:"psholder" PointId:"ps_holder"`
		PsId              valueTypes.Integer  `json:"psid" PointId:"ps_id"`
		PsLocation        valueTypes.String   `json:"pslocation" PointId:"remind_type"`
		PsName            valueTypes.String   `json:"psname" PointId:"remind_type"`
		PsNameEnus        Common.Unknown      `json:"psnameenus" PointId:"ps_name_enus"`
		PsOrgId           Common.Unknown      `json:"psorgid" PointId:"ps_org_id"`
		PsOrgName         Common.Unknown      `json:"psorgname" PointId:"ps_org_name"`
		PsType            valueTypes.Integer  `json:"pstype" PointId:"ps_type"`
		SafeStartDate     valueTypes.DateTime `json:"safe_start_date"`
		SchedulingType    valueTypes.Integer  `json:"schedulingtype" PointId:"scheduling_type"`
		Shortname         valueTypes.String   `json:"shortname"`
		Sn                valueTypes.String   `json:"sn"`
		Street            Common.Unknown      `json:"street"`
		SysScheme         valueTypes.Integer  `json:"sysscheme" PointId:"sys_scheme"`
		Timezone          valueTypes.Integer  `json:"timezone"`
		TownCode          Common.Unknown      `json:"town_code"`
		UploadProtocol    valueTypes.Integer  `json:"uploadprotocol" PointId:"upload_protocol"`
		UserAccount       valueTypes.String   `json:"user_account"`
		ValidFlag         valueTypes.Bool     `json:"validflag" PointId:"valid_flag"`
		VideoPath         Common.Unknown      `json:"videopath" PointId:"video_path"`
		VillageCode       Common.Unknown      `json:"village_code"`
		WgsLatitude       valueTypes.Float    `json:"wgs_latitude"`
		WgsLongitude      valueTypes.Float    `json:"wgs_longitude"`
	} `json:"psMap" PointId:"ps_map" PointIdFromChild:"PsId" PointIdReplace:"true"`

	SnInfoList []struct {
		ChannelDesc  Common.Unknown      `json:"chnnl_desc" PointId:"channel_description"`
		ChannelId    valueTypes.Integer  `json:"chnnl_id" PointId:"channel_id"`
		ChannelName  valueTypes.String   `json:"chnnl_name" PointId:"channel_name"`
		CrtDate      valueTypes.DateTime `json:"crt_date" PointId:"create_date"`
		CrtUserName  Common.Unknown      `json:"crt_user_name" PointId:"create_username"`
		DataFlag     valueTypes.Integer  `json:"data_flag"`
		FlagServer   Common.Unknown      `json:"flag_server"`
		HostIP       Common.Unknown      `json:"host_ip"`
		ID           valueTypes.Integer  `json:"id"`
		IsEnable     valueTypes.Bool     `json:"is_enable"`
		ProtocolType Common.Unknown      `json:"protocol_type"`
		PsGUID       Common.Unknown      `json:"ps_guid"`
		PsId         valueTypes.PsId     `json:"ps_id"`
		Secret       valueTypes.String   `json:"secrit" PointId:"secret"`
		Sn           valueTypes.String   `json:"sn"`
		TcpMode      Common.Unknown      `json:"tcp_mode"`
		TcpPort      Common.Unknown      `json:"tcp_port"`
	} `json:"snInfoList" PointId:"sn_info_list" PointIdFromChild:"PsId" PointIdReplace:"true"`
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
