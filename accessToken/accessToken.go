package accessToken

import (
	"encoding/json"
)

//Data is the token used everywhere
type Data struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//Failed is the structure when fetch access token failed
type Failed struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//ToString is used to stringfify this token
func (token *Data) ToString() string {
	data, err := json.Marshal(token)
	if err != nil {
		return ""
	}
	return string(data)
}
