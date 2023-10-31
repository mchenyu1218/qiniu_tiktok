package service

import (
	common "Projectdouy/commom"
	"Projectdouy/repository"
	"fmt"
	"log"
)

//服务类

var (
	videoDao = repository.NewVideoDaoInstance()
)

// Feed 返回指定投稿时间之后的amount个视屏
func Feedbyusername(Username string) ([]common.Video, error) {

	fmt.Println(Username)
	videos, error := videoDao.GetVideosbyusername(Username)
	if error != nil {
		log.Println("videoDao.Feedbytag 出错")
	}

	return videos, error
}

func Feedbytag(amount int, Tag string) ([]common.Video, error) {

	fmt.Println(Tag)
	videos, error := videoDao.GetVideosbytag(Tag)
	if error != nil {
		log.Println("videoDao.Feedbytag 出错")
	}

	return videos, error
}

func Feed() ([]common.Video, error) {

	videos, error := videoDao.GetallVideos()
	if error != nil {
		log.Println("videoDao.Feed 出错")
	}

	return videos, error
}

func DeleteVideo(videoID string) error {
	fmt.Println(videoID)
	error := videoDao.DeleteVideo(videoID)
	if(error != nil){
		log.Println("videoDao.Delete 出错")
	}
	return error
}

func UpdateVideo(videoID string, title, description, tag string) error {
	fmt.Println(videoID)
	error := videoDao.UpdateVideoDetails(videoID, title, description, tag)
	if(error != nil){
		log.Println("videoDao.Delete 出错")
	}
	return error
}
