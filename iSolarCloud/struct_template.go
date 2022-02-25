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
type TemplatePoints []TemplatePoint


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
		ret += fmt.Sprintf("%s,", p.PointId)
	}
	ret = strings.TrimSuffix(ret, ",")
	return ret
}

func (t *TemplatePoints) GetPoint(pskey string, point string) TemplatePoint {
	var ret TemplatePoint
	for _, k := range *t {
		if k.PsKey != pskey {
			continue
		}
		if k.PointId != point {
			continue
		}
		ret = k
		break
	}
	return ret
}

func CreatePoints(points []string) TemplatePoints {
	var ret TemplatePoints
	for range Only.Once {
		// Feed in a string array and generate points data.
		// strings can be either "pskey/point_id", "pskey.point_id", "pskey:point_id",
		for _, p := range points {
			pa := strings.Split(p, ".")
			if len(pa) == 2 {
				ret = append(ret, TemplatePoint{
					Description: "",
					PsKey:       pa[0],
					PointId:     pa[1],
					Unit:        "",
				})
			}
		}
	}
	return ret
}

func SetPointName(pskey string, point string) string {
	point = strings.TrimPrefix(point, "p")
	return pskey + ".p" + point
}

func (sg *SunGrow) GetPointNamesFromTemplate(template string) TemplatePoints {
	var ret TemplatePoints

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
					ret = append(ret, TemplatePoint {
					PsKey:       dn,
					PointId:     "p"+pr.PointID,
					Description: pr.PointName,
					Unit:        pr.Unit,
				})
			}
		}
	}

	return ret
}

func (sg *SunGrow) GetTemplateData(date string, template string, filter string) error {
	for range Only.Once {
		if template == "" {
			template = "8042"
		}

		if date == "" {
			date = api.NewDateTime("").String()
		}
		when := api.NewDateTime(date)
		psId := sg.GetPsId()
		pointNames := sg.GetPointNamesFromTemplate(template)

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

		//
		csv := api.NewCsv()
		csv = csv.SetHeader(
			"Date/Time",
			"PointId Name",
			"Point Name",
			"Value",
			"Units",
		)

		data := queryMutiPointDataList.AssertResultData(ep)
		for deviceName, deviceRef := range data.Devices {
			for pointId, pointRef := range deviceRef.Points {
				for _, tim := range pointRef.Times {
					gp := pointNames.GetPoint(deviceName, pointId)
					csv = csv.AddRow(
						tim.Key.PrintFull(),
						fmt.Sprintf("%s.%s", deviceName, pointId),
						gp.Description,
						tim.Value,
						gp.Unit,
					)
				}
			}
		}

		switch {
			case sg.OutputType.IsNone():

			case sg.OutputType.IsHuman():
				csv.Print()

			case sg.OutputType.IsGraph():
				gr := api.JsonToGraphRequest(filter)
				if gr.Error != nil {
					sg.Error = gr.Error
					break
				}
				sg.Error = csv.Graph(gr)
				// api.GraphRequest {
				// 	Title:        "Testing 1. 2. 3.",
				// 	TimeColumn:   1,
				// 	ValueColumn:  4,
				// 	SearchColumn: 3,
				// 	SearchString: "p83106",
				// 	FileName:     "foo.png",
				// }

			case sg.OutputType.IsFile():
				a := queryMutiPointDataList.Assert(ep)
				suffix := fmt.Sprintf("%s-%s", when, template)
				fn := a.GetCsvFilename(suffix)
				sg.Error = csv.WriteFile(fn, api.DefaultFileMode)

			case sg.OutputType.IsRaw():
				fmt.Println(ep.GetData(true))

			case sg.OutputType.IsJson():
				fmt.Println(ep.GetData(false))

			default:
		}
	}

	return sg.Error
}
