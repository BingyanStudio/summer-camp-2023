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