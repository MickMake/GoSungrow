package getReportData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getReportData"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
	ReportType string `json:"report_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DataList []DataList `json:"dataList" PointId:"data_list" PointIgnoreIfNilFromChild:"UpdateTime" PointNameFromChild:"DateId" PointNameDateFormat:"20060102" PointNameFromAppend:"true"`
	Info     []struct {
		DesignCapacity         valueTypes.Float   `json:"design_capacity" PointUnit:"W"`
		InstallerPsFaultStatus valueTypes.Integer `json:"installer_ps_fault_status"`
		OwnerPsFaultStatus     valueTypes.Integer `json:"owner_ps_fault_status"`
		PsFaultStatus          valueTypes.Integer `json:"ps_fault_status"`
		PsId                   valueTypes.Integer `json:"ps_id"`
		PsName                 valueTypes.String  `json:"ps_name"`
		PsStatus               valueTypes.Integer `json:"ps_status"`
		PsType                 valueTypes.Integer `json:"ps_type"`
		PsTypeName             valueTypes.String  `json:"ps_type_name"`
		SysScheme              valueTypes.Integer `json:"sys_scheme"`
		SysSchemeName          valueTypes.String  `json:"sys_scheme_name"`
		ValidFlag              valueTypes.Bool    `json:"valid_flag"`
	} `json:"info"`
	MinDateID interface{} `json:"min_date_id"`
	Total     []DataList  `json:"total" PointId:"total" PointIgnoreIfNilFromChild:"UpdateTime" PointNameFromChild:"DateId" PointNameDateFormat:"20060102" PointNameFromAppend:"true"`
}
type DataList struct {
	DateId     valueTypes.DateTime `json:"date_id"`
	DeviceName interface{}         `json:"device_name"`
	PsId       valueTypes.Integer  `json:"ps_id"`
	TimeStamp  interface{}         `json:"time_stamp"` // Sad that this alternates between string and valueTypes.Integer.
	UpdateTime valueTypes.DateTime `json:"update_time"`
	UUID       interface{}         `json:"uuid"`
	Co2Reduce  valueTypes.Float    `json:"co2_reduce"`

	CitySubsidyCharge             valueTypes.Float  `json:"city_subsidy_charge" PointUnitFrom:"CitySubsidyChargeOriginalUnit"`
	CitySubsidyChargeOriginalUnit valueTypes.String `json:"city_subsidy_charge_original_unit" PointIgnore:"true"`
	CitySubsidyChargeTran         valueTypes.Float  `json:"city_subsidy_charge_tran" PointUnitFrom:"CitySubsidyChargeUnit" PointIgnore:"true"`
	CitySubsidyChargeUnit         valueTypes.String `json:"city_subsidy_charge_unit" PointIgnore:"true"`

	CountrySubsidyCharge             valueTypes.Float  `json:"country_subsidy_charge" PointUnitFrom:"CountrySubsidyChargeOriginalUnit"`
	CountrySubsidyChargeOriginalUnit valueTypes.String `json:"country_subsidy_charge_original_unit" PointIgnore:"true"`
	CountrySubsidyChargeTran         valueTypes.Float  `json:"country_subsidy_charge_tran" PointUnitFrom:"CountrySubsidyChargeUnit" PointIgnore:"true"`
	CountrySubsidyChargeUnit         valueTypes.String `json:"country_subsidy_charge_unit" PointIgnore:"true"`

	CountySubsidyCharge             valueTypes.Float  `json:"county_subsidy_charge" PointUnitFrom:"CountySubsidyChargeOriginalUnit"`
	CountySubsidyChargeOriginalUnit valueTypes.String `json:"county_subsidy_charge_original_unit" PointIgnore:"true"`
	CountySubsidyChargeTran         valueTypes.Float  `json:"county_subsidy_charge_tran" PointUnitFrom:"CountySubsidyChargeUnit" PointIgnore:"true"`
	CountySubsidyChargeUnit         valueTypes.String `json:"county_subsidy_charge_unit" PointIgnore:"true"`

	CuspCharge             valueTypes.Float  `json:"cusp_charge" PointUnitFrom:"CuspChargeOriginalUnit"`
	CuspChargeOriginalUnit valueTypes.String `json:"cusp_charge_original_unit" PointIgnore:"true"`
	CuspChargeTran         valueTypes.Float  `json:"cusp_charge_tran" PointUnitFrom:"CuspChargeUnit" PointIgnore:"true"`
	CuspChargeUnit         valueTypes.String `json:"cusp_charge_unit" PointIgnore:"true"`

	CuspNetPowerQuantity     valueTypes.Float  `json:"cusp_net_power_quantity" PointUnitFrom:"CuspNetPowerQuantityUnit"`
	CuspNetPowerQuantityTran valueTypes.Float  `json:"cusp_net_power_quantity_tran" PointUnitFrom:"CuspNetPowerQuantityUnit" PointIgnore:"true"`
	CuspNetPowerQuantityUnit valueTypes.String `json:"cusp_net_power_quantity_unit" PointIgnore:"true"`

	CuspPowerQuantity     valueTypes.Float  `json:"cusp_power_quantity" PointUnitFrom:"CuspPowerQuantityUnit"`
	CuspPowerQuantityTran valueTypes.Float  `json:"cusp_power_quantity_tran" PointUnitFrom:"CuspPowerQuantityUnit" PointIgnore:"true"`
	CuspPowerQuantityUnit valueTypes.String `json:"cusp_power_quantity_unit" PointIgnore:"true"`

	CuspUsePowerQuantity     valueTypes.Float  `json:"cusp_use_power_quantity" PointUnitFrom:"CuspUsePowerQuantityUnit"`
	CuspUsePowerQuantityTran valueTypes.Float  `json:"cusp_use_power_quantity_tran" PointUnitFrom:"CuspUsePowerQuantityUnit" PointIgnore:"true"`
	CuspUsePowerQuantityUnit valueTypes.String `json:"cusp_use_power_quantity_unit" PointIgnore:"true"`

	FlatCharge             valueTypes.Float  `json:"flat_charge" PointUnitFrom:"FlatChargeOriginalUnit"`
	FlatChargeOriginalUnit valueTypes.String `json:"flat_charge_original_unit" PointIgnore:"true"`

	FlatChargeTran valueTypes.Float  `json:"flat_charge_tran" PointUnitFrom:"FlatChargeUnit"`
	FlatChargeUnit valueTypes.String `json:"flat_charge_unit" PointIgnore:"true"`

	FlatNetPowerQuantity     valueTypes.Float  `json:"flat_net_power_quantity" PointUnitFrom:"FlatNetPowerQuantityUnit"`
	FlatNetPowerQuantityTran valueTypes.Float  `json:"flat_net_power_quantity_tran" PointUnitFrom:"FlatNetPowerQuantityUnit" PointIgnore:"true"`
	FlatNetPowerQuantityUnit valueTypes.String `json:"flat_net_power_quantity_unit" PointIgnore:"true"`

	FlatPowerQuantity     valueTypes.Float  `json:"flat_power_quantity" PointUnitFrom:"FlatPowerQuantityUnit"`
	FlatPowerQuantityTran valueTypes.Float  `json:"flat_power_quantity_tran" PointUnitFrom:"FlatPowerQuantityUnit" PointIgnore:"true"`
	FlatPowerQuantityUnit valueTypes.String `json:"flat_power_quantity_unit" PointIgnore:"true"`

	FlatUsePowerQuantity     valueTypes.Float  `json:"flat_use_power_quantity" PointUnitFrom:"FlatUsePowerQuantityUnit"`
	FlatUsePowerQuantityTran valueTypes.Float  `json:"flat_use_power_quantity_tran" PointUnitFrom:"FlatUsePowerQuantityUnit" PointIgnore:"true"`
	FlatUsePowerQuantityUnit valueTypes.String `json:"flat_use_power_quantity_unit" PointIgnore:"true"`

	NetPowerProfit             valueTypes.Float  `json:"net_power_profit" PointUnitFrom:"NetPowerProfitOriginalUnit"`
	NetPowerProfitOriginalUnit valueTypes.String `json:"net_power_profit_original_unit" PointIgnore:"true"`
	NetPowerProfitTran         valueTypes.Float  `json:"net_power_profit_tran" PointUnitFrom:"NetPowerProfitUnit" PointIgnore:"true"`
	NetPowerProfitUnit         valueTypes.String `json:"net_power_profit_unit" PointIgnore:"true"`

	NetPowerQuantityTotal     valueTypes.Float  `json:"net_power_quantity_total" PointUnitFrom:"NetPowerQuantityTotalUnit"`
	NetPowerQuantityTotalTran valueTypes.Float  `json:"net_power_quantity_total_tran" PointUnitFrom:"NetPowerQuantityTotalUnit" PointIgnore:"true"`
	NetPowerQuantityTotalUnit valueTypes.String `json:"net_power_quantity_total_unit" PointIgnore:"true"`

	PeakCharge             valueTypes.Float  `json:"peak_charge" PointUnitFrom:"PeakChargeOriginalUnit"`
	PeakChargeOriginalUnit valueTypes.String `json:"peak_charge_original_unit" PointIgnore:"true"`

	PeakChargeTran valueTypes.Float  `json:"peak_charge_tran" PointUnitFrom:"PeakChargeUnit"`
	PeakChargeUnit valueTypes.String `json:"peak_charge_unit" PointIgnore:"true"`

	PeakNetPowerQuantity     valueTypes.Float  `json:"peak_net_power_quantity" PointUnitFrom:"PeakNetPowerQuantityUnit"`
	PeakNetPowerQuantityTran valueTypes.Float  `json:"peak_net_power_quantity_tran" PointUnitFrom:"PeakNetPowerQuantityUnit" PointIgnore:"true"`
	PeakNetPowerQuantityUnit valueTypes.String `json:"peak_net_power_quantity_unit" PointIgnore:"true"`

	PeakPowerQuantity     valueTypes.Float  `json:"peak_power_quantity" PointUnitFrom:"PeakPowerQuantityUnit"`
	PeakPowerQuantityTran valueTypes.Float  `json:"peak_power_quantity_tran" PointUnitFrom:"PeakPowerQuantityUnit" PointIgnore:"true"`
	PeakPowerQuantityUnit valueTypes.String `json:"peak_power_quantity_unit" PointIgnore:"true"`

	PeakUsePowerQuantity     valueTypes.Float  `json:"peak_use_power_quantity" PointUnitFrom:"PeakUsePowerQuantityUnit"`
	PeakUsePowerQuantityTran valueTypes.Float  `json:"peak_use_power_quantity_tran" PointUnitFrom:"PeakUsePowerQuantityUnit" PointIgnore:"true"`
	PeakUsePowerQuantityUnit valueTypes.String `json:"peak_use_power_quantity_unit" PointIgnore:"true"`

	PowerQuantityTotal     valueTypes.Float  `json:"power_quantity_total" PointUnitFrom:"PowerQuantityTotalUnit"`
	PowerQuantityTotalTran valueTypes.Float  `json:"power_quantity_total_tran" PointUnitFrom:"PowerQuantityTotalUnit" PointIgnore:"true"`
	PowerQuantityTotalUnit valueTypes.String `json:"power_quantity_total_unit" PointIgnore:"true"`

	ProvinceSubsidyCharge             valueTypes.Float  `json:"province_subsidy_charge" PointUnitFrom:"ProvinceSubsidyChargeOriginalUnit"`
	ProvinceSubsidyChargeOriginalUnit valueTypes.String `json:"province_subsidy_charge_original_unit" PointIgnore:"true"`

	ProvinceSubsidyChargeTran valueTypes.Float  `json:"province_subsidy_charge_tran" PointUnitFrom:"ProvinceSubsidyChargeUnit"`
	ProvinceSubsidyChargeUnit valueTypes.String `json:"province_subsidy_charge_unit" PointIgnore:"true"`

	SubsidyProfit             valueTypes.Float  `json:"subsidy_profit" PointUnitFrom:"SubsidyProfitOriginalUnit"`
	SubsidyProfitOriginalUnit valueTypes.String `json:"subsidy_profit_original_unit" PointIgnore:"true"`

	SubsidyProfitTran valueTypes.Float  `json:"subsidy_profit_tran" PointUnitFrom:"SubsidyProfitUnit"`
	SubsidyProfitUnit valueTypes.String `json:"subsidy_profit_unit" PointIgnore:"true"`

	TotalProfit             valueTypes.Float  `json:"total_profit" PointUnitFrom:"TotalProfitOriginalUnit"`
	TotalProfitOriginalUnit valueTypes.String `json:"total_profit_original_unit" PointIgnore:"true"`
	TotalProfitTran         valueTypes.Float  `json:"total_profit_tran" PointUnitFrom:"TotalProfitUnit" PointIgnore:"true"`
	TotalProfitUnit         valueTypes.String `json:"total_profit_unit" PointIgnore:"true"`

	UsePowerByDiscountProfit             valueTypes.Float  `json:"use_power_by_discount_profit" PointUnitFrom:"UsePowerByDiscountProfitOriginalUnit"`
	UsePowerByDiscountProfitOriginalUnit valueTypes.String `json:"use_power_by_discount_profit_original_unit" PointIgnore:"true"`
	UsePowerByDiscountProfitTran         valueTypes.Float  `json:"use_power_by_discount_profit_tran" PointUnitFrom:"UsePowerByDiscountProfitUnit" PointIgnore:"true"`
	UsePowerByDiscountProfitUnit         valueTypes.String `json:"use_power_by_discount_profit_unit" PointIgnore:"true"`

	UsePowerProfit             valueTypes.Float  `json:"use_power_profit" PointUnitFrom:"UsePowerProfitOriginalUnit"`
	UsePowerProfitOriginalUnit valueTypes.String `json:"use_power_profit_original_unit" PointIgnore:"true"`
	UsePowerProfitTran         valueTypes.Float  `json:"use_power_profit_tran" PointUnitFrom:"UsePowerProfitUnit" PointIgnore:"true"`
	UsePowerProfitUnit         valueTypes.String `json:"use_power_profit_unit" PointIgnore:"true"`

	UsePowerQuantityTotal     valueTypes.Float  `json:"use_power_quantity_total" PointUnitFrom:"UsePowerQuantityTotalUnit"`
	UsePowerQuantityTotalTran valueTypes.Float  `json:"use_power_quantity_total_tran" PointUnitFrom:"UsePowerQuantityTotalUnit" PointIgnore:"true"`
	UsePowerQuantityTotalUnit valueTypes.String `json:"use_power_quantity_total_unit" PointIgnore:"true"`

	ValleyCharge             valueTypes.Float  `json:"valley_charge" PointUnitFrom:"ValleyChargeOriginalUnit"`
	ValleyChargeOriginalUnit valueTypes.String `json:"valley_charge_original_unit" PointIgnore:"true"`
	ValleyChargeTran         valueTypes.Float  `json:"valley_charge_tran" PointUnitFrom:"ValleyChargeUnit" PointIgnore:"true"`
	ValleyChargeUnit         valueTypes.String `json:"valley_charge_unit" PointIgnore:"true"`

	ValleyNetPowerQuantity     valueTypes.Float  `json:"valley_net_power_quantity" PointUnitFrom:"ValleyNetPowerQuantityUnit"`
	ValleyNetPowerQuantityTran valueTypes.Float  `json:"valley_net_power_quantity_tran" PointUnitFrom:"ValleyNetPowerQuantityUnit" PointIgnore:"true"`
	ValleyNetPowerQuantityUnit valueTypes.String `json:"valley_net_power_quantity_unit" PointIgnore:"true"`

	ValleyPowerQuantity     valueTypes.Float  `json:"valley_power_quantity" PointUnitFrom:"ValleyPowerQuantityUnit"`
	ValleyPowerQuantityTran valueTypes.Float  `json:"valley_power_quantity_tran" PointUnitFrom:"ValleyPowerQuantityUnit" PointIgnore:"true"`
	ValleyPowerQuantityUnit valueTypes.String `json:"valley_power_quantity_unit" PointIgnore:"true"`

	ValleyUsePowerQuantity     valueTypes.Float  `json:"valley_use_power_quantity" PointUnitFrom:"ValleyUsePowerQuantityUnit"`
	ValleyUsePowerQuantityTran valueTypes.Float  `json:"valley_use_power_quantity_tran" PointUnitFrom:"ValleyUsePowerQuantityUnit" PointIgnore:"true"`
	ValleyUsePowerQuantityUnit valueTypes.String `json:"valley_use_power_quantity_unit" PointIgnore:"true"`
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
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		name := pkg + "." + e.Request.PsId.String()
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), dt)
	}

	return entries
}
