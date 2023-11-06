package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoFeedRequest struct {
	UserID     int64  `form:"user_id" json:"user_id" validator:"omitempty"`
	LatestTime int64  `form:"latest_time" json:"latest_time" validator:"omitempty"`
	Token      string `form:"token" json:"token" validator:"omitempty"`
	Tag        string `form':"tag" json:"tag"    validator:"omitempty"`
}

type VideoFeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"` //视频列表
	NextTime  int64          `json:"next_time"`
}

// @Summary 获取所有视频
// @Produce	json
// @Success	200		{object}	VideoFeedResponse        "成功"
// @Failure	400		{object}	VideoFeedResponse	"请求错误"
// @Failure	500		{object}	VideoFeedResponse	"内部错误"
// @Router		/feedallvideo  [get]
func Feed(c *gin.Context) {

	var response = &VideoFeedResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	//查询当前用户下所有视频
	VideoList, error := service.Feed()
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	c.JSON(200, response)
}

// @Summary 根绝tag获取视频
// @Produce	json
// @Param		tag	query		string		false	"类别"
// @Success	200		{object}	VideoFeedResponse        "成功"
// @Failure	400		{object}	VideoFeedResponse	"请求错误"
// @Failure	500		{object}	VideoFeedResponse	"内部错误"
// @Router		/feedbytag  [get]
func FeedByTag(c *gin.Context) {

	var response = &VideoFeedResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"
	//获取用户id
	//UserID, err := c.Get("UserID")
	//if !err {
	//	UserID = -1
	//}
	Tag := c.Query("tag")

	//查询当前用户下所有视频
	VideoList, error := service.Feedbytag(15, Tag)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	c.JSON(200, response)
}

// @Summary 根绝username获取视频
// @Produce	json
// @Param		username	query		string		false	"用户名"
// @Success	200		{object}	VideoFeedResponse        "成功"
// @Failure	400		{object}	VideoFeedResponse	"请求错误"
// @Failure	500		{object}	VideoFeedResponse	"内部错误"
// @Router		/feedbyusername  [get]
func Feedbyusername(c *gin.Context) {

	var response = &VideoFeedResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"
	//获取用户id
	//UserID, err := c.Get("UserID")
	//if !err {
	//	UserID = -1
	//}
	Username := c.Query("username")

	//查询当前用户下所有视频
	VideoList, error := service.Feedbyusername(Username)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	c.JSON(200, response)
}

// @Summary 根绝username获取视频
// @Produce	json
// @Param		user_id	query		int64		false	"用户名"
// @Param		latest_time	query	int64		false	"最后时间"
// @Param		token	query	string		false	"验证信息"
// @Param		tag	query	string		false	"验证信息"
// @Success	200		{object}	VideoFeedResponse        "成功"
// @Failure	400		{object}	VideoFeedResponse	"请求错误"
// @Failure	500		{object}	VideoFeedResponse	"内部错误"
// @Router		/feedalluser  [get]
func Feedalluser(c *gin.Context) {
	var request VideoFeedRequest
	//参数绑定
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("request参数绑定失败")
		return
	}
	var response = &VideoFeedResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	fmt.Print(request.UserID)
	//查询当前用户下所有视频
	VideoList, error := service.Feedalluser(request.UserID)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	c.JSON(200, response)
}
