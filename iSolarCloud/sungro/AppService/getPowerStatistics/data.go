package getPowerStatistics

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPowerStatistics"

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("ps_id: Can be fetched from getPsList")
	return ret
}


type ResultData struct {
	PRVlaue  string      `json:"PRVlaue"`
	City     interface{} `json:"city"`
	DayPower struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"dayPower"`
	DesignCapacity struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"design_capacity"`
	EqVlaue     string `json:"eqVlaue"`
	NowCapacity struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"nowCapacity"`
	PsName      string `json:"ps_name"`
	PsShortName string `json:"ps_short_name"`
	Status1     string `json:"status1"`
	Status2     string `json:"status2"`
	Status3     string `json:"status3"`
}
