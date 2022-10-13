package getPowerDeviceSetTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceSetTaskList"
const Disabled = false

type RequestData struct {
	Size    api.Integer `json:"size" required:"true"`
	CurPage api.Integer `json:"curPage" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		CommandStatus        api.Integer `json:"command_status"`
		CommandType          api.Integer `json:"command_type"`
		CreateTime           string      `json:"create_time"`
		OperateUserID        api.Integer `json:"operate_user_id"`
		OverTime             string      `json:"over_time"`
		PsID                 api.Integer `json:"ps_id"`
		SetCancelNum         api.Integer `json:"set_cancel_num"`
		SetFailNum           api.Integer `json:"set_fail_num"`
		SetFinishNum         api.Integer `json:"set_finish_num"`
		SetOvertimeNum       api.Integer `json:"set_overtime_num"`
		SetSuccessNum        api.Integer `json:"set_success_num"`
		SetTotalNum          api.Integer `json:"set_total_num"`
		SweepDevParamSetType api.Integer `json:"sweep_dev_param_set_type"`
		TaskID               api.Integer `json:"task_id"`
		TaskName             string      `json:"task_name"`
		TaskType             api.Integer `json:"task_type"`
		TemplateType         api.Integer `json:"template_type"`
		UpdateTime           string      `json:"update_time"`
		UserEnglishName      interface{} `json:"user_english_name"`
		UserName             string      `json:"user_name"`
		UUID                 api.Integer `json:"uuid"`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
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

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
// }

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", api.NewDateTime(""))
	}

	return entries
}
