package getPsListByName

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/powerStationService/getPsListByName"
const Disabled = false
const EndPointName = "AppService.getPsListByName"

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
	PsId         valueTypes.PsId    `json:"ps_id"`
	PsName       valueTypes.String  `json:"ps_name"`
	PsShortName  valueTypes.String  `json:"ps_short_name"`
	PsTimezone   valueTypes.String  `json:"ps_timezone"`
	PsTimezoneId valueTypes.Integer `json:"ps_timezone_id"`
	ShareType    valueTypes.Integer `json:"share_type"`
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
