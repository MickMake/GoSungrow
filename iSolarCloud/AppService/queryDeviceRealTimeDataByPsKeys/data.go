package queryDeviceRealTimeDataByPsKeys

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/queryDeviceRealTimeDataByPsKeys"
const Disabled = false
const EndPointName = "AppService.queryDeviceRealTimeDataByPsKeys"

type RequestData struct {
	PsKeyList valueTypes.String `json:"ps_key_list" required:"true"`
	PsKeyList2 valueTypes.String `json:"ps_ke_list" required:"true"`
}

func (rd *RequestData) IsValid() error {
	rd.PsKeyList = valueTypes.SetStringValue("1171348_14_1_2")
	rd.PsKeyList2 = valueTypes.SetStringValue("1171348_14_1_2")
	return GoStruct.VerifyOptionsRequired(*rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
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
