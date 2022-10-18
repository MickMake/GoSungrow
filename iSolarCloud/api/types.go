package api

import (
	"GoSungrow/iSolarCloud/api/output"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"sort"
)


// func UnitValueToPoint(endpoint string, t *valueTypes.UnitValue, parentId string, pid valueTypes.PointId, name string) *Point {
// 	uv := t.UnitValueFix()
//
// 	if name == "" {
// 		name = PointToName(pid)
// 	}
//
// 	ret := GetPoint(parentId, pid)
// 	if !ret.Valid {
// 		ret = &Point {
// 			// EndPoint:  endpoint,
// 			// FullId:    "",
// 			// Parent:    ParentDevice{ Key: psId },
// 			Id:        pid,
// 			GroupName: "",
// 			Name:      name,
// 			Unit:      uv.Unit(),
// 			UpdateFreq:  "PointUpdateFreqInstant",
// 			Valid:     true,
// 			States:    nil,
// 		}
// 	}
//
// 	return ret
// }


func (e *Web) GetDataTable(endpoint EndPoint) output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(endpoint.GetJsonData(false)))
		table.SetRaw([]byte(endpoint.GetJsonData(true)))

		_ = table.SetHeader(
			"Date",
			"Point Id",
			"Value",
			"Unit",
			"Unit Type",
			"Group Name",
			"Description",
			"Update Freq",
		)

		data := endpoint.GetEndPointData()
		var sorted []string
		for p := range data.Map {
			sorted = append(sorted, string(p))
		}
		sort.Strings(sorted)

		for _, p := range sorted {
			entries := data.Map[p].Entries
			for _, de := range entries {
				if de.Hide {
					continue
				}

				_ = table.AddRow(
					de.Date.Format(valueTypes.DateTimeLayout),
					p,
					de.Value,
					de.Point.Unit,
					de.Point.ValueType,
					de.Point.GroupName,
					de.Point.Description,
					de.Point.UpdateFreq,
				)
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(4),
		// 	ValueColumn:  output.SetInteger(5),
		// 	UnitsColumn:  output.SetInteger(6),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}
	return table
}
