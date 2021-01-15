package browser

import (
	"github.com/gin-gonic/gin"
)

func Execute() error {
	r := gin.Default()
	r.GET("/ping")
	r.GET("list", ListHandler)
	r.GET("/random", RandomHandler)
	return r.Run()
}
