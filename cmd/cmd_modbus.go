package cmd

import (
	"GoSungrow/cmdModbus"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// https://github.com/simonvetter/modbus
// https://github.com/goburrow/modbus
	// -> fork -> https://github.com/activeshadow/modbus
// https://github.com/btfak/modbus
	// -> fork -> https://github.com/dpapathanasiou/go-modbus


//goland:noinspection GoNameStartsWithPackageName
type CmdModbus struct {
	CmdDefault

	// Modbus client
	ModbusUsername   string
	ModbusPassword   string
	ModbusHost       string
	ModbusPort       string

	Client         *cmdModbus.ModBus
	// points         getDevicePointAttrs.PointsMap
	// previous       map[string]*api.DataEntries

	optionLogLevel      int
	optionSleepDelay    time.Duration
	optionFetchSchedule time.Duration
	// optionCacheTimeout  time.Duration
}

func NewCmdModbus() *CmdModbus {
	var ret *CmdModbus

	for range Only.Once {
		ret = &CmdModbus {
			CmdDefault: CmdDefault {
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},

			optionLogLevel:      LogLevelInfo,
			optionSleepDelay:    time.Second * 40,		// Takes up to 40 seconds for data to come in.
			optionFetchSchedule: time.Minute * 5,
			// previous:            make(map[string]*api.DataEntries, 0),
		}
	}

	return ret
}

func (c *CmdModbus) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		var cmdModbus = &cobra.Command{
			Use:                   "modbus",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "ModBus"},
			Short:                 "Connect directly to a Sungrow inverter.",
			Long:                  "Connect directly to a Sungrow inverter.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  c.CmdModbus,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdModbus)
		cmdModbus.Example = cmdHelp.PrintExamples(cmdModbus, "run", "sync")
	}
	return c.SelfCmd
}

func (c *CmdModbus) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		// cmd.PersistentFlags().StringVarP(&c.DbusUsername, flagDbusUsername, "", "", "HASSIO: mqtt username.")
		// viper.SetDefault(flagDbusUsername, "")
		// cmd.PersistentFlags().StringVarP(&c.DbusPassword, flagDbusPassword, "", "", "HASSIO: mqtt password.")
		// viper.SetDefault(flagDbusPassword, "")
		// cmd.PersistentFlags().StringVarP(&c.DbusHost, flagDbusHost, "", "", "HASSIO: mqtt host.")
		// viper.SetDefault(flagDbusHost, "")
		// cmd.PersistentFlags().StringVarP(&c.DbusPort, flagDbusPort, "", "", "HASSIO: mqtt port.")
		// viper.SetDefault(flagDbusPort, "")
	}
}

func (c *CmdModbus) ModbusArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// c.LogInfo("Connecting to MQTT HASSIO Service...\n")
		// c.Client = mmHa.New(mmHa.Dbus {
		// 	ClientId:     DefaultServiceName,
		// 	EntityPrefix: DefaultServiceName,
		// 	Username:     c.DbusUsername,
		// 	Password:     c.DbusPassword,
		// 	Host:         c.DbusHost,
		// 	Port:         c.DbusPort,
		// })
		// c.Error = c.Client.GetError()
		// if c.Error != nil {
		// 	break
		// }
		//
		// c.LogInfo("Connecting to SunGrow...\n")
		// c.Client.SungrowDevices, c.Error = cmds.Api.SunGrow.GetDeviceList()
		// if c.Error != nil {
		// 	break
		// }
	}

	return c.Error
}

func (c *CmdModbus) CmdModbus(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
