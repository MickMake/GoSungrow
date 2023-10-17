package cmdconfig

import (
	"fmt"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/spf13/cobra"
)

// func bindFlags(cmd *cobra.Command, v *viper.Viper) error {
// 	var err error
//
// 	prefix := cmdVersion.GetEnvPrefix()
// 	cmd.Flags().VisitAll(func(f *pflag.Flag) {
// 		// Environment variables can't have dashes in them, so bind them to their equivalent
// 		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
// 		if strings.Contains(f.Name, "-") {
// 			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
// 			// err = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", defaults.EnvPrefix, envVarSuffix))
// 			err = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", prefix, envVarSuffix))
// 		}
//
// 		// Apply the viper config value to the flag when the flag is not set and viper has a value
// 		if !f.Changed && v.IsSet(f.Name) {
// 			val := v.Get(f.Name)
// 			err = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
// 		}
// 	})
//
// 	return err
// }

//goland:noinspection GoUnusedFunction
func showArgs(cmd *cobra.Command, args []string) {
	for range only.Once {
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

func FillArray(count int, args []string) []string {
	var ret []string
	for range only.Once {
		if len(args) > count {
			count = len(args)
		}
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
