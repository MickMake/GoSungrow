package mmGit

import (
	"GoSungro/Only"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)


func IsDirExists(path string) (bool, error) {
	var ok bool
	var err error

	for range Only.Once {
		if path == "" {
			err = errors.New("empty git dir path")
			break
		}

		var fi os.FileInfo
		fi, err = os.Stat(path)
		if os.IsNotExist(err) {
			err = nil
			break
		}

		if !fi.IsDir() {
			err = errors.New("git path is not a dir")
			break
		}

		ok = true
	}

	return ok, err
}

func PrintError(err error) {
	if err == nil {
		fmt.Println("OK")
		return
	}
	fmt.Printf("\nERROR: %s\n", err)
}

func GetHash(data string) string {
	var ret string

	for range Only.Once {
			h := md5.New()
			h.Write([]byte(data))
			ret = hex.EncodeToString(h.Sum(nil))
	}

	return ret
}

func WriteTempFile(path string, data string) (string, error) {
	var fn string
	var err error

	for range Only.Once {
		fn = fmt.Sprintf("%s/%s", os.TempDir(), path)
		fmt.Printf("Writing file %s ...", fn)
		var fh *os.File
		fh, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
		defer fh.Close()

		_, err = fh.Write([]byte(data))
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
		fmt.Println("OK")

		//fmt.Printf("%v\n", data)
	}

	return fn, err
}
