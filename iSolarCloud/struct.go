package iSolarCloud

import (
	"GoSungro/iSolarCloud/api"
	"GoSungro/iSolarCloud/domain"
	"GoSungro/iSolarCloud/web"
)


type SunGro struct {
	Web          web.Web
	Error        error
	Areas        Areas
	AreaNames    api.TypeAreaNames
	// EndPoints    api.TypeEndPoints
	OutputString string
	OutputArray  [][]interface{}
	OutputType
}

type Areas struct {
	domain.Domain
}

const (
	TypeGit    = iota
	TypeJson   = iota
	TypeHuman  = iota
	TypeGoogle = iota

	StringTypeGit = "git"
	StringTypeJson = "json"
	StringTypeHuman = "human"
	StringTypeGoogle = "google"
)


func NewSunGro(url string) *SunGro {
	var p SunGro

	p.Error = p.Web.SetUrl(url)
	p.OutputType = TypeHuman

	p.Areas.Domain.Web = &p.Web

	return &p
}

func (p *SunGro) SetAuth(auth web.SunGroAuth) error {
	return p.Web.Login(&auth)
}

func (p *SunGro) SetOutput(outType string) error {

	switch {
		case p.IsStrGit(outType):
			p.OutputType = TypeGit
		case p.IsStrJson(outType):
			p.OutputType = TypeJson
		case p.IsStrHuman(outType):
			p.OutputType = TypeHuman
	}

	return p.Error
}

func (p *SunGro) SetOutputString(f string)  {
	p.OutputString = f
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


type OutputType int
func (out *OutputType) IsGit() bool {
	if *out == TypeGit {
		return true
	}
	return false
}
func (out *OutputType) IsJson() bool {
	if *out == TypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsHuman() bool {
	if *out == TypeHuman {
		return true
	}
	return false
}
func (out *OutputType) IsGoogle() bool {
	if *out == TypeGoogle {
		return true
	}
	return false
}

func (out *OutputType) IsStrGit(t string) bool {
	if t == StringTypeGit {
		return true
	}
	return false
}
func (out *OutputType) IsStrJson(t string) bool {
	if t == StringTypeJson {
		return true
	}
	return false
}
func (out *OutputType) IsStrHuman(t string) bool {
	if t == StringTypeHuman {
		return true
	}
	return false
}
func (out *OutputType) IsStrGoogle(t string) bool {
	if t == StringTypeGoogle {
		return true
	}
	return false
}
