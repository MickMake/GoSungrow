package cmd

import (
	"GoSungrow/Only"
	"github.com/spf13/cobra"
	"strings"
)


var DefaultAreas = []string{"all"}

// type CommandArgs struct {
// 	// SunGrow *iSolarCloud.SunGrow
// 	// Git *mmGit.Git
// 	// Mqtt *mmHa.Mqtt
// 	//
// 	// ConfigDir   string
// 	// CacheDir    string
// 	// ConfigFile  string
// 	// WriteConfig bool
// 	// Quiet       bool
// 	// Debug       bool
// 	// OutputType  string
// 	// OutputFile  string
// 	//
// 	// // iSolarCloud api
// 	// ApiTimeout    time.Duration
// 	// ApiUrl        string
// 	// ApiUsername   string
// 	// ApiPassword   string
// 	// ApiAppKey     string
// 	// ApiLastLogin  string
// 	// ApiToken      string
// 	// ApiTokenFile  string
// 	// ApiOutputType string
// 	//
// 	// // HASSIO MQTT
// 	// MqttUsername   string
// 	// MqttPassword   string
// 	// MqttHost       string
// 	// MqttPort       string
// 	//
// 	// // Google sheets
// 	// GoogleSheet       string
// 	// GoogleSheetUpdate bool
// 	//
// 	// // GitHub api
// 	// GitRepo     string
// 	// GitRepoDir  string
// 	// GitUsername string
// 	// GitPassword string
// 	// GitKeyFile  string
// 	// GitToken    string
// 	// GitDiffCmd  string
// 	//
// 	// Args []string
//
// 	Valid bool
// 	Error error
// }

// var Cmd CommandArgs


// func (ca *CommandArgs) IsValid() error {
// 	for range Only.Once {
// 		if !ca.Valid {
// 			ca.Error = errors.New("args are not valid")
// 			break
// 		}
// 	}
//
// 	return ca.Error
// }

func (ca *Cmds) ProcessArgs(_ *cobra.Command, args []string) error {
	for range Only.Once {
		ca.Args = args

		ca.ConfigDir = cmds.Unify.GetConfigDir()
		ca.ConfigFile = cmds.Unify.GetConfigFile()
		ca.CacheDir = cmds.Unify.GetCacheDir()
		ca.Debug = cmds.Unify.Flags.Debug
		ca.Quiet = cmds.Unify.Flags.Quiet

		// if Cmd.GoogleSheetUpdate {
		// 	SunGrow.OutputType = iSolarCloud.TypeGoogle
		// }
		//
		// Git.Error = Cmd.GitSet()
		// if Cmd.Error != nil {
		//	break
		// }
		//
		// ca.Valid = true
	}

	return ca.Error
}

// func (ca *CommandArgs) GitSet() error {
// 	for range Only.Once {
// 		if ca.Git != nil {
// 			break
// 		}
//
// 		ca.Git = mmGit.New()
// 		if ca.Git.Error != nil {
// 			ca.Error = ca.Git.Error
// 			break
// 		}
//
// 		// Cmd.Error = Git.SetAuth(ca.GitUsername, ca.GitPassword)
// 		// if Cmd.Error != nil {
// 		//	break
// 		// }
//
// 		ca.Error = ca.Git.SetKeyFile(ca.GitKeyFile)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.SetToken(ca.GitToken)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.SetRepo(ca.GitRepo)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.SetDir(ca.GitRepoDir)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.SetDiffCmd(ca.GitDiffCmd)
// 		if ca.Error != nil {
// 			break
// 		}
// 	}
//
// 	return ca.Error
// }
//
// func (ca *CommandArgs) GitLs(options ...string) error {
// 	for range Only.Once {
// 		os.Args = []string{"GitLs"}
// 		os.Args = append(os.Args, options...)
// 		ca.Error = os.Chdir(ca.GitRepoDir)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		// ls-go is a standalone GoLang executable,
// 		// but I've modified it to be a package and so directly callable.
// 		ca.Error = lsgo.LsGo()
// 		if ca.Error != nil {
// 			break
// 		}
// 	}
//
// 	return ca.Error
// }
//
// func (ca *CommandArgs) GitSync(msg string, entities ...string) error {
// 	for range Only.Once {
// 		ca.Error = ca.Git.Pull()
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.GitSave(entities...)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.Add(".")
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		if msg == "" {
// 			msg = fmt.Sprintf("Updated %d files.", len(entities))
// 		}
// 		ca.Error = ca.Git.Commit(msg)
// 		if ca.Error != nil {
// 			break
// 		}
//
// 		ca.Error = ca.Git.Push()
// 		if ca.Error != nil {
// 			break
// 		}
// 	}
//
// 	return ca.Error
// }
//
// func (ca *CommandArgs) GitSave(entities ...string) error {
// 	for range Only.Once {
// 		if len(entities) == 0 {
// 			entities = DefaultAreas
// 		}
// 		fmt.Printf("Saving %d entities from the SunGrow to Git...\n", len(entities))
//
// 		// SunGrow.OutputType = iSolarCloud.StringTypeJson
// 		// SunGrow.OutputType = iSolarCloud.TypeJson
//
// 		for _, entity := range entities {
// 			// Remove plurals.
// 			entity = strings.TrimSuffix(entity, "s")
// 			// SunGrow.OutputString = ""
//
// 			switch entity {
// 			case "domain":
// 				ca.Api.Api.SunGrow.Error = ca.Api.Api.SunGrow.Init()
// 			}
// 			if ca.Api.Api.SunGrow.Error != nil {
// 				break
// 			}
//
// 			jf := AddJsonSuffix(entity)
// 			ca.Error = ca.Git.SaveFile(jf, []byte(ca.OutputFile))
// 			if ca.Error != nil {
// 				break
// 			}
// 		}
//
// 		fmt.Printf("Saved %d files.", len(entities))
// 	}
//
// 	return ca.Error
// }
//
// func (ca *CommandArgs) GoogleUpdate(entities ...string) error {
//
// 	for range Only.Once {
// 		// SunGrow.OutputType = iSolarCloud.TypeGoogle
//
// 		if len(entities) == 0 {
// 			entities = DefaultAreas
// 		}
// 		fmt.Printf("Saving %d entities from the SunGrow to Google Docs...\n", len(entities))
//
// 		for _, entity := range entities {
// 			switch entity {
// 			case "domain":
// 				ca.Error = ca.Api.SunGrow.Init()
// 				if ca.Error != nil {
// 					break
// 				}
// 			}
//
// 			// sheet := google.Sheet {
// 			// 	Id:          "",
// 			// 	Credentials: nil,
// 			// 	SheetName:   entity,
// 			// 	Range:       "",
// 			// 	Data:        ca.OutputFile,
// 			// }
// 			// sheet.Set(sheet)
// 			// ca.Error = sheet.WriteSheet()
// 			if ca.Error != nil {
// 				break
// 			}
// 		}
// 	}
//
// 	return ca.Error
// }

// func (ca *CommandArgs) UpdateGoogleSheet(name string, data [][]interface{}) error {
//
//	for range Only.Once {
//		sheet := google.Sheet{
//			Id:          "",
//			Credentials: nil,
//			SheetName:   name,
//			Range:       "",
//			Data:        data,
//		}
//		sheet.Set(sheet)
//		ca.Error = sheet.WriteSheet()
//		if ca.Error != nil {
//			break
//		}
//	}
//
//	return p.Error
// }

func AddJsonSuffix(fn string) string {
	return strings.TrimSuffix(fn, ".json") + ".json"
}
