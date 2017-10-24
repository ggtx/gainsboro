package handler

import (
	"bytes"
	"encoding/xml"
	"handler"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	req := &handler.UserMsg{
		ToUserName:   "testdevuser",
		FromUserName: "openiduser",
		MsgType:      "text",
		Content:      "new content",
	}
	breq, err := xml.Marshal(req)
	if err != nil {
		t.Log(err)
		return
	}

	resp, err := http.Post("http://127.0.0.1:9090/wx", "text/xml", bytes.NewReader(breq))
	if err != nil {
		t.Log(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(body))
}
