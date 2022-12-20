package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelButton = "button"


func (m *Mqtt) ButtonPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsButton() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Button {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelButton, m.ClientId, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			Qos:          0,
			Retain:       true,
			Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelButton, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) ButtonPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsButton() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelButton, m.ClientId, id, "state")

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

type Button struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// Defines a template to generate the payload to send to command_topic.
	CommandTemplate Template `json:"command_template,omitempty"`

	// The MQTT topic to publish commands to trigger the button.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this button is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The type/class of the button to set the icon in the frontend.
	DeviceClass DeviceClass `json:"device_class,omitempty" default:"None"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the published messages.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The name to use when displaying this button.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload To send to trigger the button.
	PayloadPress String `json:"payload_press,omitempty" default:"PRESS"`

	// The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// An ID that uniquely identifies this button entity. If two buttons have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}

// type Button struct {
// 	Availability           *Availability `json:"availability,omitempty" required:"false"`
//
// 	AvailabilityTopic      string       `json:"availability_topic,omitempty" required:"false"`
//
// 	AvailabilityMode       string       `json:"availability_mode,omitempty" required:"false"`
//
// 	AvailabilityTemplate   string       `json:"availability_template,omitempty" required:"false"`
//
// 	CommandTemplate        string `json:"command_template" required:"false"`
//
// 	CommandTopic           string `json:"command_topic" required:"true"`
//
// 	Device                 Device `json:"device,omitempty" required:"false"`
//
// 	Icon                   string `json:"icon,omitempty" required:"false"`
//
// 	JsonAttributesTemplate string       `json:"json_attributes_template,omitempty" required:"false"`
//
// 	JsonAttributesTopic    string       `json:"json_attributes_topic,omitempty" required:"false"`
//
// 	Name                   string `json:"name,omitempty" required:"false"`
//
// 	ObjectId               string `json:"object_id,omitempty" required:"false"`
//
// 	Optimistic             bool   `json:"optimistic,omitempty" required:"false"`
//
// 	PayloadAvailable       string `json:"pl_avail,omitempty"`
//
// 	PayloadNotAvailable    string `json:"pl_not_avail,omitempty"`
//
// 	PayloadOff             string `json:"pl_off,omitempty"`
//
// 	PayloadOn              string `json:"pl_on,omitempty"`
//
// 	Qos                    int    `json:"qos,omitempty" required:"false"`
//
// 	Retain                 bool   `json:"ret,omitempty" required:"false"`
//
// 	StateOff               string `json:"stat_off,omitempty"`
//
// 	StateOn                string `json:"stat_on,omitempty"`
//
// 	StateTopic             string `json:"state_topic,omitempty" required:"false"`
//
// 	UniqueId               string `json:"unique_id,omitempty" required:"false"`
//
// 	ValueTemplate          string `json:"value_template,omitempty" required:"false"`
//
// 	// CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// StateFunc   func() string                   `json:"-"`
// 	//
// 	// UpdateInterval  float64 `json:"-"`
// 	// ForceUpdateMQTT bool    `json:"-"`
// 	//
// 	// messageHandler mqtt.MessageHandler
// }

func (c *Button) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsButton() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelButton {
			ok = true
			break
		}
	}

	return ok
}
