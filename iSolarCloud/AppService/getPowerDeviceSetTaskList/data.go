package getPowerDeviceSetTaskList

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceSetTaskList"
const Disabled = false

type RequestData struct {
	Size    int64 `json:"size" required:"true"`
	CurPage int64 `json:"curPage" required:"true"`
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
		CommandStatus        int64       `json:"command_status"`
		CommandType          int64       `json:"command_type"`
		CreateTime           string      `json:"create_time"`
		OperateUserID        int64       `json:"operate_user_id"`
		OverTime             string      `json:"over_time"`
		PsID                 int64       `json:"ps_id"`
		SetCancelNum         int64       `json:"set_cancel_num"`
		SetFailNum           int64       `json:"set_fail_num"`
		SetFinishNum         int64       `json:"set_finish_num"`
		SetOvertimeNum       int64       `json:"set_overtime_num"`
		SetSuccessNum        int64       `json:"set_success_num"`
		SetTotalNum          int64       `json:"set_total_num"`
		SweepDevParamSetType int64       `json:"sweep_dev_param_set_type"`
		TaskID               int64       `json:"task_id"`
		TaskName             string      `json:"task_name"`
		TaskType             int64       `json:"task_type"`
		TemplateType         int64       `json:"template_type"`
		UpdateTime           string      `json:"update_time"`
		UserEnglishName      interface{} `json:"user_english_name"`
		UserName             string      `json:"user_name"`
		UUID                 int64       `json:"uuid"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
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
