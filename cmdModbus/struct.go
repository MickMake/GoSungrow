package cmdModbus

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdLog"
	"github.com/simonvetter/modbus"
	"net/url"
	"os"
	"time"
)


type ModBus struct {
	ClientId      string `json:"client_id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Timeout       time.Duration `json:"timeout"`
	ServerCert    string `json:"server_cert"`
	ClientCert    string `json:"client_cert"`
	ClientKey     string `json:"client_key"`

	url            *url.URL
	client         *modbus.ModbusClient
	config         *modbus.ClientConfiguration
	clientKeyPair  tls.Certificate
	serverCertPool *x509.CertPool

	firstRun bool
	err      error
	debug    bool
}


func New(req ModBus) *ModBus {
	var ret ModBus

	for range Only.Once {
		ret.config = &modbus.ClientConfiguration{}

		ret.err = ret.setUrl(req)
		if ret.err != nil {
			break
		}
		ret.firstRun = true

		// ret.MqttDevices = make(map[string]Device)
		// ret.SungrowPsIds = make(map[valueTypes.PsId]bool)
		ret.Timeout = time.Second * 5
		// ret.UserOptions.New()
	}

	return &ret
}

func (m *ModBus) IsDebug() bool {
	return m.debug
}

func (m *ModBus) SetDebug(debug bool) {
	m.debug = debug
}

func (m *ModBus) LogDebug(format string, args ...interface{})  {
	if !m.debug {
		return
	}
	cmdLog.LogPrintDate(format, args...)
}

func (m *ModBus) GetError() error {
	return m.err
}

func (m *ModBus) IsError() bool {
	if m.err != nil {
		return true
	}
	return false
}

func (m *ModBus) setUrl(req ModBus) error {

	for range Only.Once {
		m.Host = req.Host
		if m.Host == "" {
			m.err = errors.New("HASSIO mqtt host not defined")
			break
		}
		m.config.URL = m.Host

		m.Port = req.Port
		if m.Port != "" {
			m.config.URL = m.Host + ":" + m.Port
		}

		m.Username = req.Username
		if m.Username != "" {
			m.config.URL = m.Username + "@" + m.Host + ":" + m.Port
		}

		m.Password = req.Password
		if m.Password != "" {
			m.config.URL = m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
		}

		if (req.ClientCert != "") && (req.ClientKey != "") && (req.ServerCert != "") {
			// load the client certificate and its associated private key, which
			// are used to authenticate the client to the server
			// 				"certs/client.cert.pem", "certs/client.key.pem")
			m.clientKeyPair, m.err = tls.LoadX509KeyPair(req.ClientCert, req.ClientKey)
			if m.err != nil {
				fmt.Printf("failed to load client key pair: %v\n", m.err)
				break
			}

			// load either the server certificate or the certificate of the CA
			// (Certificate Authority) which signed the server certificate
			// "certs/server.cert.pem"
			m.serverCertPool, m.err = modbus.LoadCertPool(req.ServerCert)
			if m.err != nil {
				fmt.Printf("failed to load server certificate/CA: %v\n", m.err)
				break
			}

			// tcp+tls is the moniker for MBAPS (modbus/tcp encapsulated in TLS),
			// 802/tcp is the IANA-registered port for MBAPS.
			// set the client-side cert and key
			m.config.URL = "tcp+tls://" + m.config.URL

			m.config.TLSClientCert = &m.clientKeyPair
			// set the server/CA certificate
			m.config.TLSRootCAs = m.serverCertPool
		} else {
			m.config.URL = "tcp://" + m.config.URL
		}

		var u *url.URL
		u, m.err = url.Parse(m.config.URL)
		if m.err != nil {
			break
		}

		fmt.Printf("Modbus URL before: %s\n", m.config.URL)
		req.config.URL = u.String()
		fmt.Printf("Modbus URL after: %s\n", m.config.URL)
	}

	return m.err
}

func (m *ModBus) SetAuth(username string, password string) error {

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

func (m *ModBus) Connect() error {
	for range Only.Once {
		m.client, m.err = modbus.NewClient(m.config)
		if m.err != nil {
			fmt.Printf("failed to create modbus client: %v\n", m.err)
			os.Exit(1)
		}

		// now that the client is created and configured, attempt to connect
		m.err = m.client.Open()
		if m.err != nil {
			fmt.Printf("failed to connect: %v\n", m.err)
			os.Exit(2)
		}


		// m.client = mqtt.NewClient(m.clientOptions)
		// token := m.client.Connect()
		// for !token.WaitTimeout(3 * time.Second) {
		// }
		// if m.err = token.Error(); m.err != nil {
		// 	break
		// }
		// if m.ClientId == "" {
		// 	m.ClientId = "GoSungrow"
		// }
		//
		// device := Config {
		// 	Entry:      JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId),	// m.servicePrefix
		// 	Name:       m.ClientId,
		// 	UniqueId:   m.ClientId, 	// + "_Service",
		// 	StateTopic:   "~/state",
		// 	DeviceConfig: DeviceConfig {
		// 		Identifiers:  []string{"GoSungrow"},
		// 		SwVersion:    "GoSungrow https://github.com/MickMake/GoSungrow",
		// 		Name:         m.ClientId + " Service",
		// 		Manufacturer: "MickMake",
		// 		Model:        "SunGrow",
		// 	},
		// }
		//
		// m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "config"), 0, true, device.Json())
		// if m.err != nil {
		// 	break
		// }
		// m.err = m.Publish(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, "state"), 0, true, "ON")
		// if m.err != nil {
		// 	break
		// }
		//
		// _, m.err = m.SetDeviceConfig(
		// 	m.DeviceName, m.DeviceName,
		// 	"options", "Options", "", m.DeviceName,
		// 	m.DeviceName,
		// )
		// if m.err != nil {
		// 	break
		// }
		//
		// m.err = m.SetOption("mqtt_debug", "ModBus Debug", m.funcModBusDebug)
		// if m.err != nil {
		// 	break
		// }
		//
		// v := OptionDisabled
		// if m.debug {
		// 	v = OptionEnabled
		// }
		// m.err = m.SetOptionValue("mqtt_debug", v)
		// if m.err != nil {
		// 	break
		// }
	}

	return m.err
}

// func (m *ModBus) funcModBusDebug(_ mqtt.Client, msg mqtt.Message) {
// 	for range Only.Once {
// 		request := strings.ToLower(string(msg.Payload()))
// 		cmdLog.LogPrintDate("Option[%s] set to '%s'\n", msg.Topic(), request)
// 		if request == strings.ToLower(OptionEnabled) {
// 			m.err = m.SetOptionValue("mqtt_debug", OptionEnabled)
// 			m.debug = true
// 			break
// 		}
// 		m.err = m.SetOptionValue("mqtt_debug", OptionDisabled)
// 		m.debug = false
// 	}
// }

func (m *ModBus) Disconnect() error {
	for range Only.Once {
		// close the connection
		m.err = m.client.Close()
		if m.err != nil {
			fmt.Printf("failed to close connection: %v\n", m.err)
		}

		// time.Sleep(1 * time.Second)
	}
	return m.err
}
