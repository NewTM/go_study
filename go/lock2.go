package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg2 *sync.WaitGroup
//读写锁
/*
可以随便读，其他goroutine可以随时读
写的时候，其他goroutine 什么也不能干，不能读也不能写
*/
func main(){
	rwMutex = new(sync.RWMutex)
	wg2 = new(sync.WaitGroup)
	wg2.Add(3)
	go writeData(1)
	go readData(1)
	go writeData(3)

	//wg2.Add(2)
	//go readData(1)
	//go readData(2)
	wg2.Wait()
}

func readData(i int){
	defer wg2.Done()
	fmt.Println(i,"开始读取。。。")
	rwMutex.RLock()
	fmt.Println(i,"读取中。。。")
	time.Sleep(1*time.Second)
	rwMutex.RUnlock()
	fmt.Println(i,"结束----")
}

func writeData(i int){
	defer wg2.Done()
	fmt.Println(i,"开始写---")
	rwMutex.Lock()
	fmt.Println(i,"写入中---")
	rwMutex.Unlock()
	fmt.Println(i,"写结束----")

}