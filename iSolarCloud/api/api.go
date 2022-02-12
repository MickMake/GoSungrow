// go install github.com/ungerik/pkgreflect@latest

package api


var ApiFoo = (*Api)(nil)

func LoadMe() bool {
	return true
}

type Api struct {
	User     string
	Password string
	Url      string
}

var _api *Api

func init() {
	_api = &Api{}
}

func SetUrl(url string) {
	_api.Url = url
}

func SetBasicAuth(user, password string) {
	_api.User = user
	_api.Password = password
}
