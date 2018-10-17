// 简单的web服务器的三种启动方式
package main

import (
	"fmt"
	"net/http"
)

// ----------------- 1. 默认路由方式 ---------------
func helloHandle1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world! 1. 默认路由方式"))
}

func byeHandle1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye Bye! [1]"))
}

func defaultMuxServer() {
	http.HandleFunc("/", helloHandle1)
	http.HandleFunc("/bye", byeHandle1)
	fmt.Println("localhost:8001 start up......\n")
	fmt.Println(http.ListenAndServe(":8001", nil))
}

// ----------------- 2. 自定义路由方式 --------------
type myMux struct{}

func (mux *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("Hello world! 2. 自定义路由方式"))
	case "/bye":
		w.Write([]byte("Bye Bye! [2]"))
	default:
		http.NotFound(w, r)
	}
}

func defineMuxServer() {
	fmt.Println("localhost:8002 start up......\n")
	fmt.Println(http.ListenAndServe(":8002", new(myMux)))
}

// ----------------- 3. 路由映射 --------------
var Mux = make(map[string]func(http.ResponseWriter, *http.Request))

type myHandle struct{}

func (mh *myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f, ok := Mux[r.URL.String()]; ok {
		f(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func helloHandle3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world! 3. 路由映射"))
}

func byeHandle3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye!"))
}

func mapMuxServer() {
	Mux["/"] = helloHandle3
	Mux["/bye"] = byeHandle3

	server := &http.Server{
		Addr:    ":8003",
		Handler: &myHandle{},
	}
	fmt.Println("localhost:8003 start up......\n")
	fmt.Println(server.ListenAndServe())
}

func main() {
	go defaultMuxServer()
	go defineMuxServer()
	mapMuxServer()
}
