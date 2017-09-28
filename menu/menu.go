package menu

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/duckLab/wechatApi/response"

	"errors"
	"io/ioutil"
)

//Model is the Menu
type Model struct {
	Buttons []Button `json:"button"`
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

//ClickButtonType ..
var ClickButtonType = "click"

//ViewButtonType ..
var ViewButtonType = "view"

//ScanCodePushButtonType ...
var ScanCodePushButtonType = "scancode_push"

//ScanCodeWaitMsgButtonType ...
var ScanCodeWaitMsgButtonType = "scancode_waitmsg"

//PicSysPhotoButtonType ...
var PicSysPhotoButtonType = "pic_sysphoto"

//PicPhotoOrAlbumButtonType ...
var PicPhotoOrAlbumButtonType = "pic_photo_or_album"

//PicWechatButtonType ...
var PicWechatButtonType = "pic_weixin"

//LocationSelectButtonType ...
var LocationSelectButtonType = "location_select"

//MediaIDButtonType ...
var MediaIDButtonType = "media_id"

//MiniProgaramButtonType ...
var MiniProgaramButtonType = "miniprogram"

//CreateAPIName is the unique name of this API
var CreateAPIName = "CREATE_MENU"

func (model *Model) isValid() error {
	if model.Buttons == nil || len(model.Buttons) == 0 {
		return errors.New("Empty menu")
	}
	//TODO: vaildate every button type

	return nil
}

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
		var respMsg response.CodeAndMessage
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

//DeleteAllAPIName ...
var DeleteAllAPIName = "DELETE_ALL_MENU"

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
		var respMsg response.CodeAndMessage
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
