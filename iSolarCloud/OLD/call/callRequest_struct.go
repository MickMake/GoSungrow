package call

const CallRequestCountUrl = "?format=json&object=callrequest&action=count"

type CallRequestCountRequest struct {
}
type CallRequestCountResponse struct {
	Total int `json:"total"` // Total number of CallRequests for the specified ApiAppKey/User.
}

const CallRequestReadUrl = "?format=json&object=callrequest&action=read"

type CallRequestReadRequest struct {
}
type CallRequestReadResponse struct {
}
