package queryDevicePointsDayMonthYearDataList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/commonService/queryDevicePointsDayMonthYearDataList"
	Disabled     = false
	EndPointName = "AppService.queryDevicePointsDayMonthYearDataList"
)

type RequestData struct {
	PsKey     valueTypes.PsKey    `json:"ps_key" required:"true"`
	DataPoint valueTypes.String   `json:"data_point" required:"true"`
	StartTime valueTypes.DateTime `json:"start_time" required:"true"`
	EndTime   valueTypes.DateTime `json:"end_time" required:"true"`
	DataType  valueTypes.String   `json:"data_type" required:"true"`
	QueryType valueTypes.String   `json:"query_type" required:"true"`
	Points    valueTypes.String   `json:"points" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	One []struct {
		TimeStamp valueTypes.DateTime `json:"time_stamp"`
	} `json:"1"`
	Two []struct {
		TimeStamp valueTypes.DateTime `json:"time_stamp"`
	} `json:"2"`
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
