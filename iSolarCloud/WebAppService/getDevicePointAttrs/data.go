package getDevicePointAttrs

import (
	"github.com/MickMake/GoSungrow/iSolarCloud/api"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"

	"fmt"
)

const Url = "/v1/devService/getDevicePointAttrs"
const Disabled = false
const EndPointName = "WebAppService.getDevicePointAttrs"

type RequestData struct {
	Uuid        valueTypes.Integer `json:"uuid,omitempty"`
	DeviceType2 valueTypes.Integer `json:"deviceType" required:"true"`
	PsId2       valueTypes.PsId    `json:"psId" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData []Point

func (e *ResultData) IsValid() error {
	var err error
	return err
}

// powerDevicePointList
// ┏━━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━┓
// ┃ Device Type    ┃ Id     ┃ Period    ┃ Point Id    ┃ Name                                      ┃ Type Name                ┃
// ┣━━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━┫
// ┃ 11             ┃ 29     ┃ 1         ┃ p83001      ┃ AC Power Normalization (Inverter)         ┃ Plant                    ┃
//
// ┏━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━┳━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━┳━━━━━━━━━━┳━━━━━━━━━━━━━━━━┳━━━━━━━━┓
// ┃ Ps Id   ┃ Device Type ┃ Channel Id ┃ Station Name ┃ Station Short Name ┃ Is Parent  ┃ Device Model Id ┃ Device Name                 ┃ A Type ┃ Code Id ┃ C Type ┃ Id     ┃ Node Key ┃ Name                            ┃ Target Unit ┃ Unit   ┃ Unit Type ┃ Level ┃ Order Id ┃ Point Group Id ┃ Relate ┃
// ┣━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━╇━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━╇━━━━━━━╇━━━━━━━━━━╇━━━━━━━━━━━━━━━━╇━━━━━━━━┫
// ┃ 1171348 ┃ 11          ┃ 0          ┃ MickMake42   ┃ A2281658237        ┃ false      ┃ 366316          ┃ Energy Storage System 02_01 ┃ 0      ┃ 0       ┃ 1      ┃ p83001 ┃ p83001   ┃ Inverter AC Power Normalization ┃ kW/kWp      ┃ kW/kWp ┃ 43        ┃ 4     ┃ 10021    ┃ 9999           ┃ 0      ┃

type Point struct {
	GoStruct.GoStructParent `json:"-" DataTable:"true" DataTableSortOn:"Id"`
	// GoStruct.GoStruct `json:"-" PointIdFrom:"PsId" PointIdReplace:"false"`

	PsId       valueTypes.PsId    `json:"psid" PointId:"ps_id"`
	DeviceType valueTypes.Integer `json:"pid" PointId:"device_type"`
	ChannelId  valueTypes.Integer `json:"chnnlid" PointId:"channel_id"`

	Unit         valueTypes.String  `json:"unit"`
	OrderId      valueTypes.Integer `json:"orderid" PointId:"order_id"`
	PointGroupId valueTypes.Integer `json:"point_group_id"`
	Relate       valueTypes.Integer `json:"relate"`
	CodeId       valueTypes.Integer `json:"code_id"`

	StationName      valueTypes.String  `json:"stationname" PointId:"station_name"`
	StationShortName valueTypes.String  `json:"stationshortname" PointId:"station_short_name"`
	IsParent         valueTypes.Bool    `json:"isparent" PointId:"is_parent"`
	DeviceModelId    valueTypes.Integer `json:"device_model_id"`
	DeviceName       valueTypes.String  `json:"devicename" PointId:"device_name"`
	AType            valueTypes.Integer `json:"atype" PointId:"a_type"`
	CType            valueTypes.Integer `json:"ctype" PointId:"c_type"`

	Id         valueTypes.PointId `json:"id"`
	NodeKey    valueTypes.PointId `json:"nodekey" PointId:"node_key"`
	Name       valueTypes.String  `json:"name"`
	TargetUnit valueTypes.String  `json:"target_unit"`
	UnitType   valueTypes.Integer `json:"unit_type"`
	Level      valueTypes.Integer `json:"level"`

	// Added.
	PointGroupName string `json:"point_group_name"`
	UnitTypeName   string `json:"unit_type_name"`
}

func (p *Point) UnmarshalJSON(data []byte) error {
	var err error

	for range Only.Once {
		if len(data) == 0 {
			break
		}
		type decode Point
		var pd decode

		// Store PointsData.Order
		err = json.Unmarshal(data, &pd)
		if err != nil {
			fmt.Printf("getDevicePointAttrs[pd] - err:%s data: %s\n", err, data)
			break
		}

		pd.PointGroupName = pointGroupNames.Get(pd.PointGroupId.String())
		// pd.UnitTypeName = unitTypes.Get(pd.UnitType.String())	// Not working.
		*p = Point(pd)
	}

	return err
}

type Points []Point

type PointsMap map[string]*Point

func (m *PointsMap) Exists(point string) bool {
	if _, ok := (*m)[point]; ok {
		return true
	}
	return false
}

func (m *PointsMap) Get(point string) *Point {
	if m.Exists(point) {
		return (*m)[point]
	}
	return nil
}


type UnitTypes map[string]string

var unitTypes = UnitTypes {
	"":       "Unknown",              //
	"NULL":   "Unknown",              //

	"%":      "Percent",              // 10
	"%RH":    "Humidity",             // 12

	"A":      "Current",              //
	"Ah":     "Current Capacity",     //
	"mA":     "Current",              //

	"Bar":    "Pressure",             //
	"hPa":    "Pressure",             // 19

	"Day":    "Date/Time",            // 15
	"H":      "Date/Time",            // 15
	"Hour":   "Date/Time",            // 15
	"h":      "Date/Time",            // 15
	"Year":   "Date/Time",            // 15
	"Min":    "Date/Time",            // 15
	"Mon":    "Date/Time",            // 15
	"min":    "Date/Time",            // 15
	"s":      "Date/Time",            // 15

	"Hz":     "Frequency",            //

	"KV":     "Voltage",              //
	"kV ":    "Voltage",              //
	"kV":     "Voltage",              //
	"V":      "Voltage",              //
	"mV":     "Voltage",              //

	"MAh":    "Energy",               // 13
	"MWh":    "Energy",               //
	"MWp":    "Energy",               //
	"Wh":     "Energy",               //
	"kWh":    "Energy",               //

	"MW":     "Power",                //
	"kW":     "Power",                // 0 / 21
	"W":      "Power",                //

	"Mvar":   "Reactive Power",       //
	"kVar":   "Reactive Power",       //
	"kvar":   "Reactive Power",       //
	"Var":    "Reactive Power",       //

	"VA":     "Apparent Power",       // 21
	"kVA":    "Apparent Power",       // 21

	"kW/kWp": "Power Normalization",  //

	"kVarh":  "Reactive Electricity", //
	"kvarh":  "Reactive Electricity", //

	"PCS":    "PCS",                  //

	"V/mA":   "Output",               // 0

	"W/㎡":    "Irradiation",          //
	"Wh/㎡":   "Irradiation",          // 1

	"dec":    "dec",                  //

	"g":      "Weight",               // 17
	"kg":     "Weight",               // 17

	"kΩ":     "Resistance",           //

	"m/s":    "Speed",                //
	"r/min":  "Speed",                //

	"mm":     "Length",               //

	"°":      "Angle",                //

	"℃":      "Temperature",          // 2

	"时":     "Date/Time",                   //  15
	"分":     "Date/Time",                 // Operation Minutes
	"台":     "Count",                   // Number of Online PCSs
	"排":     "Count",                   //
	"次":     "Count",                   // Discharging Count
}

// 21 = kW
// 22 = kVarh
// 22 = kvarh
// 23 = kΩ
// 25 = m/s
// 28 = mm
// 3 = kVA
// 3 = kW
// 3 = kWh
// 31 = mV
// 32 = Mvar
// 33 =
// 33 = PCS
// 33 = 分
// 33 = 台
// 33 = 排
// 33 = 次
// 37 = KV
// 37 = kV
// 39 = VA
// 4 =
// 4 = V
// 40 = W/㎡
// 43 = kW/kWp
// 44 = °
// 5 = A
// 6 = Hz
// 7 = kWh
// 8 = kvar
// 999 =
// 999 = %
// 999 = %RH
// 999 = A
// 999 = Ah
// 999 = Bar
// 999 = Hz
// 999 = MAh
// 999 = MW
// 999 = MWh
// 999 = MWp
// 999 = Mvar
// 999 = NULL
// 999 = V
// 999 = VA
// 999 = Var
// 999 = W
// 999 = Wh
// 999 = dec
// 999 = h
// 999 = kV
// 999 = kV
// 999 = kVA
// 999 = kVar
// 999 = kW
// 999 = kWh
// 999 = kvar
// 999 = kΩ
// 999 = mA
// 999 = mV
// 999 = min
// 999 = r/min
// 999 = ℃

func (ut *UnitTypes) Get(name string) string {
	if n, ok := (*ut)[name]; ok {
		return n
	}
	return name
}


type PointGroupNames map[string]string

var pointGroupNames = PointGroupNames{
	"1":    "1",
	"2":    "2",
	"7":    "7",
	"9":    "Overview",
	"10":   "Battery Information",
	"11":   "11",
	"12":   "Grid Information",
	"13":   "Load Information",
	"14":   "14",
	"16":   "MPPT Information",
	"17":   "17",
	"19":   "19",
	"20":   "20",
	"21":   "21",
	"9999": "Other Information",
}

func (pgn *PointGroupNames) Get(name string) string {
	if n, ok := (*pgn)[name]; ok {
		return n
	}
	return name
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()
	entries.StructToDataMap(*e, e.Request.PsId2.String(), GoStruct.NewEndPointPath(e.Request.PsId2.String()))
	return entries
}

func (e *EndPoint) Points() Points {
	return Points(e.Response.ResultData)
}
