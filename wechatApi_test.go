package wechatApi

import "testing"
import "os"

func TestInvalid(t *testing.T) {
	api := New("aa", "bb")
	_, err := api.getCurrentToken()
	if err != nil {
		t.Log("Failed as expected.")
	} else {
		t.Error("Success unexpectlly.")
	}
}

func TestValid(t *testing.T) {
	if os.Getenv("WECHAT_APPID") == "" {
		t.Error("You should set the WECHAT_APPID & WECHAT_APPSECRET in your env")
	}
	api := New(os.Getenv("WECHAT_APPID"), os.Getenv("WECHAT_APPSECRET"))
	_, err := api.getCurrentToken()
	if err != nil {
		t.Error("Fetch token failed with reason: ", err.Error())
	} else {
		t.Log("Fetch token success.")
	}
}
