package customerServiceMessage

import "github.com/duck-lab/wechat-api/common"

type messageHead struct {
	FollowerOpenID string `json:"touser"`
	MsgType        string `json:"msgtype"`
}

//TextMessage ...
type TextMessage struct {
	messageHead
	Text common.MediaIDContainer `json:"text"`
}

//ImageMessage ...
type ImageMessage struct {
	messageHead
	Image common.MediaIDContainer `json:"image"`
}

//VoiceMessage ...
type VoiceMessage struct {
	messageHead
	Image common.MediaIDContainer `json:"image"`
}

//VideoMessage ...
type VideoMessage struct {
	messageHead
	Video struct {
		common.MediaIDContainer
		ThumbMediaID string `json:"thumb_media_id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
	} `json:"video"`
}

//MusicMessage ...
type MusicMessage struct {
	messageHead
	Music struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		MusicURL     string `json:"musicurl"`
		HQMusicURL   string `json:"hqmusicurl"`
		ThumbMediaID string `json:"thumb_media_id"`
	} `json:"music"`
}

//NewsMessage ...
type NewsMessage struct {
	messageHead
	MPNews common.MediaIDContainer `json:"mpnews"`
}

//ExternalNewsMessage ...
type ExternalNewsMessage struct {
	messageHead
	News common.ExternalNews `json:"news"`
}

//CardMessage ...
type CardMessage struct {
	messageHead
	WXCard common.CardIDContainer `json:"wxcard"`
}
