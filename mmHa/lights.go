package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelLight = "light"


func (m *Mqtt) PublishLightConfig(config EntityConfig) error {
	// func (m *Mqtt) PublishLightConfig(id string, name string, subName string, units string, valueName string, class string) error {
	for range Only.Once {
		if !config.IsLight() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Light {
			Device:                 newDevice,
			Name:                   String(JoinStrings(m.DeviceName, config.Name)),
			StateTopic:             String(JoinStringsForTopic(m.Prefix, LabelLight, m.ClientId, id, "state")),
			UniqueId:               String(id),
			ObjectId:               String(id),

			// StateClass:             config.StateClass,
			// UniqueId:               id,
			// UnitOfMeasurement:      config.Units,
			// DeviceClass:            config.DeviceClass,
			// Qos:                    0,
			// ForceUpdate:            true,
			// ExpireAfter:            0,
			// Encoding:               "utf-8",
			// EnabledByDefault:       true,
			PayloadOn:              "true",
			PayloadOff:             "false",
			// LastResetValueTemplate: config.LastResetValueTemplate,
			StateValueTemplate:     String(config.ValueTemplate),
			Icon:                   Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelLight, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}
	return m.err
}

func (m *Mqtt) LightPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsLight() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelLight, m.ClientId, id, "state")

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
			LastReset: config.LastReset,	// m.GetLastReset(config.FullId),
			Value:     value,
		}
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

type Light struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to change the light’s brightness.
	BrightnessCommandTopic String `json:"brightness_command_topic,omitempty"`

	// Defines a template to compose message which will be sent to brightness_command_topic. Available variables: value.
	BrightnessCommandTemplate String `json:"brightness_command_template,omitempty"`

	// Defines the maximum brightness value (i.e., 100%) of the MQTT device.
	BrightnessScale Integer `json:"brightness_scale,omitempty" default:"255"`

	// The MQTT topic subscribed to receive brightness state updates.
	BrightnessStateTopic String `json:"brightness_state_topic,omitempty"`

	// Defines a template to extract the brightness value.
	BrightnessValueTemplate String `json:"brightness_value_template,omitempty"`

	// The MQTT topic subscribed to receive color mode updates. If this is not configured, color_mode will be automatically set according to the last received valid color or color temperature
	ColorModeStateTopic String `json:"color_mode_state_topic,omitempty"`

	// Defines a template to extract the color mode.
	ColorModeValueTemplate String `json:"color_mode_value_template,omitempty"`

	// Defines a template to compose message which will be sent to color_temp_command_topic. Available variables: value.
	ColorTempCommandTemplate String `json:"color_temp_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s color temperature state. The color temperature command slider has a range of 153 to 500 mireds (micro reciprocal degrees).
	ColorTempCommandTopic String `json:"color_temp_command_topic,omitempty"`

	// The MQTT topic subscribed to receive color temperature state updates.
	ColorTempStateTopic String `json:"color_temp_state_topic,omitempty"`

	// Defines a template to extract the color temperature value.
	ColorTempValueTemplate String `json:"color_temp_value_template,omitempty"`

	// The MQTT topic to publish commands to change the switch state.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this light is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// The MQTT topic to publish commands to change the light’s effect state.
	EffectCommandTopic String `json:"effect_command_topic,omitempty"`

	// Defines a template to compose message which will be sent to effect_command_topic. Available variables: value.
	EffectCommandTemplate String `json:"effect_command_template,omitempty"`

	// effect_list string | list (optional)
	// The list of effects the light supports.

	// The MQTT topic subscribed to receive effect state updates.
	EffectStateTopic String `json:"effect_state_topic,omitempty"`

	// Defines a template to extract the effect value.
	EffectValueTemplate String `json:"effect_value_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s color state in HS format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation: 0..100. Note: Brightness is sent separately in the brightness_command_topic.
	HsCommandTopic String `json:"hs_command_topic,omitempty"`

	// The MQTT topic subscribed to receive color state updates in HS format. Note: Brightness is received separately in the brightness_state_topic.
	HsStateTopic String `json:"hs_state_topic,omitempty"`

	// Defines a template to extract the HS value.
	HsValueTemplate String `json:"hs_value_template,omitempty"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The maximum color temperature in mireds.
	MaxMireds Integer `json:"max_mireds,omitempty"`

	// The minimum color temperature in mireds.
	MinMireds Integer `json:"min_mireds,omitempty"`

	// The name of the light.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Defines when on the payload_on is sent. Using last (the default) will send any style (brightness, color, etc) topics first and then a payload_on to the command_topic. Using first will send the payload_on and then any style topics. Using brightness will only send brightness commands instead of the payload_on to turn the light on.
	OnCommandType String `json:"on_command_type,omitempty"`

	// Flag that defines if switch works in optimistic mode.
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: true if no state topic defined, else false.
	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload that represents disabled state.
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The payload that represents enabled state.
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// Defines a template to compose message which will be sent to rgb_command_topic. Available variables: red, green and blue.
	RgbCommandTemplate String `json:"rgb_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s RGB state.
	RgbCommandTopic String `json:"rgb_command_topic,omitempty"`

	// The MQTT topic subscribed to receive RGB state updates. The expected payload is the RGB values separated by commas, for example, 255,0,127.
	RgbStateTopic String `json:"rgb_state_topic,omitempty"`

	// Defines a template to extract the RGB value.
	RgbValueTemplate String `json:"rgb_value_template,omitempty"`

	// Defines a template to compose message which will be sent to rgbw_command_topic. Available variables: red, green, blue and white.
	RgbwCommandTemplate String `json:"rgbw_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s RGBW state.
	RgbwCommandTopic String `json:"rgbw_command_topic,omitempty"`

	// The MQTT topic subscribed to receive RGBW state updates. The expected payload is the RGBW values separated by commas, for example, 255,0,127,64.
	RgbwStateTopic String `json:"rgbw_state_topic,omitempty"`

	// Defines a template to extract the RGBW value.
	RgbwValueTemplate String `json:"rgbw_value_template,omitempty"`

	// Defines a template to compose message which will be sent to rgbww_command_topic. Available variables: red, green, blue, cold_white and warm_white.
	RgbwwCommandTemplate String `json:"rgbww_command_template,omitempty"`

	// The MQTT topic to publish commands to change the light’s RGBWW state.
	RgbwwCommandTopic String `json:"rgbww_command_topic,omitempty"`

	// The MQTT topic subscribed to receive RGBWW state updates. The expected payload is the RGBWW values separated by commas, for example, 255,0,127,64,32.
	RgbwwStateTopic String `json:"rgbww_state_topic,omitempty"`

	// Defines a template to extract the RGBWW value.
	RgbwwValueTemplate String `json:"rgbww_value_template,omitempty"`

	// The schema to use. Must be default or omitted to select the default schema.
	Schema String `json:"schema,omitempty" default:"default"`

	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty"`

	// Defines a template to extract the state value. The template should match the payload on and off values, so if your light uses power on to turn on, your state_value_template string should return power on when the switch is on. For example if the message is just on, your state_value_template should be power .
	StateValueTemplate String `json:"state_value_template,omitempty"`

	// An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// The MQTT topic to publish commands to change the light to white mode with a given brightness.
	WhiteCommandTopic String `json:"white_command_topic,omitempty"`

	// Defines the maximum white level (i.e., 100%) of the MQTT device.
	WhiteScale Integer `json:"white_scale,omitempty" default:"255"`

	// The MQTT topic to publish commands to change the light’s XY state.
	XyCommandTopic String `json:"xy_command_topic,omitempty"`

	// The MQTT topic subscribed to receive XY state updates.
	XyStateTopic String `json:"xy_state_topic,omitempty"`

	// Defines a template to extract the XY value.
	XyValueTemplate String `json:"xy_value_template,omitempty"`
}

// type Light struct {
// 	// availability list (optional)
// 	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
// 	Availability *Availability `json:"availability,omitempty" required:"false"`
//
// 	// availability_mode string (optional, default: latest)
// 	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
// 	AvailabilityMode string `json:"availability_mode,omitempty" required:"false"`
//
// 	// availability_template template (optional)
// 	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
// 	AvailabilityTemplate string `json:"availability_template,omitempty" required:"false"`
//
// 	// availability_topic string (optional)
// 	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
// 	AvailabilityTopic string `json:"availability_topic,omitempty" required:"false"`
//
// 	// brightness_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s brightness.
// 	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`
//
// 	// brightness_command_template string (optional)
// 	// Defines a template to compose message which will be sent to brightness_command_topic. Available variables: value.
// 	BrightnessCommandTemplate string `json:"brightness_command_template,omitempty"`
//
// 	// brightness_scale integer (optional, default: 255)
// 	// Defines the maximum brightness value (i.e., 100%) of the MQTT device.
// 	BrightnessScale int `json:"brightness_scale,omitempty"`
//
// 	// brightness_state_topic string (optional)
// 	// The MQTT topic subscribed to receive brightness state updates.
// 	BrightnessStateTopic string `json:"brightness_state_topic,omitempty"`
//
// 	// brightness_value_template string (optional)
// 	// Defines a template to extract the brightness value.
// 	BrightnessValueTemplate string `json:"brightness_value_template,omitempty"`
//
// 	// color_mode_state_topic string (optional)
// 	// The MQTT topic subscribed to receive color mode updates. If this is not configured, color_mode will be automatically set according to the last received valid color or color temperature
// 	ColorModeStateTopic string `json:"color_mode_state_topic,omitempty"`
//
// 	// color_mode_value_template string (optional)
// 	// Defines a template to extract the color mode.
// 	ColorModeValueTemplate string `json:"color_mode_value_template,omitempty"`
//
// 	// color_temp_command_template string (optional)
// 	// Defines a template to compose message which will be sent to color_temp_command_topic. Available variables: value.
// 	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`
//
// 	// color_temp_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s color temperature state. The color temperature command slider has a range of 153 to 500 mireds (micro reciprocal degrees).
// 	ColorTempCommandTopic string `json:"color_temp_command_topic,omitempty"`
//
// 	// color_temp_state_topic string (optional)
// 	// The MQTT topic subscribed to receive color temperature state updates.
// 	ColorTempStateTopic string `json:"color_temp_state_topic,omitempty"`
//
// 	// color_temp_value_template string (optional)
// 	// Defines a template to extract the color temperature value.
// 	ColorTempValueTemplate string `json:"color_temp_value_template,omitempty"`
// 	// CommandTemplate        string `json:"command_template" required:"false"`
//
// 	// command_topic string REQUIRED
// 	// The MQTT topic to publish commands to change the switch state.
// 	CommandTopic string `json:"command_topic" required:"true"`
//
// 	// device map (optional)
// 	// Information about the device this light is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
// 	Device Device `json:"device,omitempty" required:"false"`
//
// 	// enabled_by_default boolean (optional, default: true)
// 	// Flag which defines if the entity should be enabled when first added.
// 	EnabledByDefault bool `json:"enabled_by_default,omitempty" required:"false"`
//
// 	// encoding string (optional, default: utf-8)
// 	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
// 	Encoding string `json:"encoding,omitempty" required:"false"`
//
// 	// entity_category string (optional, default: None)
// 	// The category of the entity.
// 	EntityCategory string `json:"entity_category,omitempty" required:"false"`
//
// 	// effect_command_template string (optional)
// 	// Defines a template to compose message which will be sent to effect_command_topic. Available variables: value.
// 	EffectCommandTemplate string `json:"effect_command_template" required:"true"`
//
// 	// effect_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s effect state.
// 	EffectCommandTopic string `json:"effect_command_topic,omitempty"`
//
// 	// effect_list string | list (optional)
// 	// The list of effects the light supports.
// 	EffectList []string `json:"effect_list,omitempty"`
//
// 	// effect_state_topic string (optional)
// 	// The MQTT topic subscribed to receive effect state updates.
// 	EffectStateTopic string `json:"effect_state_topic,omitempty"`
//
// 	// effect_value_template string (optional)
// 	// Defines a template to extract the effect value.
// 	EffectValueTemplate string `json:"effect_value_template,omitempty"`
//
// 	// hs_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s color state in HS format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation: 0..100. Note: Brightness is sent separately in the brightness_command_topic.
// 	HsCommandTopic string `json:"hs_command_topic,omitempty"`
//
// 	// hs_state_topic string (optional)
// 	// The MQTT topic subscribed to receive color state updates in HS format. Note: Brightness is received separately in the brightness_state_topic.
// 	HsStateTopic string `json:"hs_state_topic,omitempty"`
//
// 	// hs_value_template string (optional)
// 	// Defines a template to extract the HS value.
// 	HsValueTemplate string `json:"hs_value_template,omitempty"`
//
// 	// icon string (optional)
// 	// Icon for the entity.
// 	Icon string `json:"icon,omitempty" required:"false"`
//
// 	// json_attributes_template template (optional)
// 	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
// 	JsonAttributesTemplate string `json:"json_attributes_template,omitempty" required:"false"`
//
// 	// json_attributes_topic string (optional)
// 	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
// 	JsonAttributesTopic string `json:"json_attributes_topic,omitempty" required:"false"`
//
// 	// max_mireds integer (optional)
// 	// The maximum color temperature in mireds.
// 	MaxMireds int `json:"max_mireds,omitempty"`
//
// 	// min_mireds integer (optional)
// 	// The minimum color temperature in mireds.
// 	MinMireds int `json:"min_mireds,omitempty"`
//
// 	// name string (optional, default: MQTT Light)
// 	// The name of the light.
// 	Name string `json:"name,omitempty" required:"false"`
//
// 	// object_id string (optional)
// 	// Used instead of name for automatic generation of entity_id
// 	ObjectId string `json:"object_id,omitempty" required:"false"`
//
// 	// on_command_type string (optional)
// 	// Defines when on the payload_on is sent. Using last (the default) will send any style (brightness, color, etc) topics first and then a payload_on to the command_topic. Using first will send the payload_on and then any style topics. Using brightness will only send brightness commands instead of the payload_on to turn the light on.
// 	OnCommandType string `json:"on_command_type,omitempty"`
//
// 	// optimistic boolean (optional)
// 	// Flag that defines if switch works in optimistic mode.
// 	// Default: true if no state topic defined, else false.
// 	Optimistic bool `json:"optimistic,omitempty" required:"false"`
//
// 	// payload_available string (optional, default: online)
// 	// The payload that represents the available state.
// 	PayloadAvailable string `json:"payload_available,omitempty"`
//
// 	// payload_not_available string (optional, default: offline)
// 	// The payload that represents the unavailable state.
// 	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
//
// 	// payload_off string (optional, default: OFF)
// 	// The payload that represents disabled state.
// 	PayloadOff string `json:"payload_off,omitempty"`
//
// 	// payload_on string (optional, default: ON)
// 	// The payload that represents enabled state.
// 	PayloadOn string `json:"payload_on,omitempty"`
//
// 	// qos integer (optional, default: 0)
// 	// The maximum QoS level of the state topic.
// 	Qos int `json:"qos,omitempty" required:"false"`
//
// 	// retain boolean (optional, default: false)
// 	// If the published message should have the retain flag on or not.
// 	Retain bool `json:"retain,omitempty" required:"false"`
//
// 	// rgb_command_template string (optional)
// 	// Defines a template to compose message which will be sent to rgb_command_topic. Available variables: red, green and blue.
// 	RgbCommandTemplate string `json:"rgb_command_template,omitempty"`
//
// 	// rgb_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s RGB state.
// 	RgbCommandTopic string `json:"rgb_command_topic,omitempty"`
//
// 	// rgb_state_topic string (optional)
// 	// The MQTT topic subscribed to receive RGB state updates. The expected payload is the RGB values separated by commas, for example, 255,0,127.
// 	RgbStateTopic string `json:"rgb_state_topic,omitempty"`
//
// 	// rgb_value_template string (optional)
// 	// Defines a template to extract the RGB value.
// 	RgbValueTemplate string `json:"rgb_value_template,omitempty"`
//
// 	// rgbw_command_template string (optional)
// 	// Defines a template to compose message which will be sent to rgbw_command_topic. Available variables: red, green, blue and white.
// 	RgbwCommandTemplate string `json:"rgbw_command_template,omitempty"`
//
// 	// rgbw_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s RGBW state.
// 	RgbwCommandTopic string `json:"rgbw_command_topic,omitempty"`
//
// 	// rgbw_state_topic string (optional)
// 	// The MQTT topic subscribed to receive RGBW state updates. The expected payload is the RGBW values separated by commas, for example, 255,0,127,64.
// 	RgbwStateTopic string `json:"rgbw_state_topic,omitempty"`
//
// 	// rgbw_value_template string (optional)
// 	// Defines a template to extract the RGBW value.
// 	RgbwValueTemplate string `json:"rgbw_value_template,omitempty"`
//
// 	// rgbww_command_template string (optional)
// 	// Defines a template to compose message which will be sent to rgbww_command_topic. Available variables: red, green, blue, cold_white and warm_white.
// 	RgbwwCommandTemplate string `json:"rgbww_command_template,omitempty"`
//
// 	// rgbww_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s RGBWW state.
// 	RgbwwCommandTopic string `json:"rgbww_command_topic,omitempty"`
//
// 	// rgbww_state_topic string (optional)
// 	// The MQTT topic subscribed to receive RGBWW state updates. The expected payload is the RGBWW values separated by commas, for example, 255,0,127,64,32.
// 	RgbwwStateTopic string `json:"rgbww_state_topic,omitempty"`
//
// 	// rgbww_value_template string (optional)
// 	// Defines a template to extract the RGBWW value.
// 	RgbwwValueTemplate string `json:"rgbww_value_template,omitempty"`
//
// 	// schema string (optional, default: default)
// 	// The schema to use. Must be default or omitted to select the default schema.
// 	Schema string `json:"schema,omitempty"`
//
// 	// state_topic string (optional)
// 	// The MQTT topic subscribed to receive state updates.
// 	StateTopic string `json:"state_topic,omitempty" required:"false"`
//
// 	// state_value_template string (optional)
// 	// Defines a template to extract the state value. The template should match the payload on and off values, so if your light uses power on to turn on, your state_value_template string should return power on when the switch is on. For example if the message is just on, your state_value_template should be power .
// 	StateValueTemplate string `json:"state_value_template,omitempty"`
//
// 	// unique_id string (optional)
// 	// An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception.
// 	UniqueId string `json:"unique_id,omitempty" required:"false"`
//
// 	// white_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light to white mode with a given brightness.
// 	WhiteCommandTopic string `json:"white_command_topic,omitempty"`
//
// 	// white_scale integer (optional, default: 255)
// 	// Defines the maximum white level (i.e., 100%) of the MQTT device.
// 	WhiteScale int `json:"white_scale,omitempty"`
//
// 	// xy_command_topic string (optional)
// 	// The MQTT topic to publish commands to change the light’s XY state.
// 	XyCommandTopic string `json:"xy_command_topic,omitempty"`
//
// 	// xy_state_topic string (optional)
// 	// The MQTT topic subscribed to receive XY state updates.
// 	XyStateTopic string `json:"xy_state_topic,omitempty"`
//
// 	// xy_value_template string (optional)
// 	// Defines a template to extract the XY value.
// 	XyValueTemplate string `json:"xy_value_template,omitempty"`
//
// 	// WhiteValueCommandTopic   string   `json:"white_value_command_topic,omitempty"`
// 	// WhiteValueScale          int      `json:"white_value_scale,omitempty"`
// 	// WhiteValueStateTopic     string   `json:"white_value_state_topic,omitempty"`
// 	// WhiteValueTemplate       string   `json:"white_value_template,omitempty"`
// 	// ValueTemplate          string `json:"value_template,omitempty" required:"false"`
//
// 	// BrightnessStateFunc func() string `json:"-"`
// 	// ColorTempStateFunc  func() string `json:"-"`
// 	// EffectStateFunc     func() string `json:"-"`
// 	// HsStateFunc         func() string `json:"-"`
// 	// RgbStateFunc        func() string `json:"-"`
// 	// StateFunc           func() string `json:"-"`
// 	// WhiteValueStateFunc func() string `json:"-"`
// 	// XyStateFunc         func() string `json:"-"`
// 	//
// 	// BrightnessCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// ColorTempCommandFunc  func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// CommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// EffectCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// HsCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// RgbCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// WhiteValueCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
// 	// XyCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
// 	//
// 	// UpdateInterval  float64 `json:"-"`
// 	// ForceUpdateMQTT bool    `json:"-"`
// 	//
// 	// messageHandler mqtt.MessageHandler
// }

func (c *Light) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsLight() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelLight {
			ok = true
			break
		}
	}

	return ok
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