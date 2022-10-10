package reportList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/reportList"
const Disabled = false

type RequestData struct {
	PsId       api.Integer `json:"ps_id" required:"true"`
	ReportType string `json:"report_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("report_type: ")
	return ret
}

type ResultData struct {
	DataList []struct {
		CitySubsidyCharge                    string      `json:"city_subsidy_charge"`
		CitySubsidyChargeOriginalUnit        string      `json:"city_subsidy_charge_original_unit"`
		CitySubsidyChargeTran                string      `json:"city_subsidy_charge_tran"`
		CitySubsidyChargeUnit                string      `json:"city_subsidy_charge_unit"`
		Co2Reduce                            api.Float   `json:"co2_reduce"`
		CountrySubsidyCharge                 string      `json:"country_subsidy_charge"`
		CountrySubsidyChargeOriginalUnit     string      `json:"country_subsidy_charge_original_unit"`
		CountrySubsidyChargeTran             string      `json:"country_subsidy_charge_tran"`
		CountrySubsidyChargeUnit             string      `json:"country_subsidy_charge_unit"`
		CountySubsidyCharge                  string      `json:"county_subsidy_charge"`
		CountySubsidyChargeOriginalUnit      string      `json:"county_subsidy_charge_original_unit"`
		CountySubsidyChargeTran              string      `json:"county_subsidy_charge_tran"`
		CountySubsidyChargeUnit              string      `json:"county_subsidy_charge_unit"`
		CuspCharge                           string      `json:"cusp_charge"`
		CuspChargeOriginalUnit               string      `json:"cusp_charge_original_unit"`
		CuspChargeTran                       string      `json:"cusp_charge_tran"`
		CuspChargeUnit                       string      `json:"cusp_charge_unit"`
		CuspNetPowerQuantity                 string      `json:"cusp_net_power_quantity"`
		CuspNetPowerQuantityTran             string      `json:"cusp_net_power_quantity_tran"`
		CuspNetPowerQuantityUnit             string      `json:"cusp_net_power_quantity_unit"`
		CuspPowerQuantity                    string      `json:"cusp_power_quantity"`
		CuspPowerQuantityTran                string      `json:"cusp_power_quantity_tran"`
		CuspPowerQuantityUnit                string      `json:"cusp_power_quantity_unit"`
		CuspUsePowerQuantity                 string      `json:"cusp_use_power_quantity"`
		CuspUsePowerQuantityTran             string      `json:"cusp_use_power_quantity_tran"`
		CuspUsePowerQuantityUnit             string      `json:"cusp_use_power_quantity_unit"`
		DateID                               string      `json:"date_id"`
		DeviceName                           interface{} `json:"device_name"`
		FlatCharge                           string      `json:"flat_charge"`
		FlatChargeOriginalUnit               string      `json:"flat_charge_original_unit"`
		FlatChargeTran                       string      `json:"flat_charge_tran"`
		FlatChargeUnit                       string      `json:"flat_charge_unit"`
		FlatNetPowerQuantity                 string      `json:"flat_net_power_quantity"`
		FlatNetPowerQuantityTran             string      `json:"flat_net_power_quantity_tran"`
		FlatNetPowerQuantityUnit             string      `json:"flat_net_power_quantity_unit"`
		FlatPowerQuantity                    string      `json:"flat_power_quantity"`
		FlatPowerQuantityTran                string      `json:"flat_power_quantity_tran"`
		FlatPowerQuantityUnit                string      `json:"flat_power_quantity_unit"`
		FlatUsePowerQuantity                 string      `json:"flat_use_power_quantity"`
		FlatUsePowerQuantityTran             string      `json:"flat_use_power_quantity_tran"`
		FlatUsePowerQuantityUnit             string      `json:"flat_use_power_quantity_unit"`
		NetPowerProfit                       string      `json:"net_power_profit"`
		NetPowerProfitOriginalUnit           string      `json:"net_power_profit_original_unit"`
		NetPowerProfitTran                   string      `json:"net_power_profit_tran"`
		NetPowerProfitUnit                   string      `json:"net_power_profit_unit"`
		NetPowerQuantityTotal                string      `json:"net_power_quantity_total"`
		NetPowerQuantityTotalTran            string      `json:"net_power_quantity_total_tran"`
		NetPowerQuantityTotalUnit            string      `json:"net_power_quantity_total_unit"`
		PeakCharge                           string      `json:"peak_charge"`
		PeakChargeOriginalUnit               string      `json:"peak_charge_original_unit"`
		PeakChargeTran                       string      `json:"peak_charge_tran"`
		PeakChargeUnit                       string      `json:"peak_charge_unit"`
		PeakNetPowerQuantity                 string      `json:"peak_net_power_quantity"`
		PeakNetPowerQuantityTran             string      `json:"peak_net_power_quantity_tran"`
		PeakNetPowerQuantityUnit             string      `json:"peak_net_power_quantity_unit"`
		PeakPowerQuantity                    string      `json:"peak_power_quantity"`
		PeakPowerQuantityTran                string      `json:"peak_power_quantity_tran"`
		PeakPowerQuantityUnit                string      `json:"peak_power_quantity_unit"`
		PeakUsePowerQuantity                 string      `json:"peak_use_power_quantity"`
		PeakUsePowerQuantityTran             string      `json:"peak_use_power_quantity_tran"`
		PeakUsePowerQuantityUnit             string      `json:"peak_use_power_quantity_unit"`
		PowerQuantityTotal                   string      `json:"power_quantity_total"`
		PowerQuantityTotalTran               string      `json:"power_quantity_total_tran"`
		PowerQuantityTotalUnit               string      `json:"power_quantity_total_unit"`
		ProvinceSubsidyCharge                string      `json:"province_subsidy_charge"`
		ProvinceSubsidyChargeOriginalUnit    string      `json:"province_subsidy_charge_original_unit"`
		ProvinceSubsidyChargeTran            string      `json:"province_subsidy_charge_tran"`
		ProvinceSubsidyChargeUnit            string      `json:"province_subsidy_charge_unit"`
		PsID                                 api.Integer `json:"ps_id"`
		SubsidyProfit                        string      `json:"subsidy_profit"`
		SubsidyProfitOriginalUnit            string      `json:"subsidy_profit_original_unit"`
		SubsidyProfitTran                    string      `json:"subsidy_profit_tran"`
		SubsidyProfitUnit                    string      `json:"subsidy_profit_unit"`
		TimeStamp                            interface{} `json:"time_stamp"`		// Sad that this alternates between string and api.Integer.
		TotalProfit                          string      `json:"total_profit"`
		TotalProfitOriginalUnit              string      `json:"total_profit_original_unit"`
		TotalProfitTran                      string      `json:"total_profit_tran"`
		TotalProfitUnit                      string      `json:"total_profit_unit"`
		UpdateTime                           string      `json:"update_time"`
		UsePowerByDiscountProfit             string      `json:"use_power_by_discount_profit"`
		UsePowerByDiscountProfitOriginalUnit string      `json:"use_power_by_discount_profit_original_unit"`
		UsePowerByDiscountProfitTran         string      `json:"use_power_by_discount_profit_tran"`
		UsePowerByDiscountProfitUnit         string      `json:"use_power_by_discount_profit_unit"`
		UsePowerProfit                       string      `json:"use_power_profit"`
		UsePowerProfitOriginalUnit           string      `json:"use_power_profit_original_unit"`
		UsePowerProfitTran                   string      `json:"use_power_profit_tran"`
		UsePowerProfitUnit                   string      `json:"use_power_profit_unit"`
		UsePowerQuantityTotal                string      `json:"use_power_quantity_total"`
		UsePowerQuantityTotalTran            string      `json:"use_power_quantity_total_tran"`
		UsePowerQuantityTotalUnit            string      `json:"use_power_quantity_total_unit"`
		UUID                                 interface{} `json:"uuid"`
		ValleyCharge                         string      `json:"valley_charge"`
		ValleyChargeOriginalUnit             string      `json:"valley_charge_original_unit"`
		ValleyChargeTran                     string      `json:"valley_charge_tran"`
		ValleyChargeUnit                     string      `json:"valley_charge_unit"`
		ValleyNetPowerQuantity               string      `json:"valley_net_power_quantity"`
		ValleyNetPowerQuantityTran           string      `json:"valley_net_power_quantity_tran"`
		ValleyNetPowerQuantityUnit           string      `json:"valley_net_power_quantity_unit"`
		ValleyPowerQuantity                  string      `json:"valley_power_quantity"`
		ValleyPowerQuantityTran              string      `json:"valley_power_quantity_tran"`
		ValleyPowerQuantityUnit              string      `json:"valley_power_quantity_unit"`
		ValleyUsePowerQuantity               string      `json:"valley_use_power_quantity"`
		ValleyUsePowerQuantityTran           string      `json:"valley_use_power_quantity_tran"`
		ValleyUsePowerQuantityUnit           string      `json:"valley_use_power_quantity_unit"`
	} `json:"dataList"`
	Info []struct {
		DesignCapacity         api.Float `json:"design_capacity"`
		InstallerPsFaultStatus api.Integer   `json:"installer_ps_fault_status"`
		OwnerPsFaultStatus     api.Integer   `json:"owner_ps_fault_status"`
		PsFaultStatus          api.Integer   `json:"ps_fault_status"`
		PsID                   api.Integer   `json:"ps_id"`
		PsName                 string  `json:"ps_name"`
		PsStatus               api.Integer   `json:"ps_status"`
		PsType                 api.Integer   `json:"ps_type"`
		PsTypeName             string  `json:"ps_type_name"`
		SysScheme              api.Integer   `json:"sys_scheme"`
		SysSchemeName          string  `json:"sys_scheme_name"`
		ValidFlag              api.Integer   `json:"valid_flag"`
	} `json:"info"`
	MinDateID interface{} `json:"min_date_id"`
	Total     []struct {
		CitySubsidyCharge                    string `json:"city_subsidy_charge"`
		CitySubsidyChargeOriginalUnit        string `json:"city_subsidy_charge_original_unit"`
		CitySubsidyChargeTran                string `json:"city_subsidy_charge_tran"`
		CitySubsidyChargeUnit                string `json:"city_subsidy_charge_unit"`
		CountrySubsidyCharge                 string `json:"country_subsidy_charge"`
		CountrySubsidyChargeOriginalUnit     string `json:"country_subsidy_charge_original_unit"`
		CountrySubsidyChargeTran             string `json:"country_subsidy_charge_tran"`
		CountrySubsidyChargeUnit             string `json:"country_subsidy_charge_unit"`
		CountySubsidyCharge                  string `json:"county_subsidy_charge"`
		CountySubsidyChargeOriginalUnit      string `json:"county_subsidy_charge_original_unit"`
		CountySubsidyChargeTran              string `json:"county_subsidy_charge_tran"`
		CountySubsidyChargeUnit              string `json:"county_subsidy_charge_unit"`
		CuspCharge                           string `json:"cusp_charge"`
		CuspChargeOriginalUnit               string `json:"cusp_charge_original_unit"`
		CuspChargeTran                       string `json:"cusp_charge_tran"`
		CuspChargeUnit                       string `json:"cusp_charge_unit"`
		CuspNetPowerQuantity                 string `json:"cusp_net_power_quantity"`
		CuspNetPowerQuantityTran             string `json:"cusp_net_power_quantity_tran"`
		CuspNetPowerQuantityUnit             string `json:"cusp_net_power_quantity_unit"`
		CuspPowerQuantity                    string `json:"cusp_power_quantity"`
		CuspPowerQuantityTran                string `json:"cusp_power_quantity_tran"`
		CuspPowerQuantityUnit                string `json:"cusp_power_quantity_unit"`
		CuspUsePowerQuantity                 string `json:"cusp_use_power_quantity"`
		CuspUsePowerQuantityTran             string `json:"cusp_use_power_quantity_tran"`
		CuspUsePowerQuantityUnit             string `json:"cusp_use_power_quantity_unit"`
		FlatCharge                           string `json:"flat_charge"`
		FlatChargeOriginalUnit               string `json:"flat_charge_original_unit"`
		FlatChargeTran                       string `json:"flat_charge_tran"`
		FlatChargeUnit                       string `json:"flat_charge_unit"`
		FlatNetPowerQuantity                 string `json:"flat_net_power_quantity"`
		FlatNetPowerQuantityTran             string `json:"flat_net_power_quantity_tran"`
		FlatNetPowerQuantityUnit             string `json:"flat_net_power_quantity_unit"`
		FlatPowerQuantity                    string `json:"flat_power_quantity"`
		FlatPowerQuantityTran                string `json:"flat_power_quantity_tran"`
		FlatPowerQuantityUnit                string `json:"flat_power_quantity_unit"`
		FlatUsePowerQuantity                 string `json:"flat_use_power_quantity"`
		FlatUsePowerQuantityTran             string `json:"flat_use_power_quantity_tran"`
		FlatUsePowerQuantityUnit             string `json:"flat_use_power_quantity_unit"`
		NetPowerProfit                       string `json:"net_power_profit"`
		NetPowerProfitOriginalUnit           string `json:"net_power_profit_original_unit"`
		NetPowerProfitTran                   string `json:"net_power_profit_tran"`
		NetPowerProfitUnit                   string `json:"net_power_profit_unit"`
		NetPowerQuantityTotal                string `json:"net_power_quantity_total"`
		NetPowerQuantityTotalTran            string `json:"net_power_quantity_total_tran"`
		NetPowerQuantityTotalUnit            string `json:"net_power_quantity_total_unit"`
		PeakCharge                           string `json:"peak_charge"`
		PeakChargeOriginalUnit               string `json:"peak_charge_original_unit"`
		PeakChargeTran                       string `json:"peak_charge_tran"`
		PeakChargeUnit                       string `json:"peak_charge_unit"`
		PeakNetPowerQuantity                 string `json:"peak_net_power_quantity"`
		PeakNetPowerQuantityTran             string `json:"peak_net_power_quantity_tran"`
		PeakNetPowerQuantityUnit             string `json:"peak_net_power_quantity_unit"`
		PeakPowerQuantity                    string `json:"peak_power_quantity"`
		PeakPowerQuantityTran                string `json:"peak_power_quantity_tran"`
		PeakPowerQuantityUnit                string `json:"peak_power_quantity_unit"`
		PeakUsePowerQuantity                 string `json:"peak_use_power_quantity"`
		PeakUsePowerQuantityTran             string `json:"peak_use_power_quantity_tran"`
		PeakUsePowerQuantityUnit             string `json:"peak_use_power_quantity_unit"`
		PowerQuantityTotal                   string `json:"power_quantity_total"`
		PowerQuantityTotalTran               string `json:"power_quantity_total_tran"`
		PowerQuantityTotalUnit               string `json:"power_quantity_total_unit"`
		ProvinceSubsidyCharge                string `json:"province_subsidy_charge"`
		ProvinceSubsidyChargeOriginalUnit    string `json:"province_subsidy_charge_original_unit"`
		ProvinceSubsidyChargeTran            string `json:"province_subsidy_charge_tran"`
		ProvinceSubsidyChargeUnit            string `json:"province_subsidy_charge_unit"`
		PsID                                 api.Integer  `json:"ps_id"`
		SubsidyProfit                        string `json:"subsidy_profit"`
		SubsidyProfitOriginalUnit            string `json:"subsidy_profit_original_unit"`
		SubsidyProfitTran                    string `json:"subsidy_profit_tran"`
		SubsidyProfitUnit                    string `json:"subsidy_profit_unit"`
		TotalProfit                          string `json:"total_profit"`
		TotalProfitOriginalUnit              string `json:"total_profit_original_unit"`
		TotalProfitTran                      string `json:"total_profit_tran"`
		TotalProfitUnit                      string `json:"total_profit_unit"`
		UpdateTime                           string `json:"update_time"`
		UsePowerByDiscountProfit             string `json:"use_power_by_discount_profit"`
		UsePowerByDiscountProfitOriginalUnit string `json:"use_power_by_discount_profit_original_unit"`
		UsePowerByDiscountProfitTran         string `json:"use_power_by_discount_profit_tran"`
		UsePowerByDiscountProfitUnit         string `json:"use_power_by_discount_profit_unit"`
		UsePowerProfit                       string `json:"use_power_profit"`
		UsePowerProfitOriginalUnit           string `json:"use_power_profit_original_unit"`
		UsePowerProfitTran                   string `json:"use_power_profit_tran"`
		UsePowerProfitUnit                   string `json:"use_power_profit_unit"`
		UsePowerQuantityTotal                string `json:"use_power_quantity_total"`
		UsePowerQuantityTotalTran            string `json:"use_power_quantity_total_tran"`
		UsePowerQuantityTotalUnit            string `json:"use_power_quantity_total_unit"`
		ValleyCharge                         string `json:"valley_charge"`
		ValleyChargeOriginalUnit             string `json:"valley_charge_original_unit"`
		ValleyChargeTran                     string `json:"valley_charge_tran"`
		ValleyChargeUnit                     string `json:"valley_charge_unit"`
		ValleyNetPowerQuantity               string `json:"valley_net_power_quantity"`
		ValleyNetPowerQuantityTran           string `json:"valley_net_power_quantity_tran"`
		ValleyNetPowerQuantityUnit           string `json:"valley_net_power_quantity_unit"`
		ValleyPowerQuantity                  string `json:"valley_power_quantity"`
		ValleyPowerQuantityTran              string `json:"valley_power_quantity_tran"`
		ValleyPowerQuantityUnit              string `json:"valley_power_quantity_unit"`
		ValleyUsePowerQuantity               string `json:"valley_use_power_quantity"`
		ValleyUsePowerQuantityTran           string `json:"valley_use_power_quantity_tran"`
		ValleyUsePowerQuantityUnit           string `json:"valley_use_power_quantity_unit"`
	} `json:"total"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
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
