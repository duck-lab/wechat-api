package menu

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/duck-lab/wechat-api/httpHelper"
)

//APINameDeleteAll ...
var APINameDeleteAll = "DELETE_ALL_MENU"

//DeleteAll ...
func DeleteAll(baseURL string, accessToken string) error {
	resp, err := http.Post(baseURL+"/menu/delete?access_token="+accessToken, "", nil)
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
