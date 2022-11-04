package getPowerStationInfo

import (
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
	PsList []struct {
		BatteryPlateArea      interface{}        `json:"battery_plate_area"`
		BatteryType           valueTypes.Integer `json:"battery_type"`
		ConnectGrid           interface{}        `json:"connect_grid"`
		ContactMobile         interface{}        `json:"contact_mobile"`
		ContactMobileBak      interface{}        `json:"contact_mobile_bak"`
		ContactName           valueTypes.String  `json:"contact_name"`
		ContactTel            interface{}        `json:"contact_tel"`
		CreateTime            valueTypes.String  `json:"create_time"`
		CreateUserID          valueTypes.Integer `json:"create_user_id"`
		DesignCapacity        valueTypes.Integer `json:"design_capacity"`
		DesignCapacityBattery valueTypes.Integer `json:"design_capacity_battery"`
		EquivalentHour        valueTypes.Integer `json:"equivalent_hour"`
		GetCostCycle          valueTypes.Integer `json:"get_cost_cycle"`
		InverterLoadSum       interface{}        `json:"inverter_load_sum"`
		IsNewVersion          valueTypes.String  `json:"isNewVersion"`
		IsAgreeGdpr           valueTypes.String  `json:"is_agree_gdpr"`
		IsGdpr                valueTypes.String  `json:"is_gdpr"`
		IsOpenProtocol        valueTypes.String  `json:"is_open_protocol"`
		IsReceiveNotice       valueTypes.String  `json:"is_receive_notice"`
		IsSharePosition       valueTypes.String  `json:"is_share_position"`
		MobleTel              interface{}        `json:"moble_tel" PointId:"mobile_tel"`
		Monetary              interface{}        `json:"monetary"`
		ParamCo2              float64            `json:"param_co2"`
		ParamConverRate       valueTypes.Integer `json:"param_conver_rate"`
		ParamIncome           valueTypes.Integer `json:"param_income"`
		ParamIncomeUnit       valueTypes.Integer `json:"param_income_unit"`
		ParamTemperature      interface{}        `json:"param_temperature"`
		ParamTree             valueTypes.Integer `json:"param_tree"`
		PrAddition            valueTypes.Integer `json:"pr_addition"`
		PrMax                 float64            `json:"pr_max"`
		PrMin                 valueTypes.Integer `json:"pr_min"`
		PrRatio               valueTypes.Integer `json:"pr_ratio"`
		PsEmail               valueTypes.String  `json:"ps_email"`
		PsID                  valueTypes.Integer `json:"ps_id"`
		PsType                valueTypes.Integer `json:"ps_type"`
		PwCost                float64            `json:"pw_cost"`
		RadiationMax          valueTypes.Integer `json:"radiation_max"`
		UserAccount           valueTypes.String  `json:"user_account"`
		UserEmail             valueTypes.String  `json:"user_email"`
		UserName              valueTypes.String  `json:"user_name"`
	} `json:"psList"`
	PsMap struct {
		AccessType        interface{}        `json:"access_type"`
		AreaType          interface{}        `json:"area_type"`
		Areaid            interface{}        `json:"areaid"`
		BuildDate         valueTypes.String  `json:"build_date"`
		Buildstatus       valueTypes.Integer `json:"buildstatus"`
		Capitaltype       valueTypes.Integer `json:"capitaltype"`
		City              interface{}        `json:"city"`
		Country           interface{}        `json:"country"`
		CountyCode        interface{}        `json:"county_code"`
		DistrictFlag      valueTypes.Integer `json:"district_flag"`
		DivisionCode      interface{}        `json:"division_code"`
		Email             valueTypes.String  `json:"email"`
		ExpectInstallDate valueTypes.String  `json:"expect_install_date"`
		FaultSendType     interface{}        `json:"fault_send_type"`
		GcjLatitude       valueTypes.String  `json:"gcj_latitude"`
		GcjLongitude      valueTypes.String  `json:"gcj_longitude"`
		GprsLatitude      interface{}        `json:"gprs_latitude"`
		GprsLongitude     interface{}        `json:"gprs_longitude"`
		Installdate       valueTypes.String  `json:"installdate"`
		InvestmentType    valueTypes.Integer `json:"investment_type"`
		Latitude          valueTypes.String  `json:"latitude"`
		Longitude         valueTypes.String  `json:"longitude"`
		MapLatitude       valueTypes.String  `json:"map_latitude"`
		MapLongitude      valueTypes.String  `json:"map_longitude"`
		Montorurl         interface{}        `json:"montorurl"`
		Name              interface{}        `json:"name"`
		NameCode          interface{}        `json:"name_code"`
		Nation            interface{}        `json:"nation"`
		NationCode        interface{}        `json:"nation_code"`
		Nmi               valueTypes.String  `json:"nmi"`
		OperateYear       interface{}        `json:"operate_year"`
		Operationbusname  interface{}        `json:"operationbusname"`
		OrgIndexCode      interface{}        `json:"org_index_code"`
		OrgIndexCodeName  valueTypes.String  `json:"org_index_code_name"`
		OrganizationID    interface{}        `json:"organization_id"`
		OrganizationName  interface{}        `json:"organization_name"`
		PanoramaLevel     interface{}        `json:"panorama_level"`
		Producer          interface{}        `json:"producer"`
		Prov              interface{}        `json:"prov"`
		ProvCode          interface{}        `json:"prov_code"`
		PsBuildDate       valueTypes.String  `json:"ps_build_date"`
		PsCountryID       valueTypes.Integer `json:"ps_country_id"`
		Pscode            valueTypes.String  `json:"pscode"`
		Psdesc            interface{}        `json:"psdesc"`
		Psguid            valueTypes.String  `json:"psguid"`
		Psholder          valueTypes.String  `json:"psholder"`
		Psid              valueTypes.Integer `json:"psid"`
		Pslocation        valueTypes.String  `json:"pslocation"`
		Psname            valueTypes.String  `json:"psname"`
		Psnameenus        interface{}        `json:"psnameenus"`
		Psorgid           interface{}        `json:"psorgid"`
		Psorgname         interface{}        `json:"psorgname"`
		Pstype            valueTypes.Integer `json:"pstype"`
		SafeStartDate     valueTypes.String  `json:"safe_start_date"`
		Schedulingtype    valueTypes.Integer `json:"schedulingtype"`
		Shortname         valueTypes.String  `json:"shortname"`
		Sn                valueTypes.String  `json:"sn"`
		Street            interface{}        `json:"street"`
		Sysscheme         valueTypes.Integer `json:"sysscheme"`
		Timezone          valueTypes.Integer `json:"timezone"`
		TownCode          interface{}        `json:"town_code"`
		Uploadprotocol    valueTypes.Integer `json:"uploadprotocol"`
		UserAccount       valueTypes.String  `json:"user_account"`
		Validflag         valueTypes.Integer `json:"validflag"`
		Videopath         interface{}        `json:"videopath"`
		VillageCode       interface{}        `json:"village_code"`
		WgsLatitude       float64            `json:"wgs_latitude"`
		WgsLongitude      float64            `json:"wgs_longitude"`
	} `json:"psMap"`
	RemindType interface{} `json:"remindType"`
	SnInfoList []struct {
		ChnnlDesc    interface{}        `json:"chnnl_desc" PointId:"channel_description"`
		ChnnlID      valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
		ChnnlName    valueTypes.String  `json:"chnnl_name PointId:"channel_name"`
		CrtDate      valueTypes.String  `json:"crt_date" PointId:"create_date"`
		CrtUserName  interface{}        `json:"crt_user_name" PointId:"create_username"`
		DataFlag     valueTypes.Integer `json:"data_flag"`
		FlagServer   interface{}        `json:"flag_server"`
		HostIP       interface{}        `json:"host_ip"`
		ID           valueTypes.Integer `json:"id"`
		IsEnable     valueTypes.Integer `json:"is_enable"`
		ProtocolType interface{}        `json:"protocol_type"`
		PsGUID       interface{}        `json:"ps_guid"`
		PsID         valueTypes.Integer `json:"ps_id"`
		Secrit       valueTypes.String  `json:"secrit"`
		Sn           valueTypes.String  `json:"sn"`
		TcpMode      interface{}        `json:"tcp_mode"`
		TcpPort      interface{}        `json:"tcp_port"`
	} `json:"snInfoList"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
