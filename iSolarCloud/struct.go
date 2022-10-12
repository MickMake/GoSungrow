package iSolarCloud

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AliSmsService"
	"GoSungrow/iSolarCloud/AppService"
	"GoSungrow/iSolarCloud/AppService/login"
	"GoSungrow/iSolarCloud/MttvScreenService"
	"GoSungrow/iSolarCloud/NullArea"
	"GoSungrow/iSolarCloud/PowerPointService"
	"GoSungrow/iSolarCloud/WebAppService"
	"GoSungrow/iSolarCloud/WebIscmAppService"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"errors"
	"fmt"
	"strings"
	"time"
)


type SunGrow struct {
	ApiRoot api.Web
	Auth    login.EndPoint
	Areas   api.Areas
	Error   error
	NeedLogin bool

	OutputType output.OutputType
	SaveAsFile bool
}

func NewSunGro(baseUrl string, cacheDir string) *SunGrow {
	var p SunGrow

	for range Only.Once {
		p.Error = p.ApiRoot.SetUrl(baseUrl)
		if p.Error != nil {
			break
		}

		p.Error = p.ApiRoot.SetCacheDir(cacheDir)
		if p.Error != nil {
			break
		}
	}
	return &p
}

func (sg *SunGrow) Init() error {
	for range Only.Once {
		sg.Areas = make(api.Areas)

		sg.Areas[api.GetArea(NullArea.Area{})] = api.AreaStruct(NullArea.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(AliSmsService.Area{})] = api.AreaStruct(AliSmsService.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(AppService.Area{})] = api.AreaStruct(AppService.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(MttvScreenService.Area{})] = api.AreaStruct(MttvScreenService.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(PowerPointService.Area{})] = api.AreaStruct(PowerPointService.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(WebAppService.Area{})] = api.AreaStruct(WebAppService.Init(sg.ApiRoot))
		sg.Areas[api.GetArea(WebIscmAppService.Area{})] = api.AreaStruct(WebIscmAppService.Init(sg.ApiRoot))
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

		ep = sg.Areas.GetEndPoint(area, endpoint)
		if ep.IsError() {
			sg.Error = ep.GetError()
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

func (sg *SunGrow) GetByJson(endpoint string, request string) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		sg.Error = ret.GetError()
		if sg.Error != nil {
			break
		}

		if request != "" {
			ret = ret.SetRequestByJson(output.Json(request))
			sg.Error = ret.GetError()
			if sg.Error != nil {
				fmt.Println(ret.Help())
				break
			}
		}

		ret = ret.Call()
		sg.Error = ret.GetError()
		if sg.Error != nil {
			if strings.Contains(sg.Error.Error(), "er_token_login_invalid") {
				_ = sg.ApiRoot.WebCacheRemove(sg.Auth)
				_ = sg.Auth.RemoveToken()
				// _ = sg.Auth.ApiRemoveDataFile()
				break
			}
		}

		switch {
			case sg.OutputType.IsNone():
				if sg.Error != nil {
					fmt.Println(ret.Help())
					break
				}

			case sg.OutputType.IsRaw():
				// if sg.Error != nil {
				// 	fmt.Println(ret.Help())
				// 	break
				// }
				if sg.SaveAsFile {
					sg.Error = ret.WriteDataFile()
					break
				}
				fmt.Println(ret.GetJsonData(true))

			case sg.OutputType.IsJson():
				if sg.Error != nil {
					fmt.Println(ret.Help())
					break
				}
				if sg.SaveAsFile {
					sg.Error = ret.WriteDataFile()
					break
				}
				fmt.Println(ret.GetJsonData(false))

			default:
				if sg.Error != nil {
					fmt.Println(ret.Help())
					break
				}
		}
	}
	return ret
}

func (sg *SunGrow) GetByStruct(endpoint string, request interface{}, cache time.Duration) api.EndPoint {
	var ret api.EndPoint
	for range Only.Once {
		ret = sg.GetEndpoint(endpoint)
		if sg.Error != nil {
			break
		}
		if ret.IsError() {
			sg.Error = ret.GetError()
			break
		}

		if request != nil {
			ret = ret.SetRequest(request)
			if ret.IsError() {
				sg.Error = ret.GetError()
				break
			}
		}

		ret = ret.SetCacheTimeout(cache)

		ret = ret.Call()
		if !ret.IsError() {
			break
		}

		sg.Error = ret.GetError()
		if strings.Contains(sg.Error.Error(), "er_token_login_invalid") {
			_ = sg.ApiRoot.WebCacheRemove(sg.Auth)
			_ = sg.Auth.RemoveToken()
			// _ = sg.Auth.ApiRemoveDataFile()
		}
	}

	return ret
}

func (sg *SunGrow) SplitEndPoint(ae string) (api.AreaName, api.EndPointName) {
	var area api.AreaName
	var endpoint api.EndPointName

	for range Only.Once {
		s := strings.Split(ae, ".")
		switch len(s) {
			case 0:
				sg.Error = errors.New("empty endpoint")

			case 1:
				area = "AppService"
				endpoint = api.EndPointName(s[0])

			case 2:
				area = api.AreaName(s[0])
				endpoint = api.EndPointName(s[1])

			default:
				sg.Error = errors.New("too many delimiters defined, (only one '.' allowed)")
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

func (sg *SunGrow) Login(auth login.SunGrowAuth) error {
	for range Only.Once {
		a := sg.GetEndpoint(AppService.GetAreaName() + ".login")
		sg.Auth = login.Assert(a)

		sg.Error = sg.Auth.Login(&auth)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGrow) GetToken() string {
	return sg.Auth.Token()
}

func (sg *SunGrow) GetUserId() string {
	return sg.Auth.UserId()
}

func (sg *SunGrow) GetAppKey() string {
	return sg.Auth.AppKey()
}

func (sg *SunGrow) GetLastLogin() string {
	return sg.Auth.LastLogin().Format(login.DateTimeFormat)
}

func (sg *SunGrow) GetUserName() string {
	return sg.Auth.UserName()
}

func (sg *SunGrow) GetUserEmail() string {
	return sg.Auth.Email()
}

func (sg *SunGrow) HasTokenChanged() bool {
	return sg.Auth.HasTokenChanged()
}

func (sg *SunGrow) SetOutputType(outputType string) {
	sg.OutputType.Set(outputType)
}
