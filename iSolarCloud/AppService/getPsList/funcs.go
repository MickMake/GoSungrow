package getPsList

import (
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/MickMake/GoUnify/Only"
)


// type Device struct {
// 	PsId                   valueTypes.PsId
// 	PsType                 valueTypes.Integer
// 	PsName                 valueTypes.String
// 	PsShortName            valueTypes.String
// 	PsHolder               valueTypes.String
// 	PsStatus               valueTypes.Bool
// 	PsFaultStatus          valueTypes.Integer
// 	PsHealthStatus         valueTypes.Integer
// }
//
// type Devices []Device
//
// func (e *ResultData) GetPsDevices() []Common.Device {
// 	return e.PageList
// }

func (e *ResultData) GetPsIds() []valueTypes.PsId {
	var ret []valueTypes.PsId
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsId)
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsName() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsName.Value())
			}
		}
	}
	return ret
}

func (e *ResultData) GetPsSerial() []string {
	var ret []string
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.PsId.Value() != 0 {
				ret = append(ret, p.PsShortName.Value())
			}
		}
	}
	return ret
}

func (e *EndPoint) GetPsIds() []valueTypes.PsId {
	return e.Response.ResultData.GetPsIds()
}
