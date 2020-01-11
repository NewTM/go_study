package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//读取文件内容
	data, err := ioutil.ReadFile("C:/go_project/go_study/io/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//写内容到文件中，没有该文件时自动创建
	str := "hello world"
	err = ioutil.WriteFile("a.txt", []byte(str), os.ModePerm)
	if err != nil {
		panic(err)
	}

}
