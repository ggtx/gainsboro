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
