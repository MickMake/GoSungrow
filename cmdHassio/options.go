package cmdHassio

import (
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"errors"
	"github.com/MickMake/GoUnify/Only"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)


type Options struct {
	mu  sync.Mutex
	Map map[string]Option
}

type Option struct {
	Config  *EntityConfig
	Handler mqtt.MessageHandler
	Values  []string
}


func (m *Mqtt) SetOption(id string, name string, fn mqtt.MessageHandler, options ...string) error {
	for range Only.Once {
		m.err = m.UserOptions.SetOption(id, name, fn, options...)
		if m.err != nil {
			break
		}

		ec := m.UserOptions.EntityConfig(id)
		m.err = m.SelectPublishConfig(*ec, fn)
		if m.err != nil {
			break
		}
	}

	return m.err
}

func (m *Mqtt) SetOptionValue(id string, value string) error {
	for range Only.Once {
		m.err = m.UserOptions.SetOptionValue(id, value)
		if m.err != nil {
			break
		}

		ec := m.UserOptions.EntityConfig(id)
		m.err = m.SelectPublishValue(*ec)
		if m.err != nil {
			break
		}
	}

	return m.err
}

func (m *Mqtt) GetOption(id string) string {
	return m.UserOptions.GetOption(id)
}


const OptionEnabled = "Enabled"
const OptionDisabled = "Disabled"

func (m *Options) New() {
	for range Only.Once {
		m.mu.Lock()
		//goland:noinspection GoDeferInLoop
		defer m.mu.Unlock()

		m.Map = make(map[string]Option, 0)
	}
}

func (m *Options) SetOption(id string, name string, handler mqtt.MessageHandler, values ...string) error {
	var err error
	for range Only.Once {
		if len(values) == 0 {
			values = []string{OptionEnabled, OptionDisabled}
		}

		m.mu.Lock()
		//goland:noinspection GoDeferInLoop
		defer m.mu.Unlock()

		m.Map[id] = Option {
			Config:  &EntityConfig {
				Name:          "Option " + name,
				FullId:        JoinStringsForId("option", id),
				Icon:          "mdi:format-list-group",
				// ValueTemplate: `{"value": "{{ value }}"}`,
				ValueTemplate: `{{ value }}`,
				Units:         LabelSelect,
				ParentName:    "options",
				Options:       values,
			},
			Handler: handler,
			Values:  values,
		}
	}

	return err
}

func (m *Options) SetOptionValue(id string, value string) error {
	var err error
	for range Only.Once {
		m.mu.Lock()
		//goland:noinspection GoDeferInLoop
		defer m.mu.Unlock()

		if _, ok := m.Map[id]; !ok {
			err = errors.New("not exist")
			break
		}

		if m.Map[id].Config.Value == nil {
			uv := valueTypes.SetUnitValueString("", "", value)
			m.Map[id].Config.Value = &uv
			break
		}

		m.Map[id].Config.Value.SetString(value)
	}
	return err
}

func (m *Options) GetOption(id string) string {
	var ret string
	for range Only.Once {
		m.mu.Lock()
		//goland:noinspection GoDeferInLoop
		defer m.mu.Unlock()

		if value, ok := m.Map[id]; ok {
			ret = value.Config.Value.String()
			break
		}
	}
	return ret
}

func (m *Options) EntityConfig(id string) *EntityConfig {
	var ret *EntityConfig
	for range Only.Once {
		m.mu.Lock()
		//goland:noinspection GoDeferInLoop
		defer m.mu.Unlock()

		if r, ok := m.Map[id]; ok {
			ret = r.Config
			break
		}
	}
	return ret
}
