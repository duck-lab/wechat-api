package templateMessage

// Model of template message
type Model struct {
	FollowerOpenID string `json:"touser"`
	TemplateID     string `json:"template_id"`
	URL            string `json:"url"`
	MiniProgram    struct {
		AppID    string `json:"appid"`
		PagePath string `json:"pagepath"`
	}
	Data interface{} `json:"data"` //provided by outside
}
