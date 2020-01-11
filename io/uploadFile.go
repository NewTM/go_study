package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

//断点续传
func main() {
	// 1、注意权限问题
	// 2、
	srcFile := "C:/go_project/go_study/io/abc.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)

	tempFile, _ := filepath.Abs(destFile + "temp.txt")

	file1, err := os.Open(srcFile)
	handleErr(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	handleErr(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	handleErr(err)
	fmt.Println(err, "---==")

	defer file1.Close()
	defer file2.Close()

	file3.Seek(0, io.SeekStart)

	bs := make([]byte, 100, 100)

	n1, err := file3.Read(bs)
	fmt.Println(err)
	// handleErr(err)
	countStr := string(bs[:n1])

	count, err := strconv.ParseInt(countStr, 10, 64)
	// handleErr(err)

	fmt.Println("count=", count)

	file1.Seek(0, io.SeekStart)
	file2.Seek(0, io.SeekStart)

	data := make([]byte, 1024, 1024)

	n2 := -1
	n3 := -1
	total := int(count)

	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Println("已复制数据总量", total)
		// if total > 8000 {
		// 	panic("--~假装断电")
		// }
	}
}
