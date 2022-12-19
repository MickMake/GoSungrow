package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


const LabelVacuum = "vacuum"

func (m *Mqtt) VacuumPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsVacuum() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Vacuum {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelVacuum, m.ClientId, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelVacuum, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) VacuumPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsVacuum() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelVacuum, m.ClientId, id, "state")

		value := config.Value.String()
		if value == "--" {
			value = ""
		}

		// @TODO - Real hack here. Need to properly check for JSON.
		if strings.Contains(config.Value.String(), `{`) || strings.Contains(config.Value.String(), `":`) {
			t := m.client.Publish(tag, 0, true, config.Value.String())
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

type Vacuum struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to control the vacuum.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this switch is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// List of possible fan speeds for the vacuum.
	FanSpeedList List `json:"fan_speed_list,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The name of the vacuum.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload to send to the command_topic to begin a spot cleaning cycle.
	PayloadCleanSpot String `json:"payload_clean_spot,omitempty" default:"clean_spot"`

	// The payload to send to the command_topic to locate the vacuum (typically plays a song).
	PayloadLocate String `json:"payload_locate,omitempty" default:"locate"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload to send to the command_topic to pause the vacuum.
	PayloadPause String `json:"payload_pause,omitempty" default:"pause"`

	// The payload to send to the command_topic to tell the vacuum to return to base.
	PayloadReturnToBase String `json:"payload_return_to_base,omitempty" default:"return_to_base"`

	// The payload to send to the command_topic to begin the cleaning cycle.
	PayloadStart String `json:"payload_start,omitempty" default:"start"`

	// The payload to send to the command_topic to stop cleaning.
	PayloadStop String `json:"payload_stop,omitempty" default:"stop"`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The schema to use. Must be state to select the state schema.
	Schema String `json:"schema,omitempty" default:"legacy"`

	// The MQTT topic to publish custom commands to the vacuum.
	SendCommandTopic String `json:"send_command_topic,omitempty"`

	// The MQTT topic to publish commands to control the vacuum’s fan speed.
	SetFanSpeedTopic String `json:"set_fan_speed_topic,omitempty"`

	// The MQTT topic subscribed to receive state messages from the vacuum. Messages received on the state_topic must be a valid JSON dictionary, with a mandatory state key and optionally battery_level and fan_speed keys as shown in the example.
	StateTopic String `json:"state_topic,omitempty"`

	// supported_features string | list (optional)
	// List of features that the vacuum supports (possible values are start, stop, pause, return_home, battery, status, locate, clean_spot, fan_speed, send_command).
	SupportedFeatures List `json:"supported_features,omitempty"`

	// Default: start, stop, return_home, status, battery, clean_spot
	// An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}


func (c *Vacuum) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsVacuum() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelVacuum {
			ok = true
			break
		}
	}

	return ok
}
