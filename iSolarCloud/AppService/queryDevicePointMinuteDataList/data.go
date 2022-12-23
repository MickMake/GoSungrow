package queryDevicePointMinuteDataList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/commonService/queryDevicePointMinuteDataList"
const Disabled = false
const EndPointName = "AppService.queryDevicePointMinuteDataList"

// ../bin/GoSungrow api raw AppService.queryDevicePointMinuteDataList '{"data_point":"p13148","data_type":"1","points":"'$(perl -e 'foreach(10000 .. 20999){print"p$_,"}')'","query_type":"1","ps_key":"1171348_14_1_2","start_time_stamp":"20221201223000","end_time_stamp":"20221201223000"}'

type RequestData struct {
	PsKey          valueTypes.PsKey    `json:"ps_key" required:"true"`
	DataPoint      valueTypes.String   `json:"data_point" required:"true"`
	StartTimeStamp valueTypes.DateTime `json:"start_time_stamp" required:"true"`
	EndTimeStamp   valueTypes.DateTime `json:"end_time_stamp" required:"true"`
	DataType       valueTypes.String   `json:"data_type" required:"true"`
	QueryType      valueTypes.String   `json:"query_type" required:"true"`
	Points         valueTypes.String   `json:"points" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := api.HelpQueryType()
	return ret
}


type ResultData struct {
	Data []DataPoint `json:"data" PointIdReplace:"false" DataTable:"true" DataTableSortOn:"TimeStamp"`
}

func (e *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
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
	GoStruct.GoStructParent `json:"-" PointIdReplace:"true" PointIdFrom:"TimeStamp" PointNameDateFormat:"DateTimeAltLayout" PointTimestampFrom:"TimeStamp"`

	TimeStamp               valueTypes.DateTime             `json:"time_stamp" PointNameDateFormat:"DateTimeLayout"`
	Points                  map[string]valueTypes.UnitValue `json:"points"`
	IsPlatformDefaultUnit   valueTypes.Bool                 `json:"is_platform_default_unit"`
}

func (e *DataPoint) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
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

		d.Points = make(map[string]valueTypes.UnitValue)	// @TODO - change to UVS
		for k, v := range mp {
			if k == "is_platform_default_unit" {
				continue
			}

			if k == "time_stamp" {
				continue
			}

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
