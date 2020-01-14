package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	filePath1 := "/Users/mt/study/go_study/file/a.txt"
	filePath2 := "a.txt"

	fmt.Println(filepath.IsAbs(filePath1)) //是否绝对路径
	fmt.Println(filepath.IsAbs(filePath2))
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2))

	fmt.Println(path.Join(filePath1, "..")) //获取父目录

	err := os.Mkdir("/Users/mt/study/mkdirtest", os.ModePerm) //创建文件夹
	if err != nil {
		panic(err)
	}

}
