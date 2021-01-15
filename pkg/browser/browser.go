package browser

import (
	"github.com/gin-gonic/gin"
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

func Execute() error {
	r := gin.Default()
	r.GET("/ping")
	r.GET("list", ListHandler)
	r.GET("/random", RandomHandler)
	return r.Run()
}
