package claws

import (
	"bytes"
	"common"
	"github.com/PuerkitoBio/goquery"
	//"os"
)

func GetQueryResponse(cat string) string {
	/*
		f, err := os.Open("claws/first.html")
		if err != nil {
			common.Log.Error("[GetQueryResponse] open file err:%v", err)
			return ""
		}

		defer f.Close()
		doc, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			common.Log.Error("[GetQueryResponse] make doc fail, err:%s", err.Error())
			return ""
		}
	*/

	doc, err := goquery.NewDocument("https://www.douban.com/group/blabla/discussion?start=0")
	if err != nil {
		common.Log.Error("[GetQueryResponse] make doc fail, err:%s", err.Error())
		return ""
	}

	var buf bytes.Buffer
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		h, err := s.Html()
		if err != nil {
			common.Log.Error("[GetQueryResponse] get elements err:%s, text:%s", err.Error(), s.Text())
			return
		}
		buf.WriteString(h)
		buf.WriteString("</br>")
	})

	return buf.String()
}
