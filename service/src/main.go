package main

import (
	"common"
	"flag"
	"handler"
	"net/http"
)

const (
	lwx = "/wx"
)

var (
	port string = "9090"
)

func main() {
	flag.StringVar(&port, "p", "9090", "port")
	flag.Parse()
	common.Log.Info("[server] listen at port %s", port)

	http.HandleFunc(lwx, handler.WxHandler)
	http.ListenAndServe(common.KColon+port, nil)
}
