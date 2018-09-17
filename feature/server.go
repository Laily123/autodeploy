package feature

import (
	"autodeploy/config"
	"autodeploy/handler"
	"io"
	"net/http"
	"os"
	"strings"

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

func Server(params []string) {
	initLog()

	if len(params) == 0 {
		log.Fatal("params err")
	}
	configPath := params[0]
	if strings.HasPrefix(configPath, "-config=") {
		configPath = strings.Replace(params[0], "-config=", "", -1)
	} else {
		log.Fatal("config params err")
	}
	config.ParseConfig(configPath)

	http.HandleFunc("/gitea", handler.Gitea)
	log.Info("server start: ", config.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, nil))
}
