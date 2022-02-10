package conference

const ConferenceCdrCountUrl = "?format=json&object=conferencecdr&action=count"

type ConferenceCdrCountRequest struct {
}
type ConferenceCdrCountResponse struct {
	Total int `json:"total"` // Total number of ConferenceCdrs for the specified ApiAppKey/User.
}

const ConferenceCdrReadUrl = "?format=json&object=conferencecdr&action=read"

type ConferenceCdrReadRequest struct {
}
type ConferenceCdrReadResponse struct {
}
