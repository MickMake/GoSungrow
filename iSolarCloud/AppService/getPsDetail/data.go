package getPsDetail

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsDetail"
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
	ActualEnergy          []valueTypes.Float   `json:"actual_energy" PointUnitFrom:"actual_energy_unit"`
	ActualEnergyUnit      valueTypes.String    `json:"actual_energy_unit" PointId:"actual_energy_unit"`
	AlarmCount            valueTypes.Integer   `json:"alarm_count"`
	AreaID                interface{}   `json:"area_id"`
	AreaType              valueTypes.Integer   `json:"area_type"`
	BuildDate             valueTypes.DateTime  `json:"build_date"`
	Co2Reduce             valueTypes.UnitValue `json:"co2_reduce"`
	Co2ReduceTotal        valueTypes.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	CoalReduce            valueTypes.UnitValue `json:"coal_reduce"`
	CoalReduceTotal       valueTypes.UnitValue `json:"coal_reduce_total" PointId:"coal_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	ConnectGrid           string        `json:"connect_grid"`
	ConnectType           valueTypes.Integer   `json:"connect_type"`
	ContactPerson         valueTypes.String    `json:"contact_person"`
	CurrPower             valueTypes.UnitValue `json:"curr_power"` // Pv Power
	DataLastUpdateTime    valueTypes.DateTime  `json:"data_last_update_time"`
	DayEqHours            string        `json:"day_eq_hours"`
	Description           interface{}   `json:"description"`
	DesignCapacity        valueTypes.UnitValue `json:"design_capacity"`
	DesignCapacityBattery valueTypes.UnitValue `json:"design_capacity_battery"`
	DiagramURL            valueTypes.String    `json:"diagram_url"`
	EnergyScheme          interface{}   `json:"energy_scheme"`
	ExpectInstallDate     valueTypes.DateTime  `json:"expect_install_date"`
	FaultCount            valueTypes.Integer   `json:"fault_count"`
	FaultSendType         string        `json:"fault_send_type"`
	GcjLatitude           valueTypes.Float     `json:"gcj_latitude"`
	GcjLongitude          valueTypes.Float     `json:"gcj_longitude"`
	GprsLatitude          valueTypes.Float     `json:"gprs_latitude"`
	GprsLongitude         valueTypes.Float     `json:"gprs_longitude"`
	HasAmmeter            valueTypes.Bool      `json:"has_ammeter"`
	Images                []struct {
		FileID      valueTypes.Integer `json:"file_id"`
		ID          valueTypes.Integer `json:"id"`
		PicLanguage valueTypes.Integer `json:"pic_language"`
		PicType     valueTypes.Integer `json:"pic_type"`
		PictureName valueTypes.String  `json:"picture_name"`
		PictureURL  valueTypes.String  `json:"picture_url"`
		PsUnitUUID  interface{} `json:"ps_unit_uuid"`
	} `json:"images"`
	InstallDate            valueTypes.DateTime  `json:"install_date"`
	InstallerPsFaultStatus string        `json:"installer_ps_fault_status"`
	IsHaveEsInverter       valueTypes.Bool      `json:"is_have_es_inverter"`
	IsTransformSystem      valueTypes.Bool      `json:"is_transform_system"`
	IsTuv                  valueTypes.Bool      `json:"is_tuv"`
	Latitude               valueTypes.Float     `json:"latitude"`
	Longitude              valueTypes.Float     `json:"longitude"`
	MapLatitude            valueTypes.Float     `json:"map_latitude"`
	MapLongitude           valueTypes.Float     `json:"map_longitude"`
	MeterReduce            valueTypes.UnitValue `json:"meter_reduce"`
	MeterReduceTotal       valueTypes.UnitValue `json:"meter_reduce_total" PointId:"meter_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	MobileTel              valueTypes.String    `json:"mibile_tel" PointId:"mobile_tel"`
	MonthPr                string        `json:"monthPr" PointId:"monthPr" PointValueType:"" PointTimeSpan:"PointTimeSpanMonth"`
	MonthEnergy            valueTypes.UnitValue `json:"month_energy" PointId:"month_energy" PointValueType:"" PointTimeSpan:"PointTimeSpanMonth"`
	MonthEnergyVirgin      valueTypes.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthEqHours           valueTypes.Float     `json:"month_eq_hours" PointId:"month_eq_hours" PointValueType:"" PointTimeSpan:"PointTimeSpanMonth"`
	MonthIncome            valueTypes.UnitValue `json:"month_income" PointId:"month_income" PointValueType:"" PointTimeSpan:"PointTimeSpanMonth"`
	NoxReduce              valueTypes.UnitValue `json:"nox_reduce"`
	NoxReduceTotal         valueTypes.UnitValue `json:"nox_reduce_total" PointId:"nox_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	OperateYear            string        `json:"operate_year"`
	OperationBusName       valueTypes.String    `json:"operation_bus_name"`
	OwnerPsFaultStatus     string        `json:"owner_ps_fault_status"`
	P83012Value            valueTypes.Float     `json:"p83012_value" PointId:"p83012" PointUnitFrom:"p83012_unit"`
	P83012Unit             valueTypes.String    `json:"p83012_unit" PointId:"p83012_unit"`
	P83013Value            valueTypes.Float     `json:"p83013_value" PointId:"p83013" PointUnitFrom:"p83013_unit"`
	P83013Unit             valueTypes.String    `json:"p83013_unit" PointId:"p83013_unit"`
	P83036Value            valueTypes.Float     `json:"p83036_value" PointId:"p83036" PointUnitFrom:"p83036_unit"`
	P83036Unit             valueTypes.String    `json:"p83036_unit" PointId:"p83036_unit"`
	P83016                 valueTypes.Float     `json:"p83016" PointId:"p83016" PointUnitFrom:"p83016_unit"`
	P83016Unit             valueTypes.String    `json:"p83016_unit" PointId:"p83016_unit"`
	P83017                 valueTypes.Float     `json:"p83017" PointId:"p83017" PointUnitFrom:"p83017_unit"`
	P83017Unit             valueTypes.String    `json:"p83017_unit" PointId:"p83017_unit"`
	P83023                 valueTypes.Float     `json:"p83023" PointId:"p83023"`
	P83023y                string        `json:"p83023y" PointId:"p83023y"`
	P83023year             string        `json:"p83023year" PointId:"p83023y"`
	P83023ym               string        `json:"p83023ym" PointId:"p83023ym"`
	P83043                 valueTypes.Float     `json:"p83043" PointId:"p83043"`
	P83044                 valueTypes.Float     `json:"p83044" PointId:"p83044"`
	P83045                 valueTypes.Float     `json:"p83045" PointId:"p83045"`
	P83072Map              valueTypes.UnitValue `json:"p83072_map" PointId:"p83072"`
	P83072MapVirgin        valueTypes.UnitValue `json:"p83072_map_virgin"  PointIgnore:"true"`
	P83073Map              valueTypes.UnitValue `json:"p83073_map" PointId:"p83073"`
	P83073MapVirgin        valueTypes.UnitValue `json:"p83073_map_virgin"  PointIgnore:"true"`
	P83074Map              valueTypes.UnitValue `json:"p83074_map" PointId:"p83074"`
	P83074MapVirgin        valueTypes.UnitValue `json:"p83074_map_virgin"  PointIgnore:"true"`
	P83075Map              valueTypes.UnitValue `json:"p83075_map" PointId:"p83075"`
	P83075MapVirgin        valueTypes.UnitValue `json:"p83075_map_virgin"  PointIgnore:"true"`
	P83076Map              valueTypes.UnitValue `json:"p83076_map" PointId:"p83076"`           // Pv Power
	P83076MapVirgin        valueTypes.UnitValue `json:"p83076_map_virgin"  PointIgnore:"true"` // Pv Power
	P83077Map              valueTypes.UnitValue `json:"p83077_map" PointId:"p83077"`           // Pv Energy
	P83077MapVirgin        valueTypes.UnitValue `json:"p83077_map_virgin"  PointIgnore:"true"` // Pv Energy
	P83078Map              valueTypes.UnitValue `json:"p83078_map" PointId:"p83078"`
	P83078MapVirgin        valueTypes.UnitValue `json:"p83078_map_virgin"  PointIgnore:"true"`
	P83079Map              valueTypes.UnitValue `json:"p83079_map" PointId:"p83079"`
	P83079MapVirgin        valueTypes.UnitValue `json:"p83079_map_virgin"  PointIgnore:"true"`
	P83080Map              valueTypes.UnitValue `json:"p83080_map" PointId:"p83080"`
	P83080MapVirgin        valueTypes.UnitValue `json:"p83080_map_virgin"  PointIgnore:"true"`
	P83088Map              valueTypes.UnitValue `json:"p83088_map" PointId:"p83088"`           // Es Energy
	P83088MapVirgin        valueTypes.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"` // Es Energy
	P83089Map              valueTypes.UnitValue `json:"p83089_map" PointId:"p83089"`           // Es Discharge Energy
	P83089MapVirgin        valueTypes.UnitValue `json:"p83089_map_virgin"  PointIgnore:"true"` // Es Discharge Energy
	P83094Map              valueTypes.UnitValue `json:"p83094_map" PointId:"p83094"`
	P83094MapVirgin        valueTypes.UnitValue `json:"p83094_map_virgin"  PointIgnore:"true"`
	P83095Map              valueTypes.UnitValue `json:"p83095_map" PointId:"p83095" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"` // Es Total Discharge Energy
	P83095MapVirgin        valueTypes.UnitValue `json:"p83095_map_virgin"  PointIgnore:"true"`                                            // Es Total Discharge Energy
	P83097Map              valueTypes.UnitValue `json:"p83097_map" PointId:"p83097"`
	P83097MapVirgin        valueTypes.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83100Map              valueTypes.UnitValue `json:"p83100_map" PointId:"p83100"`
	P83100MapVirgin        valueTypes.UnitValue `json:"p83100_map_virgin"  PointIgnore:"true"`
	P83101Map              valueTypes.UnitValue `json:"p83101_map" PointId:"p83101"`
	P83101MapVirgin        valueTypes.UnitValue `json:"p83101_map_virgin"  PointIgnore:"true"`
	P83102Map              valueTypes.UnitValue `json:"p83102_map" PointId:"p83102"`
	P83102MapVirgin        valueTypes.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent          valueTypes.Float     `json:"p83102_percent" PointId:"p83102"`
	P83105Map              valueTypes.UnitValue `json:"p83105_map" PointId:"p83105"`
	P83105MapVirgin        valueTypes.UnitValue `json:"p83105_map_virgin"  PointIgnore:"true"`
	P83106Map              valueTypes.UnitValue `json:"p83106_map" PointId:"p83106"`
	P83106MapVirgin        valueTypes.UnitValue `json:"p83106_map_virgin"  PointIgnore:"true"`
	P83107Map              valueTypes.UnitValue `json:"p83107_map" PointId:"p83107"`
	P83107MapVirgin        valueTypes.UnitValue `json:"p83107_map_virgin"  PointIgnore:"true"`
	P83118Map              valueTypes.UnitValue `json:"p83118_map" PointId:"p83118"`
	P83118MapVirgin        valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map              valueTypes.UnitValue `json:"p83119_map" PointId:"p83119"`
	P83119MapVirgin        valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map              valueTypes.UnitValue `json:"p83120_map" PointId:"p83120"`
	P83120MapVirgin        valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121                 valueTypes.Float     `json:"p83121" PointId:"p83121"`
	P83122                 valueTypes.Float     `json:"p83122" PointId:"p83122"`
	P83123Map              valueTypes.UnitValue `json:"p83123_map" PointId:"p83123"`
	P83123MapVirgin        valueTypes.UnitValue `json:"p83123_map_virgin"  PointIgnore:"true"`
	P83124Map              valueTypes.UnitValue `json:"p83124_map" PointId:"p83124"`
	P83124MapVirgin        valueTypes.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83125                 valueTypes.Float     `json:"p83125" PointId:"p83125"`
	P83126                 valueTypes.Float     `json:"p83126" PointId:"p83126"`
	P83128MapVirgin        valueTypes.UnitValue `json:"p83128_map_virgin"  PointIgnore:"true"`
	P83202Map              valueTypes.UnitValue `json:"p83202_map" PointId:"p83202"`
	P83202MapVirgin        valueTypes.UnitValue `json:"p83202_map_virgin"  PointIgnore:"true"`
	PercentPlanYear        valueTypes.Float     `json:"percent_plan_year"`
	PlanEnergy             []valueTypes.Float   `json:"plan_energy" PointUnitFrom:"plan_energy_unit"`
	PlanEnergyUnit         valueTypes.String    `json:"plan_energy_unit" PointId:"plan_energy_unit"`
	PlanEnergyYear         valueTypes.UnitValue `json:"plan_energy_year"`
	PowderReduce           valueTypes.UnitValue `json:"powder_reduce"`
	PowderReduceTotal      valueTypes.UnitValue `json:"powder_reduce_total" PointId:"powder_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	PowerChargeSetted      valueTypes.Bool      `json:"power_charge_setted" PointId:"power_charge_setted" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	Producer               string        `json:"producer"`
	PsCountryID            valueTypes.Integer   `json:"ps_country_id"`
	PsFaultStatus          string        `json:"ps_fault_status"`
	PsHealthStatus         string        `json:"ps_health_status"`
	PsHolder               valueTypes.String    `json:"ps_holder"`
	PsLocation             valueTypes.String    `json:"ps_location"`
	PsName                 valueTypes.String    `json:"ps_name"`
	PsShortName            valueTypes.String    `json:"ps_short_name"`
	PsState                string        `json:"ps_state"`
	PsType                 valueTypes.Integer   `json:"ps_type"`
	PsTypeName             valueTypes.String    `json:"ps_type_name"`
	PsWindLevel            string        `json:"ps_wind_level"`
	PsWindPos              string        `json:"ps_wind_pos"`
	RecordCreateTime       valueTypes.DateTime  `json:"recore_create_time"`
	ReportType             string        `json:"report_type"`
	RobotNumSweepCapacity  struct {
		Num           valueTypes.Integer `json:"num"`
		SweepCapacity valueTypes.Float   `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`
	SafeStartDate                 valueTypes.DateTime  `json:"safe_start_date"`
	SelfConsumptionOffsetReminder valueTypes.Integer   `json:"self_consumption_offset_reminder"`
	ShippingAddress               valueTypes.String    `json:"shipping_address"`
	ShippingZipCode               valueTypes.String    `json:"shipping_zip_code"`
	So2Reduce                     valueTypes.UnitValue `json:"so2_reduce"`
	So2ReduceTotal                valueTypes.UnitValue `json:"so2_reduce_total" PointId:"so2_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	StorageInverterData           []struct {
		CommunicationDevSn      valueTypes.String    `json:"communication_dev_sn"`
		DevFaultStatus          valueTypes.Integer   `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer   `json:"dev_status"`
		DeviceCode              valueTypes.Integer   `json:"device_code"`
		DeviceModelCode         valueTypes.String    `json:"device_model_code"`
		DeviceName              valueTypes.String    `json:"device_name"`
		DeviceState             string        `json:"device_state"`
		DeviceType              valueTypes.Integer   `json:"device_type"`
		DrmStatus               valueTypes.Integer   `json:"drm_status"`
		DrmStatusName           valueTypes.String    `json:"drm_status_name"`
		EnergyFlow              []valueTypes.Integer `json:"energy_flow"`
		HasAmmeter              valueTypes.Bool      `json:"has_ammeter"`
		InstallerDevFaultStatus valueTypes.Integer   `json:"installer_dev_fault_status"`
		InverterSn              valueTypes.String    `json:"inverter_sn"`
		OwnerDevFaultStatus     valueTypes.Integer   `json:"owner_dev_fault_status"`
		P13003Map               valueTypes.UnitValue `json:"p13003_map" PointId:"p13003" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13003MapVirgin         valueTypes.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13011Map               valueTypes.UnitValue `json:"p13011_map" PointId:"p13011" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13011MapVirgin         valueTypes.UnitValue `json:"p13011_map_virgin"  PointIgnore:"true"`
		P13115Map               valueTypes.UnitValue `json:"p13115_map" PointId:"p13115" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13115MapVirgin         valueTypes.UnitValue `json:"p13115_map_virgin"  PointIgnore:"true"`
		P13119Map               valueTypes.UnitValue `json:"p13119_map" PointId:"p13119" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13119MapVirgin         valueTypes.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map               valueTypes.UnitValue `json:"p13121_map" PointId:"p13121" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13121MapVirgin         valueTypes.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map               valueTypes.UnitValue `json:"p13126_map" PointId:"p13126" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13126MapVirgin         valueTypes.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141                  valueTypes.Float     `json:"p13141" PointId:"p13141" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13142                  valueTypes.Float     `json:"p13142" PointId:"p13142" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13149Map               valueTypes.UnitValue `json:"p13149_map" PointId:"p13149" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13149MapVirgin         valueTypes.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map               valueTypes.UnitValue `json:"p13150_map" PointId:"p13150" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		P13150MapVirgin         valueTypes.UnitValue `json:"p13150_map_virgin"  PointIgnore:"true"`
		P13155                  valueTypes.Float     `json:"p13155" PointId:"p13155" PointValueType:"" PointTimeSpan:"PointTimeSpanInstant"`
		PsKey                   valueTypes.PsKey     `json:"ps_key"`
		UpdateTime              valueTypes.DateTime  `json:"update_time"`
		UUID                    valueTypes.Integer   `json:"uuid"`
	} `json:"storage_inverter_data"`
	SysScheme            valueTypes.Integer   `json:"sys_scheme"`
	TimeZoneID           valueTypes.Integer   `json:"time_zone_id"`
	Timezone             valueTypes.String    `json:"timezone"`
	TodayEnergy          valueTypes.UnitValue `json:"today_energy" PointId:"today_energy" PointValueType:"" PointTimeSpan:"PointTimeSpanDaily"`
	TodayEnergyVirgin    valueTypes.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome          valueTypes.UnitValue `json:"today_income" PointId:"today_income" PointValueType:"" PointTimeSpan:"PointTimeSpanDaily"`
	TotalEnergy          valueTypes.UnitValue `json:"total_energy" PointId:"total_energy" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	TotalEnergyVirgin    valueTypes.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalEnergyYear      valueTypes.UnitValue `json:"total_energy_year" PointId:"total_energy_year" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	TotalIncome          valueTypes.UnitValue `json:"total_income" PointId:"total_income" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	TreeReduce           valueTypes.UnitValue `json:"tree_reduce"`
	TreeReduceTotal      valueTypes.UnitValue `json:"tree_reduce_total" PointId:"tree_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	TuvLevel             string        `json:"tuv_level"`
	ValidFlag            valueTypes.Bool      `json:"valid_flag"`
	WaitAssignOrderCount valueTypes.Integer   `json:"wait_assign_order_count"`
	WaterReduce          valueTypes.UnitValue `json:"water_reduce"`
	WaterReduceTotal     valueTypes.UnitValue `json:"water_reduce_total" PointId:"water_reduce_total" PointValueType:"" PointTimeSpan:"PointTimeSpanTotal"`
	WgsLatitude          valueTypes.Float     `json:"wgs_latitude"`
	WgsLongitude         valueTypes.Float     `json:"wgs_longitude"`
	Year                 valueTypes.Integer   `json:"year"`
	ZfzyMap              valueTypes.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin        valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZipCode              valueTypes.String    `json:"zip_code"`
	ZjzzMap              valueTypes.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin        valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
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
		pkg := apiReflect.GetName("", *e)
		name := fmt.Sprintf("%s.%s", pkg, e.Request.PsId.String())
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))

		for _, sid := range e.Response.ResultData.StorageInverterData {
			name = fmt.Sprintf("%s.%s", pkg, sid.PsKey.Value())
			entries.StructToPoints(sid, name, sid.PsKey.Value(), valueTypes.NewDateTime(""))
		}
	}

	// for range Only.Once {
	// 	// name := fmt.Sprintf("getPsDetail.%s", e.Request.PsId.String())
	// 	//
	// 	// uv := api.SetUnitValueFloat(e.Response.ResultData.P83012Value.Value(), e.Response.ResultData.P83012Unit.Value())
	// 	// entries.AddUnitValue(name + ".p83012", e.Request.PsId.String(), "p83012", "", "", valueTypes.NewDateTime(""), uv)
	// 	//
	// 	// uv = api.SetUnitValueFloat(e.Response.ResultData.P83013Value.Value(), e.Response.ResultData.P83013Unit.Value())
	// 	// entries.AddUnitValue(name + ".p83013", e.Request.PsId.String(), "p83013", "", "", valueTypes.NewDateTime(""), uv)
	// 	//
	// 	// uv = api.SetUnitValueFloat(e.Response.ResultData.P83036Value.Value(), e.Response.ResultData.P83036Unit.Value())
	// 	// entries.AddUnitValue(name + ".p83036", e.Request.PsId.String(), "p83036", "", "", valueTypes.NewDateTime(""), uv)
	// 	//
	// 	// uv = api.SetUnitValueFloat(e.Response.ResultData.P83016.Value(), e.Response.ResultData.P83016Unit.Value())
	// 	// entries.AddUnitValue(name + ".p83016", e.Request.PsId.String(), "p83016", "", "", valueTypes.NewDateTime(""), uv)
	// 	//
	// 	// uv = api.SetUnitValueFloat(e.Response.ResultData.P83017.Value(), e.Response.ResultData.P83017Unit.Value())
	// 	// entries.AddUnitValue(name + ".p83017", e.Request.PsId.String(), "p83017", "", "", valueTypes.NewDateTime(""), uv)
	// 	//
	// 	// entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
	// 	//
	// 	// for _, sid := range e.Response.ResultData.StorageInverterData {
	// 	// 	entries.StructToPoints(sid, name + ".StorageInverterData." + sid.PsKey.Value(), sid.PsKey.Value(), valueTypes.NewDateTime(""))
	// 	// }
	// }

	return entries
}
