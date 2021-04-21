package main

import (
	"database/sql"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
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
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
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
		return
	}
	fmt.Println("connnect success")

	DB.Query("select * from person")
	SelectUserById()
	//var items = []*Book{{"Live", 23}, {"Earth", 30}}
	//data := database{item: items}
	////startServer("localhost:8000",&data)
	//
	//http.HandleFunc("/list", data.list)
	//http.HandleFunc("/price", data.price)
	//http.ListenAndServe("localhost:8000", nil)
	//
	//fmt.Println("test rebase++")
	//fmt.Println("test rebase ++++")

	//listener, err := net.Listen("tcp", "localhost:8000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Print(err) // e.g., connection aborted
	//		continue
	//	}
	//	handleConn(conn) // handle one connection at a time
	//}

}

func SelectUserById() {

	rows, err := DB.Query("SELECT * FROM person ")

	if err != nil {
		fmt.Println("查询出错了")
	}

	//获取列名
	columns, _ := rows.Columns()

	//定义一个切片,长度是字段的个数,切片里面的元素类型是sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	//定义一个切片,元素类型是interface{} 接口
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		//把sql.RawBytes类型的地址存进去了
		scanArgs[i] = &values[i]
	}
	//获取字段值
	var result []map[string]string
	for rows.Next() {
		res := make(map[string]string)
		rows.Scan(scanArgs...)
		for i, col := range values {
			res[columns[i]] = string(col)
		}
		result = append(result, res)
	}

	//遍历结果
	for _, r := range result {
		for k, v := range r {
			fmt.Printf("%s==%s", k, v)
		}
	}
	rows.Close()
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func startServer(addr string, h http.Handler) {
	fmt.Println("2021")
	http.ListenAndServe(addr, h)
}

type Book struct {
	name  string
	price float32
}

type database struct {
	item []*Book
}

func (d *database) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, it := range d.item {
		fmt.Fprintf(writer, "%s: %s\n", it.price, it.name)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.item {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	//item := req.URL.Query().Get("item")
	//price, ok := db[item]
	//if !ok {
	//	w.WriteHeader(http.StatusNotFound) // 404
	//	fmt.Fprintf(w, "no such item: %q\n", item)
	//	return
	//}
	//fmt.Fprintf(w, "%s\n", price)
}
