package mmHa

import (
	"GoSungrow/Only"
	"encoding/json"
	"strings"
)


func (m *Mqtt) BinarySensorPublishConfig(config EntityConfig) error {

	for range Only.Once {
		config.FixConfig()
		if !config.IsBinarySensor() {
			break
		}

		// LastReset := m.GetLastReset(config.UniqueId)
		// LastResetValueTemplate := ""
		// if LastReset != "" {
		// 	LastResetValueTemplate = "{{ value_json.last_reset | as_datetime() }}"
		// }

		// device := m.Device
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

		payload := BinarySensor {
			Device:                 newDevice,
			Name:                   JoinStrings(m.DeviceName, config.Name),
			StateTopic:             JoinStringsForTopic(m.binarySensorPrefix, id, "state"),
			StateClass:             config.StateClass,
			UniqueId:               id,
			UnitOfMeasurement:      config.Units,
			DeviceClass:            config.DeviceClass,
			Qos:                    0,
			ForceUpdate:            true,
			ExpireAfter:            0,
			Encoding:               "utf-8",
			EnabledByDefault:       true,
			PayloadOn:              "true",
			PayloadOff:             "false",
			LastResetValueTemplate: config.LastResetValueTemplate,
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
			// OffDelay:               0,
		}

		tag := JoinStringsForTopic(m.binarySensorPrefix, id, "config")
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}

func (m *Mqtt) BinarySensorPublishValue(config EntityConfig) error {

	for range Only.Once {
		// if config.Units != LabelBinarySensor {
		if !config.IsBinarySensor() {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		// tagId := JoinStringsForId(m.Device.Name, config.ParentName, config.Name, config.UniqueId),
		tag := JoinStringsForTopic(m.binarySensorPrefix, id, "state")

		// @TODO - Real hack here. Need to properly check for JSON.
		if strings.Contains(config.Value, `{`) || strings.Contains(config.Value, `":`) {
			t := m.client.Publish(tag, 0, true, config.Value)
			if !t.WaitTimeout(m.Timeout) {
				m.err = t.Error()
			}
			break
		}

		payload := MqttState {
			LastReset: m.GetLastReset(config.UniqueId),
			Value:     config.Value,
		}
		t := m.client.Publish(tag, 0, true, payload.Json())
		if !t.WaitTimeout(m.Timeout) {
			m.err = t.Error()
		}
	}

	return m.err
}


type BinarySensor struct {
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

	OffDelay               int          `json:"off_delay,omitempty" required:"false"`
	PayloadOff             string       `json:"pl_off,omitempty" required:"false"`
	PayloadOn              string       `json:"pl_on,omitempty" required:"false"`
}
func (c *BinarySensor) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}
