package cmdHassio

import (
	"encoding/json"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
)

const LabelText = "text"

func (m *Mqtt) TextPublishConfig(config EntityConfig) error {
	for range only.Once {
		if !config.IsText() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Text{
			Device: newDevice,
			Name:   String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelText, m.ClientId, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			Qos:          0,
			Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			// Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelText, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) TextPublishValue(config EntityConfig) error {
	for range only.Once {
		if !config.IsText() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelText, m.ClientId, id, "state")

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

type Text struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// Defines a template to generate the payload to send to command_topic.
	CommandTemplate Template `json:"command_template,omitempty"`

	// The MQTT topic to publish the text value that is set.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this of text capability is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as entity attributes. Implies force_update of the current select state when a message is received on this topic.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The maximum size of a text being set or received (maximum is 255).
	Max Integer `json:"max,omitempty" default:"255"`

	// The minimum size of a text being set or received.
	Min Integer `json:"min,omitempty" default:"0"`

	// The mode off the text entity. Must be either text or password.
	Mode String `json:"mode,omitempty" default:"text"`

	// The name of the text entity.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// A valid regular expression the text being set or received must match with.
	Pattern String `json:"pattern,omitempty"`

	// The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The MQTT topic subscribed to receive text state updates. Text state updates should match the pattern (if set) and meet the size constraints min and max. Can be used with value_template to render the incoming payload to a text update.
	StateTopic String `json:"state_topic,omitempty"`

	// An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template to extract the text state value from the payload received on state_topic.
	ValueTemplate Template `json:"value_template,omitempty"`
}

func (c *Text) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsText() bool {
	var ok bool

	for range only.Once {
		if config.Units == LabelText {
			ok = true
			break
		}
	}

	return ok
}
