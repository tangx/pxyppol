package checker

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func get(u string, pxy string) bool {

	proxy, _ := url.Parse(pxy)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 1,
	}

	resp, err := client.Get(u)
	if err != nil {
		// fmt.Printf("%s not work\n", pxy)
		logrus.Debugf("%s not work\n", pxy)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		logrus.Infof("%s ok\n", pxy)
		return true
	}
	return false
}

func Check(pxy string) bool {
	// u := "https://ip.sb"
	u := "http://ip.cip.cc"
	return get(u, pxy)
}

func Filter(pxy string) {
	if Check(pxy) {
		pxyctx.PxyReadyCh <- pxy
	}
}

func Initial() {
	for {
		candidate := <-pxyctx.PxyCandidateCh
		Filter(candidate)
	}
}