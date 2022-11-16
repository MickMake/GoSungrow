package getPsDeviceListValue

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPsDeviceListValue"
const Disabled = false

type RequestData struct {
	PsId valueTypes.PsId    `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	List []struct {
		AttrID                  valueTypes.Integer `json:"attr_id"`
		ComponentAmount         valueTypes.Integer `json:"component_amount"`
		DevFaultStatus          valueTypes.Integer `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer `json:"dev_status"`
		DeviceID                valueTypes.Integer `json:"device_id"`
		DeviceName              valueTypes.String  `json:"device_name"`
		DeviceStatus            valueTypes.String  `json:"device_status"`
		DeviceType              valueTypes.Integer `json:"device_type"`
		Id                      valueTypes.Integer `json:"id"`
		InstallerDevFaultStatus valueTypes.Integer `json:"installer_dev_fault_status"`
		OrderId                 valueTypes.Integer `json:"orderid" PointId:"order_id"`
		OwnerDevFaultStatus     valueTypes.Integer `json:"owner_dev_fault_status"`
		Posx                    interface{}        `json:"posx"`
		Posy                    interface{}        `json:"posy"`
		PsId                    valueTypes.Integer `json:"ps_id"`
		PsKey                   valueTypes.String  `json:"pskey" PointId:"ps_key"`
		StringAmount            valueTypes.Integer `json:"string_amount"`
		UpUUID                  valueTypes.Integer `json:"up_uuid"`
		UUID                    valueTypes.Integer `json:"uuid"`
		UUIDIndexCode           valueTypes.String  `json:"uuid_index_code"`
	} `json:"list" PointIdFromChild:"PsId.PsKey" PointIdReplace:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
