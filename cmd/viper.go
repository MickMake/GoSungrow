package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)


var rootViper *viper.Viper


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
