package reportList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/reportList"
const Disabled = false

type RequestData struct {
	PsId       valueTypes.Integer `json:"ps_id"`
	ReportType string `json:"report_type"`
	// DateID   string `json:"date_id,omitempty"`
	// DateType string `json:"date_type,omitempty"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("report_type: ")
	return ret
}

type ResultData struct {
	DataList []DataList `json:"dataList"`
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
	Total     []DataList  `json:"total"`
}
type DataList struct {
	DateId     valueTypes.DateTime `json:"date_id"`
	DeviceName interface{}  `json:"device_name"`
	PsId       valueTypes.Integer  `json:"ps_id"`
	TimeStamp  interface{}  `json:"time_stamp"` // Sad that this alternates between string and valueTypes.Integer.
	UpdateTime valueTypes.DateTime `json:"update_time"`
	UUID       interface{}  `json:"uuid"`
	Co2Reduce  valueTypes.Float    `json:"co2_reduce"`

	CitySubsidyCharge             valueTypes.Float `json:"city_subsidy_charge" PointUnitFrom:"city_subsidy_charge_original_unit"`
	CitySubsidyChargeOriginalUnit string    `json:"city_subsidy_charge_original_unit"`
	CitySubsidyChargeTran         valueTypes.Float `json:"city_subsidy_charge_tran" PointUnitFrom:"city_subsidy_charge_unit" PointIgnore:"true"`
	CitySubsidyChargeUnit         string    `json:"city_subsidy_charge_unit" PointIgnore:"true"`

	CountrySubsidyCharge             valueTypes.Float `json:"country_subsidy_charge" PointUnitFrom:"country_subsidy_charge_original_unit"`
	CountrySubsidyChargeOriginalUnit string    `json:"country_subsidy_charge_original_unit"`
	CountrySubsidyChargeTran         valueTypes.Float `json:"country_subsidy_charge_tran" PointUnitFrom:"country_subsidy_charge_unit" PointIgnore:"true"`
	CountrySubsidyChargeUnit         string    `json:"country_subsidy_charge_unit" PointIgnore:"true"`

	CountySubsidyCharge             valueTypes.Float `json:"county_subsidy_charge" PointUnitFrom:"county_subsidy_charge_original_unit"`
	CountySubsidyChargeOriginalUnit string    `json:"county_subsidy_charge_original_unit"`
	CountySubsidyChargeTran         valueTypes.Float `json:"county_subsidy_charge_tran" PointUnitFrom:"county_subsidy_charge_unit" PointIgnore:"true"`
	CountySubsidyChargeUnit         string    `json:"county_subsidy_charge_unit" PointIgnore:"true"`

	CuspCharge             valueTypes.Float `json:"cusp_charge" PointUnitFrom:"cusp_charge_original_unit"`
	CuspChargeOriginalUnit string    `json:"cusp_charge_original_unit"`
	CuspChargeTran         valueTypes.Float `json:"cusp_charge_tran" PointUnitFrom:"cusp_charge_unit" PointIgnore:"true"`
	CuspChargeUnit         string    `json:"cusp_charge_unit" PointIgnore:"true"`

	CuspNetPowerQuantity     valueTypes.Float `json:"cusp_net_power_quantity" PointUnitFrom:"cusp_net_power_quantity_unit"`
	CuspNetPowerQuantityTran valueTypes.Float `json:"cusp_net_power_quantity_tran" PointUnitFrom:"cusp_net_power_quantity_unit" PointIgnore:"true"`
	CuspNetPowerQuantityUnit string    `json:"cusp_net_power_quantity_unit"`

	CuspPowerQuantity     valueTypes.Float `json:"cusp_power_quantity" PointUnitFrom:"cusp_power_quantity_unit"`
	CuspPowerQuantityTran valueTypes.Float `json:"cusp_power_quantity_tran" PointUnitFrom:"cusp_power_quantity_unit" PointIgnore:"true"`
	CuspPowerQuantityUnit string    `json:"cusp_power_quantity_unit"`

	CuspUsePowerQuantity     valueTypes.Float `json:"cusp_use_power_quantity" PointUnitFrom:"cusp_use_power_quantity_unit"`
	CuspUsePowerQuantityTran valueTypes.Float `json:"cusp_use_power_quantity_tran" PointUnitFrom:"cusp_use_power_quantity_unit" PointIgnore:"true"`
	CuspUsePowerQuantityUnit string    `json:"cusp_use_power_quantity_unit"`

	FlatCharge             valueTypes.Float `json:"flat_charge" PointUnitFrom:"flat_charge_original_unit"`
	FlatChargeOriginalUnit string    `json:"flat_charge_original_unit"`

	FlatChargeTran valueTypes.Float `json:"flat_charge_tran" PointUnitFrom:"flat_charge_unit"`
	FlatChargeUnit string    `json:"flat_charge_unit"`

	FlatNetPowerQuantity     valueTypes.Float `json:"flat_net_power_quantity" PointUnitFrom:"flat_net_power_quantity_unit"`
	FlatNetPowerQuantityTran valueTypes.Float `json:"flat_net_power_quantity_tran" PointUnitFrom:"flat_net_power_quantity_unit" PointIgnore:"true"`
	FlatNetPowerQuantityUnit string    `json:"flat_net_power_quantity_unit"`

	FlatPowerQuantity     valueTypes.Float `json:"flat_power_quantity" PointUnitFrom:"flat_power_quantity_unit"`
	FlatPowerQuantityTran valueTypes.Float `json:"flat_power_quantity_tran" PointUnitFrom:"flat_power_quantity_unit" PointIgnore:"true"`
	FlatPowerQuantityUnit string    `json:"flat_power_quantity_unit"`

	FlatUsePowerQuantity     valueTypes.Float `json:"flat_use_power_quantity" PointUnitFrom:"flat_use_power_quantity_unit"`
	FlatUsePowerQuantityTran valueTypes.Float `json:"flat_use_power_quantity_tran" PointUnitFrom:"flat_use_power_quantity_unit" PointIgnore:"true"`
	FlatUsePowerQuantityUnit string    `json:"flat_use_power_quantity_unit"`

	NetPowerProfit             valueTypes.Float `json:"net_power_profit" PointUnitFrom:"net_power_profit_original_unit"`
	NetPowerProfitOriginalUnit string    `json:"net_power_profit_original_unit"`
	NetPowerProfitTran         valueTypes.Float `json:"net_power_profit_tran" PointUnitFrom:"net_power_profit_unit" PointIgnore:"true"`
	NetPowerProfitUnit         string    `json:"net_power_profit_unit" PointIgnore:"true"`

	NetPowerQuantityTotal     valueTypes.Float `json:"net_power_quantity_total" PointUnitFrom:"net_power_quantity_total_unit"`
	NetPowerQuantityTotalTran valueTypes.Float `json:"net_power_quantity_total_tran" PointUnitFrom:"net_power_quantity_total_unit" PointIgnore:"true"`
	NetPowerQuantityTotalUnit string    `json:"net_power_quantity_total_unit"`

	PeakCharge             valueTypes.Float `json:"peak_charge" PointUnitFrom:"peak_charge_original_unit"`
	PeakChargeOriginalUnit string    `json:"peak_charge_original_unit"`

	PeakChargeTran valueTypes.Float `json:"peak_charge_tran" PointUnitFrom:"peak_charge_unit"`
	PeakChargeUnit string    `json:"peak_charge_unit"`

	PeakNetPowerQuantity     valueTypes.Float `json:"peak_net_power_quantity" PointUnitFrom:"peak_net_power_quantity_unit"`
	PeakNetPowerQuantityTran valueTypes.Float `json:"peak_net_power_quantity_tran" PointUnitFrom:"peak_net_power_quantity_unit" PointIgnore:"true"`
	PeakNetPowerQuantityUnit string    `json:"peak_net_power_quantity_unit"`

	PeakPowerQuantity     valueTypes.Float `json:"peak_power_quantity" PointUnitFrom:"peak_power_quantity_unit"`
	PeakPowerQuantityTran valueTypes.Float `json:"peak_power_quantity_tran" PointUnitFrom:"peak_power_quantity_unit" PointIgnore:"true"`
	PeakPowerQuantityUnit string    `json:"peak_power_quantity_unit"`

	PeakUsePowerQuantity     valueTypes.Float `json:"peak_use_power_quantity" PointUnitFrom:"peak_use_power_quantity_unit"`
	PeakUsePowerQuantityTran valueTypes.Float `json:"peak_use_power_quantity_tran" PointUnitFrom:"peak_use_power_quantity_unit" PointIgnore:"true"`
	PeakUsePowerQuantityUnit string    `json:"peak_use_power_quantity_unit"`

	PowerQuantityTotal     valueTypes.Float `json:"power_quantity_total" PointUnitFrom:"power_quantity_total_unit"`
	PowerQuantityTotalTran valueTypes.Float `json:"power_quantity_total_tran" PointUnitFrom:"power_quantity_total_unit" PointIgnore:"true"`
	PowerQuantityTotalUnit string    `json:"power_quantity_total_unit"`

	ProvinceSubsidyCharge             valueTypes.Float `json:"province_subsidy_charge" PointUnitFrom:"province_subsidy_charge_original_unit"`
	ProvinceSubsidyChargeOriginalUnit string    `json:"province_subsidy_charge_original_unit"`

	ProvinceSubsidyChargeTran valueTypes.Float `json:"province_subsidy_charge_tran" PointUnitFrom:"province_subsidy_charge_unit"`
	ProvinceSubsidyChargeUnit string    `json:"province_subsidy_charge_unit"`

	SubsidyProfit             valueTypes.Float `json:"subsidy_profit" PointUnitFrom:"subsidy_profit_original_unit"`
	SubsidyProfitOriginalUnit string    `json:"subsidy_profit_original_unit"`

	SubsidyProfitTran valueTypes.Float `json:"subsidy_profit_tran" PointUnitFrom:"subsidy_profit_unit"`
	SubsidyProfitUnit string    `json:"subsidy_profit_unit"`

	TotalProfit             valueTypes.Float `json:"total_profit" PointUnitFrom:"total_profit_original_unit"`
	TotalProfitOriginalUnit string    `json:"total_profit_original_unit"`
	TotalProfitTran         valueTypes.Float `json:"total_profit_tran" PointUnitFrom:"total_profit_unit" PointIgnore:"true"`
	TotalProfitUnit         string    `json:"total_profit_unit" PointIgnore:"true"`

	UsePowerByDiscountProfit             valueTypes.Float `json:"use_power_by_discount_profit" PointUnitFrom:"use_power_by_discount_profit_original_unit"`
	UsePowerByDiscountProfitOriginalUnit string    `json:"use_power_by_discount_profit_original_unit"`
	UsePowerByDiscountProfitTran         valueTypes.Float `json:"use_power_by_discount_profit_tran" PointUnitFrom:"use_power_by_discount_profit_unit" PointIgnore:"true"`
	UsePowerByDiscountProfitUnit         string    `json:"use_power_by_discount_profit_unit" PointIgnore:"true"`

	UsePowerProfit             valueTypes.Float `json:"use_power_profit" PointUnitFrom:"use_power_profit_original_unit"`
	UsePowerProfitOriginalUnit string    `json:"use_power_profit_original_unit"`
	UsePowerProfitTran         valueTypes.Float `json:"use_power_profit_tran" PointUnitFrom:"use_power_profit_unit" PointIgnore:"true"`
	UsePowerProfitUnit         string    `json:"use_power_profit_unit" PointIgnore:"true"`

	UsePowerQuantityTotal     valueTypes.Float `json:"use_power_quantity_total" PointUnitFrom:"use_power_quantity_total_unit"`
	UsePowerQuantityTotalTran valueTypes.Float `json:"use_power_quantity_total_tran" PointUnitFrom:"use_power_quantity_total_unit" PointIgnore:"true"`
	UsePowerQuantityTotalUnit string    `json:"use_power_quantity_total_unit"`

	ValleyCharge             valueTypes.Float `json:"valley_charge" PointUnitFrom:"valley_charge_original_unit"`
	ValleyChargeOriginalUnit string    `json:"valley_charge_original_unit"`
	ValleyChargeTran         valueTypes.Float `json:"valley_charge_tran" PointUnitFrom:"valley_charge_unit" PointIgnore:"true"`
	ValleyChargeUnit         string    `json:"valley_charge_unit" PointIgnore:"true"`

	ValleyNetPowerQuantity     valueTypes.Float `json:"valley_net_power_quantity" PointUnitFrom:"valley_net_power_quantity_unit"`
	ValleyNetPowerQuantityTran valueTypes.Float `json:"valley_net_power_quantity_tran" PointUnitFrom:"valley_net_power_quantity_unit" PointIgnore:"true"`
	ValleyNetPowerQuantityUnit string    `json:"valley_net_power_quantity_unit"`

	ValleyPowerQuantity     valueTypes.Float `json:"valley_power_quantity" PointUnitFrom:"valley_power_quantity_unit"`
	ValleyPowerQuantityTran valueTypes.Float `json:"valley_power_quantity_tran" PointUnitFrom:"valley_power_quantity_unit" PointIgnore:"true"`
	ValleyPowerQuantityUnit string    `json:"valley_power_quantity_unit"`

	ValleyUsePowerQuantity     valueTypes.Float `json:"valley_use_power_quantity" PointUnitFrom:"valley_use_power_quantity_unit"`
	ValleyUsePowerQuantityTran valueTypes.Float `json:"valley_use_power_quantity_tran" PointUnitFrom:"valley_use_power_quantity_unit" PointIgnore:"true"`
	ValleyUsePowerQuantityUnit string    `json:"valley_use_power_quantity_unit"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		entries.StructToPoints(e.Response.ResultData, pkg, e.Request.PsId.String(), valueTypes.NewDateTime(""))

		s := valueTypes.SizeOfArrayLength(e.Response.ResultData.DataList)
		for _, v := range e.Response.ResultData.DataList {
			entries.StructToPoints(v, api.JoinWithDots(s, valueTypes.DateTimeLayoutDay, pkg, "DataList", v.PsId, v.DateId), v.PsId.String(), valueTypes.NewDateTime(""))
		}

		s = valueTypes.SizeOfArrayLength(e.Response.ResultData.Info)
		for i, v := range e.Response.ResultData.Info {
			entries.StructToPoints(v, api.JoinWithDots(s, valueTypes.DateTimeLayoutDay, pkg, "Info", v.PsId, i), v.PsId.String(), valueTypes.NewDateTime(""))
		}

		s = valueTypes.SizeOfArrayLength(e.Response.ResultData.Total)
		for i, v := range e.Response.ResultData.Total {
			entries.StructToPoints(v, api.JoinWithDots(s, valueTypes.DateTimeLayoutDay, pkg, "Total", v.PsId, i), v.PsId.String(), valueTypes.NewDateTime(""))
		}
	}

	return entries
}
