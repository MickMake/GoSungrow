package getPsListStaticData

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/powerStationService/getPsListStaticData"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		AreaID                 interface{} `json:"area_id"`
		DesignCapacity         valueTypes.Float   `json:"design_capacity"`
		GprsLatitude           valueTypes.Float   `json:"gprs_latitude"`
		GprsLongitude          valueTypes.Float   `json:"gprs_longitude"`
		InstallDate            valueTypes.DateTime      `json:"install_date"`
		InstallerPsFaultStatus valueTypes.Integer `json:"installer_ps_fault_status"`
		Latitude               valueTypes.Float   `json:"latitude"`
		Location               valueTypes.String      `json:"location"`
		Longitude              valueTypes.Float   `json:"longitude"`
		MapLatitude            valueTypes.Float   `json:"map_latitude"`
		MapLongitude           valueTypes.Float   `json:"map_longitude"`
		OwnerPsFaultStatus     valueTypes.Integer `json:"owner_ps_fault_status"`
		PsFaultStatus          valueTypes.Integer `json:"ps_fault_status"`
		PsID                   valueTypes.Integer `json:"ps_id"`
		PsName                 valueTypes.String      `json:"ps_name"`
		PsShortName            valueTypes.String      `json:"ps_short_name"`
		PsStatus               valueTypes.Integer `json:"ps_status"`
		PsType                 valueTypes.Integer `json:"ps_type"`
		ValidFlag              valueTypes.Integer `json:"valid_flag"`
		WaitAssignOrderCount   valueTypes.Integer `json:"wait_assign_order_count"`
	} `json:"pageList"`
	RowCount valueTypes.Integer `json:"rowCount"`
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

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
