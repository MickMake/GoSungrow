package iSolarCloud

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/AliSmsService"
	"GoSungro/iSolarCloud/AppService"
	"GoSungro/iSolarCloud/MttvScreenService"
	"GoSungro/iSolarCloud/PowerPointService"
	"GoSungro/iSolarCloud/WebAppService"
	"GoSungro/iSolarCloud/WebIscmAppService"
	"GoSungro/iSolarCloud/api"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/url"
	"os"
)


//var AreaNames TypeAreaNames
//var EndPoints TypeEndPoints
func (p *SunGro) Init() error {
	for range Only.Once {
		p.AreaNames = make(api.TypeAreaNames)
		// p.EndPoints = make(api.TypeEndPoints)

		p.AreaNames[AliSmsService.Area.Name] = AliSmsService.Area.EndPoints
		p.AreaNames[AppService.Area.Name] = AppService.Area.EndPoints
		p.AreaNames[MttvScreenService.Area.Name] = MttvScreenService.Area.EndPoints
		p.AreaNames[PowerPointService.Area.Name] = PowerPointService.Area.EndPoints
		p.AreaNames[WebAppService.Area.Name] = WebAppService.Area.EndPoints
		p.AreaNames[WebIscmAppService.Area.Name] = WebIscmAppService.Area.EndPoints

		// foo := p.AreaNames[AppService.Area.Name]
		// fmt.Printf("[%s] => %v", AppService.Area.Name, foo.SortEndPoints())

		// for _, e := range api.SunGro {
		// 	if len(e) != 3 {
		// 		continue
		// 	}
		// 	area := api.AreaName(e[0])
		// 	name := api.EndPointName(e[1])
		// 	u := p.AppendUrl(e[2])
		//
		// 	if _, ok := p.AreaNames[area]; !ok {
		// 		p.AreaNames[area] = make(api.TypeEndPoints)
		// 	}
		//
		// 	p.AreaNames[area][name] = api.TypeEndPoint {
		// 		Area:     area,
		// 		Name:     name,
		// 		Url:      u,
		// 		Request:  nil,
		// 		Response: nil,
		// 	}
		//
		// }
	}

	// for range Only.Once {
	// 	api.Initialize(ContentTypes)
	// 	Initialize(ContentTypes)
	//
	// 	println("Content Types\n---------")
	// 	for _, ct := range api.GetContentTypes() {
	// 		println(ct.GetSlug())
	// 	}
	// 	println("\nResource Types\n---------")
	// 	for _, rt := range GetResourceTypes() {
	// 		println(rt.GetSlug())
	// 	}
	//
	// }

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


func (p *SunGro) GetEndpoint(area string, endpoint string) *api.TypeEndPoint {
	var ret *api.TypeEndPoint
	for range Only.Once {
		if area == "" {
			p.Error = errors.New("empty area name")
			break
		}
		if endpoint == "" {
			p.Error = errors.New("empty endpoint name")
			break
		}

		ret = p.AreaNames.GetEndPoint(api.AreaName(area), api.EndPointName(endpoint))
		if ret != nil {
			break
		}

		p.Error = errors.New("endpoint not found")
	}
	return ret
}

// func (p *SunGro) AddEndpoint(endpoint *api.TypeEndPoint) error {
// 	for range Only.Once {
// 		if endpoint == nil {
// 			p.Error = errors.New("empty endpoint")
// 			break
// 		}
//
// 		p.EndPoints[endpoint.Name] = *endpoint
// 	}
// 	return p.Error
// }

func (p *SunGro) ListEndpoints(area string) error {
	for range Only.Once {
		if area == "" {
			fmt.Printf("Listing all endpoints from all areas:\n")
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Area", "Resource", "Url"})
			table.SetBorder(true)
			for _, area := range p.AreaNames.SortAreas() {
				endpoints := p.AreaNames[area]
				for _, endpoint := range endpoints.SortEndPoints() {
					table.Append([]string{
						string(p.AreaNames[area][api.EndPointName(endpoint)].Area),
						string(p.AreaNames[area][api.EndPointName(endpoint)].Name),
						p.AreaNames[area][api.EndPointName(endpoint)].Url.String(),
					})
				}
			}
			table.Render()
			break
		}

		if p.AreaNotExists(area) {
			p.Error = errors.New("unknown area name")
			break
		}

		fmt.Printf("Listing all endpoints from area '%s':\n", area)
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Resource", "Url"})
		table.SetBorder(true)
		endpoints := p.AreaNames[api.AreaName(area)]
		for _, endpoint := range endpoints.SortEndPoints() {
			table.Append([]string{
				string(p.AreaNames[api.AreaName(area)][api.EndPointName(endpoint)].Name),
				p.AreaNames[api.AreaName(area)][api.EndPointName(endpoint)].Url.String(),
			})
		}
		table.Render()
	}

	return p.Error
}

func (p *SunGro) ListAreas() error {
	for range Only.Once {
		fmt.Println("Listing all endpoint areas:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Areas", "EndPoints"})
		table.SetBorder(true)

		for _, area := range p.AreaNames.SortAreas() {
			size := fmt.Sprintf("%d", len(p.AreaNames[api.AreaName(area)]))
			table.Append([]string{string(area), size})
		}

		table.Render()
	}
	return p.Error
}


func (p *SunGro) AreaExists(area string) bool {
	var ok bool
	_, ok = p.AreaNames[api.AreaName(area)]
	return ok
}
func (p *SunGro) AreaNotExists(area string) bool {
	return !p.AreaExists(area)
}

// func (p *SunGro) EndPointExists(area string, endpoint string) bool {
// 	var ok bool
// 	_, ok = p.EndPoints[api.EndPointName(endpoint)]
// 	return ok
// }
// func (p *SunGro) EndPointNotExists(area string, endpoint string) bool {
// 	return !p.EndPointExists(area, endpoint)
// }
