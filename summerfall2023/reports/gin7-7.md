# gin7.7

## 路由

### 获取POST参数

1.curl命令

    curl http://localhost:9999/form  -X POST -d 'username=geektutu&password=1234'

-X选项用于指定HTTP请求的方法。在您的示例中，-X POST表示使用POST方法发送请求。

-d选项用于指定要发送的数据。在您的示例中，-d 'username=geektutu&password=1234'表示将用户名和密码作为表单数据发送到服务器。

c.JSON()是gin.Context结构体的方法，用于将数据以JSON格式作为响应发送给客户端。它接受两个参数：HTTP状态码和要发送的数据。

2.JSON

c.JSON()是gin.Context结构体的方法，用于将数据以JSON格式作为响应发送给客户端。它接受两个参数：HTTP状态码和要发送的数据。

### 混合

在Gin框架中，`Query`和`PostForm`方法都用于从HTTP请求中获取参数值，但它们的使用方式和获取参数的位置有所不同。

1. `Query`方法：
   - 用途：用于获取URL查询参数（query parameters）的值。
   - 位置：查询参数通常位于URL的问号后面，如`http://example.com/path?param1=value1&param2=value2`。
   - 使用方法：通过`c.Query("param")`来获取指定查询参数的值。
   - 示例：`id := c.Query("id")`表示获取名为"id"的查询参数的值。

2. `PostForm`方法：
   - 用途：用于获取POST请求中的表单参数（form parameters）的值。
   - 位置：表单参数通常位于HTTP请求的正文部分，而不是URL中。
   - 使用方法：通过`c.PostForm("param")`来获取指定表单参数的值。
   - 示例：`username := c.PostForm("username")`表示获取名为"username"的表单参数的值。

在这两种情况下，如果参数不存在或没有提供默认值，可以使用`DefaultQuery`和`DefaultPostForm`方法设置默认值。

例如，`page := c.DefaultQuery("page", "0")`表示获取名为"page"的查询参数的值，如果该参数不存在，则使用默认值"0"。

而`password := c.DefaultPostForm("password", "000000")`表示获取名为"password"的表单参数的值，如果该参数不存在，则使用默认值"000000"。

总结：

- `Query`方法用于获取URL查询参数的值。
- `PostForm`方法用于获取POST请求中的表单参数的值。
- `DefaultQuery`方法用于获取查询参数的值，并设置默认值。
- `DefaultPostForm`方法用于获取表单参数的值，并设置默认值。

### Map参数查询

*前面也使用了gin.H
gin.H{}是Gin框架中的一种便捷方式，用于创建一个map[string]interface{}类型的字面量，即一个键值对的集合。这个gin.H类型是Gin框架提供的一种用于构建JSON响应的数据结构。

### 重定向

- 访问顺序问题

- POST和GET有什么区别
POST和GET是HTTP协议中的两种常见请求方法，它们在以下几个方面有区别：

1. 数据传输方式：
   - GET：使用URL的查询参数进行数据传输，将数据附加在URL的末尾。数据可以在URL中被浏览器缓存、书签和历史记录中保存。
   - POST：将数据作为HTTP请求的正文发送，而不是作为URL的一部分。数据不会被缓存或保存在浏览器的历史记录中。

2. 数据传输大小限制：
   - GET：由于数据附加在URL中，URL的长度有限制。不同的浏览器和服务器对URL长度有不同的限制，通常约为2KB至8KB。
   - POST：数据作为请求的正文发送，没有固定的长度限制。可以发送大量的数据。

3. 安全性：
   - GET：参数以明文形式出现在URL中，可以被拦截和查看。适合传递非敏感信息。
   - POST：参数作为请求的正文发送，不会以明文形式出现在URL中。适合传递敏感信息，如密码或用户数据。

4. 幂等性：
   - GET：GET请求是幂等的，即多次重复请求会产生相同的结果，不会对资源状态产生影响。
   - POST：POST请求通常用于向服务器发送数据并对资源进行修改，多次重复请求可能会导致多次修改资源。

5. 缓存：
   - GET：可以被浏览器缓存，以便提高性能。相同的GET请求可以从缓存中获取响应，而无需再次发送请求到服务器。
   - POST：默认情况下不会被浏览器缓存，每次请求都会发送到服务器。

根据具体的需求和使用场景，选择适合的请求方法是很重要的。通常情况下，GET用于获取资源、查询操作，POST用于创建、更新资源或发送数据。

- 请求方法不匹配：请确保您的请求实际上是使用POST方法发送的。如果您使用的是curl命令进行测试，请确保使用-X POST选项来指定请求方法，例如：

```arduino
curl -X POST http://localhost:9999/
```


`curl`命令对不同的请求方法有不同的写法。`curl`命令可以通过选项来指定请求方法、请求头和请求体等。

以下是一些常见的`curl`命令选项用法示例：

- GET请求：
  ```
  curl http://localhost:9999/
  ```

- POST请求：
  ```
  curl -X POST http://localhost:9999/
  ```

- 携带请求体的POST请求：
  ```
  curl -X POST -d 'param1=value1&param2=value2' http://localhost:9999/
  ```

- 设置请求头：
  ```
  curl -H "Content-Type: application/json" http://localhost:9999/
  ```

- 携带认证信息的请求：
  ```
  curl -u username:password http://localhost:9999/
  ```

这些示例只是`curl`命令的一小部分用法，您可以根据需要使用不同的选项来满足特定的请求需求。可以使用`curl --help`命令或查阅`curl`命令的文档以了解更多选项和用法。

请注意，具体的请求方法和请求参数的写法可能会根据不同的应用程序和API设计而有所变化，请根据实际情况进行调整。

### 分组路由(Grouping Routes)

在Gin框架中，可以使用分组路由（Grouping Routes）来对相关的路由进行分组和管理。通过分组路由，可以在一组共享相同前缀或中间件的路由上应用相同的配置。

## 上传文件

### 单个文件

遇到的问题

在常规的浏览器中，直接通过链接访问发送 POST 请求是不支持的。浏览器通常只能发起 GET 请求或通过表单提交进行 POST 请求。

要测试和发送 POST 请求，您可以使用以下方法之一：

1. 使用 `curl` 命令：通过命令行使用 `curl` 命令发送 POST 请求。示例如下：
   ```bash
   curl -X POST -F "file=@/path/to/your/file.jpg" http://localhost:8080/upload1
   ```
   请将 `/path/to/your/file.jpg` 替换为您要上传的文件的实际路径，并确保使用的是 POST 请求。

2. 使用表单提交：在 HTML 页面中创建一个表单，并将表单的 `action` 属性设置为 `/upload1`，方法设置为 `POST`。然后在表单中添加一个文件输入字段，并提交表单。这样浏览器会自动以 POST 请求的方式发送数据到指定的 URL。

请注意，直接通过浏览器地址栏访问发送 POST 请求是不支持的。您需要使用其他工具或编程方式来发送 POST 请求，如 `curl` 命令、Postman 等。
### 多个文件

## HTML模板(Template)

## 中间件(Middleware)

## 热加载调试 Hot Reload