package queryFaultTypeByDevicePage

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/faultService/queryFaultTypeByDevicePage"
const Disabled = false
const EndPointName = "AppService.queryFaultTypeByDevicePage"

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
	PageList []struct {
		DevFaultTypeCode valueTypes.String  `json:"dev_fault_type_code"`
		FaultTypeCode    valueTypes.Integer `json:"fault_type_code"`
		FaultTypeCode2   valueTypes.Integer `json:"faulttypecode" PointId:"fault_type_code2"`
		FaultValue       valueTypes.String  `json:"faultvalue" PointId:"fault_value"`
		IsAllowOwnerView valueTypes.Bool    `json:"is_allow_owner_view"`
	} `json:"pageList" PointId:"page_list" DataTable:"true"`
	CurPage    valueTypes.Integer `json:"curPage" PointId:"cur_page"`
	IsMore     valueTypes.Bool    `json:"isMore" PointId:"is_more"`
	RowCount   valueTypes.Integer `json:"rowCount" PointId:"row_count"`
	Size       valueTypes.Integer `json:"size"`
	StartIndex valueTypes.Integer `json:"startIndex" PointId:"start_index"`
	TotalPage  valueTypes.Integer `json:"totalPage" PointId:"total_page"`
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
