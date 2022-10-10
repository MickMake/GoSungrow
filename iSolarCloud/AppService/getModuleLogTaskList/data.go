package getModuleLogTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/integrationService/getModuleLogTaskList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	CurPage  api.Integer `json:"curPage"`
	IsMore   interface{} `json:"isMore"`
	PageList []struct {
		BatchID        string      `json:"batch_id"`
		CommandStatus  api.Integer `json:"command_status"`
		CommandType    api.Integer `json:"command_type"`
		CreateTime     string      `json:"create_time"`
		DeviceCode     string      `json:"device_code"`
		DeviceModel    interface{} `json:"device_model"`
		DeviceModelID  interface{} `json:"device_model_id"`
		ExpireSecond   api.Integer `json:"expire_second"`
		LogType        string      `json:"log_type"`
		LoggerCode     string      `json:"logger_code"`
		OperateUserID  api.Integer `json:"operate_user_id"`
		OverTime       string      `json:"over_time"`
		Remark         string      `json:"remark"`
		SetCancelNum   api.Integer `json:"set_cancel_num"`
		SetFailNum     api.Integer `json:"set_fail_num"`
		SetFinishNum   api.Integer `json:"set_finish_num"`
		SetOvertimeNum api.Integer `json:"set_overtime_num"`
		SetSuccessNum  api.Integer `json:"set_success_num"`
		SetTotalNum    api.Integer `json:"set_total_num"`
		Sn             string      `json:"sn"`
		TaskID         api.Integer `json:"task_id"`
		TaskName       string      `json:"task_name"`
		UpdateTime     string      `json:"update_time"`
	} `json:"pageList"`
	RowCount   api.Integer `json:"rowCount"`
	Size       api.Integer `json:"size"`
	StartIndex interface{} `json:"startIndex"`
	TotalPage  interface{} `json:"totalPage"`
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
