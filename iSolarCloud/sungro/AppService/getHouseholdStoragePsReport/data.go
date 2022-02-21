package getHouseholdStoragePsReport

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	"strings"
)

const Url = "/v1/powerStationService/getHouseholdStoragePsReport"

type RequestData struct {
	DateID   string `json:"date_id" required:"true"`
	DateType string `json:"date_type" required:"true"`
	PsID     string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintln("ps_id: Can be fetched from getPsList.")
	ret += fmt.Sprintln("date_type: Day = 1")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmmdd")
	ret += fmt.Sprintln("date_type: Month = 2")
	ret += fmt.Sprintln("\tdate_id: Format YYYYmm")
	ret += fmt.Sprintln("date_type: Year = 3")
	ret += fmt.Sprintln("\tdate_id: Format YYYY")
	return ret
}

type ResultData struct {
	ConnectType       string `json:"connect_type"`
	HasAmmeter        int64  `json:"has_ammeter"`
	IsHaveEsInverter  int64  `json:"is_have_es_inverter"`
	IsTransformSystem string `json:"is_transform_system"`

	DayData   *DayData   `json:"day_data"`
	MonthData *MonthData `json:"month_data"`
	YearData  *YearData  `json:"year_data"`
	TotalData *TotalData `json:"total_data"`
}

type DayData struct {
	JthdMap         KeyValue `json:"jthd_map"`
	JthdMapVirgin   KeyValue `json:"jthd_map_virgin"`
	JtydMap         KeyValue `json:"jtyd_map"`
	JtydMapVirgin   KeyValue `json:"jtyd_map_virgin"`
	P83072Map       KeyValue `json:"p83072_map"`
	P83072MapVirgin KeyValue `json:"p83072_map_virgin"`
	P83077Map       KeyValue `json:"p83077_map"`
	P83077MapVirgin KeyValue `json:"p83077_map_virgin"`
	P83088Map       KeyValue `json:"p83088_map"`
	P83088MapVirgin KeyValue `json:"p83088_map_virgin"`
	P83089Map       KeyValue `json:"p83089_map"`
	P83089MapVirgin KeyValue `json:"p83089_map_virgin"`
	P83097Map       KeyValue `json:"p83097_map"`
	P83097MapVirgin KeyValue `json:"p83097_map_virgin"`
	P83102Map       KeyValue `json:"p83102_map"`
	P83102MapVirgin KeyValue `json:"p83102_map_virgin"`
	P83118Map       KeyValue `json:"p83118_map"`
	P83118MapVirgin KeyValue `json:"p83118_map_virgin"`
	P83119Map       KeyValue `json:"p83119_map"`
	P83119MapVirgin KeyValue `json:"p83119_map_virgin"`
	P83120Map       KeyValue `json:"p83120_map"`
	P83120MapVirgin KeyValue `json:"p83120_map_virgin"`
	P83121          string   `json:"p83121"`
	P83122          string   `json:"p83122"`
	PointData15List []struct {
		P83076     string `json:"p83076"`
		P83076Unit string `json:"p83076_unit"`
		P83080     string `json:"p83080"`
		P83080Unit string `json:"p83080_unit"`
		P83086     string `json:"p83086"`
		P83086Unit string `json:"p83086_unit"`
		P83087     string `json:"p83087"`
		P83087Unit string `json:"p83087_unit"`
		P83096     string `json:"p83096"`
		P83096Unit string `json:"p83096_unit"`
		P83101     string `json:"p83101"`
		P83101Unit string `json:"p83101_unit"`
		P83106     string `json:"p83106"`
		P83106Unit string `json:"p83106_unit"`
		P83128     string `json:"p83128"`
		P83128Unit string `json:"p83128_unit"`
		TimeStamp  string `json:"time_stamp"`
		Zfzy       string `json:"zfzy"`
		ZfzyUnit   string `json:"zfzy_unit"`
	} `json:"point_data_15_list"`
	ZfzyMap       KeyValue `json:"zfzy_map"`
	ZfzyMapVirgin KeyValue `json:"zfzy_map_virgin"`
	ZjzzMap       KeyValue `json:"zjzz_map"`
	ZjzzMapVirgin KeyValue `json:"zjzz_map_virgin"`
}

type MonthData struct {
	JthdMap          KeyValue `json:"jthd_map"`
	JthdMapVirgin    KeyValue `json:"jthd_map_virgin"`
	JtydMap          KeyValue `json:"jtyd_map"`
	JtydMapVirgin    KeyValue `json:"jtyd_map_virgin"`
	MonthDataDayList []struct {
		DateID                   int64       `json:"date_id"`
		Jthd                     string      `json:"jthd"`
		JthdUnit                 string      `json:"jthd_unit"`
		Jtyd                     string      `json:"jtyd"`
		JtydUnit                 string      `json:"jtyd_unit"`
		P83022                   string      `json:"p83022"`
		P83072                   string      `json:"p83072"`
		P83072Unit               string      `json:"p83072_unit"`
		P83077                   string      `json:"p83077"`
		P83077Unit               string      `json:"p83077_unit"`
		P83088                   string      `json:"p83088"`
		P83088Unit               string      `json:"p83088_unit"`
		P83089                   string      `json:"p83089"`
		P83089Unit               string      `json:"p83089_unit"`
		P83097                   string      `json:"p83097"`
		P83097Unit               string      `json:"p83097_unit"`
		P83102                   string      `json:"p83102"`
		P83102Unit               string      `json:"p83102_unit"`
		P83118                   string      `json:"p83118"`
		P83118Unit               string      `json:"p83118_unit"`
		P83119                   string      `json:"p83119"`
		P83119Unit               string      `json:"p83119_unit"`
		P83120                   string      `json:"p83120"`
		P83121                   string      `json:"p83121"`
		P83122                   string      `json:"p83122"`
		PsID                     interface{} `json:"ps_id"`
		SelfConsumptionYield     string      `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string      `json:"self_consumption_yield_unit"`
		TimeStamp                string      `json:"time_stamp"`
	} `json:"month_data_day_list"`
	P83073Map       KeyValue `json:"p83073_map"`
	P83073MapVirgin KeyValue `json:"p83073_map_virgin"`
	P83078Map       KeyValue `json:"p83078_map"`
	P83078MapVirgin KeyValue `json:"p83078_map_virgin"`
	P83088Map       KeyValue `json:"p83088_map"`
	P83088MapVirgin KeyValue `json:"p83088_map_virgin"`
	P83091Map       KeyValue `json:"p83091_map"`
	P83091MapVirgin KeyValue `json:"p83091_map_virgin"`
	P83097Map       KeyValue `json:"p83097_map"`
	P83097MapVirgin KeyValue `json:"p83097_map_virgin"`
	P83103Map       KeyValue `json:"p83103_map"`
	P83103MapVirgin KeyValue `json:"p83103_map_virgin"`
	P83118Map       KeyValue `json:"p83118_map"`
	P83118MapVirgin KeyValue `json:"p83118_map_virgin"`
	P83119Map       KeyValue `json:"p83119_map"`
	P83119MapVirgin KeyValue `json:"p83119_map_virgin"`
	P83120Map       KeyValue `json:"p83120_map"`
	P83120MapVirgin KeyValue `json:"p83120_map_virgin"`
	P83121          string   `json:"p83121"`
	P83122          string   `json:"p83122"`
	ZfzyMap         KeyValue `json:"zfzy_map"`
	ZfzyMapVirgin   KeyValue `json:"zfzy_map_virgin"`
	ZjzzMap         KeyValue `json:"zjzz_map"`
	ZjzzMapVirgin   KeyValue `json:"zjzz_map_virgin"`
}

type YearData struct {
	JthdMap           KeyValue `json:"jthd_map"`
	JthdMapVirgin     KeyValue `json:"jthd_map_virgin"`
	JtydMap           KeyValue `json:"jtyd_map"`
	JtydMapVirgin     KeyValue `json:"jtyd_map_virgin"`
	P83074Map         KeyValue `json:"p83074_map"`
	P83074MapVirgin   KeyValue `json:"p83074_map_virgin"`
	P83079Map         KeyValue `json:"p83079_map"`
	P83079MapVirgin   KeyValue `json:"p83079_map_virgin"`
	P83088Map         KeyValue `json:"p83088_map"`
	P83088MapVirgin   KeyValue `json:"p83088_map_virgin"`
	P83093Map         KeyValue `json:"p83093_map"`
	P83093MapVirgin   KeyValue `json:"p83093_map_virgin"`
	P83097Map         KeyValue `json:"p83097_map"`
	P83097MapVirgin   KeyValue `json:"p83097_map_virgin"`
	P83104Map         KeyValue `json:"p83104_map"`
	P83104MapVirgin   KeyValue `json:"p83104_map_virgin"`
	P83118Map         KeyValue `json:"p83118_map"`
	P83118MapVirgin   KeyValue `json:"p83118_map_virgin"`
	P83119Map         KeyValue `json:"p83119_map"`
	P83119MapVirgin   KeyValue `json:"p83119_map_virgin"`
	P83120Map         KeyValue `json:"p83120_map"`
	P83120MapVirgin   KeyValue `json:"p83120_map_virgin"`
	P83121            string   `json:"p83121"`
	P83122            string   `json:"p83122"`
	YearDataMonthList []struct {
		DateID                   int64  `json:"date_id"`
		Jthd                     string `json:"jthd"`
		JthdUnit                 string `json:"jthd_unit"`
		Jtyd                     string `json:"jtyd"`
		JtydUnit                 string `json:"jtyd_unit"`
		P83037                   string `json:"p83037"`
		P83073                   string `json:"p83073"`
		P83073Unit               string `json:"p83073_unit"`
		P83078                   string `json:"p83078"`
		P83078Unit               string `json:"p83078_unit"`
		P83088                   string `json:"p83088"`
		P83088Unit               string `json:"p83088_unit"`
		P83091                   string `json:"p83091"`
		P83091Unit               string `json:"p83091_unit"`
		P83098                   string `json:"p83098"`
		P83098Unit               string `json:"p83098_unit"`
		P83103                   string `json:"p83103"`
		P83103Unit               string `json:"p83103_unit"`
		P83118                   string `json:"p83118"`
		P83118Unit               string `json:"p83118_unit"`
		P83119                   string `json:"p83119"`
		P83119Unit               string `json:"p83119_unit"`
		P83120                   string `json:"p83120"`
		P83121                   string `json:"p83121"`
		P83122                   string `json:"p83122"`
		PsID                     int64  `json:"ps_id"`
		SelfConsumptionYield     string `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string `json:"self_consumption_yield_unit"`
		TimeStamp                string `json:"time_stamp"`
	} `json:"year_data_month_list"`
	ZfzyMap       KeyValue `json:"zfzy_map"`
	ZfzyMapVirgin KeyValue `json:"zfzy_map_virgin"`
	ZjzzMap       KeyValue `json:"zjzz_map"`
	ZjzzMapVirgin KeyValue `json:"zjzz_map_virgin"`
}

type TotalData struct {
	JthdMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"jthd_map"`
	JthdMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"jthd_map_virgin"`
	JtydMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"jtyd_map"`
	JtydMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"jtyd_map_virgin"`
	P83075Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83075_map"`
	P83075MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83075_map_virgin"`
	P83094Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83094_map"`
	P83094MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83094_map_virgin"`
	P83095Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83095_map"`
	P83095MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83095_map_virgin"`
	P83105Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83105_map"`
	P83105MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83105_map_virgin"`
	P83107Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83107_map"`
	P83107MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83107_map_virgin"`
	P83123Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83123_map"`
	P83123MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83123_map_virgin"`
	P83124Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83124_map"`
	P83124MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83124_map_virgin"`
	P83125    string `json:"p83125"`
	P83126    string `json:"p83126"`
	P83127Map struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83127_map"`
	P83127MapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"p83127_map_virgin"`
	TotalDataYearList []struct {
		DateID                   int64  `json:"date_id"`
		Jthd                     string `json:"jthd"`
		JthdUnit                 string `json:"jthd_unit"`
		Jtyd                     string `json:"jtyd"`
		JtydUnit                 string `json:"jtyd_unit"`
		P83038                   int64  `json:"p83038"`
		P83074                   string `json:"p83074"`
		P83074Unit               string `json:"p83074_unit"`
		P83079                   string `json:"p83079"`
		P83079Unit               string `json:"p83079_unit"`
		P83088                   string `json:"p83088"`
		P83088Unit               string `json:"p83088_unit"`
		P83093                   string `json:"p83093"`
		P83093Unit               string `json:"p83093_unit"`
		P83099                   string `json:"p83099"`
		P83099Unit               string `json:"p83099_unit"`
		P83104                   string `json:"p83104"`
		P83104Unit               string `json:"p83104_unit"`
		P83118                   string `json:"p83118"`
		P83118Unit               string `json:"p83118_unit"`
		P83119                   string `json:"p83119"`
		P83119Unit               string `json:"p83119_unit"`
		P83120                   int64  `json:"p83120"`
		P83121                   string `json:"p83121"`
		P83122                   string `json:"p83122"`
		PsID                     int64  `json:"ps_id"`
		SelfConsumptionYield     string `json:"self_consumption_yield"`
		SelfConsumptionYieldUnit string `json:"self_consumption_yield_unit"`
		TimeStamp                string `json:"time_stamp"`
	} `json:"total_data_year_list"`
	ZfzyMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zfzy_map"`
	ZfzyMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zfzy_map_virgin"`
	ZjzzMap struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zjzz_map"`
	ZjzzMapVirgin struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"zjzz_map_virgin"`
}

type KeyValue struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
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
