package BilibiliApi

//登录接口 返回qrcode 和登录的key
const LoginBilibiliApi = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate?source=main-mini"
const SendCommentMessageApi = "https://api.bilibili.com/x/v2/reply/add"
const PrivateChatMsgApi = "https://api.vc.bilibili.com/session_svr/v1/session_svr/new_sessions"
const SendPrivateMsgApi = "https://api.vc.bilibili.com/web_im/v1/web_im/send_msg"