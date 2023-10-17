package queryFirmwareFilesPage

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/commonService/queryFirmwareFilesPage"
	Disabled     = false
	EndPointName = "AppService.queryFirmwareFilesPage"
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
		FileId        valueTypes.Integer  `json:"file_id"`
		FileName      valueTypes.String   `json:"file_name"`
		FileSize      valueTypes.Integer  `json:"file_size"`
		FileType      valueTypes.Integer  `json:"file_type"`
		UploadTime    valueTypes.DateTime `json:"upload_time" PointNameDateFormat:"DateTimeLayout"`
		URL           valueTypes.String   `json:"url"`
		Operation     valueTypes.Integer  `json:"operation"`
		OperationTime valueTypes.DateTime `json:"operation_time" PointNameDateFormat:"DateTimeLayout"`
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

	for range only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
