// Excel文件的写入、读取、合并
package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

type Excel struct {
	File     *xlsx.File
	FileName string
}

// 新建文件
func NewExcel(fileName string) (file *Excel) {
	return &Excel{File: xlsx.NewFile(), FileName: fileName}
}

// 写入文件
func (ef *Excel) WriteAll(content [][]string) error {
	sheet := ef.File.AddSheet("sheet1") // 新建一张表
	for _, cells := range content {     // 遍历每一行
		row := sheet.AddRow()
		for _, cell := range cells { // 遍历每个单元
			row.AddCell().Value = cell // 添加一个单元并赋值
		}
	}
	return ef.File.Save(ef.FileName) // 保存文件
}

// 读出文件所有内容
func (ef *Excel) ReadAll() ([][]string, error) {
	xf, err := xlsx.OpenFile(ef.FileName)
	if err != nil {
		return nil, err
	}

	var content [][]string
	for _, sheet := range xf.Sheets {
		for _, row := range sheet.Rows {
			var cellContent []string
			for _, cell := range row.Cells {
				cellContent = append(cellContent, fmt.Sprintf("%v", cell))
			}
			content = append(content, cellContent)
		}
	}
	return content, nil
}

// 指定位置写入内容，参数分别是行、列、内容
func (ef *Excel) WritePst(row int, col int, content []string) error {
	ctVl, err := ef.ReadAll()
	if err != nil {
		return err
	}
	if ctVl == nil {
		ctVl = make([][]string, 1)
	}

	if row == 0 {
		cells := ctVl[0]
		Dcell := col - len(cells)
		if Dcell > 0 {
			for i := 0; i < Dcell; i++ {
				cells = append(cells, "")
			}
			cells = append(cells, content...)
			ctVl[0] = cells
		} else {
			cells = append(cells[:col], content...)
			ctVl[0] = cells
		}

		return ef.WriteAll(ctVl)
	}

	Drow := row - len(ctVl)
	if Drow > 0 { // 指定行小于存在的行
		for i := 0; i < Drow; i++ {
			ctVl = append(ctVl, []string{})
		}

		cells := make([]string, col)
		cells = append(cells, content...)
		ctVl[row-1] = cells
	} else { // 指定行在范围内
		cells := ctVl[row-1]
		Dcell := col - len(cells)
		if Dcell > 0 {
			for i := 0; i < Dcell; i++ {
				cells = append(cells, "")
			}
			cells = append(cells, content...)
			ctVl[row-1] = cells
		} else {
			cells = append(cells[:col], content...)
			ctVl[row-1] = cells
		}
	}
	return ef.WriteAll(ctVl)
}

// 合并exel文件
func Merge(fileNames ...string) error {
	var allContent [][]string
	for _, fileName := range fileNames {
		ef := NewExcel(fileName)
		content, err := ef.ReadAll()
		if err != nil {
			return err
		}
		allContent = append(allContent, content...)
	}
	allef := NewExcel("Merge.xlsx")
	return allef.WriteAll(allContent)
}

func main() {
	// 在新exel文件写入所有内容，如果存在则覆盖
	ef := NewExcel("test.xlsx")
	ctVl := [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}, {"7", "8"}, {"9", "10"}}
	err := ef.WriteAll(ctVl)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Printf("save [%s] success.\n", ef.FileName)
	}

	// 读取exel文件所有内容
	content, err := ef.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Printf("%q\n", content)
	}

	// 修改exel内容
	err = ef.WritePst(10, 0, []string{"20", "30", "40"})
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Printf("change [%s] success.\n", ef.FileName)
		fmt.Println(ef.ReadAll())
	}

	// 合并exel文件内容
	err = Merge("test1.xlsx", "test2.xlsx", "test3.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Merge success.\n")
	}
}
