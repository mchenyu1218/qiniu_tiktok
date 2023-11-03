package service

import (
	common "Projectdouy/commom"
	"Projectdouy/repository"
	"fmt"
	"github.com/huichen/wukong/types"
	"log"
)

//服务类

var (
	videoDao = repository.NewVideoDaoInstance()
)

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

func Feedalluser(UserId any) ([]common.Video, error) {

	fmt.Println(UserId)
	videos, error := videoDao.Feedalluser(UserId)
	if error != nil {
		log.Println("videoDao.Feedalluser 出错")
	}

	return videos, error
}

func DeleteVideo(videoID string) error {
	fmt.Println(videoID)
	error := videoDao.DeleteVideo(videoID)
	if error != nil {
		log.Println("videoDao.Delete 出错")
	}
	return error
}

func UpdateVideo(videoID string, title, description, tag string) error {
	fmt.Println(videoID)
	error := videoDao.UpdateVideoDetails(videoID, title, description, tag)
	if error != nil {
		log.Println("videoDao.Delete 出错")
	}
	return error
}

func InsertVideo(title, description, tag, playURL, coverURL string, authorID int64) (int64, error) {
	videoID, error := videoDao.AddVideo(title, description, tag, playURL, coverURL, authorID)
	if error != nil {
		log.Println("videoDao.Insert 出错")
	}
	return videoID, error
}

func SelectVideo(videoID int64) (common.Video, error) {
	video, error := videoDao.SelectVideoById(videoID)
	if error != nil {
		log.Println("videoDao.select 出错")
	}
	return video, error
}

func SearchVideo(query string) ([]common.Video, error) {
	response := repository.Searcher.Search(types.SearchRequest{
		Text:      query,
		Orderless: false,
	})

	var VideoList = make([]common.Video, len(response.Docs))
	for i := 0; i < len(response.Docs); i++ {
		temp, _ := videoDao.SelectVideoById(int64(response.Docs[i].DocId))
		VideoList = append(VideoList, temp)
	}
	return VideoList, nil
}
