## mongodb�Ĳ���

��������ϲ�����Docker��������MongoDb����
ʹ��`sudo docker run -itd -p 27017:27017 mongo --name mongodb --auth`��������ӳ��27017�˿ڲ����������֤��Ȼ�����`sudo docker exec -it mongodb mongosh admin`����������mongodbʹ��JavaScript Shell���н���������admin�û�
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
����һ������0000��admin�û���ʹ��`db.auth("admin", "0000")`���������֤

## mongodb������

һ��ģ�ʹ��`mongodb://<name>:<password>@<address>:<port>`��Ϊuri����db��������go��
```Go
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:0000@mongodb:27017"))
if err != nil {
    panic(err)
}
db := client.Database("test") // ��ȡ���ݿ�
```

## mongodb��ʹ��

ʹ��mongo shell�������ݿ⣬��������һЩ���õĲ���
```JavaScript
db.auth("<user_name>", "<password>")     	// ���������֤
db                                  		// ��ʾ��ǰ���ݿ�

show dbs									// ��ӡ�������ݿ�  
show collections/tables						// ��ӡ��ǰ���ݿ��µı�
show users									// ��ӡ�����û�

use <database>								// ת��Ŀ�����ݿ⣬���û�о��½�һ��
db.dropDatabase()							// ɾ����ǰ���ڵ����ݿ�
db.createCollection("<collection_name>", options)	// �����±�
db.<collection_name>.drop()					// ɾ�������
db.<collection_name>.insertOne(document)	// �ڱ��в���һ���ĵ��������������ھʹ���һ��
db.<collection_name>.update(query, document)// ���±��е��ĵ�
db.<collection_name>.remove(query, document)// ɾ�����е��ĵ������ֻɾһ���ͼӲ���{justOne: true}
db.<collection_name>.find(query)			// ��ѯ���е��ĵ�

// db.<collection_name>��find()��remove({}) �ֱ��ʾ��ѯ���к�ɾ������
```

��һЩ�������������ڱȽϱ��ʽ����db�л�ȡ����
+  ($ \gt $)  ���� ---- `$gt`
+  ($ \lt $)  С�� ---- `$lt`
+  ($ \geq $)  ���ڵ��� ---- `$gte`
+  ($ \leq $)  С�ڵ��� ---- `$lte`

���磬��`students`����ɸѡ $ 80 \lt score \leq 100 $ ��ѧ��
```JavaScript
db.students.find({ score : { $gt 80, $lte 100 }})
```

## mongo-go-driver

����mongodb�󣬿���ͨ��һϵ�в��������ݿ������ ɾ �� ����������磬���ڻ�ȡ��һ������
``` Go
var col *mongo.Collection = db.Collection("test")
```
����һ��S�ṹ������
``` Go
type S struct {
    Name string
    ID uint32
}
//�Ѿ����˼����ṹ��׼�����뵽���ݿ���
var (
    p1 S = S{"person01", "1001"}
    p2 S = S{"person02", "1002"}
    p3 S = S{"person03", "1003"}
)
```
1. �����ĵ�
    ʹ��`<collection>.InsertOne()`���뵥���ĵ���ʹ��`InsertMany()`������

    ``` GO
    _, err := col.InsertOne(context.TODO(), p1)
    if err != nil {
        log.Fatal(err)
        return err
    }
    ps := []interface{}{p2, p3}
    _, err = col.InsertMany(context.TODO(), ps)
    // ... ������
    ```
2. ����һ���ĵ�
    ʹ��`bson.D`������ɸѡ�͸�������
    ``` Go
    filter := bson.D{{"name", "person1"}}
    // filter := bson.D{{}}����ƥ�����е�����
    // ��������db.col.find()
    update := bson.D{
	    {"$set", bson.D{
		    {"ID", 1004},
	    }},
    }

    _, err := col.UpdateOne(context.TODO(), filter, update)
    // ... ������
    ```
3. ��ѯ�ĵ�
    һ��ʹ��һ��`bson.D`������ɸѡ����
    ``` Go
    var s S
    err := col.FindOne(context.TODO(), fitler).Decode(&s)
    // ... ������
    ```
4. ɾ���ĵ�
    ʹ��`<collection>.DeleteOne()`ɾ��һ���ĵ���`DeleteMany`ɾ������ƥ����ĵ�������ζ�ŵ����е�ƥ�䣬��`filter == bson.D{{}}`ʱ����ɾ�������ĵ���
    ``` Go
    err := col.DeleteOne(context.TODO(), filter)
    err = col.DeleteMany(context.TODO(), filter)
    // ... ������
    ```

### bson
MongoDB�е�JSON�ĵ��洢����ΪBSON(�����Ʊ����JSON)�Ķ����Ʊ�ʾ�С�����MongoDB��Go�������������������ͱ�ʾBSON���ݣ�D��Raw������D���屻�������ع���ʹ�ñ���Go���͵�BSON��������ڹ��촫�ݸ�MongoDB�������ر����á�D�����������:
+ D��һ��BSON�ĵ�����������Ӧ����˳����Ҫ�������ʹ�ã�����MongoDB���
+ M��һ�������map������D��һ���ģ�ֻ����������˳��
+ A��һ��BSON���顣
+ E��D�����һ��Ԫ�ء�

ʹ��` import "go.mongodb.org/mongo-driver/bson" `��ʹ��bson���͡�