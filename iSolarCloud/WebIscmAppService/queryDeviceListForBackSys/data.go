package queryDeviceListForBackSys

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/queryDeviceListForBackSys"
const Disabled = false
const EndPointName = "WebIscmAppService.queryDeviceListForBackSys"

type RequestData struct {
	PsId       valueTypes.PsId     `json:"ps_id" require:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []Device

type Device struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"UUID"`

	UUID       valueTypes.Integer `json:"uuid"`

	DeviceCode valueTypes.Integer `json:"device_code"`
	DeviceType valueTypes.Integer `json:"device_type"`
	ChannelId  valueTypes.Integer `json:"chnnl_id" PointId:"channel_id"`
	DeviceId   valueTypes.Integer `json:"device_id"`

	TypeName   valueTypes.String  `json:"type_name"`
	DeviceName valueTypes.String  `json:"device_name"`
	IsPublic   valueTypes.Bool    `json:"is_public"`
	RelState   valueTypes.Integer `json:"rel_state"`
	UpUUID     valueTypes.Integer `json:"up_uuid"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	}

	return entries
}
