package yugigo

import "fmt"

//GetRandomCard returns information about a random card
func GetRandomCard() (CardInfo, error) {
	var cardInfo CardInfo
	err := getData(randomUrl, &cardInfo)
	if err != nil {
		return CardInfo{}, err
	}
	setImageID(cardInfo)
	return cardInfo, nil
}

//GetCards Method parses the search query provided by the user by turning the struct into a map[string]interface.
//It then makes some sanity checks, and then calls a helper method to build the full query url.
//It then uses the query url to fetch the data which is then returns if there were no errors returned.
func GetCards(search CardSearch) ([]CardInfo, MetaInfo, error) {
	var data Data
	//use custom decoder to get a map[string]interface{}
	params, err := marshalMap(&search, "search")
	if err != nil {
		return nil, MetaInfo{}, err
	}

	//check for offsets and num
	checkOffsetNum(&search, params)
	//call urlHelper to make a url with given queries
	url := makeUrl(cardInfoUrl, params)

	//call the getData function
	err = getData(url, &data)
	if err != nil {
		return nil, MetaInfo{}, err
	}
	//check for api specific error
	if data.ApiError != "" {
		return nil, MetaInfo{}, fmt.Errorf("getting %q returned %q", url, data.ApiError)
	}
	return data.Data, data.Meta, nil
}

/*
GetCardWithFullUrl assumes that the user has the complete url (query and such included).
This can be used with the Num and Offset to separate the cards
*/
func GetCardWithFullUrl(url string) ([]CardInfo, MetaInfo, error) {
	var data Data
	err := getData(url, &data)
	if err != nil {
		return nil, MetaInfo{}, err
	}
	return data.Data, data.Meta, nil
}
