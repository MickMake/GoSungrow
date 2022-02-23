package getPsDetail

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPsDetail"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
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
	AlarmCount            int64         `json:"alarm_count"`
	AreaID                interface{}   `json:"area_id"`
	AreaType              string        `json:"area_type"`
	BuildDate             string        `json:"build_date"`
	Co2Reduce             api.UnitValue `json:"co2_reduce"`
	Co2ReduceTotal        api.UnitValue `json:"co2_reduce_total"`
	CoalReduce            api.UnitValue `json:"coal_reduce"`
	CoalReduceTotal       api.UnitValue `json:"coal_reduce_total"`
	ConnectGrid           string        `json:"connect_grid"`
	ConnectType           int64         `json:"connect_type"`
	ContactPerson         string        `json:"contact_person"`
	CurrPower             api.UnitValue `json:"curr_power"` // Pv Power
	DataLastUpdateTime    string        `json:"data_last_update_time"`
	DayEqHours            string        `json:"day_eq_hours"`
	Description           interface{}   `json:"description"`
	DesignCapacity        api.UnitValue `json:"design_capacity"`
	DesignCapacityBattery api.UnitValue `json:"design_capacity_battery"`
	DiagramURL            string        `json:"diagram_url"`
	EnergyScheme          interface{}   `json:"energy_scheme"`
	ExpectInstallDate     string        `json:"expect_install_date"`
	FaultCount            int64         `json:"fault_count"`
	FaultSendType         string        `json:"fault_send_type"`
	GcjLatitude           string        `json:"gcj_latitude"`
	GcjLongitude          string        `json:"gcj_longitude"`
	GprsLatitude          string        `json:"gprs_latitude"`
	GprsLongitude         string        `json:"gprs_longitude"`
	HasAmmeter            int64         `json:"has_ammeter"`
	Images                []struct {
		FileID      int64       `json:"file_id"`
		ID          int64       `json:"id"`
		PicLanguage int64       `json:"pic_language"`
		PicType     int64       `json:"pic_type"`
		PictureName string      `json:"picture_name"`
		PictureURL  string      `json:"picture_url"`
		PsUnitUUID  interface{} `json:"ps_unit_uuid"`
	} `json:"images"`
	InstallDate            string        `json:"install_date"`
	InstallerPsFaultStatus string        `json:"installer_ps_fault_status"`
	IsHaveEsInverter       int64         `json:"is_have_es_inverter"`
	IsTransformSystem      string        `json:"is_transform_system"`
	IsTuv                  int64         `json:"is_tuv"`
	Latitude               float64       `json:"latitude"`
	Longitude              float64       `json:"longitude"`
	MapLatitude            string        `json:"map_latitude"`
	MapLongitude           string        `json:"map_longitude"`
	MeterReduce            api.UnitValue `json:"meter_reduce"`
	MeterReduceTotal       api.UnitValue `json:"meter_reduce_total"`
	MibileTel              string        `json:"mibile_tel"`
	MonthPr                string        `json:"monthPr"`
	MonthEnergy            api.UnitValue `json:"month_energy"`
	MonthEnergyVirgin      api.UnitValue `json:"month_energy_virgin"`
	MonthEqHours           string        `json:"month_eq_hours"`
	MonthIncome            api.UnitValue `json:"month_income"`
	NoxReduce              api.UnitValue `json:"nox_reduce"`
	NoxReduceTotal         api.UnitValue `json:"nox_reduce_total"`
	OperateYear            string        `json:"operate_year"`
	OperationBusName       string        `json:"operation_bus_name"`
	OwnerPsFaultStatus     string        `json:"owner_ps_fault_status"`
	P83012Unit             string        `json:"p83012_unit"`
	P83012Value            string        `json:"p83012_value"`
	P83013Unit             string        `json:"p83013_unit"`
	P83013Value            string        `json:"p83013_value"`
	P83016                 string        `json:"p83016"`
	P83016Unit             string        `json:"p83016_unit"`
	P83017                 string        `json:"p83017"`
	P83017Unit             string        `json:"p83017_unit"`
	P83023                 string        `json:"p83023"`
	P83023y                string        `json:"p83023y"`
	P83023year             string        `json:"p83023year"`
	P83023ym               string        `json:"p83023ym"`
	P83036Unit             string        `json:"p83036_unit"`
	P83036Value            string        `json:"p83036_value"`
	P83043                 string        `json:"p83043"`
	P83044                 string        `json:"p83044"`
	P83045                 string        `json:"p83045"`
	P83072Map              api.UnitValue `json:"p83072_map"`
	P83072MapVirgin        api.UnitValue `json:"p83072_map_virgin"`
	P83073Map              api.UnitValue `json:"p83073_map"`
	P83073MapVirgin        api.UnitValue `json:"p83073_map_virgin"`
	P83074Map              api.UnitValue `json:"p83074_map"`
	P83074MapVirgin        api.UnitValue `json:"p83074_map_virgin"`
	P83075Map              api.UnitValue `json:"p83075_map"`
	P83075MapVirgin        api.UnitValue `json:"p83075_map_virgin"`
	P83076Map              api.UnitValue `json:"p83076_map"`        // Pv Power
	P83076MapVirgin        api.UnitValue `json:"p83076_map_virgin"` // Pv Power
	P83077Map              api.UnitValue `json:"p83077_map"`        // Pv Energy
	P83077MapVirgin        api.UnitValue `json:"p83077_map_virgin"` // Pv Energy
	P83078Map              api.UnitValue `json:"p83078_map"`
	P83078MapVirgin        api.UnitValue `json:"p83078_map_virgin"`
	P83079Map              api.UnitValue `json:"p83079_map"`
	P83079MapVirgin        api.UnitValue `json:"p83079_map_virgin"`
	P83080Map              api.UnitValue `json:"p83080_map"`
	P83080MapVirgin        api.UnitValue `json:"p83080_map_virgin"`
	P83088Map              api.UnitValue `json:"p83088_map"`        // Es Energy
	P83088MapVirgin        api.UnitValue `json:"p83088_map_virgin"` // Es Energy
	P83089Map              api.UnitValue `json:"p83089_map"`        // Es Discharge Energy
	P83089MapVirgin        api.UnitValue `json:"p83089_map_virgin"` // Es Discharge Energy
	P83094Map              api.UnitValue `json:"p83094_map"`
	P83094MapVirgin        api.UnitValue `json:"p83094_map_virgin"`
	P83095Map              api.UnitValue `json:"p83095_map"`        // Es Total Discharge Energy
	P83095MapVirgin        api.UnitValue `json:"p83095_map_virgin"` // Es Total Discharge Energy
	P83097Map              api.UnitValue `json:"p83097_map"`
	P83097MapVirgin        api.UnitValue `json:"p83097_map_virgin"`
	P83100Map              api.UnitValue `json:"p83100_map"`
	P83100MapVirgin        api.UnitValue `json:"p83100_map_virgin"`
	P83101Map              api.UnitValue `json:"p83101_map"`
	P83101MapVirgin        api.UnitValue `json:"p83101_map_virgin"`
	P83102Map              api.UnitValue `json:"p83102_map"`
	P83102MapVirgin        api.UnitValue `json:"p83102_map_virgin"`
	P83102Percent          string        `json:"p83102_percent"`
	P83105Map              api.UnitValue `json:"p83105_map"`
	P83105MapVirgin        api.UnitValue `json:"p83105_map_virgin"`
	P83106Map              api.UnitValue `json:"p83106_map"`
	P83106MapVirgin        api.UnitValue `json:"p83106_map_virgin"`
	P83107Map              api.UnitValue `json:"p83107_map"`
	P83107MapVirgin        api.UnitValue `json:"p83107_map_virgin"`
	P83118Map              api.UnitValue `json:"p83118_map"`
	P83118MapVirgin        api.UnitValue `json:"p83118_map_virgin"`
	P83119Map              api.UnitValue `json:"p83119_map"`
	P83119MapVirgin        api.UnitValue `json:"p83119_map_virgin"`
	P83120Map              api.UnitValue `json:"p83120_map"`
	P83120MapVirgin        api.UnitValue `json:"p83120_map_virgin"`
	P83121                 string        `json:"p83121"`
	P83122                 string        `json:"p83122"`
	P83123Map              api.UnitValue `json:"p83123_map"`
	P83123MapVirgin        api.UnitValue `json:"p83123_map_virgin"`
	P83124Map              api.UnitValue `json:"p83124_map"`
	P83124MapVirgin        api.UnitValue `json:"p83124_map_virgin"`
	P83125                 string        `json:"p83125"`
	P83126                 string        `json:"p83126"`
	P83128MapVirgin        api.UnitValue `json:"p83128_map_virgin"`
	P83202Map              api.UnitValue `json:"p83202_map"`
	P83202MapVirgin        api.UnitValue `json:"p83202_map_virgin"`
	PercentPlanYear        float64       `json:"percent_plan_year"`
	PlanEnergy             []string      `json:"plan_energy"`
	PlanEnergyUnit         string        `json:"plan_energy_unit"`
	PlanEnergyYear         api.UnitValue `json:"plan_energy_year"`
	PowderReduce           api.UnitValue `json:"powder_reduce"`
	PowderReduceTotal      api.UnitValue `json:"powder_reduce_total"`
	PowerChargeSetted      int64         `json:"power_charge_setted"`
	Producer               string        `json:"producer"`
	PsCountryID            int64         `json:"ps_country_id"`
	PsFaultStatus          string        `json:"ps_fault_status"`
	PsHealthStatus         string        `json:"ps_health_status"`
	PsHolder               string        `json:"ps_holder"`
	PsLocation             string        `json:"ps_location"`
	PsName                 string        `json:"ps_name"`
	PsShortName            string        `json:"ps_short_name"`
	PsState                string        `json:"ps_state"`
	PsType                 int64         `json:"ps_type"`
	PsTypeName             string        `json:"ps_type_name"`
	PsWindLevel            string        `json:"ps_wind_level"`
	PsWindPos              string        `json:"ps_wind_pos"`
	RecoreCreateTime       string        `json:"recore_create_time"`
	ReportType             string        `json:"report_type"`
	RobotNumSweepCapacity  struct {
		Num           int64   `json:"num"`
		SweepCapacity float64 `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`
	SafeStartDate                 string        `json:"safe_start_date"`
	SelfConsumptionOffsetReminder int64         `json:"self_consumption_offset_reminder"`
	ShippingAddress               string        `json:"shipping_address"`
	ShippingZipCode               string        `json:"shipping_zip_code"`
	So2Reduce                     api.UnitValue `json:"so2_reduce"`
	So2ReduceTotal                api.UnitValue `json:"so2_reduce_total"`
	StorageInverterData           []struct {
		CommunicationDevSn      string        `json:"communication_dev_sn"`
		DevFaultStatus          int64         `json:"dev_fault_status"`
		DevStatus               int64         `json:"dev_status"`
		DeviceCode              int64         `json:"device_code"`
		DeviceModelCode         string        `json:"device_model_code"`
		DeviceName              string        `json:"device_name"`
		DeviceState             string        `json:"device_state"`
		DeviceType              int64         `json:"device_type"`
		DrmStatus               string        `json:"drm_status"`
		DrmStatusName           string        `json:"drm_status_name"`
		EnergyFlow              []string      `json:"energy_flow"`
		HasAmmeter              int64         `json:"has_ammeter"`
		InstallerDevFaultStatus int64         `json:"installer_dev_fault_status"`
		InverterSn              string        `json:"inverter_sn"`
		OwnerDevFaultStatus     int64         `json:"owner_dev_fault_status"`
		P13003Map               api.UnitValue `json:"p13003_map"`
		P13003MapVirgin         api.UnitValue `json:"p13003_map_virgin"`
		P13011Map               api.UnitValue `json:"p13011_map"`
		P13011MapVirgin         api.UnitValue `json:"p13011_map_virgin"`
		P13115Map               api.UnitValue `json:"p13115_map"`
		P13115MapVirgin         api.UnitValue `json:"p13115_map_virgin"`
		P13119Map               api.UnitValue `json:"p13119_map"`
		P13119MapVirgin         api.UnitValue `json:"p13119_map_virgin"`
		P13121Map               api.UnitValue `json:"p13121_map"`
		P13121MapVirgin         api.UnitValue `json:"p13121_map_virgin"`
		P13126Map               api.UnitValue `json:"p13126_map"`
		P13126MapVirgin         api.UnitValue `json:"p13126_map_virgin"`
		P13141                  string        `json:"p13141"`
		P13142                  string        `json:"p13142"`
		P13149Map               api.UnitValue `json:"p13149_map"`
		P13149MapVirgin         api.UnitValue `json:"p13149_map_virgin"`
		P13150Map               api.UnitValue `json:"p13150_map"`
		P13150MapVirgin         api.UnitValue `json:"p13150_map_virgin"`
		P13155                  string        `json:"p13155"`
		PsKey                   string        `json:"ps_key"`
		UpdateTime              string        `json:"update_time"`
		UUID                    int64         `json:"uuid"`
	} `json:"storage_inverter_data"`
	SysScheme            int64         `json:"sys_scheme"`
	TimeZoneID           int64         `json:"time_zone_id"`
	Timezone             string        `json:"timezone"`
	TodayEnergy          api.UnitValue `json:"today_energy"`
	TodayEnergyVirgin    api.UnitValue `json:"today_energy_virgin"`
	TodayIncome          api.UnitValue `json:"today_income"`
	TotalEnergy          api.UnitValue `json:"total_energy"`
	TotalEnergyVirgin    api.UnitValue `json:"total_energy_virgin"`
	TotalEnergyYear      api.UnitValue `json:"total_energy_year"`
	TotalIncome          api.UnitValue `json:"total_income"`
	TreeReduce           api.UnitValue `json:"tree_reduce"`
	TreeReduceTotal      api.UnitValue `json:"tree_reduce_total"`
	TuvLevel             string        `json:"tuv_level"`
	ValidFlag            string        `json:"valid_flag"`
	WaitAssignOrderCount int64         `json:"wait_assign_order_count"`
	WaterReduce          api.UnitValue `json:"water_reduce"`
	WaterReduceTotal     api.UnitValue `json:"water_reduce_total"`
	WgsLatitude          float64       `json:"wgs_latitude"`
	WgsLongitude         float64       `json:"wgs_longitude"`
	Year                 string        `json:"year"`
	ZfzyMap              api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin        api.UnitValue `json:"zfzy_map_virgin"`
	ZipCode              string        `json:"zip_code"`
	ZjzzMap              api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin        api.UnitValue `json:"zjzz_map_virgin"`
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

// func (e *ResultData) GetCsv(name string) api.Csv {
// 	var ret api.Csv
// 	for range Only.Once {
// 		points := e.GetDataByName(name)
// 		ret = ret.SetHeader([]string{
// 			"Date",
// 			"PointGroupName",
// 			"PointName",
// 			"Value",
// 			"Unit",
// 		})
//
// 		for _, p := range points {
// 			t := api.NewDateTime(p.TimeStamp)
// 			ret = ret.AddRow([]string{
// 				t.Format(api.DtLayout),
// 				p.PointGroupName,
// 				p.PointName,
// 				p.Value,
// 				p.Unit,
// 			})
// 		}
// 	}
// 	return ret
// }
