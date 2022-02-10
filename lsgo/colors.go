package lsgo

import (
	"fmt"
	"strconv"
	"strings"
)

// see the color codes
// http://i.stack.imgur.com/UQVe5.png

// Fg wraps an 8-bit foreground color code in the ANSI escape sequence
func Fg(code int) string {
	colored := []string{"\x1b[38;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

// Bg wraps an 8-bit background color code in the ANSI escape sequence
func Bg(code int) string {
	colored := []string{"\x1b[48;5;", strconv.Itoa(code), "m"}
	return strings.Join(colored, "")
}

// Rgb2code converts RGB values (up to 5) to an 8-bit color code
func Rgb2code(r int, g int, b int) int {
	code := 36*r + 6*g + b + 16
	if code < 16 || 231 < code {
		panic(fmt.Errorf("Invalid RGB values (%d, %d, %d)", r, g, b))
	}
	return code
}

// Gray2code converts a scalar of "grayness" to an 8-bit color code
func Gray2code(lightness int) int {
	code := lightness + 232
	if code < 232 || 255 < code {
		panic(fmt.Errorf("Invalid lightness value (%d) for gray", lightness))
	}
	return code
}

// FgRGB converts RGB values (up to 5) to an ANSI-escaped foreground 8-bit color code
func FgRGB(r int, g int, b int) string {
	return Fg(Rgb2code(r, g, b))
}

// BgRGB converts RGB values (up to 5) to an ANSI-escaped background 8-bit color code
func BgRGB(r int, g int, b int) string {
	return Bg(Rgb2code(r, g, b))
}

// FgGray converts a scalar of "grayness" to an ANSI-escaped foreground 8-bit color code
func FgGray(lightness int) string {
	return Fg(Gray2code(lightness))
}

// BgGray converts a scalar of "grayness" to an ANSI-escaped background 8-bit color code
func BgGray(lightness int) string {
	return Bg(Gray2code(lightness))
}

const (
	// Reset undoes ANSI color codes
	Reset = "\x1b[0m"
	// Bold makes text bold
	Bold = "\x1b[1m"
)

var (
	// FileColor is a mapping of filetypes to colors
	FileColor = map[string][2]string{
		"as":      [2]string{Fg(196), Fg(124)},
		"asm":     [2]string{Fg(223), Fg(215)},
		"bf":      [2]string{Fg(223), Fg(215)},
		"c":       [2]string{Fg(39), Fg(27)},
		"clj":     [2]string{Fg(204), Fg(162)},
		"coffee":  [2]string{Fg(136), Fg(94)},
		"cr":      [2]string{Fg(82), Fg(70)},
		"cson":    [2]string{Fg(136), Fg(94)},
		"css":     [2]string{Fg(219), Fg(207)},
		"dart":    [2]string{Fg(43), Fg(31)},
		"diff":    [2]string{Fg(10), Fg(9)},
		"elm":     [2]string{Fg(51), Fg(39)},
		"erl":     [2]string{Fg(161), Fg(89)},
		"ex":      [2]string{Fg(99), Fg(57)},
		"f":       [2]string{Fg(208), Fg(94)},
		"fs":      [2]string{Fg(45), Fg(32)},
		"gb":      [2]string{Fg(43), Fg(29)},
		"go":      [2]string{Fg(121), Fg(109)},
		"graphql": [2]string{Fg(219), Fg(207)},
		"groovy":  [2]string{Fg(223), Fg(215)},
		"gv":      [2]string{Fg(141), Fg(99)},
		"hs":      [2]string{Fg(99), Fg(57)},
		"html":    [2]string{Fg(87), Fg(73)},
		"ino":     [2]string{Fg(43), Fg(29)},
		"java":    [2]string{Fg(136), Fg(94)},
		"jl":      [2]string{Fg(141), Fg(99)},
		"js":      [2]string{FgRGB(4, 4, 0), FgRGB(2, 2, 0)},
		"jsx":     [2]string{Fg(87), Fg(73)},
		"lock":    [2]string{FgGray(8), FgGray(5)},
		"log":     [2]string{FgGray(8), FgGray(5)},
		"lua":     [2]string{Fg(39), Fg(27)},
		"m":       [2]string{Fg(208), Fg(196)},
		"md":      [2]string{Fg(87), Fg(73)},
		"ml":      [2]string{Fg(208), Fg(94)},
		"php":     [2]string{Fg(30), Fg(22)},
		"pl":      [2]string{Fg(99), Fg(57)},
		"py":      [2]string{Fg(34), Fg(28)},
		"r":       [2]string{Fg(51), Fg(39)},
		"rb":      [2]string{FgRGB(5, 1, 0), FgRGB(3, 1, 1)},
		"rs":      [2]string{Fg(208), Fg(94)},
		"scala":   [2]string{Fg(196), Fg(124)},
		"sh":      [2]string{FgRGB(4, 0, 4), FgRGB(2, 0, 2)},
		"sol":     [2]string{Fg(39), Fg(27)},
		"sql":     [2]string{Fg(193), Fg(148)},
		"svelte":  [2]string{Fg(208), Fg(196)},
		"swift":   [2]string{Fg(223), Fg(215)},
		"vim":     [2]string{Fg(34), Fg(28)},
		"vue":     [2]string{Fg(43), Fg(29)},
		"xml":     [2]string{Fg(87), Fg(73)},

		"compiled": [2]string{FgGray(8), FgGray(5)},
		"compress": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"document": [2]string{FgRGB(5, 0, 0), FgRGB(3, 0, 0)},
		"media":    [2]string{Fg(141), Fg(99)},
		"_default": [2]string{FgGray(23), FgGray(12)},
	}
	// FileAliases converts alternative extensions to their canonical mapping in FileColor
	FileAliases = map[string]string{
		"s":        "asm",
		"b":        "bf",
		"c++":      "c",
		"cc":       "c",
		"cpp":      "c",
		"cs":       "c",
		"cxx":      "c",
		"d":        "c",
		"h":        "c",
		"h++":      "c",
		"hh":       "c",
		"hpp":      "c",
		"hxx":      "c",
		"pxd":      "c",
		"cljc":     "clj",
		"cljs":     "clj",
		"class":    "compiled",
		"elc":      "compiled",
		"hi":       "compiled",
		"o":        "compiled",
		"pyc":      "compiled",
		"7z":       "compress",
		"Z":        "compress",
		"bz2":      "compress",
		"deb":      "compress",
		"dmg":      "compress",
		"dpkg":     "compress",
		"gz":       "compress",
		"iso":      "compress",
		"jar":      "compress",
		"lzma":     "compress",
		"par":      "compress",
		"rar":      "compress",
		"rpm":      "compress",
		"tar":      "compress",
		"tc":       "compress",
		"tgz":      "compress",
		"txz":      "compress",
		"whl":      "compress",
		"xz":       "compress",
		"z":        "compress",
		"zip":      "compress",
		"less":     "css",
		"sass":     "css",
		"scss":     "css",
		"styl":     "css",
		"djvu":     "document",
		"doc":      "document",
		"docx":     "document",
		"dvi":      "document",
		"eml":      "document",
		"fotd":     "document",
		"odp":      "document",
		"odt":      "document",
		"pdf":      "document",
		"ppt":      "document",
		"pptx":     "document",
		"rtf":      "document",
		"xls":      "document",
		"xlsx":     "document",
		"f03":      "f",
		"f77":      "f",
		"f90":      "f",
		"f95":      "f",
		"for":      "f",
		"fpp":      "f",
		"ftn":      "f",
		"fsi":      "fs",
		"fsscript": "fs",
		"fsx":      "fs",
		"dna":      "gb",
		"gsh":      "groovy",
		"gvy":      "groovy",
		"gy":       "groovy",
		"htm":      "html",
		"xhtml":    "html",
		"bson":     "js",
		"json":     "js",
		"mjs":      "js",
		"ts":       "js",
		"tsx":      "jsx",
		"cjsx":     "jsx",
		"mat":      "m",
		"markdown": "md",
		"mkd":      "md",
		"rst":      "md",
		"sml":      "ml",
		"mli":      "ml",
		"aac":      "media",
		"alac":     "media",
		"audio":    "media",
		"avi":      "media",
		"bmp":      "media",
		"cbr":      "media",
		"cbz":      "media",
		"eps":      "media",
		"flac":     "media",
		"flv":      "media",
		"gif":      "media",
		"ico":      "media",
		"image":    "media",
		"jpeg":     "media",
		"jpg":      "media",
		"m2v":      "media",
		"m4a":      "media",
		"mka":      "media",
		"mkv":      "media",
		"mov":      "media",
		"mp3":      "media",
		"mp4":      "media",
		"mpeg":     "media",
		"mpg":      "media",
		"nef":      "media",
		"ogg":      "media",
		"ogm":      "media",
		"ogv":      "media",
		"opus":     "media",
		"orf":      "media",
		"pbm":      "media",
		"pgm":      "media",
		"png":      "media",
		"pnm":      "media",
		"ppm":      "media",
		"pxm":      "media",
		"stl":      "media",
		"svg":      "media",
		"tif":      "media",
		"tiff":     "media",
		"video":    "media",
		"vob":      "media",
		"wav":      "media",
		"webm":     "media",
		"webp":     "media",
		"wma":      "media",
		"wmv":      "media",
		"xpm":      "media",
		"php3":     "php",
		"php4":     "php",
		"php5":     "php",
		"phpt":     "php",
		"phtml":    "php",
		"ipynb":    "py",
		"pickle":   "py",
		"pkl":      "py",
		"pyx":      "py",
		"bash":     "sh",
		"csh":      "sh",
		"fish":     "sh",
		"ksh":      "sh",
		"zsh":      "sh",
		"plpgsql":  "sql",
		"plsql":    "sql",
		"psql":     "sql",
		"tsql":     "sql",
		"vimrc":    "vim",
	}
	// SizeColor has the color mappings for ranges of file sizes
	SizeColor = map[string]string{
		"B": Fg(27),
		"K": Fg(33),
		"M": Fg(81),
		"G": Fg(123),
		"T": Fg(159),
	}
	// ConfigColor holds mappings for other various colors
	ConfigColor = map[string]map[string]string{
		"dir": map[string]string{
			"name": Bold + BgRGB(0, 0, 2) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		".dir": map[string]string{
			"name": Bold + BgRGB(0, 0, 1) + FgGray(23),
			"ext":  FgRGB(2, 2, 5),
		},
		"folderHeader": map[string]string{
			"arrow":      FgRGB(3, 2, 0),
			"main":       BgGray(2) + FgRGB(3, 2, 0),
			"slash":      FgGray(5),
			"lastFolder": Bold + FgRGB(5, 5, 0),
			"error":      BgRGB(5, 0, 0) + FgRGB(5, 5, 0),
		},
		"link": map[string]string{
			"name":    Bold + FgRGB(0, 5, 0),
			"nameDir": Bold + BgRGB(0, 0, 2) + FgRGB(0, 5, 5),
			"arrow":   FgRGB(1, 0, 1),
			"path":    FgRGB(4, 0, 4),
			"broken":  FgRGB(5, 0, 0),
		},
		"device": map[string]string{
			"name": Bold + BgGray(3) + Fg(220),
		},
		"socket": map[string]string{
			"name": Bold + Bg(53) + Fg(15),
		},
		"pipe": map[string]string{
			"name": Bold + Bg(94) + Fg(15),
		},
		"stats": map[string]string{
			"text":   BgGray(2) + FgGray(15),
			"number": Fg(24),
			"ms":     Fg(39),
		},
	}
	// PermsColor holds color mappings for users and groups
	PermsColor = map[string]map[string]string{
		"user": map[string]string{
			"root":     FgRGB(5, 0, 2),
			"daemon":   FgRGB(4, 2, 1),
			"_self":    FgRGB(0, 4, 0),
			"_default": FgRGB(0, 3, 3),
		},
		"group": map[string]string{
			"wheel":    FgRGB(3, 0, 0),
			"staff":    FgRGB(0, 2, 0),
			"admin":    FgRGB(2, 2, 0),
			"_default": FgRGB(2, 0, 2),
		},
		"other": map[string]string{
			"_default": BgGray(2) + FgGray(15),
		},
	}
)
