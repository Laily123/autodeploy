package feature

import (
	"autodeploy/config"
	"autodeploy/handler"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLog() {
	file, err := os.OpenFile("./logs/autodeploy.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("open log file err: ", err)
	}
	out := io.MultiWriter(file, os.Stdout)
	log.SetOutput(out)
}

func Server(configPath string) {
	initLog()

	if configPath == "" {
		log.Fatal("config path error")
	}
	config.ParseConfig(configPath)

	//http.HandleFunc("/gitea", handler.Gitea)
	http.HandleFunc("/ding", handler.Ding)
	log.Info("server start: ", config.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, nil))
}
