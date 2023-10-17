package cmdHassio

import (
	"encoding/json"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
)

const LabelBinarySensor = "binary_sensor"

func (m *Mqtt) BinarySensorPublishConfig(config EntityConfig) error {
	for range only.Once {
		if !config.IsBinarySensor() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		// StateClass := config.Point.UpdateFreq
		// switch {
		// 	case config.Point.IsTotal():
		// 		StateClass = "total"
		// 	default:
		// 		StateClass = "measurement"
		// }
		// ValueName := config.Point.Description
		// UniqueId := config.Point.Id
		// Units := config.Point.Unit
		// LastReset := config.Point.WhenReset()
		//
		// re := mmHa.EntityConfig {
		// 	Name:        name,	// mmHa.JoinStringsForName(" - ", id), // r.Point.Name, // PointName,
		// 	SubName:     "",
		// 	ParentId:    r.EndPoint,
		// 	ParentName:  r.Parent.Key,
		// 	FullId: id, // string(r.FullId),	// WAS r.Point.FullId
		// 	DeviceClass: "",
		// }

		payload := BinarySensor{
			Device:     newDevice,
			Name:       String(JoinStrings(m.DeviceName, config.Name)),
			StateTopic: String(JoinStringsForTopic(m.Prefix, LabelBinarySensor, m.ClientId, id, "state")),
			// StateClass:             config.StateClass,
			ObjectId: String(id),
			UniqueId: String(id),
			// UnitOfMeasurement:      config.Units,
			DeviceClass:      DeviceClass(config.DeviceClass),
			Qos:              0,
			ForceUpdate:      true,
			ExpireAfter:      0,
			Encoding:         "utf-8",
			EnabledByDefault: true,
			PayloadOn:        "true",
			PayloadOff:       "false",
			// LastResetValueTemplate: config.LastResetValueTemplate,
			ValueTemplate: String(config.ValueTemplate),
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
			// EntityCategory:         "",
			// JsonAttributesTemplate: "",
			// JsonAttributesTopic:    "",
			// ObjectId:               "",
			// PayloadAvailable:       "",
			// PayloadNotAvailable:    "",
			// OffDelay:               0,
		}

		tag := JoinStringsForTopic(m.Prefix, LabelBinarySensor, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) BinarySensorPublishValue(config EntityConfig) error {
	for range only.Once {
		if !config.IsBinarySensor() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelBinarySensor, m.ClientId, id, "state")

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
		if config.Units == LabelBinarySensor {
			payload = MqttState{
				Value: value,
			}
		}

		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

type BinarySensor struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive birth and LWT messages from the MQTT device. If availability is not defined, the binary sensor will always be considered available and its state will be on, off or unknown. If availability is defined, the binary sensor will be considered as unavailable by default and the sensor’s initial state will be unavailable. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// Information about the device this binary sensor is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Sets the class of the device, changing the device state and icon that is displayed on the frontend.
	DeviceClass DeviceClass `json:"device_class,omitempty"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" default:"true"`

	// The encoding of the payloads received. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" default:"None"`

	// If set, it defines the number of seconds after the sensor’s state expires, if it’s not updated. After expiry, the sensor’s state becomes unavailable. Default the sensors state never expires.
	ExpireAfter Integer `json:"expire_after,omitempty"`

	// Sends update events (which results in update of state object’s last_changed) even if the sensor’s state hasn’t changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on every incoming state message (not only when the sensor’s new state is different to the current one).
	ForceUpdate Boolean `json:"force_update,omitempty" default:"false"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in MQTT sensor documentation.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// The name of the binary sensor.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// For sensors that only send on state updates (like PIRs), this variable sets a delay in seconds after which the sensor’s state will be updated back to off.
	OffDelay Integer `json:"off_delay,omitempty"`

	// The string that represents the online state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The string that represents the offline state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The string that represents the off state. It will be compared to the message in the state_topic (see value_template for details)
	PayloadOff String `json:"payload_off,omitempty" default:"OFF"`

	// The string that represents the on state. It will be compared to the message in the state_topic (see value_template for details)
	PayloadOn String `json:"payload_on,omitempty" default:"ON"`

	// The maximum QoS level to be used when receiving messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// The MQTT topic subscribed to receive sensor’s state.
	StateTopic String `json:"state_topic,omitempty" required:"true"`

	// An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template that returns a string to be compared to payload_on/payload_off or an empty string, in which case the MQTT message will be removed. Remove this option when payload_on and payload_off are sufficient to match your payloads (i.e no pre-processing of original message is required).
	ValueTemplate String `json:"value_template,omitempty"`
}

// type BinarySensor struct {
// 	Availability           *Availability `json:"availability,omitempty"`
//
// 	AvailabilityTopic      string       `json:"availability_topic,omitempty"`
//
// 	AvailabilityMode       string       `json:"availability_mode,omitempty"`
//
// 	AvailabilityTemplate   string       `json:"availability_template,omitempty"`
//
// 	Device                 Device       `json:"device,omitempty"`
//
// 	DeviceClass            string       `json:"device_class,omitempty"`
//
// 	EnabledByDefault       bool         `json:"enabled_by_default,omitempty"`
//
// 	Encoding               string       `json:"encoding,omitempty"`
//
// 	EntityCategory         string       `json:"entity_category,omitempty"`
//
// 	ExpireAfter            int          `json:"expire_after,omitempty"`
//
// 	ForceUpdate            bool         `json:"force_update,omitempty"`
//
// 	Icon                   string `json:"icon,omitempty"`
//
// 	JsonAttributesTemplate string       `json:"json_attributes_template,omitempty"`
//
// 	JsonAttributesTopic    string       `json:"json_attributes_topic,omitempty"`
//
// 	LastResetValueTemplate string       `json:"last_reset_value_template,omitempty"`
//
// 	Name                   string `json:"name,omitempty"`
//
// 	ObjectId               string `json:"object_id,omitempty"`
//
// 	PayloadAvailable       string       `json:"payload_available,omitempty"`
//
// 	PayloadNotAvailable    string       `json:"payload_not_available,omitempty"`
//
// 	Qos                    int    `json:"qos,omitempty"`
//
// 	Retain                 bool   `json:"ret,omitempty"`
//
// 	StateClass             string       `json:"state_class,omitempty"`
//
// 	StateTopic             string `json:"state_topic,omitempty" required:"true"`
//
// 	UniqueId               string `json:"unique_id,omitempty"`
//
// 	UnitOfMeasurement      string       `json:"unit_of_measurement,omitempty"`
//
// 	Topic                  string       `json:"topic,omitempty" required:"true"`
//
// 	ValueTemplate          string       `json:"value_template,omitempty"`
//
// 	OffDelay               int          `json:"off_delay,omitempty"`
//
// 	PayloadOff             string       `json:"pl_off,omitempty"`
//
// 	PayloadOn              string       `json:"pl_on,omitempty"`
// }

func (c *BinarySensor) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsBinarySensor() bool {
	var ok bool

	for range only.Once {
		if config.Value.IsBool() {
			ok = true
			break
		}
		if config.Point.IsBool() {
			ok = true
			break
		}
		if config.Units == LabelBinarySensor {
			ok = true
			break
		}
		if config.Units == "binary" {
			ok = true
			break
		}
		if config.Units == "Bool" {
			ok = true
			break
		}
		if config.Units == "valueTypes.Bool" {
			ok = true
			break
		}
	}

	return ok
}
