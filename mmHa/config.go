package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
)


type Config struct {
	Entry        string       `json:"~,omitempty" required:"false"`
	Name         string       `json:"name,omitempty" required:"false"`
	UniqueId     string       `json:"unique_id,omitempty" required:"false"`
	StateTopic   string       `json:"state_topic,omitempty" required:"true"`
	DeviceConfig DeviceConfig `json:"device,omitempty" required:"false"`
}
type DeviceConfig struct {
	Identifiers  []string `json:"identifiers,omitempty" required:"false"`
	SwVersion    string   `json:"sw_version,omitempty" required:"false"`
	Name         string   `json:"name,omitempty" required:"false"`
	Manufacturer string   `json:"manufacturer,omitempty" required:"false"`
	Model        string   `json:"model,omitempty" required:"false"`
}

func (c *Config) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}


type Device struct {
	ConfigurationUrl string     `json:"configuration_url,omitempty" required:"false"`
	Connections      [][]string `json:"connections,omitempty" required:"false"`
	Identifiers      []string   `json:"identifiers,omitempty" required:"false"`
	Manufacturer     string     `json:"manufacturer,omitempty" required:"false"`
	Model            string     `json:"model,omitempty" required:"false"`
	Name             string     `json:"name,omitempty" required:"false"`
	SuggestedArea    string     `json:"suggested_area,omitempty" required:"false"`
	SwVersion        string     `json:"sw_version,omitempty" required:"false"`
	ViaDevice        string     `json:"via_device,omitempty" required:"false"`
}

func (m *Mqtt) NewDevice(config EntityConfig) (bool, Device) {
	var ok bool
	var ret Device

	for range Only.Once {
		var parent Device
		if parent, ok = m.MqttDevices[config.ParentName]; !ok {
			break
		}

		manu := parent.Manufacturer
		if manu == "" {
			manu = m.DeviceName
		}
		modl := parent.Model
		if modl == "" {
			modl = m.DeviceName
		}

		ret = Device {
			ConfigurationUrl: parent.ConfigurationUrl,
			Connections:      [][]string {
				{ m.EntityPrefix, JoinStringsForId(m.EntityPrefix, config.ParentName) },
				{ JoinStringsForId(m.EntityPrefix, config.ParentName), JoinStringsForId(m.EntityPrefix, config.ParentId) },
			},
			Identifiers:      []string{ JoinStringsForId(m.EntityPrefix, config.ParentId) },
			Manufacturer:     manu,
			Model:            modl,
			Name:             JoinStrings(m.EntityPrefix, config.ParentName, "-", parent.Name),
			SuggestedArea:    parent.SuggestedArea,
			SwVersion:        parent.SwVersion,
			ViaDevice:        parent.ViaDevice,
		}
		ok = true
	}

	return ok, ret

}

// func (d *Device) NewDevice(config EntityConfig) Device {
// 	var ret Device
//
// 	for range Only.Once {
// 		ret = Device{
// 			ConfigurationUrl: d.ConfigurationUrl,
// 			Connections:      [][]string {
// 				{ d.Name, JoinStringsForId(d.Name, config.ParentName) },
// 				{ JoinStringsForId(d.Name, config.ParentName), JoinStringsForId(d.Name, config.ParentId) },
// 			},
// 			Identifiers:      []string{ JoinStringsForId(d.Name, config.ParentId) },
// 			Manufacturer:     d.Manufacturer,
// 			Model:            d.Model,
// 			Name:             JoinStrings(d.Name, config.ParentName),
// 			SuggestedArea:    d.SuggestedArea,
// 			SwVersion:        d.SwVersion,
// 			ViaDevice:        d.ViaDevice,
// 		}
//
// 		// // device.Name = JoinStrings(m.Device.Name, config.ParentId)
// 		// device.Name = JoinStrings(m.Device.Name, config.ParentName)	// , config.ValueName)
// 		// device.Connections = [][]string {
// 		// 	{ m.Device.Name, JoinStringsForId(m.Device.Name, config.ParentName) },
// 		// 	{ JoinStringsForId(m.Device.Name, config.ParentName), JoinStringsForId(m.Device.Name, config.ParentId) },
// 		// 	// { JoinStringsForId(m.Device.Name, config.ParentId), JoinStringsForId(m.Device.Name, config.ParentId, config.Name) },
// 		// }
// 		// // device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId, config.Name) }
// 		// device.Identifiers = []string{ JoinStringsForId(m.Device.Name, config.ParentId) }
// 	}
//
// 	return ret
// }

// {
//	"device": {
//		"identifiers": [
//			"sungrow"
//		],
//		"manufacturer": "MickMake",
//		"model": "GoLang",
//		"name": "sungrow",
//		"sw_version": "sungrow https://github.com/MickMake/GoSungrow"
//	},
//	"name": "sungrow",
//	"stat_t": "~/state",
//	"unique_id": "sungrow",
//	"~": "homeassistant/binary_sensor/SunGrow"
// }
