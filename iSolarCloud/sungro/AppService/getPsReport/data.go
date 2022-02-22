package getPsReport

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/reportService/getPsReport"
const Disabled = false

type RequestData struct {
	ReportType string `json:"report_type" required:"true"`
	DateID     string `json:"date_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	PsDateRangeTotalData struct {
		Co2Reduce    float64 `json:"co2_reduce"`
		Co2ReduceMap struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"co2_reduce_map"`
		Co2ReduceOriginalUnit string  `json:"co2_reduce_original_unit"`
		PowerQuantityTotal    float64 `json:"power_quantity_total"`
		PowerQuantityTotalMap struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"power_quantity_total_map"`
		PowerQuantityTotalOriginalUnit string `json:"power_quantity_total_original_unit"`
		ProfitList                     []struct {
			IncomeUnitID   int64   `json:"income_unit_id"`
			IncomeUnitName string  `json:"income_unit_name"`
			TotalProfit    float64 `json:"total_profit"`
		} `json:"profit_list"`
		TotalProfit float64 `json:"total_profit"`
	} `json:"ps_date_range_total_data"`
	PsTotalData struct {
		Co2Reduce    float64 `json:"co2_reduce"`
		Co2ReduceMap struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"co2_reduce_map"`
		Co2ReduceOriginalUnit string  `json:"co2_reduce_original_unit"`
		MinDateID             int64   `json:"min_date_id"`
		PowerQuantityTotal    float64 `json:"power_quantity_total"`
		PowerQuantityTotalMap struct {
			Unit  string `json:"unit"`
			Value string `json:"value"`
		} `json:"power_quantity_total_map"`
		PowerQuantityTotalOriginalUnit string `json:"power_quantity_total_original_unit"`
		ProfitList                     []struct {
			IncomeUnitID   int64   `json:"income_unit_id"`
			IncomeUnitName string  `json:"income_unit_name"`
			TotalProfit    float64 `json:"total_profit"`
		} `json:"profit_list"`
		TotalProfit float64 `json:"total_profit"`
	} `json:"ps_total_data"`
	ReportListData struct {
		Co2ReduceUnit string `json:"co2_reduce_unit"`
		PageList      []struct {
			Co2Reduce          float64 `json:"co2_reduce"`
			IncomeUnitName     string  `json:"income_unit_name"`
			PowerQuantityTotal float64 `json:"power_quantity_total"`
			PsID               int64   `json:"ps_id"`
			PsInstalledPower   float64 `json:"ps_installed_power"`
			PsName             string  `json:"ps_name"`
			PsTotalCo2Reduce   float64 `json:"ps_total_co2_reduce"`
			PsTotalPower       float64 `json:"ps_total_power"`
			PsType             int64   `json:"ps_type"`
			ShareType          string  `json:"share_type"`
			TotalProfit        float64 `json:"total_profit"`
		} `json:"pageList"`
		PowerQuantityTotalUnit string `json:"power_quantity_total_unit"`
		PsInstalledPowerUnit   string `json:"ps_installed_power_unit"`
		PsTotalCo2ReduceUnit   string `json:"ps_total_co2_reduce_unit"`
		PsTotalPowerUnit       string `json:"ps_total_power_unit"`
		RowCount               int64  `json:"rowCount"`
	} `json:"report_list_data"`
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
