package iSolarCloud

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsDetail"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/getPsList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"github.com/MickMake/GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"github.com/MickMake/GoSungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

// GetPowerDevicePointInfo - AppService.getPowerDevicePointInfo
func (sg *SunGrow) GetPowerDevicePointInfo(min valueTypes.Integer, max valueTypes.Integer) ([]getPowerDevicePointInfo.ResultData, error) {
	var ret []getPowerDevicePointInfo.ResultData
	for range Only.Once {
		_, _ = fmt.Fprintf(os.Stderr, "Scanning from %s to %s.", min, max)
		for id := min.Value(); id < max.Value(); id++ {
			PrintPause(id, 20)

			ep := sg.GetByStruct(getPowerDevicePointInfo.EndPointName,
				getPowerDevicePointInfo.RequestData{
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
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	}

	return ret, sg.Error
}

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
				getPowerStationData.RequestData{PsId: psId, DateType: valueTypes.SetStringValue("1"), DateId: valueTypes.SetDateTimeString("20221007")},
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
				getPowerStationData.RequestData{PsId: psId, DateType: valueTypes.SetStringValue("2"), DateId: valueTypes.SetDateTimeString("202210")},
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
				getPowerStationData.RequestData{PsId: psId, DateType: valueTypes.SetStringValue("3"), DateId: valueTypes.SetDateTimeString("2022")},
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
			ep = sg.GetByStruct(queryDeviceList.EndPointName, queryDeviceList.RequestData{PsId: psId}, DefaultCacheTimeout)
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
				getHouseholdStoragePsReport.RequestData{DateId: valueTypes.SetDateTimeString("20221001"), DateType: valueTypes.SetStringValue("1"), PsId: psId},
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
			ep = sg.GetByStruct(getPsDetailWithPsType.EndPointName, getPsDetailWithPsType.RequestData{PsId: psId}, DefaultCacheTimeout)
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
			ep = sg.GetByStruct(getPsDetail.EndPointName, getPsDetail.RequestData{PsId: psId}, DefaultCacheTimeout)
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
			ep = sg.GetByStruct(getKpiInfo.EndPointName, getKpiInfo.RequestData{}, DefaultCacheTimeout)
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

		fmt.Println(points.String())
	}

	return sg.Error
}

// DevicePointAttrs - Return all points associated with psIds and device_type filter.
func (sg *SunGrow) DevicePointAttrs(psIds []string, deviceType string) (getDevicePointAttrs.Points, error) {
	var points getDevicePointAttrs.Points

	for range Only.Once {
		var pids valueTypes.PsIds
		pids = sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, pid := range pids {
			var p getDevicePointAttrs.Points
			p, sg.Error = sg.GetDevicePointAttrs(pid)
			if sg.Error != nil {
				break
			}

			for _, point := range p {
				if deviceType == "" {
					points = append(points, point)
					continue
				}
				if point.DeviceType.MatchString(deviceType) {
					points = append(points, point)
					continue
				}
			}

			// @TODO - Include AppService.getPowerDevicePointNames
			// points2 := cmds.Api.SunGrow.GetDevicePointNames(pid)
			// if c.Error != nil {
			// 	break
			// }
			// if len(points) == 0 {
			// 	continue
			// }
		}
		if sg.Error != nil {
			break
		}
	}

	return points, sg.Error
}

// DevicePointAttrsMap - Return all points associated with psIds and device_type filter.
func (sg *SunGrow) DevicePointAttrsMap(psIds []string, deviceType string) (getDevicePointAttrs.PointsMap, error) {
	points := make(getDevicePointAttrs.PointsMap)

	for range Only.Once {
		var pa getDevicePointAttrs.Points
		pa, sg.Error = sg.DevicePointAttrs(psIds, deviceType)
		if sg.Error != nil {
			break
		}

		for index := range pa {
			points[pa[index].Id.String()] = &pa[index]
		}
	}

	return points, sg.Error
}

// GetDevicePointAttrs - WebAppService.getDevicePointAttrs Uuid: PsId: DeviceType
func (sg *SunGrow) GetDevicePointAttrs(psId valueTypes.PsId) (getDevicePointAttrs.Points, error) {
	var ret getDevicePointAttrs.Points

	for range Only.Once {
		var tree PsTree
		tree, sg.Error = sg.PsTreeMenu(psId.String())
		if sg.Error != nil {
			break
		}

		for _, pid := range tree.Devices {
			ep := sg.GetByStruct(getDevicePointAttrs.EndPointName,
				getDevicePointAttrs.RequestData{
					Uuid:        pid.UUID,
					PsId2:       pid.PsId,
					DeviceType2: pid.DeviceType,
				},
				time.Hour*24,
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

func (sg *SunGrow) PointData(startDate string, endDate string, interval string, points ...string) error {
	for range Only.Once {
		// _, _ = sg.QueryMultiPointDataList(
		// 	valueTypes.SetDateTimeString(startDate),
		// 	valueTypes.SetDateTimeString(endDate),
		// 	valueTypes.SetIntegerString(interval),
		// 	valueTypes.SetPointIdsString(points...),
		// )

		data := sg.NewSunGrowData()
		// req := iSolarCloud.RequestArgs {
		// 	StartTimeStamp:           startDate,
		// 	EndTimeStamp:   endDate,
		// }
		// var req iSolarCloud.RequestArgs
		// data.Request.SetPoints(points)

		sd := valueTypes.NewDateTime(startDate)
		var ed valueTypes.DateTime
		if endDate == "" {
			ed = sd
			ed.SetDayEnd()
		} else {
			ed = valueTypes.NewDateTime(endDate)
		}

		// _, _ = fmt.Fprintf(os.Stderr,"Points: %s\n", strings.Join(points, " "))
		data.SetArgs(
			"StartTimeStamp:"+sd.Format(valueTypes.DateTimeLayoutSecond),
			"EndTimeStamp:"+ed.Format(valueTypes.DateTimeLayoutSecond),
			"MinuteInterval:"+interval,
			"Points:"+strings.Join(points, ","),
		)
		data.SetEndpoints(queryMutiPointDataList.EndPointName)

		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) PointScan(min string, max string) (string, error) {
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

	return ret, sg.Error
}

func PrintPause(index int64, max int) {
	for range Only.Once {
		if index == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "\n%.3d ", index)
			break
		}

		m := math.Mod(float64(index), float64(max))
		if m == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "PAUSE")
			time.Sleep(time.Millisecond * 500)
			// fmt.Printf("\r%s%.3d ", strings.Repeat(" ", 4), pni)
			_, _ = fmt.Fprintf(os.Stderr, "\r%.3d ", index)
		} else {
			time.Sleep(time.Millisecond * 100)
			_, _ = fmt.Fprintf(os.Stderr, ".")
		}
	}
}

// QueryMultiPointDataList - AppService.queryMutiPointDataList MinuteInterval:5 StartTimeStamp:20221001000000 EndTimeStamp:20221001235900 PsId:1129147 Points:1129147_14_1_1.p13148,1129147_14_1_1.p13147,1129147_14_1_1.p13146,1129147_14_1_1.p13145,1129147_14_1_1.p13144,1129147_14_1_1.p13143
func (sg *SunGrow) QueryMultiPointDataList(startDate valueTypes.DateTime, endDate valueTypes.DateTime, interval valueTypes.Integer, points valueTypes.PointIds) (queryMutiPointDataList.Data, error) {
	var ret queryMutiPointDataList.Data
	for range Only.Once {
		if len(points.PointIds) == 0 {
			break
		}

		psids := points.PsIds()
		if len(psids) == 0 {
			break
		}
		psId := psids[0]

		if !interval.Valid {
			interval.SetValue(5)
		}

		if startDate.IsZero() {
			startDate.SetValue(time.Now())
			startDate.SetDayStart()
		}
		startDate.SetDateType(valueTypes.DateTimeLayoutSecond)

		if endDate.IsZero() {
			endDate.SetValue(startDate.Value())
			endDate.SetDayEnd()
		}
		endDate.SetDateType(valueTypes.DateTimeLayoutSecond)

		if !interval.Valid {
			interval.SetValue(5)
		}

		ep := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData{
				PsId:           psId,
				StartTimeStamp: startDate,
				EndTimeStamp:   endDate,
				MinuteInterval: interval,
				PsKeys:         *points.PsKeys(),
				Points:         points,
			},
			time.Hour*24,
		)
		if sg.IsError() {
			break
		}

		var response SunGrowDataResponse
		response.Data = ep.GetEndPointData()

		response.Data.ProcessMap()
		if response.Data.Error != nil {
			sg.Error = response.Data.Error
			break
		}

		response.Options = OutputOptions{
			OutputType:   sg.OutputType,
			SaveAsFile:   sg.SaveAsFile,
			GraphRequest: output.GraphRequest{},
		}

		// // Outputs that can drop through to DataTables.
		// if sg.OutputType.IsTable() || sg.OutputType.IsXLSX() || sg.OutputType.IsCsv() || sg.OutputType.IsXML() {
		// 	table := response.Data.CreateResultTable(false)
		// 	table.OutputType = sg.OutputType
		// 	table.SetSaveFile(sg.SaveAsFile)
		// 	// table.AppendTitle(" - %s", sg.TitleSuffix)
		// 	// table.AppendFilePrefix(sg.FileSuffix)
		// 	sg.Error = table.Output()
		// 	if sg.Error != nil {
		// 		break
		// 	}
		// }

		sg.Error = response.OutputDataTables()
		if sg.IsError() {
			break
		}

	}

	return ret, sg.Error
}
