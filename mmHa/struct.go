package mmHa

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
)


type Mqtt struct {
	ClientId      string `json:"client_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Timeout  time.Duration `json:"timeout"`
	EntityPrefix string `json:"entity_prefix"`

	url           *url.URL
	client        mqtt.Client
	pubClient     mqtt.Client
	clientOptions *mqtt.ClientOptions
	LastRefresh   time.Time `json:"-"`
	PsId          api.Integer `json:"-"`

	Device Device

	servicePrefix string
	sensorPrefix string
	lightPrefix string
	switchPrefix string
	binarySensorPrefix string

	token    mqtt.Token
	firstRun bool
	err      error
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

		ret.servicePrefix = "homeassistant/sensor/" + req.ClientId
		ret.sensorPrefix = "homeassistant/sensor/" + req.ClientId
		ret.lightPrefix = "homeassistant/light/" + req.ClientId
		ret.switchPrefix = "homeassistant/switch/" + req.ClientId
		ret.binarySensorPrefix = "homeassistant/binary_sensor/" + req.ClientId
	}

	return &ret
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
			m.ClientId = "GoSunGrow"
		}

		device := Config {
			Entry:      m.servicePrefix,
			Name:       m.ClientId,
			UniqueId:   m.ClientId, 	// + "_Service",
			StateTopic:   "~/state",
			DeviceConfig: DeviceConfig {
				Identifiers:  []string{"GoSunGrow"},
				SwVersion:    "GoSunGrow https://github.com/MickMake/GoSungrow",
				Name:         m.ClientId + " Service",
				Manufacturer: "MickMake",
				Model:        "SunGrow",
			},
		}

		m.err = m.Publish(JoinStringsForTopic(m.servicePrefix, "config"), 0, true, device.Json())
		if m.err != nil {
			break
		}

		m.err = m.Publish(JoinStringsForTopic(m.servicePrefix, "state"), 0, true, "ON")
		if m.err != nil {
			break
		}

	}

	return m.err
}

func (m *Mqtt) Disconnect() error {
	for range Only.Once {
		m.client.Disconnect(250)
		time.Sleep(1 * time.Second)
	}
	return m.err
}

// const ServiceBaseName = "GoSunGrow"
// const ServiceBaseUniqueId = ServiceBaseName + "_Service"
// const ServiceBaseTopic = "homeassistant/sensor/" + ServiceBaseName
// const SensorBaseTopic = "homeassistant/sensor/" + ServiceBaseName

func (m *Mqtt) createClientOptions() error {
	for range Only.Once {
		m.clientOptions = mqtt.NewClientOptions()
		m.clientOptions.AddBroker(fmt.Sprintf("tcp://%s", m.url.Host))
		m.clientOptions.SetUsername(m.url.User.Username())
		password, _ := m.url.User.Password()
		m.clientOptions.SetPassword(password)
		m.clientOptions.SetClientID(m.ClientId)

		m.clientOptions.WillTopic = JoinStringsForTopic(m.servicePrefix, "state")
		m.clientOptions.WillPayload = []byte("OFF")
		m.clientOptions.WillQos = 0
		m.clientOptions.WillRetained = true
		m.clientOptions.WillEnabled = true
	}
	return m.err
}

func (m *Mqtt) Subscribe(topic string) error {
	for range Only.Once {
		t := m.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		})
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
			// m.err = errors.New("mqtt subscribe timeout")
		}
	}
	return m.err
}

func (m *Mqtt) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	for range Only.Once {
		t := m.client.Publish(topic, qos, retained, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
			// m.err = errors.New("mqtt publish timeout")
		}
	}
	return m.err
}


// func (m *Mqtt) PublishConfig(config EntityConfig) error {
// 	// func (m *Mqtt) PublishConfig(t string, id string, name string, subName string, units string, valueName string, class string) error {
// 	switch config.Type {
// 	// case "sensor":
// 	// 	m.err = m.PublishSensorConfig(id, name, subName, units, valueName, class)
// 	case "binary_sensor":
// 		m.err = m.PublishBinarySensorConfig(config)	// (id, name, subName, units, valueName, class)
// 	// case "lights":
// 	// 	m.err = m.PublishLightConfig(id, name, subName, units, valueName, class)
// 	// case "switch":
// 	// 	m.err = m.PublishSwitchConfig(id, name, subName, units, valueName, class)
// 	// default:
// 	// 	m.err = m.PublishSensorConfig(config)
// 	}
//
// 	return m.err
// }

func (m *Mqtt) PublishState(Type string, subtopic string, payload interface{}) error {
	for range Only.Once {
		// topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
		// topic = JoinStringsForTopic(m.sensorPrefix, topic, "state")
		// st := JoinStringsForTopic(m.sensorPrefix, JoinStringsForId(m.EntityPrefix, m.Device.FullName, strings.ReplaceAll(subName, "/", ".")), "state")
		topic := ""
		switch Type {
			case "sensor":
				topic = JoinStringsForTopic(m.sensorPrefix, subtopic, "state")
			case "binary_sensor":
				topic = JoinStringsForTopic(m.binarySensorPrefix, subtopic, "state")
			case "lights":
				topic = JoinStringsForTopic(m.lightPrefix, subtopic, "state")
			case "switch":
				topic = JoinStringsForTopic(m.switchPrefix, subtopic, "state")
			default:
				topic = JoinStringsForTopic(m.sensorPrefix, subtopic, "state")
		}

		t := m.client.Publish(topic, 0, true, payload)
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
		case "sensor":
			topic = JoinStringsForTopic(m.sensorPrefix, subtopic, "state")
			// state := MqttState {
			// 	LastReset: "", // m.GetLastReset(point.PointId),
			// 	Value:     value,
			// }
			// value = state.Json()

		case "binary_sensor":
			topic = JoinStringsForTopic(m.binarySensorPrefix, subtopic, "state")
			// state := MqttState {
			// 	LastReset: "", // m.GetLastReset(point.PointId),
			// 	Value:     value,
			// }
			// value = state.Json()

		case "lights":
			topic = JoinStringsForTopic(m.lightPrefix, subtopic, "state")
			// state := MqttState {
			// 	LastReset: "", // m.GetLastReset(point.PointId),
			// 	Value:     value,
			// }
			// value = state.Json()

		case "switch":
			topic = JoinStringsForTopic(m.switchPrefix, subtopic, "state")
			// state := MqttState {
			// 	LastReset: "", // m.GetLastReset(point.PointId),
			// 	Value:     value,
			// }
			// value = state.Json()

		default:
			topic = JoinStringsForTopic(m.sensorPrefix, subtopic, "state")
		}

		// t = JoinStringsForId(m.EntityPrefix, m.Device.Name, t)
		// st := JoinStringsForTopic(m.sensorPrefix, JoinStringsForId(m.EntityPrefix, m.Device.FullName, strings.ReplaceAll(subName, "/", ".")), "state")
		// payload := MqttState {
		// 	LastReset: "", // m.GetLastReset(point.PointId),
		// 	Value:     value,
		// }
		// m.client.Publish(JoinStringsForTopic(m.sensorPrefix, t, "state"), 0, true, payload.Json())
		t := m.client.Publish(topic, 0, true, value)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

// func (m *Mqtt) PublishValue(t string, topic string, value string) error {
// 	switch t {
// 		case "sensor":
// 			m.err = m.PublishSensorValue(topic, value)
// 		case "binary_sensor":
// 			m.err = m.PublishBinarySensorState(topic, value)
// 		case "lights":
// 			m.err = m.PublishLightState(topic, value)
// 		case "switch":
// 			m.err = m.PublishSwitchState(topic, value)
// 		default:
// 			m.err = m.PublishSensorState(topic, value)
// 	}
//
// 	return m.err
// }

func (m *Mqtt) SetDeviceConfig(swname string, id string, name string, model string, vendor string, area string) error {
	for range Only.Once {
		id = JoinStringsForId(m.EntityPrefix, id)

		m.Device = Device {
			Connections:  [][]string{
				{swname, id},
			},
			Identifiers:  []string{id},
			Manufacturer: vendor,
			Model:        model,
			Name:         name,
			SwVersion:    swname + " https://github.com/MickMake/" + swname,
			ViaDevice:    swname,
			SuggestedArea: area,
		}
	}
	return m.err
}


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


type Availability struct {
	PayloadAvailable    string `json:"payload_available,omitempty" required:"false"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty" required:"false"`
	Topic               string `json:"topic,omitempty" required:"true"`
	ValueTemplate       string `json:"value_template,omitempty" required:"false"`
}
type SensorState string


func (m *Mqtt) GetLastReset(pointType string) string {
	var ret string

	for range Only.Once {
		pt := api.GetDevicePoint(pointType)
		if !pt.Valid {
			break
		}
		if pt.Type == "" {
			break
		}
		ret = pt.WhenReset()
	}

	return ret
}


type EntityConfig struct {
	// Type          string
	Name          string
	SubName       string

	ParentId      string
	ParentName    string

	UniqueId string
	FullId   string
	Units    string
	ValueName     string
	DeviceClass   string
	StateClass    string
	Icon          string

	Value         string
	ValueTemplate string

	LastReset     string
	LastResetValueTemplate string

	haType        string
}

func (config *EntityConfig) IsSensor() bool {
	var ok bool

	for range Only.Once {
		if config.IsBinarySensor() {
			break
		}
		if config.IsSwitch() {
			break
		}
		if config.IsLight() {
			break
		}
		// if config.Value != "" {
		// 	ok = true
		// 	break
		// }
		// if config.Units == "" {
		// 	break
		// }

		ok = true
	}

	return ok
}

func (config *EntityConfig) IsBinarySensor() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelBinarySensor {
			ok = true
			break
		}
	}

	return ok
}

func (config *EntityConfig) IsSwitch() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelSwitch {
			ok = true
			break
		}
	}

	return ok
}

func (config *EntityConfig) IsLight() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelBinarySensor {
			ok = true
			break
		}
	}

	return ok
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

		switch config.Units {
			case LabelBinarySensor:
				config.DeviceClass = SetDefault(config.DeviceClass, "power")
				config.Icon = SetDefault(config.Icon, "mdi:check-circle-outline")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")

			case "MW":
				fallthrough
			case "kW":
				fallthrough
			case "W":
				config.DeviceClass = SetDefault(config.DeviceClass, "power")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")
				// config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s | float }}", config.ValueName))
				// - Used with merged values.

			case "MWh":
				fallthrough
			case "kWh":
				fallthrough
			case "Wh":
				config.DeviceClass = SetDefault(config.DeviceClass, "energy")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "kvar":
				config.DeviceClass = SetDefault(config.DeviceClass, "reactive_power")
				config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "Hz":
				config.DeviceClass = SetDefault(config.DeviceClass, "frequency")
				config.Icon = SetDefault(config.Icon, "mdi:sine-wave")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "V":
				config.DeviceClass = SetDefault(config.DeviceClass, "voltage")
				config.Icon = SetDefault(config.Icon, "mdi:current-dc")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "A":
				config.DeviceClass = SetDefault(config.DeviceClass, "current")
				config.Icon = SetDefault(config.Icon, "mdi:current-ac")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "°C":
				fallthrough
			case "C":
				fallthrough
			case "℃":
				config.DeviceClass = SetDefault(config.DeviceClass, "temperature")
				config.Units = "°C"
				config.Icon = SetDefault(config.Icon, "mdi:thermometer")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			case "%":
				config.DeviceClass = SetDefault(config.DeviceClass, "battery")
				config.Icon = SetDefault(config.Icon, "mdi:home-battery-outline")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value | float }}")

			default:
				config.DeviceClass = SetDefault(config.DeviceClass, "")
				config.Icon = SetDefault(config.Icon, "")
				config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")
		}

		if config.LastReset != "" {
			break
		}

		pt := api.GetDevicePoint(config.FullId)
		if !pt.Valid {
			break
		}

		if config.StateClass == "instant" {
			config.StateClass = "measurement"
			break
		}

		if config.StateClass == "" {
			config.StateClass = "measurement"
			break
		}

		config.LastReset = pt.WhenReset()
		config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | as_datetime() }}")
		// config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | int | timestamp_local | as_datetime }}")

		if config.LastReset == "" {
			config.StateClass = "measurement"
			break
		}
		config.StateClass = "total"
	}
}

func SetDefault(value string, def string) string {
	if value == "" {
		value = def
	}
	return value
}
