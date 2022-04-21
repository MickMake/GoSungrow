package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
)


func (m *Mqtt) PublishSwitchConfig(id string, name string, subName string, units string, valueName string, class string) error {
	for range Only.Once {
		id = JoinStringsForId(m.EntityPrefix, m.Device.Name, id)

		payload := Switch {
			Device:                 m.Device,
			Name:                   JoinStrings(m.Device.ViaDevice, name),
			StateTopic:             JoinStringsForTopic(m.switchPrefix, id, "state"),
			// StateClass:             "measurement",
			// UniqueId:               id,
			// UnitOfMeasurement:      units,
			// DeviceClass:            class,
			// Qos:                    0,
			// ForceUpdate:            true,
			// ExpireAfter:            0,
			// Encoding:               "utf-8",
			// EnabledByDefault:       true,
			// LastResetValueTemplate: LastResetValueTemplate,
			// LastReset:              LastReset,
			// ValueTemplate:          "{{ value_json.value | float }}",
			// LastReset: time.Now().Format("2006-01-02T00:00:00+00:00"),
			// LastResetValueTemplate: "{{entity_id}}",
			// LastResetValueTemplate: "{{ (as_datetime((value_json.last_reset | int | timestamp_utc)|string+'Z')).isoformat() }}",
		}

		m.client.Publish(JoinStringsForTopic(m.switchPrefix, id, "config"), 0, true, payload.Json())
	}
	return m.err
}

func (m *Mqtt) PublishSwitch(subtopic string, payload interface{}) error {
	for range Only.Once {
		t := m.client.Publish(JoinStringsForTopic(m.switchPrefix, subtopic), 0, true, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}

func (m *Mqtt) PublishSwitchState(topic string, payload interface{}) error {
	for range Only.Once {
		topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
		t := m.client.Publish(JoinStringsForTopic(m.switchPrefix, topic, "state"), 0, true, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}


type Switch struct {
	AvailabilityTopic      string `json:"avty_t,omitempty"`
	CommandTopic           string `json:"cmd_t"`
	Device                 Device `json:"dev,omitempty"`
	Icon                   string `json:"ic,omitempty"`
	JSONAttributesTemplate string `json:"json_attr_tpl,omitempty"`
	JSONAttributesTopic    string `json:"json_attr_t,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"opt,omitempty"`
	PayloadAvailable       string `json:"pl_avail,omitempty"`
	PayloadNotAvailable    string `json:"pl_not_avail,omitempty"`
	PayloadOff             string `json:"pl_off,omitempty"`
	PayloadOn              string `json:"pl_on,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"ret,omitempty"`
	StateOff               string `json:"stat_off,omitempty"`
	StateOn                string `json:"stat_on,omitempty"`
	StateTopic             string `json:"stat_t,omitempty"`
	UniqueID               string `json:"uniq_id,omitempty"`
	ValueTemplate          string `json:"val_tpl,omitempty"`

	// CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	// StateFunc   func() string                   `json:"-"`
	//
	// UpdateInterval  float64 `json:"-"`
	// ForceUpdateMQTT bool    `json:"-"`
	//
	// messageHandler mqtt.MessageHandler
}

func (c *Switch) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}
