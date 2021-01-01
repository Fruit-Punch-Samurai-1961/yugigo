package yugigo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

/*
All constants of this package
*/

const (
	archetypesUrl  = "https://db.ygoprodeck.com/api/v7/archetypes.php"
	randomUrl      = "https://db.ygoprodeck.com/api/v7/randomcard.php"
	databaseurl    = "https://db.ygoprodeck.com/api/v7/checkDBVer.php"
	allCardSetsUrl = "https://db.ygoprodeck.com/api/v7/cardsets.php"
	specificUrl    = "https://db.ygoprodeck.com/api/v7/cardsetsinfo.php?setcode="
	cardInfoUrl    = "https://db.ygoprodeck.com/api/v7/cardinfo.php?"
)

/*
	getData is a helper function which helps reduce repeated code
*/
func getData(url string, target interface{}) error {
	//Perform a Get operation
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//decode the JSON message into the given interface
	return json.NewDecoder(resp.Body).Decode(target)
}

/*
The decoder ignores fields that have the same name, so we traverse trough the image json and grab the ID from the ImageUrl
*/
func setImageID(info CardInfo) {
	for i := range info.CardImages {
		str := info.CardImages[i].ImageUrl
		info.CardImages[i].Id, _ = strconv.Atoi(str[strings.LastIndex(str, "/")+1 : strings.LastIndex(str, ".")])
	}
}

func marshalMap(input interface{}, tag string) (map[string]interface{}, error) {
	//make the map[string]interface we're returning
	result := make(map[string]interface{})
	//Get the value of the input
	inputVal := reflect.ValueOf(input)

	//If the input is a pointer, set inputVal to the ptr's elements
	if inputVal.Kind() == reflect.Ptr {
		inputVal = inputVal.Elem()
	}

	//Check if the input is a struct
	if inputVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("marshalMap only accepts structs; got %T which is type %T", inputVal, inputVal.Type())
	}

	//Get the Type of input which will be used to get the struct field
	inputType := inputVal.Type()
	//loops over each of the fields of the structs
	//It checks if the field's value is zero for the tag provided
	//If it's not, then it adds it to the result
	for i := 0; i < inputVal.NumField(); i++ {
		field := inputType.Field(i)
		if tagValue := field.Tag.Get(tag); !inputVal.Field(i).IsZero() && tagValue != "" {
			result[tagValue] = inputVal.Field(i).Interface()
		}
	}
	return result, nil
}

/*
checkOffsetNum checks if the user provided both offset and num or none at all
If they provided one or the other, the default value for the other is set
*/
func checkOffsetNum(search *CardSearch, params map[string]interface{}) {
	search.Num = strings.TrimSpace(search.Num)
	search.Offset = strings.TrimSpace(search.Offset)
	if search.Num != "" && search.Offset == "" {
		params["offset"] = "0"
	} else if search.Num == "" && search.Offset != "" {
		params["num"] = "50"
	}
}

/*
makeUrl adds the given queries into the url string
*/
func makeUrl(defaulturl string, params map[string]interface{}) string {
	//loop through the map
	//the only thing we have to check for is if it's a string or a []string
	for key, value := range params {
		if _, ok := value.([]string); !ok {
			//encode the query
			value = url.QueryEscape(value.(string))
			defaulturl += key + "=" + value.(string) + "&"
		} else {
			for _, str := range value.([]string) {
				str = strings.ReplaceAll(strings.TrimSpace(str), " ", "%20")
				defaulturl += key + "=" + str + "&"
			}
		}
	}
	return defaulturl
}
