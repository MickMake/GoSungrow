package sungro

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/sungro/AliSmsService"
	"GoSungrow/iSolarCloud/sungro/AppService"
	"GoSungrow/iSolarCloud/sungro/MttvScreenService"
	"GoSungrow/iSolarCloud/sungro/PowerPointService"
	"GoSungrow/iSolarCloud/sungro/WebAppService"
	"GoSungrow/iSolarCloud/sungro/WebIscmAppService"
	"errors"
	"strings"
)

func (sg *SunGrow) Init() error {
	for range Only.Once {
		sg.Areas = make(api.Areas)

		sg.Areas[api.GetArea(AliSmsService.Area{})] = api.AreaStruct(AliSmsService.Init(&sg.ApiRoot))
		sg.Areas[api.GetArea(AppService.Area{})] = api.AreaStruct(AppService.Init(&sg.ApiRoot))
		sg.Areas[api.GetArea(MttvScreenService.Area{})] = api.AreaStruct(MttvScreenService.Init(&sg.ApiRoot))
		sg.Areas[api.GetArea(PowerPointService.Area{})] = api.AreaStruct(PowerPointService.Init(&sg.ApiRoot))
		sg.Areas[api.GetArea(WebAppService.Area{})] = api.AreaStruct(WebAppService.Init(&sg.ApiRoot))
		sg.Areas[api.GetArea(WebIscmAppService.Area{})] = api.AreaStruct(WebIscmAppService.Init(&sg.ApiRoot))
	}

	return sg.Error
}

func (sg *SunGrow) AppendUrl(endpoint string) api.EndPointUrl {
	return sg.ApiRoot.AppendUrl(endpoint)
}

func (sg *SunGrow) GetEndpoint(ae string) api.EndPoint {
	var ep api.EndPoint
	for range Only.Once {
		area, endpoint := sg.SplitEndPoint(ae)
		if sg.Error != nil {
			break
		}

		ep = sg.Areas.GetEndPoint(api.AreaName(area), api.EndPointName(endpoint))
		if ep == nil {
			sg.Error = errors.New("EndPoint not found")
			break
		}

		if ep.IsDisabled() {
			sg.Error = errors.New("API EndPoint is not implemented")
			break
		}

		if sg.Auth.Token() != "" {
			ep = ep.SetRequest(api.RequestCommon{
				Appkey:    sg.GetAppKey(), // sg.Auth.RequestCommon.Appkey
				Lang:      "_en_US",
				SysCode:   "200",
				Token:     sg.GetToken(),
				UserID:    sg.GetUserId(),
				ValidFlag: "1,3",
			})
		}
	}
	return ep
}

func (sg *SunGrow) SplitEndPoint(ae string) (string, string) {
	var area string
	var endpoint string

	for range Only.Once {
		s := strings.Split(ae, ".")
		switch len(s) {
		case 0:
			sg.Error = errors.New("empty endpoint")

		case 1:
			area = "AppService"
			endpoint = s[0]

		case 2:
			area = s[0]
			endpoint = s[1]

		default:
			sg.Error = errors.New("too many delimeters defined, (only one '.' allowed)")
		}
	}

	return area, endpoint
}

func (sg *SunGrow) ListEndpoints(area string) error {
	return sg.Areas.ListEndpoints(area)
}

func (sg *SunGrow) ListAreas() {
	sg.Areas.ListAreas()
}

func (sg *SunGrow) AreaExists(area string) bool {
	return sg.Areas.Exists(area)
}
func (sg *SunGrow) AreaNotExists(area string) bool {
	return sg.Areas.NotExists(area)
}
