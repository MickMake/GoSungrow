package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"GoSungrow/iSolarCloud/WebAppService/queryUserCurveTemplateData"
	"GoSungrow/iSolarCloud/api"
	"errors"
	"fmt"
	"strings"
	"time"
)

type TemplatePoint struct {
	Description string
	PsKey       string
	PointId     string
	Unit        string
}
type TemplatePoints map[string]TemplatePoint

// type TemplateDevices map[string]TemplatePoints

// func (t *TemplateDevices) PrintKeys() string {
// 	var ret string
// 	for _, p := range *t {
// 		ret += fmt.Sprintf("%s,", p.PrintKeys())
// 	}
// 	ret = strings.TrimSuffix(ret, ",")
// 	return ret
// }
//
// func (t *TemplateDevices) PrintPoints() string {
// 	var ret string
// 	for p := range *t {
// 		ret += fmt.Sprintf("%s,", p)
// 	}
// 	ret = strings.TrimSuffix(ret, ",")
// 	return ret
// }

func (t *TemplatePoints) PrintKeys() string {
	var ret string
	for _, p := range *t {
		ret += fmt.Sprintf("%s,", p.PsKey)
	}
	ret = strings.TrimSuffix(ret, ",")
	return ret
}

func (t *TemplatePoints) PrintPoints() string {
	var ret string
	for _, p := range *t {
		ret += fmt.Sprintf("p%s,", p.PointId)
	}
	ret = strings.TrimSuffix(ret, ",")
	return ret
}

func (t TemplatePoints) GetPoint(pskey string, point string) TemplatePoint {
	var ret TemplatePoint
	if _, ok := t[SetPointName(pskey, point)]; ok {
		ret = t[SetPointName(pskey, point)]
	}
	return ret
}

func SetPointName(pskey string, point string) string {
	return pskey + "." + point
}

func (sg *SunGrow) GetPointNamesFromTemplate(template string) TemplatePoints {
	ret := make(TemplatePoints)

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
			// if _, ok := ret[dn]; !ok {
			// 	ret[dn] = make(TemplateDevices)
			// }
			for _, pr := range dr.Points {
				ret[SetPointName(dn, pr.PointID)] = TemplatePoint{
					PsKey:       dn,
					PointId:     pr.PointID,
					Description: pr.PointName,
					Unit:        pr.Unit,
				}
			}
		}
	}

	return ret
}

func (sg *SunGrow) GetTemplateData(date string, template string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		pointNames := sg.GetPointNamesFromTemplate(template)
		fmt.Printf("%v", pointNames)

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)
		psId := sg.GetPsId()

		ep2 := sg.GetByStruct(
			"AppService.queryMutiPointDataList",
			queryMutiPointDataList.RequestData{
				PsID:           psId,
				PsKey:          pointNames.PrintKeys(),
				Points:         pointNames.PrintPoints(),
				MinuteInterval: "5",
				StartTimeStamp: when.GetDayStartTimestamp(),
				EndTimeStamp:   when.GetDayEndTimestamp(),
			},
			DefaultCacheTimeout,
		)
		if ep2.IsError() {
			sg.Error = ep2.GetError()
			break
		}

		//
		csv := api.NewCsv()
		csv = csv.SetHeader([]string{
			"Date/Time",
			"PointId Name",
			"Point Name",
			"Value",
			"Units",
		})

		data2 := queryMutiPointDataList.AssertResultData(ep2)
		for deviceName, deviceRef := range data2.Devices {
			for pointId, pointRef := range deviceRef.Points {
				for _, tim := range pointRef.Times {
					foo := []string{
						tim.Key.PrintFull(),
						deviceName,
						fmt.Sprintf("%s (%s)", pointNames[SetPointName(deviceName, pointId)].Description, pointId),
						tim.Value,
						pointNames[SetPointName(deviceName, pointId)].Unit,
					}
					csv = csv.AddRow(foo)
				}
			}
		}

		switch {
		case sg.OutputType.IsNone():

		case sg.OutputType.IsHuman():
			csv.Print()

		case sg.OutputType.IsFile():
			a := queryMutiPointDataList.Assert(ep2)
			suffix := fmt.Sprintf("%s-%s", when, template)
			fn := a.GetCsvFilename(suffix)
			sg.Error = csv.WriteFile(fn, api.DefaultFileMode)

		case sg.OutputType.IsRaw():
			fmt.Println(ep2.GetData(true))

		case sg.OutputType.IsJson():
			fmt.Println(ep2.GetData(false))

		default:
		}
	}

	return sg.Error
}
