
package main

import (
//	"bufio"
//	"compress/gzip"
	"fmt"
//	"io"
	"io/ioutil"
	"log"
	"net/http"
        "encoding/json"
	
	"os"
)



//req, err := http.NewRequest("GET", "http://localhost:8080/", pr)
//client := &http.Client{}
//req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
func main(){



url := "http://localhost:8080/version"
request, err := http.NewRequest("GET", url, nil)

//resp, err := http.DefaultClient.Do(req)
resp, err := http.DefaultClient.Do(request)

//response, err := (&http.Client{}).Do(request)

if err != nil {
	log.Fatal(err)
}
log.Println(resp)
      body, _ := ioutil.ReadAll(resp.Body)
       u := map[string]interface{}{}
       err = json.Unmarshal(body, &u)
       if err != nil {
              panic(err)
       }
	fmt.Println(u["test"])
	fmt.Println(resp.StatusCode)
//        bodystr := string(body);
//        fmt.Println(bodystr)
}
