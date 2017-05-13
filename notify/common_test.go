package notify

import "testing"
import "encoding/xml"

import "log"

func TestNotifyHead(t *testing.T) {

	var head Head
	xmlBytes := []byte(`<xml>
 <ToUserName><![CDATA[123]]></ToUserName>
 <FromUserName><![CDATA[321]]></FromUserName>
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[image]]></MsgType>
 <PicUrl><![CDATA[this is a url]]></PicUrl>
 <MediaId><![CDATA[media_id]]></MediaId>
 <MsgId>1234567890123456</MsgId>
 </xml>`)
	err := xml.Unmarshal(xmlBytes, &head)
	log.Println(head, err)
}
