package handler

import (
	"common"
	x "encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
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

type Resp struct {
	RespXml *xml
}

type xml struct {
	ToUserName   *ToUserNameS
	FromUserName *FromUserNameS
	CreateTime   int64
	MsgType      *MsgTypeS
	Content      *ContentS
}

type ToUserNameS struct {
	UserName string `xml:",cdata"`
}

type FromUserNameS struct {
	UserName string `xml:",cdata"`
}

type MsgTypeS struct {
	MType string `xml:",cdata"`
}

type ContentS struct {
	Data string `xml:",cdata"`
}

type MsgIdS struct {
	Id string `xml:",cdata"`
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
		common.Log.Debug("request:%+v", r)
		req := &UserMsg{}
		err = x.Unmarshal(body, req)
		if err != nil {
			common.Log.Warn("unmarshal body err:%v", err)
			w.WriteHeader(http.StatusNoContent)
			return
		}
		common.Log.Debug("request:%v", req)
		resp := &Resp{
			RespXml: &xml{
				FromUserName: &FromUserNameS{UserName: "gh_283616b98eee"},
				ToUserName:   &ToUserNameS{UserName: req.FromUserName},
				CreateTime:   time.Now().Unix(),
				MsgType:      &MsgTypeS{MType: req.MsgType},
				Content:      &ContentS{Data: "I am working!"},
			},
		}
		bresp, err := x.Marshal(resp.RespXml)
		if err != nil {
			common.Log.Debug("marshal resp err:%v", err)
			return
		}
		common.Log.Debug("resp:%s", bresp)
		w.Header().Set("Content-Type", "text/xml")
		w.Header().Set("Content-Length", strconv.Itoa(len(bresp)))
		w.Write(bresp)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
