### mongodb的部署
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