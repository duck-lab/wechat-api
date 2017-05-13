package notify

import (
	"encoding/xml"
)

type NotifyHead struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          CString  `xml:"ToUserName"`
	UserOpenID     CString  `xml:"FromUserName"`
	CreateTimeUnix int64    `xml:"CreateTime"`
	MsgType        CString  `xml:"MsgType"`
}

type CString string

func (c CString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
