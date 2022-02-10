package contact

import (
	"GoSungro/iSolarCloud/web"
	"net/url"
)

type Contact struct {
	count  Count
	create Create
	delete Delete
	list   List
	read   Read
	update Update

	Web   *web.Web
	Error error
}

//const CountUrl = "?format=json&object=contact&action=count"
type Count struct {
	Url      *url.URL
	Request  CountRequest
	Response CountResponse
	Error    error
}
type CountRequest struct {
	Domain    string `json:"domain" required:"true"` // Identifies Domain from which to count Contacts.
	User      string `json:"user" required:"true"`   // Identifies User from which to count Contacts.
	FirstName string `json:"first_name"`             // Filters count based on Caontacts' first name.
	LastName  string `json:"last_name"`              // Filters count based on Caontacts' last name.
	ContactId string `json:"contact_id"`             // Contact_id hash of the contact.
}
type CountResponse struct {
	Total int `json:"total"` // Total number of Departments for the specified ApiAppKey/User.
}

//const CreateUrl = "?format=json&object=subscriber&action=create&directory_match=contacts"
type Create struct {
	Url      *url.URL
	Request  CreateRequest
	Response CreateResponse
	Error    error
}
type CreateRequest struct {
	Domain    string `json:"domain" required:"true"`     // Identifies new Contact's Domain.
	User      string `json:"user" required:"true"`       // Specifies user/extension for the User whose Contact you want.
	FirstName string `json:"first_name" required:"true"` // Contact's first name
	LastName  string `json:"last_name" required:"true"`  // Contact's last name
	HomePhone string `json:"home_phone"`                 // Contact's home number
	CellPhone string `json:"cell_phone"`                 // Contact's mobile number
	WorkPhone string `json:"work_phone"`                 // Contact's work number
	Email     string `json:"email"`                      // Contact's email address
	Fax       string `json:"fax"`                        // Fax number
}
type CreateResponse struct {
}

//const DeleteUrl = "?format=json&object=contact&action=delete"
type Delete struct {
	Url      *url.URL
	Request  DeleteRequest
	Response DeleteResponse
	Error    error
}
type DeleteRequest struct {
	Domain    string `json:"domain" required:"true"` // Identifies Contact to delete by Domain.
	User      string `json:"user" required:"true"`   // Specifies user/extension for the User whose Contact you want.
	FirstName string `json:"first_name"`             // Contact's first name.
	LastName  string `json:"last_name"`              // Contact's last name.
	ContactId string `json:"contact_id"`             // Hash that is contact's id.
}
type DeleteResponse struct {
}

// [
//    {
//        "last_name": "Angus mobile",
//        "first_name": "James",
//        "middle_name": "",
//        "company": "",
//        "work_phone": "",
//        "cell_phone": "0428048169",
//        "fax": "",
//        "email": "",
//        "home_phone": "",
//        "tags": "shared_domain,,post,",
//        "ts": "2019-06-10 04:53:22",
//        "contact_id": "a2bf6d39d2feaa1ee7469693506e8cce",
//        "shared": true
//    }
// ]

//const ReadUrl = "?format=json&object=contact&action=read"
type Read struct {
	Url      *url.URL
	Request  ReadRequest
	Response ReadResponse
	Error    error
}
type ReadRequest struct {
	Domain        string `json:"domain" required:"true"` // Identifies Domain from which to read Contacts.
	User          string `json:"user" required:"true"`   // Identifies User from which to read Contacts.
	FirstName     string `json:"first_name"`             // Filters search by Contacts' first name.
	LastName      string `json:"last_name"`              // Filters search by Contact's last_name.
	Tags          string `json:"tags"`                   // Ability to search by a specific tag or group of contacts.
	IncludeDomain string `json:"includeDomain"`          // (yes/no) Include the presence data set for non residential domains.
	Limit         string `json:"limit"`                  // Specifies limit for number of Contacts returned.
	Department    string `json:"department"`             // Filters search by Contacts' department.
	Order         string `json:"order"`                  // Specifies parameter by which the output is ordered.
	ContactId     string `json:"contact_id"`             // Specific id of the contact to be read.
}
type ReadResponse []ReadResponseSingle
type ReadResponseSingle struct {
	Domain     string `json:"domain"`      //
	User       string `json:"user"`        //
	LastName   string `json:"last_name"`   //
	FirstName  string `json:"first_name"`  //
	MiddleName string `json:"middle_name"` //
	Company    string `json:"company"`     //
	WorkPhone  string `json:"work_phone"`  //
	CellPhone  string `json:"cell_phone"`  //
	Email      string `json:"email"`       //
	HomePhone  string `json:"home_phone"`  //
	Tags       string `json:"tags"`        //
	Ts         string `json:"ts"`          //
}

//const ListUrl = "?format=json&object=contact&action=list"
type List struct {
	Url      *url.URL
	Request  ListRequest
	Response ListResponse
	Error    error
}
type ListRequest struct {
	Domain string `json:"domain"` // Identifies Domain for which to list Departments.
	User   string `json:"user"`   //
}
type ListResponse []ListResponseSingle
type ListResponseSingle string

//const UpdateUrl = "?format=json&object=contact&action=update"
type Update struct {
	Url      *url.URL
	Request  UpdateRequest
	Response UpdateResponse
	Error    error
}
type UpdateRequest struct {
	Domain     string `json:"domain" required:"true"`     // Identifies Contacts to update by Domain.
	User       string `json:"user" required:"true"`       // Identifies Contacts to update by User/Extension.
	FirstName  string `json:"first_name" required:"true"` // Identifies Contacts to update by first name.
	LastName   string `json:"last_name" required:"true"`  // Identifies Contacts to update by last name.
	Company    string `json:"company"`                    // New value for Contact.
	WorkPhone  string `json:"work_phone"`                 // New value for Contact.
	CellPhone  string `json:"cell_phone"`                 // New value for Contact.
	Fax        string `json:"fax"`                        // New value for Contact.
	Email      string `json:"email"`                      // New value for Contact.
	HomePhone  string `json:"home_phone"`                 // New value for Contact.
	TimeAnswer string `json:"time_answer"`                // New value for Contact.
	Tags       string `json:"tags"`                       // New value for Contact.
	Ts         string `json:"ts"`                         // New value for Contact.
	ContactId  string `json:"contact_id"`                 // Contact_id hash for contact, should not change, it is the index.
}
type UpdateResponse struct {
}
