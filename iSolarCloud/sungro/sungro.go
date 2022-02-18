package sungro

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService"
)


func (sg *SunGro) Init() error {
	for range Only.Once {
		sg.Areas = make(api.Areas)

		// sg.Areas[AliSmsService.Area.Name] = AliSmsService.Area.EndPoints
		sg.Areas[api.GetArea(AppService.Area{})] = api.AreaStruct(AppService.Init(&sg.ApiRoot))
		// sg.Areas[MttvScreenService.Area.Name] = MttvScreenService.Area.EndPoints
		// sg.Areas[PowerPointService.Area.Name] = PowerPointService.Area.EndPoints
		// sg.Areas[WebAppService.Area.Name] = WebAppService.Area.EndPoints
		// sg.Areas[WebIscmAppService.Area.Name] = WebIscmAppService.Area.EndPoints
	}

	return sg.Error
}

func (sg *SunGro) AppendUrl(endpoint string) api.EndPointUrl {
	return sg.ApiRoot.AppendUrl(endpoint)
}

func (sg *SunGro) GetEndpoint(area string, endpoint string) api.EndPoint {

	ep := sg.Areas.GetEndPoint(api.AreaName(area), api.EndPointName(endpoint))
	if sg.Auth.GetToken() != "" {
		ep = ep.SetRequest(api.RequestCommon {
			Appkey:    sg.GetAppKey(),	// sg.Auth.RequestCommon.Appkey
			Lang:      "_en_US",
			SysCode:   "200",
			Token:     sg.GetToken(),
			UserID:    sg.GetUserId(),
			ValidFlag: "1,3",
		})
	}

	return ep
}

func (sg *SunGro) ListEndpoints(area string) error {
	return sg.Areas.ListEndpoints(area)
}

func (sg *SunGro) ListAreas() {
	sg.Areas.ListAreas()
}

func (sg *SunGro) AreaExists(area string) bool {
	return sg.Areas.Exists(area)
}
func (sg *SunGro) AreaNotExists(area string) bool {
	return sg.Areas.NotExists(area)
}
