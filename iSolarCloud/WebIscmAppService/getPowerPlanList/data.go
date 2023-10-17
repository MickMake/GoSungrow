package getPowerPlanList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/powerStationService/getPowerPlanList"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPowerPlanList"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"CodeId"`

	PsId valueTypes.Float `json:"ps_id"`

	January   valueTypes.Float `json:"one" PointId:"january"`
	February  valueTypes.Float `json:"two" PointId:"february"`
	March     valueTypes.Float `json:"three" PointId:"march"`
	April     valueTypes.Float `json:"four" PointId:"april"`
	May       valueTypes.Float `json:"five" PointId:"may"`
	June      valueTypes.Float `json:"six" PointId:"june"`
	July      valueTypes.Float `json:"seven" PointId:"july"`
	August    valueTypes.Float `json:"eight" PointId:"august"`
	September valueTypes.Float `json:"nine" PointId:"september"`
	October   valueTypes.Float `json:"ten" PointId:"october"`
	November  valueTypes.Float `json:"eleven" PointId:"november"`
	December  valueTypes.Float `json:"twelve" PointId:"december"`

	YearPlanPower valueTypes.Float `json:"year_plan_power"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
