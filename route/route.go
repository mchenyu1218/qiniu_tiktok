package routes

import (
	"Projectdouy/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	/**
	  设置路由组tiktok
	*/
	apiRouter := r.Group("/tiktok")

	// basic apis
	//apiRouter.GET("/feedbyuserid", controller.Feedbyuserid)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter.GET("/feedallvideo", controller.Feed)
	apiRouter.GET("/feedbytag", controller.FeedByTag)
	//查询这个用户发布的视频
	apiRouter.GET("/feedbyusername", controller.Feedbyusername)
	//查询视频的同时查询关注列表和是否关注该博主
	apiRouter.GET("/feedalluser", controller.Feedalluser)
	//进行关注

	apiRouter.POST("/deleteVideo", controller.DeleteByVideoID)
	apiRouter.POST("/updateVideo", controller.UpdateByVideoID)
	apiRouter.POST("/insertVideo", controller.InsertVideo)
	apiRouter.GET("/searchVideo", controller.SearchVideo)
	apiRouter.GET("/selectVideo", controller.SelectByVideoID)

	apiRouter.GET("/testCache", controller.GetCache)

	apiRouter.POST("/user/register", controller.Register)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.POST("/user/modeify", controller.Usermodify)

	apiRouter.POST("/publish/action/:UserID", controller.Publish)

	apiRouter.GET("/user/search/:UserID", controller.UserInfo)
	//收藏该视频

	apiRouter.POST("/collect/action", controller.CollectAction)
	apiRouter.GET("/collect/list", controller.CollectList)
	//点赞该视频
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)

	//关注该用户
	apiRouter.POST("/relation/action", controller.RelationAction)
	//apiRouter.POST("/favorite/Collect/", controller.Collect)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	return r
}
