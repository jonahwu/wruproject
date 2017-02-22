package main

import (
	"fmt"
	"strconv"
)

func main() {
	//str := "3.1415"
	str := "3.1415"
	f, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("something wrong", err)
	}
	fmt.Println(f)
}
