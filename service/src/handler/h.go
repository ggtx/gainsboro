package handler

import (
	"claws"
	"common"
	"html/template"
	"net/http"
)

type info struct {
	Content string
}

func WxHandler(w http.ResponseWriter, r *http.Request) {
	_req := r.URL.Query()
	signature, ts, nonce, echostr := _req.Get("signature"), _req.Get("timestamp"), _req.Get("nonce"), _req.Get("echostr")
	common.Log.Info("[WxHandler]signature:%s, ts:%s, nonce:%s, echostr:%s", signature, ts, nonce, echostr)
	if !common.CheckInitRequest(signature, ts, nonce) {
		return
	}

	inf := &info{Content: claws.GetQueryResponse("douban")}

	t, err := template.ParseFiles("view/outer.html", "view/content.html")
	if err != nil {
		common.Log.Error("[WxHandler] template parse fail, err:%v", err)
		return
	}
	err = t.ExecuteTemplate(w, "outer", inf)
	if err != nil {
		common.Log.Error("[WxHandler] execute template content err:%v", err)
		return
	}
	w.Write([]byte(inf.Content))

	/*
		err = t.Execute(w, inf)
		if err != nil {
			common.Log.Error("[WxHandler] execute template err:%v", err)
			return
		}
	*/
}
