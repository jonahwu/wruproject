package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, _ := goconfig.LoadConfigFile("./conf.ini")
	cfg.MustValue("server", "address", "localhost")

	value, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	fmt.Println(value)
	value1, _ := cfg.GetValue("server", "port")
	fmt.Println(value1)
	addr, _ := cfg.GetValue("server", "address")
	fmt.Println(addr)

}
