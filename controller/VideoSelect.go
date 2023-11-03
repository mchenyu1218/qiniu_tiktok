package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoSelectRequest struct {
	VideoID int64 `form:"video_id" json:"video_id" validator:"omitempty"`
}

type VideoSelectResponse struct {
	common.Response
	common.Video
}

func SelectByVideoID(c *gin.Context) {
	var request VideoSelectRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "视频ID错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &VideoSelectResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	v, error := service.SelectVideo(request.VideoID)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "select error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.Video = v
	c.JSON(200, response)
}
