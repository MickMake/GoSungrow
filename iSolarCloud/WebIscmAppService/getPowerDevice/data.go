package getPowerDevice

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPowerDevice"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ChannelID        valueTypes.Integer `json:"channelId"`
	ClassCode        valueTypes.Integer `json:"classCode"`
	DeviceCode       valueTypes.Integer `json:"deviceCode"`
	DeviceID         valueTypes.Integer `json:"deviceId"`
	DeviceName       valueTypes.String  `json:"deviceName"`
	DeviceStatus     interface{}        `json:"deviceStatus"`
	IsDisplay        valueTypes.Integer `json:"isDisplay"`
	IsUse            valueTypes.Integer `json:"isUse"`
	ParentDeviceCode interface{}        `json:"parentDeviceCode"`
	ParentGateWay    interface{}        `json:"parentGateWay"`
	Producer         interface{}        `json:"producer"`
	PsGUID           valueTypes.String  `json:"psGuid"`
	SnCode           interface{}        `json:"snCode"`
	SyncDate         interface{}        `json:"syncDate"`
	UpdateDate       valueTypes.String  `json:"updateDate"`
	UpdateUserCode   interface{}        `json:"updateUserCode"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
