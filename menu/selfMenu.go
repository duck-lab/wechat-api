package menu

//SelfMenu is the menu by conditions
type SelfMenu struct {
	Model
	MatchRule MatchRule
}

//MatchRule is the conditons for SelfMenu
type MatchRule struct {
	TagID              string `json:"tag_id"`
	Gender             string `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType string `json:"client_platform_type"`
	Language           string `json:"language"`
}
