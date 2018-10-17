// 搭建一个简单的网页服务器，处理浏览器的请求和提交表单。
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// "/"路径对应的处理函数
func homeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tl, err := template.ParseFiles("./home.html")
		if err != nil {
			http.NotFound(w, r)
		} else {
			tl.Execute(w, nil)
		}
	} else {
		w.Write([]byte("提示：请求方法应该为GET。"))
	}
}

// "/login"路径对应的处理函数
func loginHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tl, err := template.ParseFiles("./login.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}

		tl.Execute(w, nil)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			fmt.Println("用户名 =", r.Form["usr"], "\n密  码 =", r.Form["pwd"])

			// 处理完毕后可以设置跳转到其他网页
			http.Redirect(w, r, "/content/index", http.StatusFound)
		}
	default:
		w.Write([]byte("提示：请求方法应该为GET或POST。"))
	}
}

func contentHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tl, err := template.ParseFiles("./content/index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}

		tl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/content/index", contentHandle)
	fmt.Println("server localhost:8080 start up......\ncan use url = http://localhost:8080 or url = http://localhost:8080/login")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
