package getPsDetail

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"time"
)

const Url = "/v1/powerStationService/getPsDetail"
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
	ActualEnergy          []string      `json:"actual_energy"`
	ActualEnergyUnit      string        `json:"actual_energy_unit"`
	AlarmCount            api.Integer   `json:"alarm_count"`
	AreaID                interface{}   `json:"area_id"`
	AreaType              api.Integer   `json:"area_type"`
	BuildDate             api.DateTime  `json:"build_date"`
	Co2Reduce             api.UnitValue `json:"co2_reduce"`
	Co2ReduceTotal        api.UnitValue `json:"co2_reduce_total"   PointId:"co2_reduce_total" PointType:"PointTypeTotal"`
	CoalReduce            api.UnitValue `json:"coal_reduce"`
	CoalReduceTotal       api.UnitValue `json:"coal_reduce_total"   PointId:"coal_reduce_total" PointType:"PointTypeTotal"`
	ConnectGrid           string        `json:"connect_grid"`
	ConnectType           api.Integer   `json:"connect_type"`
	ContactPerson         string        `json:"contact_person"`
	CurrPower             api.UnitValue `json:"curr_power"` // Pv Power
	DataLastUpdateTime    string        `json:"data_last_update_time"`
	DayEqHours            string        `json:"day_eq_hours"`
	Description           interface{}   `json:"description"`
	DesignCapacity        api.UnitValue `json:"design_capacity"`
	DesignCapacityBattery api.UnitValue `json:"design_capacity_battery"`
	DiagramURL            string        `json:"diagram_url"`
	EnergyScheme          interface{}   `json:"energy_scheme"`
	ExpectInstallDate     api.DateTime  `json:"expect_install_date"`
	FaultCount            api.Integer   `json:"fault_count"`
	FaultSendType         string        `json:"fault_send_type"`
	GcjLatitude           api.Float     `json:"gcj_latitude"`
	GcjLongitude          api.Float     `json:"gcj_longitude"`
	GprsLatitude          api.Float     `json:"gprs_latitude"`
	GprsLongitude         api.Float     `json:"gprs_longitude"`
	HasAmmeter            api.Integer   `json:"has_ammeter"`
	Images                []struct {
		FileID      api.Integer `json:"file_id"`
		ID          api.Integer `json:"id"`
		PicLanguage api.Integer `json:"pic_language"`
		PicType     api.Integer `json:"pic_type"`
		PictureName string      `json:"picture_name"`
		PictureURL  string      `json:"picture_url"`
		PsUnitUUID  interface{} `json:"ps_unit_uuid"`
	} `json:"images"`
	InstallDate            api.DateTime  `json:"install_date"`
	InstallerPsFaultStatus string        `json:"installer_ps_fault_status"`
	IsHaveEsInverter       api.Integer   `json:"is_have_es_inverter"`
	IsTransformSystem      string        `json:"is_transform_system"`
	IsTuv                  api.Integer   `json:"is_tuv"`
	Latitude               api.Float     `json:"latitude"`
	Longitude              api.Float     `json:"longitude"`
	MapLatitude            api.Float     `json:"map_latitude"`
	MapLongitude           api.Float     `json:"map_longitude"`
	MeterReduce            api.UnitValue `json:"meter_reduce"`
	MeterReduceTotal       api.UnitValue `json:"meter_reduce_total"   PointId:"meter_reduce_total" PointType:"PointTypeTotal"`
	MobileTel              string        `json:"mibile_tel"   PointId:"mobile_tel"`
	MonthPr                string        `json:"monthPr"   PointId:"monthPr" PointType:"PointTypeMonth"`
	MonthEnergy            api.UnitValue `json:"month_energy"   PointId:"month_energy" PointType:"PointTypeMonth"`
	MonthEnergyVirgin      api.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthEqHours           string        `json:"month_eq_hours"   PointId:"month_eq_hours" PointType:"PointTypeMonth"`
	MonthIncome            api.UnitValue `json:"month_income"   PointId:"month_income" PointType:"PointTypeMonth"`
	NoxReduce              api.UnitValue `json:"nox_reduce"`
	NoxReduceTotal         api.UnitValue `json:"nox_reduce_total"   PointId:"nox_reduce_total" PointType:"PointTypeTotal"`
	OperateYear            string        `json:"operate_year"`
	OperationBusName       string        `json:"operation_bus_name"`
	OwnerPsFaultStatus     string        `json:"owner_ps_fault_status"`
	P83012Unit             string        `json:"p83012_unit"   PointId:"P83012Unit" PointType:""  PointIgnore:"true"`
	P83012Value            string        `json:"p83012_value"   PointId:"P83012Value" PointType:""  PointIgnore:"true"`
	P83013Unit             string        `json:"p83013_unit"   PointId:"P83013Unit" PointType:""  PointIgnore:"true"`
	P83013Value            string        `json:"p83013_value"   PointId:"P83013Value" PointType:""  PointIgnore:"true"`
	P83036Unit             string        `json:"p83036_unit"   PointId:"P83036Unit" PointType:""  PointIgnore:"true"`
	P83036Value            string        `json:"p83036_value"   PointId:"P83036Value" PointType:""  PointIgnore:"true"`
	P83016                 string        `json:"p83016"   PointId:"P83016" PointType:""  PointIgnore:"true"`
	P83016Unit             string        `json:"p83016_unit"   PointId:"P83016Unit" PointType:""  PointIgnore:"true"`
	P83017                 string        `json:"p83017"   PointId:"P83017" PointType:""  PointIgnore:"true"`
	P83017Unit             string        `json:"p83017_unit"   PointId:"P83017Unit" PointType:""  PointIgnore:"true"`
	P83023                 string        `json:"p83023"   PointId:"P83023" PointType:""`
	P83023y                string        `json:"p83023y"   PointId:"P83023y" PointType:""`
	P83023year             string        `json:"p83023year"   PointId:"P83023year" PointType:""`
	P83023ym               string        `json:"p83023ym"   PointId:"P83023ym" PointType:""`
	P83043                 string        `json:"p83043"   PointId:"P83043" PointType:""`
	P83044                 string        `json:"p83044"   PointId:"P83044" PointType:""`
	P83045                 string        `json:"p83045"   PointId:"P83045" PointType:""`
	P83072Map              api.UnitValue `json:"p83072_map"   PointId:"p83072" PointType:""`
	P83072MapVirgin        api.UnitValue `json:"p83072_map_virgin"  PointIgnore:"true"`
	P83073Map              api.UnitValue `json:"p83073_map"   PointId:"p83073" PointType:""`
	P83073MapVirgin        api.UnitValue `json:"p83073_map_virgin"  PointIgnore:"true"`
	P83074Map              api.UnitValue `json:"p83074_map"   PointId:"p83074" PointType:""`
	P83074MapVirgin        api.UnitValue `json:"p83074_map_virgin"  PointIgnore:"true"`
	P83075Map              api.UnitValue `json:"p83075_map"   PointId:"p83075" PointType:""`
	P83075MapVirgin        api.UnitValue `json:"p83075_map_virgin"  PointIgnore:"true"`
	P83076Map              api.UnitValue `json:"p83076_map"   PointId:"p83076" PointType:""`        // Pv Power
	P83076MapVirgin        api.UnitValue `json:"p83076_map_virgin"  PointIgnore:"true"` // Pv Power
	P83077Map              api.UnitValue `json:"p83077_map"   PointId:"p83077" PointType:""`        // Pv Energy
	P83077MapVirgin        api.UnitValue `json:"p83077_map_virgin"  PointIgnore:"true"` // Pv Energy
	P83078Map              api.UnitValue `json:"p83078_map"   PointId:"p83078" PointType:""`
	P83078MapVirgin        api.UnitValue `json:"p83078_map_virgin"  PointIgnore:"true"`
	P83079Map              api.UnitValue `json:"p83079_map"   PointId:"p83079" PointType:""`
	P83079MapVirgin        api.UnitValue `json:"p83079_map_virgin"  PointIgnore:"true"`
	P83080Map              api.UnitValue `json:"p83080_map"   PointId:"p83080" PointType:""`
	P83080MapVirgin        api.UnitValue `json:"p83080_map_virgin"  PointIgnore:"true"`
	P83088Map              api.UnitValue `json:"p83088_map"   PointId:"p83088" PointType:""`        // Es Energy
	P83088MapVirgin        api.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"` // Es Energy
	P83089Map              api.UnitValue `json:"p83089_map"   PointId:"p83089" PointType:""`        // Es Discharge Energy
	P83089MapVirgin        api.UnitValue `json:"p83089_map_virgin"  PointIgnore:"true"` // Es Discharge Energy
	P83094Map              api.UnitValue `json:"p83094_map"   PointId:"p83094" PointType:""`
	P83094MapVirgin        api.UnitValue `json:"p83094_map_virgin"  PointIgnore:"true"`
	P83095Map              api.UnitValue `json:"p83095_map"   PointId:"p83095" PointType:"PointTypeTotal"`        // Es Total Discharge Energy
	P83095MapVirgin        api.UnitValue `json:"p83095_map_virgin"  PointIgnore:"true"` // Es Total Discharge Energy
	P83097Map              api.UnitValue `json:"p83097_map"   PointId:"p83097" PointType:""`
	P83097MapVirgin        api.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83100Map              api.UnitValue `json:"p83100_map"   PointId:"p83100" PointType:""`
	P83100MapVirgin        api.UnitValue `json:"p83100_map_virgin"  PointIgnore:"true"`
	P83101Map              api.UnitValue `json:"p83101_map"   PointId:"p83101" PointType:""`
	P83101MapVirgin        api.UnitValue `json:"p83101_map_virgin"  PointIgnore:"true"`
	P83102Map              api.UnitValue `json:"p83102_map"   PointId:"p83102" PointType:""`
	P83102MapVirgin        api.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent          string        `json:"p83102_percent"   PointId:"P83102Percent" PointType:""`
	P83105Map              api.UnitValue `json:"p83105_map"   PointId:"p83105" PointType:""`
	P83105MapVirgin        api.UnitValue `json:"p83105_map_virgin"  PointIgnore:"true"`
	P83106Map              api.UnitValue `json:"p83106_map"   PointId:"p83106" PointType:""`
	P83106MapVirgin        api.UnitValue `json:"p83106_map_virgin"  PointIgnore:"true"`
	P83107Map              api.UnitValue `json:"p83107_map"   PointId:"p83107" PointType:""`
	P83107MapVirgin        api.UnitValue `json:"p83107_map_virgin"  PointIgnore:"true"`
	P83118Map              api.UnitValue `json:"p83118_map"   PointId:"p83118" PointType:""`
	P83118MapVirgin        api.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map              api.UnitValue `json:"p83119_map"   PointId:"p83119" PointType:""`
	P83119MapVirgin        api.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map              api.UnitValue `json:"p83120_map"   PointId:"p83120" PointType:""`
	P83120MapVirgin        api.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121                 string        `json:"p83121"   PointId:"P83121" PointType:""`
	P83122                 string        `json:"p83122"   PointId:"P83122" PointType:""`
	P83123Map              api.UnitValue `json:"p83123_map"   PointId:"p83123" PointType:""`
	P83123MapVirgin        api.UnitValue `json:"p83123_map_virgin"  PointIgnore:"true"`
	P83124Map              api.UnitValue `json:"p83124_map"   PointId:"p83124" PointType:""`
	P83124MapVirgin        api.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83125                 string        `json:"p83125"   PointId:"P83125" PointType:""`
	P83126                 string        `json:"p83126"   PointId:"P83126" PointType:""`
	P83128MapVirgin        api.UnitValue `json:"p83128_map_virgin"  PointIgnore:"true"`
	P83202Map              api.UnitValue `json:"p83202_map"   PointId:"p83202" PointType:""`
	P83202MapVirgin        api.UnitValue `json:"p83202_map_virgin"  PointIgnore:"true"`
	PercentPlanYear        api.Float     `json:"percent_plan_year"`
	PlanEnergy             []string      `json:"plan_energy"`
	PlanEnergyUnit         string        `json:"plan_energy_unit"`
	PlanEnergyYear         api.UnitValue `json:"plan_energy_year"`
	PowderReduce           api.UnitValue `json:"powder_reduce"`
	PowderReduceTotal      api.UnitValue `json:"powder_reduce_total"   PointId:"powder_reduce_total" PointType:"PointTypeTotal"`
	PowerChargeSetted      api.Integer   `json:"power_charge_setted"   PointId:"power_charge_setted" PointType:"PointTypeTotal"`
	Producer               string        `json:"producer"`
	PsCountryID            api.Integer   `json:"ps_country_id"`
	PsFaultStatus          string        `json:"ps_fault_status"`
	PsHealthStatus         string        `json:"ps_health_status"`
	PsHolder               string        `json:"ps_holder"`
	PsLocation             string        `json:"ps_location"`
	PsName                 string        `json:"ps_name"`
	PsShortName            string        `json:"ps_short_name"`
	PsState                string        `json:"ps_state"`
	PsType                 api.Integer   `json:"ps_type"`
	PsTypeName             string        `json:"ps_type_name"`
	PsWindLevel            string        `json:"ps_wind_level"`
	PsWindPos              string        `json:"ps_wind_pos"`
	RecoreCreateTime       string        `json:"recore_create_time"`
	ReportType             string        `json:"report_type"`
	RobotNumSweepCapacity  struct {
		Num           api.Integer   `json:"num"`
		SweepCapacity api.Float `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`
	SafeStartDate                 api.DateTime  `json:"safe_start_date"`
	SelfConsumptionOffsetReminder api.Integer   `json:"self_consumption_offset_reminder"`
	ShippingAddress               string        `json:"shipping_address"`
	ShippingZipCode               string        `json:"shipping_zip_code"`
	So2Reduce                     api.UnitValue `json:"so2_reduce"`
	So2ReduceTotal                api.UnitValue `json:"so2_reduce_total"   PointId:"so2_reduce_total" PointType:"PointTypeTotal"`
	StorageInverterData           []struct {
		CommunicationDevSn      string        `json:"communication_dev_sn"`
		DevFaultStatus          api.Integer   `json:"dev_fault_status"`
		DevStatus               api.Integer   `json:"dev_status"`
		DeviceCode              api.Integer   `json:"device_code"`
		DeviceModelCode         string        `json:"device_model_code"`
		DeviceName              string        `json:"device_name"`
		DeviceState             string        `json:"device_state"`
		DeviceType              api.Integer   `json:"device_type"`
		DrmStatus               string        `json:"drm_status"`
		DrmStatusName           string        `json:"drm_status_name"`
		EnergyFlow              []string      `json:"energy_flow"`
		HasAmmeter              api.Integer   `json:"has_ammeter"`
		InstallerDevFaultStatus api.Integer   `json:"installer_dev_fault_status"`
		InverterSn              string        `json:"inverter_sn"`
		OwnerDevFaultStatus     api.Integer   `json:"owner_dev_fault_status"`
		P13003Map               api.UnitValue `json:"p13003_map"   PointId:"p13003" PointType:"PointTypeInstant"`
		P13003MapVirgin         api.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13011Map               api.UnitValue `json:"p13011_map"   PointId:"p13011" PointType:"PointTypeInstant"`
		P13011MapVirgin         api.UnitValue `json:"p13011_map_virgin"  PointIgnore:"true"`
		P13115Map               api.UnitValue `json:"p13115_map"   PointId:"p13115" PointType:"PointTypeInstant"`
		P13115MapVirgin         api.UnitValue `json:"p13115_map_virgin"  PointIgnore:"true"`
		P13119Map               api.UnitValue `json:"p13119_map"   PointId:"p13119" PointType:"PointTypeInstant"`
		P13119MapVirgin         api.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map               api.UnitValue `json:"p13121_map"   PointId:"p13121" PointType:"PointTypeInstant"`
		P13121MapVirgin         api.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map               api.UnitValue `json:"p13126_map"   PointId:"p13126" PointType:"PointTypeInstant"`
		P13126MapVirgin         api.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141                  string        `json:"p13141"   PointId:"P13141" PointType:"PointTypeInstant"`
		P13142                  string        `json:"p13142"   PointId:"P13142" PointType:"PointTypeInstant"`
		P13149Map               api.UnitValue `json:"p13149_map"   PointId:"p13149" PointType:"PointTypeInstant"`
		P13149MapVirgin         api.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map               api.UnitValue `json:"p13150_map"   PointId:"p13150" PointType:"PointTypeInstant"`
		P13150MapVirgin         api.UnitValue `json:"p13150_map_virgin"  PointIgnore:"true"`
		P13155                  string        `json:"p13155"   PointId:"P13155" PointType:"PointTypeInstant"`
		PsKey                   string        `json:"ps_key"`
		UpdateTime              string        `json:"update_time"`
		UUID                    api.Integer   `json:"uuid"`
	} `json:"storage_inverter_data"`
	SysScheme            api.Integer   `json:"sys_scheme"`
	TimeZoneID           api.Integer   `json:"time_zone_id"`
	Timezone             string        `json:"timezone"`
	TodayEnergy          api.UnitValue `json:"today_energy"   PointId:"today_energy" PointType:"PointTypeDaily"`
	TodayEnergyVirgin    api.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome          api.UnitValue `json:"today_income"   PointId:"today_income" PointType:"PointTypeDaily"`
	TotalEnergy          api.UnitValue `json:"total_energy"   PointId:"total_energy" PointType:"PointTypeTotal"`
	TotalEnergyVirgin    api.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalEnergyYear      api.UnitValue `json:"total_energy_year"   PointId:"total_energy_year" PointType:"PointTypeTotal"`
	TotalIncome          api.UnitValue `json:"total_income"   PointId:"total_income" PointType:"PointTypeTotal"`
	TreeReduce           api.UnitValue `json:"tree_reduce"`
	TreeReduceTotal      api.UnitValue `json:"tree_reduce_total"   PointId:"tree_reduce_total" PointType:"PointTypeTotal"`
	TuvLevel             string        `json:"tuv_level"`
	ValidFlag            string        `json:"valid_flag"`
	WaitAssignOrderCount api.Integer   `json:"wait_assign_order_count"`
	WaterReduce          api.UnitValue `json:"water_reduce"`
	WaterReduceTotal     api.UnitValue `json:"water_reduce_total"   PointId:"water_reduce_total" PointType:"PointTypeTotal"`
	WgsLatitude          api.Float     `json:"wgs_latitude"`
	WgsLongitude         api.Float     `json:"wgs_longitude"`
	Year                 string        `json:"year"`
	ZfzyMap              api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin        api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZipCode              string        `json:"zip_code"`
	ZjzzMap              api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin        api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
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
// }


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		name := fmt.Sprintf("getPsDetail.%s", e.Request.PsId.String())

		uv := api.SetUnitValueString(e.Response.ResultData.P83012Value, e.Response.ResultData.P83012Unit)
		entries.AddUnitValue(name + ".p83012", e.Request.PsId.String(), "p83012", "", api.NewDateTime(""), uv)

		uv = api.SetUnitValueString(e.Response.ResultData.P83013Value, e.Response.ResultData.P83013Unit)
		entries.AddUnitValue(name + ".p83013", e.Request.PsId.String(), "p83013", "", api.NewDateTime(""), uv)

		uv = api.SetUnitValueString(e.Response.ResultData.P83036Value, e.Response.ResultData.P83036Unit)
		entries.AddUnitValue(name + ".p83036", e.Request.PsId.String(), "p83036", "", api.NewDateTime(""), uv)

		uv = api.SetUnitValueString(e.Response.ResultData.P83016, e.Response.ResultData.P83016Unit)
		entries.AddUnitValue(name + ".p83016", e.Request.PsId.String(), "p83016", "", api.NewDateTime(""), uv)

		uv = api.SetUnitValueString(e.Response.ResultData.P83017, e.Response.ResultData.P83017Unit)
		entries.AddUnitValue(name + ".p83017", e.Request.PsId.String(), "p83017", "", api.NewDateTime(""), uv)


		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), time.Time{})

		for _, sid := range e.Response.ResultData.StorageInverterData {
			entries.StructToPoints(sid, name + ".StorageInverterData." + sid.PsKey, sid.PsKey, time.Time{})
		}
	}
	return entries
}
