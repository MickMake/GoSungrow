package getModuleLogTaskList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
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
	CurPage  valueTypes.Integer `json:"curPage"`
	IsMore   valueTypes.Bool    `json:"isMore"`
	PageList []struct {
		BatchId        valueTypes.String      `json:"batch_id"`
		CommandStatus  valueTypes.Integer `json:"command_status"`
		CommandType    valueTypes.Integer `json:"command_type"`
		CreateTime     valueTypes.DateTime      `json:"create_time"`
		DeviceCode     valueTypes.Integer      `json:"device_code"`
		DeviceModel    valueTypes.String `json:"device_model"`
		DeviceModelId  valueTypes.String `json:"device_model_id"`
		ExpireSecond   valueTypes.Integer `json:"expire_second"`
		LogType        valueTypes.Integer      `json:"log_type"`
		LoggerCode     valueTypes.Integer      `json:"logger_code"`
		OperateUserId  valueTypes.Integer `json:"operate_user_id"`
		OverTime       valueTypes.DateTime      `json:"over_time"`
		Remark         valueTypes.String      `json:"remark"`
		SetCancelNum   valueTypes.Integer `json:"set_cancel_num"`
		SetFailNum     valueTypes.Integer `json:"set_fail_num"`
		SetFinishNum   valueTypes.Integer `json:"set_finish_num"`
		SetOvertimeNum valueTypes.Integer `json:"set_overtime_num"`
		SetSuccessNum  valueTypes.Integer `json:"set_success_num"`
		SetTotalNum    valueTypes.Integer `json:"set_total_num"`
		Sn             valueTypes.String      `json:"sn"`
		TaskId         valueTypes.Integer `json:"task_id"`
		TaskName       valueTypes.Integer      `json:"task_name"`
		UpdateTime     valueTypes.DateTime      `json:"update_time"`
	} `json:"pageList" PointNameFromChild:"TaskId" PointNameAppend:"false" PointArrayFlatten:"false"`
	RowCount   valueTypes.Integer `json:"rowCount"`
	Size       valueTypes.Integer `json:"size"`
	StartIndex valueTypes.Integer `json:"startIndex"`
	TotalPage  valueTypes.Integer `json:"totalPage"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
