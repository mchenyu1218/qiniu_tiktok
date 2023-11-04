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

// @Summary 修改用户密码
// @Produce	json
// @Param		username	query		string				false	"用户名"
// @Param		oldpassword	query		string				false	"旧的密码"
// @Param		newpassword	query		string				false	"新的密码"
// @Success	200		{object}	UserModifyResponse        "成功"
// @Failure	400		{object}	UserModifyResponse	"请求错误"
// @Failure	500		{object}	UserModifyResponse	"内部错误"
// @Router		/user/modeify  [post]
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
