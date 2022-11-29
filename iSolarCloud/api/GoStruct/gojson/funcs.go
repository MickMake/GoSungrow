// Package gojson - generates go struct defintions from JSON documents
//
// Reads from stdin and prints to stdout
//
// Output:
// 	package main
//
// 	type User struct {
// 		AvatarURL         string      `json:"avatar_url"`
// 		Bio               interface{} `json:"bio"`
// 		Blog              string      `json:"blog"`
// 		Company           string      `json:"company"`
// 		CreatedAt         string      `json:"created_at"`
// 		Email             string      `json:"email"`
// 		EventsURL         string      `json:"events_url"`
// 		Followers         float64     `json:"followers"`
// 		FollowersURL      string      `json:"followers_url"`
// 		Following         float64     `json:"following"`
// 		FollowingURL      string      `json:"following_url"`
// 		GistsURL          string      `json:"gists_url"`
// 		GravatarID        string      `json:"gravatar_id"`
// 		Hireable          bool        `json:"hireable"`
// 		HtmlURL           string      `json:"html_url"`
// 		ID                float64     `json:"id"`
// 		Location          string      `json:"location"`
// 		Login             string      `json:"login"`
// 		Name              string      `json:"name"`
// 		OrganizationsURL  string      `json:"organizations_url"`
// 		PublicGists       float64     `json:"public_gists"`
// 		PublicRepos       float64     `json:"public_repos"`
// 		ReceivedEventsURL string      `json:"received_events_url"`
// 		ReposURL          string      `json:"repos_url"`
// 		StarredURL        string      `json:"starred_url"`
// 		SubscriptionsURL  string      `json:"subscriptions_url"`
// 		Type              string      `json:"type"`
// 		UpdatedAt         string      `json:"updated_at"`
// 		URL               string      `json:"url"`
// 	}
package gojson

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"io"
	"os"
	// . "github.com/ChimeraCoder/gojson"
)


// var (
// 	name        = flag.String("name", "Foo", "the name of the struct")
// 	pkg         = flag.String("pkg", "main", "the name of the package for the generated code")
// 	inputName   = flag.String("input", "", "the name of the input file containing JSON (if input not provided via STDIN)")
// 	outputName  = flag.String("o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
// 	format      = flag.String("fmt", "json", "the format of the input data (json or yaml, defaults to json)")
// 	tags        = flag.String("tags", "fmt", "comma seperated list of the tags to put on the struct, default is the same as fmt")
// 	forceFloats = flag.Bool("forcefloats", false, "[experimental] force float64 type for integral values")
// 	subStruct   = flag.Bool("subStruct", false, "create types for sub-structs (default is false)")
// )

var formatYaml = "yaml"	// true
var formatJson = "json"	// false

type Options struct {
	inputName     *string
	outputName    *string

	saveOutput    bool
	structureName *string
	packageName   *string
	format        string
	tags          []string
	forceFloats   bool
	subStructs    bool
}

func (o *Options) StructureName(name string)  {
	o.structureName = &name
}

func (o *Options) PackageName(name string)  {
	o.packageName = &name
}

func (o *Options) InputIsJson()  {
	o.format = formatJson
}

func (o *Options) InputIsYaml()  {
	o.format = formatYaml
}

func (o *Options) Tags(names []string)  {
	o.tags = names
}

func (o *Options) ForceFloats()  {
	o.forceFloats = true
}

func (o *Options) SubStructs()  {
	o.subStructs = true
}

func (o *Options) SetOutput(name string)  {
	o.outputName = &name
}

func (o *Options) SaveOutput()  {
	o.saveOutput = true
}


func Parse(options Options, data []byte) (string, error) {
	var ret string
	var err error

	for range Only.Once {
		if options.format == "" {
			options.format = "json"
		}

		if options.packageName == nil {
			*options.packageName = "Package"
		}
		if *options.packageName == "" {
			*options.packageName = "Package"
		}

		if options.structureName == nil {
			*options.structureName = "Response"
		}
		if *options.structureName == "" {
			*options.structureName = "Response"
		}

		if options.format != "json" && options.format != "yaml" {
			err = errors.New("fmt must be json or yaml")
			break
		}

		tagList := make([]string, 0)
		tagList = append(tagList, options.format)
		if len(options.tags) > 0 {
			tagList = append(tagList, options.tags...)
		}

		var input io.Reader
		for range Only.Once {
			if options.inputName == nil {
				input = bytes.NewReader(data)
				break
			}

			var f *os.File
			f, err = os.Open(*options.inputName)
			if err != nil {
				err = errors.New(fmt.Sprintf("reading input file: %s", err))
			}
			//goland:noinspection GoUnhandledErrorResult
			defer f.Close()
			input = f
		}
		if err != nil {
			break
		}

		var convertFloats bool
		var parser Parser
		switch options.format {
			case "json":
				parser = ParseJson
				convertFloats = true
			case "yaml":
				parser = ParseYaml
		}

		var output []byte
		if output, err = Generate(input, parser, *options.structureName, *options.packageName, tagList, options.subStructs, convertFloats); err != nil {
			err = errors.New(fmt.Sprintf("error parsing: %s", err))
			break
		}

		ret = string(output)

		// if options.outputName == nil {
		// 	fmt.Print(string(output))
		// 	break
		// }
		//
		// err = ioutil.WriteFile(*options.outputName, output, 0644)
		// if err != nil {
		// 	err = errors.New(fmt.Sprintf("writing output: %s", err))
		// 	break
		// }
	}

	return ret, err
}

// // Return true if os.Stdin appears to be interactive
// func isInteractive() bool {
// 	fileInfo, err := os.Stdin.Stat()
// 	if err != nil {
// 		return false
// 	}
// 	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
// }
