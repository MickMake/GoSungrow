package getPsListStaticData

import (
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
		DesignCapacity         float64     `json:"design_capacity"`
		GprsLatitude           interface{} `json:"gprs_latitude"`
		GprsLongitude          interface{} `json:"gprs_longitude"`
		InstallDate            string      `json:"install_date"`
		InstallerPsFaultStatus int64       `json:"installer_ps_fault_status"`
		Latitude               float64     `json:"latitude"`
		Location               string      `json:"location"`
		Longitude              float64     `json:"longitude"`
		MapLatitude            string      `json:"map_latitude"`
		MapLongitude           string      `json:"map_longitude"`
		OwnerPsFaultStatus     int64       `json:"owner_ps_fault_status"`
		PsFaultStatus          int64       `json:"ps_fault_status"`
		PsID                   int64       `json:"ps_id"`
		PsName                 string      `json:"ps_name"`
		PsShortName            string      `json:"ps_short_name"`
		PsStatus               int64       `json:"ps_status"`
		PsType                 int64       `json:"ps_type"`
		ValidFlag              int64       `json:"valid_flag"`
		WaitAssignOrderCount   int64       `json:"wait_assign_order_count"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
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
