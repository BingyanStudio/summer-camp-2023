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

# 07-08

-   最爱的周末

# 07-10

- 新发现
  - 发现了可以学习的开源项目 [link](https://github.com/newbee-ltd/newbee-mall-api-go)
  - 发现`copier`库，可以将两个结构体之间的同名字段进行复制（方便DTO和DAO互转）
- GORM 的使用
  - 外键的使用
    - 定义结构体时：添加
    - 在模型结构体中添加`TableName()`方法，可以自定义表名——我们应该将同一张表的模型定义在一个文件当中
    - 示例项目中使用了手写 SQL，没有使用外键
    - 在最后一次试验中，实现了外键查询，但是没有搞明白原理
- 模型的设计
  - 函数的参数可以封装为`xxxParam`结构体

# 07-11

- 数据库设计相关

  - 价格使用 int 存储，单位为分
  - 外键不是必须的，只是在数据库层面保证约束条件奏效；也可以在程序逻辑层面实现，而不使用外键
    - 但是定义外键可以在一定程度上简化查询，无需指定联合条件（使用 GORM 时）
  - 参考闲鱼“买卖”功能：可以在 Order 表中添加`userID`和`buyerID`实现数据共享；每个用户同时是买家&卖家

- GORM 使用相关

  - ```go
    // 指定了模型：从全局的 DB 到 local 的 db
    db := global.GVA_DB.Model(&manage.MallGoodsInfo{})
    // 指定了查询条件
    db.Where("goods_category_id= ?", goodsCategoryId)
    // 进行查询
    err = db.Count(&total).Error
    // 进行排序
    db.Order("goods_id desc")
    // 存储结果 *** 注意使用存储结果到返回值当中！！！否则只是改变了 db 的属性而已
    err = db.Limit(limit).Offset(offset).Find(&goodsList).Error
    ```

    上述代码实际转化为 MySQL 语句：

    ```mysql
    -- 查询符合条件的记录数
    SELECT COUNT(*) FROM `mall_goods_info` WHERE (goods_category_id=goodsCategoryId);
    -- 查询符合条件的记录，并按照 `goods_id` 字段降序排序，限制返回记录数和偏移量
    SELECT * FROM `mall_goods_info` WHERE (goods_category_id=goodsCategoryId) ORDER BY `goods_id` DESC LIMIT limit OFFSET offset;
    ```

    可以看到，WHERE 语句得到了保留，而 COUNT 语句没有继承。

- 将注册登陆功能（SESSIONID）进行了迁移，后期考虑使用 JWT 进行重构

- 商品查询功能（疑问：正则是怎么运用的）（model 层）

TODO：测试！测试！测试！

# 07-13

- 新知
  - RFC 7807 关于 RESTful API 中，返回错误时的格式规范
  - SQL 中的锁
- 开发
  - 完成商品查询功能（service 层），完成对数组类型的 Query 参数解析
  - 对登陆注册，分类查询，商品查询的接口进行了测试
  - 核心功能：下单（开发中）（商品的修改完成，还差新建订单）
- TODO
  - 下单时价格核验
  - 完成更新商品图片功能（对象存储 SDK 的操作，文件上传和下载）

# 07-14

- 开发
  - 完善订单提交事务 - 完成提交订单功能
    - 订单的设计：Order 存储买卖双方和支付信息；OrderItem 存储订单中的商品；每个 Order 中的商品都来自同一个卖家（前端负责）
    - TODO：进行价格校验（商品总价== $\sum$商品单价 ）
    - golang 中结构体=>`JSON`时（以`ABigApple`为例），默认将命名方式改为***小驼峰***`aBigApple`，而`GORM`默认改为***蛇形***(`a_big_apple`)，`form`字段（负责 JSON 和 XML 的解析）会转为***全小写***`abigapple`
      - 这给我们一个什么启示呢：尽管 convention better than configuration 的理念简化了代码，但是也要求对我们对 convention 足够的了解，否则还是自行指定 config 为妙

  - TODO:模拟对接支付接口（修改订单状态，对应修改商品状态；支付成功：商品销量++；支付失败：商品库存回滚）

    - 请求字段（节选）：

      | 字段名（*可选） | 描述                             |
      | --------------- | -------------------------------- |
      | notify_url      | 商城后端接口，接收支付状态信息   |
      | return_url      | 支付成功后，将用户重定向到的位置 |
      | out_trade_no    | 商家的自定义的订单号             |
      | total_amount    | 总金额                           |
      | *quit_url       | 用户取消支付后重定向的位置       |

    - 支付宝的响应字段：
  
      | 字段名            | 描述                                   |
      | ----------------- | -------------------------------------- |
      | code,msg          | 状态码，状态信息                       |
      | out_trade_no      | 商家的自定义的订单号                   |
      | trade_no          | 支付宝流水号（最长 64 位的交易流水号） |
      | total_amount      | 总金额                                 |
      | seller_id         | 收款支付宝账号                         |
      | merchant_order_no | 商户原始订单号（最长32位）             |
  
  - 上传图片并进行压缩
  
    - 进度：接受图片并保存到本地
  
- 新知
  - jetbrains 系列软件中文字体的设置：设置中搜索`font`，调整`fallback font`（用于补充主字体不支持的字符所使用的字体）
  - 设置 Google 搜索结果语言：給浏览器添加搜索引擎`https://www.google.com/search?q=%s&lr=lang_zh-CN|lang_en`
  - 支付信息的回调：支付宝：向用户指定的 URL POST 一个支付信息表单
  - JSON Schema 中 array 类型