# 极简抖音后端API设计
## 项目说明
2022年字节青训营后端实践项目之一。

项目源地址：https://github.com/vlinglandy/qxy-dou-yin

[题目链接](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)

[接口文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

本项目为团队合作开发，我负责视频流接口、投稿接口以及发布列表接口的开发。

团队为6人小组，代码贡献者为4人。
|功能项|说明|
|:--|:--|
|视频Feed流、视频投稿、个人信息|支持所有用户刷抖音,按投稿时间倒序推出，登录用户可以自己拍视频投稿，查看自己的基本信息和投稿列表。|
|点赞列表、用户评论|登录用户可以对视频点赞，并在视频下进行评论，在个人项能够查看点赞视频列表。|
|关注列表、粉丝列表|登录用户可以关注其他驴，能够在个人信息页查看本人的关注数和粉丝数,击打开关注列表和粉丝列表。|





## 目录结构


```
│  .env
│  .gitignore
│  e-r.png
│  ffmpeg.exe
│  go.mod
│  go.sum
│  main.go
│  README.md
│
├─api            
│      comment.go
│      demo.go
│      favorite.go
│      feed.go
│      main.go
│      publish.go
│      relation.go
│      user.go
│
├─auth
│      .keep
│
├─conf
│      conf.go
│
├─middleware
│      auth.go
│      cors.go
│
├─model
│      comment.go
│      favorite.go
│      follow.go
│      init.go
│      migration.go
│      user.go
│      video.go
│
├─public
│      0_VID_20220613_150918.jpeg
│      0_VID_20220613_150918.mp4
│
├─serializer
│      common.go
│      entity.go
│
├─server
│      router.go
│
└─util
        common.go
        logger.go
```
1. api           MVC框架的controller，负责协调各部件完成任务
2. model         存储数据库模型和数据库操作相关的代码
3. util          通用的工具
4. conf          环境配置初始化
5. public        静态资源目录
6. e-r.png       数据库的e-r图
7. .env          环境配置
8. middleware    中间件，实现了登陆验证
9. serializer    一些序列化统一格式的操作
10. server       要挂载的路由
11. ffmpeg.exe   视频工具（上传视频时会用到）

## 项目本地运行
1. 安卓端下载安装[极简抖音APP](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
2. 按照APP使用说明设置服务端地址
3. 下载[视频工具](https://ffbinaries.com/downloads)并解压到根目录
4. 启动后端服务
  - 初次运行，将`.env`中数据库连接设置为本地数据库
  - 将model/init.go中的`migration()`取消注释（首次运行后可注释）
  - 根目录下执行
 ```go
 go run main.go
 ```

