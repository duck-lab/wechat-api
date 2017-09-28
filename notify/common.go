package notify

// Head the common fields of wechat notify
type Head struct {
	AppID          string `xml:"ToUserName"`
	UserOpenID     string `xml:"FromUserName"`
	CreateTimeUnix int64  `xml:"CreateTime"`
	MsgType        string `xml:"MsgType"`
}
