package getBatchNewestPointData

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"

	"fmt"
)

const Url = "/v1/devService/getBatchNewestPointData"
const Disabled = false
const EndPointName = "AppService.getBatchNewestPointData"

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

type ResultData struct {
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	return entries
}
