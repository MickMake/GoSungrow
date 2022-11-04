package queryMutiPointDataList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"

	"encoding/json"
	"fmt"
	"sort"
)

const Url = "/v1/commonService/queryMutiPointDataList"
const Disabled = false

type RequestData struct {
	PsId           valueTypes.PsId  `json:"ps_id" required:"true"`
	PsKey          valueTypes.PsKey `json:"ps_key" required:"true"`
	Points         valueTypes.String `json:"points" required:"true"`
	MinuteInterval valueTypes.String `json:"minute_interval" required:"true"`
	StartTimeStamp valueTypes.String `json:"start_time_stamp" required:"true"`
	EndTimeStamp   valueTypes.String `json:"end_time_stamp" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Devices Devices `json:"devices"`
}

type Devices map[string]Device
type Device struct {
	Points Points `json:"points"`
}
type Points map[valueTypes.PointId]Point
type Point struct {
	Name  valueTypes.String `json:"name"`
	Units valueTypes.String `json:"units"`
	Times Times  `json:"times"`
}
type Times []Time
type Time struct {
	Key   valueTypes.DateTime `json:"key"`
	Value string       `json:"value"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

type dDevices map[string]dPoints
type dPoints map[string]dTimes
type dTimes map[string]string

type DecodeResultData ResultData

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		d := make(dDevices)

		// Store DeviceData.Points.
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		e.Devices = make(Devices)
		for deviceName, deviceRef := range d {
			points := Points{}
			for pointName, pointRef := range deviceRef {
				times := Times{}
				for time, value := range pointRef {
					times = append(times, Time{
						Key:   valueTypes.NewDateTime(time),
						Value: value,
					})
				}

				sort.Slice(times, func(i, j int) bool {
					return times[i].Key.Before(times[j].Key.Time)
				})
				points[valueTypes.SetPointIdString(pointName)] = Point{
					Name:  valueTypes.SetStringValue(""),
					Units: valueTypes.SetStringValue(""),
					Times: times,
				}
			}

			e.Devices[deviceName] = Device{
				Points: points,
			}
		}
	}

	return err
}

func (e *ResultData) String() string {
	var ret string

	// for range Only.Once {
	// 	if len(data) == 0 {
	// 		break
	// 	}
	//
	// 	e.Points = make(Points)
	//
	// 	var du dDevices
	// 	// Store DeviceData.Points
	// 	_ = json.Unmarshal(data, &du)
	// 	for deviceName, deviceRef := range du {
	// 		fmt.Printf("%s =>\n", deviceName)
	//
	// 		for pointName, pointRef := range deviceRef {
	//
	// 			for time, value := range pointRef {
	// 				e.Points[pointName].List = append(e.Points[pointName].List)
	// 				fmt.Printf("%s, %s, %s, %s\n", deviceName, pointName, time, value)
	// 				// if k == "" {
	// 				// 	delete(du, i)
	// 				// 	continue
	// 				// }
	// 			}
	// 		}
	// 	}
	//
	// 	// var dp DecodePoints
	// 	// // Store DeviceData.Points
	// 	// _ = json.Unmarshal(data, &dp)
	// 	// for i, k := range dp {
	// 	// 	if k == nil {
	// 	// 		delete(dp, i)
	// 	// 		continue
	// 	// 	}
	// 	// 	i = strings.TrimSuffix(i, "List")
	// 	// 	u := du[i+"_unit"]
	// 	// 	e.Points[i] = Point{
	// 	// 		List: k,
	// 	// 		Unit: u,
	// 	// 	}
	// 	// }
	// }

	return ret
}

func (e *EndPoint) GetPointDataTable(points api.TemplatePoints) output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable(
			"Date/Time",
			"Point Id",
			"Point Name",
			"Value",
			"Units",
		)
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		// e.Error = table.SetHeader(
		// 	"Date/Time",
		// 	"Point Id",
		// 	"Point Name",
		// 	"Value",
		// 	"Units",
		// )
		if e.Error != nil {
			break
		}

		t := e.Request.RequestData.StartTimeStamp
		e.SetFilenamePrefix(t.String())
		table.SetFilePrefix(t.String())

		for deviceName, deviceRef := range e.Response.ResultData.Devices {
			for pointId, pointRef := range deviceRef.Points {
				for _, tim := range pointRef.Times {
					gp := points.GetPoint(deviceName, pointId)
					_ = table.AddRow(tim.Key.PrintFull(),
						fmt.Sprintf("%s.%s", deviceName, pointId),
						gp.Name,
						tim.Value,
						gp.Unit,
					)
					if table.Error != nil {
						continue
					}
				}
			}
		}

		table.InitGraph(output.GraphRequest {
			Title:        "",
			TimeColumn:   output.SetString("Date/Time"),
			SearchColumn: output.SetString("Point Id"),
			NameColumn:   output.SetString("Point Name"),
			ValueColumn:  output.SetString("Value"),
			UnitsColumn:  output.SetString("Units"),
			SearchString: output.SetString(""),
			MinLeftAxis:  output.SetFloat(0),
			MaxLeftAxis:  output.SetFloat(0),
		})

	}

	return table
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, e.Request.PsId.String(), GoStruct.NewEndPointPath(e.Request.PsId.String()))

		// entries.StructToDataMap(*e, "system", GoStruct.EndPointPath{})
	}

	return entries
}
