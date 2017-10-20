package main

import (
	"bytes"
	"common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type MenuItem struct {
	Button []*Menu_Button `json:"button"`
}

type Menu_Button struct {
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Key       string            `json:"key"`
	Url       string            `json:"url"`
	SubButton []*Menu_SubButton `json:"sub_button"`
}

type Menu_SubButton struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Key  string `json:"key"`
	Url  string `json:"url"`
}

func main() {
	makeMenu()
}

func makeMenu() {
	at := common.GetNewAccessToken(time.Now().Unix())
	common.Log.Warn("at:%s", at)
	menuApi := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + at
	mi := &MenuItem{
		Button: []*Menu_Button{
			{
				Name: "bagua",
				Type: "view",
				Url:  "http://todayisnow.ink/content/bg",
			},
		},
	}
	miByte, err := json.Marshal(mi)
	if err != nil {
		common.Log.Warn("Make menu err:%v", err)
		return
	}
	common.Log.Info("req body:%s", miByte)
	resp, err := http.Post(menuApi, "application/json", bytes.NewReader(miByte))
	if err != nil {
		common.Log.Warn("Post err:%v", err)
		return
	}
	common.Log.Debug("resp:%v", resp)
	b, _ := ioutil.ReadAll(resp.Body)
	common.Log.Debug("menu resp:%+v", string(b))
}
