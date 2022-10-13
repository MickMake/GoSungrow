package getPsReport

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"
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
		Co2Reduce                      api.Float     `json:"co2_reduce"`
		Co2ReduceMap                   api.UnitValue `json:"co2_reduce_map"`
		Co2ReduceOriginalUnit          api.String        `json:"co2_reduce_original_unit"`
		PowerQuantityTotal             api.Float     `json:"power_quantity_total"`
		PowerQuantityTotalMap          api.UnitValue `json:"power_quantity_total_map"`
		PowerQuantityTotalOriginalUnit api.String        `json:"power_quantity_total_original_unit"`
		ProfitList                     ProfitList    `json:"profit_list"`
		TotalProfit                    api.Float     `json:"total_profit"`
	} `json:"ps_date_range_total_data"`
	PsTotalData struct {
		Co2Reduce                      api.Float     `json:"co2_reduce"`
		Co2ReduceMap                   api.UnitValue `json:"co2_reduce_map"`
		Co2ReduceOriginalUnit          api.String        `json:"co2_reduce_original_unit"`
		MinDateID                      api.DateTime   `json:"min_date_id"`
		PowerQuantityTotal             api.Float     `json:"power_quantity_total"`
		PowerQuantityTotalMap          api.UnitValue `json:"power_quantity_total_map"`
		PowerQuantityTotalOriginalUnit api.String        `json:"power_quantity_total_original_unit"`
		ProfitList                     ProfitList    `json:"profit_list"`
		TotalProfit                    api.Float     `json:"total_profit"`
	} `json:"ps_total_data"`
	ReportListData struct {
		Co2ReduceUnit api.String `json:"co2_reduce_unit"`
		PageList      []struct {
			Co2Reduce          api.Float  `json:"co2_reduce"`
			IncomeUnitName     api.String     `json:"income_unit_name"`
			PowerQuantityTotal api.Float  `json:"power_quantity_total"`
			PsID               api.Integer`json:"ps_id"`
			PsInstalledPower   api.Float  `json:"ps_installed_power"`
			PsName             api.String     `json:"ps_name"`
			PsTotalCo2Reduce   api.Float  `json:"ps_total_co2_reduce"`
			PsTotalPower       api.Float  `json:"ps_total_power"`
			PsType             api.Integer`json:"ps_type"`
			ShareType          string     `json:"share_type"`
			ProfitList         ProfitList `json:"profit_list"`
			TotalProfit        api.Float  `json:"total_profit"`
		} `json:"pageList"`
		PowerQuantityTotalUnit api.String `json:"power_quantity_total_unit"`
		PsInstalledPowerUnit   api.String `json:"ps_installed_power_unit"`
		PsTotalCo2ReduceUnit   api.String `json:"ps_total_co2_reduce_unit"`
		PsTotalPowerUnit       api.String `json:"ps_total_power_unit"`
		RowCount               api.Integer  `json:"rowCount"`
	} `json:"report_list_data"`
}

type ProfitList []struct {
	IncomeUnitID   api.Integer   `json:"income_unit_id"`
	IncomeUnitName api.String  `json:"income_unit_name"`
	TotalProfit    api.Float `json:"total_profit"`
}

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

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
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
// }

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}
