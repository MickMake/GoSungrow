package api

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
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
