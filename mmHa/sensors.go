package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
	"fmt"
)


// func (m *Mqtt) PublishSensorConfig(config EntityConfig) error {
// 	// func (m *Mqtt) PublishSensorConfig(id string, name string, subName string, units string, valueName string, class string) error {
// 	for range Only.Once {
// 		// name = strings.ReplaceAll(name, "/", ".")
// 		// config.SubName = strings.ReplaceAll(config.SubName, "/", ".")
// 		// id = JoinStringsForId(m.Device.FullName, config.SubName, id)
// 		// st := JoinStringsForTopic(m.sensorPrefix, JoinStringsForId(m.EntityPrefix, m.Device.FullName, strings.ReplaceAll(config.SubName, "/", ".")), "state")
// 		// st := m.GetSensorStateTopic(name, config.SubName)
//
// 		device := m.Device
// 		device.Name = JoinStrings(m.Device.Name, config.ParentId)
// 		device.Connections = [][]string{{m.Device.Name, device.Name}}
// 		device.Identifiers = []string{device.Name}
// 		st := JoinStringsForId(m.Device.Name, config.ParentName, config.Name)
//
// 		payload := Sensor {
// 			Device:                 device,
// 			DeviceClass:            config.Class,
// 			Name:                   JoinStrings(m.Device.Name, config.ParentName, config.Name, config.SubName), // config.FullName, // JoinStrings(m.Device.ViaDevice, config.SubName),
// 			StateTopic:             JoinStringsForTopic(m.sensorPrefix, st, "state"), // m.GetSensorStateTopic(name, config.SubName),m.EntityPrefix, m.Device.FullName, config.SubName
// 			StateClass:             "measurement",
// 			UniqueId:               JoinStringsForId(m.Device.Name, config.ParentName, config.Name, config.UniqueId),
// 			UnitOfMeasurement:      config.Units,
// 			Qos:                    0,
// 			ForceUpdate:            true,
// 			ExpireAfter:            0,
// 			Encoding:               "utf-8",
// 			EnabledByDefault:       true,
// 			// LastResetValueTemplate: LastResetValueTemplate,
// 			// LastReset:              LastReset,
// 			// "{%if is_state(value_json.ifadminstatus,\"up\")-%}ON{%-else-%}OFF{%-endif%}"
// 			ValueTemplate:          fmt.Sprintf("{{ value_json.%s | is_defined }}", config.ValueName),
// 			// LastReset: time.Now().Format("2006-01-02T00:00:00+00:00"),
// 			// LastResetValueTemplate: "{{entity_id}}",
// 			// LastResetValueTemplate: "{{ (as_datetime((value_json.last_reset | int | timestamp_utc)|string+'Z')).isoformat() }}",
// 		}
//
// 		ct := JoinStringsForId(m.Device.Name, config.ParentName, config.Name, config.SubName)
// 		m.client.Publish(JoinStringsForTopic(m.sensorPrefix, ct, "config"), 0, true, payload.Json())
// 		// m.client.Publish(JoinStringsForTopic(m.sensorPrefix, config.UniqueId, "config"), 0, true, payload.Json())
// 	}
//
// 	return m.err
// }

type Fields map[string]string
func (m *Mqtt) PublishSensorValues(configs []EntityConfig) error {
	for range Only.Once {
		cs := make(map[string]Fields)
		topic := ""
		for _, oid := range configs {
			if topic == "" {
				topic = JoinStringsForId(m.Device.Name, oid.ParentName, oid.Name)
			}
			if _, ok := cs[oid.Type]; !ok {
				cs[oid.Type] = make(Fields)
			}
			cs[oid.Type][oid.ValueName] = oid.Value
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
	st := JoinStringsForId(m.Device.Name, config.ParentName, config.Name)
	st = JoinStringsForTopic(m.sensorPrefix, st, "state")		// m.GetSensorStateTopic(name, config.SubName),m.EntityPrefix, m.Device.FullName, config.SubName
	return st
}

func (m *Mqtt) SensorPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if config.Units == "binary" {
			break
		}

		ValueTemplate := "{{ value_json.value | float }}"
		// ValueTemplate := fmt.Sprintf("{{ value_json.%s | float }}", config.ValueName)	// <- Used with merged values.

		LastReset := m.GetLastReset(config.UniqueId)
		LastResetValueTemplate := ""
		if LastReset != "" {
			LastResetValueTemplate = "{{ value_json.last_reset | as_datetime() }}"
			// LastResetValueTemplate = "{{ value_json.last_reset | int | timestamp_local | as_datetime }}"
		}

		switch config.Units {
			case "MW":
				fallthrough
			case "kW":
				fallthrough
			case "W":
				config.Class = "power"
	
			case "MWh":
				fallthrough
			case "kWh":
				fallthrough
			case "Wh":
				config.Class = "energy"
	
			case "kvar":
				config.Class = "reactive_power"
	
			case "Hz":
				config.Class = "frequency"
	
			case "V":
				config.Class = "voltage"
	
			case "A":
				config.Class = "current"
	
			case "℃":
				config.Class = "temperature"
				// point.Unit = "C"
	
			case "C":
				config.Class = "temperature"
				config.Units = "℃"
	
			case "%":
				config.Class = "battery"

			default:
				ValueTemplate = "{{ value_json.value }}"
		}

		device := m.Device
		device.Name = JoinStrings(m.Device.Name, config.ParentId)
		device.Connections = [][]string {
			{ m.Device.Name, JoinStringsForId(m.Device.Name, config.ParentId) },
			{ JoinStringsForId(m.Device.Name, config.ParentId), JoinStringsForId(m.Device.Name, config.ParentId, config.Name) },
		}
		device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId) }
		st := JoinStringsForId(m.Device.Name, config.ParentId, config.Name)	// , config.ParentName, config.Name)
		// UniqueId:               JoinStringsForId(m.Device.Name, config.ParentName, config.Name, config.UniqueId),

		payload := Sensor {
			Device:                 device,
			Name:                   JoinStrings(m.Device.Name, config.ParentName, config.FullName),
			StateTopic:             JoinStringsForTopic(m.sensorPrefix, st, "state"),
			// m.GetSensorStateTopic(name, config.SubName),m.EntityPrefix, m.Device.FullName, config.SubName
			StateClass:             "measurement",
			UniqueId:               st,
			UnitOfMeasurement:      config.Units,
			DeviceClass:            config.Class,
			Qos:                    0,
			ForceUpdate:            true,
			ExpireAfter:            0,
			Encoding:               "utf-8",
			EnabledByDefault:       true,
			LastResetValueTemplate: LastResetValueTemplate,
			LastReset:              LastReset,
			ValueTemplate:          ValueTemplate,
		}

		ct := JoinStringsForTopic(m.sensorPrefix, st, "config")
		t := m.client.Publish(ct, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}


// func (m *Mqtt) PublishSensor(subtopic string, payload interface{}) error {
// 	for range Only.Once {
// 		t := m.client.Publish(subtopic, 0, true, payload)
// 		if !t.WaitTimeout(m.Timeout) {
// 			m.err = t.Error()
// 		}
// 	}
// 	return m.err
// }
// func (m *Mqtt) SensorPublish(subtopic string, payload interface{}) error {
// 	for range Only.Once {
// 		m.client.Publish(SensorBaseTopic + "/" + subtopic, 0, true, payload)
// 	}
// 	return m.err
// }
//
// func (m *Mqtt) PublishSensorState(topic string, payload interface{}) error {
// 	for range Only.Once {
// 		// topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
// 		// topic = JoinStringsForTopic(m.sensorPrefix, topic, "state")
// 		// st := JoinStringsForTopic(m.sensorPrefix, JoinStringsForId(m.EntityPrefix, m.Device.FullName, strings.ReplaceAll(subName, "/", ".")), "state")
// 		t := m.client.Publish(topic, 0, true, payload)
// 		if !t.WaitTimeout(m.Timeout) {
// 			m.err = t.Error()
// 		}
// 	}
// 	return m.err
// }
// func (m *Mqtt) SensorPublishState(id string, payload interface{}) error {
// 	for range Only.Once {
// 		id = strings.ReplaceAll("sungrow_" + id, ".", "-")
// 		m.client.Publish(SensorBaseTopic + "/" + id + "/state", 0, true, payload)
// 	}
// 	return m.err
// }
//
// func (m *Mqtt) PublishSensorValue(topic string, value string) error {
// 	for range Only.Once {
// 		topic = JoinStringsForId(m.EntityPrefix, m.Device.Name, topic)
// 		// st := JoinStringsForTopic(m.sensorPrefix, JoinStringsForId(m.EntityPrefix, m.Device.FullName, strings.ReplaceAll(subName, "/", ".")), "state")
// 		payload := MqttState {
// 			LastReset: "", // m.GetLastReset(point.PointId),
// 			Value:     value,
// 		}
// 		m.client.Publish(JoinStringsForTopic(m.sensorPrefix, topic, "state"), 0, true, payload.Json())
// 	}
// 	return m.err
// }

func (m *Mqtt) SensorPublishValue(config EntityConfig) error {
	// func (m *Mqtt) SensorPublishValue(point api.DataEntry) error {

	for range Only.Once {
		if config.Units == "binary" {
			break
		}

		st := JoinStringsForId(m.Device.Name, config.ParentId, config.Name)
		payload := MqttState {
			LastReset: m.GetLastReset(config.UniqueId),
			Value:     config.Value,
		}
		st = JoinStringsForTopic(m.sensorPrefix, st, "state")
		t := m.client.Publish(st, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}
	return m.err
}

type Sensor struct {
	Availability           *Availability `json:"availability,omitempty" required:"false"`
	AvailabilityMode       string       `json:"availability_mode,omitempty" required:"false"`
	AvailabilityTemplate   string       `json:"availability_template,omitempty" required:"false"`
	AvailabilityTopic      string       `json:"availability_topic,omitempty" required:"false"`
	Device                 Device       `json:"device,omitempty" required:"false"`
	DeviceClass            string       `json:"device_class,omitempty" required:"false"`
	EnabledByDefault       bool         `json:"enabled_by_default,omitempty" required:"false"`
	Encoding               string       `json:"encoding,omitempty" required:"false"`
	EntityCategory         string       `json:"entity_category,omitempty" required:"false"`
	ExpireAfter            int          `json:"expire_after,omitempty" required:"false"`
	ForceUpdate            bool         `json:"force_update,omitempty" required:"false"`
	Icon                   string       `json:"icon,omitempty" required:"false"`
	JsonAttributesTemplate string       `json:"json_attributes_template,omitempty" required:"false"`
	JsonAttributesTopic    string       `json:"json_attributes_topic,omitempty" required:"false"`
	LastResetValueTemplate string       `json:"last_reset_value_template,omitempty" required:"false"`
	Name                   string       `json:"name,omitempty" required:"false"`
	ObjectId               string       `json:"object_id,omitempty" required:"false"`
	PayloadAvailable       string       `json:"payload_available,omitempty" required:"false"`
	PayloadNotAvailable    string       `json:"payload_not_available,omitempty" required:"false"`
	Qos                    int          `json:"qos,omitempty" required:"false"`
	StateClass             string       `json:"state_class,omitempty" required:"false"`
	StateTopic             string       `json:"state_topic" required:"true"`
	UniqueId               string       `json:"unique_id,omitempty" required:"false"`
	UnitOfMeasurement      string       `json:"unit_of_measurement,omitempty" required:"false"`
	ValueTemplate          string       `json:"value_template,omitempty" required:"false"`
	LastReset              string       `json:"last_reset,omitempty" required:"false"`

	// StateFunc func() string `json:"-"`
	//
	// UpdateInterval  float64 `json:"-"`
	// ForceUpdateMQTT bool    `json:"-"`
}

func (c *Sensor) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}
