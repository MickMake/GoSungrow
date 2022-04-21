package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/mmHa"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)


func AttachCmdMqtt(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdMqtt = &cobra.Command{
		Use:                   "mqtt",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Connect to a HASSIO broker."),
		Long:                  fmt.Sprintf("Connect to a HASSIO broker."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.MqttArgs,
		RunE:                  cmdMqttFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(cmdMqtt)
	cmdMqtt.Example = PrintExamples(cmdMqtt, "run", "sync")


	// ******************************************************************************** //
	var cmdMqttRun = &cobra.Command{
		Use:                   "run",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("One-off sync to a HASSIO broker."),
		Long:                  fmt.Sprintf("One-off sync to a HASSIO broker."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.MqttArgs,
		RunE:                  cmdMqttRunFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdMqtt.AddCommand(cmdMqttRun)
	cmdMqttRun.Example = PrintExamples(cmdMqttRun, "")

	// ******************************************************************************** //
	var cmdMqttSync = &cobra.Command{
		Use:                   "sync",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Sync to a HASSIO MQTT broker."),
		Long:                  fmt.Sprintf("Sync to a HASSIO MQTT broker."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.MqttArgs,
		RunE:                  cmdMqttSyncFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdMqtt.AddCommand(cmdMqttSync)
	cmdMqttSync.Example = PrintExamples(cmdMqttSync, "", "all")

	return cmdMqtt
}


func (ca *CommandArgs) MqttArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		ca.Error = ca.ProcessArgs(cmd, args)
		if ca.Error != nil {
			break
		}

		LogPrintDate("Connecting to SunGrow...\n")
		Cmd.Error = Cmd.SunGrowArgs(cmd, args)
		if Cmd.Error != nil {
			break
		}

		var id int64
		id, Cmd.Error = Cmd.SunGrow.GetPsId()
		if Cmd.Error != nil {
			break
		}

		var model string
		model, Cmd.Error = Cmd.SunGrow.GetPsModel()
		if Cmd.Error != nil {
			break
		}

		var serial string
		serial, Cmd.Error = Cmd.SunGrow.GetPsSerial()
		if Cmd.Error != nil {
			break
		}
		LogPrintDate("Found SunGrow device %s id:%d serial:%s\n", model, id, serial)

		LogPrintDate("Connecting to MQTT HASSIO Service...\n")
		Cmd.Mqtt = mmHa.New(mmHa.Mqtt {
			ClientId: "GoSunGrow",
			Username: Cmd.MqttUsername,
			Password: Cmd.MqttPassword,
			Host:     Cmd.MqttHost,
			Port:     Cmd.MqttPort,
		})
		Cmd.Error = Cmd.Mqtt.GetError()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Mqtt.SetDeviceConfig("GoSunGrow", strconv.FormatInt(id, 10), "GoSungrow", model, "Sungrow", "Roof")
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Mqtt.Connect()
		if Cmd.Error != nil {
			break
		}

		// if Cmd.Mqtt.PsId == 0 {
		// 	Cmd.Mqtt.PsId, Cmd.Error = Cmd.SunGrow.GetPsId()
		// 	if Cmd.Error != nil {
		// 		break
		// 	}
		// 	LogPrintDate("Found SunGrow device %d\n", Cmd.Mqtt.PsId)
		// }
	}

	return Cmd.Error
}


func cmdMqttFunc(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func cmdMqttRunFunc(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// switch1 := mmMqtt.BinarySensor {
		// 	Device: mmMqtt.Device {
		// 		Connections:  [][]string{{"sungrow_address", "0"}},
		// 		Identifiers:  []string{"sungrow_bin_sensor_0"},
		// 		Manufacturer: "MickMake",
		// 		Model:        "SunGrow inverter",
		// 		Name:         "SunGrow inverter online",
		// 		SwVersion:    "GoSunGrow https://github.com/MickMake/GoSungrow",
		// 		ViaDevice:    "GoSunGrow",
		// 	},
		// 	Name:         "SunGrow inverter online",
		// 	StateTopic:   "homeassistant/binary_sensor/GoSunGrow_0/state",
		// 	UniqueId:     "sungrow_bin_sensor_0",
		// }
		// err = foo.Publish("homeassistant/binary_sensor/GoSunGrow_0/config", 0, true, switch1.Json())
		// if err != nil {
		// 	break
		// }
		// err = foo.Publish("homeassistant/binary_sensor/GoSunGrow_0/state", 0, true, "OFF")
		// if err != nil {
		// 	break
		// }

		Cmd.Error = MqttCron()
		if Cmd.Error != nil {
			break
		}

		LogPrintDate("Starting ticker...\n")
		updateCounter := 0
		timer := time.NewTicker(60 * time.Second)
		for t := range timer.C {
			if updateCounter < 5 {
				updateCounter++
				LogPrintDate("Sleeping: %d\n", updateCounter)
				continue
			}

			updateCounter = 0
			LogPrintDate("Update: %s\n", t.String())
			Cmd.Error = MqttCron()
			if Cmd.Error != nil {
				break
			}

			// ep = Cmd.SunGrow.QueryDevice(psId)
			// if ep.IsError() {
			// 	Cmd.Error = ep.GetError()
			// 	break
			// }
			//
			// data = ep.GetData()
			// for _, r := range data.Entries {
			// 	// fmt.Printf("%s ", r.PointId)
			// 	Cmd.Error = foo.SensorPublishState(r.PointId, r.Value)
			// 	if err != nil {
			// 		break
			// 	}
			// }
			// // fmt.Println()
		}
	}

	return Cmd.Error
}

func cmdMqttSyncFunc(_ *cobra.Command, args []string) error {

	for range Only.Once {
		// */1 * * * * /dir/command args args
		cronString := "*/5 * * * *"
		if len(args) > 0 {
			cronString = strings.Join(args[0:5], " ")
			cronString = strings.ReplaceAll(cronString, ".", "*")
		}

		Cron.Scheduler = gocron.NewScheduler(time.UTC)
		Cron.Scheduler = Cron.Scheduler.Cron(cronString)
		Cron.Scheduler = Cron.Scheduler.SingletonMode()

		Cmd.Error = MqttCron()
		if Cmd.Error != nil {
			break
		}

		Cron.Job, Cmd.Error = Cron.Scheduler.Do(MqttCron)
		if Cmd.Error != nil {
			break
		}

		LogPrintDate("Created job schedule using '%s'\n", cronString)
		Cron.Scheduler.StartBlocking()
		if Cmd.Error != nil {
			break
		}
	}

	return Cmd.Error
}

func MqttCron() error {
	for range Only.Once {
		if Cmd.Mqtt == nil {
			Cmd.Error = errors.New("mqtt not available")
			break
		}

		if Cmd.SunGrow == nil {
			Cmd.Error = errors.New("sungrow not available")
			break
		}

		if Cmd.Mqtt.IsFirstRun() {
			Cmd.Mqtt.UnsetFirstRun()
		} else {
			time.Sleep(time.Second * 40)	// Takes up to 40 seconds for data to come in.
		}

		newDay := false
		if Cmd.Mqtt.IsNewDay() {
			newDay = true
		}

		Cmd.Error = Update1(newDay)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Update2(newDay)
		if Cmd.Error != nil {
			break
		}

		Cmd.Mqtt.LastRefresh = time.Now()
	}

	if Cmd.Error != nil {
		LogPrintDate("Error: %s\n", Cmd.Error)
	}
	return Cmd.Error
}

func Update1(newDay bool) error {
	for range Only.Once {
		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint,
		ep := Cmd.SunGrow.QueryDevice(Cmd.Mqtt.PsId)
		if ep.IsError() {
			Cmd.Error = ep.GetError()
			break
		}
		data := ep.GetData()

		if newDay {
			LogPrintDate("New day: Configuring %d entries in HASSIO.\n", len(data.Entries))
			for _, r := range data.Entries {
				fmt.Printf("C")
				re := mmHa.EntityConfig {
					Type:        r.ValueType.Type,
					Name:        r.ValueType.Id, // PointName,
					SubName:     "",
					ParentId:    r.ValueType.PsKey,
					ParentName:  "",
					UniqueId:    r.PointId,
					FullName:    r.ValueType.Description,
					Units:       r.Unit,
					ValueName:   r.PointId,
					DeviceClass: "",
					Value:       r.Value,
				}

				Cmd.Error = Cmd.Mqtt.BinarySensorPublishConfig(re)
				if Cmd.Error != nil {
					break
				}

				Cmd.Error = Cmd.Mqtt.SensorPublishConfig(re)
				if Cmd.Error != nil {
					break
				}
			}
			fmt.Println()
		}

		LogPrintDate("Updating %d entries to HASSIO.\n", len(data.Entries))
		for _, r := range data.Entries {
			fmt.Printf("U")
			re := mmHa.EntityConfig {
				Type:        r.ValueType.Type,
				Name:        r.ValueType.Id, // PointName,
				SubName:     "",
				ParentId:    r.ValueType.PsKey,
				ParentName:  "",
				UniqueId:    r.PointId,
				FullName:    r.ValueType.Description,
				Units:       r.Unit,
				ValueName:   r.PointId,
				DeviceClass: "",
				Value:       r.Value,
			}

			Cmd.Error = Cmd.Mqtt.BinarySensorPublishValue(re)
			if Cmd.Error != nil {
				break
			}

			Cmd.Error = Cmd.Mqtt.SensorPublishValue(re)
			if Cmd.Error != nil {
				break
			}
		}
		fmt.Println()
	}

	if Cmd.Error != nil {
		LogPrintDate("Error: %s\n", Cmd.Error)
	}
	return Cmd.Error
}

func Update2(newDay bool) error {
	for range Only.Once {
		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint,
		ep := Cmd.SunGrow.QueryPs(Cmd.Mqtt.PsId)
		if ep.IsError() {
			Cmd.Error = ep.GetError()
			break
		}
		data := ep.GetData()

		if newDay {
			LogPrintDate("New day: Configuring %d entries in HASSIO.\n", len(data.Entries))
			for _, r := range data.Entries {
				fmt.Printf("C")

				re := mmHa.EntityConfig {
					Type:        r.ValueType.Type,
					Name:        r.ValueType.Id, // PointName,
					SubName:     "",
					ParentId:    r.ValueType.PsKey,
					ParentName:  "",
					UniqueId:    r.PointId,
					FullName:    r.ValueType.Description,
					Units:       r.Unit,
					ValueName:   r.PointId,
					DeviceClass: "",
					Value:       r.Value,
				}

				Cmd.Error = Cmd.Mqtt.BinarySensorPublishConfig(re)
				if Cmd.Error != nil {
					break
				}

				Cmd.Error = Cmd.Mqtt.SensorPublishConfig(re)
				if Cmd.Error != nil {
					break
				}
			}
			fmt.Println()
		}

		LogPrintDate("Updating %d entries to HASSIO.\n", len(data.Entries))
		for _, r := range data.Entries {
			fmt.Printf("U")

			re := mmHa.EntityConfig {
				Type:        r.ValueType.Type,
				Name:        r.ValueType.Id, // PointName,
				SubName:     "",
				ParentId:    r.ValueType.PsKey,
				ParentName:  "",
				UniqueId:    r.PointId,
				FullName:    r.ValueType.Description,
				Units:       r.Unit,
				ValueName:   r.PointId,
				DeviceClass: "",
				Value:       r.Value,
			}

			Cmd.Error = Cmd.Mqtt.BinarySensorPublishValue(re)
			if Cmd.Error != nil {
				break
			}

			Cmd.Error = Cmd.Mqtt.SensorPublishValue(re)
			if Cmd.Error != nil {
				break
			}
		}
		fmt.Println()
	}

	if Cmd.Error != nil {
		LogPrintDate("Error: %s\n", Cmd.Error)
	}
	return Cmd.Error
}


// func toggle(v string) string {
// 	switch v {
// 		case "OFF":
// 			v = "ON"
// 		case "ON":
// 			v = "OFF"
// 	}
// 	return v
// }
//
// func randoPercent() string {
// 	t := time.Now()
// 	min := 0
// 	max := t.Second()
// 	i := (rand.Intn(max - min) + min) * t.Minute()	// / float64(max)
// 	return fmt.Sprintf("%.2f", (float64(i) / 3600) * 100)
// }
//
// func randoKWh() string {
// 	t := time.Now()
// 	min := 0
// 	max := t.Minute()
// 	i := (rand.Intn(max - min) + min) * t.Second()	// / float64(max)
// 	return fmt.Sprintf("%.2f", (float64(i) / 3600) * 11000)
// }
