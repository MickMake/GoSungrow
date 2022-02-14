package api

import (
	"GoSungro/Only"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)


// This is used to trim the sub-packages imported under the API.
var thisPackagePath string
func init() {
	for range Only.Once {
		val := reflect.ValueOf(Api{})
		if val.Kind() == reflect.Ptr {
			thisPackagePath = val.Elem().Type().PkgPath()
			break
		}
		thisPackagePath = strings.TrimSuffix(val.Type().PkgPath(), "api")
	}
}


func AppendUrl(host string, endpoint string) *url.URL {
	var ret *url.URL
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", host, endpoint)
		ret, _ = url.Parse(endpoint)
	}
	return ret
}

// func CreateEndPoints(list [][]string) TypeEndPoints {
// 	endpoints := make(TypeEndPoints, len(list))
//
// 	for range Only.Once {
// 		for _, e := range list {
// 			if len(e) != 3 {
// 				continue
// 			}
// 			area := AreaName(e[0])
// 			name := EndPointName(e[1])
// 			u, _ := url.Parse(e[2])
//
// 			// endpoints[name] = EndPointStruct {
// 			// 	Area:     area,
// 			// 	Name:     name,
// 			// 	Url:      u,
// 			// 	Request:  nil,
// 			// 	Response: nil,
// 			// 	// EndPoint: EndPoint(),
// 			// }
// 			fmt.Printf("EndPoint: %s\t%s\t%s\n", area, name, u)
//
// 			if name == "getPowerDevicePointNames" {
// 				fmt.Println("HEYHEYHEY")
// 				point := endpoints[name]
// 				fmt.Printf("EndPoint: %s\t%s\t%s\n",
// 					point.GetArea(),
// 					point.GetName(),
// 					point.GetUrl(),
// 					)
//
// 				// point.SetResource(getPowerDevicePointNames.Init())
// 				// fmt.Printf("EndPoint: %v\n", point.GetResource())
// 				continue
// 			}
// 		}
// 	}
//
// 	return endpoints
// }

// func getType(myvar interface{}) string {
// 	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
// 		return "*" + t.Elem().Name()
// 	} else {
// 		return t.Name()
// 	}
// }

// func SetEndPointReference(endpoint string, ref interface{}) *EndPointStruct {
// 	var endpoint EndPointStruct
//
// 	for range Only.Once {
// 		for _, e := range list {
// 			if len(e) != 3 {
// 				continue
// 			}
// 			area := AreaName(e[0])
// 			name := EndPointName(e[1])
// 			u, _ := url.Parse(e[2])
//
// 			endpoint = EndPointStruct {
// 				Area:     area,
// 				Name:     name,
// 				Url:      u,
// 				Request:  nil,
// 				Response: nil,
// 				// EndPoint: EndPoint(),
// 			}
//
// 			// if name == "getPowerDevicePointNames" {
// 			// 	fmt.Println("HEYHEYHEY")
// 			// 	point := endpoints[name]
// 			// 	point.SetResource(getPowerDevicePointNames.Init())
// 			// 	fmt.Printf("EndPoint: %v\n", point.GetResource())
// 			// 	continue
// 			// }
// 		}
// 	}
//
// 	return &endpoint
// }

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

		ret2 := val.Type().Name()
		ret3 := val.Type().String()
		ret = val.Type().PkgPath()
		ret = strings.TrimPrefix(ret, thisPackagePath)

		fmt.Printf("%s\t%s\t%s\n", ret, ret2, ret3)
		fmt.Println("")
	}
	return ret
}

func GetArea(v interface{}) AreaName {
	var ret AreaName
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret2 := val.Type().Name()
		ret3 := val.Type().String()
		fmt.Printf("%s\t%s\t%s\n", ret1, ret2, ret3)

		if strings.HasSuffix(ret3, "Area") {
			ret = AreaName(strings.TrimSuffix(ret3, ".Area"))
			break
		}

		if strings.HasSuffix(ret3, "EndPoint") {
			s := strings.Split(ret1, "/")
			ret = AreaName(s[len(s)-2])
			break
		}

		ret = AreaName(strings.TrimPrefix(ret1, thisPackagePath))

		fmt.Println("")
	}
	return ret
}

func GetName(v interface{}) EndPointName {
	var ret EndPointName
	for range Only.Once {
		if v == nil {
			break
		}

		val := reflect.ValueOf(v)
		ret1 := val.Type().PkgPath()
		ret2 := val.Type().Name()
		ret3 := val.Type().String()
		fmt.Printf("%s\t%s\t%s\n", ret1, ret2, ret3)

		if strings.HasSuffix(ret3, "Area") {
			ret = EndPointName(strings.TrimSuffix(ret3, ".Area"))
			break
		}

		if strings.HasSuffix(ret3, "EndPoint") {
			ret = EndPointName(strings.TrimSuffix(ret3, ".EndPoint"))
			break
		}

		ret = EndPointName(strings.TrimPrefix(ret1, thisPackagePath))

		fmt.Println("")
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


func GetAsJson(r interface{}) Json {
	var ret Json
	for range Only.Once {
		j, err := json.Marshal(r)
		if err != nil {
			break
		}
		ret = Json(j)
	}
	return ret
}


// func MakeNewFrom(i interface{}) (new interface{}) {
// 	for range OnlyOnce {
// 		t := reflect.TypeOf(i)
// 		if t == nil {
// 			break
// 		}
// 		new = reflect.New(t.Elem()).Interface()
// 	}
// 	return new
// }
//
// func ToSnakeCase(s string) string {
// 	slen := len(s)
// 	b := []byte(s)
// 	pos := make([]int, 0)
// 	for i := 1; i < len(b); i++ {
// 		if b[i] >= 'A' && b[i] <= 'Z' {
// 			pos = append(pos, i)
// 			slen++
// 		}
// 	}
// 	if len(pos) > 0 {
// 		sc := make([]byte, slen)
// 		n := 0
// 		p := 0
// 		for i := 0; i < len(b); i++ {
// 			if pos[p] == i {
// 				sc[n] = '_'
// 				n++
// 			}
// 			sc[n] = b[i]
// 			n++
// 		}
// 		s = string(sc)
// 	}
// 	return strings.ToLower(s)
// }
