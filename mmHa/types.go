package mmHa


// Device - device map (optional)
// Information about the device this Select is a part of to tie it into the device registry.
// Only works through MQTT discovery and when unique_id is set.
// At least one of identifiers or connections must be present to identify the device.
type Device struct {
	// configuration_url string (optional)
	// A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link.
	ConfigurationUrl string     `json:"configuration_url,omitempty" required:"false"`

	// hw_version string (optional)
	// The hardware version of the device.
	HwVersion        string     `json:"hw_version,omitempty" required:"false"`

	// connections list (optional)
	// A list of connections of the device to the outside world as a list of tuples [connection_type, connection_identifier].
	// For example the MAC address of a network interface: "connections": ["mac", "02:5b:26:a8:dc:12"].
	Connections      [][]string `json:"connections,omitempty" required:"false"`

	// identifiers list | string (optional)
	// A list of IDs that uniquely identify the device. For example a serial number.
	Identifiers      []string   `json:"identifiers,omitempty" required:"false"`

	// manufacturer string (optional)
	// The manufacturer of the device.
	Manufacturer     string     `json:"manufacturer,omitempty" required:"false"`

	// model string (optional)
	// The model of the device.
	Model            string     `json:"model,omitempty" required:"false"`

	// name string (optional)
	// The name of the device.
	Name             string     `json:"name,omitempty" required:"false"`

	// suggested_area string (optional)
	// Suggest an area if the device isn’t in one yet.
	SuggestedArea    string     `json:"suggested_area,omitempty" required:"false"`

	// sw_version string (optional)
	// The firmware version of the device.
	SwVersion        string     `json:"sw_version,omitempty" required:"false"`

	// via_device string (optional)
	// Identifier of a device that routes messages between this device and Home Assistant.
	// Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant.
	ViaDevice        string     `json:"via_device,omitempty" required:"false"`
}


// Availability - availability list (optional)
// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
type Availability struct {
	// payload_available string (optional, default: online)
	// The payload that represents the available state.
	PayloadAvailable    string `json:"payload_available,omitempty" required:"false"`

	// payload_not_available string (optional, default: offline)
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty" required:"false"`

	// topic string REQUIRED
	// An MQTT topic subscribed to receive availability (online/offline) updates.
	Topic               string `json:"topic,omitempty" required:"true"`

	// value_template template (optional)
	// Defines a template to extract device’s availability from the topic.
	// To determine the devices’ availability result of this template will be compared to payload_available and payload_not_available.
	ValueTemplate       string `json:"value_template,omitempty" required:"false"`
}

type List []string
type String string
type Template string
type Boolean bool
type Integer int64
type Icon string
type DeviceClass string
type Float float64
