package getLoadCurveList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/reportService/getLoadCurveList"
	Disabled     = false
	EndPointName = "WebAppService.getLoadCurveList"
)

type RequestData struct {
	PsId       valueTypes.PsId    `json:"ps_id" required:"true"`
	MonthDate2 valueTypes.Integer `json:"monthDate" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DayList []interface{}     `json:"dayList"`
	PsKey   valueTypes.String `json:"psKey"`
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
