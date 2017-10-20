package handler

import (
	"common"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type UserMsg struct {
	ToUserName string `xml:"ToUserName"`
}

func WxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_req := r.URL.Query()
		signature, ts, nonce, echostr := _req.Get("signature"), _req.Get("timestamp"), _req.Get("nonce"), _req.Get("echostr")
		common.Log.Debug("[WxHandler]signature:%s, ts:%s, nonce:%s, echostr:%s", signature, ts, nonce, echostr)
		if !common.CheckInitRequest(signature, ts, nonce) {
			common.Log.Warn("[WxHandler] token verify fails")
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			w.Write([]byte(echostr))
			return
		}
	} else if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			common.Log.Warn("request body read err:%v", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.Log.Debug("body:%v", string(body))
		req := &UserMsg{}
		err = xml.Unmarshal(body, req)
		if err != nil {
			common.Log.Warn("unmarshal body err:%v", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.Log.Debug("request:%v", req)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
