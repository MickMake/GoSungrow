package api

import (
	"GoSungrow/Only"
	"fmt"
	"strings"
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
