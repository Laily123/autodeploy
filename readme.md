## audodeploy

### watch
监听文件改变，自动重新运行程序
在项目main包所在目录运行 `autodeploy watch`，则会监听在目录下的文件变化，当发现变化是自动执行 `go run *.go`，结果输出在控制台。  
如果需要执行自定义的命令，则运行 `autodeploy watch <your cmd>`。

### server
配合 webhook 执行仓库里脚本文件




