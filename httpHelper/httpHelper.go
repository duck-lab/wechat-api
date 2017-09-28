package httpHelper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func sendPost(url string, body []byte) {

}

//CodeAndMessage is the structure when fetch access token failed
type CodeAndMessage struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//Format the msg
func (msg *CodeAndMessage) Format() string {
	return fmt.Sprintf("Error from wechat server: %d : %s", msg.ErrCode, msg.ErrMsg)
}

//FetchAndParse ...
func FetchAndParse(baseURL string, path string, accessToken string, parsed interface{}) error {
	url := baseURL + path + "?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return Parse(resp, parsed)
	}
	return errors.New("Network error")
}

//PostAndParse ...
func PostAndParse(baseURL string, path string, accessToken string, body io.Reader, parsed interface{}) error {
	url := baseURL + path + "?access_token=" + accessToken
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var respMsg CodeAndMessage
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

// Parse ...
func Parse(resp *http.Response, parsed interface{}) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	var codeAndMsg CodeAndMessage
	err = json.Unmarshal(respBody, codeAndMsg)
	if err != nil {
		return err
	}
	if codeAndMsg.ErrCode != 0 && codeAndMsg.ErrMsg != "" && codeAndMsg.ErrMsg != "ok" {
		return errors.New(codeAndMsg.Format())
	}
	err = json.Unmarshal(respBody, parsed)
	if err != nil {
		return err
	}
	return nil
}
