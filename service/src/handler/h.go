package handler

import (
	"net/http"
	"time"
)

func WxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now()))
}
