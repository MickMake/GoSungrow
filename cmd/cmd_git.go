package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)


func AttachCmdGit(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdGit = &cobra.Command{
		Use: "git",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Git related commands."),
		Long:                  fmt.Sprintf("Git related commands. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(cmdGit)
	cmdGit.Example = PrintExamples(cmdGit, "clone", "ls -la", "diff foo.json", "commit")


	// ******************************************************************************** //
	var cmdSave = &cobra.Command{
		Use:                   "save [area]",
		Aliases:               []string{},
		Short:                 fmt.Sprintf("Save SunGrow areas as JSON files."),
		Long:                  fmt.Sprintf("Save SunGrow areas as JSON files."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdSaveFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdGit.AddCommand(cmdSave)
	cmdSave.Example = PrintExamples(cmdSave, "all", "unit", "device")

	// ******************************************************************************** //
	var cmdGitClone = &cobra.Command{
		Use: "clone",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Clone git repo."),
		Long:                  fmt.Sprintf("Clone git repo. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitCloneFunc,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmdGit.AddCommand(cmdGitClone)
	cmdGitClone.Example = PrintExamples(cmdGitClone, "")

	// ******************************************************************************** //
	var cmdGitCommit = &cobra.Command{
		Use: "commit [message]",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Commit git changes."),
		Long:                  fmt.Sprintf("Commit git changes. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitCommitFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdGit.AddCommand(cmdGitCommit)
	cmdGitCommit.Example = PrintExamples(cmdGitCommit, "this is a commit message")

	// ******************************************************************************** //
	var cmdGitAdd = &cobra.Command{
		Use: "add <filename>",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Add files to Git repo."),
		Long:                  fmt.Sprintf("Add files to Git repo. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitAddFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdGit.AddCommand(cmdGitAdd)
	cmdGitAdd.Example = PrintExamples(cmdGitAdd, ".", "foo.json")

	// ******************************************************************************** //
	var cmdGitPull = &cobra.Command{
		Use: "pull",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Pull Git repo."),
		Long:                  fmt.Sprintf("Pull Git repo. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitPullFunc,
		Args:                  cobra.MaximumNArgs(0),
	}
	cmdGit.AddCommand(cmdGitPull)
	cmdGitPull.Example = PrintExamples(cmdGitPull, "")

	// ******************************************************************************** //
	var cmdGitPush = &cobra.Command{
		Use: "push",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Push Git repo."),
		Long:                  fmt.Sprintf("Push Git repo. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitPushFunc,
		Args:                  cobra.MaximumNArgs(0),
	}
	cmdGit.AddCommand(cmdGitPush)
	cmdGitPush.Example = PrintExamples(cmdGitPush, "")

	// ******************************************************************************** //
	var cmdGitDiff = &cobra.Command{
		Use: "diff <filename>",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Show diffs between current and last version."),
		Long:                  fmt.Sprintf("Show diffs between current and last version. "),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitDiffFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdGit.AddCommand(cmdGitDiff)
	cmdGitDiff.Example = PrintExamples(cmdGitDiff, "foo.json")

	// ******************************************************************************** //
	var cmdGitLs = &cobra.Command{
		Use:                   "ls <filename>",
		Aliases:               []string{"list"},
		Short:                 fmt.Sprintf("List files in Git repo."),
		Long:                  fmt.Sprintf("List files in Git repo. "),
		DisableFlagParsing:    true,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitLsFunc,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmdGit.AddCommand(cmdGitLs)
	cmdGitLs.Example = PrintExamples(cmdGitLs, "-la", "-l README.md")

	// ******************************************************************************** //
	var cmdGitSync = &cobra.Command{
		Use:                   "sync <commit msg | default> [area]",
		Aliases:               []string{},
		Short:                 fmt.Sprintf("Sync SunGrow to Git repo."),
		Long:                  fmt.Sprintf("Sync SunGrow to Git repo."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdGitSyncFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdGit.AddCommand(cmdGitSync)
	cmdGitSync.Example = PrintExamples(cmdGitSync, "default", "'updated everything'", "'this is an update' users")

	return cmdGit
}


func cmdGitFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			_ = cmd.Help()
		}
	}
}

func cmdSaveFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		return

		// if len(args) == 0 {
		// 	args = DefaultAreas
		// }
		// if args[0] == "all" {
		// 	args = DefaultAreas
		// }
		//
		// Cmd.Error = Cmd.Git.Connect()
		// if Cmd.Error != nil {
		// 	break
		// }
		//
		// Cmd.Error = Cmd.GitSave(args...)
		// if Cmd.Error != nil {
		// 	break
		// }
	}
}

func cmdGitCloneFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		Cmd.Error = Cmd.Git.Clone()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitCommitFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Cmd.Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Git.Commit(strings.Join(args, " "))
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitAddFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Cmd.Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Git.Add(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitPullFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		Cmd.Error = Cmd.Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Git.Pull()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitPushFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		Cmd.Error = Cmd.Git.Connect()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Cmd.Git.Push()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitDiffFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		if args[0] == "" {
			_ = cmd.Help()
			break
		}

		Cmd.Error = Cmd.Git.Connect()
		if Cmd.Error != nil {
			break
		}

		path := AddJsonSuffix(args[0])
		Cmd.Error = Cmd.Git.Diff(path)
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitLsFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		// args = fillArray(4, args)
		Cmd.Error = Cmd.GitLs(args...)
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdGitSyncFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		return

		// if len(args) < 1 {
		// 	Cmd.Error = cmd.Help()
		// 	break
		// }
		//
		// var msg string
		// switch {
		// case args[0] == "":
		// 	fallthrough
		// case args[0] == "default":
		// 	u, _ := user.Current()
		// 	msg = fmt.Sprintf("Regular sync by %s", u.Username)
		// default:
		// 	msg = args[0]
		// }
		//
		// args = args[1:]
		// if len(args) == 0 {
		// 	args = DefaultAreas
		// }
		//
		// Cmd.Error = Cmd.Git.Connect()
		// if Cmd.Error != nil {
		// 	break
		// }
		//
		// Cmd.Error = Cmd.GitSync(msg, args...)
		// if Cmd.Error != nil {
		// 	break
		// }
	}
}
