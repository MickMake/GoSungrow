package cmd

import (
	"GoSungrow/iSolarCloud"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/MickMake/GoUnify/cmdPath"
	"github.com/spf13/cobra"
	"strings"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdHa CmdDefault

func NewCmdHa() *CmdHa {
	var ret *CmdHa

	for range Only.Once {
		ret = &CmdHa{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *CmdHa) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "ha",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Ha"},
			Short:                 fmt.Sprintf("Home Assistant commands."),
			Long:                  fmt.Sprintf("Home Assistant commands."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "")

		// ********************************************************************************
		var cmdHaGet = &cobra.Command{
			Use:                   "lovelace",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Ha"},
			Short:                 fmt.Sprintf("Produce Lovelace card"),
			Long:                  fmt.Sprintf("Produce Lovelace card"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.CmdLovelace,
			Args: cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdHaGet)
		cmdHaGet.Example = cmdHelp.PrintExamples(cmdHaGet, "[area.]<endpoint>")
	}
	return c.SelfCmd
}


func (c *CmdHa) CmdLovelace(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		// if len(args) == 0 {
		// 	fmt.Println("One of: basic graphs stats")
		// 	_ = cmd.Help()
		// 	break
		// }
		//
		// var pids valueTypes.PsIds
		// pids, c.Error = cmds.Api.SunGrow.GetPsIds()
		// if c.Error != nil {
		// 	break
		// }

		pids := cmds.Api.SunGrow.SetPsIds(args...)
		if cmds.Api.SunGrow.Error != nil {
			c.Error = cmds.Api.SunGrow.Error
			break
		}

		for _, pid := range pids {
			var tree iSolarCloud.PsTree
			tree, c.Error = cmds.Api.SunGrow.PsTreeMenu(pid.String())
			if c.Error != nil {
				break
			}

			var DeviceType14 string
			var DeviceType22 string
			var DeviceType43 string

			for _, device := range tree.Devices {
				if device.DeviceType.Match(14) {
					DeviceType14 = device.PsKey.String()
					continue
				}
				if device.DeviceType.Match(22) {
					DeviceType22 = device.PsKey.String()
					continue
				}
				if device.DeviceType.Match(43) {
					DeviceType43 = device.PsKey.String()
					continue
				}
			}

			if DeviceType14 == "" {
				fmt.Printf("Can't find DeviceType 14 attached to ps_id %s.\n", pid)
				continue
			}

			if DeviceType22 == "" {
				fmt.Printf("Can't find DeviceType 22 attached to ps_id %s.\n", pid)
				continue
			}

			if DeviceType43 == "" {
				fmt.Printf("Can't find DeviceType 43 attached to ps_id %s.\n", pid)
				continue
			}

			data := lovelaceBasic
			data = strings.ReplaceAll(data, "{{ PsId }}", pid.String())
			data = strings.ReplaceAll(data, "{{ DeviceType:14 }}", DeviceType14)
			data = strings.ReplaceAll(data, "{{ DeviceType:22 }}", DeviceType22)
			data = strings.ReplaceAll(data, "{{ DeviceType:43 }}", DeviceType43)
			c.Error = cmdPath.PlainFileWrite(fmt.Sprintf("Lovelace-Basic-%s.yaml", pid.String()), []byte(data), 0644)
			if c.Error != nil {
				break
			}

			data = lovelaceGraphs
			data = strings.ReplaceAll(data, "{{ PsId }}", pid.String())
			data = strings.ReplaceAll(data, "{{ DeviceType:14 }}", DeviceType14)
			data = strings.ReplaceAll(data, "{{ DeviceType:22 }}", DeviceType22)
			data = strings.ReplaceAll(data, "{{ DeviceType:43 }}", DeviceType43)
			c.Error = cmdPath.PlainFileWrite(fmt.Sprintf("Lovelace-Graphs-%s.yaml", pid.String()), []byte(data), 0644)
			if c.Error != nil {
				break
			}

			data = lovelaceStats
			data = strings.ReplaceAll(data, "{{ PsId }}", pid.String())
			data = strings.ReplaceAll(data, "{{ DeviceType:14 }}", DeviceType14)
			data = strings.ReplaceAll(data, "{{ DeviceType:22 }}", DeviceType22)
			data = strings.ReplaceAll(data, "{{ DeviceType:43 }}", DeviceType43)
			c.Error = cmdPath.PlainFileWrite(fmt.Sprintf("Lovelace-Stats-%s.yaml", pid.String()), []byte(data), 0644)
			if c.Error != nil {
				break
			}
		}

		// switch args[0] {
		// 	case "basic":
		// 	case "graphs":
		// 	case "stats":
		// }

	}

	return c.Error
}

const lovelaceBasic = `views:
  - theme: Backend-selected
    title: Status
    path: status
    type: panel
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: entities
            entities:
              - entity: binary_sensor.gosungrow_getpslist_devices_{{ PsId }}_ps_status
            state_color: true
          - square: false
            columns: 1
            type: grid
            cards:
              - type: picture-elements
                image: /local/SungrowEnergy3.png
                elements:
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power_active
                    tap_action: none
                    state_image:
                      'off': /local/SungrowEnergy2-GridOff.png
                    state_filter:
                      'off': opacity(100%)
                      'on': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power_active
                    tap_action: none
                    state_image:
                      'off': /local/SungrowEnergy2-BatteryOff.png
                    state_filter:
                      'off': opacity(100%)
                      'on': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power_active
                    tap_action: none
                    state_image:
                      'off': /local/SungrowEnergy2-LoadOff.png
                    state_filter:
                      'off': opacity(100%)
                      'on': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power_active
                    tap_action: none
                    state_image:
                      'off': /local/SungrowEnergy2-PVOff.png
                    state_filter:
                      'off': opacity(100%)
                      'on': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
                    tap_action: none
                    state_image:
                      'on': /local/SungrowEnergy2-PVToBattery.png
                    state_filter:
                      'on': opacity(100%)
                      'off': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
                    tap_action: none
                    state_image:
                      'on': /local/SungrowEnergy2-PVToLoad.png
                    state_filter:
                      'on': opacity(100%)
                      'off': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
                    tap_action: none
                    state_image:
                      'on': /local/SungrowEnergy2-PVToGrid.png
                    state_filter:
                      'on': opacity(100%)
                      'off': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
                    tap_action: none
                    state_image:
                      'on': /local/SungrowEnergy2-BatteryToLoad.png
                    state_filter:
                      'on': opacity(100%)
                      'off': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: image
                    entity: >-
                      binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
                    tap_action: none
                    state_image:
                      'on': /local/SungrowEnergy2-GridToLoad.png
                    state_filter:
                      'on': opacity(100%)
                      'off': opacity(0%)
                    style:
                      top: 50%
                      left: 50%
                      width: 100%
                  - type: state-label
                    entity: binary_sensor.gosungrow
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
                        style:
                          background-color: rgba(0, 102, 204, 1)
                          line-height: 8px
                          padding: 0px 0px
                          margin: 0px 0px
                          color: white
                          top: 4%
                          left: 50%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power
                        style:
                          background-color: rgba(0, 102, 204, 1)
                          line-height: 8px
                          padding: 0px 0px
                          margin: 0px 0px
                          color: white
                          top: 85%
                          left: 15.5%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
                        style:
                          background-color: rgba(0, 102, 204, 1)
                          line-height: 8px
                          padding: 0px 0px
                          margin: 0px 0px
                          color: white
                          top: 85%
                          left: 85%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13119
                        style:
                          background-color: rgba(0, 102, 204, 1)
                          line-height: 8px
                          padding: 0px 0px
                          margin: 0px 0px
                          color: white
                          top: 65%
                          left: 50%
                  - type: state-label
                    entity: binary_sensor.gosungrow
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13126
                        style:
                          color: blue
                          top: 27%
                          left: 10%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: >-
                          sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power
                        style:
                          color: blue
                          top: 30%
                          left: 42%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13121
                        style:
                          color: blue
                          top: 27%
                          left: 90%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13029
                        style:
                          color: blue
                          top: 61%
                          left: 27%
                  - type: conditional
                    conditions:
                      - entity: >-
                          binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
                        state: 'on'
                    elements:
                      - type: state-label
                        entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
                        style:
                          color: blue
                          top: 61%
                          left: 73%
                  - type: state-label
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13141
                    style:
                      color: blue
                      top: 75.5%
                      left: 15.5%
          - type: entities
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13011
              - entity: sensor.gosungrow_getpslist_devices_{{ PsId }}_installed_power_map
            state_color: true
  - theme: Backend-selected
    title: Today Yield
    path: today-yield
    type: panel
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: entity
            entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_daily_energy
            name: Today Yield
            state_color: true
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy_percent
                card_height: 200
                gauge:
                  type: radial-gauge
                  width: 200
                  height: 180
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(254,141,75,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  highlights: false
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxHeight: 100
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - ''
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: false
              - type: vertical-stack
                cards:
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
                    name: Self-consumption of PV
                    state_color: true
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy_percent
                    name: Percent
                    state_color: true
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy_percent
                card_height: 200
                gauge:
                  type: radial-gauge
                  width: 200
                  height: 180
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  highlights: false
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxHeight: 100
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - ''
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: false
              - type: vertical-stack
                cards:
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
                    name: Battery Charging Energy from PV
                    state_color: true
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy_percent
                    name: Percent
                    state_color: true
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_energy_percent
                card_height: 200
                gauge:
                  type: radial-gauge
                  width: 200
                  height: 180
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  highlights: false
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxHeight: 100
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - ''
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: false
              - type: vertical-stack
                cards:
                  - type: entity
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_energy
                    name: Feed-in Energy from PV
                    state_color: true
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_energy_percent
                    name: Percent
                    state_color: true
          - type: entity
            entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_total_energy
            name: Total Yield
            state_color: true
  - theme: Backend-selected
    title: E-use Today
    path: e-use-today
    type: panel
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: entity
            entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_total_daily_energy
            name: E-use Today
            state_color: true
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_daily_energy_percent
                card_height: 200
                gauge:
                  type: radial-gauge
                  width: 200
                  height: 180
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(7,205,205,1)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  highlights: false
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxHeight: 100
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - ''
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: false
              - type: vertical-stack
                cards:
                  - type: entity
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_energy
                    name: Self-sufficiency
                    state_color: true
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_daily_energy_percent
                    name: Percent
                    state_color: true
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_energy_percent
                card_height: 200
                gauge:
                  type: radial-gauge
                  width: 200
                  height: 180
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,1)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  highlights: false
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxHeight: 100
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - ''
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: false
              - type: vertical-stack
                cards:
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_energy
                    name: Purchased Energy
                    state_color: true
                  - type: entity
                    entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_energy_percent
                    name: Percent
                    state_color: true
          - type: entity
            entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13130
            name: Cumulative Electricity Consumption
            state_color: true
          - type: horizontal-stack
            cards:
              - type: entity
                entity: sensor.gosungrow_getpslist_devices_{{ PsId }}_today_income
                name: Today Revenue
                icon: mdi:currency-usd
                unit: $
                state_color: true
              - type: entity
                entity: sensor.gosungrow_getpslist_devices_{{ PsId }}_total_income
                name: Total Revenue
                icon: mdi:currency-usd
                unit: $
                state_color: true
  - theme: Backend-selected
    title: graphs
    path: graphs
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: custom:mini-graph-card
            show:
              labels: true
              points: false
              icon: false
              name: false
              average: true
              extrema: true
              state: false
            hour24: true
            animate: false
            unit: kW
            points_per_hour: 20
            line_width: 1
            hours_to_show: 18
            height: 260
            hover_action: hover-info
            tap_action: none
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                name: PV
                color: rgb(255, 152, 48, .75)
                show_points: false
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power
                name: Grid
                color: rgb(87, 148, 242, .75)
                show_points: false
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power
                name: Battery
                color: rgb(115, 191, 105, .75)
                show_points: false
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power
                name: Load
                color: rgb(184, 119, 217, .75)
                show_points: false
  - theme: Backend-selected
    title: Activity
    path: activity
    type: panel
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: history-graph
            hours_to_show: 18
            entities:
              - entity: binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
            title: PV
          - type: history-graph
            hours_to_show: 18
            entities:
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
            title: Battery
          - type: history-graph
            hours_to_show: 18
            entities:
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
            title: Grid
          - type: history-graph
            hours_to_show: 18
            entities:
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
            title: Load
  - theme: Backend-selected
    title: Inverter
    path: inverter
    type: panel
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: grid
            cards:
              - type: gauge
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13019
                min: 0
                max: 90
                needle: true
                severity:
                  green: 10
                  yellow: 40
                  red: 70
                name: Inverter Temperature
              - type: gauge
                entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58603
                min: 0
                max: 90
                needle: true
                severity:
                  green: 10
                  yellow: 40
                  red: 60
                name: Battery Temperature
              - type: gauge
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
                min: 0
                max: 9
                needle: true
                severity:
                  green: 6
                  yellow: 3
                  red: 0
                name: PV Yield
              - type: gauge
                entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58604
                min: 0
                max: 100
                needle: true
                severity:
                  green: 50
                  yellow: 25
                  red: 0
                name: Battery Level
            columns: 2
            square: false
          - type: entities
            entities:
              - entity: sensor.gosungrow
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13007
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13160
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13012
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13013
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13144
              - entity: sensor.energy_production_today
              - entity: sensor.power_production_now
              - entity: sensor.energy_current_hour
              - entity: sensor.power_production_next_24hours
              - entity: sensor.power_production_next_12hours
            state_color: true
            title: SunGrow Inverter
`

const lovelaceGraphs = `views:
  - theme: Backend-selected
    title: Power
    path: power
    type: panel
    icon: mdi:solar-power
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: markdown
            content: ' '
            title: Power (kW)
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV Yield (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(254,141,75,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
                    name: PV Yield
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Battery (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power
                    name: PV to Battery
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Grid (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power
                    name: PV to Grid
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Load (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(7,205,205,1)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power
                    name: PV to Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13119
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Load (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(254,141,75,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(50, 200, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(200, 50, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13119
                    name: Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery to Load (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 200, 50, .75)
                  - value: 6
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power
                    name: Battery to Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Grid to Load (kW)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(50, 200, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(200, 50, 50, .75)
                hour24: true
                animate: false
                unit: kW
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
                    name: Grid to Load
  - theme: Backend-selected
    title: Energy
    path: energy
    type: panel
    icon: mdi:home-lightning-bolt-outline
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: markdown
            content: ' '
            title: Energy (kWh)
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13112
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV Yield (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(254,141,75,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 60
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '15'
                    - '30'
                    - '45'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 15
                      color: rgba(200, 50, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 200, 50, .75)
                    - from: 30
                      to: 60
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 15
                    color: rgba(200, 200, 50, .75)
                  - value: 30
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 60
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13112
                    name: PV Yield
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Battery (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
                    name: PV to Battery
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Load (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(7,205,205,1)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 60
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '15'
                    - '30'
                    - '45'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 15
                      color: rgba(200, 50, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 200, 50, .75)
                    - from: 30
                      to: 60
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 10
                    color: rgba(200, 200, 50, .75)
                  - value: 20
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 60
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
                    name: PV to Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: PV to Grid (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 60
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '15'
                    - '30'
                    - '45'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 15
                      color: rgba(200, 50, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 200, 50, .75)
                    - from: 30
                      to: 60
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 15
                    color: rgba(200, 200, 50, .75)
                  - value: 30
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 60
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
                    name: PV to Grid
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Load (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(254,141,75,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 60
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '15'
                    - '30'
                    - '45'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 15
                      color: rgba(50, 200, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 200, 50, .75)
                    - from: 30
                      to: 60
                      color: rgba(200, 50, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 25
                    color: rgba(200, 200, 50, .75)
                  - value: 35
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 60
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
                    name: Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: >-
                  sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_discharge_energy
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery to Load (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 12
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 200, 50, .75)
                  - value: 5
                    color: rgba(50, 200, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_discharge_energy
                    name: Battery to Load
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Grid to Load (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(74,176,249,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 30
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '5'
                    - '10'
                    - '15'
                    - '20'
                    - '25'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 5
                      color: rgba(50, 200, 50, .75)
                    - from: 5
                      to: 15
                      color: rgba(200, 200, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 50, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(50, 200, 50, .75)
                  - value: 6
                    color: rgba(200, 200, 50, .75)
                  - value: 12
                    color: rgba(200, 50, 50, .75)
                hour24: true
                animate: false
                unit: kWh
                lower_bound: 0
                upper_bound: 24
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
                    name: Grid to Load
  - theme: Backend-selected
    type: panel
    title: Battery
    icon: mdi:home-battery
    path: battery
    badges: []
    cards:
      - type: vertical-stack
        cards:
          - type: markdown
            content: ' '
            title: Battery
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Charge (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '2'
                    - '4'
                    - '6'
                    - '8'
                    - '10'
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 4
                      color: rgba(200, 50, 50, .75)
                    - from: 4
                      to: 8
                      color: rgba(200, 200, 50, .75)
                    - from: 8
                      to: 10.36
                      color: rgba(50, 200, 50, .75)
                    - from: 10.36
                      to: 12
                      color: rgba(128, 128, 128, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                unit: kWh
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: >-
                      sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
                    name: Battery Charge
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13140
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Capacity (kWh)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 12
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 3
                  majorTicks:
                    - '0'
                    - '3'
                    - '6'
                    - '9'
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 10.36
                      color: rgba(50, 200, 50, .75)
                    - from: 10.36
                      to: 25
                      color: rgba(128, 128, 128, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 4
                    color: rgba(200, 200, 50, .75)
                  - value: 8
                    color: rgba(50, 200, 50, .75)
                hour24: true
                unit: kWh
                lower_bound: 0
                upper_bound: 12
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13140
                    name: Battery Capacity
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13141
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Level (%)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - '0'
                    - '25'
                    - '50'
                    - '75'
                    - ''
                  minorTicks: 0
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 25
                      color: rgba(200, 50, 50, .75)
                    - from: 25
                      to: 50
                      color: rgba(200, 200, 50, .75)
                    - from: 50
                      to: 100
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 25
                    color: rgba(200, 200, 50, .75)
                  - value: 50
                    color: rgba(50, 200, 50, .75)
                hour24: true
                unit: '%'
                lower_bound: 0
                upper_bound: 100
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13141
                    name: Battery Level
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13142
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Health (%)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 100
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - '0'
                    - '25'
                    - '50'
                    - '75'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 25
                      color: rgba(200, 50, 50, .75)
                    - from: 25
                      to: 50
                      color: rgba(200, 200, 50, .75)
                    - from: 50
                      to: 100
                      color: rgba(50, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 50, 50, .75)
                  - value: 25
                    color: rgba(200, 200, 50, .75)
                  - value: 50
                    color: rgba(50, 200, 50, .75)
                hour24: true
                unit: '%'
                lower_bound: 0
                upper_bound: 100
                icon: mdi:battery-heart-outline
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13142
                    name: Battery Health
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58601
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Voltage (V)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: 0
                  maxValue: 500
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - '0'
                    - '100'
                    - '200'
                    - '300'
                    - '400'
                    - ''
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: 0
                      to: 150
                      color: rgba(200, 200, 50, .75)
                    - from: 150
                      to: 350
                      color: rgba(50, 200, 50, .75)
                    - from: 350
                      to: 500
                      color: rgba(200, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: 0
                    color: rgba(200, 200, 50, .75)
                  - value: 150
                    color: rgba(50, 200, 50, .75)
                  - value: 250
                    color: rgba(200, 200, 50, .75)
                hour24: true
                unit: V
                lower_bound: 0
                upper_bound: 500
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58601
                    name: Battery Voltage
          - type: horizontal-stack
            cards:
              - type: custom:canvas-gauge-card
                entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58602
                style:
                  top: 50%
                  left: 50%
                  width: null
                  heigth: null
                  box-shadow: none
                  background-color: transparent
                  transform: scale(1,1) translate(-50%,-50%)
                gauge:
                  type: radial-gauge
                  title: Battery Current (A)
                  height: 190
                  width: 190
                  barWidth: 20
                  barShadow: 1
                  colorBarProgress: rgba(125,195,77,0.8)
                  colorBar: '#0000'
                  borderShadowWidth: 0
                  borderInnerWidth: 0
                  borderOuterWidth: 0
                  borderMiddleWidth: 0
                  colorValueBoxShadow: 0
                  colorValueBoxBackground: transparent
                  needle: false
                  needleCircleSize: 15
                  needleCircleOuter: false
                  minValue: -30
                  maxValue: 30
                  startAngle: 180
                  ticksAngle: 360
                  valueBox: true
                  valueBoxWidth: 40
                  valueBoxStroke: 1
                  valueBoxBorderRadius: 5
                  valueTextShadow: true
                  valueInt: 1
                  valueDec: 1
                  majorTicks:
                    - '-30'
                    - '20'
                    - '-10'
                    - '0'
                    - '10'
                    - '20'
                    - '30'
                  minorTicks: 2
                  strokeTicks: false
                  borders: true
                  highlights:
                    - from: -30
                      to: -15
                      color: rgba(200, 200, 50, .75)
                    - from: -15
                      to: 15
                      color: rgba(50, 200, 50, .75)
                    - from: 15
                      to: 30
                      color: rgba(200, 200, 50, .75)
              - type: custom:mini-graph-card
                show:
                  labels: true
                  points: false
                  icon: false
                  name: false
                  average: true
                  extrema: true
                color_thresholds:
                  - value: -25
                    color: rgba(200, 200, 50, .75)
                  - value: 0
                    color: rgba(50, 200, 50, .75)
                  - value: 25
                    color: rgba(200, 200, 50, .75)
                hour24: true
                unit: A
                lower_bound: -50
                upper_bound: 50
                points_per_hour: 20
                line_width: 1
                hours_to_show: 16
                height: 150
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58602
                    name: Battery Current
  - title: VA
    path: va
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - square: false
            columns: 2
            type: grid
            cards:
              - type: grid
                cards:
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13138
                    min: 0
                    max: 400
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13161
                    min: 0
                    max: 800
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13001
                    min: 0
                    max: 400
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13105
                    min: 0
                    max: 400
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13157
                    min: 0
                    max: 400
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13158
                    min: 0
                    max: 400
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13159
                    min: 0
                    max: 400
                    needle: true
                columns: 1
                square: false
              - type: grid
                cards:
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13139
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13161
                    min: 0
                    max: 800
                    needle: true
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13002
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13106
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13008
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13009
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                  - type: gauge
                    entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13010
                    min: 0
                    max: 10
                    needle: true
                    severity:
                      green: 0
                      yellow: 8
                      red: 9
                columns: 1
                square: false
  - theme: Backend-selected
    title: Direction
    path: direction
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - entity: binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
            state_color: true
            title: PV
          - type: entities
            entities:
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
            state_color: true
            title: Battery
          - type: entities
            entities:
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_battery_power_active
              - entity: >-
                  binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
            state_color: true
            title: Grid
          - type: entities
            entities:
              - binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power_active
            state_color: true
            title: Load
  - theme: Backend-selected
    title: Yield (Week)
    path: yield-week
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: custom:mini-graph-card
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: max
                name: Max
                color: null
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: min
                name: Min
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: avg
                name: Avg
                color: green
            name: PV Yield (week)
            hours_to_show: 168
            group_by: hour
            points_per_hour: 20
            line_width: 1
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    type: panel
    title: Yield (fortnight)
    path: yield-fortnight
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: custom:mini-graph-card
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: max
                name: Max
                color: null
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: min
                name: Min
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                aggregate_func: avg
                name: Avg
                color: green
            name: PV Yield (fortnight)
            hours_to_show: 336
            group_by: hour
            points_per_hour: 20
            line_width: 1
          - chart_type: line
            period: hour
            days_to_show: 14
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Yield (Month)
    type: panel
    path: yield-month
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - square: false
            columns: 1
            type: grid
            cards:
              - type: custom:mini-graph-card
                entities:
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                    aggregate_func: max
                    name: Max
                    color: null
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                    aggregate_func: min
                    name: Min
                  - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                    aggregate_func: avg
                    name: Avg
                    color: green
                name: PV Yield (monthly)
                hours_to_show: 744
                group_by: hour
                points_per_hour: 20
                line_width: 1
              - chart_type: line
                period: hour
                days_to_show: 31
                type: statistics-graph
                entities:
                  - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
                stat_types:
                  - mean
                  - min
                  - max
  - theme: Backend-selected
    title: Energy
    path: info-energy
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13125
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13175
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13130
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13137
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13148
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13176
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13174
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13122
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13199
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13116
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13144
            title: Energy
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13125
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13175
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13130
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13137
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13148
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13176
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13174
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13199
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13122
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13144
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13116
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Power
    path: info-power
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13126
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13150
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
            title: Power
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13126
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13150
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Battery
    path: info-battery
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_energy_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_grid_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power_active
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_charge_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_discharge_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_discharge_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_to_load_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_battery_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13034
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13035
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13176
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13028
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13029
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13140
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13174
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13126
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13150
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13138
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13139
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13162
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13163
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13143
            title: Battery Information
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13034
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13035
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13176
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13028
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13029
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13140
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13174
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13126
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13150
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13138
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13139
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13162
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13163
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13143
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Grid
    path: info-grid
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_battery_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power_active
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_to_load_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_grid_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13121
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13125
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13148
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13175
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13122
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
            state_color: true
            title: Grid Information
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13125
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13148
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13175
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13122
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13147
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13173
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13121
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13149
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Load
    path: info-load
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power_active
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_daily_total_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_load_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_consumption_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_total_daily_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_total_load_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13119
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13130
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13137
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13116
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13199
            title: Load Information
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13119
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13130
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13137
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13116
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13199
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: MPPT
    path: info-mppt
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13001
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13105
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13002
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13106
            title: MPPT Information
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13001
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13105
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13002
              - entity: sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13106
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Overview
    path: info-overview
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power_active
              - >-
                binary_sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power_active
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_daily_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_total_energy
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_pv_to_load_power
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13011
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18065
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18066
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18067
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18068
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13112
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13134
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13008
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13009
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13010
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18062
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18063
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18064
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13157
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13158
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13159
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13161
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13007
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13012
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13013
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13019
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13160
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13155
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18061
            title: Overview
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13003
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13011
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18065
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18066
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18067
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18068
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13112
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13134
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13008
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13009
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13010
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18062
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18063
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p18064
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13157
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13158
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13159
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13161
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13007
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13012
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13013
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13019
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13160
            stat_types:
              - mean
              - min
              - max
  - theme: Backend-selected
    title: Other
    path: info-other
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - type: entities
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:22 }}_p23014
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58601
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58602
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58603
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58606
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58607
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58608
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58609
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58610
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58611
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58612
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58613
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58614
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58615
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58616
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58617
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58618
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58619
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58620
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58621
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58622
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58623
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58624
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58625
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58626
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58627
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58628
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58629
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58630
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58631
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58632
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58633
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58635
              - sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58636
            title: Other Information
            state_color: true
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58601
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58602
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58603
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: hour
            days_to_show: 7
            type: statistics-graph
            entities:
              - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58606
              - entity: sensor.gosungrow_virtual_{{ DeviceType:43 }}_p58607
            stat_types:
              - mean
              - min
              - max
`

const lovelaceStats = `views:
  - theme: Backend-selected
    title: Home
    type: panel
    badges: []
    cards:
      - square: false
        columns: 1
        type: grid
        cards:
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_pv_daily_yield
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13112
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_battery_power
              - sensor.gosungrow_battery_power
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_grid_power
              - sensor.gosungrow_grid_power
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13160
              - sensor.gosungrow_array_insulation_resistance
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13007
              - sensor.gosungrow_grid_frequency
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_internal_air_temperature
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13019
            stat_types:
              - mean
              - min
              - max
          - chart_type: line
            period: day
            days_to_show: 90
            type: statistics-graph
            entities:
              - sensor.gosungrow_battery_temperature
              - sensor.gosungrow_virtual_{{ DeviceType:14 }}_p13143
            stat_types:
              - mean
              - min
              - max
`
