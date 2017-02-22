package main

import (
	"fmt"
	"reflect"
)

func retarb(it interface{}) interface{} {
	fmt.Println(reflect.TypeOf(it))
	if fmt.Sprintf("%v", reflect.TypeOf(it)) == "string" {
		return "3"
	}
	if fmt.Sprintf("%v", reflect.TypeOf(it)) == "float64" {
		return 3.14
	}
	return nil
}

func main() {

	arb := retarb("3")
	fmt.Println("result sting", reflect.TypeOf(arb))

	ii := retarb(3.14)
	fmt.Println("result float", reflect.TypeOf(ii))
}
