package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelAlarmControlPanel = "alarm_control_panel"


func (m *Mqtt) AlarmControlPanelPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsAlarmControlPanel() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := AlarmControlPanel {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelAlarmControlPanel, m.ClientId, id, "cmd")),
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
		}

		tag := JoinStringsForTopic(m.Prefix, LabelAlarmControlPanel, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) AlarmControlPanelPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsAlarmControlPanel() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelAlarmControlPanel, m.ClientId, id, "state")

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

type AlarmControlPanel struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values REMOTE_CODE (numeric code) or REMOTE_CODE_TEXT (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use command_template to send the code to the remote device. Example configurations for remote code validation can be found here.
	Code String `json:"code,omitempty"`

	// If true the code is required to arm the alarm. If false the code is not validated.
	CodeArmRequired Boolean `json:"code_arm_required,omitempty" default:"true"`

	// If true the code is required to disarm the alarm. If false the code is not validated.
	CodeDisarmRequired Boolean `json:"code_disarm_required,omitempty" default:"true"`

	// If true the code is required to trigger the alarm. If false the code is not validated.
	CodeTriggerRequired Boolean `json:"code_trigger_required,omitempty" default:"true"`

	// The template used for the command payload. Available variables: action and code.
	CommandTemplate String `json:"command_template,omitempty" default:"action"`

	// The MQTT topic to publish commands to change the alarm state.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this alarm panel is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
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

	// The name of the alarm.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The payload to set armed-away mode on your Alarm Panel.
	PayloadArmAway String `json:"payload_arm_away,omitempty" default:"ARM_AWAY"`

	// The payload to set armed-home mode on your Alarm Panel.
	PayloadArmHome String `json:"payload_arm_home,omitempty" default:"ARM_HOME"`

	// The payload to set armed-night mode on your Alarm Panel.
	PayloadArmNight String `json:"payload_arm_night,omitempty" default:"ARM_NIGHT"`

	// The payload to set armed-vacation mode on your Alarm Panel.
	PayloadArmVacation String `json:"payload_arm_vacation,omitempty" default:"ARM_VACATION"`

	// The payload to set armed-custom-bypass mode on your Alarm Panel.
	PayloadArmCustomBypass String `json:"payload_arm_custom_bypass,omitempty" default:"ARM_CUSTOM_BYPASS"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload to disarm your Alarm Panel.
	PayloadDisarm String `json:"payload_disarm,omitempty" default:"DISARM"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload to trigger the alarm on your Alarm Panel.
	PayloadTrigger String `json:"payload_trigger,omitempty" default:"TRIGGER"`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty" required:"true"`

	// An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template to extract the value.
	ValueTemplate Template `json:"value_template,omitempty"`
}


func (c *AlarmControlPanel) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsAlarmControlPanel() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelAlarmControlPanel {
			ok = true
			break
		}
	}

	return ok
}
