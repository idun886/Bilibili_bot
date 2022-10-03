package Plugin

import (
	"Bili_go/BilibiliApi"
	"Bili_go/ClientMethod"
	"Bili_go/util"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"time"
)

type ReplyRrivateMsg struct {
	TalkerID int64			//消息框的对方的id   如果此id和 senderuid不一样 那么那条消息是自己发送的 不用理会
	SenderUID int64         	//发送消息的id
	ReceiverType int64		//接收消息的类型
	ReceiverID int64			//接收消息的id
	Content string
	Msg string
}

func mapgotojson() map[string]string{

	jsoncontext,err1 := ioutil.ReadFile("Config/ciku.json")
	if err1 != nil {
		fmt.Println("read fail", err1)
	}
	m := make(map[string]string)
	err := json.Unmarshal([]byte(string(jsoncontext)), &m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}


func (r *ReplyRrivateMsg)ReplyPrivate(header string,userid string,csrf string)  {
	bilibiliapirequest := ClientMethod.BilibiliApiRequest{}
	createpostdata := util.CreatePostData{}
	for {
		contextbody := bilibiliapirequest.GetApi(BilibiliApi.PrivateChatMsgApi, header)
		jsondata := gjson.Get(contextbody, "data").String()
		testarray := gjson.Get(jsondata, "session_list").Array()

		//fmt.Println(testarray)

		for msg := range testarray {
			msgjson := testarray[msg].String()
			r.TalkerID = gjson.Get(msgjson, "talker_id").Int()
			last_msg := gjson.Get(msgjson, "last_msg").String()
			r.SenderUID = gjson.Get(last_msg, "sender_uid").Int()
			r.ReceiverID = gjson.Get(last_msg, "receiver_id").Int()
			r.Content = gjson.Get(last_msg, "content").String()
			r.Msg = gjson.Get(r.Content, "content").String()



			//fmt.Printf("当前对话框对方的uid:%d\t,消息发送者的uid:%d\t,消息接收者的uid:%d\t,最后一条消息:%s\t,最后一条消息的文本格式%s\n", r.TalkerID, r.SenderUID, r.ReceiverID, r.Content, r.Msg)

			for mapkey := range mapgotojson() {
				if mapkey == r.Msg && r.SenderUID != r.ReceiverID {
					//fmt.Println(mapgotojson()[mapkey])
					postdata := createpostdata.SendPrivtePostData(userid,r.TalkerID,1,1,1,"send_msg",csrf,mapgotojson()[mapkey])
					fmt.Printf("向uid:%d回复了一条消息:%s\n", r.TalkerID, mapgotojson()[mapkey])

					contextbody := bilibiliapirequest.PostApi(BilibiliApi.SendPrivateMsgApi,postdata,header)
					fmt.Println(contextbody)
				}
			}

		}
		time.Sleep(time.Duration(2)*time.Second)
	}
}