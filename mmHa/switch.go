package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
)


func (m *Mqtt) SwitchPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsSwitch() {
			break
		}

		device := m.Device
		device.Name = JoinStrings(m.Device.Name, config.ParentId)
		device.Connections = [][]string{
			{ m.Device.Name, JoinStringsForId(m.Device.Name, config.ParentId) },
			{ JoinStringsForId(m.Device.Name, config.ParentId), JoinStringsForId(m.Device.Name, config.ParentId, config.Name) },
		}
		device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId) }
		st := JoinStringsForId(m.Device.Name, config.ParentId, config.Name)

		payload := Switch {
			Device:                 device,
			Name:                   JoinStrings(m.Device.Name, config.ParentName, config.FullName),
			StateTopic:             JoinStringsForTopic(m.binarySensorPrefix, st, "state"),
			CommandTopic:           JoinStringsForTopic(m.binarySensorPrefix, st, "cmd"),
			UniqueId:               st,
			Qos:                    0,
			Retain:                 true,

			PayloadOn:              "true",
			PayloadOff:             "false",
			ValueTemplate:          config.ValueTemplate,
			Icon:                   config.Icon,

			// Availability:           &Availability {
			// 	PayloadAvailable:    "",
			// 	PayloadNotAvailable: "",
			// 	Topic:               "",
			// 	ValueTemplate:       "",
			// },
			// AvailabilityMode:       "",
			// AvailabilityTemplate:   "",
			// AvailabilityTopic:      "",
			// CommandTopic:           "",
			// JSONAttributesTemplate: "",
			// JSONAttributesTopic:    "",
			// Optimistic:             false,
			// PayloadAvailable:       "",
			// PayloadNotAvailable:    "",
			// StateOff:               "",
			// StateOn:                "",
		}

		ct := JoinStringsForTopic(m.binarySensorPrefix, st, "config")
		t := m.client.Publish(ct, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) SwitchPublishValue(config EntityConfig) error {

	for range Only.Once {
		if config.Units != LabelSwitch {
			break
		}

		st := JoinStringsForId(m.Device.Name, config.ParentId, config.Name)
		payload := MqttState {
			LastReset: m.GetLastReset(config.UniqueId),
			Value:     config.Value,
		}
		st = JoinStringsForTopic(m.binarySensorPrefix, st, "state")
		t := m.client.Publish(st, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}


// func (m *Mqtt) PublishSwitchConfig(id string, name string, subName string, units string, valueName string, class string) error {
// 	for range Only.Once {
// 		id = JoinStringsForId(m.EntityPrefix, m.Device.Name, id)
//
// 		payload := Switch {
// 			Device:                 m.Device,
// 			Name:                   JoinStrings(m.Device.ViaDevice, name),
// 			StateTopic:             JoinStringsForTopic(m.switchPrefix, id, "state"),
// 			// StateClass:             "measurement",
// 			// UniqueId:               id,
// 			// UnitOfMeasurement:      units,
// 			// DeviceClass:            class,
// 			// Qos:                    0,
// 			// ForceUpdate:            true,
// 			// ExpireAfter:            0,
// 			// Encoding:               "utf-8",
// 			// EnabledByDefault:       true,
// 			// LastResetValueTemplate: LastResetValueTemplate,
// 			// LastReset:              LastReset,
// 			// ValueTemplate:          "{{ value_json.value | float }}",
// 			// LastReset: time.Now().Format("2006-01-02T00:00:00+00:00"),
// 			// LastResetValueTemplate: "{{entity_id}}",
// 			// LastResetValueTemplate: "{{ (as_datetime((value_json.last_reset | int | timestamp_utc)|string+'Z')).isoformat() }}",
// 		}
//
// 		m.client.Publish(JoinStringsForTopic(m.switchPrefix, id, "config"), 0, true, payload.Json())
// 	}
// 	return m.err
// }
//
// func (m *Mqtt) PublishSwitch(subtopic string, payload interface{}) error {
// 	for range Only.Once {
// 		t := m.client.Publish(JoinStringsForTopic(m.switchPrefix, subtopic), 0, true, payload)
// 		if !t.WaitTimeout(m.Timeout) {
// 			m.err = t.Error()
// 		}
// 	}
// 	return m.err
// }
//
// func (m *Mqtt) PublishSwitchState(topic string, payload interface{}) error {
// 	for range Only.Once {
// 		topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
// 		t := m.client.Publish(JoinStringsForTopic(m.switchPrefix, topic, "state"), 0, true, payload)
// 		if !t.WaitTimeout(m.Timeout) {
// 			m.err = t.Error()
// 		}
// 	}
// 	return m.err
// }


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
	Qos                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"ret,omitempty"`
	StateOff               string `json:"stat_off,omitempty"`
	StateOn                string `json:"stat_on,omitempty"`
	StateTopic             string `json:"stat_t,omitempty"`
	UniqueId               string `json:"uniq_id,omitempty"`
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
