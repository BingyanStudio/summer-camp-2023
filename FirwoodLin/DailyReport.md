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

# 07-06

-   实现注册功能
    -   SessionID 使用 UUID,Session 中存储用户 ID，过期信息
    -   在 redis 中存储哈希表：`HSET <HashKeyVal> <FieldName> <FieldVal>`。比如 `HSET person name laoyang`
-   登陆
    -   gin 中中间件的使用
        -   `c.Next()/c.Abort()`用于控制流程，`c.Get()/c.Set()`用于传递信息。
        -   `c.Next()`可以实现类似于“回调”的作用，在执行完后续函数后，会返回继续执行`.Next`后面的语句。
        -   `r.use()`使用全局中间键，`r.GET("地址"，中间件，函数)`使用局部中间键 
        -   `type HandlerFunc func(*Context)`可以看出，gin 中间件的的实质是具有特定参数类型的函数
    -   关于邮箱/手机号两种登陆方式的实现
        -   前端发送请求时添加 “LoginType” 字段，后端根据字段进行判断，并且进行校验
-   修改信息
    -   APIFox 中如何设置自动获取 Cookie [link](https://apifox.com/blog/cookies-and-token/)
        -   登陆后设置`后置操作`，将 Cookie 保存到环境变量中
        -   后续操作中，将需要 Cookie 的操作处，填为环境变量
    -   使用中间件校验 SessionId
    -   更新个人信息的时候使用`Updates`进行不确定 key 的数量的更新
    -   需要将 UserID 进行一系列传递，调试时可以打断点观察
        -   统一变量命名规则！！！（`UserID` vs `UserId`）


**杂记**
-   疑惑：golang 中使用文件夹名还是 package 名进行导入
-   TODO：调用 validate 函数进行数据校验

# 07-07

-   阶段性成果：热身项目基本完成 [link](https://github.com/FirwoodLin/Projects-BingyanSummer2023/tree/main/WarmUp)
    -   Session 相关
        -   中间件校验时，如果 Session 过期，进行重定向
        -   每次成功校验 Session,就延长 Session 的有效期（TODO）   
    -   查询相关
        -   查询单条记录要添加`Where`条件。批量查询不加条件
        -   使用`Table`指定查询表，`Select`方法指定要查询的字段，避免创建新模型和模型间的转化
    -   小功能：Session 延期；`viper`配置读取；完善了 API 文档的 Reponse 部分
    -   要继续学习的小点：项目中 Error 的规范；日志（log）系统的规范
    -   一些反思：先设计 API 接口，再进行实现。比如返回值的设计，在后期开发过程中产生了对`IsAdmin`字段的需求，但前期没有考虑到，导致需要进行重构。
    -   还可以改进的地方：使用邮箱**或者**手机号登陆

-   商城项目进展
    -   ~~新建文件夹~~
    -   了解商城项目数据库的设计