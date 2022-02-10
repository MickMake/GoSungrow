package main

import (
	"GoSungro/cmd"
	"fmt"
	"os"
)

// https://augateway.isolarcloud.com/v1/

func main() {
	var err error

	for range OnceOnly {
		err = cmd.Execute()
		if err != nil {
			break
		}

	}

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
}
