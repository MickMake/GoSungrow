package getTheoryAndActualPower

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/powerStationService/getTheoryAndActualPower"
	Disabled     = false
	EndPointName = "MttvScreenService.getTheoryAndActualPower"
)

type RequestData struct {
	// ws.missing-parameter:user_id or ps_id
	// @TODO - Figure out why duplicate UserId entries in the structure causes json.Marshall to bork in iSolarCloud/api/web.go:128
	PsId   valueTypes.PsId   `json:"ps_id" required:"true"`
	UserId valueTypes.String `json:"user_id"` // required:"true"`
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

	for range only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
