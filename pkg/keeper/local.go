package keeper

import (
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func Local() {
	for {
		pxy := <-pxyctx.PxyReadyCh
		pxyctx.Pool[pxy] = true
	}
}

func Initial() {
	Local()
}
