package yugigo

import (
	"fmt"
)

//GetAllCardSets gets all the current cardsets
func GetAllCardSets() ([]CardSetInfo, error) {
	//Decode the JSON Message into an array of CardSetInfo
	var cardsets []CardSetInfo
	err := getData(allCardSetsUrl, &cardsets)
	if err != nil {
		return nil, err
	}
	return cardsets, nil
}

//GetCardFromCardSet takes in a setcode and returns a match of a card if found
//The format of the setcode follows the following pattern: [SetName]-[CardNumber] (EX: SDY-045)
func GetCardFromCardSet(setcode string) (SpecificCard, error) {
	var specificCard SpecificCard
	err := getData(specificUrl+setcode, &specificCard)
	if err != nil {
		return SpecificCard{}, err
	}
	if specificCard.Error != "" {
		return SpecificCard{}, fmt.Errorf("getting setcode %q returned %q", setcode, specificCard.Error)
	}
	return specificCard, nil
}