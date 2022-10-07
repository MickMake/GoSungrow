package getPowerChargeSettingInfo

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPowerChargeSettingInfo"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	ParamIncomeUnit          int64  `json:"param_income_unit"`
	ParamIncomeUnitName      string `json:"param_income_unit_name"`
	PowerElectricalChargeMap struct {
		DefaultCharge      float64     `json:"default_charge"`
		IntervalTimeCharge interface{} `json:"interval_time_charge"`
	} `json:"powerElectricalChargeMap"`
	PowerSelfUseTimesChargeMap struct {
		DefaultCharge      float64 `json:"default_charge"`
		IntervalTimeCharge string  `json:"interval_time_charge"`
	} `json:"powerSelfUseTimesChargeMap"`
	PsID string `json:"ps_id"`
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
