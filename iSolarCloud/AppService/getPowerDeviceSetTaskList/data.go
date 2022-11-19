package getPowerDeviceSetTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getPowerDeviceSetTaskList"
const Disabled = false

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
		CreateTime           valueTypes.DateTime `json:"create_time"`
		UpdateTime           valueTypes.DateTime `json:"update_time"`
		OverTime             valueTypes.DateTime `json:"over_time"`
		SetCancelNum         valueTypes.Count    `json:"set_cancel_num"`
		SetFailNum           valueTypes.Count    `json:"set_fail_num"`
		SetFinishNum         valueTypes.Count    `json:"set_finish_num"`
		SetOvertimeNum       valueTypes.Count    `json:"set_overtime_num"`
		SetSuccessNum        valueTypes.Count    `json:"set_success_num"`
		SetTotalNum          valueTypes.Count    `json:"set_total_num"`
		SweepDevParamSetType valueTypes.Integer  `json:"sweep_dev_param_set_type"`
		TemplateType         valueTypes.Integer  `json:"template_type"`
	} `json:"pageList" PointId:"page_list" PointIdFromChild:"PsId" PointIdReplace:"true" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount" PointId:"row_count"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
