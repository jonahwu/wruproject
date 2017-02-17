package main

import (
	"fmt"
	"time"
)

type gpslocation struct {
	lati float64
	long float64
}

//note channel is blocking
//func callinearloc() (gpslocation, error) {

func callinearloc() chan gpslocation {
	gloc := make(chan gpslocation)
	gg := gpslocation{}
	gg.lati = 23.3333333
	gg.long = 123.333333
	step := 0.0001
	dist := 0.002
	gpsx := gg.lati
	gpsy := gg.lati
	distx := 0.0
	disty := 0.0
	direct := 1.0
	go func() {
		for {
			//fmt.Println("dist loop")
			distx = distx + step*direct
			disty = disty + step*direct
			gpsx = gg.lati + distx
			gpsy = gg.lati + disty
			//fmt.Println("gpsloc", gpsx, gpsy)
			if distx > dist {
				direct = direct * -1
			}
			if distx <= 0 {
				direct = direct * -1
			}

			gggloc := gpslocation{gpsx, gpsy}
			//		time.Sleep(time.Second * 10)
			gloc <- gggloc

		}
	}()
	//fmt.Println("return chan first")
	return gloc
	//	newloc := gpslocation{lati: gpsx, long: gpsy}
	//return newloc, nil
}

func gencalllinearloc(ggloc chan gpslocation) gpslocation {
	//here we reuturn value not chan, you can think as type transformation
	return <-ggloc
}

func gpsloc() {
	ggloc := make(chan gpslocation)
	ggloc = callinearloc()
	for {
		gpsloc := gencalllinearloc(ggloc)
		fmt.Println("the gpslocation is ", gpsloc.lati, gpsloc.long)
		time.Sleep(time.Second * 1)
	}
}
