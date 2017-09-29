package menu

import "github.com/duck-lab/wechat-api/httpHelper"

//FetchedList ...
type FetchedList struct {
	Menu            Model             `json:"menu"`
	ConditionalMenu []ConditionalMenu `json:"conditionalmenu"`
}

//APINameFetchedList ..
var APINameFetchedList = "FETCH_MENU_LIST"

//FetchList ...
func FetchList(baseURL string, accessToken string) (FetchedList, error) {
	path := "/menu/get"
	var fetchedList FetchedList
	err := httpHelper.FetchAndParse(baseURL, path, accessToken, &fetchedList)
	return fetchedList, err
}
