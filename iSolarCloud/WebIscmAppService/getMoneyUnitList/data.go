package getMoneyUnitList

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/commonService/getMoneyUnitList"
	Disabled     = false
	EndPointName = "WebIscmAppService.getMoneyUnitList"
)

type RequestData struct{}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []struct {
	GoStructParent GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"CodeId"`

	CodeId   valueTypes.String `json:"code_id"`
	CodeName valueTypes.String `json:"code_name"`
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

	return entries
}
