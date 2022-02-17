package api

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api/apiReflect"
	"encoding/json"
	"errors"
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


// func PackageName(v interface{}) string {
// 	var ret string
// 	for range Only.Once {
// 		if v == nil {
// 			break
// 		}
//
// 		val := reflect.ValueOf(v)
// 		if val.Kind() == reflect.Ptr {
// 			ret = val.Elem().Type().PkgPath()
// 			break
// 		}
//
// 		ret2 := val.Type().Name()
// 		ret3 := val.Type().String()
// 		ret = val.Type().PkgPath()
// 		ret = strings.TrimPrefix(ret, thisPackagePath)
//
// 		fmt.Printf("%s\t%s\t%s\n", ret, ret2, ret3)
// 		fmt.Println("")
// 	}
// 	return ret
// }

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

func GetArea(v interface{}) AreaName {
	return AreaName(apiReflect.GetArea(thisPackagePath, v))
}

func GetName(v interface{}) EndPointName {
	return EndPointName(apiReflect.GetName(thisPackagePath, v))
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


// Response checks
func CheckResultCode(rc string) error {
	var err error
	for range Only.Once {
		switch rc {
			case "1":
				err = nil
			case "-1":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "010":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "000":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "201":
				err = errors.New(fmt.Sprintf("error '%s'", rc))
			case "E00003":
				err = errors.New(fmt.Sprintf("need to login again '%s'", rc))
			default:
				err = errors.New(fmt.Sprintf("unknown error '%s'", rc))
		}
	}
	return err
}

func CheckResultMessage(rc string) error {
	var err error
	for range Only.Once {
		switch {
			case rc == "success":
				err = nil
			case rc == "er_token_login_invalid":
				err = errors.New(fmt.Sprintf("need to login again '%s'", rc))
			case rc == "er_parameter_value_invalid":
				err = errors.New(fmt.Sprintf("incorrect request data '%s'", rc))
			case rc == "er_unknown_exception":
				err = errors.New(fmt.Sprintf("API error '%s'", rc))
			case strings.HasPrefix(rc, "Parameter:"):
				err = errors.New(fmt.Sprintf("incorrect request data '%s'", rc))
			// case err == nil:
			// 	err = errors.New(fmt.Sprintf("unknown '%s'", rc))
			default:
				err = errors.New(fmt.Sprintf("unknown error '%s'", rc))
		}
	}
	return err
}


// Request checks
func CheckString(name string, rc string) error {
	var err error
	for range Only.Once {
		if rc == "" {
			err = errors.New(name + ": empty string")
			break
		}
		if strings.TrimSpace(rc) == "" {
			err = errors.New(name + ": empty string with spaces")
			break
		}
	}
	return err
}

func (req RequestCommon) IsValid() error {
	var err error
	for range Only.Once {
		err = CheckString("Appkey", req.Appkey)
		if err != nil {
			break
		}
		err = CheckString("Lang", req.Lang)
		if err != nil {
			break
		}
		err = CheckString("SysCode", req.SysCode)
		if err != nil {
			break
		}
		err = CheckString("Auth", req.Token)
		if err != nil {
			break
		}
		err = CheckString("UserID", req.UserID)
		if err != nil {
			break
		}
		err = CheckString("ValidFlag", req.ValidFlag)
		if err != nil {
			break
		}
	}
	return err
}

func (req ResponseCommon) IsValid() error {
	var err error
	for range Only.Once {
		err = CheckResultMessage(req.ResultMsg)
		if err != nil {
			break
		}
		err = CheckResultCode(req.ResultCode)
		if err != nil {
			break
		}
		err = CheckString("ReqSerialNum", req.ReqSerialNum)
		if err != nil {
			break
		}
		// if req.ResultData == nil {
		// 	err = errors.New("zero results")
		// 	break
		// }
	}
	return err
}


// func CheckEmpty(rc string) error {
// 	var err error
// 	for range Only.Once {
// 		// switch rc {
// 		// 	case "1":
// 		// 		err = nil
// 		// 	case "-1":
// 		// 		err = errors.New(fmt.Sprintf("error '%s'", rc))
// 		// 	case "010":
// 		// 		err = errors.New(fmt.Sprintf("error '%s'", rc))
// 		// 	case "000":
// 		// 		err = errors.New(fmt.Sprintf("error '%s'", rc))
// 		// 	case "201":
// 		// 		err = errors.New(fmt.Sprintf("error '%s'", rc))
// 		// 	case "E00003":
// 		// 		err = errors.New(fmt.Sprintf("need to login again '%s'", rc))
// 		// 	default:
// 		// 		err = errors.New(fmt.Sprintf("unknown error '%s'", rc))
// 		// }
// 	}
// 	return err
// }

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
