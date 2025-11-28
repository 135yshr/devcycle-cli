package main

import (
	"os"

	"github.com/135yshr/devcycle-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
