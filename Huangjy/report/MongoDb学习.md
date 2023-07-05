## mongodb的部署

在虚拟机上部署了Docker，并创建MongoDb容器
使用`sudo docker run -itd -p 27017:27017 mongo --name mongodb --auth`创建容器映射27017端口并开启身份验证，然后键入`sudo docker exec -it mongodb mongosh admin`进入容器，mongodb使用JavaScript Shell进行交互，创建admin用户
```JavaScript
db.createUser({ 
    user:'admin',
    pwd:'0000',
    roles:[ { 
        role:'userAdminAnyDatabase', 
        db: 'admin'
        },
        "readWriteAnyDatabase"
        ]
    }
);
```
创建一个密码0000的admin用户，使用`db.auth("admin", "0000")`进行身份验证

## mongodb的连接

一般的，使用`mongodb://<name>:<password>@<address>:<port>`作为uri连接db，具体在go中
```Go
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:0000@mongodb:27017"))
if err != nil {
    panic(err)
}
db := client.Database("test") // 获取数据库
```

## mongodb的使用

使用mongo shell连接数据库，则有以下一些常用的操作
```JavaScript
db.auth("<user_name>", "<password>")     	// 进行身份验证
db                                  		// 显示当前数据库

show dbs									// 打印所有数据库  
show collections/tables						// 打印当前数据库下的表
show users									// 打印所有用户

use <database>								// 转到目标数据库，如果没有就新建一个
db.dropDatabase()							// 删除当前所在的数据库
db.createCollection("<collection_name>", options)	// 创建新表
db.<collection_name>.drop()					// 删除这个表
db.<collection_name>.insertOne(document)	// 在表中插入一个文档，如果这个表不存在就创建一个
db.<collection_name>.update(query, document)// 更新表中的文档
db.<collection_name>.remove(query, document)// 删除表中的文档，如果只删一个就加参数{justOne: true}
db.<collection_name>.find(query)			// 查询表中的文档

// db.<collection_name>下find()和remove({}) 分别表示查询所有和删除所有
```

有一些条件操作符用于比较表达式并从db中获取数据
+  ($ \gt $)  大于 ---- `$gt`
+  ($ \lt $)  小于 ---- `$lt`
+  ($ \geq $)  大于等于 ---- `$gte`
+  ($ \leq $)  小于等于 ---- `$lte`

例如，从`students`表中筛选 $ 80 \lt score \leq 100 $ 的学生
```JavaScript
db.students.find({ score : { $gt 80, $lte 100 }})
```

## mongo-go-driver

连接mongodb后，可以通过一系列操作对数据库进行增 删 改 查操作。例如，现在获取了一个集合
``` Go
var col *mongo.Collection = db.Collection("test")
```
定义一个S结构体如下
``` Go
type S struct {
    Name string
    ID uint32
}
//已经有了几个结构，准备插入到数据库中
var (
    p1 S = S{"person01", "1001"}
    p2 S = S{"person02", "1002"}
    p3 S = S{"person03", "1003"}
)
```
1. 插入文档
    使用`<collection>.InsertOne()`插入单个文档，使用`InsertMany()`插入多个

    ``` GO
    _, err := col.InsertOne(context.TODO(), p1)
    if err != nil {
        log.Fatal(err)
        return err
    }
    ps := []interface{}{p2, p3}
    _, err = col.InsertMany(context.TODO(), ps)
    // ... 错误处理
    ```
2. 更新一个文档
    使用`bson.D`来描述筛选和更新内容
    ``` Go
    filter := bson.D{{"name", "person1"}}
    // filter := bson.D{{}}将会匹配所有的内容
    // 这类似于db.col.find()
    update := bson.D{
	    {"$set", bson.D{
		    {"ID", 1004},
	    }},
    }

    _, err := col.UpdateOne(context.TODO(), filter, update)
    // ... 错误处理
    ```
3. 查询文档
    一样使用一个`bson.D`来表述筛选内容
    ``` Go
    var s S
    err := col.FindOne(context.TODO(), fitler).Decode(&s)
    // ... 错误处理
    ```
4. 删除文档
    使用`<collection>.DeleteOne()`删除一个文档，`DeleteMany`删除所有匹配的文档，这意味着当所有的匹配，即`filter == bson.D{{}}`时，会删除所有文档。
    ``` Go
    err := col.DeleteOne(context.TODO(), filter)
    err = col.DeleteMany(context.TODO(), filter)
    // ... 错误处理
    ```

### bson
MongoDB中的JSON文档存储在名为BSON(二进制编码的JSON)的二进制表示中。连接MongoDB的Go驱动程序中有两大类型表示BSON数据：D和Raw。类型D家族被用来简洁地构建使用本地Go类型的BSON对象。这对于构造传递给MongoDB的命令特别有用。D家族包括四类:
+ D：一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongoDB命令。
+ M：一张无序的map。它和D是一样的，只是它不保持顺序。
+ A：一个BSON数组。
+ E：D里面的一个元素。

使用` import "go.mongodb.org/mongo-driver/bson" `来使用bson类型。