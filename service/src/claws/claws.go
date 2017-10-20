package claws

import (
	"bytes"
	"common"
	"github.com/PuerkitoBio/goquery"
)

func GetQueryResponse(cat string) string {
	doc, err := goquery.NewDocument("https://www.douban.com/group/blabla/discussion?start=0")
	if err != nil {
		common.Log.Error("[GetQueryResponse] make doc fail, err:%s", err.Error())
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString("<html>\n<body>")
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		h, err := s.Html()
		if err != nil {
			common.Log.Error("[GetQueryResponse] get elements err:%s, text:%s", err.Error(), s.Text())
			return
		}
		buf.WriteString(h)
		buf.WriteString("</br>")
	})
	buf.WriteString("</body>\n</html>")

	return buf.String()
}
