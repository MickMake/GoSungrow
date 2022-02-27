package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"time"
)


func (sg *SunGrow) GetPointNamesFromTemplate(template string) api.TemplatePoints {
	var ret api.TemplatePoints

	for range Only.Once {
		if template == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		ep := sg.GetByStruct(
			"WebAppService.queryUserCurveTemplateData",
			queryUserCurveTemplateData.RequestData{TemplateID: template},
			time.Hour,
		)
		if sg.Error != nil {
			break
		}

		data := queryUserCurveTemplateData.AssertResultData(ep)
		for dn, dr := range data.PointsData.Devices {
			for _, pr := range dr.Points {
					ret = append(ret, api.TemplatePoint {
					PsKey:       dn,
					PointId:     "p" + pr.PointID,
					Description: pr.PointName,
					Unit:        pr.Unit,
				})
			}
		}
	}

	return ret
}

func (sg *SunGrow) GetTemplateData(template string, date string, filter string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)

		psId := sg.GetPsId()
		if sg.Error != nil {
			break
		}

		pointNames := sg.GetPointNamesFromTemplate(template)
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

		// data := queryMutiPointDataList.AssertResultData(ep)
		data := queryMutiPointDataList.Assert(ep)
		table := data.GetDataTable(pointNames)
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		fn := data.SetFilenamePrefix("%s-%s", when.String(), template)
		sg.Error = table.SetFilePrefix(fn)
		if sg.Error != nil {
			break
		}

		sg.Error = sg.Output(ep, table, filter)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetTemplatePoints(template string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		table := output.NewTable()
		sg.Error = table.SetHeader(
			"PointId",
			"Description",
			"Unit",
			)
		if sg.Error != nil {
			break
		}

		ss := sg.GetPointNamesFromTemplate(template)
		for _, s := range ss {
			sg.Error = table.AddRow(
				s.PointId,
				s.Description,
				s.Unit,
			)
			if sg.Error != nil {
				break
			}
		}
		if sg.Error != nil {
			break
		}

		table.Print()
	}

	return sg.Error
}
