package main

import (
	"flag"

	"bufio"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port = ""
)

func init() {
	flag.StringVar(&port, "port", "8000", "server port")
	flag.Parse()
}

func handlerGitea(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("get gitea request body err: ", err, r.URL.String())
	}
	log.Info("body: ", string(req))
	writer := bufio.NewWriter(w)
	writer.WriteString("hello")
	writer.Flush()
}

func main() {
	http.HandleFunc("/gitea", handlerGitea)
	log.Info("server start: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
