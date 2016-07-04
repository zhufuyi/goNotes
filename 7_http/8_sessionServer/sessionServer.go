// 通过session第三方包来简单使用session的get、set和delete方法

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/astaxie/beego/session"
)

// 始化一个全局的变量用来存储 session 控制器
var globalSessions *session.Manager

// 初始化时配置session
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"mySessionID", "enableSetCookie,omitempty": true, "gclifetime":300, "maxLifetime": 300, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 300, "providerConfig": ""}`)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	sess, err := globalSessions.SessionStart(w, r) // 如果sessionID存在，直接返回，如果sessionID不存在,会生成新的sessionID存放在cookie
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer sess.SessionRelease(w)
	v := sess.Get("username")
	fmt.Printf("%5s  sessionID = %s\n", fmt.Sprintf("%v", v), sess.SessionID())

	switch r.Method {
	case "GET":
		if v == nil { // 当第一次请求网页或网页过期时，v值为nil，把它设置为1
			sess.Set("username", int(1))
		}

		// 解析和输出网页模板
		tl, err := template.ParseFiles("./login.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tl.Execute(w, nil)

	case "POST":
		// session处理
		if v == nil {
			sess.Set("username", int(1))
			io.WriteString(w, "警告：页面已过期！")
			return
		} else {
			sess.Set("username", v.(int)+1)
			if v.(int) > 1 { // 测试
				sess.Delete("username") // 删除session
				io.WriteString(w, "警告：检测到重复登录！")
				return
			}
		}

		// 解析表单并处理
		err = r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		if r.Form.Get("usr") != "admin" || r.Form.Get("pwd") != "123456" {
			w.Write([]byte("用户名或密码错误！"))
			return
		}
		w.Write([]byte("登录成功！"))
	}
}

func main() {
	http.HandleFunc("/login", login)
	fmt.Println("server localhost:8080 start up......\nonly use url = http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
