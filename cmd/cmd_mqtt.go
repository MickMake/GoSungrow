package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/mmMqtt"
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)


func AttachCmdMqtt(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdMqtt = &cobra.Command{
		Use:                   "mqtt",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("All things MQTT related."),
		Long:                  fmt.Sprintf("All things MQTT related."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		RunE:                  cmdMqttFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmd.AddCommand(cmdMqtt)
	cmdMqtt.Example = PrintExamples(cmdMqtt, "sync", "sync all")


	// ******************************************************************************** //
	var cmdMqttSync = &cobra.Command{
		Use:                   "update",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Sync to an MQTT broker."),
		Long:                  fmt.Sprintf("Sync to an MQTT broker."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdMqttSyncFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdMqtt.AddCommand(cmdMqttSync)
	cmdMqttSync.Example = PrintExamples(cmdMqttSync, "", "all")

	return cmdMqtt
}


func cmdMqttFunc(cmd *cobra.Command, args []string) error {
	var err error

	for range Only.Once {
		foo := mmMqtt.New(mmMqtt.Mqtt{
			ClientId: "SunGrow",
			Username: "mickmake",
			Password: "rvsrzdd0",
			Host:     "10.0.5.21",
			Port:     "11883",
		})
		err = foo.GetError()
		if err != nil {
			break
		}

		err = foo.Connect()
		if err != nil {
			break
		}


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

		err = Cmd.SunGrowArgs(cmd, args)
		if err != nil {
			break
		}

		var psId int64
		psId, err = Cmd.SunGrow.GetPsId()
		if err != nil {
			break
		}

		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint, 
		ep := Cmd.SunGrow.QueryDevice(psId)
		if ep.IsError() {
			err = ep.GetError()
			break
		}

		data := ep.GetData()
		for i, r := range data.Entries {
			fmt.Printf("%s ", r.PointId)
			err = foo.SensorPublishConfig(r.PointId, r.PointName, r.Unit, i)
			if err != nil {
				break
			}
			err = foo.SensorPublishState(r.PointId, r.Value)
			if err != nil {
				break
			}
		}
		fmt.Println()
		if err != nil {
			break
		}

		updateCounter := 0
		timer := time.NewTicker(10 * time.Second)
		for t := range timer.C {
			if updateCounter < 6 {
				updateCounter++
				fmt.Printf("Wait: %d - %s\n", updateCounter, t.String())
				continue
			}

			updateCounter = 0
			fmt.Printf("Update: %s\n", t.String())
			ep = Cmd.SunGrow.QueryDevice(psId)
			if ep.IsError() {
				err = ep.GetError()
				break
			}

			data = ep.GetData()
			for _, r := range data.Entries {
				fmt.Printf("%s ", r.PointId)
				err = foo.SensorPublishState(r.PointId, r.Value)
				if err != nil {
					break
				}
			}
			fmt.Println()
		}
		if err != nil {
			break
		}


		// switch {
		// 	case len(args) == 0:
		// 		Cmd.Error = cmd.Help()
		//
		// 	case args[0] == "all":
		// 		// Cmd.Error = Cmd.GoogleUpdate()
		//
		// 	default:
		// 		fmt.Println("Unknown sub-command.")
		// 		_ = cmd.Help()
		// }
	}

	return err
}

func toggle(v string) string {
	switch v {
		case "OFF":
			v = "ON"
		case "ON":
			v = "OFF"
	}
	return v
}

func randoPercent() string {
	t := time.Now()
	min := 0
	max := t.Second()
	i := (rand.Intn(max - min) + min) * t.Minute()	// / float64(max)
	return fmt.Sprintf("%.2f", (float64(i) / 3600) * 100)
}

func randoKWh() string {
	t := time.Now()
	min := 0
	max := t.Minute()
	i := (rand.Intn(max - min) + min) * t.Second()	// / float64(max)
	return fmt.Sprintf("%.2f", (float64(i) / 3600) * 11000)
}

func cmdMqttSyncFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				Cmd.Error = cmd.Help()

			case args[0] == "all":
				// Cmd.Error = Cmd.GoogleUpdate()

			default:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()
		}
	}
}
