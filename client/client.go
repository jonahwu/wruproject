package main

import (
	//"bufio"
	//	"compress/gzip"
	"fmt"
	//	"io"
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	//      "compress/gzip"
	//      "fmt"
	"io"
	//      "io/ioutil"
	"os"
	//"os"
)

func upload(token string) {
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
		//              gw.Close()
		pw.Close()
		log.Printf("Done writing: %d", n)
	}()

	url := "http://localhost:8080/upload"
	request, err := http.NewRequest("POST", url, pr)
	//request.Header.Set("Auth-Token","f184f5d1-87b9-4bc5-9274-5f86e4bf907d")
	request.Header.Set("Auth-Token", token)

	//resp, err := http.DefaultClient.Do(req)
	resp, err := http.DefaultClient.Do(request)

	//response, err := (&http.Client{}).Do(request)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

}

//req, err := http.NewRequest("GET", "http://localhost:8080/", pr)
//client := &http.Client{}
//req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
func login() (string, error) {

	url := "http://localhost:8080/login"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("name", "kala")
	request.Header.Add("passwd", "kala")
	resp, _ := http.DefaultClient.Do(request)
	log.Println(resp)
	log.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	//u := map[string]interface{}{}
	u := map[string]string{}
	err := json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u["auth-token"])
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return "", errors.New("can not get token")
	}
	strr := string(u["auth-token"])
	return strr, nil
}

func main() {
	token, err := login()
	if err != nil {
		fmt.Println(err)
		return
	}
	upload(token)
}
