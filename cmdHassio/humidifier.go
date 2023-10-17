package cmdHassio

import (
	"encoding/json"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
)

const LabelHumidifier = "humidifier"

func (m *Mqtt) HumidifierPublishConfig(config EntityConfig) error {
	for range only.Once {
		if !config.IsHumidifier() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Humidifier{
			Device: newDevice,
			Name:   String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelHumidifier, m.ClientId, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			Qos:          0,
			Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			Icon: Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelHumidifier, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) HumidifierPublishValue(config EntityConfig) error {
	for range only.Once {
		if !config.IsHumidifier() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelHumidifier, m.ClientId, id, "state")

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

type Humidifier struct {
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

	// The MQTT topic to publish commands to change the humidifier state.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this humidifier is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The device class of the MQTT device. Must be either humidifier or dehumidifier.
	DeviceClass String `json:"device_class,omitempty" default:"humidifier"`

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

	// The minimum target humidity percentage that can be set.
	MaxHumidity Integer `json:"max_humidity,omitempty" default:"100"`

	// The maximum target humidity percentage that can be set.
	MinHumidity Integer `json:"min_humidity,omitempty" default:"0"`

	// The name of the humidifier.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if humidifier works in optimistic mode
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: true if no state topic defined, else false.
	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload that represents the stop state.
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The payload that represents the running state.
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// A special payload that resets the target_humidity state attribute to None when received at the target_humidity_state_topic.
	PayloadResetHumidity String `json:"payload_reset_humidity,omitempty" default:"None"`

	// A special payload that resets the mode state attribute to None when received at the mode_state_topic.
	PayloadResetMode String `json:"payload_reset_mode,omitempty" default:"None"`

	// Defines a template to generate the payload to send to target_humidity_command_topic.
	TargetHumidityCommandTemplate Template `json:"target_humidity_command_template,omitempty"`

	// The MQTT topic to publish commands to change the humidifier target humidity state based on a percentage.
	TargetHumidityCommandTopic String `json:"target_humidity_command_topic,omitempty" required:"true"`

	// The MQTT topic subscribed to receive humidifier target humidity.
	TargetHumidityStateTopic String `json:"target_humidity_state_topic,omitempty"`

	// Defines a template to extract a value for the humidifier target_humidity state.
	TargetHumidityStateTemplate String `json:"target_humidity_state_template,omitempty"`

	// Defines a template to generate the payload to send to mode_command_topic.
	ModeCommandTemplate Template `json:"mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the mode on the humidifier. This attribute ust be configured together with the modes attribute.
	ModeCommandTopic String `json:"mode_command_topic,omitempty"`

	// The MQTT topic subscribed to receive the humidifier mode.
	ModeStateTopic String `json:"mode_state_topic,omitempty"`

	// Defines a template to extract a value for the humidifier mode state.
	ModeStateTemplate String `json:"mode_state_template,omitempty"`

	// List of available modes this humidifier is capable of running at. Common examples include normal, eco, away, boost, comfort, home, sleep, auto and baby. These examples offer built-in translations but other custom modes are allowed as well. This attribute ust be configured together with the mode_command_topic attribute.
	Modes List `json:"modes,omitempty"(optional, default: [])`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"true"`

	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty"`

	// Defines a template to extract a value from the state.
	StateValueTemplate String `json:"state_value_template,omitempty"`

	// An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}

func (c *Humidifier) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsHumidifier() bool {
	var ok bool

	for range only.Once {
		if config.Units == LabelHumidifier {
			ok = true
			break
		}
	}

	return ok
}
