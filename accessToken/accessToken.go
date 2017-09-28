package accessToken

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/duck-lab/wechat-api/httpHelper"
)

//APINameFetch is the name of API
var APINameFetch = "FETCH_ACCESS_TOKEN"

//StoreInMemory means store in a variable
const StoreInMemory = 1

//StoreInFile means store in a specified file
const StoreInFile = 2

//StoreInRedis means store in the specified redis with given key
const StoreInRedis = 3

//Data is the token used everywhere
type Data struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresTime time.Time
}

//StoreInFileParams ...
type StoreInFileParams struct {
	name string
}

//StoreInRedisParams ...
type StoreInRedisParams struct {
	host     string
	port     int
	password string
}

//ToString is used to stringfify this token
func (token *Data) ToString() string {
	data, err := json.Marshal(token)
	if err != nil {
		return ""
	}
	return string(data)
}

//Fetch from wechat server
func Fetch(appID string, appSecret string, baseURL string) (Data, error) {
	resp, err := http.Get(baseURL + "/token?grant_type=client_credential&appid=" + appID + "&secret=" + appSecret)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Data{}, errors.New("Get response failed")
	}

	// parse to errormsg
	var failed httpHelper.CodeAndMessage
	json.Unmarshal(body, &failed)
	if failed.ErrCode != 0 {
		return Data{}, errors.New(failed.Format())
	}
	var data Data
	json.Unmarshal(body, &data)
	data.ExpiresTime = time.Now().Add(time.Duration(data.ExpiresIn) * time.Second)
	return data, nil
}
