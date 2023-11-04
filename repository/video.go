package repository

import (
	common "Projectdouy/commom"
	"fmt"
	"github.com/huichen/wukong/types"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	videoOnce sync.Once
	videoDao  *VideoDao //实现多态调用
)

// 相关数据访问对象,封装增删改查
type result = struct {
	Id            int64  `json:"id"`
	AuthorId      int64  `json:"author_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	IsCollect     bool   `json:"is_collect"`
	Title         string `json:"title"`
	CollectCount  int64  `json:"collect_count"`
}

var video_result []struct {
	Id            int64  `json:"id"`
	AuthorId      int64  `json:"author_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	IsCollect     bool   `json:"is_collect"`

	Title        string `json:"title"`
	CollectCount int64  `json:"collect_count"`
}

// VideoDao 即数据访问对象，直接对指定的“某个数据源”的增删改查的封装（这里是对video的增删改查）
type VideoDao struct {
}

// NewVideoDaoInstance 返回一个UserVideoDao实例
func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

// 查，获取视频列表
func (videoDao *VideoDao) GetPublishList(UserID int64) ([]common.Video, error) {
	VideoListSQL := " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count," +
		" IFNULL( (SELECT 1 FROM	favorite WHERE favorite.user_id = " + fmt.Sprintf("%v", UserID) + " and favorite.video_id = video.id LIMIT 1) , false ) as is_favorite," +
		" IFNULL( (SELECT 1 FROM	fans WHERE fans.fans_id = " + fmt.Sprintf("%v", UserID) + " and fans.blogger_id = 1 LIMIT 1) , false ) as is_follow," +
		" IFNULL( (SELECT 1 FROM	collect WHERE collect.user_id = " + fmt.Sprintf("%v", UserID) + " and collect.video_id = video.id LIMIT 1) , false ) as is_collect" +
		" from user join video" +
		" on video.author_id = user.id" +
		" order by video.create_time"
	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = video_result[i].IsFollow
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = video_result[i].IsFavorite
		VideoList[i].IsCollect = video_result[i].IsCollect

		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount

	}
	return VideoList, nil
}

func (videoDao VideoDao) Feedalluser(UserID any) ([]common.Video, error) {

	var VideoListSQL string
	VideoListSQL = " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count," +
		" IFNULL( (SELECT 1 FROM	favorite WHERE favorite.user_id = " + fmt.Sprintf("%v", UserID) + " and favorite.video_id = video.id LIMIT 1) , false ) as is_favorite," +
		" IFNULL( (SELECT 1 FROM	fans WHERE fans.fans_id = " + fmt.Sprintf("%v", UserID) + " and fans.blogger_id = 1 LIMIT 1) , false ) as is_follow," +
		" IFNULL( (SELECT 1 FROM	collect WHERE collect.user_id = " + fmt.Sprintf("%v", UserID) + " and collect.video_id = video.id LIMIT 1) , false ) as is_collect" +
		" from user join video" +
		" on video.author_id = user.id" +
		" order by video.create_time "

	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = video_result[i].IsFollow
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = video_result[i].IsFavorite
		VideoList[i].IsCollect = video_result[i].IsCollect
		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount
	}
	return VideoList, nil
}

// 获取收藏的视频
func (videoDao VideoDao) GetCollectList(UserID int64) ([]common.Video, error) {

	var res []result
	VideoListSQL := "select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count ,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count,1 as is_collect" +
		" from user join video" +
		" on video.author_id = user.id" +
		" where IFNULL( (SELECT 1 FROM collect WHERE collect.user_id =  " + fmt.Sprintf("%v", UserID) +
		" and collect.video_id = video.id LIMIT 1) , false )  = 1"
	Db.Raw(VideoListSQL).Scan(&res)
	var VideoList = make([]common.Video, len(res))
	for i := 0; i < len(res); i++ {
		VideoList[i].Author.Id = res[i].AuthorId
		VideoList[i].Author.FollowCount = res[i].FollowCount
		VideoList[i].Author.FollowerCount = res[i].FollowerCount
		VideoList[i].Author.Name = res[i].Name
		VideoList[i].Author.IsFollow = false
		VideoList[i].Id = res[i].Id
		VideoList[i].PlayUrl = res[i].PlayUrl
		VideoList[i].CoverUrl = res[i].CoverUrl
		VideoList[i].IsFavorite = false
		VideoList[i].IsCollect = true
		VideoList[i].Title = res[i].Title
		VideoList[i].CommentCount = res[i].CommentCount
		VideoList[i].FavoriteCount = res[i].FavoriteCount
		VideoList[i].CollectCount = res[i].CollectCount
	}
	return VideoList, nil
}

func (videoDao *VideoDao) Addvideo(authorId int64, playUrl string, coverUrl string, title string) error {
	newVideo := &Video{
		AuthorID: authorId,
		CoverUrl: "http://" + coverUrl,
		PlayUrl:  "http://" + playUrl,
		Title:    title,
	}

	err := Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Table("video").Create(newVideo).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		var maxID int64
		result := tx.Table("video").Select("MAX(id)").Scan(&maxID)
		if result.Error != nil {
			panic(result.Error)

		}
		newVideo2 := &UserVideo{
			UserID:  authorId,
			VideoID: maxID,
		}

		result2 := tx.Table("user_video").Create(newVideo2)

		if result2.Error != nil {
			panic(result2.Error)
		}
		return nil

	})

	return err

}

// 查，获取视频列表
func (videoDao *VideoDao) GetVideosbytag(Tag string) ([]common.Video, error) {
	VideoListSQL := " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count" +
		" from user join video" +
		" on video.author_id = user.id" +
		" where video.tag = '" + Tag + "' " +
		" order by video.create_time"
	fmt.Sprint(VideoListSQL)
	fmt.Sprint(Tag)

	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = false
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = false
		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount
	}
	return VideoList, nil
}

// 查，获取视频列表
func (videoDao *VideoDao) GetallVideos() ([]common.Video, error) {
	VideoListSQL := " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count" +
		" from user join video" +
		" on video.author_id = user.id" +
		" order by video.create_time"

	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = false
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = false
		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount
	}
	return VideoList, nil
}

// 查，获取视频列表
func (videoDao *VideoDao) GetVideosbyuserid(Tag string) ([]common.Video, error) {
	VideoListSQL := " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count" +
		" from user join video" +
		" on video.author_id = user.id" +
		" where video.tag = '" + Tag + "' " +
		" order by video.create_time"
	fmt.Sprint(VideoListSQL)
	fmt.Sprint(Tag)

	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = false
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = false
		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount
	}
	return VideoList, nil
}

func (videoDao *VideoDao) GetVideosbyusername(Username string) ([]common.Video, error) {
	VideoListSQL := " select video.id,video.play_url,video.cover_url,video.title,video.comment_count,video.favorite_count,video.collect_count," +
		" video.author_id,user.username as name,user.follow_count,user.follower_count" +
		" from user join video" +
		" on video.author_id = user.id" +
		" where user.username = '" + Username + "' " +
		" order by video.create_time"
	fmt.Sprint(Username)

	Db.Raw(VideoListSQL).Scan(&video_result)
	var VideoList = make([]common.Video, len(video_result))
	for i := 0; i < len(video_result); i++ {
		VideoList[i].Author.Id = video_result[i].AuthorId
		VideoList[i].Author.FollowCount = video_result[i].FollowCount
		VideoList[i].Author.FollowerCount = video_result[i].FollowerCount
		VideoList[i].Author.Name = video_result[i].Name
		VideoList[i].Author.IsFollow = false
		VideoList[i].Id = video_result[i].Id
		VideoList[i].PlayUrl = video_result[i].PlayUrl
		VideoList[i].CoverUrl = video_result[i].CoverUrl
		VideoList[i].IsFavorite = false
		VideoList[i].Title = video_result[i].Title
		VideoList[i].CommentCount = video_result[i].CommentCount
		VideoList[i].FavoriteCount = video_result[i].FavoriteCount
		VideoList[i].CollectCount = video_result[i].CollectCount
		VideoList[i].IsCollect = false
	}
	return VideoList, nil
}

// DeleteVideo 根据 videoID 删除视频记录
func (videoDao *VideoDao) DeleteVideo(videoID string) error {
	// 将is_deleted字段的值修改为1
	deleteSQL := "UPDATE video SET is_deleted = 1 WHERE id = ?"
	result := Db.Exec(deleteSQL, videoID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (videoDao *VideoDao) UpdateVideoDetails(videoID string, title, description, tag string) error {
	// 准备 SQL 语句以更新视频详情，同时将 update_time 设为当前时间
	updateSQL := "UPDATE video SET title = ?, description = ?, tag = ?, update_time = NOW() WHERE id = ?"
	// 执行 SQL 语句来更新视频详情，传递 title, description, tag 和 videoID 作为参数
	result := Db.Exec(updateSQL, title, description, tag, videoID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (videoDao *VideoDao) SelectVideoById(videoID int64) (common.Video, error) {
	var video common.Video
	var temp result
	selectSQL := "SELECT * FROM video WHERE id = ?"
	// 使用Scan方法将查询结果填充到video结构体中
	err := Db.Raw(selectSQL, videoID).Scan(&temp).Error
	if err != nil {
		return common.Video{}, err
	}
	video.Author.Id = temp.AuthorId
	video.Author.FollowCount = temp.FollowCount
	video.Author.FollowerCount = temp.FollowerCount
	video.Author.Name = temp.Name
	video.Author.IsFollow = temp.IsFollow
	video.Id = temp.Id
	video.PlayUrl = temp.PlayUrl
	video.CoverUrl = temp.CoverUrl
	video.IsFavorite = temp.IsFavorite
	video.Title = temp.Title
	video.CommentCount = temp.CommentCount
	video.FavoriteCount = temp.FavoriteCount
	video.CollectCount = temp.CollectCount
	video.IsCollect = temp.IsCollect
	return video, nil
}

func (videoDao *VideoDao) AddVideo(Title, Description, Tag, PlayURL, CoverURL string, AuthorID int64) (int64, error) {
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

	video := &Video{Title: Title, Description: Description, Tag: Tag, PlayUrl: PlayURL, CoverUrl: CoverURL, AuthorID: AuthorID}
	if err := tx.Table("video").Create(video).Error; err != nil {
		log.Println("插入 video 数据时出错：", err)
		tx.Rollback()
		return -1, err
	}
	videoID := video.ID

	uservideo := &UserVideo{UserID: AuthorID, VideoID: videoID}
	if err := tx.Table("user_video").Create(uservideo).Error; err != nil {
		log.Println("插入 user_video 数据时出错：", err)
		tx.Rollback()
		return -1, err
	}
	tx.Commit()

	Searcher.IndexDocument(uint64(videoID), types.DocumentIndexData{
		Content: video.Description, // Weibo结构体见上文的定义。必须是UTF-8格式。
	}, false)
	Searcher.FlushIndex()

	return videoID, nil
}
