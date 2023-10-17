package queryMutiPointDataList

import (
	"encoding/json"
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/commonService/queryMutiPointDataList"
	Disabled     = false
	EndPointName = "AppService.queryMutiPointDataList"
)

type RequestData struct {
	PsId           valueTypes.PsId     `json:"ps_id" required:"true"`
	StartTimeStamp valueTypes.DateTime `json:"start_time_stamp" required:"true"`
	EndTimeStamp   valueTypes.DateTime `json:"end_time_stamp" required:"true"`
	MinuteInterval valueTypes.Integer  `json:"minute_interval" required:"true"`
	PsKeys         valueTypes.PsKeys   `json:"ps_key" required:"true"`
	Points         valueTypes.PointIds `json:"points" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Data Data `json:"data" DataTable:"true" DataTableSortOn:"Timestamp"`
}

type (
	Data  map[valueTypes.DateTime]Value
	Value struct {
		GoStructParent GoStruct.GoStructParent `json:"-" PointIdFrom:"PsKey.Timestamp" PointIdReplace:"true" PointDeviceFrom:"PsKey"`
		// GoStruct  GoStruct.GoStruct   `json:"-" PointDeviceFrom:"PsKey"`

		Timestamp valueTypes.DateTime           `json:"timestamp" PointNameDateFormat:"DateTimeLayout"`
		PsKey     valueTypes.PsKey              `json:"ps_key"`
		Points    map[string]valueTypes.Generic `json:"points" PointDeviceFrom:"PsKey"`
	}
)

func (e *ResultData) IsValid() error {
	var err error
	return err
}

// func (e *ResultData) String() string {
// 	var ret string
//
// 	// for range only.Once {
// 	// 	if len(data) == 0 {
// 	// 		break
// 	// 	}
// 	//
// 	// 	e.Points = make(Points)
// 	//
// 	// 	var du dDevices
// 	// 	// Store DeviceData.Points
// 	// 	_ = json.Unmarshal(data, &du)
// 	// 	for deviceName, deviceRef := range du {
// 	// 		fmt.Printf("%s =>\n", deviceName)
// 	//
// 	// 		for pointName, pointRef := range deviceRef {
// 	//
// 	// 			for time, value := range pointRef {
// 	// 				e.Points[pointName].List = append(e.Points[pointName].List)
// 	// 				fmt.Printf("%s, %s, %s, %s\n", deviceName, pointName, time, value)
// 	// 				// if k == "" {
// 	// 				// 	delete(du, i)
// 	// 				// 	continue
// 	// 				// }
// 	// 			}
// 	// 		}
// 	// 	}
// 	//
// 	// 	// var dp DecodePoints
// 	// 	// // Store DeviceData.Points
// 	// 	// _ = json.Unmarshal(data, &dp)
// 	// 	// for i, k := range dp {
// 	// 	// 	if k == nil {
// 	// 	// 		delete(dp, i)
// 	// 	// 		continue
// 	// 	// 	}
// 	// 	// 	i = strings.TrimSuffix(i, "List")
// 	// 	// 	u := du[i+"_unit"]
// 	// 	// 	e.Points[i] = Point{
// 	// 	// 		List: k,
// 	// 	// 		Unit: u,
// 	// 	// 	}
// 	// 	// }
// 	// }
//
// 	return ret
// }

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		type scan scanDevices
		d := make(scan, 0)
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		e.Data = make(Data, 0)
		for device := range d {
			for point := range d[device] {
				for value := range d[device][point] {
					d[device][point][value].PointId = valueTypes.SetPointIdString(point)
					ts := d[device][point][value].Timestamp

					if _, ok := e.Data[ts]; !ok {
						rdv := Value{
							Timestamp: ts,
							PsKey:     valueTypes.SetPsKeyString(device),
							Points:    make(map[string]valueTypes.Generic),
						}
						rdv.Points[point] = d[device][point][value].Value
						e.Data[ts] = rdv
						continue
					}

					e.Data[ts].Points[point] = d[device][point][value].Value
				}
			}
		}
	}

	return err
}

// Scan incoming JSON.

type scanDevices map[string]scanPoints

func (e *scanDevices) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		type scan scanDevices
		d := make(scan, 0)
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		*e = scanDevices(d)
	}

	return err
}

type scanPoints map[string]scanValues

func (e *scanPoints) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		type scan scanPoints
		d := make(scan, 0)
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		for point := range d {
			for value := range d[point] {
				d[point][value].PointId = valueTypes.SetPointIdString(point)
			}
		}
		*e = scanPoints(d)
	}

	return err
}

type scanValues []scanValue

func (e *scanValues) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		type scan map[string]string
		d := make(scan, 0)

		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		*e = make(scanValues, 0)
		for k, v := range d {
			*e = append(*e, scanValue{
				Timestamp: valueTypes.SetDateTimeString(k),
				// PointId:   valueTypes.SetPointIdString(v),
				Value: valueTypes.SetGenericString(v),
			})
		}
	}

	return err
}

type scanValue struct {
	Timestamp valueTypes.DateTime `json:"timestamp"`
	PointId   valueTypes.PointId  `json:"point_id"`
	Value     valueTypes.Generic  `json:"value"`
}

// func (e *EndPoint) GetPointDataTable(points api.TemplatePoints) output.Table {
// 	var table output.Table
//
// 	// for range only.Once {
// 	// 	table = output.NewTable(
// 	// 		"Date/Time",
// 	// 		"Point Id",
// 	// 		"Point Name",
// 	// 		"Value",
// 	// 		"Units",
// 	// 	)
// 	// 	table.SetTitle("")
// 	// 	table.SetJson([]byte(e.GetJsonData(false)))
// 	// 	table.SetRaw([]byte(e.GetJsonData(true)))
// 	//
// 	// 	// e.Error = table.SetHeader(
// 	// 	// 	"Date/Time",
// 	// 	// 	"Point Id",
// 	// 	// 	"Point Name",
// 	// 	// 	"Value",
// 	// 	// 	"Units",
// 	// 	// )
// 	// 	if e.Error != nil {
// 	// 		break
// 	// 	}
// 	//
// 	// 	t := e.Request.RequestData.StartTimeStamp
// 	// 	e.SetFilenamePrefix(t.String())
// 	// 	table.SetFilePrefix(t.String())
// 	//
// 	// 	for deviceName, deviceRef := range e.Response.ResultData.Devices {
// 	// 		for pointId, pointRef := range deviceRef.Points {
// 	// 			for _, tim := range pointRef.Times {
// 	// 				gp := points.GetPoint(deviceName, pointId)
// 	// 				_ = table.AddRow(tim.Key.PrintFull(),
// 	// 					fmt.Sprintf("%s.%s", deviceName, pointId),
// 	// 					gp.Name,
// 	// 					tim.Value,
// 	// 					gp.Unit,
// 	// 				)
// 	// 				if table.Error != nil {
// 	// 					continue
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}
// 	//
// 	// 	table.InitGraph(output.GraphRequest {
// 	// 		Title:        "",
// 	// 		TimeColumn:   output.SetString("Date/Time"),
// 	// 		SearchColumn: output.SetString("Point Id"),
// 	// 		NameColumn:   output.SetString("Point Name"),
// 	// 		ValueColumn:  output.SetString("Value"),
// 	// 		UnitsColumn:  output.SetString("Units"),
// 	// 		SearchString: output.SetString(""),
// 	// 		MinLeftAxis:  output.SetFloat(0),
// 	// 		MaxLeftAxis:  output.SetFloat(0),
// 	// 	})
// 	//
// 	// }
//
// 	return table
// }

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	// entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))
	entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.EndPointPath{})

	// table.InitGraph(output.GraphRequest {
	// 	Title:        "",
	// 	TimeColumn:   output.SetString("Date/Time"),
	// 	SearchColumn: output.SetString("Point Id"),
	// 	NameColumn:   output.SetString("Point Name"),
	// 	ValueColumn:  output.SetString("Value"),
	// 	UnitsColumn:  output.SetString("Units"),
	// 	SearchString: output.SetString(""),
	// 	MinLeftAxis:  output.SetFloat(0),
	// 	MaxLeftAxis:  output.SetFloat(0),
	// })

	return entries
}
