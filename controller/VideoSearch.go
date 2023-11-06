package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoSearchRequest struct {
	QueryString string `form:"query" json:"query" validator:"omitempty"`
}

type VideoSearchResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"` //视频列表
}

// @Summary 进行偶所
// @Produce	json
// @Param		title	query		string		false	"标签"
// @Param		description	query	string		false	"描述"
// @Param		tag	query	query   string		false	"标签"
// @Param		play_url	query	string		false	"视频url"
// @Param		cover_url	query	string		false	"封面url"
// @Param		author_id	query	int64		false	"用户id"
// @Success	200		{object}	VideoSearchResponse        "成功"
// @Failure	400		{object}	VideoSearchResponse	"请求错误"
// @Failure	500		{object}	VideoSearchResponse	"内部错误"
// @Router		/insertVideo  [post]
func SearchVideo(c *gin.Context) {
	var request VideoSearchRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "搜索错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &VideoSearchResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	VideoList, error := service.SearchVideo(request.QueryString)
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "update error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	response.VideoList = VideoList
	c.JSON(200, response)
}
