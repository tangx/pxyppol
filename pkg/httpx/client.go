package httpx

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/tangx/pxypool/pkg/pxyctx"
)

func GET(target, pxy string) (*http.Response, error) {

	tr := &http.Transport{
		// Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 设置 proxy
	if len(pxy) != 0 {
		proxy, err := url.Parse(pxy)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	return client.Get(target)

}

func GETx(target string) (*http.Response, error) {
	pxy := pxyctx.RandomPxy()
	return GET(target, pxy)
}
