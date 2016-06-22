/* 结论：
三种方式读文件效率比较：
read+buff和bufio这两种方式读文件效率与缓存大小有关,速度差不多
ioutil方式整体效率最高。
*/
package rwEg

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// 方式一：read+buff读文件
func readBuff_ts(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fileContent := make([]byte, 0, 1024000) // 缓存不够大时append次数越多就越慢
	buf := make([]byte, 4096)               // 读速度与缓存大小有关
	for {
		size, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if size == 0 {
			break
		}
		fileContent = append(fileContent, buf[:size]...)
	}
	return fileContent, nil
}

// 方式二：bufio读文件
func bufio_ts(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	fileContent := make([]byte, 0, 2048000) // 缓存不够大时append次数越多就越慢
	buf := make([]byte, 8192)               // 读速度与缓存大小有关
	for {
		size, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			return fileContent, err
		}
		if size == 0 {
			break
		}
		fileContent = append(fileContent, buf[:size]...)
	}
	return fileContent, nil
}

// 方式三：ioutil读文件
func ioutil_ts(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ReadRateCompare() {
	var file = "./conf.ini"
	t := time.Now()
	_, err := readBuff_ts(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("    read+buff: %v\n", time.Now().Sub(t))

	time.Sleep(1 * time.Second)

	t = time.Now()
	bufio_ts(file)
	fmt.Printf("    bufio: %v\n", time.Now().Sub(t))

	time.Sleep(1 * time.Second)

	t = time.Now()
	ioutil_ts(file)
	fmt.Printf("    ioutil: %v\n", time.Now().Sub(t))

	time.Sleep(1 * time.Second)
}
