package getPsList

import (
	"GoSungrow/iSolarCloud/Common"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/powerStationService/getPsList"
const Disabled = false
const EndPointName = "AppService.getPsList"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []Common.Device  `json:"pageList" PointId:"page_list"`
	RowCount valueTypes.Count `json:"rowCount" PointId:"row_count"`
}

// type Device struct {
// 	GoStruct GoStruct.GoStruct `json:"GoStruct" PointIdFrom:"PsId" PointIdReplace:"true" PointDeviceFrom:"PsId"`
//
// 	PsId                      valueTypes.PsId           `json:"ps_id"`
// 	PsName                    valueTypes.String         `json:"ps_name"`
// 	PsStatus                  valueTypes.Bool           `json:"ps_status"`
// 	PsIsNotInit               valueTypes.Bool           `json:"ps_is_not_init"`
// 	IsBankPs                  valueTypes.Bool           `json:"is_bank_ps"`
// 	IsTuv                     valueTypes.Bool           `json:"is_tuv"`
// 	ValidFlag                 valueTypes.Bool           `json:"valid_flag"`
// 	PsType                    valueTypes.Integer        `json:"ps_type"`
// 	PsCode                    valueTypes.String         `json:"ps_code"`
// 	PsShortName               valueTypes.String         `json:"ps_short_name"`
// 	PsTimezone                valueTypes.String         `json:"ps_timezone"`
// 	PsFaultStatus             valueTypes.Integer        `json:"ps_fault_status"`
// 	PsHealthStatus            valueTypes.Integer        `json:"ps_health_status"`
// 	PsCountryId               valueTypes.Integer        `json:"ps_country_id"`
// 	PsHolder                  valueTypes.String         `json:"ps_holder"`
// 	AlarmCount                valueTypes.Count          `json:"alarm_count" PointUpdateFreq:"UpdateFreqBoot"`
// 	AlarmDevCount             valueTypes.Count          `json:"alarm_dev_count" PointUpdateFreq:"UpdateFreqBoot"`
// 	AreaId                    valueTypes.String         `json:"area_id"`
// 	AreaType                  valueTypes.Integer        `json:"area_type"`
// 	ArrearsStatus             valueTypes.Integer        `json:"arrears_status"`
// 	BuildDate                 valueTypes.DateTime       `json:"build_date" PointUpdateFreq:"UpdateFreqBoot" PointNameDateFormat:"2006/01/02 15:04:05"`
// 	ExpectInstallDate         valueTypes.DateTime       `json:"expect_install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
// 	InstallDate               valueTypes.DateTime       `json:"install_date" PointNameDateFormat:"2006/01/02 15:04:05"`
// 	BuildStatus               valueTypes.Integer        `json:"build_status" PointUpdateFreq:"UpdateFreqBoot"`
// 	Co2Reduce                 valueTypes.UnitValue      `json:"co2_reduce"`
// 	Co2ReduceTotal            valueTypes.UnitValue      `json:"co2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
// 	CurrPower                 valueTypes.UnitValue      `json:"curr_power"`
// 	DailyIrradiation          valueTypes.UnitValue      `json:"daily_irradiation" PointUpdateFreq:"UpdateFreqDay"`
// 	DailyIrradiationVirgin    valueTypes.Float          `json:"daily_irradiation_virgin" PointIgnore:"true"`
// 	Description               valueTypes.String         `json:"description"`
// 	DesignCapacity            valueTypes.Float          `json:"design_capacity" PointUnitFrom:"design_capacity_unit"`
// 	DesignCapacityUnit        valueTypes.String         `json:"design_capacity_unit"`
// 	DesignCapacityVirgin      valueTypes.Float          `json:"design_capacity_virgin" PointIgnore:"true"`
// 	EquivalentHour            valueTypes.UnitValue      `json:"equivalent_hour" PointUpdateFreq:"UpdateFreqDay"`
// 	FaultAlarmOfflineDevCount valueTypes.Count          `json:"fault_alarm_offline_dev_count"`
// 	FaultCount                valueTypes.Count          `json:"fault_count"`
// 	FaultDevCount             valueTypes.Count          `json:"fault_dev_count"`
// 	GcjLatitude               valueTypes.Float          `json:"gcj_latitude"`
// 	GcjLongitude              valueTypes.Float          `json:"gcj_longitude"`
// 	GprsLatitude              valueTypes.Float          `json:"gprs_latitude"`
// 	GprsLongitude             valueTypes.Float          `json:"gprs_longitude"`
// 	InstalledPowerMap         valueTypes.UnitValue      `json:"installed_power_map"`
// 	InstalledPowerVirgin      valueTypes.Float          `json:"installed_power_virgin" PointIgnore:"true"`
// 	InstallerAlarmCount       valueTypes.Count          `json:"installer_alarm_count"`
// 	InstallerFaultCount       valueTypes.Count          `json:"installer_fault_count"`
// 	InstallerPsFaultStatus    valueTypes.Integer        `json:"installer_ps_fault_status"`
// 	JoinYearInitElec          valueTypes.Float          `json:"join_year_init_elec"`
// 	Latitude                  valueTypes.Float          `json:"latitude"`
// 	Location                  valueTypes.String         `json:"location"`
// 	Longitude                 valueTypes.Float          `json:"longitude"`
// 	MapLatitude               valueTypes.Float          `json:"map_latitude"`
// 	MapLongitude              valueTypes.Float          `json:"map_longitude"`
// 	MlpeFlag                  valueTypes.Integer        `json:"mlpe_flag"`
// 	Nmi                       valueTypes.String         `json:"nmi"`
// 	OfflineDevCount           valueTypes.Count          `json:"offline_dev_count"`
// 	OperateYear               valueTypes.String         `json:"operate_year"`
// 	OperationBusName          valueTypes.String         `json:"operation_bus_name"`
// 	OwnerAlarmCount           valueTypes.Count          `json:"owner_alarm_count"`
// 	OwnerFaultCount           valueTypes.Count          `json:"owner_fault_count"`
// 	OwnerPsFaultStatus        valueTypes.Integer        `json:"owner_ps_fault_status"`
// 	Producer                  valueTypes.String         `json:"producer"`
// 	RecordCreateTime          valueTypes.DateTime       `json:"recore_create_time" PointId:"record_create_time" PointNameDateFormat:"2006/01/02 15:04:05"`
// 	SafeStartDate             valueTypes.DateTime       `json:"safe_start_date" PointNameDateFormat:"2006/01/02 15:04:05"`
// 	ShareType                 valueTypes.Integer        `json:"share_type"`
// 	ShippingAddress           valueTypes.String         `json:"shipping_address"`
// 	ShippingZipCode           valueTypes.String         `json:"shipping_zip_code"`
// 	SysScheme                 valueTypes.Integer        `json:"sys_scheme"`
// 	WgsLatitude               valueTypes.Float          `json:"wgs_latitude"`
// 	WgsLongitude              valueTypes.Float          `json:"wgs_longitude"`
// 	ZipCode                   valueTypes.String         `json:"zip_code"`
// 	Images                    Common.PowerStationImages `json:"images" PointArrayFlatten:"false"`
//
// 	P83022y                valueTypes.String    `json:"p83022y" PointId:"p83022" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83046                 valueTypes.Float     `json:"p83046" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83048                 valueTypes.Float     `json:"p83048" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83049                 valueTypes.Float     `json:"p83049" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83050                 valueTypes.Float     `json:"p83050" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83051                 valueTypes.Float     `json:"p83051" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83054                 valueTypes.Float     `json:"p83054" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83055                 valueTypes.Float     `json:"p83055" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83067                 valueTypes.Float     `json:"p83067" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83070                 valueTypes.Float     `json:"p83070" PointUpdateFreq:"UpdateFreq5Mins"`
// 	P83076                 valueTypes.Float     `json:"p83076" PointId:"_p83076" PointName:"Pv Power" PointIgnore:"true"`                  // Dupe of PvPower
// 	P83077                 valueTypes.Float     `json:"p83077" PointId:"_p83077" PointName:"Pv Energy" PointIgnore:"true"`                 // Dupe of PvEnergy
// 	P83081                 valueTypes.Float     `json:"p83081" PointId:"_p83081" PointName:"Es Power" PointIgnore:"true"`                  // Dupe of EsPower
// 	P83089                 valueTypes.Float     `json:"p83089" PointId:"_p83089" PointName:"Es Discharge Energy" PointIgnore:"true"`       // Dupe of EsDischargeEnergy
// 	P83095                 valueTypes.Float     `json:"p83095" PointId:"_p83095" PointName:"Es Total Discharge Energy" PointIgnore:"true"` // Dupe of EsTotalDischargeEnergy
// 	P83118                 valueTypes.Float     `json:"p83118" PointId:"_p83118" PointName:"Use Energy" PointIgnore:"true"`                // Dupe of UseEnergy
// 	P83120                 valueTypes.Float     `json:"p83120" PointId:"_p83120" PointName:"Es Energy" PointIgnore:"true"`                 // Dupe of EsEnergy
// 	P83127                 valueTypes.Float     `json:"p83127" PointId:"_p83127" PointName:"Es Total Energy" PointIgnore:"true"`           // Dupe of EsTotalEnergy
// 	PvPower                valueTypes.UnitValue `json:"pv_power" PointId:"p83076" PointName:"Pv Power" PointUpdateFreq:"UpdateFreq5Mins"`
// 	PvEnergy               valueTypes.UnitValue `json:"pv_energy" PointId:"p83077" PointName:"Pv Energy" PointUpdateFreq:"UpdateFreq5Mins"`
// 	EsPower                valueTypes.UnitValue `json:"es_power" PointId:"p83081" PointName:"ES Power" PointUpdateFreq:"UpdateFreq5Mins"`
// 	EsDischargeEnergy      valueTypes.UnitValue `json:"es_disenergy" PointId:"p83089" PointName:"ES Discharge Energy" PointUpdateFreq:"UpdateFreq5Mins"`
// 	EsTotalDischargeEnergy valueTypes.UnitValue `json:"es_total_disenergy" PointId:"p83095" PointName:"ES Total Discharge Energy" PointUpdateFreq:"UpdateFreqTotal"`
// 	UseEnergy              valueTypes.UnitValue `json:"use_energy" PointId:"p83118" PointName:"Use Energy" PointUpdateFreq:"UpdateFreq5Mins"`
// 	EsEnergy               valueTypes.UnitValue `json:"es_energy" PointId:"p83120" PointName:"ES Energy" PointUpdateFreq:"UpdateFreq5Mins"`
// 	EsTotalEnergy          valueTypes.UnitValue `json:"es_total_energy" PointId:"p83127" PointName:"ES Total Energy" PointUpdateFreq:"UpdateFreqTotal"`
// 	ParamCo2               valueTypes.Float     `json:"param_co2"`
// 	ParamCoal              valueTypes.Float     `json:"param_coal"`
// 	ParamIncome            valueTypes.Float     `json:"param_income"`
// 	ParamMeter             valueTypes.Float     `json:"param_meter"`
// 	ParamNox               valueTypes.Float     `json:"param_nox"`
// 	ParamPowder            valueTypes.Float     `json:"param_powder"`
// 	ParamSo2               valueTypes.Float     `json:"param_so2"`
// 	ParamTree              valueTypes.Float     `json:"param_tree"`
// 	ParamWater             valueTypes.Float     `json:"param_water"`
// 	PrScale                string               `json:"pr_scale"`
// 	Radiation              valueTypes.UnitValue `json:"radiation"`
// 	RadiationVirgin        valueTypes.Float     `json:"radiation_virgin" PointIgnore:"true"`
// 	TodayEnergy            valueTypes.UnitValue `json:"today_energy" PointUpdateFreq:"UpdateFreqDay"`
// 	TodayIncome            valueTypes.UnitValue `json:"today_income" PointUpdateFreq:"UpdateFreqDay"`
// 	TotalCapacity          valueTypes.UnitValue `json:"total_capcity" PointId:"total_capacity" PointUpdateFreq:"UpdateFreqTotal"`
// 	TotalEnergy            valueTypes.UnitValue `json:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`
// 	TotalIncome            valueTypes.UnitValue `json:"total_income" PointUpdateFreq:"UpdateFreqTotal"`
// 	TotalInitCo2Accelerate valueTypes.Float     `json:"total_init_co2_accelerate" PointUpdateFreq:"UpdateFreqTotal"`
// 	TotalInitElec          valueTypes.Float     `json:"total_init_elec" PointUpdateFreq:"UpdateFreqTotal"`
// 	TotalInitProfit        valueTypes.Float     `json:"total_init_profit" PointUpdateFreq:"UpdateFreqTotal"`
// }

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
