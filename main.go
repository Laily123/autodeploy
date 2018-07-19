package main

import (
	"autodeploy/config"
	"autodeploy/handler"
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port       = ""
	configFile = ""
)

func init() {
	flag.StringVar(&port, "port", "8000", "server port")
	flag.StringVar(&configFile, "config", "./config.toml", "config file path")
	flag.Parse()
}

func main() {
	http.HandleFunc("/gitea", handler.Gitea)
	config.ParseConfig(configFile)
	log.Info("server start: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
