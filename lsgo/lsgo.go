package lsgo

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/acarl005/textcol"
	colorable "github.com/mattn/go-colorable"
	"github.com/willf/pad"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// DisplayItem wraps the file stat info and string to be printed
type DisplayItem struct {
	display  string
	info     os.FileInfo
	basename string
	ext      string
	link     *LinkInfo
}

// LinkInfo wraps link stat info and whether the link points to valid file
type LinkInfo struct {
	path   string
	info   os.FileInfo
	broken bool
}

var (
	// True is a helper varable to help make pointers to `true`
	True       = true
	sizeUnits  = []string{"B", "K", "M", "G", "T"}
	dateFormat = "02.Jan'06" // uses the "reference time" https://golang.org/pkg/time/#Time.Format
	timeFormat = "15:04"
	start      int64                            // keep track of execution time
	stdout     = colorable.NewColorableStdout() // write to this to allow ANSI color codes to be compatible on Windows
)

// func LsGo(cmd *cobra.Command, args []string) {
func LsGo() error {
	var err error

	textcol.Output = stdout

	start = time.Now().UnixNano()
	// auto-generate help text for the command with -h
	kingpin.CommandLine.HelpFlag.Short('h')

	// parse the arguments and populate the struct
	kingpin.Parse()
	argsPostParse()

	// separate the directories from the regular files
	dirs := []string{}
	files := []os.FileInfo{}
	for _, pathStr := range *args.paths {
		var fileStat os.FileInfo
		fileStat, err = os.Stat(pathStr)
		if err != nil && strings.Contains(err.Error(), "no such file or directory") {
			printErrorHeader(err, prettifyPath(pathStr))
			continue
		} else {
			check(err)
		}
		if fileStat.IsDir() {
			dirs = append(dirs, pathStr)
		} else {
			files = append(files, fileStat)
		}
	}

	// list files first
	if len(files) > 0 {
		pwd := os.Getenv("PWD")
		listFiles(pwd, &files, true)
	}

	// then list the contents of each directory
	for i, dir := range dirs {
		// print a blank line between directories, but not before the first one
		if i > 0 {
			fmt.Fprintln(stdout, "")
		}
		listDir(dir)
	}

	return err
}

func listDir(pathStr string) {
	items, err := ioutil.ReadDir(pathStr)
	// if we couldn't read the folder, print a "header" with error message and use error-looking colors
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") || strings.Contains(err.Error(), "permission denied") {
			printErrorHeader(err, prettifyPath(pathStr))
			return
		}
		check(err)
	}

	// filter by the regexp if one was passed
	if len(*args.find) > 0 {
		filteredItems := []os.FileInfo{}
		for _, fileInfo := range items {
			re, err := regexp.Compile(*args.find)
			check(err)
			if re.MatchString(fileInfo.Name()) {
				filteredItems = append(filteredItems, fileInfo)
			}
		}
		items = filteredItems
	}

	if !(len(*args.find) > 0 && len(items) == 0) &&
		!(len(*args.paths) == 1 && (*args.paths)[0] == "." && !*args.recurse) {
		printFolderHeader(pathStr)
	}

	if len(items) > 0 {
		listFiles(pathStr, &items, false)
	}

	if *args.recurse {
		for _, item := range items {
			if item.IsDir() && (item.Name()[0] != '.' || *args.all) {
				fmt.Fprintln(stdout, "") // put a blank line between directories
				listDir(path.Join(pathStr, item.Name()))
			}
		}
	}
}

func listFiles(parentDir string, items *[]os.FileInfo, forceDotfiles bool) {
	absPath, err := filepath.Abs(parentDir)
	check(err)

	// collect all the contents here
	files := []*DisplayItem{}
	dirs := []*DisplayItem{}

	// to help with formatting, we need to know the length of the longest name to add appropriate padding
	longestOwnerName := 0
	longestGroupName := 0
	if *args.owner {
		for _, fileInfo := range *items {
			owner, group := getOwnerAndGroup(&fileInfo)
			longestOwnerName = max(longestOwnerName, len(owner))
			longestGroupName = max(longestGroupName, len(group))
		}
	}

	for _, fileInfo := range *items {
		// if this is a dotfile (hidden file)
		if fileInfo.Name()[0] == '.' {
			// we can skip everything with this file if we aren't using the `all` option
			if !*args.all && !forceDotfiles {
				continue
			}
		}

		basename, ext := splitExt(fileInfo.Name())

		displayItem := DisplayItem{
			info:     fileInfo,
			ext:      ext,
			basename: basename,
		}

		// read some info about linked file if this item is a symlink
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			getLinkInfo(&displayItem, absPath)
		}

		if fileInfo.IsDir() || (fileInfo.Mode()&os.ModeSymlink != 0 &&
			displayItem.link.info != nil &&
			displayItem.link.info.IsDir()) {
			if *args.files {
				continue
			} else {
				dirs = append(dirs, &displayItem)
			}
		} else {
			if *args.dirs {
				continue
			} else {
				files = append(files, &displayItem)
			}
		}

		owner, group := getOwnerAndGroup(&fileInfo)
		ownerColor, groupColor := getOwnerAndGroupColors(owner, group)

		if *args.perms {
			displayItem.display += permString(fileInfo, ownerColor, groupColor)
		}

		if *args.owner {
			paddedOwner := pad.Right(owner, longestOwnerName, " ")
			ownerInfo := []string{Reset + ownerColor + paddedOwner}
			if !*args.nogroup {
				paddedGroup := pad.Right(group, longestGroupName, " ")
				ownerInfo = append(ownerInfo, groupColor+paddedGroup)
			}
			ownerInfo = append(ownerInfo, Reset)
			displayItem.display += strings.Join(ownerInfo, " ")
		}

		if *args.bytes {
			if fileInfo.Mode()&os.ModeDevice != 0 {
				displayItem.display += deviceNumbers(path.Join(absPath, fileInfo.Name()))
			} else {
				displayItem.display += sizeString(fileInfo.Size())
			}
		}

		if *args.mdate {
			displayItem.display += timeString(fileInfo.ModTime())
		}

		displayItem.display += nameString(&displayItem)

		if *args.links && fileInfo.Mode()&os.ModeSymlink != 0 {
			displayItem.display += linkString(&displayItem, absPath)
		}
	}

	if *args.sortTime {
		sort.Sort(ByTime(dirs))
		sort.Sort(ByTime(files))
		if *args.backwards {
			reverse(dirs)
			reverse(files)
		}
	}

	if *args.sortSize {
		sort.Sort(BySize(files))
		if *args.backwards {
			reverse(files)
		}
	}

	if *args.sortKind {
		sort.Sort(ByKind(files))
		if *args.backwards {
			reverse(files)
		}
	}

	// combine the items together again after sorting
	allItems := append(dirs, files...)

	// if using "long" display, just print one item per line
	if *args.bytes || *args.mdate || *args.owner || *args.perms || *args.long {
		for _, item := range allItems {
			fmt.Fprintln(stdout, item.display)
		}
	} else {
		// but if not, try to format in columns, link `ls` would
		strs := []string{}
		for _, item := range allItems {
			strs = append(strs, item.display)
		}
		textcol.PrintColumns(&strs, 2)
	}

	if *args.stats {
		printStats(len(files), len(dirs))
	}
}

func getLinkInfo(item *DisplayItem, absPath string) {
	fullPath := path.Join(absPath, item.info.Name())
	linkPath, err1 := os.Readlink(fullPath)
	check(err1)

	linkFullPath := linkPath
	if linkPath[0] != '/' {
		linkFullPath = path.Join(absPath, linkPath)
	}

	linkInfo, err2 := os.Stat(linkFullPath)
	if *args.linkRel {
		linkRel, _ := filepath.Rel(absPath, linkPath)
		if linkRel != "" && len(linkRel) <= len(linkPath) {
			// i prefer the look of these relative paths prepended with ./
			if linkRel[0] != '.' {
				linkPath = "./" + linkRel
			} else {
				linkPath = linkRel
			}
		}
	}

	link := LinkInfo{
		path: linkPath,
	}
	item.link = &link
	if linkInfo != nil {
		link.info = linkInfo
	} else if strings.Contains(err2.Error(), "no such file or directory") {
		link.broken = true
	} else if !strings.Contains(err2.Error(), "permission denied") {
		check(err2)
	}
}

func nameString(item *DisplayItem) string {
	mode := item.info.Mode()
	name := item.info.Name()
	if mode&os.ModeDir != 0 {
		return dirString(item)
	} else if mode&os.ModeSymlink != 0 {
		if !item.link.broken && item.link.info.IsDir() {
			color := ConfigColor["link"]["nameDir"]
			if *args.nerdfont {
				var linkIcon string
				if item.link.broken {
					linkIcon = otherIcons["brokenLink"]
				} else {
					linkIcon = otherIcons["linkDir"]
				}
				return color + linkIcon + " " + name + " " + Reset
			} else if *args.icons {
				return color + "🔗 " + name + " " + Reset
			} else {
				return color + " " + name + " " + Reset
			}
		} else {
			color := ConfigColor["link"]["name"]
			if *args.nerdfont {
				var linkIcon string
				if item.link.broken {
					linkIcon = otherIcons["brokenLink"]
				} else {
					linkIcon = otherIcons["link"]
				}
				return color + linkIcon + " " + name + " " + Reset
			} else if *args.icons {
				return color + "🔗 " + name + " " + Reset
			} else {
				return color + name + " " + Reset
			}
		}
	} else if mode&os.ModeDevice != 0 {
		color := ConfigColor["device"]["name"]
		if *args.nerdfont {
			return color + otherIcons["device"] + " " + name + " " + Reset
		} else if *args.icons {
			return color + "💽 " + name + " " + Reset
		} else {
			return color + " " + name + " " + Reset
		}
	} else if mode&os.ModeNamedPipe != 0 {
		return ConfigColor["pipe"]["name"] + " " + name + " " + Reset
	} else if mode&os.ModeSocket != 0 {
		return ConfigColor["socket"]["name"] + " " + name + " " + Reset
	}
	return fileString(item)
}

func linkString(item *DisplayItem, absPath string) string {
	colors := ConfigColor["link"]
	displayStrings := []string{colors["arrow"] + "►"}
	if item.link.info == nil && item.link.broken {
		displayStrings = append(displayStrings, colors["broken"]+item.link.path+Reset)
	} else if item.link.info != nil {
		linkname, linkext := splitExt(item.link.path)
		displayItem := DisplayItem{
			info:     item.link.info,
			basename: linkname,
			ext:      linkext,
		}
		displayStrings = append(displayStrings, nameString(&displayItem))
	} else {
		displayStrings = append(displayStrings, item.link.path)
	}
	return strings.Join(displayStrings, " ")
}

func fileString(item *DisplayItem) string {
	key := strings.ToLower(item.ext)
	// figure out which color to choose
	colors := FileColor["_default"]
	alias, hasAlias := FileAliases[key]
	if hasAlias {
		key = alias
	}
	betterColor, hasBetterColor := FileColor[key]
	if hasBetterColor {
		colors = betterColor
	}

	ext := item.ext
	if ext != "" {
		ext = "." + ext
	}

	// in some cases files have icons if front
	// if nerd font enabled, then it'll be a file-specific icon, or if its an executable script, a little shell icon
	// if the regular --icons flag is used instead, then it will show a ">_" only if the file is executable
	icon := ""
	executable := isExecutableScript(item)
	if *args.nerdfont {
		if executable {
			icon = colors[0] + getIconForFile("", "shell") + " "
		} else {
			icon = colors[0] + getIconForFile(item.basename, item.ext) + " "
		}
	} else if *args.icons {
		if executable {
			icon = BgGray(1) + FgRGB(0, 5, 0) + ">_" + Reset + " "
		}
	}
	displayStrings := []string{icon, colors[0], item.basename, colors[1], ext, Reset}
	return strings.Join(displayStrings, "")
}

// check for executable permissions
func isExecutableScript(item *DisplayItem) bool {
	if item.info.Mode()&0111 != 0 && item.info.Mode().IsRegular() {
		return true
	}
	return false
}

func dirString(item *DisplayItem) string {
	colors := ConfigColor["dir"]
	if item.basename == "" {
		colors = ConfigColor[".dir"]
	}
	displayStrings := []string{colors["name"]}
	icon := ""
	if *args.icons {
		displayStrings = append(displayStrings, "📂 ")
	} else if *args.nerdfont {
		icon = getIconForFolder(item.info.Name()) + " "
		displayStrings = append(displayStrings, icon)
	} else {
		displayStrings = append(displayStrings, " ")
	}
	ext := item.ext
	if ext != "" {
		ext = "." + ext
	}
	displayStrings = append(displayStrings, item.basename, colors["ext"], ext, " ", Reset)
	return strings.Join(displayStrings, "")
}

func rwxString(mode os.FileMode, i uint, color string) string {
	bits := mode >> (i * 3)
	coloredStrings := []string{color}
	if bits&4 != 0 {
		coloredStrings = append(coloredStrings, "r")
	} else {
		coloredStrings = append(coloredStrings, "-")
	}
	if bits&2 != 0 {
		coloredStrings = append(coloredStrings, "w")
	} else {
		coloredStrings = append(coloredStrings, "-")
	}
	if i == 0 && mode&os.ModeSticky != 0 {
		if bits&1 != 0 {
			coloredStrings = append(coloredStrings, "t")
		} else {
			coloredStrings = append(coloredStrings, "T")
		}
	} else {
		if bits&1 != 0 {
			coloredStrings = append(coloredStrings, "x")
		} else {
			coloredStrings = append(coloredStrings, "-")
		}
	}
	return strings.Join(coloredStrings, "")
}

// generates the permissions string, ya know like "drwxr-xr-x" and stuff like that
func permString(info os.FileInfo, ownerColor string, groupColor string) string {
	defaultColor := PermsColor["other"]["_default"]

	// info.Mode().String() does not produce the same output as `ls`, so we must build that string manually
	mode := info.Mode()
	// this "type" is not the file extension, but type as far as the OS is concerned
	filetype := "-"
	if mode&os.ModeDir != 0 {
		filetype = "d"
	} else if mode&os.ModeSymlink != 0 {
		filetype = "l"
	} else if mode&os.ModeDevice != 0 {
		if mode&os.ModeCharDevice == 0 {
			filetype = "b" // block device
		} else {
			filetype = "c" // character device
		}
	} else if mode&os.ModeNamedPipe != 0 {
		filetype = "p"
	} else if mode&os.ModeSocket != 0 {
		filetype = "s"
	}
	coloredStrings := []string{defaultColor, filetype}
	coloredStrings = append(coloredStrings, rwxString(mode, 2, ownerColor))
	coloredStrings = append(coloredStrings, rwxString(mode, 1, groupColor))
	coloredStrings = append(coloredStrings, rwxString(mode, 0, defaultColor), Reset, Reset)
	return strings.Join(coloredStrings, " ")
}

func sizeString(size int64) string {
	sizeFloat := float64(size)
	for i, unit := range sizeUnits {
		base := math.Pow(1024, float64(i))
		if sizeFloat < base*1024 {
			var sizeStr string
			if i == 0 {
				sizeStr = strconv.FormatInt(size, 10)
			} else {
				sizeStr = fmt.Sprintf("%.2f", sizeFloat/base)
			}
			return SizeColor[unit] + pad.Left(sizeStr, 6, " ") + unit + " " + Reset
		}
	}
	return strconv.Itoa(int(size))
}

func timeString(modtime time.Time) string {
	dateStr := modtime.Format(dateFormat)
	timeStr := modtime.Format(timeFormat)
	hour, err := strconv.Atoi(timeStr[0:2])
	check(err)
	// generate a color based on the hour of the day. darkest around midnight and whitest around noon
	timeColor := 14 - int(8*math.Cos(math.Pi*float64(hour)/12))
	colored := []string{FgGray(22) + dateStr, FgGray(timeColor) + timeStr, Reset}
	return strings.Join(colored, " ")
}

// when we list out any subdirectories, print those paths conspicuously above the contents
// this helps with visual separation
func printFolderHeader(pathStr string) {
	colors := ConfigColor["folderHeader"]
	headerString := colors["arrow"] + "►" + colors["main"] + " "
	prettyPath := prettifyPath(pathStr)

	if prettyPath == "/" {
		headerString += "/"
	} else {
		folders := strings.Split(prettyPath, "/")
		coloredFolders := make([]string, 0, len(folders))
		for i, folder := range folders {
			if i == len(folders)-1 { // different color for the last folder in the path
				coloredFolders = append(coloredFolders, colors["lastFolder"]+folder)
			} else {
				coloredFolders = append(coloredFolders, colors["main"]+folder)
			}
		}
		headerString += strings.Join(coloredFolders, colors["slash"]+"/")
	}

	fmt.Fprintln(stdout, headerString+" "+Reset)
}

func printErrorHeader(err error, pathStr string) {
	fmt.Fprintln(stdout, ConfigColor["folderHeader"]["error"]+"► "+pathStr+Reset)
	fmt.Fprintln(stdout, err.Error())
}

func prettifyPath(pathStr string) string {
	prettyPath, err := filepath.Abs(pathStr)
	check(err)
	pwd := os.Getenv("PWD")
	home := os.Getenv("HOME")

	if strings.HasPrefix(prettyPath, pwd) {
		prettyPath = "." + prettyPath[len(pwd):]
	} else if strings.HasPrefix(prettyPath, home) {
		prettyPath = "~" + prettyPath[len(home):]
	}
	return prettyPath
}

func getOwnerAndGroupColors(owner string, group string) (string, string) {
	if owner == os.Getenv("USER") {
		owner = "_self"
	}
	ownerColor := PermsColor["user"][owner]
	if ownerColor == "" {
		ownerColor = PermsColor["user"]["_default"]
	}
	groupColor := PermsColor["group"][group]
	if groupColor == "" {
		groupColor = PermsColor["group"]["_default"]
	}
	return ownerColor, groupColor
}

func printStats(numFiles, numDirs int) {
	colors := ConfigColor["stats"]
	end := time.Now().UnixNano()
	microSeconds := (end - start) / int64(time.Microsecond)
	milliSeconds := float64(microSeconds) / 1000
	statStrings := []string{
		colors["text"],
		colors["number"] + strconv.Itoa(numDirs),
		colors["text"] + "dirs",
		colors["number"] + strconv.Itoa(numFiles),
		colors["text"] + "files",
		colors["ms"] + fmt.Sprintf("%.2f", milliSeconds),
		colors["text"] + "ms",
		Reset,
	}
	fmt.Fprintln(stdout, strings.Join(statStrings, " "))
}

func splitExt(filename string) (basepath, ext string) {
	basename := filepath.Base(filename)
	if basename[0] == '.' {
		ext = basename[1:]
		basepath = filename[:len(filename)-len(ext)-1]
	} else {
		ext = filepath.Ext(filename)
		basepath = filename[:len(filename)-len(ext)]
		if ext != "" {
			ext = ext[1:]
		}
	}
	return
}

// Go doesn't provide a `Max` function for ints like it does for floats (wtf?)
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
