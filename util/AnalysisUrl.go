package util

import "strings"

type AnalysisUrl struct {
	ApiUrl string
	Session string
	Crsf string
	UserID string
	UserIDMD5 string
	Expires string
}
//解析登录后要跳转的 个人信息Url url里有 session csrf  uid uidmd5 等重要信息
//获取session
func (a *AnalysisUrl)GetSession(LoginUrl string) string{
	sessstart := strings.Index(LoginUrl, "SESSDATA=")
	sesslast := strings.Index(LoginUrl,"&bili_jct")
	a.Session = LoginUrl[sessstart:sesslast]
	return LoginUrl[sessstart:sesslast]
}
//获取crsf
func (a *AnalysisUrl)GetCrsf(LoginUrl string) string {
	crefstart := strings.Index(LoginUrl,"&bili_jct=")
	creflast := strings.Index(LoginUrl,"&gourl")
	a.Crsf = LoginUrl[crefstart+10:creflast]
	return LoginUrl[crefstart+10:creflast]
}
//获取userid(用户uid)
func (a *AnalysisUrl)GetUserID(LoginUrl string) string  {
	useridstart := strings.Index(LoginUrl,"DedeUserID=")+11
	useridlast := strings.Index(LoginUrl,"&DedeUserID__ckMd5=")
	a.UserID = LoginUrl[useridstart:useridlast]
	return LoginUrl[useridstart:useridlast]
}
//获取userid的md5格式
func (a *AnalysisUrl)GetUserIDMD5(LoginUrl string) string  {
	useridmd5start := strings.Index(LoginUrl,"DedeUserID__ckMd5=")+18
	useridmd5last := strings.Index(LoginUrl,"&Expires")
	a.UserIDMD5 =LoginUrl[useridmd5start:useridmd5last]
	return LoginUrl[useridmd5start:useridmd5last]
}
//获取登录的时间戳
func (a *AnalysisUrl)GetExpires(LoginUrl string) string  {
	expiresstart := strings.Index(LoginUrl,"Expires=")+8
	expireslast := strings.Index(LoginUrl,"&SESSDATA=")
	a.Expires = LoginUrl[expiresstart:expireslast]
	return LoginUrl[expiresstart:expireslast]
}
