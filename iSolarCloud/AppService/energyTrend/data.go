package energyTrend

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/energyTrend"
const Disabled = false

type RequestData struct {
	// DeviceType valueTypes.String `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Echartunit valueTypes.String   `json:"echartunit" PointId:"echart_unit"`
	EndTime    valueTypes.DateTime `json:"endTime" PointId:"end_time"`
	EnergyMap  struct {
		ValStr valueTypes.String `json:"valStr" PointId:"val_str"`
	} `json:"energyMap" PointId:"energy_map"`
	Energyunit valueTypes.String `json:"energyunit" PointId:"energy_unit"`
	PowerMap   struct {
		Dates  []valueTypes.DateTime `json:"dates"`
		Units  valueTypes.String     `json:"units"`
		ValStr valueTypes.String     `json:"valStr" PointId:"val_str"`
	} `json:"powerMap" PointId:"power_map"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// pkg := reflection.GetName("", *e)
		// dt := valueTypes.NewDateTime(valueTypes.Now)
		// name := pkg + "." + e.Request.PsId.String()
		entries.StructToDataMap(*e, "", GoStruct.EndPointPath{})
	}

	return entries
}
