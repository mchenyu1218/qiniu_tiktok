package service

import (
	common "Projectdouy/commom"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
)

var (
	accessKey = "HmWLfOMTQVlzn42VUOBMwiI0rPSu9Pc4xxYtl1ow"
	secretKey = "vXLUqqeNhr_Qa7BBqiQ3zuSpLM_XlLydQT8393ou"
	bucket    = "douyintik"
	// 域名
	domain = "s38oif2dm.hn-bkt.clouddn.com"
)

func AddVideo(data []byte, videoName string, coverName string, authorId int64, title string) error {

	// 存入七牛云oss对象存储
	err := UploadDataToOSS(data, videoName)
	if err != nil {
		return err
	}
	videoUrl := domain + "/" + videoName
	coverUrl := domain + "/" + coverName

	// 存入数据库
	err = videoDao.Addvideo(authorId, videoUrl, coverUrl, title)
	if err != nil {
		log.Println("发布文件失败，请稍后再试")
	}
	return err

}

// UploadDataToOSS 将本地文件上传到七牛云oss中
func UploadDataToOSS(data []byte, videoName string) error {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	//data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, videoName, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		fmt.Println("shu")

		return err
	}
	return nil

}

func PublishList(userId int64) ([]common.Video, error) {
	videoList, err := videoDao.GetPublishList(userId)
	if err != nil {
		return videoList, err
	}
	return videoList, nil
}
