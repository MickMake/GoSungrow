package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"GoSungrow/iSolarCloud/AppService/getPsList"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	datatable "go.pennock.tech/tabular/auto"
	"math"
	"os"
	"strings"
	"time"
)


// GetDevicePointAttrs - WebAppService.getDevicePointAttrs Uuid: PsId: DeviceType
func (sg *SunGrow) GetDevicePointAttrs(psId valueTypes.PsId) ([]getDevicePointAttrs.Points, error) {
	var ret []getDevicePointAttrs.Points

	for range Only.Once {
		var tree PsTree
		tree, sg.Error = sg.PsTreeMenu(psId.String())
		if sg.Error != nil {
			break
		}

		for _, pid := range tree.Devices {
			ep := sg.GetByStruct(getDevicePointAttrs.EndPointName,
				getDevicePointAttrs.RequestData {
					Uuid:        pid.UUID,
					PsId2:       pid.PsId,
					DeviceType2: pid.DeviceType,
				},
				time.Hour * 24,
			)
			if sg.IsError() {
				break
			}

			data := getDevicePointAttrs.Assert(ep)

			ret = append(ret, data.Points()...)

			// more := sg.GetDevicePointNames(pid.DeviceType)
		}
	}

	return ret, sg.Error
}

// GetPowerDevicePointInfo - AppService.getPowerDevicePointInfo
func (sg *SunGrow) GetPowerDevicePointInfo(min valueTypes.Integer, max valueTypes.Integer) ([]getPowerDevicePointInfo.ResultData, error) {
	var ret []getPowerDevicePointInfo.ResultData
	for range Only.Once {
		_, _ = fmt.Fprintf(os.Stderr,"Scanning from %s to %s.", min, max)
		for id := min.Value(); id < max.Value(); id++ {
			PrintPause(id, 20)

			ep := sg.GetByStruct(getPowerDevicePointInfo.EndPointName,
				getPowerDevicePointInfo.RequestData {
					Id: valueTypes.SetIntegerValue(id),
				},
				time.Hour*24,
			)
			if sg.IsError() {
				break
			}

			data := getPowerDevicePointInfo.Assert(ep)
			ret = append(ret, data.Response.ResultData)
		}
		_, _ = fmt.Fprintf(os.Stderr,"\n")
	}

	return ret, sg.Error
}

// // QueryMultiPointDataList - AppService.queryMutiPointDataList MinuteInterval:5 StartTimeStamp:20221001000000 EndTimeStamp:20221001235900 PsId:1129147 Points:1129147_14_1_1.p13148,1129147_14_1_1.p13147,1129147_14_1_1.p13146,1129147_14_1_1.p13145,1129147_14_1_1.p13144,1129147_14_1_1.p13143
// func (sg *SunGrow) QueryMultiPointDataList(startDate valueTypes.DateTime, endDate valueTypes.DateTime, interval valueTypes.Integer, points valueTypes.PointIds) (queryMutiPointDataList.Data, error) {
// 	var ret queryMutiPointDataList.Data
// 	for range Only.Once {
// 		if len(points.PointIds) == 0 {
// 			break
// 		}
//
// 		var psId valueTypes.PsId
// 		psids := points.PsIds()
// 		if len(psids) != 1 {
// 			// Should only have one ps_id.
// 			break
// 		}
// 		psId := psids[0]
//
// 		if !psId.Valid {
// 			interval.SetValue(5)
// 		}
//
// 		if startDate.IsZero() {
// 			startDate.SetValue(time.Now())
// 			startDate.SetDayStart()
// 		}
//
// 		if endDate.IsZero() {
// 			endDate.SetValue(startDate.Value())
// 			endDate.SetDayEnd()
// 		}
//
// 		if !interval.Valid {
// 			interval.SetValue(5)
// 		}
//
// 		ep := sg.GetByStruct(
// 			"WebAppService.queryMutiPointDataList",
// 			queryMutiPointDataList.RequestData {
// 				PsId: psId,
// 				StartTimeStamp: startDate,
// 				EndTimeStamp: endDate,
// 				MinuteInterval: interval,
// 				PsKeys: (points.PsKeys().String()),
// 				Points: points,
// 			},
// 			time.Hour * 24,
// 		)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		data := queryMutiPointDataList.Assert(ep)
// 		ret = data.Response.ResultData.Data
//
// 		// tables := data.GetEndPointDataTables()
// 		// for _, table := range tables {
// 		// 	table.OutputType = sg.OutputType
// 		// 	table.Output()
// 		// }
// 		//
// 		// // ret = data.Response.ResultData
// 		// //
// 		// // table := output.NewTable(
// 		// // 	"Date/Time",
// 		// // 	"Point Id",
// 		// // 	"Point Name",
// 		// // 	"Value",
// 		// // 	"Units",
// 		// // )
// 		// // table.SetTitle("")
// 		// //
// 		// // t := data.Request.RequestData.StartTimeStamp
// 		// // data.SetFilenamePrefix(t.String())
// 		// // table.SetFilePrefix(t.String())
// 		// //
// 		// // for deviceName, deviceRef := range data.Response.ResultData.Data {
// 		// // 	for pointId, pointRef := range deviceRef.Points {
// 		// // 		for _, tim := range pointRef.Times {
// 		// // 			gp := points.GetPoint(deviceName, pointId)
// 		// // 			sg.Error = table.AddRow(tim.Key.PrintFull(),
// 		// // 				fmt.Sprintf("%s.%s", deviceName, pointId),
// 		// // 				gp.Name,
// 		// // 				tim.Value,
// 		// // 				gp.Unit,
// 		// // 			)
// 		// // 			if sg.Error != nil {
// 		// // 				continue
// 		// // 			}
// 		// // 		}
// 		// // 	}
// 		// // }
// 		// //
// 		// // table.InitGraph(output.GraphRequest {
// 		// // 	Title:        "",
// 		// // 	TimeColumn:   output.SetString("Date/Time"),
// 		// // 	SearchColumn: output.SetString("Point Id"),
// 		// // 	NameColumn:   output.SetString("Point Name"),
// 		// // 	ValueColumn:  output.SetString("Value"),
// 		// // 	UnitsColumn:  output.SetString("Units"),
// 		// // 	SearchString: output.SetString(""),
// 		// // 	MinLeftAxis:  output.SetFloat(0),
// 		// // 	MaxLeftAxis:  output.SetFloat(0),
// 		// // })
// 		// //
// 		// // table.SetTitle("Point Data %s", psId)
// 		// // table.SetFilePrefix(data.SetFilenamePrefix("%d", psId))
// 		// // table.SetGraphFilter("")
// 		// // table.SetSaveFile(sg.SaveAsFile)
// 		// // table.OutputType = sg.OutputType
// 		// // sg.Error = table.Output()
// 		// if sg.IsError() {
// 		// 	break
// 		// }
//
// 	}
//
// 	return ret, sg.Error
// }

func (sg *SunGrow) GetAllPointsData(psIds ...string) error {
	for range Only.Once {
		var points api.DataMap

		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		// getPsList
		ep := sg.GetByStruct(getPsList.EndPointName, getPsList.RequestData{}, DefaultCacheTimeout)
		if sg.IsError() {
			break
		}
		PsList := getPsList.Assert(ep)
		data := PsList.GetEndPointData()
		if PsList.Error != nil {
			sg.Error = PsList.Error
			break
		}
		points.AppendMap(data)

		for _, psId := range pids {
			// pid := strconv.Itoa(int(psId))

			// GoSungrow api raw getIncomeSettingInfos '{"ps_id":"1171348"}'
			// // getPowerStationForHousehold
			// // GoSungrow api raw getPowerStationForHousehold '{"ps_id":"1171348"}'

			// // getAreaList
			// getAreaList.Disabled
			//
			// // getPsReport
			// getPsReport.Disabled

			// // getPowerStatistics
			// // api raw getPowerStatistics '{"ps_id":"1171348"}'
			// ep = sg.GetByStruct(getPowerStatistics.EndPointName, getPowerStatistics.RequestData{ PsId: pid }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// PowerStatistics := getPowerStatistics.Assert(ep)
			// data = PowerStatistics.GetEndPointData()
			// if PowerStatistics.Error != nil {
			// 	sg.Error = PowerStatistics.Error
			// 	break
			// }
			// points.AppendMap(data)

			// getPowerStationData
			// api raw getPowerStationData '{"date_id":"20221007","date_type":"1","ps_id":"1171348"}'
			ep = sg.GetByStruct(getPowerStationData.EndPointName,
				getPowerStationData.RequestData{ PsId: psId, DateType: valueTypes.SetStringValue("1"), DateId: valueTypes.SetDateTimeString("20221007")},
				DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			PowerStationData := getPowerStationData.Assert(ep)
			data = PowerStationData.GetEndPointData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)

			// api raw getPowerStationData '{"date_id":"202210","date_type":"2","ps_id":"1171348"}'
			ep = sg.GetByStruct(getPowerStationData.EndPointName,
				getPowerStationData.RequestData{ PsId: psId, DateType: valueTypes.SetStringValue("2"), DateId: valueTypes.SetDateTimeString("202210")},
				DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			PowerStationData = getPowerStationData.Assert(ep)
			data = PowerStationData.GetEndPointData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)

			// api raw getPowerStationData '{"date_id":"2022","date_type":"3","ps_id":"1171348"}'
			ep = sg.GetByStruct(getPowerStationData.EndPointName,
				getPowerStationData.RequestData{ PsId: psId, DateType: valueTypes.SetStringValue("3"), DateId: valueTypes.SetDateTimeString("2022")},
				DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			PowerStationData = getPowerStationData.Assert(ep)
			data = PowerStationData.GetEndPointData()
			if PowerStationData.Error != nil {
				sg.Error = PowerStationData.Error
				break
			}
			points.AppendMap(data)


			// queryDeviceList
			// api get AppService.queryDeviceList '{"ps_id":"1171348"}'
			ep = sg.GetByStruct(queryDeviceList.EndPointName, queryDeviceList.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			DeviceList := queryDeviceList.Assert(ep)
			data = DeviceList.GetEndPointData()
			if DeviceList.Error != nil {
				sg.Error = DeviceList.Error
				break
			}
			points.AppendMap(data)


			// getHouseholdStoragePsReport
			// api get getHouseholdStoragePsReport '{"date_id":"20221001","date_type":"1","ps_id":"1129147"}'
			ep = sg.GetByStruct(getHouseholdStoragePsReport.EndPointName,
				getHouseholdStoragePsReport.RequestData{ DateId: valueTypes.SetDateTimeString("20221001"), DateType: valueTypes.SetStringValue("1"), PsId: psId },
				DefaultCacheTimeout,
			)
			if sg.IsError() {
				break
			}
			HouseholdStoragePsReport := getHouseholdStoragePsReport.Assert(ep)
			data = HouseholdStoragePsReport.GetEndPointData()
			if HouseholdStoragePsReport.Error != nil {
				sg.Error = HouseholdStoragePsReport.Error
				break
			}
			points.AppendMap(data)


			// getPsDetailWithPsType
			ep = sg.GetByStruct(getPsDetailWithPsType.EndPointName, getPsDetailWithPsType.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			PsDetailWithPsType := getPsDetailWithPsType.Assert(ep)
			data = PsDetailWithPsType.GetEndPointData()
			if PsDetailWithPsType.Error != nil {
				sg.Error = PsDetailWithPsType.Error
				break
			}
			points.AppendMap(data)


			// getPsDetail
			ep = sg.GetByStruct(getPsDetail.EndPointName, getPsDetail.RequestData{ PsId: psId }, DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			PsDetail := getPsDetail.Assert(ep)
			data = PsDetail.GetEndPointData()
			if PsDetail.Error != nil {
				sg.Error = PsDetailWithPsType.Error
				break
			}
			points.AppendMap(data)


			// getKpiInfo
			ep = sg.GetByStruct(getKpiInfo.EndPointName, getKpiInfo.RequestData{ }, DefaultCacheTimeout)
			if sg.IsError() {
				break
			}
			KpiInfo := getKpiInfo.Assert(ep)
			data = KpiInfo.GetEndPointData()
			if KpiInfo.Error != nil {
				sg.Error = KpiInfo.Error
				break
			}
			points.AppendMap(data)


			// // getPowerDevicePointInfo
			// ep = sg.GetByStruct(getPowerDevicePointInfo.EndPointName, getPowerDevicePointInfo.RequestData{ Id: psId }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// ep4 := getPowerDevicePointInfo.Assert(ep)
			// data = ep4.GetEndPointData()
			// if ep4.Error != nil {
			// 	sg.Error = ep4.Error
			// 	break
			// }
			// points.AppendMap(data)


			// // queryDeviceRealTimeDataByPsKeys
			// ep = sg.GetByStruct(queryDeviceRealTimeDataByPsKeys.EndPointName, queryDeviceRealTimeDataByPsKeys.RequestData{ PsKeyList: psId }, DefaultCacheTimeout)
			// if sg.Error != nil {
			// 	break
			// }
			// DeviceRealTimeDataByPsKeys := queryDeviceRealTimeDataByPsKeys.Assert(ep)
			// data = DeviceRealTimeDataByPsKeys.GetEndPointData()
			// if DeviceRealTimeDataByPsKeys.Error != nil {
			// 	sg.Error = DeviceRealTimeDataByPsKeys.Error
			// 	break
			// }
			// points.AppendMap(data)

		}

		points.Print()
	}

	return sg.Error
}

// PointNames - Return all points associated with psIds.
func (sg *SunGrow) PointNames(psIds ...string) string {
	var ret string

	for range Only.Once {
		var pids valueTypes.PsIds
		pids = sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, pid := range pids {
			var points []getDevicePointAttrs.Points
			points, sg.Error = sg.GetDevicePointAttrs(pid)
			if sg.Error != nil {
				break
			}

			if len(points) == 0 {
				continue
			}

			ret += fmt.Sprintf("# Available points for ps_id %s:\n", pid.String())
			table := datatable.New("utf8-heavy")
			table.AddHeaders("Id", "Name", "Unit", "UnitType", "PsId", "DeviceType", "DeviceName")
			for _, point := range points {
				table.AddRowItems(point.Id, point.Name, point.Unit, point.UnitType, point.PsId, point.DeviceType, point.DeviceName)
			}

			var r string
			r, sg.Error = table.Render()
			if sg.Error != nil {
				break
			}
			ret += r

			// @TODO - Include AppService.getPowerDevicePointNames
			// points2 := cmds.Api.SunGrow.GetDevicePointNames(pid)
			// if c.Error != nil {
			// 	break
			// }
			// if len(points) == 0 {
			// 	continue
			// }

		}
	}

	return ret
}

func (sg *SunGrow) PointData(startDate string, endDate string, interval string, points ...string) string {
	var ret string

	for range Only.Once {
		data := sg.NewSunGrowData()
		data.SetEndpoints(queryMutiPointDataList.EndPointName)
		// req := iSolarCloud.RequestArgs{
		// 	StartTimeStamp:           startDate,
		// 	EndTimeStamp:   endDate,
		// }
		// var req iSolarCloud.RequestArgs
		// data.Request.SetPoints(points)

		startDate = valueTypes.NewDateTime(startDate).Format(valueTypes.DateTimeLayoutSecond)
		endDate = valueTypes.NewDateTime(endDate).Format(valueTypes.DateTimeLayoutSecond)

		data.SetArgs(
			"StartTimeStamp:" + startDate,
			"EndTimeStamp:" + endDate,
			"MinuteInterval:" + interval,
			"Points:" + strings.Join(points, ","),
		)

		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.Process()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}

		// var argStartDate valueTypes.DateTime
		// switch {
		// 	case startDate == valueTypes.Now:
		// 		fallthrough
		// 	case startDate == "":
		// 		fallthrough
		// 	case startDate == ".":
		// 		argStartDate = valueTypes.NewDateTime(valueTypes.Now)
		// 		argStartDate.SetDayStart()
		// 	default:
		// 		argStartDate = valueTypes.NewDateTime(startDate)
		// }
		// if argStartDate.Error != nil {
		// 	c.Error = argStartDate.Error
		// 	break
		// }
		//
		// var argEndDate valueTypes.DateTime
		// switch {
		// 	case endDate == valueTypes.Now:
		// 		fallthrough
		// 	case endDate == "":
		// 		fallthrough
		// 	case endDate == ".":
		// 		argEndDate = valueTypes.NewDateTime(valueTypes.Now)
		// 		argEndDate.SetDayEnd()
		// 	default:
		// 		argEndDate = valueTypes.NewDateTime(endDate)
		// }
		// if argEndDate.Error != nil {
		// 	c.Error = argEndDate.Error
		// 	break
		// }
		//
		// var argInterval valueTypes.Integer
		// switch {
		// 	case endDate == "":
		// 		fallthrough
		// 	case endDate == ".":
		// 		argInterval = valueTypes.SetIntegerString("5")
		// 	default:
		// 		argInterval = valueTypes.SetIntegerString(interval)
		// }
		// if argInterval.Error != nil {
		// 	c.Error = argInterval.Error
		// 	break
		// }
		//
		// argPsId := valueTypes.SetPsIdString(psId)
		// if argPsId.Error != nil {
		// 	c.Error = argPsId.Error
		// 	break
		// }
		//
		// var argPsKey valueTypes.PsKey
		// argPsKey = valueTypes.SetPsKeyString(psKey)
		// if argPsKey.Error != nil {
		// 	c.Error = argPsKey.Error
		// 	break
		// }
		//
		// var argPoints valueTypes.PointIds
		// argPoints.Set(points...)
		// if argPoints.Error != nil {
		// 	c.Error = argPoints.Error
		// 	break
		// }
		// var data queryMutiPointDataList.Data
		// data, c.Error = cmds.Api.SunGrow.QueryMultiPointDataList(argStartDate, argEndDate, argInterval, argPsId, argPsKey, argPoints)
		// if c.Error != nil {
		// 	break
		// }
		//
		// table := output.NewTable(
		// 	"Timestamp",
		// 	"Ps Key",
		// 	"Key",
		// 	"Value",
		// )
		//
		// for _, d := range data {
		// 	for k, v := range d.Points {
		// 		c.Error = table.AddRow(
		// 			d.Timestamp.String(),
		// 			d.PsKey.String(),
		// 			k,
		// 			v.String(),
		// 		)
		// 		if c.Error != nil {
		// 			break
		// 		}
		// 	}
		// }
		// if c.Error != nil {
		// 	break
		// }
		//
		// ret = table.String()
	}

	return ret
}

func (sg *SunGrow) PointScan(min string, max string) string {
	var ret string

	for range Only.Once {
		if min == "" || max == "" {
			min = "0"
			max = "1000"
		}

		minI := valueTypes.SetIntegerString(min)
		if minI.Error != nil {
			sg.Error = minI.Error
			break
		}

		maxI := valueTypes.SetIntegerString(max)
		if maxI.Error != nil {
			sg.Error = maxI.Error
			break
		}

		var data []getPowerDevicePointInfo.ResultData
		data, sg.Error = sg.GetPowerDevicePointInfo(minI, maxI)
		if sg.Error != nil {
			break
		}

		table := output.NewTable(
			"DeviceType",
			"Id",
			"Period",
			"Point Id",
			"Point Name",
			"Show Point Name",
			"Translation Id",
		)

		for _, d := range data {
			sg.Error = table.AddRow(d.DeviceType.String(),
				d.Id.String(),
				d.Period.String(),
				d.PointId.String(),
				d.PointName.String(),
				d.ShowPointName.String(),
				d.TranslationId.String(),
			)
			if sg.Error != nil {
				break
			}
		}
		if sg.Error != nil {
			break
		}

		ret = table.String()
	}

	return ret
}


func PrintPause(index int64, max int) {
	for range Only.Once {
		if index == 0 {
			_, _ = fmt.Fprintf(os.Stderr,"\n%.3d ", index)
			break
		}

		m := math.Mod(float64(index), float64(max))
		if m == 0 {
			_, _ = fmt.Fprintf(os.Stderr,"PAUSE")
			time.Sleep(time.Millisecond * 500)
			// fmt.Printf("\r%s%.3d ", strings.Repeat(" ", 4), pni)
			_, _ = fmt.Fprintf(os.Stderr,"\r%.3d ", index)
		} else {
			time.Sleep(time.Millisecond * 100)
			_, _ = fmt.Fprintf(os.Stderr,".")
		}
	}
}
