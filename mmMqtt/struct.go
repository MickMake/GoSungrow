package mmMqtt

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"encoding/json"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"strconv"
	"strings"
	"time"
)


type Mqtt struct {
	ClientId      string `json:"client_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Host          string `json:"host"`
	Port          string `json:"port"`

	url           *url.URL
	client        mqtt.Client
	pubClient     mqtt.Client
	clientOptions *mqtt.ClientOptions
	LastRefresh   time.Time `json:"-"`
	PsId          int64 `json:"-"`

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

		device := Config {
			Entry:        ServiceBaseTopic,
			Name:         ServiceBaseName,
			UniqueId:     ServiceBaseUniqueId,
			StateTopic:   "~/state",
			DeviceConfig: DeviceConfig {
				Identifiers:  []string{"GoSunGrow", "SunGrow"},
				SwVersion:    "GoSunGrow https://github.com/MickMake/GoSungrow",
				Name:         ServiceBaseName,
				Manufacturer: "MickMake",
				Model:        "SunGrow",
			},
		}

		m.err = m.Publish(ServiceBaseTopic + "/config", 0, true, device.Json())
		if m.err != nil {
			break
		}

		m.err = m.Publish(ServiceBaseTopic + "/state", 0, true, "ON")
		if m.err != nil {
			break
		}

	}

	return m.err
}

const ServiceBaseName = "GoSunGrow"
const ServiceBaseUniqueId = ServiceBaseName + "_Service"
const ServiceBaseTopic = "homeassistant/sensor/" + ServiceBaseName
const SensorBaseTopic = "homeassistant/sensor/" + ServiceBaseName

func (m *Mqtt) createClientOptions() error {
	for range Only.Once {
		m.clientOptions = mqtt.NewClientOptions()
		m.clientOptions.AddBroker(fmt.Sprintf("tcp://%s", m.url.Host))
		m.clientOptions.SetUsername(m.url.User.Username())
		password, _ := m.url.User.Password()
		m.clientOptions.SetPassword(password)
		m.clientOptions.SetClientID(m.ClientId)

		m.clientOptions.WillTopic = ServiceBaseTopic + "/state"
		m.clientOptions.WillPayload = []byte("OFF")
		m.clientOptions.WillQos = 0
		m.clientOptions.WillRetained = true
		m.clientOptions.WillEnabled = true
	}
	return m.err
}

func (m *Mqtt) Subscribe(topic string) error {
	for range Only.Once {
		m.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		})
	}
	return m.err
}

func (m *Mqtt) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	for range Only.Once {
		m.client.Publish(topic, qos, retained, payload)
	}
	return m.err
}

func (m *Mqtt) SensorPublish(subtopic string, payload interface{}) error {
	for range Only.Once {
		m.client.Publish(SensorBaseTopic + "/" + subtopic, 0, true, payload)
	}
	return m.err
}

// func (m *Mqtt) SensorPublishConfig(id string, name string, units string, address int) error {
func (m *Mqtt) SensorPublishConfig(point api.DataEntry) error {
	for range Only.Once {
		a := strconv.Itoa(point.Index)
		id := strings.ReplaceAll("sungrow_" + point.PointId, ".", "-")

		class := ""
		switch point.Unit {
			case "MW":
				fallthrough
			case "kW":
				fallthrough
			case "W":
				class = "power"

			case "MWh":
				fallthrough
			case "kWh":
				fallthrough
			case "Wh":
				class = "energy"

			case "kvar":
				class = "reactive_power"

			case "Hz":
				class = "frequency"

			case "V":
				class = "voltage"

			case "A":
				class = "current"

			case "℃":
				class = "temperature"

			case "%":
				class = "battery"
		}

		LastReset := m.GetLastReset(point.PointId)
		LastResetValueTemplate := ""
		if LastReset != "" {
			LastResetValueTemplate = "{{ value_json.last_reset | as_datetime() }}"
			// LastResetValueTemplate = "{{ value_json.last_reset | int | timestamp_local | as_datetime }}"
		}

		payload := Sensor {
			Device: Device {
				Connections:  [][]string{{"sungrow_address", a}},
				Identifiers:  []string{id, "sungrow_address_" + a},
				Manufacturer: "MickMake",
				Model:        "SunGrow inverter",
				Name:         point.PointName,
				SwVersion:    "GoSunGrow https://github.com/MickMake/GoSungrow",
				ViaDevice:    "gosungrow",
			},
			Name:                   "SunGrow " + point.PointName,
			StateClass:             "measurement",
			StateTopic:             SensorBaseTopic + "/" + id + "/state",
			UniqueId:               id,
			UnitOfMeasurement:      point.Unit,
			DeviceClass:            class,
			Qos:                    0,
			ForceUpdate:            true,
			ExpireAfter:            0,
			Encoding:               "utf-8",
			EnabledByDefault:       true,
			LastResetValueTemplate: LastResetValueTemplate,
			LastReset:              LastReset,
			ValueTemplate:          "{{ value_json.value | float }}",
			// LastReset: time.Now().Format("2006-01-02T00:00:00+00:00"),
			// LastResetValueTemplate: "{{entity_id}}",
			// LastResetValueTemplate: "{{ (as_datetime((value_json.last_reset | int | timestamp_utc)|string+'Z')).isoformat() }}",
		}

		m.client.Publish(SensorBaseTopic + "/" + id + "/config", 0, true, payload.Json())
	}
	return m.err
}

// func (m *Mqtt) SensorPublishValue(id string, value string) error {
func (m *Mqtt) SensorPublishValue(point api.DataEntry) error {
	for range Only.Once {
		id := strings.ReplaceAll("sungrow_" + point.PointId, ".", "-")
		payload := MqttState {
			LastReset: m.GetLastReset(point.PointId),
			Value:     point.Value,
		}
		m.client.Publish(SensorBaseTopic + "/" + id + "/state", 0, true, payload.Json())
	}
	return m.err
}

func (m *Mqtt) GetLastReset(pointType string) string {
	var ret string

	for range Only.Once {
		pt := api.GetDevicePoint(pointType)
		if pt == nil {
			break
		}
		ret = pt.WhenReset()
	}

	return ret
}

func (m *Mqtt) SensorPublishState(id string, payload interface{}) error {
	for range Only.Once {
		id = strings.ReplaceAll("sungrow_" + id, ".", "-")
		m.client.Publish(SensorBaseTopic + "/" + id + "/state", 0, true, payload)
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