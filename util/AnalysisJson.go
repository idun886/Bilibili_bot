package util

import (
	"Bili_go/BilibiliApi"
	"Bili_go/ClientMethod"
	"github.com/tidwall/gjson"
)

type AnalysisJson struct {
	JsonContext string
	LoginInfoUrl string
}
//获取qrcode响应的方法
func (a *AnalysisJson) GetQrCodeJson() {
	bilibiliapirequest := ClientMethod.BilibiliApiRequest{}
	contextbody := bilibiliapirequest.GetApi(BilibiliApi.LoginBilibiliApi,"")
	contextdata := gjson.Get(contextbody,"data").String()
	a.JsonContext = contextdata
	//fmt.Println(a.JsonContext)
}

func (a*AnalysisJson)GetQrCodeKey() string{

	qrcodekey := gjson.Get(a.JsonContext,"qrcode_key").String()
	//fmt.Println("==========log==========")
	//fmt.Println(qrcodekey)
	//fmt.Println("调用getqrcodekey方法")
	//fmt.Println("==========log==========")
	return qrcodekey
}
//获取qrcodeurl 的解析apiJSON的方法
func (a *AnalysisJson)GetQrCodeUrl() string {

	qrcodeurl := gjson.Get(a.JsonContext,"url").String()
	//fmt.Println("==========log==========")
	//fmt.Println("调用getqrcodeurl方法")
	//fmt.Println(qrcodeurl)
	//fmt.Println("==========log==========")
	return qrcodeurl
}
// 获取 qrcode状态的 api 字符串方法
func (a *AnalysisJson)GetQrCodeStateApiUrl() string{
	return "https://passport.bilibili.com/x/passport-login/web/qrcode/poll?qrcode_key=" + a.GetQrCodeKey() + "&source=main_mini"
}
//使用get调用api然后获取qrcode的状态码  0为已扫码 其他为未扫码或扫码未确认 返回int64类型
func (a *AnalysisJson)GetQrCodeStateCode() int64{
	bilibiliapirequest := ClientMethod.BilibiliApiRequest{}
	contextbody := bilibiliapirequest.GetApi(a.GetQrCodeStateApiUrl(),"")
	contextdata := gjson.Get(contextbody,"data").String()
	contexturl := gjson.Get(contextdata,"url").String()
	a.LoginInfoUrl = contexturl
	contextcode := gjson.Get(contextdata,"code").Int()
	return contextcode
}





