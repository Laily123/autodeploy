## watch  
允许参数 自定义命令。
监听文件，文件改动就执行指定命令，默认是 `go run *.go`，可以指定。
现在监听文件的改动收到的事件太多，所以执行一个命令后 3 秒才可以再次执行命令。

## server
允许参数 -config=xxx 配置文件地址
配置
[[project]]
url = ""
secret = ""
branch = ""
script = ""


