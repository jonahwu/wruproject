package main

import (
	"fmt"
	//"reflect"
	"time"
)

func run(cb chan int) {
	for {
		for i := 0; i < 15; i++ {
			cb <- i
			fmt.Println("sending data", i)
		}
		time.Sleep(time.Second * 3)
		for i := 15; i < 30; i++ {
			cb <- i
		}

		time.Sleep(time.Second * 10)
	}
}

func send(data int) {
	fmt.Println("send data:", data)
}
func main() {
	cb := make(chan int, 10)
	go run(cb)
	time.Sleep(time.Second * 5) //test block
	for cbb := range cb {       //.where cbb is type of int not chan
		fmt.Println(cbb)
		go send(cbb)
	}
}
