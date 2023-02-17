package getPowerDevicePointInfo

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/reportService/getPowerDevicePointInfo"
const Disabled = false
const EndPointName = "AppService.getPowerDevicePointInfo"

type RequestData struct {
	Id valueTypes.Integer `json:"id" required:"true"`
	// Id valueTypes.String `json:"id"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"Id"`

	DeviceType    valueTypes.Integer `json:"device_type" PointId:"device_type" PointValueType:"" PointUpdateFreq:""`
	Id            valueTypes.Integer `json:"id" PointId:"id" PointValueType:"" PointUpdateFreq:""`
	Period        valueTypes.Integer `json:"period" PointId:"period" PointValueType:"" PointUpdateFreq:""`
	PointId       valueTypes.Integer `json:"point_id" PointId:"point_id" PointValueType:"" PointUpdateFreq:""`
	PointName     valueTypes.String  `json:"point_name" PointId:"point_name" PointValueType:"" PointUpdateFreq:""`
	ShowPointName valueTypes.String  `json:"show_point_name" PointId:"show_point_name" PointValueType:"" PointUpdateFreq:""`
	TranslationId valueTypes.Integer `json:"translation_id" PointId:"translation_id" PointValueType:"" PointUpdateFreq:""`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) AddDataTable(table output.Table) output.Table {

	for range Only.Once {
		rd := e.Response.ResultData

		if rd.Id.Value() == 0 {
			break
		}
		_ = table.AddRow(rd.DeviceType.String(),
			rd.Id.String(),
			rd.Period.String(),
			rd.PointId.String(),
			rd.PointName.String(),
			rd.ShowPointName.String(),
			rd.TranslationId.String(),
		)
	}
	return table
}

func (e *EndPoint) GetPointDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable(
			"DeviceType",
			"Id",
			"Period",
			"Point Id",
			"Point Name",
			"Show Point Name",
			"Translation Id",
		)
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		// _ = table.SetHeader(
		// 	"DeviceType",
		// 	"Id",
		// 	"Period",
		// 	"Point Id",
		// 	"Point Name",
		// 	"Show Point Name",
		// 	"Translation Id",
		// )
		rd := e.Response.ResultData
		_ = table.AddRow(rd.DeviceType.String(),
			rd.Id.String(),
			rd.Period.String(),
			rd.PointId.String(),
			rd.PointName.String(),
			rd.ShowPointName.String(),
			rd.TranslationId.String(),
		)
	}
	return table
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
