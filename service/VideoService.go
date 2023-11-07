package service

import (
	common "Projectdouy/commom"
	"Projectdouy/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/huichen/wukong/types"
	"log"
	"time"
)

//服务类

var (
	videoDao = repository.NewVideoDaoInstance()
	ctx      = context.Background()
)

func Feedbyusername(Username string) ([]common.Video, error) {

	fmt.Println(Username)
	videos, error := videoDao.GetVideosbyusername(Username)
	if error != nil {
		log.Println("videoDao.Feedbytag 出错")
	}

	return videos, error
}

func TestCache(test string) (string, error) {
	err := Rdb.Set(ctx, "score", 50, 30).Err()
	if err != nil {
		log.Printf("set score failed, err:%v\n", err)

	}
	val, err := Rdb.Get(ctx, "score").Result()
	if err != nil {
		log.Printf("get score failed, err:%v\n", err)
	}
	fmt.Println("score", val)
	return val, err
}

func Feedbytag(amount int, Tag string) ([]common.Video, error) {
	// 构建 Redis 缓存键名
	cacheKey := "videos:" + Tag

	// 尝试从 Redis 缓存中获取数据
	cachedVideos, err := Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("cache hit!!!")
		// 缓存命中，将缓存的数据反序列化为 []common.Video
		var videos []common.Video
		if err := json.Unmarshal([]byte(cachedVideos), &videos); err != nil {
			log.Printf("Failed to unmarshal cached data: %v", err)
		}
		return videos, nil
	}

	// 缓存未命中，从数据库中获取数据
	videos, err := videoDao.GetVideosbytag(Tag)
	if err != nil {
		log.Println("videoDao.FeedByTag 出错")
		return nil, err
	}

	// 将数据库查询结果存入 Redis 缓存
	if len(videos) > 0 {
		videosJSON, err := json.Marshal(videos)
		if err == nil {
			Rdb.Set(ctx, cacheKey, videosJSON, time.Minute*30) // 缓存数据，有效期 30 分钟
		} else {
			log.Printf("Failed to marshal data for caching: %v", err)
		}
	}

	return videos, nil
}

func Feed() ([]common.Video, error) {
	// 构建 Redis 缓存键名
	cacheKey := "allVideos"

	// 尝试从 Redis 缓存中获取数据
	cachedVideos, err := Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// 缓存命中，将缓存的数据反序列化为 []common.Video
		var videos []common.Video
		if err := json.Unmarshal([]byte(cachedVideos), &videos); err != nil {
			log.Printf("Failed to unmarshal cached data: %v", err)
		}
		return videos, nil
	}

	// 缓存未命中，从数据库中获取数据
	videos, err := videoDao.GetallVideos()
	if err != nil {
		log.Println("videoDao.Feed 出错")
		return nil, err
	}

	// 将数据库查询结果存入 Redis 缓存
	if len(videos) > 0 {
		videosJSON, err := json.Marshal(videos)
		if err == nil {
			Rdb.Set(ctx, cacheKey, videosJSON, time.Minute*30) // 缓存数据，有效期 30 分钟
		} else {
			log.Printf("Failed to marshal data for caching: %v", err)
		}
	}

	return videos, nil
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

	// 删除成功后，删除相关缓存
	cacheKey := "allVideos"
	_, error = Rdb.Del(ctx, cacheKey).Result()
	if error != nil {
		log.Printf("Failed to delete cache for video %s: %v", videoID, error)
		// 不过这个错误不会影响主要功能，可以忽略
	}

	return error
}

func UpdateVideo(videoID string, title, description, tag string) error {
	fmt.Println(videoID)
	error := videoDao.UpdateVideoDetails(videoID, title, description, tag)
	if error != nil {
		log.Println("videoDao.Delete 出错")
	}

	// 更新成功后，删除相关缓存
	cacheKey := "allVideos"
	_, error = Rdb.Del(ctx, cacheKey).Result()
	if error != nil {
		log.Printf("Failed to delete cache for video %s: %v", videoID, error)
		// 不过这个错误不会影响主要功能，可以忽略
	}

	return error
}

func InsertVideo(title, description, tag, playURL, coverURL string, authorID int64) (int64, error) {
	videoID, error := videoDao.AddVideo(title, description, tag, playURL, coverURL, authorID)
	if error != nil {
		log.Println("videoDao.Insert 出错")
	}
	// 增加成功后，删除相关缓存
	cacheKey := "allVideos"
	_, error = Rdb.Del(ctx, cacheKey).Result()
	if error != nil {
		log.Printf("Failed to delete cache for video %s: %v", videoID, error)
		// 不过这个错误不会影响主要功能，可以忽略
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
