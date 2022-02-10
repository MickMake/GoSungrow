package department

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Department struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=department&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain string `json:"domain" required:"true"` // Domain where new Department will be created.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Departments for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=subscriber&action=create&directory_match=departments"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	User           string `json:"user"`                                  // Extension of Department. Example: "Sales"
	Domain         string `json:"domain" required:"true"`                // Domain where new Department will be created.
	DirectoryMatch string `json:"directory_match" default:"departments"` // SetAuth for the User. Default value: departments
	FirstName      string `json:"first_name"`                            // First name of new User.
	LastName       string `json:"last_name"`                             // Last name of new User.
	DirList        string `json:"dir_list"`                              // Include in directory. Default value: no
	DirAnc         string `json:"dir_anc"`                               // Include in audible directories. Default value: no
	SrvCode        string `json:"srv_code"`                              // Service code for Department. Default value: system-department
	PasswordLength string `json:"passwordLength"`                        // Length for Random new password
	PwdHash        string `json:"pwd_hash"`                              // Portal password for new User.
	SubscriberPin  string `json:"subscriber_pin"`                        // Voicemail PIN for new User.
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=department&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain string `json:"domain" required:"true"` // Identifies Domain from which to read info by name.
}
type DeleteResponse struct {
}

//const ReadUrl = "?format=json&object=department&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain string `json:"domain"` // Identifies Domain for which to list Departments.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle string

//const ListUrl = "?format=json&object=department&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain" required:"true"` // Identifies Domain for which to list Departments.
}
type ListResponse []ListResponseSingle
type ListResponseSingle string

//const UpdateUrl = "?format=json&object=department&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain string `json:"domain"` // Identifies Domain from which to read info by name.
}
type UpdateResponse struct {
}
