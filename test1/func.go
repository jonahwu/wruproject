package main

import "fmt"

type aa struct {
	s float64
}

func (a aa) Reset1() {
	a.s = 0.0
}
func (a *aa) Reset2() error {
	a.s = 0.0
	return nil
}

func main() {
	sa := aa{}
	sa.s = 1
	fmt.Println(sa.s)
	sa.Reset1()
	fmt.Println(sa.s)
	sa.Reset2()
	fmt.Println(sa.s)

}
