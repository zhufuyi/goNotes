package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 无论什么请求方法，都把设置好的cookie发送给客户端
	cookie := &http.Cookie{
		Name:   "usr",
		Value:  "admin",
		Path:   "/",
		MaxAge: 30, // 30秒过期
	}
	http.SetCookie(w, cookie) // 发送cookie给客户端

	switch r.Method {
	case "GET":
		// get的请求，把解析网页并把网页内容发送给客户端
		tl, err := template.ParseFiles("./login.html")
		if err != nil {
			http.NotFound(w, r)
		} else {
			tl.Execute(w, nil)
		}
	case "POST":
		// 读取客户端发过来的cookie并比较处理
		cookie, err := r.Cookie("usr")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if cookie.Name != "usr" || cookie.Value != "admin" {
			w.Write([]byte("\ncookie名或值不正确！\n"))
			return
		}

		// 解析表单并处理
		err = r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if r.Form.Get("usr") == "admin" && r.Form.Get("pwd") == "123456" {
			w.Write([]byte("\n登录成功！\n"))
		} else {
			w.Write([]byte("\n用户名或密码错误！\n"))
		}
	default:
		w.Write([]byte("\n提示：只允许GET或POST的请求方法。\n"))
	}
}

func main() {
	http.HandleFunc("/login", Login)
	fmt.Println("server localhost:8080 start up......\nonly use url = http://localhost:8080/login")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
