package kuaidaili

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/httpx"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func craw(page int) {
	url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d/", page)
	logrus.Debugln(url)

	// Request the HTML page.
	resp, err := httpx.GETx(url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	// goquery parse table
	// https://gist.github.com/salmoni/27aee5bb0d26536391aabe7f13a72494
	doc.Find("#list>table>tbody>tr").Each(func(index int, trhtml *goquery.Selection) {
		p := pxyctx.Proxy{}

		trhtml.Find("td").Each(func(idx int, cell *goquery.Selection) {
			switch idx {
			case 0:
				p.Address = cell.Text()
			case 1:
				p.Port = cell.Text()
			case 3:
				p.Protocol = cell.Text()
			}
		})

		pxy := fmt.Sprintf("%s://%s:%s", p.Protocol, p.Address, p.Port)
		pxy = strings.ToLower(pxy)
		logrus.Debugln(pxy)

		// 发送到队列等待处理
		pxyctx.PxyCandidateCh <- pxy
	})
}

func start() {
	for i := 1; i <= 50; i++ {
		craw(i)
		time.Sleep(2 * time.Second)

		// cycle
		if i == 20 {
			i = 1
			time.Sleep(60 * time.Second)
		}
	}
}

func Initial() {
	go start()
}
