package routers

import (
	"github.com/gin-gonic/gin"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/config"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/serializer"
)

func InitRouter() *gin.Engine {
	if config.Conf().DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	api(r)

	return r
}

func api(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("/api")
	{
		api.Any("/ping", func(c *gin.Context) {
			c.JSON(200, &serializer.Response{
				Code: 200,
				Msg:  "OK",
				Data: "pong",
			})
		})
	}
	return api
}
