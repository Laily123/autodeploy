package main

import (
	"autodeploy/config"
	"autodeploy/handler"
	"flag"
	"io"
	"net/http"
	"os"

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
	file, err := os.OpenFile("./logs/autodeploy.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("open log file err: ", err)
	}
	out := io.MultiWriter(file, os.Stdout)
	log.SetOutput(out)

}

func main() {
	http.HandleFunc("/gitea", handler.Gitea)
	config.ParseConfig(configFile)
	log.Info("server start: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
