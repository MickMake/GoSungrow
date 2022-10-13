package getReportData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/powerStationService/getReportData"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
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
	DataList []struct {
		CitySubsidyCharge                    api.Float    `json:"city_subsidy_charge"`
		CitySubsidyChargeOriginalUnit        api.String   `json:"city_subsidy_charge_original_unit"`
		CitySubsidyChargeTran                api.Float    `json:"city_subsidy_charge_tran"`
		CitySubsidyChargeUnit                api.String   `json:"city_subsidy_charge_unit"`
		CountrySubsidyCharge                 api.Float    `json:"country_subsidy_charge"`
		CountrySubsidyChargeOriginalUnit     api.String   `json:"country_subsidy_charge_original_unit"`
		CountrySubsidyChargeTran             api.Float    `json:"country_subsidy_charge_tran"`
		CountrySubsidyChargeUnit             api.String   `json:"country_subsidy_charge_unit"`
		CountySubsidyCharge                  api.Float    `json:"county_subsidy_charge"`
		CountySubsidyChargeOriginalUnit      api.String   `json:"county_subsidy_charge_original_unit"`
		CountySubsidyChargeTran              api.Float    `json:"county_subsidy_charge_tran"`
		CountySubsidyChargeUnit              api.String   `json:"county_subsidy_charge_unit"`
		CuspCharge                           api.Float    `json:"cusp_charge"`
		CuspChargeOriginalUnit               api.String   `json:"cusp_charge_original_unit"`
		CuspChargeTran                       api.Float    `json:"cusp_charge_tran"`
		CuspChargeUnit                       api.String   `json:"cusp_charge_unit"`
		CuspNetPowerQuantity                 api.Float    `json:"cusp_net_power_quantity"`
		CuspNetPowerQuantityTran             api.Float    `json:"cusp_net_power_quantity_tran"`
		CuspNetPowerQuantityUnit             api.String   `json:"cusp_net_power_quantity_unit"`
		CuspPowerQuantity                    api.Float    `json:"cusp_power_quantity"`
		CuspPowerQuantityTran                api.Float    `json:"cusp_power_quantity_tran"`
		CuspPowerQuantityUnit                api.String   `json:"cusp_power_quantity_unit"`
		CuspUsePowerQuantity                 api.Float    `json:"cusp_use_power_quantity"`
		CuspUsePowerQuantityTran             api.Float    `json:"cusp_use_power_quantity_tran"`
		CuspUsePowerQuantityUnit             api.String   `json:"cusp_use_power_quantity_unit"`
		DateID                               api.DateTime `json:"date_id"`
		DeviceName                           api.String   `json:"device_name"`
		FlatCharge                           api.Float    `json:"flat_charge"`
		FlatChargeOriginalUnit               api.String   `json:"flat_charge_original_unit"`
		FlatChargeTran                       api.Float    `json:"flat_charge_tran"`
		FlatChargeUnit                       api.String   `json:"flat_charge_unit"`
		FlatNetPowerQuantity                 api.Float    `json:"flat_net_power_quantity"`
		FlatNetPowerQuantityTran             api.Float    `json:"flat_net_power_quantity_tran"`
		FlatNetPowerQuantityUnit             api.String   `json:"flat_net_power_quantity_unit"`
		FlatPowerQuantity                    api.Float    `json:"flat_power_quantity"`
		FlatPowerQuantityTran                api.Float    `json:"flat_power_quantity_tran"`
		FlatPowerQuantityUnit                api.String   `json:"flat_power_quantity_unit"`
		FlatUsePowerQuantity                 api.Float    `json:"flat_use_power_quantity"`
		FlatUsePowerQuantityTran             api.Float    `json:"flat_use_power_quantity_tran"`
		FlatUsePowerQuantityUnit             api.String   `json:"flat_use_power_quantity_unit"`
		NetPowerProfit                       api.Float    `json:"net_power_profit"`
		NetPowerProfitOriginalUnit           api.String   `json:"net_power_profit_original_unit"`
		NetPowerProfitTran                   api.Float    `json:"net_power_profit_tran"`
		NetPowerProfitUnit                   api.String   `json:"net_power_profit_unit"`
		NetPowerQuantityTotal                api.Float    `json:"net_power_quantity_total"`
		NetPowerQuantityTotalTran            api.Float    `json:"net_power_quantity_total_tran"`
		NetPowerQuantityTotalUnit            api.String   `json:"net_power_quantity_total_unit"`
		PeakCharge                           api.Float    `json:"peak_charge"`
		PeakChargeOriginalUnit               api.String   `json:"peak_charge_original_unit"`
		PeakChargeTran                       api.Float    `json:"peak_charge_tran"`
		PeakChargeUnit                       api.String   `json:"peak_charge_unit"`
		PeakNetPowerQuantity                 api.Float    `json:"peak_net_power_quantity"`
		PeakNetPowerQuantityTran             api.Float    `json:"peak_net_power_quantity_tran"`
		PeakNetPowerQuantityUnit             api.String   `json:"peak_net_power_quantity_unit"`
		PeakPowerQuantity                    api.Float    `json:"peak_power_quantity"`
		PeakPowerQuantityTran                api.Float    `json:"peak_power_quantity_tran"`
		PeakPowerQuantityUnit                api.String   `json:"peak_power_quantity_unit"`
		PeakUsePowerQuantity                 api.Float    `json:"peak_use_power_quantity"`
		PeakUsePowerQuantityTran             api.Float    `json:"peak_use_power_quantity_tran"`
		PeakUsePowerQuantityUnit             api.String   `json:"peak_use_power_quantity_unit"`
		PowerQuantityTotal                   api.Float    `json:"power_quantity_total"`
		PowerQuantityTotalTran               api.Float    `json:"power_quantity_total_tran"`
		PowerQuantityTotalUnit               api.String   `json:"power_quantity_total_unit"`
		ProvinceSubsidyCharge                api.Float    `json:"province_subsidy_charge"`
		ProvinceSubsidyChargeOriginalUnit    api.String   `json:"province_subsidy_charge_original_unit"`
		ProvinceSubsidyChargeTran            api.Float    `json:"province_subsidy_charge_tran"`
		ProvinceSubsidyChargeUnit            api.String   `json:"province_subsidy_charge_unit"`
		PsID                                 api.Integer  `json:"ps_id"`
		SubsidyProfit                        api.Float    `json:"subsidy_profit"`
		SubsidyProfitOriginalUnit            api.String   `json:"subsidy_profit_original_unit"`
		SubsidyProfitTran                    api.Float    `json:"subsidy_profit_tran"`
		SubsidyProfitUnit                    api.String   `json:"subsidy_profit_unit"`
		TimeStamp                            int64        `json:"time_stamp"`
		TotalProfit                          api.Float    `json:"total_profit"`
		TotalProfitOriginalUnit              api.String   `json:"total_profit_original_unit"`
		TotalProfitTran                      api.Float    `json:"total_profit_tran"`
		TotalProfitUnit                      api.String   `json:"total_profit_unit"`
		UpdateTime                           api.DateTime `json:"update_time"`
		UsePowerByDiscountProfit             api.Float    `json:"use_power_by_discount_profit"`
		UsePowerByDiscountProfitOriginalUnit api.String   `json:"use_power_by_discount_profit_original_unit"`
		UsePowerByDiscountProfitTran         api.Float    `json:"use_power_by_discount_profit_tran"`
		UsePowerByDiscountProfitUnit         api.String   `json:"use_power_by_discount_profit_unit"`
		UsePowerProfit                       api.Float    `json:"use_power_profit"`
		UsePowerProfitOriginalUnit           api.String   `json:"use_power_profit_original_unit"`
		UsePowerProfitTran                   api.Float    `json:"use_power_profit_tran"`
		UsePowerProfitUnit                   api.String   `json:"use_power_profit_unit"`
		UsePowerQuantityTotal                api.Float    `json:"use_power_quantity_total"`
		UsePowerQuantityTotalTran            api.Float    `json:"use_power_quantity_total_tran"`
		UsePowerQuantityTotalUnit            api.String   `json:"use_power_quantity_total_unit"`
		UUID                                 interface{}  `json:"uuid"`
		ValleyCharge                         api.Float    `json:"valley_charge"`
		ValleyChargeOriginalUnit             api.String   `json:"valley_charge_original_unit"`
		ValleyChargeTran                     api.Float    `json:"valley_charge_tran"`
		ValleyChargeUnit                     api.String   `json:"valley_charge_unit"`
		ValleyNetPowerQuantity               api.Float    `json:"valley_net_power_quantity"`
		ValleyNetPowerQuantityTran           api.Float    `json:"valley_net_power_quantity_tran"`
		ValleyNetPowerQuantityUnit           api.String   `json:"valley_net_power_quantity_unit"`
		ValleyPowerQuantity                  api.Float    `json:"valley_power_quantity"`
		ValleyPowerQuantityTran              api.Float    `json:"valley_power_quantity_tran"`
		ValleyPowerQuantityUnit              api.String   `json:"valley_power_quantity_unit"`
		ValleyUsePowerQuantity               api.Float    `json:"valley_use_power_quantity"`
		ValleyUsePowerQuantityTran           api.Float    `json:"valley_use_power_quantity_tran"`
		ValleyUsePowerQuantityUnit           api.String   `json:"valley_use_power_quantity_unit"`
	} `json:"dataList"`
	Info []struct {
		DesignCapacity         api.Float   `json:"design_capacity"`
		InstallerPsFaultStatus int64       `json:"installer_ps_fault_status"`
		OwnerPsFaultStatus     int64       `json:"owner_ps_fault_status"`
		PsFaultStatus          int64       `json:"ps_fault_status"`
		PsID                   api.Integer `json:"ps_id"`
		PsName                 api.String  `json:"ps_name"`
		PsStatus               int64       `json:"ps_status"`
		PsType                 api.Integer `json:"ps_type"`
		PsTypeName             api.String  `json:"ps_type_name"`
		SysScheme              int64       `json:"sys_scheme"`
		SysSchemeName          api.String  `json:"sys_scheme_name"`
		ValidFlag              api.Bool    `json:"valid_flag"`
	} `json:"info"`
	MinDateID interface{} `json:"min_date_id"`
	Total     []struct {
		CitySubsidyCharge                    api.Float    `json:"city_subsidy_charge"`
		CitySubsidyChargeOriginalUnit        api.String   `json:"city_subsidy_charge_original_unit"`
		CitySubsidyChargeTran                api.Float    `json:"city_subsidy_charge_tran"`
		CitySubsidyChargeUnit                api.String   `json:"city_subsidy_charge_unit"`
		CountrySubsidyCharge                 api.Float    `json:"country_subsidy_charge"`
		CountrySubsidyChargeOriginalUnit     api.String   `json:"country_subsidy_charge_original_unit"`
		CountrySubsidyChargeTran             api.Float    `json:"country_subsidy_charge_tran"`
		CountrySubsidyChargeUnit             api.String   `json:"country_subsidy_charge_unit"`
		CountySubsidyCharge                  api.Float    `json:"county_subsidy_charge"`
		CountySubsidyChargeOriginalUnit      api.String   `json:"county_subsidy_charge_original_unit"`
		CountySubsidyChargeTran              api.Float    `json:"county_subsidy_charge_tran"`
		CountySubsidyChargeUnit              api.String   `json:"county_subsidy_charge_unit"`
		CuspCharge                           api.Float    `json:"cusp_charge"`
		CuspChargeOriginalUnit               api.String   `json:"cusp_charge_original_unit"`
		CuspChargeTran                       api.Float    `json:"cusp_charge_tran"`
		CuspChargeUnit                       api.String   `json:"cusp_charge_unit"`
		CuspNetPowerQuantity                 api.Float    `json:"cusp_net_power_quantity"`
		CuspNetPowerQuantityTran             api.Float    `json:"cusp_net_power_quantity_tran"`
		CuspNetPowerQuantityUnit             api.String   `json:"cusp_net_power_quantity_unit"`
		CuspPowerQuantity                    api.Float    `json:"cusp_power_quantity"`
		CuspPowerQuantityTran                api.Float    `json:"cusp_power_quantity_tran"`
		CuspPowerQuantityUnit                api.String   `json:"cusp_power_quantity_unit"`
		CuspUsePowerQuantity                 api.Float    `json:"cusp_use_power_quantity"`
		CuspUsePowerQuantityTran             api.Float    `json:"cusp_use_power_quantity_tran"`
		CuspUsePowerQuantityUnit             api.String   `json:"cusp_use_power_quantity_unit"`
		FlatCharge                           api.Float    `json:"flat_charge"`
		FlatChargeOriginalUnit               api.String   `json:"flat_charge_original_unit"`
		FlatChargeTran                       api.Float    `json:"flat_charge_tran"`
		FlatChargeUnit                       api.String   `json:"flat_charge_unit"`
		FlatNetPowerQuantity                 api.Float    `json:"flat_net_power_quantity"`
		FlatNetPowerQuantityTran             api.Float    `json:"flat_net_power_quantity_tran"`
		FlatNetPowerQuantityUnit             api.String   `json:"flat_net_power_quantity_unit"`
		FlatPowerQuantity                    api.Float    `json:"flat_power_quantity"`
		FlatPowerQuantityTran                api.Float    `json:"flat_power_quantity_tran"`
		FlatPowerQuantityUnit                api.String   `json:"flat_power_quantity_unit"`
		FlatUsePowerQuantity                 api.Float    `json:"flat_use_power_quantity"`
		FlatUsePowerQuantityTran             api.Float    `json:"flat_use_power_quantity_tran"`
		FlatUsePowerQuantityUnit             api.String   `json:"flat_use_power_quantity_unit"`
		NetPowerProfit                       api.Float    `json:"net_power_profit"`
		NetPowerProfitOriginalUnit           api.String   `json:"net_power_profit_original_unit"`
		NetPowerProfitTran                   api.Float    `json:"net_power_profit_tran"`
		NetPowerProfitUnit                   api.String   `json:"net_power_profit_unit"`
		NetPowerQuantityTotal                api.Float    `json:"net_power_quantity_total"`
		NetPowerQuantityTotalTran            api.Float    `json:"net_power_quantity_total_tran"`
		NetPowerQuantityTotalUnit            api.String   `json:"net_power_quantity_total_unit"`
		PeakCharge                           api.Float    `json:"peak_charge"`
		PeakChargeOriginalUnit               api.String   `json:"peak_charge_original_unit"`
		PeakChargeTran                       api.Float    `json:"peak_charge_tran"`
		PeakChargeUnit                       api.String   `json:"peak_charge_unit"`
		PeakNetPowerQuantity                 api.Float    `json:"peak_net_power_quantity"`
		PeakNetPowerQuantityTran             api.Float    `json:"peak_net_power_quantity_tran"`
		PeakNetPowerQuantityUnit             api.String   `json:"peak_net_power_quantity_unit"`
		PeakPowerQuantity                    api.Float    `json:"peak_power_quantity"`
		PeakPowerQuantityTran                api.Float    `json:"peak_power_quantity_tran"`
		PeakPowerQuantityUnit                api.String   `json:"peak_power_quantity_unit"`
		PeakUsePowerQuantity                 api.Float    `json:"peak_use_power_quantity"`
		PeakUsePowerQuantityTran             api.Float    `json:"peak_use_power_quantity_tran"`
		PeakUsePowerQuantityUnit             api.String   `json:"peak_use_power_quantity_unit"`
		PowerQuantityTotal                   api.Float    `json:"power_quantity_total"`
		PowerQuantityTotalTran               api.Float    `json:"power_quantity_total_tran"`
		PowerQuantityTotalUnit               api.String   `json:"power_quantity_total_unit"`
		ProvinceSubsidyCharge                api.Float    `json:"province_subsidy_charge"`
		ProvinceSubsidyChargeOriginalUnit    api.String   `json:"province_subsidy_charge_original_unit"`
		ProvinceSubsidyChargeTran            api.Float    `json:"province_subsidy_charge_tran"`
		ProvinceSubsidyChargeUnit            api.String   `json:"province_subsidy_charge_unit"`
		PsID                                 api.Integer  `json:"ps_id"`
		SubsidyProfit                        api.Float    `json:"subsidy_profit"`
		SubsidyProfitOriginalUnit            api.String   `json:"subsidy_profit_original_unit"`
		SubsidyProfitTran                    api.Float    `json:"subsidy_profit_tran"`
		SubsidyProfitUnit                    api.String   `json:"subsidy_profit_unit"`
		TotalProfit                          api.Float    `json:"total_profit"`
		TotalProfitOriginalUnit              api.String   `json:"total_profit_original_unit"`
		TotalProfitTran                      api.Float    `json:"total_profit_tran"`
		TotalProfitUnit                      api.String   `json:"total_profit_unit"`
		UpdateTime                           api.DateTime `json:"update_time"`
		UsePowerByDiscountProfit             api.Float    `json:"use_power_by_discount_profit"`
		UsePowerByDiscountProfitOriginalUnit api.String   `json:"use_power_by_discount_profit_original_unit"`
		UsePowerByDiscountProfitTran         api.Float    `json:"use_power_by_discount_profit_tran"`
		UsePowerByDiscountProfitUnit         api.String   `json:"use_power_by_discount_profit_unit"`
		UsePowerProfit                       api.Float    `json:"use_power_profit"`
		UsePowerProfitOriginalUnit           api.String   `json:"use_power_profit_original_unit"`
		UsePowerProfitTran                   api.Float    `json:"use_power_profit_tran"`
		UsePowerProfitUnit                   api.String   `json:"use_power_profit_unit"`
		UsePowerQuantityTotal                api.Float    `json:"use_power_quantity_total"`
		UsePowerQuantityTotalTran            api.Float    `json:"use_power_quantity_total_tran"`
		UsePowerQuantityTotalUnit            api.String   `json:"use_power_quantity_total_unit"`
		ValleyCharge                         api.Float    `json:"valley_charge"`
		ValleyChargeOriginalUnit             api.String   `json:"valley_charge_original_unit"`
		ValleyChargeTran                     api.Float    `json:"valley_charge_tran"`
		ValleyChargeUnit                     api.String   `json:"valley_charge_unit"`
		ValleyNetPowerQuantity               api.Float    `json:"valley_net_power_quantity"`
		ValleyNetPowerQuantityTran           api.Float    `json:"valley_net_power_quantity_tran"`
		ValleyNetPowerQuantityUnit           api.String   `json:"valley_net_power_quantity_unit"`
		ValleyPowerQuantity                  api.Float    `json:"valley_power_quantity"`
		ValleyPowerQuantityTran              api.Float    `json:"valley_power_quantity_tran"`
		ValleyPowerQuantityUnit              api.String   `json:"valley_power_quantity_unit"`
		ValleyUsePowerQuantity               api.Float    `json:"valley_use_power_quantity"`
		ValleyUsePowerQuantityTran           api.Float    `json:"valley_use_power_quantity_tran"`
		ValleyUsePowerQuantityUnit           api.String   `json:"valley_use_power_quantity_unit"`
	} `json:"total"`
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

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
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
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", api.NewDateTime(""))
	}

	return entries
}
