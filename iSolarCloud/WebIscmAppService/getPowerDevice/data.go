package getPowerDevice

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/devService/getPowerDevice"
	Disabled     = false
	EndPointName = "WebIscmAppService.getPowerDevice"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ChannelId        valueTypes.Integer  `json:"channelId" PointId:"channel_id"`
	ClassCode        valueTypes.Integer  `json:"classCode" PointId:"class_code"`
	DeviceCode       valueTypes.Integer  `json:"deviceCode" PointId:"device_code"`
	DeviceId         valueTypes.Integer  `json:"deviceId" PointId:"device_id"`
	DeviceName       valueTypes.String   `json:"deviceName" PointId:"device_name"`
	DeviceStatus     interface{}         `json:"deviceStatus" PointId:"device_status"`
	IsDisplay        valueTypes.Bool     `json:"isDisplay" PointId:"is_display"`
	IsUse            valueTypes.Bool     `json:"isUse" PointId:"is_use"`
	ParentDeviceCode interface{}         `json:"parentDeviceCode" PointId:"parent_device_code"`
	ParentGateWay    interface{}         `json:"parentGateWay" PointId:"parent_gate_way"`
	Producer         interface{}         `json:"producer"`
	PsGUID           valueTypes.String   `json:"psGuid" PointId:"ps_guid"`
	SnCode           interface{}         `json:"snCode" PointId:"sn_code"`
	SyncDate         interface{}         `json:"syncDate" PointId:"sync_date"`
	UpdateDate       valueTypes.DateTime `json:"updateDate" PointId:"update_date" PointNameDateFormat:"DateTimeLayout"`
	UpdateUserCode   interface{}         `json:"updateUserCode" PointId:"update_user_code"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
