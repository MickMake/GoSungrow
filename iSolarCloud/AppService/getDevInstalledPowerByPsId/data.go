package getDevInstalledPowerByPsId

import (
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

const (
	Url          = "/v1/devService/getDevInstalledPowerByPsId"
	Disabled     = false
	EndPointName = "AppService.getDevInstalledPowerByPsId"
)

type RequestData struct {
	PsId valueTypes.PsId `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		PsKey         valueTypes.String  `json:"ps_key"`
		PsId          valueTypes.Integer `json:"ps_id"`
		DeviceType    valueTypes.Integer `json:"device_type"`
		DeviceName    valueTypes.String  `json:"device_name"`
		UUID          valueTypes.Integer `json:"uuid"`
		CountryId     valueTypes.Integer `json:"country_id"`
		GridId        interface{}        `json:"grid_id"`
		PropertyCode  valueTypes.String  `json:"property_code"`
		PropertyValue valueTypes.String  `json:"property_value"`
		RatedPower    valueTypes.String  `json:"rated_power"`
	} `json:"pageList" PointId:"devices" DataTable:"true" DataTableSortOn:"PsKey"`
	RowCount   valueTypes.Integer `json:"rowCount" PointId:"row_count"`
	CurPage    valueTypes.Integer `json:"curPage" PointId:"cur_page"`
	IsMore     valueTypes.Integer `json:"isMore" PointId:"is_more"`
	Size       valueTypes.Integer `json:"size"`
	StartIndex valueTypes.Integer `json:"startIndex" PointId:"start_index"`
	TotalPage  valueTypes.Integer `json:"totalPage" PointId:"total_page"`
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
