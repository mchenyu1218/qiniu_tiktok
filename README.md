# qiniu_tiktok
## 项目描述
这个项目是一个仿抖音的短视频应用，旨在为用户提供一个社交媒体平台，让他们可以轻松创建、分享和发现有趣的短视频内容。该应用具有类似抖音的特性，包括用户注册和登录、浏览和互动等功能。用户可以通过应用上传自己的短视频，还可以与其他用户互动，点赞、评论和关注他们的内容。

## 如何运行
### 先决条件
在开始之前，确保已经安装以下软件：
Go, Mysql, Redis, Node18
### 步骤
1. 克隆项目仓库到本地：
   ```bash
   git clone https://github.com/mchenyu1218/qiniu_tiktok.git
2. 执行data目录中的douyin.sql文件
3. 在front_end目录下安装前端依赖
    ```base
    npm install
4. 启动前端
    ```bash
    npm run serve
5. 修改repository/init.go和service/init.go文件当中的数据库地址和密码 
6. 启动项目：
    ```base
    go run main.go

### 架构图
![image](structure.png)

### API文档
1. /tiktok/feedallvideo 不用传入参数，查找出所有的视频
2. /tiktok/feedbytag?tag=hot
3. /tiktok/feedbyusername 查询这个用户发布的视频
4. /tiktok/feedalluser  //查询视频的同时查询关注列表和是否关注该博主
5. /tiktok/deleteVideo 根据视频id删除相关视频
6. /tiktok/updateVideo 根据视频id更新相关视频
7. /tiktok/insertVideo 根据视频id插入相关视频
8. /tiktok/selectVideo 根据视频id搜索相关视频
9. /tiktok/searchVideo 根据视频描述搜索相关视频
10. /tiktok/user/register 用户注册
11. /tiktok/user/login 用户登录
12. /tiktok/user/modeify 用户修改对应密码
13. /tiktok/publish/action/:UserID 发布视频
14. /tiktok/collect/action 收藏视频 1的时候收藏，2的时候取消收藏
15. /tiktok/collect/list 查看自己收藏哪些视频
16. /tiktok/favorite/action/ 点赞视频 1的时候点赞，2的时候取消点赞
17. /tiktok/relation/action 关注该用户， 1的时候关注，2的时候取消关注

### 分工
前端：黄义行
后端：胡净皓，毛辰宇
### demo视频地址
Youtube: https://youtu.be/DkmrDBqbIVs

