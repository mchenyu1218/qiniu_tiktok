package controller

import (
	common "Projectdouy/commom"
	"Projectdouy/service"
	"fmt"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//type UserPublishRequest struct {
//	Token string `json:"token" binding:"required"`
//	Data  []byte `form:"data" json:"data" binding:"required"`
//	Title string `form:"title" json:"title" binding:"required"`
//}

var (
	// 视频图片格式判断

	pictureIndexMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

// @Summary 发布用户视频
// @Produce	json
// @Param		username	query		string				false	"用户名"
// @Param		oldpassword	query		string				false	"旧的密码"
// @Param		newpassword	query		string				false	"新的密码"
// @Success	200		{object}	common.Response        "成功"
// @Failure	400		{object}	common.Response	"请求错误"
// @Failure	500		{object}	common.Response	"内部错误"
// @Router		/publish/action/:UserID [post]
func Publish(c *gin.Context) {

	title := c.PostForm("title")
	form, err := c.MultipartForm()
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}
	files := form.File["data"]

	// 多个文件 文件信息计入数据库
	for _, file := range files {

		open, err := file.Open()
		if err != nil {
			PublishVideoError(c, "上传文件数据有误，无法读取")
			return
		}
		defer open.Close()
		size := file.Size
		bytes := make([]byte, size)
		if _, err := open.Read(bytes); err != nil {
			PublishVideoError(c, "文件读取错误")
			return
		}
		index := strings.LastIndex(file.Filename, ".")
		newfileName := file.Filename[0:index]

		UserIDAny := c.Param("UserID")
		UserID, _ := strconv.ParseInt(fmt.Sprintf("%v", UserIDAny), 0, 64)

		videoName := strconv.FormatInt(UserID, 10) + "-" + newfileName + ".mp4"
		coverName := strconv.FormatInt(UserID, 10) + "-" + newfileName + ".jpg"

		// 上传到数据库
		err = service.AddVideo(bytes, videoName, coverName, UserID, title)

		// 视频剪切
		fileUrl := "http://s38oif2dm.hn-bkt.clouddn.com/" + videoName
		tmpCoverUrl := "tmpCover/" + coverName
		err2 := ffmpeg.Input(fileUrl, ffmpeg.KwArgs{"ss": "1"}).
			Output(tmpCoverUrl, ffmpeg.KwArgs{"s": "368x208", "pix_fmt": "rgb24", "t": "1", "r": "1"}).
			OverWriteOutput().
			ErrorToStdOut().
			Run()
		if err2 != nil {
			log.Fatal(err)
		}
		// 文件转换为字节流文件
		openFile, err := os.Open(tmpCoverUrl)
		if err != nil {
			PublishVideoError(c, err.Error())
			return
		}
		var data []byte
		buf := make([]byte, 1024)
		for {
			// 将文件中读取的byte存储到buf中
			n, err := openFile.Read(buf)
			if err != nil && err != io.EOF {
				PublishVideoError(c, err.Error())
				return
			}
			if n == 0 {
				break
			}
			// 将读取到的结果追加到data切片中
			data = append(data, buf[:n]...)
		}

		// 视频封面上传到oss
		service.UploadDataToOSS(data, coverName)
		openFile.Close()
		err = os.Remove(tmpCoverUrl)
		if err != nil {
			log.Println("临时图片移除失败")
			return
		}
	}

	// 成功
	var response = &common.Response{}
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(200, response)
}

// 返回错误
func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, common.Response{StatusCode: 1, StatusMsg: msg})
}
