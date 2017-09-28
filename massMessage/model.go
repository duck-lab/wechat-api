package massMessage

//News 图文消息
type News struct {
	Articles []struct {
		ThumbMediaID     string `json:"thumb_media_id"`
		Author           string `json:"author"`
		Title            string `json:"title"`
		ContentSourceURL string `json:"content_source_url"`
		Content          string `json:"content"`
		Digest           string `json:"digest"`
		ShowCoverPic     string `json:"show_cover_pic"`
	} `json:"articles"`
}
