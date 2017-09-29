package notifyFromWechat

//CommonHead ...
type CommonHead struct {
	AppID          string `xml:"ToUserName"`
	FollowerOpenID string `xml:"FromUserName"`
	CreateTimeUnix int64  `xml:"CreateTime"`
	MsgType        string `xml:"MsgType"`
}

// MessageHead the common fields of wechat notify
type MessageHead struct {
	CommonHead
	MsgID int64 `xml:"MsgId"`
}

//TextMessage ...
type TextMessage struct {
	MessageHead
	Content string `xml:"Content"`
}

//ImageMessage ...
type ImageMessage struct {
	MessageHead
	PicURL  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`
}

//VoiceMessage ...
type VoiceMessage struct {
	MessageHead
	Format      string `xml:"Format"`
	MediaID     string `xml:"MediaId"`
	Recognition string `xml:"Recognition"`
}

//VideoMessage ...
type VideoMessage struct {
	MessageHead
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

//ShortVideoMessage ...
type ShortVideoMessage struct {
	MessageHead
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

//LocationMessage ...
type LocationMessage struct {
	MessageHead
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
}

//LinkMessage ...
type LinkMessage struct {
	MessageHead
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
}

//EventHead ...
type EventHead struct {
	CommonHead
	Event string `xml:"Event"`
}

//SubscribeEvent ...
type SubscribeEvent struct {
	EventHead
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

//UnsubscribeEvent ...
type UnsubscribeEvent struct {
	EventHead
}

//ScanEvent ...
type ScanEvent struct {
	EventHead
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

//LocationEvent ...
type LocationEvent struct {
	EventHead
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
}

//MenuClickEvent ...
type MenuClickEvent struct {
	EventHead
	EventKey string `xml:"EventKey"`
}

//MenuViewEvent ...
type MenuViewEvent struct {
	EventHead
	EventKey string `xml:"EventKey"`
	MenuID   string `xml:"MenuID"`
}

//ScanCodeInfo ...
type ScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

//MenuScanCodePushEvent ...
type MenuScanCodePushEvent struct {
	EventHead
	EventKey     string       `xml:"EventKey"`
	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
}

//MenuScanCodeWaitMsgEvent ...
type MenuScanCodeWaitMsgEvent struct {
	MenuScanCodePushEvent
}

//SendPicsInfo ...
type SendPicsInfo struct {
	Count   int `xml:"Count"`
	PicList struct {
		Items []struct {
			PicMd5Sum string `xml:"PicMd5Sum"`
		} `xml:"item"`
	} `xml:"PicList"`
}

//MenuPicSysPhotoEvent ...
type MenuPicSysPhotoEvent struct {
	EventHead
	EventKey     string       `xml:"EventKey"`
	SendPicsInfo SendPicsInfo `xml:"SendPicsInfo"`
}

//MenuPicPhotoOrAlbumEvent ...
type MenuPicPhotoOrAlbumEvent struct {
	MenuPicSysPhotoEvent
}

//MenuPicWeixinEvent ...
type MenuPicWeixinEvent struct {
	MenuPicSysPhotoEvent
}

//SendLocationInfo ..
type SendLocationInfo struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	POIname   string  `xml:"Poiname"`
}

//MenuLocationSelectEvent ...
type MenuLocationSelectEvent struct {
	EventHead
	EventKey         string           `xml:"EventKey"`
	SendLocationInfo SendLocationInfo `xml:"SendLocationInfo"`
}
