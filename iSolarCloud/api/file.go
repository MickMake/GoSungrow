package api

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const DefaultFileMode = 0644

func (ep *EndPointStruct) GetCsvFilename(postfix string) string {
	if postfix == "" {
		return fmt.Sprintf("%s_%s.csv", ep.Area, ep.Name)
	}
	return fmt.Sprintf("%s_%s-%s.csv", ep.Area, ep.Name, postfix)
}

func (ep *EndPointStruct) GetFilename(postfix string) string {
	if postfix == "" {
		return fmt.Sprintf("%s_%s.json", ep.Area, ep.Name)
	}
	return fmt.Sprintf("%s_%s-%s.json", ep.Area, ep.Name, postfix)
}

func (ep *EndPointStruct) GetFilePath() string {
	var ret string

	for range Only.Once {
		var hd string
		hd, ep.Error = os.UserHomeDir()
		if ep.Error != nil {
			break
		}

		ret = filepath.Join(hd, ".GoSungrow", ep.GetFilename(""))
	}

	return ret
}

//
//
//

// Retrieves data from a local file.
func (ep *EndPointStruct) FileExists(fn string) bool {
	var ok bool

	for range Only.Once {
		if fn == "" {
			fn = ep.GetFilename("")
			if ep.Error != nil {
				break
			}
		}

		var f os.FileInfo
		f, ep.Error = os.Stat(fn)
		if ep.Error != nil {
			if os.IsNotExist(ep.Error) {
				ep.Error = nil
			}
			break
		}
		if f.IsDir() {
			ep.Error = errors.New("file is a directory")
			break
		}

		ok = true
	}

	return ok
}

// Retrieves data from a local file.
func (ep *EndPointStruct) FileRead(fn string, ref interface{}) error {
	// var ret []byte

	for range Only.Once {
		if fn == "" {
			fn = ep.GetFilename("")
			if ep.Error != nil {
				break
			}
		}

		var f *os.File
		f, ep.Error = os.Open(fn)
		if ep.Error != nil {
			if os.IsNotExist(ep.Error) {
				ep.Error = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()

		ep.Error = json.NewDecoder(f).Decode(&ref)
	}

	// for range Only.Once {
	//	fn := ep.GetFilename()
	//	if ep.Error != nil {
	//		break
	//	}
	//
	//	ret, ep.Error = os.FileRead(fn)
	//	if ep.Error != nil {
	//		break
	//	}
	// }

	return ep.Error
}

// Saves data to a file path.
func (ep *EndPointStruct) FileWrite(fn string, ref interface{}, perm os.FileMode) error {
	for range Only.Once {
		if fn == "" {
			fn = ep.GetFilename("")
			if ep.Error != nil {
				break
			}
		}

		var f *os.File
		f, ep.Error = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if ep.Error != nil {
			ep.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, ep.Error))
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		ep.Error = json.NewEncoder(f).Encode(ref)

		// fn := ep.GetFilename()
		// if ep.Error != nil {
		//	break
		// }
		//
		// ep.Error = os.FileWrite(fn, data, perm)
		// if ep.Error != nil {
		//	break
		// }
	}

	return ep.Error
}

//
//
//

func (ep *EndPointStruct) GetCacheFilename(request interface{}) string {
	postfix := apiReflect.GetFingerprint(request)
	return fmt.Sprintf("%s_%s-%s.json", ep.Area, ep.Name, postfix)
}

func (ep *EndPointStruct) GetCacheFilePath(request interface{}) string {
	return filepath.Join(ep.ApiRoot.GetCacheDir(), ep.GetCacheFilename(request))
}

// Retrieves cache data from a local file.
func (ep *EndPointStruct) IsCacheFileOk(request interface{}) bool {
	var ok bool
	for range Only.Once {
		fn := ep.GetCacheFilePath(request)

		var f os.FileInfo
		f, ep.Error = os.Stat(fn)
		if ep.Error != nil {
			if os.IsNotExist(ep.Error) {
				ep.Error = nil
			}
			break
		}

		if f.IsDir() {
			ep.Error = errors.New("file is a directory")
			break
		}

		duration := ep.ApiRoot.GetCacheTimeout()
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

// Retrieves cache data from a local file.
func (ep *EndPointStruct) CacheRead(request interface{}, ref interface{}) error {
	for range Only.Once {
		fn := ep.GetCacheFilePath(request)

		ep.Error = ep.FileRead(fn, ref)
	}

	return ep.Error
}

// Saves cache data to a file path.
func (ep *EndPointStruct) CacheWrite(request interface{}, ref interface{}) error {
	for range Only.Once {
		fn := ep.GetCacheFilePath(request)

		var f *os.File
		f, ep.Error = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, DefaultFileMode)
		if ep.Error != nil {
			ep.Error = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, ep.Error))
			break
		}

		//goland:noinspection GoUnhandledErrorResult
		defer f.Close()
		ep.Error = json.NewEncoder(f).Encode(ref)
	}

	return ep.Error
}
