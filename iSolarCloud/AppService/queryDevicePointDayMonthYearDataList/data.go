package queryDevicePointDayMonthYearDataList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/commonService/queryDevicePointDayMonthYearDataList"
const Disabled = false

type RequestData struct {
	PsKey     valueTypes.PsKey    `json:"ps_key" required:"true"`
	DataPoint valueTypes.String   `json:"data_point" required:"true"`
	StartTime valueTypes.DateTime `json:"start_time" required:"true"`
	EndTime   valueTypes.DateTime `json:"end_time" required:"true"`
	DataType  valueTypes.String   `json:"data_type" required:"true"`
	QueryType valueTypes.String   `json:"query_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("Parameter:query_type must be in (1,2,3)\n")
	ret += fmt.Sprintf("1 - YYYYMMDD\n")
	ret += fmt.Sprintf("2 - YYYYMM\n")
	ret += fmt.Sprintf("3 - YYYY\n")
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

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		name := pkg + "." + e.Request.PsKey.String()
		entries.StructToPoints(e.Response.ResultData, name, e.Request.PsKey.String(), dt)
	}

	return entries
}
