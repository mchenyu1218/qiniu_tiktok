package service

import (
	"Projectdouy/repository"
)

var (
	likeDao = repository.NewLikeDaoInstance()
)

// FavoriteAction 点赞与取消赞操作
func FavoriteAction(userId int64, videoId int64, actionType int32) error {
	// _, err := likeDao.GetLikeByUserIDAndVideoID(userId, videoId)
	if actionType == 1 { //点赞
		err := likeDao.InsertLike(userId, videoId)
		return err
	} else if actionType == 2 { //取消点赞
		err := likeDao.DeleteLike(userId, videoId)
		return err
	}
	return nil
}

// 查询Like接口
