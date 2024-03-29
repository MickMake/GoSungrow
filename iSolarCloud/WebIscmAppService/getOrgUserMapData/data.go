package getOrgUserMapData

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/faultService/getOrgUserMapData"
const Disabled = false
const EndPointName = "WebIscmAppService.getOrgUserMapData"

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
	CurPage    valueTypes.Integer `json:"curPage" PointId:"cur_page"`
	IsMore     interface{}        `json:"isMore" PointId:"is_more"`
	PageList   []interface{}      `json:"pageList" PointId:"page_list"`
	RowCount   valueTypes.Integer `json:"rowCount" PointId:"row_count"`
	Size       valueTypes.Integer `json:"size"`
	StartIndex interface{}        `json:"startIndex" PointId:"start_index"`
	TotalPage  interface{}        `json:"totalPage" PointId:"total_page"`
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
