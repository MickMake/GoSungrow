package getModuleLogTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/MickMake/GoUnify/Only"
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
	IsMore   api.Bool    `json:"isMore"`
	PageList []struct {
		BatchID        api.String      `json:"batch_id"`
		CommandStatus  api.Integer `json:"command_status"`
		CommandType    api.Integer `json:"command_type"`
		CreateTime     api.DateTime      `json:"create_time"`
		DeviceCode     api.Integer      `json:"device_code"`
		DeviceModel    api.String `json:"device_model"`
		DeviceModelID  api.String `json:"device_model_id"`
		ExpireSecond   api.Integer `json:"expire_second"`
		LogType        api.Integer      `json:"log_type"`
		LoggerCode     api.Integer      `json:"logger_code"`
		OperateUserID  api.Integer `json:"operate_user_id"`
		OverTime       api.DateTime      `json:"over_time"`
		Remark         api.String      `json:"remark"`
		SetCancelNum   api.Integer `json:"set_cancel_num"`
		SetFailNum     api.Integer `json:"set_fail_num"`
		SetFinishNum   api.Integer `json:"set_finish_num"`
		SetOvertimeNum api.Integer `json:"set_overtime_num"`
		SetSuccessNum  api.Integer `json:"set_success_num"`
		SetTotalNum    api.Integer `json:"set_total_num"`
		Sn             api.String      `json:"sn"`
		TaskID         api.Integer `json:"task_id"`
		TaskName       api.Integer      `json:"task_name"`
		UpdateTime     api.DateTime      `json:"update_time"`
	} `json:"pageList"`
	RowCount   api.Integer `json:"rowCount"`
	Size       api.Integer `json:"size"`
	StartIndex api.Integer `json:"startIndex"`
	TotalPage  api.Integer `json:"totalPage"`
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
