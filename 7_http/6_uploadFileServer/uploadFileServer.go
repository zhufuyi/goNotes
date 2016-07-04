// 测试上传文件，解析从客户段批量上传过来的表单文件并保存到磁盘里
package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

var token string

// 解析保存单个文件表单
func parseSaveOnefile(w http.ResponseWriter, r *http.Request) bool {
	filer, fileHeader, err := r.FormFile("uploadFile") // 解析文件
	if err != nil {
		w.Write([]byte(err.Error()))
		return false
	}
	defer filer.Close()

	f, err := os.OpenFile(fileHeader.Filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		w.Write([]byte(err.Error()))
		return false
	}
	defer f.Close()

	_, err = io.Copy(f, filer)
	if err != nil {
		w.Write([]byte(err.Error()))
		return false
	}

	return true
}

// 解析保存批量文件表单
func parseSaveMultiFile(w http.ResponseWriter, r *http.Request) bool {
	if files, ok := r.MultipartForm.File["uploadFile"]; ok { // 多部件表单文件slice
		for _, header := range files {
			file, err := header.Open()
			if err != nil {
				w.Write([]byte(err.Error()))
				return false
			}

			f, err := os.OpenFile(header.Filename, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				w.Write([]byte(err.Error()))
				return false
			}

			io.Copy(f, file)
			file.Close()
			f.Close()
		}
	} else {
		return false
	}
	return true
}

func generateToken() string {
	digest := md5.New()
	_, err := io.WriteString(digest, fmt.Sprintf("%d", time.Now().Unix()))
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", digest.Sum(nil))
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// 生成随机的token令牌
		token = generateToken()
		if token == "" {
			w.Write([]byte("错误：token为空!"))
			return
		}
		fmt.Println("token =", token)

		// 解析html模板并发送给客户端
		tl, err := template.ParseFiles("./upload.html")
		if err != nil {
			http.NotFound(w, r)
		}
		err = tl.Execute(w, token)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

	case "POST":
		// 解析表单
		err := r.ParseMultipartForm(32 << 20) // 当form的编码类型为multipart/form-data时，使用r.ParseMultipartForm，否则可以使用r.ParseForm,固定最大缓存为10M
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		// 验证token令牌
		postToken := r.Form.Get("token")
		if postToken == "" {
			w.Write([]byte("警告：token值为空。"))
			return
		}
		if token != postToken {
			w.Write([]byte("警告：token不合法。"))
			return
		}

		token = "" // 防止多次重传

		// 复制文件
		ok := parseSaveMultiFile(w, r)
		if ok {
			w.Write([]byte("上传文件成功。"))
		} else {
			w.Write([]byte("上传文件失败。"))
		}

	default:
		w.Write([]byte("\n警告：只允许GET或POST的请求方法。\n"))
	}
}

func main() {
	http.HandleFunc("/upload", UploadFile)
	fmt.Println("server localhost:8080 start up......\nonly use url = http://localhost:8080/upload")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
