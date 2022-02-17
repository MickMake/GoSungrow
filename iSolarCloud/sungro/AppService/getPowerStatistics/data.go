package getPowerStatistics

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"errors"
)

const Url = "/v1/powerStationService/getPowerStatistics"

type RequestData struct {
	PsId string `json:"ps_id"`
}

func (rd *RequestData) IsValid() error {
	var err error
	for range Only.Once {
		if rd == nil {
			err = errors.New("empty device type")
			break
		}
		err = api.CheckString("PsId", rd.PsId)
		if err != nil {
			err = errors.New("empty device type")
			break
		}
	}
	return err
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
