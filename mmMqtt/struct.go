package mmMqtt

import (
	"GoSungrow/Only"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"strconv"
	"time"
)


type Mqtt struct {
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`

	url       *url.URL
	client    mqtt.Client
	pubClient mqtt.Client
	clientOptions *mqtt.ClientOptions
	err           error
}

func New(req Mqtt) *Mqtt {
	var ret Mqtt

	for range Only.Once {
		ret.err = ret.setUrl(req)
		if ret.err != nil {
			break
		}
	}

	return &ret
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
			m.err = errors.New("host empty")
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

func (m *Mqtt) SensorPublishConfig(id string, name string, units string, address int) error {
	for range Only.Once {
		a := strconv.Itoa(address)
		class := ""
		switch units {
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

		payload := Sensor {
			Device: Device {
				Connections:  [][]string{{"sungrow_address", a}},
				Identifiers:  []string{id, "sungrow_address_" + a},
				Manufacturer: "MickMake",
				Model:        "SunGrow inverter",
				Name:         name,
				SwVersion:    "GoSunGrow https://github.com/MickMake/GoSungrow",
				ViaDevice:    "gosungrow",
			},
			Name:              "SunGrow " + name,
			StateClass:        "measurement",
			StateTopic:        SensorBaseTopic + "/" + id + "/state",
			UniqueId:          id,
			UnitOfMeasurement: units,
			DeviceClass:       class,
			Qos:               0,
			ForceUpdate: true,
			ExpireAfter: 0,
			Encoding: "utf-8",
			EnabledByDefault: true,
		}

		topic := SensorBaseTopic + "/" + id + "/config"
		m.client.Publish(topic, 0, true, payload.Json())
	}
	return m.err
}

func (m *Mqtt) SensorPublishState(sensor string, payload interface{}) error {
	for range Only.Once {
		m.client.Publish(SensorBaseTopic + "/" + sensor + "/state", 0, true, payload)
	}
	return m.err
}


// func (m *Mqtt) SetKeyFile(path string) error {
//
// 	for range Only.Once {
// 		if path == "" {
// 			break
// 		}
//
// 		m.err = checkKeyFile(path)
// 		if m.err != nil {
// 			break
// 		}
//
// 		m.KeyFile = path
// 	}
//
// 	return m.err
// }
//
// func (m *Mqtt) SetToken(t string) error {
//
// 	for range Only.Once {
// 		if t == "" {
// 			break
// 		}
//
// 		m.Token = t
// 	}
//
// 	return m.err
// }
//
// func (m *Mqtt) SetRepo(repo string) error {
//
// 	for range Only.Once {
// 		if repo == "" {
// 			m.err = errors.New("repo empty")
// 			break
// 		}
// 		m.RepoUrl = repo
// 	}
//
// 	return m.err
// }
//
// func (m *Mqtt) SetDir(dir string) error {
//
// 	for range Only.Once {
// 		if dir == "" {
// 			m.err = errors.New("dir empty")
// 			break
// 		}
// 		m.RepoDir = dir
// 	}
//
// 	return m.err
// }
//
// func (m *Mqtt) SetDiffCmd(cmd string) error {
//
// 	for range Only.Once {
// 		if cmd == "" {
// 			cmd = "tkdiff"
// 		}
// 		m.DiffCmd = cmd
// 	}
//
// 	return m.err
// }
//
// func (m *Mqtt) IsOk() bool {
// 	var ok bool
//
// 	for range Only.Once {
// 		//if m.ApiUsername == "" {
// 		//	m.Error = errors.New("username empty")
// 		//	break
// 		//}
// 		//
// 		//if m.ApiPassword == "" {
// 		//	m.Error = errors.New("password empty")
// 		//	break
// 		//}
//
// 		if m.RepoUrl == "" {
// 			m.err = errors.New("repo empty")
// 			break
// 		}
//
// 		if m.RepoDir == "" {
// 			m.err = errors.New("repo dir empty")
// 			break
// 		}
//
// 		ok = true
// 	}
//
// 	return ok
// }
// func (m *Mqtt) IsNotOk() bool {
// 	return !m.IsOk()
// }
