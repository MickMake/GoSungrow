package mac

//goland:noinspection GoNameStartsWithPackageName

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Mac struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=mac&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain    string `json:"domain" required:"true"` // Identifies Domain for which to count MAC addresses.
	Territory string `json:"territory"`              // Identifies territory for which to count MAC addresses.
	Mac       string `json:"mac"`                    // Filters count by targeting addresses with the specified MAC address.
	MacLIKE   string `json:"mac_LIKE"`               //
	Model     string `json:"model"`                  //
	Notes     string `json:"notes"`                  //
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Devices for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=mac&action=create"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Mac         string `json:"mac" required:"true"`    // Specifies MAC address of new Device. Just hex numbers, no -,:. needed. Size range: 12
	Domain      string `json:"domain" required:"true"` // Specifies Domain of new Device.
	Model       string `json:"model" required:"true"`  // Specifies Model of new Device. Query available options via Devicemodel
	Transport   string `json:"transport"`              // The transport type for the device if configured through the NDP
	Server      string `json:"server"`                 //
	LastPull    string `json:"last_pull"`              //
	DateCreated string `json:"date_created"`           //
	Device1     string `json:"device1"`                //
	Device2     string `json:"device2"`                //
	Device3     string `json:"device3"`                //
	Device4     string `json:"device4"`                //
	Device5     string `json:"device5"`                //
	Device6     string `json:"device6"`                //
	Device7     string `json:"device7"`                //
	Device8     string `json:"device8"`                //
	Territory   string `json:"territory"`              //
	Notes       string `json:"notes"`                  //
	Line1Ext    string `json:"line1_ext"`              //
	Line2Ext    string `json:"line2_ext"`              //
	Redundancy  string `json:"redundancy"`             //
	Line1Enable string `json:"line1_enable"`           //
	Line2Enable string `json:"line2_enable"`           //
	Line3Enable string `json:"line3_enable"`           //
	Line4Enable string `json:"line4_enable"`           //
	Line5Enable string `json:"line5_enable"`           //
	Line6Enable string `json:"line6_enable"`           //
	Line7Enable string `json:"line7_enable"`           //
	Line8Enable string `json:"line8_enable"`           //
	Line1Share  string `json:"line1_share"`            //
	Line2Share  string `json:"line2_share"`            //
	Line3Share  string `json:"line3_share"`            //
	Line4Share  string `json:"line4_share"`            //
	Line5Share  string `json:"line5_share"`            //
	Line6Share  string `json:"line6_share"`            //
	Line7Share  string `json:"line7_share"`            //
	Line8Share  string `json:"line8_share"`            //
	Overrides   string `json:"overrides"`              //
	DirInc      string `json:"dir_inc"`                //
	Presence    string `json:"presence"`               //
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=mac&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain" required:"true"` // Identifies Domain of Device to delete.
	Mac    string `json:"mac" required:"true"`    // Identifies MAC address of Device to delete. The MAC address should be a numerical string with ':' removed.
}
type DeleteResponse struct {
}

//const ListUrl = "?format=json&object=mac&action=list"
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
	Mac                     string `json:"mac"`                       // MAC Address
	Server                  string `json:"server"`                    // Server
	Territory               string `json:"territory"`                 // Territory
	DirInc                  string `json:"dir_inc"`                   // dir_inc
	Presence                string `json:"presence"`                  // Presence
	Domain                  string `json:"domain"`                    // Domain
	Model                   string `json:"model"`                     // Model
	Device1                 string `json:"device1"`                   // Device1
	Device2                 string `json:"device2"`                   // Device2
	Device3                 string `json:"device3"`                   // Device3
	Device4                 string `json:"device4"`                   // Device4
	Device5                 string `json:"device5"`                   // Device5
	Device6                 string `json:"device6"`                   // Device6
	Device7                 string `json:"device7"`                   // Device7
	Device8                 string `json:"device8"`                   // Device8
	Notes                   string `json:"notes"`                     // notes
	Line1Share              string `json:"line1_share"`               // line1_share
	Line2Share              string `json:"line2_share"`               // line2_share
	Line3Share              string `json:"line3_share"`               // line3_share
	Line4Share              string `json:"line4_share"`               // line4_share
	Line5Share              string `json:"line5_share"`               // line5_share
	Line6Share              string `json:"line6_share"`               // line6_share
	Overrides               string `json:"overrides"`                 // overrides
	PhoneExt                string `json:"phone_ext"`                 // Phone Extension
	Transport               string `json:"transport"`                 // The transport type for the device if configured through the NDP
	Fxs                     string `json:"fxs"`                       // fxs
	Sla                     string `json:"sla"`                       // sla
	Sidecar                 string `json:"sidecar"`                   // sidecar
	Resync                  string `json:"resync"`                    // resync
	DirectorySupport        string `json:"directory_support"`         // Directory Support
	UserAgent               string `json:"user_agent"`                // User Agent
	Contact                 string `json:"contact"`                   // Contact
	RegistrationExpiresTime string `json:"registration_expires_time"` // Registration Expiration Time
	RegistrationTime        string `json:"registration_time"`         // Registration Time
}

//const ReadUrl = "?format=json&object=mac&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Mac       string `json:"mac" required:"unknown"` // Identifies MAC address of Device to read. The MAC address should be a numerical string with ':' removed.
	Extension string `json:"extension"`              // numerical string, parameter searches lines containing this extension associated with the mac address
	//CheckExistance bool   `json:"checkExistance"`         // @TODO - doesn't seem to work.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Mac                     string `json:"mac"`                       // MAC Address
	Server                  string `json:"server"`                    // Server
	Territory               string `json:"territory"`                 // Territory
	DirInc                  string `json:"dir_inc"`                   // dir_inc
	Presence                string `json:"presence"`                  // Presence
	Domain                  string `json:"domain"`                    // Domain
	Model                   string `json:"model"`                     // Model
	Device1                 string `json:"device1"`                   // Device1
	Device2                 string `json:"device2"`                   // Device2
	Device3                 string `json:"device3"`                   // Device3
	Device4                 string `json:"device4"`                   // Device4
	Device5                 string `json:"device5"`                   // Device5
	Device6                 string `json:"device6"`                   // Device6
	Device7                 string `json:"device7"`                   // Device7
	Device8                 string `json:"device8"`                   // Device8
	Notes                   string `json:"notes"`                     // notes
	Line1Share              string `json:"line1_share"`               // line1_share
	Line2Share              string `json:"line2_share"`               // line2_share
	Line3Share              string `json:"line3_share"`               // line3_share
	Line4Share              string `json:"line4_share"`               // line4_share
	Line5Share              string `json:"line5_share"`               // line5_share
	Line6Share              string `json:"line6_share"`               // line6_share
	Overrides               string `json:"overrides"`                 // overrides
	PhoneExt                string `json:"phone_ext"`                 // Phone Extension
	Transport               string `json:"transport"`                 // The transport type for the device if configured through the NDP
	Fxs                     string `json:"fxs"`                       // fxs
	Sla                     string `json:"sla"`                       // sla
	Sidecar                 string `json:"sidecar"`                   // sidecar
	Resync                  string `json:"resync"`                    // resync
	DirectorySupport        string `json:"directory_support"`         // Directory Support
	UserAgent               string `json:"user_agent"`                // User Agent
	Contact                 string `json:"contact"`                   // Contact
	RegistrationExpiresTime string `json:"registration_expires_time"` // Registration Expiration Time
	RegistrationTime        string `json:"registration_time"`         // Registration Time
}

//const UpdateUrl = "?format=json&object=device&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Mac         string `json:"mac" required:"true"` // Identifies MAC address of Device to update. The MAC address should be a numerical string with ':' removed.
	Model       string `json:"model"`               // line	string	// device	string	// transport	string	// The transport type for the device if configured through the NDP
	Server      string `json:"server"`              //
	LastPull    string `json:"last_pull"`           //
	DateCreated string `json:"date_created"`        //
	Device1     string `json:"device1"`             //
	Device2     string `json:"device2"`             //
	Device3     string `json:"device3"`             //
	Device4     string `json:"device4"`             //
	Device5     string `json:"device5"`             //
	Device6     string `json:"device6"`             //
	Device7     string `json:"device7"`             //
	Device8     string `json:"device8"`             //
	Territory   string `json:"territory"`           //
	Notes       string `json:"notes"`               //
	Line1Ext    string `json:"line1_ext"`           //
	Line2Ext    string `json:"line2_ext"`           //
	Redundancy  string `json:"redundancy"`          //
	Line1Enable string `json:"line1_enable"`        //
	Line2Enable string `json:"line2_enable"`        //
	Line3Enable string `json:"line3_enable"`        //
	Line4Enable string `json:"line4_enable"`        //
	Line5Enable string `json:"line5_enable"`        //
	Line6Enable string `json:"line6_enable"`        //
	Line7Enable string `json:"line7_enable"`        //
	Line8Enable string `json:"line8_enable"`        //
	Line1Share  string `json:"line1_share"`         //
	Line2Share  string `json:"line2_share"`         //
	Line3Share  string `json:"line3_share"`         //
	Line4Share  string `json:"line4_share"`         //
	Line5Share  string `json:"line5_share"`         //
	Line6Share  string `json:"line6_share"`         //
	Line7Share  string `json:"line7_share"`         //
	Line8Share  string `json:"line8_share"`         //
	Overrides   string `json:"overrides"`           //
	DirInc      string `json:"dir_inc"`             //
	Presence    string `json:"presence"`            //
}
type UpdateResponse struct {
}
