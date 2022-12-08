package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"net/url"
)


type EndPointUrl struct {
	EndPoint *url.URL	`json:"endpoint"`
	Error error	`json:"error"`
}


func (u EndPointUrl) AppendPath(endpoint string) EndPointUrl {
	var ret EndPointUrl
	for range Only.Once {
		endpoint = fmt.Sprintf("%s%s", u.String(), endpoint)
		// ret.URL, ret.Error = url.Parse(endpoint)
		// u, e := url.Parse(endpoint)
		// ret.URL = *u
		// ret.Error = e
		ret = SetUrl(endpoint)
	}
	return ret
}

func (u *EndPointUrl) IsValid() error {
	var err error
	for range Only.Once {
		if u == nil {
			err = errors.New("empty url")
			break
		}
		if u.EndPoint == nil {
			err = errors.New("empty url")
			break
		}
		if u.EndPoint.String() == "" {
			err = errors.New("empty url")
			break
		}
		err = u.Error
	}
	return err
}

func (u EndPointUrl) String() string {
	if u.EndPoint == nil {
		return ""
	}
	return u.EndPoint.String()
}

func (u *EndPointUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Url    string `json:"url"`
	}{
		Url:    u.String(),
	})
}

func (u *EndPointUrl) UnmarshalJSON(data []byte) error {
	// type Alias EndPointUrl
	var err error
	for range Only.Once {
		aux := &struct {
			EndPoint string	`json:"url"`
			// Scheme string `json:"schema"`
			// Host   string `json:"host"`
			// Path   string `json:"path"`
			// Url    string `json:"url"`
		}{
			EndPoint:    u.String(),
		}
		err = json.Unmarshal(data, &aux)
		if err != nil {
			break
		}
	}
	return err
}


func SetUrl(endpoint string) EndPointUrl {
	var ret EndPointUrl
	// ret.URL, ret.Error = url.Parse(endpoint)
	u, e := url.Parse(endpoint)
	ret.EndPoint = u
	ret.Error = e
	return ret
}
