package handler

import (
	"autodeploy/config"
	"bufio"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func Gitea(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("get gitea request body err: %s, url: %s\n", err, r.URL.String())
		ResponseErr(w)
		return
	}
	pushData := string(req)
	pushInfo := getPushData(pushData)
	configInfo, ok := config.Config[pushInfo.ProjectAddr]
	if !ok {
		log.Errorf("request repository addr is not match, get %s\n", pushInfo.ProjectAddr)
		ResponseErr(w)
		return
	}
	if pushInfo.Secret != configInfo.Secret {
		log.Errorf("secret not match\n")
		ResponseErr(w)
		return
	}

	ExecShell(configInfo.Dir, configInfo.ShellName)

	writer := bufio.NewWriter(w)
	writer.WriteString("hello")
	writer.Flush()
}

// getPushData 从请求里获取需要的数据
// gitea 的 webhook 数据在 https://docs.gitea.io/en-us/webhooks/
func getPushData(data string) *PushInfoStruct {
	info := &PushInfoStruct{}
	info.Secret = gjson.Get(data, "secret").String()
	info.ProjectAddr = gjson.Get(data, "repository.html_url").String()
	branchInfo := gjson.Get(data, "refs").String()
	arr := strings.Split(branchInfo, "/")
	info.Branch = arr[len(arr)-1]
	log.Debugf("get push info: %+v\n", info)
	return info
}
