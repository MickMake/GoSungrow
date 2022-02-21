package sungro

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/sungro/AppService"
	"GoSungrow/iSolarCloud/sungro/AppService/login"
)

type SunGrow struct {
	ApiRoot api.Web
	Auth    login.EndPoint
	Areas   api.Areas
	Error   error

	// OutputString string
	// OutputArray  [][]interface{}
	// OutputType
}

func NewSunGro(url string) *SunGrow {
	var p SunGrow

	p.Error = p.ApiRoot.SetUrl(url)
	// p.OutputType = TypeHuman

	// p.Areas.Domain.ApiRoot = &p.ApiRoot

	return &p
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
