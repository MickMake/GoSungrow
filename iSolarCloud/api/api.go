package api

import (
	"GoSungro/Only"
	"fmt"
	"net/url"
)


const DefaultHost     = "https://augateway.isolarcloud.com"


func CreateEndPoints(list [][]string) TypeEndPoints {
	endpoints := make(TypeEndPoints, len(list))

	for range Only.Once {
		for _, e := range list {
			if len(e) != 3 {
				continue
			}
			area := AreaName(e[0])
			name := EndPointName(e[1])
			u, _ := url.Parse(e[2])

			endpoints[name] = TypeEndPoint {
				Area:     area,
				Name:     name,
				Url:      u,
				Request:  nil,
				Response: nil,
			}

		}
	}

	return endpoints
}

func AppendUrl(host string, endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", host, endpoint)
		ret, _ = url.Parse(endpoint)
	}
	return ret
}

func (an *TypeAreaNames) GetRequest(area AreaName, endpoint EndPointName) interface{} {
	var ret interface{}

	for range Only.Once {
		err := an.Exists(area, endpoint)
		if err != nil {
			break
		}
		ret = an.GetEndPoint(area, endpoint).Request
	}

	return ret
}

func (an *TypeAreaNames) GetResponse(area AreaName, endpoint EndPointName) interface{} {
	var ret interface{}

	for range Only.Once {
		err := an.Exists(area, endpoint)
		if err != nil {
			break
		}
		ret = an.GetEndPoint(area, endpoint).Response
	}

	return ret
}

