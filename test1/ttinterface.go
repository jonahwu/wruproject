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

func ConvertToFloat(it interface{}) float64 {
	aa := retarb(it)
	if fmt.Sprintf("%v", reflect.TypeOf(aa)) == "float64" {
		return aa.(float64)
	}
	return aa.(float64)
}

func ConvertToString(it interface{}) string {
	aa := retarb(it)
	if fmt.Sprintf("%v", reflect.TypeOf(aa)) == "string" {
		return aa.(string)
	}
	return aa.(string)
}

func main() {

	//	arb := retarb("3")
	arb := ConvertToString("3")
	fmt.Println("result string", reflect.TypeOf(arb))
	fmt.Println(arb)

	ii := ConvertToFloat(3.14)
	//ii := retarb(3.14)
	fmt.Println("result float", reflect.TypeOf(ii))
	fmt.Println(2.0 + ii)
}
