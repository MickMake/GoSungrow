package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud"
	"GoSungrow/iSolarCloud/AppService/login"
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"fmt"
	"github.com/MickMake/GoUnify/cmdConfig"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)


const (
	flagApiUrl        = "host"
	flagApiTimeout    = "timeout"
	flagApiUsername   = "user"
	flagApiPassword   = "password"
	flagApiAppKey     = "appkey"
	flagApiLastLogin  = "token-expiry"
	flagApiOutputType = "out"
	flagApiSaveFile   = "save"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdApi struct {
	CmdDefault

	// iSolarCloud api
	ApiTimeout  time.Duration
	Url      string
	Username string
	Password string
	AppKey      string
	LastLogin   string
	ApiToken  string
	ApiTokenFile string
	OutputType string
	SaveFile   bool

	SunGrow *iSolarCloud.SunGrow
}


func NewCmdApi() *CmdApi {
	var ret *CmdApi

	for range Only.Once {
		ret = &CmdApi {
			CmdDefault: CmdDefault {
				Error:   nil,
				cmd:     nil,
				SelfCmd: nil,
			},
			ApiTimeout:   defaultTimeout,
			Url:          defaultHost,
			Username:     "",
			Password:     "",
			AppKey:       defaultApiAppKey,
			LastLogin:    "",
			ApiToken:     "",
			ApiTokenFile: "",
			OutputType:   "",
			SunGrow:      nil,
		}
	}

	return ret
}

func (c *CmdApi) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		var cmdApi = &cobra.Command{
			Use:                   "api",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Interact with the low-level SunGrow api."),
			Long:                  fmt.Sprintf("Interact with the low-level SunGrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               nil,
			Run:                   c.CmdApi,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(cmdApi)
		cmdApi.Example = cmdHelp.PrintExamples(cmdApi, "get <endpoint>", "put <endpoint>")

		// ******************************************************************************** //
		var cmdApiList = &cobra.Command{
			Use:                   "ls",
			Aliases:               []string{"list"},
			Short:                 fmt.Sprintf("List iSolarCloud api endpoints/areas"),
			Long:                  fmt.Sprintf("List iSolarCloud api endpoints/areas"),
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
				return nil
			},
			Run:                   c.CmdApiList,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdApi.AddCommand(cmdApiList)
		cmdApiList.Example = cmdHelp.PrintExamples(cmdApiList, "", "areas", "endpoints", "<area name>")

		// ******************************************************************************** //
		var cmdApiLogin = &cobra.Command{
			Use:                   "login",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Login to iSolarCloud"),
			Long:                  fmt.Sprintf("Login to iSolarCloud"),
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
				return nil
			},
			Run:                   c.CmdApiLogin,
			Args:                  cobra.MinimumNArgs(0),
		}
		cmdApi.AddCommand(cmdApiLogin)
		cmdApiLogin.Example = cmdHelp.PrintExamples(cmdApiLogin, "")

		// ******************************************************************************** //
		var cmdApiGet = &cobra.Command{
			Use:                   "get",
			Aliases:               []string{output.StringTypeTable},
			Short:                 fmt.Sprintf("Get details from iSolarCloud"),
			Long:                  fmt.Sprintf("Get details from iSolarCloud"),
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
				return nil
			},
			RunE:                  func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = false
				c.SunGrow.OutputType.SetJson()
				return c.CmdApiGet(cmd, args)
			},
			Args:                  cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiGet)
		cmdApiGet.Example = cmdHelp.PrintExamples(cmdApiGet, "[area].<endpoint>")

		// ******************************************************************************** //
		var cmdApiRaw = &cobra.Command{
			Use:                   output.StringTypeRaw,
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Raw details from iSolarCloud"),
			Long:                  fmt.Sprintf("Raw details from iSolarCloud"),
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
				return nil
			},
			RunE:                  func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = false
				c.SunGrow.OutputType.SetRaw()
				return c.CmdApiGet(cmd, args)
			},
			Args:                  cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiRaw)
		cmdApiRaw.Example = cmdHelp.PrintExamples(cmdApiRaw, "[area].<endpoint>")

		// ******************************************************************************** //
		var cmdApiSave = &cobra.Command{
			Use:                   "save",
			Aliases:               []string{},
			Short:                 fmt.Sprintf("Save details from iSolarCloud"),
			Long:                  fmt.Sprintf("Save details from iSolarCloud"),
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
				return nil
			},
			RunE:                  func(cmd *cobra.Command, args []string) error {
				c.SunGrow.SaveAsFile = true
				c.SunGrow.OutputType.SetJson()
				return c.CmdApiGet(cmd, args)
			},
			Args:                  cobra.MinimumNArgs(1),
		}
		cmdApi.AddCommand(cmdApiSave)
		cmdApiSave.Example = cmdHelp.PrintExamples(cmdApiSave, "[area].<endpoint>")

		// ******************************************************************************** //
		var cmdApiPut = &cobra.Command{
			Use:                   "put",
			Aliases:               []string{"write"},
			Short:                 fmt.Sprintf("Put details onto iSolarCloud"),
			Long:                  fmt.Sprintf("Put details onto iSolarCloud"),
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
				return nil
			},
			Run:                   c.CmdApiPut,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmdApi.AddCommand(cmdApiPut)
		cmdApiPut.Example = cmdHelp.PrintExamples(cmdApiPut, "[area].<endpoint> <value>")
	}
	return c.SelfCmd
}

func (c *CmdApi) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
	for range Only.Once {
		cmd.PersistentFlags().StringVarP(&c.Username, flagApiUsername, "u", "", fmt.Sprintf("SunGrow: api username."))
		viper.SetDefault(flagApiUsername, "")
		cmd.PersistentFlags().StringVarP(&c.Password, flagApiPassword, "p", "", fmt.Sprintf("SunGrow: api password."))
		viper.SetDefault(flagApiPassword, "")
		cmd.PersistentFlags().StringVarP(&c.AppKey, flagApiAppKey, "", defaultApiAppKey, fmt.Sprintf("SunGrow: api application key."))
		viper.SetDefault(flagApiAppKey, defaultApiAppKey)
		cmd.PersistentFlags().StringVarP(&c.Url, flagApiUrl, "", defaultHost, fmt.Sprintf("SunGrow: Provider API URL."))
		viper.SetDefault(flagApiUrl, defaultHost)
		// cmd.PersistentFlags().DurationVarP(&c.ApiTimeout, flagApiTimeout, "", defaultTimeout, fmt.Sprintf("SunGrow: API timeout."))
		// viper.SetDefault(flagApiTimeout, defaultTimeout)
		c.ApiTimeout = defaultTimeout
		cmd.PersistentFlags().StringVar(&c.LastLogin, flagApiLastLogin, "", "SunGrow: last login.")
		viper.SetDefault(flagApiLastLogin, "")
		// _ = cmd.PersistentFlags().MarkHidden(flagApiLastLogin)

		cmd.PersistentFlags().StringVarP(&c.OutputType, flagApiOutputType, "o", "", fmt.Sprintf("Output type: 'json', 'raw', 'file'"))
		_ = cmd.PersistentFlags().MarkHidden(flagApiOutputType)
		cmd.PersistentFlags().BoolVarP(&c.SaveFile, flagApiSaveFile, "s", false, "Save output as a file.")
		viper.SetDefault(flagApiSaveFile, false)
	}
}


func (ca *Cmds) SunGrowArgs(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		ca.Api.SunGrow = iSolarCloud.NewSunGro(ca.Api.Url, ca.CacheDir)
		if ca.Api.SunGrow.Error != nil {
			ca.Error = ca.Api.SunGrow.Error
			break
		}

		ca.Error = ca.Api.SunGrow.Init()
		if ca.Error != nil {
			break
		}

		ca.Api.SunGrow.SetOutputType(ca.Api.OutputType)
		ca.Api.SunGrow.SaveAsFile = ca.Api.SaveFile

		if ca.Api.AppKey == "" {
			ca.Api.AppKey = defaultApiAppKey
		}

		ca.Error = ca.Api.SunGrow.Login(login.SunGrowAuth{
			AppKey:       ca.Api.AppKey,
			UserAccount:  ca.Api.Username,
			UserPassword: ca.Api.Password,
			TokenFile:    ca.Api.ApiTokenFile,
			Force:        false,
		})
		if ca.Error != nil {
			break
		}

		if ca.Debug {
			ca.Api.SunGrow.Auth.Print()
		}

		if ca.Api.SunGrow.HasTokenChanged() {
			ca.Api.LastLogin = ca.Api.SunGrow.GetLastLogin()
			ca.Api.ApiToken = ca.Api.SunGrow.GetToken()
			ca.Error = cmds.Unify.WriteConfig()
		}

		// if Cmd.GoogleSheetUpdate {
		// 	SunGrow.OutputType = iSolarCloud.TypeGoogle
		// }

		// Git.Error = Cmd.GitSet()
		// if Cmd.Error != nil {
		//	break
		// }
		//
		// ca.Valid = true
	}

	return ca.Error
}

func (ca *Cmds) SetOutputType(cmd *cobra.Command) error {
	var err error
	for range Only.Once {
		foo := cmd.Parent()
		ca.Api.SunGrow.SetOutputType(foo.Use)
	}

	return err
}


func (c *CmdApi) CmdApi(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			c.Error = cmd.Help()
			break
		}
	}
}

func (c *CmdApi) CmdApiList(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()

			case args[0] == "endpoints":
				c.Error = c.SunGrow.ListEndpoints("")

			case args[0] == "areas":
				c.SunGrow.ListAreas()

			default:
				c.Error = c.SunGrow.ListEndpoints(args[0])
		}
	}
}

func (c *CmdApi) CmdApiLogin(_ *cobra.Command, _ []string) {
	for range Only.Once {
		c.Error = c.SunGrow.Login(login.SunGrowAuth{
			AppKey:       c.AppKey,
			UserAccount:  c.Username,
			UserPassword: c.Password,
			TokenFile:    c.ApiTokenFile,
			Force:        true,
		})
		if c.Error != nil {
			break
		}

		c.SunGrow.Auth.Print()

		if c.SunGrow.HasTokenChanged() {
			c.LastLogin = c.SunGrow.GetLastLogin()
			c.ApiToken = c.SunGrow.GetToken()
			c.Error = cmds.Unify.WriteConfig()
		}
	}
}

func (c *CmdApi) CmdApiGet(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(2, args)
		if args[0] == "all" {
			c.Error = c.SunGrow.AllCritical()
			break
		}

		ep := c.SunGrow.GetByJson(args[0], args[1])
		if ep.IsError() {
			c.Error = ep.GetError()
			break
		}

		if c.SunGrow.Error != nil {
			c.Error = c.SunGrow.Error
			break
		}

		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *CmdApi) CmdApiPut(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// args = fillArray(1, args)
		// c.Error = SunGrow.Init()
		// if c.Error != nil {
		// 	break
		// }
	}
}
