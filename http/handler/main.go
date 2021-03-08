package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{} // 处理器和处理器函数

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world !")
}

func main() { // Server服务器&hander处理器进行了绑定：用同一个处理器处理所有的请求
	handler := MyHandler{}
	server := http.Server{ // 服务器
		Addr:    "127.0.0.1:8080",
		Handler: &handler, // 处理器绑定服务器http://127.0.0.1:8080/anything/at/all
	}
	server.ListenAndServe()
}
