package keeper

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func Local() {
	logrus.Infoln("keeper.Local")
	for {
		pxy := <-pxyctx.PxyReadyCh
		pxyctx.Pool[pxy] = true
		// spew.Dump(pxyctx.Pool)
	}
}

func Initial() {
	Local()
}
