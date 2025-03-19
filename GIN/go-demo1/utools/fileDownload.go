package utools

import (
	"log"
	"os"
	"path/filepath"
)

/**
实现前端文件传入的功能
 1. 读取前端传入的文件
 2. 获取文件名，文件大小，文件类型
 3. 存放地址
 4. 提供下载
**/

func CreateIfExistNo(path string) (string, bool) {

	_, err := os.Stat(path)
	if err == nil {
		log.Println("current path file exist !")
	}

	parentPath := filepath.Dir(path)

	os.MkdirAll(parentPath, os.ModePerm)

	return path, true
}
