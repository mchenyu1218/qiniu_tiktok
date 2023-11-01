package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"github.com/gin-gonic/gin"
	"log"
)

type CollectListRequest struct {
	Token  string `form:"token" json:"token" validator:"required"`
	UserID int64  `form:"user_id" json:"user_id" validator:"required,gt=0"`
}

type CollectListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// @Summary	获取收藏的视频列表
// @Produce	json
// @Param		user_id	query		int64				false	"用户id"
// @Param		Token	query		string				false	"验证信息"
// @Success	200		{object}	CollectListResponse	 "成功"
// @Failure	400		{object}	CollectListResponse	 "请求错误"
// @Failure	500		{object}	CollectListResponse	 "内部错误"
// @Router		/collect/list [get]
func CollectList(c *gin.Context) {
	var request CollectListRequest
	var response = &CollectListResponse{}
	if err := c.Bind(&request); err != nil {
		response.StatusCode = 1
		response.StatusMsg = "参数解析失败"
		c.JSON(400, response)
		log.Println("request参数绑定失败")
		return
	}
	VideoList, error := service.CollectList(request.UserID)
	if error != nil {
		log.Println("赞操作失败", error)
		response.StatusCode = 1
		response.StatusMsg = "赞操作失败"
		c.JSON(400, response)
		return
	}
	response.VideoList = VideoList
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}
