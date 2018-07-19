package handler

import (
	"autodeploy/config"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type GitPushStruct struct {
	Secret     string
	Repository *repositoryStruct
}

type repositoryStruct struct {
	Name string
}

func Gitea(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("get gitea request body err: ", err, r.URL.String())
	}
	var push GitPushStruct
	err = json.Unmarshal(req, &push)
	if err != nil {
		log.Error("parse request err: ", err)
		ResponseErr(w)
	}

	configInfo, ok := config.Config[push.Repository.Name]
	if !ok {
		log.Error("request repository name is wrong")
		ResponseErr(w)
	}

	ExecShell(configInfo.ShellName)

	writer := bufio.NewWriter(w)
	writer.WriteString("hello")
	writer.Flush()
}
