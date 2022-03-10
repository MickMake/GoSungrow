package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
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

		if w.CheckCache(endpoint) {
			w.Body, w.Error = w.CacheRead(endpoint)
			if w.Error != nil {
				break
			}
		} else {
			w.Body, w.Error = w.getApi(endpoint)
			if w.Error != nil {
				break
			}
			w.Error = w.CacheWrite(endpoint, w.Body)
			if w.Error != nil {
				break
			}
		}

		if len(w.Body) == 0 {
			w.Error = errors.New("empty httpResponse")
			break
		}

		endpoint = endpoint.SetResponse(w.Body)
		w.Error = endpoint.IsResponseValid()
		if w.Error != nil {
			_ = w.CacheRemove(endpoint)
			// fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}
	}

	if w.Error != nil {
		endpoint = endpoint.SetError("%s", w.Error)
	}
	return endpoint
}

func (w *Web) getApi(endpoint EndPoint) ([]byte, error) {
	for range Only.Once {
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
	}

	return w.Body, w.Error
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

// CheckCache Retrieves cache data from a local file.
func (w *Web) CheckCache(endpoint EndPoint) bool {
	var ok bool
	for range Only.Once {
		fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())

		var f os.FileInfo
		f, w.Error = os.Stat(fn)
		if w.Error != nil {
			if os.IsNotExist(w.Error) {
				w.Error = nil
			}
			break
		}

		if f.IsDir() {
			w.Error = errors.New("file is a directory")
			break
		}

		duration := w.GetCacheTimeout()
		then := f.ModTime()
		then = then.Add(duration)
		now := time.Now()
		if then.Before(now) {
			break
		}

		ok = true
	}

	return ok
}

// CacheRead Retrieves cache data from a local file.
func (w *Web) CacheRead(endpoint EndPoint) ([]byte, error) {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileRead(fn)
}

// CacheRemove Removes a cache file.
func (w *Web) CacheRemove(endpoint EndPoint) error {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.FileRemove(fn)
}

// CacheWrite Saves cache data to a file path.
func (w *Web) CacheWrite(endpoint EndPoint, data []byte) error {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileWrite(fn, data, output.DefaultFileMode)
}
