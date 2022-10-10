package getHouseholdStoragePsReport

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const Url = "/v1/powerStationService/getHouseholdStoragePsReport"
const Disabled = false

type RequestData struct {
	DateID   string `json:"date_id" required:"true"`
	DateType string `json:"date_type" required:"true"`
	PsID     api.Integer `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("ps_id: Can be fetched from getPsList.")
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	ConnectType       string `json:"connect_type"`
	HasAmmeter        api.Integer  `json:"has_ammeter"`
	IsHaveEsInverter  api.Integer  `json:"is_have_es_inverter"`
	IsTransformSystem string `json:"is_transform_system"`

	DayData   *DayData   `json:"day_data"`
	MonthData *MonthData `json:"month_data"`
	YearData  *YearData  `json:"year_data"`
	TotalData *TotalData `json:"total_data"`
}

type DayData struct {
	JthdMap         api.UnitValue `json:"jthd_map"`
	JthdMapVirgin   api.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap         api.UnitValue `json:"jtyd_map"`
	JtydMapVirgin   api.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83072Map       api.UnitValue `json:"p83072_map"  PointId:"p83072" PointType:"PointTypeDaily"`
	P83072MapVirgin api.UnitValue `json:"p83072_map_virgin"  PointIgnore:"true"`
	P83077Map       api.UnitValue `json:"p83077_map"  PointId:"p83077" PointType:"PointTypeDaily"`
	P83077MapVirgin api.UnitValue `json:"p83077_map_virgin"  PointIgnore:"true"`
	P83088Map       api.UnitValue `json:"p83088_map"  PointId:"p83088" PointType:"PointTypeDaily"`
	P83088MapVirgin api.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83089Map       api.UnitValue `json:"p83089_map"  PointId:"p83089" PointType:"PointTypeDaily"`
	P83089MapVirgin api.UnitValue `json:"p83089_map_virgin"  PointIgnore:"true"`
	P83097Map       api.UnitValue `json:"p83097_map"  PointId:"p83097" PointType:"PointTypeDaily"`
	P83097MapVirgin api.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83102Map       api.UnitValue `json:"p83102_map"  PointId:"p83102" PointType:"PointTypeDaily"`
	P83102MapVirgin api.UnitValue `json:"p83102_map_virgin"  PointIgnore:"true"`
	P83118Map       api.UnitValue `json:"p83118_map"  PointId:"p83118" PointType:"PointTypeDaily"`
	P83118MapVirgin api.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map       api.UnitValue `json:"p83119_map"  PointId:"p83119" PointType:"PointTypeDaily"`
	P83119MapVirgin api.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map       api.UnitValue `json:"p83120_map"  PointId:"p83120" PointType:"PointTypeDaily"`
	P83120MapVirgin api.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121          string        `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeDaily"`
	P83122          string        `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeDaily"`
	PointData15List []struct {
		P83076     string `json:"p83076"  PointId:"p83076" PointUnit:"W" PointType:"PointTypeDaily"`
		P83076Unit string `json:"p83076_unit"`
		P83080     string `json:"p83080"  PointId:"p83080" PointUnit:"W" PointType:"PointTypeDaily"`
		P83080Unit string `json:"p83080_unit"`
		P83086     string `json:"p83086"  PointId:"p83086" PointUnit:"W" PointType:"PointTypeDaily"`
		P83086Unit string `json:"p83086_unit"`
		P83087     string `json:"p83087"  PointId:"p83087" PointUnit:"W" PointType:"PointTypeDaily"`
		P83087Unit string `json:"p83087_unit"`
		P83096     string `json:"p83096"  PointId:"p83096" PointUnit:"W" PointType:"PointTypeDaily"`
		P83096Unit string `json:"p83096_unit"`
		P83101     string `json:"p83101"  PointId:"p83101" PointUnit:"W" PointType:"PointTypeDaily"`
		P83101Unit string `json:"p83101_unit"`
		P83106     string `json:"p83106"  PointId:"p83106" PointUnit:"W" PointType:"PointTypeDaily"`
		P83106Unit string `json:"p83106_unit"`
		P83128     string `json:"p83128"  PointId:"p83128" PointUnit:"W" PointType:"PointTypeDaily"`
		P83128Unit string `json:"p83128_unit"`
		TimeStamp  string `json:"time_stamp"`
		Zfzy       string `json:"zfzy"`
		ZfzyUnit   string `json:"zfzy_unit"`
	} `json:"point_data_15_list"`
	ZfzyMap       api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

type MonthData struct {
	JthdMap          api.UnitValue `json:"jthd_map"`
	JthdMapVirgin    api.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap          api.UnitValue `json:"jtyd_map"`
	JtydMapVirgin    api.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	MonthDataDayList []struct {
		DateID                   api.Integer `json:"date_id"`
		Jthd                     string      `json:"jthd"`
		JthdUnit                 string      `json:"jthd_unit"`
		Jtyd                     string      `json:"jtyd"`
		JtydUnit                 string      `json:"jtyd_unit"`
		P83022                   string      `json:"p83022"  PointId:"p83022" PointUnit:"" PointType:"PointTypeMonthly"`
		P83072                   string      `json:"p83072"  PointId:"p83072" PointUnit:"" PointType:"PointTypeMonthly"`
		P83072Unit               string      `json:"p83072_unit"`
		P83077                   string      `json:"p83077"  PointId:"p83077" PointUnit:"" PointType:"PointTypeMonthly"`
		P83077Unit               string      `json:"p83077_unit"`
		P83088                   string      `json:"p83088"  PointId:"p83088" PointUnit:"" PointType:"PointTypeMonthly"`
		P83088Unit               string      `json:"p83088_unit"`
		P83089                   string      `json:"p83089"  PointId:"p83089" PointUnit:"" PointType:"PointTypeMonthly"`
		P83089Unit               string      `json:"p83089_unit"`
		P83097                   string      `json:"p83097"  PointId:"p83097" PointUnit:"" PointType:"PointTypeMonthly"`
		P83097Unit               string      `json:"p83097_unit"`
		P83102                   string      `json:"p83102"  PointId:"p83102" PointUnit:"" PointType:"PointTypeMonthly"`
		P83102Unit               string      `json:"p83102_unit"`
		P83118                   string      `json:"p83118"  PointId:"p83118" PointUnit:"" PointType:"PointTypeMonthly"`
		P83118Unit               string      `json:"p83118_unit"`
		P83119                   string      `json:"p83119"  PointId:"p83119" PointUnit:"" PointType:"PointTypeMonthly"`
		P83119Unit               string      `json:"p83119_unit"`
		P83120                   string      `json:"p83120"  PointId:"p83120" PointUnit:"" PointType:"PointTypeMonthly"`
		P83121                   string      `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeMonthly"`
		P83122                   string      `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeMonthly"`
		PsID                     interface{} `json:"ps_id"`
		SelfConsumptionYield     string      `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string      `json:"self_consumption_yield_unit"`
		TimeStamp                string      `json:"time_stamp"`
	} `json:"month_data_day_list"`
	P83073Map       api.UnitValue `json:"p83073_map"  PointId:"p83073" PointType:"PointTypeMonthly"`
	P83073MapVirgin api.UnitValue `json:"p83073_map_virgin"  PointIgnore:"true"`
	P83078Map       api.UnitValue `json:"p83078_map"  PointId:"p83078" PointType:"PointTypeMonthly"`
	P83078MapVirgin api.UnitValue `json:"p83078_map_virgin"  PointIgnore:"true"`
	P83088Map       api.UnitValue `json:"p83088_map"  PointId:"p83088" PointType:"PointTypeMonthly"`
	P83088MapVirgin api.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83091Map       api.UnitValue `json:"p83091_map"  PointId:"p83091" PointType:"PointTypeMonthly"`
	P83091MapVirgin api.UnitValue `json:"p83091_map_virgin"  PointIgnore:"true"`
	P83097Map       api.UnitValue `json:"p83097_map"  PointId:"p83097" PointType:"PointTypeMonthly"`
	P83097MapVirgin api.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83103Map       api.UnitValue `json:"p83103_map"  PointId:"p83103" PointType:"PointTypeMonthly"`
	P83103MapVirgin api.UnitValue `json:"p83103_map_virgin"  PointIgnore:"true"`
	P83118Map       api.UnitValue `json:"p83118_map"  PointId:"p83118" PointType:"PointTypeMonthly"`
	P83118MapVirgin api.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119Map       api.UnitValue `json:"p83119_map"  PointId:"p83119" PointType:"PointTypeMonthly"`
	P83119MapVirgin api.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120Map       api.UnitValue `json:"p83120_map"  PointId:"p83120" PointType:"PointTypeMonthly"`
	P83120MapVirgin api.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121          string        `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeMonthly"`
	P83122          string        `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeMonthly"`
	ZfzyMap         api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin   api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap         api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin   api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

type YearData struct {
	JthdMap           api.UnitValue `json:"jthd_map"`
	JthdMapVirgin     api.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap           api.UnitValue `json:"jtyd_map"`
	JtydMapVirgin     api.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83074         api.UnitValue `json:"p83074_map"  PointId:"p83074" PointType:"PointTypeYearly"`
	P83074MapVirgin   api.UnitValue `json:"p83074_map_virgin"  PointIgnore:"true"`
	P83079         api.UnitValue `json:"p83079_map"  PointId:"p83079" PointType:"PointTypeYearly"`
	P83079MapVirgin   api.UnitValue `json:"p83079_map_virgin"  PointIgnore:"true"`
	P83088         api.UnitValue `json:"p83088_map"  PointId:"p83088" PointType:"PointTypeYearly"`
	P83088MapVirgin   api.UnitValue `json:"p83088_map_virgin"  PointIgnore:"true"`
	P83093         api.UnitValue `json:"p83093_map"  PointId:"p83093" PointType:"PointTypeYearly"`
	P83093MapVirgin   api.UnitValue `json:"p83093_map_virgin"  PointIgnore:"true"`
	P83097         api.UnitValue `json:"p83097_map"  PointId:"p83097" PointType:"PointTypeYearly"`
	P83097MapVirgin   api.UnitValue `json:"p83097_map_virgin"  PointIgnore:"true"`
	P83104         api.UnitValue `json:"p83104_map"  PointId:"p83104" PointType:"PointTypeYearly"`
	P83104MapVirgin   api.UnitValue `json:"p83104_map_virgin"  PointIgnore:"true"`
	P83118         api.UnitValue `json:"p83118_map"  PointId:"p83118" PointType:"PointTypeYearly"`
	P83118MapVirgin   api.UnitValue `json:"p83118_map_virgin"  PointIgnore:"true"`
	P83119         api.UnitValue `json:"p83119_map"  PointId:"p83119" PointType:"PointTypeYearly"`
	P83119MapVirgin   api.UnitValue `json:"p83119_map_virgin"  PointIgnore:"true"`
	P83120         api.UnitValue `json:"p83120_map"  PointId:"p83120" PointType:"PointTypeYearly"`
	P83120MapVirgin   api.UnitValue `json:"p83120_map_virgin"  PointIgnore:"true"`
	P83121            string        `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeYearly"`
	P83122            string        `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeYearly"`
	YearDataMonthList []struct {
		DateID                   api.Integer  `json:"date_id"`
		Jthd                     string `json:"jthd"`
		JthdUnit                 string `json:"jthd_unit"`
		Jtyd                     string `json:"jtyd"`
		JtydUnit                 string `json:"jtyd_unit"`
		P83037                   string `json:"p83037"  PointId:"p83037" PointUnit:"" PointType:"PointTypeYearly"`
		P83073                   string `json:"p83073"  PointId:"p83073" PointUnit:"" PointType:"PointTypeYearly"`
		P83073Unit               string `json:"p83073_unit"`
		P83078                   string `json:"p83078"  PointId:"p83078" PointUnit:"" PointType:"PointTypeYearly"`
		P83078Unit               string `json:"p83078_unit"`
		P83088                   string `json:"p83088"  PointId:"p83088" PointUnit:"" PointType:"PointTypeYearly"`
		P83088Unit               string `json:"p83088_unit"`
		P83091                   string `json:"p83091"  PointId:"p83091" PointUnit:"" PointType:"PointTypeYearly"`
		P83091Unit               string `json:"p83091_unit"`
		P83098                   string `json:"p83098"  PointId:"p83098" PointUnit:"" PointType:"PointTypeYearly"`
		P83098Unit               string `json:"p83098_unit"`
		P83103                   string `json:"p83103"  PointId:"p83103" PointUnit:"" PointType:"PointTypeYearly"`
		P83103Unit               string `json:"p83103_unit"`
		P83118                   string `json:"p83118"  PointId:"p83118" PointUnit:"" PointType:"PointTypeYearly"`
		P83118Unit               string `json:"p83118_unit"`
		P83119                   string `json:"p83119"  PointId:"p83119" PointUnit:"" PointType:"PointTypeYearly"`
		P83119Unit               string `json:"p83119_unit"`
		P83120                   string `json:"p83120"  PointId:"p83120" PointUnit:"" PointType:"PointTypeYearly"`
		P83121                   string `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeYearly"`
		P83122                   string `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeYearly"`
		PsID                     api.Integer  `json:"ps_id"`
		SelfConsumptionYield     string `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string `json:"self_consumption_yield_unit"`
		TimeStamp                string `json:"time_stamp"`
	} `json:"year_data_month_list"`
	ZfzyMap       api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

type TotalData struct {
	JthdMap           api.UnitValue `json:"jthd_map"`
	JthdMapVirgin     api.UnitValue `json:"jthd_map_virgin"  PointIgnore:"true"`
	JtydMap           api.UnitValue `json:"jtyd_map"`
	JtydMapVirgin     api.UnitValue `json:"jtyd_map_virgin"  PointIgnore:"true"`
	P83075            api.UnitValue `json:"p83075_map"  PointId:"p83075" PointType:"PointTypeTotal"`
	P83075MapVirgin   api.UnitValue `json:"p83075_map_virgin"  PointIgnore:"true"`
	P83094            api.UnitValue `json:"p83094_map"  PointId:"p83094" PointType:"PointTypeTotal"`
	P83094MapVirgin   api.UnitValue `json:"p83094_map_virgin"  PointIgnore:"true"`
	P83095            api.UnitValue `json:"p83095_map"  PointId:"p83095" PointType:"PointTypeTotal"`
	P83095MapVirgin   api.UnitValue `json:"p83095_map_virgin"  PointIgnore:"true"`
	P83105            api.UnitValue `json:"p83105_map"  PointId:"p83105" PointType:"PointTypeTotal"`
	P83105MapVirgin   api.UnitValue `json:"p83105_map_virgin"  PointIgnore:"true"`
	P83107            api.UnitValue `json:"p83107_map"  PointId:"p83107" PointType:"PointTypeTotal"`
	P83107MapVirgin   api.UnitValue `json:"p83107_map_virgin"  PointIgnore:"true"`
	P83123            api.UnitValue `json:"p83123_map"  PointId:"p83123" PointType:"PointTypeTotal"`
	P83123MapVirgin   api.UnitValue `json:"p83123_map_virgin"  PointIgnore:"true"`
	P83124            api.UnitValue `json:"p83124_map"  PointId:"p83124" PointType:"PointTypeTotal"`
	P83124MapVirgin   api.UnitValue `json:"p83124_map_virgin"  PointIgnore:"true"`
	P83125            string        `json:"p83125"  PointId:"p83125" PointUnit:"" PointType:"PointTypeTotal"`
	P83126            string        `json:"p83126"  PointId:"p83126" PointUnit:"" PointType:"PointTypeTotal"`
	P83127            api.UnitValue `json:"p83127_map"  PointId:"p83127" PointType:"PointTypeTotal"`
	P83127MapVirgin   api.UnitValue `json:"p83127_map_virgin"  PointIgnore:"true"`
	TotalDataYearList []struct {
		DateID                   api.Integer  `json:"date_id"`
		Jthd                     string `json:"jthd"`
		JthdUnit                 string `json:"jthd_unit"`
		Jtyd                     string `json:"jtyd"`
		JtydUnit                 string `json:"jtyd_unit"`
		P83038                   api.Integer  `json:"p83038"  PointId:"p83038" PointUnit:"" PointType:"PointTypeTotal"`
		P83074                   string `json:"p83074"  PointId:"p83074" PointUnit:"" PointType:"PointTypeTotal"`
		P83074Unit               string `json:"p83074_unit"`
		P83079                   string `json:"p83079"  PointId:"p83079" PointUnit:"" PointType:"PointTypeTotal"`
		P83079Unit               string `json:"p83079_unit"`
		P83088                   string `json:"p83088"  PointId:"p83088" PointUnit:"" PointType:"PointTypeTotal"`
		P83088Unit               string `json:"p83088_unit"`
		P83093                   string `json:"p83093"  PointId:"p83093" PointUnit:"" PointType:"PointTypeTotal"`
		P83093Unit               string `json:"p83093_unit"`
		P83099                   string `json:"p83099"  PointId:"p83099" PointUnit:"" PointType:"PointTypeTotal"`
		P83099Unit               string `json:"p83099_unit"`
		P83104                   string `json:"p83104"  PointId:"p83104" PointUnit:"" PointType:"PointTypeTotal"`
		P83104Unit               string `json:"p83104_unit"`
		P83118                   string `json:"p83118"  PointId:"p83118" PointUnit:"" PointType:"PointTypeTotal"`
		P83118Unit               string `json:"p83118_unit"`
		P83119                   string `json:"p83119"  PointId:"p83119" PointUnit:"" PointType:"PointTypeTotal"`
		P83119Unit               string `json:"p83119_unit"`
		P83120                   api.Integer  `json:"p83120"  PointId:"p83120" PointUnit:"" PointType:"PointTypeTotal"`
		P83121                   string `json:"p83121"  PointId:"p83121" PointUnit:"" PointType:"PointTypeTotal"`
		P83122                   string `json:"p83122"  PointId:"p83122" PointUnit:"" PointType:"PointTypeTotal"`
		PsID                     api.Integer  `json:"ps_id"`
		SelfConsumptionYield     string `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string `json:"self_consumption_yield_unit"`
		TimeStamp                string `json:"time_stamp"`
	} `json:"total_data_year_list"`
	ZfzyMap       api.UnitValue `json:"zfzy_map"`
	ZfzyMapVirgin api.UnitValue `json:"zfzy_map_virgin"  PointIgnore:"true"`
	ZjzzMap       api.UnitValue `json:"zjzz_map"`
	ZjzzMapVirgin api.UnitValue `json:"zjzz_map_virgin"  PointIgnore:"true"`
}

type DecodeResultData ResultData

func (p *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		var pd DecodeResultData

		// Store ResultData
		_ = json.Unmarshal(data, &pd)
		p.ConnectType = pd.ConnectType
		p.HasAmmeter = pd.HasAmmeter
		p.IsHaveEsInverter = pd.IsHaveEsInverter
		p.IsTransformSystem = pd.IsTransformSystem

		if strings.Contains(string(data), "day_data") {
			// Store ResultData.DayData
			err = json.Unmarshal(data, &pd.DayData)
			p.DayData = pd.DayData
			break
		}

		if strings.Contains(string(data), "month_data") {
			// Store ResultData.MonthData
			_ = json.Unmarshal(data, &pd.MonthData)
			p.MonthData = pd.MonthData
			break
		}

		if strings.Contains(string(data), "year_data") {
			// Store ResultData.YearData
			_ = json.Unmarshal(data, &pd.YearData)
			p.YearData = pd.YearData
			break
		}

		if strings.Contains(string(data), "total_data") {
			// Store ResultData.TotalData
			_ = json.Unmarshal(data, &pd.TotalData)
			p.TotalData = pd.TotalData
			break
		}
	}

	return err
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		if e.Response.ResultData.DayData != nil {
			name := "getHouseholdStoragePsReport.Day." + e.Request.PsID.String()
			entries.StructToPoints(*e.Response.ResultData.DayData, name, e.Request.PsID.String(), time.Time{})

			// 83072
			// 83077
			// 83088
			// 83089
			// 83097
			// 83102
			// 83118
			// 83119
			// 83120

			for _, d := range e.Response.ResultData.DayData.PointData15List {
				// pid := api.SetPoint(strconv.FormatInt(p.PointID, 10))
				// name2 := fmt.Sprintf("getHouseholdStoragePsReport.Fart.%d", i)
				// si := fmt.Sprintf("%.2d", i)
				// parent := api.ParentDevice{Key: e.Request.PsID}

				uv := api.SetUnitValueString(d.P83076, d.P83076Unit)
				entries.AddUnitValue(name + ".p83076", e.Request.PsID.String(), "p83076", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83080, d.P83080Unit)
				entries.AddUnitValue(name + ".p83080", e.Request.PsID.String(), "p83080", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83086, d.P83086Unit)
				entries.AddUnitValue(name + ".p83086", e.Request.PsID.String(), "p83086", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83087, d.P83087Unit)
				entries.AddUnitValue(name + ".p83087", e.Request.PsID.String(), "p83087", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83096, d.P83096Unit)
				entries.AddUnitValue(name + ".p83096", e.Request.PsID.String(), "p83096", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83101, d.P83101Unit)
				entries.AddUnitValue(name + ".p83101", e.Request.PsID.String(), "p83101", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83106, d.P83106Unit)
				entries.AddUnitValue(name + ".p83106", e.Request.PsID.String(), "p83106", "", api.NewDateTime(d.TimeStamp), uv)
				uv = api.SetUnitValueString(d.P83128, d.P83128Unit)
				entries.AddUnitValue(name + ".p83128", e.Request.PsID.String(), "p83128", "", api.NewDateTime(d.TimeStamp), uv)
			}
		}
	}

	return entries
}
