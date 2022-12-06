package queryPsProfit

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/queryPsProfit"
const Disabled = false
const EndPointName = "AppService.queryPsProfit"

type RequestData struct {
	PsId     valueTypes.PsId `json:"ps_id" required:"true"`
	DateId   valueTypes.DateTime `json:"date_id" required:"true"`
	DateType valueTypes.String `json:"date_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += api.HelpDateId()
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	ActualList []struct {
		GoStruct                 GoStruct.GoStruct   `json:"-" PointIdReplace:"true" PointIdFrom:"DateId" PointNameDateFormat:"20060102" PointTimestampFrom:"UpdateTime"`

		DateId                   valueTypes.DateTime `json:"date_id" PointNameDateFormat:"2006-01-02" PointTimestampFrom:"UpdateTime"`
		UpdateTime               valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006-01-02 15:04:05"`

		NetPowerProfit           valueTypes.Float    `json:"net_power_profit" PointTimestampFrom:"UpdateTime"`
		SubsidyProfit            interface{}         `json:"subsidy_profit" PointTimestampFrom:"UpdateTime"`
		UsePowerProfit           valueTypes.Float    `json:"use_power_profit" PointTimestampFrom:"UpdateTime"`
		UsePowerByDiscountProfit valueTypes.Float    `json:"use_power_by_discount_profit" PointTimestampFrom:"UpdateTime"`
		TotalProfit              valueTypes.Float    `json:"total_profit" PointTimestampFrom:"UpdateTime"`

		NetPowerQuantityTotal    valueTypes.Float    `json:"net_power_quantity_total" PointTimestampFrom:"UpdateTime"`
		PowerQuantityTotal       valueTypes.Float    `json:"power_quantity_total" PointTimestampFrom:"UpdateTime"`
		UsePowerQuantityTotal    valueTypes.Float    `json:"use_power_quantity_total" PointTimestampFrom:"UpdateTime"`

		CuspNetPowerQuantity     interface{}         `json:"cusp_net_power_quantity" PointTimestampFrom:"UpdateTime"`
		CuspPowerQuantity        interface{}         `json:"cusp_power_quantity" PointTimestampFrom:"UpdateTime"`
		CuspUsePowerQuantity     interface{}         `json:"cusp_use_power_quantity" PointTimestampFrom:"UpdateTime"`
		FlatNetPowerQuantity     valueTypes.Float    `json:"flat_net_power_quantity" PointTimestampFrom:"UpdateTime"`
		FlatPowerQuantity        valueTypes.Float    `json:"flat_power_quantity" PointTimestampFrom:"UpdateTime"`
		FlatUsePowerQuantity     valueTypes.Float    `json:"flat_use_power_quantity" PointTimestampFrom:"UpdateTime"`
		PeakNetPowerQuantity     interface{}         `json:"peak_net_power_quantity" PointTimestampFrom:"UpdateTime"`
		PeakPowerQuantity        interface{}         `json:"peak_power_quantity" PointTimestampFrom:"UpdateTime"`
		PeakUsePowerQuantity     interface{}         `json:"peak_use_power_quantity" PointTimestampFrom:"UpdateTime"`
		ValleyNetPowerQuantity   interface{}         `json:"valley_net_power_quantity" PointTimestampFrom:"UpdateTime"`
		ValleyPowerQuantity      interface{}         `json:"valley_power_quantity" PointTimestampFrom:"UpdateTime"`
		ValleyUsePowerQuantity   interface{}         `json:"valley_use_power_quantity" PointTimestampFrom:"UpdateTime"`
	} `json:"actual_list" PointId:"actual" DataTable:"true"`
	// Need to fix this output - PointIdFromChild:"DateId" isn't working.
	PlanList []interface{} `json:"plan_list" PointArrayFlatten:"false"`
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
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
