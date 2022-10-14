package getPowerStationBasicInfo

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/powerStationService/getPowerStationBasicInfo"
const Disabled = false

type RequestData struct {
	PsId valueTypes.Integer `json:"ps_id" required:"true"`
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
	ConnectType       valueTypes.Integer `json:"connect_type"`
	County            interface{} `json:"county"`
	CountyCode        interface{} `json:"county_code"`
	DesignCapacity    valueTypes.Integer `json:"design_capacity"`
	Email             string      `json:"email"`
	EnergyScheme      interface{} `json:"energy_scheme"`
	ExpectInstallDate valueTypes.DateTime      `json:"expect_install_date"`
	FaultSendType     interface{} `json:"fault_send_type"`
	GcjLatitude       valueTypes.Float   `json:"gcj_latitude"`
	GcjLongitude      valueTypes.Float   `json:"gcj_longitude"`
	Latitude          valueTypes.Float   `json:"latitude"`
	Longitude         valueTypes.Float   `json:"longitude"`
	Name              interface{} `json:"name"`
	NameCode          interface{} `json:"name_code"`
	Nation            interface{} `json:"nation"`
	NationCode        interface{} `json:"nation_code"`
	ParamIncome       valueTypes.Integer `json:"param_income"`
	Prov              interface{} `json:"prov"`
	ProvCode          interface{} `json:"prov_code"`
	PsBuildDate       valueTypes.DateTime      `json:"ps_build_date"`
	PsCountryID       valueTypes.Integer `json:"ps_country_id"`
	PsDesc            interface{} `json:"ps_desc"`
	PsHolder          string      `json:"ps_holder"`
	PsID              valueTypes.Integer `json:"ps_id"`
	PsLocation        string      `json:"ps_location"`
	PsName            string      `json:"ps_name"`
	PsType            valueTypes.Integer `json:"ps_type"`
	RecoreCreateTime  string      `json:"recore_create_time"`
	ReportType        interface{} `json:"report_type"`
	ShippingAddress   string      `json:"shipping_address"`
	ShippingZipCode   string      `json:"shipping_zip_code"`
	TimeZoneID        valueTypes.Integer `json:"time_zone_id"`
	ValidFlag         valueTypes.Integer `json:"valid_flag"`
	Village           interface{} `json:"village"`
	VillageCode       interface{} `json:"village_code"`
	WgsLatitude       valueTypes.Float   `json:"wgs_latitude"`
	WgsLongitude      valueTypes.Float   `json:"wgs_longitude"`
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

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToPoints(e.Response.ResultData, apiReflect.GetName("", *e), "system", valueTypes.NewDateTime(""))
	}

	return entries
}
