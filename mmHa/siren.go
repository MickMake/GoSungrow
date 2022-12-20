package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelSiren = "siren"


func (m *Mqtt) SirenPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsSiren() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Siren {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelSiren, m.ClientId, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			Qos:          0,
			Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelSiren, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) SirenPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsSiren() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelSiren, m.ClientId, id, "state")

		value := config.Value.String()
		if value == "--" {
			value = ""
		}

		// @TODO - Real hack here. Need to properly check for JSON.
		if strings.Contains(value, `{`) || strings.Contains(value, `":`) {
			t := m.client.Publish(tag, 0, true, value)
			if !t.WaitTimeout(m.Timeout) {
				m.err = t.Error()
			}
			break
		}

		payload := MqttState{
			LastReset: config.LastReset, // m.GetLastReset(config.FullId),
			Value:     value,
		}

		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

type Siren struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// A list of available tones the siren supports. When configured, this enables the support for setting a tone and enables the tone state attribute.
	AvailableTones List `json:"available_tones,omitempty"`

	// Defines a template to generate a custom payload to send to command_topic. The variable value will be assigned with the configured payload_on or payload_off setting. The siren turn on service parameters tone, volume_level or duration can be used as variables in the template. When operation in optimistic mode the corresponding state attributes will be set. Turn on parameters will be filtered if a device misses the support.
	CommandTemplate Template `json:"command_template,omitempty"`

	// Defines a template to generate a custom payload to send to command_topic when the siren turn off service is called. By default command_template will be used as template for service turn off. The variable value will be assigned with the configured payload_off setting.
	CommandOffTemplate Template `json:"command_off_template,omitempty"`

	// The MQTT topic to publish commands to change the siren state. Without command templates, a default JSON payload like {"state":"ON", "tone": "bell", "duration": 10, "volume_level": 0.5 } is published. When the siren turn on service is called, the startup parameters will be added to the JSON payload. The state value of the JSON payload will be set to the the payload_on or payload_off configured payload.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this siren is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The name to use when displaying this siren.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if siren works in optimistic mode.
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: true if no state_topic defined, else false.
	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload that represents off state. If specified, will be used for both comparing to the value in the state_topic (see value_template and state_off for details) and sending as off command to the command_topic.
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The payload that represents on state. If specified, will be used for both comparing to the value in the state_topic (see value_template and state_on for details) and sending as on command to the command_topic.
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The payload that represents the off state. Used when value that represents off state in the state_topic is different from value that should be sent to the command_topic to turn the device off.
	StateOff String `json:"state_off,omitempty"`

	// Default: payload_off if defined, else 'OFF'
	// The payload that represents the on state. Used when value that represents on state in the state_topic is different from value that should be sent to the command_topic to turn the device on.
	StateOn String `json:"state_on,omitempty"`

	// Default: payload_on if defined, else 'ON'
	// The MQTT topic subscribed to receive state updates. The state update may be either JSON or a simple string. When a JSON payload is detected, the state value of the JSON payload should supply the payload_on or payload_off defined payload to turn the siren on or off. Additionally, the state attributes duration, tone and volume_level can be updated. Use value_template to transform the received state udpate to a compliant JSON payload. Attributes will only be set if the function is supported by the device and a valid value is supplied. When a non JSON payload is detected, it should be either of the payload_on or payload_off defined payloads or None to reset the siren’s state to unknown. The initial state will be unknown. The state will be reset to unknown if a None payload or null JSON value is received as a state update.
	StateTopic String `json:"state_topic,omitempty"`

	// Defines a template to extract device’s state from the state_topic. To determine the siren’s state result of this template will be compared to state_on and state_off. Alternatively value_template can be used to render to a valid JSON payload.
	StateValueTemplate String `json:"state_value_template,omitempty"`

	// Set to true if the MQTT siren supports the duration service turn on parameter and enables the duration state attribute.
	SupportDuration Boolean `json:"support_duration,omitempty" default:"true"`

	// Set to true if the MQTT siren supports the volume_set service turn on parameter and enables the volume_level state attribute.
	SupportVolumeSet Boolean `json:"support_volume_set,omitempty" default:"true"`

	// An ID that uniquely identifies this siren device. If two sirens have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}


func (c *Siren) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsSiren() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelSiren {
			ok = true
			break
		}
	}

	return ok
}
