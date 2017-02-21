package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
	"time"
)

var tokenDir = "/token/"
var gpsDir = "/gps/"
var ErrorPassword = errors.New("error passowrd")

type gpslocation struct {
	lati float64
	long float64
}

type AuthInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"userid"`
	Email    string `json:"email"`
}

func getTokenExist(kAPI client.KeysAPI, passwordToken string) error {

	_, err := kAPI.Get(context.Background(), tokenDir+passwordToken, nil)
	if err != nil {
		fmt.Println("error to get token")
		return errors.New("token not existed")
	}
	return nil

}

func createUser(kAPI client.KeysAPI, user string, password string) error {
	fmt.Println("go to createUser")

	userpath := fmt.Sprintf("/user/%s", user)
	var userinfo AuthInfo
	userinfo.UserName = user
	userinfo.Password = password
	userinfo.UserID = uuid.New()
	charuserinfo, _ := json.Marshal(userinfo)
	fmt.Println(charuserinfo)
	_, err := kAPI.Create(context.Background(), userpath, string(charuserinfo))
	//_, err := kAPI.Create(context.Background(), userpath, password)
	if err != nil {
		// handle error
		fmt.Println("error to create User")
	}

	//_, err = kAPI.Create(context.Background(), tokenDir, passwordToken)
	passwordToken := uuid.New()
	fmt.Println("print token:", passwordToken)
	sec, _ := cfg.Int64("token", "expiretime")
	expireTime := client.CreateInOrderOptions{TTL: time.Duration(time.Second * time.Duration(sec))}
	_, err = kAPI.CreateInOrder(context.Background(), tokenDir+passwordToken, passwordToken, &expireTime)
	if err != nil {
		fmt.Println("error to create Token")
	}
	if err := getTokenExist(kAPI, passwordToken); err != nil {
		fmt.Println("haha", err)
		return err
	}

	return nil
}

func getValueResponse(res *client.Response) (string, error) {

	return res.Node.Value, nil
}

//func checkUser( kAPI client.KeysAPI, user string)(*client.Response, error){

func getUserInfo(kAPI client.KeysAPI, username string, usertype string) (string, error) {
	//userinfo := AuthInfo{}
	userinfo := map[string]interface{}{}
	userpath := fmt.Sprintf("/user/%s", username)
	res, _ := kAPI.Get(context.Background(), userpath, nil)
	_ = json.Unmarshal([]byte(res.Node.Value), &userinfo)
	fmt.Println("get userinfo:", userinfo)

	fmt.Println("get userinfo UserID:", userinfo["userid"])
	if usertype == "userid" {
		return fmt.Sprintf("%v", userinfo[usertype]), nil
	}
	return "", errors.New("not on the right target of userinfo")

}

func setToken(kAPI client.KeysAPI, username string) (string, error) {
	//username = "lala"
	token := uuid.New()
	token1 := tokenDir + token
	userID, _ := getUserInfo(kAPI, username, "userID")
	fmt.Println("into setToken", username)
	var ss = client.SetOptions{TTL: time.Duration(100000 * time.Hour)}
	//todo: set the following token as username, means we store token in key and store userinfo in value
	//_, err := kAPI.Set(context.Background(), token1, token, &ss)
	//_, err := kAPI.Set(context.Background(), token1, username, &ss)
	_, err := kAPI.Set(context.Background(), token1, userID, &ss)
	if err != nil {
		fmt.Println("set error")
	}
	res, err := kAPI.Get(context.Background(), token1, nil)
	fmt.Println(res.Node.Value)
	return token, nil

}
func getNameFromToken(token string) (string, error) {
	token1 := tokenDir + token
	res, err := kAPI.Get(context.Background(), token1, nil)
	if err != nil {
		return "", err
	}
	username := res.Node.Value
	fmt.Println("getNameFromToken", res.Node.Value)
	return username, nil
}

func getToken(kAPI client.KeysAPI, user string) (string, error) {
	fmt.Println("go to checkUser")
	userpath := fmt.Sprintf("/user/%s", user)
	res, err := kAPI.Get(context.Background(), userpath, nil)
	value, _ := getValueResponse(res)
	return value, nil
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		return value, err
	}
	return value, nil

}

func login(kAPI client.KeysAPI, username string, password string) (string, error) {

	err := checkUserExisted(kAPI, username, password)
	fmt.Println("checkUserExisted", err)
	if err != nil {
		fmt.Println("error")
		return "", err
	}
	token, _ := setToken(kAPI, username)
	return token, nil

}

func checkUserExisted(kAPI client.KeysAPI, user string, password string) error {

	userpath := fmt.Sprintf("/user/%s", user)
	res, err := kAPI.Get(context.Background(), userpath, nil)
	if err != nil {
		return err
	}
	jsonres := AuthInfo{}
	_ = json.Unmarshal([]byte(res.Node.Value), &jsonres)
	fmt.Println("the json user password", jsonres.Password)
	if password != jsonres.Password {
		return errors.New("not right password")
	}
	fmt.Println(res)
	return nil

}

func connectETCD(ip string) (client.KeysAPI, error) {
	fmt.Println("start to connect")
	target := fmt.Sprintf("http://%s:2379", ip)
	cfgetcd := client.Config{
		Endpoints: []string{target},
		//Endpoints: []string{"http://127.0.0.1:2379"},
		//Transport: DefaultTransport,
	}
	c, err := client.New(cfgetcd)
	if err != nil {
		// handle error
	}
	fmt.Println("ready to create User")
	kAPI := client.NewKeysAPI(c)
	return kAPI, nil
}

func setGpsLoc(kAPI client.KeysAPI, token string, gpsloc gpslocation) (string, error) {
	//username = "lala"
	//token := uuid.New()
	//username := "kala"
	username, _ := getNameFromToken(token)
	gpsDir1 := gpsDir + username
	fmt.Println("into setToken", username)
	sec, _ := cfg.Int64("gps", "expiretime")
	var ss = client.SetOptions{TTL: time.Duration(time.Duration(sec) * time.Hour)}
	strgpsloc := fmt.Sprintf("%f//%f", gpsloc.lati, gpsloc.long)
	_, err := kAPI.Set(context.Background(), gpsDir1, strgpsloc, &ss)
	if err != nil {
		fmt.Println("set error")
	}
	res, err := kAPI.Get(context.Background(), gpsDir1, nil)
	fmt.Println(res.Node.Value)
	return strgpsloc, nil
}

func GetGpsLoc(kAPI client.KeysAPI, clientuserid string) (string, error) {
	gpsDir1 := gpsDir + clientuserid

	res, err := kAPI.Get(context.Background(), gpsDir1, nil)
	if err != nil {
		fmt.Println("set error")
	}
	strgpsloc := string(res.Node.Value)
	return strgpsloc, nil
}

func GetUserID(kAPI client.KeysAPI, username string, password string) (string, error) {

	err := checkUserExisted(kAPI, username, password)
	fmt.Println("checkUserExisted", err)
	if err != nil {
		fmt.Println("error")
		return "", err
	}

	userid, _ := getUserInfo(kAPI, username, "userid")

	return userid, nil
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
