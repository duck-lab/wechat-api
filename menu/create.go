package menu

import (
	"bytes"
	"encoding/json"
	"errors"

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
	path := "/menu/addconditional"
	var menuIDContainer IDContainer
	err = httpHelper.PostAndParse(baseURL, path, accessToken, bytes.NewReader(bodyBytes), &menuIDContainer)
	return menuIDContainer.MenuID, err
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
	var codeAndMessage httpHelper.CodeAndMessage
	return httpHelper.PostAndParse(baseURL, "/menu/create", accessToken, body, &codeAndMessage)
}
