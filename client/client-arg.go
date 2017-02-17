package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fileupload(token string) {
	in, err := os.Open("./filerandom")
	//in, err := os.Open("./filerandom")
	if err != nil {
		log.Fatal(err)
	}
	pr, pw := io.Pipe()
	bufin := bufio.NewReader(in)
	//gw := gzip.NewWriter(pw)
	go func() {
		log.Printf("Start writing")
		//n, err := bufin.WriteTo(gw)
		n, err := bufin.WriteTo(pw)
		if err != nil {
			log.Fatal(err)
		}
		//      gw.Close()
		pw.Close()
		log.Printf("Done writing: %d", n)
	}()

	url := "http://localhost:8080/upload"
	request, err := http.NewRequest("POST", url, pr)
	//userToken := "1cbc02af-f4c5-4cad-a0bc-5da5fe70af4b"
	userToken := token
	request.Header.Set("Auth-Token", userToken)

	//resp, err := http.DefaultClient.Do(req)
	resp, err := http.DefaultClient.Do(request)

	//response, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}

func createuser(username string, password string) {
	url := "http://localhost:8080/adduser"
	request, _ := http.NewRequest("POST", url, nil)
	request.Header.Add("name", username)
	request.Header.Add("passwd", password)
	resp, _ := http.DefaultClient.Do(request)
	log.Println(resp)
	log.Println(resp.StatusCode)
}

func runlogin(username string, password string) {
	s := 1
	if s != 0 {
		url := "http://localhost:8080/login"
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Add("name", username)
		request.Header.Add("passwd", password)
		//request.Header.Add("name", username)
		//request.Header.Add("passwd", password)
		resp, _ := http.DefaultClient.Do(request)
		log.Println(resp)
		log.Println(resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		u := map[string]interface{}{}
		err := json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}
		fmt.Println(u["auth-token"])
		fmt.Println(resp.StatusCode)
	}
}

func main() {
	var baudrate int
	flag.IntVar(&baudrate, "baudrate", 1200, "help message for flagname")
	var databits int
	flag.IntVar(&databits, "databits", 10, "number of data bits")
	var runcommand string
	flag.StringVar(&runcommand, "runcommand", "", "number of data bits")
	var token string
	flag.StringVar(&token, "token", "111111111", "token in string")

	flag.Parse()
	switch runcommand {
	case "":
		fmt.Println("please adding runcommand: -runcommand")
		os.Exit(3)

	case "runlogin":
		if len(flag.Args()) == 0 {
			fmt.Println("please adding username and password")
			os.Exit(3)
		}
		username := flag.Arg(0)
		password := flag.Arg(1)

		fmt.Println("adding new user with username and password", username, password)
		runlogin(username, password)
	case "createuser":
		fmt.Println("create user")
		if len(flag.Args()) == 0 {
			fmt.Println("please adding username and password")
			os.Exit(3)
		}
		username := flag.Arg(0)
		password := flag.Arg(1)
		createuser(username, password)
	case "fileupload":
		fmt.Println("upload file")
		fmt.Println("create user")
		if len(flag.Args()) == 0 {
			fmt.Println("please adding token ")
			os.Exit(3)
		}
		token := flag.Arg(0)
		fileupload(token)
	case "run":
		fmt.Println("run")

	case "gpsloc":
		fmt.Println("gpsloc")
		if len(flag.Args()) != 1 {
			fmt.Println("not enough argv")
			os.Exit(3)
		}
		token := flag.Arg(0)
		gpsloc(token)
	}
}
