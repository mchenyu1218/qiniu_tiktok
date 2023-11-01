package main

import (
	_ "Projectdouy/docs"
	"Projectdouy/route"
	"gorm.io/gorm"
)

// 定义全局的db对象，我们执行数据库操作主要通过他实现。
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
var Db *gorm.DB

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server tiktok server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		http://localhost:8081/
//	@BasePath	/v2

func main() {
	//初始化操作
	r := routes.SetRouter()

	r.MaxMultipartMemory = 128 << 20 //设置视频最大上传容量
	//添加中间件
	//r.Use(utils.JwtVerify)

	//启动
	r.Run(":8081")
}
