package cmd

import (
	"GoSungrow/iSolarCloud/WebAppService/getDevicePointAttrs"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"GoSungrow/mmHa"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/MickMake/GoUnify/cmdLog"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)


const (
	DefaultServiceName = "GoSungrow"
	DefaultServiceArea = "Roof"
	DefaultVendor      = "MickMake"
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

	Client         *mmHa.Mqtt
	endpoints      MqttEndPoints
	points         getDevicePointAttrs.PointsMap
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
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.Mqtt.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  cmds.Mqtt.CmdMqttRun,
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
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.Mqtt.MqttArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  cmds.Mqtt.CmdMqttSync,
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

func (c *CmdMqtt) MqttArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		cmdLog.LogPrintDate("Connecting to MQTT HASSIO Service...\n")
		c.Client = mmHa.New(mmHa.Mqtt {
			ClientId:     DefaultServiceName,
			EntityPrefix: DefaultServiceName,
			Username:     c.MqttUsername,
			Password:     c.MqttPassword,
			Host:         c.MqttHost,
			Port:         c.MqttPort,
		})
		c.Error = c.Client.GetError()
		if c.Error != nil {
			break
		}

		cmdLog.LogPrintDate("Connecting to SunGrow...\n")
		c.Client.SungrowDevices, c.Error = cmds.Api.SunGrow.GetDeviceList()
		// ca.Mqtt.Client.SungrowDevices, ca.Error = ca.Api.SunGrow.GetPsKeys()
		// ca.Mqtt.Client.SungrowDevices, ca.Error = ca.Api.SunGrow.GetPsIds()
		// ca.Mqtt.Client.SungrowDevices, ca.Error = ca.Api.SunGrow.GetPsTreeMenu()
		if c.Error != nil {
			break
		}

		cmdLog.LogPrintDate("Found SunGrow %d devices\n", len(c.Client.SungrowDevices))
		c.Client.DeviceName = DefaultServiceName
		c.Error = c.Client.SetDeviceConfig(
			c.Client.DeviceName,
			c.Client.DeviceName,
			"virtual",
			"virtual",
			"",
			"",
			DefaultServiceArea,
		)
		if c.Error != nil {
			break
		}

		c.Error = c.Client.SetDeviceConfig(
			c.Client.DeviceName,
			c.Client.DeviceName,
			"system",
			"system",
			"",
			"",
			DefaultServiceArea,
		)
		if c.Error != nil {
			break
		}

		for _, psId := range c.Client.SungrowDevices {
			// ca.Error = ca.Mqtt.Mqtt.SetDeviceConfig(DefaultServiceName, strconv.FormatInt(id, 10), DefaultServiceName, model[0], "Sungrow", DefaultServiceArea)
			parent := psId.PsId.String()
			if parent == psId.PsKey.Value() {
				parent = c.Client.DeviceName
			}

			c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName,
				parent,
				psId.PsKey.Value(),
				psId.DeviceName.Value(),
				psId.DeviceModel.Value(),
				psId.FactoryName.Value(),
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}

			c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName,
				DefaultServiceName,
				psId.PsId.String(),
				psId.FactoryName.Value(),
				psId.FactoryName.Value(),
				psId.FactoryName.Value(),
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}

			c.Client.SungrowPsIds[psId.PsId] = true
		}

		c.Error = c.Client.Connect()
		if c.Error != nil {
			break
		}

		cmdLog.LogPrintDate("Caching Sungrow metadata...\n")
		c.Error = c.GetEndPoints()
		if c.Error != nil {
			break
		}

		c.points, c.Error = cmds.Api.SunGrow.DevicePointAttrsMap(nil, "")
		if c.Error != nil {
			break
		}
		cmdLog.LogPrintDate("Cached %d Sungrow data points...\n", len(c.points))
	}

	return c.Error
}

func (c *CmdMqtt) CmdMqtt(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func (c *CmdMqtt) CmdMqttRun(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// switch1 := mmMqtt.BinarySensor {
		// 	Device: mmMqtt.Device {
		// 		Connections:  [][]string{{"sungrow_address", "0"}},
		// 		Identifiers:  []string{"sungrow_bin_sensor_0"},
		// 		Manufacturer: DefaultVendor,
		// 		Model:        "SunGrow inverter",
		// 		Name:         "SunGrow inverter online",
		// 		SwVersion:    "GoSungrow https://github.com/MickMake/GoSungrow",
		// 		ViaDevice:    DefaultServiceName,
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

		c.Error = c.Cron()
		if c.Error != nil {
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
			c.Error = c.Cron()
			if c.Error != nil {
				break
			}
		}
	}

	return c.Error
}

func (c *CmdMqtt) CmdMqttSync(_ *cobra.Command, args []string) error {
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

		c.Error = c.Cron()
		if c.Error != nil {
			break
		}

		var job *gocron.Job
		job, c.Error = cron.Do(c.Cron)
		if c.Error != nil {
			break
		}
		job.IsRunning()

		cmdLog.LogPrintDate("Created job schedule using '%s'\n", cronString)
		cron.StartBlocking()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}


func (c *CmdMqtt) Cron() error {
	for range Only.Once {
		if c == nil {
			c.Error = errors.New("mqtt not available")
			break
		}

		if cmds.Api.SunGrow == nil {
			c.Error = errors.New("sungrow not available")
			break
		}

		if c.Client.IsFirstRun() {
			c.Client.UnsetFirstRun()
		} else {
			time.Sleep(time.Second * 40)	// Takes up to 40 seconds for data to come in.
		}

		newDay := false
		if c.Client.IsNewDay() {
			newDay = true
		}

		data := cmds.Api.SunGrow.NewSunGrowData()
		data.SetPsIds()
		if data.Error != nil {
			c.Error = data.Error
			break
		}

		data.SetEndpoints(c.endpoints.Names()...)
		c.Error = data.GetData()
		if c.Error != nil {
			break
		}

		for _, result := range data.Results {
			c.Error = c.Update(result.EndPointName.String(), result.Response.Data, newDay)
			if c.Error != nil {
				break
			}
		}

		c.Client.LastRefresh = time.Now()
	}

	if c.Error != nil {
		cmdLog.LogPrintDate("Error: %s\n", c.Error)
	}
	return c.Error
}

func (c *CmdMqtt) Update(endpoint string, data api.DataMap, newDay bool) error {
	for range Only.Once {
		// Also getPowerStatistics, getHouseholdStoragePsReport, getPsList, getUpTimePoint,
		cmdLog.LogPrintDate("Syncing %d entries with HASSIO from %s.\n", len(data.Map), endpoint)

		for o := range data.Map {
			entries := data.Map[o]
			r := entries.GetEntry(api.LastEntry) // Gets the last entry

			if !r.Point.Valid {
				// cmdLog.LogPrintDate("\n[%s] - invalid value - %s ...\n", r.Current.FieldPath.String(), r.Value.String())
				fmt.Printf("?")
				continue
			}

			if !c.endpoints.IsOK(r) {
				continue
			}

			_ = c.UpdatePoint(r)
			r.Value.UnitValueFix()	// @TODO - Fix this up properly

			var id string
			var name string
			switch {
				case r.Point.GroupName == "alias":
					id = mmHa.JoinStringsForId(r.Parent.Key, r.Point.Parents.Index[0], r.Point.Id)
					name = mmHa.JoinStringsForName(" - ", r.Parent.Key, r.Point.Parents.Index[0], r.Point.Id)
				case r.Point.GroupName != "":
					id = r.EndPoint
					name = mmHa.JoinStringsForName(" - ", r.Parent.Key, r.Point.Id, r.Point.GroupName, r.Point.Description)
				default:
					id = r.EndPoint
					name = r.EndPoint
			}

			if r.Point.Unit == "" {
				r.Point.Unit = r.Point.ValueType
			}
			if r.Point.Unit == "Bool" {
				r.Point.Unit = mmHa.LabelBinarySensor
			}
			if r.Point.ValueType == "Bool" {
				r.Point.Unit = mmHa.LabelBinarySensor
			}

			re := mmHa.EntityConfig {
				Name:        name,	// mmHa.JoinStringsForName(" - ", id), // r.Point.Name, // PointName,
				SubName:     "",
				ParentId:    r.EndPoint,
				ParentName:  r.Parent.Key,
				UniqueId:    r.Point.Id,
				// UniqueId:    r.Id,
				FullId: id, // string(r.FullId),	// WAS r.Point.FullId
				// FullName:    r.Point.Name,
				Units:       r.Point.Unit,
				ValueName:   r.Point.Description,
				// ValueName:   r.Id,
				DeviceClass: "",
				StateClass:  r.Point.UpdateFreq,
				Value:       r.Value,
				Point:       r.Point,

				LastReset:              r.Point.WhenReset(),
				// LastResetValueTemplate: "",
			}

			// if strings.Contains(r.EndPoint, "13149") || strings.Contains(r.Current.DataStructure.Endpoint.String(), "13149") ||
			// 	strings.Contains(r.EndPoint, "13119") || strings.Contains(r.Current.DataStructure.Endpoint.String(), "13119") {
			// 	fmt.Sprintf("")
			// }

			switch {
				case r.Point.IsTotal():
					re.StateClass = "total"
				default:
					re.StateClass = "measurement"
			}

			if newDay {
				fmt.Printf("C")
				c.Error = c.Client.BinarySensorPublishConfig(re)
				if c.Error != nil {
					break
				}

				c.Error = c.Client.SensorPublishConfig(re)
				if c.Error != nil {
					break
				}
			}

			// if strings.Contains(r.EndPoint, "p13115") || strings.Contains(r.Current.DataStructure.Endpoint.String(), "p13115") {
			// 	fmt.Sprintf("")
			// }

			fmt.Printf("U")
			c.Error = c.Client.BinarySensorPublishValue(re)
			if c.Error != nil {
				break
			}

			c.Error = c.Client.SensorPublishValue(re)
			if c.Error != nil {
				break
			}
		}
		fmt.Println()
	}

	if c.Error != nil {
		cmdLog.LogPrintDate("Error: %s\n", c.Error)
	}
	return c.Error
}

func (c *CmdMqtt) GetEndPoints() error {
	for range Only.Once {
		fn := filepath.Join(cmds.ConfigDir, "mqtt_endpoints.json")
		if !output.FileExists(fn) {
			c.Error = output.PlainFileWrite(fn, []byte(DefaultMqttFile), 0644)
			if c.Error != nil {
				break
			}
		}

		c.Error = output.FileRead(fn, &c.endpoints)
		if c.Error != nil {
			break
		}

		// All := []string{ "queryDeviceList", "getPsList", "getPsDetailWithPsType", "getPsDetail", "getKpiInfo"}
		// All := []string{ "queryDeviceList", "getPsList", "getPsDetailWithPsType", "getPsDetail", "getKpiInfo"}	//, queryMutiPointDataList, getDevicePointMinuteDataList }
		// All := []string{ "WebIscmAppService.getDeviceModel" }
		for name := range c.endpoints {
			c.Error = c.Client.SetDeviceConfig(
				DefaultServiceName,
				DefaultServiceName,
				name,
				DefaultServiceName + "." + name,
			 	DefaultServiceName,
				DefaultVendor,
				DefaultServiceArea,
			)
			if c.Error != nil {
				break
			}
		}
	}
	return c.Error
}

func (c *CmdMqtt) UpdatePoint(entry *api.DataEntry) error {
	for range Only.Once {
		if !c.points.Exists(entry.Point.Id) {
			// fmt.Printf("entry.Point: %s - NOT FOUND\n", entry.Point.Id)
			break
		}

		p := c.points.Get(entry.Point.Id)
		if p == nil {
			// fmt.Printf("entry.Point: %s - FOUND - EMPTY\n", entry.Point.Id)
			break
		}

		// {
		// 	fmt.Printf("entry.Point: %s - FOUND - %v\n", entry.Point.Id, p)
		// 	// fmt.Printf("\tValue   - Description:'-'\t\tUnit:'%s'\tGroupName:'-'\tValueType:'%s'\n",
		// 	// 	r.Value.UnitValue, r.Point.ValueType)
		// 	fmt.Printf("\tDescription:'%s'\tPointName:'%s' - SAME:%t\n",
		// 		entry.Point.Description, entry.Current.DataStructure.PointName, entry.Current.DataStructure.PointName == entry.Point.Description)
		// 	fmt.Printf("\tUnit:'%s'\tPointUnit:'%s' - SAME:%t\n",
		// 		entry.Point.Unit, entry.Current.DataStructure.PointUnit, entry.Current.DataStructure.PointUnit == entry.Point.Unit)
		// 	fmt.Printf("\tGroupName:'%s'\tPointGroupName:'%s' - SAME:%t\n",
		// 		entry.Point.GroupName, entry.Current.DataStructure.PointGroupName, entry.Current.DataStructure.PointGroupName == entry.Point.GroupName)
		// 	fmt.Printf("\tValueType:'%s'\tValueType:'%s' - SAME:%t\n",
		// 		entry.Point.ValueType, entry.Current.DataStructure.ValueType, entry.Current.DataStructure.ValueType == entry.Point.ValueType)
		// }

		// If Point description matches ...
		if p.Name.String() == entry.Point.Description {
			break
		}

		// If Point unit matches ...
		if p.Unit.String() == entry.Point.Unit {
			break
		}

		// If Point group name is set ...
		if entry.Point.GroupName != "" {
			break
		}

		// ... update the Point.
		entry.Point.Description = p.Name.String()
		entry.Point.Unit = p.Unit.String()
		entry.Point.GroupName = p.PointGroupName
		entry.Point.ValueType = p.UnitType.String()

		// if (p.Name.String() != entry.Point.Description) && (p.Unit.String() != entry.Point.Unit) && (entry.Point.GroupName == "") {
		// 	// fmt.Printf("\nNOT SAME\n")
		// 	// fmt.Println("BEFORE:")
		// 	// fmt.Printf("\tName:'%s'\tDescription:'%s'\n", p.Name, entry.Point.Description)
		// 	// fmt.Printf("\tPointGroupId:'%s'\tGroupName:'%s'\n", p.PointGroupId, entry.Point.GroupName)
		// 	// fmt.Printf("\tUnitType:'%s'\tValueType:'%s'\n", p.UnitType, entry.Point.ValueType)
		// 	// fmt.Printf("\tUnit:'%s'\tUnit:'%s'\n", p.Unit, entry.Point.Unit)
		//
		// 	entry.Point.Description = p.Name.String()
		// 	entry.Point.Unit = p.Unit.String()
		// 	entry.Point.GroupName = p.PointGroupName
		// 	entry.Point.ValueType = p.UnitType.String()
		//
		// 	// fmt.Println("AFTER:")
		// 	// fmt.Printf("\tName:'%s'\tDescription:'%s'\n", p.Name, entry.Point.Description)
		// 	// fmt.Printf("\tPointGroupId:'%s'\tGroupName:'%s'\n", p.PointGroupId, entry.Point.GroupName)
		// 	// fmt.Printf("\tUnitType:'%s'\tValueType:'%s'\n", p.UnitType, entry.Point.ValueType)
		// 	// fmt.Printf("\tUnit:'%s'\tUnit:'%s'\n", p.Unit, entry.Point.Unit)
		// 	// fmt.Println("")
		// }

	}
	return c.Error
}


type MqttEndPoints map[string]MqttEndPoint
type MqttEndPoint struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}
func (c *MqttEndPoints) Names() []string {
	var ret []string
	for name := range *c {
		ret = append(ret, name)
	}
	return ret
}

func (c *MqttEndPoints) IsOK(check *api.DataEntry) bool {
	var yes bool
	for range Only.Once {
		field := check.Current.GetFieldPath()
		name := field.First()

		var ep MqttEndPoint
		if ep, yes = (*c)[name]; !yes {
			yes = false
			break
		}

		if len(ep.Include) == 0 {
			yes = false
			break
		}

		for _, reStr := range ep.Exclude {
			reStr = strings.ReplaceAll(reStr, `.`, `\.`)
			reStr = strings.ReplaceAll(reStr, `*`, `.*?`)
			reStr = "^" + strings.TrimPrefix(reStr, "^")
			re := regexp.MustCompile(reStr)
			if re.MatchString(check.EndPoint) {
				yes = false
				break
			}
			if re.MatchString(check.Current.FieldPath.String()) {
				yes = false
				break
			}
			if re.MatchString(check.Current.DataStructure.Endpoint.String()) {
				yes = false
				break
			}
		}

		for _, reStr := range ep.Include {
			reStr = strings.ReplaceAll(reStr, `.`, `\.`)
			reStr = strings.ReplaceAll(reStr, `*`, `.*`)
			reStr = "^" + strings.TrimPrefix(reStr, "^")
			re := regexp.MustCompile(reStr)
			if re.MatchString(check.EndPoint) {
				yes = true
				break
			}
			if re.MatchString(check.Current.FieldPath.String()) {
				yes = true
				break
			}
			if re.MatchString(check.Current.DataStructure.Endpoint.String()) {
				yes = true
				break
			}
		}
	}
	return yes
}

const DefaultMqttFile = `{
	"queryDeviceList": {
		"include": [
			"virtual.*"
		],
		"exclude": [
			"queryDeviceList.*.devices.*",
			"queryDeviceList.*.device_types.*"
		]
	},
	"getPsList": {
		"include": [
			"virtual.*"
		],
		"exclude": [
		]
	},
	"getPsDetail": {
		"include": [
			"virtual.*"
		],
		"exclude": [
		]
	}
}
`
