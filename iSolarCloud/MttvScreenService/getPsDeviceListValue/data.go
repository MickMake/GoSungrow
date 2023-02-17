package getPsDeviceListValue

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/powerStationService/getPsDeviceListValue"
const Disabled = false
const EndPointName = "MttvScreenService.getPsDeviceListValue"

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
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
		GoStruct GoStruct.GoStruct `json:"-" PointIdFromChild:"PsKey" PointIdReplace:"true" PointDeviceFrom:"PsKey"`

		PsKey      valueTypes.String  `json:"pskey" PointId:"ps_key"`
		PsId       valueTypes.Integer `json:"ps_id"`
		DeviceType valueTypes.Integer `json:"device_type"`
		DeviceCode valueTypes.Integer `json:"device_code"`
		ChannelId  valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`

		AttrId                  valueTypes.Integer `json:"attr_id"`
		ComponentAmount         valueTypes.Integer `json:"component_amount"`
		DevFaultStatus          valueTypes.Integer `json:"dev_fault_status"`
		DevStatus               valueTypes.Integer `json:"dev_status"`
		DeviceId                valueTypes.Integer `json:"device_id"`
		DeviceName              valueTypes.String  `json:"device_name"`
		DeviceStatus            valueTypes.String  `json:"device_status"`
		Id                      valueTypes.Integer `json:"id"`
		InstallerDevFaultStatus valueTypes.Integer `json:"installer_dev_fault_status"`
		OrderId                 valueTypes.Integer `json:"orderid" PointId:"order_id"`
		OwnerDevFaultStatus     valueTypes.Integer `json:"owner_dev_fault_status"`
		Posx                    interface{}        `json:"posx"`
		Posy                    interface{}        `json:"posy"`
		StringAmount            valueTypes.Integer `json:"string_amount"`
		UpUUID                  valueTypes.Integer `json:"up_uuid"`
		UUID                    valueTypes.Integer `json:"uuid"`
		UUIDIndexCode           valueTypes.String  `json:"uuid_index_code"`
	} `json:"list" PointId:"device" DataTable:"true" DataTableSortOn:"PsKey"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	return entries
}
