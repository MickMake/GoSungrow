package cmdconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anicoll/gosungrow/pkg/cmdpath"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	BinaryName    string
	BinaryVersion string
	Dir           string
	File          string
	EnvPrefix     string
	Error         error

	viper   *viper.Viper
	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

func New(name string, version string, envPrefix string) *Config {
	var ret *Config

	for range only.Once {
		ret = &Config{
			BinaryName:    name,
			BinaryVersion: version,
			Dir:           ".",
			File:          defaultConfigFile,
			EnvPrefix:     envPrefix, // cmdVersion.GetEnvPrefix(),
			Error:         nil,

			viper:   viper.New(),
			cmd:     nil,
			SelfCmd: nil,
		}

		ret.Dir, ret.Error = os.UserHomeDir()
		if ret.Error != nil {
			break
		}
		ret.SetDir(filepath.Join(ret.Dir, "."+name))

		// Cmd.CacheDir, ret.Error = os.UserHomeDir()
		// if ret.Error != nil {
		// 	break
		// }
		// Cmd.CacheDir = filepath.Join(Cmd.CacheDir, "." + defaults.BinaryName, "cache")
		// _, ret.Error = os.Stat(Cmd.CacheDir)
		// if os.IsExist(ret.Error) {
		// 	break
		// }
		// ret.Error = os.MkdirAll(Cmd.CacheDir, 0700)
		// if ret.Error != nil {
		// 	break
		// }

		ret.SetFile(filepath.Join(ret.Dir, defaultConfigFile))
	}

	return ret
}

func (c *Config) GetViper() *viper.Viper {
	return c.viper
}

func (c *Config) GetCmd() *cobra.Command {
	return c.SelfCmd
}

func (c *Config) GetBinaryName() string {
	return c.BinaryName
}

func (c *Config) GetBinaryVersion() string {
	return c.BinaryVersion
}

func (c *Config) SetDir(path string) {
	for range only.Once {
		if path == "" {
			break
		}
		c.Dir = path
		c.viper.AddConfigPath(c.Dir)

		p := cmdpath.NewPath(c.Dir)
		if p.DirExists() {
			break
		}

		c.Error = p.MkdirAll()
		if c.Error != nil {
			break
		}

		// _, c.Error = os.Stat(c.Dir)
		// if !os.IsNotExist(c.Error) {
		// 	break
		// }
		// c.Error = os.MkdirAll(c.Dir, 0700)
		// if c.Error != nil {
		// 	break
		// }
	}
}

func (c *Config) SetFile(fn string) {
	for range only.Once {
		if fn == "" {
			break
		}
		c.File = fn
		c.viper.SetConfigFile(c.File)
		// c.viper.SetConfigName("config")
	}
}

func (c *Config) SetPath(path string) {
	for range only.Once {
		if path == "" {
			break
		}

		var dir string
		dir, c.Error = filepath.Abs(filepath.Dir(path))
		if c.Error != nil {
			break
		}
		c.SetDir(dir)

		// path = filepath.Base(path)
		c.SetFile(path)
	}
}

// Init reads in config file and ENV variables if set.
func (c *Config) Init(_ *cobra.Command) error {
	var err error

	for range only.Once {
		// Check for change to config file specified in flags.
		config := c.cmd.Flag(ConfigFileFlag).Value.String()
		// fmt.Printf("config: %s\n", config)
		if config != "" {
			c.SetPath(config)
		}
		// fmt.Printf("config: %s\n", c.viper.ConfigFileUsed())

		// If a config file is found, read it in.
		err = c.Open()
		if err != nil {
			break
		}
	}

	return err
}

// UpdateFlags - reads in config file and ENV variables if set.
func (c *Config) UpdateFlags() error {
	var err error

	for range only.Once {
		c.viper.SetEnvPrefix(c.EnvPrefix)
		c.viper.AutomaticEnv() // read in environment variables that match
		// c.viper.BindFlagValues(c.cmd.PersistentFlags())

		c.cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			// Environment variables can't have dashes in them, so bind them to their equivalent
			// keys with underscores, e.g. --favorite-color to STRING_FAVOURITE_COLOR
			if strings.Contains(f.Name, "-") {
				envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
				err = c.viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", c.EnvPrefix, envVarSuffix))
			}

			// Apply the viper config value to the flag when the flag is not set and viper has a value
			if !f.Changed && c.viper.IsSet(f.Name) {
				// val := c.viper.Get(f.Name)	// Doesn't handle time.Duration well.
				// val := c.cmd.Flag(f.Name).Value.String()
				// val := f.Value.String()

				val := fmt.Sprintf("%v", c.viper.Get(f.Name))
				err = c.cmd.Flags().Set(f.Name, val)
			} else {
				err = c.cmd.Flags().Set(f.Name, c.cmd.Flag(f.Name).Value.String())
			}
		})

		// c.cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// 	// Environment variables can't have dashes in them, so bind them to their equivalent
		// 	// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		// 	if strings.Contains(f.Name, "-") {
		// 		envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
		// 		err = c.viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", c.EnvPrefix, envVarSuffix))
		// 	}
		//
		// 	// fmt.Printf("FLAG: %s => %s\n", f.Name, f.Value.String())
		// 	// Apply the viper config value to the flag when the flag is not set and viper has a value
		// 	if !f.Changed && c.viper.IsSet(f.Name) {
		// 		// val := c.viper.Get(f.Name)	// Doesn't handle time.Duration well.
		// 		// val := c.cmd.Flag(f.Name).Value.String()
		// 		val := f.Value.String()
		// 		err = c.cmd.Flags().Set(f.Name, val)
		// 	}
		// })

		if err != nil {
			break
		}
	}

	return err
}

func (c *Config) Open() error {
	for range only.Once {
		c.Error = c.viper.ReadInConfig()
		if _, ok := c.Error.(viper.UnsupportedConfigError); ok {
			break
		}

		if _, ok := c.Error.(viper.ConfigParseError); ok {
			break
		}

		if _, ok := c.Error.(viper.ConfigMarshalError); ok {
			break
		}

		if os.IsNotExist(c.Error) {
			c.cmd.Flags().VisitAll(func(f *pflag.Flag) {
				switch f.Value.Type() {
				case "duration":
					c.viper.SetDefault(f.Name, f.Value.String())
				default:
					c.viper.SetDefault(f.Name, f.Value)
				}
			})

			c.Error = c.viper.WriteConfig()
			if c.Error != nil {
				break
			}

			c.Error = c.viper.ReadInConfig()
		}
		if c.Error != nil {
			break
		}

		c.Error = c.viper.MergeInConfig()
		if c.Error != nil {
			break
		}

		c.Error = c.UpdateFlags()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *Config) Write() error {
	for range only.Once {
		c.Error = c.viper.MergeInConfig()
		if c.Error != nil {
			break
		}

		c.cmd.Flags().VisitAll(func(f *pflag.Flag) {
			switch f.Value.Type() {
			case "stringArray":
				va, _ := c.cmd.Flags().GetStringArray(f.Name)
				c.viper.Set(f.Name, []string{va[0]})
			case "duration":
				c.viper.Set(f.Name, f.Value.String())
			default:
				c.viper.Set(f.Name, f.Value)
			}
		})

		c.Error = c.viper.WriteConfig()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *Config) Read() error {
	for range only.Once {
		c.Error = c.viper.ReadInConfig()
		if c.Error != nil {
			break
		}

		_, _ = fmt.Fprintln(os.Stderr, "Config file settings:")
		c.cmd.Flags().VisitAll(func(f *pflag.Flag) {
			_, _ = fmt.Fprintf(os.Stderr, "%s:			%s\n", strings.ToTitle(f.Name), f.Value.String())
		})
	}

	return c.Error
}

func (c *Config) SetDefault(key string, value interface{}) {
	for range only.Once {
		c.viper.SetDefault(key, value)
		// c.Flags[key] = value
	}
}
