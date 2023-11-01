package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"github.com/gin-gonic/gin"
	"log"
)

type LikeRequest struct {
	UserID     int64  `form:"user_id" json:"user_id" validator:"omitempty"`
	Token      string `form:"token" json:"token" validator:"required "`
	VideoID    int64  `form:"video_id" json:"video_id" validator:"required,gt=0"`
	ActionType int32  `form:"action_type" json:"action_type" validator:"required,gte=1,lte=2"` //1点咱2取消
}

type LikeResponse struct {
	common.Response
}

// @Summary	进行点赞操作，1点赞2取消
// @Produce	json
// @Param		user_id	query		int64				false	"用户id"
// @Param		token	query		string				false	"验证信息"
// @Param		VideoID		query		int64			false	"视频id"
// @Param		ActionType	query		int				false	"行动目标"
// @Success	200		{object}	LikeResponse        "成功"
// @Failure	400		{object}	LikeResponse	"请求错误"
// @Failure	500		{object}	LikeResponse	"内部错误"
// @Router		/favorite/action/  [post]
func FavoriteAction(c *gin.Context) {
	var request LikeRequest
	var response = &LikeResponse{}
	if err := c.Bind(&request); err != nil {
		response.StatusCode = 1
		response.StatusMsg = "参数解析失败"
		c.JSON(400, response)
		log.Println("赞操作request参数绑定失败")
		return
	}

	err := service.FavoriteAction(request.UserID, request.VideoID, request.ActionType)
	if err != nil {
		log.Println("赞操作失败", err)
		response.StatusCode = 1
		response.StatusMsg = "赞操作失败"
		c.JSON(400, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}
