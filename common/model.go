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
	Articles []ExternalNewsArticle `json:"articles"`
}
