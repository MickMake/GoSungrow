module GoSungrow

go 1.19

//replace github.com/MickMake/GoUnify => ../../GoUnify
//replace github.com/MickMake/GoUnify/cmdConfig => ../../GoUnify/cmdConfig
//replace github.com/MickMake/GoUnify/cmdLog => ../../GoUnify/cmdLog
//replace github.com/MickMake/GoUnify/cmdHelp => ../../GoUnify/cmdHelp
//replace github.com/MickMake/GoUnify/Unify => ../../GoUnify/Unify
//replace github.com/MickMake/GoUnify/Only => ../../GoUnify/Only
//replace github.com/MickMake/GoUnify/cmdCron => ../../GoUnify/cmdCron
//replace github.com/MickMake/GoUnify/cmdDaemon => ../../GoUnify/cmdDaemon
//replace github.com/MickMake/GoUnify/cmdShell => ../../GoUnify/cmdShell
//replace github.com/MickMake/GoUnify/cmdVersion => ../../GoUnify/cmdVersion
//replace github.com/MickMake/GoUnify/cmdExec => ../../GoUnify/cmdExec
//replace github.com/MickMake/GoUnify/cmdPath => ../../GoUnify/cmdPath

require (
	github.com/MickMake/GoUnify/Only v0.0.0-20221125023651-ff4a37b1928a
	github.com/MickMake/GoUnify/Unify v0.0.0-20221125023651-ff4a37b1928a
	github.com/MickMake/GoUnify/cmdConfig v0.0.0-20221125023651-ff4a37b1928a
	github.com/MickMake/GoUnify/cmdHelp v0.0.0-20221125023651-ff4a37b1928a
	github.com/MickMake/GoUnify/cmdLog v0.0.0-20221125023651-ff4a37b1928a
	github.com/MickMake/GoUnify/cmdPath v0.0.0-20221125023651-ff4a37b1928a
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee
	github.com/eclipse/paho.mqtt.golang v1.4.2
	github.com/go-co-op/gocron v1.18.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.14.0
	github.com/wcharczuk/go-chart/v2 v2.1.0
	go.pennock.tech/tabular v1.1.3
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/MichaelMure/go-term-markdown v0.1.4 // indirect
	github.com/MichaelMure/go-term-text v0.3.1 // indirect
	github.com/MickMake/GoUnify/cmdCron v0.0.0-20221125020154-1b15b8d20735 // indirect
	github.com/MickMake/GoUnify/cmdDaemon v0.0.0-20221125023651-ff4a37b1928a // indirect
	github.com/MickMake/GoUnify/cmdExec v0.0.0-20221125015223-b8c165efd0ec // indirect
	github.com/MickMake/GoUnify/cmdShell v0.0.0-20221125023651-ff4a37b1928a // indirect
	github.com/MickMake/GoUnify/cmdVersion v0.0.0-20221125023651-ff4a37b1928a // indirect
	github.com/abiosoft/ishell v2.0.0+incompatible // indirect
	github.com/abiosoft/ishell/v2 v2.0.2 // indirect
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/alecthomas/chroma v0.7.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/briandowns/spinner v1.19.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.1.6 // indirect
	github.com/eliukblau/pixterm/pkg/ansimage v0.0.0-20191210081756-9fb6cf8c2f75 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomarkdown/markdown v0.0.0-20191123064959-2c17d62f5098 // indirect
	github.com/google/go-github/v30 v30.1.0 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/ivanpirog/coloredcobra v1.0.1 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kyokomi/emoji/v2 v2.2.8 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/rhysd/go-github-selfupdate v1.2.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sevlyar/go-daemon v0.1.6 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/tcnksm/go-gitconfig v0.1.2 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/ulikunitz/xz v0.5.9 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5 // indirect
	golang.org/x/net v0.0.0-20221014081412-f15817d10f9b // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.0.0-20220908164124-27713097b956 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
