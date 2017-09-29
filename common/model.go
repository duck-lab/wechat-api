package common

//MediaIDContainer ..
type MediaIDContainer struct {
	MediaID string `json:"media_id"`
}

//CardIDContainer ..
type CardIDContainer struct {
	CardID string `json:"card_id"`
}

//ExternalNewsArticle ...
type ExternalNewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

//ExternalNews ...
type ExternalNews struct {
	Articles []ExternalNewsArticle `xml:"item" json:"articles"`
}

//News 图文消息
type News struct {
	Articles []NewsArticle `json:"articles"`
}

//NewsArticle ...
type NewsArticle struct {
	ThumbMediaID       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Title              string `json:"title"`
	ContentSourceURL   string `json:"content_source_url"`
	Content            string `json:"content"`
	Digest             string `json:"digest"`
	ShowCoverPic       string `json:"show_cover_pic"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
}
