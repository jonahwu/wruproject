package main

import (
	"fmt"
	"net/http"
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

func gpsloc(token string) {
	ggloc := make(chan gpslocation)
	ggloc = callinearloc()
	for {
		gpsloc := gencalllinearloc(ggloc)
		fmt.Println("the gpslocation is ", gpsloc.lati, gpsloc.long)
		setgpsloc(token, gpsloc)
		time.Sleep(time.Second * 1)

	}
}

func setgpsloc(token string, gpsloc gpslocation) {
	url := "http://localhost:8080/setgpsloc"
	request, _ := http.NewRequest("POST", url, nil)
	userToken := token
	request.Header.Set("Auth-Token", userToken)
	request.Header.Set("lati", fmt.Sprintf("%f", gpsloc.lati))
	request.Header.Set("long", fmt.Sprintf("%f", gpsloc.long))

	resp, _ := http.DefaultClient.Do(request)
	fmt.Println(resp)
	fmt.Println(resp.StatusCode)
}
