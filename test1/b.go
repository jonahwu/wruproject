package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile

func init() {
	//	cfg, _ = goconfig.LoadConfigFile("./conf.ini")
	cfg, _ = goconfig.LoadConfigFile("./conf.ini")
	cfg.MustValue("server", "address", "localhost")
}

func main() {

	value, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	fmt.Println(value)
	value1, _ := cfg.GetValue("server", "port")
	fmt.Println(value1)
	addr, _ := cfg.GetValue("server", "address")
	fmt.Println(addr)
	test1()

}
