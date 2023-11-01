package service

import (
	common "Projectdouy/commom"
	"log"
)

// Feed 返回指定投稿时间之后的amount个视屏
func CollectList(UserId int64) ([]common.Video, error) {
	videos, error := videoDao.GetCollectList(UserId)
	if error != nil {
		log.Println("videoDao.GetLikeList 出错")
	}
	return videos, error
}
