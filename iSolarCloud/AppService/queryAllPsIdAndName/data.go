package queryAllPsIdAndName

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/powerStationService/queryAllPsIdAndName"
const Disabled = false
const EndPointName = "AppService.queryAllPsIdAndName"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		GoStruct GoStruct.GoStruct `json:"-" PointIdReplace:"true" PointIdFrom:"PsId"`

		PsId   valueTypes.PsId   `json:"ps_id"`
		PsName valueTypes.String `json:"ps_name"`
	} `json:"pageList" PointId:"page_list" DataTable:"true"`
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
