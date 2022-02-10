package device

//goland:noinspection GoNameStartsWithPackageName

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Device struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=device&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain    string `json:"domain" required:"true"` // Identifies Domain for which to count Devices.
	Territory string `json:"territory"`              // Identifies territory for which to count Devices.
	User      string `json:"user"`                   // Identifies User for which to count Devices.
	Aor       string `json:"aor"`                    // Filter count by Address of Record.
	Device    string `json:"device"`                 // Filter count by Device name.
	Limit     string `json:"limit"`                  // Set limit for number of returned outputs.
	Start     string `json:"start"`                  // Set starting point for returned outputs.
	Sort      string `json:"sort"`                   // Specifies parameter by which the output is sorted.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Devices for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=device&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain string `json:"domain" required:"true"` // Specifies Domain of new Device.
	Device string `json:"device" required:"true"` // Specifies name of new Device.
	User   string `json:"user" required:"true"`   // Specifies owner of new Device. Will be a valid user extension.
	Mac    string `json:"mac" required:"true"`    // Specifies MAC address of new Device. Just hex numbers, no -,:. needed. Size range: 12
	Model  string `json:"model" required:"true"`  // Specifies Model of new Device. Query available options via Devicemodel
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=device&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Device           string `json:"device" required:"true"` // Identifies Device to delete by name.
	TerminationMatch string `json:"termination_match"`      // Identifies Device to delete by termination match.
}
type DeleteResponse struct {
}

//const ListUrl = "?format=json&object=device&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read Devices.
}
type ListResponse []ListResponseSingle
type ListResponseSingle struct {
	Aor                     string `json:"aor"`                       // AOR
	Mac                     string `json:"mac"`                       // MAC Address
	Model                   string `json:"model"`                     // Brand and model
	AorUser                 string `json:"aor_user"`                  // AOR User
	Mode                    string `json:"mode"`                      // Mode
	Server                  string `json:"server"`                    // Preferred Server
	UserAgent               string `json:"user_agent"`                // User Agent
	Contact                 string `json:"contact"`                   //
	ReceivedFrom            string `json:"received_from"`             // Where Received From
	RegistrationTime        string `json:"registration_time"`         // Time of Registration
	Expires                 string `json:"expires"`                   // Time to Expiration
	RegistrationExpiresTime string `json:"registration_expires_time"` // Registration Expiration Time
	SubscriberName          string `json:"subscriber_name"`           // Name of Subscriber
	SubscriberDomain        string `json:"subscriber_domain"`         // ApiAppKey of Subscriber
	AuthenticationKey       string `json:"authentication_key"`        // Authentication Key/ApiPassword
	Transport               string `json:"transport"`                 // The transport type for the device if configured through the NDP
	Overrides               string `json:"overrides"`                 // Custom Configuration Overrides
	CallProcessingRule      string `json:"call_processing_rule"`      // Call Processing Rule
	CallidEmgr              string `json:"callid_emgr"`               // Caller ID for Emergency Calls
	Line                    int    `json:"line"`                      // Line assignment
	AuthUser                string `json:"auth_user"`                 // Device Authorization User
	AuthPass                string `json:"auth_pass"`                 // Device Authorization ApiPassword
	SubFullname             string `json:"sub_fullname"`              // Fullname of Subscriber
	SubScope                string `json:"sub_scope"`                 // Scope of Subscriber
	SubLogin                string `json:"sub_login"`                 // SetAuth of Subscriber
	Ndperror                string `json:"ndperror"`                  // NDP Error

	Owner string `json:"owner"`
}

//const ReadUrl = "?format=json&object=device&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain  string `json:"domain" required:"true"` // Identifies Domain from which to read Devices.
	Device  string `json:"device"`                 // Identifies name of device to read.
	User    string `json:"user"`                   //
	Owner   string `json:"owner"`                  //
	Start   string `json:"start"`                  //
	Mode    string `json:"mode"`                   //
	Limit   string `json:"limit"`                  //
	Sort    string `json:"sort"`                   //
	List    string `json:"list"`                   //
	Fields  string `json:"fields"`                 //
	Filters string `json:"filters"`                //
	NoNDP   string `json:"noNDP"`                  //
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Aor                     string `json:"aor"`                       // AOR
	Mac                     string `json:"mac"`                       // MAC Address
	Model                   string `json:"model"`                     // Brand and model
	AorUser                 string `json:"aor_user"`                  // AOR User
	Mode                    string `json:"mode"`                      // Mode
	Server                  string `json:"server"`                    // Preferred Server
	UserAgent               string `json:"user_agent"`                // User Agent
	Contact                 string `json:"contact"`                   //
	ReceivedFrom            string `json:"received_from"`             // Where Received From
	RegistrationTime        string `json:"registration_time"`         // Time of Registration
	Expires                 string `json:"expires"`                   // Time to Expiration
	RegistrationExpiresTime string `json:"registration_expires_time"` // Registration Expiration Time
	SubscriberName          string `json:"subscriber_name"`           // Name of Subscriber
	SubscriberDomain        string `json:"subscriber_domain"`         // ApiAppKey of Subscriber
	AuthenticationKey       string `json:"authentication_key"`        // Authentication Key/ApiPassword
	Transport               string `json:"transport"`                 // The transport type for the device if configured through the NDP
	Overrides               string `json:"overrides"`                 // Custom Configuration Overrides
	CallProcessingRule      string `json:"call_processing_rule"`      // Call Processing Rule
	CallidEmgr              string `json:"callid_emgr"`               // Caller ID for Emergency Calls
	Line                    string `json:"line"`                      // Line assignment
	AuthUser                string `json:"auth_user"`                 // Device Authorization User
	AuthPass                string `json:"auth_pass"`                 // Device Authorization ApiPassword
	SubFullname             string `json:"sub_fullname"`              // Fullname of Subscriber
	SubScope                string `json:"sub_scope"`                 // Scope of Subscriber
	SubLogin                string `json:"sub_login"`                 // SetAuth of Subscriber
	Ndperror                string `json:"ndperror"`                  // NDP Error
}

//const UpdateUrl = "?format=json&object=device&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain                  string `json:"domain" required:"true"`    // New value for Device's Domain.
	Device                  string `json:"device" required:"true"`    // Identifies Devices to update by aor.
	User                    string `json:"user"`                      // The user of the device in the domain.
	CheckSync               string `json:"check-sync"`                // Sets Device's new check-sync value, set to yes or 1
	EvtCheckSync            string `json:"evtCheckSync"`              // If check-sync is used it allows you to customize the check-sync Event Header. Default value: check-sync
	Mac                     string `json:"mac"`                       // New value for Device's MAC address.
	Transport               string `json:"transport"`                 // The transport type for the device if configured through the NDP
	Model                   string `json:"model"`                     // New value for Device's model.
	Mode                    string `json:"mode"`                      // Mode
	UserAgent               string `json:"user_agent"`                // User Agent
	AcceptAgent             string `json:"accept_agent"`              // Accept Agent field. if set to not empty the core may filter registrations where user_agent must match this string.
	ReceivedFrom            string `json:"received_from"`             // Where Received From
	RegistrationTime        string `json:"registration_time"`         // Time of Registration
	SubscriberName          string `json:"subscriber_name"`           // Name of Subscriber
	SubscriberDomain        string `json:"subscriber_domain"`         // Domain of Subscriber
	AuthenticationKey       string `json:"authentication_key"`        // Authentication Key/ApiPassword
	CallProcessingRule      string `json:"call_processing_rule"`      // Call Processing Rule
	RegistrationExpiresTime string `json:"registration_expires_time"` // Registration Expiration Time
	SubFullname             string `json:"sub_fullname"`              // Fullname of Subscriber
	SubScope                string `json:"sub_scope"`                 // Scope of Subscriber
	SubLogin                string `json:"sub_login"`                 // SetAuth of Subscriber
	Ndperror                string `json:"ndperror"`                  // ndperror
}
type UpdateResponse struct {
}
