package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
	"strings"
)

func test01()  {
	url := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate?source=main-mini"

	req, _ := http.NewRequest("GET", url, nil)

	response, err := http.DefaultClient.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	context,err := ioutil.ReadAll(response.Body)
	fmt.Println(string(context))
	jsondata := gjson.Get(string(context) ,"data").String()
	fmt.Println(jsondata)
	jsonurl := gjson.Get(jsondata,"url").String()

	//var zhuanyi string = "https://passport.bilibili.com/h5-app/passport/login/scan?navhide=1\\u0026qrcode_key=128c6636fe2ebd92ca5f0631c2170f2b\\u0026from=main-mini"
	jieguo := strings.Replace(jsonurl,"\u0026","&",-1)
	fmt.Println(jieguo)
}
func main() {
	test01()
}

//package main
//
//import (
//"github.com/Baozisoftware/qrcode-terminal-go"
//)
//func main() {
//	Test1()
//}
//func Test1(){
//	content := "https://passport.bilibili.com/h5-app/passport/login/scan?navhide=1&qrcode_key=bcb29a95fbba159d79e919dcf964faaa&from=main-mini"
//	obj := qrcodeTerminal.New()
//	obj.Get(content).Print()
//
//}