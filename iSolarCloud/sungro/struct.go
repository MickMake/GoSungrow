package sungro

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService"
	"GoSungro/iSolarCloud/sungro/AppService/login"
)


type SunGro struct {
	ApiRoot        api.Web
	// Auth           api.EndPoint
	Auth  login.EndPoint
	Error error

	Areas          api.Areas
	ApiTokenExpiry string

	// OutputString string
	// OutputArray  [][]interface{}
	// OutputType
}


func NewSunGro(url string) *SunGro {
	var p SunGro

	p.Error = p.ApiRoot.SetUrl(url)
	// p.OutputType = TypeHuman

	// p.Areas.Domain.ApiRoot = &p.ApiRoot

	return &p
}

func (sg *SunGro) Login(auth login.SunGroAuth) error {
	for range Only.Once {
		// sg.Auth = sg.GetEndpoint(AppService.GetAreaName(), "login")
		// sg.Auth = sg.Auth.SetRequest(login.RequestCommon {
		// 	Appkey:       auth.AppKey,
		// 	SysCode:      "600",
		// 	UserAccount:  auth.UserAccount,
		// 	UserPassword: auth.UserPassword,
		// })
		//
		// sg.Auth = sg.Auth.Call()
		// sg.Error = sg.Auth.GetError()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// r := sg.ApiRoot.Get(sg.Auth)

		a := sg.GetEndpoint(AppService.GetAreaName(), "login")
		sg.Auth = login.Assert(a)

		sg.Error = sg.Auth.Login(&auth)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGro) GetToken() string {
	return sg.Auth.GetToken()
}

func (sg *SunGro) GetUserId() string {
	return sg.Auth.GetUserId()
}

func (sg *SunGro) GetAppKey() string {
	return sg.Auth.GetAppKey()
}

func (sg *SunGro) GetTokenExpiry() string {
	return sg.Auth.GetTokenExpiry().Format(login.DateTimeFormat)
}

func (sg *SunGro) GetUserName() string {
	return sg.Auth.GetUserName()
}

func (sg *SunGro) GetUserEmail() string {
	return sg.Auth.GetUserEmail()
}

func (sg *SunGro) HasTokenChanged() bool {
	return sg.Auth.HasTokenChanged()
}
