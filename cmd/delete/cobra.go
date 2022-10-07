package cmd


// var rootCmd *cobra.Command
//
//
// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() error {
// 	for range Only.Once {
// 		Cmd.ConfigDir, Cmd.Error = os.UserHomeDir()
// 		if Cmd.Error != nil {
// 			break
// 		}
// 		Cmd.ConfigDir = filepath.Join(Cmd.ConfigDir, "."+DefaultBinaryName)
// 		_, Cmd.Error = os.Stat(Cmd.ConfigDir)
// 		if os.IsExist(Cmd.Error) {
// 			break
// 		}
// 		Cmd.Error = os.MkdirAll(Cmd.ConfigDir, 0700)
// 		if Cmd.Error != nil {
// 			break
// 		}
//
// 		Cmd.CacheDir, Cmd.Error = os.UserHomeDir()
// 		if Cmd.Error != nil {
// 			break
// 		}
// 		Cmd.CacheDir = filepath.Join(Cmd.CacheDir, "."+DefaultBinaryName, "cache")
// 		_, Cmd.Error = os.Stat(Cmd.CacheDir)
// 		if os.IsExist(Cmd.Error) {
// 			break
// 		}
// 		Cmd.Error = os.MkdirAll(Cmd.CacheDir, 0700)
// 		if Cmd.Error != nil {
// 			break
// 		}
//
// 		Cmd.ConfigFile = filepath.Join(Cmd.ConfigDir, defaultConfigFile)
// 		Cmd.ApiTokenFile = filepath.Join(Cmd.ConfigDir, defaultTokenFile)
//
// 		rootViper = viper.New()
// 		rootCmd = AttachRootCmd(nil)
// 		_ = AttachCmdConfig(rootCmd)
// 		_ = AttachCmdApi(rootCmd)
// 		_ = AttachCmdData(rootCmd)
// 		_ = AttachCmdMqtt(rootCmd)
// 		_ = AttachCmdGit(rootCmd)
// 		_ = AttachCmdGoogle(rootCmd)
// 		_ = AttachCmdCron(rootCmd)
// 		_ = AttachCmdVersion(rootCmd)
// 		_ = AttachCmdHelpFlags(rootCmd)
//
// 		// cobra.OnInitialize(initConfig)	// Bound to rootCmd now.
// 		cobra.EnableCommandSorting = false
//
// 		err := rootCmd.Execute()
// 		if err != nil {
// 			break
// 		}
// 		if Cmd.Error != nil {
// 			break
// 		}
// 	}
//
// 	return Cmd.Error
// }
//
// //goland:noinspection GoUnusedFunction
// func showArgs(cmd *cobra.Command, args []string) {
// 	for range Only.Once {
// 		flargs := cmd.Flags().Args()
// 		if flargs != nil {
// 			fmt.Printf("'%s' called with '%s'\n", cmd.CommandPath(), strings.Join(flargs, " "))
// 			break
// 		}
//
// 		fmt.Printf("'%s' called with '%s'\n", cmd.CommandPath(), strings.Join(args, " "))
// 		break
// 	}
//
// 	fmt.Println("")
// }
//
// func fillArray(count int, args []string) []string {
// 	var ret []string
// 	for range Only.Once {
// 		//
// 		// if len(args) == 0 {
// 		//	break
// 		// }
// 		ret = make([]string, count)
// 		for i, e := range args {
// 			ret[i] = e
// 		}
// 	}
// 	return ret
// }
//
// func AddJsonSuffix(fn string) string {
// 	return strings.TrimSuffix(fn, ".json") + ".json"
// }
