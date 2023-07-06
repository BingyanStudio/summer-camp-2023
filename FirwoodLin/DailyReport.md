# 07-03

-   Git
    -   clone 失败的原因：网络/Git credential配置错误/clone地址错误
    -   `git switch -c`的使用
-   Github
    -   在新 Branch Push 之后，可以在网页中提出 Pull Request
-   Golang 
    -   Golang 圣经阅读 - 结构体，函数参数，协程与通道；锁看了一部分

# 07-04

计划：翻完 Golang 圣经；了解设计模式；部署 MongoDB 和 Redis 环境。
实际情况：
-   了解了单次执行锁 `sync.once`
-   修改 Git Config，使 PR 界面展示直达链接
-   了解 MVC 架构
-   学习 session 和 cookie 的认证方式
-   配置 Arch Linux 中的 Golang 开发环境;修复了字符集缺失的问题;重装了 Linux QQ,闪退问题暂时解决
-   新建了热身项目文件夹（雾）

# 07-05

计划：实现登陆注册功能

实际情况：
-  了解了 Validator 的使用
-  Gorm 的使用
   -  检验用户是否存在：进行查询，通过 `.RowsAffected == 0` 判断查询的结果
-  Golang 开发规范
   -  bool 类型的函数，直接返回判断表达式的值；不使用 if 语句
-  （未完成）注册成功后返回 ~~cookie~~ 包含 sessionId 的 cookie
   -  Session 的生成、存储：由于高频读写，选择 Redis 作为数据库