package controller

import (
	common "Projectdouy/commom"
)

//调用service层的服务

type VideoFeedRequest struct {
	LatestTime int64  `form:"latest_time" json:"latest_time" validator:"omitempty"`
	Token      string `form:"token" json:"token" validator:"omitempty"`
}

type VideoFeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
	NextTime  int64          `json:"next_time"`
}
