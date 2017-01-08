/*
获取指定文件夹的文件列表
	fileList,err:=ioutil.ReadDir(path)
	fileList的方法
	Name()	// 文件名
	Size()	// 文件大小
	Mode() FileMode	// 文件权限
	ModTime() time.Time 	// 文件修改时间
*/

package rwEg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 获取指定文件夹所有文件名，不包括文件夹，(注：文件名不包括路径)
func ListFiles(path string) ([]string, error) {
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, v := range fileList {
		if !v.IsDir() {
			files = append(files, v.Name())
		}
	}
	return files, nil
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤，(注：文件名包括路径)
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。(注：文件名没有包括根路径)
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func GetFileList() {
	files, err := ListFiles(".")
	fmt.Println("\n获取指定文件夹所有文件名。\n", files, len(files), err)

	files, err = ListDir(".", ".txt")
	fmt.Println("\n获取指定目录下的所有文件，匹配后缀过滤，不进入下一级目录搜索\n", files, len(files), err)

	files, err = WalkDir(".", ".go")
	fmt.Println("\n获取指定目录下的所有文件，可以匹配后缀过滤，包括索引子目录文件\n", files, len(files), err)
}
