package showPSView

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)


const Url = "/v1/powerStationService/showPSView"


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
	BatteryPlateCom       string `json:"batteryPlateCom"`
	EnvironmentCom        string `json:"environmentCom"`
	IsPlatformDefaultUnit string `json:"is_platform_default_unit"`
	Normalization         string `json:"normalization"`
	Todayvalue            string `json:"todayvalue"`
	Todayweather          string `json:"todayweather"`
	TotalAllPower         struct {
		Unit  string `json:"unit"`
		Value string `json:"value"`
	} `json:"totalAllPower"`
	Yestedayvalue   string `json:"yestedayvalue"`
	Yestedayweather string `json:"yestedayweather"`
}
