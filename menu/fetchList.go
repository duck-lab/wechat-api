package menu

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//FetchedList ...
type FetchedList struct {
	Menu            Model             `json:"menu"`
	ConditionalMenu []ConditionalMenu `json:"conditionalmenu"`
}

//FetchedListAPIName ..
var FetchedListAPIName = "FETCH_MENU_LIST"

//FetchList ...
func FetchList(baseURL string, accessToken string) (FetchedList, error) {
	url := baseURL + "/menu/get?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		return FetchedList{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		var fetchedList FetchedList
		err = json.Unmarshal(respBody, &fetchedList)
		if err != nil {
			return FetchedList{}, err
		}
		return fetchedList, nil
	}
	return FetchedList{}, errors.New("Network error")
}
