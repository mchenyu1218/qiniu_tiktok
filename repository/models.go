package repository

import (
	"time"
)

type User struct {
	ID            int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	Username      string    `gorm:"column:username;NOT NULL" json:"username"`
	Password      string    `gorm:"column:password;NOT NULL" json:"password"`
	Gender        int       `gorm:"column:gender;default:0;NOT NULL" json:"gender"`                                                                                                                                                                                           // 0 男 1女
	FollowerCount int       `gorm:"column:follower_count;default:0;NOT NULL" json:"follower_count"`                                                                                                                                                                           // 粉丝数
	FollowCount   int       `gorm:"column:follow_count;default:0;NOT NULL" json:"follow_count"`                                                                                                                                                                               // 关注数
	Face          string    `gorm:"column:face;default:https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg/800px-Everest_North_Face_toward_Base_Camp_Tibet_Luca_Galuzzi_2006.jpg;NOT NULL" json:"face"` // 头像地址
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	IsDeleted     int       `gorm:"column:is_deleted;default:0;NOT NULL" json:"is_deleted"` // (0-, 1-)

}

// `video`
type Video struct {
	ID            int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`       // id
	AuthorID      int64     `gorm:"column:author_id;NOT NULL" json:"author_id"`           // 作者id
	Description   string    `gorm:"column:description" json:"description"`                // 描述
	PlayUrl       string    `gorm:"column:play_url;NOT NULL" json:"play_url"`             // 播放地址
	CoverUrl      string    `gorm:"column:cover_url;NOT NULL" json:"cover_url"`           // 封面地址
	Title         string    `gorm:"column:title;NOT NULL" json:"title"`                   // 标题
	FavoriteCount int       `gorm:"column:favorite_count;NOT NULL" json:"favorite_count"` // 点赞数
	PlayCounts    int       `gorm:"column:play_counts;NOT NULL" json:"play_counts"`       // 播放次数
	CommentCount  int       `gorm:"column:comment_count;NOT NULL" json:"comment_count"`   // 评论次数
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	IsDeleted     int       `gorm:"column:is_deleted;default:0;NOT NULL" json:"is_deleted"` // (0-, 1-)
	CollectCount  int       `gorm:"column:collect_count;NOT NULL" json:"collect_count"`     // 收藏数
}

// 点赞表(当用户对某个视屏点赞，则添加一条数据。取消点赞则删除对应数据)
type Favorite struct {
	ID         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	UserID     int64     `gorm:"column:user_id;NOT NULL" json:"user_id"`   // id
	VideoID    int64     `gorm:"column:video_id;NOT NULL" json:"video_id"` // id
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
}

// 收藏表(当用户对某个视屏点赞，则添加一条数据。取消点赞则删除对应数据)
type Collect struct {
	ID         int64     `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	UserID     int64     `gorm:"column:user_id;NOT NULL" json:"user_id"`   // id
	VideoID    int64     `gorm:"column:video_id;NOT NULL" json:"video_id"` // id
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
}
