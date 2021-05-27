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

func main() {
	connectDB()
	http.HandleFunc("/list", handSqlData)
	http.ListenAndServe("localhost:8000", nil)
}
func connectDB() bool {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		return true
	}
	return false
}

type Book struct {
	name  string
	price string
}

func handSqlData(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("execuhand")
	rows, err := DB.Query("SELECT * FROM person ")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.name, &book.price)
		fmt.Fprintf(writer, "id = %s name = %s", book.name, book.price)
		if err != nil {
			return
		}
	}
}
