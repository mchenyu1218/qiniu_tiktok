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
	//查询这个用户发布的视频
	apiRouter.GET("/feedbyusername", controller.Feedbyusername)
	//查询视频的同时查询关注列表和是否关注该博主
	apiRouter.GET("/feedalluser", controller.Feedalluser)
	//进行关注
	apiRouter.GET("/collectfeed", controller.CollectAction)

	//apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register", controller.Register)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/modeify", controller.Usermodify)
	//收藏该视频

	apiRouter.POST("/collect/action", controller.CollectAction)

	//apiRouter.POST("/favorite/Collect/", controller.Collect)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	return r
}
