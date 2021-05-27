package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	type Addr struct {
		city     string
		district string
	}
	type Person struct {
		Name    string
		Age     uint8
		Address Addr
	}

	personChan := make(chan Person, 1)

	person := Person{"xiaoming", 10, Addr{"shenzhen", "longgang"}}
	personChan <- person

	person.Address = Addr{"guangzhou", "huadu"}
	fmt.Printf("src person : %+v \n", person)

	newPerson := <-personChan
	fmt.Printf("new person : %+v \n", newPerson)

	//fmt.Printf(Get("http://localhost:8000/list"))
	//fmt.Printf(Get("http://localhost:8000/list"))
}

func Get(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}
