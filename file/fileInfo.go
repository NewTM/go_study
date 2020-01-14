package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("/Users/mt/study/go_study/file/a.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())  //文件名
	fmt.Println(fileInfo.Size())  //文件大小 字节
	fmt.Println(fileInfo.IsDir()) //是否文件夹
}
