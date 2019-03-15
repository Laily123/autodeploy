package handler

import (
	"autodeploy/config"
	"net/http"
	"os"
	"strings"
)

/*
	处理通用 webhook 触发
	url 的 query 参数 key 代表触发 key，然后从配置文件获得 shell 脚本的执行路径，执行 shell 脚本
*/
var (
	shellMap = make(map[string]string, 0)
)

func initShellMap() {
	for _, v := range config.App.Dings {
		arr := strings.Split(v, ";")
		if len(arr) == 2 {
			shellMap[arr[0]] = arr[1]
		}
	}
}

func Ding(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		ResponseErr(w, "")
		return
	}
	initShellMap()
	if filepath, ok := shellMap[key]; !ok {
		ResponseErr(w, "key is not exist")
	} else {
		arr := strings.Split(filepath, string(os.PathSeparator))
		if len(arr) <= 1 {
			ResponseErr(w, "shell path error")
			return
		}
		shellName := arr[len(arr)-1]
		path := strings.TrimRight(filepath, shellName)
		err := ExecShell(path, shellName)
		if err != nil {
			ResponseErr(w, err.Error())
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
