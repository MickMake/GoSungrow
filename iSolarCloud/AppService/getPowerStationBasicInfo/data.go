package getPowerStationBasicInfo

import (
	"GoSungrow/iSolarCloud/api/apiReflect"
	"fmt"
)

const Url = "/v1/powerStationService/getPowerStationBasicInfo"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData   struct {
	AreaID            interface{} `json:"area_id"`
	ConnectType       int64       `json:"connect_type"`
	County            interface{} `json:"county"`
	CountyCode        interface{} `json:"county_code"`
	DesignCapacity    int64       `json:"design_capacity"`
	Email             string      `json:"email"`
	EnergyScheme      interface{} `json:"energy_scheme"`
	ExpectInstallDate string      `json:"expect_install_date"`
	FaultSendType     interface{} `json:"fault_send_type"`
	GcjLatitude       string      `json:"gcj_latitude"`
	GcjLongitude      string      `json:"gcj_longitude"`
	Latitude          string      `json:"latitude"`
	Longitude         string      `json:"longitude"`
	Name              interface{} `json:"name"`
	NameCode          interface{} `json:"name_code"`
	Nation            interface{} `json:"nation"`
	NationCode        interface{} `json:"nation_code"`
	ParamIncome       int64       `json:"param_income"`
	Prov              interface{} `json:"prov"`
	ProvCode          interface{} `json:"prov_code"`
	PsBuildDate       string      `json:"ps_build_date"`
	PsCountryID       int64       `json:"ps_country_id"`
	PsDesc            interface{} `json:"ps_desc"`
	PsHolder          string      `json:"ps_holder"`
	PsID              int64       `json:"ps_id"`
	PsLocation        string      `json:"ps_location"`
	PsName            string      `json:"ps_name"`
	PsType            int64       `json:"ps_type"`
	RecoreCreateTime  string      `json:"recore_create_time"`
	ReportType        interface{} `json:"report_type"`
	ShippingAddress   string      `json:"shipping_address"`
	ShippingZipCode   string      `json:"shipping_zip_code"`
	TimeZoneID        int64       `json:"time_zone_id"`
	ValidFlag         int64       `json:"valid_flag"`
	Village           interface{} `json:"village"`
	VillageCode       interface{} `json:"village_code"`
	WgsLatitude       float64     `json:"wgs_latitude"`
	WgsLongitude      float64     `json:"wgs_longitude"`
	ZipCode           string      `json:"zip_code"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
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
