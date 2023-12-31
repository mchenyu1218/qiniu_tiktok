package repository

import (
	"errors"
	"log"
	"sync"

	"gorm.io/gorm"
)

type CollectDao struct{}

var (
	collectOnce sync.Once
	collectDao  *CollectDao
)

func NewCollectDaoInstance() *CollectDao {
	//不论NewCollectDaoInstance()被调用多少次，Do中的内容只会调用一次
	collectOnce.Do(
		func() {
			//在Go语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
			collectDao = &CollectDao{}
		})
	return collectDao
}

func (collectDao *CollectDao) GetCollectByUserIDAndVideoID(UserID int64, VideoId int64) (Collect, error) {
	f := Collect{}
	result := Db.Table("collect").Where("user_id = ? and video_id = ?", UserID, VideoId).Take(&f)
	//错误处理
	if result.Error != nil {
		//当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return f, errors.New("找不到指定的记录")
		}
		return f, errors.New("发生未知错误")
	}
	return f, nil
}

// InsertLike 插入点赞数据
func (collectDao *CollectDao) InserCollect(UserID int64, VideoId int64) error {
	// 一个事务
	tx := Db.Begin()

	defer func() {
		if r := recover(); r != nil {
			log.Println("回滚")
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Println("事务开启异常")
	}

	collect := &Collect{UserID: UserID, VideoID: VideoId}
	if err := tx.Table("collect").Select("user_id", "video_id").Create(&collect).Error; err != nil {
		log.Println("添加收藏回滚！错误信息：", err)
		tx.Rollback()
	}
	video := &Video{ID: VideoId}
	if err := tx.Table("video").Model(&video).UpdateColumn("collect_count", gorm.Expr("collect_count + 1")).Error; err != nil {
		log.Println("更新收藏回滚！错误信息：", err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

// DeleteLike 取消点赞
func (collectDao *CollectDao) DeleteCollect(UserID int64, VideoId int64) error {
	// 一个事务
	tx := Db.Begin()

	defer func() {
		if r := recover(); r != nil {
			log.Println("回滚")
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		log.Println("事务开启异常")
	}

	collect := &Collect{VideoID: VideoId, UserID: UserID}
	if err := tx.Table("collect").Where("user_id = ? and video_id = ?", UserID, VideoId).Delete(&collect).Error; err != nil {
		log.Println("取消收藏回滚！错误信息：", err)
		tx.Rollback()
	}

	video := &Video{ID: VideoId}
	if err := tx.Table("video").Model(&video).Where("collect_count != 0").UpdateColumn("collect_count", gorm.Expr("collect_count - 1")).Error; err != nil {
		log.Println("更新视频收藏回滚！错误信息：", err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
