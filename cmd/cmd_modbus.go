package cmd

import (
	"fmt"
	"time"

	"github.com/MickMake/GoSungrow/cmdModbus"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/MickMake/GoUnify/cmdLog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// https://github.com/simonvetter/modbus
// https://github.com/goburrow/modbus
// -> fork -> https://github.com/activeshadow/modbus
// https://github.com/btfak/modbus
// -> fork -> https://github.com/dpapathanasiou/go-modbus

const (
	flagModbusUsername = "modbus-user"
	flagModbusPassword = "modbus-password"
	flagModbusHost     = "modbus-host"
	flagModbusPort     = "modbus-port"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdModbus struct {
	CmdDefault

	Username string
	Password string
	Host     string
	Port     string

	Client              cmdModbus.Modbus
	log                 cmdLog.Log
	optionSleepDelay    time.Duration
	optionFetchSchedule time.Duration
	// optionCacheTimeout  time.Duration
}

func NewCmdModbus(logLevel string) *CmdModbus {
	var ret *CmdModbus

	for range Only.Once {
		if logLevel == "" {
			logLevel = cmdLog.LogLevelInfoStr
		}

		ret = &CmdModbus{
			CmdDefault: CmdDefault{
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},

			log:                 cmdLog.New(logLevel),
			optionSleepDelay:    time.Second * 40, // Takes up to 40 seconds for data to come in.
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
		var cmdRoot = &cobra.Command{
			Use:                   "modbus",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "ModBus"},
			Short:                 "Connect directly to a Sungrow inverter via Modbus.",
			Long:                  "Connect directly to a Sungrow inverter via Modbus.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			RunE:                  c.CmdModbus,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdRoot)
		cmdRoot.Example = cmdHelp.PrintExamples(cmdRoot, "run", "sync")

		// ******************************************************************************** //
		var cmdRootGet = &cobra.Command{
			Use:                   "get <address> [quantity] [type]",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "MQTT"},
			Short:                 "Get value from a Modbus register.",
			Long:                  "Get value from a Modbus register.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				// cmds.Error = cmds.SunGrowArgs(cmd, args)
				// if cmds.Error != nil {
				// 	return cmds.Error
				// }
				cmds.Error = cmds.Modbus.ModbusArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: cmds.Modbus.CmdModbusGet,
			Args: cobra.RangeArgs(1, 3),
		}
		cmdRoot.AddCommand(cmdRootGet)
		cmdRootGet.Example = cmdHelp.PrintExamples(cmdRootGet, "")

		// ******************************************************************************** //
		var cmdRootScan = &cobra.Command{
			Use:                   "scan <start address> <end address> [type]",
			Aliases:               []string{""},
			Annotations:           map[string]string{"group": "MQTT"},
			Short:                 "Scan value from a Modbus register.",
			Long:                  "Scan value from a Modbus register.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE: func(cmd *cobra.Command, args []string) error {
				// cmds.Error = cmds.SunGrowArgs(cmd, args)
				// if cmds.Error != nil {
				// 	return cmds.Error
				// }
				cmds.Error = cmds.Modbus.ModbusArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: cmds.Modbus.CmdModbusScan,
			Args: cobra.RangeArgs(2, 3),
		}
		cmdRoot.AddCommand(cmdRootScan)
		cmdRootScan.Example = cmdHelp.PrintExamples(cmdRootScan, "")
	}
	return c.SelfCmd
}

func (c *CmdModbus) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		cmd.PersistentFlags().StringVarP(&c.Username, flagModbusUsername, "", "", "Modbus username.")
		viper.SetDefault(flagModbusUsername, "")
		cmd.PersistentFlags().StringVarP(&c.Password, flagModbusPassword, "", "", "Modbus password.")
		viper.SetDefault(flagModbusPassword, "")
		cmd.PersistentFlags().StringVarP(&c.Host, flagModbusHost, "", "", "Modbus host.")
		viper.SetDefault(flagModbusHost, "")
		cmd.PersistentFlags().StringVarP(&c.Port, flagModbusPort, "", cmdModbus.DefaultPort, "Modbus port.")
		viper.SetDefault(flagModbusPort, cmdModbus.DefaultPort)
	}
}

func (c *CmdModbus) ModbusArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		c.Client = cmdModbus.New(cmdModbus.Modbus{
			ClientId: DefaultServiceName,
			Username: c.Username,
			Password: c.Password,
			Host:     c.Host,
			Port:     c.Port,
			LogLevel: c.log.GetLogLevel(),
		})
		c.Error = c.Client.GetError()
		if c.Error != nil {
			break
		}

		fmt.Println("WARNING! This feature is still in alpha state.")
		c.log.Info("Connecting to Modbus...\n")
		c.Error = c.Client.Connect()
		if c.Error != nil {
			break
		}

		c.Client.SetLog(c.log)

		// c.log.Info("Connecting to SunGrow...\n")
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

func (c *CmdModbus) CmdModbusGet(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(3, args)

		var address cmdModbus.Address
		address, c.Error = cmdModbus.StringToAddress(args[0])
		if c.Error != nil {
			break
		}

		var quantity cmdModbus.Quantity
		quantity, c.Error = cmdModbus.StringToQuantity(args[1])
		if c.Error != nil {
			quantity = 1
			c.Error = nil
		}
		if quantity == 0 {
			quantity = 1
		}

		c.log.Info("Get Modbus value from %v\n", address)

		fmt.Println("################################################################################")
		fmt.Println("# Read Input Register #")
		retI := c.Client.Read(address, quantity, args[2])
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", retI)

		fmt.Println("################################################################################")
		fmt.Println("# Read Holding Register #")
		retI = c.Client.ReadHolding(address, quantity, args[2])
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", retI)

		fmt.Println("################################################################################")
		fmt.Println("# Read Coil #")
		retB := c.Client.ReadBool(address, quantity)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", retB)

		fmt.Println("################################################################################")
		fmt.Println("# Read Discrete Input #")
		retB = c.Client.ReadDiscreteInput(address, quantity)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", retB)

	}

	return c.Error
}

func (c *CmdModbus) CmdModbusScan(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(3, args)

		var start cmdModbus.Address
		start, c.Error = cmdModbus.StringToAddress(args[0])
		if c.Error != nil {
			break
		}

		var end cmdModbus.Address
		end, c.Error = cmdModbus.StringToAddress(args[1])
		if c.Error != nil {
			c.Error = nil
			end = 0
		}
		if end == 0 {
			end = start + 100
		}

		size := args[2]
		if size == "" {
			size = cmdModbus.TypeUnsigned8Bit
		}

		c.log.Info("Scanning Modbus addresses from %v to %v\n", start, end)

		limit := cmdModbus.Quantity(100) // 123
		var quantity cmdModbus.Quantity

		fmt.Println("################################################################################")
		fmt.Println("# Read Input Register #")
		for address := start; address < end; {
			quantity = cmdModbus.Quantity(end - address)
			if quantity > limit {
				quantity = limit
			}

			ret := c.Client.Read(address, quantity, size)
			if c.Client.IsError() {
				address++
				continue
			}
			fmt.Printf("# Address[%d]: #\n%s", address, ret)
			address += cmdModbus.Address(quantity)
		}

		fmt.Println("################################################################################")
		fmt.Println("# Read Holding Register #")
		for address := start; address < end; {
			quantity = cmdModbus.Quantity(end - start)
			if quantity > limit {
				quantity = limit
			}

			ret := c.Client.ReadHolding(address, quantity, size)
			if c.Client.IsError() {
				address++
				continue
			}
			fmt.Printf("# Address[%d]: #\n%s", address, ret)
			address += cmdModbus.Address(quantity)
		}

		// fmt.Println("################################################################################")
		// fmt.Println("# Read Coil #")
		// for address := start; address < end; {
		// 	quantity = cmdModbus.Quantity(end - start)
		// 	if quantity > limit {
		// 		quantity = limit
		// 	}
		//
		// 	ret := c.Client.ReadBool(address, quantity)
		// 	if c.Client.IsError() {
		// 		address++
		// 		continue
		// 	}
		// 	fmt.Printf("# Address[%d]: #\n%s", address, ret)
		// 	address += cmdModbus.Address(quantity)
		// }
		//
		// fmt.Println("################################################################################")
		// fmt.Println("# Read Discrete Input #")
		// for address := start; address < end; {
		// 	quantity = cmdModbus.Quantity(end - start)
		// 	if quantity > limit {
		// 		quantity = limit
		// 	}
		//
		// 	ret := c.Client.ReadDiscreteInput(address, quantity)
		// 	if c.Client.IsError() {
		// 		address++
		// 		continue
		// 	}
		// 	fmt.Printf("# Address[%d]: #\n%s", address, ret)
		// 	address += cmdModbus.Address(quantity)
		// }
	}

	return c.Error
}
