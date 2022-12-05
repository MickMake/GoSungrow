package checkDeviceModel

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/checkDeviceModel"
const Disabled = false
const EndPointName = "WebIscmAppService.checkDeviceModel"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Id                valueTypes.Integer  `json:"id"`
	ComType           valueTypes.String   `json:"com_type"`
	DeviceFactoryId   valueTypes.String   `json:"device_factory_id"`
	DeviceModel       valueTypes.String   `json:"device_model"`
	DeviceModelCode   valueTypes.String   `json:"device_model_code"`
	DeviceSupplier    valueTypes.String   `json:"device_supplier"`
	DeviceType        valueTypes.Integer  `json:"device_type"`
	DeviceTypeName    valueTypes.String   `json:"device_typename" PointId:"device_type_name"`
	InverterModelType valueTypes.Integer  `json:"inverter_model_type"`
	IsDataValid       valueTypes.Bool     `json:"isvalid" PointId:"is_valid"`
	Phone             valueTypes.String   `json:"phone"`
	Protocol          valueTypes.String   `json:"protocol"`
	Remark            valueTypes.String   `json:"remark"`
	SysId             valueTypes.String   `json:"sys_id"`
	SysType           valueTypes.String   `json:"sys_type"`
	UpdateDate        valueTypes.DateTime `json:"updatedate" PointId:"update_date" PointNameDateFormat:"2006/01/02 15:04:05"`
	UpdateUserCode    valueTypes.Bool     `json:"updateusercode" PointId:"update_user_code"`
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
