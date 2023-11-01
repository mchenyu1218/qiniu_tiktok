package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoDeleteRequest struct {
	VideoID string `form:"video_id" json:"video_id" validator:"omitempty"`
}

type VideoDeleteResponse struct {
	common.Response
}

func DeleteByVideoID(c *gin.Context) {
	var request VideoDeleteRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "视频ID错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &VideoDeleteResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	error := service.DeleteVideo(request.VideoID)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "delete error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	c.JSON(200, response)
}
