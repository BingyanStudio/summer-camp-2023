package main

import (
	// "./encryption"
	"fmt"
	"myproject/encryption"
	"strconv"

	// "log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model     //id、created_at、updated_at、deleted_at
	Password       string
	NickName       string
	PhoneNumber    string
	EmailAddress   string
	WhetherDeleted string //可以用gorm.Model里的deleted_at替代

}

type Administer struct {
	gorm.Model
	AdminKey string
}

// 判断指定表tableName的字段column是否存在值value
func Exists(db *gorm.DB, tableName string, column string, value interface{}) bool {

	// 使用map进行动态查询
	condition := map[string]interface{}{
		column: value,
	}

	var count int64
	db.Table(tableName).Where(condition).Count(&count)
	fmt.Println("exist number: ", count)
	return count > 0 //bool count == 0 false-->不存在；count > 0 true-->已存在
}

/* //调用示例
if Exists("users", "name", "张三") {
  // 存在
} else {
  // 不存在
}
*/

func checkerr(err error) {
	if err != nil {
		fmt.Println("an err occurred: ", err)
		panic(err)
	}
}

func main() {

	//连接数据库
	db, err := gorm.Open("mysql", "root:mysql002003004@tcp(127.0.0.1:3306)/project1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//模型与数据表对应
	db.AutoMigrate(&User{})       //用空结构体指针
	db.AutoMigrate(&Administer{}) //用空结构体指针

	r := gin.Default()
	//创建路由实例

	//1.用户注册(创建)
	r.POST("/register", func(c *gin.Context) {
		var user User
		err2 := c.ShouldBind(&user)
		if err2 != nil {
			panic(err2)
		}
		//id自动递增，作为主键，是登录凭证，要返回给用户
		user.Password = c.PostForm("Password")
		user.Password = encryption.Myencrypt(user.Password)
		user.NickName = c.PostForm("Nickname")
		user.PhoneNumber = c.PostForm("phone")
		user.EmailAddress = c.PostForm("email")
		user.WhetherDeleted = "0"
		//WhtherDeleted设置默认值为0，不返回给用户
		//*Phone和Email默认值调为空
		//fmt.Println("1234")
		fmt.Println("the register information: ", user.Password, user.NickName, user.PhoneNumber, user.EmailAddress, user.WhetherDeleted, "\nall information is shown")

		ph_exist := Exists(db, "users", "phone_number", user.PhoneNumber) //bug : 调用函数的时候已经创建了数据
		// ph_exist1 :=Exists(db,"users","phone_number","12222")
		// fmt.Println(ph_exist1)
		fmt.Println("whether the phone is used: ", ph_exist)
		if ph_exist == true {
			// err1 := fmt.Errorf("y")
			// panic(err1)
			fmt.Println("The phone is already registered.Change a phone number or log in directly.")
			return
		} else {
			if user.Password != "" && user.PhoneNumber != "" {
				db.Create(&user)
				fmt.Println("oh created")
			}
		}
		//              if ph_exist != 0 {
		//   err := fmt.Errorf("phone already exists")
		//   panic(err)
		// }
		//**或邮箱注册
	})

	//2.用户登录(查询)
	r.POST("/login", func(c *gin.Context) {
		var user User
		user.PhoneNumber = c.Query("Phone")
		user.Password = c.Query("Password")
		user.Password = encryption.Myencrypt(user.Password)
		fmt.Println("log in information : ", user.PhoneNumber, user.Password)
		var being_logged User

		fmt.Println("user password: ", user.Password, "user Nickname: ", user.NickName, "right password: ", being_logged.Password, "right nickname: ", being_logged.NickName)

		//扫描
		db.Table("users").Where("phone_number = ?", user.PhoneNumber).First(&being_logged)

		err3 := db.Table("users").Where("phone_number = ?", user.PhoneNumber).First(&being_logged).Error

		checkerr(err3)

		fmt.Println("user password: ", user.Password, "user Nickname: ", user.NickName, "right password: ", being_logged.Password, "right nickname: ", being_logged.NickName)
		fmt.Println("the user who is being logged : ", being_logged.PhoneNumber, being_logged.Password)
		fmt.Println("user password: ", user.Password, "right password: ", being_logged.Password)
		switch {
		case user.Password == being_logged.Password:
			fmt.Println("log in successfully")
			//登录成功标识......
			//......
		case being_logged.PhoneNumber != "":
			fmt.Println("The phone and the password is not matched.")
		default:
			fmt.Println("The phone number is not registered.Please register first.")
		}
	})
	//3.用户修改个人信息(更新)
	r.POST("/update", func(c *gin.Context) {
		//用户鉴权
		//用户在登录状态下可以修改自己的数据

	})

	//4.管理员登录(查询)
	r.POST("/access", func(c *gin.Context) {
		var administer Administer

		id, err4 := strconv.Atoi(c.PostForm("ID")) //string转int
		checkerr(err4)
		administer.ID = uint(id) //int转uint
		administer.AdminKey = c.PostForm("Key")
		administer.AdminKey = encryption.Myencrypt(administer.AdminKey)

		fmt.Println("the administer ID : ", administer.ID, "the adminkey : ", administer.AdminKey)

		var accessing Administer
		db.Table("administers").Where("id = ?", administer.ID).First(&accessing)

		err5 := db.Table("administers").Where("id = ?", administer.ID).First(&accessing).Error
		checkerr(err5)

		switch {
		case administer.AdminKey == accessing.AdminKey:
			fmt.Println("Get access.")
			//管理员身份状态
			//......
		default:
			fmt.Println("Chech your id and key.")
		}

	})

	//5.管理员获取用户信息(查询)
	r.POST("/checkall", func(c *gin.Context) {

		//管理员鉴权......
		//鉴权通过......
		//var results []Result
		//db.Table("users").Select("name, age").Where("id > ?", 0).Scan(&results)
		var results []User
		db.Find(&results)
		err6 := db.Find(&results).Error
		checkerr(err6)
		fmt.Println(results)
	})

	r.POST("/checkone", func(c *gin.Context) {

		//管理员鉴权......
		//鉴权通过......
		//var results []Result
		//db.Table("users").Select("name, age").Where("id > ?", 0).Scan(&results)
		var result User
		result.PhoneNumber = c.PostForm("phone")
		fmt.Println("want to check :", result.PhoneNumber)
		db.Table("users").Where("phone_number = ?", result.PhoneNumber).First(&result)

		err7 := db.Table("users").Where("phone_number = ?", result.PhoneNumber).First(&result).Error

		checkerr(err7)
		fmt.Println(result)

	})

	//6.管理员删除用户(删)
	r.POST("/delete", func(c *gin.Context) {

		//管理员鉴权......
		//鉴权通过......
		//var results []Result
		//db.Table("users").Select("name, age").Where("id > ?", 0).Scan(&results)
		var result User
		result.PhoneNumber = c.PostForm("phone")
		fmt.Println("want to check :", result.PhoneNumber)
		db.Table("users").Where("phone_number = ?", result.PhoneNumber).First(&result)

		err9 := db.Table("users").Where("phone_number = ?", result.PhoneNumber).First(&result).Error

		checkerr(err9)
		fmt.Println(result)
		//软删除
		db.Delete(&result)

		err10 := db.Delete(&result).Error
		checkerr(err10)
		fmt.Println("soft delete is done")

		// var result1 User

		// db.Table("users").Select("deleted_at").Where("phone_number = ?", result.PhoneNumber).First(&result1)

		// fmt.Println("the deleted status:",result1)
		// err8:=db.Table("users").Select("deleted_at").Where("phone = ?", result.PhoneNumber).First(&result1).Error
		// checkerr(err8)

	})

	//监听
	r.Run()

}
