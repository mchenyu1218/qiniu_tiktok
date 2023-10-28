package main

import (
	"Projectdouy/route"
	"gorm.io/gorm"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
var Db *gorm.DB

func main() {
	//初始化操作
	r := routes.SetRouter()

	r.MaxMultipartMemory = 128 << 20 //设置视频最大上传容量

	//启动
	r.Run(":8081")
}
