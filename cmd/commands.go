package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/defaults"
	"github.com/MickMake/GoUnify/Unify"
	"github.com/spf13/cobra"
)


type Cmds struct {
	Unify  *Unify.Unify
	Api    *CmdApi
	Data   *CmdData
	Info   *CmdInfo
	Mqtt   *CmdMqtt

	// SunGrow *iSolarCloud.SunGrow
	// Git *mmGit.Git
	// Mqtt *mmHa.Mqtt

	ConfigDir   string
	CacheDir    string
	ConfigFile  string
	WriteConfig bool
	Quiet       bool
	Debug       bool
	// OutputType  string
	// OutputFile  string
	// ApiOutputType string

	// // Google sheets
	// GoogleSheet       string
	// GoogleSheetUpdate bool
	//
	// // GitHub api
	// GitRepo     string
	// GitRepoDir  string
	// GitUsername string
	// GitPassword string
	// GitKeyFile  string
	// GitToken    string
	// GitDiffCmd  string

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

		cmds.Mqtt = NewCmdMqtt()
		cmds.Mqtt.AttachCommand(cmdRoot)

		// cmds.Git = cmdGit.NewCmdGit()
		// cmds.Git.AttachCommands(cmdRoot)
		// cmds.Git.AttachFlags(cmdRoot, cmds.Unify.GetViper())
		// _ = cmds.Git.SetDefaultEntities(defaults.DefaultAreas...)
		//
		// cmds.Google = cmdGoogle.NewCmdGoogle()
		// cmds.Google.AttachCommands(cmdRoot)
		//
		// cmds.AttachFlags(cmdRoot, cmds.Unify.GetViper())
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
