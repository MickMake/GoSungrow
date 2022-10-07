package cmd


// const (
// 	flagFormat = "format"
// 	flagGoFormat = "go"
// 	flagCppFormat = "cpp"
// 	flagJavaFormat = "java"
// )
//
// func (cs *Cmds) AttachFlags(cmd *cobra.Command, viper *viper.Viper) {
// 	// for range Only.Once {
// 	// 	cmd.PersistentFlags().StringVarP(&cs.Data.FormatType, flagFormat, "", "go", fmt.Sprintf("Which layout to use for parsing and formatting."))
// 	// 	viper.SetDefault(flagFormat, "go")
// 	// 	// cmd.PersistentFlags().StringArrayVarP(&cs.Data.FormatType, flagFormat, "", []string{"go"}, fmt.Sprintf("Which layout to use for parsing and formatting."))
// 	// 	// viper.SetDefault(flagFormat, "go")
// 	// 	// cmd.PersistentFlags().BoolVarP(&cs.Data.CppFormat, flagCppFormat, "", false, fmt.Sprintf("Use CPP layout for parsing."))
// 	// 	// viper.SetDefault(flagCppFormat, false)
// 	// 	// cmd.PersistentFlags().BoolVarP(&cs.Data.JavaFormat, flagJavaFormat, "", false, fmt.Sprintf("Use Java layout for parsing."))
// 	// 	// viper.SetDefault(flagJavaFormat, false)
// 	//
// 	//
// 	// 	// var rootCmd = &cobra.Command{
// 	// 	// 	Use:              DefaultBinaryName,
// 	// 	// 	Short:            fmt.Sprintf("%s - Manage a SunGrow instance", DefaultBinaryName),
// 	// 	// 	Long:             fmt.Sprintf("%s - Manage a SunGrow instance", DefaultBinaryName),
// 	// 	// 	Run:              gbRootFunc,
// 	// 	// 	TraverseChildren: true,
// 	// 	// 	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
// 	// 	// 		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
// 	// 	// 		return initConfig(cmd)
// 	// 	// 	},
// 	// 	// }
// 	// 	// if cmd != nil {
// 	// 	// 	cmd.AddCommand(rootCmd)
// 	// 	// }
// 	// 	// rootCmd.Example = PrintExamples(rootCmd, "")
// 	// 	//
// 	// 	//
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GoogleSheet, flagGoogleSheet, "", "", fmt.Sprintf("Google: Sheet URL for updates."))
// 	// 	// rootViper.SetDefault(flagGoogleSheet, "")
// 	// 	// rootCmd.PersistentFlags().BoolVarP(&Cmd.GoogleSheetUpdate, flagGoogleSheetUpdate, "", false, fmt.Sprintf("Update Google sheets."))
// 	// 	// rootViper.SetDefault(flagGoogleSheetUpdate, false)
// 	// 	// _ = rootCmd.PersistentFlags().MarkHidden(flagGoogleSheetUpdate)
// 	//
// 	//
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitRepo, flagGitRepo, "", "", fmt.Sprintf("Git: Repo url for updates."))
// 	// 	// rootViper.SetDefault(flagGitRepo, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitRepoDir, flagGitRepoDir, "", "", fmt.Sprintf("Git: Local repo directory."))
// 	// 	// rootViper.SetDefault(flagGitRepoDir, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitUsername, flagGitUsername, "", "", fmt.Sprintf("Git: Repo username."))
// 	// 	// rootViper.SetDefault(flagGitUsername, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitPassword, flagGitPassword, "", "", fmt.Sprintf("Git: Repo password."))
// 	// 	// rootViper.SetDefault(flagGitPassword, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitKeyFile, flagGitKeyFile, "", "", fmt.Sprintf("Git: Repo SSH keyfile."))
// 	// 	// rootViper.SetDefault(flagGitKeyFile, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitToken, flagGitToken, "", "", fmt.Sprintf("Git: Repo token string."))
// 	// 	// rootViper.SetDefault(flagGitToken, "")
// 	// 	// rootCmd.PersistentFlags().StringVarP(&Cmd.GitDiffCmd, flagGitDiffCmd, "", "tkdiff", fmt.Sprintf("Git: Command for diffs."))
// 	// 	// rootViper.SetDefault(flagGitDiffCmd, "tkdiff")
// 	//
// 	// 	rootCmd = AttachRootCmd(nil)
// 	// 	_ = AttachCmdConfig(rootCmd)
// 	// 	_ = AttachCmdApi(rootCmd)
// 	// 	_ = AttachCmdData(rootCmd)
// 	// 	_ = AttachCmdMqtt(rootCmd)
// 	// 	_ = AttachCmdGit(rootCmd)
// 	// 	_ = AttachCmdGoogle(rootCmd)
// 	// 	_ = AttachCmdCron(rootCmd)
// 	// 	_ = AttachCmdVersion(rootCmd)
// 	// 	_ = AttachCmdHelpFlags(rootCmd)
// 	//
// 	// }
// }
//
// func (cs *Cmds) InitArgs(cmd *cobra.Command, _ []string) error {
// 	for range Only.Once {
// 		cs.Error = nil
// 		if cs.reparse {
// 			// We can re-execute commands inline, (particularly in a shell).
// 			break
// 		}
// 		cs.reparse = true
//
// 		cs.Data.SetDateIfNil()
// 		cs.Data.SetCmd(cmd.Name())
// 		switch cmd.Flag(flagFormat).Value.String() {
// 		case flagGoFormat:
// 			cs.Data.GoFormat = true
// 			cs.Data.CppFormat = false
// 			cs.Data.JavaFormat = false
// 		case flagCppFormat:
// 			cs.Data.GoFormat = false
// 			cs.Data.CppFormat = true
// 			cs.Data.JavaFormat = false
// 		case flagJavaFormat:
// 			cs.Data.GoFormat = false
// 			cs.Data.CppFormat = false
// 			cs.Data.JavaFormat = true
// 		default:
// 			cs.Error = errors.New(fmt.Sprintf("Unknown format layout. Can be only one of \"%s\", \"%s\" or \"%s\"",
// 				flagGoFormat, flagCppFormat, flagJavaFormat))
// 			break
// 		}
//
// 		if cs.Data.Convert == nil {
// 			fn := filepath.Join(cs.Unify.Commands.CmdConfig.Dir, "convert.json")
// 			cs.Data.Convert, cs.Error = cal.ReadConvert(fn)
// 			if cs.Error != nil {
// 				break
// 			}
// 		}
// 	}
//
// 	return cs.Error
// }
