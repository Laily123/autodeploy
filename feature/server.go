package feature

import (
	"autodeploy/handler"
	"net/http"
)

func Server(params []string) {
	http.HandleFunc("/gitea", handler.Gitea)
	//config.ParseConfig(configFile)
	//log.Info("server start: ", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	//file, err := os.OpenFile("./logs/autodeploy.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//if err != nil {
	//	log.Fatal("open log file err: ", err)
	//}
	//out := io.MultiWriter(file, os.Stdout)
	//log.SetOutput(out)

}
