package main

import "fmt"

func aa(ch chan <- int){
	ch<-10
}

func main(){
	ch1 :=make(chan int)

	go aa(ch1)

	data :=<-ch1

	fmt.Println(data)
}