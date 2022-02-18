package sungro

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/AppService"
	"GoSungro/iSolarCloud/sungro/AppService/login"
)


type SunGro struct {
	ApiRoot api.Web
	Auth    login.EndPoint
	Areas   api.Areas
	Error   error

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
		a := sg.GetEndpoint(AppService.GetAreaName() + ".login")
		sg.Auth = login.Assert(a)

		sg.Error = sg.Auth.Login(&auth)
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}

func (sg *SunGro) GetToken() string {
	return sg.Auth.Token()
}

func (sg *SunGro) GetUserId() string {
	return sg.Auth.UserId()
}

func (sg *SunGro) GetAppKey() string {
	return sg.Auth.AppKey()
}

func (sg *SunGro) GetLastLogin() string {
	return sg.Auth.LastLogin().Format(login.DateTimeFormat)
}

func (sg *SunGro) GetUserName() string {
	return sg.Auth.UserName()
}

func (sg *SunGro) GetUserEmail() string {
	return sg.Auth.Email()
}

func (sg *SunGro) HasTokenChanged() bool {
	return sg.Auth.HasTokenChanged()
}
