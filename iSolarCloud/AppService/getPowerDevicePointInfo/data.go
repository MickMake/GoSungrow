package getPowerDevicePointInfo

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
)

const Url = "/v1/reportService/getPowerDevicePointInfo"
const Disabled = false

type RequestData struct {
	Id string `json:"id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	DeviceType    int64  `json:"device_type" PointId:"device_type" PointType:""`
	ID            int64  `json:"id" PointId:"id" PointType:""`
	Period        int64  `json:"period" PointId:"period" PointType:""`
	PointID       int64  `json:"point_id" PointId:"point_id" PointType:""`
	PointName     string `json:"point_name" PointId:"point_name" PointType:""`
	ShowPointName string `json:"show_point_name" PointId:"show_point_name" PointType:""`
	TranslationID int64  `json:"translation_id" PointId:"translation_id" PointType:""`
}

func (e *ResultData) IsValid() error {
	var err error
		// switch {
		// 	case e.Dummy == "":
		// 		break
		// 	default:
		// 		err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
		// }
	return err
}

func (e *EndPoint) GetData() ResultData {
	return e.Response.ResultData
}

// func (e *EndPoint) GetData() api.DataMap {
// 	return e.Response.ResultData.GetData()
// }
//
// func (e *ResultData) GetData() api.DataMap {
// 	entries := api.NewDataMap()
//
// 	for range Only.Once {
// 		entries.StructToPoints("", *e)
// 	}
// 	return entries
// }

func (e *EndPoint) AddDataTable(table output.Table) output.Table {

	for range Only.Once {
		rd := e.Response.ResultData

		if rd.ID == 0 {
			break
		}
		_ = table.AddRow(rd.DeviceType, rd.ID, rd.Period, rd.PointID,rd.PointName, rd.ShowPointName, rd.TranslationID)
	}
	return table
}

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		_ = table.SetHeader(
			"DeviceType",
			"Id",
			"Period",
			"Point Id",
			"Point Name",
			"Show Point Name",
			"Translation Id",
		)
		rd := e.Response.ResultData
		_ = table.AddRow(rd.DeviceType, rd.ID, rd.Period, rd.PointID,rd.PointName, rd.ShowPointName, rd.TranslationID)
	}
	return table
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
