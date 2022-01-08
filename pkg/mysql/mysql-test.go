package mysql

import (
	"Database/configs"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// MysqlDB 定义全局mysql客户端
var db *sqlx.DB

// 初始化mysql数据库连接
func init() {
	mysqlHost := configs.MysqlHost
	database, err := sqlx.Open("mysql", mysqlHost)
	if err != nil {
		log.Printf("创建数据库连接%s失败,错误信息 %e\n", mysqlHost, err)
		return
	}
	db = database
	// 关闭mysql数据库链接
	//defer db.Close()
}

// InsertOperation 数据库连接初始化
func InsertOperation() {
	result, err := db.Exec("insert into person(name, sex, email) values (?,?,?)", "stu001", "man", "asdfasdf@163.com")
	if err != nil {
		log.Println("执行sql插入失败,错误信息", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("执行失败,错误信息", err)
		return
	}
	fmt.Println("向数据库表person中插入数据成功,数据插入ID号", id)
}

// SelectOperation mysql查询操作
func SelectOperation() {
	// 定义Person结构体
	type Person struct {
		Id    int    `db:"id""`
		Name  string `db:"name"`
		Sex   string `db:"sex"`
		Email string `db:"email"`
	}
	var person []Person
	err := db.Select(&person, "select id, name, sex, email from person")
	if err != nil {
		log.Println("执行查询失败, 错误信息", err)
		return
	}
	fmt.Println("执行查询成功", person)
}
