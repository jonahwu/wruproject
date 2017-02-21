package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"reflect"
)

var cfg *goconfig.ConfigFile

func init() {
	//	cfg, _ = goconfig.LoadConfigFile("./conf.ini")
	//cfg, _ = goconfig.LoadConfigFile("./conf.ini")
	cfg, _ = goconfig.LoadConfigFile("/etc/testconf/conf.ini")
	_, _ = cfg.MustValueSet("server", "address", "localhost")

}

func main() {

	value, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	fmt.Println(value)
	value1, _ := cfg.GetValue("server", "port")
	value2, _ := cfg.Int64("server", "port")
	fmt.Println(value1)
	fmt.Println(reflect.TypeOf(value1))
	fmt.Println("the value2", value2)
	fmt.Println(reflect.TypeOf(value2))
	addr, _ := cfg.GetValue("server", "address")
	fmt.Println("print address", addr)
	//	test1()

}
