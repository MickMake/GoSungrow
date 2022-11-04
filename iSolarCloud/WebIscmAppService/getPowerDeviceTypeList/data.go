package getPowerDeviceTypeList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPowerDeviceTypeList"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData []struct {
	IsRemoteUpgrade valueTypes.Integer `json:"is_remote_upgrade"`
	SysID           valueTypes.String  `json:"sys_id"`
	SysName         valueTypes.String  `json:"sys_name"`
	TypeCode        valueTypes.Integer `json:"type_code"`
	TypeID          valueTypes.Integer `json:"type_id"`
	TypeName        valueTypes.String  `json:"type_name"`
	TypeNameEn      valueTypes.String  `json:"type_name_en"`
	UpdateDate      valueTypes.String  `json:"update_date"`
	ValidFlag       valueTypes.Integer `json:"valid_flag"`
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
		entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
