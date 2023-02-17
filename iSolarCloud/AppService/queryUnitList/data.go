package queryUnitList

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/userService/queryUnitList"
const Disabled = false
const EndPointName = "AppService.queryUnitList"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"UnitConvertId"`

	UnitConvertId valueTypes.Integer `json:"unit_conver_id" PointId:"unit_convert_id"`
	UnitName      valueTypes.String  `json:"unit_name"`
	UnitType      valueTypes.Integer `json:"unit_type"`
	IsBasicUnit   valueTypes.Bool    `json:"is_basic_unit"`
	TargetUnit    valueTypes.String  `json:"target_unit"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}
	return entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
}
