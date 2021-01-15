package browser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func RandomHandler(c *gin.Context) {

	n := len(pxyctx.Pool)
	if n == 0 {
		c.String(http.StatusOK, "")
		return
	}

	for pxy, ok := range pxyctx.Pool {
		if ok {
			c.String(http.StatusOK, pxy)
			return
		}
	}
}

func ListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, pxyctx.Pool)
}
