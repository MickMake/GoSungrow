package queryDeviceRealTimeDataByPsKeys

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
)

const Url = "/v1/devService/queryDeviceRealTimeDataByPsKeys"
const Disabled = false

type RequestData struct {
	PsKeyList string `json:"ps_key_list" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	Dummy string `json:"dummy"`
}

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

	// for range Only.Once {
	// 	table = output.NewTable()
	// 	e.Error = table.SetTitle("")
	// 	if e.Error != nil {
	// 		break
	// 	}
	//
	// 	_ = table.SetHeader(
	// 		"Date",
	// 		"Point Id",
	// 		"Point Group Name",
	// 		"Point Name",
	// 		"Value",
	// 		"Unit",
	// 	)
	//
	// 	for _, d := range e.Response.ResultData.PageList {
	// 		for _, p := range d.PointData {
	// 			_ = table.AddRow(
	// 				api.NewDateTime(p.TimeStamp).PrintFull(),
	// 				fmt.Sprintf("%s.%d", d.PsKey, p.PointID),
	// 				p.PointGroupName,
	// 				p.PointName,
	// 				p.Value,
	// 				p.Unit,
	// 			)
	// 		}
	// 	}
	//
	// 	table.InitGraph(output.GraphRequest {
	// 		Title:        "",
	// 		TimeColumn:   output.SetInteger(1),
	// 		SearchColumn: output.SetInteger(2),
	// 		NameColumn:   output.SetInteger(4),
	// 		ValueColumn:  output.SetInteger(5),
	// 		UnitsColumn:  output.SetInteger(6),
	// 		SearchString: output.SetString(""),
	// 		MinLeftAxis:  output.SetFloat(0),
	// 		MaxLeftAxis:  output.SetFloat(0),
	// 	})
	//
	// }

	return table
}

func (e *EndPoint) GetData() api.DataMap {
	var ret api.DataMap

	// for range Only.Once {
	// 	index := 0
	// 	for _, d := range e.Response.ResultData.PageList {
	// 		for _, p := range d.PointData {
	// 			if p.Unit == "W" {
	// 				fv, err := strconv.ParseFloat(p.Value, 64)
	// 				fv = fv / 1000
	// 				if err == nil {
	// 					p.Value = fmt.Sprintf("%.3f", fv)
	// 					p.Unit = "kW"
	// 				}
	// 			}
	//
	// 			ret.Entries = append(ret.Entries, api.DataEntry {
	// 				Date:           api.NewDateTime(p.TimeStamp),
	// 				PointId:        api.GetAreaPointName(d.PsKey, p.PointID),
	// 				PointGroupName: p.PointGroupName,
	// 				PointName:      p.PointName,
	// 				Value:          p.Value,
	// 				Unit:           p.Unit,
	// 				ValueType:      api.GetPointType(d.PsKey, p.PointID),
	// 				Index:          index,
	// 			})
	//
	// 			index++
	// 		}
	// 	}
	// }

	return ret
}
