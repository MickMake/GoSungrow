package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelFan = "fan"


func (m *Mqtt) FanPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsFan() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Fan {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelFan, m.ClientId, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelFan, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) FanPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsFan() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelFan, m.ClientId, id, "state")

		value := config.Value.String()
		if value == "--" {
			value = ""
		}

		// @TODO - Real hack here. Need to properly check for JSON.
		if strings.Contains(value, `{`) || strings.Contains(value, `":`) {
			m.err = m.Publish(tag, 0, true, value)
			break
		}

		payload := MqttState {
			LastReset: config.LastReset, // m.GetLastReset(config.FullId),
			Value:     value,
		}
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

type Fan struct {
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

	// The MQTT topic to publish commands to change the fan state.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this fan is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
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

	// The name of the fan.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if fan works in optimistic mode
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: true if no state topic defined, else false.
	// Defines a template to generate the payload to send to oscillation_command_topic.
	OscillationCommandTemplate Template `json:"oscillation_command_template,omitempty"`

	// The MQTT topic to publish commands to change the oscillation state.
	OscillationCommandTopic String `json:"oscillation_command_topic,omitempty"`

	// The MQTT topic subscribed to receive oscillation state updates.
	OscillationStateTopic String `json:"oscillation_state_topic,omitempty"`

	// Defines a template to extract a value from the oscillation.
	OscillationValueTemplate String `json:"oscillation_value_template,omitempty"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload that represents the stop state.
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The payload that represents the running state.
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// The payload that represents the oscillation off state.
	PayloadOscillationOff String `json:"payload_oscillation_off,omitempty" default:"oscillate_off"`

	// The payload that represents the oscillation on state.
	PayloadOscillationOn String `json:"payload_oscillation_on,omitempty" default:"oscillate_on"`

	// A special payload that resets the percentage state attribute to None when received at the percentage_state_topic.
	PayloadResetPercentage String `json:"payload_reset_percentage,omitempty" default:"None"`

	// A special payload that resets the preset_mode state attribute to None when received at the preset_mode_state_topic.
	PayloadResetPresetMode String `json:"payload_reset_preset_mode,omitempty" default:"None"`

	// Defines a template to generate the payload to send to percentage_command_topic.
	PercentageCommandTemplate Template `json:"percentage_command_template,omitempty"`

	// The MQTT topic to publish commands to change the fan speed state based on a percentage.
	PercentageCommandTopic String `json:"percentage_command_topic,omitempty"`

	// The MQTT topic subscribed to receive fan speed based on percentage.
	PercentageStateTopic String `json:"percentage_state_topic,omitempty"`

	// Defines a template to extract the percentage value from the payload received on percentage_state_topic.
	PercentageValueTemplate String `json:"percentage_value_template,omitempty"`

	// Defines a template to generate the payload to send to preset_mode_command_topic.
	PresetModeCommandTemplate Template `json:"preset_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the preset mode.
	PresetModeCommandTopic String `json:"preset_mode_command_topic,omitempty"`

	// The MQTT topic subscribed to receive fan speed based on presets.
	PresetModeStateTopic String `json:"preset_mode_state_topic,omitempty"`

	// Defines a template to extract the preset_mode value from the payload received on preset_mode_state_topic.
	PresetModeValueTemplate String `json:"preset_mode_value_template,omitempty"`

	// List of preset modes this fan is capable of running at. Common examples include auto, smart, whoosh, eco and breeze.
	PresetModes List `json:"preset_modes,omitempty"(optional, default: [])`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"true"`

	// The maximum of numeric output range (representing 100 %). The number of speeds within the speed_range / 100 will determine the percentage_step.
	SpeedRangeMax Integer `json:"speed_range_max,omitempty" default:"100"`

	// The minimum of numeric output range (off not included, so speed_range_min - 1 represents 0 %). The number of speeds within the speed_range / 100 will determine the percentage_step.
	SpeedRangeMin Integer `json:"speed_range_min,omitempty" default:"1"`

	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty"`

	// Defines a template to extract a value from the state.
	StateValueTemplate String `json:"state_value_template,omitempty"`

	// An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}


func (c *Fan) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsFan() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelFan {
			ok = true
			break
		}
	}

	return ok
}
