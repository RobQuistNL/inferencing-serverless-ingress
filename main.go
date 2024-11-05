package main

import (
	log2 "github.com/ipfs/go-log/v2"
	"sii/cmd"
)

var log = log2.Logger("cli")

func main() {
	// Set default log level to info
	// @todo set to debug by default, info only on built versions?
	lvl, _ := log2.LevelFromString("info")
	log2.SetAllLoggers(lvl)

	cmd.Execute()
}
