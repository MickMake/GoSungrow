package getCommunicationDeviceConfigInfo

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const Url = "/v1/devService/getCommunicationDeviceConfigInfo"
const Disabled = false
const EndPointName = "AppService.getCommunicationDeviceConfigInfo"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"TypeId"`
	GoStruct       GoStruct.GoStruct       `json:"-" PointIdFrom:"TypeId" PointIdReplace:"false"`

	TypeId                 valueTypes.Integer `json:"type_id"`
	DeviceName             valueTypes.String  `json:"device_name"`
	CustomTopic            valueTypes.String  `json:"custom_topic"`
	Remark                 valueTypes.String  `json:"remark"`
	IsNeedModbus           valueTypes.Bool    `json:"is_need_modbus"`
	IsSupportParamSet      valueTypes.Bool    `json:"is_support_param_set"`
	IsSupportRemoteUpgrade valueTypes.Bool    `json:"is_support_remote_upgrade"`
	IsSupportTimezoneSet   valueTypes.Bool    `json:"is_support_timezone_set"`
	IsThirdParty           valueTypes.Bool    `json:"is_third_party"`
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
