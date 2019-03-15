package handler

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

/*
	执行 shell 文件
*/

func ExecShell(path, shellName string) error {
	shellPath := path
	if path[len(path)-1] != byte(os.PathSeparator) {
		shellPath += string(os.PathSeparator)
	}
	shellPath += shellName
	if exist := fileExist(shellPath); !exist {
		return fmt.Errorf("shell file not exist, %s\n", shellPath)
	}
	cmd := exec.Command("/bin/bash", shellPath)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("exec comment %s err: %s, %s", shellPath, string(output), err)
	}
	return nil
}

func fileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ResponseErr(w http.ResponseWriter, msg string) {
	w.WriteHeader(501)
	if msg == "" {
		w.Write([]byte{})
	} else {
		w.Write([]byte(msg))
	}
}

// PushInfo push 过来的数据
type PushInfoStruct struct {
	Secret      string
	Branch      string
	ProjectAddr string
}
