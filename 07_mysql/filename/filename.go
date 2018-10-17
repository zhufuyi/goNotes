package filename

import (
	"strings"
	"os"
	"path/filepath"
)

// 获取exel类型文件接口
type Exeler interface {
	GetExelFiles() ([]string, error)
}

// 实例化
func Getfiles(path string) Exeler {
	return &fil{path:path}
}

// 获取文件的对象
type fil struct {
	path string
}

// 获取文件夹下所有的文件
func (fl *fil)getFilelist() ([]string, error) {
	var files []string
	err := filepath.Walk(fl.path, func(path string, f os.FileInfo, err error) error {
		if (f == nil ) {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}


// 获取exel类型文件
func (fl *fil)GetExelFiles() ([]string, error) {
	files, err := fl.getFilelist()
	if err != nil {
		return nil, err
	}

	var exelFiles  []string
	for _, file := range files {
		if len(file) > 4 {
			tmp := strings.ToLower(file)
			if strings.HasSuffix(tmp, ".xlsx") || strings.HasSuffix(tmp, ".xls") {
				// 后缀为.xlsx或.xls
				exelFiles = append(exelFiles, file)
			}
		}
	}
	return exelFiles, nil
}
