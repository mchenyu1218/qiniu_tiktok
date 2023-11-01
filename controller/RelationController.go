package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"github.com/gin-gonic/gin"
	"log"
)

type RelationActionRequest struct {
	UserID   int64  `form:"user_id" json:"user_id" validator:"omitempty"`
	Token    string `form:"token" json:"token" validator:"required"`
	ToUserID int64  `form:"to_user_id" json:"to_user_id" validator:"required,gt=0"`
	// 1关注 2取消
	ActionType int `form:"action_type" json:"action_type" validator:"required,gt=0,lt=3"`
}

// RelationAction 关注
func RelationAction(c *gin.Context) {
	var request RelationActionRequest
	var response = &common.Response{}
	//参数绑定
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("request参数绑定失败")
		return
	}

	err := service.FollowRelationAction(request.ToUserID, request.UserID, request.ActionType)
	if err != nil {
		log.Println("关注/取关操作失败", err)
		response.StatusCode = 1
		response.StatusMsg = "关注/取关操作失败"
		c.JSON(400, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}
