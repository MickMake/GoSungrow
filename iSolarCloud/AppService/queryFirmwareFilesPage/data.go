package queryFirmwareFilesPage

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/commonService/queryFirmwareFilesPage"
const Disabled = false

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
		FileId        valueTypes.Integer  `json:"file_id"`
		FileName      valueTypes.String   `json:"file_name"`
		FileSize      valueTypes.Integer  `json:"file_size"`
		FileType      valueTypes.Integer  `json:"file_type"`
		UploadTime    valueTypes.DateTime `json:"upload_time" PointNameDateFormat:"2006/01/02 15:04:05"`
		URL           valueTypes.String   `json:"url"`
		Operation     valueTypes.Integer  `json:"operation"`
		OperationTime valueTypes.DateTime `json:"operation_time" PointNameDateFormat:"2006/01/02 15:04:05"`
		OperatorId    valueTypes.String   `json:"operator_id"`
		OperatorName  valueTypes.String   `json:"operator_name"`
		System        valueTypes.String   `json:"system"`
	} `json:"pageList" PointId:"page_list" DataTable:"true"`
	Code     valueTypes.String  `json:"code"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
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
