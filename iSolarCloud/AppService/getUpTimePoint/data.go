package getUpTimePoint

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/devService/getUpTimePoint"
const Disabled = false
const EndPointName = "AppService.getUpTimePoint"

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PointTimeRelation []struct {
		UpTimePointId valueTypes.Integer `json:"up_time_point_id"`
		Is24Hour      valueTypes.Bool    `json:"is_24_hour"`

		PointList []struct {
			PointId  valueTypes.Integer `json:"point_id"`
			TimeType valueTypes.Integer `json:"time_type"`
		} `json:"point_list" DataTable:"true"`	// DataTablePivot:"true"`
	} `json:"point_time_relation"`	// DataTable:"true"`
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
