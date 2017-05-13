package message

import "encoding/xml"

//Model ...
type Model struct {
	XMLName xml.Name `xml:"xml"`
	head
}

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
}

//Head is the common fields
type head struct {
	AppID          CString `xml:"ToUserName"`
	UserOpenID     CString `xml:"FromUserName"`
	CreateTimeUnix int64   `xml:"CreateTime"`
	MsgType        CString `xml:"MsgType"`
}

//CString will be marshalled to CDATA
type CString string

//MarshalXML ...
func (c CString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
