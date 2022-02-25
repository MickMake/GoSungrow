package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"os"
	"path/filepath"
	"time"

	// "GoSungrow/iSolarCloud/AppService/login"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Web struct {
	Url   EndPointUrl
	Body  []byte
	Error error

	cacheDir     string
	cacheTimeout time.Duration
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
			w.Error = errors.New("SUNGRO API EndPoint not yet implemented")
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

		u := endpoint.GetUrl()
		w.Error = u.IsValid()
		if w.Error != nil {
			break
		}

		postUrl := w.Url.AppendPath(u.String()).String()
		j, _ := json.Marshal(request)

		w.httpResponse, w.Error = http.Post(postUrl, "application/json", bytes.NewBuffer(j))
		if w.Error != nil {
			break
		}

		if w.httpResponse.StatusCode == 401 {
			w.Error = errors.New(w.httpResponse.Status)
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
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

		// fmt.Printf("URL: %s\n\n", postUrl)
		// fmt.Printf("REQUEST: %s\n\n", j)
		// fmt.Printf("RESPONSE: %s\n\n", w.Body)

		endpoint = endpoint.SetResponse(w.Body)
		// w.Error = endpoint.GetError()
		// if w.Error != nil {
		//	break
		// }

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

func (w *Web) SetCacheDir(basedir string) error {
	for range Only.Once {
		w.cacheDir = filepath.Join(basedir)
		_, w.Error = os.Stat(w.cacheDir)
		if os.IsExist(w.Error) {
			w.Error = nil
			break
		}

		w.Error = os.MkdirAll(w.cacheDir, 0700)
		if w.Error != nil {
			break
		}
	}

	return w.Error
}

func (w *Web) GetCacheDir() string {
	return w.cacheDir
}

func (w *Web) SetCacheTimeout(duration time.Duration) {
	w.cacheTimeout = duration
}

func (w *Web) GetCacheTimeout() time.Duration {
	return w.cacheTimeout
}
