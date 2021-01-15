package browser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/pxypool/pkg/pxyctx"
)

func RandomHandler(c *gin.Context) {
	pxy := pxyctx.RandomPxy()
	c.String(http.StatusOK, pxy)
}

func ListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, pxyctx.Pool)
}
