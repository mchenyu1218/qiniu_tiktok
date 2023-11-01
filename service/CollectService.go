package service

import (
	"Projectdouy/repository"
)

var (
	collectDao = repository.NewCollectDaoInstance()
)

// FavoriteAction 点赞与取消赞操作
func CollectAction(userId int64, videoId int64, actionType int32) error {
	// _, err := likeDao.GetLikeByUserIDAndVideoID(userId, videoId)
	if actionType == 1 { //收藏
		err := collectDao.InserCollect(userId, videoId)
		return err
	} else if actionType == 2 { //取消收藏
		err := collectDao.DeleteCollect(userId, videoId)
		return err
	}
	return nil
}

// 查询Like接口
