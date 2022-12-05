package getTheoryAndActualPower

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getTheoryAndActualPower"
const Disabled = false
const EndPointName = "MttvScreenService.getTheoryAndActualPower"

type RequestData struct {
	// ws.missing-parameter:user_id or ps_id
	// @TODO - Figure out why duplicate UserId entries in the structure causes json.Marshall to bork in iSolarCloud/api/web.go:128
	PsId   valueTypes.PsId   `json:"ps_id" required:"true"`
	UserId valueTypes.String `json:"user_id"`	// required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
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
