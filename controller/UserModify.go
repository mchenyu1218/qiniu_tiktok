package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"Projectdouy/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type UserModifyRequest struct {
	UserName    string `form:"username" json:"username" validator:"required,min=6,max = 20"`
	OldPassword string `form:"oldpassword" json:"olepassword" validator:"required,min=6,max = 20"`
	NewPassword string `form:"newpassword" json:"olepassword" validator:"required,min=6,max = 20"`
}

type UserModifyResponse struct {
	common.Response
	UserID int64  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

func Usermodify(c *gin.Context) {
	var request UserModifyRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("request参数绑定失败")
		return
	}

	var response = &UserModifyResponse{}
	UserID, err := service.Usermodify(request.UserName, request.OldPassword, request.NewPassword)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenToken(UserID)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response.UserID = UserID
	response.Token = token
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}
