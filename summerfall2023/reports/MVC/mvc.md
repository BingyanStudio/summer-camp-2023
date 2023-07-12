# mvc

## 概览

我们将创建一个作为模型的 Student 对象。StudentView 是一个把学生详细信息输出到控制台的视图类，StudentController 是负责存储数据到 Student 对象中的控制器类，并相应地更新视图 StudentView。


MVC 将代码分层，易于调试和管理

代码的组织结构：  

- 易于添加新功能
- 易于修复漏洞
- 易于其他各种操作

## Model View and Controller

### Model ——the most important

包括数据库交互

- a good model:your models package doesn’t know anything about the rest of your application.仅限于仅与数据库实体交互

- what logic should be found in Model  


    - Code that defines the User type as it is stored in the database.
定义存储在数据库中的 User 类型的代码。
    - Logic that will check a user’s email address when a new user is being created to verify that it isn’t already taken.
创建新用户时检查用户的电子邮件地址以验证该地址是否尚未被占用的逻辑。
    - If your particular SQL variant is case sensitive, code that will convert a user’s email address to lowercase before it is inserted into the DB.
如果您的特定 SQL 变体区分大小写，则在将用户的电子邮件地址插入数据库之前将其转换为小写的代码。
    - Code for inserting users into the users table in the DB.
用于将用户插入到数据库中的用户表中的代码。
    - Code that retrieves relational data; eg a User along with all of their Reviews.
检索关系数据的代码；例如 User 及其所有 Reviews 。

- what logic should not be found in Model  

    - HTML responses.
    - HTTP status codes.

### Veiw

- what is in View:

    - A glorified wrapper around the html/template.Template package html/template.Template 包周围的美化包装
    - A wrapper (or direct use of) the encoding/json package
encoding/json 包的包装器（或直接使用）

- what should we do:  
The main way I tend to customize this is by adding a bunch of template.FuncMap functions to the template BEFORE parsing it (so it parses correctly), and then I’ll override them using request-specific data if I need to. Eg:
我倾向于自定义它的主要方法是在解析模板之前向模板添加一堆 template.FuncMap 函数（因此它可以正确解析），然后如果需要，我将使用请求特定的数据覆盖它们到。

MVC 中的控制器与空中交通管制员非常相似。它们不会直接负责写入数据库或创建 HTML 响应，而是会将传入的数据定向到适当的模型，视图和其他可用于完成请求工作的程序包。

与视图类似，控制器不应包含太多业务逻辑。相反，它们应该只解析数据并将其交付给其他函数，类型等进行处理。

————————————————
原文作者：Summer
转自链接：https://learnku.com/go/t/48112
版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。

### Controller

- what does controller do  
创建 http.Handler 来解析传入数据，调用 models 包提供的 UserStore 或 CourseStore 等方法，然后最终通过视图呈现结果（或在某些情况下将用户重定向到适当的页面）。

## 扁平结构

    Using that same line of reasoning, it is possible to use a flat structure and MVC at the same time.
    使用相同的推理思路，可以同时使用扁平结构和 MVC。

    在扁平结构应用程序中，我们经常将代码分解为数据库、处理程序和渲染层。所有这三个都很好地映射到模型、控制器和视图。


建议在更了解APP功能时将flat structure 改为 MVC structure BY 创建有明显好处的包

Another big reason to consider a flat structure is that it is much easier for your structure to evolve as your application grows in complexity. When it becomes apparent that you could benefit from breaking code into a separate package, all you often need to do is move a few source files into a subdirectory, change their package, and update any reference to use the new package prefix. Eg if we had SqlUser and decided we would benefit from having a separate sql package to handle all our database related logic, we would update any references to now use sql.User after moving the type to the new package. I have found that structures like MVC are a bit more challenging to refactor, albeit not impossible or as hard as it might be in other programming languages.
考虑扁平结构的另一个重要原因是，随着应用程序复杂性的增加，您的结构更容易发展。当您明显可以从将代码分解到单独的包中受益时，您通常需要做的就是将一些源文件移动到子目录中，更改它们的包，并更新任何引用以使用新的包前缀。例如，如果我们有 SqlUser 并决定我们将受益于拥有一个单独的 sql 包来处理所有与数据库相关的逻辑，我们将更新所有引用以现在使用 sql.User 将类型移至新包后。我发现像 MVC 这样的结构重构起来更具挑战性，尽管这并非不可能，也不像其他编程语言那样困难。

A flat structure can be especially useful for beginners who are often too quick to create packages. I can’t really say why this phenomenon happens, but newcomers to Go love to create tons of packages and this almost always leads to stuttering (user.User), cyclical dependencies, or some other issue.
扁平结构对于初学者来说尤其有用，因为他们创建包的速度往往太快。我真的不能说为什么会发生这种现象，但是 Go 的新手喜欢创建大量的包，这几乎总是会导致口吃 ( user.User )、循环依赖或其他一些问题。

In the next article on MVC we will explore how this phenomenon of creating too many packages can make MVC seem impossible in Go, despite that being far from the truth.
在下一篇关于 MVC 的文章中，我们将探讨这种创建过多包的现象如何使 MVC 在 Go 中显得不可能，尽管这与事实相去甚远。

By putting off decisions to create new packages until our application grows a bit and we understand it better, budding Gophers are far less likely to make this mistake.
通过推迟创建新包的决定，直到我们的应用程序增长一点并且我们更好地理解它，崭露头角的地鼠犯这个错误的可能性要小得多。

This is also why many people will encourage developers to avoid breaking their code into microservices too early - you often don’t have enough knowledge to really know what should and shouldn’t be split into a microservice early on and preemptive microservicing (I kinda hope that becomes a saying) will just lead to more work in the future.
这也是为什么许多人会鼓励开发人员避免过早地将代码分解为微服务 - 您通常没有足够的知识来真正知道什么应该和不应该尽早拆分为微服务以及抢占式微服务（我有点希望这成为了一种说法）只会导致未来更多的工作。


## some suggestion

- When we use MCL,we should pay attention to those problems:  

1. it is okay to create additional packages in your codebase.
我们可以使用额外的包

2. Don’t break things up too much  
不要分解得过于细致

3. cyclical dependencies.
循环依赖问题  
如果您发现自己存在循环依赖关系，那么您很可能将事情分解得太多，或者将类型/接口/函数放在错误的位置。

- YOU might do :  

1. You will need to define resources more than once
您将需要多次定义资源

2. Don't use Globals  
全局变量在 goroutine 中导致竞争和混乱  
insert when you use it  
kill it when you finish using it  

3. Don’t embed DB connections to make relational queries possible  
不要嵌入数据库连接来进行关系查询
