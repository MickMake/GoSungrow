package getPsListStaticData

import (
	"time"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"github.com/MickMake/GoUnify/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
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
		DesignCapacity         api.Float   `json:"design_capacity"`
		GprsLatitude           api.Float   `json:"gprs_latitude"`
		GprsLongitude          api.Float   `json:"gprs_longitude"`
		InstallDate            api.DateTime      `json:"install_date"`
		InstallerPsFaultStatus api.Integer `json:"installer_ps_fault_status"`
		Latitude               api.Float   `json:"latitude"`
		Location               api.String      `json:"location"`
		Longitude              api.Float   `json:"longitude"`
		MapLatitude            api.Float   `json:"map_latitude"`
		MapLongitude           api.Float   `json:"map_longitude"`
		OwnerPsFaultStatus     api.Integer `json:"owner_ps_fault_status"`
		PsFaultStatus          api.Integer `json:"ps_fault_status"`
		PsID                   api.Integer `json:"ps_id"`
		PsName                 api.String      `json:"ps_name"`
		PsShortName            api.String      `json:"ps_short_name"`
		PsStatus               api.Integer `json:"ps_status"`
		PsType                 api.Integer `json:"ps_type"`
		ValidFlag              api.Integer `json:"valid_flag"`
		WaitAssignOrderCount   api.Integer `json:"wait_assign_order_count"`
	} `json:"pageList"`
	RowCount api.Integer `json:"rowCount"`
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
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", time.Time{})
	}

	return entries
}
