package main

import (
	"GoSungrow/cmd"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"os"
)


// https://augateway.isolarcloud.com/v1/


func main() {
	var err error

	for range Only.Once {
		err = cmd.Execute()
		if err != nil {
			break
		}
	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
