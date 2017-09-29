package notifyFromWechat

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageMessage(t *testing.T) {
	var head CommonHead
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
	if err != nil {
		t.Error(err)
	}
	pt := GetEmptyMessage(head.MsgType)
	err = xml.Unmarshal(xmlBytes, pt)
	if err != nil {
		t.Error(err)
	}
	imageMessage := pt.(*ImageMessage)
	assert.Equal(t, int64(1234567890123456), imageMessage.MsgID)
	assert.Equal(t, "this is a url", imageMessage.PicURL)
}

func TestMenuScanCodePushEvent(t *testing.T) {
	var head CommonHead
	xmlBytes := []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090502</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[scancode_push]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<ScanCodeInfo>
		<ScanType><![CDATA[qrcode]]></ScanType>
		<ScanResult><![CDATA[1]]></ScanResult>
		</ScanCodeInfo>
		</xml>`)
	err := xml.Unmarshal(xmlBytes, &head)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, MsgTypeEvent, head.MsgType)
	var eventHead EventHead
	err = xml.Unmarshal(xmlBytes, &eventHead)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, EventScanCodePush, eventHead.Event)

	pt := GetEmptyEvent(eventHead.Event)
	err = xml.Unmarshal(xmlBytes, pt)
	if err != nil {
		t.Error(err)
	}

	menuScanCodePushEvent := pt.(*MenuScanCodePushEvent)
	assert.Equal(t, "qrcode", menuScanCodePushEvent.ScanCodeInfo.ScanType)
}

func TestMenuPicSysPhotoEvent(t *testing.T) {
	var head CommonHead
	xmlBytes := []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
		<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
		<CreateTime>1408090651</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[pic_sysphoto]]></Event>
		<EventKey><![CDATA[6]]></EventKey>
		<SendPicsInfo>
			<Count>2</Count>
			<PicList>
				<item>
					<PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum>
				</item>
				<item>
					<PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e186]]></PicMd5Sum>
				</item>
			</PicList>
		</SendPicsInfo>
		</xml>`)
	err := xml.Unmarshal(xmlBytes, &head)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, MsgTypeEvent, head.MsgType)
	var eventHead EventHead
	err = xml.Unmarshal(xmlBytes, &eventHead)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, EventPicSysPhoto, eventHead.Event)

	pt := GetEmptyEvent(eventHead.Event)
	err = xml.Unmarshal(xmlBytes, pt)
	if err != nil {
		t.Error(err)
		return
	}

	menuPicSysPhotoEvent := pt.(*MenuPicSysPhotoEvent)
	assert.Equal(
		t,
		menuPicSysPhotoEvent.SendPicsInfo.Count,
		len(menuPicSysPhotoEvent.SendPicsInfo.PicList.Items),
	)

	assert.Equal(
		t,
		"1b5f7c23b5bf75682a53e7b6d163e186",
		menuPicSysPhotoEvent.SendPicsInfo.PicList.Items[1].PicMd5Sum,
	)
	j, _ := json.Marshal(menuPicSysPhotoEvent)
	log.Println(string(j))
}
