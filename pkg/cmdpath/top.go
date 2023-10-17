package cmdpath

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// FileRead Retrieves data from a local file.
func FileRead(fn string, ref interface{}) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if fn == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			return
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		err = json.NewDecoder(f).Decode(&ref)
	})

	// var once sync.Once
	// once.Do(func() {
	//	fn := ep.GetFilename()
	//	if err != nil {
	//		return
	//	}
	//
	//	ret, err = os.FileRead(fn)
	//	if err != nil {
	//		return
	//	}
	// }

	return err
}

// FileWrite Saves data to a file path.
func FileWrite(fn string, ref interface{}, perm os.FileMode) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if fn == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			return
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()
		err = json.NewEncoder(f).Encode(ref)

		// fn := ep.GetFilename()
		// if err != nil {
		//	return
		// }
		//
		// err = os.FileWrite(fn, data, perm)
		// if err != nil {
		//	return
		// }
	})

	return err
}

// PlainFileRead Retrieves data from a local file.
func PlainFileRead(fn string) ([]byte, error) {
	var data []byte
	var err error
	var once sync.Once
	once.Do(func() {
		if fn == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.Open(fn)
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			return
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		data, err = ioutil.ReadAll(f)
	})

	return data, err
}

// PlainFileWrite Saves data to a file path.
func PlainFileWrite(fn string, data []byte, perm os.FileMode) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if fn == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", fn, err))
			return
		}
		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		_, err = f.Write(data)
	})

	return err
}

// FileRemove Removes a file path.
func FileRemove(fn string) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if fn == "" {
			err = errors.New("empty file")
			return
		}

		var f os.FileInfo
		f, err = os.Stat(fn)
		if os.IsNotExist(err) {
			err = nil
			return
		}
		if err != nil {
			return
		}
		if f.IsDir() {
			err = errors.New("file is a directory")
			return
		}

		err = os.Remove(fn)
	})

	return err
}
