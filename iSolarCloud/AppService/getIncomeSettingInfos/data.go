package getIncomeSettingInfos

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getIncomeSettingInfos"
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
	CodeType                   valueTypes.Integer   `json:"code_type"`
	EnvironmentPowerChargeList []interface{} `json:"enviormentPowerChargeList" PointId:"environmentPowerChargeList"`
	ParamIncomeUnit            valueTypes.Integer   `json:"param_income_unit"`
	PowerElectricalChargeMap   struct {
		CityAllowanceMoney     interface{}  `json:"city_allowance_money" PointValueType:""`
		CodeType               valueTypes.Integer  `json:"code_type"`
		CountyAllowanceMoney   interface{}  `json:"county_allowance_money"`
		DefaultCharge          valueTypes.Float    `json:"default_charge"`
		ElectricChargeID       valueTypes.Integer  `json:"electric_charge_id"`
		EndTime                valueTypes.DateTime `json:"end_time"`
		IncomeStyle            interface{}  `json:"income_style"`
		IntervalTimeCharge     interface{}  `json:"interval_time_charge"`
		NationAllowanceMoney   interface{}  `json:"nation_allowance_money"`
		ParamIncomeUnit        valueTypes.Integer  `json:"param_income_unit"`
		ProvinceAllowanceMoney interface{}  `json:"province_allowance_money"`
		PsID                   valueTypes.Integer  `json:"ps_id"`
		StartTime              valueTypes.DateTime `json:"start_time"`
		UseSharpPeekValleyFlat interface{}  `json:"use_sharp_peek_valley_flat"`
		ValidFlag              valueTypes.Bool     `json:"valid_flag"`
	} `json:"powerElectricalChargeMap"`
	PowerIntervalTimesChargeMap interface{} `json:"powerIntevalTimesChargeMap" PointId:"powerIntervalTimesChargeMap"`
	PowerSelfUseTimesChargeMap  struct {
		DefaultCharge            valueTypes.Float    `json:"default_charge"`
		EndTime                  valueTypes.DateTime `json:"end_time"`
		IntervalTimeCharge       valueTypes.String   `json:"interval_time_charge"`
		OnlineElectricityPercent valueTypes.Float    `json:"online_electricity_percent"`
		PsID                     valueTypes.Integer  `json:"ps_id"`
		StartTime                valueTypes.DateTime `json:"start_time"`
		UseElectricityDiscount   valueTypes.Float    `json:"use_electricity_discount"`
	} `json:"powerSelfUseTimesChargeMap"`
	PsID valueTypes.Integer `json:"ps_id"`
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
		name := api.JoinWithDots(0, "", pkg, e.Response.ResultData.PsID)
		// entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		entries.StructToPoints(e.Response.ResultData, name, e.Response.ResultData.PsID.String(), valueTypes.NewDateTime(""))

		// for _, d := range e.Response.ResultData.EnvironmentPowerChargeList {
		// 	name = fmt.Sprintf("getIncomeSettingInfos.%s", e.Response.ResultData.PsID.String())
		// 	entries.StructToPoints(d, name, e.Request.PsId.String(), valueTypes.NewDateTime(""))
		// }
	}

	return entries
}
