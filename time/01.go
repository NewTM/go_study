package main

import (
	"fmt"
	"time"
)

func main() {
	//1\获取当前时间
	now := time.Now()

	//2、获取指定时间

	year, month, day := now.Date() //返回 年月日 ，月是英语的
	hour, min, sec := now.Clock()  //时 分 秒
	fmt.Println(year, month, day)
	fmt.Println(hour, min, sec)
	fmt.Println(now)

	fmt.Println(now.Year())    //年
	fmt.Println(now.YearDay()) //今年已经过了多少天
	fmt.Println(now.Month())   //月
	fmt.Println(now.Day())     //日
	fmt.Println(now.Hour())    //时
	fmt.Println(now.Minute())  //分
	fmt.Println(now.Second())  //秒
	fmt.Println(now.Unix())    //时间戳
	time.Sleep(3 * time.Second)
	fmt.Println("睡了3秒程序结束---")
}
