package api

import (
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)


type DataEntry struct {
	Current    *GoStruct.Reflect    `json:"-"`
	EndPoint   string               `json:"endpoint"`
	Point      *Point               `json:"point"`
	Parent     ParentDevice         `json:"parent"`
	Date       valueTypes.DateTime  `json:"date"`
	Value      valueTypes.UnitValue `json:"value"`

	Valid      bool                 `json:"valid"`
	Hide       bool                 `json:"hide"`
	// ListHide   bool                 `json:"list_hide"`
}

func (de *DataEntry) IsValid() bool {
	var ok bool
	for range Only.Once {
		if de == nil {
			break
		}
		if de.Point == nil {
			break
		}
		if de.Point.Valid == false {
			break
		}
		ok = true
	}
	return ok
}
func (de *DataEntry) IsNotValid() bool {
	return !de.IsValid()
}

// func (de *DataEntry) FullId() string {
// 	return de.EndPoint	// + "." + de.Point.Id.String()
// }

// func (de *DataEntry) SetEndpoint(endpoint string, pointId string) {
// 	de.EndPoint = endpoint + "." + pointId
// 	de.Point.Id.SetString(pointId)
// }

// func (de *DataEntry) SetPointId(pointId string) {
// 	de.Point.Id.SetString(pointId)
// }

// func (de *DataEntry) SetPointName(name string) {
// 	if name != "" {
// 		de.Point.SetName(name)
// 	}
// }

// func (de *DataEntry) MakeState(state bool) DataEntry {
// 	var ret DataEntry
// 	for range Only.Once {
// 		// uv := valueTypes.SetUnitValueBool(state)
// 		// de.Value = uv.String()
// 		// de.ValueFloat = uv.Value()
// 		de.Value = valueTypes.SetUnitValueBool(state)
// 		de.Point.Unit = ""
// 		de.Point.ValueType = "Bool"
// 		de.Point.Valid = true
// 		de.Valid = true
// 		// de.EndPoint += ".state"
// 		de.Hide = false
// 	}
//
// 	return ret
// }

// func (de *DataEntry) MakeFloat(value float64, unit string, Type string) {
// 	for range Only.Once {
// 		if unit == "" {
// 			unit = de.Point.Unit
// 		}
// 		if Type == "" {
// 			Type = de.Point.ValueType
// 		}
// 		// uv := valueTypes.SetUnitValueFloat(value, unit, Type)
// 		// de.Value = uv.String()
// 		// de.ValueFloat = uv.Value()
// 		de.Value = valueTypes.SetUnitValueFloat(value, unit, Type)
// 		de.Valid = true
// 		de.Hide = false
// 	}
// }

// func (de *DataEntry) Copy() DataEntry {
// 	var ret DataEntry
// 	ret = *de
// 	var point Point
// 	point = *de.Point
// 	ret.Point = &point
// 	return ret
// }
