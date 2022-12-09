package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/queryUnitList"
	"github.com/MickMake/GoUnify/Only"
)


func (sg *SunGrow) MetaUnitList() error {
	for range Only.Once {
		data := sg.NewSunGrowData()
		data.SetArgs()
		data.SetEndpoints(queryUnitList.EndPointName)

		sg.Error = data.GetData()
		if sg.Error != nil {
			break
		}

		sg.Error = data.OutputDataTables()
		if sg.Error != nil {
			break
		}
	}

	return sg.Error
}
