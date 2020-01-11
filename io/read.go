package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 1、打开文件
// 2、读取数据
// 3、断开链接
// 4、关闭文件资源
const fname string = "C:/go_project/go_study/io/test.txt"

//读取 文件数据的操作
func readTest() {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 1024, 1024)
	n := -1
	for {
		n, err = f.Read(b)
		if err == io.EOF || n == 0 {
			fmt.Println("已读取到文件末尾")
			break
		}
		fmt.Println(string(b[:n]))
	}
	str := string(b)
	fmt.Println(str)
}

// 1、打开文件
// 2、写入数据
// 3、关闭文件
func writeTest() {
	//O_CREATE 没有文件的时候创建新文件
	//O_WRONLY  让文件可写
	//O_APPEND 写操作时是否在文件末尾写，不加该参数会在文件的开头写数据，可能造成数据的覆盖
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs := []byte{65, 66, 67, 68, 69}
	n, err := f.Write(bs)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "....")

	f.WriteString("哈哈哈 哈哈哈") //直接写入字符串 有可能乱码，原因还未找到

	i, err := f.Write([]byte("\n hello world"))
	fmt.Println(i, err)

}

//使用io.copy拷贝文件
func copyFile(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

//使用ioutil包拷贝文件
//一次性读，和一次性写，适合小文件，大文件可能内存会溢出
func copyFile2(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func main() {
	// |os.O_APPEND
	// readTest()
	// writeTest()
	// copyFile(fname, "copy.txt")
	copyFile2(fname, "copy2.txt")
}
