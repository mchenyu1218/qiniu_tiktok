package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoUpdateRequest struct {
	VideoID     string `form:"video_id" json:"video_id" validator:"omitempty"`
	Title       string `form:"title" json:"title" validator:"omitempty"`
	Description string `form:"description" json:"description" validator:"omitempty"`
	Tag         string `form:"tag" json:"tag" validator:"omitempty"`
}

type VideoUpdateResponse struct {
	common.Response
}

func UpdateByVideoID(c *gin.Context) {
	var request VideoUpdateRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "视频ID错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &VideoUpdateResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	error := service.UpdateVideo(request.VideoID, request.Title, request.Description, request.Tag)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "update error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	c.JSON(200, response)
}
