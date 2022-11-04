package getReportPsTree

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/getReportPsTree"
const Disabled = false

type RequestData struct {
	DeviceType valueTypes.String `json:"device_type" required:"true"`
	PsId       valueTypes.PsId   `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	Dummy valueTypes.String `json:"dummy"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.DateId.Format(valueTypes.DateTimeLayoutDay)
		entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
