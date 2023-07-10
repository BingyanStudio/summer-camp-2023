// 用sql实现基本的CRUD
package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 定义记录的结构体方便接受数据
type Student struct {
	Id   int
	Name string
	Age  int
}

// 定义指针
var db *sql.DB

// 查询一行数据
func queryonedemo() (Student, error) {
	var student Student
	err := db.QueryRow("SELECT * FROM students WHERE id = ?", 1).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return student, err
	}
	return student, nil
}

// 查询多行数据
func queryalldemo() ([]Student, error) {
	sqlStr := "SELECT * FROM students"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

// insert
func insertdemo() {
	//prepare预处理用于之后的多次执行，减少sql语句的编译次数
	stmt, err := db.Prepare("insert into students(name,age) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("klutz", 22)
	stmt.Close() //关闭预处理
	if err != nil {
		log.Fatal(err)
	}

	/*不用prepare也可以
	res, err := db.Exec("insert into students(name,age) values(?,?)", "klutz", 22)
	if err != nil {
		log.Fatal(err)
	}
	*/

	id, err := res.LastInsertId() //lastinsertid返回插入数据的id
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

// update
func updatedemo() {
	stmt, err := db.Prepare("update students set age = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(100, 2)
	stmt.Close() //关闭预处理
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected() //rowsaffected返回影响的行数
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(affect)
}

// delete
func deletedemo() {
	stmt, err := db.Prepare("delete from students where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1)
	stmt.Close() //关闭预处理
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(affect)
}

func main() {
	var err error

	//连接数据库  用户名:密码@tcp(ip:端口)/数据库名称
	db, err = sql.Open("mysql", "ljx:123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	//defer保证数据库连接的关闭
	defer db.Close()
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//判断连接是否成功，注意sql.Open()函数只是验证参数格式是否正确，不会验证账号密码是否正确
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	//查询一行数据
	student666, err := queryonedemo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student666)

	//查询多行数据
	students, err := queryalldemo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)

	//插入数据
	insertdemo()

	//更新数据
	updatedemo()

	//删除数据
	deletedemo()
}
