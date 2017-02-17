package main

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"fmt"
	"errors"
	"time"
	"github.com/pborman/uuid"
//	"encoding/json"
)

var tokenDir="/token/"
var ErrorPassword= errors.New("error passowrd")



func getTokenExist(kAPI client.KeysAPI, passwordToken string)error{

	_, err:= kAPI.Get(context.Background(), tokenDir+passwordToken, nil)
	if err!=nil{
		fmt.Println("error to get token")
		return errors.New("token not existed")
	}
	return nil

}


func createUser( kAPI client.KeysAPI, user string, password string)error{
	fmt.Println("go to createUser")

	userpath := fmt.Sprintf("/user/%s",user)
	_, err := kAPI.Create(context.Background(), userpath, password)
	if err != nil {
		// handle error
		fmt.Println("error to create User")
	}

	//_, err = kAPI.Create(context.Background(), tokenDir, passwordToken)
	passwordToken:=uuid.New()
	fmt.Println("print token:",passwordToken)
	expireTime:=client.CreateInOrderOptions{TTL:time.Duration(time.Second*10)}
	_, err = kAPI.CreateInOrder(context.Background(), tokenDir+passwordToken, passwordToken, &expireTime)
	if err != nil {
		fmt.Println("error to create Token")
	}
	if err := getTokenExist(kAPI, passwordToken); err!=nil{
		fmt.Println("haha",err)
		return err
	}
	

	return nil
}


func getValueResponse( res *client.Response)(string, error){

	return res.Node.Value, nil
}


//func checkUser( kAPI client.KeysAPI, user string)(*client.Response, error){

func setToken(kAPI client.KeysAPI, username string)(string, error){
	//username = "lala"
	token:=uuid.New()
	token1:=tokenDir+token
	fmt.Println("into setToken", username)
	var ss=client.SetOptions{TTL:time.Duration(100*time.Second)}
	_, err:= kAPI.Set(context.Background(), token1, token, &ss)
	if err!=nil{
	fmt.Println("set error")
	}
	res, err:= kAPI.Get(context.Background(), token1, nil)
	fmt.Println(res.Node.Value)
	return token, nil
	
}

func login(kAPI client.KeysAPI, username string, password string)(string, error){

	err := checkUserExisted(kAPI, username, password)
	fmt.Println("checkUserExisted",err)
	if err!=nil{
		fmt.Println("error")
		return "", err
		}
	token,_:=setToken(kAPI, username)
	return token, nil

}

func checkUserExisted(kAPI client.KeysAPI, user string, password string)(error){

	userpath := fmt.Sprintf("/user/%s",user)
	res, err:= kAPI.Get(context.Background(), userpath, nil)
	if err!=nil{
		return err
	}
	fmt.Println("the user's password",res.Node.Value)
	if password!=res.Node.Value{
		return errors.New("not right password")
	}
	fmt.Println(res)
	return nil

}

func getToken( kAPI client.KeysAPI, user string)(string, error){
	fmt.Println("go to checkUser")
	userpath := fmt.Sprintf("/user/%s",user)
	res, err:= kAPI.Get(context.Background(), userpath, nil)
	value,_:=getValueResponse(res)
	return value, nil
	fmt.Println(res)
	if err !=nil{
		fmt.Println(err)
		return value, err
	}
	return value, nil	

}

func connectETCD(ip string)(client.KeysAPI, error){
	fmt.Println("start to connect")
	target:=fmt.Sprintf("http://%s:2379",ip)
	cfg := client.Config{
		Endpoints: []string{target},
		//Endpoints: []string{"http://127.0.0.1:2379"},
		//Transport: DefaultTransport,
	}
	c, err := client.New(cfg)
	if err != nil {
		// handle error
	}
	fmt.Println("ready to create User")
	kAPI := client.NewKeysAPI(c)
	return kAPI, nil
}

/*
func main(){
// create a new key /foo with the value "bar"
//fmt.Println(kAPI.(type))
fmt.Println("ready to create User")

kAPI,_:=connectETCD("127.0.0.1")

createUser( kAPI, "lala", "haha")

res, err:= checkUser(kAPI, "lala")
fmt.Println(res)
}

*/
