package getBatchNewestPointData

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getBatchNewestPointData"
	Disabled     = false
	EndPointName = "AppService.getBatchNewestPointData"
)

type RequestData struct {
	PointIds valueTypes.String `json:"ps_key_points" required:"true"`
}

func (rd RequestData) IsValid() error {
	// rd.PointIds = valueTypes.SetStringValue("1129147_14_1_1.p13143")
	rd.PointIds = valueTypes.SetStringValue("p13143")
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct{}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
