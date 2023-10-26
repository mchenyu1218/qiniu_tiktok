package routes

import (
	"Projectdouy/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	/**
	  设置路由组tiktok
	*/
	apiRouter := r.Group("/tiktok")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)

	return r
}
