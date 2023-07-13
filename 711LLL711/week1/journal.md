# week1 

## 使用GORM操作mysql数据库  
~~~go
//声明模型 
type User struct {
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time 
  //ID、CreatedAt、UpdatedAt、DeletedAt是gorm.Model的默认字段
}

//创建单条记录
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 通过数据的指针来创建

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数

//创建多条记录--指针数组切片
users := []*User{
    User{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
    User{Name: "Jackson", Age: 19, Birthday: time.Now()},
}
result := db.Create(users) // pass a slice to insert multiple row 

//查询
result := db.First(&user)//first()返回第一条记录，last()返回最后一条记录，take()返回一条记录，where()返回所有匹配的记录
result.RowsAffected // 返回找到的记录数
result.Error        // returns error or nil

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)
//update
db.First(&user)
user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
//更新单个列
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
//delete
db.Where("name = ?", "jinzhu").Delete(&email) 
~~~

## 加密算法
1. 对称加密--加密和解密密钥相同   
2. 非对称加密--加密和解密密钥不同    
3. 哈希算法--不可逆，不可解密，单向函数，可以用来加密存在数据库中的密码  
    ~~~go
    password := "user_password" // 用户输入的密码

    // 生成哈希密码
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    //DefaultCost是哈希密码的cost factor，值越高，计算成本越高，但是也更安全
    if err != nil {
        // 处理错误
    }
    // 将哈希密码存储到数据库中
    // ...
    // 检查密码是否匹配
    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))  
    if err != nil {
	// 密码不匹配
    } else {
        // 密码匹配，登录成功
    }
    ~~~