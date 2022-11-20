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
	}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"CodeId"`

	Eight         valueTypes.Float `json:"eight"`
	Eleven        valueTypes.Float `json:"eleven"`
	Five          valueTypes.Float `json:"five"`
	Four          valueTypes.Float `json:"four"`
	Nine          valueTypes.Float `json:"nine"`
	One           valueTypes.Float `json:"one"`
	PsID          valueTypes.Float `json:"ps_id"`
	Seven         valueTypes.Float `json:"seven"`
	Six           valueTypes.Float `json:"six"`
	Ten           valueTypes.Float `json:"ten"`
	Three         valueTypes.Float `json:"three"`
	Twelve        valueTypes.Float `json:"twelve"`
	Two           valueTypes.Float `json:"two"`
	YearPlanPower valueTypes.Float `json:"year_plan_power"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
