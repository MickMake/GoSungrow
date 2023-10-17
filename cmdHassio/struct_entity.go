package cmdHassio

import (
	"fmt"

	"github.com/MickMake/GoUnify/Only"
	"github.com/anicoll/gosungrow/iSolarCloud/api"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/valueTypes"
)

type EntityConfig struct {
	// Type          string
	Name    string
	SubName string

	ParentId   string
	ParentName string

	UniqueId    string
	FullId      string
	Units       string
	ValueName   string
	DeviceClass string
	StateClass  string
	Icon        string

	Value         *valueTypes.UnitValue
	Point         *api.Point
	ValueTemplate string

	UpdateFreq             string
	LastReset              string
	LastResetValueTemplate string

	IgnoreUpdate bool

	haType  string
	Options []string
}

func (config *EntityConfig) FixConfig() {
	for range Only.Once {
		// mdi:power-socket-au
		// mdi:solar-power
		// mdi:home-lightning-bolt-outline
		// mdi:transmission-tower
		// mdi:transmission-tower-export
		// mdi:transmission-tower-import
		// mdi:transmission-tower-off
		// mdi:home-battery-outline
		// mdi:lightning-bolt
		// mdi:check-circle-outline | mdi:arrow-right-bold
		// mdi:transmission-tower

		// Set default ValueTemplate
		switch {
		case config.Point.IsBool():
			fallthrough
		case config.Value.IsBool():
			fallthrough
		case config.IsBinarySensor():
			config.ValueTemplate = SetDefault(config.ValueTemplate, "{{ value_json.value }}")

		case config.Value.IsFloat():
			if !config.Value.Valid {
				config.IgnoreUpdate = true
			}
			cnv := "| float"
			if config.Value.String() == "" {
				cnv = ""
			}
			vj := "value"
			if config.ValueName != "" {
				vj = config.ValueName
			}
			config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s %s }}", vj, cnv))

		case config.Value.IsTypeDateTime():
			vj := "value"
			if config.ValueName != "" {
				vj = config.ValueName
			}
			value, _, err := valueTypes.ParseDateTime(config.Value.String())
			if err == nil {
				config.Value.SetString(value.Local().Format(valueTypes.DateTimeFullLayout))
				config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s | as_datetime }}", vj))
			} else {
				config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s }}", vj))
			}

		case config.Value.IsInt():
			vj := "value"
			if config.ValueName != "" {
				vj = config.ValueName
			}
			config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s | int }}", vj))

		default:
			vj := "value"
			if config.ValueName != "" {
				vj = config.ValueName
			}
			config.ValueTemplate = SetDefault(config.ValueTemplate, fmt.Sprintf("{{ value_json.%s }}", vj))
		}

		// Set DeviceClass & Icon
		switch {
		case config.Point.IsBool():
			fallthrough
		case config.Value.IsBool():
			fallthrough
		case config.IsBinarySensor():
			config.DeviceClass = SetDefault(config.DeviceClass, "power")
			config.Icon = SetDefault(config.Icon, "mdi:check-circle-outline")

		case config.Value.TypeValue == "Power":
			fallthrough
		case config.Units == "MW":
			fallthrough
		case config.Units == "kW":
			fallthrough
		case config.Units == "W":
			config.DeviceClass = SetDefault(config.DeviceClass, "power")
			config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

		case config.Value.TypeValue == "Energy":
			fallthrough
		case config.Units == "MWh":
			fallthrough
		case config.Units == "kWh":
			fallthrough
		case config.Units == "Wh":
			config.DeviceClass = SetDefault(config.DeviceClass, "energy")
			config.Icon = SetDefault(config.Icon, "mdi:transmission-tower")

		case config.Units == "var":
			fallthrough
		case config.Units == "kvar":
			config.DeviceClass = SetDefault(config.DeviceClass, "reactive_power")
			config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

		case config.Units == "VA":
			config.DeviceClass = SetDefault(config.DeviceClass, "apparent_power")
			config.Icon = SetDefault(config.Icon, "mdi:lightning-bolt")

		case config.Units == "Hz":
			config.DeviceClass = SetDefault(config.DeviceClass, "frequency")
			config.Icon = SetDefault(config.Icon, "mdi:sine-wave")

		case config.Units == "V":
			config.DeviceClass = SetDefault(config.DeviceClass, "voltage")
			config.Icon = SetDefault(config.Icon, "mdi:current-dc")

		case config.Units == "A":
			config.DeviceClass = SetDefault(config.DeviceClass, "current")
			config.Icon = SetDefault(config.Icon, "mdi:current-ac")

		case config.Units == "°F":
			fallthrough
		case config.Units == "F":
			fallthrough
		case config.Units == "℉":
			config.DeviceClass = SetDefault(config.DeviceClass, "temperature")
			config.Units = "℉"
			config.Icon = SetDefault(config.Icon, "mdi:thermometer")

		case config.Units == "°C":
			fallthrough
		case config.Units == "C":
			fallthrough
		case config.Units == "℃":
			config.DeviceClass = SetDefault(config.DeviceClass, "temperature")
			config.Units = "°C"
			config.Icon = SetDefault(config.Icon, "mdi:thermometer")

		case config.Icon == "mdi:home-battery-outline":
			fallthrough
		case config.Icon == "mdi:battery":
			config.DeviceClass = SetDefault(config.DeviceClass, "battery")
			// config.Icon = SetDefault(config.Icon, "mdi:percent") // mdi:home-battery-outline

		case config.Value.TypeValue == "Percent":
			fallthrough
		case config.Units == "%":
			// @TODO - Not supported in older versions of HA.
			// config.DeviceClass = SetDefault(config.DeviceClass, "battery")
			config.Icon = SetDefault(config.Icon, "mdi:percent") // mdi:home-battery-outline

		case config.Value.TypeValue == "DateTime":
			config.DeviceClass = SetDefault(config.DeviceClass, "timestamp") // date
			config.Icon = SetDefault(config.Icon, "mdi:clock-outline")

		case config.Units == "h":
			// config.DeviceClass = SetDefault(config.DeviceClass, "timestamp") // date
			config.Icon = SetDefault(config.Icon, "mdi:clock-outline")

		case config.Units == "kg":
			config.DeviceClass = SetDefault(config.DeviceClass, "weight")
			config.Icon = SetDefault(config.Icon, "mdi:weight")

		case config.Units == "km":
			config.DeviceClass = SetDefault(config.DeviceClass, "distance")
			config.Icon = SetDefault(config.Icon, "mdi:map-marker-distance")

		case config.Units == "Wh/㎡":
			fallthrough
		case config.Units == "W/㎡":
			// @TODO - Not supported in older versions of HA.
			config.DeviceClass = SetDefault(config.DeviceClass, "irradiance")
			config.Icon = SetDefault(config.Icon, "mdi:weather-sunny")

		case config.Value.TypeValue == "Currency":
			fallthrough
		case config.Units == "AUD":
			fallthrough
		case config.Units == "$":
			config.DeviceClass = SetDefault(config.DeviceClass, "monetary")
			config.Icon = SetDefault(config.Icon, "mdi:currency-usd")

			// p13013 - power_factor

		case config.Units == "GPS":
			// config.DeviceClass = SetDefault(config.DeviceClass, "")
			config.Icon = SetDefault(config.Icon, "mdi:crosshairs-gps")

		default:
			config.DeviceClass = SetDefault(config.DeviceClass, "")
			config.Icon = SetDefault(config.Icon, "")
		}

		switch {
		case config.Point.IsBoot():
			config.StateClass = "measurement"
			config.LastReset = ""
			config.LastResetValueTemplate = ""

		case config.Point.IsDaily():
			fallthrough
		case config.Point.IsMonthly():
			fallthrough
		case config.Point.IsYearly():
			fallthrough
		case config.Point.IsTotal():
			config.StateClass = "total"
			config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | as_datetime }}")
			// config.LastReset = config.Point.WhenReset(config.Date)

		case config.Point.Is5Minute():
			fallthrough
		case config.Point.Is15Minute():
			fallthrough
		case config.Point.Is30Minute():
			fallthrough
		case config.Point.IsInstant():
			fallthrough
		default:
			config.StateClass = "measurement"
			config.LastReset = ""
			config.LastResetValueTemplate = ""
		}

		// if config.LastReset == "" {
		// 	break
		// }
		//
		// pt := api.GetDevicePoint(config.FullId)
		// if !pt.Valid {
		// 	break
		// }
		//
		// config.LastReset = pt.WhenReset()
		// config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | as_datetime }}")
		// config.LastResetValueTemplate = SetDefault(config.LastResetValueTemplate, "{{ value_json.last_reset | int | timestamp_local | as_datetime }}")
		// config.StateClass = "total"
		// config.StateClass = "measurement"
	}
}

func SetDefault(value string, def string) string {
	if value == "" {
		value = def
	}
	return value
}
