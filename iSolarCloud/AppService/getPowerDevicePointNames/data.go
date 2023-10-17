package getPowerDevicePointNames

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/reportService/getPowerDevicePointNames"
	Disabled     = false
	EndPointName = "AppService.getPowerDevicePointNames"
)

const (
	DeviceType1  = 1
	DeviceType3  = 3
	DeviceType4  = 4
	DeviceType5  = 5
	DeviceType7  = 7
	DeviceType11 = 11
	DeviceType14 = 14
	DeviceType17 = 17
	DeviceType22 = 22
	DeviceType23 = 23
	DeviceType26 = 26
	DeviceType37 = 37
	DeviceType41 = 41
	DeviceType43 = 43
	DeviceType47 = 47
)

var DeviceTypes = []int{
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
	DeviceType valueTypes.Integer `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("device_type can be one of:\n")
	ret += fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d\n",
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

type ResultData []Point

type Point struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true"`

	PointId      valueTypes.Integer `json:"point_id"`
	PointName    valueTypes.String  `json:"point_name"`
	PointCalType valueTypes.Integer `json:"point_cal_type"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

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
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
