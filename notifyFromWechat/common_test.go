package notifyFromWechat

import "testing"

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

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
	jstr, _ := json.Marshal(head)
	log.Println(string(jstr), err)
}
