package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)
import _ "github.com/go-sql-driver/mysql"

const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "mydb"
)

var DB *sql.DB
var books []*Book

func main() {
	resolveServerReq()
}

func resolveServerReq() {
	queryDataFromDB()
	data := database{item: books}
	http.HandleFunc("/list", data.haha)
	http.ListenAndServe("localhost:8000", nil)
}

func queryDataFromDB() {
	if connectDB() {
		return
	}
	rows, err := DB.Query("SELECT * FROM person ")
	if err != nil {
		fmt.Println("查询出错了")
		return
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.name)
		if err != nil {
			fmt.Printf("query erro %v", err)
			return
		}
		books = append(books, &book)
	}
}

func connectDB() bool {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return true
	}
	fmt.Println("connnect success")
	return false
}

type Book struct {
	name  string
	price float32
}

type database struct {
	item []*Book
}

func (db *database) haha(writer http.ResponseWriter, request *http.Request) {
	for _, it := range db.item {
		fmt.Fprintf(writer, "id = %s\n", it.name)
	}
}
