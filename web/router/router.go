package router

import (
	"github.com/fanke15/hftoc/pkg/lib/bolt"
	"github.com/fanke15/hftoc/web/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(e *gin.Engine) {
	// 设置默认页面
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/page/login")
	})

	//-----------------------WEB-----------------------
	webPage := e.Group("page")
	{
		webPage.GET("login", middleware.InitAce("tmpl", "login", map[string]interface{}{
			"title":  "用户登录",
			"config": string(bolt.InitBolt().Query("login")),
		}))
		webPage.GET("dashboard", middleware.InitAce("tmpl", "default", map[string]interface{}{
			"title":  "仪表盘",
			"config": string(bolt.InitBolt().Query("dashboard")),
		}))
	}

	//-----------------------API-----------------------

}
