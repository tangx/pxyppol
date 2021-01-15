package keeper

import (
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func Local() {
	for {
		select {
		case pxy := <-pxyctx.PxyExpiredCh:
			// logrus.Infof("proxy expired: %s", pxy)
			delete(pxyctx.Pool, pxy)
		case pxy := <-pxyctx.PxyReadyCh:
			pxyctx.Pool[pxy] = true
		}
	}
}

func Initial() {
	Local()
}
