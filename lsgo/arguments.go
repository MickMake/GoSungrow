package lsgo

import (
	"log"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// declare the struct that holds all the arguments
type arguments struct {
	paths     *[]string
	all       *bool
	bytes     *bool
	mdate     *bool
	owner     *bool
	nogroup   *bool
	perms     *bool
	long      *bool
	dirs      *bool
	files     *bool
	links     *bool
	linkRel   *bool
	sortSize  *bool
	sortTime  *bool
	sortKind  *bool
	backwards *bool
	stats     *bool
	icons     *bool
	nerdfont  *bool
	recurse   *bool
	find      *string
}

var args = arguments{
	kingpin.Arg("paths", "the files(s) and/or folder(s) to display").Default(".").Strings(),
	kingpin.Flag("all", "show hidden files").Short('a').Bool(),
	kingpin.Flag("bytes", "include size").Short('b').Bool(),
	kingpin.Flag("mdate", "include modification date").Short('m').Bool(),
	kingpin.Flag("owner", "include owner and group").Short('o').Bool(),
	kingpin.Flag("nogroup", "hide group").Short('N').Bool(),
	kingpin.Flag("perms", "include permissions for owner, group, and other").Short('p').Bool(),
	kingpin.Flag("long", "include size, date, owner, and permissions").Short('l').Bool(),
	kingpin.Flag("dirs", "only show directories").Short('d').Bool(),
	kingpin.Flag("files", "only show files").Short('f').Bool(),
	kingpin.Flag("links", "show paths for symlinks").Short('L').Bool(),
	kingpin.Flag("link-rel", "show symlinks as relative paths if shorter than absolute path").Short('R').Bool(),
	kingpin.Flag("size", "sort items by size").Short('s').Bool(),
	kingpin.Flag("time", "sort items by time").Short('t').Bool(),
	kingpin.Flag("kind", "sort items by extension").Short('k').Bool(),
	kingpin.Flag("backwards", "reverse the sort order of --size, --time, or --kind").Short('B').Bool(),
	kingpin.Flag("stats", "show statistics").Short('S').Bool(),
	kingpin.Flag("icons", "show folder icon before dirs").Short('i').Bool(),
	kingpin.Flag("nerd-font", "show nerd font glyphs before file names").Short('n').Bool(),
	kingpin.Flag("recurse", "traverse all dirs recursively").Short('r').Bool(),
	kingpin.Flag("find", "filter items with a regexp").Short('F').String(),
}

func argsPostParse() {
	if *args.long {
		args.bytes = &True
		args.mdate = &True
		args.owner = &True
		args.perms = &True
		args.links = &True
	}
	if *args.dirs && *args.files {
		log.Fatal("--dirs and --files cannot both be set")
	}
	if *args.nerdfont && *args.icons {
		log.Fatal("--nerd-font and --icons cannot both be set")
	}
}
