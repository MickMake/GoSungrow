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
func (p *Path) FileRead(ref interface{}) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if p.name == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.Open(p.name)
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
func (p *Path) FileWrite(ref interface{}, perm os.FileMode) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if p.name == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.OpenFile(p.name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", p.name, err))
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
func (p *Path) PlainFileRead() ([]byte, error) {
	var data []byte
	var err error
	var once sync.Once
	once.Do(func() {
		if p.name == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.Open(p.name)
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
func (p *Path) PlainFileWrite(data []byte, perm os.FileMode) error {
	var err error
	var once sync.Once
	once.Do(func() {
		if p.name == "" {
			err = errors.New("empty file")
			return
		}

		var f *os.File
		f, err = os.OpenFile(p.name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", p.name, err))
			return
		}
		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		_, err = f.Write(data)
	})

	return err
}

// FileRemove Removes a file path.
func (p *Path) FileRemove() error {
	var err error
	var once sync.Once
	once.Do(func() {
		if p.name == "" {
			err = errors.New("empty file")
			return
		}

		var f os.FileInfo
		f, err = os.Stat(p.name)
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

		err = os.Remove(p.name)
	})

	return err
}
