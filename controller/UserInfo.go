package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfoRequest struct {
	UserId int64  `form:"user_id" json:"user_id" validator:"required,gt=0"`
	Token  string `form:"token" json:"token" validator:"required"`
}

type UserInfoResponse struct {
	common.Response
	User common.UserIn
}

// @Summary	获取用户信息，同时查询是否关注该用户
// @Produce	json
// @Param		user_id	query		int64				false	"用户id"
// @Param		token	query		string				false	"验证信息"
// @Success	200		{object}	common.Response        "成功"
// @Failure	400		{object}	common.Response	"请求错误"
// @Failure	500		{object}	common.Response	"内部错误"
// @Router		/user/search/:UserID  [get]
func UserInfo(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	var response = &UserInfoResponse{}
	user, err := service.GetUserByID(int64(userId))
	if err != nil {
		log.Println(err)
		response := common.Response{StatusCode: 1, StatusMsg: "用户不存在"}
		c.JSON(404, response)
		return
	}
	userVO := common.UserIn{}
	userVO.Id = user.ID
	userVO.Name = user.Username
	userVO.FollowCount = int64(user.FollowCount)
	userVO.FollowerCount = int64(user.FollowerCount)
	userVO.Face = user.Face

	// 设置is_follow属性：
	//   1.获取 发起当前请求的用户的信息
	// 		a.如果是未登录的用户，is_follow的值应该为false；
	// 		b.如果是已经登录的用户，is_follow的值根据fans表中的数据决定
	curUserIdStr := c.Param("UserID")
	curUserId, err := strconv.ParseInt(curUserIdStr, 10, 64)

	if curUserId == int64(userId) {
		userVO.IsFollow = false
	} else {
		userVO.IsFollow = service.HasFollowed(int64(userId), curUserId)
	}

	response.StatusCode = 0
	response.StatusMsg = "success"
	response.User = userVO
	c.JSON(200, response)
}
