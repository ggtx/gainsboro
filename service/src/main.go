package main

import (
	"flag"
	"handler"
	"net/http"
)

const (
	lwx    = "/wx"
	kcolon = ":"
)

var (
	port string = "9090"
)

func main() {
	flag.StringVar(&port, "p", "9090", "port")
	flag.Parse()

	http.HandleFunc(lwx, handler.WxHandler)
	http.ListenAndServe(kcolon+port, nil)
}
