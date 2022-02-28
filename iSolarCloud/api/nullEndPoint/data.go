package nullEndPoint

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
)


const Url = "%URL%"
const Disabled = true

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

// IsValid Checks for validity of results data.
func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

// Help provides more info to the user on request JSON fields.
func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

// ResultData holds data returned from the API.
type ResultData struct {
	Dummy string `json:"dummy"`
}

// IsValid Checks for validity of results data.
func (e *ResultData) IsValid() error {
	var err error
	switch {
		case e.Dummy == "":
			break
		default:
			err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	}
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable()
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		// e.Error = table.SetHeader(
		// 	"Template Id",
		// 	"Template Name",
		// 	"Update On",
		// )
		// if e.Error != nil {
		// 	break
		// }

		// for _, p := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		p.TemplateID,
		// 		p.TemplateName,
		// 		api.NewDateTime(p.UpdateTime).PrintFull(),
		// 	)
		// 	if table.Error != nil {
		// 		continue
		// 	}
		// }

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(3),
		// 	ValueColumn:  output.SetInteger(4),
		// 	UnitsColumn:  output.SetInteger(5),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}

	return table
}
