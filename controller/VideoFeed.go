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
}

type VideoFeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"` //视频列表
	NextTime  int64          `json:"next_time"`
}

func Feed(c *gin.Context) {
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
	//获取用户id
	//UserID, err := c.Get("UserID")
	//if !err {
	//	UserID = -1
	//}
	UserID := request.UserID
	fmt.Print(UserID)
	//查询当前用户下所有视频
	VideoList, NextTime, error := service.Feed(15, UserID, request.LatestTime)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	response.NextTime = NextTime
	c.JSON(200, response)
}
