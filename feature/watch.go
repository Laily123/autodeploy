package feature

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

/*
	监听文件变化重新执行程序
*/

var (
	processing = false
)

type Watcher struct {
	cmd       *exec.Cmd // exec.Cmd 结构体
	cmdCancel context.CancelFunc
	dir       string // 工作目录
	cmdStr    string // cmd 执行的 shell 命令
	oldPID    int
}

// StartWater 开启一个新的监控
// workDir 工作目录，go main 包所在的目录
// args 额外参数，这里是指定执行的 shell 命令
func StartWatcher(workDir string, args []string) {
	watcher := Watcher{}
	watcher.init(workDir, args)
}
func (w *Watcher) init(pwdDir string, args []string) {
	defer w.kill()
	dir, err := filepath.Abs(filepath.Dir(pwdDir))
	if err != nil {
		log.Fatal("Get dir err: ", err)
	}
	log.Info("Watching directory: ", dir)
	w.dir = dir
	if len(args) > 0 {
		w.cmdStr = strings.Join(args, " ")
	} else {
		w.cmdStr = "go run *.go"
	}
	w.initCmd()
	w.watch()
}

// initCmd 初始化命令
func (w *Watcher) initCmd() *exec.Cmd {
	cmd := exec.Command("/bin/sh", "-c", w.cmdStr)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	go printCmdResult(stdout)
	go printCmdResult(stderr)
	w.kill()
	return cmd
}

func (w *Watcher) watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Get watcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Debug("event: ", event.String())
				go w.execCMD()
				w.addDir(watcher)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Info("watch err: ", err)
			}
		}
	}()

	w.addDir(watcher)
	<-done
}

// 将根目录下的子目录加入监控
func (w *Watcher) addDir(watcher *fsnotify.Watcher) {
	err := filepath.Walk(w.dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Error("walk dir err: ", err)
	}
}

// exec 当文件修改时执行命令
func (w *Watcher) execCMD() {
	// 一个执行后 3s 才能进行新的执行
	if processing {
		return
	}
	processing = true
	go func() {
		time.Sleep(3 * time.Second)
		processing = false
	}()

	// 清屏
	clear()
	log.Infof("Rerun: [%s]", w.cmdStr)
	cmd := w.initCmd()
	err := cmd.Start()
	if err == nil {
		w.oldPID = cmd.Process.Pid
	}
	cmd.Wait()
	w.kill()
}

func (w *Watcher) kill() {
	if w.oldPID > 0 {
		err := syscall.Kill(-w.oldPID, syscall.SIGKILL)
		if err != nil {
			log.Info("kill process err: ", err)
		}
		w.oldPID = 0
	}
}

func printCmdResult(r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Debug("printCmdResult err: ", err)
			return
		}
		fmt.Println(string(line))
	}
}

// 清屏
func clear() {
	var clearCMD *exec.Cmd
	if runtime.GOOS == "windows" {
		clearCMD = exec.Command("cmd", "/c", "cls")
	} else {
		clearCMD = exec.Command("clear")
	}
	clearCMD.Stdout = os.Stdout
	clearCMD.Run()
}
