package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
	"fmt"
)


func (m *Mqtt) SensorPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsSensor() {
			break
		}

		// if config.LastReset == "" {
		// 	config.LastReset = m.GetLastReset(config.UniqueId)
		// }
		// if config.LastReset != "" {
		// 	if config.LastResetValueTemplate == "" {
		// 		config.LastResetValueTemplate = "{{ value_json.last_reset | as_datetime() }}"
		// 	}
		// 	// LastResetValueTemplate = "{{ value_json.last_reset | int | timestamp_local | as_datetime }}"
		// }

		// switch config.Units {
		// 	case "MW":
		// 		fallthrough
		// 	case "kW":
		// 		fallthrough
		// 	case "W":
		// 		config.DeviceClass = "power"
		//
		// 	case "MWh":
		// 		fallthrough
		// 	case "kWh":
		// 		fallthrough
		// 	case "Wh":
		// 		config.DeviceClass = "energy"
		//
		// 	case "kvar":
		// 		config.DeviceClass = "reactive_power"
		//
		// 	case "Hz":
		// 		config.DeviceClass = "frequency"
		//
		// 	case "V":
		// 		config.DeviceClass = "voltage"
		//
		// 	case "A":
		// 		config.DeviceClass = "current"
		//
		// 	case "℃":
		// 		config.DeviceClass = "temperature"
		// 		// point.Unit = "C"
		//
		// 	case "C":
		// 		config.DeviceClass = "temperature"
		// 		config.Units = "℃"
		//
		// 	case "%":
		// 		config.DeviceClass = "battery"
		//
		// 	default:
		// 		ValueTemplate = "{{ value_json.value }}"
		// }

		// var ok bool
		// var device Device
		// if device, ok = m.Devices[config.ParentName]; !ok {
		// 	break
		// }
		//
		// newDevice := Device{
		// 	ConfigurationUrl: device.ConfigurationUrl,
		// 	Connections:      [][]string {
		// 		{ device.Name, JoinStringsForId(device.Name, config.ParentName) },
		// 		{ JoinStringsForId(device.Name, config.ParentName), JoinStringsForId(device.Name, config.ParentId) },
		// 	},
		// 	Identifiers:      []string{ JoinStringsForId(device.Name, config.ParentId) },
		// 	Manufacturer:     device.Manufacturer,
		// 	Model:            device.Model,
		// 	Name:             JoinStrings(device.Name, config.ParentName),
		// 	SuggestedArea:    device.SuggestedArea,
		// 	SwVersion:        device.SwVersion,
		// 	ViaDevice:        device.ViaDevice,
		// }

		// // device.Name = JoinStrings(m.Device.Name, config.ParentId)
		// device.Name = JoinStrings(m.Device.Name, config.ParentName)	// , config.ValueName)
		// device.Connections = [][]string {
		// 	{ m.Device.Name, JoinStringsForId(m.Device.Name, config.ParentName) },
		// 	{ JoinStringsForId(m.Device.Name, config.ParentName), JoinStringsForId(m.Device.Name, config.ParentId) },
		// 	// { JoinStringsForId(m.Device.Name, config.ParentId), JoinStringsForId(m.Device.Name, config.ParentId, config.Name) },
		// }
		// // device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId, config.Name) }
		// device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId) }

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		// id := JoinStringsForId(m.Device.Name, config.ParentName, config.Name, config.UniqueId),

		payload := Sensor {
			Device:                 newDevice,
			Name:                   JoinStrings(m.DeviceName, config.Name),
			StateTopic:             JoinStringsForTopic(m.sensorPrefix, id, "state"),
			StateClass:             config.StateClass,
			UniqueId:               id,
			UnitOfMeasurement:      config.Units,
			DeviceClass:            config.DeviceClass,
			Qos:                    0,
			ForceUpdate:            true,
			ExpireAfter:            0,
			Encoding:               "utf-8",
			EnabledByDefault:       true,
			LastResetValueTemplate: config.LastResetValueTemplate,
			LastReset:              config.LastReset,
			ValueTemplate:          config.ValueTemplate,
			Icon:                   config.Icon,

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
		}

		tag := JoinStringsForTopic(m.sensorPrefix, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) SensorPublishValue(config EntityConfig) error {

	for range Only.Once {
		// if config.Units == LabelBinarySensor {
		if !config.IsSensor() {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		payload := MqttState {
			LastReset: m.GetLastReset(config.FullId),
			Value:     config.Value,
		}
		tag := JoinStringsForTopic(m.sensorPrefix, id, "state")
		t := m.client.Publish(tag, 0, true, payload.Json())
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
			cs[oid.StateClass][oid.ValueName] = oid.Value
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
	st = JoinStringsForTopic(m.sensorPrefix, st, "state")		// m.GetSensorStateTopic(name, config.SubName),m.EntityPrefix, m.Device.FullName, config.SubName
	return st
}
