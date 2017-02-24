package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type transfer interface {
	ToString() string
	ToInt() int64
	ToFloat() float64
	//ToInt() int
}

//impliment Float Function
type FloatTo struct {
	f float64
}

func (ci FloatTo) ToString() string {
	return fmt.Sprintf("%f", ci.f)
}

func (ci FloatTo) ToInt() int64 {
	return int64(ci.f)
}

func (ci FloatTo) ToFloat() float64 {
	return ci.f
}

//impliment Int Function
type IntTo struct {
	f int64
}

func (ci IntTo) ToString() string {
	return fmt.Sprintf("%d", ci.f)
}

func (ci IntTo) ToInt() int64 {
	return ci.f
}

func (ci IntTo) ToFloat() float64 {
	return float64(ci.f)
}

//impliment ByteList Function
type CharListTo struct {
	f []byte
}

func (ci CharListTo) ToString() string {
	return string(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

func (ci CharListTo) ToInt() int64 {
	value, _ := strconv.ParseInt(string(ci.f), 10, 64)
	return value
	//	return int(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

func (ci CharListTo) ToFloat() float64 {
	value, _ := strconv.ParseFloat(string(ci.f), 64)
	return value
	//	return int(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

//impliment String Function
type StringTo struct {
	f string
}

func (str StringTo) ToInt() int64 {
	value, _ := strconv.ParseInt(str.f, 10, 64)
	return value
	//	return int(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

func (str StringTo) ToString() string {

	return str.f
	//	return int(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

func (str StringTo) ToFloat() float64 {
	//value, _ := strconv.ParseInt(str.f, 10, 64)
	value, _ := strconv.ParseFloat(str.f, 64)
	return value
	//	return int(ci.f)
	//	return fmt.Sprintf("%f", ci.f)
}

func ConvertTo(t transfer, convtype string) interface{} {
	if convtype == "string" {
		ct := t.ToString()
		return ct
	}
	if convtype == "int64" {
		ct := t.ToInt()
		return ct
	}

	if convtype == "float64" {
		ct := t.ToFloat()
		return ct
	}
	var ii interface{}
	return ii
}

func DetectType(aa interface{}) string {

	return fmt.Sprintf("%v", reflect.TypeOf(aa))
}

func ConvertToString(in interface{}) string {
	inputtype := DetectType(in)
	if inputtype == "float64" {
		ft := FloatTo{}
		ft.f = in.(float64)
		ret := ConvertTo(ft, "string")
		return ret.(string)
	}
	return ""
}

func Convert(in interface{}, convtype string) interface{} {
	//ConvertTo()
	//transfer.ToString()
	inputtype := DetectType(in)
	if inputtype == convtype {
		fmt.Println("two types are the same")
		return in
	}
	fmt.Println("detect in input type ", inputtype)
	if inputtype == "float64" {
		ft := FloatTo{}
		//transfer interface to float
		ft.f = in.(float64)
		ret := ConvertTo(ft, convtype)
		fmt.Println("shows in convert", ret, "withType", DetectType(ret))
		return ret
	}

	if inputtype == "int64" {
		fmt.Println("running in int converter")
		ft := IntTo{}
		//transfer interface to float
		ft.f = in.(int64)
		ret := ConvertTo(ft, convtype)
		fmt.Println("shows in convert", ret, "withType", DetectType(ret))
		return ret
	}

	//if inputtype == "[]char" {
	// here is []byte
	if inputtype == "[]uint8" {
		fmt.Println("running in int converter")
		ft := CharListTo{}
		//transfer interface to float
		ft.f = in.([]byte)
		ret := ConvertTo(ft, convtype)
		fmt.Println("shows in convert", ret, "withType", DetectType(ret))
		return ret
	}
	if inputtype == "string" {
		fmt.Println("running in int converter")
		ft := StringTo{}
		//transfer interface to float
		ft.f = in.(string)
		ret := ConvertTo(ft, convtype)
		fmt.Println("shows in convert", ret, "withType", DetectType(ret))
		return ret
	}

	var ii interface{}
	return ii
}

func main() {
	//var ff FloatTo
	//ff.f = 1.2345678
	//aa := ff.ToString()
	fmt.Println("----------  running float to string ------------")
	f := 3.1415
	a := Convert(f, "string")
	fmt.Println("result", a, reflect.TypeOf(a))
	fmt.Println("----------  running int to string ------------")
	var i int64
	i = 2
	ii := Convert(i, "string")
	fmt.Println("result", ii, reflect.TypeOf(ii))
	fmt.Println("----------  running charlist to string ------------")
	//	ss := "aaaaaa"
	c := []byte("3.1415")
	cc := Convert(c, "string")
	fmt.Println("result", cc, reflect.TypeOf(cc))

	fmt.Println("----------  running string to string ------------")
	s := "3.1415"
	ss := Convert(s, "string")
	fmt.Println("result", ss, reflect.TypeOf(ss))

	fmt.Println("----------  running interface to string ------------")
	var it interface{}
	//	it = "aaa"
	fmt.Println(it, reflect.TypeOf(it))

	fmt.Println("----------running float to int  ------------")
	fff := 3.1415
	iii := Convert(fff, "int64")
	fmt.Println(iii, reflect.TypeOf(iii))

	fmt.Println("----------int int to int  ------------")
	var iid int64
	iid = 3
	iiii := Convert(iid, "int64")
	fmt.Println(iiii, reflect.TypeOf(iiii))

	fmt.Println("----------bytelist to int  ------------")
	ccc := []byte("3.1415")
	iccc := Convert(ccc, "int64")
	fmt.Println(iccc, reflect.TypeOf(iccc))

	fmt.Println("----------string to int  ------------")
	sti := "3.1415"
	isti := Convert(sti, "int64")
	fmt.Println(isti, reflect.TypeOf(isti))

	fmt.Println("----------string to float  ------------")
	stf := "3.1415"
	sstf := Convert(stf, "float64")
	fmt.Println(sstf, reflect.TypeOf(sstf))

	fmt.Println("----------int to float  ------------")

	var itf int64
	itf = 3
	iitf := Convert(itf, "float64")
	fmt.Println(iitf, reflect.TypeOf(iitf))

	fmt.Println("----------bytelist to float  ------------")
	btf := []byte("3.1415")
	bbtf := Convert(btf, "float64")
	fmt.Println(bbtf, reflect.TypeOf(bbtf))

	fmt.Println("----------float to float  ------------")
	ftf := 3.14159
	fftf := Convert(ftf, "float64")
	fmt.Println(fftf, reflect.TypeOf(fftf))
	fmt.Println("-------------------------------------")
	fmt.Println("----------ConvertToString test---------------------")
	cf := 3.1515
	scf := ConvertToString(cf)
	fmt.Println(scf + "aaa")
}
