package yugigo

import (
	"encoding/json"
	"net/http"
)



//GetAllArchetypes fetches all the current archetype names in an array
func GetAllArchetypes() ([]ArcheType, error) {
	//Perform a get operation
	resp, err := http.Get(archetypesUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Decode the JSON response into the []ArcheType Struct
	var archetypes []ArcheType
	err = json.NewDecoder(resp.Body).Decode(&archetypes)
	if err != nil {
		return nil, err
	}
	return archetypes, nil
}
