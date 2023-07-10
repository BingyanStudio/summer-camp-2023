
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	type Doctor struct {
	ID      int
	Name    string
	Age     int
	Sex     int
	AddTime string
}


	db, err := sql.Open("mysql", "root:mysql002003004@tcp(127.0.0.1:3306)/testdb")//使用 sql.Open() 方法来连接到 MySQL 数据库。第一个参数是数据库驱动名称 "mysql"，第二个参数是连接字符串，指定用户名、密码、主机和端口以及要连接的数据库名称。
	if err != nil {
		log.Fatal(err)
	}
	//错误处理：在连接数据库和执行 SQL 语句的过程中，如果出现错误，会使用 log.Fatal() 方法将错误信息打印到控制台并终止程序的执行。
	defer db.Close()



	//它包含了列的定义以及主键的指定
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS doctor_tb (
			id int(11) NOT NULL AUTO_INCREMENT,
			name varchar(50) DEFAULT '' COMMENT '姓名',
			age int(11) DEFAULT '0' COMMENT '年龄',
			sex int(11) DEFAULT '0' COMMENT '性别',
			addTime datetime DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='医生表';
	`
//id 列是一个整数类型的列，长度为 11，NOT NULL 表示该列不能为空，AUTO_INCREMENT 表示该列的值将自动递增。name 列是一个最大长度为 50 的字符串类型的列，DEFAULT '' 表示如果没有提供具体值，默认值为空字符串。age 列是一个整数类型的列，长度为 11，DEFAULT '0' 表示如果没有提供具体值，默认值为 0。sex 列是一个整数类型的列，长度为 11，DEFAULT '0' 表示如果没有提供具体值，默认值为 0。
//PRIMARY KEY (id)：指定 id 列作为主键，它将唯一标识表中的每一行数据。
/*ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='医生表'：这是表的选项部分。它定义了表的存储引擎、自动递增的起始值和字符集等。

ENGINE=InnoDB：指定表的存储引擎为 InnoDB，它是 MySQL 的一种事务安全存储引擎。

AUTO_INCREMENT=1：指定自动递增的起始值为 1，这意味着每次插入新行时，id 列的值将从 1 开始递增。

DEFAULT CHARSET=utf8：指定表的字符集为 utf8，这将影响表中字符串列的存储和排序规则。

COMMENT='医生表'：为表提供一个注释，用于描述表的用途或特点。*/


	_, err = db.Exec(createTableSQL)
	//使用了 db.Exec() 方法执行 SQL 语句来创建名为 "doctor_tb" 的表。CREATE TABLE 语句定义了表的结构，包括字段名称、数据类型、约束等。
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")
	

	//-------4、新增数据--------  

/*	
result, err := db.Exec("insert into doctor_tb(name,age,sex,addTime) values(?,?,?,Now())", "全医生", 40, 1)
if err != nil {  
   fmt.Println("新增数据错误", err)  
   return  
}  
newID, _ := result.LastInsertId() //新增数据的ID  
i, _ := result.RowsAffected()     //受影响行数  
fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
*/


/*
//-------2、查询单条数据--------  
//定义接收数据的结构  
var doc Doctor  
//执行单条查询  
rows := db.QueryRow("select * from doctor_tb where id = ?", 2) 
 
err = rows.Scan(&doc.ID, &doc.Name, &doc.Age, &doc.Sex, &doc.AddTime) // 用于把读取的数据赋值到Doctor对象的属性上，要注意字段顺序。
//err 是一个错误类型的值 (error)。在调用 rows.Scan() 方法时，如果出现错误，会将该错误赋值给 err 变量。如果没有错误发生，err 将为 nil。
//使用 rows.Scan() 方法将查询结果的列值扫描并存储到 doc 变量对应的字段中。在这里，我们传入了 &doc.ID、&doc.Name、&doc.Age、&doc.Sex 和 &doc.AddTime，将查询结果的各个列值扫描到对应的 Doctor 结构体字段中。
fmt.Println("单条数据结果：", doc,rows,err)
*/

//-------3、查询数据列表--------  
rows2, err := db.Query("select * from doctor_tb where age > ?", 30)  //查询
if err != nil {  
   fmt.Println("多条数据查询错误", err)  
   return  
}  
//定义对象数组,用于接收数据  
var docList []Doctor  
for rows2.Next() {  
    var doc2 Doctor  
    rows2.Scan(&doc2.ID, &doc2.Name, &doc2.Age, &doc2.Sex, &doc2.AddTime)  
    //加入数组  
    docList = append(docList, doc2)  
}  
fmt.Println("多条数据查询结果", docList)
}







