// +build !windows

package lsgo

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/willf/pad"
	"golang.org/x/sys/unix"
)

func getOwnerAndGroup(fileInfo *os.FileInfo) (string, string) {
	statT := (*fileInfo).Sys().(*syscall.Stat_t)
	uid := fmt.Sprint(statT.Uid)
	gid := fmt.Sprint(statT.Gid)
	owner, err := user.LookupId(uid)
	var ownerName string
	if err == nil {
		ownerName = owner.Username
	} else {
		ownerName = uid
	}

	group, err := user.LookupGroupId(gid)
	var groupName string
	if err == nil {
		groupName = group.Name
	} else {
		groupName = gid
	}
	return ownerName, groupName
}

func deviceNumbers(absPath string) string {
	stat := syscall.Stat_t{}
	err := syscall.Stat(absPath, &stat)
	check(err)
	major := strconv.FormatInt(int64(unix.Major(uint64(stat.Rdev))), 10)
	minor := strconv.FormatInt(int64(unix.Minor(uint64(stat.Rdev))), 10)
	return pad.Left(strings.Join([]string{major, minor}, ","), 7, " ") + " " + Reset
}
