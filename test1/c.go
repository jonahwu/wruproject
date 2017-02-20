package main

import "fmt"

func test1() {
	addr, _ := cfg.GetValue("server", "address")
	fmt.Println(addr)
}
