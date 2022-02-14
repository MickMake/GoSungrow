package sungro

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/sungro/web"
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

func (p *SunGro) SetAuth(auth web.SunGroAuth) error {
	return p.Web.Login(&auth)
}

func (p *SunGro) GetToken() string {
	return p.Web.GetToken()
}

func (p *SunGro) GetTokenExpiry() string {
	return p.Web.GetTokenExpiry().Format(web.DateTimeFormat)
}

func (p *SunGro) GetDomain() string {
	return p.Web.GetDomain()
}
func (p *SunGro) VerifyDomain(domain string) string {
	return p.Web.VerifyDomain(domain)
}

func (p *SunGro) GetUser() string {
	return p.Web.GetUser()
}
func (p *SunGro) VerifyUser(user string) string {
	return p.Web.VerifyUser(user)
}

func (p *SunGro) GetUserMac() string {
	return p.Web.GetUser()
}
func (p *SunGro) VerifyUserMac(user string) string {
	return p.Web.VerifyUser(user)
}

func (p *SunGro) GetUsername() string {
	return p.Web.GetUsername()
}
func (p *SunGro) VerifyUsername(user string) string {
	return p.Web.VerifyUsername(user)
}

func (p *SunGro) GetUserEmail() string {
	return p.Web.GetUserEmail()
}
func (p *SunGro) VerifyUserEmail(user string) string {
	return p.Web.VerifyUserEmail(user)
}

func (p *SunGro) GetDisplayName() string {
	return p.Web.GetDisplayName()
}
func (p *SunGro) VerifyDisplayName(user string) string {
	return p.Web.VerifyDisplayName(user)
}

func (p *SunGro) HasTokenChanged() bool {
	return p.Web.HasTokenChanged()
}
