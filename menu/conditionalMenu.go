package menu

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/duckLab/wechatApi/response"
)

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

//CreateConditionalAPIName is the unique name of this API
var CreateConditionalAPIName = "CREATE_CONDITIONAL_MENU"

//CreateConditional returns menuId and error
func CreateConditional(selfMenu ConditionalMenu, baseURL string, accessToken string) (string, error) {
	err := selfMenu.isValid()
	if err != nil {
		return "", err
	}
	bodyBytes, err := json.Marshal(selfMenu)
	if err != nil {
		return "", err
	}
	url := baseURL + "/menu/addconditional?access_token=" + accessToken
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var menuIDContainer IDContainer
		err = json.Unmarshal(respBody, &menuIDContainer)
		if err != nil {
			return "", err
		}
		if menuIDContainer.MenuID != "" {
			return menuIDContainer.MenuID, nil
		}
		var respMsg response.CodeAndMessage
		err = json.Unmarshal(respBody, &respMsg)
		if err != nil {
			return "", err
		}
		if respMsg.ErrCode != 0 {
			return "", errors.New(respMsg.Format())
		}
		return "", errors.New("Unknow error")
	}
	return "", errors.New("Network error")
}
