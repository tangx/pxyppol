package browser

import (
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func RandomPxy() string {

	for pxy, ok := range pxyctx.Pool {
		if ok {
			return pxy
		}
	}

	return ""
}
