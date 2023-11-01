package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// 视频类
type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author" gorm:"references:ID"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	IsCollect     bool   `json:"is_collect"`

	Title        string `json:"title"`
	CollectCount int64  `json:"collect_count"`
}

// 用户类
type User struct {
	Id            int64  `json:"id" gorm:""`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`

	IsFollow bool `json:"is_follow"`
}
