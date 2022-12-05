package queryDutyType

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/otherService/queryDutyType"
const Disabled = false
const EndPointName = "WebIscmAppService.queryDutyType"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent  `json:"-" DataTable:"true" DataTableSortOn:"CodeId"`

	CodeId     valueTypes.Integer `json:"code_id"`
	CodeName   valueTypes.String  `json:"code_name"`
	CodeType   valueTypes.Integer `json:"code_type"`
	CodeValue  valueTypes.String  `json:"code_value"`
	CodeValue2 interface{}        `json:"code_value2"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
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
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
