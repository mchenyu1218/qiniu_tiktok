package repository

import (
	"sync"
)

var (
	videoOnce sync.Once
	videoDao  *VideoDao //实现多态调用
)

//相关数据访问对象,封装增删改查

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
