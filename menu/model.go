package menu

import "errors"

//ConditionalMenu is the menu by conditions
type ConditionalMenu struct {
	Model
	MatchRule MatchRule
}

//MatchRule is the conditons for ConditionalMenu
type MatchRule struct {
	TagID              string `json:"tag_id"`
	Gender             string `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType string `json:"client_platform_type"`
	Language           string `json:"language"`
}

//IDContainer used for http request and http response
type IDContainer struct {
	MenuID string `json:"menuid"`
}

func (selfMenu *ConditionalMenu) isValid() error {
	//TODO: validate the parameters
	return nil
}

//Model is the Menu
type Model struct {
	Buttons []Button `json:"button"`
}

func (model *Model) isValid() error {
	if model.Buttons == nil || len(model.Buttons) == 0 {
		return errors.New("Empty menu")
	}
	//TODO: vaildate every button type

	return nil
}

//Button ...
type Button struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Key       string   `json:"key"`
	URL       string   `json:"url"`
	MediaID   string   `json:"media_id"`
	AppID     string   `json:"appid"`
	PagePath  string   `json:"pagepath"`
	SubBttons []Button `json:"sub_button"`
}
