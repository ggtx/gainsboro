package handler

import (
	"net/http"
	"time"
)

func WxHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05") + "\n"

	w.Write([]byte(nowStr))
}
