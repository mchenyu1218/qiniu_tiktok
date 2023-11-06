package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

// 获取视频流类
type VideoInsertRequest struct {
	Title       string `form:"title" json:"title" validator:"omitempty"`
	Description string `form:"description" json:"description" validator:"omitempty"`
	Tag         string `form:"tag" json:"tag" validator:"omitempty"`
	PlayURL     string `form:"play_url" json:"play_url" validator:"omitempty"`
	CoverURL    string `form:"cover_url" json:"cover_url" validator:"omitempty"`
	AuthorID    int64  `form:"author_id" json:"author_id" validator:"omitempty"`
}

type VideoInsertResponse struct {
	common.Response
}

// @Summary 增加视频
// @Produce	json
// @Param		title	query		string		false	"标签"
// @Param		description	query	string		false	"描述"
// @Param		tag	query	query   string		false	"标签"
// @Param		play_url	query	string		false	"视频url"
// @Param		cover_url	query	string		false	"封面url"
// @Param		author_id	query	int64		false	"用户id"
// @Success	200		{object}	VideoInsertResponse        "成功"
// @Failure	400		{object}	VideoInsertResponse	"请求错误"
// @Failure	500		{object}	VideoInsertResponse	"内部错误"
// @Router		/insertVideo  [post]
func InsertVideo(c *gin.Context) {
	var request VideoInsertRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "上传视频请求参数错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &VideoInsertResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"
	videoID, error := service.InsertVideo(request.Title, request.Description, request.Tag, request.PlayURL, request.CoverURL, request.AuthorID)
	if error != nil || videoID == -1 {
		response.StatusCode = 1
		response.StatusMsg = "insert error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	c.JSON(200, response)
}
