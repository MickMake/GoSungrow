package queryDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryDeviceListForBackSys"
const Disabled = false

type RequestData struct {
	PsID       valueTypes.PsId     `json:"ps_id" require:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	ChnnlID    valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
	DeviceCode valueTypes.Integer `json:"device_code"`
	DeviceID   valueTypes.Integer `json:"device_id"`
	DeviceName valueTypes.String  `json:"device_name"`
	DeviceType valueTypes.Integer `json:"device_type"`
	IsPublic   valueTypes.Integer `json:"is_public"`
	RelState   valueTypes.Integer `json:"rel_state"`
	TypeName   valueTypes.String  `json:"type_name"`
	UpUUID     valueTypes.Integer `json:"up_uuid"`
	UUID       valueTypes.Integer `json:"uuid"`
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
