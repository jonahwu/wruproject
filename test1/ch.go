package main

import (
	"fmt"
	"time"
)

type tla struct {
	server bool
}

type ts struct {
	name    string
	oserver tla
}

func runSecond(a chan ts) {
	go func() {
		la := ts{}
		la.name = "haha"
		la.oserver.server = false
		for i := 100; i < 200; i++ {
			a <- la
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()
}

func test(a chan ts) {

	runSecond(a)

	is := ts{}
	is.name = "test"
	is.oserver.server = true
	for i := 0; i < 100; i++ {
		a <- is
		time.Sleep(time.Second * time.Duration(1))
	}
}

func main() {
	a := make(chan ts)
	go test(a)
	for {
		b := <-a
		fmt.Println(b)
	}
}
