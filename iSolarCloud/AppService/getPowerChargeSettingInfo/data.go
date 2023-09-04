package getPowerChargeSettingInfo

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerChargeSettingInfo"
const Disabled = false
const EndPointName = "AppService.getPowerChargeSettingInfo"

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
	ParamIncomeUnit          valueTypes.Integer `json:"param_income_unit"`
	ParamIncomeUnitName      valueTypes.String  `json:"param_income_unit_name"`
	PowerElectricalChargeMap struct {
		DefaultCharge      valueTypes.Float `json:"default_charge" PointUnitFromParent:"ParamIncomeUnitName"`
		IntervalTimeCharge interface{}      `json:"interval_time_charge"`
	} `json:"powerElectricalChargeMap" PointId:"power_electrical_charge_map"`
	PowerSelfUseTimesChargeMap struct {
		DefaultCharge      valueTypes.Float  `json:"default_charge" PointUnitFromParent:"ParamIncomeUnitName"`
		IntervalTimeCharge valueTypes.String `json:"interval_time_charge"`
	} `json:"powerSelfUseTimesChargeMap" PointId:"power_self_use_charge_map"`
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

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
