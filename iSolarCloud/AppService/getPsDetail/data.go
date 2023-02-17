package getPsDetail

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/Common"
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/powerStationService/getPsDetail"
const Disabled = false
const EndPointName = "AppService.getPsDetail"

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
	ActualEnergy []valueTypes.Float `json:"actual_energy" PointId:"actual_energy" PointName:"Actual Energy" PointUnitFrom:"ActualEnergyUnit" PointArrayFlatten:"true"`
	PlanEnergy   []valueTypes.Float `json:"plan_energy" PointUnitFrom:"PlanEnergyUnit" PointArrayFlatten:"true"`

	BuildDate                     valueTypes.DateTime  `json:"build_date" PointNameDateFormat:"DateTimeLayout"`
	DataLastUpdateTime            valueTypes.DateTime  `json:"data_last_update_time" PointNameDateFormat:"DateTimeLayout"`
	ExpectInstallDate             valueTypes.DateTime  `json:"expect_install_date" PointNameDateFormat:"DateTimeLayout"`
	InstallDate                   valueTypes.DateTime  `json:"install_date" PointNameDateFormat:"DateTimeLayout"`
	RecordCreateTime              valueTypes.DateTime  `json:"recore_create_time" PointId:"record_create_time" PointNameDateFormat:"DateTimeLayout"`
	SafeStartDate                 valueTypes.DateTime  `json:"safe_start_date" PointNameDateFormat:"DateTimeLayout"`
	ActualEnergyUnit              valueTypes.String    `json:"actual_energy_unit" PointId:"actual_energy_unit"  PointIgnore:"true"`
	AlarmCount                    valueTypes.Count     `json:"alarm_count"`
	AreaId                        interface{}          `json:"area_id"`
	AreaType                      valueTypes.Integer   `json:"area_type"`
	Co2Reduce                     valueTypes.UnitValue `json:"co2_reduce"`
	Co2ReduceTotal                valueTypes.UnitValue `json:"co2_reduce_total" PointId:"co2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	CoalReduce                    valueTypes.UnitValue `json:"coal_reduce"`
	CoalReduceTotal               valueTypes.UnitValue `json:"coal_reduce_total" PointId:"coal_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	ConnectGrid                   valueTypes.String    `json:"connect_grid"`
	ConnectType                   valueTypes.Integer   `json:"connect_type"`
	ContactPerson                 valueTypes.String    `json:"contact_person"`
	CurrPower                     valueTypes.UnitValue `json:"curr_power"` // Pv Power
	DayEqHours                    valueTypes.Float     `json:"day_eq_hours" PointUnit:"h" PointUpdateFreq:"UpdateFreqDay"`
	Description                   interface{}          `json:"description"`
	DesignCapacity                valueTypes.UnitValue `json:"design_capacity"`
	DesignCapacityBattery         valueTypes.UnitValue `json:"design_capacity_battery"`
	DiagramURL                    valueTypes.String    `json:"diagram_url"`
	EnergyScheme                  interface{}          `json:"energy_scheme"`
	FaultCount                    valueTypes.Count     `json:"fault_count"`
	FaultSendType                 valueTypes.String    `json:"fault_send_type"`
	GcjLatitude                   valueTypes.Float     `json:"gcj_latitude"`
	GcjLongitude                  valueTypes.Float     `json:"gcj_longitude"`
	GprsLatitude                  valueTypes.Float     `json:"gprs_latitude"`
	GprsLongitude                 valueTypes.Float     `json:"gprs_longitude"`
	HasAmmeter                    valueTypes.Bool      `json:"has_ammeter"`
	InstallerPsFaultStatus        valueTypes.Integer   `json:"installer_ps_fault_status"`
	IsHaveEsInverter              valueTypes.Bool      `json:"is_have_es_inverter"`
	IsTransformSystem             valueTypes.Bool      `json:"is_transform_system"`
	IsTuv                         valueTypes.Bool      `json:"is_tuv"`
	Latitude                      valueTypes.Float     `json:"latitude"`
	Longitude                     valueTypes.Float     `json:"longitude"`
	MapLatitude                   valueTypes.Float     `json:"map_latitude"`
	MapLongitude                  valueTypes.Float     `json:"map_longitude"`
	MeterReduce                   valueTypes.UnitValue `json:"meter_reduce"`
	MeterReduceTotal              valueTypes.UnitValue `json:"meter_reduce_total" PointId:"meter_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	MobileTel                     valueTypes.String    `json:"mibile_tel" PointId:"mobile_tel"`
	MonthPr                       valueTypes.String    `json:"monthPr" PointId:"month_pr" PointUpdateFreq:"UpdateFreqMonth"`
	MonthEnergy                   valueTypes.UnitValue `json:"month_energy" PointId:"month_energy" PointUpdateFreq:"UpdateFreqMonth"`
	MonthEnergyVirgin             valueTypes.UnitValue `json:"month_energy_virgin"  PointIgnore:"true"`
	MonthEqHours                  valueTypes.Float     `json:"month_eq_hours" PointId:"month_eq_hours" PointUnit:"h" PointUpdateFreq:"UpdateFreqMonth"`
	MonthIncome                   valueTypes.UnitValue `json:"month_income" PointId:"month_income" PointUpdateFreq:"UpdateFreqMonth"`
	NoxReduce                     valueTypes.UnitValue `json:"nox_reduce"`
	NoxReduceTotal                valueTypes.UnitValue `json:"nox_reduce_total" PointId:"nox_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	OperateYear                   valueTypes.String    `json:"operate_year"`
	OperationBusName              valueTypes.String    `json:"operation_bus_name"`
	OwnerPsFaultStatus            valueTypes.Integer   `json:"owner_ps_fault_status"`
	PercentPlanYear               valueTypes.Float     `json:"percent_plan_year"`
	PlanEnergyUnit                valueTypes.String    `json:"plan_energy_unit" PointId:"plan_energy_unit"  PointIgnore:"true"`
	PlanEnergyYear                valueTypes.UnitValue `json:"plan_energy_year"`
	PowderReduce                  valueTypes.UnitValue `json:"powder_reduce"`
	PowderReduceTotal             valueTypes.UnitValue `json:"powder_reduce_total" PointId:"powder_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	PowerChargeSet                valueTypes.Bool      `json:"power_charge_setted" PointId:"power_charge_set" PointUpdateFreq:"UpdateFreqTotal"`
	Producer                      valueTypes.String    `json:"producer"`
	PsCountryId                   valueTypes.Integer   `json:"ps_country_id"`
	PsFaultStatus                 valueTypes.Integer   `json:"ps_fault_status"`
	PsHealthStatus                valueTypes.Integer   `json:"ps_health_status"`
	PsHolder                      valueTypes.String    `json:"ps_holder"`
	PsLocation                    valueTypes.String    `json:"ps_location"`
	PsName                        valueTypes.String    `json:"ps_name"`
	PsShortName                   valueTypes.String    `json:"ps_short_name"`
	PsState                       valueTypes.Integer   `json:"ps_state"`
	PsType                        valueTypes.Integer   `json:"ps_type"`
	PsTypeName                    valueTypes.String    `json:"ps_type_name"`
	PsWindLevel                   valueTypes.String    `json:"ps_wind_level"`
	PsWindPos                     valueTypes.String    `json:"ps_wind_pos"`
	ReportType                    valueTypes.String    `json:"report_type"`
	SelfConsumptionOffsetReminder valueTypes.Integer   `json:"self_consumption_offset_reminder"`
	ShippingAddress               valueTypes.String    `json:"shipping_address"`
	ShippingZipCode               valueTypes.String    `json:"shipping_zip_code"`
	So2Reduce                     valueTypes.UnitValue `json:"so2_reduce"`
	So2ReduceTotal                valueTypes.UnitValue `json:"so2_reduce_total" PointId:"so2_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	SysScheme                     valueTypes.Integer   `json:"sys_scheme"`
	TimeZoneId                    valueTypes.Integer   `json:"time_zone_id"`
	Timezone                      valueTypes.String    `json:"timezone"`
	TodayEnergy                   valueTypes.UnitValue `json:"today_energy" PointId:"today_energy" PointUpdateFreq:"UpdateFreqDay"`
	TodayEnergyVirgin             valueTypes.UnitValue `json:"today_energy_virgin"  PointIgnore:"true"`
	TodayIncome                   valueTypes.UnitValue `json:"today_income" PointId:"today_income" PointUpdateFreq:"UpdateFreqDay"`
	TotalEnergy                   valueTypes.UnitValue `json:"total_energy" PointId:"total_energy" PointUpdateFreq:"UpdateFreqTotal"`
	TotalEnergyVirgin             valueTypes.UnitValue `json:"total_energy_virgin"  PointIgnore:"true"`
	TotalEnergyYear               valueTypes.UnitValue `json:"total_energy_year" PointId:"total_energy_year" PointUpdateFreq:"UpdateFreqTotal"`
	TotalIncome                   valueTypes.UnitValue `json:"total_income" PointId:"total_income" PointUpdateFreq:"UpdateFreqTotal"`
	TreeReduce                    valueTypes.UnitValue `json:"tree_reduce"`
	TreeReduceTotal               valueTypes.UnitValue `json:"tree_reduce_total" PointId:"tree_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	TuvLevel                      valueTypes.String    `json:"tuv_level"`
	ValidFlag                     valueTypes.Bool      `json:"valid_flag"`
	WaitAssignOrderCount          valueTypes.Count     `json:"wait_assign_order_count"`
	WaterReduce                   valueTypes.UnitValue `json:"water_reduce"`
	WaterReduceTotal              valueTypes.UnitValue `json:"water_reduce_total" PointId:"water_reduce_total" PointUpdateFreq:"UpdateFreqTotal"`
	WgsLatitude                   valueTypes.Float     `json:"wgs_latitude"`
	WgsLongitude                  valueTypes.Float     `json:"wgs_longitude"`
	Year                          valueTypes.Integer   `json:"year"`
	ZipCode                       valueTypes.String    `json:"zip_code"`

	P83012Value     valueTypes.Float     `json:"p83012_value" PointId:"p83012" PointUnitFrom:"P83012Unit" PointVirtual:"true"`
	P83012Unit      valueTypes.String    `json:"p83012_unit" PointId:"p83012_unit"  PointIgnore:"true"`
	P83013Value     valueTypes.Float     `json:"p83013_value" PointId:"p83013" PointUnitFrom:"P83013Unit" PointVirtual:"true"`
	P83013Unit      valueTypes.String    `json:"p83013_unit" PointId:"p83013_unit"  PointIgnore:"true"`
	P83036Value     valueTypes.Float     `json:"p83036_value" PointId:"p83036" PointUnitFrom:"P83036Unit" PointVirtual:"true"`
	P83036Unit      valueTypes.String    `json:"p83036_unit" PointId:"p83036_unit"  PointIgnore:"true"`
	P83016          valueTypes.Float     `json:"p83016" PointId:"p83016" PointUnitFrom:"P83016Unit" PointVirtual:"true"`
	P83016Unit      valueTypes.String    `json:"p83016_unit" PointId:"p83016_unit"  PointIgnore:"true"`
	P83017          valueTypes.Float     `json:"p83017" PointId:"p83017" PointUnitFrom:"P83017Unit" PointVirtual:"true"`
	P83017Unit      valueTypes.String    `json:"p83017_unit" PointId:"p83017_unit"  PointIgnore:"true"`
	P83023          valueTypes.Float     `json:"p83023" PointId:"p83023" PointVirtual:"true"`
	P83023y         valueTypes.String    `json:"p83023y" PointId:"p83023y" PointVirtual:"true"`
	P83023year      valueTypes.String    `json:"p83023year" PointId:"p83023y" PointVirtual:"true"`
	P83023ym        valueTypes.String    `json:"p83023ym" PointId:"p83023ym" PointVirtual:"true"`
	P83043          valueTypes.Float     `json:"p83043" PointId:"p83043" PointVirtual:"true"`
	P83044          valueTypes.Float     `json:"p83044" PointId:"p83044" PointVirtual:"true"`
	P83045          valueTypes.Float     `json:"p83045" PointId:"p83045" PointVirtual:"true"`
	P83072Map       valueTypes.UnitValue `json:"p83072_map" PointId:"p83072" PointVirtual:"true"`
	P83072MapVirgin valueTypes.UnitValue `json:"p83072_map_virgin"  PointIgnore:"true"`
	P83073Map       valueTypes.UnitValue `json:"p83073_map" PointId:"p83073" PointVirtual:"true"`
	P83073MapVirgin valueTypes.UnitValue `json:"p83073_map_virgin"  PointIgnore:"true"`
	P83074Map       valueTypes.UnitValue `json:"p83074_map" PointId:"p83074" PointVirtual:"true"`
	P83074MapVirgin valueTypes.UnitValue `json:"p83074_map_virgin"  PointIgnore:"true"`
	P83075Map       valueTypes.UnitValue `json:"p83075_map" PointId:"p83075" PointVirtual:"true"`
	P83075MapVirgin valueTypes.UnitValue `json:"p83075_map_virgin"  PointIgnore:"true"`
	P83076Map       valueTypes.UnitValue `json:"p83076_map" PointName:"Pv Power" PointId:"p83076" PointVirtual:"true"`
	P83076MapVirgin valueTypes.UnitValue `json:"p83076_map_virgin"  PointIgnore:"true"`
	P83077Map       valueTypes.UnitValue `json:"p83077_map" PointName:"Pv Energy" PointId:"p83077" PointVirtual:"true"`
	P83077MapVirgin valueTypes.UnitValue `json:"p83077_map_virgin"  PointIgnore:"true"`
	P83078Map       valueTypes.UnitValue `json:"p83078_map" PointId:"p83078" PointVirtual:"true"`
	P83078MapVirgin valueTypes.UnitValue `json:"p83078_map_virgin"  PointIgnore:"true"`
	P83079Map       valueTypes.UnitValue `json:"p83079_map" PointId:"p83079" PointVirtual:"true"`
	P83079MapVirgin valueTypes.UnitValue `json:"p83079_map_virgin"  PointIgnore:"true"`
	P83080Map       valueTypes.UnitValue `json:"p83080_map" PointId:"p83080" PointVirtual:"true"`
	P83080MapVirgin valueTypes.UnitValue `json:"p83080_map_virgin"  PointIgnore:"true"`
	P83088Map       valueTypes.UnitValue `json:"p83088_map" PointName:"ES Energy" PointId:"p83088" PointVirtual:"true"`
	P83088MapVirgin valueTypes.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83089Map       valueTypes.UnitValue `json:"p83089_map" PointName:"ES Discharge Energy" PointId:"p83089" PointVirtual:"true"`
	P83089MapVirgin valueTypes.UnitValue `json:"p83089_map_virgin"  PointIgnore:"true"`
	P83094Map       valueTypes.UnitValue `json:"p83094_map" PointId:"p83094" PointVirtual:"true"`
	P83094MapVirgin valueTypes.UnitValue `json:"p83094_map_virgin"  PointIgnore:"true"`
	P83095Map       valueTypes.UnitValue `json:"p83095_map" PointId:"p83095" PointName:"ES Total Discharge Energy" PointUpdateFreq:"UpdateFreqTotal" PointVirtual:"true"`
	P83095MapVirgin valueTypes.UnitValue `json:"p83095_map_virgin"  PointIgnore:"true"`
	P83097Map       valueTypes.UnitValue `json:"p83097_map" PointId:"p83097" PointVirtual:"true"`
	P83097MapVirgin valueTypes.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83100Map       valueTypes.UnitValue `json:"p83100_map" PointId:"p83100" PointVirtual:"true"`
	P83100MapVirgin valueTypes.UnitValue `json:"p83100_map_virgin"  PointIgnore:"true"`
	P83101Map       valueTypes.UnitValue `json:"p83101_map" PointId:"p83101" PointVirtual:"true"`
	P83101MapVirgin valueTypes.UnitValue `json:"p83101_map_virgin"  PointIgnore:"true"`
	P83102Map       valueTypes.UnitValue `json:"p83102_map" PointId:"p83102" PointName:"Energy Purchased" PointVirtual:"true"`
	P83102MapVirgin valueTypes.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83102Percent   valueTypes.Float     `json:"p83102_percent" PointId:"p83102_percent" PointName:"Energy Purchased Percent" PointUnit:"%" PointVirtual:"true"`
	P83105Map       valueTypes.UnitValue `json:"p83105_map" PointId:"p83105" PointVirtual:"true"`
	P83105MapVirgin valueTypes.UnitValue `json:"p83105_map_virgin"  PointIgnore:"true"`
	P83106Map       valueTypes.UnitValue `json:"p83106_map" PointId:"p83106" PointVirtual:"true"`
	P83106MapVirgin valueTypes.UnitValue `json:"p83106_map_virgin"  PointIgnore:"true"`
	P83107Map       valueTypes.UnitValue `json:"p83107_map" PointId:"p83107" PointVirtual:"true"`
	P83107MapVirgin valueTypes.UnitValue `json:"p83107_map_virgin"  PointIgnore:"true"`
	P83118Map       valueTypes.UnitValue `json:"p83118_map" PointId:"p83118" PointName:"Energy Used" PointVirtual:"true"`
	P83118MapVirgin valueTypes.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map       valueTypes.UnitValue `json:"p83119_map" PointId:"p83119" PointName:"Energy Feed-In" PointVirtual:"true"`
	P83119MapVirgin valueTypes.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map       valueTypes.UnitValue `json:"p83120_map" PointId:"p83120" PointName:"Energy Battery Charge" PointVirtual:"true"`
	P83120MapVirgin valueTypes.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121          valueTypes.Float     `json:"p83121" PointId:"p83121" PointVirtual:"true"`
	P83122          valueTypes.Float     `json:"p83122" PointName:"Self Sufficiency Percent" PointUnit:"%" PointVirtual:"true" PointUpdateFreq:"UpdateFreq5Mins"`
	P83123Map       valueTypes.UnitValue `json:"p83123_map" PointId:"p83123" PointVirtual:"true"`
	P83123MapVirgin valueTypes.UnitValue `json:"p83123_map_virgin"  PointIgnore:"true"`
	P83124Map       valueTypes.UnitValue `json:"p83124_map" PointId:"p83124" PointVirtual:"true"`
	P83124MapVirgin valueTypes.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83125          valueTypes.Float     `json:"p83125" PointId:"p83125" PointVirtual:"true"`
	P83126          valueTypes.Float     `json:"p83126" PointId:"p83126" PointVirtual:"true"`
	P83128MapVirgin valueTypes.UnitValue `json:"p83128_map_virgin"  PointIgnore:"true"`
	P83202Map       valueTypes.UnitValue `json:"p83202_map" PointId:"p83202" PointName:"Installed Power" PointVirtual:"true"`
	P83202MapVirgin valueTypes.UnitValue `json:"p83202_map_virgin"  PointIgnore:"true"`
	ZfzyMap         valueTypes.UnitValue `json:"zfzy_map" PointName:"Self Consumption Of PV" PointVirtual:"true"`
	ZfzyMapVirgin   valueTypes.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap         valueTypes.UnitValue `json:"zjzz_map" PointName:"Self Sufficiency" PointVirtual:"true"`
	ZjzzMapVirgin   valueTypes.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`

	Images Common.PowerStationImages `json:"images"`

	RobotNumSweepCapacity struct {
		Num           valueTypes.Integer `json:"num"`
		SweepCapacity valueTypes.Float   `json:"sweep_capacity"`
	} `json:"robot_num_sweep_capacity"`

	StorageInverterData []struct {
		GoStruct GoStruct.GoStruct `json:"-" PointIdReplace:"true" PointIdFrom:"PsKey" PointDeviceFrom:"PsKey"`

		UpdateTime              valueTypes.DateTime  `json:"update_time" PointNameDateFormat:"DateTimeLayout"`
		PsKey                   valueTypes.PsKey     `json:"ps_key"`
		CommunicationDevSn      valueTypes.String    `json:"communication_dev_sn"`
		DevFaultStatus          valueTypes.Integer   `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer   `json:"dev_status"`
		DeviceCode              valueTypes.Integer   `json:"device_code"`
		DeviceModelCode         valueTypes.String    `json:"device_model_code"`
		DeviceName              valueTypes.String    `json:"device_name"`
		DeviceState             valueTypes.String    `json:"device_state"`
		DeviceType              valueTypes.Integer   `json:"device_type"`
		DrmStatus               valueTypes.Integer   `json:"drm_status"`
		DrmStatusName           valueTypes.String    `json:"drm_status_name"`
		HasAmmeter              valueTypes.Bool      `json:"has_ammeter"`
		InstallerDevFaultStatus valueTypes.Integer   `json:"installer_dev_fault_status"`
		InverterSn              valueTypes.String    `json:"inverter_sn"`
		OwnerDevFaultStatus     valueTypes.Integer   `json:"owner_dev_fault_status"`
		UUID                    valueTypes.Integer   `json:"uuid"`
		EnergyFlow              []valueTypes.Integer `json:"energy_flow" PointName:"Energy Flow" PointIgnoreZero:"false" PointArrayFlatten:"true"`

		P13003Map       valueTypes.UnitValue `json:"p13003_map" PointId:"p13003" PointName:"PV Power To Load" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13003MapVirgin valueTypes.UnitValue `json:"p13003_map_virgin"  PointIgnore:"true"`
		P13011Map       valueTypes.UnitValue `json:"p13011_map" PointId:"p13011" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13011MapVirgin valueTypes.UnitValue `json:"p13011_map_virgin"  PointIgnore:"true"`
		P13115Map       valueTypes.UnitValue `json:"p13115_map" PointId:"p13115" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13115MapVirgin valueTypes.UnitValue `json:"p13115_map_virgin"  PointIgnore:"true"`
		P13119Map       valueTypes.UnitValue `json:"p13119_map" PointId:"p13119" PointName:"Load Power" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13119MapVirgin valueTypes.UnitValue `json:"p13119_map_virgin"  PointIgnore:"true"`
		P13121Map       valueTypes.UnitValue `json:"p13121_map" PointId:"p13121" PointName:"PV Power To Grid" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13121MapVirgin valueTypes.UnitValue `json:"p13121_map_virgin"  PointIgnore:"true"`
		P13126Map       valueTypes.UnitValue `json:"p13126_map" PointId:"p13126" PointName:"PV Power To Battery" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13126MapVirgin valueTypes.UnitValue `json:"p13126_map_virgin"  PointIgnore:"true"`
		P13141          valueTypes.Float     `json:"p13141" PointName:"Battery Charge Percent" PointUnit:"%" PointIcon:"mdi:battery" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13142          valueTypes.Float     `json:"p13142" PointName:"Battery Health" PointUnit:"%" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13149Map       valueTypes.UnitValue `json:"p13149_map" PointId:"p13149" PointName:"Grid Power To Load" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13149MapVirgin valueTypes.UnitValue `json:"p13149_map_virgin"  PointIgnore:"true"`
		P13150Map       valueTypes.UnitValue `json:"p13150_map" PointId:"p13150" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
		P13150MapVirgin valueTypes.UnitValue `json:"p13150_map_virgin"  PointIgnore:"true"`
		P13155          valueTypes.Float     `json:"p13155" PointDeviceFrom:"PsKey" PointUpdateFreq:"UpdateFreq5Mins" PointVirtual:"true" PointVirtualShift:"3"`
	} `json:"storage_inverter_data" DataTable:"false"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
