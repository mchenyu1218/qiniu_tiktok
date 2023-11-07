package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"

	"github.com/gin-gonic/gin"
	"log"
)

//调用service层的服务

type CacheGetRequest struct {
	Score string `form:"score" json:"score" validator:"omitempty"`
}

type CacheGetResponse struct {
	common.Response
}

func GetCache(c *gin.Context) {
	var request VideoSelectRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, &common.Response{StatusCode: 1, StatusMsg: "视频ID错误"})
		log.Println("request参数绑定失败")
		return
	}

	var response = &CacheGetResponse{}
	response.StatusCode = 0
	response.StatusMsg = "success"

	val, error := service.TestCache("string")
	if error != nil {
		response.StatusCode = 1
		response.StatusMsg = "update error"
		log.Println(error)
		c.JSON(200, response)
		return
	}
	log.Println(val)
	c.JSON(200, response)
}
