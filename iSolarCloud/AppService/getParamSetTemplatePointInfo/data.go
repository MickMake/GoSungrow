package getParamSetTemplatePointInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getParamSetTemplatePointInfo"
const Disabled = false
const EndPointName = "AppService.getParamSetTemplatePointInfo"

type RequestData struct {
	UuidList          valueTypes.String  `json:"uuid_list" required:"true"`
	SetType           valueTypes.String  `json:"set_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	fmt.Sprintf("")
	rd.UuidList = valueTypes.SetStringValue("1179860")
	rd.SetType = valueTypes.SetStringValue("1")
	return GoStruct.VerifyOptionsRequired(rd)
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
