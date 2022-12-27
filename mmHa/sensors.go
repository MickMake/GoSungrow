package mmHa

import (
	"encoding/json"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


const LabelSensor = "sensor"

func (m *Mqtt) SensorPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsSensor() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Sensor {
			Device:                 newDevice,
			Name:                   String(JoinStrings(m.DeviceName, config.Name)),
			StateTopic:             String(JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, id, "state")),
			StateClass:             String(config.StateClass),
			ObjectId:               String(id),
			UniqueId:               String(id),
			UnitOfMeasurement:      String(config.Units),
			DeviceClass:            DeviceClass(config.DeviceClass),
			Qos:                    0,
			ForceUpdate:            true,
			ExpireAfter:            0,
			Encoding:               "utf-8",
			EnabledByDefault:       true,
			LastResetValueTemplate: String(config.LastResetValueTemplate),
			// LastReset:              config.LastReset,
			ValueTemplate:          Template(config.ValueTemplate),
			Icon:                   Icon(config.Icon),
		}

		tag := JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) SensorPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsSensor() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, id, "state")

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
		if config.StateClass != "total" {
			payload = MqttState {
				Value:     value,
			}
		}

		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}


type Sensor struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability           *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" Default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’ availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive availability (online/offline) updates.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// Information about the device this sensor is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// The type/class of the sensor to set the icon in the frontend.
	DeviceClass DeviceClass `json:"device_class,omitempty" Default:"None"`

	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault Boolean `json:"enabled_by_default,omitempty" Default:"true"`

	// The encoding of the payloads received. Set to "" to disable decoding of incoming payload.
	Encoding String `json:"encoding,omitempty" Default:"utf"`

	// The category of the entity.
	EntityCategory String `json:"entity_category,omitempty" Default:"None"`

	// If set, it defines the number of seconds after the sensor’s state expires, if it’s not updated. After expiry, the sensor’s state becomes unavailable. Default the sensors state never expires.
	ExpireAfter Integer `json:"expire_after,omitempty" Default:"0"`

	// Sends update events even if the value hasn’t changed. Useful if you want to have meaningful value graphs in history.
	ForceUpdate Boolean `json:"force_update,omitempty" Default:"false"`

	// Icon for the entity.
	Icon Icon `json:"icon,omitempty"`

	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic.
	JsonAttributesTemplate Template `json:"json_attributes_template,omitempty"`

	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies force_update of the current sensor state when a message is received on this topic.
	JsonAttributesTopic String `json:"json_attributes_topic,omitempty"`

	// Defines a template to extract the last_reset. Available variables: entity_id. The entity_id can be used to reference the entity’s attributes.
	LastResetValueTemplate String `json:"last_reset_value_template,omitempty"`

	// The name of the MQTT sensor.
	Name String `json:"name,omitempty" Default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// The payload that represents the available state.
	PayloadAvailable String `json:"payload_available,omitempty" Default:"online"`

	// The payload that represents the unavailable state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" Default:"offline"`

	// The maximum QoS level of the state topic.
	Qos Integer `json:"qos,omitempty" Default:"0"`

	// The state_class of the sensor.
	StateClass String `json:"state_class,omitempty" Default:"None"`

	// The MQTT topic subscribed to receive sensor values.
	StateTopic String `json:"state_topic,omitempty" required:"true"`

	// An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines the units of measurement of the sensor, if any.
	UnitOfMeasurement String `json:"unit_of_measurement,omitempty"`

	// Defines a template to extract the value. If the template throws an error, the current state will be used instead.
	ValueTemplate Template `json:"value_template,omitempty"`
}

// type Sensor struct {
// 	// availability list (optional)
// 	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
// 	Availability           *Availability `json:"availability,omitempty" required:"false"`
//
// 	// availability_template template (optional)
// 	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
// 	AvailabilityTemplate   string       `json:"availability_template,omitempty" required:"false"`
//
// 	// availability_topic string (optional)
// 	// The MQTT topic subscribed to receive availability (online/offline) updates.
// 	AvailabilityTopic      string       `json:"availability_topic,omitempty" required:"false"`
//
// 	// availability_mode string (optional, default: latest)
// 	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
// 	AvailabilityMode       string       `json:"availability_mode,omitempty" required:"false"`
//
// 	// device map (optional)
// 	// Information about the device this sensor is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
// 	Device                 Device       `json:"device,omitempty" required:"false"`
//
// 	// device_class device_class (optional, default: None)
// 	// The type/class of the sensor to set the icon in the frontend.
// 	DeviceClass            string       `json:"device_class,omitempty" required:"false"`
//
// 	// enabled_by_default boolean (optional, default: true)
// 	// Flag which defines if the entity should be enabled when first added.
// 	EnabledByDefault       bool         `json:"enabled_by_default,omitempty" required:"false"`
//
// 	// encoding string (optional, default: utf-8)
// 	// The encoding of the payloads received. Set to "" to disable decoding of incoming payload.
// 	Encoding               string       `json:"encoding,omitempty" required:"false"`
//
// 	// entity_category string (optional, default: None)
// 	// The category of the entity.
// 	EntityCategory         string       `json:"entity_category,omitempty" required:"false"`
//
// 	// expire_after integer (optional, default: 0)
// 	// If set, it defines the number of seconds after the sensor’s state expires, if it’s not updated. After expiry, the sensor’s state becomes unavailable. Default the sensors state never expires.
// 	ExpireAfter            int          `json:"expire_after,omitempty" required:"false"`
//
// 	// force_update boolean (optional, default: false)
// 	// Sends update events even if the value hasn’t changed. Useful if you want to have meaningful value graphs in history.
// 	ForceUpdate            bool         `json:"force_update,omitempty" required:"false"`
//
// 	// icon string (optional)
// 	// Icon for the entity.
// 	Icon                   string `json:"icon,omitempty" required:"false"`
//
// 	// json_attributes_template template (optional)
// 	// Defines a template to extract the JSON dictionary from messages received on the json_attributes_topic.
// 	JsonAttributesTemplate string       `json:"json_attributes_template,omitempty" required:"false"`
//
// 	// json_attributes_topic string (optional)
// 	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies force_update of the current sensor state when a message is received on this topic.
// 	JsonAttributesTopic    string       `json:"json_attributes_topic,omitempty" required:"false"`
//
// 	// last_reset_value_template string (optional)
// 	// Defines a template to extract the last_reset. Available variables: entity_id. The entity_id can be used to reference the entity’s attributes.
// 	LastResetValueTemplate string       `json:"last_reset_value_template,omitempty" required:"false"`
//
// 	// name string (optional, default: MQTT Sensor)
// 	// The name of the MQTT sensor.
// 	Name                   string `json:"name,omitempty" required:"false"`
//
// 	// object_id string (optional)
// 	// Used instead of name for automatic generation of entity_id
// 	ObjectId               string `json:"object_id,omitempty" required:"false"`
//
// 	// payload_available string (optional, default: online)
// 	// The payload that represents the available state.
// 	PayloadAvailable       string       `json:"payload_available,omitempty" required:"false"`
//
// 	// payload_not_available string (optional, default: offline)
// 	// The payload that represents the unavailable state.
// 	PayloadNotAvailable    string       `json:"payload_not_available,omitempty" required:"false"`
//
// 	// qos integer (optional, default: 0)
// 	// The maximum QoS level of the state topic.
// 	Qos                    int    `json:"qos,omitempty" required:"false"`
//
// 	// retain boolean (optional, default: false)
// 	// If the published message should have the retain flag on or not.
// 	Retain                 bool   `json:"retain,omitempty" required:"false"`
//
// 	// state_class string (optional, default: None)
// 	// The state_class of the sensor.
// 	StateClass             string       `json:"state_class,omitempty" required:"false"`
//
// 	// state_topic string REQUIRED
// 	// The MQTT topic subscribed to receive sensor values.
// 	StateTopic             string `json:"state_topic,omitempty" required:"false"`
//
// 	// unique_id string (optional)
// 	// An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception.
// 	UniqueId               string `json:"unique_id,omitempty" required:"false"`
//
// 	// unit_of_measurement string (optional)
// 	// Defines the units of measurement of the sensor, if any.
// 	UnitOfMeasurement      string       `json:"unit_of_measurement,omitempty" required:"false"`
//
// 	// value_template template (optional)
// 	// Defines a template to extract the value. If the template throws an error, the current state will be used instead.
// 	ValueTemplate          string `json:"value_template,omitempty" required:"false"`
//
// 	LastReset              string       `json:"last_reset,omitempty" required:"false"`
//
// 	// StateFunc func() string `json:"-"`
// 	//
// 	// UpdateInterval  float64 `json:"-"`
// 	// ForceUpdateMQTT bool    `json:"-"`
// }

func (c *Sensor) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsSensor() bool {
	var ok bool

	for range Only.Once {
		if config.IsBinarySensor() {
			break
		}
		if config.IsSwitch() {
			break
		}
		if config.IsLight() {
			break
		}
		if config.IsSelect() {
			break
		}

		ok = true
	}

	return ok
}


type Fields map[string]string

func (m *Mqtt) PublishSensorValues(configs []EntityConfig) error {
	for range Only.Once {
		cs := make(map[string]Fields)
		topic := ""
		for _, oid := range configs {
			if topic == "" {
				topic = JoinStringsForId(m.DeviceName, oid.ParentName, oid.Name)
			}
			if _, ok := cs[oid.StateClass]; !ok {
				cs[oid.StateClass] = make(Fields)
			}
			cs[oid.StateClass][oid.ValueName] = oid.Value.String()
		}

		for n, c := range cs {
			j, _ := json.Marshal(c)
			fmt.Printf("%s (%s) -> %s\n", topic, n, string(j))

			m.err = m.PublishValue(n, topic, string(j))
		}
	}
	return m.err
}

func (m *Mqtt) GetSensorStateTopic(config EntityConfig) string {
	st := JoinStringsForId(m.DeviceName, config.FullId)
	st = JoinStringsForTopic(m.Prefix, LabelSensor, m.ClientId, st, "state")		// m.GetSensorStateTopic(name, config.SubName),m.EntityPrefix, m.Device.FullName, config.SubName
	return st
}
