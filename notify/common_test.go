package notify

import "testing"
import "encoding/xml"
import "time"
import "log"

func TestNotifyHead(t *testing.T) {
	xmlBytes, err := xml.Marshal(NotifyHead{
		AppID:          "123",
		UserOpenID:     "aaa",
		CreateTimeUnix: time.Now().Unix(),
		MsgType:        "text",
	})
	log.Println(string(xmlBytes), err)
	var head NotifyHead
	xmlBytes = []byte(`<xml>
 <ToUserName><![CDATA[123]]></ToUserName>
 <FromUserName><![CDATA[321]]></FromUserName>
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[image]]></MsgType>
 <PicUrl><![CDATA[this is a url]]></PicUrl>
 <MediaId><![CDATA[media_id]]></MediaId>
 <MsgId>1234567890123456</MsgId>
 </xml>`)
	err = xml.Unmarshal(xmlBytes, &head)
	log.Println(head, err)
}
