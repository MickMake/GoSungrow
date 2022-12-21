package defaults

import _ "embed"

// Need to execute `go generate -v -x defaults/const.go` OR `go generate -v -x ./...`
//go:generate cp ../README.md README.md
//go:generate cp ../EXAMPLES.md EXAMPLES.md
// //go:generate cp ../CHANGELOG.md CHANGELOG.md

//go:embed README.md
var Readme string

//go:embed EXAMPLES.md
var Examples string

// //go:embed CHANGELOG.md
// var Changelog string

const (
	Description   = "GoSungrow - GoLang implementation to access the iSolarCloud API updated by SunGrow inverters"
	BinaryName    = "GoSungrow"
	BinaryVersion = "3.0.2"
	SourceRepo    = "github.com/MickMake/" + BinaryName
	BinaryRepo    = "github.com/MickMake/" + BinaryName

	EnvPrefix     = "GOSUNGROW"

	HelpSummary   = `
# GoSungrow - GoLang implementation to access the iSolarCloud API updated by SunGrow inverters.

This GoLang package has a complete implementation of the iSolarCloud API.
There's been no published specs on this, so I've had to figure it all out based on the Android app, using javascript IDEs and various other means.

`
)
