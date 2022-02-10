package domain

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Domain struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

const CountUrl = "?format=json&object=domain&action=count"

type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	BaseUrl  string
	Error    error
}
type CountRequest struct {
	Domain    string `json:"domain"`    // Name of Domain.
	Territory string `json:"territory"` // Territory/Reseller for the Domain.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Domains in the specified domain.
}

const CreateUrl = "?format=json&object=domain&action=create"

type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	BaseUrl  string
	Error    error
}
type CreateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type CreateResponse struct {
}

const DeleteUrl = "?format=json&object=domain&action=delete"

type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	BaseUrl  string
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type DeleteResponse struct {
}

const ReadUrl = "?format=json&object=domain&action=read"

type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	BaseUrl  string
	Error    error
}
type ReadRequest struct {
	Domain string `json:"domain"` // Name of Domain.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Domain                   string `json:"domain"`                      // Name of Domain.
	Territory                string `json:"territory"`                   // Territory/Reseller for the Domain.
	DialMatch                string `json:"dial_match"`                  // (deprecated)
	Description              string `json:"description"`                 // Description for the Domain.
	Moh                      string `json:"moh"`                         // Music on hold
	Mor                      string `json:"mor"`                         // Music on ring
	Mot                      string `json:"mot"`                         // Music on transfer (deprecated)
	Rmoh                     string `json:"rmoh"`                        // Randomize Music on hold
	Rating                   string `json:"rating"`                      // Call Rating Prepaid (deprecated)
	Resi                     string `json:"resi"`                        // Residential Domain, Limited business features
	Mksdir                   string `json:"mksdir"`                      // Setup file structure for Users (should always be yes).
	DialPlan                 string `json:"dial_plan"`                   // Default DialPlan for the Domain.
	DialPolicy               string `json:"dial_policy"`                 // Default Dial Permission for the Domain.
	CallLimit                int    `json:"call_limit"`                  // The active Call Limit for this Domain.
	CallLimitExt             int    `json:"call_limit_ext"`              // The Call Limit for external calls.
	SubLimit                 int    `json:"sub_limit"`                   // Subscriber limit (includes system users)
	TimeZone                 string `json:"time_zone"`                   // Timezone for domain
	MaxUser                  int    `json:"max_user"`                    // The Maximum user count that the Domain can have.
	MaxCallQueue             int    `json:"max_call_queue"`              // The Maximum call queue count that the Domain can have.
	MaxAa                    int    `json:"max_aa"`                      // The Maximum Auto Attendant count that the Domain can have.
	MaxConference            int    `json:"max_conference"`              // The Maximum conference bridge count that the Domain can have.
	MaxDepartment            int    `json:"max_department"`              // The Maximum department count that the Domain can have.
	MaxSite                  int    `json:"max_site"`                    // The Maximum site count that the Domain can have.
	MaxDevice                int    `json:"max_device"`                  // The Maximum device count per user.
	Policies                 string `json:"policies"`                    // Policy on the Domain active/locked.
	EmailSender              string `json:"email_sender"`                // Email address use as 'from' address for system-dispatched emails.
	CallidNmbr               int64  `json:"callid_nmbr"`                 // Default CallerID setting for the Domain.
	CallidName               string `json:"callid_name"`                 // Default Cname setting for the Domain.
	CallidEmgr               int64  `json:"callid_emgr"`                 // Default 911 CallerID setting for the Domain.
	AreaCode                 int    `json:"area_code"`                   // Default area code setting for the Domain.
	SmsInboundLast           int    `json:"sms_inbound_last"`            // Total number of inbound sms for last the month .
	SmsOutboundLast          int    `json:"sms_outbound_last"`           // Total number of outbound sms for last the month.
	SmsOutboundCurrent       int    `json:"sms_outbound_current"`        // Total number of outbound sms for this month.
	SmsInboundCurrent        int    `json:"sms_inbound_current"`         // Total number of inbound sms for this month.
	ActiveCallsOnnetLast     int    `json:"active_calls_onnet_last"`     // Total number of onnet active calls for the last month.
	ActiveCallsOffnetLast    int    `json:"active_calls_offnet_last"`    // Total number of offnet active calls for the last month.
	ActiveCallsOffnetCurrent int    `json:"active_calls_offnet_current"` // Total number of offnet active calls for this month.
	ActiveCallsOnnetCurrent  int    `json:"active_calls_onnet_current"`  // Total number of onnet active calls for this month.
	SmtpHost                 string `json:"smtp_host"`                   // Email setting: host server
	SmtpPort                 string `json:"smtp_port"`                   // Email setting: host port
	SmtpUid                  string `json:"smtp_uid"`                    // Email setting: user
	SmtpPwd                  string `json:"smtp_pwd"`                    // Email setting: password
	FromUser                 string `json:"from_user"`                   // (deprecated)
	FromHost                 string `json:"from_host"`                   // (deprecated)
	ActiveCall               int    `json:"active_call"`                 // Current number of active calls
	CountForLimit            int    `json:"countForLimit"`               // Count against the Call Limit
	CountExternal            int    `json:"countExternal"`               // Count against the External Call Limit
	SubCount                 int    `json:"sub_count"`                   // Total number of subscribers for the Domain.

	Sso              string `json:"sso"`               // Added
	Language         string `json:"language"`          // Added
	MaxFax           int    `json:"max_fax"`           // Added
	VmailProvisioned string `json:"vmail_provisioned"` // Added
	VmailTranscribe  string `json:"vmail_transcribe"`  // Added
}

const BillingUrl = "?format=json&object=domain&action=read&billing=yes"

type Billing struct {
	Request  BillingRequest
	Response BillingResponse
	BaseUrl  string
	Error    error
}
type BillingRequest struct {
	Domain string `json:"domain"` // Name of Domain.
}
type BillingResponse []BillingResponseSingle
type BillingResponseSingle struct {
	Domain                     string `json:"domain"`                       // Name of Domain.
	Territory                  string `json:"territory"`                    // Territory/Reseller for the Domain.
	Description                string `json:"description"`                  // Description for the Domain.
	CallLimit                  int    `json:"call_limit"`                   // The active Call Limit for this Domain.
	CallLimitExt               int    `json:"call_limit_ext"`               // The active Call Limit for this Domain only counting calls external to the domain.
	MaxUser                    int    `json:"max_user"`                     // The Maximum user count that the Domain can have.
	MaxCallQueue               int    `json:"max_call_queue"`               // The Maximum call queue count that the Domain can have.
	MaxAa                      int    `json:"max_aa"`                       // The Maximum Auto Attendant count that the Domain can have.
	MaxConference              int    `json:"max_conference"`               // The Maximum conference bridge count that the Domain can have.
	MaxDepartment              int    `json:"max_department"`               // The Maximum department count that the Domain can have.
	MaxSite                    int    `json:"max_site"`                     // The Maximum site count that the Domain can have.
	MaxDevice                  int    `json:"max_device"`                   // The Maximum device count per user.
	CurrentUser                int    `json:"current_user"`                 // The current user count that the Domain can have.
	CurrentCallQueue           int    `json:"current_call_queue"`           // The current call queue count that the Domain can have.
	CurrentPark                int    `json:"current_park"`                 // The current number of call parks enabled for the domain.
	CurrentAa                  int    `json:"current_aa"`                   // The current Auto Attendant count that the Domain can have.
	CurrentConference          int    `json:"current_conference"`           // The current conference bridge count that the Domain can have.
	CurrentDepartment          int    `json:"current_department"`           // The current department count that the Domain can have.
	CurrentRegisteredDevice    int    `json:"current_registered_device"`    // The current count of registered devices.
	CurrentDevice              int    `json:"current_device"`               // The current total device count.
	CurrentPhonenumbers        int    `json:"current_phonenumbers"`         // The current count of phone numbers assigned to domain
	CurrentTollfree            int    `json:"current_tollfree"`             // The current count of toll free phonenumbers
	CurrentScope_SCOPE         int    `json:"current_scope_SCOPE"`          // The current count the SCOPE for this domain. will show 1 entry per scope type.
	CurrentTranscribe_VENDOR   int    `json:"current_transcribe_VENDOR"`    // The current count of users assigned to vmail transcription by VENDOR for this domain. will show 1 entry per VENDOR enabled.
	CurrentServiceCode_SRVCODE int    `json:"current_service_code_SRVCODE"` // The current count of users tagged with server code SRVCODE for this domain. will show 1 entry per SRVCODE used.
	SmsInboundLast             int    `json:"sms_inbound_last"`             // Total number of inbound sms for last the month .
	SmsOutboundLast            int    `json:"sms_outbound_last"`            // Total number of outbound sms for last the month.
	SmsOutboundCurrent         int    `json:"sms_outbound_current"`         // Total number of outbound sms for this month.
	SmsInboundCurrent          int    `json:"sms_inbound_current"`          // Total number of inbound sms for this month.
	ActiveCallsOnnetLast       int    `json:"active_calls_onnet_last"`      // Total number of onnet active calls for the last month.
	ActiveCallsOffnetLast      int    `json:"active_calls_offnet_last"`     // Total number of offnet active calls for the last month.
	ActiveCallsOffnetCurrent   int    `json:"active_calls_offnet_current"`  // Total number of offnet active calls for this month.
	ActiveCallsOnnetCurrent    int    `json:"active_calls_onnet_current"`   //Total number of onnet active calls for this month.
	CalculationTimeMs          int    `json:"calculation_time_ms"`          //The time it too the API to generate this data.
}

const ListUrl = "?format=json&object=domain&action=list"

type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	BaseUrl  string
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain"` // Domain containing Users to list.
}
type ListResponse []ListResponseSingle
type ListResponseSingle string

const UpdateUrl = "?format=json&object=domain&action=update"

type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	BaseUrl  string
	Error    error
}
type UpdateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
	Site   string `json:"site"`   // Identifies site from which to read info.
}
type UpdateResponse struct {
}
