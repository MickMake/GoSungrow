package subscriber

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Subscriber struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=subscriber&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain    string `json:"domain"`    // Name of Domain.
	Territory string `json:"Territory"` // Territory/Reseller for the Domain.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Domains in the specified domain.
}

//const CreateUrl = "?format=json&object=subscriber&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Uid              string `json:"uid"`               // User ID of new User.
	User             string `json:"user"`              // Extension of new User. Cannot be changed once set.
	Domain           string `json:"domain"`            // Domain where new User will be created.
	SubscriberLogin  string `json:"subscriber_login"`  // SetAuth for the User. If not set, user@domain will be used for login.
	FirstName        string `json:"first_name"`        // First name of new User.
	LastName         string `json:"last_name"`         // Last name of new User.
	PasswordLength   string `json:"passwordLength"`    // Length for Random new password
	PwdHash          string `json:"pwd_hash"`          // Portal password for new User.
	SubscriberPin    string `json:"subscriber_pin"`    // Voicemail PIN for new User.
	DirAnc           string `json:"dir_anc"`           // Determines if the user will be in any dial by name directory options. Allowed values: yes, no
	DirList          string `json:"dir_list"`          // Determines if the user will be in shown in any visual company directories. Allowed values: yes, no
	Message          string `json:"message"`           // A status message for a user
	Pwd              string `json:"pwd"`               // The password for the user to use on the vmail pin or portal access.
	Group            string `json:"group"`             // The department or group for this user.
	Dir              string `json:"dir"`               // A 3 digit number string matching the fist three digits of the user's first or last name.
	Email            string `json:"email"`             // email address for the user.
	VmailProvisioned string `json:"vmail_provisioned"` // Voicemail box allowed for the user. This one is generally not user configurable. Allowed values: yes, no
	VmailFrdTo       string `json:"vmail_frd_to"`      // A comma-separated list of email addresses to distribute voicemails to. Allowed values: yes, no
	VmailEnabled     string `json:"vmail_enabled"`     // Voicemail enabled for the user. This one is generally user configurable. Allowed values: yes, no
	VmailGreeting    string `json:"vmail_greeting"`    // Current voicemail greeting index in use.
	VmailNotify      string `json:"vmail_notify"`      // Option to send new voicemails to email. Allowed values: no, yes, attnew, attsave, atttrash
	VmailAnncTime    string `json:"vmail_annc_time"`   // Announce the time during vmail playback.
	VmailAnncCid     string `json:"vmail_annc_cid"`    // Announce the caller id during vmail playback.
	VmailSortLifo    string `json:"vmail_sort_lifo"`   // Sort by last in first out during vmail playback
	VmailTranscribe  string `json:"vmail_transcribe"`  // Transcribe voicemails (when enabled by snapped in services).
	DataLimit        string `json:"data_limit"`        // Max size of audio for a user in Kb, example "5000" would be roughly 5MB.
	CallLimit        string `json:"call_limit"`        // Max active calls for the user.
	DialPlan         string `json:"dial_plan"`         // The dial plan for the user.
	DialPolicy       string `json:"dial_policy"`       // The dial permission for the user.
	AreaCode         string `json:"area_code"`         // A 3 digit area code for the user.
	CallidName       string `json:"callid_name"`       // A callerid Name for CNAM purposes
	CallidNmbr       string `json:"callid_nmbr"`       // A callerid for this user on outbound calls.
	CallidEmgr       string `json:"callid_emgr"`       // A callerid for this user on emergency calls.
	NoAnswerTimeout  string `json:"no_answer_timeout"` // The timeout for his user between user features and before going to voicemail
	TimeZone         string `json:"time_zone"`         // Timezone for this user. Valid values are in "TZ" column of this chart, https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Scope            string `json:"scope"`             // The Scope for this user. Allowed values: "Basic User", "Office Manager", "Call Center Agent", "Call Center Supervisor", "No Portal"
	RejAnony         string `json:"rej_anony"`         // Reject any inbound anonymous call. Allowed values: yes, no
	Screen           string `json:"screen"`            // Allow call screening application. Allowed values: yes, no
	SrvCode          string `json:"srv_code"`          // Configurable service code used for 3rd party purposes. some system level user tagged with system-X where x is configurable. Allowed values: yes, no
	NtfyMissedCall   string `json:"ntfy_missed_call"`  // Send email when call is missed. Allowed values: yes, no
	NtfyDataLimit    string `json:"ntfy_data_limit"`   // Send email when data limit is reached. Allowed values: yes, no
	Language         string `json:"language"`          // Localized language for this user.
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=subscriber&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Uid              string `json:"uid"`               // Identifies User to delete by User ID.
	Domain           string `json:"domain"`            // Identifies User to delete by the domain to which a User belongs.
	User             string `json:"user"`              // Identifies User to delete by extension.
	TerminationMatch string `json:"termination_match"` //
	Device           string `json:"device"`            //
}
type DeleteResponse struct {
}

//const ReadUrl = "?format=json&object=subscriber&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Uid            string `json:"uid"`                    // Identifies User to read by User ID.
	User           string `json:"user" required:"true"`   // Identifies a User to read by User ID when a Domain is specified.
	Domain         string `json:"domain" required:"true"` // Domain containing Users to read.
	Login          string `json:"login"`                  // Identifies User to read by login info.
	Limit          string `json:"limit"`                  // Identifies number of Users to read when a Domain is specified.
	FirstName      string `json:"first_name"`             // Identifies User to read by first name.
	LastName       string `json:"last_name"`              // Identifies User to read by last name.
	Group          string `json:"group"`                  //
	Fields         string `json:"fields"`                 // Comma separated list of fields to filter the return data.
	SrvCode        string `json:"srv_code"`               // Allows for searching for specific service codes.
	Email          string `json:"email"`                  //
	Dir            string `json:"dir"`                    //
	FilterUsers    string `json:"filter_users"`           //
	DirectoryMatch string `json:"directory_match"`        //
	Owner          string `json:"owner"`                  //
	Scope          string `json:"scope"`                  //
	Start          string `json:"start"`                  //
	Sort           string `json:"sort"`                   // Sorts User output by the specified parameter. Example: domain
}
type ReadResponseSingle struct {
	Domain           string `json:"domain"`            //
	User             string `json:"user"`              //
	FirstName        string `json:"first_name"`        //
	LastName         string `json:"last_name"`         //
	SubscriberLogin  string `json:"subscriber_login"`  //
	SubscriberPin    string `json:"subscriber_pin"`    //
	PwdHash          string `json:"pwd_hash"`          //
	Group            string `json:"group"`             //
	Site             string `json:"site"`              //
	Dir              string `json:"dir"`               //
	Email            string `json:"email"`             //
	Presence         string `json:"presence"`          //
	Message          string `json:"message"`           //
	Accept           string `json:"accept"`            //
	Reject           string `json:"reject"`            //
	VmailProvisioned string `json:"vmail_provisioned"` //
	VmailEnabled     string `json:"vmail_enabled"`     //
	VmailGreeting    string `json:"vmail_greeting"`    //
	VmailNotify      string `json:"vmail_notify"`      //
	VmailAnncTime    string `json:"vmail_annc_time"`   //
	VmailAnncCid     string `json:"vmail_annc_cid"`    //
	VmailSortLifo    string `json:"vmail_sort_lifo"`   //
	VmailTranscribe  string `json:"vmail_transcribe"`  //
	VmailFwdTo       string `json:"vmail_fwd_to"`      //
	DataLimit        string `json:"data_limit"`        //
	CallLimit        string `json:"call_limit"`        //
	DialPlan         string `json:"dial_plan"`         //
	DialPolicy       string `json:"dial_policy"`       //
	AreaCode         string `json:"area_code"`         //
	CallidName       string `json:"callid_name"`       //
	CallidNmbr       string `json:"callid_nmbr"`       //
	CallidEmgr       string `json:"callid_emgr"`       //
	NoAnswerTimeout  string `json:"no_answer_timeout"` //
	TimeZone         string `json:"time_zone"`         //
	DirAnc           string `json:"dir_anc"`           //
	DirList          string `json:"dir_list"`          //
	DateCreated      string `json:"date_created"`      //
	Scope            string `json:"scope"`             //
	RejAnony         string `json:"rej_anony"`         //
	DirectoryOrder   string `json:"directory_order"`   //
	Screen           string `json:"screen"`            //
	SrvCode          string `json:"srv_code"`          //
	NtfyMissedCall   string `json:"ntfy_missed_call"`  //
	NtfyDataLimit    string `json:"ntfy_data_limit"`   //
	Language         string `json:"language"`          //
	GauSession       string `json:"gauSession"`        //
	LastUpdate       string `json:"last_update"`       //
	AccountStatus    string `json:"account_status"`    //
}
type ReadResponse []ReadResponseSingle

//const ListUrl = "?format=json&object=subscriber&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain" required:"true"` // Domain containing Users to list.
	Fields string `json:"fields"`                 // Comma delimited string of specific return fields to list.
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	Domain    string `json:"domain"`     //
	Site      string `json:"site"`       //
	Group     string `json:"group"`      //
	User      string `json:"user"`       //
	FirstName string `json:"first_name"` //
	LastName  string `json:"last_name"`  //
	Dir       string `json:"dir"`        //
}

//const UpdateUrl = "?format=json&object=subscriber&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Uid              string `json:"uid"`               // Identifies User to update by User ID. A User's UID cannot be updated once it is created. To update a User's UID a new User must be created. Example: user@domain
	User             string `json:"user"`              // Identifies User to update by extension. A User's extension cannot be updated once it is created. To update a User's extension a new User must be created.
	Domain           string `json:"domain"`            // Identifies User to update by Domain. A User's Domain cannot be updated once it is created.
	FirstName        string `json:"first_name"`        // Identifies User to update by first name.
	LastName         string `json:"last_name"`         // Identifies User to update by last name.
	DirAnc           string `json:"dir_anc"`           // Determines if the user will be in any dial by name directory options. Allowed values: yes, no
	DirList          string `json:"dir_list"`          // Determines if the user will be in shown in any visual company directories. Allowed values: yes, no
	Message          string `json:"message"`           // A status message for a user
	SubscriberLogin  string `json:"subscriber_login"`  // The login for the use via the portal or API
	Pwd              string `json:"pwd"`               // The password for the user to use on the vmail pin.
	CurrentPassword  string `json:"current_password"`  // The password for the user to use on the vmail pin.
	PwdHash          string `json:"pwd_hash"`          // The password for portal access.
	Group            string `json:"group"`             // The department or group for this user.
	Dir              string `json:"dir"`               // A 3 digit number string matching the fist three digits of the user's first or last name.
	Email            string `json:"email"`             // email address for the user.
	VmailProvisioned string `json:"vmail_provisioned"` // Voicemail box allowed for the user. This one is generally not user configurable. Allowed values: yes, no
	VmailFrdTo       string `json:"vmail_frd_to"`      // A comma-separated list of email addresses to distribute voicemails to. Allowed values: yes, no
	VmailEnabled     string `json:"vmail_enabled"`     // Voicemail enabled for the user. This one is generally user configurable. Allowed values: yes, no
	VmailGreeting    string `json:"vmail_greeting"`    // Current voicemail greeting index in use.
	VmailNotify      string `json:"vmail_notify"`      // Option to send new voicemails to email. Allowed values: no, yes, attnew, attsave, atttrash
	VmailAnncTime    string `json:"vmail_annc_time"`   // Announce the time during vmail playback.
	VmailAnncCid     string `json:"vmail_annc_cid"`    // Announce the caller id during vmail playback.
	VmailSortLifo    string `json:"vmail_sort_lifo"`   // Sort by last in first out during vmail playback
	VmailTranscribe  string `json:"vmail_transcribe"`  // Transcribe voicemails (when enabled by snapped in services).
	DataLimit        string `json:"data_limit"`        // Max size of audio for a user in Kb, example "5000" would be roughly 5MB.
	CallLimit        string `json:"call_limit"`        // Max active calls for the user.
	DialPlan         string `json:"dial_plan"`         // The dial plan for the user.
	DialPolicy       string `json:"dial_policy"`       // The dial permission for the user.
	AreaCode         string `json:"area_code"`         // A 3 digit area code for the user.
	CallidName       string `json:"callid_name"`       // A callerid Name for CNAM purposes
	CallidNmbr       string `json:"callid_nmbr"`       // A callerid for this user on outbound calls.
	CallidEmgr       string `json:"callid_emgr"`       // A callerid for this user on emergency calls.
	NoAnswerTimeout  string `json:"no_answer_timeout"` // The timeout for his user between user features and before going to voicemail
	TimeZone         string `json:"time_zone"`         // Timezone for this user. Valid values are in "TZ" column of this chart, https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	Scope            string `json:"scope"`             // The Scope for this user. Allowed values: "Basic User", "Office Manager", "Call Center Agent", "Call Center Supervisor", "No Portal"
	RejAnony         string `json:"rej_anony"`         // Reject any inbound anonymous call. Allowed values: yes, no
	Screen           string `json:"screen"`            // Allow call screening application Allowed values: yes, no
	SrvCode          string `json:"srv_code"`          // Configurable service code used for 3rd party purposes. some system level user tagged with system-X where x is configurable. Allowed values: yes, no
	NtfyMissedCall   string `json:"ntfy_missed_call"`  // Send email when call is missed. Allowed values: yes, no
	NtfyDataLimit    string `json:"ntfy_data_limit"`   // Send email when data limit is reached. Allowed values: yes, no
	Language         string `json:"language"`          // Localized language for this user.
}
type UpdateResponse struct {
}
