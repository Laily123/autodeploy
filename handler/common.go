package handler

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func ExecShell(shellPath string) {
	if exist := fileExist(shellPath); !exist {
		log.Error("shell file not exist")
		return
	}
	cmd := exec.Command("/bin/sh", shellPath)
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("exec comment %s err:]\n%s\n", shellPath, err)
		return
	}
	log.Infof("exec command %s succ:\n%s\n", shellPath, string(output))
}

func fileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
