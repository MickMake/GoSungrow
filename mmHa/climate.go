package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelClimate = "climate"


func (m *Mqtt) ClimatePublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsClimate() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Climate {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			// CommandTopic: String(JoinStringsForTopic(m.switchPrefix, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelClimate, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) ClimatePublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsClimate() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelClimate, m.ClientId, id, "state")

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

type Climate struct {
	// A template to render the value received on the action_topic with.
	ActionTemplate Template `json:"action_template,omitempty"`

	// The MQTT topic to subscribe for changes of the current action. If this is set, the climate graph uses the value received as data source. Valid values: off, heating, cooling, drying, idle, fan.
	ActionTopic String `json:"action_topic,omitempty"`

	// The MQTT topic to publish commands to switch auxiliary heat.
	AuxCommandTopic String `json:"aux_command_topic,omitempty"`

	// A template to render the value received on the aux_state_topic with.
	AuxStateTemplate Template `json:"aux_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the auxiliary heat mode. If this is not set, the auxiliary heat mode works in optimistic mode (see below).
	AuxStateTopic String `json:"aux_state_topic,omitempty"`

	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// A template with which the value received on current_temperature_topic will be rendered.
	CurrentTemperatureTemplate Template `json:"current_temperature_template,omitempty"`

	// The MQTT topic on which to listen for the current temperature.
	CurrentTemperatureTopic String `json:"current_temperature_topic,omitempty"`

	// Information about the device this HVAC device is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// A template to render the value sent to the fan_mode_command_topic with.
	FanModeCommandTemplate Template `json:"fan_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the fan mode.
	FanModeCommandTopic String `json:"fan_mode_command_topic,omitempty"`

	// A template to render the value received on the fan_mode_state_topic with.
	FanModeStateTemplate Template `json:"fan_mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not set, the fan mode works in optimistic mode (see below).
	FanModeStateTopic String `json:"fan_mode_state_topic,omitempty"`

	// A list of supported fan modes.
	FanModes List `json:"fan_modes,omitempty"`

	// Default: [“auto”, “low”, “medium”, “high”]
	// Set the initial target temperature.
	Initial Integer `json:"initial,omitempty" default:"21"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// Maximum set point available.
	MaxTemp Float `json:"max_temp,omitempty"`

	// Minimum set point available.
	MinTemp Float `json:"min_temp,omitempty"`

	// A template to render the value sent to the mode_command_topic with.
	ModeCommandTemplate Template `json:"mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the HVAC operation mode.
	ModeCommandTopic String `json:"mode_command_topic,omitempty"`

	// A template to render the value received on the mode_state_topic with.
	ModeStateTemplate Template `json:"mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC operation mode. If this is not set, the operation mode works in optimistic mode (see below).
	ModeStateTopic String `json:"mode_state_topic,omitempty"`

	// A list of supported modes. Needs to be a subset of the default values.
	Modes List `json:"modes,omitempty"`

	// Default: [“auto”, “off”, “cool”, “heat”, “dry”, “fan_only”]
	// The name of the HVAC.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload that represents disabled state.
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The payload that represents enabled state.
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// The MQTT topic to publish commands to change the power state. This is useful if your device has a separate power toggle in addition to mode.
	PowerCommandTopic String `json:"power_command_topic,omitempty"`

	// The desired precision for this device. Can be used to match your actual thermostat’s precision. Supported values are 0.1, 0.5 and 1.0.
	Precision Float `json:"precision,omitempty"`

	// Default: 0.1 for Celsius and 1.0 for Fahrenheit.
	// Defines a template to generate the payload to send to preset_mode_command_topic.
	PresetModeCommandTemplate Template `json:"preset_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the preset mode.
	PresetModeCommandTopic String `json:"preset_mode_command_topic,omitempty"`

	// The MQTT topic subscribed to receive climate speed based on presets. When preset ‘none’ is received or None the preset_mode will be reset.
	PresetModeStateTopic String `json:"preset_mode_state_topic,omitempty"`

	// Defines a template to extract the preset_mode value from the payload received on preset_mode_state_topic.
	PresetModeValueTemplate String `json:"preset_mode_value_template,omitempty"`

	// List of preset modes this climate is supporting. Common examples include eco, away, boost, comfort, home, sleep and activity.
	PresetModes List `json:"preset_modes,omitempty"`

	// The maximum QoS level to be used when receiving and publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// Defines if published messages should have the retain flag set.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// A template to render the value sent to the swing_mode_command_topic with.
	SwingModeCommandTemplate Template `json:"swing_mode_command_template,omitempty"`

	// The MQTT topic to publish commands to change the swing mode.
	SwingModeCommandTopic String `json:"swing_mode_command_topic,omitempty"`

	// A template to render the value received on the swing_mode_state_topic with.
	SwingModeStateTemplate Template `json:"swing_mode_state_template,omitempty"`

	// The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not set, the swing mode works in optimistic mode (see below).
	SwingModeStateTopic String `json:"swing_mode_state_topic,omitempty"`

	// A list of supported swing modes - (optional, default: [“on”, “off”])
	SwingModes List `json:"swing_modes,omitempty"`

	// A template to render the value sent to the temperature_command_topic with.
	TemperatureCommandTemplate Template `json:"temperature_command_template,omitempty"`

	// The MQTT topic to publish commands to change the target temperature.
	TemperatureCommandTopic String `json:"temperature_command_topic,omitempty"`

	// A template to render the value sent to the temperature_high_command_topic with.
	TemperatureHighCommandTemplate Template `json:"temperature_high_command_template,omitempty"`

	// The MQTT topic to publish commands to change the high target temperature.
	TemperatureHighCommandTopic String `json:"temperature_high_command_topic,omitempty"`

	// A template to render the value received on the temperature_high_state_topic with.
	TemperatureHighStateTemplate Template `json:"temperature_high_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target high temperature. If this is not set, the target high temperature works in optimistic mode (see below).
	TemperatureHighStateTopic String `json:"temperature_high_state_topic,omitempty"`

	// A template to render the value sent to the temperature_low_command_topic with.
	TemperatureLowCommandTemplate Template `json:"temperature_low_command_template,omitempty"`

	// The MQTT topic to publish commands to change the target low temperature.
	TemperatureLowCommandTopic String `json:"temperature_low_command_topic,omitempty"`

	// A template to render the value received on the temperature_low_state_topic with.
	TemperatureLowStateTemplate Template `json:"temperature_low_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target low temperature. If this is not set, the target low temperature works in optimistic mode (see below).
	TemperatureLowStateTopic String `json:"temperature_low_state_topic,omitempty"`

	// A template to render the value received on the temperature_state_topic with.
	TemperatureStateTemplate Template `json:"temperature_state_template,omitempty"`

	// The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below).
	TemperatureStateTopic String `json:"temperature_state_topic,omitempty"`

	// Defines the temperature unit of the device, C or F. If this is not set, the temperature unit is set to the system temperature unit.
	TemperatureUnit String `json:"temperature_unit,omitempty"`

	// Step size for temperature set point.
	TempStep Float `json:"temp_step,omitempty" default:"1"`

	// An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Default template to render the payloads on all *_state_topics with.
	ValueTemplate Template `json:"value_template,omitempty"`
}


func (c *Climate) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsClimate() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelClimate {
			ok = true
			break
		}
	}

	return ok
}
