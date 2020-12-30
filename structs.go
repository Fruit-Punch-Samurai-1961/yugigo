package yugigo

//ArcheType struct contains the name of the archetype
type ArcheType struct {
	ArchetypeName string `json:"archetype_name"`
}

//The api returns an array of cards which are under the data field
//Looks like this {"data" : [CardInfo,CardInfo...]}
type Data struct {
	Data     []CardInfo `json:"data"`
	Meta     MetaInfo   `json:"meta,omitempty"`
	ApiError string     `json:"error"`
}

//Struct used to parse card(s) information
type CardInfo struct {
	Id          int             `card:"id"`
	Name        string          `json:"name"`
	CardType    string          `json:"type"`
	Desc        string          `json:"desc"`
	Atk         int             `json:"atk,omitempty"`
	Def         int             `json:"def,omitempty"`
	Level       int             `json:"level,omitempty"`
	Race        string          `json:"race"`
	Attribute   string          `json:"attribute,omitempty"`
	Archetype   string          `json:"archetype,omitempty"`
	Scale       int             `json:"scale,omitempty"`
	Linkval     int             `json:"linkval,omitempty"`
	Linkmarkers []string        `json:"linkmarkers,omitempty"`
	CardSets    []CardSet       `json:"card_sets"`
	CardImages  []CardImageInfo `json:"card_images"`
	CardPrices  []CardPriceInfo `json:"card_prices"`
	BanlistInfo Banlist         `json:"banlist_info,omitempty"`
	BetaName    string          `json:"beta_name,omitempty"`
	TreatedAs   string          `json:"treated_as,omitempty"`
	TCGDate     string          `json:"tcg_date,omitempty"`
	OCGDate     string          `json:"ocg_date,omitempty"`
	HasEffect   string          `json:"has_effect,omitempty"`
}

//If a user wants an offset, then the API will return meta information which includes the next-page's link
//This can be used for the GetCardWithFullUrl function
type MetaInfo struct {
	CurrentRows    int    `json:"current_rows"`
	TotalRows      int    `json:"total_rows"`
	RowsRemaining  int    `json:"rows_remaining"`
	TotalPages     int    `json:"total_pages"`
	PagesRemaining int    `json:"pages_remaining"`
	NextPage       string `json:"next_page"`
	NextPageOffset int    `json:"next_page_offset"`
}

//CardSet struct contains all information regarding the card set where the card can be found
type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

//Struct that contains the ID, imageUrl, and the small Image url of requested card
type CardImageInfo struct {
	Id            int    `json:"id"`
	ImageUrl      string `json:"image_url"`
	ImageUrlSmall string `json:"image_url_small"`
}

//Struct info which contains all the prices for the requested card
type CardPriceInfo struct {
	CardmarketPrice   string `json:"cardmarket_price"`
	CoolstuffincPrice string `json:"coolstuffinc_price"`
	TCGPlayerPrice    string `json:"tcgplayer_price"`
	EbayPrice         string `json:"ebay_price"`
	AmazonPrice       string `json:"amazon_price"`
}

//Certain cards have some restrictions which this struct helps display to the user
type Banlist struct {
	BanTCG  string `json:"ban_tcg"`
	BanOCG  string `json:"ban_ocg"`
	BanGoat string `json:"ban_goat,omitempty"`
}

//Card Search Struct which is accessible to the user to specify their query.
//For ATK and DEF, once can add "lt" (less than), "lte" (less than equals to), "gt" (greater than), or "gte" (greater than equals to) in front of the value.
//If both Name and FName are specified, Name will take priority.
//If EndDate is specified, then you need to specify StartDate and vice versa.
//Num is the total number of cards to return per request.
//Offset is the number of cards to skip.
//If Num is specified, then you need to specify Offset and vice versa. If only Num is specified, the default Offset is 0. If only Offset is specified, the default Num becomes 50.
type CardSearch struct {
	Name           string   `search:"name"`
	FName          string   `search:"fname"`
	ID             string   `search:"id"`
	Type           []string `search:"type"`
	ATK            string   `search:"atk"`
	DEF            string   `search:"def"`
	Level          string   `search:"level"`
	Race           []string `search:"race"`
	Attribute      string   `search:"attribute"`
	Link           string   `search:"link"`
	Linkmarker     string   `search:"linkmarker"`
	Scale          string   `search:"scale"`
	CardSet        string   `search:"cardset"`
	Archetype      string   `search:"archetype"`
	Banlist        string   `search:"banlist"`
	Sort           string   `search:"sort"`
	Format         string   `search:"format"`
	Staple         string   `search:"staple"`
	HasEffect      string   `search:"has_effect"`
	Misc           string   `search:"misc"`
	IncludeAliased string   `search:"include_aliased"`
	StartDate      string   `search:"startdate"`
	EndDate        string   `search:"enddate"`
	Region         string   `search:"dateregion"`
	Num            string   `search:"num"`
	Offset         string   `search:"offset"`
}

//CardSetInfo struct contains the name of the set, the code, the number of cards, and when it was released
type CardSetInfo struct {
	SetName    string `json:"set_name"`
	SetCode    string `json:"set_code"`
	NumOfCards int    `json:"num_of_cards"`
	TCGDate    string `json:"tcg_date"`
}

//SpecificCard contains information of a specific card from a specific card set
type SpecificCard struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SetName   string `json:"set_name"`
	SetCode   string `json:"set_code"`
	SetRarity string `json:"set_rarity"`
	SetPrice  string `json:"set_price"`
	Error     string `json:"error,omitempty"`
}
