package queryMutiPointDataList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	"sort"
)

const Url = "/v1/commonService/queryMutiPointDataList"
const Disabled = false

type RequestData struct {
	PsID           int64  `json:"ps_id" required:"true"`
	PsKey          string `json:"ps_key" required:"true"`
	Points         string `json:"points" required:"true"`
	MinuteInterval string `json:"minute_interval" required:"true"`
	StartTimeStamp string `json:"start_time_stamp" required:"true"`
	EndTimeStamp   string `json:"end_time_stamp" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
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
type Points map[string]Point
type Point struct {
	Name  string `json:"name"`
	Units string `json:"units"`
	Times Times  `json:"times"`
}
type Times []Time
type Time struct {
	Key   api.DateTime `json:"key"`
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

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &d)

		e.Devices = make(Devices)
		for deviceName, deviceRef := range d {

			points := Points{}
			for pointName, pointRef := range deviceRef {

				times := Times{}
				for time, value := range pointRef {

					times = append(times, Time{
						Key:   api.NewDateTime(time),
						Value: value,
					})
				}

				sort.Slice(times, func(i, j int) bool {
					return times[i].Key.Before(times[j].Key.Time)
				})
				points[pointName] = Point{
					Name:  "",
					Units: "",
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
