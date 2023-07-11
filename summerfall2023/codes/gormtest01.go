package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//结构体-->数据表（gorm）
type UserInfo struct{
	ID uint
	Name string
	Gender string
	Hobby string
}

func main(){
//连接数据库
db,err := gorm.Open("mysql","root:mysql002003004@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
if err != nil{
	panic(err)
}
defer db.Close()//直到函数返回前执行

//创建表 自动迁移（把结构体和数据表对应）
db.AutoMigrate(&UserInfo{})

//创建数据行
u1 := UserInfo{ID:1,Name: "xixi",Gender: "male",Hobby: "smim"}
db.Create(u1)
//查询
var u UserInfo
db.First(&u)//传指针 ,查询第一行数据
fmt.Printf("u:%#v\n",u)
//更新
db.Model(&u).Update("hobby","pingpong")
fmt.Printf("u%#v\n",u)
//删除
db.Delete(&u)

fmt.Printf("u%#v\n",u)
}