package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
	"os/user"
)


// ******************************************************************************** //
var cmdSync = &cobra.Command{
	Use:                   "sync <commit msg | default> [area]",
	Aliases:               []string{},
	Short:                 fmt.Sprintf("Sync SunGro to Git repo."),
	Long:                  fmt.Sprintf("Sync SunGro to Git repo."),
	Example:               PrintExamples("sync", "default", "'updated everything'", "'this is an update' users"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdSyncFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdSyncFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) < 1 {
			Cmd.Error = cmd.Help()
			break
		}

		var msg string
		switch {
			case args[0] == "":
				fallthrough
			case args[0] == "default":
				u, _ := user.Current()
				msg = fmt.Sprintf("Regular sync by %s", u.Username)
			default:
				msg = args[0]
		}

		args = args[1:]
		if len(args) == 0 {
			args = DefaultAreas
		}

		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.GitSync(msg, args...)
		if Cmd.Error != nil {
			break
		}
	}
}
