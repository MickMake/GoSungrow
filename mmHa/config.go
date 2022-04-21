package mmHa

import "encoding/json"


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
