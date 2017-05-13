package accessToken

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
