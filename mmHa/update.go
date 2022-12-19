package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


const LabelUpdate = "update"


func (m *Mqtt) UpdatePublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsUpdate() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Update {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelUpdate, m.ClientId, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelUpdate, m.ClientId, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) UpdatePublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsUpdate() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelUpdate, m.ClientId, id, "state")

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

type Update struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic to publish payload_install to start installing process.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this Update is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The type/class of the update to set the icon in the frontend.
	DeviceClass DeviceClass `json:"device_class,omitempty" default:"None"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received and published messages. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// Picture URL for the entity.
	EntityPicture String `json:"entity_picture,omitempty"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as entity attributes. Implies force_update of the current select state when a message is received on this topic.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// Defines a template to extract the latest version value.
	LatestVersionTemplate Template `json:"latest_version_template,omitempty"`

	// The MQTT topic subscribed to receive an update of the latest version.
	LatestVersionTopic String `json:"latest_version_topic,omitempty"`

	// The name of the Select.
	Name String `json:"name,omitempty"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The MQTT payload to start installing process.
	PayloadInstall String `json:"payload_install,omitempty"`

	// The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters.
	ReleaseSummary String `json:"release_summary,omitempty"`

	// URL to the full release notes of the latest version available.
	ReleaseUrl String `json:"release_url,omitempty"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The MQTT topic subscribed to receive state updates. The state update may be either JSON or a simple string with installed_version value. When a JSON payload is detected, the state value of the JSON payload should supply the installed_version and can optional supply: latest_version, title, release_summary, release_url or an entity_picture URL.
	StateTopic String `json:"state_topic,omitempty"`

	// Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed.
	Title String `json:"title,omitempty"`

	// An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template to extract the installed_version state value or to render to a valid JSON payload on from the payload received on state_topic.
	ValueTemplate Template `json:"value_template,omitempty"`
}


func (c *Update) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsUpdate() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelUpdate {
			ok = true
			break
		}
	}

	return ok
}
