package getPowerDeviceSetTaskDetailList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/devService/getPowerDeviceSetTaskDetailList"
const Disabled = false
const EndPointName = "AppService.getPowerDeviceSetTaskDetailList"

type RequestData struct {
	QueryType valueTypes.String `json:"query_type" required:"true"`
	TaskId    valueTypes.String `json:"task_id" required:"true"`
	Uuid      valueTypes.String `json:"uuid" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("query_type: 2, 3, 4")
	return ret
}

type ResultData struct {
	DeviceList     []interface{} `json:"device_list"`
	PageList       []interface{} `json:"pageList" PointId:"page_list" PointIdReplace:"true" PointArrayFlatten:"false"`
	PsNameInfoList []interface{} `json:"ps_name_info_list"`
	RowCount       valueTypes.Integer   `json:"rowCount" PointId:"row_count"`
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
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
