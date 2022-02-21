package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/sungro"
	"GoSungro/mmGit"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var SunGro *sungro.SunGro
var Git *mmGit.Git
var Cmd CommandArgs
var rootViper *viper.Viper

func init() {
	for range Only.Once {
		rootViper = viper.New()

		Cmd.ConfigDir, Cmd.Error = os.UserHomeDir()
		if Cmd.Error != nil {
			break
		}
		Cmd.ConfigDir = filepath.Join(Cmd.ConfigDir, "."+DefaultBinaryName)

		_, Cmd.Error = os.Stat(Cmd.ConfigDir)
		if os.IsExist(Cmd.Error) {
			break
		}

		Cmd.Error = os.MkdirAll(Cmd.ConfigDir, 0700)
		if Cmd.Error != nil {
			break
		}

		Cmd.ConfigFile = filepath.Join(Cmd.ConfigDir, defaultConfigFile)
		Cmd.ApiTokenFile = filepath.Join(Cmd.ConfigDir, defaultTokenFile)

		rootCmd.PersistentFlags().StringVarP(&Cmd.ApiUsername, flagApiUsername, "u", "", fmt.Sprintf("SunGro: api username."))
		rootViper.SetDefault(flagApiUsername, "")
		rootCmd.PersistentFlags().StringVarP(&Cmd.ApiPassword, flagApiPassword, "p", "", fmt.Sprintf("SunGro: api password."))
		rootViper.SetDefault(flagApiPassword, "")
		rootCmd.PersistentFlags().StringVarP(&Cmd.ApiAppKey, flagApiAppKey, "", "", fmt.Sprintf("SunGro: api application key."))
		rootViper.SetDefault(flagApiAppKey, "")

		rootCmd.PersistentFlags().StringVarP(&Cmd.ApiUrl, flagApiUrl, "", defaultHost, fmt.Sprintf("SunGro: Provider API URL."))
		rootViper.SetDefault(flagApiUrl, "")
		rootCmd.PersistentFlags().DurationVarP(&Cmd.ApiTimeout, flagApiTimeout, "", defaultTimeout, fmt.Sprintf("SunGro: API timeout."))
		rootViper.SetDefault(flagApiTimeout, "")
		rootCmd.PersistentFlags().StringVar(&Cmd.ApiLastLogin, flagApiLastLogin, "", "SunGro: last login.")
		rootViper.SetDefault(flagApiLastLogin, "")
		_ = rootCmd.PersistentFlags().MarkHidden(flagApiLastLogin)

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

		// cobra.OnInitialize(initConfig)
		cobra.EnableCommandSorting = false
		rootCmd.PersistentFlags().SortFlags = false
		rootCmd.Flags().SortFlags = false
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	for range Only.Once {
		Cmd.Error = InitCommands()
		if Cmd.Error != nil {
			break
		}

		err := rootCmd.Execute()
		if err != nil {
			break
		}
		if Cmd.Error != nil {
			break
		}
	}

	return Cmd.Error
}

func InitCommands() error {
	var err error

	for range Only.Once {
		rootCmd.AddCommand(cmdConfig, cmdApi, cmdGit, cmdMqtt, cmdCron, cmdVersion, cmdHelpFlags)
		cmdConfig.AddCommand(cmdConfigWrite, cmdConfigRead)
		cmdGit.AddCommand(cmdGitSync, cmdSave, cmdGitLs, cmdGitDiff, cmdGitClone, cmdGitPull, cmdGitPush, cmdGitCommit, cmdGitAdd)
		cmdMqtt.AddCommand(cmdMqttSync)
		cmdCron.AddCommand(cmdCronRun, cmdCronAdd, cmdCronRemove, cmdCronList)

		cmdApi.AddCommand(cmdApiList, cmdApiLogin, cmdApiPut, cmdApiGet)
		cmdApi.PersistentFlags().BoolVarP(&Cmd.ApiGetRaw, flagApiGetRaw, "", false, fmt.Sprintf("SunGro: api raw data, (not parsed nor evaluated)."))
		rootViper.SetDefault(flagApiGetRaw, false)

		// foo := rootCmd.HelpTemplate()
		// foo := rootCmd.UsageTemplate()
		// foo := rootCmd.VersionTemplate()
		// fmt.Println(foo)

		rootCmd.SetHelpTemplate(DefaultHelpTemplate)
		rootCmd.SetUsageTemplate(DefaultUsageTemplate)
		rootCmd.SetVersionTemplate(DefaultVersionTemplate)

		foo := rootCmd.Commands()
		foo[0].CommandPath()
	}

	return err
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	var err error

	for range Only.Once {
		rootViper = viper.New()
		rootViper.AddConfigPath(Cmd.ConfigDir)
		rootViper.SetConfigFile(Cmd.ConfigFile)
		// rootViper.SetConfigName("config")

		// If a config file is found, read it in.
		err = openConfig()
		if err != nil {
			break
		}

		rootViper.SetEnvPrefix(EnvPrefix)
		rootViper.AutomaticEnv() // read in environment variables that match
		err = bindFlags(cmd, rootViper)
		if err != nil {
			break
		}
	}

	return err
}

func openConfig() error {
	var err error

	for range Only.Once {
		err = rootViper.ReadInConfig()
		if _, ok := err.(viper.UnsupportedConfigError); ok {
			break
		}

		if _, ok := err.(viper.ConfigParseError); ok {
			break
		}

		if _, ok := err.(viper.ConfigMarshalError); ok {
			break
		}

		if os.IsNotExist(err) {
			rootViper.SetDefault(flagDebug, Cmd.Debug)
			rootViper.SetDefault(flagQuiet, Cmd.Quiet)
			rootViper.SetDefault(flagApiLastLogin, Cmd.ApiLastLogin)

			rootViper.SetDefault(flagApiUrl, Cmd.ApiUrl)
			rootViper.SetDefault(flagApiTimeout, Cmd.ApiTimeout)
			rootViper.SetDefault(flagApiUsername, defaultUsername)
			rootViper.SetDefault(flagApiPassword, defaultPassword)
			rootViper.SetDefault(flagApiAppKey, Cmd.ApiAppKey)

			rootViper.SetDefault(flagGoogleSheet, Cmd.GoogleSheet)
			rootViper.SetDefault(flagGoogleSheetUpdate, Cmd.GoogleSheetUpdate)

			rootViper.SetDefault(flagGitRepo, Cmd.GitRepo)
			rootViper.SetDefault(flagGitRepoDir, Cmd.GitRepoDir)
			rootViper.SetDefault(flagGitUsername, Cmd.GitUsername)
			rootViper.SetDefault(flagGitPassword, Cmd.GitPassword)
			rootViper.SetDefault(flagGitKeyFile, Cmd.GitKeyFile)
			rootViper.SetDefault(flagGitToken, Cmd.GitToken)
			rootViper.SetDefault(flagGitDiffCmd, Cmd.GitDiffCmd)

			err = rootViper.WriteConfig()
			if err != nil {
				break
			}

			err = rootViper.ReadInConfig()
			break
		}
		if err != nil {
			break
		}

		err = rootViper.MergeInConfig()
		if err != nil {
			break
		}

		// err = viper.Unmarshal(Cmd)
	}

	return err
}

func writeConfig() error {
	var err error

	for range Only.Once {
		err = rootViper.MergeInConfig()
		if err != nil {
			break
		}

		rootViper.Set(flagDebug, Cmd.Debug)
		rootViper.Set(flagQuiet, Cmd.Quiet)
		rootViper.Set(flagApiLastLogin, Cmd.ApiLastLogin)

		rootViper.Set(flagApiUrl, Cmd.ApiUrl)
		rootViper.Set(flagApiTimeout, Cmd.ApiTimeout)
		rootViper.Set(flagApiUsername, Cmd.ApiUsername)
		rootViper.Set(flagApiPassword, Cmd.ApiPassword)
		rootViper.Set(flagApiAppKey, Cmd.ApiAppKey)

		rootViper.Set(flagGoogleSheet, Cmd.GoogleSheet)
		rootViper.Set(flagGoogleSheetUpdate, Cmd.GoogleSheetUpdate)

		rootViper.Set(flagGitRepo, Cmd.GitRepo)
		rootViper.Set(flagGitRepoDir, Cmd.GitRepoDir)
		rootViper.Set(flagGitUsername, Cmd.GitUsername)
		rootViper.Set(flagGitPassword, Cmd.GitPassword)
		rootViper.Set(flagGitKeyFile, Cmd.GitKeyFile)
		rootViper.Set(flagGitToken, Cmd.GitToken)
		rootViper.Set(flagGitDiffCmd, Cmd.GitDiffCmd)

		err = rootViper.WriteConfig()
		if err != nil {
			break
		}
	}

	return err
}

func readConfig() error {
	var err error

	for range Only.Once {
		err = rootViper.ReadInConfig()
		if err != nil {
			break
		}

		_, _ = fmt.Fprintln(os.Stderr, "Config file settings:")

		_, _ = fmt.Fprintf(os.Stderr, "Api Url:	%v\n", rootViper.Get(flagApiUrl))
		_, _ = fmt.Fprintf(os.Stderr, "Api AppKey:	%v\n", rootViper.Get(flagApiAppKey))
		_, _ = fmt.Fprintf(os.Stderr, "Api UserAccount:	%v\n", rootViper.Get(flagApiUsername))
		_, _ = fmt.Fprintf(os.Stderr, "Api UserPassword:	%v\n", rootViper.Get(flagApiPassword))
		_, _ = fmt.Fprintln(os.Stderr)

		_, _ = fmt.Fprintf(os.Stderr, "Git Repo URL:		%v\n", rootViper.Get(flagGitRepo))
		_, _ = fmt.Fprintf(os.Stderr, "Git Repo Dir:		%v\n", rootViper.Get(flagGitRepoDir))
		_, _ = fmt.Fprintf(os.Stderr, "Git Repo User:	%v\n", rootViper.Get(flagGitUsername))
		_, _ = fmt.Fprintf(os.Stderr, "Git Repo ApiPassword:	%v\n", rootViper.Get(flagGitPassword))
		_, _ = fmt.Fprintf(os.Stderr, "Git SSH keyfile:	%v\n", rootViper.Get(flagGitKeyFile))
		_, _ = fmt.Fprintf(os.Stderr, "Git Auth:	%v\n", rootViper.Get(flagGitToken))
		_, _ = fmt.Fprintf(os.Stderr, "Git Diff Command:	%v\n", rootViper.Get(flagGitDiffCmd))
		_, _ = fmt.Fprintln(os.Stderr)

		_, _ = fmt.Fprintf(os.Stderr, "Debug:		%v\n", rootViper.Get(flagDebug))
		_, _ = fmt.Fprintf(os.Stderr, "Quiet:		%v\n", rootViper.Get(flagQuiet))
		_, _ = fmt.Fprintf(os.Stderr, "ApiTimeout:	%v\n", rootViper.Get(flagApiTimeout))
	}

	return err
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) error {
	var err error

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			err = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", EnvPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			err = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	return err
}

//goland:noinspection GoUnusedFunction
func showArgs(cmd *cobra.Command, args []string) {
	for range Only.Once {
		flargs := cmd.Flags().Args()
		if flargs != nil {
			fmt.Printf("'%s' called with '%s'\n", cmd.CommandPath(), strings.Join(flargs, " "))
			break
		}

		fmt.Printf("'%s' called with '%s'\n", cmd.CommandPath(), strings.Join(args, " "))
		break
	}

	fmt.Println("")
}

func fillArray(count int, args []string) []string {
	var ret []string
	for range Only.Once {
		//
		// if len(args) == 0 {
		//	break
		// }
		ret = make([]string, count)
		for i, e := range args {
			ret[i] = e
		}
	}
	return ret
}

func AddJsonSuffix(fn string) string {
	return strings.TrimSuffix(fn, ".json") + ".json"
}
