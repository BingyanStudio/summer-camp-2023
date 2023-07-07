# 🚀7.5 task

## 💡Things I Learned  
1. a tour of go   
2. [http crash course](https://www.youtube.com/watch?v=iYM2zFP3Zn0)   
3. [mozilla:overview of http](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview)   
4. Gin framework  
5. MVC model  

## 🖥️Go    

### channel  

1. definition   
    Channels in Go are a powerful feature that enable communication and synchronization between goroutines. A channel is a typed conduit through which you can send and receive values of a specified type. They provide a safe and efficient way for goroutines to exchange data and coordinate their execution.  
2. Buffered Channels  
    ~~~go 
    ch := make(chan int, 100)
    ~~~
    Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.  
3. unbuffered channels  
    the send and receive operations will always block until both the sender and receiver are ready to perform the operation.    
4. Close  
    A `sender` can close a channel to indicate that no more values will be sent.  
    check if it is closed:   
    ~~~go
    v, ok := <-ch   
    ~~~
    The loop for i := range c receives values from the channel repeatedly until it is closed.  

### select  
1. definition  
    The select statement lets a goroutine wait on multiple communication operations. A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.  
    ~~~go
    select {
    case <-channel1:
        // Code to execute when channel1 is ready
    case data := <-channel2:
        // Code to execute when channel2 is ready and receive data from channel2
    case channel3 <- value:
        // Code to execute when channel3 is ready to send value
    default:
        // Code to execute if no channel operation is ready
    }
    ~~~
    note: If multiple cases are ready, one of them is chosen at random. The default case is executed if no other case is ready.   
    ~~~go
    func fibonacci(c, quit chan int) {
        x, y := 0, 1
        for {
            select {
            case c <- x:
                x, y = y, x+y
            case <-quit:
                fmt.Println("quit")
                return
            }
        }
    }

    func main() {
        c := make(chan int)
        quit := make(chan int)
        go func() {
            for i := 0; i < 10; i++ {
                fmt.Println(<-c)
            }
            quit <- 0
        }()
        fibonacci(c, quit)
    }
    ~~~
### sync.Mutex  
1. definition  
    互斥锁，不想要不同的channel间进行通信   
2. code example  
    ~~~go 
    package main

    import (
        "fmt"
        "sync"
        "time"
    )

    // SafeCounter is safe to use concurrently.
    type SafeCounter struct {
        mu sync.Mutex
        v  map[string]int
    }

    // Inc increments the counter for the given key.
    func (c *SafeCounter) Inc(key string) {
        c.mu.Lock()
        // Lock so only one goroutine at a time can access the map c.v.
        c.v[key]++
        c.mu.Unlock()
    }

    // Value returns the current value of the counter for the given key.
    func (c *SafeCounter) Value(key string) int {
        c.mu.Lock()
        // Lock so only one goroutine at a time can access the map c.v.
        defer c.mu.Unlock()
        return c.v[key]
    }

    func main() {
        c := SafeCounter{v: make(map[string]int)}
        for i := 0; i < 1000; i++ {
            go c.Inc("somekey")
        }

        time.Sleep(time.Second)
        fmt.Println(c.Value("somekey"))
    }
    ~~~


## 👩‍💻http  

### http,https  

HTTP (Hypertext Transfer Protocol)is an application-layer protocol that allows the retrieval and transfer of resources between clients and servers.  
https is a secure version of HTTP. (ssl,tls)  

### http method 

1. GET   retrieve a resource from the server   
2. POST  o submit data to the server to create a new resource   
3. PUT  to update a resource  
4. DELETE  to delete a resource  

### http 请求的构成

1. 请求行（Request Line）：请求行包含了请求方法、请求的URL和HTTP协议的版本。它的格式通常是：<请求方法> <URL> <HTTP协议版本>。例如：GET /example HTTP/1.1。  
2. 请求头部（Request Headers）：请求头部包含了关于请求的附加信息，如客户端信息、认证凭据、所期望的响应格式等。请求头部以键值对的形式出现，每个键值对占据一行。常见的请求头包括User-Agent（标识客户端）、Content-Type（请求体的数据类型）、Authorization（身份验证凭据）等。  
3. 空行：请求头部和请求体之间由一个空行分隔，空行是一个仅包含回车换行符的行，用于标识请求头部的结束。  
4. 请求体（Request Body）：请求体是可选的，仅在使用POST、PUT等方法发送数据时才存在。它包含了请求的实际数据，如表单字段、JSON数据等。  

### http 响应的构成

1. 状态行（Status Line）：状态行包含了HTTP协议的版本、状态码和对应的状态信息。它的格式通常是：<HTTP协议版本> <状态码> <状态信息>。例如：HTTP/1.1 200 OK。  
2. 响应头部（Response Headers）：响应头部包含了关于响应的附加信息，如服务器信息、响应内容的类型、响应时间等。响应头部也以键值对的形式出现，每个键值对占据一行。常见的响应头包括Content-Type（响应体的数据类型）、Content-Length（响应体的长度）、Server（服务器信息）等。  
3. 空行：响应头部和响应体之间由一个空行分隔，空行是一个仅包含回车换行符的行，用于标识响应头部的结束。  
4. 响应体（Response Body）：响应体包含了服务器返回的实际数据，如HTML文档、JSON数据、图片等。响应体的内容和格式由Content-Type头部字段指定。 

### http status 

1. 1xx Informational
2. 2xx Success
3. 3xx Redirection
4. 4xx Client Error
5. 5xx Server Error

### 典型的http会话
在像 HTTP 这样的客户端——服务器（Client-Server）协议中，会话分为三个阶段：  
客户端建立一条 TCP 连接（如果传输层不是 TCP，也可以是其他适合的连接）。   
客户端发送请求并等待应答。  
服务器处理请求并送回应答，回应包括一个状态码和对应的数据。 

### 前端和后端如何发送、处理请求
1. 前端发送请求   
    前端可以使用JavaScript代码发送HTTP请求到后端服务器。  
2. 后端处理请求  
    后端服务器收到请求后，会根据请求的内容和目的，进行相应的处理。 
    具体如何判断请求：路由配置，http方法，请求头，请求体等（具体取决于后端语言和框架）  
3. 后端后续操作  
    查询数据库、调用api等  
4. 后端构造和发送响应   


## 🔺gin framework  

### 常用的结构、方法

1. gin.Context  
    用于获取请求的参数，构造响应等  
    常用的方法有Params(用于找到来自url的某个参数),query(查询参数),setcookies,.json/.html/(生成特定格式的response)  

## 🗺️MVC model  

1. definition  
    - Model:retrieving and manipulating data  
    - View:presenting the data to the user  
    - Controller:the intermediary between the model and the view   

## 🚩总结 

今天学的有点吃力，go的channel,goroutine这些半懂不懂，看例子勉强能看懂，让我自己写应该写不出来。  
gin的文档不知道从何下手，没看过官方文档这种东西，不知道从哪开始看起，今天晚上或者明天找一些视频教程看看吧。   
MVC和http只做了简单了解，MVC举的例子看不太懂。   
明天再多研究一下gin吧！