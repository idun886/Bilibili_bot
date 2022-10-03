package util

import "strconv"

type CreatePostData struct {
	Csrf string				//bilibili所有与登录状态相关请求都需要的postdata
	Message string			//发送消息需要的post
	Oid int					//发送的视频oid  oid是由bv号转av号转过来的
	Plat int
	Posttype int
	//BiliCSRF string 		//哔哩哔哩注销登录需要的postdata

}
func (c *CreatePostData)NewCreatePostData(){

}
//获取 发送评论api的post参数函数
func (c *CreatePostData)SendCommentPostData(datacrsf string,datamessage string,dataoid int, dataplat int, datatype int) string {
	return "csrf="+datacrsf+"&message="+datamessage+"&oid="+strconv.Itoa(dataoid)+"&plat="+strconv.Itoa(dataplat)+"&type="+strconv.Itoa(datatype)
}

func (c *CreatePostData)SendPrivtePostData(sender_uid string,receiver_id int64,receiver_type int64,msg_type int64,timestamp int64,dev_id string,csrf string,msg string) string{
	return "csrf="+csrf+
		"&msg[sender_uid]="+sender_uid+
		"&msg[receiver_id]="+strconv.FormatInt(receiver_id,10)+
		"&msg[receiver_type]="+ strconv.FormatInt(receiver_type,10)+
		"&msg[msg_type]="+strconv.FormatInt(msg_type,10)+
		"&msg[dev_id]="+dev_id+
		"&msg[timestamp]="+strconv.FormatInt(timestamp,10)+
		"&msg[content]="+`{"content":"`+msg+`"}`
}

