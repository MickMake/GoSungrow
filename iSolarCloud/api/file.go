package api

import (
	"github.com/MickMake/GoUnify/Only"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/iSolarCloud/api/GoStruct/reflection"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/cmdPath"
	"os"
	"path/filepath"
)


// func (ep *EndPointStruct) ApiSetFilenamePrefix2(ref interface{}, format string, args ...interface{}) string {
// 	f := strings.Join(GoStruct.GetStructValuesAsArray(ref), "-")
// 	fmt.Printf("[%s]\n", f)
// 	if format != "" {
// 		ep.FileNamePrefix = fmt.Sprintf(format, args...)
// 		// ep.FileNamePrefix = fmt.Sprintf("%s_%s-%s", ep.Area, ep.Name, ep.FileNamePrefix)
// 		ep.FileNamePrefix = string(ep.Area) + "_" + string(ep.Name) + "-" + ep.FileNamePrefix
// 	} else {
// 		ep.FileNamePrefix = string(ep.Area) + "_" + string(ep.Name)
// 	}
// 	return ep.FileNamePrefix
// }

func (ep *EndPointStruct) ApiSetFilenamePrefix(format string, args ...interface{}) string {
	if format != "" {
		ep.FileNamePrefix = fmt.Sprintf(format, args...)
		// ep.FileNamePrefix = fmt.Sprintf("%s_%s-%s", ep.Area, ep.Name, ep.FileNamePrefix)
		ep.FileNamePrefix = string(ep.Area) + "_" + string(ep.Name) + "-" + ep.FileNamePrefix
	} else {
		ep.FileNamePrefix = string(ep.Area) + "_" + string(ep.Name)
	}
	return ep.FileNamePrefix
}

// func (ep *EndPointStruct) ApiGetCsvFilename() string {
// 	if ep.FileNamePrefix == "" {
// 		ep.FileNamePrefix = ep.ApiSetFilenamePrefix("")
// 	}
// 	return ep.FileNamePrefix + ".csv"
// }
//
// func (ep *EndPointStruct) ApiGetTableFilename() string {
// 	if ep.FileNamePrefix == "" {
// 		ep.FileNamePrefix = ep.ApiSetFilenamePrefix("")
// 	}
// 	return ep.FileNamePrefix + "-table.txt"
// }
//
// func (ep *EndPointStruct) ApiGetListFilename() string {
// 	if ep.FileNamePrefix == "" {
// 		ep.FileNamePrefix = ep.ApiSetFilenamePrefix("")
// 	}
// 	return ep.FileNamePrefix + ".txt"
// }

func (ep *EndPointStruct) ApiGetJsonFilename() string {
	if ep.FileNamePrefix == "" {
		ep.FileNamePrefix = ep.ApiSetFilenamePrefix("")
	}
	return ep.FileNamePrefix + ".json"
}

// func (ep *EndPointStruct) GetImageFilename() string {
// 	if ep.FileNamePrefix == "" {
// 		ep.FileNamePrefix = ep.ApiSetFilenamePrefix("")
// 	}
// 	return ep.FileNamePrefix + ".png"
// }

func (ep *EndPointStruct) GetFilePath() string {
	var ret string

	for range Only.Once {
		var hd string
		hd, ep.Error = os.UserHomeDir()
		if ep.Error != nil {
			break
		}

		ret = filepath.Join(hd, ".GoSungrow", ep.ApiGetJsonFilename())
	}

	return ret
}


// FileExists Checks for existence of a local file.
func (ep *EndPointStruct) FileExists(fn string) bool {
	var ok bool

	for range Only.Once {
		if fn == "" {
			fn = ep.ApiGetJsonFilename()
			if ep.Error != nil {
				break
			}
		}

		p := cmdPath.NewPath(fn)
		if p.DirExists() {
			ep.Error = errors.New("file is a directory")
			ok = false
			break
		}
		if p.FileExists() {
			ok = true
			break
		}

		ok = false
	}

	return ok
}

// ApiReadDataFile - Retrieves data from a local file.
func (ep *EndPointStruct) ApiReadDataFile(ref interface{}) error {
	return output.FileRead(ep.ApiGetJsonFilename(), ref)
}

// ApiWriteDataFile - Saves data to a file path.
func (ep *EndPointStruct) ApiWriteDataFile(ref interface{}) error {
	return output.FileWrite(ep.ApiGetJsonFilename(), ref, output.DefaultFileMode)
}

// ApiRemoveDataFile - Saves data to a file path.
func (ep *EndPointStruct) ApiRemoveDataFile() error {
	return output.FileRemove(ep.ApiGetJsonFilename())
}

func (ep *EndPointStruct) ApiCacheFilename(request interface{}) string {
	postfix := reflection.GetFingerprint(request)
	return fmt.Sprintf("%s_%s-%s.json", ep.Area, ep.Name, postfix)
}

func (ep *EndPointStruct) ApiFingerprint(request interface{}) string {
	return reflection.GetFingerprint(request)
}
