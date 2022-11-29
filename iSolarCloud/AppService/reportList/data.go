package reportList

import (
	"GoSungrow/iSolarCloud/Common"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/powerStationService/reportList"
const Disabled = false

type RequestData struct {
	PsId       valueTypes.PsId   `json:"ps_id" required:"true"`
	ReportType valueTypes.String `json:"report_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := api.HelpReportType()
	return ret
}

type ResultData struct {
	MinDateId interface{}       `json:"min_date_id"`
	Info      Common.ReportInfo `json:"info" PointArrayFlatten:"false"`
	DataList  []DataList        `json:"dataList" PointId:"data_list" PointIdReplace:"true" DataTable:"true"`
	Total     []DataList        `json:"total" PointId:"total" PointIdReplace:"true" DataTable:"true"`
	// `PointIgnoreIfChildFromNil:"UpdateTime" PointIdFromChild:"DateId" PointNameDateFormat:"20060102" PointIdReplace:"false" DataTable:"true"`
}

type DataList struct {
	GoStruct GoStruct.GoStruct `json:"-" PointIgnoreIfNil:"UpdateTime" PointIdFromChild:"DateId" PointNameDateFormat:"20060102" PointIdReplace:"false"`

	DateId     valueTypes.DateTime `json:"date_id" PointTimestampFrom:"UpdateTime" PointNameDateFormat:"20060102"`
	PsId       valueTypes.PsId     `json:"ps_id" PointTimestampFrom:"UpdateTime"`
	TimeStamp  valueTypes.Generic  `json:"time_stamp" PointTimestampFrom:"UpdateTime"` // Sad that this alternates between string and valueTypes.Integer.
	UpdateTime valueTypes.DateTime `json:"update_time" PointTimestampFrom:"UpdateTime"`
	DeviceName interface{}         `json:"device_name" PointTimestampFrom:"UpdateTime"`
	UUID       interface{}         `json:"uuid" PointTimestampFrom:"UpdateTime"`
	Co2Reduce  valueTypes.Float    `json:"co2_reduce" PointTimestampFrom:"UpdateTime"`

	CitySubsidyCharge                 valueTypes.Float  `json:"city_subsidy_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CitySubsidyChargeOriginalUnit"`
	CitySubsidyChargeOriginalUnit     valueTypes.String `json:"city_subsidy_charge_original_unit" PointIgnore:"true"`
	CitySubsidyChargeTran             valueTypes.Float  `json:"city_subsidy_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CitySubsidyChargeUnit" PointIgnore:"true"`
	CitySubsidyChargeUnit             valueTypes.String `json:"city_subsidy_charge_unit" PointIgnore:"true"`
	CountrySubsidyCharge              valueTypes.Float  `json:"country_subsidy_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CountrySubsidyChargeOriginalUnit"`
	CountrySubsidyChargeOriginalUnit  valueTypes.String `json:"country_subsidy_charge_original_unit" PointIgnore:"true"`
	CountrySubsidyChargeTran          valueTypes.Float  `json:"country_subsidy_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CountrySubsidyChargeUnit" PointIgnore:"true"`
	CountrySubsidyChargeUnit          valueTypes.String `json:"country_subsidy_charge_unit" PointIgnore:"true"`
	CountySubsidyCharge               valueTypes.Float  `json:"county_subsidy_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CountySubsidyChargeOriginalUnit"`
	CountySubsidyChargeOriginalUnit   valueTypes.String `json:"county_subsidy_charge_original_unit" PointIgnore:"true"`
	CountySubsidyChargeTran           valueTypes.Float  `json:"county_subsidy_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CountySubsidyChargeUnit" PointIgnore:"true"`
	CountySubsidyChargeUnit           valueTypes.String `json:"county_subsidy_charge_unit" PointIgnore:"true"`
	CuspCharge                        valueTypes.Float  `json:"cusp_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspChargeOriginalUnit"`
	CuspChargeOriginalUnit            valueTypes.String `json:"cusp_charge_original_unit" PointIgnore:"true"`
	CuspChargeTran                    valueTypes.Float  `json:"cusp_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspChargeUnit" PointIgnore:"true"`
	CuspChargeUnit                    valueTypes.String `json:"cusp_charge_unit" PointIgnore:"true"`
	FlatCharge                        valueTypes.Float  `json:"flat_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatChargeOriginalUnit"`
	FlatChargeOriginalUnit            valueTypes.String `json:"flat_charge_original_unit" PointIgnore:"true"`
	FlatChargeTran                    valueTypes.Float  `json:"flat_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatChargeUnit"`
	FlatChargeUnit                    valueTypes.String `json:"flat_charge_unit" PointIgnore:"true"`
	PeakCharge                        valueTypes.Float  `json:"peak_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakChargeOriginalUnit"`
	PeakChargeOriginalUnit            valueTypes.String `json:"peak_charge_original_unit" PointIgnore:"true"`
	PeakChargeTran                    valueTypes.Float  `json:"peak_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakChargeUnit"`
	PeakChargeUnit                    valueTypes.String `json:"peak_charge_unit" PointIgnore:"true"`
	ProvinceSubsidyCharge             valueTypes.Float  `json:"province_subsidy_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ProvinceSubsidyChargeOriginalUnit"`
	ProvinceSubsidyChargeOriginalUnit valueTypes.String `json:"province_subsidy_charge_original_unit" PointIgnore:"true"`
	ProvinceSubsidyChargeTran         valueTypes.Float  `json:"province_subsidy_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ProvinceSubsidyChargeUnit"`
	ProvinceSubsidyChargeUnit         valueTypes.String `json:"province_subsidy_charge_unit" PointIgnore:"true"`
	ValleyCharge                      valueTypes.Float  `json:"valley_charge" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyChargeOriginalUnit"`
	ValleyChargeOriginalUnit          valueTypes.String `json:"valley_charge_original_unit" PointIgnore:"true"`
	ValleyChargeTran                  valueTypes.Float  `json:"valley_charge_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyChargeUnit" PointIgnore:"true"`
	ValleyChargeUnit                  valueTypes.String `json:"valley_charge_unit" PointIgnore:"true"`

	NetPowerProfit                       valueTypes.Float  `json:"net_power_profit" PointTimestampFrom:"UpdateTime" PointUnitFrom:"NetPowerProfitOriginalUnit"`
	NetPowerProfitOriginalUnit           valueTypes.String `json:"net_power_profit_original_unit" PointIgnore:"true"`
	NetPowerProfitTran                   valueTypes.Float  `json:"net_power_profit_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"NetPowerProfitUnit" PointIgnore:"true"`
	NetPowerProfitUnit                   valueTypes.String `json:"net_power_profit_unit" PointIgnore:"true"`
	SubsidyProfit                        valueTypes.Float  `json:"subsidy_profit" PointTimestampFrom:"UpdateTime" PointUnitFrom:"SubsidyProfitOriginalUnit"`
	SubsidyProfitOriginalUnit            valueTypes.String `json:"subsidy_profit_original_unit" PointIgnore:"true"`
	SubsidyProfitTran                    valueTypes.Float  `json:"subsidy_profit_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"SubsidyProfitUnit"`
	SubsidyProfitUnit                    valueTypes.String `json:"subsidy_profit_unit" PointIgnore:"true"`
	TotalProfit                          valueTypes.Float  `json:"total_profit" PointTimestampFrom:"UpdateTime" PointUnitFrom:"TotalProfitOriginalUnit"`
	TotalProfitOriginalUnit              valueTypes.String `json:"total_profit_original_unit" PointIgnore:"true"`
	TotalProfitTran                      valueTypes.Float  `json:"total_profit_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"TotalProfitUnit" PointIgnore:"true"`
	TotalProfitUnit                      valueTypes.String `json:"total_profit_unit" PointIgnore:"true"`
	UsePowerByDiscountProfit             valueTypes.Float  `json:"use_power_by_discount_profit" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerByDiscountProfitOriginalUnit"`
	UsePowerByDiscountProfitOriginalUnit valueTypes.String `json:"use_power_by_discount_profit_original_unit" PointIgnore:"true"`
	UsePowerByDiscountProfitTran         valueTypes.Float  `json:"use_power_by_discount_profit_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerByDiscountProfitUnit" PointIgnore:"true"`
	UsePowerByDiscountProfitUnit         valueTypes.String `json:"use_power_by_discount_profit_unit" PointIgnore:"true"`
	UsePowerProfit                       valueTypes.Float  `json:"use_power_profit" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerProfitOriginalUnit"`
	UsePowerProfitOriginalUnit           valueTypes.String `json:"use_power_profit_original_unit" PointIgnore:"true"`
	UsePowerProfitTran                   valueTypes.Float  `json:"use_power_profit_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerProfitUnit" PointIgnore:"true"`
	UsePowerProfitUnit                   valueTypes.String `json:"use_power_profit_unit" PointIgnore:"true"`

	CuspNetPowerQuantity       valueTypes.Float  `json:"cusp_net_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspNetPowerQuantityUnit"`
	CuspNetPowerQuantityTran   valueTypes.Float  `json:"cusp_net_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspNetPowerQuantityUnit" PointIgnore:"true"`
	CuspNetPowerQuantityUnit   valueTypes.String `json:"cusp_net_power_quantity_unit" PointIgnore:"true"`
	CuspPowerQuantity          valueTypes.Float  `json:"cusp_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspPowerQuantityUnit"`
	CuspPowerQuantityTran      valueTypes.Float  `json:"cusp_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspPowerQuantityUnit" PointIgnore:"true"`
	CuspPowerQuantityUnit      valueTypes.String `json:"cusp_power_quantity_unit" PointIgnore:"true"`
	CuspUsePowerQuantity       valueTypes.Float  `json:"cusp_use_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspUsePowerQuantityUnit"`
	CuspUsePowerQuantityTran   valueTypes.Float  `json:"cusp_use_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"CuspUsePowerQuantityUnit" PointIgnore:"true"`
	CuspUsePowerQuantityUnit   valueTypes.String `json:"cusp_use_power_quantity_unit" PointIgnore:"true"`
	FlatNetPowerQuantity       valueTypes.Float  `json:"flat_net_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatNetPowerQuantityUnit"`
	FlatNetPowerQuantityTran   valueTypes.Float  `json:"flat_net_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatNetPowerQuantityUnit" PointIgnore:"true"`
	FlatNetPowerQuantityUnit   valueTypes.String `json:"flat_net_power_quantity_unit" PointIgnore:"true"`
	FlatPowerQuantity          valueTypes.Float  `json:"flat_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatPowerQuantityUnit"`
	FlatPowerQuantityTran      valueTypes.Float  `json:"flat_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatPowerQuantityUnit" PointIgnore:"true"`
	FlatPowerQuantityUnit      valueTypes.String `json:"flat_power_quantity_unit" PointIgnore:"true"`
	FlatUsePowerQuantity       valueTypes.Float  `json:"flat_use_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatUsePowerQuantityUnit"`
	FlatUsePowerQuantityTran   valueTypes.Float  `json:"flat_use_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"FlatUsePowerQuantityUnit" PointIgnore:"true"`
	FlatUsePowerQuantityUnit   valueTypes.String `json:"flat_use_power_quantity_unit" PointIgnore:"true"`
	NetPowerQuantityTotal      valueTypes.Float  `json:"net_power_quantity_total" PointTimestampFrom:"UpdateTime" PointUnitFrom:"NetPowerQuantityTotalUnit"`
	NetPowerQuantityTotalTran  valueTypes.Float  `json:"net_power_quantity_total_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"NetPowerQuantityTotalUnit" PointIgnore:"true"`
	NetPowerQuantityTotalUnit  valueTypes.String `json:"net_power_quantity_total_unit" PointIgnore:"true"`
	PeakNetPowerQuantity       valueTypes.Float  `json:"peak_net_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakNetPowerQuantityUnit"`
	PeakNetPowerQuantityTran   valueTypes.Float  `json:"peak_net_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakNetPowerQuantityUnit" PointIgnore:"true"`
	PeakNetPowerQuantityUnit   valueTypes.String `json:"peak_net_power_quantity_unit" PointIgnore:"true"`
	PeakPowerQuantity          valueTypes.Float  `json:"peak_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakPowerQuantityUnit"`
	PeakPowerQuantityTran      valueTypes.Float  `json:"peak_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakPowerQuantityUnit" PointIgnore:"true"`
	PeakPowerQuantityUnit      valueTypes.String `json:"peak_power_quantity_unit" PointIgnore:"true"`
	PeakUsePowerQuantity       valueTypes.Float  `json:"peak_use_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakUsePowerQuantityUnit"`
	PeakUsePowerQuantityTran   valueTypes.Float  `json:"peak_use_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PeakUsePowerQuantityUnit" PointIgnore:"true"`
	PeakUsePowerQuantityUnit   valueTypes.String `json:"peak_use_power_quantity_unit" PointIgnore:"true"`
	PowerQuantityTotal         valueTypes.Float  `json:"power_quantity_total" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PowerQuantityTotalUnit"`
	PowerQuantityTotalTran     valueTypes.Float  `json:"power_quantity_total_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"PowerQuantityTotalUnit" PointIgnore:"true"`
	PowerQuantityTotalUnit     valueTypes.String `json:"power_quantity_total_unit" PointIgnore:"true"`
	UsePowerQuantityTotal      valueTypes.Float  `json:"use_power_quantity_total" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerQuantityTotalUnit"`
	UsePowerQuantityTotalTran  valueTypes.Float  `json:"use_power_quantity_total_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"UsePowerQuantityTotalUnit" PointIgnore:"true"`
	UsePowerQuantityTotalUnit  valueTypes.String `json:"use_power_quantity_total_unit" PointIgnore:"true"`
	ValleyNetPowerQuantity     valueTypes.Float  `json:"valley_net_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyNetPowerQuantityUnit"`
	ValleyNetPowerQuantityTran valueTypes.Float  `json:"valley_net_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyNetPowerQuantityUnit" PointIgnore:"true"`
	ValleyNetPowerQuantityUnit valueTypes.String `json:"valley_net_power_quantity_unit" PointIgnore:"true"`
	ValleyPowerQuantity        valueTypes.Float  `json:"valley_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyPowerQuantityUnit"`
	ValleyPowerQuantityTran    valueTypes.Float  `json:"valley_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyPowerQuantityUnit" PointIgnore:"true"`
	ValleyPowerQuantityUnit    valueTypes.String `json:"valley_power_quantity_unit" PointIgnore:"true"`
	ValleyUsePowerQuantity     valueTypes.Float  `json:"valley_use_power_quantity" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyUsePowerQuantityUnit"`
	ValleyUsePowerQuantityTran valueTypes.Float  `json:"valley_use_power_quantity_tran" PointTimestampFrom:"UpdateTime" PointUnitFrom:"ValleyUsePowerQuantityUnit" PointIgnore:"true"`
	ValleyUsePowerQuantityUnit valueTypes.String `json:"valley_use_power_quantity_unit" PointIgnore:"true"`
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
