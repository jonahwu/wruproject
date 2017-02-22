package main

import (
	"fmt"
	"reflect"
)

type aa struct {
	ff int
}

type face interface {
	Show1() string
}

func transfer(f face) {
	f.Show1()
}

// remember you have use Capital
func (c *aa) Show1() string {
	fmt.Println("now in show")
	fmt.Println(c.ff)
	a := "aaaaaa"
	return a
}

func main() {
	var t aa
	commits := map[string]string{
		"ss":  "Show1",
		"r":   "2138",
		"gri": "1908",
		"adg": "912",
	}

	fmt.Println(commits["rsc"])
	s1 := commits["ss"]
	//	s := "Show1"
	aa := reflect.ValueOf(&t).MethodByName(s1).Call([]reflect.Value{})
	fmt.Println("result", aa[0])

	// it cannot use in interface
	var tt face
	aa1 := reflect.ValueOf(&tt).MethodByName(s1).Call([]reflect.Value{})
	fmt.Println("result1", aa1)
	//reflect.ValueOf(&t).MethodByName("Foo").Call([]reflect.Value{})

}

//reflect.ValueOf(&t).MethodByName(s).Call([]reflect.Value{})
//one can use reflect.TypeOf(aa)
//aa := reflect.ValueOf(&t).MethodByName(s).Call([]reflect.Value{})[0]
