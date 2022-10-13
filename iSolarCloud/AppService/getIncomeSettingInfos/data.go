package getIncomeSettingInfos

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getIncomeSettingInfos"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CodeType                   api.Integer   `json:"code_type"`
	EnvironmentPowerChargeList []interface{} `json:"enviormentPowerChargeList"`
	ParamIncomeUnit            api.Integer   `json:"param_income_unit"`
	PowerElectricalChargeMap   struct {
		CityAllowanceMoney     interface{}  `json:"city_allowance_money"`
		CodeType               api.Integer  `json:"code_type"`
		CountyAllowanceMoney   interface{}  `json:"county_allowance_money"`
		DefaultCharge          api.Float    `json:"default_charge"`
		ElectricChargeID       api.Integer  `json:"electric_charge_id"`
		EndTime                api.DateTime `json:"end_time"`
		IncomeStyle            interface{}  `json:"income_style"`
		IntervalTimeCharge     interface{}  `json:"interval_time_charge"`
		NationAllowanceMoney   interface{}  `json:"nation_allowance_money"`
		ParamIncomeUnit        api.Integer  `json:"param_income_unit"`
		ProvinceAllowanceMoney interface{}  `json:"province_allowance_money"`
		PsID                   api.Integer  `json:"ps_id"`
		StartTime              api.DateTime `json:"start_time"`
		UseSharpPeekValleyFlat interface{}  `json:"use_sharp_peek_valley_flat"`
		ValidFlag              api.Bool     `json:"valid_flag"`
	} `json:"powerElectricalChargeMap"`
	PowerIntervalTimesChargeMap interface{} `json:"powerIntevalTimesChargeMap"`
	PowerSelfUseTimesChargeMap  struct {
		DefaultCharge            api.Float    `json:"default_charge"`
		EndTime                  api.DateTime `json:"end_time"`
		IntervalTimeCharge       api.String   `json:"interval_time_charge"`
		OnlineElectricityPercent api.Float    `json:"online_electricity_percent"`
		PsID                     api.Integer  `json:"ps_id"`
		StartTime                api.DateTime `json:"start_time"`
		UseElectricityDiscount   api.Float    `json:"use_electricity_discount"`
	} `json:"powerSelfUseTimesChargeMap"`
	PsID api.Integer `json:"ps_id"`
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
		pkg := apiReflect.GetName("", *e)
		name := fmt.Sprintf("%s.%s", pkg, e.Response.ResultData.PsID.String())
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsId.String(), api.NewDateTime(""))

		// for _, d := range e.Response.ResultData.EnvironmentPowerChargeList {
		// 	name = fmt.Sprintf("getIncomeSettingInfos.%s", e.Response.ResultData.PsID.String())
		// 	entries.StructToPoints(d, name, e.Request.PsId.String(), api.NewDateTime(""))
		// }
	}

	return entries
}
