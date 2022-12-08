package api

import (
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdPath"
	"io"
	"path/filepath"
	"time"

	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
			w.Error = errors.New("Sungrow API EndPoint not yet implemented")
			fmt.Println(w.Error)
			break
		}

		isCached := false
		if w.WebCacheCheck(endpoint) {
			isCached = true
		}


		if isCached {
			w.Body, w.Error = w.WebCacheRead(endpoint)
			if w.Error != nil {
				break
			}

		} else {
			w.Body, w.Error = w.getApi(endpoint)
			if w.Error != nil {
				break
			}
		}


		if len(w.Body) == 0 {
			w.Error = errors.New("empty http response")
			break
		}
		endpoint = endpoint.SetResponse(w.Body)
		if endpoint.GetError() != nil {
			w.Error = endpoint.GetError()
			break
		}

		w.Error = endpoint.IsResponseValid()
		if w.Error != nil {
			_ = w.WebCacheRemove(endpoint)
			// fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}


		if isCached {

		} else {
			w.Error = w.WebCacheWrite(endpoint, w.Body)
			if w.Error != nil {
				break
			}
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
		w.Error = GoStruct.VerifyOptionsRequired(request)
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
		var j []byte
		j, w.Error = json.Marshal(request)
		if w.Error != nil {
			break
		}

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

		w.Body, w.Error = io.ReadAll(w.httpResponse.Body)
		if w.Error != nil {
			break
		}
	}

	return w.Body, w.Error
}

func (w *Web) SetCacheDir(basedir string) error {
	for range Only.Once {
		w.cacheDir = filepath.Join(basedir)

		p := cmdPath.NewPath(basedir)
		if p.DirExists() {
			break
		}

		w.Error = p.MkdirAll()
		if w.Error != nil {
			break
		}

		// _, w.Error = os.Stat(w.cacheDir)
		// if w.Error != nil {
		// 	if os.IsNotExist(w.Error) {
		// 		w.Error = nil
		// 	}
		// 	break
		// }
		//
		// w.Error = os.MkdirAll(w.cacheDir, 0700)
		// if w.Error != nil {
		// 	break
		// }
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

// WebCacheCheck Retrieves cache data from a local file.
func (w *Web) WebCacheCheck(endpoint EndPoint) bool {
	var ok bool
	for range Only.Once {
		// fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
		//
		// var f os.FileInfo
		// f, w.Error = os.Stat(fn)
		// if w.Error != nil {
		// 	if os.IsNotExist(w.Error) {
		// 		w.Error = nil
		// 	}
		// 	break
		// }
		//
		// if f.IsDir() {
		// 	w.Error = errors.New("file is a directory")
		// 	break
		// }

		p := cmdPath.NewPath(w.cacheDir, endpoint.CacheFilename())
		if p.DirExists() {
			w.Error = errors.New("file is a directory")
			ok = false
			break
		}
		if !p.FileExists() {
			ok = false
			break
		}

		duration := w.GetCacheTimeout()
		then := p.ModTime()
		then = then.Add(duration)
		now := time.Now()
		if then.Before(now) {
			ok = false
			break
		}

		ok = true
	}

	return ok
}

// WebCacheRead Retrieves cache data from a local file.
func (w *Web) WebCacheRead(endpoint EndPoint) ([]byte, error) {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileRead(fn)
}

// WebCacheRemove Removes a cache file.
func (w *Web) WebCacheRemove(endpoint EndPoint) error {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.FileRemove(fn)
}

// WebCacheWrite Saves cache data to a file path.
func (w *Web) WebCacheWrite(endpoint EndPoint, data []byte) error {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileWrite(fn, data, output.DefaultFileMode)
}


// PointCacheCheck Retrieves cache data from a local file.
func (w *Web) PointCacheCheck(data DataMap) bool {
	var ok bool
	for range Only.Once {
		p := cmdPath.NewPath(w.cacheDir, "Points.json")
		if p.DirExists() {
			w.Error = errors.New("file is a directory")
			ok = false
			break
		}
		if p.FileExists() {
			ok = true
			break
		}

		duration := w.GetCacheTimeout()
		then := p.ModTime()
		then = then.Add(duration)
		now := time.Now()
		if then.Before(now) {
			break
		}

		ok = true
	}

	return ok
}

// PointCacheRead Retrieves cache data from a local file.
func (w *Web) PointCacheRead(endpoint EndPoint) ([]byte, error) {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileRead(fn)
}

// PointCacheWrite Saves cache data to a file path.
func (w *Web) PointCacheWrite(endpoint EndPoint, data []byte) error {
	fn := filepath.Join(w.cacheDir, endpoint.CacheFilename())
	return output.PlainFileWrite(fn, data, output.DefaultFileMode)
}
