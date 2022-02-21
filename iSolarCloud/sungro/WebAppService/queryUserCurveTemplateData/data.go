package queryUserCurveTemplateData

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
)

const Url = "/v1/devService/queryUserCurveTemplateData"

// {"template_id":"","date_type":"","start_time":"","end_time":""}
type RequestData struct {
	TemplateID string `json:"template_id" required:"true"`
	DateType   string `json:"date_type" required:"true"`
	StartTime  string `json:"start_time" required:"true"`
	EndTime    string `json:"end_time" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("template_id: Use AppService.getTemplateList for ids.")
	ret += fmt.Sprintln("date_type: Day = 1")
	ret += fmt.Sprintln("\tstart_time: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("\tend_time: Format YYYYmmddHHMM00")
	ret += fmt.Sprintln("date_type: Month = 2")
	ret += fmt.Sprintln("\tstart_time: Format YYYYmm")
	ret += fmt.Sprintln("\tend_time: Format YYYYmm")
	ret += fmt.Sprintln("date_type: Year = 3")
	ret += fmt.Sprintln("\tstart_time: Format YYYY")
	ret += fmt.Sprintln("\tend_time: Format YYYY")
	return ret
}

// ResultData -> PointsData -> []Devices -> []Points
type ResultData struct {
	CreateTime   string     `json:"create_time"`
	Cycle        string     `json:"cycle"`
	DateType     string     `json:"date_type"`
	EndTime      string     `json:"end_time"`
	OpenTime     string     `json:"open_time"`
	PointsData   PointsData `json:"points_data"`
	StartTime    string     `json:"start_time"`
	TemplateID   string     `json:"template_id"`
	TemplateName string     `json:"template_name"`
	UpdateTime   string     `json:"update_time"`
}

type PointsData struct {
	Devices Devices `json:"devices"`
	Order   string  `json:"order"`
}

type DecodePointsData PointsData

func (p *PointsData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		var pd DecodePointsData

		// Store PointsData.Order
		err = json.Unmarshal(data, &pd)
		p.Order = pd.Order

		// Store PointsData.Devices
		_ = json.Unmarshal(data, &pd.Devices)
		for i, k := range pd.Devices {
			if k.DeviceName == "" {
				delete(pd.Devices, i)
			}
			// fmt.Printf("%v - %v - %v\n", k.DeviceName, k.DeviceUUID, k.DeviceType)
		}
		p.Devices = pd.Devices
	}

	return err
}

type Devices map[string]DeviceData

type DeviceData struct {
	DeviceName string `json:"device_name"`
	DeviceType int64  `json:"device_type"`
	DeviceUUID int64  `json:"device_uuid"`
	Points     Points `json:"points"`
}

type DecodeDeviceData DeviceData

func (p *DeviceData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		var dd DecodeDeviceData

		// Store DeviceData extras
		_ = json.Unmarshal(data, &dd)
		p.DeviceName = dd.DeviceName
		p.DeviceType = dd.DeviceType
		p.DeviceUUID = dd.DeviceUUID

		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dd.Points)
		for i, k := range dd.Points {
			if k.PointID == "" {
				delete(dd.Points, i)
				continue
			}
		}
		p.Points = dd.Points
	}

	return err
}

type Points map[string]Point

type Point struct {
	Color    string `json:"color"`
	DataList []struct {
		TimeStamp string `json:"time_stamp"`
		Value     string `json:"value"`
	} `json:"data_list"`
	DetailID   string `json:"detail_id"`
	PointID    string `json:"point_id"`
	PointName  string `json:"point_name"`
	PsID       string `json:"ps_id"`
	PsKey      string `json:"ps_key"`
	PsName     string `json:"ps_name"`
	Statistics string `json:"statistics"`
	Style      string `json:"style"`
	Unit       string `json:"unit"`
}
