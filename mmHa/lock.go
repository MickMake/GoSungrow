package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)

const LabelLock = "lock"


func (m *Mqtt) LockPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsLock() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Lock {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelLock, m.ClientId, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelLock, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) LockPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsLock() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelLock, m.ClientId, id, "state")

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

type Lock struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to change the lock state.
	CommandTopic String `json:"command_topic,omitempty" required:"true"`

	// Information about the device this lock is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
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

	// The name of the lock.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if lock works in optimistic mode.
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: true if no state_topic defined, else false.
	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The payload sent to the lock to lock it.
	PayloadLock String `json:"payload_lock,omitempty" default:"LOCK"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The payload sent to the lock to unlock it.
	PayloadUnlock String `json:"payload_unlock,omitempty" default:"UNLOCK"`

	// The payload sent to the lock to open it.
	PayloadOpen String `json:"payload_open,omitempty" default:"OPEN"`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// If the published message should have the retain flag on or not.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// The payload sent to by the lock when it’s locked.
	StateLocked String `json:"state_locked,omitempty" default:"LOCKED"`

	// The MQTT topic subscribed to receive state updates.
	StateTopic String `json:"state_topic,omitempty"`

	// The payload sent to by the lock when it’s unlocked.
	StateUnlocked String `json:"state_unlocked,omitempty" default:"UNLOCKED"`

	// An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template to extract a value from the payload.
	ValueTemplate String `json:"value_template,omitempty"`
}


func (c *Lock) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsLock() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelLock {
			ok = true
			break
		}
	}

	return ok
}
