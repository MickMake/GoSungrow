package mmHa

import (
	"GoSungrow/iSolarCloud/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdLog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"strings"
	"time"
)


type Mqtt struct {
	ClientId      string `json:"client_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Timeout       time.Duration `json:"timeout"`
	EntityPrefix  string `json:"entity_prefix"`

	url           *url.URL
	client        mqtt.Client
	pubClient     mqtt.Client
	clientOptions *mqtt.ClientOptions

	DeviceName     string
	LastRefresh    time.Time             `json:"-"`
	SungrowDevices getDeviceList.Devices `json:"-"`
	SungrowPsIds   map[valueTypes.PsId]bool
	MqttDevices    map[string]Device
	Prefix         string
	UserOptions    Options

	token    mqtt.Token
	firstRun bool
	err      error
	debug    bool
}


func New(req Mqtt) *Mqtt {
	var ret Mqtt

	for range Only.Once {
		ret.err = ret.setUrl(req)
		if ret.err != nil {
			break
		}
		ret.firstRun = true
		ret.EntityPrefix = req.EntityPrefix
		ret.Prefix = "homeassistant"

		// ret.selectPrefix = fmt.Sprintf("homeassistant/%s/%s", LabelSelect, req.ClientId)
		// ret.servicePrefix = fmt.Sprintf("homeassistant/%s/%s", LabelSensor, req.ClientId)
		// ret.sensorPrefix = fmt.Sprintf("homeassistant/%s/%s", LabelSensor, req.ClientId)
		// ret.lightPrefix = fmt.Sprintf("homeassistant/%s/%s", LabelLight, req.ClientId)
		// ret.switchPrefix = fmt.Sprintf("homeassistant/%s/%s", LabelSwitch, req.ClientId)
		// ret.binarySensorPrefix = fmt.Sprintf("homeassistant/%s/%s", LabelBinarySensor, req.ClientId)

		ret.MqttDevices = make(map[string]Device)
		ret.SungrowPsIds = make(map[valueTypes.PsId]bool)
		ret.Timeout = time.Second * 5
		ret.UserOptions.New()
	}

	return &ret
}

func (m *Mqtt) IsDebug() bool {
	return m.debug
}

func (m *Mqtt) LogDebug(format string, args ...interface{})  {
	if !m.debug {
		return
	}
	cmdLog.LogPrintDate(format, args...)
}

func (m *Mqtt) IsFirstRun() bool {
	return m.firstRun
}

func (m *Mqtt) IsNotFirstRun() bool {
	return !m.firstRun
}

func (m *Mqtt) UnsetFirstRun() {
	m.firstRun = false
}

func (m *Mqtt) GetError() error {
	return m.err
}

func (m *Mqtt) IsError() bool {
	if m.err != nil {
		return true
	}
	return false
}

func (m *Mqtt) IsNewDay() bool {
	var yes bool
	for range Only.Once {
		last := m.LastRefresh.Format("20060102")
		now := time.Now().Format("20060102")

		if last != now {
			yes = true
			break
		}
	}
	return yes
}

func (m *Mqtt) setUrl(req Mqtt) error {

	for range Only.Once {
		// if req.Username == "" {
		// 	m.err = errors.New("username empty")
		// 	break
		// }
		m.Username = req.Username

		// if req.Password == "" {
		// 	m.err = errors.New("password empty")
		// 	break
		// }
		m.Password = req.Password

		if req.Host == "" {
			m.err = errors.New("HASSIO mqtt host not defined")
			break
		}
		m.Host = req.Host

		if req.Port == "" {
			req.Port = "1883"
		}
		m.Port = req.Port

		u := fmt.Sprintf("tcp://%s:%s@%s:%s",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			)
		m.url, m.err = url.Parse(u)
	}

	return m.err
}

func (m *Mqtt) SetAuth(username string, password string) error {

	for range Only.Once {
		if username == "" {
			m.err = errors.New("username empty")
			break
		}
		m.Username = username

		if password == "" {
			m.err = errors.New("password empty")
			break
		}
		m.Password = password
	}

	return m.err
}

func (m *Mqtt) Connect() error {
	for range Only.Once {
		m.err = m.createClientOptions()
		if m.err != nil {
			break
		}

		m.client = mqtt.NewClient(m.clientOptions)
		token := m.client.Connect()
		for !token.WaitTimeout(3 * time.Second) {
		}
		if m.err = token.Error(); m.err != nil {
			break
		}
		if m.ClientId == "" {
			m.ClientId = "GoSungrow"
		}

		device := Config {
			Entry:      JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId),	// m.servicePrefix
			Name:       m.ClientId,
			UniqueId:   m.ClientId, 	// + "_Service",
			StateTopic:   "~/state",
			DeviceConfig: DeviceConfig {
				Identifiers:  []string{"GoSungrow"},
				SwVersion:    "GoSungrow https://github.com/MickMake/GoSungrow",
				Name:         m.ClientId + " Service",
				Manufacturer: "MickMake",
				Model:        "SunGrow",
			},
		}

		m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "config"), 0, true, device.Json())
		if m.err != nil {
			break
		}
		m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "state"), 0, true, "ON")
		if m.err != nil {
			break
		}

		_, m.err = m.SetDeviceConfig(
			m.DeviceName, m.DeviceName,
			"options", "Options", "", m.DeviceName,
			m.DeviceName,
		)
		if m.err != nil {
			break
		}

		m.err = m.SetOption("mqtt_debug", "Mqtt Debug", m.funcMqttDebug)
		if m.err != nil {
			break
		}

		v := OptionDisabled
		if m.debug {
			v = OptionEnabled
		}
		m.err = m.SetOptionValue("mqtt_debug", v)
		if m.err != nil {
			break
		}
	}

	return m.err
}

func (m *Mqtt) funcMqttDebug(_ mqtt.Client, msg mqtt.Message) {
	for range Only.Once {
		request := strings.ToLower(string(msg.Payload()))
		cmdLog.LogPrintDate("Option[%s] set to '%s'\n", msg.Topic(), request)
		if request == strings.ToLower(OptionEnabled) {
			m.err = m.SetOptionValue("mqtt_debug", OptionEnabled)
			m.debug = true
			break
		}
		m.err = m.SetOptionValue("mqtt_debug", OptionDisabled)
		m.debug = false
	}
}

func (m *Mqtt) Disconnect() error {
	for range Only.Once {
		m.client.Disconnect(250)
		time.Sleep(1 * time.Second)
	}
	return m.err
}

func (m *Mqtt) createClientOptions() error {
	for range Only.Once {
		m.clientOptions = mqtt.NewClientOptions()
		m.clientOptions.AddBroker(fmt.Sprintf("tcp://%s", m.url.Host))
		m.clientOptions.SetUsername(m.url.User.Username())
		password, _ := m.url.User.Password()
		m.clientOptions.SetPassword(password)
		m.clientOptions.SetClientID(m.ClientId)

		m.clientOptions.WillTopic = JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "state")
		m.clientOptions.WillPayload = []byte("OFF")
		m.clientOptions.WillQos = 0
		m.clientOptions.WillRetained = true
		m.clientOptions.WillEnabled = true
	}
	return m.err
}

// type SubscribeFunc func(client mqtt.Client, msg mqtt.Message)
func (m *Mqtt) subscribeDefault(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("*%t> [%s] %s\n", client.IsConnected(), msg.Topic(), string(msg.Payload()))
}

func (m *Mqtt) Subscribe(topic string, fn mqtt.MessageHandler) error {
	for range Only.Once {
		if fn == nil {
			fn = m.subscribeDefault
		}
		t := m.client.Subscribe(topic, 0, fn)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()

		}
	}
	return m.err
}

func (m *Mqtt) Publish(topic string, qos byte, retained bool, payload string) error {
	for range Only.Once {
		m.LogDebug("MQTT[%s] -> %v\n", topic, payload)
		t := m.client.Publish(topic, qos, retained, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}

func (m *Mqtt) PublishValue(Type string, subtopic string, value string) error {
	for range Only.Once {
		topic := ""
		switch Type {
			case LabelSensor:
				topic = JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, subtopic, "state")
				// state := MqttState {
				// 	LastReset: "", // m.GetLastReset(point.PointId),
				// 	Value:     value,
				// }
				// value = state.Json()

			case "binary_sensor":
				topic = JoinStringsForTopic(m.Prefix, LabelBinarySensor, m.ClientId, subtopic, "state")
				// state := MqttState {
				// 	LastReset: "", // m.GetLastReset(point.PointId),
				// 	Value:     value,
				// }
				// value = state.Json()

			case "lights":
				topic = JoinStringsForTopic(m.Prefix, LabelLight, m.ClientId, subtopic, "state")
				// state := MqttState {
				// 	LastReset: "", // m.GetLastReset(point.PointId),
				// 	Value:     value,
				// }
				// value = state.Json()

			case LabelSwitch:
				topic = JoinStringsForTopic(m.Prefix, LabelSwitch, m.ClientId, subtopic, "state")
				// state := MqttState {
				// 	LastReset: "", // m.GetLastReset(point.PointId),
				// 	Value:     value,
				// }
				// value = state.Json()

			default:
				topic = JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, subtopic, "state")
		}

		m.LogDebug("PublishValue - topic: '%s'\tpayload: '%s'\n", topic, value)
		t := m.client.Publish(topic, 0, true, value)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) SetDeviceConfig(swname string, parentId string, id string, name string, model string, vendor string, area string) (Device, error) {
	var ret Device

	for range Only.Once {
		// id = JoinStringsForId(m.EntityPrefix, id)

		c := [][]string{
			{swname, JoinStringsForId(m.EntityPrefix, parentId)},
			{JoinStringsForId(m.EntityPrefix, parentId), JoinStringsForId(m.EntityPrefix, id)},
		}
		if swname == parentId {
			c = [][]string{
				{parentId, JoinStringsForId(m.EntityPrefix, id)},
			}
		}

		ret = Device {
			Connections:  c,
			Identifiers:  []string{JoinStringsForId(m.EntityPrefix, id)},
			Manufacturer: vendor,
			Model:        model,
			Name:         name,
			SwVersion:    swname + " https://github.com/MickMake/" + swname,
			ViaDevice:    swname,
			SuggestedArea: area,
		}
		m.MqttDevices[id] = ret
	}

	return ret, m.err
}

// func (m *Mqtt) GetLastReset(pointType string) string {
// 	var ret string
//
// 	for range Only.Once {
// 		pt := api.GetDevicePoint(pointType)
// 		if !pt.Valid {
// 			break
// 		}
// 		if pt.UpdateFreq == "" {
// 			break
// 		}
// 		ret = pt.WhenReset()
// 	}
//
// 	return ret
// }


type MqttState struct {
	LastReset string `json:"last_reset,omitempty"`
	Value string `json:"value"`
}

func (mq *MqttState) Json() string {
	var ret string
	for range Only.Once {
		j, err := json.Marshal(*mq)
		if err != nil {
			ret = fmt.Sprintf("{ \"error\": \"%s\"", err)
			break
		}
		ret = string(j)
	}
	return ret
}

type SensorState string

type EntityConfig struct {
	// Type          string
	Name          string
	SubName       string

	ParentId      string
	ParentName    string

	UniqueId      string
	FullId        string
	Units         string
	ValueName     string
	DeviceClass   string
	StateClass    string
	Icon          string

	Value         *valueTypes.UnitValue
	Point         *api.Point
	ValueTemplate string

	UpdateFreq    string
	LastReset     string
	LastResetValueTemplate string

	IgnoreUpdate  bool

	haType        string
	Options       []string
}

func (config *EntityConfig) FixConfig() {

	for range Only.Once {
		// mdi:power-socket-au
		// mdi:solar-power
		// mdi:home-lightning-bolt-outline
		// mdi:transmission-tower
		// mdi:transmission-tower-export
		// mdi:transmission-tower-import
		// mdi:transmission-tower-off
		// mdi:home-battery-outline
		// mdi:lightning-bolt
		// mdi:check-circle-outline | mdi:arrow-right-bold
		// mdi:transmission-tower

		// Set ValueTemplate
		switch {
			// 	fallthrough
			//
			// case config.Value.TypeValue == "Energy":
			// 	fallthrough
			//
			// case config.Units == "MW":
			// 	fallthrough
			// case config.Units == "kW":
			// 	fallthrough
			// case config.Units == "W":
			// 	fallthrough
			//
			// case config.Units == "kWp":
			// 	fallthrough
			// case config.Units == "MWh":
			// 	fallthrough
			// case config.Units == "kWh":
			// 	fallthrough
			// case config.Units == "Wh":
			// 	fallthrough
			//
			// case config.Units == "kvar":
			// 	fallthrough
			// case config.Units == "Hz":
			// 	fallthrough
			// case config.Units == "V":
			// 	fallthrough
			// case config.Units == "A":
			// 	fallthrough
			// case config.Units == "°F":
			// 	fallthrough
			// case config.Units == "F":
			// 	fallthrough
			// case config.Units == "℉":
			// 	fallthrough
			// case config.Units == "°C":
			// 	fallthrough
			// case config.Units == "C":
			// 	fallthrough
			// case config.Units == "℃":
			// 	fallthrough
			// case config.Units == "%":
			case config.Value.IsFloat():
				if !config.Value.Valid {
					config.IgnoreUpdate = true
				}
				cnv := "| float"
				if config.Value.String() == "" {
					cnv = ""
				}
				if config.ValueName == "" {
					config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.value %s }}", cnv))
				} else {
					config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s %s }}", config.ValueName, cnv))
				}

			case config.Value.IsBool():
				fallthrough
			case config.Value.Unit() == LabelBinarySensor:
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")

			case config.Value.Unit() == "DateTime":
				fallthrough
			case config.Value.TypeValue == "DateTime":
				value, _, err := valueTypes.ParseDateTime(config.Value.String())
				if err == nil {
					config.Value.SetString(value.Local().Format(valueTypes.DateTimeFullLayout))
					config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | as_datetime }}")
				} else {
					config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")
				}

			case config.Value.IsInt():
				fallthrough
			default:
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")
		}

		// Set DeviceClass & Icon
		switch {
			case config.Units == "Bool":
				fallthrough
			case config.Units == LabelBinarySensor:
				config.DeviceClass = SetDefault(config.DeviceClass, "power")
				config.Icon = SetDefault(config.Icon, "mdi:check-circle-outline")
				// if !config.Value.Valid {
				// 	config.Value = "false"
				// }

			case config.Value.TypeValue == "Power":
				fallthrough
			case config.Units == "MW":
				fallthrough
			case config.Units == "kW":
				fallthrough
			case config.Units == "W":
				config.DeviceClass = SetDefault(config.DeviceClass, "power")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

			case config.Value.TypeValue == "Energy":
				fallthrough
			case config.Units == "MWh":
				fallthrough
			case config.Units == "kWh":
				fallthrough
			case config.Units == "Wh":
				config.DeviceClass = SetDefault(config.DeviceClass, "energy")
				config.Icon = SetDefault(config.Icon, "mdi:transmission-tower")

			case config.Units == "var":
				fallthrough
			case config.Units == "kvar":
				config.DeviceClass = SetDefault(config.DeviceClass, "reactive_power")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

			case config.Units == "VA":
				config.DeviceClass = SetDefault(config.DeviceClass, "apparent_power")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

			case config.Units == "Hz":
				config.DeviceClass = SetDefault(config.DeviceClass, "frequency")
				config.Icon = SetDefault(config.Icon, "mdi:sine-wave")

			case config.Units == "V":
				config.DeviceClass = SetDefault(config.DeviceClass, "voltage")
				config.Icon = SetDefault(config.Icon, "mdi:current-dc")

			case config.Units == "A":
				config.DeviceClass = SetDefault(config.DeviceClass, "current")
				config.Icon = SetDefault(config.Icon, "mdi:current-ac")

			case config.Units == "°F":
				fallthrough
			case config.Units == "F":
				fallthrough
			case config.Units == "℉":
				config.DeviceClass = SetDefault(config.DeviceClass, "temperature")
				config.Units = "℉"
				config.Icon = SetDefault(config.Icon, "mdi:thermometer")

			case config.Units == "°C":
				fallthrough
			case config.Units == "C":
				fallthrough
			case config.Units == "℃":
				config.DeviceClass = SetDefault(config.DeviceClass, "temperature")
				config.Units = "°C"
				config.Icon = SetDefault(config.Icon, "mdi:thermometer")

			case config.Icon == "mdi:home-battery-outline":
				fallthrough
			case config.Icon == "mdi:battery":
				config.DeviceClass = SetDefault(config.DeviceClass, "battery")
				// config.Icon = SetDefault(config.Icon, "mdi:percent") // mdi:home-battery-outline

			case config.Value.TypeValue == "Percent":
				fallthrough
			case config.Units == "%":
				// @TODO - Not supported in older versions of HA.
				// config.DeviceClass = SetDefault(config.DeviceClass, "battery")
				config.Icon = SetDefault(config.Icon, "mdi:percent") // mdi:home-battery-outline

			case config.Value.TypeValue == "DateTime":
				config.DeviceClass = SetDefault(config.DeviceClass, "timestamp") // date
				config.Icon = SetDefault(config.Icon, "mdi:clock-outline")

			case config.Units == "h":
				// config.DeviceClass = SetDefault(config.DeviceClass, "timestamp") // date
				config.Icon = SetDefault(config.Icon, "mdi:clock-outline")

			case config.Units == "kg":
				config.DeviceClass = SetDefault(config.DeviceClass, "weight")
				config.Icon = SetDefault(config.Icon, "mdi:weight")

			case config.Units == "km":
				config.DeviceClass = SetDefault(config.DeviceClass, "distance")
				config.Icon = SetDefault(config.Icon, "mdi:map-marker-distance")

			case config.Units == "Wh/㎡":
				fallthrough
			case config.Units == "W/㎡":
				// @TODO - Not supported in older versions of HA.
				config.DeviceClass = SetDefault(config.DeviceClass, "irradiance")
				config.Icon = SetDefault(config.Icon, "mdi:weather-sunny")

			case config.Value.TypeValue == "Currency":
				fallthrough
			case config.Units == "AUD":
				fallthrough
			case config.Units == "$":
				config.DeviceClass = SetDefault(config.DeviceClass, "monetary")
				config.Icon = SetDefault(config.Icon, "mdi:currency-usd")

				// p13013 - power_factor

			case config.Units == "GPS":
				// config.DeviceClass = SetDefault(config.DeviceClass, "")
				config.Icon = SetDefault(config.Icon, "mdi:crosshairs-gps")

			default:
				config.DeviceClass = SetDefault(config.DeviceClass, "")
				config.Icon = SetDefault(config.Icon, "")
		}

		switch {
			case config.Point.IsBoot():
				config.StateClass = "measurement"
				config.LastReset = ""
				config.LastResetValueTemplate = ""

			case config.Point.IsDaily():
				fallthrough
			case config.Point.IsMonthly():
				fallthrough
			case config.Point.IsYearly():
				fallthrough
			case config.Point.IsTotal():
				config.StateClass = "total"
				config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | as_datetime }}")
				// config.LastReset = config.Point.WhenReset(config.Date)

			case config.Point.Is5Minute():
				fallthrough
			case config.Point.Is15Minute():
				fallthrough
			case config.Point.Is30Minute():
				fallthrough
			case config.Point.IsInstant():
				fallthrough
			default:
				config.StateClass = "measurement"
				config.LastReset = ""
				config.LastResetValueTemplate = ""
		}

		// if config.LastReset == "" {
		// 	break
		// }
		//
		// pt := api.GetDevicePoint(config.FullId)
		// if !pt.Valid {
		// 	break
		// }
		//
		// config.LastReset = pt.WhenReset()
		// config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | as_datetime }}")
		// config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | int | timestamp_local | as_datetime }}")
		// config.StateClass = "total"
		// config.StateClass = "measurement"
	}
}

func SetDefault(value string, def string) string {
	if value == "" {
		value = def
	}
	return value
}
