package cmdHassio

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelTag = "tag"


func (m *Mqtt) TagPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsTag() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Tag {
			Device:       newDevice,
			// Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			// CommandTopic: String(JoinStringsForTopic(m.switchPrefix, id, "cmd")),
			// ObjectId:     String(id),
			// UniqueId:     String(id),
			// Qos:          0,
			// Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			// Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelTag, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) TagPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsTag() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelTag, m.ClientId, id, "state")

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

type Tag struct {
	// The MQTT topic subscribed to receive tag scanned events.
	Topic String `json:"topic,omitempty" required:"true"`

	// Defines a template that returns a tag ID.
	ValueTemplate String `json:"value_template,omitempty"`

	// Information about the device this device trigger is a part of to tie it into the device registry. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty" required:"true"`
}


func (c *Tag) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsTag() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelTag {
			ok = true
			break
		}
	}

	return ok
}
