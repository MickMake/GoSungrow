package mmHa

import (
	"encoding/json"
	"github.com/MickMake/GoUnify/Only"
	"strings"
)


const LabelCover = "cover"


func (m *Mqtt) CoverPublishConfig(config EntityConfig) error {

	for range Only.Once {
		if !config.IsCover() {
			break
		}

		ok, newDevice := m.NewDevice(config)
		if !ok {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)

		payload := Cover {
			Device:       newDevice,
			Name:         String(JoinStrings(m.DeviceName, config.Name)),
			// StateTopic:   JoinStringsForTopic(m.switchPrefix, id, "state"),
			CommandTopic: String(JoinStringsForTopic(m.Prefix, LabelCover, m.ClientId, id, "cmd")),
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

		tag := JoinStringsForTopic(m.Prefix, LabelCover, m.ClientId, id, "config")
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

func (m *Mqtt) CoverPublishValue(config EntityConfig) error {

	for range Only.Once {
		if !config.IsCover() {
			break
		}

		if config.IgnoreUpdate {
			break
		}

		id := JoinStringsForId(m.DeviceName, config.FullId)
		tag := JoinStringsForTopic(m.Prefix, LabelCover, m.ClientId, id, "state")

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
			LastReset: config.LastReset, // m.GetLastReset(config.FullId),
			Value:     value,
		}
		m.err = m.Publish(tag, 0, true, payload.Json())
	}

	return m.err
}

type Cover struct {
	// A list of MQTT topics subscribed to receive availability (online/offline) updates. Must not be used together with availability_topic.
	Availability *Availability `json:"availability,omitempty"`

	// When availability is configured, this controls the conditions needed to set the entity to available. Valid entries are all, any, and latest. If set to all, payload_available must be received on all configured availability topics before the entity is marked as online. If set to any, payload_available must be received on at least one configured availability topic before the entity is marked as online. If set to latest, the last payload_available or payload_not_available received on any configured availability topic controls the availability.
	AvailabilityMode String `json:"availability_mode,omitempty" default:"latest"`

	// Defines a template to extract device’s availability from the availability_topic. To determine the devices’s availability result of this template will be compared to payload_available and payload_not_available.
	AvailabilityTemplate Template `json:"availability_template,omitempty"`

	// The MQTT topic subscribed to receive birth and LWT messages from the MQTT cover device. If an availability topic is not defined, the cover availability state will always be available. If an availability topic is defined, the cover availability state will be unavailable by default. Must not be used together with availability.
	AvailabilityTopic String `json:"availability_topic,omitempty"`

	// The MQTT topic to publish commands to control the cover.
	CommandTopic String `json:"command_topic,omitempty"`

	// Information about the device this cover is a part of to tie it into the device registry. Only works through MQTT discovery and when unique_id is set. At least one of identifiers or connections must be present to identify the device.
	Device Device `json:"device,omitempty"`

	// Sets the class of the device, changing the device state and icon that is displayed on the frontend.
	DeviceClass String `json:"device_class,omitempty"`

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

	// The name of the cover.
	Name String `json:"name,omitempty" default:"MQTT"`

	// Used instead of name for automatic generation of entity_id
	ObjectId String `json:"object_id,omitempty"`

	// Flag that defines if switch works in optimistic mode.
	Optimistic Boolean `json:"optimistic,omitempty"`

	// Default: false if state or position topic defined, else true.
	// The payload that represents the online state.
	PayloadAvailable String `json:"payload_available,omitempty" default:"online"`

	// The command payload that closes the cover.
	PayloadClose String `json:"payload_close,omitempty" default:"CLOSE"`

	// The payload that represents the offline state.
	PayloadNotAvailable String `json:"payload_not_available,omitempty" default:"offline"`

	// The command payload that opens the cover.
	PayloadOpen String `json:"payload_open,omitempty" default:"OPEN"`

	// The command payload that stops the cover.
	PayloadStop String `json:"payload_stop,omitempty" default:"STOP"`

	// Number which represents closed position.
	PositionClosed Integer `json:"position_closed,omitempty" default:"0"`

	// Number which represents open position.
	PositionOpen Integer `json:"position_open,omitempty" default:"100"`

	// Defines a template that can be used to extract the payload for the position_topic topic. Within the template the following variables are available: entity_id, position_open; position_closed; tilt_min; tilt_max. The entity_id can be used to reference the entity’s attributes with help of the states template function;
	PositionTemplate String `json:"position_template,omitempty"`

	// The MQTT topic subscribed to receive cover position messages.
	PositionTopic String `json:"position_topic,omitempty"`

	// The maximum QoS level to be used when receiving and publishing messages.
	Qos Integer `json:"qos,omitempty" default:"0"`

	// Defines if published messages should have the retain flag set.
	Retain Boolean `json:"retain,omitempty" default:"false"`

	// Defines a template to define the position to be sent to the set_position_topic topic. Incoming position value is available for use in the template {{ position }}. Within the template the following variables are available: entity_id, position, the target position in percent; position_open; position_closed; tilt_min; tilt_max. The entity_id can be used to reference the entity’s attributes with help of the states template function;
	SetPositionTemplate String `json:"set_position_template,omitempty"`

	// The MQTT topic to publish position commands to. You need to set position_topic as well if you want to use position topic. Use template if position topic wants different values than within range position_closed - position_open. If template is not defined and position_closed != 100 and position_open != 0 then proper position value is calculated from percentage position.
	SetPositionTopic String `json:"set_position_topic,omitempty"`

	// The payload that represents the closed state.
	StateClosed String `json:"state_closed,omitempty" default:"closed"`

	// The payload that represents the closing state.
	StateClosing String `json:"state_closing,omitempty" default:"closing"`

	// The payload that represents the open state.
	StateOpen String `json:"state_open,omitempty" default:"open"`

	// The payload that represents the opening state.
	StateOpening String `json:"state_opening,omitempty" default:"opening"`

	// The payload that represents the stopped state (for covers that do not report open/closed state).
	StateStopped String `json:"state_stopped,omitempty" default:"stopped"`

	// The MQTT topic subscribed to receive cover state messages. State topic can only read (open, opening, closed, closing or stopped) state.
	StateTopic String `json:"state_topic,omitempty"`

	// The value that will be sent on a close_cover_tilt command.
	TiltClosedValue Integer `json:"tilt_closed_value,omitempty" default:"0"`

	// Defines a template that can be used to extract the payload for the tilt_command_topic topic. Within the template the following variables are available: entity_id, tilt_position, the target tilt position in percent; position_open; position_closed; tilt_min; tilt_max. The entity_id can be used to reference the entity’s attributes with help of the states template function;
	TiltCommandTemplate String `json:"tilt_command_template,omitempty"`

	// The MQTT topic to publish commands to control the cover tilt.
	TiltCommandTopic String `json:"tilt_command_topic,omitempty"`

	// The maximum tilt value.
	TiltMax Integer `json:"tilt_max,omitempty" default:"100"`

	// The minimum tilt value.
	TiltMin Integer `json:"tilt_min,omitempty" default:"0"`

	// The value that will be sent on an open_cover_tilt command.
	TiltOpenedValue Integer `json:"tilt_opened_value,omitempty" default:"100"`

	// Flag that determines if tilt works in optimistic mode.
	TiltOptimistic Boolean `json:"tilt_optimistic,omitempty"`

	// Default: true if tilt_status_topic is not defined, else false
	// Defines a template that can be used to extract the payload for the tilt_status_topic topic. Within the template the following variables are available: entity_id, position_open; position_closed; tilt_min; tilt_max. The entity_id can be used to reference the entity’s attributes with help of the states template function;
	TiltStatusTemplate String `json:"tilt_status_template,omitempty"`

	// The MQTT topic subscribed to receive tilt status update values.
	TiltStatusTopic String `json:"tilt_status_topic,omitempty"`

	// An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception.
	UniqueId String `json:"unique_id,omitempty"`

	// Defines a template that can be used to extract the payload for the state_topic topic.
	ValueTemplate String `json:"value_template,omitempty"`
}

func (c *Cover) Json() string {
	j, _ := json.Marshal(*c)
	return string(j)
}

func (config *EntityConfig) IsCover() bool {
	var ok bool

	for range Only.Once {
		if config.Units == LabelCover {
			ok = true
			break
		}
	}

	return ok
}
