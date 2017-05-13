package notify

import (
	"encoding/xml"
)

type NotifyHead struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          cstring  `xml:"ToUserName"`
	UserOpenID     cstring  `xml:"FromUserName"`
	CreateTimeUnix int64    `xml:"CreateTime"`
	MsgType        cstring  `xml:"MsgType"`
}

type cstring string

func (c cstring) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
