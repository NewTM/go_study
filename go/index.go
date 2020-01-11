package main

// import "fmt"
func a(){
	for index := 0; index < 100; index++ {
		print(index)
	}
}

func b(){
	for index := 0; index < 100; index++ {
		print(index,"----")
	}
}

func main(){
	go a()
	go b()

}