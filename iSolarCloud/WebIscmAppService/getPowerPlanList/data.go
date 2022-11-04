package getPowerPlanList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerPlanList"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	Eight         valueTypes.Integer `json:"eight"`
	Eleven        valueTypes.Integer `json:"eleven"`
	Five          valueTypes.Integer `json:"five"`
	Four          valueTypes.Integer `json:"four"`
	Nine          valueTypes.Integer `json:"nine"`
	One           valueTypes.Integer `json:"one"`
	PsID          valueTypes.Integer `json:"ps_id"`
	Seven         valueTypes.Integer `json:"seven"`
	Six           valueTypes.Integer `json:"six"`
	Ten           valueTypes.Integer `json:"ten"`
	Three         valueTypes.Integer `json:"three"`
	Twelve        valueTypes.Integer `json:"twelve"`
	Two           valueTypes.Integer `json:"two"`
	YearPlanPower valueTypes.Integer `json:"year_plan_power"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
