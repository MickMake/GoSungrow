package queryDevicePointDayMonthYearDataList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/commonService/queryDevicePointDayMonthYearDataList"
const Disabled = false
const EndPointName = "AppService.queryDevicePointDayMonthYearDataList"

type RequestData struct {
	PsKey     valueTypes.PsKey    `json:"ps_key" required:"true"`
	DataPoint valueTypes.String   `json:"data_point" required:"true"`
	StartTime valueTypes.DateTime `json:"start_time" required:"true"`
	EndTime   valueTypes.DateTime `json:"end_time" required:"true"`
	DataType  valueTypes.String   `json:"data_type" required:"true"`
	QueryType valueTypes.String   `json:"query_type" required:"true"`
}

func (rd *RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(*rd)
}

func (rd RequestData) Help() string {
	ret := api.HelpQueryType()
	// ret += api.HelpDataType()
	return ret
}

type ResultData []struct {
	TimeStamp string `json:"time_stamp"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsKey.String(), GoStruct.NewEndPointPath(e.Request.PsKey.String()))
	return entries
}
