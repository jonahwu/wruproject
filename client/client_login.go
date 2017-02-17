package main

import (
	//"bufio"
	//	"compress/gzip"
	"fmt"
	//	"io"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
)

//req, err := http.NewRequest("GET", "http://localhost:8080/", pr)
//client := &http.Client{}
//req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
func main() {

	url := "http://localhost:8080/login"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("name", "kala")
	request.Header.Add("passwd", "kala")
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
