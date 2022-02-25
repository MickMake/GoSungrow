package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)


func AttachCmdHelpFlags(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdHelpFlags = &cobra.Command{
		Use: "help-all",
		// Aliases:               []string{"flags"},
		Short:                 fmt.Sprintf("Extended help"),
		Long:                  fmt.Sprintf("Extended help"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdHelpFlagsFunc,
		Args:                  cobra.RangeArgs(0, 0),
	}
	cmd.AddCommand(cmdHelpFlags)
	cmdHelpFlags.Example = PrintExamples(cmdHelpFlags, "")

	return cmdHelpFlags
}


func cmdHelpFlagsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) > 0 {
			fmt.Println("Unknown sub-command.")
		}

		ExtendedHelp()

		// cmd.SetUsageTemplate(DefaultFlagHelpTemplate)
		cmd.SetUsageTemplate("")
		_ = cmd.Help()

		PrintFlags(rootCmd)
	}
}


func PrintExamples(cmd *cobra.Command, examples ...string) string {
	var ret string

	c := BuildCmd(cmd)
	for _, example := range examples {
		ret += fmt.Sprintf("\t%s %s\n", c, example)
	}

	return ret
}

func BuildCmd(cmd *cobra.Command) string {
	var ret string
	if cmd.HasParent() {
		ret += BuildCmd(cmd.Parent())
	}
	ret += cmd.Name() + " "
	return ret
}

func PrintFlags(cmd *cobra.Command) {
	fmt.Printf("\nUsing environment variables instad of flags.\n")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flag", "Short flag", "Environment", "Description", "Default"})
	table.SetBorder(true)

	cmd.PersistentFlags().SortFlags = false
	cmd.Flags().SortFlags = false
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		sh := ""
		if flag.Shorthand != "" {
			sh = "-" + flag.Shorthand
		}

		table.Append([]string{
			"--" + flag.Name,
			sh,
			PrintFlagEnv(flag.Name),
			flag.Usage,
			flag.DefValue,
			// flag.Value.String(),
		})
	})

	table.Render()
}

func PrintConfig(cmd *cobra.Command) {
	fmt.Printf("Config file '%s':\n", Cmd.ConfigFile)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flag", "Short flag", "Environment", "Description", "Value"})
	table.SetBorder(true)

	cmd.PersistentFlags().SortFlags = false
	cmd.Flags().SortFlags = false
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		sh := ""
		if flag.Shorthand != "" {
			sh = "-" + flag.Shorthand
		}

		table.Append([]string{
			"--" + flag.Name,
			sh,
			PrintFlagEnv(flag.Name),
			flag.Usage,
			flag.Value.String(),
			// flag.Value.String(),
		})
	})

	table.Render()
}

func PrintFlagEnv(flag string) string {
	fenv := strings.ReplaceAll(flag, "-", "_")
	fenv = strings.ToUpper(fenv)

	// ret := fmt.Sprintf("--%s\t%s_%s\n", flag, EnvPrefix, fenv)
	ret := fmt.Sprintf("%s_%s", EnvPrefix, fenv)
	return ret
}

func ExtendedHelp() {
	fmt.Println(strings.ReplaceAll(ExtendedHelpText, "DefaultBinaryName", DefaultBinaryName))
}
