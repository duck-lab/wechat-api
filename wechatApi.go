package wechatApi

import (
	"net/http"

	"errors"

	"io/ioutil"

	"fmt"

	"encoding/json"

	"strconv"

	"github.com/duckLab/wechatApi/accessToken"
	"github.com/duckLab/wechatApi/accessToken/storage"
)

//API is the outlet of all APIs
type API struct {
	endpoint         string
	baseURL          string
	appID            string
	appSecret        string
	tokenStorageMode int
	currentToken     accessToken.Data
}

//New is the start point
func New(appID string, appSecret string) API {
	api := API{
		appID:     appID,
		appSecret: appSecret,
		endpoint:  "https://api.weixin.qq.com",
	}
	api.baseURL = api.endpoint + "/cgi-bin"
	api.tokenStorageMode = storage.InMemory
	return api
}

//SetEndpoint let the user can chose which api server to use
func (api *API) SetEndpoint(endpoint string) {
	api.endpoint = endpoint
	api.baseURL = api.endpoint + "/cgi-bin"
}

//SetTokenStorageMode ...
func (api *API) SetTokenStorageMode(mode int, params interface{}) bool {
	if mode == storage.InMemory {

	}
	if mode == storage.InFile {
		file, ok := params.(storage.File)
		fmt.Println(file, ok)
	}
	if mode == storage.InRedis {
		redis, ok := params.(storage.Redis)
		fmt.Println(redis, ok)
	}
	return false
}

func (api *API) getCurrentToken() (accessToken.Data, error) {
	if api.tokenStorageMode == storage.InMemory {
		if api.currentToken.AccessToken == "" {
			resp, err := http.Get(api.baseURL + "/token?grant_type=client_credential&appid=" + api.appID + "&secret=" + api.appSecret)
			if err != nil {
				return accessToken.Data{}, err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return accessToken.Data{}, errors.New("Get response failed")
			}

			// parse to errormsg
			var failed accessToken.Failed
			json.Unmarshal(body, &failed)
			if failed.ErrCode != 0 {
				return accessToken.Data{}, errors.New(strconv.Itoa(failed.ErrCode) + ":" + failed.ErrMsg)
			}
			var data accessToken.Data
			json.Unmarshal(body, &data)
			api.currentToken = data
			return data, nil
		}
		return api.currentToken, errors.New("")
	}
	if api.tokenStorageMode == storage.InFile {
		//TODO: read file
	}

	if api.tokenStorageMode == storage.InRedis {
		//TODO: read redis
	}
	return accessToken.Data{}, errors.New("")
}

//GetAccessToken can let user fetch the access token
func (api *API) GetAccessToken() (accessToken.Data, error) {
	return api.getCurrentToken()
}
