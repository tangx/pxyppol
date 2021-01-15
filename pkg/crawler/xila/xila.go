package xiladaili

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func craw(page int) {
	url := fmt.Sprintf("http://www.xiladaili.com/gaoni/%d/", page)
	logrus.Infoln(url)

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	// goquery parse table
	// https://gist.github.com/salmoni/27aee5bb0d26536391aabe7f13a72494
	// goquery class div with whitespace
	// https://stackoverflow.com/a/27933907
	doc.Find(".mt-0.mb-2.table-responsive>table>tbody>tr").Each(func(index int, trhtml *goquery.Selection) {
		p := pxyctx.Proxy{}

		trhtml.Find("td").Each(func(idx int, cell *goquery.Selection) {
			switch idx {
			case 0:
				p.Address = cell.Text()
			case 1:
				p.Protocol = pTrim(cell.Text())
			}
		})

		pxy := fmt.Sprintf("%s://%s", p.Protocol, p.Address)
		pxy = strings.ToLower(pxy)
		logrus.Debugln(pxy)

		// 发送到队列等待处理
		pxyctx.PxyCandidateCh <- pxy
	})
}

func pTrim(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Trim(s, "代理")
	return strings.Split(s, ",")[0]
}

func start() {
	for i := 1; i <= 50; i++ {
		craw(i)
		time.Sleep(2 * time.Second)

		// cycle
		if i == 10 {
			i = 1
			time.Sleep(60 * time.Second)
		}
	}
}

func Initial() {
	go start()
}
