package sungro

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService"
	"errors"
	"fmt"
	"net/url"
)


func (p *SunGro) Init() error {
	for range Only.Once {
		p.Areas = make(api.Areas)

		// p.Areas[AliSmsService.Area.Name] = AliSmsService.Area.EndPoints
		p.Areas[api.GetArea(AppService.Area{})] = api.AreaStruct(AppService.Init())
		// p.Areas[MttvScreenService.Area.Name] = MttvScreenService.Area.EndPoints
		// p.Areas[PowerPointService.Area.Name] = PowerPointService.Area.EndPoints
		// p.Areas[WebAppService.Area.Name] = WebAppService.Area.EndPoints
		// p.Areas[WebIscmAppService.Area.Name] = WebIscmAppService.Area.EndPoints
	}

	return p.Error
}

func (p *SunGro) AppendUrl(endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", p.Web.Url.String(), endpoint)
		ret, p.Error = url.Parse(endpoint)
		if p.Error != nil {
			break
		}
	}
	return ret
}


func (p *SunGro) GetEndpoint(area string, endpoint string) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		if area == "" {
			p.Error = errors.New("empty area name")
			break
		}
		if endpoint == "" {
			p.Error = errors.New("empty endpoint name")
			break
		}

		ret = p.Areas.GetEndPoint(api.AreaName(area), api.EndPointName(endpoint))
		if ret != nil {
			break
		}

		p.Error = errors.New("endpoint not found")
	}
	return ret
}

func (p *SunGro) ListEndpoints(area string) error {
	return p.Areas.ListEndpoints(area)
}

func (p *SunGro) ListAreas() {
	p.Areas.ListAreas()
}

func (p *SunGro) AreaExists(area string) bool {
	return p.Areas.Exists(area)
}
func (p *SunGro) AreaNotExists(area string) bool {
	return p.Areas.NotExists(area)
}
