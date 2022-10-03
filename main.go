package main

import (
	"Bili_go/Plugin"
	"Bili_go/util"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"time"
)
func login(){
	//废话部分
	fmt.Println("==========================================================================================")
	fmt.Println("			 哔哩哔哩Up小助手											")
	fmt.Println("         	 您可以在config文件夹下配置您的插件信息						           ")
	fmt.Println("==========================================================================================")
	fmt.Println("请把窗口调到全屏10秒后扫描二维码（因为二维码太大控制台显示不完全（不会调二维码大小qwq））")
	//倒计时部分
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(1)*time.Second)
		if i == 9 {
			fmt.Printf("%d\n",i+1)
		}else{
			fmt.Printf("%d\t",i+1)
		}
	}
	//开始调用接口输出二维码
	analysisjson := util.AnalysisJson{}
	analysisurl := util.AnalysisUrl{}
	analysisjson.GetQrCodeJson()
	analysisjson.GetQrCodeKey()
	analysisjson.GetQrCodeUrl()
	//开始打印二维码
	obj := qrcodeTerminal.New()
	obj.Get(analysisjson.GetQrCodeUrl()).Print()
	//开始3秒调用一次 qrcode查看是否确认登录
	for true{
		time.Sleep(time.Duration(3)*time.Second)
		if analysisjson.GetQrCodeStateCode() == 0{
			fmt.Println("已登录（修改密码可删除cookie,防止他人盗号）")
			fmt.Printf("登录账号uid为:%s\n",analysisurl.GetUserID(analysisjson.LoginInfoUrl))
			fmt.Println("===============个人重要信息=============")
			fmt.Println("session:"+analysisurl.GetSession(analysisjson.LoginInfoUrl))
			fmt.Println("crsf:"+analysisurl.GetCrsf(analysisjson.LoginInfoUrl))
			fmt.Println("===============个人重要信息=============")
			break
		}

	}
	replyprivatemsg := Plugin.ReplyRrivateMsg{}
	replyprivatemsg.ReplyPrivate(analysisurl.Session,analysisurl.UserID,analysisurl.Crsf)

}

func main() {

	login()

}

