package main

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/pborman/uuid"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	//"time"
	//	"bufio"
	//"log"
)

var kAPI client.KeysAPI

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var fileloc string

func init() {

	fileloc = "/mnt/dataloc/"
	err := os.Mkdir(fileloc, 0777)
	if err != nil {
		fmt.Println("data file existed")
	}
}

//curl -v -X POST -N --data-binary @"./haproxy.cfg" http://localhost:8080/webUploadHandler
func webUploadHandler(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New()
	fmt.Println(uuid)
	out, err := os.Create(fileloc + uuid + ".txt")
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}
	defer out.Close()

	// write the content from POST to the file
	//_, err = io.Copy(out, file)
	fmt.Println("start to copy")
	io.Copy(out, r.Body)
	fmt.Println("File uploaded successfully : ")

}

func webDownloadHandler(w http.ResponseWriter, r *http.Request) {
	in, err := os.Open("./filerandom")
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}
	defer in.Close()
	//pr, pw := io.Pipe()
	//bufin := bufio.NewReader(in)

	// write the content from POST to the file
	//_, err = io.Copy(out, file)
	buf := make([]byte, 1024)
	for {
		n, _ := in.Read(buf)
		if n == 0 {
			break
		}
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}

	}

	fmt.Println("File uploaded successfully : ")

}
func webDownloadHandler1(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New()
	fmt.Println(uuid)
	out, err := os.Open("./file")
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}
	defer out.Close()

	// write the content from POST to the file
	//_, err = io.Copy(out, file)
	fmt.Println("start to copy")
	buffer := make([]byte, 100)
	_, _ = out.Read(buffer)
	//io.Copy(out, buffer)
	w.Write(buffer)
	fmt.Println("File uploaded successfully : ")

}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("into middle1")

		// put process here

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
func tokenHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("into middle recover")

		// put process here
		token := r.Header.Get("Auth-Token")
		fmt.Println(token)
		if token == "powertoken" {
			next.ServeHTTP(w, r)
		}
		err := getTokenExist(kAPI, token)
		if err == nil {
			next.ServeHTTP(w, r)
		} else {

			http.Error(w, http.StatusText(401), 401)
		}
		// error oocured
	}
	return http.HandlerFunc(fn)
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//a := make(map[string]string)
	a := map[string]interface{}{}
	a["test"] = "haha"
	jsonString, _ := json.Marshal(a)
	//w.WriteHeader(200)
	w.Write(jsonString)
}

func getInfoHandler1(str string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am here with str:", str)
		w.Header().Set("Content-Type", "application/json")
		//a := make(map[string]string)
		a := map[string]interface{}{}
		a["test"] = "haha"
		jsonString, _ := json.Marshal(a)
		//w.WriteHeader(200)
		w.Write(jsonString)
	})
}

func addUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("aaa")
	accname := r.Header.Get("name")
	accpasswd := r.Header.Get("passwd")
	fmt.Println("user -------", accname, accpasswd)

	createUser(kAPI, accname, accpasswd)
	/*
		a := map[string]interface{}{}
		a["auth-token"] = "haha"
		jsonString, _ := json.Marshal(a)
		w.Write(jsonString)
	*/
	w.WriteHeader(200)
}

func authLoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uname := r.Header.Get("name")
	upass := r.Header.Get("passwd")
	as, _ := login(kAPI, uname, upass)
	fmt.Println("the http token:", as)
	a := map[string]interface{}{}
	a["auth-token"] = as
	jsonString, _ := json.Marshal(a)
	w.Write(jsonString)
}

func ttHandler(aa string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("in tt handler", aa)

		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

	}
}

func Func1(foo, foo2, timeoutMessage string) alice.Constructor {
	//... something to do with foo and foo2
	fmt.Println("into Func1")
	return func(h http.Handler) http.Handler {
		return http.TimeoutHandler(h, 1*time.Second, timeoutMessage)
	}
}

func middlewareGenerator(foo, foo2 string) (mw func(http.Handler) http.Handler) {

	mw = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use foo1 & foo2
			fmt.Println("in to middleware with parameters with ", foo, foo2)
			h.ServeHTTP(w, r)
		})
	}
	return
}

func setGpsLocHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("into gps loc handler")
	gpsloc := gpslocation{}
	gpsloc.lati, _ = strconv.ParseFloat(r.Header.Get("lati"), 64)
	gpsloc.long, _ = strconv.ParseFloat(r.Header.Get("long"), 64)
	token := r.Header.Get("Auth-Token")
	as, _ := setGpsLoc(kAPI, token, gpsloc)
	fmt.Println("the http token:", as)
	//a := map[string]interface{}{}
	//a["auth-token"] = as
	//jsonString, _ := json.Marshal(a)
	//w.Write(jsonString)

}

func startWeb() {
	commonHandlers := alice.New(loggingHandler, tokenHandler, middlewareGenerator("foo", "foo2"))
	router := httprouter.New()
	router.POST("/upload", wrapHandler(commonHandlers.ThenFunc(webUploadHandler)))
	router.GET("/version", wrapHandler(commonHandlers.ThenFunc(getInfoHandler1("haha"))))
	router.GET("/download", wrapHandler(commonHandlers.ThenFunc(webDownloadHandler)))
	router.POST("/adduser", addUserHandler)
	router.GET("/login", authLoginHandler)
	router.POST("/setgpsloc", wrapHandler(commonHandlers.ThenFunc(setGpsLocHandler)))

	aa := "strrrrr"
	router.GET("/tt", ttHandler(aa))
	http.ListenAndServe(":8080", router)
}

func main() {
	//println("start web")
	kAPI, _ = connectETCD("127.0.0.1")
	go startWeb()
	fmt.Println("start over")
	var input string
	fmt.Scanln(&input)
}
