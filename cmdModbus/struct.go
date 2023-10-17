//go:build !(freebsd && amd64)

package cmdModbus

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/anicoll/gosungrow/pkg/cmdlog"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/simonvetter/modbus"
)

type Modbus struct {
	ClientId   string        `json:"client_id"`
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	Host       string        `json:"host"`
	Port       string        `json:"port"`
	Timeout    time.Duration `json:"timeout"`
	ServerCert string        `json:"server_cert"`
	ClientCert string        `json:"client_cert"`
	ClientKey  string        `json:"client_key"`
	LogLevel   string        `json:"log_level"`

	url            *url.URL
	client         *modbus.ModbusClient
	config         *modbus.ClientConfiguration
	clientKeyPair  tls.Certificate
	serverCertPool *x509.CertPool
	endianness     modbus.Endianness
	wordOrder      modbus.WordOrder

	firstRun bool
	err      error
	logger   cmdlog.Log
}

func New(req Modbus) Modbus {
	var ret Modbus

	for range only.Once {
		ret = Modbus{
			ClientId:   req.ClientId,
			Username:   req.Username,
			Password:   req.Password,
			Host:       req.Host,
			Port:       req.Port,
			Timeout:    req.Timeout,
			ServerCert: req.ServerCert,
			ClientCert: req.ClientCert,
			ClientKey:  req.ClientKey,
			LogLevel:   req.LogLevel,

			// url:            nil,
			client: nil,
			config: &modbus.ClientConfiguration{},
			// clientKeyPair:  tls.Certificate{},
			// serverCertPool: nil,
			endianness: modbus.BIG_ENDIAN,
			wordOrder:  modbus.HIGH_WORD_FIRST,

			firstRun: true,

			// err:            nil,
			logger: cmdlog.New(req.LogLevel),
		}

		ret.err = ret.setUrl()
		if ret.err != nil {
			break
		}

		ret.Timeout = time.Second * 5
	}

	return ret
}

func (m *Modbus) SetLog(log cmdlog.Log) {
	m.logger = log
}

func (m *Modbus) GetError() error {
	return m.err
}

func (m *Modbus) IsError() bool {
	if m.err != nil {
		return true
	}
	return false
}

func (m *Modbus) setUrl() error {
	for range only.Once {
		if m.Host == "" {
			m.err = errors.New("Modbus host not defined")
			break
		}
		m.config.URL = m.Host

		if m.Port == "" {
			m.Port = DefaultPort
		}
		m.config.URL = m.Host + ":" + m.Port

		if m.Username != "" {
			m.config.URL = m.Username + "@" + m.Host + ":" + m.Port
		}

		if m.Password != "" {
			m.config.URL = m.Username + ":" + m.Password + "@" + m.Host + ":" + m.Port
		}

		if (m.ClientCert != "") && (m.ClientKey != "") && (m.ServerCert != "") {
			// load the client certificate and its associated private key, which
			// are used to authenticate the client to the server
			// 				"certs/client.cert.pem", "certs/client.key.pem")
			m.clientKeyPair, m.err = tls.LoadX509KeyPair(m.ClientCert, m.ClientKey)
			if m.err != nil {
				fmt.Printf("failed to load client key pair: %v\n", m.err)
				break
			}

			// load either the server certificate or the certificate of the CA
			// (Certificate Authority) which signed the server certificate
			// "certs/server.cert.pem"
			m.serverCertPool, m.err = modbus.LoadCertPool(m.ServerCert)
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

		m.config.URL = u.String()
	}

	return m.err
}

func (m *Modbus) SetAuth(username string, password string) error {
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

func (m *Modbus) Connect() error {
	for range only.Once {
		if m.config == nil {
			m.logger.Debug("Modbus: failed to create modbus client: %v\n", m.err)
			break
		}

		m.client, m.err = modbus.NewClient(m.config)
		if m.err != nil {
			m.logger.Debug("Modbus: failed to create modbus client: %v\n", m.err)
			break
		}

		// now that the client is created and configured, attempt to connect
		m.err = m.client.Open()
		if m.err != nil {
			fmt.Printf("Modbus: failed to connect: %v\n", m.err)
			break
		}
	}

	return m.err
}

func (m *Modbus) Disconnect() error {
	for range only.Once {
		// close the connection
		m.err = m.client.Close()
		if m.err != nil {
			fmt.Printf("failed to close connection: %v\n", m.err)
		}

		// time.Sleep(1 * time.Second)
	}
	return m.err
}

func (m *Modbus) SetLittleEndian() {
	m.endianness = modbus.LITTLE_ENDIAN
	m.err = m.client.SetEncoding(m.endianness, m.wordOrder)
}

func (m *Modbus) SetBigEndian() {
	m.endianness = modbus.BIG_ENDIAN
	m.err = m.client.SetEncoding(m.endianness, m.wordOrder)
}

func (m *Modbus) SetLowWordFirst() {
	m.wordOrder = modbus.LOW_WORD_FIRST
	m.err = m.client.SetEncoding(m.endianness, m.wordOrder)
}

func (m *Modbus) SetHighWordFirst() {
	m.wordOrder = modbus.HIGH_WORD_FIRST
	m.err = m.client.SetEncoding(m.endianness, m.wordOrder)
}
