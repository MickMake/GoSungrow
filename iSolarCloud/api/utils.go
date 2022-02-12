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

func PackageName(v interface{}) string {
	var ret string
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			ret = val.Elem().Type().PkgPath()
			break
		}

		ret = val.Type().PkgPath()
	}
	return ret
}

func GetArea(v interface{}) AreaName {
	var ret AreaName
	for range Only.Once {
		s := strings.Split(PackageName(v), "/")
		if len(s) < 2 {
			break
		}
		ret = AreaName(s[len(s)-1])
	}
	return ret
}

func GetEndPoint(v interface{}) EndPointName {
	var ret EndPointName
	for range Only.Once {
		s := strings.Split(PackageName(v), "/")
		if len(s) < 2 {
			break
		}
		ret = EndPointName(s[len(s)-1])
	}
	return ret
}

func GetUrl(u string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		var err error
		ret, err = url.Parse(u)
		if err != nil {
			ret = nil
		}
	}
	return ret
}
