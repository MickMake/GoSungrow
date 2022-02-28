package getPowerDevicePointNames

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
)

const Url = "/v1/reportService/getPowerDevicePointNames"
const Disabled = false

const (
	DeviceType1  = "1"
	DeviceType3  = "3"
	DeviceType4  = "4"
	DeviceType5  = "5"
	DeviceType7  = "7"
	DeviceType11 = "11"
	DeviceType14 = "14"
	DeviceType17 = "17"
)

var DeviceTypes = []string{
	DeviceType1,
	DeviceType3,
	DeviceType4,
	DeviceType5,
	DeviceType7,
	DeviceType11,
	DeviceType14,
	DeviceType17,
}

type RequestData struct {
	DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("device_type can be one of:\n")
	ret += fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s\n",
		DeviceType1,
		DeviceType3,
		DeviceType4,
		DeviceType5,
		DeviceType7,
		DeviceType11,
		DeviceType14,
		DeviceType17,
	)
	return ret
}

type ResultData []struct {
	PointCalType int64  `json:"point_cal_type"`
	PointID      int64  `json:"point_id"`
	PointName    string `json:"point_name"`
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
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		e.Error = table.SetHeader(
			"Device Type",
			"Point Type",
			"Point Id",
			"Point Name",
		)
		if e.Error != nil {
			break
		}

		e.Error = table.SetFilePrefix(e.Request.DeviceType)
		if e.Error != nil {
			break
		}

		for _, p := range e.Response.ResultData {
			_ = table.AddRow(
				e.Request.DeviceType,
				p.PointCalType,
				p.PointID,
				p.PointName,
			)
			if table.Error != nil {
				continue
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	ValueColumn:  output.SetInteger(4),
		// 	UnitsColumn:  output.SetInteger(5),
		// 	SearchColumn: output.SetInteger(2),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}

	return table
}
