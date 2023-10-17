package iSolarCloud

import (
	"errors"
	"strings"
	"time"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getTemplateList"
	"github.com/anicoll/gosungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

// TemplateList - Return all templates.
func (sg *SunGrow) TemplateList() error {
	for range Only.Once {
		data := sg.NewSunGrowData()
		data.SetEndpoints(getTemplateList.EndPointName)
		data.SetArgs()
		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}

		// var data getTemplateList.ResultData
		// data, sg.Error = sg.GetTemplateList()
		//
		// table := data.GetEndPointResultTable()
		// if table.Error != nil {
		// 	sg.Error = table.Error
		// 	break
		// }
		//
		// table.SetTitle("getTemplateList")
		// table.SetFilePrefix(data.SetFilenamePrefix(""))
		// table.SetGraphFilter("")
		// table.SetSaveFile(sg.SaveAsFile)
		// table.OutputType = sg.OutputType
		// sg.Error = table.Output()
		// if sg.IsError() {
		// 	break
		// }
	}

	return sg.Error
}

// TemplatePoints - Return all points associated with a template_id.
func (sg *SunGrow) TemplatePoints(template string) error {
	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}
		// ep := sg.GetByStruct(
		// 	"WebAppService.queryUserCurveTemplateData",
		// 	queryUserCurveTemplateData.RequestData{TemplateId: valueTypes.SetStringValue(template)},
		// 	time.Hour,
		// )
		// if sg.IsError() {
		// 	break
		// }
		// data := queryUserCurveTemplateData.Assert(ep)

		data := sg.NewSunGrowData()
		data.SetEndpoints(queryUserCurveTemplateData.EndPointName)
		data.SetArgs(
			"TemplateId:" + template,
		)
		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}

		// result := sg.QueryUserCurveTemplateData(template)
		// if sg.IsError() {
		// 	break
		// }
		// table := output.NewTable(
		// 	"PointStruct Id",
		// 	"Description",
		// 	"Unit",
		// )
		// for device := range result.PointsData.Devices {
		// 	for point := range result.PointsData.Devices[device].Points {
		// 		p := result.PointsData.Devices[device].Points[point]
		// 		sg.Error = table.AddRow(
		// 			s.PsKey+"."+s.PointId.String(),
		// 			s.Name,
		// 			s.Unit,
		// 		)
		// 		if sg.IsError() {
		// 			break
		// 		}
		// 	}
		// }
		// if sg.IsError() {
		// 	break
		// }
		//
		// table.SetTitle("Template %s", template)
		// table.SetFilePrefix(template)
		// table.SetGraphFilter("")
		// table.SetSaveFile(sg.SaveAsFile)
		// table.OutputType = sg.OutputType
		// sg.Error = table.Output()
		// if sg.IsError() {
		// 	break
		// }
	}

	return sg.Error
}

// TemplateData - Return all point data associated with a template_id.
func (sg *SunGrow) TemplateData(template string, startDate string, endDate string, interval string) error {
	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		var data queryUserCurveTemplateData.ResultData
		data, sg.Error = sg.QueryUserCurveTemplateData(template)
		if sg.IsError() {
			break
		}

		var points []string
		// for an := range data.PointsData.Devices {
		// 	for _, b := range data.PointsData.Devices[an].Points {
		// 		points = append(points, b.PointId.Full())
		// 	}
		// }
		// Alternative - this maintains the original order defined in the template.
		ps := strings.ReplaceAll(data.PointsData.Order.String(), "&", ".p")
		points = strings.Split(ps, ",")

		// data2 := sg.NewSunGrowData()
		// data2.SetEndpoints(queryMutiPointDataList.EndPointName)
		// // req := iSolarCloud.RequestArgs{
		// // 	StartTimeStamp:           startDate,
		// // 	EndTimeStamp:   endDate,
		// // }
		// // var req iSolarCloud.RequestArgs
		// // data.Request.SetPoints(points)
		//
		// startDate = valueTypes.NewDateTime(startDate).Format(valueTypes.DateTimeLayoutSecond)
		// endDate = valueTypes.NewDateTime(endDate).Format(valueTypes.DateTimeLayoutSecond)
		//
		// data2.SetArgs(
		// 	"StartTimeStamp:" + startDate,
		// 	"EndTimeStamp:" + endDate,
		// 	"MinuteInterval:" + interval,
		// 	"Points:" + strings.Join(points, ","),
		// )
		//
		// sg.Error = data2.GetData()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// sg.Error = data2.Process()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// // @TODO - Figure out a way to push the Unit values from QueryUserCurveTemplateData to this table.
		// // result := queryMutiPointDataList.Assert(data2.Results["queryMutiPointDataList/1129147"].EndPoint)
		// // for nr, r := range result.Response.ResultData.Data {
		// // 	for nr2, r2 := range r.Points {
		// // 		result.Response.ResultData.Data[nr].Points[r2].
		// // 	}
		// // }
		// sg.Error = data2.OutputDataTables()
		// if sg.Error != nil {
		// 	break
		// }

		// @TODO - Figure out a way to push the Unit values from QueryUserCurveTemplateData to this table.
		// @TODO - Maybe use a point cache?!

		sg.Error = sg.PointData(startDate, endDate, interval, points...)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

// TemplateDataSave - Return all point data associated with a template_id and save to files.
func (sg *SunGrow) TemplateDataSave(template string, startDate string, endDate string, interval string) error {
	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		var data queryUserCurveTemplateData.ResultData
		data, sg.Error = sg.QueryUserCurveTemplateData(template)
		if sg.IsError() {
			break
		}

		var points []string
		// for an := range data.PointsData.Devices {
		// 	for _, b := range data.PointsData.Devices[an].Points {
		// 		points = append(points, b.PointId.Full())
		// 	}
		// }
		// Alternative - this maintains the original order defined in the template.
		ps := strings.ReplaceAll(data.PointsData.Order.String(), "&", ".p")
		points = strings.Split(ps, ",")

		// data2 := sg.NewSunGrowData()
		// data2.SetEndpoints(queryMutiPointDataList.EndPointName)
		// // req := iSolarCloud.RequestArgs{
		// // 	StartTimeStamp:           startDate,
		// // 	EndTimeStamp:   endDate,
		// // }
		// // var req iSolarCloud.RequestArgs
		// // data.Request.SetPoints(points)
		//
		// startDate = valueTypes.NewDateTime(startDate).Format(valueTypes.DateTimeLayoutSecond)
		// endDate = valueTypes.NewDateTime(endDate).Format(valueTypes.DateTimeLayoutSecond)
		//
		// data2.SetArgs(
		// 	"StartTimeStamp:" + startDate,
		// 	"EndTimeStamp:" + endDate,
		// 	"MinuteInterval:" + interval,
		// 	"Points:" + strings.Join(points, ","),
		// )
		//
		// sg.Error = data2.GetData()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// sg.Error = data2.Process()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// // @TODO - Figure out a way to push the Unit values from QueryUserCurveTemplateData to this table.
		// // result := queryMutiPointDataList.Assert(data2.Results["queryMutiPointDataList/1129147"].EndPoint)
		// // for nr, r := range result.Response.ResultData.Data {
		// // 	for nr2, r2 := range r.Points {
		// // 		result.Response.ResultData.Data[nr].Points[r2].
		// // 	}
		// // }
		// sg.Error = data2.OutputDataTables()
		// if sg.Error != nil {
		// 	break
		// }

		// @TODO - Figure out a way to push the Unit values from QueryUserCurveTemplateData to this table.
		// @TODO - Maybe use a point cache?!

		sg.Error = sg.PointDataSave(startDate, endDate, interval, points...)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

// GetTemplateList - AppService.getTemplateList
func (sg *SunGrow) GetTemplateList() (getTemplateList.ResultData, error) {
	var ret getTemplateList.ResultData
	for range Only.Once {
		ep := sg.GetByStruct(getTemplateList.EndPointName,
			getTemplateList.RequestData{},
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}

		data := getTemplateList.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}

// QueryUserCurveTemplateData - WebAppService.queryUserCurveTemplateData
func (sg *SunGrow) QueryUserCurveTemplateData(template string) (queryUserCurveTemplateData.ResultData, error) {
	var ret queryUserCurveTemplateData.ResultData

	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		ep := sg.GetByStruct(queryUserCurveTemplateData.EndPointName,
			queryUserCurveTemplateData.RequestData{TemplateId: valueTypes.SetStringValue(template)},
			time.Hour,
		)
		if sg.IsError() {
			break
		}

		data := queryUserCurveTemplateData.Assert(ep)
		ret = data.Response.ResultData
		// foo := data.GetData()
		// fmt.Printf("%v\n", foo)
		// tables := foo.CreateDataTables()
		//
		// for _, table := range tables {
		// 	table.SetOutputType(output.StringTypeTable)
		// 	table.Output()
		// }

		// for an, a := range ret.PointsData.Devices {
		// 	fmt.Println(an)
		// 	for bn, b := range ret.PointsData.Devices[an].Points {
		// 		fmt.Println(bn)
		// 		fmt.Print("%v\n", b)
		// 	}
		// }

		// data := queryUserCurveTemplateData.AssertResultData(ep)
		//
		// for dn, dr := range data.PointsData.Devices {
		// 	for _, pr := range dr.Points {
		// 		if pr.Unit.String() == "null" {
		// 			pr.Unit.SetString("")
		// 		}
		// 		ret = append(ret, api.TemplatePoint {
		// 			PsKey:   dn,
		// 			PointId: pr.PointId,
		// 			Name:    pr.PointName.String(),
		// 			Unit:    pr.Unit.String(),
		// 		})
		// 	}
		// }
	}

	return ret, sg.Error
}
