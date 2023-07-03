# 冰岩作坊 2023 夏令营任务书

## 前言

欢迎参加冰岩作坊夏令营！在本次夏令营中，我们为你安排了从常用开发工具、语言、框架到部署等一系列从理论到实战的任务，你将从简单的 Hello World 开始，一步步学习后端开发的相关知识，最后从 0 到 1 实现一个企业级在线商品交易平台项目。我们也将基于冰岩作坊多年技术经验，为你提供专属学习建议，助你掌握全流程开发能力。

不在夏令营的同学，如果对任务内容感兴趣，也可以自行学习。

## 要求

- 请认真看清楚任务书上的每一个字，按照要求完成。
- 遇到问题请先尝试自行查找资料，无法解决再来提问。
- 沟通交流很重要，如果你遇到困难，或者你需要一些指导等，请及时告诉我们。
- 写日报和周报，介绍每天学习了什么，以及适当记录你认为的重点即可。
- 日报、周报和代码都放在本仓库下你的 Github ID 的目录中，通过 PR 提交合并。
- 学习安排里的部分内容你可能已经掌握，直接跳过就好。
- 不一定要按顺序学习，有时候跳过一些内容/多线程学习效果会更好。
- 学习安排里的资料仅供参考，你可以选择其它你认为更好的教程。
- 夏令营时间有限，只需要按要求学习就好了，参考资料和书籍等可以在以后进一步学习。

## 学习安排

1. ### Markdown

学习使用 Markdown 语法，它非常的简单好用，是绝大多数开发者的选择。之后你的日报、周报等文档都应该使用 Markdown 编写。

参考资料：

- [Markdown Guide](https://www.markdownguide.org/)

中文资料：

- [Markdown 教程](https://markdown.com.cn/basic-syntax/)

2. ### Git

学习和整理 Git 的常用命令和基本操作流程，尝试在 Github 中创建一个仓库并进行管理。学会如何在 Github 上贡献代码（Fork、修改、提交、Pull Request）。

参考资料：

- [Git Official Document](https://git-scm.com/docs/gittutorial)
- [Make your first contributions on Github](https://github.com/firstcontributions/first-contributions/blob/main/README.md)

中文资料：

- [廖雪峰的 Git 教程](https://www.liaoxuefeng.com/wiki/896043488029600)
- [Commit message 和 Change log 编写指南 - 阮一峰的网络日志](https://www.ruanyifeng.com/blog/2016/01/commit_message_change_log.html)
- [在 Github 上做出你的第一个贡献](https://github.com/firstcontributions/first-contributions/blob/main/translations/README.zh-cn.md)

3. ### Golang

学习 Golang 的基础语法，挑选一个趁手的 IDE，使用 Go 编写一些简单的程序。

参考资料：

- [Go Official Website](https://golang.org/)
- [A Tour of Go](https://go.dev/tour/)

中文资料：

- [Go 指南](https://tour.go-zh.org/welcome/1)
- [Go 菜鸟教程](https://www.runoob.com/go/go-tutorial.html)

相关书籍：

- [前言 · Go语言圣经](http://books.studygolang.com/gopl-zh/)
- The Go Programming Language
- Effective Go
- Go Web 编程

**时间安排：Markdown 语法 + Git 简单使用 + Golang 语法预计 3 天**

4. ### HTTP

了解并学习 HTTP 的相关知识，如请求、状态码等。了解前端如何获取后端返回的数据，如何发送请求，后端如何根据前端发过来的请求，回应请求，如何辨别不同的请求。

> 这部分内容简单了解即可，做完了项目可以再慢慢看这本书。

参考资料：

- 图解 HTTP

5. ### Web 后端框架

在 Go Web 后端框架中选择一个学习即可，推荐 Gin 和 Echo。尝试用这个框架处理简单的网页请求。

参考资料：

- https://github.com/gin-gonic/gin
- https://echo.labstack.com/guide

6. ### 设计模式

了解一下 MVC 架构、Go 项目文件布局、编码规范等，这些有助于你更好的组织代码。

了解什么是 API 系统，如何设计 RESTful API，拓展 GraphQL。

参考资料：

- [Go 项目结构:如何在 Go 项目中使用 MVC ? | Go优质外文翻译 | Go 技术论坛](https://learnku.com/go/t/48112) 推荐阅读英文原文
- [GitHub - josephspurrier/gowebapp: Basic MVC Web Application in Go](https://github.com/josephspurrier/gowebapp)
- [RESTful API 设计指南 - 阮一峰的网络日志](https://www.ruanyifeng.com/blog/2014/05/restful_api.html)

延伸阅读：

- [GitHub - xxjwxc/uber_go_guide_cn: Uber Go 语言编码规范中文版. The Uber Go Style Guide .](https://github.com/xxjwxc/uber_go_guide_cn)

**时间安排：4+5+6 建议结合下面的项目任务边学边实现，不用过于深入的研究，预计 5 天**

7. ### 数据库

- MySQL
- MongoDB
- Redis
- PostgreSQL

Golang 使用 GORM：[GORM 指南](https://gorm.io/zh_CN/docs/index.html)

8. ### 身份验证、权限校验与加密

**认证：**

熟悉以下三种前后端认证方式，一般在登录时使用

- cookie
- session
- JWT：[User Authentication in Go Echo with JWT | WebDevStation](https://webdevstation.com/posts/user-authentication-with-go-using-jwt-token/)

**加密算法：**

- 对称加密
- 非对称加密
- 哈希算法

**时间安排：数据库 +** **JWT** **同样结合项目完成，预计 3 天**

## 热身项目 - **成员管理系统**

实现内容：

- 管理员和普通用户
- 用户注册和登录

  用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址

- 管理员

  - 删除普通用户

  - 获取一个成员、所有成员信息

- 普通用户

  - 更改个人信息

**时间安排：热身项目 结合上面边学边做其实挺快的，预计 2 天**

## 实战项目 - 商城系统

> 先做能做的，不必按顺序做。

基本功能：

- 用户可以买卖商品 
- 登录注册

  - 用户密码加密存储

  - 用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址等，可自定义。

- 商品按照类别查询

  如：商品类别：电子设备、书籍资料、宿舍百货、美妆护肤、女装、男装、鞋帽配饰、门票卡券、其他

- 根据关键词搜索商品
- 推荐用正则表达式
- 商品页面

  - 商品详细信息

​    标题、简介、价格等

- 图片

​    图片可以存在本地，或者使用七牛云存储

- 个人信息页（类似于名片，自己和他人都可以查看）

  - 个人基本信息

  - 浏览量（个人信息页被他人访问次数，可考虑去重）

进阶功能：

- 图片压缩

  浏览时显示压缩的小图片，详细页显示大一点的图片

- 收藏夹(类似于购物车)
- 商品浏览量、收藏量等
- 热门查询、最新查询

  **时间安排：实战项目 预计** **5** **天**

## 项目部署

### 1. 配置nginx

学习配置 nginx 做中间代理层，具体可从以下链接中选取部分学习，作为示例，夏令营之后可以好好研究，当然夏令营期间有时间也可以自行研究，遇到坑可以问我们。

- [nginx 配置简介 - 掘金](https://juejin.im/post/5ad96864f265da0b8f62188f)
- [openresty/nginx 实践 - 掘金](https://juejin.im/post/5aae659c6fb9a028d375308b)

### 2. 配置 docker

- [Docker 从入门到实践 - Ubuntu](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html)
- [Docker 实践 - 掘金](https://juejin.im/post/5b34f0ac51882574ec30afce)

### 3. 配置域名https (不要求)

前提：有已经备案的域名，有服务器

- [Let's Encrypt 给网站加 HTTPS 完全指南](https://ksmx.me/letsencrypt-ssl-https/?utm_source=v2ex&utm_medium=forum&utm_campaign=20160529)