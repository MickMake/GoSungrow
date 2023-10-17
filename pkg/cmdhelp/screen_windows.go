//go:build windows

package cmdhelp

import (
	"os"

	"github.com/MickMake/GoUnify/Only"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWidth() int {
	// ws := &winsize{}
	// retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
	// 	uintptr(syscall.Stdin),
	// 	uintptr(syscall.TIOCGWINSZ),
	// 	uintptr(unsafe.Pointer(ws)))
	//
	// if int(retCode) == -1 {
	// 	panic(errno)
	// }
	return int(256)
}

func isTerminal() bool {
	var yes bool
	for range Only.Once {
		o, _ := os.Stdout.Stat()
		if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			// Is terminal
			yes = true
			break
		}
		// Is pipe
		yes = false
	}
	return yes
}

func isPipe() bool {
	return !isTerminal()
}
