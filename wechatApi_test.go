package wechatApi

import "testing"
import "os"
import "fmt"
import "github.com/duck-lab/wechat-api/accessToken"
import "github.com/stretchr/testify/assert"

func TestInvalid(t *testing.T) {
	api := New("aa", "bb")
	_, err := api.GetAccessToken()
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
	token, err := api.GetAccessToken()
	fmt.Println(token.ExpiresTime)
	if err != nil {
		t.Error("Fetch token failed with reason: ", err.Error())
	} else {
		t.Log("Fetch token success.")
	}
}

func TestCallCount(t *testing.T) {
	if os.Getenv("WECHAT_APPID") == "" {
		t.Error("You should set the WECHAT_APPID & WECHAT_APPSECRET in your env")
	}
	api := New(os.Getenv("WECHAT_APPID"), os.Getenv("WECHAT_APPSECRET"))
	assert.Equal(t, 0, api.getCall(accessToken.APIName))
	token, err := api.GetAccessToken()
	fmt.Println(token, err)
	assert.Equal(t, 1, api.getCall(accessToken.APIName))
	token, err = api.GetAccessToken()
	fmt.Println(token, err)
	assert.Equal(t, 1, api.getCall(accessToken.APIName))
}
