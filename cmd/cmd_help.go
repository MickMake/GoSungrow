package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)


const DefaultHelpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

const DefaultUsageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags: Use "{{.Root.CommandPath}} help-all" for more info.

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand }}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} help [command]" for more information about a command.{{end}}
`

const DefaultVersionTemplate = `
{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
`

const DefaultFlagHelpTemplate = `{{if .HasAvailableInheritedFlags}}Flags available for all commands:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`

// ******************************************************************************** //
var cmdHelpFlags = &cobra.Command{
	Use:                   "help-all",
	//Aliases:               []string{"flags"},
	Short:                 fmt.Sprintf("Extended help"),
	Long:                  fmt.Sprintf("Extended help"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdHelpFlagsFunc,
	Args:                  cobra.RangeArgs(0, 0),
}
//goland:noinspection GoUnusedParameter
func cmdHelpFlagsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) > 0 {
			fmt.Println("Unknown sub-command.")
		}

		ExtendedHelp()

		//cmd.SetUsageTemplate(DefaultFlagHelpTemplate)
		cmd.SetUsageTemplate("")
		_ = cmd.Help()

		PrintFlags(rootCmd)
	}
}

func PrintExamples(cmd string, examples ...string) string {
	var ret string

	for _, example := range examples {
		ret += fmt.Sprintf("\t%s %s %s\n", DefaultBinaryName, cmd, example)
	}

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
			//flag.Value.String(),
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
			//flag.Value.String(),
		})
	})

	table.Render()
}

func PrintFlagEnv(flag string) string {
	fenv := strings.ReplaceAll(flag, "-", "_")
	fenv = strings.ToUpper(fenv)

	//ret := fmt.Sprintf("--%s\t%s_%s\n", flag, EnvPrefix, fenv)
	ret := fmt.Sprintf("%s_%s", EnvPrefix, fenv)
	return ret
}

func ExtendedHelp()  {
	var str = `
DefaultBinaryName - Over The Wire SunGro to Gitlab syncing tool.

This tool does several things:
1. Pull a Gitlab repo that holds SunGro data.
2. Fetch all data available from the SunGro.
3. Save this data as a JSON file.
4. Commit changes to the Gitlab repo.
5. Push changes to remote.

It is intended to provide full revision history for any changes made to the SunGro.

Use case example:
# Record changes made to user data on SunGro. (Will clone if not existing.)
	% DefaultBinaryName sync 'Updated users' users

# Record changes made to all SunGro manually.
	% DefaultBinaryName sync 'Updated all zones'

# Record changes made to the SunGro with default commit message.
	% DefaultBinaryName sync default

# Record changes made to the SunGro via every 30 minutes.
	% DefaultBinaryName cron run ./30 . . . . sync default

# Show changes made to a zone JSON file.
	% DefaultBinaryName git diff devices.json

# Other available Gitlab commands.
	Clone repo.
	% DefaultBinaryName git clone

	Pull repo.
	% DefaultBinaryName git pull

	Add files to repo.
	% DefaultBinaryName git add .

	Push repo.
	% DefaultBinaryName git push

	Commit changes to repo.
	% DefaultBinaryName git commit 'this is a commit message'

# Config file.
	Show current config.
	% DefaultBinaryName config read

	Change diff command used in compares.
	% DefaultBinaryName --diff-cmd='sdiff' config write

	Change Git repo directory.
	% DefaultBinaryName --git-dir=/some/other/directory config write

	Change Git repo url.
	% DefaultBinaryName --git-url=https://github.com/MickMake/iSolarData config write

	Change SunGro API token.
	% DefaultBinaryName --cf-token='this is a token string' config write

`

	str = strings.ReplaceAll(str, "DefaultBinaryName", DefaultBinaryName)
	fmt.Println(str)
}
