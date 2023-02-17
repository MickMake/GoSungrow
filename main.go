package main

import (
	"fmt"
	"os"

	"github.com/MickMake/GoSungrow/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
