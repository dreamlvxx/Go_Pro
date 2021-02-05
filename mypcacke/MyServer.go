package main

import (
	"fmt"
	"net/http"
)

func main() {
	var items = []*Book{{"Live", 23}, {"Earth", 30}}
	data := database{item: items}
	//startServer("localhost:8000",&data)

	http.HandleFunc("/list", data.list)
	http.HandleFunc("/price", data.price)
	http.ListenAndServe("localhost:8000", nil)

}

func startServer(addr string, h http.Handler) {
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

func mm() {

}
