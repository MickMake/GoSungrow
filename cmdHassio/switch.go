package cmdHassio

import (
	"encoding/json"
	"strings"

	"github.com/MickMake/GoUnify/Only"
)

const LabelSwitch = "switch"

func (m *Mqtt) SwitchPublishConfig(config EntityConfig) error {
	for range Only.Once {
		if !config.IsSwitch() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Switch{
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			StateTopic:   String(JoinStringsForTopic(m.Prefix, LabelSwitch, m.ClientId, id, "state")),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelSwitch, m.ClientId, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			Qos:          0,
			Retain:       true,

			PayloadOn:     "true",
			PayloadOff:    "false",
			StateOn:       "true",
			StateOff:      "false",
			ValueTemplate: String(config.ValueTemplate),
			Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelSwitch, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) SwitchPublishValue(config EntityConfig) error {
	for range Only.Once {
		if !config.IsSwitch() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelSwitch, m.ClientId, id, "state")

		value := config.Value.String()
		if value == "--" {
			value = ""
		}

		// @TODO - Real hack here. Need to properly check for JSON.
		if strings.Contains(value, `{`) || strings.Contains(value, `":`) {
			m.err = m.Publish(tag, 0, true, value)
			break
		}

		payload := MqttState{
			LastReset: config.LastReset, // m.GetLastReset(config.FullId),
			Value:     value,
		}
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

type Switch struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to change the switch state.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this switch is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The type/class of the switch to set the icon in the frontend.
	DeviceClass DeviceClass `json:"device_class,omitempty" default:"None"`

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

	// The name to use when displaying this switch.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if switch works in optimistic mode.
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

	// Default: payload_off if defined, else OFF
	// The payload that represents the on state. Used when value that represents on state in the state_topic is different from value that should be sent to the command_topic to turn the device on.
	StateOn String `json:"state_on,omitempty"`

	// Default: payload_on if defined, else ON
	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty"`

	// An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template to extract device’s state from the state_topic. To determine the switches’s state result of this template will be compared to state_on and state_off.
	ValueTemplate String `json:"value_template,omitempty"`
}

// type Switch struct {
// 	// availability list (optional)
// 	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
// 	Availability           *Availability `json:"availability,omitempty" required:"false"`
//
// 	// availability_template template (optional)
// 	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
// 	AvailabilityTemplate   string       `json:"availability_template,omitempty" required:"false"`
//
// 	// availability_topic string (optional)
// 	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
// 	AvailabilityTopic      string       `json:"availability_topic,omitempty" required:"false"`
//
// 	// availability_mode string (optional, default: latest)
// 	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
// 	AvailabilityMode       string       `json:"availability_mode,omitempty" required:"false"`
//
// 	// command_topic string (optional)
// 	// The MQTT topic to publish commands to change the switch state.
// 	CommandTopic           string `json:"command_topic" required:"true"`
//
// 	// device map (optional)
// 	// Information about the device this switch is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
// 	Device                 Device `json:"device,omitempty" required:"false"`
//
// 	// device_class device_class (optional, default: None)
// 	// The type/class of the switch to set the icon in the frontend.
// 	DeviceClass            string       `json:"device_class,omitempty" required:"false"`
//
// 	// enabled_by_default boolean (optional, default: true)
// 	// Flag which defines if the entity should be enabled when first added.
// 	EnabledByDefault       bool         `json:"enabled_by_default,omitempty" required:"false"`
//
// 	// encoding string (optional, default: utf-8)
// 	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
// 	Encoding               string       `json:"encoding,omitempty" required:"false"`
//
// 	// entity_category string (optional, default: None)
// 	// The category of the entity.
// 	EntityCategory         string       `json:"entity_category,omitempty" required:"false"`
//
// 	// icon string (optional)
// 	// Icon for the entity.
// 	Icon                   string `json:"icon,omitempty" required:"false"`
//
// 	// json_attributes_template template (optional)
// 	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
// 	JsonAttributesTemplate string       `json:"json_attributes_template,omitempty" required:"false"`
//
// 	// json_attributes_topic string (optional)
// 	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
// 	JsonAttributesTopic    string       `json:"json_attributes_topic,omitempty" required:"false"`
//
// 	// name string (optional, default: MQTT Switch)
// 	// The name to use when displaying this switch.
// 	Name                   string `json:"name,omitempty" required:"false"`
//
// 	// object_id string (optional)
// 	// Used instead of name for automatic generation of entity_id
// 	ObjectId               string `json:"object_id,omitempty" required:"false"`
//
// 	// optimistic boolean (optional)
// 	// Flag that defines if switch works in optimistic mode.
// 	// Default: true if no state_topic defined, else false.
// 	Optimistic             bool   `json:"optimistic,omitempty" required:"false"`
//
// 	// payload_available string (optional, default: online)
// 	// The payload that represents the available state.
// 	PayloadAvailable       string `json:"payload_available,omitempty"`
//
// 	// payload_not_available string (optional, default: offline)
// 	// The payload that represents the unavailable state.
// 	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
//
// 	// payload_off string (optional, default: OFF)
// 	// The payload that represents off state. If specified, will be used for both comparing to the value in the state_topic (see value_template and state_off for details) and sending as off command to the command_topic.
// 	PayloadOff             string `json:"payload_off,omitempty"`
//
// 	// payload_on string (optional, default: ON)
// 	// The payload that represents on state. If specified, will be used for both comparing to the value in the state_topic (see value_template and state_on for details) and sending as on command to the command_topic.
// 	PayloadOn              string `json:"payload_on,omitempty"`
//
// 	// qos integer (optional, default: 0)
// 	// The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
// 	Qos                    int    `json:"qos,omitempty" required:"false"`
//
// 	// retain boolean (optional, default: false)
// 	// If the published message should have the retain flag on or not.
// 	Retain                 bool   `json:"retain,omitempty" required:"false"`
//
// 	// state_off string (optional)
// 	// The payload that represents the off state. Used when value that represents off state in the state_topic is different from value that should be sent to the command_topic to turn the device off.
// 	// Default: payload_off if defined, else OFF
// 	StateOff               string `json:"state_off,omitempty"`
//
// 	// state_on string (optional)
// 	// The payload that represents the on state. Used when value that represents on state in the state_topic is different from value that should be sent to the command_topic to turn the device on.
// 	// Default: payload_on if defined, else ON
// 	StateOn                string `json:"state_on,omitempty"`
//
// 	// state_topic string (optional)
// 	// The MQTT topic subscribed to receive state updates.
// 	StateTopic             string `json:"state_topic,omitempty" required:"false"`
//
// 	// unique_id string (optional)
// 	// An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception.
// 	UniqueId               string `json:"unique_id,omitempty" required:"false"`
//
// 	// value_template value_templatestring (optional)
// 	// Defines a template to extract device’s state from the state_topic. To determine the switches’s state result of this template will be compared to state_on and state_off.
// 	ValueTemplate          string `json:"value_template,omitempty" required:"false"`
//
// 	// CommandTemplate        string `json:"command_template" required:"false"`
//
// 	// CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// StateFunc   func() string                   `json:"-"`
// 	//
// 	// UpdateInterval  float64 `json:"-"`
// 	// ForceUpdateMQTT bool    `json:"-"`
// 	//
// 	// messageHandler mqtt.MessageHandler
// }

func (c *Switch) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsSwitch() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelSwitch {
			ok = true
			break
		}
	}

	return ok
}
