### mongodb�Ĳ���
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