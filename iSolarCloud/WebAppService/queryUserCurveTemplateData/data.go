package queryUserCurveTemplateData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/queryUserCurveTemplateData"
const Disabled = false


// {"template_id":"","date_type":"","start_time":"","end_time":""}

type RequestData struct {
	TemplateId valueTypes.String `json:"template_id" required:"true"`
	// DateType   valueTypes.String `json:"date_type" required:"true"`
	// StartTime  valueTypes.String `json:"start_time" required:"true"`
	// EndTime    valueTypes.String `json:"end_time" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += fmt.Sprintln("date_type: Day = 1")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("date_type: Month = 2")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYYmm")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYYmm")
	ret += fmt.Sprintln("date_type: Year = 3")
	ret += fmt.Sprintln("\tstart_time_stamp: Format YYYY")
	ret += fmt.Sprintln("\tend_time_stamp: Format YYYY")
	return ret
}


// ResultData -> PointsData -> []dDevices -> []Points
type ResultData struct {
	CreateTime   valueTypes.DateTime `json:"create_time"`
	OpenTime     valueTypes.DateTime `json:"open_time"`
	UpdateTime   valueTypes.DateTime `json:"update_time"`
	StartTime    valueTypes.DateTime `json:"start_time"`
	EndTime      valueTypes.DateTime `json:"end_time"`
	DateType     valueTypes.Integer  `json:"date_type"`
	Cycle        valueTypes.Integer  `json:"cycle"`
	TemplateId   valueTypes.Integer  `json:"template_id"`
	TemplateName valueTypes.String   `json:"template_name"`
	PointsData   PointsData          `json:"points_data" PointNameAppend:"false"`	// PointNameFromChild:"Order"`
}


type PointsData struct {
	Devices map[string]DeviceData `json:"devices" PointNameAppend:"false"`	// PointNameFromChild:"DateId" PointNameAppend:"false"`
	Order   valueTypes.String  `json:"order" PointSplitOn:","`
}

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
	Points     Points             `json:"points" PointNameAppend:"false"`
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
			fmt.Printf("queryUserCurveTemplateData[dd] - err:%s data: %s\n", err, data)
			// break
		}
		p.DeviceName = dd.DeviceName
		p.DeviceType = dd.DeviceType
		p.DeviceUUID = dd.DeviceUUID

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dd.Points)
		if err != nil {
			fmt.Printf("queryUserCurveTemplateData[dd.Points] - err:%s data: %s\n", err, data)
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
	Color      valueTypes.String  `json:"color"`
	DetailId   valueTypes.String  `json:"detail_id"`
	PointId    valueTypes.PointId `json:"point_id"`
	PointName  valueTypes.String  `json:"point_name"`
	PsId       valueTypes.PsId    `json:"ps_id"`
	PsKey      valueTypes.PsKey   `json:"ps_key"`
	PsName     valueTypes.String  `json:"ps_name"`
	Statistics valueTypes.String  `json:"statistics"`
	Style      valueTypes.String  `json:"style"`
	Unit       valueTypes.String  `json:"unit"`
	DataList   []struct {
		TimeStamp valueTypes.String `json:"time_stamp"`
		Value     valueTypes.String `json:"value"`
	} `json:"data_list" PointArrayFlatten:"false"`
}

// type Foo struct {
// 	ReqSerialNum valueTypes.String `json:"req_serial_num"`
// 	ResultCode   valueTypes.String `json:"result_code"`
// 	ResultData   struct {
// 		CreateTime valueTypes.String `json:"create_time"`
// 		Cycle      valueTypes.String `json:"cycle"`
// 		DateType   valueTypes.String `json:"date_type"`
// 		EndTime    valueTypes.String `json:"end_time"`
// 		OpenTime   valueTypes.String `json:"open_time"`
// 		PointsData struct {
// 			One171348_43_2_2 struct {
// 				DeviceName valueTypes.String `json:"device_name"`
// 				DeviceType int64  `json:"device_type"`
// 				DeviceUUID int64  `json:"device_uuid"`
// 				P58629     struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58629"`
// 				P58630 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58630"`
// 				P58631 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58631"`
// 				P58632 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58632"`
// 				P58633 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58633"`
// 				P58635 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58635"`
// 				P58636 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58636"`
// 				P58815 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58815"`
// 				P58816 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58816"`
// 				P58817 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58817"`
// 				P58818 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58818"`
// 				P58819 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58819"`
// 				P58820 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58820"`
// 				P58821 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58821"`
// 				P58822 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58822"`
// 				P58823 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58823"`
// 				P58824 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58824"`
// 				P58825 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58825"`
// 				P58826 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58826"`
// 				P58827 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58827"`
// 				P58828 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58828"`
// 				P58829 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58829"`
// 				P58830 struct {
// 					Color      valueTypes.String `json:"color"`
// 					DetailID   valueTypes.String `json:"detail_id"`
// 					PointID    valueTypes.String `json:"point_id"`
// 					PointName  valueTypes.String `json:"point_name"`
// 					PsID       valueTypes.String `json:"ps_id"`
// 					PsKey      valueTypes.String `json:"ps_key"`
// 					PsName     valueTypes.String `json:"ps_name"`
// 					Statistics valueTypes.String `json:"statistics"`
// 					Style      valueTypes.String `json:"style"`
// 					Unit       valueTypes.String `json:"unit"`
// 				} `json:"p58830"`
// 			} `json:"1171348_43_2_2"`
// 			Order valueTypes.String `json:"order"`
// 		} `json:"points_data"`
// 		StartTime    valueTypes.String `json:"start_time"`
// 		TemplateID   valueTypes.String `json:"template_id"`
// 		TemplateName valueTypes.String `json:"template_name"`
// 		UpdateTime   valueTypes.String `json:"update_time"`
// 	} `json:"result_data"`
// 	ResultMsg valueTypes.String `json:"result_msg"`
// }


func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		dt := valueTypes.NewDateTime(valueTypes.Now)
		name := pkg + "." + e.Request.TemplateId.String()
		entries.StructToPoints(e.Response.ResultData, name, e.Request.TemplateId.String(), dt)
	}

	return entries
}
