package getOwnerFaultConfigList

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/faultService/getOwnerFaultConfigList"
	Disabled     = false
	EndPointName = "AppService.getOwnerFaultConfigList"
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
	PageList []struct {
		GoStruct GoStruct.GoStruct `json:"-" PointIdFrom:"FaultTypeId" PointIdReplace:"true"`

		FaultTypeId       valueTypes.Integer `json:"fault_type_id"`
		FaultTypeCode     valueTypes.Integer `json:"fault_type_code"`
		FaultTypeCodeName valueTypes.String  `json:"fault_type_code_name"`
		FaultLevel        valueTypes.Integer `json:"fault_level"`
		FaultType         valueTypes.Integer `json:"fault_type"`

		DeviceType       valueTypes.Integer `json:"device_type"`
		DeviceTypeName   valueTypes.String  `json:"device_type_name"`
		DevFaultTypeCode valueTypes.String  `json:"dev_fault_type_code"`
		NewFaultTypeCode valueTypes.Integer `json:"new_fault_type_code"`

		IsAllowOwnerView     valueTypes.Bool `json:"is_allow_owner_view"`
		IsAllowUserAdd       valueTypes.Bool `json:"is_allow_user_add"`
		IsSupportFaultRecord valueTypes.Bool `json:"is_support_fault_record"`
	} `json:"pageList" PointId:"faults" DataTable:"true" DataTableSortOn:"FaultTypeId"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
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
