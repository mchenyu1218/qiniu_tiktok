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

// @Summary 根据video_id删除相关视频
// @Produce	json
// @Param		username	query		string				false	"用户名"
// @Param		password	query		string				false	"密码"
// @Success	200		{object}	UserRegisterResponse        "成功"
// @Failure	400		{object}	UserRegisterResponse	"请求错误"
// @Failure	500		{object}	UserRegisterResponse	"内部错误"
// @Router		/deleteVideo  [post]
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
