package handler

import (
	"common"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type UserMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
}

type RespMsg struct {
	xmldata *xmlRespMsg `xml:"xml"`
}

type xmlRespMsg struct {
	ToUserName   string `xml:"cdata,ToUserName"`
	FromUserName string `xml:"cdata,FromUserName"`
	CreateTime   int64  `xml:"cdata,CreateTime"`
	MsgType      string `xml:"cdata,MsgType"`
	Content      string `xml:"cdata,Content"`
	MsgId        string `xml:"cdata,MsgId"`
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
		resp := &RespMsg{
			xmldata: &xmlRespMsg{
				FromUserName: "gh_283616b98eee",
				ToUserName:   req.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      req.MsgType,
				Content:      "I am working!",
			},
		}
		bresp, err := xml.Marshal(resp)
		if err != nil {
			common.Log.Debug("marshal resp err:%v", err)
			return
		}
		w.Write(bresp)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
