package cmdHassio

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/anicoll/gosungrow/iSolarCloud/AppService/getDeviceList"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
	"github.com/anicoll/gosungrow/pkg/cmdlog"
	"github.com/anicoll/gosungrow/pkg/only"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	ClientId     string        `json:"client_id"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	Host         string        `json:"host"`
	Port         string        `json:"port"`
	Timeout      time.Duration `json:"timeout"`
	EntityPrefix string        `json:"entity_prefix"`

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
	logger   cmdlog.Log
}

const (
	OptionLogLevel = "mqtt_loglevel"
	OptionDebug    = "mqtt_debug"
)

func New(req Mqtt) *Mqtt {
	var ret Mqtt

	for range only.Once {
		ret = Mqtt{
			ClientId:       req.ClientId,
			Username:       req.Username,
			Password:       req.Password,
			Host:           req.Host,
			Port:           req.Port,
			Timeout:        time.Second * 5,
			EntityPrefix:   req.EntityPrefix,
			url:            nil,
			client:         nil,
			pubClient:      nil,
			clientOptions:  nil,
			DeviceName:     "",
			LastRefresh:    time.Time{},
			SungrowDevices: nil,
			SungrowPsIds:   make(map[valueTypes.PsId]bool),
			MqttDevices:    make(map[string]Device),
			Prefix:         "homeassistant",
			UserOptions:    Options{},
			token:          nil,
			firstRun:       true,
			err:            nil,
			logger:         cmdlog.New(cmdlog.LogLevelInfoStr),
			// logger:         cmdlog.New(req.UserOptions.Get(OptionLogLevel)),
		}

		ret.err = ret.setUrl()
		if ret.err != nil {
			break
		}

		ret.UserOptions.New()
		// ret.logger = cmdlog.New(req.UserOptions.Get(OptionLogLevel))
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
	for range only.Once {
		last := m.LastRefresh.Format("20060102")
		now := time.Now().Format("20060102")

		if last != now {
			yes = true
			break
		}
	}
	return yes
}

func (m *Mqtt) setUrl() error {
	for range only.Once {
		var u string

		if m.Host == "" {
			m.err = errors.New("HASSIO mqtt host not defined")
			break
		}
		u = m.Host

		if m.Port == "" {
			m.Port = "1883"
		}
		u = m.Host + ":" + m.Port

		if m.Username != "" {
			u = m.Username + "@" + m.Host + ":" + m.Port
		}

		if m.Password != "" {
			u = m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
		}

		// if (m.ClientCert != "") && (m.ClientKey != "") && (m.ServerCert != "") {
		// 	// load the client certificate and its associated private key, which
		// 	// are used to authenticate the client to the server
		// 	// 				"certs/client.cert.pem", "certs/client.key.pem")
		// 	m.clientKeyPair, m.err = tls.LoadX509KeyPair(m.ClientCert, m.ClientKey)
		// 	if m.err != nil {
		// 		fmt.Printf("failed to load client key pair: %v\n", m.err)
		// 		break
		// 	}
		//
		// 	// load either the server certificate or the certificate of the CA
		// 	// (Certificate Authority) which signed the server certificate
		// 	// "certs/server.cert.pem"
		// 	m.serverCertPool, m.err = modbus.LoadCertPool(m.ServerCert)
		// 	if m.err != nil {
		// 		fmt.Printf("failed to load server certificate/CA: %v\n", m.err)
		// 		break
		// 	}
		//
		// 	// tcp+tls is the moniker for MBAPS (modbus/tcp encapsulated in TLS),
		// 	// 802/tcp is the IANA-registered port for MBAPS.
		// 	// set the client-side cert and key
		// 	u = "tcp+tls://" + u
		//
		// 	m.config.TLSClientCert = &m.clientKeyPair
		// 	// set the server/CA certificate
		// 	m.config.TLSRootCAs = m.serverCertPool
		// } else {
		u = "tcp://" + u
		// }

		m.url, m.err = url.Parse(u)
		if m.err != nil {
			break
		}
	}

	return m.err
}

func (m *Mqtt) SetAuth(username string, password string) error {
	for range only.Once {
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
	for range only.Once {
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

		device := Config{
			Entry:      JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId), // m.servicePrefix
			Name:       m.ClientId,
			UniqueId:   m.ClientId, // + "_Service",
			StateTopic: "~/state",
			DeviceConfig: DeviceConfig{
				Identifiers:  []string{"GoSungrow"},
				SwVersion:    "GoSungrow https://github.com/anicoll/gosungrow",
				Name:         m.ClientId + " Service",
				Manufacturer: "MickMake",
				Model:        "SunGrow",
			},
		}

		m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "config"), 0, true, device.Json())
		if m.err != nil {
			break
		}
		m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelBinarySensor, m.ClientId, "state"), 0, true, "ON")
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

		m.err = m.CreateOption(OptionLogLevel, "Mqtt LogLevel", m.funcMqttLogLevel)
		if m.err != nil {
			break
		}

		// v := OptionDisabled
		// if m.logger.IsDebug() {
		// 	v = OptionEnabled
		// }
		m.err = m.SetOption(OptionLogLevel, m.logger.GetLogLevel())
		if m.err != nil {
			break
		}
	}

	return m.err
}

func (m *Mqtt) funcMqttLogLevel(_ mqtt.Client, msg mqtt.Message) {
	for range only.Once {
		request := strings.ToLower(string(msg.Payload()))
		m.logger.Info("Option[%s] set to '%s'\n", msg.Topic(), request)
		m.err = m.SetOption(OptionLogLevel, request)
		if m.err != nil {
			break
		}
		m.logger.SetLogLevel(request)
	}
}

// func (m *Mqtt) funcMqttDebug(_ mqtt.Client, msg mqtt.Message) {
// 	for range only.Once {
// 		request := strings.ToLower(string(msg.Payload()))
// 		cmdlog.LogPrintDate("Option[%s] set to '%s'\n", msg.Topic(), request)
// 		if request == strings.ToLower(OptionEnabled) {
// 			m.err = m.SetOption("mqtt_debug", OptionEnabled)
// 			m.debug = true
// 			break
// 		}
// 		m.err = m.SetOption("mqtt_debug", OptionDisabled)
// 		m.debug = false
// 	}
// }

func (m *Mqtt) Disconnect() error {
	for range only.Once {
		m.client.Disconnect(250)
		time.Sleep(1 * time.Second)
	}
	return m.err
}

func (m *Mqtt) createClientOptions() error {
	for range only.Once {
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
	m.logger.Debug("*%t> [%s] %s\n", client.IsConnected(), msg.Topic(), string(msg.Payload()))
}

func (m *Mqtt) Subscribe(topic string, fn mqtt.MessageHandler) error {
	for range only.Once {
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
	for range only.Once {
		m.logger.Debug("MQTT[%s] -> %v\n", topic, payload)
		t := m.client.Publish(topic, qos, retained, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}

func (m *Mqtt) PublishValue(Type string, subtopic string, value string) error {
	for range only.Once {
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

		m.logger.Debug("PublishValue - topic: '%s'\tpayload: '%s'\n", topic, value)
		t := m.client.Publish(topic, 0, true, value)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) SetDeviceConfig(swname string, parentId string, id string, name string, model string, vendor string, area string) (Device, error) {
	var ret Device

	for range only.Once {
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

		ret = Device{
			Connections:   c,
			Identifiers:   []string{JoinStringsForId(m.EntityPrefix, id)},
			Manufacturer:  vendor,
			Model:         model,
			Name:          name,
			SwVersion:     swname + " https://github.com/MickMake/" + swname,
			ViaDevice:     swname,
			SuggestedArea: area,
		}
		m.MqttDevices[id] = ret
	}

	return ret, m.err
}

// func (m *Mqtt) GetLastReset(pointType string) string {
// 	var ret string
//
// 	for range only.Once {
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
