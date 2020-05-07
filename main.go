package main

import (
	"os"

	"github.com/imadu/blockchain_app/cli"
)

func main() {
	defer os.Exit(0)
	command := cli.CommandLine{}
	command.Run()
}
