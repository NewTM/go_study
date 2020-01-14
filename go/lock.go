package main

import (
	"fmt"
	"sync"
)

var ticket = 1000
var mutex sync.Mutex //创建锁头
var wg sync.WaitGroup //同步等待组对象

//互斥锁 demo
func main(){
	//在使用互斥锁的时候，对资源操作完一定要解锁，否则会出现 死锁或程序异常
	//defer
	wg.Add(4)
	go saleTicket("A")
	go saleTicket("B")
	go saleTicket("C")
	go saleTicket("D")
	wg.Wait()
}

func saleTicket(name string){
	defer wg.Done()
	//rand.Seed(time.Now().UnixNano())
	for {
		//上锁：该部分只会被一个goroutine访问
		mutex.Lock()
		if ticket >0{
			//time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出",ticket)
			ticket--
		}else{
			mutex.Unlock()
			fmt.Print("没票了")
			break
		}
		//解锁
		mutex.Unlock()
	}
}