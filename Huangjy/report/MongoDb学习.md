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