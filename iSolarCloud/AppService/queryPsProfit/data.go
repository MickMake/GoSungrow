package queryPsProfit

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/queryPsProfit"
const Disabled = false

type RequestData struct {
	PsId     string `json:"ps_id" required:"true"`
	DateID   string `json:"date_id" required:"true"`
	DateType string `json:"date_type" required:"true"`
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
	ret += api.HelpDataType()
	return ret
}


type ResultData   struct {
	ActualList []struct {
		CuspNetPowerQuantity     interface{} `json:"cusp_net_power_quantity"`
		CuspPowerQuantity        interface{} `json:"cusp_power_quantity"`
		CuspUsePowerQuantity     interface{} `json:"cusp_use_power_quantity"`
		DateID                   int64       `json:"date_id"`
		FlatNetPowerQuantity     int64       `json:"flat_net_power_quantity"`
		FlatPowerQuantity        int64       `json:"flat_power_quantity"`
		FlatUsePowerQuantity     float64     `json:"flat_use_power_quantity"`
		NetPowerProfit           float64     `json:"net_power_profit"`
		NetPowerQuantityTotal    int64       `json:"net_power_quantity_total"`
		PeakNetPowerQuantity     interface{} `json:"peak_net_power_quantity"`
		PeakPowerQuantity        interface{} `json:"peak_power_quantity"`
		PeakUsePowerQuantity     interface{} `json:"peak_use_power_quantity"`
		PowerQuantityTotal       int64       `json:"power_quantity_total"`
		SubsidyProfit            interface{} `json:"subsidy_profit"`
		TotalProfit              float64     `json:"total_profit"`
		UpdateTime               string      `json:"update_time"`
		UsePowerByDiscountProfit float64     `json:"use_power_by_discount_profit"`
		UsePowerProfit           float64     `json:"use_power_profit"`
		UsePowerQuantityTotal    float64     `json:"use_power_quantity_total"`
		ValleyNetPowerQuantity   interface{} `json:"valley_net_power_quantity"`
		ValleyPowerQuantity      interface{} `json:"valley_power_quantity"`
		ValleyUsePowerQuantity   interface{} `json:"valley_use_power_quantity"`
	} `json:"actual_list"`
	PlanList []interface{} `json:"plan_list"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
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
