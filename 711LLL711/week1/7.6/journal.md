# 🚀7.6 task

## 💡Things I Learned  

1. 尝试用jin写简单的前后端交互框架    
    [后端处理前端传的参数、返回页面给前端、设置中间件的代码](http_project/main.go)   
    参考：   
    - [狂神说--一小时上手gin框架](https://www.bilibili.com/video/BV1Rd4y1C7A1/?spm_id_from=333.1007.top_right_bar_window_history.content.click)   
    - [尚硅谷goweb实战](https://www.bilibili.com/video/BV1nJ411D7P4?p=1&vd_source=0eed6cf5f9626548f201d318112f2680)   
    - [教程--用gin实现http应用](https://golang2.eddycjy.com/posts/ch2/01-simple-server/)  
    - [李文周的博客--gin框架](https://www.liwenzhou.com/posts/Go/Gin_framework/)  
2. bug解决--import gin爆红  
    命令行修改了go proxy为国内代理，`go env -w GOPROXY=https://goproxy.cn,direct，`  
    但是检查了一下在vscode设置中没有改动，所以又在setting.json中添加了：   
    ~~~json
    "go.proxyToolsCommand": "GOPROXY=https://proxy.golang.org,direct",
    "go.useLanguageServer": true
    ~~~
    重启vscode后，问题解决。  
3. 简单了解restfulapi,MVC，JWT   
4. 在b站大学学习sql基础（还没学完）   

##  🍸gin  

### gin demo 分析  

1. demo源码
    ~~~go
    package main

    import (
        "net/http"
        "github.com/gin-gonic/gin"
    )

   func main() {
        r := gin.Default()//创建路由
        r.GET("/ping", func(c *gin.Context) {//调用GET方法的reponse
            c.JSON(200, gin.H{"message": "pong"})
        })//返回JSON格式的数据
        r.Run()//默认监听端口8080
    }

2. 用到的实例方法
- gin.Default  
    我们会通过调用 gin.Default 方法来创建默认的 Engine 实例，它会在初始化阶段就引入 Logger 和 Recovery 中间件，能够保障你应用程序的最基本运作，这两个中间件具有以下作用：   

    Logger：输出请求日志，并标准化日志的格式。   
    Recovery：异常捕获，也就是针对每次请求处理进行 recovery 处理，防止因为出现 panic 导致服务崩溃，并同时标准化异常日志的格式。
- gin.New   
    Engine 实例就像引擎一样，与整个应用的运行、路由、对象、模板等管理和调度都有关联  
    ~~~GO
    func New() *Engine{
        ...
    }
    ~~~
- r.GET  
    注册路由，路径对应的reponse   
- gin.H  
    一个map[string]interface{}类型的别名，用于生成JSON数据
3. 简单的前后端交互代码实现    
[后端处理前端传的参数、返回页面给前端、设置中间件的代码](http_project/main.go)   
    

## 🔥大致了解MVC,JWT,API

### MVC设计架构示例  

参考：[github--gowebapp](https://github.com/josephspurrier/gowebapp)    
~~~
config		- application settings and database schema
static		- location of statically served files like CSS and JS
template	- HTML templates

vendor/app/controller	- page logic organized by HTTP methods (GET, POST)
vendor/app/shared		- packages for templates, MySQL, cryptography, sessions, and json
vendor/app/model		- database queries
vendor/app/route		- route information and middleware
~~~

### RESTful API   

api--客户端与web上资源间的大门  
restful api组成:url,方法，数据,参数...

 
### JWT  

1. 组成部分:header（头部）,playload（有效载荷）,signature（签名）   
    - header:  
        ~~~json
        {
            "alg": "HS256", //加密算法
            "typ": "JWT"    //令牌类型
        }
        ~~~
    - playload:  
        ~~~json
        {
            "sub": "1234567890", //主题
            "name": "John Doe",  //用户名
            "iat": 1516239022 ,   //签发时间issue at
            "aud": "www.example.com",   //接收方audience
            "iss": "www.example.com",   //签发方issuer
            "exp": 1516239022,  //过期时间expire

        }
        ~~~
    - signature:   
        签名将会用于校验消息在整个过程中有没有被篡改，并且对有使用私钥进行签名的令牌，它还可以验证 JWT 的发送者是否它的真实身份。  
        ~~~
        HMACSHA256(
            base64UrlEncode(header) + "." +
            base64UrlEncode(payload),
            secret)     //secret是保存在服务器端的，用于验证签名的密钥
        ~~~
2. 安装JWT  
    ~~~
    go get -u github.com/dgrijalva/jwt-go
    ~~~
3. 在go中使用JWT(🔴TODO)


## 👩‍💻mysql简单使用

### 启动、停止、连接客户端

~~~
net start mysql80(是注册在windows服务中的名字) 
net stop mysql80
mysql -u root -p
~~~
下载并学会使用图形化工具dataGrip   

### sql语句

1. 基本语法
- 单行或多行书写，分号结尾，不区分大小写
- 注释：单行注释：#，--，多行注释：/* */
2. sql语句分类
- DDL：数据定义语言，用来定义数据库对象：数据库，表，列等。关键字：create，drop，alter等 
    ~~~sql  
    show databases;   --查看所有数据库
    create database test;  --创建数据库
    use test;--使用数据库
    drop database test;--删除数据库
    ~~~

    表操作
    ~~~sql
    show tables;--查看所有表
    desc student;--查看表结构
    show create table student;--查看创建表的语句
    --创建表
    create table student(
        id int comment '学号',--comment注释可选
        name varchar(20),
        age int,
        address varchar(20)
    )[comment表注释];
    --添加字段
    alter table 表名  add 字段名 类型 [comment '注释'];
    --修改字段名和字段类型
    alter table 表名 change 旧字段名 新字段名 新类型 [comment '注释'];
    --修改字段类型
    alter table 表名 modify 字段名 新类型 [comment '注释'];
    --删除字段
    alter table 表名 drop 字段名;
    --修改表名
    alter table 表名 rename 新表名;
    --删除表
    drop table 表名;
    --删除指定表并重新创建空表
    truncate table 表名;
    ~~~

- DML：数据操作语言，用来对数据库中表的数据进行增删改。关键字：insert，delete，update等   
    ~~~sql
    --插入数据
    insert into 表名(字段名1,字段名2,...) values(值1,值2,...);
    --删除数据
    delete from 表名 where 条件;
    --修改数据
    update 表名 set 字段名=值 where 条件;
    ~~~
- DQL：数据查询语言，用来查询数据库中表的记录（数据）。关键字：select，where等    
    ~~~sql
    select
        [distinct] 字段名1,字段名2,...
    --不重复
    select distinct 字段名1,字段名2,...
    from 
        表名1,表名2,...
    where
        条件
    -- 所有条件逻辑运算符：=,>,>=,<,<=,<>,!=,between,and,or,in,not in,like,not like,is null,is not null  
    group by 
        分组字段列表
    having
        分组后条件列表
    order by
        排序字段列表
    limit
        起始行，行数
    ~~~ 
   
- DCL：数据控制语言，用来定义数据库的访问权限和安全级别，及创建用户。关键字：GRANT，REVOKE等   

## 🌃总结  

今天主要研究了gin框架，没有刚接触时那么懵了，可以照着定义编写一些简单的前后端交互代码😀  
对于MVC,JWT,数据库这些，不确定要学到什么程度，了解得泛泛的   

🔴TODO LIST：  
- jwt在go中的使用，其他验证方式：cookie,session   
- mysql基础看完   
- gorm 
- 简单了解加密算法   
