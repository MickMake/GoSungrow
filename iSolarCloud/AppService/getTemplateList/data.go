package getTemplateList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getTemplateList"
const Disabled = false
const EndPointName = "AppService.getTemplateList"

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
	PageList []Template `json:"pageList" PointId:"page_list" DataTable:"true" DataTableSortOn:"UpdateTime"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

type Template struct {
	TemplateId   valueTypes.Integer  `json:"template_id"`
	TemplateName valueTypes.String `json:"template_name"`
	UpdateTime   valueTypes.DateTime `json:"update_time" PointNameDateFormat:"DateTimeLayout"`
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
