package getTemplateList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/getTemplateList"
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
	PageList []struct {
		TemplateId   valueTypes.Integer  `json:"template_id"`
		TemplateName valueTypes.String `json:"template_name"`
		UpdateTime   valueTypes.DateTime `json:"update_time"`
	} `json:"pageList" PointNameFromChild:"TemplateId" PointNameAppend:"false" PointArrayFlatten:"false" DataTable:"true"`
	RowCount valueTypes.Integer `json:"rowCount"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e,  "system", nil)
	}

	return entries
}
