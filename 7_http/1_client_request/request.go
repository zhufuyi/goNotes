package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func httpGet() ([]byte, error) {
	resp, err := http.Get("http://www.12306.cn")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpPost() ([]byte, error) {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpPostForm() ([]byte, error) {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpDo() ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	dspl("httpGet", httpGet)
	dspl("httpPost", httpPost)
	dspl("httpPostForm", httpPostForm)
	dspl("httpDo", httpDo)
}

func dspl(funName string, f func() ([]byte, error)) {
	fmt.Printf("\n\n--------------- %s ---------------\n", funName)
	data, err := f()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}
