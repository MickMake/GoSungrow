package queryParamSettingTask

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
)

const Url = "/v1/devService/queryParamSettingTask"
const Disabled = false
const EndPointName = "AppService.queryParamSettingTask"

type RequestData struct {
	TaskId valueTypes.String  `json:"task_id" required:"true"`
	Uuid   valueTypes.Integer `json:"uuid" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	TaskId         valueTypes.Integer  `json:"task_id"`
	TaskName       valueTypes.String   `json:"task_name"`
	CommandStatus  valueTypes.Integer  `json:"command_status"`
	CreateTime     valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	CreateTimeZone valueTypes.DateTime `json:"create_time_zone" PointNameDateFormat:"2006-01-02 15:04:05"`
	OverTime       valueTypes.DateTime `json:"over_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	OverTimeZone   valueTypes.DateTime `json:"over_time_zone" PointNameDateFormat:"2006-01-02 15:04:05"`

	ParamList      []struct {
		// GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableIndex:"true"`

		PointId                 valueTypes.PointId  `json:"point_id"`
		PointName               valueTypes.String   `json:"point_name"`
		Unit                    valueTypes.String   `json:"unit"`
		SetValue                valueTypes.String   `json:"set_value"`
		SetPrecision            valueTypes.String   `json:"set_precision"`
		ReturnValue             valueTypes.String   `json:"return_value"`
		SetValName              valueTypes.String   `json:"set_val_name"`
		SetValNameVal           valueTypes.String   `json:"set_val_name_val"`
		CommandStatus           valueTypes.Integer  `json:"command_status"`
		CreateTime              valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		CreateTimeZone          valueTypes.DateTime `json:"create_time_zone" PointNameDateFormat:"2006-01-02 15:04:05"`
		UpdateTime              valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006-01-02 15:04:05"`
		UpdateTimeZone          valueTypes.DateTime `json:"update_time_zone" PointNameDateFormat:"2006-01-02 15:04:05"`
		ParamCode               valueTypes.String   `json:"param_code"`
	} `json:"param_list" DataTable:"true" DataTableIndex:"true"`
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
