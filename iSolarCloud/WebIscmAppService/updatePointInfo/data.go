package updatePointInfo

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/updatePointInfo"
	Disabled     = false
	EndPointName = "WebIscmAppService.updatePointInfo"
)

type RequestData struct {
	PointName    valueTypes.String  `json:"point_name"`
	Id           valueTypes.String  `json:"id" required:"true"`
	DeviceType   valueTypes.Integer `json:"device_type" required:"true"`
	PointId      valueTypes.String  `json:"point_id" required:"true"`
	PointType    valueTypes.String  `json:"point_type"`
	PointGroupId valueTypes.String  `json:"point_group_id"`
}

func (rd *RequestData) IsValid() error {
	rd.PointName = valueTypes.SetStringValue("")
	rd.PointType = valueTypes.SetStringValue("")
	rd.PointGroupId = valueTypes.SetStringValue("")
	return GoStruct.VerifyOptionsRequired(*rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("Wireless Signal Strength")
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
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
