package cmd

import (
	"GoSungrow/defaults"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/Unify"
	"github.com/spf13/cobra"
)


type Cmds struct {
	Unify  *Unify.Unify
	Api    *CmdApi
	Data   *CmdData
	Info   *CmdInfo
	Show   *CmdShow
	Mqtt   *CmdMqtt
	Ha     *CmdHa

	ConfigDir   string
	CacheDir    string
	ConfigFile  string
	WriteConfig bool
	Quiet       bool
	Debug       bool

	Args []string

	Error error
}

//goland:noinspection GoNameStartsWithPackageName
type CmdDefault struct {
	Error   error
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}


var cmds Cmds


func init() {
	for range Only.Once {
		cmds.Unify = Unify.New(
			Unify.Options{
				Description:   defaults.Description,
				BinaryName:    defaults.BinaryName,
				BinaryVersion: defaults.BinaryVersion,
				SourceRepo:    defaults.SourceRepo,
				BinaryRepo:    defaults.BinaryRepo,
				EnvPrefix:     defaults.EnvPrefix,
				HelpSummary:   defaults.HelpSummary,
				ReadMe:        defaults.Readme,
				Examples:      defaults.Examples,
			},
			Unify.Flags{},
		)

		cmdRoot := cmds.Unify.GetCmd()

		cmds.Api = NewCmdApi()
		cmds.Api.AttachCommand(cmdRoot)
		cmds.Api.AttachFlags(cmdRoot, cmds.Unify.GetViper())

		cmds.Data = NewCmdData()
		cmds.Data.AttachCommand(cmdRoot)

		cmds.Info = NewCmdInfo()
		cmds.Info.AttachCommand(cmdRoot)

		cmds.Show = NewCmdShow()
		cmds.Show.AttachCommand(cmdRoot)

		cmds.Mqtt = NewCmdMqtt()
		cmds.Mqtt.AttachCommand(cmdRoot)
		cmds.Mqtt.AttachFlags(cmdRoot, cmds.Unify.GetViper())

		cmds.Ha = NewCmdHa()
		cmds.Ha.AttachCommand(cmdRoot)
	}
}

func Execute() error {
	var err error

	for range Only.Once {
		// Execute adds all child commands to the root command and sets flags appropriately.
		// This is called by main.main(). It only needs to happen once to the rootCmd.
		err = cmds.Unify.Execute()
		if err != nil {
			break
		}
	}

	return err
}


func (ca *Cmds) ProcessArgs(_ *cobra.Command, args []string) error {
	for range Only.Once {
		ca.Args = args

		ca.ConfigDir = cmds.Unify.GetConfigDir()
		ca.ConfigFile = cmds.Unify.GetConfigFile()
		ca.CacheDir = cmds.Unify.GetCacheDir()
		ca.Debug = cmds.Unify.Flags.Debug
		ca.Quiet = cmds.Unify.Flags.Quiet
	}

	return ca.Error
}
