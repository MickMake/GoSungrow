package main

import (
	"GoSungrow/cmd"
	"fmt"
	"os"
)


// https://augateway.isolarcloud.com/v1/


func main() {
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
