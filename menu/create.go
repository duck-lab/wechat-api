package menu

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/duck-lab/wechat-api/httpHelper"
)

//APINameCreateConditional is the unique name of this API
var APINameCreateConditional = "CREATE_CONDITIONAL_MENU"

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
		var respMsg httpHelper.CodeAndMessage
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

//APINameCreate is the unique name of this API
var APINameCreate = "CREATE_MENU"

//Create ...
func Create(model Model, baseURL string, accessToken string) error {
	err := model.isValid()
	if err != nil {
		return err
	}
	modelJSONBytes, err := json.Marshal(model)
	if err != nil {
		return errors.New("Marshal Menu failed")
	}
	body := bytes.NewReader(modelJSONBytes)
	resp, err := http.Post(baseURL+"/menu/create", "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var respMsg httpHelper.CodeAndMessage
		err = json.Unmarshal(respBody, &respMsg)
		if err != nil {
			return err
		}
		if respMsg.ErrCode != 0 {
			return errors.New(respMsg.Format())
		}
		return nil
	}
	return errors.New("Net error")
}
