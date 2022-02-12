package api

import (
	"GoSungro/Only"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

var OnlyOnce = "1"

func MakeNewFrom(i interface{}) (new interface{}) {
	for range OnlyOnce {
		t := reflect.TypeOf(i)
		if t == nil {
			break
		}
		new = reflect.New(t.Elem()).Interface()
	}
	return new
}

func ToSnakeCase(s string) string {
	slen := len(s)
	b := []byte(s)
	pos := make([]int, 0)
	for i := 1; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			pos = append(pos, i)
			slen++
		}
	}
	if len(pos) > 0 {
		sc := make([]byte, slen)
		n := 0
		p := 0
		for i := 0; i < len(b); i++ {
			if pos[p] == i {
				sc[n] = '_'
				n++
			}
			sc[n] = b[i]
			n++
		}
		s = string(sc)
	}
	return strings.ToLower(s)
}

func AppendUrl(host string, endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", host, endpoint)
		ret, _ = url.Parse(endpoint)
	}
	return ret
}

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
				// Resource: Resource(),
			}
		}
	}

	return endpoints
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
