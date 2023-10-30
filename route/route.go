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
	//apiRouter.GET("/feedbyuserid", controller.Feedbyuserid)

	apiRouter.GET("/feedallvideo", controller.Feed)
	apiRouter.GET("/feedbytag", controller.FeedByTag)
	apiRouter.GET("/feedbyusername", controller.Feedbyusername)

	//apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register", controller.Register)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/modeify", controller.Usermodify)

	return r
}
