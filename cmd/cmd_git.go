package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)


// ******************************************************************************** //
var cmdGit = &cobra.Command{
	Use:                   "git",
	//Aliases:               []string{""},
	Short:                 fmt.Sprintf("Git related commands."),
	Long:                  fmt.Sprintf("Git related commands. "),
	Example:               PrintExamples("git", "clone", "ls -la", "diff generate.org.au.json", "commit"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			_ = cmd.Help()
		}
	}
}


// ******************************************************************************** //
var cmdSave = &cobra.Command{
	Use:                   "save [area]",
	Aliases:               []string{},
	Short:                 fmt.Sprintf("Save SunGro areas as JSON files."),
	Long:                  fmt.Sprintf("Save SunGro areas as JSON files."),
	Example:               PrintExamples("save", "all", "site", "presence", "device"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdSaveFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdSaveFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			args = DefaultAreas
		}
		if args[0] == "all" {
			args = DefaultAreas
		}

		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.GitSave(args...)
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitClone = &cobra.Command{
	Use:                   "clone",
	//Aliases:               []string{"ls"},
	Short:                 fmt.Sprintf("Clone git repo."),
	Long:                  fmt.Sprintf("Clone git repo. "),
	Example:               PrintExamples("git clone", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitCloneFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitCloneFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = Git.Clone()
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitCommit = &cobra.Command{
	Use:                   "commit [message]",
	//Aliases:               []string{"ls"},
	Short:                 fmt.Sprintf("Commit git changes."),
	Long:                  fmt.Sprintf("Commit git changes. "),
	Example:               PrintExamples("git commit", "this is a commit message"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitCommitFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitCommitFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Commit(strings.Join(args, " "))
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitAdd = &cobra.Command{
	Use:                   "add <filename>",
	//Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Add files to Git repo."),
	Long:                  fmt.Sprintf("Add files to Git repo. "),
	Example:               PrintExamples("git add", ".", "generate.org.au.json"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitAddFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitAddFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Add(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitPull = &cobra.Command{
	Use:                   "pull",
	//Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Pull Git repo."),
	Long:                  fmt.Sprintf("Pull Git repo. "),
	Example:               PrintExamples("git pull", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitPullFunc,
	Args:                  cobra.MaximumNArgs(0),
}
//goland:noinspection GoUnusedParameter
func cmdGitPullFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Pull()
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitPush = &cobra.Command{
	Use:                   "push",
	//Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Push Git repo."),
	Long:                  fmt.Sprintf("Push Git repo. "),
	Example:               PrintExamples("git push", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitPushFunc,
	Args:                  cobra.MaximumNArgs(0),
}
//goland:noinspection GoUnusedParameter
func cmdGitPushFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Push()
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitDiff = &cobra.Command{
	Use:                   "diff <filename>",
	//Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Show diffs between current and last version."),
	Long:                  fmt.Sprintf("Show diffs between current and last version. "),
	Example:               PrintExamples("git diff", "generate.org.au.json"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitDiffFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitDiffFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Git.Connect()
		if Cmd.Error != nil {
			break
		}

		path := AddJsonSuffix(args[0])
		Cmd.Error = Git.Diff(path)
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdGitLs = &cobra.Command{
	Use:                   "ls <filename>",
	Aliases:               []string{"list"},
	Short:                 fmt.Sprintf("List files in Git repo."),
	Long:                  fmt.Sprintf("List files in Git repo. "),
	Example:               PrintExamples("git ls", "-la", "-l README.md"),
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGitLsFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdGitLsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		//args = fillArray(1, args)
		Cmd.Error = Cmd.GitLs(args...)
		if Cmd.Error != nil {
			break
		}
	}
}
