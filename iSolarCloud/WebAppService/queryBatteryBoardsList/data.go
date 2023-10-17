package queryBatteryBoardsList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"

	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/queryBatteryBoardsList"
	Disabled     = false
	EndPointName = "WebAppService.queryBatteryBoardsList"
)

type RequestData struct {
	PsId       valueTypes.PsId    `json:"ps_id" required:"true"`
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
	DeviceSn   valueTypes.String  `json:"device_sn,omitempty"`
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

	for range only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
