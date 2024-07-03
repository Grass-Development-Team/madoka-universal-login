package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/config"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/serializer"
	"projects.gdt.im/yuanzui-cf/madoka-universal-login/internal/utils"
)

func InitRouter() *gin.Engine {
	if config.Conf().DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.LoggerWithWriter(utils.Log().Writer))

	api(r)

	return r
}

func api(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("/api")
	{
		api.Any("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, &serializer.Response{
				Code: serializer.CodeOK,
				Msg:  "OK",
				Data: "pong",
			})
		})
	}
	return api
}
