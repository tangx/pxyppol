package checker

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/browser"
	"github.com/tangx/pxypool/pkg/httpx"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func get(u string, pxy string) bool {
	resp, err := httpx.GET(u, pxy)
	if err != nil {
		// fmt.Printf("%s not work\n", pxy)
		logrus.Debugf("%s not work\n", pxy)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func Check(pxy string) bool {
	// u := "https://ip.sb"
	u := "http://ip.cip.cc"
	return get(u, pxy)
}

func Hire() {
	for {
		candidate := <-pxyctx.PxyCandidateCh
		go func(pxy string) {
			if Check(pxy) {
				logrus.Infof("hire pxy : %s \n", pxy)
				pxyctx.PxyReadyCh <- pxy
			}
		}(candidate)
	}
}
func Fire() {
	for {
		pxy := browser.RandomPxy()
		if !Check(pxy) {
			logrus.Infof("fire pxy : %s \n", pxy)
			pxyctx.PxyExpiredCh <- pxy
		}
	}
}
func Initial() {
	go Hire()
	go Fire()
}
