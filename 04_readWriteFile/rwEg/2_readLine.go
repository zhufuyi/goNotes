package rwEg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadLineContent() {
	// 注：windows 下的记事本换行是'\r' '\n'，linux下的文件换行只有'\n'
	f, err := os.Open("./a.txt") //打开文件
	if err != nil {
		fmt.Println(err.Error()) //打开文件出错处理
		return
	}
	defer f.Close()

	buff := bufio.NewReader(f) //读入缓存
	for {
		lineContent, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			if lineContent != "" { // 最后一行不包括"\r\n",但是有内容
				fmt.Printf("%#v ", lineContent)
			}
			break
		}
		if len(lineContent) > 1 { // Linux：lineContent包括'\n'
			if lineContent[0] == '\r' && lineContent[1] == '\n' { // windos：lineContent包括'\r' '\n'
				continue
			}
			lineContent = strings.Replace(lineContent, "\r", "", -1) // 去掉换行
			lineContent = strings.Replace(lineContent, "\n", "", -1) // 去掉换行
			fmt.Printf("%#v ", lineContent)                          // 行处理
		}
	}
}
