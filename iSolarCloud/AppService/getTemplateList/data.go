package getTemplateList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
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
		TemplateID   api.Integer  `json:"template_id"`
		TemplateName api.String `json:"template_name"`
		UpdateTime   api.DateTime `json:"update_time"`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		e.Error = table.SetHeader(
			"Template Id",
			"Template Name",
			"Update On",
		)
		if e.Error != nil {
			break
		}

		for _, p := range e.Response.ResultData.PageList {
			_ = table.AddRow(
				p.TemplateID,
				p.TemplateName,
				// api.NewDateTime(p.UpdateTime).PrintFull(),
				p.UpdateTime.String(),
			)
			if table.Error != nil {
				continue
			}
		}
	}

	return table
}
