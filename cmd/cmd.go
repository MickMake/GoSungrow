package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)


func AttachRootCmd(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var rootCmd = &cobra.Command{
		Use:              DefaultBinaryName,
		Short:            fmt.Sprintf("%s - Manage an SunGrow  instance", DefaultBinaryName),
		Long:             fmt.Sprintf("%s - Manage an SunGrow  instance", DefaultBinaryName),
		Run:              gbRootFunc,
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			return initConfig(cmd)
		},
	}
	if cmd != nil {
		cmd.AddCommand(rootCmd)
	}
	rootCmd.Example = PrintExamples(rootCmd, "")

	rootCmd.SetHelpTemplate(DefaultHelpTemplate)
	rootCmd.SetUsageTemplate(DefaultUsageTemplate)
	rootCmd.SetVersionTemplate(DefaultVersionTemplate)

	rootCmd.PersistentFlags().StringVarP(&Cmd.ApiUsername, flagApiUsername, "u", "", fmt.Sprintf("SunGrow: api username."))
	rootViper.SetDefault(flagApiUsername, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.ApiPassword, flagApiPassword, "p", "", fmt.Sprintf("SunGrow: api password."))
	rootViper.SetDefault(flagApiPassword, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.ApiAppKey, flagApiAppKey, "", defaultApiAppKey, fmt.Sprintf("SunGrow: api application key."))
	rootViper.SetDefault(flagApiAppKey, defaultApiAppKey)
	rootCmd.PersistentFlags().StringVarP(&Cmd.ApiUrl, flagApiUrl, "", defaultHost, fmt.Sprintf("SunGrow: Provider API URL."))
	rootViper.SetDefault(flagApiUrl, defaultHost)
	rootCmd.PersistentFlags().DurationVarP(&Cmd.ApiTimeout, flagApiTimeout, "", defaultTimeout, fmt.Sprintf("SunGrow: API timeout."))
	rootViper.SetDefault(flagApiTimeout, defaultTimeout)
	rootCmd.PersistentFlags().StringVar(&Cmd.ApiLastLogin, flagApiLastLogin, "", "SunGrow: last login.")
	rootViper.SetDefault(flagApiLastLogin, "")
	// _ = rootCmd.PersistentFlags().MarkHidden(flagApiLastLogin)

	rootCmd.PersistentFlags().StringVarP(&Cmd.GoogleSheet, flagGoogleSheet, "", "", fmt.Sprintf("Google: Sheet URL for updates."))
	rootViper.SetDefault(flagGoogleSheet, "")
	rootCmd.PersistentFlags().BoolVarP(&Cmd.GoogleSheetUpdate, flagGoogleSheetUpdate, "", false, fmt.Sprintf("Update Google sheets."))
	rootViper.SetDefault(flagGoogleSheetUpdate, false)
	_ = rootCmd.PersistentFlags().MarkHidden(flagGoogleSheetUpdate)

	rootCmd.PersistentFlags().StringVarP(&Cmd.GitRepo, flagGitRepo, "", "", fmt.Sprintf("Git: Repo url for updates."))
	rootViper.SetDefault(flagGitRepo, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitRepoDir, flagGitRepoDir, "", "", fmt.Sprintf("Git: Local repo directory."))
	rootViper.SetDefault(flagGitRepoDir, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitUsername, flagGitUsername, "", "", fmt.Sprintf("Git: Repo username."))
	rootViper.SetDefault(flagGitUsername, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitPassword, flagGitPassword, "", "", fmt.Sprintf("Git: Repo password."))
	rootViper.SetDefault(flagGitPassword, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitKeyFile, flagGitKeyFile, "", "", fmt.Sprintf("Git: Repo SSH keyfile."))
	rootViper.SetDefault(flagGitKeyFile, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitToken, flagGitToken, "", "", fmt.Sprintf("Git: Repo token string."))
	rootViper.SetDefault(flagGitToken, "")
	rootCmd.PersistentFlags().StringVarP(&Cmd.GitDiffCmd, flagGitDiffCmd, "", "tkdiff", fmt.Sprintf("Git: Command for diffs."))
	rootViper.SetDefault(flagGitDiffCmd, "tkdiff")

	rootCmd.PersistentFlags().StringVar(&Cmd.ConfigFile, flagConfigFile, Cmd.ConfigFile, fmt.Sprintf("%s: config file.", DefaultBinaryName))
	// _ = rootCmd.PersistentFlags().MarkHidden(flagConfigFile)
	rootCmd.PersistentFlags().BoolVarP(&Cmd.Debug, flagDebug, "", false, fmt.Sprintf("%s: Debug mode.", DefaultBinaryName))
	rootViper.SetDefault(flagDebug, false)
	rootCmd.PersistentFlags().BoolVarP(&Cmd.Quiet, flagQuiet, "q", false, fmt.Sprintf("%s: Silence all messages.", DefaultBinaryName))
	rootViper.SetDefault(flagQuiet, false)

	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false

	return rootCmd
}


func gbRootFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			_ = cmd.Help()
			break
		}
	}
}
