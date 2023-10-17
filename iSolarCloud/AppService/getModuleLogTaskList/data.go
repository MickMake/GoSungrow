package getModuleLogTaskList

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/integrationService/getModuleLogTaskList"
	Disabled     = false
	EndPointName = "AppService.getModuleLogTaskList"
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
	CurPage  valueTypes.Integer `json:"curPage" PointId:"cur_page"`
	IsMore   valueTypes.Bool    `json:"isMore" PointId:"is_more"`
	PageList []struct {
		// PointIdFromChild:"TaskId" PointIdReplace:"true"
		GoStruct GoStruct.GoStruct `json:"-" PointIdReplace:"true" PointIdFrom:"TaskId" PointNameDateFormat:"DateTimeAltLayout" PointTimestampFrom:"CreateTime"`

		TaskId         valueTypes.Integer  `json:"task_id"`
		CreateTime     valueTypes.DateTime `json:"create_time" PointNameDateFormat:"DateTimeLayout"`
		OverTime       valueTypes.DateTime `json:"over_time" PointNameDateFormat:"DateTimeLayout"`
		UpdateTime     valueTypes.DateTime `json:"update_time" PointNameDateFormat:"DateTimeLayout"`
		ExpireSecond   valueTypes.Integer  `json:"expire_second"`
		TaskName       valueTypes.String   `json:"task_name"`
		CommandStatus  valueTypes.Integer  `json:"command_status"`
		CommandType    valueTypes.Integer  `json:"command_type"`
		BatchId        valueTypes.String   `json:"batch_id"`
		Sn             valueTypes.String   `json:"sn"`
		DeviceCode     valueTypes.Integer  `json:"device_code"`
		DeviceModel    valueTypes.String   `json:"device_model"`
		DeviceModelId  valueTypes.String   `json:"device_model_id"`
		LogType        valueTypes.Integer  `json:"log_type"`
		LoggerCode     valueTypes.Integer  `json:"logger_code"`
		OperateUserId  valueTypes.Integer  `json:"operate_user_id"`
		SetCancelNum   valueTypes.Bool     `json:"set_cancel_num"`
		SetFailNum     valueTypes.Bool     `json:"set_fail_num"`
		SetFinishNum   valueTypes.Bool     `json:"set_finish_num"`
		SetOvertimeNum valueTypes.Bool     `json:"set_overtime_num"`
		SetSuccessNum  valueTypes.Bool     `json:"set_success_num"`
		SetTotalNum    valueTypes.Bool     `json:"set_total_num"`
		Remark         valueTypes.String   `json:"remark"`
	} `json:"pageList" PointId:"page_list" DataTable:"true"`
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
