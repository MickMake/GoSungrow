package api


type Json string


func (req Json) String() string {
	return string(req)
}
