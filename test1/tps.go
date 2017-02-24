package main

import (
	"fmt"
)

func tps(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case string:
		fmt.Printf("string  %s\n", t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

}

func main() {
	/*
		var t interface{}
		t = functionOfSomeType()
	*/
	tps("string")
	tps(10)
	tps(3.14)

}
