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
		CuspNetPowerQuantity     interface{}         `json:"cusp_net_power_quantity"`
		CuspPowerQuantity        interface{}         `json:"cusp_power_quantity"`
		CuspUsePowerQuantity     interface{}         `json:"cusp_use_power_quantity"`
		DateId                   valueTypes.DateTime `json:"date_id" PointNameDateFormat:"2006/01/02 15:04:05"`
		FlatNetPowerQuantity     valueTypes.Float    `json:"flat_net_power_quantity"`
		FlatPowerQuantity        valueTypes.Float    `json:"flat_power_quantity"`
		FlatUsePowerQuantity     valueTypes.Float    `json:"flat_use_power_quantity"`
		NetPowerProfit           valueTypes.Float    `json:"net_power_profit"`
		NetPowerQuantityTotal    valueTypes.Float    `json:"net_power_quantity_total"`
		PeakNetPowerQuantity     interface{}         `json:"peak_net_power_quantity"`
		PeakPowerQuantity        interface{}         `json:"peak_power_quantity"`
		PeakUsePowerQuantity     interface{}         `json:"peak_use_power_quantity"`
		PowerQuantityTotal       valueTypes.Float    `json:"power_quantity_total"`
		SubsidyProfit            interface{}         `json:"subsidy_profit"`
		TotalProfit              valueTypes.Float    `json:"total_profit"`
		UpdateTime               valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006/01/02 15:04:05"`
		UsePowerByDiscountProfit valueTypes.Float    `json:"use_power_by_discount_profit"`
		UsePowerProfit           valueTypes.Float    `json:"use_power_profit"`
		UsePowerQuantityTotal    valueTypes.Float    `json:"use_power_quantity_total"`
		ValleyNetPowerQuantity   interface{}         `json:"valley_net_power_quantity"`
		ValleyPowerQuantity      interface{}         `json:"valley_power_quantity"`
		ValleyUsePowerQuantity   interface{}         `json:"valley_use_power_quantity"`
	} `json:"actual_list" PointId:"actual" PointIdFromChild:"DateId" PointNameDateFormat:"20060102" PointArrayFlatten:"false"`	// PointIgnoreIfChildFromNil:"UpdateTime" DataTable:"true"`
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
