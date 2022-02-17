package sungro

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/api/web"
	"GoSungro/iSolarCloud/sungro/AppService/login"
)


type SunGro struct {
	Web          web.Web
	Error error
	Areas api.Areas
	// Areas        api.Areas
	// EndPoints    api.TypeEndPoints
	// OutputString string
	// OutputArray  [][]interface{}
	// OutputType
}


func NewSunGro(url string) *SunGro {
	var p SunGro

	p.Error = p.Web.SetUrl(url)
	// p.OutputType = TypeHuman

	// p.Areas.Domain.Web = &p.Web

	return &p
}

func (p *SunGro) SetAuth(auth login.SunGroAuth) error {
	return p.Web.Auth.Login(&auth)
}

func (p *SunGro) GetToken() string {
	return p.Web.Auth.GetToken()
}

func (p *SunGro) GetTokenExpiry() string {
	return p.Web.Auth.GetTokenExpiry().Format(login.DateTimeFormat)
}

func (p *SunGro) GetUserName() string {
	return p.Web.Auth.GetUserName()
}
// func (p *SunGro) VerifyUsername(user string) string {
// 	return p.Web.Auth.VerifyUsername(user)
// }

func (p *SunGro) GetUserEmail() string {
	return p.Web.Auth.GetUserEmail()
}
// func (p *SunGro) VerifyUserEmail(user string) string {
// 	return p.Web.Auth.VerifyUserEmail(user)
// }

func (p *SunGro) HasTokenChanged() bool {
	return p.Web.Auth.HasTokenChanged()
}
