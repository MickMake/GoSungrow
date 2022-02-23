package queryMutiPointDataList

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"fmt"
	"strings"
)

const Url = "/v1/commonService/queryMutiPointDataList"
const Disabled = false

type RequestData struct {
	PsID           int64  `json:"ps_id" required:"true"`
	PsKey          string `json:"ps_key" required:"true"`
	Points         string `json:"points" required:"true"`
	MinuteInterval string `json:"minute_interval" required:"true"`
	StartTimeStamp string `json:"start_time_stamp" required:"true"`
	EndTimeStamp   string `json:"end_time_stamp" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Points Points `json:"points"`
}

type Points map[string]Point

type Point struct {
	List []string `json:"list"`
	Unit string   `json:"unit"`
}

type DecodeResultData ResultData

type DecodePoints map[string][]string
type DecodeUnit map[string]string

func (p *ResultData) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}

		p.Points = make(Points)

		var du DecodeUnit
		// Store DeviceData.Points
		_ = json.Unmarshal(data, &du)
		for i, k := range du {
			if k == "" {
				delete(du, i)
				continue
			}
		}

		var dp DecodePoints
		// Store DeviceData.Points
		_ = json.Unmarshal(data, &dp)
		for i, k := range dp {
			if k == nil {
				delete(dp, i)
				continue
			}
			i = strings.TrimSuffix(i, "List")
			u := du[i+"_unit"]
			p.Points[i] = Point{
				List: k,
				Unit: u,
			}
		}
	}

	return err
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
