package getDeviceList

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"sort"
	"time"
)

const Url = "/v1/devService/getDeviceList"
const Disabled = false

type RequestData struct {
	PsId api.Integer `json:"ps_id" required:"true"`
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
		AttrID                  api.Integer `json:"attr_id"`
		ChannelId               api.Integer `json:"chnnl_id" PointId:"channel_id"`
		CommandStatus           api.Integer `json:"command_status"`
		ConnectState            api.Integer `json:"connect_state"`
		DataFlag                api.Integer `json:"data_flag"`
		DataFlagDetail          api.Integer `json:"data_flag_detail"`
		DevFaultStatus          string      `json:"dev_fault_status"`
		DevStatus               string      `json:"dev_status"`
		DeviceArea              string      `json:"device_area"`
		DeviceCode              api.Integer `json:"device_code"`
		DeviceFactoryDate       api.DateTime `json:"device_factory_date"`
		DeviceID                api.Integer `json:"device_id"`
		DeviceModel             api.String  `json:"device_model"`
		DeviceModelCode         api.String  `json:"device_model_code"`
		DeviceModelID           api.Integer `json:"device_model_id"`
		DeviceName              api.String  `json:"device_name"`
		DeviceProSn             api.String  `json:"device_pro_sn"`
		DeviceState             string      `json:"device_state"`
		DeviceSubType           interface{} `json:"device_sub_type"`
		DeviceSubTypeName       interface{} `json:"device_sub_type_name"`
		DeviceType              api.Integer `json:"device_type"`
		FactoryName             api.String  `json:"factory_name"`
		InstallerDevFaultStatus string      `json:"installer_dev_fault_status"`
		InverterModelType       api.Integer `json:"inverter_model_type"`
		IsCountryCheck          api.Bool    `json:"is_country_check"`
		IsHasFunctionEnum       api.Bool    `json:"is_has_function_enum"`
		IsHasTheAbility         api.Bool    `json:"is_has_the_ability"`
		IsInit                  api.Bool    `json:"is_init"`
		IsReadSet               api.Bool    `json:"is_read_set"`
		IsReplacing             api.Bool    `json:"is_replacing"`
		IsReset                 api.Bool    `json:"is_reset"`
		IsSecond                api.Bool    `json:"is_second"`
		IsThirdParty            api.Bool    `json:"is_third_party"`
		ModuleUUID              api.Integer `json:"module_uuid"`
		OwnerDevFaultStatus     string      `json:"owner_dev_fault_status"`
		P24                     interface{} `json:"p24"`
		Posx                    interface{} `json:"posx"`
		Posy                    interface{} `json:"posy"`
		PsID                    api.Integer `json:"ps_id"`
		PsKey                   api.PsKey   `json:"ps_key"`
		RelState                api.Integer `json:"rel_state"`
		Sn                      api.String  `json:"sn" PointName:"Serial Number"`
		TypeName                api.String  `json:"type_name"`
		UUID                    api.Integer `json:"uuid"`
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

type Device struct {
	Vendor        api.String
	PsId          api.Integer
	PsKey         api.PsKey
	DeviceName    api.String
	DeviceProSn   api.String
	DeviceModel   api.String
	DeviceType    api.Integer
	DeviceCode    api.Integer
	ChannelId     api.Integer
	DeviceModelID api.Integer
	TypeName      api.String
	DeviceState   string
	DevStatus     string
	Uuid          api.Integer
}
type Devices []Device

func (e *EndPoint) GetDevices() Devices {
	var ret Devices
	for _, d := range e.Response.ResultData.PageList {
		ret = append(ret, Device{
			Vendor:        d.FactoryName,
			PsKey:         d.PsKey,
			PsId:          d.PsID,
			DeviceType:    d.DeviceType,
			DeviceCode:    d.DeviceCode,
			ChannelId:     d.ChannelId,
			TypeName:      d.TypeName,
			DeviceProSn:   d.DeviceProSn,
			DeviceModel:   d.DeviceModel,
			DeviceModelID: d.DeviceModelID,
			DeviceName:    d.DeviceName,
			DeviceState:   d.DeviceState,
			DevStatus:     d.DevStatus,
			Uuid:          d.ModuleUUID,
		})
	}
	return ret
}

func (e *EndPoint) GetDevicesTable() output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsID.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelID.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }

		data := e.GetDevices()
		table = GetDevicesTable(data)
	}
	return table
}

func GetDevicesTable(data Devices) output.Table {
	var table output.Table
	for range Only.Once {
		// table = output.NewTable()
		// table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))
		//
		// _ = table.SetHeader(
		// 	"Ps Key",
		// 	"Ps Id",
		// 	"Type",
		// 	"Code",
		// 	"Id",
		// 	"Type Name",
		// 	"Serial Number",
		// 	"Model",
		// 	"Model Id",
		// 	"Name",
		// 	"State",
		// 	"Status",
		// 	// "Factory Date",
		// )
		// for _, d := range e.Response.ResultData.PageList {
		// 	_ = table.AddRow(
		// 		d.PsKey.Value(),
		// 		d.PsID.Value(),
		// 		d.DeviceType.Value(),
		// 		d.DeviceCode.Value(),
		// 		d.ChannelId.Value(),
		// 		d.TypeName.Value(),
		// 		d.DeviceProSn.Value(),
		// 		d.DeviceModel.Value(),
		// 		d.DeviceModelID.Value(),
		// 		d.DeviceName.Value(),
		// 		d.DeviceState,
		// 		d.DevStatus,
		// 		// d.DeviceFactoryDate,
		// 	)
		// }

		table = output.NewTable()
		table.SetTitle("")
		// table.SetJson([]byte(e.GetJsonData(false)))
		// table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Vendor",
			"Ps Key",
			"Ps Id",
			"Type",
			"Code",
			"Id",
			"Type Name",
			"Serial Number",
			"Model",
			"Model Id",
			"Name",
			"State",
			"Status",
			"UUID",
		)
		for _, d := range data {
			_ = table.AddRow(
				d.Vendor.Value(),
				d.PsKey.Value(),
				d.PsId.Value(),
				d.DeviceType.Value(),
				d.DeviceCode.Value(),
				d.ChannelId.Value(),
				d.TypeName.Value(),
				d.DeviceProSn.Value(),
				d.DeviceModel.Value(),
				d.DeviceModelID.Value(),
				d.DeviceName.Value(),
				d.DeviceState,
				d.DevStatus,
				d.Uuid.Value(),
			)
		}
	}
	return table
}


func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		pkg := apiReflect.GetName("", *e)
		for _, d := range e.Response.ResultData.PageList {
			name := fmt.Sprintf("%s.%s", pkg, d.PsKey.Value())
			entries.StructToPoints(d, name, e.Request.PsId.String(), time.Time{})
		}
	}

	return entries
}

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		table.SetTitle("")
		table.SetJson([]byte(e.GetJsonData(false)))
		table.SetRaw([]byte(e.GetJsonData(true)))

		_ = table.SetHeader(
			"Date",
			"Point Id",
			// "Parents",
			"Group Name",
			"Description",
			"Value",
			"Unit",
		)

		data := e.GetData()
		var sorted []string
		for p := range data.DataPoints {
			sorted = append(sorted, string(p))
		}
		sort.Strings(sorted)

		for _, p := range sorted {
			entries := data.DataPoints[api.PointId(p)]
			for _, de := range entries {
				if de.Hide {
					continue
				}

				_ = table.AddRow(
					de.Date.Format(api.DtLayout),
					// api.NameDevicePointInt(de.Point.Parents, p.PointID.Value()),
					// de.Point.Id,
					p,
					// de.Point.Parents.String(),
					de.Point.GroupName,
					de.Point.Name,
					de.Value,
					de.Point.Unit,
				)
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(4),
		// 	ValueColumn:  output.SetInteger(5),
		// 	UnitsColumn:  output.SetInteger(6),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}
	return table
}
