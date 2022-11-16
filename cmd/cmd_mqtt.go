package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/mmHa"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/MickMake/GoUnify/cmdLog"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"time"
)


const (
	flagMqttUsername   = "mqtt-user"
	flagMqttPassword   = "mqtt-password"
	flagMqttHost   = "mqtt-host"
	flagMqttPort   = "mqtt-port"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdMqtt struct {
	CmdDefault

	// HASSIO MQTT
	MqttUsername   string

	MqttPassword   string
	MqttHost       string
	MqttPort       string

	Mqtt *mmHa.Mqtt
	// SunGrow *iSolarCloud.SunGrow
}

func NewCmdMqtt() *CmdMqtt {
	var ret *CmdMqtt

	for range Only.Once {
		ret = &CmdMqtt {
			CmdDefault: CmdDefault {
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},
		}
	}

	return ret
}

func (c *CmdMqtt) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		var cmdMqtt = &cobra.Command{
			Use:                   "mqtt",
			Aliases:               []string{""},
			Short:                 fmt.Sprintf("Connect to a HASSIO broker."),
			Long:                  fmt.Sprintf("Connect to a HASSIO broker."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  c.CmdMqtt,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdMqtt)
		cmdMqtt.Example = cmdHelp.PrintExamples(cmdMqtt, "run", "sync")

		// ******************************************************************************** //
		var cmdMqttRun = &cobra.Command{
			Use:                   "run",
			Aliases:               []string{""},
			Short:                 fmt.Sprintf("One-off sync to a HASSIO broker."),
			Long:                  fmt.Sprintf("One-off sync to a HASSIO broker."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  cmds.CmdMqttRun,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdMqtt.AddCommand(cmdMqttRun)
		cmdMqttRun.Example = cmdHelp.PrintExamples(cmdMqttRun, "")

		// ******************************************************************************** //
		var cmdMqttSync = &cobra.Command{
			Use:                   "sync",
			Aliases:               []string{""},
			Short:                 fmt.Sprintf("Sync to a HASSIO MQTT broker."),
			Long:                  fmt.Sprintf("Sync to a HASSIO MQTT broker."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  cmds.CmdMqttSync,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdMqtt.AddCommand(cmdMqttSync)
		cmdMqttSync.Example = cmdHelp.PrintExamples(cmdMqttSync, "", "all")
	}
	return c.SelfCmd
}

func (c *CmdMqtt) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		cmd.PersistentFlags().StringVarP(&c.MqttUsername, flagMqttUsername, "", "", fmt.Sprintf("HASSIO: mqtt username."))
		viper.SetDefault(flagMqttUsername, "")
		cmd.PersistentFlags().StringVarP(&c.MqttPassword, flagMqttPassword, "", "", fmt.Sprintf("HASSIO: mqtt password."))
		viper.SetDefault(flagMqttPassword, "")
		cmd.PersistentFlags().StringVarP(&c.MqttHost, flagMqttHost, "", "", fmt.Sprintf("HASSIO: mqtt host."))
		viper.SetDefault(flagMqttHost, "")
		cmd.PersistentFlags().StringVarP(&c.MqttPort, flagMqttPort, "", "", fmt.Sprintf("HASSIO: mqtt port."))
		viper.SetDefault(flagMqttPort, "")
	}
}

func (ca *Cmds) MqttArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		cmdLog.LogPrintDate("Connecting to MQTT HASSIO Service...\n")
		ca.Mqtt.Mqtt = mmHa.New(mmHa.Mqtt {
			ClientId: "GoSungrow",
			EntityPrefix: "GoSungrow",
			Username: ca.Mqtt.MqttUsername,
			Password: ca.Mqtt.MqttPassword,
			Host:     ca.Mqtt.MqttHost,
			Port:     ca.Mqtt.MqttPort,
		})
		ca.Error = ca.Mqtt.Mqtt.GetError()
		if ca.Error != nil {
			break
		}

		cmdLog.LogPrintDate("Connecting to SunGrow...\n")
		ca.Mqtt.Mqtt.SungrowDevices, ca.Error = ca.Api.SunGrow.GetDevices(true)
		if ca.Error != nil {
			break
		}
		cmdLog.LogPrintDate("Found SunGrow %d devices\n", len(ca.Mqtt.Mqtt.SungrowDevices))

		ca.Mqtt.Mqtt.DeviceName = "GoSungrow"
		ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig(
			ca.Mqtt.Mqtt.DeviceName,
			ca.Mqtt.Mqtt.DeviceName,
			"virtual",
			"virtual",
			"",
			"",
			"Roof",
		)
		if ca.Error != nil {
			break
		}

		ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig(
			ca.Mqtt.Mqtt.DeviceName,
			ca.Mqtt.Mqtt.DeviceName,
			"system",
			"system",
			"",
			"",
			"Roof",
		)
		if ca.Error != nil {
			break
		}

		for _, psId := range ca.Mqtt.Mqtt.SungrowDevices {
			// ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig("GoSungrow", strconv.FormatInt(id, 10), "GoSungrow", model[0], "Sungrow", "Roof")
			parent := psId.PsId.String()
			if parent == psId.PsKey.Value() {
				parent = ca.Mqtt.Mqtt.DeviceName
			}
			ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig(
				"GoSungrow",
				parent,
				psId.PsKey.Value(),
				psId.DeviceName.Value(),
				psId.DeviceModel.Value(),
				psId.Vendor.Value(),
				"Roof",
				)
			if ca.Error != nil {
				break
			}
			ca.Mqtt.Mqtt.SungrowPsIds[psId.PsId] = true
		}

		ca.Error = ca.Mqtt.Mqtt.Connect()
		if ca.Error != nil {
			break
		}

		// if c.Mqtt.PsId == 0 {
		// 	c.Mqtt.PsId, c.Error = c.Api.SunGrow.GetPsId()
		// 	if c.Error != nil {
		// 		break
		// 	}
		// 	cmdLog.LogPrintDate("Found SunGrow device %d\n", c.Mqtt.PsId)
		// }
	}

	return ca.Error
}


func (c *CmdMqtt) CmdMqtt(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func (ca *Cmds) CmdMqttRun(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// switch1 := mmMqtt.BinarySensor {
		// 	Device: mmMqtt.Device {
		// 		Connections:  [][]string{{"sungrow_address", "0"}},
		// 		Identifiers:  []string{"sungrow_bin_sensor_0"},
		// 		Manufacturer: "MickMake",
		// 		Model:        "SunGrow inverter",
		// 		Name:         "SunGrow inverter online",
		// 		SwVersion:    "GoSungrow https://github.com/MickMake/GoSungrow",
		// 		ViaDevice:    "GoSungrow",
		// 	},
		// 	Name:         "SunGrow inverter online",
		// 	StateTopic:   "homeassistant/binary_sensor/GoSungrow_0/state",
		// 	UniqueId:     "sungrow_bin_sensor_0",
		// }
		// err = foo.Publish("homeassistant/binary_sensor/GoSungrow_0/config", 0, true, switch1.Json())
		// if err != nil {
		// 	break
		// }
		// err = foo.Publish("homeassistant/binary_sensor/GoSungrow_0/state", 0, true, "OFF")
		// if err != nil {
		// 	break
		// }

		ca.Error = ca.MqttCron()
		if ca.Error != nil {
			break
		}

		cmdLog.LogPrintDate("Starting ticker...\n")
		updateCounter := 0
		timer := time.NewTicker(60 * time.Second)
		for t := range timer.C {
			if updateCounter < 5 {
				updateCounter++
				cmdLog.LogPrintDate("Sleeping: %d\n", updateCounter)
				continue
			}

			updateCounter = 0
			cmdLog.LogPrintDate("Update: %s\n", t.String())
			ca.Error = ca.MqttCron()
			if ca.Error != nil {
				break
			}
		}
	}

	return ca.Error
}

func (ca *Cmds) CmdMqttSync(_ *cobra.Command, args []string) error {

	for range Only.Once {
		// */1 * * * * /dir/command args args
		cronString := "*/5 * * * *"
		if len(args) > 0 {
			cronString = strings.Join(args[0:5], " ")
			cronString = strings.ReplaceAll(cronString, ".", "*")
		}

		cron := gocron.NewScheduler(time.UTC)
		cron = cron.Cron(cronString)
		cron = cron.SingletonMode()

		ca.Error = ca.MqttCron()
		if ca.Error != nil {
			break
		}

		var job *gocron.Job
		job, ca.Error = cron.Do(ca.MqttCron)
		if ca.Error != nil {
			break
		}
		job.IsRunning()

		cmdLog.LogPrintDate("Created job schedule using '%s'\n", cronString)
		cron.StartBlocking()
		if ca.Error != nil {
			break
		}
	}

	return ca.Error
}

func (ca *Cmds) MqttCron() error {
	for range Only.Once {
		if ca.Mqtt == nil {
			ca.Error = errors.New("mqtt not available")
			break
		}

		if ca.Api.SunGrow == nil {
			ca.Error = errors.New("sungrow not available")
			break
		}

		if ca.Mqtt.Mqtt.IsFirstRun() {
			ca.Mqtt.Mqtt.UnsetFirstRun()
		} else {
			time.Sleep(time.Second * 40)	// Takes up to 40 seconds for data to come in.
		}

		newDay := false
		if ca.Mqtt.Mqtt.IsNewDay() {
			newDay = true
		}

		data := ca.Api.SunGrow.NewSunGrowData()
		data.SetPsIds()

		// All := []string{ "queryDeviceList", "getPsList", "getPsDetailWithPsType", "getPsDetail" }
		All := []string{ "queryDeviceList" }
		data.SetEndpoints(All...)
		ca.Error = data.GetData()
		if ca.Error != nil {
			break
		}

		for _, result := range data.GetResults() {
			ca.Error = ca.Update(string(result.EndPointName), result.Response.Data, newDay)
			if ca.Error != nil {
				break
			}
		}

		ca.Mqtt.Mqtt.LastRefresh = time.Now()
	}

	if ca.Error != nil {
		cmdLog.LogPrintDate("Error: %s\n", ca.Error)
	}
	return ca.Error
}

func (ca *Cmds) Update(endpoint string, data api.DataMap, newDay bool) error {
	for range Only.Once {
		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint,
		cmdLog.LogPrintDate("Syncing %d entries with HASSIO from %s.\n", len(data.Map), endpoint)

		for o := range data.Map {
			entries := data.Map[o]
			r := entries.GetEntry(api.LastEntry) // Gets the last entry
			if !r.Point.Valid {
				fmt.Printf("\nInvalid: %v\n", r)
				continue
			}

			fullId := r.FullId()
			if r.Point.GroupName == "alias" {
				fullId = mmHa.JoinStringsForId(r.Parent.Key, r.Point.Parents.Index[0], r.Point.Id.String())
			}

			re := mmHa.EntityConfig {
				Name:        mmHa.JoinStringsForName(" - ", fullId),	// r.Point.Name, // PointName,
				SubName:     "",
				ParentId:    r.EndPoint,
				ParentName:  r.Parent.Key,
				UniqueId:    r.Point.Id.String(),
				// UniqueId:    r.Id,
				FullId:      fullId,	// string(r.FullId),	// WAS r.Point.FullId
				// FullName:    r.Point.Name,
				Units:       r.Point.Unit,
				ValueName:   r.Point.Description,
				// ValueName:   r.Id,
				DeviceClass: "",
				StateClass:  r.Point.UpdateFreq,
				Value:       r.Value.String(),

				// Icon:                   "",
				// ValueTemplate:          "",
				// LastReset:              "",
				// LastResetValueTemplate: "",
			}

			if strings.Contains(fullId, "device_type_count") {
				fmt.Printf("")
			}
			if newDay {
				fmt.Printf("C")
				ca.Error = ca.Mqtt.Mqtt.BinarySensorPublishConfig(re)
				if ca.Error != nil {
					break
				}

				ca.Error = ca.Mqtt.Mqtt.SensorPublishConfig(re)
				if ca.Error != nil {
					break
				}
			}

			fmt.Printf("U")
			ca.Error = ca.Mqtt.Mqtt.BinarySensorPublishValue(re)
			if ca.Error != nil {
				break
			}

			ca.Error = ca.Mqtt.Mqtt.SensorPublishValue(re)
			if ca.Error != nil {
				break
			}
		}
		fmt.Println()
	}

	if ca.Error != nil {
		cmdLog.LogPrintDate("Error: %s\n", ca.Error)
	}
	return ca.Error
}
