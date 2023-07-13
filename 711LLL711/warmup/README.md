# 热身项目介绍和总结 

## 结构 
~~~ 
system-----app---controller                 //登录、注册等的逻辑代码    
        |   |----route  ----router.go       //路由设置    
        |   |            |---middlaware     //定义判断登陆状态和身份的中间件    
        |   |----shared  ----database       //数据库连接    
        |                |---session        //session设置    
        |                |---user           //定义数据库中的结构体   
        |                |---server         //初始化服务器和全局中间件    
        |
        |--config                           //数据库信息     
        |--main.go                          //入口文件，负责连接数据库，初始化服务器，load前端页面    
        |--templates                        //前端页面     
~~~
## 路由设置

1. 注册、登录、更新--这三个都是GET获取前端表单，再POST传回后端  
    - /login
    - /register
    - /update
2. 仅管理员操作的页面
    - /admin/del?id=    删除用户GET  
    - admin/getalluser  获取所有用户GET     
    - admin/getuser?id= 获取单个用户DELETE   


## 中间件设置、数据库设置  

通过session判断登录状态，中间件包括验证是否登录，验证是否是管理员    
数据库信息在config里的JSON文件中，通过读取JSON文件获取数据库信息，连接数据库。  
数据库的信息：id,password(存储哈希加密过的密码),name,phone,email，role   
数据库使用的是go-mysql-driver包，为什么不用GORM？就不需要写那么多的sql语句了555   
~~~sql
create table users(
    id varchar(20) comment '用户ID',
    password varchar(80) comment '密码',
    name varchar(20) comment  '昵称',
    phone varchar(11) comment '手机号',
    email varchar(30) comment '邮箱' ,
    role ENUM('admin', 'user') comment '用户类型'
)comment '用户信息表';
~~~


## 总结写出的bug

1. 导出的变量和函数要大写开头，统一命名，避免大小写混乱导致写错的问题   
2. main包不能导出，各个包之间分的尽量减少耦合，避免循环导入，模块化设计，我感觉这个项目分的不是很好，各种包导来导去的,有的函数不知道该放在哪个包里...   
3. `c.Abort()` 可以中止请求的处理过程，但它并不会阻止用户访问页面。当调用 `c.Abort()` 后，Gin 会停止执行后续的中间件和处理函数，但仍然会返回响应给用户。当需要完全阻止用户访问或重定向时，应该使用 `c.AbortWithStatus() 或 c.Redirect()`。 
4. session or JWT?   
    这次只用了session,管理用户登录状态还是挺简单方便的，session ID通常通过Cookie在客户端进行存储。当用户发送后续请求时，服务器可以通过session ID 来识别用户，并从服务器存储中检索相关的会话数据。  
    JWT可以无状态跨系统验证用户的身份和权限，更适用于微服务、分布式的应用程序  
5. 数据库操作失败？  
    检查变量类型、大小是否匹配，sql语法是否正确，是否已经连接数据库   
    哈希加密过的密码是60字符，开始在数据库中设置的是varchar(20)，导致密码存储失败     
6. Query和Param获取参数方法搞混   


## 可以改进

1. session没有设置过期时间，使用的是默认的。gorilla.sessions可以通过设置Options字段中的MaxAge属性来定义Session的过期时间。将MaxAge设置为负值表示Session在浏览器关闭后立即失效。   
2. update逻辑有点问题，非管理员应该设置只能修改自己的信息，这一点没有注意   
3. 设置可以通过手机号验证码登录   
4. 当检测到用户未登录设置了重定向到登录页面，但是应该要记录用户访问的url，登录后要重定向回来
