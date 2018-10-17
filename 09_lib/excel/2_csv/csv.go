// csv文件的写入和读取
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Csver interface {
	WriteAll(content [][]string) error
	ReadAll() ([][]string, error)
	Write(content []string) error
	Read() ([]string, error)
	Close()
}

type CSV struct {
	File *os.File
	CsvW *csv.Writer
	CsvR *csv.Reader
}

func isExit(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// 新建文件
func OpenCSV(fileName string) (Csver, error) {
	var f *os.File
	var err error
	if isExit(fileName) { //判断文件是否存在
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM格式
	}

	return &CSV{
		File: f,
		CsvW: csv.NewWriter(f),
		CsvR: csv.NewReader(f),
	}, nil
}

// 关闭文件
func (c *CSV) Close() {
	c.File.Close()
}

// 写入所有内容
func (c *CSV) WriteAll(content [][]string) error {
	return c.CsvW.WriteAll(content)
}

// 读取所有内容
func (c *CSV) ReadAll() ([][]string, error) {
	return c.CsvR.ReadAll()
}

// 写入内容
func (c *CSV) Write(content []string) error {
	return c.CsvW.Write(content)
}

// 读取部分内容
func (c *CSV) Read() ([]string, error) {
	return c.CsvR.Read()
}

func main() {
	content := []string{"编号", "姓名", "年龄"}
	csver, _ := OpenCSV("test.csv") // 写入后自动转到下一行
	// 写入一行内容
	csver.Write(content)

	// 写入批量内容
	contentAll := [][]string{{"1", "张三", "23"}, {"2", "李四", "24"}, {"3", "王五", "25"}, {"4", "赵六", "26"}}
	csver.WriteAll(contentAll)

	csver.Close() // 关闭文件

	// 读取一行内容
	csver, _ = OpenCSV("test.csv")
	ct, _ := csver.Read()
	fmt.Println(ct)

	// 读取剩余全部内容
	cta, _ := csver.ReadAll()
	fmt.Println(cta)

	csver.Close() // 关闭文件
}
