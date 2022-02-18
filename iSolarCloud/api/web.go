package api

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/api/apiReflect"
	// "GoSungro/iSolarCloud/sungro/AppService/login"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Web struct {
	Url          EndPointUrl
	Body         []byte
	Error        error

	retry        int
	client       http.Client
	httpRequest  *http.Request
	httpResponse *http.Response
}

func (w *Web) SetUrl(u string) error {
	w.Url = SetUrl(u)
	// w.Error = w.Url.Error
	return w.Error
}

func (w *Web) AppendUrl(endpoint string) EndPointUrl {
	return w.Url.AppendPath(endpoint)
}

func (w *Web) Get(endpoint EndPoint) EndPoint {
	for range Only.Once {
		w.Error = w.Url.IsValid()
		if w.Error != nil {
			w.Error = errors.New("SUNGRO API URL is invalid")
			break
		}

		request := endpoint.RequestRef()
		w.Error = apiReflect.VerifyOptionsRequired(request)
		if w.Error != nil {
			break
		}

		w.Error = endpoint.IsRequestValid()
		if w.Error != nil {
			break
		}

		j, _ := json.Marshal(request)
		postUrl := w.Url.AppendPath(endpoint.GetUrl().String()).String()

		w.httpResponse, w.Error = http.Post(postUrl, "application/json", bytes.NewBuffer(j))
		if w.Error != nil {
			break
		}

		if w.httpResponse.StatusCode == 401 {
			w.Error = errors.New(w.httpResponse.Status)
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer w.httpResponse.Body.Close()
		if w.Error != nil {
			break
		}

		if w.httpResponse.StatusCode != 200 {
			w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
			break
		}

		w.Body, w.Error = ioutil.ReadAll(w.httpResponse.Body)
		if w.Error != nil {
			break
		}

		if len(w.Body) == 0 {
			w.Error = errors.New("empty httpResponse")
			break
		}

		endpoint = endpoint.SetResponse(w.Body)

		w.Error = endpoint.IsResponseValid()
		if w.Error != nil {
			// fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}
	}

	if w.Error != nil {
		endpoint = endpoint.SetError("%s", w.Error)
	}
	return endpoint
}
