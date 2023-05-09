package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"io"
	"os"
)


// FileRead Retrieves data from a local file.
func FileRead(fn string, ref interface{}) error {
	var err error

	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		err = json.NewDecoder(f).Decode(&ref)
	}

	return err
}

// FileWrite Saves data to a file path.
func FileWrite(fn string, ref interface{}, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()
		err = json.NewEncoder(f).Encode(ref)
	}

	return err
}

// PlainFileRead Retrieves data from a local file.
func PlainFileRead(fn string) ([]byte, error) {
	var data []byte
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		data, err = io.ReadAll(f)
	}

	return data, err
}

// PlainFileWrite Saves data to a file path.
func PlainFileWrite(fn string, data []byte, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			break
		}
		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		_, err = f.Write(data)
	}

	return err
}

// FileRemove Removes a file path.
func FileRemove(fn string) error {
	var err error
	for range Only.Once {
		if fn == "" {
			err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			err = nil
			break
		}
		if err != nil {
			break
		}
		if f.IsDir() {
			err = errors.New("file is a directory")
			break
		}

		err = os.Remove(fn)
	}

	return err
}

// FileExists - Checks if a file exists.
func FileExists(fn string) bool {
	var yes bool
	for range Only.Once {
		var err error
		if fn == "" {
			// err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			break
		}
		if f.IsDir() {
			// err = errors.New("file is a directory")
			break
		}

		yes = true
	}

	return yes
}

// DirExists - Checks if a directory exists.
func DirExists(fn string) bool {
	var yes bool
	for range Only.Once {
		var err error
		if fn == "" {
			// err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			break
		}
		if !f.IsDir() {
			// err = errors.New("file is a directory")
			break
		}

		yes = true
	}

	return yes
}

// Mkdir - Create dir.
func Mkdir(fn string) bool {
	var yes bool
	for range Only.Once {
		var err error
		if fn == "" {
			// err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			err = os.MkdirAll(fn, 755)
			if err != nil {
				yes = true
			}
			break
		}
		if err != nil {
			break
		}
		if !f.IsDir() {
			// err = errors.New("file is a directory")
			break
		}

		yes = true
	}

	return yes
}
