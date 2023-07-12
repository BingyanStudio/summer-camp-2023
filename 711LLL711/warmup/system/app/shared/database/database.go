package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      string
	Parameter string
}

// 用于解析JSON数据的结构体
type Configdata struct {
	Database struct {
		MySQL MySQLInfo `json:"mysql"`
	} `json:"Database"`
}

var Db *sql.DB

// DSN returns the Data Source Name
func DSN(ci MySQLInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		ci.Port +
		")/" +
		ci.Name + ci.Parameter
}

func loadjson() MySQLInfo {
	// Load the configuration file
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析JSON数据到结构体
	var config Configdata
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config.Database.MySQL
}

// Connect to the database
func Connect() {
	ci := loadjson()
	var err error
	Db, err = sql.Open("mysql", DSN(ci))
	if err != nil {
		log.Println("SQL Driver Error", err)
	}
	if err = Db.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	log.Println("Connected to database")
}

// 验证是否是管理员
func Isadmin(Id string) bool {
	var role string
	err := Db.QueryRow("SELECT role FROM users WHERE id = ?", Id).Scan(&role)
	if err != nil {
		log.Println(err)
		log.Println("查询失败")
	}
	if role == "admin" {
		log.Println("管理员")
		return true
	} else {
		log.Println("非管理员")
		log.Println(Id)
		log.Println(role)
		return false
	}
}

func Checkconnection() bool {
	if err := Db.Ping(); err != nil {
		return false
	}
	return true
}
