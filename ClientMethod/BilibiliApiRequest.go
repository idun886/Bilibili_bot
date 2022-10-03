package ClientMethod

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BilibiliApiRequest struct {
	ApiUrl string
	PostData string
}

//get请求方法 返回一个响应字符串
func (b *BilibiliApiRequest)GetApi(ApiUrl string,requestheader string) string {
	Client := &http.Client{
		//定义请求事件2秒
		//Timeout: time.Second * 2,
	}
	req, _ := http.NewRequest("GET", ApiUrl, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cookie",requestheader)
	response, err := Client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}

	context,err := ioutil.ReadAll(response.Body)
	return string(context)

	//Client := &http.Client{
	//	//定义请求事件2秒
	//	//Timeout: time.Second * 2,
	//}
	//req, _ := http.NewRequest("GET", ApiUrl, strings.NewReader(requestheader))
	//response, err := Client.Do(req)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//context,err := ioutil.ReadAll(response.Body)
	//return string(context)
}

//post方法 返回 接口的响应字符串
func (b *BilibiliApiRequest)PostApi(ApiUrl string,PostData string,header string) string{
	client := &http.Client{}
	request,err := http.NewRequest("POST",ApiUrl,strings.NewReader(PostData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "err"
	}
	request.Header.Add("cookie",header)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	context,_:= ioutil.ReadAll(response.Body)
	return string(context)
}