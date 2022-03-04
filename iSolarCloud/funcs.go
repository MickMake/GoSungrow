package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"strconv"
	"time"
)


func (sg *SunGrow) GetByJson(endpoint string, request string) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		if sg.Error != nil {
			break
		}
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		if request != "" {
			ret = ret.SetRequestByJson(output.Json(request))
			if ret.IsError() {
				fmt.Println(ret.Help())
				sg.Error = ret.GetError()
				break
			}
		}

		ret = ret.Call()
		if ret.IsError() {
			fmt.Println(ret.Help())
			sg.Error = ret.GetError()
			break
		}

		switch {
			case sg.OutputType.IsNone():

			case sg.OutputType.IsFile():
				sg.Error = ret.WriteDataFile()

			case sg.OutputType.IsRaw():
				fmt.Println(ret.GetJsonData(true))

			case sg.OutputType.IsJson():
				fmt.Println(ret.GetJsonData(false))

			default:
		}
	}
	return ret
}

func (sg *SunGrow) GetByStruct(endpoint string, request interface{}, cache time.Duration) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		if sg.Error != nil {
			break
		}
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		if request != nil {
			ret = ret.SetRequest(request)
			if ret.IsError() {
				sg.Error = ret.GetError()
				break
			}
		}

		ret = ret.SetCacheTimeout(cache)
		// if ret.CheckCache() {
		// 	ret = ret.ReadCache()
		// 	if !ret.IsError() {
		// 		break
		// 	}
		// }

		ret = ret.Call()
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		// sg.Error = ret.WriteCache()
		// if sg.Error != nil {
		// 	break
		// }
	}

	return ret
}

func (sg *SunGrow) GetPointData(date string, pointNames api.TemplatePoints) error {
	for range Only.Once {
		if len(pointNames) == 0 {
			sg.Error = errors.New("no points defined")
			break
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)

		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData {
				PsID:           psId,
				PsKey:          pointNames.PrintKeys(),
				Points:         pointNames.PrintPoints(),
				MinuteInterval: "5",
				StartTimeStamp: when.GetDayStartTimestamp(),
				EndTimeStamp:   when.GetDayEndTimestamp(),
			},
			DefaultCacheTimeout,
		)
		if sg.Error != nil {
			break
		}

		ep2 := queryMutiPointDataList.Assert(ep)
		table := ep2.GetDataTable(pointNames)
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		sg.Error = sg.Output(ep2, &table, "")
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) Output(endpoint api.EndPoint, table *output.Table, graphFilter string) error {
	for range Only.Once {
		switch {
			case sg.OutputType.IsNone():

			case sg.OutputType.IsHuman():
				if table == nil {
					break
				}
				table.Print()

			case sg.OutputType.IsFile():
				if table == nil {
					break
				}
				sg.Error = table.WriteCsvFile()

			case sg.OutputType.IsRaw():
				fmt.Println(endpoint.GetJsonData(true))

			case sg.OutputType.IsJson():
				fmt.Println(endpoint.GetJsonData(false))

			case sg.OutputType.IsGraph():
				if table == nil {
					break
				}
				sg.Error = table.SetGraphFromJson(output.Json(graphFilter))
				if sg.Error != nil {
					break
				}
				sg.Error = table.CreateGraph()
				if sg.Error != nil {
					break
				}

			default:
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetPsId() (int64, error) {
	var ret int64

	for range Only.Once {

		ep := sg.GetByStruct("AppService.getPsList", nil, DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		_getPsList := getPsList.AssertResultData(ep)
		ret = _getPsList.GetPsId()
	}

	return ret, sg.Error
}

func (sg *SunGrow) GetPsKeys() ([]string, error) {
	var ret []string

	for range Only.Once {
		var psId int64
		psId, sg.Error = sg.GetPsId()
		if sg.Error != nil {
			break
		}

		ep := sg.GetByStruct(
			"AppService.getPsDetailWithPsType",
			getPsDetailWithPsType.RequestData{PsId: strconv.FormatInt(psId, 10)},
			DefaultCacheTimeout)
		if ep.IsError() {
			sg.Error = ep.GetError()
			break
		}

		ep2 := getPsDetailWithPsType.Assert(ep)
		ret = ep2.GetPsKeys()
	}

	return ret, sg.Error
}

func fillArray(count int, args []string) []string {
	var ret []string
	for range Only.Once {
		if count < len(args) {
			count = len(args)
		}
		ret = make([]string, count)
		for i, e := range args {
			ret[i] = e
		}
	}
	return ret
}
