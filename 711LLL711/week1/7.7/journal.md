# 🚀7.7 task

## 💡Things I Learned  

- mysql基本语句、函数   
- 在go中操作mysql数据库    
    [代码demo](project/main.go)   
    参考:  
    [Go 语言操作 MySQL 之 CURD 操作](https://segmentfault.com/a/1190000023067651)   
- 用go使用jwt,cooke,session   

##  🍸mysql

### DQL查询语句
1. DQL常见聚合函数   

    将一列数据作为整体，进行纵向的计算，得到一个结果。  
    count,min,max,sum,avg   
    ~~~sql
    select  () count(字段) from 表名; --计算表中记录数 
    ~~~
2. 分组查询   
    ~~~SQL
    select ... from ... where...（查询条件）group by ... having ...(分组后的过滤条件)   
    ~~~
3. 排序查询   
    ~~~SQL
    select ... from ... where...（查询条件）order by ... (排序的字段) asc/desc(升序/降序) 
    ~~~ 
4. 分页查询   
    note:不同的数据库实现的语句不同
    ~~~SQL
    select ... from ... where...（查询条件）limit ... (起始行，行数)   
    ~~~
    起始索引 = （页码-1）*每页记录数     
    ~~~sql
    select * from 表名 0,10;--查询第一页，每页展示10条记录  
    ~~~

### DCL管理用户 
1. 查询用户   
    ~~~sql
    USE mysql; --切换到mysql数据库
    SELECT user,host FROM user; --查询用户  
    ~~~
2. 创建用户   
    ~~~sql
    CREATE USER '用户名'@'主机名' IDENTIFIED BY '密码'; --创建用户
    --@'主机名'可以用%代替，代表任意主机  @localhost代表本地主机   
    ~~~
3. 修改用户
    ~~~sql
    --修改用户密码
    alter user '用户名'@'主机名' identified by '新密码';
    --修改用户权限
    GRANT 权限列表 ON 数据库名.表名 TO '用户名'@'主机名';
    -- 查询权限
    SHOW GRANTS FOR '用户名'@'主机名';
    -- 撤销权限
    REVOKE 权限列表 ON 数据库名.表名 FROM '用户名'@'主机名';
    ~~~
4. 删除用户
    ~~~sql
    DROP USER '用户名'@'主机名';
    ~~~

### 常见内置函数 
1. 字符串函数    
2. 数值函数    
3. 日期函数   
4. 流程函数   
    ~~~sql
    select
        name,
        (case when age<18 then '未成年' else '成年' end) as '是否成年'
    from student;
    ~~~

### 约束

## 👩‍💻在go中操作mysql数据库 
1. 导入驱动包   
    ~~~go
    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
    )
    ~~~
2. CURD示例   
    [代码见demo](project/main.go) 

3. 常用的对象方法
- *sql.Stmt 对象具有以下特性和方法：   
    Exec(args ...interface{}) (sql.Result, error):   
        执行准备好的语句并返回 sql.Result 对象和可能的错误。   
    Query(args ...interface{}) (*sql.Rows, error):   
        执行准备好的查询语句并返回 *sql.Rows 对象和可能的错误。  
    QueryRow(args ...interface{}) *sql.Row:   
        执行准备好的查询语句并返回单行结果的 *sql.Row 对象。   
    Close() error:   
        关闭准备好的语句，释放相关的数据库资源。   

    **❗❗在使用完 *sql.Stmt 对象后，一定要调用 Close() 方法来显式关闭它，以释放相关的数据库资源** 

- db.Prepare() 方法返回的是 *sql.Stmt 对象  
    使用prepare的优点是可以预处理，参数化查询，防止sql注入，使代码可读性更高。  

## 🔥cookie
1. cookie的作用   
    cookie是服务器发送到用户浏览器并保存在本地的一小块数据，它会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上。   
    通常，它用于告知服务端两个请求是否来自同一浏览器，如保持用户的登录状态。   
 2. 在go中使用cookie-- 利用net/http包    
    ~~~go
    //设置cookie
    http.SetCookie(w, &http.Cookie{
    Name:  "username",
    Value: "john",
    })

    //读取cookie
    cookie, err := r.Cookie("username") //r 是一个 http.Request 对象，它表示一个 HTTP 请求。
    if err == nil {
    fmt.Println(cookie.Value)
    }

    //删除cookie
    http.SetCookie(w, &http.Cookie{
    Name:   "username",
    Value:  "",
    MaxAge: -1,
    })
    
    //处理http请求时获取cookie
    func handler(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("username")
        if err != nil {
            if err == http.ErrNoCookie {
                // 未设置
                ...
            }
            // 其他错误
            ...
        }
        // 获取到了
        ...
    }
    func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
    }
    ~~~

## 💭session
1. 用于跟踪用户的身份验证状态、存储用户信息和会话数据等。  
2. 在go中使用session--用gorilla/sessions
    ~~~go
    package main

    import (
        "net/http"
        "github.com/gin-gonic/gin"
        "github.com/gorilla/sessions"
    )

    var store = sessions.NewCookieStore([]byte("secret-key"))//创建session存储对象

    func main() {
        r := gin.Default()//创建路由

        r.Use(sessions.Middleware(store))//设置中间件

        r.GET("/set", func(c *gin.Context) {
            session := sessions.Default(c)
            session.Set("username", "john")
            session.Save()
            c.String(http.StatusOK, "Session set")
        })//获取会话对象，设置会话值，保存会话

        r.GET("/get", func(c *gin.Context) {
            session := sessions.Default(c)
            username := session.Get("username")
            c.String(http.StatusOK, "Username: %s", username)
        })//获取会话对象，获取会话值

        r.Run(":8080")
    }

    ~~~

## 💤JWT使用  

### JWT基本流程  
1. 客户端进行身份验证，例如提供用户名和密码。  
2. 服务器验证身份信息，并在成功验证后生成 JWT。  
3. 服务器将 JWT 返回给客户端，客户端在后续请求中携带 JWT。  
4. 服务器接收到 JWT，使用密钥验证 JWT 的签名和有效性。  
5. 如果 JWT 验证成功，服务器根据 JWT 中的信息执行相应的操作，例如授权访问受保护的资源。  

### go实现JWT--使用jwt-go包
参考：
    [golang-JWT](https://golang-jwt.github.io/jwt/usage/create/)
1. create代码
    ~~~go
    //对称加密实现
    var (
    key *ecdsa.PrivateKey
    t   *jwt.Token
    s   string
    )

    key = /* Load key from somewhere, for example a file */
    t = jwt.New(jwt.SigningMethodES256) 
    s = t.SignedString(key) 

    //with claims
    var (
    key *ecdsa.PrivateKey
    t   *jwt.Token
    s   string
    )

    key = /* Load key from somewhere, for example a file */
    t = jwt.NewWithClaims(jwt.SigningMethodES256, 
    jwt.MapClaims{ 
        "iss": "my-auth-server", 
        "sub": "john", 
        "foo": 2, 
    })
    s = t.SignedString(key) 

    ~~~
2. 选择签名方法   
HMAC对称签名，签名和验证密钥类型都是[]byte   


## 🌃总结 
学了怎么在go中用mysql,cookie,session,jwt  
TODO:  
1. 加密算法  
2. GORM  
3. 开始做热身项目  