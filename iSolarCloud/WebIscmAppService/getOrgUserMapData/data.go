package getOrgUserMapData

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/faultService/getOrgUserMapData"
	Disabled     = false
	EndPointName = "WebIscmAppService.getOrgUserMapData"
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
	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}
	return entries
}
