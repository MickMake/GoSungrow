package api

import (
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"strings"
)


type TemplatePoint struct {
	Name    string
	PsKey   string
	PointId valueTypes.PointId
	Unit    string
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

func (t *TemplatePoints) GetPoint(pskey string, point valueTypes.PointId) TemplatePoint {
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
				pa[1] = valueTypes.SetPointIdString(pa[1]).String()
				// pa[1] = "p" + strings.TrimPrefix(pa[1], "p")
				ret = append(ret, TemplatePoint{
					Name:    "",
					PsKey:   pa[0],
					PointId: valueTypes.SetPointIdString(pa[1]),
					Unit:    "",
				})
			}
		}
	}
	return ret
}
