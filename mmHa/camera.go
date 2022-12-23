package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelCamera = "camera"


func (m *Mqtt) CameraPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsCamera() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Camera {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			// CommandTopic: String(JoinStringsForTopic(m.switchPrefix, id, "cmd")),
			ObjectId:     String(id),
			UniqueId:     String(id),
			// Qos:          0,
			// Retain:       true,

			// PayloadOn:     "true",
			// PayloadOff:    "false",
			// StateOn:       "true",
			// StateOff:      "false",
			// ValueTemplate: config.ValueTemplate,
			Icon:          Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelCamera, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) CameraPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsCamera() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelCamera, m.ClientId, id, "state")

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

type Camera struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// Information about the device this camera is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received. Set to "" to disable decoding of incoming payload. Use image_encoding to enable Base64 decoding on topic.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// The encoding of the image payloads received. Set to "b64" to enable base64 decoding of image payload. If not set, the image payload must be raw binary data.
	ImageEncoding String `json:"image_encoding,omitempty" default:"None"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies force_update of the current sensor state when a message is received on this topic.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The name of the camera.
	Name String `json:"name,omitempty"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The MQTT topic to subscribe to.
	Topic String `json:"topic,omitempty" required:"true"`

	// An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`
}


func (c *Camera) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsCamera() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelCamera {
			ok = true
			break
		}
	}

	return ok
}
