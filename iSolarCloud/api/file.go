package api

import (
	"GoSungrow/Only"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func (ep *EndPointStruct) GetFilename() string {
	return fmt.Sprintf("%s_%s.json", ep.Area, ep.Name)
}

func (ep *EndPointStruct) GetFilepath() string {
	var ret string

	for range Only.Once {
		var hd string
		hd, ep.Error = os.UserHomeDir()
		if ep.Error != nil {
			break
		}

		ret = filepath.Join(hd, ".GoSungrow", ep.GetFilename())
	}

	return ret
}

// Retrieves data from a local file.
func (ep *EndPointStruct) FileRead(fn string, ref interface{}) error {
	//var ret []byte

	for range Only.Once {
		if fn == "" {
			fn = ep.GetFilename()
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

	//for range Only.Once {
	//	fn := ep.GetFilename()
	//	if ep.Error != nil {
	//		break
	//	}
	//
	//	ret, ep.Error = os.FileRead(fn)
	//	if ep.Error != nil {
	//		break
	//	}
	//}

	return ep.Error
}

const DefaultFileMode = 0644

// Saves a token to a file path.
func (ep *EndPointStruct) FileWrite(fn string, ref interface{}, perm os.FileMode) error {
	for range Only.Once {
		if fn == "" {
			fn = ep.GetFilename()
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

		//fn := ep.GetFilename()
		//if ep.Error != nil {
		//	break
		//}
		//
		//ep.Error = os.FileWrite(fn, data, perm)
		//if ep.Error != nil {
		//	break
		//}
	}

	return ep.Error
}
