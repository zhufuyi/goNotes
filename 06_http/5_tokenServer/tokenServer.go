// 通过随机token令牌，防止页面重复提交
package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

var token string
var submitCnt, handleSum int // 分别是提交次数和处理次数

// 用时间戳生成token
func generateToken() string {
	digest := md5.New()
	_, err := io.WriteString(digest, fmt.Sprintf("%d", time.Now().Unix()))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", digest.Sum(nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		submitCnt = 0
		handleSum = 0

		// 随机生成token
		token = generateToken()
		if token == "" {
			fmt.Println("生成token为空。")
			return
		}
		fmt.Println("token =", token)

		// 解析模板并发给客户端
		tl, err := template.ParseFiles("./login.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		err = tl.Execute(w, token)
		if err != nil {
			http.NotFound(w, r)
		}

	case "POST":
		time.Sleep(time.Second) // 模拟网络繁忙情况，延时1秒

		// 解析表单
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		submitCnt++
		fmt.Println("submitCnt =", submitCnt, " handleSum =", handleSum)

		// 读取token令牌并验证tonken是否合法
		postToken := r.Form.Get("token")
		if postToken == "" {
			w.Write([]byte("警告：token值为空！"))
			return
		}
		if token != postToken {
			w.Write([]byte("警告：tonken不合法"))
			return
		}

		token = "" // 清空token,防止处理重复的提交表单

		handleSum++
		fmt.Println("submitCnt =", submitCnt, " handleSum =", handleSum)

		if r.Form.Get("usr") == "admin" && r.Form.Get("pwd") == "123456" {
			w.Write([]byte("\n登录成功！\n"))
		} else {
			w.Write([]byte("\n用户名或密码错误！\n"))
		}
	default:
		w.Write([]byte("\n警告：只允许GET或POST的请求方法。\n"))
	}
}

func main() {
	http.HandleFunc("/login", Login)
	fmt.Println("server localhost:8080 start up......\nonly use url = http://localhost:8080/login")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
