package wechatApi

import "testing"

func TestBasic(t *testing.T) {
	api := New("aa", "bb")
	token, err := api.getCurrentToken()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
}
