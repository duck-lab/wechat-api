package notifyFromWechat

// MsgTypeEvent ...
var MsgTypeEvent = "event"

// MsgTypeText ...
var MsgTypeText = "text"

// MsgTypeImage ...
var MsgTypeImage = "image"

// MsgTypeVoice ...
var MsgTypeVoice = "voice"

// MsgTypeVideo ...
var MsgTypeVideo = "video"

// MsgTypeShortVideo ...
var MsgTypeShortVideo = "shortvideo"

// MsgTypeLink ...
var MsgTypeLink = "link"

//EventClick ...
var EventClick = "CLICK"

//EventView ...
var EventView = "VIEW"

//EventScanCodePush ...
var EventScanCodePush = "scancode_push"

//EventScanCodeWaitMsg ...
var EventScanCodeWaitMsg = "scancode_waitmsg"

//EventPicSysPhoto ...
var EventPicSysPhoto = "pic_sysphoto"

//EventPicPhotoOrAlbum ...
var EventPicPhotoOrAlbum = "pic_photo_or_album"

//EventPicWeixin ...
var EventPicWeixin = "pic_weixin"

//EventLocationSelect ...
var EventLocationSelect = "location_select"

//EventSubscribe ...
var EventSubscribe = "subscribe"

//EventUnsubscribe ...
var EventUnsubscribe = "unsubscribe"

//EventScan ...
var EventScan = "SCAN"

//EventLocation ..
var EventLocation = "LOCATION"

//GetEmptyMessage ...
func GetEmptyMessage(msgType string) interface{} {
	switch msgType {
	case MsgTypeEvent:
		return &EventHead{}
	case MsgTypeImage:
		return &ImageMessage{}
	case MsgTypeLink:
		return &LinkMessage{}
	case MsgTypeShortVideo:
		return &ShortVideoMessage{}
	case MsgTypeText:
		return &TextMessage{}
	case MsgTypeVideo:
		return &VideoMessage{}
	case MsgTypeVoice:
		return &VoiceMessage{}
	}
	return nil
}

//GetEmptyEvent ...
func GetEmptyEvent(eventType string) interface{} {
	switch eventType {
	case EventClick:
		return &MenuClickEvent{}
	case EventLocation:
		return &LocationEvent{}
	case EventLocationSelect:
		return &MenuLocationSelectEvent{}
	case EventPicPhotoOrAlbum:
		return &MenuPicPhotoOrAlbumEvent{}
	case EventPicSysPhoto:
		return &MenuPicSysPhotoEvent{}
	case EventPicWeixin:
		return &MenuPicWeixinEvent{}
	case EventScan:
		return &ScanEvent{}
	case EventScanCodePush:
		return &MenuScanCodePushEvent{}
	case EventScanCodeWaitMsg:
		return &MenuScanCodeWaitMsgEvent{}
	case EventSubscribe:
		return &SubscribeEvent{}
	case EventUnsubscribe:
		return &UnsubscribeEvent{}
	}
	return nil
}
