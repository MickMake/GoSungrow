package nullEndPoint

import (
	"GoSungro/Only"
	"errors"
)


const Url = ""


type RequestData struct {
}

func (rd *RequestData) IsValid() error {
	var err error
	for range Only.Once {
		if rd == nil {
			err = errors.New("empty device type")
			break
		}
		// err = api.CheckString("PsId", rd.PsId)
		// if err != nil {
		// 	err = errors.New("empty device type")
		// 	break
		// }
	}
	return err
}


type ResultData struct {
}
