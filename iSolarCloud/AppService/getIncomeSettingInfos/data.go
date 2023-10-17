package getIncomeSettingInfos

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/powerStationService/getIncomeSettingInfos"
	Disabled     = false
	EndPointName = "AppService.getIncomeSettingInfos"
)

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
	CodeType                   valueTypes.Integer `json:"code_type"`
	EnvironmentPowerChargeList []interface{}      `json:"enviormentPowerChargeList" PointId:"environment_power_charge_list"`
	ParamIncomeUnit            valueTypes.Integer `json:"param_income_unit"`
	PowerElectricalChargeMap   struct {
		CityAllowanceMoney     interface{}         `json:"city_allowance_money" PointValueType:""`
		CodeType               valueTypes.Integer  `json:"code_type"`
		CountyAllowanceMoney   interface{}         `json:"county_allowance_money"`
		DefaultCharge          valueTypes.Float    `json:"default_charge"`
		ElectricChargeId       valueTypes.Integer  `json:"electric_charge_id"`
		EndTime                valueTypes.DateTime `json:"end_time" PointNameDateFormat:"DateTimeLayout"`
		IncomeStyle            interface{}         `json:"income_style"`
		IntervalTimeCharge     interface{}         `json:"interval_time_charge"`
		NationAllowanceMoney   interface{}         `json:"nation_allowance_money"`
		ParamIncomeUnit        valueTypes.Integer  `json:"param_income_unit"`
		ProvinceAllowanceMoney interface{}         `json:"province_allowance_money"`
		PsId                   valueTypes.PsId     `json:"ps_id"`
		StartTime              valueTypes.DateTime `json:"start_time" PointNameDateFormat:"DateTimeLayout"`
		UseSharpPeekValleyFlat interface{}         `json:"use_sharp_peek_valley_flat"`
		ValidFlag              valueTypes.Bool     `json:"valid_flag"`
	} `json:"powerElectricalChargeMap" PointId:"power_electrical_charge_map"`
	PowerIntervalTimesChargeMap interface{} `json:"powerIntevalTimesChargeMap" PointId:"power_interval_times_charge_map"`
	PowerSelfUseTimesChargeMap  struct {
		DefaultCharge            valueTypes.Float    `json:"default_charge"`
		EndTime                  valueTypes.DateTime `json:"end_time" PointNameDateFormat:"DateTimeLayout"`
		IntervalTimeCharge       valueTypes.String   `json:"interval_time_charge"`
		OnlineElectricityPercent valueTypes.Float    `json:"online_electricity_percent" PointUnit:"%"`
		PsId                     valueTypes.PsId     `json:"ps_id"`
		StartTime                valueTypes.DateTime `json:"start_time" PointNameDateFormat:"DateTimeLayout"`
		UseElectricityDiscount   valueTypes.Float    `json:"use_electricity_discount" PointUnit:"%"`
	} `json:"powerSelfUseTimesChargeMap" PointId:"power_selfuse_times_charge_map"`
	PsId valueTypes.PsId `json:"ps_id"`
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

	for range only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
