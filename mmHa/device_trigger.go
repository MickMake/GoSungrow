package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelDeviceTrigger = "device_trigger"


func (m *Mqtt) DeviceTriggerPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsDeviceTrigger() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := DeviceTrigger {
			Device:       newDevice,
			// Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			// CommandTopic: String(JoinStringsForTopic(m.switchPrefix, id, "cmd")),
			// ObjectId:     String(id),
			// UniqueId:     String(id),
			Qos:          0,
			// Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			// Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelDeviceTrigger, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) DeviceTriggerPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsDeviceTrigger() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelDeviceTrigger, m.ClientId, id, "state")

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

type DeviceTrigger struct {
	// The type of automation, must be ‘trigger’.
	AutomationType String `json:"automation_type,omitempty" required:"true"`

	// Optional payload to match the payload being sent over the topic.
	Payload String `json:"payload,omitempty"`

	// The maximum QoS level to be used when receiving messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// The MQTT topic subscribed to receive trigger events.
	Topic String `json:"topic,omitempty" required:"true"`

	// The type of the trigger, e.g. button_short_press. Entries supported by the frontend: button_short_press, button_short_release, button_long_press, button_long_release, button_double_press, button_triple_press, button_quadruple_press, button_quintuple_press. If set to an unsupported value, will render as subtype type, e.g. button_1 spammed with type set to spammed and subtype set to button_1
	Type String `json:"type,omitempty" required:"true"`

	// The subtype of the trigger, e.g. button_1. Entries supported by the frontend: turn_on, turn_off, button_1, button_2, button_3, button_4, button_5, button_6. If set to an unsupported value, will render as subtype type, e.g. left_button pressed with type set to button_short_press and subtype set to left_button
	Subtype String `json:"subtype,omitempty" required:"true"`

	// Information about the device this device trigger is a part of to tie it into the device registry. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty" required:"true"`

	// Defines a template to extract the value.
	ValueTemplate Template `json:"value_template,omitempty"`
}


func (c *DeviceTrigger) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsDeviceTrigger() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelDeviceTrigger {
			ok = true
			break
		}
	}

	return ok
}
