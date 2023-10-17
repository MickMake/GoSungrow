package getDevicePointMinuteDataList

import (
	"encoding/json"
	"fmt"

	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/only"
)

const (
	Url          = "/v1/commonService/getDevicePointMinuteDataList"
	Disabled     = false
	EndPointName = "AppService.getDevicePointMinuteDataList"
)

type RequestData struct {
	PsKey          valueTypes.PsKey  `json:"ps_key" required:"true"`
	Points         valueTypes.String `json:"points" required:"true"`
	StartTimeStamp valueTypes.String `json:"start_time_stamp" required:"true"`
	EndTimeStamp   valueTypes.String `json:"end_time_stamp" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Data []DataPoint `json:"data" PointIdReplace:"false" DataTable:"true" DataTableSortOn:"TimeStamp"`
}

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		err = json.Unmarshal(data, &e.Data)
		if err != nil {
			break
		}
	}

	return err
}

type DataPoint struct {
	GoStruct.GoStruct `json:"-" PointIdReplace:"true" PointIdFrom:"TimeStamp" PointNameDateFormat:"DateTimeAltLayout" PointTimestampFrom:"TimeStamp"`

	TimeStamp valueTypes.DateTime             `json:"time_stamp" PointNameDateFormat:"DateTimeLayout"`
	Points    map[string]valueTypes.UnitValue `json:"points"`
	// P13148                valueTypes.Integer          `json:"p13148"`
	IsPlatformDefaultUnit valueTypes.Bool `json:"is_platform_default_unit"`
}

func (e *DataPoint) UnmarshalJSON(data []byte) error {
	var err error

	for range only.Once {
		if len(data) == 0 {
			break
		}

		type decode DataPoint
		var d decode
		// Store DataPoint
		err = json.Unmarshal(data, &d)
		if err != nil {
			break
		}

		mp := map[string]interface{}{}
		err = json.Unmarshal(data, &mp)
		if err != nil {
			break
		}

		d.Points = make(map[string]valueTypes.UnitValue) // @TODO - change to UVS
		for k, v := range mp {
			if k == "is_platform_default_unit" {
				continue
			}

			if k == "time_stamp" {
				continue
			}

			// key := valueTypes.SetPointIdString(k)
			// value, ok, _ := valueTypes.AnyToUnitValue(v, "", "", "")

			// value, _, _ := valueTypes.AnyToUnitValue(v, "", "", "")
			// d.Points[k] = valueTypes.SetFloatValue(value.First().Value())

			value, _, _ := valueTypes.AnyToUnitValue(v, "", "", "", "")
			d.Points[k] = *value.First()
		}

		*e = DataPoint(d)
	}

	return err
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsKey.String(), GoStruct.NewEndPointPath(e.Request.PsKey.String()))
	return entries
}
