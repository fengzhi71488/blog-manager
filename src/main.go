package main

import (
	"net/http"
	_ "router"
	"fmt"
)

func main() {

	fmt.Println("------服务器启动")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("------服务器启动出错：",err)
	}
}
