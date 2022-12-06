package getPowerDeviceSetTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceSetTaskList"
const Disabled = false
const EndPointName = "AppService.getPowerDeviceSetTaskList"

type RequestData struct {
	Size    valueTypes.Integer `json:"size" required:"true"`
	CurPage valueTypes.Integer `json:"curPage" required:"true"`
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
		GoStruct             GoStruct.GoStruct   `json:"GoStruct" PointIdFrom:"TaskId" PointIdReplace:"true" PointDeviceFrom:"PsId"`

		TaskId               valueTypes.Integer  `json:"task_id"`
		TaskType             valueTypes.Integer  `json:"task_type"`
		TaskName             valueTypes.String   `json:"task_name"`
		UUID                 valueTypes.Integer  `json:"uuid"`
		PsId                 valueTypes.PsId     `json:"ps_id"`
		UserEnglishName      interface{}         `json:"user_english_name"`
		UserName             valueTypes.String   `json:"user_name"`
		OperateUserId        valueTypes.Integer  `json:"operate_user_id"`
		CommandStatus        valueTypes.Integer  `json:"command_status"`
		CommandType          valueTypes.Integer  `json:"command_type"`
		CreateTime           valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		UpdateTime           valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		OverTime             valueTypes.DateTime `json:"over_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		SetCancelNum         valueTypes.Count    `json:"set_cancel_num"`
		SetFailNum           valueTypes.Count    `json:"set_fail_num"`
		SetFinishNum         valueTypes.Count    `json:"set_finish_num"`
		SetOvertimeNum       valueTypes.Count    `json:"set_overtime_num"`
		SetSuccessNum        valueTypes.Count    `json:"set_success_num"`
		SetTotalNum          valueTypes.Count    `json:"set_total_num"`
		SweepDevParamSetType valueTypes.Integer  `json:"sweep_dev_param_set_type"`
		TemplateType         valueTypes.Integer  `json:"template_type"`
	} `json:"pageList" PointId:"tasks" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
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
