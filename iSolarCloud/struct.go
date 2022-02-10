package iSolarCloud

import (
	//"GoSungro/iSolarCloud/agent"
	//"GoSungro/iSolarCloud/agentLog"
	//"GoSungro/iSolarCloud/answerRule"
	//"GoSungro/iSolarCloud/audio"
	//"GoSungro/iSolarCloud/call"
	//"GoSungro/iSolarCloud/callQueue"
	//"GoSungro/iSolarCloud/callQueueEmailReport"
	//"GoSungro/iSolarCloud/callQueueStat"
	//"GoSungro/iSolarCloud/callerIdEmergency"
	//"GoSungro/iSolarCloud/cdr2"
	//"GoSungro/iSolarCloud/cdrExport"
	//"GoSungro/iSolarCloud/cdrSchedule"
	//"GoSungro/iSolarCloud/chart"
	//"GoSungro/iSolarCloud/conference"
	//"GoSungro/iSolarCloud/conferenceCdr"
	//"GoSungro/iSolarCloud/connection"
	//"GoSungro/iSolarCloud/contact"
	//"GoSungro/iSolarCloud/dashboard"
	//"GoSungro/iSolarCloud/department"
	//"GoSungro/iSolarCloud/device"
	//"GoSungro/iSolarCloud/deviceModel"
	//"GoSungro/iSolarCloud/deviceProfile"
	//"GoSungro/iSolarCloud/dialPlan"
	//"GoSungro/iSolarCloud/dialPolicy"
	//"GoSungro/iSolarCloud/dialRule"
	//"GoSungro/iSolarCloud/domain"
	//"GoSungro/iSolarCloud/image"
	//"GoSungro/iSolarCloud/mac"
	//"GoSungro/iSolarCloud/meeting"
	//"GoSungro/iSolarCloud/messageSession"
	//"GoSungro/iSolarCloud/ndpServer"
	//"GoSungro/iSolarCloud/participant"
	//"GoSungro/iSolarCloud/permission"
	//"GoSungro/iSolarCloud/phoneConfiguration"
	//"GoSungro/iSolarCloud/phoneNumber"
	//"GoSungro/iSolarCloud/presence"
	//"GoSungro/iSolarCloud/queued"
	//"GoSungro/iSolarCloud/quota"
	//"GoSungro/iSolarCloud/recording"
	//"GoSungro/iSolarCloud/routeConnection"
	//"GoSungro/iSolarCloud/site"
	//"GoSungro/iSolarCloud/smsNumber"
	//"GoSungro/iSolarCloud/statistics"
	//"GoSungro/iSolarCloud/subscriber"
	//"GoSungro/iSolarCloud/timeframe"
	"GoSungro/iSolarCloud/domain"
	"GoSungro/iSolarCloud/web"
)


type SunGro struct {
	Web          web.Web
	Error        error
	Areas        Areas
	AreaNames    TypeAreaNames
	EndPoints    TypeEndPoints
	OutputString string
	OutputArray  [][]interface{}
	OutputType
}

type Areas struct {
	domain.Domain
	//site.Site
	//subscriber.Subscriber
	//department.Department
	//presence.Presence
	//device.Device
	//deviceModel.Model
	//deviceProfile.Profile
	//ndpServer.Ndp
	//contact.Contact
	//mac.Mac
	//chart.Chart
	//
	//agentLog.AgentLog
	//agent.Agent
	//answerRule.AnswerRule
	//audio.Audio
	//callQueue.CallQueue
	//call.Call
	//callerIdEmergency.CallerIdEmergency
	//cdr2.Cdr2
	//
	//callQueueEmailReport.CallQueueEmailReport
	//callQueueStat.CallQueueStat
	//cdrExport.CdrExport
	//cdrSchedule.CdrSchedule
	//
	//conference.Conference
	//conferenceCdr.ConferenceCdr
	//connection.Connection
	//dashboard.Dashboard
	//
	//dialPlan.DialPlan
	//dialPolicy.DialPolicy
	//dialRule.DialRule
	//image.Image
	//
	//meeting.Meeting
	//messageSession.MessageSession
	//participant.Participant
	//permission.Permission
	//
	//phoneConfiguration.PhoneConfiguration
	//phoneNumber.PhoneNumber
	//queued.Queued
	//quota.Quota
	//
	//recording.Recording
	//routeConnection.RouteConnection
	//smsNumber.SmsNumber
	//statistics.Statistics
	//
	//timeframe.Timeframe
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
	//p.ApiAppKey.SetBaseUrl(p.Web)
	//p.Areas.Site.Web = &p.Web
	//p.Areas.Subscriber.Web = &p.Web
	//p.Areas.Department.Web = &p.Web
	//p.Areas.Presence.Web = &p.Web
	//
	//p.Areas.Device.Web = &p.Web
	//p.Areas.Model.Web = &p.Web
	//p.Areas.Profile.Web = &p.Web
	//
	//p.Areas.Ndp.Web = &p.Web
	//p.Areas.Contact.Web = &p.Web
	//p.Areas.Mac.Web = &p.Web
	//p.Areas.Chart.Web = &p.Web
	//
	//p.Areas.AgentLog.Web = &p.Web
	//p.Areas.Agent.Web = &p.Web
	//p.Areas.AnswerRule.Web = &p.Web
	//p.Areas.Audio.Web = &p.Web
	//p.Areas.CallQueue.Web = &p.Web
	//p.Areas.Call.Web = &p.Web
	//p.Areas.CallerIdEmergency.Web = &p.Web
	//p.Areas.Cdr2.Web = &p.Web
	//
	//p.Areas.CallQueueEmailReport.Web = &p.Web
	//p.Areas.CallQueueStat.Web = &p.Web
	//p.Areas.CdrExport.Web = &p.Web
	//p.Areas.CdrSchedule.Web = &p.Web
	//
	//p.Areas.Conference.Web = &p.Web
	//p.Areas.ConferenceCdr.Web = &p.Web
	//p.Areas.Connection.Web = &p.Web
	//p.Areas.Dashboard.Web = &p.Web
	//
	//p.Areas.DialPlan.Web = &p.Web
	//p.Areas.DialPolicy.Web = &p.Web
	//p.Areas.DialRule.Web = &p.Web
	//p.Areas.Image.Web = &p.Web
	//
	//p.Areas.Meeting.Web = &p.Web
	//p.Areas.MessageSession.Web = &p.Web
	//p.Areas.Participant.Web = &p.Web
	//p.Areas.Permission.Web = &p.Web
	//
	//p.Areas.PhoneConfiguration.Web = &p.Web
	//p.Areas.PhoneNumber.Web = &p.Web
	//p.Areas.Queued.Web = &p.Web
	//p.Areas.Quota.Web = &p.Web
	//
	//p.Areas.Recording.Web = &p.Web
	//p.Areas.RouteConnection.Web = &p.Web
	//p.Areas.SmsNumber.Web = &p.Web
	//p.Areas.Statistics.Web = &p.Web
	//
	//p.Areas.Timeframe.Web = &p.Web

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

//func (p *SunGro) GetSite() string {
//	return p.Web.GetSite()
//}
//func (p *SunGro) VerifySite(user string) string {
//	return p.Web.VerifySite(user)
//}
//
//func (p *SunGro) GetClientId() string {
//	return p.Web.GetClientId()
//}
//func (p *SunGro) VerifyClientId(user string) string {
//	return p.Web.VerifyClientId(user)
//}
//
//func (p *SunGro) GetGroup() string {
//	return p.Web.GetGroup()
//}
//func (p *SunGro) VerifyGroup(user string) string {
//	return p.Web.VerifyGroup(user)
//}
//
//func (p *SunGro) GetTerritory() string {
//	return p.Web.GetTerritory()
//}
//func (p *SunGro) VerifyTerritory(user string) string {
//	return p.Web.VerifyTerritory(user)
//}
//
//func (p *SunGro) GetScope() string {
//	return p.Web.GetScope()
//}
//func (p *SunGro) VerifyScope(user string) string {
//	return p.Web.VerifyScope(user)
//}

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
