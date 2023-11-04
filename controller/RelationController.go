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

// @Summary	进行点赞操作，1点赞2取消
// @Produce	json
// @Param		user_id	query		int64				false	"用户id"
// @Param		token	query		string				false	"验证信息"
// @Param		to_user_id		query		int64			false	"博主id"
// @Param		ActionType	query		int				false	"行动目标"
// @Success	200		{object}	common.Response        "成功"
// @Failure	400		{object}	common.Response	"请求错误"
// @Failure	500		{object}	common.Response	"内部错误"
// @Router		/relation/action [post]
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
