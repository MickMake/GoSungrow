package queryBatteryBoardsList

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryBatteryBoardsList"
const Disabled = false
const EndPointName = "WebAppService.queryBatteryBoardsList"

type RequestData struct {
	PsId       valueTypes.PsId   `json:"ps_id" required:"true"`
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
	DeviceSn   valueTypes.String `json:"device_sn,omitempty"`
	Uuid       valueTypes.Integer `json:"uuid,omitempty"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData []struct {
	// Dummy valueTypes.String `json:"dummy"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
