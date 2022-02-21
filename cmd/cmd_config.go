package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


// ******************************************************************************** //
var cmdConfig = &cobra.Command{
	Use:                   "config",
	Short:                 "Create, update or show config file.",
	Long:                  "Create, update or show config file.",
	Example:               PrintExamples("config", "read", "write", "write --git-dir=/some/other/directory"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdConfigFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdConfigFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", rootViper.ConfigFileUsed())
		if len(args) == 0 {
			_ = cmd.Help()
		}
	}
}


// ******************************************************************************** //
var cmdConfigWrite = &cobra.Command{
	Use:                   "write",
	Short:                 "Update config file.",
	Long:                  "Update config file from CLI args.",
	Example:               PrintExamples("config write","", "--git-dir=/some/other/directory", "--diff-cmd=tkdiff"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdConfigWriteFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdConfigWriteFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 1 {
			Cmd.ConfigFile = args[0]
			rootViper.SetConfigFile(Cmd.ConfigFile)
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", rootViper.ConfigFileUsed())
		Cmd.Error = openConfig()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = writeConfig()
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdConfigRead = &cobra.Command{
	Use:                   "read",
	Short:                 "Read config file.",
	Long:                  "Read config file.",
	Example:               PrintExamples("config read", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdConfigReadFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdConfigReadFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 1 {
			Cmd.ConfigFile = args[0]
			rootViper.SetConfigFile(Cmd.ConfigFile)
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", rootViper.ConfigFileUsed())
		Cmd.Error = openConfig()
		if Cmd.Error != nil {
			break
		}

		PrintConfig(rootCmd)
		//Cmd.Error = readConfig()
		//if Cmd.Error != nil {
		//	break
		//}
	}
}
