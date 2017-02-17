package main

import "fmt"

func test() (int, error) {
	fmt.Println("aaa")
	return 2, nil
}

func main() {
	a, _ := test()
	fmt.Println(a)

}
