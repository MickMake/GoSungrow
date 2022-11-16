package getPowerDevicePointNames

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

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
	DeviceType22 = "22"
	DeviceType23 = "23"
	DeviceType26 = "26"
	DeviceType37 = "37"
	DeviceType41 = "41"
	DeviceType43 = "43"
	DeviceType47 = "47"
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
	DeviceType22,
	DeviceType23,
	DeviceType26,
	DeviceType37,
	DeviceType41,
	DeviceType43,
	DeviceType47,
}

type RequestData struct {
	DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
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
	PointCalType valueTypes.Integer  `json:"point_cal_type"`
	PointId      valueTypes.Integer  `json:"point_id"`
	PointName    valueTypes.String `json:"point_name"`
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

func (e *EndPoint) GetPointDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable(
			"Device Type",
			"Point Type",
			"Point Id",
			"Point Name",
		)
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		// e.Error = table.SetHeader(
		// 	"Device Type",
		// 	"Point Type",
		// 	"Point Id",
		// 	"Point Name",
		// )
		if e.Error != nil {
			break
		}

		table.SetFilePrefix(e.Request.DeviceType.String())

		for _, p := range e.Response.ResultData {
			_ = table.AddRow(e.Request.DeviceType.String(),
				p.PointCalType.String(),
				p.PointId.String(),
				p.PointName.String(),
			)
			if table.Error != nil {
				continue
			}
		}

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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
