package iSolarCloud

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"errors"
	"fmt"
	"net/url"
)


type TypeEndPoint struct {
	Area AreaName     `json:"area"`
	Name EndPointName `json:"name"`
	Url  *url.URL     `json:"url"`
	ref  interface{}
}
type TypeEndPoints map[EndPointName]TypeEndPoint	// Map of EndPoints by endpoint name.
type TypeAreaNames map[AreaName][]TypeEndPoint		// Map of EndPoints by area name.

type EndPointName string
type EndPointNames []EndPointName
type AreaName string

//type ApiEndPoint string
//type ApiName string
//type ApiEndPoints map[ApiName]ApiEndPoint

//var AreaNames TypeAreaNames
//var EndPoints TypeEndPoints
func (p *SunGro) Init() error {
	for range Only.Once {
		p.AreaNames = make(TypeAreaNames)
		p.EndPoints = make(TypeEndPoints)

		for _, e := range SunGroApi {
			if len(e) != 3 {
				continue
			}
			area := AreaName(e[0])
			name := EndPointName(e[1])
			u := p.AppendUrl(e[2])

			p.EndPoints[name] = TypeEndPoint {
				Area: area,
				Name: name,
				Url:  u,
				ref:  nil,
			}

			if _, ok := p.AreaNames[area]; !ok {
				p.AreaNames[area] = []TypeEndPoint{p.EndPoints[name]}
			} else {
				p.AreaNames[area] = append(p.AreaNames[area], p.EndPoints[name])
			}
		}
	}

	for range Only.Once {
		api.Initialize(ContentTypes)
		Initialize(ContentTypes)

		println("Content Types\n---------")
		for _, ct := range api.GetContentTypes() {
			println(ct.GetSlug())
		}
		println("\nResource Types\n---------")
		for _, rt := range GetResourceTypes() {
			println(rt.GetSlug())
		}

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


func (p *SunGro) GetEndpoint(name string) *TypeEndPoint {
	var ret *TypeEndPoint
	for range Only.Once {
		if name == "" {
			p.Error = errors.New("empty endpoint name")
			break
		}
		if e, ok := p.EndPoints[EndPointName(name)]; ok {
			ret = &e
			break
		}
		p.Error = errors.New("endpoint name not found")
	}
	return ret
}

func (p *SunGro) AddEndpoint(endpoint *TypeEndPoint) error {
	for range Only.Once {
		if endpoint == nil {
			p.Error = errors.New("empty endpoint")
			break
		}

		p.EndPoints[endpoint.Name] = *endpoint
	}
	return p.Error
}
