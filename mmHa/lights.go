package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
)


func (m *Mqtt) PublishLightConfig(id string, name string, subName string, units string, valueName string, class string) error {
	for range Only.Once {
		id = JoinStringsForId(m.EntityPrefix, m.Device.Name, id)

		payload := Light {
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

		m.client.Publish(JoinStringsForTopic(m.lightPrefix, id, "config"), 0, true, payload.Json())
	}
	return m.err
}

func (m *Mqtt) PublishLight(subtopic string, payload interface{}) error {
	for range Only.Once {
		t := m.client.Publish(JoinStringsForTopic(m.lightPrefix, subtopic), 0, true, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}

func (m *Mqtt) PublishLightState(topic string, payload interface{}) error {
	for range Only.Once {
		topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
		t := m.client.Publish(JoinStringsForTopic(m.lightPrefix, topic, "state"), 0, true, payload)
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}


type Light struct {
	AvailabilityTopic        string   `json:"availability_topic,omitempty"`
	BrightnessCommandTopic   string   `json:"brightness_command_topic,omitempty"`
	BrightnessScale          int      `json:"brightness_scale,omitempty"`
	BrightnessStateTopic     string   `json:"brightness_state_topic,omitempty"`
	BrightnessValueTemplate  string   `json:"brightness_value_template,omitempty"`
	ColorTempCommandTemplate string   `json:"color_temp_command_template,omitempty"`
	ColorTempCommandTopic    string   `json:"color_temp_command_topic,omitempty"`
	ColorTempStateTopic      string   `json:"color_temp_state_topic,omitempty"`
	ColorTempValueTemplate   string   `json:"color_temp_value_template,omitempty"`
	CommandTopic             string   `json:"command_topic"`
	Device                   Device   `json:"device,omitempty"`
	EffectCommandTopic       string   `json:"effect_command_topic,omitempty"`
	EffectList               []string `json:"effect_list,omitempty"`
	EffectStateTopic         string   `json:"effect_state_topic,omitempty"`
	EffectValueTemplate      string   `json:"effect_value_template,omitempty"`
	HsCommandTopic           string   `json:"hs_command_topic,omitempty"`
	HsStateTopic             string   `json:"hs_state_topic,omitempty"`
	HsValueTemplate          string   `json:"hs_value_template,omitempty"`
	JSONAttributesTemplate   string   `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic      string   `json:"json_attributes_topic,omitempty"`
	MaxMireds                int      `json:"max_mireds,omitempty"`
	MinMireds                int      `json:"min_mireds,omitempty"`
	Name                     string   `json:"name,omitempty"`
	OnCommandType            string   `json:"on_command_type,omitempty"`
	Optimistic               bool     `json:"opt,omitempty"`
	PayloadAvailable         string   `json:"payload_available,omitempty"`
	PayloadNotAvailable      string   `json:"payload_not_available,omitempty"`
	PayloadOff               string   `json:"pl_off,omitempty"`
	PayloadOn                string   `json:"pl_on,omitempty"`
	QOS                      int      `json:"qos,omitempty"`
	Retain                   bool     `json:"ret,omitempty"`
	RgbCommandTemplate       string   `json:"rgb_command_template,omitempty"`
	RgbCommandTopic          string   `json:"rgb_command_topic,omitempty"`
	RgbStateTopic            string   `json:"rgb_state_topic,omitempty"`
	RgbValueTemplate         string   `json:"rgb_value_template,omitempty"`
	Schema                   string   `json:"schema,omitempty"`
	StateTopic               string   `json:"state_topic,omitempty"`
	StateValueTemplate       string   `json:"state_value_template,omitempty"`
	UniqueID                 string   `json:"unique_id,omitempty"`
	WhiteValueCommandTopic   string   `json:"white_value_command_topic,omitempty"`
	WhiteValueScale          int      `json:"white_value_scale,omitempty"`
	WhiteValueStateTopic     string   `json:"white_value_state_topic,omitempty"`
	WhiteValueTemplate       string   `json:"white_value_template,omitempty"`
	XyCommandTopic           string   `json:"xy_command_topic,omitempty"`
	XyStateTopic             string   `json:"xy_state_topic,omitempty"`
	XyValueTemplate          string   `json:"xy_value_template,omitempty"`
	ValueTemplate            string   `json:"value_template,omitempty"`

	// BrightnessStateFunc func() string `json:"-"`
	// ColorTempStateFunc  func() string `json:"-"`
	// EffectStateFunc     func() string `json:"-"`
	// HsStateFunc         func() string `json:"-"`
	// RgbStateFunc        func() string `json:"-"`
	// StateFunc           func() string `json:"-"`
	// WhiteValueStateFunc func() string `json:"-"`
	// XyStateFunc         func() string `json:"-"`
	//
	// BrightnessCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	// ColorTempCommandFunc  func(mqtt.Message, mqtt.Client) `json:"-"`
	// CommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	// EffectCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	// HsCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	// RgbCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	// WhiteValueCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	// XyCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	//
	// UpdateInterval  float64 `json:"-"`
	// ForceUpdateMQTT bool    `json:"-"`
	//
	// messageHandler mqtt.MessageHandler
}

func (c *Light) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}


// {
//	"brightness": true,
//	"cmd_t": "homeassistant/light/cbus_20/set",
//	"device": {
//		"connections": [
//			[
//				"cbus_group_address",
//				"20"
//			]
//		],
//		"identifiers": [
//			"cbus_light_20"
//		],
//		"manufacturer": "Clipsal",
//		"model": "C-Bus Lighting Application",
//		"name": "C-Bus Light 020",
//		"sw_version": "cmqttd https://github.com/micolous/cbus",
//		"via_device": "cmqttd"
//	},
//	"name": "C-Bus Light 020",
//	"schema": "json",
//	"stat_t": "homeassistant/light/cbus_20/state",
//	"unique_id": "cbus_light_20"
// }
//
// type LightConfig struct {
// 	Name        string      `json:"name"`
// 	UniqueId    string      `json:"unique_id"`
// 	CmdT        string      `json:"cmd_t"`
// 	StatT       string      `json:"stat_t"`
// 	Schema      string      `json:"schema"`
// 	Brightness  bool        `json:"brightness"`
// 	LightDevice LightDevice `json:"device"`
// }
// type LightDevice struct {
// 	Identifiers  []string   `json:"identifiers"`
// 	Connections  [][]string `json:"connections"`
// 	SwVersion    string     `json:"sw_version"`
// 	Name         string     `json:"name"`
// 	Manufacturer string     `json:"manufacturer"`
// 	Model        string     `json:"model"`
// 	ViaDevice    string     `json:"via_device"`
// }


// {
//	"brightness": 255,
//	"cbus_source_addr": 7,
//	"state": "ON",
//	"transition": 0
// }

// type LightState struct {
// 	State          string `json:"state"`
// 	Brightness     int    `json:"brightness"`
// 	Transition     int    `json:"transition"`
// 	CbusSourceAddr int    `json:"cbus_source_addr"`
// }