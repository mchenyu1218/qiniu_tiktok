package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"github.com/gin-gonic/gin"
	"log"
)

type CollectRequest struct {
	UserID     int64  `form:"user_id" json:"user_id" validator:"omitempty"`
	Token      string `form:"token" json:"token" validator:"required "`
	VideoID    int64  `form:"video_id" json:"video_id" validator:"required,gt=0"`
	ActionType int32  `form:"action_type" json:"action_type" validator:"required,gte=1,lte=2"` //1点咱2取消
}

type CollectResponse struct {
	common.Response
}

// CollectAction
func CollectAction(c *gin.Context) {
	var request CollectRequest
	var response = &CollectResponse{}
	if err := c.Bind(&request); err != nil {
		response.StatusCode = 1
		response.StatusMsg = "参数解析失败"
		c.JSON(400, response)
		log.Println("收藏操作request参数绑定失败")
		return
	}

	err := service.CollectAction(request.UserID, request.VideoID, request.ActionType)
	if err != nil {
		log.Println("收藏操作失败", err)
		response.StatusCode = 1
		response.StatusMsg = "收藏操作失败"
		c.JSON(400, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}
