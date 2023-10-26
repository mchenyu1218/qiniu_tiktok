package main

import (
	"Projectdouy/route"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
var Db *gorm.DB

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	Db = _db
	sqlDB, _ := Db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	//初始化操作
	r := routes.SetRouter()

	r.MaxMultipartMemory = 128 << 20 //设置视频最大上传容量

	//启动
	r.Run(":8081")
}
