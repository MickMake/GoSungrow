package queryUserCurveTemplateData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/queryUserCurveTemplateData"
const Disabled = false
const EndPointName = "WebAppService.queryUserCurveTemplateData"


// {"template_id":"","date_type":"","start_time":"","end_time":""}

type RequestData struct {
	TemplateId valueTypes.String `json:"template_id" required:"true"`
	// DateType   valueTypes.String `json:"date_type" required:"true"`
	// StartTime  valueTypes.String `json:"start_time" required:"true"`
	// EndTime    valueTypes.String `json:"end_time" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += api.HelpDateId()
	return ret
}


// ResultData (struct) -> PointsData(struct) -> Devices(map[string]DeviceData) -> Points(map[string]Point)
type ResultData struct {
	CreateTime   valueTypes.DateTime `json:"create_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	OpenTime     valueTypes.DateTime `json:"open_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	UpdateTime   valueTypes.DateTime `json:"update_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	StartTime    valueTypes.DateTime `json:"start_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	EndTime      valueTypes.DateTime `json:"end_time" PointNameDateFormat:"2006-01-02 15:04:05"`
	DateType     valueTypes.Integer  `json:"date_type"`
	Cycle        valueTypes.Integer  `json:"cycle"`
	TemplateId   valueTypes.Integer  `json:"template_id"`
	TemplateName valueTypes.String   `json:"template_name"`
	PointsData   PointsData          `json:"points_data" PointIdReplace:"true"`
}


type PointsData struct {
	Devices map[string]DeviceData `json:"devices" PointIdReplace:"true"`	// PointIdFromChild:"DateId" PointIdReplace:"true"`
	Order   valueTypes.String  `json:"order" PointSplitOn:"," PointValueReplace:"&" PointValueReplaceWith:".p"`
}

type Devices map[string]DeviceData

func (p *PointsData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		type decode PointsData
		var pd decode

		// Store PointsData.Order
		err = json.Unmarshal(data, &pd)
		if err != nil {
			fmt.Printf("queryUserCurveTemplateData[pd] - err:%s data: %s\n", err, data)
			// break
		}

		p.Order = pd.Order

		// Store PointsData.dDevices
		err = json.Unmarshal(data, &pd.Devices)
		if err != nil {
			fmt.Printf("queryUserCurveTemplateData[pd.Devices] - err:%s data: %s\n", err, data)
			// break
		}

		for i, k := range pd.Devices {
			if k.DeviceName.String() == "" {
				delete(pd.Devices, i)
			}
			// fmt.Printf("%v - %v - %v\n", k.DeviceName, k.DeviceUUID, k.DeviceType)
		}
		p.Devices = pd.Devices
		err = nil
	}

	return err
}


type DeviceData struct {
	DeviceName valueTypes.String  `json:"device_name"`
	DeviceType valueTypes.Integer `json:"device_type"`
	DeviceUUID valueTypes.Integer `json:"device_uuid"`
	Points     Points             `json:"points" DataTable:"true" DataTableMerge:"true" DataTableSortOn:"PointId"`	// PointIdReplace:"false" PointIdFromChild:"PsId.PsKey"
}

func (p *DeviceData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		type decode DeviceData
		var dd decode

		// Store DeviceData extras
		err = json.Unmarshal(data, &dd)
		if err != nil {
			// fmt.Printf("queryUserCurveTemplateData[dd] - err:%s data: %s\n", err, data)
			// break
		}
		p.DeviceName = dd.DeviceName
		p.DeviceType = dd.DeviceType
		p.DeviceUUID = dd.DeviceUUID

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dd.Points)
		if err != nil {
			// fmt.Printf("queryUserCurveTemplateData[dd.Points] - err:%s data: %s\n", err, data)
			// break
		}

		for i, k := range dd.Points {
			if k.PointId.String() == "" {
				delete(dd.Points, i)
				continue
			}
		}
		p.Points = dd.Points
		err = nil
	}

	return err
}


type Points map[string]Point

func (p *Points) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		type decode Points
		var dd decode

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dd)
		if err != nil {
			fmt.Printf("queryUserCurveTemplateData[Point] - err:%s data: %s\n", err, data)
			// break
		}

		*p = Points(dd)
		err = nil
	}

	return err
}


type Point struct {
	PointId    valueTypes.PointId `json:"point_id"`
	PointName  valueTypes.String  `json:"point_name"`
	PsId       valueTypes.PsId    `json:"ps_id"`
	PsKey      valueTypes.PsKey   `json:"ps_key"`
	Color      valueTypes.String  `json:"color"`
	DetailId   valueTypes.String  `json:"detail_id"`
	PsName     valueTypes.String  `json:"ps_name"`
	Statistics valueTypes.String  `json:"statistics"`
	Style      valueTypes.String  `json:"style"`
	Unit       valueTypes.String  `json:"unit"`
	DataList   []struct {
		TimeStamp valueTypes.String `json:"time_stamp"`
		Value     valueTypes.String `json:"value"`
	} `json:"data_list" PointArrayFlatten:"false"`
}

func (p *Point) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		type decode Point
		var dd decode

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dd)
		if err != nil {
			fmt.Printf("queryUserCurveTemplateData[Point] - err:%s data: %s\n", err, data)
			// break
		}

		dd.PointId.PsKey = dd.PsKey

		*p = Point(dd)
		err = nil
	}

	return err
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.TemplateId.String(), GoStruct.NewEndPointPath(e.Request.TemplateId.String()))
	return entries
}
