package main

import (
	"autodeploy/feature"
	"os"
)

var (
	port       = ""
	configFile = ""
)

func init() {
	//log.SetLevel(log.DebugLevel)
}

func initParams() {
	args := os.Args[1:]
	switch args[0] {
	case "watch":
		feature.StartWatcher(os.Args[0], args[1:])
	case "server":
		feature.Server(args[1:])
	}
}

func main() {
	initParams()
}
