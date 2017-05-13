package notify

type Head struct {
	AppID          string `xml:"ToUserName"`
	UserOpenID     string `xml:"FromUserName"`
	CreateTimeUnix int64  `xml:"CreateTime"`
	MsgType        string `xml:"MsgType"`
}
