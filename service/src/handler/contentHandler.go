package handler

import (
	"claws"
	"net/http"
)

type info struct {
	Content string
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	inf := &info{Content: claws.GetQueryResponse("douban")}
	w.Write([]byte(inf.Content))
}
